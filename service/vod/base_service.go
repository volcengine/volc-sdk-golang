package vod

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/models/vod/response"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/consts"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/model"
)

func (p *Vod) CreateSha1HlsDrmAuthToken(deadline time.Time) (auth string, err error) {
	return createAuth(DSAHmacSha1, Version1, p.ServiceInfo.Credentials.AccessKeyID,
		p.ServiceInfo.Credentials.SecretAccessKey, deadline)
}

func (p *Vod) CreateSha256HlsDrmAuthToken(deadline time.Time) (auth string, err error) {
	return createAuth(DSAHmacSha256, Version1, p.ServiceInfo.Credentials.AccessKeyID,
		p.ServiceInfo.Credentials.SecretAccessKey, deadline)
}

func (p *Vod) GetPlayAuthToken(req *request.VodGetPlayInfoRequest, tokenExpireTime int) (string, error) {
	if len(req.GetVid()) == 0 {
		return "", errors.New("传入的Vid为空")
	}
	query := url.Values{
		"Vid": []string{req.GetVid()},
	}
	if len(req.GetDefinition()) > 0 {
		query.Add("Definition", req.GetDefinition())
	}
	if len(req.GetFileType()) > 0 {
		query.Add("FileType", req.GetFileType())
	}
	if len(req.GetCodec()) > 0 {
		query.Add("Codec", req.GetCodec())
	}
	if len(req.GetFormat()) > 0 {
		query.Add("Format", req.GetFormat())
	}
	if len(req.GetBase64()) > 0 {
		query.Add("Base64", req.GetBase64())
	}
	if len(req.GetLogoType()) > 0 {
		query.Add("LogoType", req.GetLogoType())
	}
	if len(req.GetSsl()) > 0 {
		query.Add("Ssl", req.GetSsl())
	}
	if tokenExpireTime > 0 {
		query.Add("X-Expires", strconv.Itoa(tokenExpireTime))
	}
	if len(req.GetNeedThumbs()) > 0 {
		query.Add("NeedThumbs", req.GetNeedThumbs())
	}
	if getPlayInfoToken, err := p.GetSignUrl("GetPlayInfo", query); err == nil {
		ret := map[string]string{}
		ret["GetPlayInfoToken"] = getPlayInfoToken
		ret["TokenVersion"] = "V2"
		b, err := json.Marshal(ret)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(b), nil
	} else {
		return "", err
	}
}

func (p *Vod) UploadMediaWithCallback(filePath string, spaceName string, callbackArgs string, funcs ...Function) (*response.VodCommitUploadInfoResponse, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, -1, err
	}
	return p.UploadMediaInner(file, stat.Size(), spaceName, callbackArgs, funcs...)

}

func (p *Vod) UploadMediaInner(rd io.Reader, size int64, spaceName string, callbackArgs string, funcs ...Function) (*response.VodCommitUploadInfoResponse, int, error) {
	_, sessionKey, err, code := p.Upload(rd, size, spaceName)
	if err != nil {
		return nil, code, err
	}

	fbts, err := json.Marshal(funcs)
	if err != nil {
		panic(err)
	}

	commitRequest := &request.VodCommitUploadInfoRequest{
		SpaceName:    spaceName,
		SessionKey:   sessionKey,
		CallbackArgs: callbackArgs,
		Functions:    string(fbts),
	}

	commitResp, code, err := p.CommitUploadInfo(commitRequest)
	if err != nil {
		return nil, code, err
	}
	return commitResp, code, nil
}

func (p *Vod) GetUploadAuthWithExpiredTime(expiredTime time.Duration) (*base.SecurityToken2, error) {
	inlinePolicy := new(base.Policy)
	actions := []string{"vod:ApplyUploadInfo", "vod:CommitUploadInfo"}
	resources := make([]string, 0)
	statement := base.NewAllowStatement(actions, resources)
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)
	return p.SignSts2(inlinePolicy, expiredTime)
}

func (p *Vod) GetUploadAuth() (*base.SecurityToken2, error) {
	return p.GetUploadAuthWithExpiredTime(time.Hour)
}

func (p *Vod) Upload(rd io.Reader, size int64, spaceName string) (string, string, error, int) {
	if size == 0 {
		return "", "", fmt.Errorf("file size is zero"), http.StatusBadRequest
	}

	applyRequest := &request.VodApplyUploadInfoRequest{SpaceName: spaceName}

	resp, code, err := p.ApplyUploadInfo(applyRequest)
	if err != nil {
		return "", "", err, code
	}

	if resp.ResponseMetadata.Error != nil && resp.ResponseMetadata.Error.Code != "0" {
		return "", "", fmt.Errorf("%+v", resp.ResponseMetadata.Error), code
	}

	uploadAddress := resp.GetResult().GetData().GetUploadAddress()
	if uploadAddress != nil {
		if len(uploadAddress.GetUploadHosts()) == 0 {
			return "", "", fmt.Errorf("no tos host found"), http.StatusBadRequest
		}
		if len(uploadAddress.GetStoreInfos()) == 0 && (uploadAddress.GetStoreInfos()[0] == nil) {
			return "", "", fmt.Errorf("no store info found"), http.StatusBadRequest
		}

		tosHost := uploadAddress.GetUploadHosts()[0]
		oid := uploadAddress.StoreInfos[0].GetStoreUri()
		sessionKey := uploadAddress.GetSessionKey()
		auth := uploadAddress.GetStoreInfos()[0].GetAuth()
		client := &http.Client{}

		if int(size) < consts.MinChunckSize {
			bts, err := ioutil.ReadAll(rd)
			if err != nil {
				return "", "", err, http.StatusBadRequest
			}
			if err := p.directUpload(tosHost, oid, auth, bts, client); err != nil {
				return "", "", err, http.StatusBadRequest
			}
		} else {
			uploadPart := model.UploadPartCommon{
				TosHost: tosHost,
				Oid:     oid,
				Auth:    auth,
			}
			isLargeFile := false
			if size > consts.LargeFileSize {
				isLargeFile = true
			}
			if err := p.chunkUpload(rd, uploadPart, client, size, isLargeFile); err != nil {
				return "", "", err, http.StatusBadRequest
			}
		}
		return oid, sessionKey, nil, http.StatusOK
	}
	return "", "", errors.New("upload address not exist"), http.StatusBadRequest
}

func (p *Vod) directUpload(tosHost string, oid string, auth string, fileBytes []byte, client *http.Client) error {
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(fileBytes))
	url := fmt.Sprintf("http://%s/%s", tosHost, oid)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", auth)

	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return err
	}
	res := &model.UploadPartCommonResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}
	if res.Success != 0 {
		return errors.New(res.Error.Message)
	}
	return nil
}

func (p *Vod) chunkUpload(rd io.Reader, uploadPart model.UploadPartCommon, client *http.Client, size int64, isLargeFile bool) error {
	uploadID, err := p.initUploadPart(uploadPart.TosHost, uploadPart.Oid, uploadPart.Auth, client, isLargeFile)
	if err != nil {
		return err
	}
	cur := make([]byte, consts.MinChunckSize)
	parts := make([]string, 0)

	num := size / consts.MinChunckSize
	cnt := 0
	lastNum := int(num) - 1

	// 读 n-1 片并上传上去
	var part string
	for i := 0; i < lastNum; i++ {
		n, err := io.ReadFull(rd, cur)
		if err != nil {
			return err
		}
		cnt += n
		err = retry.Do(func() error {
			part, err = p.uploadPart(uploadPart, uploadID, i, cur, client, isLargeFile)
			return err
		}, retry.Attempts(3))
		if err != nil {
			return err
		}
		parts = append(parts, part)
	}
	// 读 n 和 n+1片（如有）上传上去
	bts, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}
	total := len(bts) + cnt
	if total != int(size) {
		return errors.New(fmt.Sprintf("last part download size mismatch ,download %d , size %d", total, size))
	}
	err = retry.Do(func() error {
		part, err = p.uploadPart(uploadPart, uploadID, lastNum, bts, client, isLargeFile)
		return err
	}, retry.Attempts(3))
	if err != nil {
		return err
	}
	parts = append(parts, part)
	return p.uploadMergePart(uploadPart, uploadID, parts, client, isLargeFile)
}

func (p *Vod) initUploadPart(tosHost string, oid string, auth string, client *http.Client, isLargeFile bool) (string, error) {
	url := fmt.Sprintf("http://%s/%s?uploads", tosHost, oid)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", auth)
	if isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}
	res := &model.UploadPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return "", err
	}
	if res.Success != 0 {
		return "", errors.New(res.Error.Message)
	}
	return res.PayLoad.UploadID, nil
}

func (p *Vod) uploadPart(uploadPart model.UploadPartCommon, uploadID string, partNumber int, data []byte, client *http.Client, isLargeFile bool) (string, error) {
	url := fmt.Sprintf("http://%s/%s?partNumber=%d&uploadID=%s", uploadPart.TosHost, uploadPart.Oid, partNumber, uploadID)
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(data))
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", uploadPart.Auth)
	if isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}
	res := &model.UploadPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return "", err
	}
	if res.Success != 0 {
		return "", errors.New(res.Error.Message)
	}
	return checkSum, nil
}

func (p *Vod) uploadMergePart(uploadPart model.UploadPartCommon, uploadID string, checkSum []string, client *http.Client, isLargeFile bool) error {
	url := fmt.Sprintf("http://%s/%s?uploadID=%s", uploadPart.TosHost, uploadPart.Oid, uploadID)
	body, err := p.genMergeBody(checkSum)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", uploadPart.Auth)
	if isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return err
	}
	res := &model.UploadMergeResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}
	if res.Success != 0 {
		return errors.New(res.Error.Message)
	}
	return nil
}

func (p *Vod) genMergeBody(checkSum []string) (string, error) {
	if len(checkSum) == 0 {
		return "", errors.New("body crc32 empty")
	}
	s := make([]string, len(checkSum))
	for partNumber, crc := range checkSum {
		s[partNumber] = fmt.Sprintf("%d:%s", partNumber, crc)
	}
	return strings.Join(s, ","), nil
}
