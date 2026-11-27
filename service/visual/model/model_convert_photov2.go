package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type ConvertPhotoV2Request struct {
	ReqKey           string   `json:"req_key"`
	BinaryDataBase64 []string `json:"binary_data_base64,omitempty"`
	ImageUrls        []string `json:"image_urls"`
	IsColor          bool     `json:"is_color"` // Deprecated
	IfColor          int      `json:"if_color"`
}

// 响应参数

type ConvertPhotoV2Data struct {
	AlgorithmBaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
}

type ConvertPhotoV2Result struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *ConvertPhotoV2Data    `json:"data"`
}
