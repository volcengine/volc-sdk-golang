package model

import "github.com/volcengine/volc-sdk-golang/base"

// VisualBaseRespDataV2 is the public response data of visual new ability, it might be updates
type VisualBaseRespDataV2 struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp,omitempty"`
	BinaryDataBase64 []string `json:"binary_data_base64,omitempty"`
	TaskId           string   `json:"task_id,omitempty"`
	Status           string   `json:"status,omitempty"`
	RespData         string   `json:"resp_data,omitempty"`
}

// VisualPubResult is the public response of visual new ability
type VisualPubResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`

	Code        int                   `json:"code"`
	Data        *VisualBaseRespDataV2 `json:"data,omitempty"`
	Message     string                `json:"message"`
	RequestId   string                `json:"request_id"`
	Status      int                   `json:"status"`
	TimeElapsed string                `json:"time_elapsed"`
}
