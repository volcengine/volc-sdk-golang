package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type VideoOverResolutionSubmitTaskV2Request struct {
	ReqKey string `json:"req_key"`
	Vid    string `json:"vid"`
	Url    string `json:"url"`
}

type VideoOverResolutionQueryTaskV2Request struct {
	ReqKey string `json:"req_key"`
	TaskId string `json:"task_id"`
}

// 响应参数

type VideoOverResolutionSubmitTaskV2Data struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	TaskId           string   `json:"task_id"`
}

type VideoOverResolutionSubmitTaskV2Result struct {
	ResponseMetadata *base.ResponseMetadata               `json:",omitempty"`
	RequestId        string                               `json:"request_id"`
	Code             int                                  `json:"code"`
	Message          string                               `json:"message"`
	Status           int                                  `json:"status"`
	TimeElapsed      string                               `json:"time_elapsed"`
	Data             *VideoOverResolutionSubmitTaskV2Data `json:"data"`
}

type VideoOverResolutionQueryTaskV2Data struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	RespData         string   `json:"resp_data"`
	Status           string   `json:"status"`
}

type VideoOverResolutionQueryTaskV2Result struct {
	ResponseMetadata *base.ResponseMetadata              `json:",omitempty"`
	RequestId        string                              `json:"request_id"`
	Code             int                                 `json:"code"`
	Message          string                              `json:"message"`
	Status           int                                 `json:"status"`
	TimeElapsed      string                              `json:"time_elapsed"`
	Data             *VideoOverResolutionQueryTaskV2Data `json:"data"`
}
