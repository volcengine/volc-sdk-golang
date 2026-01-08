package tls

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) DescribeShards(request *DescribeShardsRequest) (r *DescribeShardsResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"TopicId": request.TopicID,
	}

	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeShards, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &DescribeShardsResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) ManualShardSplit(ctx context.Context, request *ManualShardSplitRequest) (*ManualShardSplitResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, err
	}

	headers := c.assembleHeader(request.CommonRequest, map[string]string{
		"Content-Type": "application/json",
	})

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(http.MethodPost, PathManualSplitShard, nil, headers, body)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, NewClientError(errors.New("nil http response"))
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ManualShardSplitResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	response.FillRequestId(resp)
	return &response, nil
}
