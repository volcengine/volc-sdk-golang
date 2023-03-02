package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type BodyDetectionRequest struct {
	ReqKey           string   `json:"req_key"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	MaxObjNum        int      `json:"max_obj_num"`
}

// 响应参数

type BodyDetectionData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	Result           [][]*struct {
		Box *struct {
			X1 string `json:"x1"`
			X2 string `json:"x2"`
			Y1 string `json:"y1"`
			Y2 string `json:"y2"`
		} `json:"box"`
		ChnName string `json:"chn_name"`
		EngName string `json:"eng_name"`
		Score   string `json:"score"`
	} `json:"result"`
}

type BodyDetectionResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *BodyDetectionData     `json:"data"`
}
