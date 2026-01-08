package tls

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateScheduleSqlTask 创建定时 SQL 任务
func (c *LsClient) CreateScheduleSqlTask(request *CreateScheduleSqlTaskRequest) (*CreateScheduleSqlTaskResponse, error) {
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

	rawResponse, err := c.Request(http.MethodPost, PathCreateScheduleSqlTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, errors.New("raw response is nil")
	}
	if rawResponse.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}
	defer rawResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CreateScheduleSqlTaskResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteScheduleSqlTask 删除定时 SQL 任务
func (c *LsClient) DeleteScheduleSqlTask(request *DeleteScheduleSqlTaskRequest) (*DeleteScheduleSqlTaskResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

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

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteScheduleSqlTask, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, errors.New("raw response is nil")
	}
	if rawResponse.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}
	defer rawResponse.Body.Close()

	if rawResponse.StatusCode != http.StatusOK {
		responseBody, err := ioutil.ReadAll(rawResponse.Body)
		if err != nil {
			return nil, err
		}
		if responseBody != nil {
			tlsErr := &Error{}
			if json.Unmarshal(responseBody, tlsErr) == nil && (len(tlsErr.Code) != 0 || len(tlsErr.Message) != 0) {
				tlsErr.HTTPCode = int32(rawResponse.StatusCode)
				tlsErr.RequestID = rawResponse.Header.Get(RequestIDHeader)
				return nil, tlsErr
			}
		}
		return nil, NewBadResponseError(string(responseBody), rawResponse.Header, rawResponse.StatusCode)
	}

	response := &DeleteScheduleSqlTaskResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}

// DescribeScheduleSqlTask 查询定时 SQL 任务详情
func (c *LsClient) DescribeScheduleSqlTask(request *DescribeScheduleSqlTaskRequest) (*DescribeScheduleSqlTaskResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

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

	rawResponse, err := c.Request(http.MethodGet, PathDescribeScheduleSqlTask, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, errors.New("raw response is nil")
	}
	if rawResponse.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	response := &DescribeScheduleSqlTaskResponse{}
	response.FillRequestId(rawResponse)
	if err = json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
