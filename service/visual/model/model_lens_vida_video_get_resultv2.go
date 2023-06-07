package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type LensVidaVideoGetResultV2Request struct {
	ReqKey string `json:"req_key"`
	TaskId string `json:"task_id"`
}

// 响应参数

type LensVidaVideoGetResultV2Data struct {
	BinaryDataBase64 []string `json:"binary_data_base64"`
	RespData         string   `json:"resp_data"`
	Status           string   `json:"status"`
}

type LensVidaVideoGetResultV2Result struct {
	ResponseMetadata *base.ResponseMetadata        `json:",omitempty"`
	RequestId        string                        `json:"request_id"`
	Code             int                           `json:"code"`
	Message          string                        `json:"message"`
	Status           int                           `json:"status"`
	TimeElapsed      string                        `json:"time_elapsed"`
	Data             *LensVidaVideoGetResultV2Data `json:"data"`
}
