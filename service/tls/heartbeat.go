package tls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *LsClient) ConsumerHeartbeat(request *ConsumerHeartbeatRequest) (*ConsumerHeartbeatResponse, error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"ProjectID":         request.ProjectID,
		"ConsumerGroupName": request.ConsumerGroupName,
		"ConsumerName":      request.ConsumerName,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathConsumerHeartbeat, nil, reqHeaders, bytesBody)
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

	var response = &ConsumerHeartbeatResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
