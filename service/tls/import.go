package tls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateImportTask(request *CreateImportTaskRequest) (r *CreateImportTaskResponse, err error) {
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

	rawResponse, err := c.Request(http.MethodPost, PathCreateImportTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	response := &CreateImportTaskResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteImportTask(request *DeleteImportTaskRequest) (r *DeleteImportTaskResponse, err error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteImportTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	response := &DeleteImportTaskResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) ModifyImportTask(request *ModifyImportTaskRequest) (r *ModifyImportTaskResponse, err error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyImportTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	response := &ModifyImportTaskResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeImportTask(request *DescribeImportTaskRequest) (r *DescribeImportTaskResponse, err error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TaskId": request.TaskID,
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeImportTask, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeImportTaskResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeImportTasks(request *DescribeImportTasksRequest) (r *DescribeImportTasksResponse, err error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}
	if request.TaskID != "" {
		params["TaskId"] = request.TaskID
	}
	if request.TaskName != "" {
		params["TaskName"] = request.TaskName
	}
	if request.ProjectID != "" {
		params["ProjectId"] = request.ProjectID
	}
	if request.ProjectName != "" {
		params["ProjectName"] = request.ProjectName
	}
	if request.IamProjectName != "" {
		params["IamProjectName"] = request.IamProjectName
	}
	if request.TopicID != "" {
		params["TopicId"] = request.TopicID
	}
	if request.TopicName != "" {
		params["TopicName"] = request.TopicName
	}
	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}
	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}
	if request.SourceType != "" {
		params["SourceType"] = request.SourceType
	}
	if request.Status != "" {
		params["Status"] = request.Status
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeImportTasks, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeImportTasksResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
