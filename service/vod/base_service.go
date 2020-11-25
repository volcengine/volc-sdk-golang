package vod

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/models/vod/response"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (p *Vod) GetPlayAuthToken(req *request.VodGetPlayInfoRequest) (string, error) {
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

func (p *Vod) UploadMediaWithCallback(fileBytes []byte, spaceName string, callbackArgs string, funcs ...Function) (*response.VodCommitUploadInfoResponse, error) {
	return p.UploadMediaInner(fileBytes, spaceName, callbackArgs, funcs...)
}

func (p *Vod) UploadMediaInner(fileBytes []byte, spaceName string, callbackArgs string, funcs ...Function) (*response.VodCommitUploadInfoResponse, error) {
	_, sessionKey, err := p.Upload(fileBytes, spaceName)
	if err != nil {
		return nil, err
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

	commitResp, _, err := p.CommitUploadInfo(commitRequest)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
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

func (p *Vod) Upload(fileBytes []byte, spaceName string) (string, string, error) {
	if len(fileBytes) == 0 {
		return "", "", fmt.Errorf("file size is zero")
	}

	applyRequest := &request.VodApplyUploadInfoRequest{SpaceName: spaceName}

	resp, _, err := p.ApplyUploadInfo(applyRequest)
	if err != nil {
		return "", "", err
	}

	if resp.ResponseMetadata.Error != nil && resp.ResponseMetadata.Error.Code != "0" {
		return "", "", fmt.Errorf("%+v", resp.ResponseMetadata.Error)
	}

	uploadAddress := resp.GetResult().GetData().GetUploadAddress()
	if uploadAddress != nil {
		if len(uploadAddress.GetUploadHosts()) == 0 {
			return "", "", fmt.Errorf("no tos host found")
		}
		if len(uploadAddress.GetStoreInfos()) == 0 && (uploadAddress.GetStoreInfos()[0] == nil) {
			return "", "", fmt.Errorf("no store info found")
		}

		checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(fileBytes))
		tosHost := uploadAddress.GetUploadHosts()[0]
		oid := uploadAddress.StoreInfos[0].GetStoreUri()
		sessionKey := uploadAddress.GetSessionKey()
		auth := uploadAddress.GetStoreInfos()[0].GetAuth()
		url := fmt.Sprintf("http://%s/%s", tosHost, oid)
		req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
		if err != nil {
			return "", "", err
		}
		req.Header.Set("Content-CRC32", checkSum)
		req.Header.Set("Authorization", auth)

		client := &http.Client{}
		rsp, err := client.Do(req)
		if err != nil {
			return "", "", err
		}
		if rsp.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(rsp.Body)
			return "", "", fmt.Errorf("http status=%v, body=%s, remote_addr=%v", rsp.StatusCode, string(b), req.Host)
		}
		defer rsp.Body.Close()

		var tosResp struct {
			Success int         `json:"success"`
			Payload interface{} `json:"payload"`
		}
		err = json.NewDecoder(rsp.Body).Decode(&tosResp)
		if err != nil {
			return "", "", err
		}

		if tosResp.Success != 0 {
			return "", "", fmt.Errorf("tos err:%+v", tosResp)
		}
		return oid, sessionKey, nil
	}
	return "", "", errors.New("upload address not exist")
}
