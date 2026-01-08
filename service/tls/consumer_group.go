package tls

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateConsumerGroup(request *CreateConsumerGroupRequest) (*CreateConsumerGroupResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"ProjectID":         request.ProjectID,
		"TopicIDList":       request.TopicIDList,
		"ConsumerGroupName": request.ConsumerGroupName,
		"HeartbeatTTL":      request.HeartbeatTTL,
		"OrderedConsume":    request.OrderedConsume,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateConsumerGroup, nil, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, fmt.Errorf("raw response is nil")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CreateConsumerGroupResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"ProjectID":         request.ProjectID,
		"ConsumerGroupName": request.ConsumerGroupName,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteConsumerGroup, nil, reqHeaders, jsonBody)
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

func (c *LsClient) DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (*DescribeConsumerGroupsResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := make(map[string]string)
	if len(request.ProjectID) != 0 {
		params["ProjectId"] = request.ProjectID
	}

	if len(request.ProjectName) != 0 {
		params["ProjectName"] = request.ProjectName
	}

	if len(request.ConsumerGroupName) != 0 {
		params["ConsumerGroupName"] = request.ConsumerGroupName
	}

	if len(request.TopicID) != 0 {
		params["TopicId"] = request.TopicID
	}

	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	if request.TopicName != nil && len(*request.TopicName) != 0 {
		params["TopicName"] = *request.TopicName
	}

	if request.IamProjectName != nil && len(*request.IamProjectName) != 0 {
		params["IamProjectName"] = *request.IamProjectName
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeConsumerGroups, params, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeConsumerGroupsResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) ModifyConsumerGroup(request *ModifyConsumerGroupRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"ProjectID":         request.ProjectID,
		"ConsumerGroupName": request.ConsumerGroupName,
	}

	if request.TopicIDList != nil {
		reqBody["TopicIDList"] = *request.TopicIDList
	}

	if request.HeartbeatTTL != nil {
		reqBody["HeartbeatTTL"] = *request.HeartbeatTTL
	}

	if request.OrderedConsume != nil {
		reqBody["OrderedConsume"] = request.OrderedConsume
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyConsumerGroup, nil, reqHeaders, jsonBody)
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
