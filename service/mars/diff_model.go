package mars

import (
	"encoding/json"
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

type CreateServiceResp struct {
	RespInfo
	Data *CreateServiceData `json:"data"`
}

type CreateServiceData struct {
	ServiceId int `json:"service_id"` // 新建好服务的id
}

// validate
type ValidateReq struct {
	OldUrl   string `json:"old_url"`   // 原始包url
	PatchUrl string `json:"patch_url"` // 新包url
}

// gen_by_pkg
type GenByPkgReq struct {
	ServiceId    int                    `json:"service_id"`     // 服务id
	Alg          string                 `json:"alg"`            // 差分算法
	Options      map[string]interface{} `json:"options"`        // 配置项, key-value字典
	OldUrl       string                 `json:"old_url"`        // 原始包url
	OldVersion   string                 `json:"old_version"`    // 原始包版本
	OldExtraInfo string                 `json:"old_extra_info"` // 业务方对于原始包的附带信息
	NewUrl       string                 `json:"new_url"`        // 新包url
	NewVersion   string                 `json:"new_version"`    // 新包版本
	NewExtraInfo string                 `json:"new_extra_info"` // 业务方对于新包的附带信息
}

// gen_by_count
type GenByCountReq struct {
	ServiceId    int                    `json:"service_id"`     // 服务id
	Alg          string                 `json:"alg"`            // 差分算法
	Options      map[string]interface{} `json:"options"`        // 配置项, key-value字典
	NewUrl       string                 `json:"new_url"`        // 新包url
	NewVersion   string                 `json:"new_version"`    // 新包版本
	NewExtraInfo string                 `json:"new_extra_info"` // 业务方对于新包的附带信息
	Count        int8                   `json:"count"`          // 需要生成的差分包的个数
}

// gen_by_version
type GenByVersionReq struct {
	ServiceId    int                    `json:"service_id"`     // 服务id
	Alg          string                 `json:"alg"`            // 差分算法
	Options      map[string]interface{} `json:"options"`        // 配置项
	NewUrl       string                 `json:"new_url"`        // 新包url
	NewVersion   string                 `json:"new_version"`    // 新包版本
	NewExtraInfo string                 `json:"new_extra_info"` // 业务方对于新包的附带信息
	OldVersions  []string               `json:"old_versions"`   // 需要生成差分包的原始包版本
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
	Data *checkResponseData `json:"data"`
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
	Code          int     `json:"code"`            // 差分包状态
	ID            uint    `json:"id"`              // 唯一ID
	CreatedAt     int64   `json:"created_at"`      // 创建时间(毫秒)
	Url           string  `json:"patch_url"`       // patch包地址
	Md5           string  `json:"patch_md5"`       // patch包MD5
	Size          int64   `json:"patch_size"`      // patch包Size
	Alg           string  `json:"alg"`             // 差分算法类型
	Ratio         float64 `json:"ratio"`           // 差分效率
	NewVersion    string  `json:"new_version"`     // 新包版本
	NewExtraInfo  string  `json:"new_extra_info"`  // 业务方传入的新包附带信息
	OldVersion    string  `json:"old_version"`     // 原始包版本
	OldInfo       string  `json:"old_info"`        // 真正参与差分的原始包的信息
	OldExtraInfo  string  `json:"old_extra_info"`  // 业务方传入的原始包包附带信息
	MinSdkVersion int     `json:"min_sdk_version"` // 最小支持的sdk版本号
}

// 验证差分包的返回结构
type ValidateResult struct {
	RespInfo
	Data *ValidateData `json:"data"`
}

type ValidateData struct {
	CreatedAt int64  `json:"created_at"` // 创建时间(毫秒)
	NewMd5    string `json:"new_md5"`    // 合成包md5
	NewSize   int64  `json:"new_size"`   // 合成包大小（Byte）
	Alg       string `json:"alg"`        // 差分算法
}

// 删除原始包
type DeletePackagesReq struct {
	ServiceId   int      `json:"service_id"`   // 服务id
	OldVersions []string `json:"old_versions"` // 需要生成差分包的版本
}

type DeletePackagesResp struct {
	RespInfo
	Data *DeletePackagesSinglePackage `json:"data"`
}

type DeletePackagesSinglePackage struct {
	Code            int    `json:"code"`              // 删除状态
	CreatedAt       int64  `json:"created_at"`        // 上传时间
	ServiceId       int    `json:"service_id"`        // 原始包所在Service
	Url             string `json:"package_url"`       // 原始包下载地址
	Md5             string `json:"package_md5"`       // 原始包MD5
	Size            int64  `json:"package_size"`      // 原始包Size
	Version         string `json:"version"`           // 原始包版本
	NewPackageCount int    `json:"new_package_count"` // 删除的该包作为新包的差分包个数
	OldPackageCount int    `json:"old_package_count"` // 删除的该包作为原始包的差分包个数
}

// 删除服务
type DeleteServiceReq struct {
	ServiceId int `json:"service_id"` // 需要删除的服务的ID
}

type DeleteServiceResp struct {
	RespInfo
}

// 查看服务中的包信息
type QueryPatchByServiceReq struct {
	ServiceId    int    `json:"service_id"`    // service id
	StartTime    int64  `json:"start_time"`    // 起始时间（包含）（可选）
	EndTime      int64  `json:"end_time"`      // 结束时间（包含）（可选）
	StartVersion string `json:"start_version"` // 起始版本（包含）（可选）
	EndVersion   string `json:"end_version"`   // 最终版本（包含）（可选）
	NoPatches    bool   `json:"no_patches"`    // 不包含差分包信息（可选）
}

type QueryPatchByServiceResp struct {
	RespInfo
	Data *RespQueryPatchByServiceOverview `json:"data"`
}

type RespQueryPatchByServiceOverview struct {
	ServiceName string                    `json:"service_name"`
	Packages    []RespQueryPatchByService `json:"packages"`
}

type RespQueryPatchByService struct {
	Version   string                         `json:"version"`    // 完整包版本号
	Md5       string                         `json:"md5"`        // 完整包MD5
	CreatedAt int64                          `json:"created_at"` // 上传时间
	Size      int64                          `json:"size"`       // 完整包大小
	Url       string                         `json:"url"`        // 完整包下载地址
	ExtraInfo string                         `json:"extra_info"` // 业务方随包传入的附带信息
	PatchList []RespQueryPatchByServicePatch `json:"patch_list"` // 差分包列表
}

type RespQueryPatchByServicePatch struct {
	OldVersion string `json:"old_version"` // 原始包版本
	OldMd5     string `json:"old_md5"`     // 原始包MD5
	OldSize    int64  `json:"old_size"`    // 原始包大小
	OldUrl     string `json:"old_url"`     // 原始包下载地址
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	Md5        string `json:"md5"`         // 差分包MD5
	Size       int64  `json:"size"`        // 描述
	Url        string `json:"url"`         // 差分包下载地址
	Alg        string `json:"alg"`         // 算法
}
