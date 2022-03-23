package diff

import (
	"encoding/json"
	"time"
)

// common response
type RespInfo struct {
	Code int    `json:"error_no"`  // 错误码
	Msg  string `json:"error_msg"` // 错误信息
}

// create_service
type CreateServiceReq struct {
	Name    string `json:"name"`    // 服务名称
	Comment string `json:"comment"` // 备注信息（可选）
}

type CreateServiceData struct {
	ServiceId int `json:"service_id"` // 新建好服务的id
}

type CreateServiceResp struct {
	RespInfo
	Data *CreateServiceData `json:"data"`
}

// validate
type ValidateReq struct {
	OldUrl   string `json:"old_url"`   // 原始包url
	PatchUrl string `json:"patch_url"` // 新包url
}

// gen_by_pkg
type GenByPkgReq struct {
	ServiceId  int    `json:"service_id"`  // 服务id
	Alg        string `json:"alg"`         // 差分算法
	OldUrl     string `json:"old_url"`     // 原始包url
	OldVersion string `json:"old_version"` // 原始包版本
	NewUrl     string `json:"new_url"`     // 新包url
	NewVersion string `json:"new_version"` // 新包版本
}

// gen_by_count
type GenByCountReq struct {
	ServiceId  int    `json:"service_id"`  // 服务id
	Alg        string `json:"alg"`         // 差分算法
	NewUrl     string `json:"new_url"`     // 新包url
	NewVersion string `json:"new_version"` // 新包版本
	Count      int8   `json:"count"`       // 需要生成的差分包的个数
}

// gen_by_version
type GenByVersionReq struct {
	ServiceId   int      `json:"service_id"`   // 服务id
	Alg         string   `json:"alg"`          // 差分算法
	NewUrl      string   `json:"new_url"`      // 新包url
	NewVersion  string   `json:"new_version"`  // 新包版本
	OldVersions []string `json:"old_versions"` // 需要生成差分包的原始包版本
}

// 所有异步接口的统一返回结构
type AsyncResp struct {
	RespInfo
	Data *AsyncData `json:"data"`
}

type AsyncData struct {
	TaskId int64 `json:"task_id"` // 生成差分包请求的id（后续用这个id查询生成结果）
}

// check_response
type CheckResponseReq struct {
	TaskId int64 `json:"task_id"` // 任务ID
}

type CheckResponseResp struct {
	RespInfo
	Data *checkResponseData
}

type checkResponseData struct {
	TaskId   int64           `json:"task_id"`  // 请求ID
	Api      int8            `json:"api"`      // 生成该任务的api名称
	Response json.RawMessage `json:"response"` // 请求返回值
}

// 生成差分包统一返回的结构
type GenResult struct {
	RespInfo
	Data *[]RespPatch `json:"data"` // 成功时携带数据
}

type RespPatch struct {
	Code       int     `json:"code"`        // 差分包状态
	ID         uint    `json:"id"`          // 唯一ID
	CreatedAt  int64   `json:"created_at"`  // 创建时间
	Url        string  `json:"patch_url"`   // patch包地址
	Md5        string  `json:"patch_md5"`   // patch包MD5
	Size       int64   `json:"patch_size"`  // patch包Size
	Alg        string  `json:"alg"`         // 差分算法类型
	Ratio      float64 `json:"ratio"`       // 差分效率
	NewVersion string  `json:"new_version"` // 新包版本
	OldVersion string  `json:"old_version"` // 原始包版本
}

// 验证差分包的返回结构
type ValidateResult struct {
	RespInfo
	Data *ValidateData
}

type ValidateData struct {
	CreatedAt time.Time `json:"created_at"` // 创建时间
	NewMd5    string    `json:"new_md5"`    // 合成包md5
	NewSize   int64     `json:"new_size"`   // 合成包大小（Byte）
	Alg       string    `json:"alg"`        // 差分算法
}
