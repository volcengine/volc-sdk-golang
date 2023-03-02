package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type AllAgeGenerationRequest struct {
	ReqKey           string   `json:"req_key"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	TargetAge        int      `json:"target_age"`
}

// 响应参数

type AllAgeGenerationData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
}

type AllAgeGenerationResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *AllAgeGenerationData  `json:"data"`
}
