package vod

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/volcengine/volc-sdk-golang/base"
	model_base "github.com/volcengine/volc-sdk-golang/service/base/models/base"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/business"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/response"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/volcengine/volc-sdk-golang/service/vod/upload/consts"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/model"
)

func (p *Vod) GetSubtitleAuthToken(req *request.VodGetSubtitleInfoListRequest, tokenExpireTime int) (string, error) {
	if len(req.GetVid()) == 0 {
		return "", errors.New("传入的Vid为空")
	}
	query := url.Values{
		"Vid":    []string{req.GetVid()},
		"Status": []string{"Published"},
	}

	if tokenExpireTime > 0 {
		query.Add("X-Expires", strconv.Itoa(tokenExpireTime))
	}

	if getSubtitleInfoAuthToken, err := p.GetSignUrl("GetSubtitleInfoList", query); err == nil {
		ret := map[string]string{}
		ret["GetSubtitleAuthToken"] = getSubtitleInfoAuthToken
		b, err := json.Marshal(ret)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(b), nil
	} else {
		return "", err
	}
}

func (p *Vod) GetPrivateDrmAuthToken(req *request.VodGetPrivateDrmPlayAuthRequest, tokenExpireTime int) (string, error) {
	if len(req.GetVid()) == 0 {
		return "", errors.New("传入的Vid为空")
	}
	query := url.Values{
		"Vid": []string{req.GetVid()},
	}

	if len(req.GetPlayAuthIds()) > 0 {
		query.Add("PlayAuthIds", req.GetPlayAuthIds())
	}
	if len(req.GetDrmType()) > 0 {
		query.Add("DrmType", req.GetDrmType())
		switch req.GetDrmType() {
		case "appdevice", "webdevice":
			if len(req.GetUnionInfo()) == 0 {
				return "", errors.New("invalid unionInfo")
			}
		}
	}
	if len(req.GetUnionInfo()) > 0 {
		query.Add("UnionInfo", req.GetUnionInfo())
	}
	if tokenExpireTime > 0 {
		query.Add("X-Expires", strconv.Itoa(tokenExpireTime))
	}

	if getPrivateDrmAuthToken, err := p.GetSignUrl("GetPrivateDrmPlayAuth", query); err == nil {
		return getPrivateDrmAuthToken, nil
	} else {
		return "", err
	}
}

func (p *Vod) CreateSha1HlsDrmAuthToken(expireSeconds int64) (auth string, err error) {
	return p.createHlsDrmAuthToken(DSAHmacSha1, expireSeconds)
}

func (p *Vod) createHlsDrmAuthToken(authAlgorithm string, expireSeconds int64) (string, error) {
	if expireSeconds == 0 {
		return "", errors.New("invalid expire")
	}

	token, err := createAuth(authAlgorithm, Version2, p.ServiceInfo.Credentials.AccessKeyID,
		p.ServiceInfo.Credentials.SecretAccessKey, p.ServiceInfo.Credentials.Region, expireSeconds)
	if err != nil {
		return "", err
	}

	query := url.Values{}
	query.Set("DrmAuthToken", token)
	query.Set("X-Expires", strconv.FormatInt(expireSeconds, 10))
	if getAuth, err := p.GetSignUrl("GetHlsDecryptionKey", query); err == nil {
		return getAuth, nil
	} else {
		return "", err
	}
}

func (p *Vod) GetPlayAuthToken(req *request.VodGetPlayInfoRequest, tokenExpireTime int) (string, error) {
	if len(req.GetVid()) == 0 {
		return "", errors.New("传入的Vid为空")
	}
	query := url.Values{}
	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	jsonData := marshaler.Format(req)
	reqMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonData), &reqMap)
	if err != nil {
		return "", err
	}
	for k, v := range reqMap {
		var sv string
		switch ov := v.(type) {
		case string:
			sv = ov
		case int:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int8:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint8:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int16:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint16:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int32:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint32:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int64:
			sv = strconv.FormatInt(ov, 10)
		case uint64:
			sv = strconv.FormatUint(ov, 10)
		case bool:
			sv = strconv.FormatBool(ov)
		case float32:
			sv = strconv.FormatFloat(float64(ov), 'f', -1, 32)
		case float64:
			sv = strconv.FormatFloat(ov, 'f', -1, 64)
		case []byte:
			sv = string(ov)
		default:
			v2, e2 := json.Marshal(ov)
			if e2 != nil {
				return "", e2
			}
			sv = string(v2)
		}
		query.Set(k, sv)
	}
	if tokenExpireTime > 0 {
		query.Add("X-Expires", strconv.Itoa(tokenExpireTime))
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

func (p *Vod) UploadObjectWithCallback(filePath string, spaceName string, callbackArgs string, fileName, fileExtension string, funcs string) (*response.VodCommitUploadInfoResponse, int, error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, -1, err
	}

	req := &model.VodUploadMediaInnerFuncRequest{
		FilePath:      filePath,
		Rd:            file,
		Size:          stat.Size(),
		SpaceName:     spaceName,
		FileType:      "object",
		CallbackArgs:  callbackArgs,
		Funcs:         funcs,
		FileName:      fileName,
		FileExtension: fileExtension,
	}
	return p.UploadMediaInner(req)
}

func (p *Vod) UploadMediaWithCallback(mediaRequset *request.VodUploadMediaRequest) (*response.VodCommitUploadInfoResponse, int, error) {
	file, err := os.Open(filepath.Clean(mediaRequset.GetFilePath()))
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, -1, err
	}

	req := &model.VodUploadMediaInnerFuncRequest{
		FilePath:        mediaRequset.GetFilePath(),
		Rd:              file,
		Size:            stat.Size(),
		ParallelNum:     int(mediaRequset.GetParallelNum()),
		SpaceName:       mediaRequset.GetSpaceName(),
		CallbackArgs:    mediaRequset.GetCallbackArgs(),
		Funcs:           mediaRequset.GetFunctions(),
		FileName:        mediaRequset.GetFileName(),
		FileExtension:   mediaRequset.GetFileExtension(),
		VodUploadSource: mediaRequset.GetVodUploadSource(),
		StorageClass:    mediaRequset.StorageClass,
	}
	return p.UploadMediaInner(req)
}

func (p *Vod) UploadMaterialWithCallback(materialRequest *request.VodUploadMaterialRequest) (*response.VodCommitUploadInfoResponse, int, error) {
	file, err := os.Open(filepath.Clean(materialRequest.GetFilePath()))
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, -1, err
	}

	req := &model.VodUploadMediaInnerFuncRequest{
		FilePath:      materialRequest.GetFilePath(),
		Rd:            file,
		Size:          stat.Size(),
		ParallelNum:   int(materialRequest.GetParallelNum()),
		SpaceName:     materialRequest.GetSpaceName(),
		FileType:      materialRequest.GetFileType(),
		CallbackArgs:  materialRequest.GetCallbackArgs(),
		Funcs:         materialRequest.GetFunctions(),
		FileName:      materialRequest.GetFileName(),
		FileExtension: materialRequest.GetFileExtension(),
	}
	return p.UploadMediaInner(req)
}

func (p *Vod) UploadMediaInner(uploadMediaInnerRequest *model.VodUploadMediaInnerFuncRequest) (*response.VodCommitUploadInfoResponse, int, error) {
	req := &model.VodUploadFuncRequest{
		FilePath:      uploadMediaInnerRequest.FilePath,
		Rd:            uploadMediaInnerRequest.Rd,
		Size:          uploadMediaInnerRequest.Size,
		ParallelNum:   uploadMediaInnerRequest.ParallelNum,
		SpaceName:     uploadMediaInnerRequest.SpaceName,
		FileType:      uploadMediaInnerRequest.FileType,
		FileName:      uploadMediaInnerRequest.FileName,
		FileExtension: uploadMediaInnerRequest.FileExtension,
		StorageClass:  uploadMediaInnerRequest.StorageClass,
	}
	logId, sessionKey, err, code := p.Upload(req)
	if err != nil {
		return p.fillCommitUploadInfoResponseWhenError(logId, err.Error()), code, err
	}

	commitRequest := &request.VodCommitUploadInfoRequest{
		SpaceName:       uploadMediaInnerRequest.SpaceName,
		SessionKey:      sessionKey,
		CallbackArgs:    uploadMediaInnerRequest.CallbackArgs,
		Functions:       uploadMediaInnerRequest.Funcs,
		VodUploadSource: uploadMediaInnerRequest.VodUploadSource,
	}

	commitResp, code, err := p.CommitUploadInfo(commitRequest)
	if err != nil {
		return commitResp, code, err
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

func (p *Vod) fillCommitUploadInfoResponseWhenError(logId, errMsg string) *response.VodCommitUploadInfoResponse {
	commitUploadInfoRespone := &response.VodCommitUploadInfoResponse{
		ResponseMetadata: &model_base.ResponseMetadata{
			RequestId: logId,
			Service:   "vod",
			Error:     &model_base.ResponseError{Message: errMsg},
		},
	}
	return commitUploadInfoRespone
}

func (p *Vod) Upload(vodUploadFuncRequest *model.VodUploadFuncRequest) (string, string, error, int) {
	if vodUploadFuncRequest.Size == 0 {
		return "", "", fmt.Errorf("file size is zero"), http.StatusBadRequest
	}

	applyRequest := &request.VodApplyUploadInfoRequest{
		SpaceName:     vodUploadFuncRequest.SpaceName,
		FileType:      vodUploadFuncRequest.FileType,
		FileName:      vodUploadFuncRequest.FileName,
		FileExtension: vodUploadFuncRequest.FileExtension,
		StorageClass:  vodUploadFuncRequest.StorageClass,
	}

	resp, code, err := p.ApplyUploadInfo(applyRequest)
	logId := resp.GetResponseMetadata().GetRequestId()
	if err != nil {
		return logId, "", err, code
	}

	if resp.ResponseMetadata.Error != nil && resp.ResponseMetadata.Error.Code != "0" {
		return logId, "", fmt.Errorf("%+v", resp.ResponseMetadata.Error), code
	}

	uploadAddress := resp.GetResult().GetData().GetUploadAddress()
	if uploadAddress != nil {
		if len(uploadAddress.GetUploadHosts()) == 0 {
			return logId, "", fmt.Errorf("no tos host found"), http.StatusBadRequest
		}
		if len(uploadAddress.GetStoreInfos()) == 0 && (uploadAddress.GetStoreInfos()[0] == nil) {
			return logId, "", fmt.Errorf("no store info found"), http.StatusBadRequest
		}

		tosHost := uploadAddress.GetUploadHosts()[0]
		oid := uploadAddress.StoreInfos[0].GetStoreUri()
		sessionKey := uploadAddress.GetSessionKey()
		auth := uploadAddress.GetStoreInfos()[0].GetAuth()
		client := &http.Client{}

		if int(vodUploadFuncRequest.Size) < consts.MinChunckSize {
			bts, err := ioutil.ReadAll(vodUploadFuncRequest.Rd)
			if err != nil {
				return logId, "", err, http.StatusBadRequest
			}
			if err := p.directUpload(tosHost, oid, auth, bts, client, vodUploadFuncRequest.StorageClass); err != nil {
				return logId, "", err, http.StatusBadRequest
			}
		} else {
			uploadPart := model.UploadPartCommon{
				TosHost: tosHost,
				Oid:     oid,
				Auth:    auth,
			}
			if vodUploadFuncRequest.ParallelNum == 0 {
				vodUploadFuncRequest.ParallelNum = 1
			}
			if err := p.chunkUpload(vodUploadFuncRequest.FilePath, uploadPart, client, vodUploadFuncRequest.Size, vodUploadFuncRequest.ParallelNum, vodUploadFuncRequest.StorageClass); err != nil {
				return logId, "", err, http.StatusBadRequest
			}
		}
		return oid, sessionKey, nil, http.StatusOK
	}
	return logId, "", errors.New("upload address not exist"), http.StatusBadRequest
}

func (p *Vod) directUpload(tosHost string, oid string, auth string, fileBytes []byte, client *http.Client, storageClass int32) error {
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(fileBytes))
	url := fmt.Sprintf("http://%s/%s", tosHost, oid)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", auth)

	if storageClass == int32(business.StorageClassType_Archive) {
		req.Header.Set("X-Upload-Storage-Class", "archive")
	}
	if storageClass == int32(business.StorageClassType_IA) {
		req.Header.Set("X-Upload-Storage-Class", "ia")
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
	res := &model.UploadPartCommonResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}
	if res.Success != 0 {
		return errors.New(res.Error.Message)
	}
	return nil
}

type UploadPartInfo struct {
	Number int
	OffSet int64
	Size   int64
}

func calculateUploadParts(size int64) (int, []*UploadPartInfo, error) {
	totalNum := size / consts.MinChunckSize
	if totalNum > 10000 {
		return 0, nil, errors.New("parts over 10000")
	}

	uploadPartInfos := make([]*UploadPartInfo, 0)
	for i := 0; i < int(totalNum); i++ {
		partInfo := &UploadPartInfo{
			Number: i + 1,
			OffSet: int64(i * consts.MinChunckSize),
			Size:   consts.MinChunckSize,
		}
		uploadPartInfos = append(uploadPartInfos, partInfo)
	}

	last := size % consts.MinChunckSize
	if last != 0 {
		uploadPartInfos[totalNum-1].Size += last
	}
	return int(totalNum), uploadPartInfos, nil
}

type Jobs struct {
	filePath         string
	uploadPartInfo   *UploadPartInfo
	uploadPartCommon *model.UploadPartCommon
	uploadId         string
	client           *http.Client
	storageClass     int32
}

func worker(p *Vod, jobs <-chan *Jobs, results chan<- *model.UploadPartResponse, errChan chan<- *model.UploadPartResponse, quit *int, objectContentType *string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		if *quit != 0 {
			continue
		}
		fd, err := os.Open(job.filePath)
		if err != nil {
			res := &model.UploadPartResponse{
				UploadPartCommonResponse: model.UploadPartCommonResponse{
					Success: -1,
					Error: model.UploadPartError{
						Code:    -1,
						Error:   err.Error(),
						Message: "open file failed",
					},
				},
				PartNumber: job.uploadPartInfo.Number,
			}
			errChan <- res
			*quit = -1
			return
		}

		data := make([]byte, job.uploadPartInfo.Size)

		_, err = fd.ReadAt(data, job.uploadPartInfo.OffSet)
		if err != nil {
			res := &model.UploadPartResponse{
				UploadPartCommonResponse: model.UploadPartCommonResponse{
					Success: -1,
					Error: model.UploadPartError{
						Code:    -1,
						Error:   err.Error(),
						Message: "read data error",
					},
				},
				PartNumber: job.uploadPartInfo.Number,
			}
			errChan <- res
			*quit = -1
			return
		}

		var resp *model.UploadPartResponse
		retry.Do(func() error {
			resp, err = p.uploadPart(*job.uploadPartCommon, job.uploadId, job.uploadPartInfo.Number, data, job.client, job.storageClass)
			return err
		}, retry.Attempts(3))
		if err != nil {
			res := &model.UploadPartResponse{
				UploadPartCommonResponse: model.UploadPartCommonResponse{
					Success: -1,
					Error: model.UploadPartError{
						Code:    -1,
						Error:   err.Error(),
						Message: "upload part fail",
					},
				},
				PartNumber: job.uploadPartInfo.Number,
			}
			errChan <- res
			*quit = -1
			return
		}
		results <- resp
		if job.uploadPartInfo.Number == 1 {
			*objectContentType = resp.PayLoad.Meta.ObjectContentType
		}
	}
}

func (p *Vod) chunkUpload(filePath string, uploadPartCommon model.UploadPartCommon, client *http.Client, size int64, parallelNum int, storageClass int32) error {
	// 1. 计算分片
	totalNum, uploadPartInfos, err := calculateUploadParts(size)
	if err != nil {
		return err
	}

	// 2. Init 初始化分片
	uploadID, err := p.InitUploadPart(uploadPartCommon.TosHost, uploadPartCommon.Oid, uploadPartCommon.Auth, client, storageClass)
	if err != nil {
		return err
	}

	chJobs := make(chan *Jobs, totalNum)
	chUploadPartRes := make(chan *model.UploadPartResponse, totalNum)
	errChan := make(chan *model.UploadPartResponse, totalNum)
	quitSig := 0
	wg := sync.WaitGroup{}
	wg.Add(parallelNum)
	objectContentType := ""

	// 3. StartWorker
	for w := 1; w <= parallelNum; w++ {
		go worker(p, chJobs, chUploadPartRes, errChan, &quitSig, &objectContentType, &wg)
	}

	// 4. PushJobs
	for i := 0; i < totalNum; i++ {
		job := &Jobs{
			filePath:         filePath,
			uploadPartInfo:   uploadPartInfos[i],
			uploadPartCommon: &uploadPartCommon,
			uploadId:         uploadID,
			client:           client,
			storageClass:     storageClass,
		}
		chJobs <- job
	}
	close(chJobs)

	// 5. recive results
	wg.Wait()
	select {
	case v := <-errChan:
		close(chUploadPartRes)
		close(errChan)
		return fmt.Errorf("Error=%s,Message is %s", v.UploadPartCommonResponse.Error.Error, v.UploadPartCommonResponse.Error.Message)
	default:
		// all parts upload success
		close(errChan)
	}
	uploadPartResponseList := make([]*model.UploadPartResponse, 0)
	for i := 0; i < totalNum; i++ {
		uploadPartResponseList = append(uploadPartResponseList, <-chUploadPartRes)
	}
	close(chUploadPartRes)
	uploadPartCommon.ObjectContentType = objectContentType
	return p.UploadMergePart(uploadPartCommon, uploadID, uploadPartResponseList, client, storageClass)
}

func (p *Vod) UploadMergePart(uploadPart model.UploadPartCommon, uploadID string, uploadPartResponseList []*model.UploadPartResponse, client *http.Client, storageClass int32) error {
	url := fmt.Sprintf("http://%s/%s?uploadID=%s", uploadPart.TosHost, uploadPart.Oid, uploadID)
	body, err := p.genMergeBody(uploadPartResponseList)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", uploadPart.Auth)
	req.Header.Set("X-Storage-Mode", "gateway")

	if storageClass == int32(business.StorageClassType_Archive) || storageClass == int32(business.StorageClassType_IA) {
		if storageClass == int32(business.StorageClassType_Archive) {
			req.Header.Set("X-Upload-Storage-Class", "archive")
		}
		if storageClass == int32(business.StorageClassType_IA) {
			req.Header.Set("X-Upload-Storage-Class", "ia")
		}
		if uploadPart.ObjectContentType != "" {
			q := req.URL.Query()
			q.Add("ObjectContentType", uploadPart.ObjectContentType)
			req.URL.RawQuery = q.Encode()
		}
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

func (p *Vod) genMergeBody(uploadPartResponseList []*model.UploadPartResponse) (string, error) {
	if len(uploadPartResponseList) == 0 {
		return "", errors.New("body crc32 empty")
	}
	s := make([]string, len(uploadPartResponseList))
	for _, v := range uploadPartResponseList {
		s[v.PartNumber-1] = fmt.Sprintf("%d:%s", v.PartNumber, v.CheckSum)
	}
	return strings.Join(s, ","), nil
}

func (p *Vod) uploadPart(uploadPart model.UploadPartCommon, uploadID string, partNumber int, data []byte, client *http.Client, storageClass int32) (*model.UploadPartResponse, error) {
	url := fmt.Sprintf("http://%s/%s?partNumber=%d&uploadID=%s", uploadPart.TosHost, uploadPart.Oid, partNumber, uploadID)
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(data))
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", uploadPart.Auth)
	req.Header.Set("X-Storage-Mode", "gateway")

	if storageClass == int32(business.StorageClassType_Archive) {
		req.Header.Set("X-Upload-Storage-Class", "archive")
	}
	if storageClass == int32(business.StorageClassType_IA) {
		req.Header.Set("X-Upload-Storage-Class", "ia")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}
	res := &model.UploadPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return nil, err
	}
	if res.Success != 0 {
		return nil, errors.New(res.Error.Message)
	}
	res.PartNumber = partNumber
	res.CheckSum = checkSum
	//return checkSum, res.PayLoad.Meta.ObjectContentType, nil
	return res, nil
}

func (p *Vod) InitUploadPart(tosHost string, oid string, auth string, client *http.Client, storageClass int32) (string, error) {
	url := fmt.Sprintf("http://%s/%s?uploads", tosHost, oid)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", auth)
	req.Header.Set("X-Storage-Mode", "gateway")

	if storageClass == int32(business.StorageClassType_Archive) {
		req.Header.Set("X-Upload-Storage-Class", "archive")
	}
	if storageClass == int32(business.StorageClassType_IA) {
		req.Header.Set("X-Upload-Storage-Class", "ia")
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
	res := &model.InitPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return "", err
	}
	if res.Success != 0 {
		return "", errors.New(res.Error.Message)
	}
	return res.PayLoad.UploadID, nil
}
