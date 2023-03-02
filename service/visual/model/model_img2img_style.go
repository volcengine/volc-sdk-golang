package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type Img2ImgStyleRequest struct {
	ReqKey           string   `json:"req_key"`
	Prompt           string   `json:"prompt"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	Strength         float64  `json:"strength"`
	Seed             int64    `json:"seed"`
	ImageUrl         string   `json:"image_url"`
}

// 响应参数

type Img2ImgStyleData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	ImageUrl         string   `json:"image_url"`
}

type Img2ImgStyleResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *Img2ImgStyleData      `json:"data"`
}
