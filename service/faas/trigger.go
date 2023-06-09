package faas

import (
	"encoding/json"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

type ListTriggersRequest struct {
	FunctionId string `json:"FunctionId"`
}

type Triggers struct {
	Items []*Trigger `json:"Items"`
	Total int32      `json:"Total"`
}

type ListTriggersResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *Triggers              `json:"Result,omitempty"`
}

type Trigger struct {
	Id             string  `json:"Id"`
	Type           string  `json:"Type"`
	Name           string  `json:"Name"`
	Description    string  `json:"Description"`
	CreationTime   string  `json:"CreationTime"`
	LastUpdateTime string  `json:"LastUpdateTime"`
	DetailedConfig string  `json:"DetailedConfig"`
	Enabled        *bool   `json:"Enabled,omitempty"`
	ImageVersion   *string `json:"ImageVersion,omitempty"`
	FunctionID     *string `json:"FunctionID,omitempty"`
}

func (f *FaaS) ListTriggers(req *ListTriggersRequest) (*ListTriggersResponse, int, error) {
	query := url.Values{}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, -1, err
	}

	var resp ListTriggersResponse
	respBody, status, err := f.Json("ListTriggers", query, string(body))
	if err != nil {
		return nil, status, err
	}
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, status, err
	}
	return &resp, status, nil
}
