package model

import "io"

type UploadPartCommonResponse struct {
	Version string          `json:"Version"`
	Success int             `json:"success,omitempty"`
	Error   UploadPartError `json:"error"`
}

type InitPartResponse struct {
	UploadPartCommonResponse
	PayLoad InitPartPayLoad `json:"payload,omitempty"`
}

type UploadPartResponse struct {
	UploadPartCommonResponse
	PayLoad    UploadPartPayLoad `json:"payload,omitempty"`
	PartNumber int
	CheckSum   string
}

type Meta struct {
	ObjectContentType string
}

type UploadPartPayLoad struct {
	Uploadid   string `json:"uploadid"`
	PartNumber string `json:"part_number"`
	Crc32      string `json:"crc32"`
	Etag       string `json:"etag"`
	Meta       Meta   `json:"meta"`
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
	TosHost           string
	Oid               string
	Auth              string
	ObjectContentType string
}

type VodUploadFuncRequest struct {
	FilePath          string
	Rd                io.Reader
	Size              int64
	ParallelNum       int
	SpaceName         string
	FileType          string
	FileName          string
	FileExtension     string
	StorageClass      int32
	ClientNetWorkMode string
	ClientIDCMode     string
}

type VodUploadMediaInnerFuncRequest struct {
	FilePath          string
	Rd                io.Reader
	Size              int64
	ParallelNum       int
	SpaceName         string
	FileType          string
	CallbackArgs      string
	Funcs             string
	FileName          string
	FileExtension     string
	StorageClass      int32
	VodUploadSource   string
	ClientNetWorkMode string
	ClientIDCMode     string
}
