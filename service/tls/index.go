package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (c *LsClient) CreateIndex(request *CreateIndexRequest) (*CreateIndexResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"topic_id":        request.TopicID,
		"fulltext_index":  request.FulltextIndex,
		"cas_sensitive":   request.CasSensitive,
		"include_chinese": request.IncludeChinese,
		"delimiter":       request.Delimiter,
		"key_value_index": request.KeyValueIndex,
		"key_value_list":  request.KeyValueList,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateIndex, nil, reqHeaders, bytesBody)
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

func (c *LsClient) DeleteIndex(request *DeleteIndexRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"topic_id": request.TopicID,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteIndex, nil, reqHeaders, jsonBody)
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

func (c *LsClient) DescribeIndex(request *DescribeIndexRequest) (*DescribeIndexResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"topic_id": request.TopicID,
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeIndex, params, reqHeaders, bytesBody)
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

	return response, nil
}

// ModifyIndex: 更新索引。
// 由于该接口为全量更新接口，等同于重新创建一个新的索引，因此要注意不要漏填字段

func (c *LsClient) ModifyIndex(request *ModifyIndexRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"topic_id":        request.TopicID,
		"fulltext_index":  request.FulltextIndex,
		"cas_sensitive":   request.CasSensitive,
		"include_chinese": request.IncludeChinese,
		"delimiter":       request.Delimiter,
		"key_value_index": request.KeyValueIndex,
		"key_value_list":  request.KeyValueList,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyIndex, nil, reqHeaders, bytesBody)
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

func (c *LsClient) SearchLogs(request *SearchLogsRequest) (*SearchLogsResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"topic_id":   request.TopicID,
		"query":      request.Query,
		"start_time": request.StartTime,
		"end_time":   request.EndTime,
		"limit":      request.Limit,
		"high_light": request.HighLight,
		"context":    request.Context,
		"sort":       request.Sort,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathSearchLogs, nil, reqHeaders, bytesBody)
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

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
