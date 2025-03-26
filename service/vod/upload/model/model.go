package model

import (
	"io"
	"net/http"
)

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
	Code      int    `json:"code"`
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type InitPartPayLoad struct {
	UploadID string `json:"uploadID"`
}

type UploadPartCommon struct {
	TosHost           string
	Oid               string
	Auth              string
	ObjectContentType string
	ChunkSize         int64
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
	ExpireTime        string
	UploadHostPrefer  string
	ChunkSize         int64
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
	ExpireTime        string
	UploadHostPrefer  string
	ChunkSize         int64
}

type UploadAuthOpt func(option *UploadAuthOption)

type UploadAuthOption struct {
	KeyPtn       string
	SpaceNames   []string
	UploadPolicy *UploadPolicy
}

type UploadPolicy struct {
	ContentTypeBlackList []string `json:"ContentTypeBlackList,omitempty"` // 上传文件Content-Type黑名单
	ContentTypeWhiteList []string `json:"ContentTypeWhiteList,omitempty"` // 上传文件Content-Type白名单，建议不和黑名单同时使用
	FileSizeUpLimit      string   `json:"FileSizeUpLimit,omitempty"`      // 上传文件大小上限
	FileSizeBottomLimit  string   `json:"FileSizeBottomLimit,omitempty"`  // 上传文件大小下限
}

type VpcUploadPartsInfo struct {
	Parts []*VpcUploadPartInfo `json:"Parts"`
}

type VpcUploadPartInfo struct {
	PartNumber int    `json:"PartNumber"`
	ETag       string `json:"ETag"`
}

type UploadCommonInfo struct {
	Auth               string
	Hosts              []string
	Oid                string
	Client             *http.Client
	StorageClass       int32
	SessionKey         string
	PreferredHostIndex int
}

type CreateMultipartUploadInput struct {
	UploadCommonInfo *UploadCommonInfo
}

type UploadPartInput struct {
	UploadCommonInfo *UploadCommonInfo
	PartNumber       int64
	UploadId         string
	Data             []byte
	Content          io.Reader
}

type CompleteMultipartUploadInput struct {
	UploadCommonInfo *UploadCommonInfo
	UploadId         string
	PartList         []*UploadPartResponse
}

type PutObjectInput struct {
	UploadCommonInfo *UploadCommonInfo
	Data             []byte
	Content          io.Reader
}

type VodStreamUploadRequest struct {
	// 上传的空间名
	SpaceName string `json:"SpaceName,omitempty"`
	// 上传内容
	Content io.Reader `json:"Content,omitempty"`
	// 上传内容大小
	Size int64 `json:"Size,omitempty"`
	// 业务希望上传透传的信息，会在上传成功时返回给用户
	CallbackArgs string `json:"CallbackArgs,omitempty"`
	// 上传的功能函数
	Functions string `json:"Functions,omitempty"`
	// 上传的文件在存储中的名字，即 bucket/key 中的 key
	FileName string `json:"FileName,omitempty"`
	// 上传文件类型
	FileType string `json:"FileType"`
	// 上传的文件的存储类型，1-标准存储，2-归档存储，非必填参数，默认为标准存储
	StorageClass int32 `json:"StorageClass,omitempty"`
	// 上传中文件的文件后缀
	FileExtension string `json:"FileExtension,omitempty"`
	// 上传中文件的来源
	VodUploadSource string `json:"VodUploadSource,omitempty"`
	// 上传策略
	UploadStrategy int32 `json:"UploadStrategy,omitempty"`
	// 客户端网络环境
	ClientNetWorkMode string `json:"ClientNetWorkMode,omitempty"`
	// 客户端机房环境
	ClientIDCMode string `json:"ClientIDCMode,omitempty"`
	// 媒资文件过期时间,采用ISO日期格式. 不传或传空,不修改.
	// 填"9999-12-31T23:59:59Z"表示永不过期.
	// 过期后该媒资文件及其相关资源（转码结果、封面图等）将被永久删除.
	// 示例值:2024-08-30T20:10:11+08:00
	ExpireTime string `json:"ExpireTime,omitempty"`
	// 上传域名偏好
	UploadHostPrefer string `json:"UploadHostPrefer,omitempty"`
	// 大文件上传分片大小，最小20MB
	ChunkSize int64 `json:"ChunkSize,omitempty"`
}

type UploadContentParam struct {
	UploadCommonInfo *UploadCommonInfo
	ChunkSize        int64
	Size             int64
	Content          io.Reader
}
