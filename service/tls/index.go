package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *LsClient) CreateIndex(request *CreateIndexRequest) (r *CreateIndexResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateIndex, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CreateIndexResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteIndex(request *DeleteIndexRequest) (r *CommonResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"TopicId": request.TopicID,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteIndex, nil, c.assembleHeader(request.CommonRequest, reqHeaders), jsonBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	if rawResponse.StatusCode != http.StatusOK {
		errorMessage, err := ioutil.ReadAll(rawResponse.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorMessage))
	}

	response := &CommonResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeIndex(request *DescribeIndexRequest) (r *DescribeIndexResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TopicId": request.TopicID,
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeIndex, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeIndexResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	if response.FullText != nil && len(response.FullText.Delimiter) > 0 {
		response.FullText.Delimiter = ReplaceWhiteSpaceCharacter(response.FullText.Delimiter)
	}
	if len(*response.KeyValue) > 0 {
		for k, v := range *response.KeyValue {
			v.Value.Delimiter = ReplaceWhiteSpaceCharacter(v.Value.Delimiter)
			(*response.KeyValue)[k] = v
		}
	}
	if len(*response.UserInnerKeyValue) > 0 {
		for k, v := range *response.UserInnerKeyValue {
			v.Value.Delimiter = ReplaceWhiteSpaceCharacter(v.Value.Delimiter)
			(*response.UserInnerKeyValue)[k] = v
		}
	}

	return response, nil
}

// ModifyIndex: 更新索引。
// 由于该接口为全量更新接口，等同于重新创建一个新的索引，因此要注意不要漏填字段

func (c *LsClient) ModifyIndex(request *ModifyIndexRequest) (r *CommonResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyIndex, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	if rawResponse.StatusCode != http.StatusOK {
		errorMessage, err := ioutil.ReadAll(rawResponse.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorMessage))
	}

	response := &CommonResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) SearchLogs(request *SearchLogsRequest) (r *SearchLogsResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Encoding": CompressGz,
		HeaderAPIVersion:  APIVersion2,
	}

	return c.search(request, reqHeaders)
}

func (c *LsClient) search(request *SearchLogsRequest, reqHeaders map[string]string) (*SearchLogsResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqBody := map[string]interface{}{
		"TopicId":   request.TopicID,
		"Query":     request.Query,
		"StartTime": request.StartTime,
		"EndTime":   request.EndTime,
		"Limit":     request.Limit,
		"HighLight": request.HighLight,
		"Context":   request.Context,
		"Sort":      request.Sort,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathSearchLogs, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &SearchLogsResponse{}
	response.FillRequestId(rawResponse)
	decoder := json.NewDecoder(strings.NewReader(string(responseBody)))
	decoder.UseNumber()

	if err := decoder.Decode(response); err != nil {
		return nil, err
	}
	return response, nil
}

// SearchLogsV2 搜索按照0.3.0api版本进行，和默认的0.2.0版本区别见文档https://www.volcengine.com/docs/6470/112170
func (c *LsClient) SearchLogsV2(request *SearchLogsRequest) (*SearchLogsResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Encoding": CompressGz,
		HeaderAPIVersion:  APIVersion3,
	}

	return c.search(request, reqHeaders)
}
