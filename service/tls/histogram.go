package tls

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *LsClient) DescribeHistogram(request *DescribeHistogramRequest) (r *DescribeHistogramResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathDescribeHistogram, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeHistogramResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}
