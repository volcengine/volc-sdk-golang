package mcdn

import "fmt"

type DnsScheduleRegionEnum string

type DnsScheduleStatusEnum string

type CdnTypeEnum string

const (
	CdnTypeEnumWeb      CdnTypeEnum = "Web"      // 网页/小文件加速
	CdnTypeEnumVideo    CdnTypeEnum = "Video"    // 音视频点播加速
	CdnTypeEnumDownload CdnTypeEnum = "Download" // 大文件下载加速
)

var CdnTypeEnumList = []CdnTypeEnum{CdnTypeEnumWeb, CdnTypeEnumVideo, CdnTypeEnumDownload}

func (e CdnTypeEnum) Validate() error {
	for _, item := range CdnTypeEnumList {
		if item == e {
			return nil
		}
	}
	return fmt.Errorf("unknown typedef value CdnTypeEnum(%v)", e)
}

type SyncStatus string

const (
	SyncStatusSuccessful SyncStatus = "Successful" // 成功
	SyncStatusFailed     SyncStatus = "Failed"     // 失败
	SyncStatusSyncing    SyncStatus = "Syncing"    // 同步中
)

var SyncStatusList = []SyncStatus{SyncStatusSuccessful, SyncStatusFailed, SyncStatusSyncing}

func (e SyncStatus) Validate() error {
	for _, item := range SyncStatusList {
		if item == e {
			return nil
		}
	}
	return fmt.Errorf("unknown typedef value SyncStatus(%v)", e)
}

type RegionEnum string

const (
	RegionEnumDomestic RegionEnum = "Domestic" // 国内
	RegionEnumOverSea  RegionEnum = "OverSea"  // 海外
	RegionEnumGlobal   RegionEnum = "Global"   // 全球
)

var RegionEnumList = []RegionEnum{RegionEnumDomestic, RegionEnumOverSea, RegionEnumGlobal}

func (e RegionEnum) Validate() error {
	for _, item := range RegionEnumList {
		if item == e {
			return nil
		}
	}
	return fmt.Errorf("unknown typedef value RegionEnum(%v)", e)
}

type SyncStatusEnum string

const (
	SyncStatusEnumSuccess SyncStatusEnum = "Success" // 同步成功
	SyncStatusEnumFailed  SyncStatusEnum = "Failed"  // 同步失败
	SyncStatusEnumSyncing SyncStatusEnum = "Syncing" // 同步中
)

var SyncStatusEnumList = []SyncStatusEnum{SyncStatusEnumSuccess, SyncStatusEnumFailed, SyncStatusEnumSyncing}

func (e SyncStatusEnum) Validate() error {
	for _, item := range SyncStatusEnumList {
		if item == e {
			return nil
		}
	}
	return fmt.Errorf("unknown typedef value SyncStatusEnum(%v)", e)
}

type DomainStatusEnum string

const (
	DomainStatusEnumDeploying       DomainStatusEnum = "Deploying"       // 部署中
	DomainStatusEnumOnline          DomainStatusEnum = "Online"          // 已启用
	DomainStatusEnumOffline         DomainStatusEnum = "Offline"         // 已停用
	DomainStatusEnumDeployingFailed DomainStatusEnum = "DeployingFailed" // 部署失败
	DomainStatusEnumLocked          DomainStatusEnum = "Locked"          // 已封禁
	DomainStatusEnumLocking         DomainStatusEnum = "Locking"         // 封禁中
	DomainStatusEnumStopping        DomainStatusEnum = "Stopping"        // 停用中
	DomainStatusEnumDeleting        DomainStatusEnum = "Deleting"        // 删除中
	DomainStatusEnumReviewing       DomainStatusEnum = "Reviewing"       // 审核中
	DomainStatusEnumReviewFailed    DomainStatusEnum = "ReviewFailed"    // 审核失败
)

var DomainStatusEnumList = []DomainStatusEnum{DomainStatusEnumDeploying, DomainStatusEnumOnline, DomainStatusEnumOffline, DomainStatusEnumDeployingFailed, DomainStatusEnumLocked, DomainStatusEnumLocking, DomainStatusEnumStopping, DomainStatusEnumDeleting, DomainStatusEnumReviewing, DomainStatusEnumReviewFailed}

func (e DomainStatusEnum) Validate() error {
	for _, item := range DomainStatusEnumList {
		if item == e {
			return nil
		}
	}
	return fmt.Errorf("unknown typedef value DomainStatusEnum(%v)", e)
}

type EmptyRequest struct {
}

type EmptyResponse struct {
}

type PagingMarkerOption struct {

	// 每页的数量，最大值是100。
	PageSize   int64  `form:"PageSize,required" json:"PageSize,required" query:"PageSize,required"`
	PageMarker string `form:"PageMarker,required" json:"PageMarker,required" query:"PageMarker,required"`
}

type PagingMarkerResult struct {

	// 每页的数量。
	PageSize       int64  `form:"PageSize,required" json:"PageSize,required" query:"PageSize,required"`
	NextPageMarker string `form:"NextPageMarker,required" json:"NextPageMarker,required" query:"NextPageMarker,required"`
}

type PagingOption struct {

	// 每页的数量，最大值是100。
	PageSize int64 `form:"PageSize,required" json:"PageSize,required" query:"PageSize,required"`

	// 页码，从1开始。
	PageNum int64 `form:"PageNum,required" json:"PageNum,required" query:"PageNum,required"`
}

type PagingResult struct {

	// 每页的数量。
	PageSize int64 `form:"PageSize,required" json:"PageSize,required" query:"PageSize,required"`

	// 页码，从1开始。
	PageNum int64 `form:"PageNum,required" json:"PageNum,required" query:"PageNum,required"`

	// 总页数。
	Total int64 `form:"Total,required" json:"Total,required" query:"Total,required"`
}

type ListCloudAccountsRequest struct {

	// 列出该厂商下的所有云厂商账号
	Vendor *string `form:"Vendor" json:"Vendor,omitempty" query:"Vendor"`

	// 分页设置（不传则PageSize为10，PageNum为1）
	Pagination *PagingOption `form:"Pagination" json:"Pagination,omitempty" query:"Pagination"`
}

type ListCloudAccountsResponse struct {

	// 云厂商账号列表
	CloudAccounts []*CloudAccountForList `form:"CloudAccounts,required" json:"CloudAccounts,required" query:"CloudAccounts,required"`

	// 分页信息
	Pagination *PagingResult `form:"Pagination,required" json:"Pagination,required" query:"Pagination,required"`
}

type AwsContentSetting struct {

	// 预热设置
	Preload *AwsPreloadContentSetting `form:"Preload" json:"Preload,omitempty" query:"Preload"`
}

type AwsPreloadContentSetting struct {

	// 是否开启
	Enabled bool `form:"Enabled,required" json:"Enabled,required" query:"Enabled,required"`

	// 提交接口Host
	SubmitEndpoint string `form:"SubmitEndpoint,required" json:"SubmitEndpoint,required" query:"SubmitEndpoint,required"`

	// 查询接口Host
	QueryEndpoint string `form:"QueryEndpoint,required" json:"QueryEndpoint,required" query:"QueryEndpoint,required"`

	// api key
	ApiKey string `form:"ApiKey,required" json:"ApiKey,required" query:"ApiKey,required"`
}

type AwsStatSetting struct {

	// 是否开启
	Enabled bool `form:"Enabled,required" json:"Enabled,required" query:"Enabled,required"`

	// 接口Host
	Endpoint string `form:"Endpoint,required" json:"Endpoint,required" query:"Endpoint,required"`

	// Api key
	ApiKey string `form:"ApiKey,required" json:"ApiKey,required" query:"ApiKey,required"`
}

type CloudAccountContentSettings struct {

	// AWS 内容管理设置
	Aws *AwsContentSetting `form:"Aws" json:"Aws,omitempty" query:"Aws"`
}

type CloudAccountDomainSettings struct {

	// UCloud 域名设置
	UCloud *UCloudDomainSettings `form:"UCloud" json:"UCloud,omitempty" query:"UCloud"`
}

type CloudAccountExtra struct {
	ProductId   *string `form:"ProductId" json:"ProductId,omitempty" query:"ProductId"`
	AccessToken *string `form:"AccessToken" json:"AccessToken,omitempty" query:"AccessToken"`

	// Used by akamai
	AkamaiEndpoint *string `form:"AkamaiEndpoint" json:"AkamaiEndpoint,omitempty" query:"AkamaiEndpoint"`
}

type CloudAccountForList struct {

	// 云厂商账号Id
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 云厂商账号名
	Name string `form:"Name,required" json:"Name,required" query:"Name,required"`

	// 云厂商唯一名
	Vendor string `form:"Vendor,required" json:"Vendor,required" query:"Vendor,required"`

	// 总的同步状态 (Successful, Failed, Syncing)
	SyncStatus SyncStatus `form:"SyncStatus,required" json:"SyncStatus,required" query:"SyncStatus,required"`

	// 创建时间(Unix时间戳)
	CreatedAt int64 `form:"CreatedAt,required" json:"CreatedAt,required" query:"CreatedAt,required"`

	// 更新时间(Unix时间戳)
	UpdatedAt int64 `form:"UpdatedAt,required" json:"UpdatedAt,required" query:"UpdatedAt,required"`

	// 特定云厂商的额外参数
	Extra *CloudAccountExtra `form:"Extra,required" json:"Extra,required" query:"Extra,required"`

	// 统计分析设置
	StatSettings *CloudAccountStatSettings `form:"StatSettings,required" json:"StatSettings,required" query:"StatSettings,required"`

	// 内容管理设置
	ContentSettings *CloudAccountContentSettings `form:"ContentSettings,required" json:"ContentSettings,required" query:"ContentSettings,required"`

	// 域名设置
	DomainSettings *CloudAccountDomainSettings `form:"DomainSettings,required" json:"DomainSettings,required" query:"DomainSettings,required"`

	// 当前使用的产品列表
	SubProducts []string `form:"SubProducts,required" json:"SubProducts,required" query:"SubProducts,required"`
}

type CloudAccountStatSettings struct {

	// AWS 统计分析设置
	Aws *AwsStatSetting `form:"Aws" json:"Aws,omitempty" query:"Aws"`
}

type UCloudDomainSettings struct {

	// 项目组ID
	ProjectId string `form:"ProjectId,required" json:"ProjectId,required" query:"ProjectId,required"`
}

type SubmitRefreshTaskRequest struct {

	// URL刷新500条；目录刷新20条；URL必须以http://或https://开头，且URL之间需通过换行符分隔；域名必须属于已接入纳管的域名列表
	Urls string `form:"Urls,required" json:"Urls,required" query:"Urls,required"`

	// 刷新类型，在提交刷新任务时必填。取值：file：文件刷新。dir：目录刷新。
	Type string `form:"Type,required" json:"Type,required" query:"Type,required"`

	// 指定厂商提交，用,分隔，默认为不指定，由多云自动选择
	Vendor string `form:"Vendor" json:"Vendor" query:"Vendor"`
}

type SubmitTaskResponse struct {

	// 多云任务id
	TaskId string `form:"TaskId" json:"TaskId" query:"TaskId"`
}

type SubmitPreloadTaskRequest struct {

	// URL预热20条；URL必须以http://或https://开头，且URL之间需通过换行符分隔；域名必须属于已接入纳管的域名列表
	Urls string `form:"Urls,required" json:"Urls,required" query:"Urls,required"`

	// 指定厂商提交，用,分隔，默认为不指定，由多云自动选择
	Vendor string `form:"Vendor" json:"Vendor" query:"Vendor"`
}

type ListContentTaskRequest struct {

	// 多云任务id
	TaskId string `form:"TaskId" json:"TaskId" query:"TaskId"`

	// 任务类型 刷新目录 refresh_dir; 刷新文件 refresh_file; 预热文件 preload
	TaskType string `form:"TaskType" json:"TaskType" query:"TaskType"`

	// 开始时间
	StartTime int64 `form:"StartTime" json:"StartTime" query:"StartTime"`

	// 结束时间
	EndTime int64 `form:"EndTime" json:"EndTime" query:"EndTime"`

	// 域名精确匹配
	Domain string `form:"Domain" json:"Domain" query:"Domain"`

	// url精确匹配
	Url string `form:"Url" json:"Url" query:"Url"`

	// 分页信息
	Pagination *PagingOption `form:"Pagination" json:"Pagination,omitempty" query:"Pagination"`
}

type ListContentTaskResponse struct {

	// 多云任务列表
	Tasks []*TaskInfo `form:"Tasks" json:"Tasks" query:"Tasks"`

	// 分页信息
	Pagination *PagingResult `form:"Pagination" json:"Pagination" query:"Pagination"`
}

type TaskInfo struct {

	// 多云任务id
	TaskId string `form:"TaskId" json:"TaskId" query:"TaskId"`

	// 指定厂商时对应的厂商名
	Vendor string `form:"Vendor" json:"Vendor" query:"Vendor"`

	// 任务类型 刷新目录 refresh_dir; 刷新文件 refresh_file; 预热文件 preload
	TaskType string `form:"TaskType" json:"TaskType" query:"TaskType"`

	// 提交的各个厂商任务详情
	SubTasks []*SubTaskInfo `form:"SubTasks" json:"SubTasks" query:"SubTasks"`

	// 创建时间
	CreatedAt int64 `form:"CreatedAt" json:"CreatedAt" query:"CreatedAt"`
}

type SubTaskInfo struct {

	// 厂商task id组，提交失败则没有厂商task id
	VendorTaskIds []string `form:"VendorTaskIds" json:"VendorTaskIds" query:"VendorTaskIds"`

	// 厂商名
	Vendor string `form:"Vendor" json:"Vendor" query:"Vendor"`

	// 提交状态 success：成功 failed: 失败
	SubmitStatus string `form:"SubmitStatus" json:"SubmitStatus" query:"SubmitStatus"`

	// 报错原因
	Description string `form:"Description" json:"Description" query:"Description"`

	// 厂商提交url组
	Url []string `form:"Url" json:"Url" query:"Url"`

	// 云厂商账号id
	CloudAccountId string `form:"CloudAccountId" json:"CloudAccountId" query:"CloudAccountId"`

	// 多云账号名
	CloudAccountName string `form:"CloudAccountName" json:"CloudAccountName" query:"CloudAccountName"`

	// 产品类型 静态 cdn; 动态 dcdn
	ProductType string `form:"ProductType" json:"ProductType" query:"ProductType"`
}

type DescribeContentTaskByTaskIdRequest struct {

	// 多云task id
	TaskId string `form:"TaskId,required" json:"TaskId,required" query:"TaskId,required"`
}

type DescribeContentTaskByTaskIdResponse struct {

	// 提交的各个厂商url任务详情
	SubTasks []*SubUrlInfo `form:"SubTasks" json:"SubTasks" query:"SubTasks"`

	// 云厂商查询过程中的信息
	VendorsMetaData []*ContentVendorMetaData `form:"VendorsMetaData" json:"VendorsMetaData" query:"VendorsMetaData"`
}

type SubUrlInfo struct {

	// 厂商task ids
	VendorTaskIds []string `form:"VendorTaskIds" json:"VendorTaskIds" query:"VendorTaskIds"`
	Url           string   `form:"Url" json:"Url" query:"Url"`

	// url 状态 complete: 完成，executing: 运行中，failed: 失败，pending：等待中, undefined: 未定义
	Status string `form:"Status" json:"Status" query:"Status"`

	// 创建时间
	CreatedAt int64 `form:"CreatedAt" json:"CreatedAt" query:"CreatedAt"`

	// 报错原因
	Description string `form:"Description" json:"Description" query:"Description"`

	// 多云taskId
	TaskId string `form:"TaskId" json:"TaskId" query:"TaskId"`

	// 多云账号Id
	CloudAccountId string `form:"CloudAccountId" json:"CloudAccountId" query:"CloudAccountId"`

	// 多云账号名
	CloudAccountName string `form:"CloudAccountName" json:"CloudAccountName" query:"CloudAccountName"`

	// 厂商
	Vendor string `form:"Vendor" json:"Vendor" query:"Vendor"`

	// 任务类型 刷新目录 refresh_dir; 刷新文件 refresh_file; 预热文件 preload
	TaskType string `form:"TaskType" json:"TaskType" query:"TaskType"`

	// 产品类型 静态 cdn; 动态 dcdn
	ProductType string `form:"ProductType" json:"ProductType" query:"ProductType"`
}

type ContentVendorMetaData struct {
	Vendor         string           `form:"Vendor" json:"Vendor" query:"Vendor"`
	RequestId      string           `form:"RequestId" json:"RequestId" query:"RequestId"`
	Cost           float64          `form:"Cost" json:"Cost" query:"Cost"`
	Error          *VendorMetaError `form:"Error" json:"Error,omitempty" query:"Error"`
	CloudAccountId string           `form:"CloudAccountId" json:"CloudAccountId" query:"CloudAccountId"`

	// 产品类型 静态 cdn; 动态 dcdn
	ProductType string `form:"ProductType" json:"ProductType" query:"ProductType"`
}

type VendorMetaData struct {
	Vendor    string           `form:"Vendor,required" json:"Vendor,required" query:"Vendor,required"`
	RequestId string           `form:"RequestId,required" json:"RequestId,required" query:"RequestId,required"`
	Cost      float64          `form:"Cost,required" json:"Cost,required" query:"Cost,required"`
	Error     *VendorMetaError `form:"Error" json:"Error,omitempty" query:"Error"`
}

type VendorMetaError struct {
	Detail string `form:"Detail,required" json:"Detail,required" query:"Detail,required"`
}

type ListVendorContentTaskRequest struct {

	// 厂商task id
	VendorTaskId string `form:"VendorTaskId" json:"VendorTaskId" query:"VendorTaskId"`

	// 任务类型 刷新目录 refresh_dir; 刷新文件 refresh_file; 预热文件 preload
	TaskType string `form:"TaskType,required" json:"TaskType,required" query:"TaskType,required"`

	// 云厂商账号id
	CloudAccountId string `form:"CloudAccountId,required" json:"CloudAccountId,required" query:"CloudAccountId,required"`

	// 开始时间 默认查询前一天记录
	StartTime int64 `form:"StartTime" json:"StartTime" query:"StartTime"`

	// 结束时间
	EndTime int64 `form:"EndTime" json:"EndTime" query:"EndTime"`

	// 产品类型 静态 cdn; 动态 dcdn，不填则默认请求cdn
	ProductType string `form:"ProductType" json:"ProductType" query:"ProductType"`

	// 分页信息
	Pagination *PagingOption `form:"Pagination" json:"Pagination,omitempty" query:"Pagination"`
}

type ListVendorContentTaskResponse struct {

	// 厂商url任务列表
	Tasks []*UrlInfo `form:"Tasks" json:"Tasks" query:"Tasks"`

	// 分页信息
	Pagination *PagingResult `form:"Pagination" json:"Pagination" query:"Pagination"`
}

type DescribeContentQuotaRequest struct {

	// 查询Cloud Account ID, 多个账号用,分隔
	CloudAccountId string `form:"CloudAccountId,required" json:"CloudAccountId,required" query:"CloudAccountId,required"`
}

type DescribeContentQuotaResponse struct {

	// 厂商配额
	Quotas []*QuotaDetail `form:"Quotas" json:"Quotas" query:"Quotas"`

	// 云厂商查询过程中的信息
	VendorsMetaData []*ContentVendorMetaData `form:"VendorsMetaData" json:"VendorsMetaData" query:"VendorsMetaData"`
}

type UrlInfo struct {

	// 厂商task id
	VendorTaskId string `form:"VendorTaskId" json:"VendorTaskId" query:"VendorTaskId"`

	// 提交url
	Url string `form:"Url" json:"Url" query:"Url"`

	// url 状态 complete: 完成，executing: 运行中，failed: 失败，pending：等待中, undefined: 未定义
	Status string `form:"Status" json:"Status" query:"Status"`

	// 创建时间
	CreatedAt int64 `form:"CreatedAt" json:"CreatedAt" query:"CreatedAt"`

	// 报错原因
	Description string `form:"Description" json:"Description" query:"Description"`
}

type QuotaDetail struct {

	// 当天文件刷新文件数量上限
	RefreshUrlQuota int64 `form:"RefreshUrlQuota" json:"RefreshUrlQuota" query:"RefreshUrlQuota"`

	// 当天文件刷新文件数量剩余
	RefreshUrlRemain int64 `form:"RefreshUrlRemain" json:"RefreshUrlRemain" query:"RefreshUrlRemain"`

	// 当天文件刷新目录数量上限
	RefreshDirQuota int64 `form:"RefreshDirQuota" json:"RefreshDirQuota" query:"RefreshDirQuota"`

	// 当天文件刷新目录数量剩余
	RefreshDirRemain int64 `form:"RefreshDirRemain" json:"RefreshDirRemain" query:"RefreshDirRemain"`

	// 当天预热数量上限
	PreloadQuota int64 `form:"PreloadQuota" json:"PreloadQuota" query:"PreloadQuota"`

	// 当天预热数量剩余
	PreloadRemain int64 `form:"PreloadRemain" json:"PreloadRemain" query:"PreloadRemain"`

	// 配额的区域 可以为空
	Area string `form:"Area" json:"Area" query:"Area"`

	// 厂商名
	Vendor string `form:"Vendor" json:"Vendor" query:"Vendor"`

	// 多云厂商账号
	CloudAccountId string `form:"CloudAccountId" json:"CloudAccountId" query:"CloudAccountId"`

	// 多云账号名
	CloudAccountName string `form:"CloudAccountName" json:"CloudAccountName" query:"CloudAccountName"`

	// 产品类型 静态 cdn; 动态 dcdn
	ProductType string `form:"ProductType" json:"ProductType" query:"ProductType"`
}

// 查询域名请求
type ListDomainsRequest struct {

	// 域名名称。
	Name *string `form:"Name" json:"Name,omitempty" query:"Name"`

	// 业务类型。
	CdnType []CdnTypeEnum `form:"CdnType" json:"CdnType,omitempty" query:"CdnType"`

	// 厂商名称。
	Vendor []string `form:"Vendor" json:"Vendor,omitempty" query:"Vendor"`

	// 域名状态。
	Status []DomainStatusEnum `form:"Status" json:"Status,omitempty" query:"Status"`

	// 域名服务区域。
	Region []RegionEnum `form:"Region" json:"Region,omitempty" query:"Region"`

	// 云账号ID。
	CloudAccountId *string `form:"CloudAccountId" json:"CloudAccountId,omitempty" query:"CloudAccountId"`

	// 域名名称, 精确搜索。
	// @ignore
	ExactName *string `form:"ExactName" json:"ExactName,omitempty" query:"ExactName"`

	// IAM 项目列表，不传不过滤
	VolcProjects []string `form:"VolcProjects" json:"VolcProjects,omitempty" query:"VolcProjects"`

	// 分页信息。不填则使用默认值分页。
	Pagination *PagingOption `form:"Pagination" json:"Pagination,omitempty" query:"Pagination"`
}

// 查询域名的响应
type ListDomainsResponse struct {

	// 域名列表。
	Domains []*Domain `form:"Domains,required" json:"Domains,required" query:"Domains,required"`

	// 分页信息。
	Pagination *PagingResult `form:"Pagination,required" json:"Pagination,required" query:"Pagination,required"`
}

// 域名简略信息
type Domain struct {

	// 域名在多云平台的ID。
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 域名名称。
	Name string `form:"Name,required" json:"Name,required" query:"Name,required"`

	// 域名CName。
	Cname string `form:"Cname,required" json:"Cname,required" query:"Cname,required"`

	// 域名状态。
	Status DomainStatusEnum `form:"Status,required" json:"Status,required" query:"Status,required"`

	// 业务类型。
	CdnType CdnTypeEnum `form:"CdnType,required" json:"CdnType,required" query:"CdnType,required"`

	// 厂商名称。
	Vendor string `form:"Vendor,required" json:"Vendor,required" query:"Vendor,required"`

	// 域名服务区域。
	Region RegionEnum `form:"Region,required" json:"Region,required" query:"Region,required"`

	// 火山引擎账号ID。
	TopAccountId string `form:"TopAccountId,required" json:"TopAccountId,required" query:"TopAccountId,required"`

	// 云账号ID。代表当前域名是使用该云账号从厂商侧同步过来的。
	CloudAccountId string `form:"CloudAccountId,required" json:"CloudAccountId,required" query:"CloudAccountId,required"`

	// 云账号名称，每次同步时更新。
	CloudAccountName string `form:"CloudAccountName,required" json:"CloudAccountName,required" query:"CloudAccountName,required"`

	// 标签。
	Tags []*Tag `form:"Tags,required" json:"Tags,required" query:"Tags,required"`

	// 域名所属的厂商的产品名
	SubProduct *string `form:"SubProduct" json:"SubProduct,omitempty" query:"SubProduct"`

	// 是否已创建调度
	ScheduleCreated *bool `form:"ScheduleCreated" json:"ScheduleCreated,omitempty" query:"ScheduleCreated"`

	// 同步状态。
	SyncStatus SyncStatusEnum `form:"SyncStatus,required" json:"SyncStatus,required" query:"SyncStatus,required"`

	// 同步时间。
	SyncedAt string `form:"SyncedAt,required" json:"SyncedAt,required" query:"SyncedAt,required"`
}

// 域名标签
type Tag struct {

	// 标签Key。
	Key string `form:"Key,required" json:"Key,required" query:"Key,required"`

	// 标签Value。
	Value string `form:"Value,required" json:"Value,required" query:"Value,required"`
}

// 获取调度列表
type ListDnsSchedulesRequest struct {

	// 模糊匹配调度域名
	DomainName *string `form:"DomainName" json:"DomainName,omitempty" query:"DomainName"`

	// 调度状态
	ScheduleStatus *DnsScheduleStatusEnum `form:"ScheduleStatus" json:"ScheduleStatus,omitempty" query:"ScheduleStatus"`

	// 厂商账号ID
	CloudAccountIds []string `form:"CloudAccountIds" json:"CloudAccountIds,omitempty" query:"CloudAccountIds"`

	// 域名名称, 精确搜索
	ExactDomainName *string `form:"ExactDomainName" json:"ExactDomainName,omitempty" query:"ExactDomainName"`

	// 分页信息
	Pagination *PagingOption `form:"Pagination" json:"Pagination,omitempty" query:"Pagination"`
}

// 调度信息
type DnsScheduleListItem struct {

	// 调度ID
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 调度域名，需要存在同名的厂商域名才能创建
	DomainName string `form:"DomainName,required" json:"DomainName,required" query:"DomainName,required"`

	// 调度Cname
	ScheduleCname string `form:"ScheduleCname,required" json:"ScheduleCname,required" query:"ScheduleCname,required"`

	// 调度状态
	ScheduleStatus DnsScheduleStatusEnum `form:"ScheduleStatus,required" json:"ScheduleStatus,required" query:"ScheduleStatus,required"`

	// 是否异常
	IsAbnormal bool `form:"IsAbnormal,required" json:"IsAbnormal,required" query:"IsAbnormal,required"`

	// 调度范围
	Region DnsScheduleRegionEnum `form:"Region,required" json:"Region,required" query:"Region,required"`

	// 厂商账号ID
	CloudAccountIds []string `form:"CloudAccountIds,required" json:"CloudAccountIds,required" query:"CloudAccountIds,required"`

	// 更新时间
	UpdatedAt int64 `form:"UpdatedAt,required" json:"UpdatedAt,required" query:"UpdatedAt,required"`
}

type ListDnsSchedulesResponse struct {

	// 调度列表
	DnsSchedules []*DnsScheduleListItem `form:"DnsSchedules,required" json:"DnsSchedules,required" query:"DnsSchedules,required"`

	// 分页信息
	Pagination *PagingResult `form:"Pagination,required" json:"Pagination,required" query:"Pagination,required"`
}

// 获取调度管理详情数据
type DescribeDnsScheduleRequest struct {

	// 调度ID
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`
}

type DescribeDnsScheduleResponse struct {

	// 调度详情数据
	DnsScheduleInfo *DnsScheduleInfo `form:"DnsScheduleInfo,required" json:"DnsScheduleInfo,required" query:"DnsScheduleInfo,required"`

	// 区域配置数据
	WeightInfos []*WeightInfo `form:"WeightInfos,required" json:"WeightInfos,required" query:"WeightInfos,required"`
}

// 区域配置子项数据
type WeightInfoItem struct {

	// 域名Id
	DomainId string `form:"DomainId,required" json:"DomainId,required" query:"DomainId,required"`

	// 设置比例
	Value int32 `form:"Value,required" json:"Value,required" query:"Value,required"`

	// 实际比例，该值会根据区域容灾而产生变化
	ActualValue int32 `form:"ActualValue,required" json:"ActualValue,required" query:"ActualValue,required"`
}

// 区域配置详情数据
type WeightInfo struct {

	// 区域配置ID
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 国家代码，空表示全球默认，即View数据的code字段值
	Country string `form:"Country,required" json:"Country,required" query:"Country,required"`

	// 省份代码
	Province string `form:"Province,required" json:"Province,required" query:"Province,required"`

	// 城市代码
	City string `form:"City,required" json:"City,required" query:"City,required"`

	// 运营商id
	Isp string `form:"Isp,required" json:"Isp,required" query:"Isp,required"`

	// 区域配置子项数据
	WeightInfoItems []*WeightInfoItem `form:"WeightInfoItems,required" json:"WeightInfoItems,required" query:"WeightInfoItems,required"`

	// 是否触发容灾
	IsFailover bool `form:"IsFailover,required" json:"IsFailover,required" query:"IsFailover,required"`

	// 区域配置的容灾id
	FailoverIds []string `form:"FailoverIds" json:"FailoverIds,omitempty" query:"FailoverIds"`
}

// 调度详细数据
type DnsScheduleInfo struct {

	// 调度ID
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 调度域名
	DomainName string `form:"DomainName,required" json:"DomainName,required" query:"DomainName,required"`

	// 调度Cname
	ScheduleCname string `form:"ScheduleCname,required" json:"ScheduleCname,required" query:"ScheduleCname,required"`

	// 调度状态
	ScheduleStatus DnsScheduleStatusEnum `form:"ScheduleStatus,required" json:"ScheduleStatus,required" query:"ScheduleStatus,required"`

	// 是否异常
	IsAbnormal bool `form:"IsAbnormal,required" json:"IsAbnormal,required" query:"IsAbnormal,required"`

	// 调度范围
	Region DnsScheduleRegionEnum `form:"Region,required" json:"Region,required" query:"Region,required"`

	// 全球区域绑定的域名
	GlobalDomains []*DnsScheduleDomain `form:"GlobalDomains" json:"GlobalDomains,omitempty" query:"GlobalDomains"`

	// 国内区域绑定的域名
	DomesticDomains []*DnsScheduleDomain `form:"DomesticDomains,required" json:"DomesticDomains,required" query:"DomesticDomains,required"`

	// 更新时间
	UpdatedAt int64 `form:"UpdatedAt,required" json:"UpdatedAt,required" query:"UpdatedAt,required"`

	// 项目组
	VolcProject string `form:"VolcProject,required" json:"VolcProject,required" query:"VolcProject,required"`
}

type DnsScheduleDomain struct {

	// 域名ID
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 域名名称
	Name string `form:"Name,required" json:"Name,required" query:"Name,required"`

	// 域名CNAME
	Cname string `form:"Cname,required" json:"Cname,required" query:"Cname,required"`

	// 厂商名称
	Vendor string `form:"Vendor,required" json:"Vendor,required" query:"Vendor,required"`

	// 域名服务区域
	Region RegionEnum `form:"Region,required" json:"Region,required" query:"Region,required"`

	// 云账号ID
	CloudAccountId string `form:"CloudAccountId,required" json:"CloudAccountId,required" query:"CloudAccountId,required"`

	// 是否CNAME冲突
	IsCnameConflict bool `form:"IsCnameConflict,required" json:"IsCnameConflict,required" query:"IsCnameConflict,required"`

	// 是否服务区域冲突
	IsRegionConflict bool `form:"IsRegionConflict,required" json:"IsRegionConflict,required" query:"IsRegionConflict,required"`

	// 是否状态异常
	IsStatusAbnormal bool `form:"IsStatusAbnormal,required" json:"IsStatusAbnormal,required" query:"IsStatusAbnormal,required"`

	// 是否异常
	IsAbnormal bool `form:"IsAbnormal,required" json:"IsAbnormal,required" query:"IsAbnormal,required"`
}

type DescribeCdnDataRequest struct {

	// 查询起始时间，相对于UTC 1970-01-01到当前时间相隔的秒数
	StartTime int64 `form:"StartTime,required" json:"StartTime,required" query:"StartTime,required"`

	// 查询结束时间，相对于UTC 1970-01-01到当前时间相隔的秒数。查询区间为前闭后开[StartTime, EndTime)，时间按照传入的Interval向前规整，如Interval为5min，1644163200和1644163499都会规整为1644163200
	EndTime int64 `form:"EndTime,required" json:"EndTime,required" query:"EndTime,required"`

	// 查询指标，仅支持单个，支持以下几种指标类型：1. flux：流量（单位为 byte），2. bandwidth：带宽（单位为 bps），3. request: 请求数（单位为 次），4. status_all：状态码，返回 2xx、3xx、4xx、5xx 汇总数据（单位为 个），5. status_2xx：返回 2xx 状态码汇总及各 2 开头状态码数据（单位为 个），6. status_3xx：返回 3xx 状态码汇总及各 3 开头状态码数据（单位为 个），7. status_4xx：返回 4xx 状态码汇总及各 4 开头状态码数据（单位为 个），8. status_5xx：返回 5xx 状态码汇总及各 5 开头状态码数据（单位为 个），9. hitrate: 流量命中率（单位为 %，小数点后保留两位），10. request_hitrate：请求命中率（单位为 %，小数点后保留两位）。DescribeCdnOriginData接口只能取1-8；有Region、Isp、Protocol、IpVersion参数时不支持hitrate和request_hitrate指标查询
	Metric string `form:"Metric,required" json:"Metric,required" query:"Metric,required"`

	// 时间粒度，默认5min，支持以下几种取值：1. 5min：5分钟粒度，指定查询区间3天内，3. hour：1小时粒度，指定查询区间31天内，4. day：天粒度，指定查询区间至少2天，最多90天
	Interval *string `form:"Interval" json:"Interval,omitempty" query:"Interval"`

	// 指定查询域名，多个域名逗号分隔，最多10个域名，不填充返回账号级别数据，账号级别只返回汇总数据，多域名默认返回汇总数据；不可和DomainId或CpCode同时使用
	Domain *string `form:"Domain" json:"Domain,omitempty" query:"Domain"`

	// 厂商名称，支持多个用逗号分隔，不可和DomainId同时使用
	Vendor *string `form:"Vendor" json:"Vendor,omitempty" query:"Vendor"`

	// 查询的国家代码。CN: 中国大陆，OHTER:中国境外
	Country *string `form:"Country" json:"Country,omitempty" query:"Country"`

	// 指定地域英文名，仅支持单个。地域名称列表通过DescribeCdnRegionAndIsp接口查询，DescribeCdnOriginData接口不支持分地区运营商查询
	Region *string `form:"Region" json:"Region,omitempty" query:"Region"`

	// 指定运营商英文名，仅支持单个。运营商名称列表通过DescribeCdnRegionAndIsp接口查询
	Isp *string `form:"Isp" json:"Isp,omitempty" query:"Isp"`

	// 默认aggregate，账号级别数据只支持aggregate，支持以下几种取值：1. aggregate：多域名/域名id/厂商查询返回聚合数据，2. disaggregate：多域名/域名id查询返回分域名/域名id/厂商数据
	Aggregate *string `form:"Aggregate" json:"Aggregate,omitempty" query:"Aggregate"`

	// 域名id，支持多个用逗号分隔，最多10个，不可和Domain、Vendor或CpCode参数同时使用
	DomainId *string `form:"DomainId" json:"DomainId,omitempty" query:"DomainId"`

	// 客户端请求应用层协议，不填充表示不区分协议查询。取值：http、https。不能和IpVersion/Region/Isp同时使用
	Protocol *string `form:"Protocol" json:"Protocol,omitempty" query:"Protocol"`

	// 客户端请求网络层协议，不填充表示不区分协议查询。取值：ipv4、ipv6。不能和Protocol/Region/Isp同时使用
	IpVersion *string `form:"IpVersion" json:"IpVersion,omitempty" query:"IpVersion"`

	// akamai cpcode, 不可和DomainId或domain同时使用
	CpCodeList []*CpCodeInfo `form:"CpCodeList" json:"CpCodeList,omitempty" query:"CpCodeList"`

	// 可选值：1. Web, 2. Video, 3. Download, 4. Dynamic。 动态域名必填，静态域名可选
	CdnType *string `form:"CdnType" json:"CdnType,omitempty" query:"CdnType"`

	// 默认1，可选值：1：表示从厂商api获取数据，2：从离线存储获取数据
	DataSource *string `form:"DataSource" json:"DataSource,omitempty" query:"DataSource"`

	// 云厂商账号，支持多个用逗号分隔
	CloudAccountId *string `form:"CloudAccountId" json:"CloudAccountId,omitempty" query:"CloudAccountId"`
}

type DescribeCdnDataResponse struct {

	// 查询得到的数据明细
	Resources []*ResourceStatData `form:"Resources" json:"Resources" query:"Resources"`

	// 云厂商查询过程中的信息
	VendorsMetaData []*VendorMetaData `form:"VendorsMetaData,required" json:"VendorsMetaData,required" query:"VendorsMetaData,required"`
}

type CpCodeInfo struct {

	// akamai cpcode
	CpCode string `form:"CpCode,required" json:"CpCode,required" query:"CpCode,required"`

	// 使用cpcode时必填CloudAccountId
	CloudAccountId string `form:"CloudAccountId,required" json:"CloudAccountId,required" query:"CloudAccountId,required"`
}

type ResourceStatData struct {
	Name    string            `form:"Name" json:"Name" query:"Name"`
	Metrics []*MetricStatData `form:"Metrics" json:"Metrics" query:"Metrics"`
}

type MetricStatData struct {
	Metric string            `form:"Metric" json:"Metric" query:"Metric"`
	Values []*TimeSeriesData `form:"Values" json:"Values" query:"Values"`
}

type TimeSeriesData struct {
	Timestamp int64   `form:"Timestamp" json:"Timestamp" query:"Timestamp"`
	Value     float64 `form:"Value" json:"Value" query:"Value"`
}

type DescribeCdnDataOfflineRequest struct {

	// 查询起始时间，相对于UTC 1970-01-01到当前时间相隔的秒数
	StartTime int64 `form:"StartTime,required" json:"StartTime,required" query:"StartTime,required"`

	// 查询结束时间，相对于UTC 1970-01-01到当前时间相隔的秒数。查询区间为前闭后开[StartTime, EndTime)，时间按照传入的Interval向前规整，如Interval为5min，1644163200和1644163499都会规整为1644163200
	EndTime int64 `form:"EndTime,required" json:"EndTime,required" query:"EndTime,required"`

	// 查询指标，仅支持单个，支持以下几种指标类型：1. flux：流量（单位为 byte），2. bandwidth：带宽（单位为 bps），3. request: 请求数（单位为 次），4. status_all：状态码，返回 2xx、3xx、4xx、5xx 汇总数据（单位为 个），5. status_2xx：返回 2xx 状态码汇总及各 2 开头状态码数据（单位为 个），6. status_3xx：返回 3xx 状态码汇总及各 3 开头状态码数据（单位为 个），7. status_4xx：返回 4xx 状态码汇总及各 4 开头状态码数据（单位为 个），8. status_5xx：返回 5xx 状态码汇总及各 5 开头状态码数据（单位为 个），9. hitrate: 流量命中率（单位为 %，小数点后保留两位），10. request_hitrate：请求命中率（单位为 %，小数点后保留两位）。
	Metric string `form:"Metric,required" json:"Metric,required" query:"Metric,required"`

	// 时间粒度，支持以下几种取值：1. 1min：1分钟粒度，指定查询区间1天内，2. 5min：5分钟粒度，指定查询区间31天内，3. hour：1小时粒度，指定查询区间90天内，4. day：天粒度，指定查询区间90天内
	Interval string `form:"Interval,required" json:"Interval,required" query:"Interval,required"`

	// 厂商名称
	Vendors []string `form:"Vendors" json:"Vendors,omitempty" query:"Vendors"`

	// 厂商id
	CloudAccountIds []string `form:"CloudAccountIds" json:"CloudAccountIds,omitempty" query:"CloudAccountIds"`

	// 子产品
	SubProducts []string `form:"SubProducts" json:"SubProducts,omitempty" query:"SubProducts"`

	// cdn_type
	CdnTypes []string `form:"CdnTypes" json:"CdnTypes,omitempty" query:"CdnTypes"`

	// 指定查询域名，最多50个
	Domains []string `form:"Domains" json:"Domains,omitempty" query:"Domains"`

	// 域名id，最多50个
	DomainIds []string `form:"DomainIds" json:"DomainIds,omitempty" query:"DomainIds"`

	// 标签key以及对应的标签value列表，value列表长度最大20，key必须以STAT_i_开头(i取值1-4)，如STAT_1_K1
	TagKVs map[string][]string `form:"TagKVs" json:"TagKVs,omitempty" query:"TagKVs"`

	// 不传返回汇总数据，另外支持以下几种取值：1. vendor：返回分厂商数据，2. cloud_account_id：返回分厂商id数据，3. sub_product：返回分子产品数据，4. cdn_type：返回分CdnType数据，5. domain：返回分域名数据(仅在Domain参数非空时支持)，6. domain_id：返回分域名id数据(仅在DomainId参数非空时支持)，7. TagKVs中的key，如STAT_1_K1：返回分tag数据(仅在TagKVs参数中对应key的value列表非空时支持)。
	GroupBy *string `form:"GroupBy" json:"GroupBy,omitempty" query:"GroupBy"`
}

type DescribeCdnDataOfflineResponse struct {

	// 查询得到的数据明细
	Resources []*ResourceStatData `form:"Resources" json:"Resources" query:"Resources"`
}

type ListViewsResponse struct {

	// 国家及地区列表
	Countries []*ViewCountry `form:"Countries,required" json:"Countries,required" query:"Countries,required"`

	// 当前数据版本号，格式：支持区域版本@地理数据版本
	Version string `form:"Version,required" json:"Version,required" query:"Version,required"`
}

type ViewCountry struct {

	// 国家及地区id
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 名称
	Name string `form:"Name,required" json:"Name,required" query:"Name,required"`

	// 代码
	Code string `form:"Code,required" json:"Code,required" query:"Code,required"`

	// 运营商列表
	Isps []*ViewIsp `form:"Isps" json:"Isps" query:"Isps"`

	// 省份列表
	Provinces []*ViewProvince `form:"Provinces" json:"Provinces" query:"Provinces"`
}

type ViewIsp struct {

	// 运营商id
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 运营商名称
	Name string `form:"Name,required" json:"Name,required" query:"Name,required"`
}

type ViewProvince struct {

	// 省份id
	Id string `form:"Id,required" json:"Id,required" query:"Id,required"`

	// 名称
	Name string `form:"Name,required" json:"Name,required" query:"Name,required"`

	// 代码
	Code string `form:"Code,required" json:"Code,required" query:"Code,required"`
}

type DescribeCdnRegionAndIspResponse struct {

	// 运营商列表
	Isps []*NamePair `form:"Isps" json:"Isps" query:"Isps"`

	// 国家列表
	Countries []*Country `form:"Countries" json:"Countries" query:"Countries"`
}

type NamePair struct {

	// 中文名称
	CnName string `form:"CnName" json:"CnName" query:"CnName"`

	// 英文名称
	EnName string `form:"EnName" json:"EnName" query:"EnName"`
}

type Country struct {

	// 国家
	NamePair *NamePair `form:"NamePair" json:"NamePair" query:"NamePair"`

	// 地区列表
	Regions []*NamePair `form:"Regions" json:"Regions" query:"Regions"`
}
