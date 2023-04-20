package tls

import (
	"encoding/json"
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
