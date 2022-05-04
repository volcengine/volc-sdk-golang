package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateTopic(request *CreateTopicRequest) (*CreateTopicResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"ProjectId":   request.ProjectID,
		"TopicName":   request.TopicName,
		"Ttl":         request.Ttl,
		"Description": request.Description,
	}

	if request.ShardCount != 0 {
		reqBody["ShardCount"] = request.ShardCount
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateTopic, nil, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CreateTopicResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteTopic(request *DeleteTopicRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"TopicId": request.TopicID,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteTopic, nil, reqHeaders, bytesBody)
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

// ModifyTopic: 更新日志主题信息

func (c *LsClient) ModifyTopic(request *ModifyTopicRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"TopicId": request.TopicID,
	}

	if request.TopicName != nil {
		reqBody["TopicName"] = *request.TopicName
	}

	if request.Ttl != nil {
		reqBody["Ttl"] = *request.Ttl
	}

	if request.Description != nil {
		reqBody["Description"] = *request.Description
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyTopic, nil, reqHeaders, bytesBody)
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

// DescribeTopic 获取一个日志项目的信息

func (c *LsClient) DescribeTopic(request *DescribeTopicRequest) (*DescribeTopicResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TopicId": request.TopicID,
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeTopic, params, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeTopicResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

// DescribeTopics: 获取日志主题列表

func (c *LsClient) DescribeTopics(request *DescribeTopicsRequest) (*DescribeTopicsResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"ProjectId": request.ProjectID,
	}

	if len(request.TopicName) != 0 {
		params["TopicName"] = request.TopicName
	}

	if len(request.TopicID) != 0 {
		params["TopicID"] = request.TopicID
	}

	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeTopics, params, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeTopicsResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
