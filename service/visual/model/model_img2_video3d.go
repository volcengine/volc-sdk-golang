package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type Img2Video3DRenderSpec struct {
	Mode       int64     `json:"mode"`
	LongSide   int64     `json:"long_side"`
	FrameNum   int64     `json:"frame_num"`
	Fps        int64     `json:"fps"`
	UseFlow    int64     `json:"use_flow"`
	SpeedShift []float64 `json:"speed_shift"`
}

type Img2Video3DRequest struct {
	ReqKey           string                 `json:"req_key"`
	BinaryDataBase64 []string               `json:"binary_data_base64"`
	RenderSpec       *Img2Video3DRenderSpec `json:"render_spec"`
}

// 响应参数

type Img2Video3DData struct {
	AlgorithmBaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
}

type Img2Video3DResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *Img2Video3DData       `json:"data"`
}
