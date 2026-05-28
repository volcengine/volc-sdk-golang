package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (c *LsClient) CreateLogBackFlowTask(request *CreateLogBackFlowTaskRequest) (r *CreateLogBackFlowTaskResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	reqHeaders := map[string]string{"Content-Type": "application/json"}
	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	rawResponse, err := c.Request(http.MethodPost, PathCreateLogBackFlowTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, NewClientError(errors.New("nil http response"))
	}
	defer rawResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}
	response := &CreateLogBackFlowTaskResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DeleteLogBackFlowTask(request *DeleteLogBackFlowTaskRequest) (r *DeleteLogBackFlowTaskResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	reqHeaders := map[string]string{"Content-Type": "application/json"}
	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	rawResponse, err := c.Request(http.MethodDelete, PathDeleteLogBackFlowTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, NewClientError(errors.New("nil http response"))
	}
	defer rawResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}
	response := &DeleteLogBackFlowTaskResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DescribeLogBackFlowTasks(request *DescribeLogBackFlowTasksRequest) (r *DescribeLogBackFlowTasksResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	reqHeaders := map[string]string{"Content-Type": "application/json"}
	params := map[string]string{}
	if request.PageNumber != nil {
		params["PageNumber"] = strconv.Itoa(*request.PageNumber)
	}
	if request.PageSize != nil {
		params["PageSize"] = strconv.Itoa(*request.PageSize)
	}
	if request.TaskID != nil {
		params["TaskId"] = *request.TaskID
	}
	if request.TaskName != nil {
		params["TaskName"] = *request.TaskName
	}
	if request.Status != nil {
		params["Status"] = strconv.Itoa(*request.Status)
	}
	if request.ScheduleSQLTaskID != nil {
		params["ScheduleSQLTaskId"] = *request.ScheduleSQLTaskID
	}
	if request.ShipperID != nil {
		params["ShipperId"] = *request.ShipperID
	}
	rawResponse, err := c.Request(http.MethodGet, buildDescribeLogBackFlowTasksURI(params, request.TopicIDList), nil, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, NewClientError(errors.New("nil http response"))
	}
	defer rawResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}
	response := &DescribeLogBackFlowTasksResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	return response, nil
}

func buildDescribeLogBackFlowTasksURI(params map[string]string, topicIDs []string) string {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}
	for _, topicID := range topicIDs {
		if topicID != "" {
			query.Add("TopicIDList", topicID)
		}
	}
	if len(query) == 0 {
		return PathDescribeLogBackFlowTasks
	}
	return PathDescribeLogBackFlowTasks + "?" + query.Encode()
}

func (c *LsClient) ModifyLogBackFlowTask(request *ModifyLogBackFlowTaskRequest) (r *ModifyLogBackFlowTaskResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	reqHeaders := map[string]string{"Content-Type": "application/json"}
	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	rawResponse, err := c.Request(http.MethodPut, PathModifyLogBackFlowTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, NewClientError(errors.New("nil http response"))
	}
	defer rawResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}
	response := &ModifyLogBackFlowTaskResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	return response, nil
}
