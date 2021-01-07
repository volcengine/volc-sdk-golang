package model

type UploadPartCommonResponse struct {
	Version string          `json:"Version"`
	Success int             `json:"success,omitempty"`
	Error   UploadPartError `json:"error"`
}

type UploadPartResponse struct {
	UploadPartCommonResponse
	PayLoad InitPartPayLoad `json:"payload,omitempty"`
}

type UploadMergeResponse struct {
	UploadPartCommonResponse
	PayLoad MergePartPayLoad `json:"payload,omitempty"`
}

type MergePartPayLoad struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type UploadPartError struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type InitPartPayLoad struct {
	UploadID string `json:"uploadID"`
}

type UploadPartCommon struct {
	TosHost string
	Oid     string
	Auth    string
}
