package tls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *LsClient) OpenKafkaConsumer(request *OpenKafkaConsumerRequest) (r *OpenKafkaConsumerResponse, e error) {
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

	rawResponse, err := c.Request(http.MethodPut, PathOpenKafkaConsumer, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &OpenKafkaConsumerResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) CloseKafkaConsumer(request *CloseKafkaConsumerRequest) (r *CloseKafkaConsumerResponse, e error) {
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

	rawResponse, err := c.Request(http.MethodPut, PathCloseKafkaConsumer, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CloseKafkaConsumerResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeKafkaConsumer(request *DescribeKafkaConsumerRequest) (r *DescribeKafkaConsumerResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TopicId": request.TopicID,
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeKafkaConsumer, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeKafkaConsumerResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}
