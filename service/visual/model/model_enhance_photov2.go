package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type EnhancePhotoV2Request struct {
	ReqKey             string   `json:"req_key"`
	BinaryDataBase64   []string `json:"binary_data_base64"`
	ResolutionBoundary string   `json:"resolution_boundary"`
	EnableHdr          bool     `json:"enable_hdr"`
	EnableWb           bool     `json:"enable_wb"`
	ResultFormat       int64    `json:"result_format"`
	JpgQuality         int64    `json:"jpg_quality"`
}

// 响应参数

type EnhancePhotoV2Data struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
}

type EnhancePhotoV2Result struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *EnhancePhotoV2Data    `json:"data"`
}
