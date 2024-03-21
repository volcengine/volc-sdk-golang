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

const (
	ResourceServiceIdTRN = "trn:ImageX:*:*:ServiceId/%s"
	ResourceStoreKeyTRN  = "trn:ImageX:*:*:StoreKeys/%s"

	MinChunkSize   = 1024 * 1024 * 20
	LargeFileSize  = 1024 * 1024 * 1024
	UploadRoutines = 4
)

func (c *Imagex) ImageXGet(action string, query url.Values, result interface{}) error {
	respBody, _, err := c.Client.Query(action, query)
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := unmarshalInto(respBody, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

func (c *Imagex) ImageXPost(action string, query url.Values, req, result interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: fail to marshal request, %v", action, err)
	}
	data, _, err := c.Client.Json(action, query, string(body))
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := unmarshalInto(data, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

// ApplyImageUpload 获取图片上传地址
func (c *Imagex) ApplyUploadImage(params *ApplyUploadImageParam) (*ApplyUploadImageResult, error) {
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
	query.Add("Overwrite", strconv.FormatBool(params.Overwrite))

	respBody, _, err := c.Client.Query("ApplyImageUpload", query)
	if err != nil {
		return nil, fmt.Errorf("ApplyUploadImage: fail to do request, %v", err)
	}

	result := new(ApplyUploadImageResult)
	if err := unmarshalResultInto(respBody, result); err != nil {
		return nil, fmt.Errorf("ApplyUploadImage: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// CommitImageUpload 图片上传完成上报
func (c *Imagex) CommitUploadImage(params *CommitUploadImageParam) (*CommitUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	query.Add("SkipMeta", fmt.Sprintf("%v", params.SkipMeta))

	bts, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("CommitUploadImage: fail to marshal request, %v", err)
	}

	respBody, _, err := c.Client.Json("CommitImageUpload", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("CommitUploadImage: fail to do request, %v", err)
	}

	result := new(CommitUploadImageResult)
	if err := unmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Imagex) segmentedUpload(set *uploadTaskSet, item *uploadTaskElement) error {
	if item.size <= MinChunkSize {
		//goland:noinspection GoDeprecation
		bts, err := ioutil.ReadAll(item.content)
		if err != nil {
			return err
		}
		err = c.directUpload(item.ctx, item.host, item.idx, set, item.info, bts)
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
			idx:         item.idx,
			set:         set,
		}
		err := arg.chunkUpload()
		if err != nil {
			return err
		}
	}
	set.result[item.idx].success = true
	return nil
}

// 上传图片
// 请确保 content 长度和 size 长度一致
func (c *Imagex) SegmentedUploadImages(ctx context.Context, params *ApplyUploadImageParam, content []io.Reader, size []int64) (*CommitUploadImageResult, error) {
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
				uploadTaskSet.result[item.idx] = uploadTaskResult{
					uri: item.info.StoreUri,
				}
				err := c.segmentedUpload(uploadTaskSet, item)
				if err == nil {
					uploadTaskSet.result[item.idx].success = true
					uploadTaskSet.addSuccess(item.info.StoreUri)
				} else {
					uploadTaskSet.result[item.idx].errMsg = err.Error()
				}
			}
		}()
	}
	wg.Wait()
	if len(uploadTaskSet.successOids) == 0 || params.SkipCommit {
		return &CommitUploadImageResult{
			Results: uploadTaskSet.GetResult(),
		}, nil
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:   params.ServiceId,
		SessionKey:  uploadAddr.SessionKey,
		SuccessOids: uploadTaskSet.successOids,
		SkipMeta:    params.SkipMeta,
	}
	if params.CommitParam != nil {
		commitParams.Functions = params.CommitParam.Functions
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	uploadTaskSet.fill(commitResp)
	return commitResp, nil
}

// 上传图片
func (c *Imagex) UploadImages(params *ApplyUploadImageParam, images [][]byte) (*CommitUploadImageResult, error) {
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
	host := uploadAddr.UploadHosts[0]
	uploadTaskSet := &uploadTaskSet{
		ctx:     context.Background(),
		host:    host,
		info:    uploadAddr.StoreInfos,
		content: make([]io.Reader, 0),
		size:    make([]int64, 0),
	}
	uploadTaskSet.init()
	for i, image := range images {
		imageCopy := image
		info := uploadAddr.StoreInfos[i]
		idx := i
		uploadTaskSet.result[idx] = uploadTaskResult{
			uri: info.StoreUri,
		}
		err = retry.Do(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), c.ServiceInfo.Timeout)
			defer cancel()
			return c.directUpload(ctx, host, idx, uploadTaskSet, info, imageCopy)
		}, retry.Attempts(2))
		if err != nil {
			uploadTaskSet.result[idx].errMsg = err.Error()
		} else {
			uploadTaskSet.result[idx].success = true
			uploadTaskSet.result[idx].errMsg = ""
			uploadTaskSet.result[idx].putErr = nil
			uploadTaskSet.addSuccess(info.StoreUri)
		}
	}
	if len(uploadTaskSet.successOids) == 0 || params.SkipCommit {
		return &CommitUploadImageResult{
			Results: uploadTaskSet.GetResult(),
		}, nil
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:   params.ServiceId,
		SessionKey:  uploadAddr.SessionKey,
		SuccessOids: uploadTaskSet.successOids,
		SkipMeta:    params.SkipMeta,
	}
	if params.CommitParam != nil {
		commitParams.Functions = params.CommitParam.Functions
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	uploadTaskSet.fill(commitResp)
	return commitResp, nil
}

// 获取临时上传凭证
func (c *Imagex) GetUploadAuthToken(query url.Values) (string, error) {
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
	keyPtn     string
	conditions map[string]string
}

func WithUploadKeyPtn(ptn string) UploadAuthOpt {
	return func(o *uploadAuthOption) {
		o.keyPtn = ptn
	}
}

func WithUploadOverwrite(overwrite bool) UploadAuthOpt {
	return func(op *uploadAuthOption) {
		op.conditions["UploadOverwrite"] = strconv.FormatBool(overwrite)
	}
}

func WithUploadPolicy(policy *UploadPolicy) UploadAuthOpt {
	res, _ := json.Marshal(policy)
	return func(op *uploadAuthOption) {
		op.conditions["UploadPolicy"] = string(res)
	}
}

// 获取上传临时密钥
func (c *Imagex) GetUploadAuth(serviceIds []string, opt ...UploadAuthOpt) (*base.SecurityToken2, error) {
	return c.GetUploadAuthWithExpire(serviceIds, time.Hour, opt...)
}

// 获取上传临时密钥
func (c *Imagex) GetUploadAuthWithExpire(serviceIds []string, expire time.Duration, opt ...UploadAuthOpt) (*base.SecurityToken2, error) {
	serviceIdRes := make([]string, 0)
	if len(serviceIds) == 0 {
		serviceIdRes = append(serviceIdRes, fmt.Sprintf(ResourceServiceIdTRN, "*"))
	} else {
		for _, sid := range serviceIds {
			serviceIdRes = append(serviceIdRes, fmt.Sprintf(ResourceServiceIdTRN, sid))
		}
	}
	op := &uploadAuthOption{
		conditions: map[string]string{},
	}
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
	for k, v := range op.conditions {
		inlinePolicy.Statement = append(inlinePolicy.Statement, base.NewAllowStatement([]string{k}, []string{v}))
	}

	return c.Client.SignSts2(inlinePolicy, expire)
}
