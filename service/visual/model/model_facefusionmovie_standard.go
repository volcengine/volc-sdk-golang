package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type FaceFusionMovieSubmitTaskRequest struct {
	ReqKey             string `json:"req_key"`
	ImageUrl           string `json:"image_url"`
	VideoUrl           string `json:"video_url"`
	EnableFaceBeautify bool   `json:"enable_face_beautify"`
}

type FaceFusionMovieGetResultRequest struct {
	ReqKey string `json:"req_key"`
	TaskId string `json:"task_id"`
}

// 响应参数

type FaceFusionMovieSubmitTaskData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	TaskId           string   `json:"task_id"`
}

type FaceFusionMovieSubmitTaskResult struct {
	ResponseMetadata *base.ResponseMetadata         `json:",omitempty"`
	RequestId        string                         `json:"request_id"`
	Code             int                            `json:"code"`
	Message          string                         `json:"message"`
	Status           int                            `json:"status"`
	TimeElapsed      string                         `json:"time_elapsed"`
	Data             *FaceFusionMovieSubmitTaskData `json:"data"`
}

type FaceFusionMovieGetResultData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	RespData         string   `json:"resp_data"`
	Status           string   `json:"status"`
}

type FaceFusionMovieGetResultResult struct {
	ResponseMetadata *base.ResponseMetadata        `json:",omitempty"`
	RequestId        string                        `json:"request_id"`
	Code             int                           `json:"code"`
	Message          string                        `json:"message"`
	Status           int                           `json:"status"`
	TimeElapsed      string                        `json:"time_elapsed"`
	Data             *FaceFusionMovieGetResultData `json:"data"`
}
