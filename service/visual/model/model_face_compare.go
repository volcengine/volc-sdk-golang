package model

import "github.com/volcengine/volc-sdk-golang/base"

type FaceCompareData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	Confidence       string   `json:"confidence"`
	IsMatch          bool     `json:"is_match"`
	RectAList        []string `json:"rect_a_list"`
	RectBList        []string `json:"rect_b_list"`
	Thresholds       []string `json:"thresholds"`
}

type FaceCompareResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *FaceCompareData       `json:"data"`
}
