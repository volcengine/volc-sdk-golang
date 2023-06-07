package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type LensVidaVideoSubmitTaskV2Request struct {
	ReqKey   string `json:"req_key"`
	Url      string `json:"url"`
	VidaMode string `json:"vida_mode"`
}

// 响应参数

type LensVidaVideoSubmitTaskV2Data struct {
	TaskId string `json:"task_id"`
}

type LensVidaVideoSubmitTaskV2Result struct {
	ResponseMetadata *base.ResponseMetadata         `json:",omitempty"`
	RequestId        string                         `json:"request_id"`
	Code             int                            `json:"code"`
	Message          string                         `json:"message"`
	Status           int                            `json:"status"`
	TimeElapsed      string                         `json:"time_elapsed"`
	Data             *LensVidaVideoSubmitTaskV2Data `json:"data"`
}
