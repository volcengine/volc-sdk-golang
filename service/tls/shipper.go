package tls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) CreateShipper(request *CreateShipperRequest) (*CreateShipperResponse, error) {
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

	rawResponse, err := c.Request(http.MethodPost, PathCreateShipper, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	response := &CreateShipperResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DeleteShipper(request *DeleteShipperRequest) (*DeleteShipperResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteShipper, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	response := &DeleteShipperResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) ModifyShipper(request *ModifyShipperRequest) (*ModifyShipperResponse, error) {
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

	rawResponse, err := c.Request(http.MethodPut, PathModifyShipper, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	response := &ModifyShipperResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeShipper(request *DescribeShipperRequest) (*DescribeShipperResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"ShipperId": request.ShipperId,
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeShipper, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeShipperResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeShippers(request *DescribeShippersRequest) (*DescribeShippersResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	if request.IamProjectName != "" {
		params["IamProjectName"] = request.IamProjectName
	}
	if request.ProjectId != "" {
		params["ProjectId"] = request.ProjectId
	}
	if request.ProjectName != "" {
		params["ProjectName"] = request.ProjectName
	}
	if request.TopicId != "" {
		params["TopicId"] = request.TopicId
	}
	if request.TopicName != "" {
		params["TopicName"] = request.TopicName
	}
	if request.ShipperId != "" {
		params["ShipperId"] = request.ShipperId
	}
	if request.ShipperName != "" {
		params["ShipperName"] = request.ShipperName
	}
	if request.ShipperType != "" {
		params["ShipperType"] = request.ShipperType
	}
	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}
	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeShippers, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeShippersResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
