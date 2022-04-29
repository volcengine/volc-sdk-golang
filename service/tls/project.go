package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateProject(request *CreateProjectRequest) (*CreateProjectResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"project_name": request.ProjectName,
		"description":  request.Description,
		"region":       request.Region,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateProject, nil, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
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

func (c *LsClient) DeleteProject(request *DeleteProjectRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"project_id": request.ProjectID,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteProject, nil, reqHeaders, jsonBody)
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

func (c *LsClient) DescribeProject(request *DescribeProjectRequest) (*DescribeProjectResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"project_id": request.ProjectID,
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeProject, params, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
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

func (c *LsClient) DescribeProjects(request *DescribeProjectsRequest) (*DescribeProjectsResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := make(map[string]string)
	if len(request.ProjectName) != 0 {
		params["project_name"] = request.ProjectName
	}

	if len(request.ProjectID) != 0 {
		params["project_id"] = request.ProjectID
	}

	params["offset"] = strconv.Itoa(request.Offset)

	if request.Limit != 0 {
		params["limit"] = strconv.Itoa(request.Limit)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeProjects, params, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
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

func (c *LsClient) ModifyProject(request *ModifyProjectRequest) (*CommonResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"project_id": request.ProjectID,
	}

	if request.ProjectName != nil {
		reqBody["project_name"] = *request.ProjectName
	}

	if request.Description != nil {
		reqBody["description"] = *request.Description
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyProject, nil, reqHeaders, jsonBody)
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
