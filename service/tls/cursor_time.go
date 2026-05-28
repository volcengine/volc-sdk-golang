package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DescribeCursorTimeRequest struct {
	CommonRequest
	TopicID string `json:"TopicId,omitempty"`
	ShardID *int   `json:"ShardId,omitempty"`
	Cursor  string `json:"Cursor"`
}

func (v *DescribeCursorTimeRequest) CheckValidation() error {
	if v.TopicID == "" {
		return errors.New("Invalid argument, empty TopicID")
	}
	if v.ShardID == nil {
		return errors.New("Invalid argument, empty ShardID")
	}
	if v.Cursor == "" {
		return errors.New("Invalid argument, empty Cursor")
	}
	return nil
}

type DescribeCursorTimeResponse struct {
	CommonResponse
	CursorTime int64 `json:"CursorTime"`
}

func (c *LsClient) DescribeCursorTime(request *DescribeCursorTimeRequest) (r *DescribeCursorTimeResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	reqHeaders := map[string]string{"Content-Type": "application/json"}
	params := map[string]string{"Cursor": request.Cursor}
	if request.TopicID != "" {
		params["TopicId"] = request.TopicID
	}
	if request.ShardID != nil {
		params["ShardId"] = strconv.Itoa(*request.ShardID)
	}
	rawResponse, err := c.Request(http.MethodGet, PathDescribeCursorTime, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
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
	response := &DescribeCursorTimeResponse{}
	response.FillRequestId(rawResponse)
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	return response, nil
}
