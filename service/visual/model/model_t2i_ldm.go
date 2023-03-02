package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type T2ILDMRequest struct {
	ReqKey    string `json:"req_key"`
	Text      string `json:"text"`
	StyleTerm string `json:"style_term"`
}

// 响应参数

type T2ILDMData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
}

type T2ILDMResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *T2ILDMData            `json:"data"`
}
