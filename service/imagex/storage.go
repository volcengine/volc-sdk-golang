package imagex

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/volcengine/volc-sdk-golang/base"
)

type uploadTaskSet struct {
	ctx     context.Context
	host    string
	info    []StoreInfo
	content []io.Reader
	size    []int64

	lock        sync.Mutex
	taskChan    chan *uploadTaskElement
	successOids []string
}

type uploadTaskElement struct {
	ctx     context.Context
	host    string
	info    StoreInfo
	content io.Reader
	size    int64
}

func (r *uploadTaskSet) init() {
	r.lock = sync.Mutex{}
	r.successOids = make([]string, 0)

	r.taskChan = make(chan *uploadTaskElement, len(r.size))

	for idx := range r.size {
		r.taskChan <- &uploadTaskElement{
			ctx:     r.ctx,
			host:    r.host,
			info:    r.info[idx],
			content: r.content[idx],
			size:    r.size[idx],
		}
	}

	close(r.taskChan)
}

func (r *uploadTaskSet) addSuccess(oid string) {
	r.lock.Lock()
	r.successOids = append(r.successOids, oid)
	r.lock.Unlock()
}

func (c *ImageX) ImageXGet(action string, query url.Values, result interface{}) error {
	respBody, _, err := c.Client.Query(action, query)
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

func (c *ImageX) ImageXPost(action string, query url.Values, req, result interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: fail to marshal request, %v", action, err)
	}
	data, _, err := c.Client.Json(action, query, string(body))
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := UnmarshalResultInto(data, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

// GetAllImageServices 获取所有服务信息
func (c *ImageX) GetImageServices(searchPtn string) (*GetServicesResult, error) {
	query := url.Values{}
	if searchPtn != "" {
		query.Add("SearchPtn", searchPtn)
	}

	respBody, _, err := c.Client.Query("GetAllImageServices", query)
	if err != nil {
		return nil, fmt.Errorf("GetAllImageServices: fail to do request, %v", err)
	}

	result := new(GetServicesResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, fmt.Errorf("GetAllImageServices: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// GetServiceDomains 获取服务下的所有域名
func (c *ImageX) GetImageDomains(serviceId string) ([]DomainResult, error) {
	query := url.Values{}
	query.Add("ServiceId", serviceId)

	respBody, _, err := c.Client.Query("GetServiceDomains", query)
	if err != nil {
		return nil, fmt.Errorf("GetServiceDomains: fail to do request, %v", err)
	}

	result := make([]DomainResult, 0)
	if err := UnmarshalResultInto(respBody, &result); err != nil {
		return nil, fmt.Errorf("GetServiceDomains: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// DeleteImageUploadFiles 删除图片
func (c *ImageX) DeleteImages(serviceId string, uris []string) (*DeleteImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", serviceId)
	param := new(DeleteImageParam)
	param.StoreUris = uris

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("DeleteImages: fail to marshal request, %v", err)
	}

	data, _, err := c.Client.Json("DeleteImageUploadFiles", query, string(body))
	if err != nil {
		return nil, fmt.Errorf("DeleteImages: fail to do request, %v", err)
	}

	result := new(DeleteImageResult)
	if err := UnmarshalResultInto(data, result); err != nil {
		return nil, fmt.Errorf("DeleteImages: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// ApplyImageUpload 获取图片上传地址
func (c *ImageX) ApplyUploadImage(params *ApplyUploadImageParam) (*ApplyUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	if params.SessionKey != "" {
		query.Add("SessionKey", params.SessionKey)
	}
	if params.UploadNum > 0 {
		query.Add("UploadNum", strconv.Itoa(params.UploadNum))
	}
	for _, key := range params.StoreKeys {
		query.Add("StoreKeys", key)
	}

	respBody, _, err := c.Client.Query("ApplyImageUpload", query)
	if err != nil {
		return nil, fmt.Errorf("ApplyUploadImage: fail to do request, %v", err)
	}

	result := new(ApplyUploadImageResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, fmt.Errorf("ApplyUploadImage: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// CommitImageUpload 图片上传完成上报
func (c *ImageX) CommitUploadImage(params *CommitUploadImageParam) (*CommitUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)

	bts, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("CommitUploadImage: fail to marshal request, %v", err)
	}

	respBody, _, err := c.Client.Json("CommitImageUpload", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("CommitUploadImage: fail to do request, %v", err)
	}

	result := new(CommitUploadImageResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) segmentedUpload(item *uploadTaskElement) error {
	if item.size <= MinChunkSize {
		//goland:noinspection GoDeprecation
		bts, err := ioutil.ReadAll(item.content)
		if err != nil {
			return err
		}
		err = c.directUpload(item.ctx, item.host, item.info, bts)
		if err != nil {
			return err
		}
	} else {
		arg := &segmentedUploadParam{
			host:        item.host,
			StoreInfo:   item.info,
			content:     item.content,
			size:        item.size,
			isLargeFile: item.size > LargeFileSize,
		}
		err := arg.chunkUpload()
		if err != nil {
			return err
		}
	}
	return nil
}

// 上传图片
// 请确保 content 长度和 size 长度一致
func (c *ImageX) SegmentedUploadImages(ctx context.Context, params *ApplyUploadImageParam, content []io.Reader, size []int64) (*CommitUploadImageResult, error) {
	if len(content) != len(size) {
		return nil, fmt.Errorf("expect len(size) == len(content), but len(size) = %d, len(content) = %d", len(size), len(content))
	}

	params.UploadNum = len(content)

	for idx, item := range size {
		if item == 0 {
			return nil, fmt.Errorf("size[%d] is zero", idx)
		}
	}

	// 1. apply
	applyResp, err := c.ApplyUploadImage(params)
	if err != nil {
		return nil, err
	}

	uploadAddr := applyResp.UploadAddress
	if len(uploadAddr.UploadHosts) == 0 {
		return nil, fmt.Errorf("UploadImages: no upload host found, request id %s", applyResp.RequestId)
	}
	if len(uploadAddr.StoreInfos) != params.UploadNum {
		return nil, fmt.Errorf("UploadImages: store infos num %d != upload num %d, request id %s",
			len(uploadAddr.StoreInfos), params.UploadNum, applyResp.RequestId)
	}

	// 2. upload
	host := uploadAddr.UploadHosts[0]

	wg := &sync.WaitGroup{}
	uploadTaskSet := &uploadTaskSet{
		ctx:     ctx,
		host:    host,
		info:    uploadAddr.StoreInfos,
		content: content,
		size:    size,
	}
	uploadTaskSet.init()

	for i := 0; i < UploadRoutines; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				recover()
				wg.Done()
			}()

			for item := range uploadTaskSet.taskChan {
				err := c.segmentedUpload(item)
				if err == nil {
					uploadTaskSet.addSuccess(item.info.StoreUri)
				}
			}
		}()
	}
	wg.Wait()

	if len(uploadTaskSet.successOids) == 0 {
		return nil, fmt.Errorf("upload failed, no file uploaded")
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:   params.ServiceId,
		SessionKey:  uploadAddr.SessionKey,
		SuccessOids: uploadTaskSet.successOids,
	}
	if params.CommitParam != nil {
		commitParams.Functions = params.CommitParam.Functions
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}

// 上传图片
func (c *ImageX) UploadImages(params *ApplyUploadImageParam, images [][]byte) (*CommitUploadImageResult, error) {
	params.UploadNum = len(images)

	// 1. apply
	applyResp, err := c.ApplyUploadImage(params)
	if err != nil {
		return nil, err
	}
	uploadAddr := applyResp.UploadAddress
	if len(uploadAddr.UploadHosts) == 0 {
		return nil, fmt.Errorf("UploadImages: no upload host found, request id %s", applyResp.RequestId)
	} else if len(uploadAddr.StoreInfos) != params.UploadNum {
		return nil, fmt.Errorf("UploadImages: store infos num %d != upload num %d, request id %s",
			len(uploadAddr.StoreInfos), params.UploadNum, applyResp.RequestId)
	}

	// 2. upload
	success := make([]string, 0)
	host := uploadAddr.UploadHosts[0]
	for i, image := range images {
		imageCopy := image
		info := uploadAddr.StoreInfos[i]
		err = retry.Do(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), c.ServiceInfo.Timeout)
			defer cancel()
			return c.directUpload(ctx, host, info, imageCopy)
		}, retry.Attempts(3))
		if err != nil {
			fmt.Printf("UploadImages: fail to do upload, %v", err)
		} else {
			success = append(success, info.StoreUri)
		}
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:   params.ServiceId,
		SessionKey:  uploadAddr.SessionKey,
		SuccessOids: success,
	}
	if params.CommitParam != nil {
		commitParams.Functions = params.CommitParam.Functions
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}

// 获取临时上传凭证
func (c *ImageX) GetUploadAuthToken(query url.Values) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}

	applyUploadToken, err := c.Client.GetSignUrl("ApplyImageUpload", query)
	if err != nil {
		return "", fmt.Errorf("GetUploadAuthToken: fail to get apply token, %v", err)
	}
	ret["ApplyUploadToken"] = applyUploadToken

	commitUploadToken, err := c.Client.GetSignUrl("CommitImageUpload", query)
	if err != nil {
		return "", fmt.Errorf("GetUploadAuthToken: fail to get commit token, %v", err)
	}
	ret["CommitUploadToken"] = commitUploadToken

	b, err := json.Marshal(ret)
	if err != nil {
		return "", fmt.Errorf("GetUploadAuthToken: fail to marshal token, %v", err)
	}
	token := strings.ReplaceAll(string(b), "\\u0026", "&")
	return base64.StdEncoding.EncodeToString([]byte(token)), nil
}

type UploadAuthOpt func(option *uploadAuthOption)

type uploadAuthOption struct {
	keyPtn string
}

func WithUploadKeyPtn(ptn string) UploadAuthOpt {
	return func(o *uploadAuthOption) {
		o.keyPtn = ptn
	}
}

// 获取上传临时密钥
func (c *ImageX) GetUploadAuth(serviceIds []string, opt ...UploadAuthOpt) (*base.SecurityToken2, error) {
	return c.GetUploadAuthWithExpire(serviceIds, time.Hour, opt...)
}

// 获取上传临时密钥
func (c *ImageX) GetUploadAuthWithExpire(serviceIds []string, expire time.Duration, opt ...UploadAuthOpt) (*base.SecurityToken2, error) {
	serviceIdRes := make([]string, 0)
	if len(serviceIds) == 0 {
		serviceIdRes = append(serviceIdRes, fmt.Sprintf(ResourceServiceIdTRN, "*"))
	} else {
		for _, sid := range serviceIds {
			serviceIdRes = append(serviceIdRes, fmt.Sprintf(ResourceServiceIdTRN, sid))
		}
	}
	op := new(uploadAuthOption)
	for _, o := range opt {
		o(op)
	}
	storeKeyRes := []string{
		fmt.Sprintf(ResourceStoreKeyTRN, op.keyPtn),
	}

	inlinePolicy := new(base.Policy)
	inlinePolicy.Statement = append(inlinePolicy.Statement, base.NewAllowStatement([]string{
		"ImageX:ApplyImageUpload",
	}, append(serviceIdRes, storeKeyRes...)))
	inlinePolicy.Statement = append(inlinePolicy.Statement, base.NewAllowStatement([]string{
		"ImageX:CommitImageUpload",
	}, serviceIdRes))

	return c.Client.SignSts2(inlinePolicy, expire)
}

func (c *ImageX) CreateImageContentTask(req *CreateImageContentTaskReq) (*CreateImageContentTaskResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &CreateImageContentTaskResp{}
	err = c.ImageXPost("CreateImageContentTask", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageContentTaskDetail(req *GetImageContentTaskDetailReq) (*GetImageContentTaskDetailResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageContentTaskDetailResp{}
	err = c.ImageXPost("GetImageContentTaskDetail", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageContentBlockList(req *GetImageContentBlockListReq) (*GetImageContentBlockListResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageContentBlockListResp{}
	err = c.ImageXPost("GetImageContentBlockList", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) FetchImageUrl(req *FetchUrlReq) (*FetchUrlResp, error) {
	resp := new(FetchUrlResp)
	err := c.ImageXPost("FetchImageUrl", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetImageUploadFile(param *GetImageUploadFileParam) (*GetImageUploadFileResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)
	query.Add("StoreUri", param.StoreUri)

	result := new(GetImageUploadFileResult)
	err := c.ImageXGet("GetImageUploadFile", query, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImageUploadFiles(param *GetImageUploadFilesParam) (*GetImageUploadFilesResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)
	query.Add("Limit", strconv.Itoa(param.Limit))
	query.Add("Marker", strconv.FormatInt(param.Marker, 10))

	result := new(GetImageUploadFilesResult)
	err := c.ImageXGet("GetImageUploadFiles", query, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
