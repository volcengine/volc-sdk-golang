package tls

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateETLTask(request *CreateETLTaskRequest) (*CreateETLTaskResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateETLTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &CreateETLTaskResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteETLTask(request *DeleteETLTaskRequest) (*CommonResponse, error) {

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"TaskId": request.TaskID,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteETLTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

func (c *LsClient) ModifyETLTask(request *ModifyETLTaskRequest) (*CommonResponse, error) {

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyETLTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

func (c *LsClient) DescribeETLTask(request *DescribeETLTaskRequest) (*DescribeETLTaskResponse, error) {

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TaskId": request.TaskID,
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeETLTask, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeETLTaskResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeETLTasks(request *DescribeETLTasksRequest) (*DescribeETLTasksResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	if len(request.ProjectID) != 0 {
		params["ProjectId"] = request.ProjectID
	}

	if len(request.ProjectName) != 0 {
		params["ProjectName"] = request.ProjectName
	}

	if len(request.IamProjectName) != 0 {
		params["IamProjectName"] = request.IamProjectName
	}

	if len(request.SourceTopicID) != 0 {
		params["SourceTopicId"] = request.SourceTopicID
	}

	if len(request.SourceTopicName) != 0 {
		params["SourceTopicName"] = request.SourceTopicName
	}

	if len(request.TaskID) != 0 {
		params["TaskId"] = request.TaskID
	}

	if len(request.TaskName) != 0 {
		params["TaskName"] = request.TaskName
	}

	if len(request.Status) != 0 {
		params["Status"] = request.Status
	}

	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeETLTasks, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeETLTasksResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) ModifyETLTaskStatus(request *ModifyETLTaskStatusRequest) (*CommonResponse, error) {

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyETLTaskStatus, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

func (c *LsClient) ModifyScheduleSqlTask(request *ModifyScheduleSqlTaskRequest) (*ModifyScheduleSqlTaskResponse, error) {
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

	rawResponse, err := c.Request(http.MethodPut, PathModifyScheduleSqlTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, fmt.Errorf("raw response is nil")
	}
	if rawResponse.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	response := &ModifyScheduleSqlTaskResponse{}
	response.FillRequestId(rawResponse)

	if err = json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
