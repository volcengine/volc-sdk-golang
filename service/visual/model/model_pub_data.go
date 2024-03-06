package model

type VisualBaseRespDataV2 struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp,omitempty"`
	BinaryDataBase64 []string `json:"binary_data_base64,omitempty"`
}
