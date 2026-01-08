package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateTraceInstance(request *CreateTraceInstanceRequest) (*CreateTraceInstanceResponse, error) {
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

	rawResponse, err := c.Request(http.MethodPost, PathCreateTraceInstance, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &CreateTraceInstanceResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeTraceInstances(request *DescribeTraceInstancesRequest) (*DescribeTraceInstancesResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := make(map[string]string)
	if request.PageNumber != nil {
		params["PageNumber"] = strconv.Itoa(*request.PageNumber)
	}
	if request.PageSize != nil {
		params["PageSize"] = strconv.Itoa(*request.PageSize)
	}
	if request.TraceInstanceName != nil {
		params["TraceInstanceName"] = *request.TraceInstanceName
	}
	if request.TraceInstanceID != nil {
		params["TraceInstanceId"] = *request.TraceInstanceID
	}
	if request.ProjectID != nil {
		params["ProjectId"] = *request.ProjectID
	}
	if request.ProjectName != nil {
		params["ProjectName"] = *request.ProjectName
	}
	if request.Status != nil {
		params["Status"] = *request.Status
	}
	if request.IamProjectName != nil {
		params["IamProjectName"] = *request.IamProjectName
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeTraceInstances, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, errors.New("request failed: empty response")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeTraceInstancesResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) SearchTraces(request *SearchTracesRequest) (*SearchTracesResponse, error) {
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

	rawResponse, err := c.Request(http.MethodPost, PathSearchTraces, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &SearchTracesResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeTrace(request *DescribeTraceRequest) (*DescribeTraceResponse, error) {
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

	rawResponse, err := c.Request(http.MethodPost, PathDescribeTrace, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var resp = &DescribeTraceResponse{}
	resp.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
