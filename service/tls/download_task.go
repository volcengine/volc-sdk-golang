package tls

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateDownloadTask(request *CreateDownloadTaskRequest) (r *CreateDownloadTaskResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateDownloadTask, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &CreateDownloadTaskResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeDownloadTasks(request *DescribeDownloadTasksRequest) (r *DescribeDownloadTasksResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TopicId": request.TopicID,
	}
	if request.PageNumber != nil {
		params["PageNumber"] = strconv.Itoa(*request.PageNumber)
	}
	if request.PageSize != nil {
		params["PageSize"] = strconv.Itoa(*request.PageSize)
	}
	if request.TaskName != nil {
		params["TaskName"] = *request.TaskName
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeDownloadTasks, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeDownloadTasksResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeDownloadUrl(request *DescribeDownloadUrlRequest) (r *DescribeDownloadUrlResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TaskId": request.TaskId,
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeDownloadUrl, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeDownloadUrlResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) CancelDownloadTask(request *CancelDownloadTaskRequest) (r *CancelDownloadTaskResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCancelDownloadTask, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &CancelDownloadTaskResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}
