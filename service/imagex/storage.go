package imagex

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

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
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) upload(host string, storeInfo StoreInfo, imageBytes []byte) error {
	if len(imageBytes) == 0 {
		return fmt.Errorf("file size is zero")
	}

	checkSum := fmt.Sprintf("%x", crc32.ChecksumIEEE(imageBytes))
	url := fmt.Sprintf("https://%s/%s", host, storeInfo.StoreUri)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(imageBytes))
	if err != nil {
		return fmt.Errorf("fail to new put request, %v", err)
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", storeInfo.Auth)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("fail to do request to %s, %v", url, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("fail to read response body %s, %v", url, err)
	}

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status=%v, body=%s, url=%s", rsp.StatusCode, string(body), url)
	}
	defer rsp.Body.Close()

	var putResp struct {
		Success int         `json:"success"`
		Payload interface{} `json:"payload"`
	}
	if err = json.Unmarshal(body, &putResp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	if putResp.Success != 0 {
		return fmt.Errorf("put to host %s err:%+v", url, putResp)
	}
	return nil
}

// 上传图片
func (c *ImageX) UploadImages(params *ApplyUploadImageParam, images [][]byte) (*CommitUploadImageResult, error) {
	if params.UploadNum == 0 {
		params.UploadNum = 1
	}
	if len(images) != params.UploadNum {
		return nil, fmt.Errorf("UploadImages: images num %d != upload num %d", len(images), params.UploadNum)
	}

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
	for i, image := range images {
		err := c.upload(uploadAddr.UploadHosts[0], uploadAddr.StoreInfos[i], image)
		if err != nil {
			return nil, fmt.Errorf("UploadImages: fail to do upload, %v", err)
		}
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:  params.ServiceId,
		SessionKey: uploadAddr.SessionKey,
	}
	if params.CommitParam != nil {
		commitParams.SkipMeta = params.CommitParam.SkipMeta
		commitParams.OptionInfos = params.CommitParam.OptionInfos
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

// 获取上传临时密钥
func (c *ImageX) GetUploadAuth(serviceIds []string) (*base.SecurityToken2, error) {
	return c.GetUploadAuthWithExpire(serviceIds, time.Hour)
}

func (c *ImageX) GetUploadAuthWithExpire(serviceIds []string, expire time.Duration) (*base.SecurityToken2, error) {
	inlinePolicy := new(base.Policy)
	actions := []string{
		"ImageX:ApplyImageUpload",
		"ImageX:CommitImageUpload",
	}

	resources := make([]string, 0)
	if len(serviceIds) == 0 {
		resources = append(resources, fmt.Sprintf(ResourceServiceIdTRN, "*"))
	} else {
		for _, sid := range serviceIds {
			resources = append(resources, fmt.Sprintf(ResourceServiceIdTRN, sid))
		}
	}

	statement := base.NewAllowStatement(actions, resources)
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)
	return c.Client.SignSts2(inlinePolicy, expire)
}

func (c *ImageX) updateImageUrls(serviceId string, req *UpdateImageUrlPayload) ([]string, error) {
	query := url.Values{}
	query.Add("ServiceId", serviceId)

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("updateImageUrls: fail to marshal request, %v", err)
	}

	respBody, _, err := c.Client.Json("UpdateImageUploadFiles", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("updateImageUrls: fail to do request, %v", err)
	}

	result := new(UpdateImageUrlPayload)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result.ImageUrls, nil
}

func (c *ImageX) RefreshImageUrls(serviceId string, urls []string) ([]string, error) {
	req := &UpdateImageUrlPayload{
		Action:    ActionRefresh,
		ImageUrls: urls,
	}
	return c.updateImageUrls(serviceId, req)
}

func (c *ImageX) EnableImageUrls(serviceId string, urls []string) ([]string, error) {
	req := &UpdateImageUrlPayload{
		Action:    ActionEnable,
		ImageUrls: urls,
	}
	return c.updateImageUrls(serviceId, req)
}

func (c *ImageX) DisableImageUrls(serviceId string, urls []string) ([]string, error) {
	req := &UpdateImageUrlPayload{
		Action:    ActionDisable,
		ImageUrls: urls,
	}
	return c.updateImageUrls(serviceId, req)
}
