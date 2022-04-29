package tls

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) DescribeShards(request *DescribeShardsRequest) (*DescribeShardsResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{
		"topic_id": request.TopicID,
	}

	params["offset"] = strconv.Itoa(request.Offset)

	if request.Limit != 0 {
		params["limit"] = strconv.Itoa(request.Limit)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeShards, params, reqHeaders, bytesBody)
	if err != nil {
		return nil, err
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
