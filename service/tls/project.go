package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateProject(request *CreateProjectRequest) (r *CreateProjectResponse, e error) {
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

	rawResponse, err := c.Request(http.MethodPost, PathCreateProject, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &CreateProjectResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteProject(request *DeleteProjectRequest) (r *CommonResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"ProjectId": request.ProjectID,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteProject, nil, c.assembleHeader(request.CommonRequest, reqHeaders), jsonBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, errors.New("request failed: empty response")
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

func (c *LsClient) DeleteTraceInstance(request *DeleteTraceInstanceRequest) (r *DeleteTraceInstanceResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteTraceInstance, nil, c.assembleHeader(request.CommonRequest, reqHeaders), jsonBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, NewClientError(errors.New("nil http response"))
	}
	defer rawResponse.Body.Close()

	if rawResponse.StatusCode != http.StatusOK {
		errorMessage, err := ioutil.ReadAll(rawResponse.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorMessage))
	}

	response := &DeleteTraceInstanceResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeProject(request *DescribeProjectRequest) (r *DescribeProjectResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"ProjectId": request.ProjectID,
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeProject, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &DescribeProjectResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeProjects(request *DescribeProjectsRequest) (r *DescribeProjectsResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := make(map[string]string)
	if len(request.ProjectName) != 0 {
		params["ProjectName"] = request.ProjectName
	}

	if len(request.ProjectID) != 0 {
		params["ProjectId"] = request.ProjectID
	}

	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	if request.IsFullName {
		params["IsFullName"] = strconv.FormatBool(request.IsFullName)
	}

	if request.IamProjectName != nil {
		params["IamProjectName"] = *request.IamProjectName
	}

	if request.Tags != nil && len(request.Tags) > 0 {
		tagsBytes, err := json.Marshal(request.Tags)
		if err != nil {
			return nil, err
		}
		params["Tags"] = string(tagsBytes)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeProjects, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &DescribeProjectsResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) ModifyProject(request *ModifyProjectRequest) (r *CommonResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"ProjectId": request.ProjectID,
	}

	if request.ProjectName != nil {
		reqBody["ProjectName"] = *request.ProjectName
	}

	if request.Description != nil {
		reqBody["Description"] = *request.Description
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyProject, nil, c.assembleHeader(request.CommonRequest, reqHeaders), jsonBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, errors.New("request failed: empty response")
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

func (c *LsClient) GetAccountStatus(request *GetAccountStatusRequest) (r *GetAccountStatusResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathGetAccountStatus, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &GetAccountStatusResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
