package ipaas

type ArrayItemschema int32

type ControlApplicationBodyOperation string

type CreateDevicesBodyNetConfigISP int32

type DetailInstanceResResultIsp int32

type DetailInstanceResResultSecurityGroupSecurityRuleListItemExpose int32

type DetailInstanceResResultSecurityGroupSecurityRuleListItemNatType int32

type DetailInstanceResResultSecurityGroupSecurityRuleListItemProtocol int32

type DetailInstanceResResultSgBoundSt int32

type DetailInstanceResResultStatus int32

type DetailSecurityGroupResResultSecurityRuleListItemExpose int32

type DetailSecurityGroupResResultSecurityRuleListItemNatType int32

type DetailSecurityGroupResResultSecurityRuleListItemProtocol int32

type DistributeFileBodyAfterDistributionActionAction int32

type DistributeFileBodyFileFormat string

type DistributeFileToInstancesBodyAfterDistributionActionAction int32

type DistributeFileToInstancesBodyFileFormat string

type Enum0 int32

type Enum1 int32

type Enum13 string

type Enum14 int32

type Enum15 int32

type Enum16 int32

type Enum17 int32

type Enum21 int32

type Enum22 int32

type Enum23 int32

type Enum35 string

type Enum36 int32

type Enum37 int32

type Enum38 string

type Enum39 int32

type Enum40 int32

type Enum41 int32

type Enum42 string

type Enum43 string

type Enum44 string

type Enum45 string

type Enum8 int32

type Enum9 int32

type ExportInstanceQueryIsp int32

type ExportInstanceQuerySgBoundSt int32

type GetFileDistributionJobDetailResResultJobStatus int32

type ImportContainerImageBodyImageFileType string

type ListContainerImagesResResultRowItemStatusCode int32

type ListDcCapacityBodyIsp int32

type ListDcCapacityResResultItemIsp int32

type ListHostMetricDataBodyMetricNameItem string

type ListHostMetricDataBodyMetricType string

type ListHostMetricDataResResultMetricNameItem string

type ListHostMetricDataResResultMetricType string

type ListHostQueryIsp int32

type ListHostQueryOrderBy string

type ListHostResResultRowItemIsp int32

type ListHostResResultRowItemStatus int32

type ListInstanceMetricDataBodyMetricNameItem string

type ListInstanceMetricDataBodyMetricType string

type ListInstanceMetricDataResResultMetricNameItem string

type ListInstanceMetricDataResResultMetricType string

type ListInstanceQueryIsp int32

type ListInstanceQuerySgBoundSt int32

type ListInstanceResResultRowItemIsp int32

type ListInstanceResResultRowItemSecurityGroupSecurityRuleListItemExpose int32

type ListInstanceResResultRowItemSecurityGroupSecurityRuleListItemNatType int32

type ListInstanceResResultRowItemSecurityGroupSecurityRuleListItemProtocol int32

type ListInstanceResResultRowItemSgBoundSt int32

type ListInstanceResResultRowItemStatus int32

type ListPortMappingQueryIsp int32

type ListPortMappingQueryProtocolEnum int32

type ListPortMappingQueryState int32

type ListPortMappingResResultRowItemIsp int32

type ListPortMappingResResultRowItemProtocol string

type ListPortMappingResResultRowItemProtocolEnum int32

type ListPortMappingResResultRowItemState int32

type ListProductResResultRowItemProductType int32

type ListSecurityGroupResResultRowItemSecurityRuleListItemExpose int32

type ListSecurityGroupResResultRowItemSecurityRuleListItemNatType int32

type ListSecurityGroupResResultRowItemSecurityRuleListItemProtocol int32

type ModifyInstanceWindowDisplaySpecBodyResolutionLevel string

type RecordScreenBodyOption string
type AcquireIdempotentTokenBody struct {
	// 令牌有效期, 单位: 秒, 最大值: 120
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty"`
}

type AcquireIdempotentTokenRes struct {
	// REQUIRED
	ResponseMetadata AcquireIdempotentTokenResResponseMetadata `json:"ResponseMetadata"`
	Result           *AcquireIdempotentTokenResResult          `json:"Result,omitempty"`
}

type AcquireIdempotentTokenResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *AcquireIdempotentTokenResResponseMetadataError `json:"Error,omitempty"`
}

type AcquireIdempotentTokenResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type AcquireIdempotentTokenResResult struct {
	// 令牌
	Token *string `json:"Token,omitempty"`
}

type AdbCommandBody struct {
	// REQUIRED; 执行的命令
	Command string `json:"command"`

	// REQUIRED; 目标 Id 列表
	ObjectIDList []string `json:"object_id_list"`

	// REQUIRED; 对象的类别，可选枚举值包括： instance（对指定实例为对象进行操作） host（对云机下的所有实例进行操作)
	ObjectType string `json:"object_type"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`
}

type AdbCommandRes struct {
	// REQUIRED
	ResponseMetadata AdbCommandResResponseMetadata `json:"ResponseMetadata"`
	Result           *AdbCommandResResult          `json:"Result,omitempty"`
}

type AdbCommandResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *AdbCommandResResponseMetadataError `json:"Error,omitempty"`
}

type AdbCommandResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type AdbCommandResResult struct {
	// 失败的ID列表
	FailedIDList []*AdbCommandResResultFailedIDListItem `json:"failed_id_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 废弃
	PassedIDList []*string `json:"passed_id_list,omitempty"`
}

type AdbCommandResResultFailedIDListItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type BindInstanceAdbKeyBody struct {
	// REQUIRED
	AdbKeyID int32 `json:"adb_key_id"`

	// REQUIRED
	InstanceID string `json:"instance_id"`

	// REQUIRED
	ProductID string `json:"product_id"`
}

type BindInstanceAdbKeyRes struct {
	// REQUIRED
	ResponseMetadata BindInstanceAdbKeyResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                               `json:"Result,omitempty"`
}

type BindInstanceAdbKeyResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *BindInstanceAdbKeyResResponseMetadataError `json:"Error,omitempty"`
}

type BindInstanceAdbKeyResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type BindInstancesAdbKeyBody struct {
	// REQUIRED
	AdbKeyID int32 `json:"adb_key_id"`

	// REQUIRED
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED
	ProductID string `json:"product_id"`
}

type BindInstancesAdbKeyRes struct {
	// REQUIRED
	ResponseMetadata BindInstancesAdbKeyResResponseMetadata `json:"ResponseMetadata"`
	Result           *BindInstancesAdbKeyResResult          `json:"Result,omitempty"`
}

type BindInstancesAdbKeyResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *BindInstancesAdbKeyResResponseMetadataError `json:"Error,omitempty"`
}

type BindInstancesAdbKeyResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type BindInstancesAdbKeyResResult struct {
	JobID *string `json:"job_id,omitempty"`
}

type BindInstancesSecurityGroupBody struct {
	// REQUIRED; 实例 ID 列表
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED; 业务 ID
	ProductID string `json:"product_id"`

	// REQUIRED; 安全组 ID
	SecurityGroupID int32 `json:"security_group_id"`
}

type BindInstancesSecurityGroupRes struct {
	// REQUIRED
	ResponseMetadata BindInstancesSecurityGroupResResponseMetadata `json:"ResponseMetadata"`
	Result           *BindInstancesSecurityGroupResResult          `json:"Result,omitempty"`
}

type BindInstancesSecurityGroupResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                              `json:"Version"`
	Error   *BindInstancesSecurityGroupResResponseMetadataError `json:"Error,omitempty"`
}

type BindInstancesSecurityGroupResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type BindInstancesSecurityGroupResResult struct {
	// 异步 JobID
	JobID *string `json:"job_id,omitempty"`
}

type ColdRebootInstanceBody struct {
	// REQUIRED; 实例 Id 列表
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED; 产品 Id
	ProductID string `json:"product_id"`
}

type ColdRebootInstanceRes struct {
	// REQUIRED
	ResponseMetadata ColdRebootInstanceResResponseMetadata `json:"ResponseMetadata"`
	Result           *ColdRebootInstanceResResult          `json:"Result,omitempty"`
}

type ColdRebootInstanceResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *ColdRebootInstanceResResponseMetadataError `json:"Error,omitempty"`
}

type ColdRebootInstanceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ColdRebootInstanceResResult struct {
	// Deprecated
	FailIDList []*string `json:"fail_id_list,omitempty"`

	// 失败的记录
	FailedList []*ColdRebootInstanceResResultFailedListItem `json:"failed_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated
	SuccessIDList []*string `json:"success_id_list,omitempty"`
}

type ColdRebootInstanceResResultFailedListItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type ControlApplicationBody struct {
	// REQUIRED; 目标 Id 列表
	ObjectIDList []string `json:"object_id_list"`

	// REQUIRED; 对象的类别，可选枚举值包括： instance（对指定实例为对象进行操作） host（对云机下的所有实例进行操作)
	ObjectType string `json:"object_type"`

	// REQUIRED
	Operation ControlApplicationBodyOperation `json:"operation"`

	// REQUIRED; 应用包名，例如：com.abc.xyz uninstall, 卸载, 执行pm uninstall [packagename], e.g. com.abc.xyz enable, 启用, 执行pm enable [packagename],
	// e.g. com.abc.xyz disable, 停用, 执行pm disable-user [packagename], e.g. com.abc.xyz
	// stop, 停止, 执行am force-stop [packagename], e.g. com.abc.xyz start, 启动, 执行am start [-n] [package_name], 支持传入或者不传入具体的 activity
	// case 1: 不传入 activity, 例如 com.abc.xyz, 启动 [android.intent.action.MAIN]
	// activity, 如果 apk 不存在该 activity, 则报错. case 2: 传入 activity 时, 例如 com.abc.xyz/.MainActivity. 举例来说, 对于有前台 UI 的应用, 可以在云手机内部,
	// 启动应用后, 命令行输入 dumpsys activity |grep -i mResumedActivity | awk -F " " '{print
	// $4}' 来确认具体的字段内容
	PackageName string `json:"package_name"`

	// REQUIRED; 产品ID`
	ProductID string `json:"product_id"`
}

type ControlApplicationRes struct {
	// REQUIRED
	ResponseMetadata ControlApplicationResResponseMetadata `json:"ResponseMetadata"`
	Result           *ControlApplicationResResult          `json:"Result,omitempty"`
}

type ControlApplicationResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *ControlApplicationResResponseMetadataError `json:"Error,omitempty"`
}

type ControlApplicationResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ControlApplicationResResult struct {
	// 失败的ID列表
	FailedIDList []*ControlApplicationResResultFailedIDListItem `json:"failed_id_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 废弃
	PassedIDList []*string `json:"passed_id_list,omitempty"`
}

type ControlApplicationResResultFailedIDListItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type CreateDevicesBody struct {
	// REQUIRED
	DeviceCharge CreateDevicesBodyDeviceCharge `json:"DeviceCharge"`

	// REQUIRED; 需要订购的设备数量
	DeviceCount int32 `json:"DeviceCount"`

	// REQUIRED; 套餐规格
	DevicePackageID string `json:"DevicePackageId"`

	// REQUIRED; container容器订购，vm虚拟机订购， bm 裸node订购
	DeviceType string `json:"DeviceType"`

	// REQUIRED
	ImageConfig CreateDevicesBodyImageConfig `json:"ImageConfig"`

	// REQUIRED
	NetConfig CreateDevicesBodyNetConfig `json:"NetConfig"`

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`

	// 机房列表,region和dc至少一个不为空
	DC []*string `json:"DC,omitempty"`

	// 设备名称，不填默认和deviceId保持一致，若订购多个，会加后缀递增
	DeviceName *string `json:"DeviceName,omitempty"`

	InstanceDisk *int32 `json:"InstanceDisk,omitempty"` // 共享存储，MB，新增参数，针对云游戏做的特殊逻辑

	// 订购设备的区域，cn-north,cn-middle,cn-south
	Region *string `json:"Region,omitempty"`
}

type CreateDevicesBodyDeviceCharge struct {
	// REQUIRED
	DeviceChargeMode CreateDevicesBodyDeviceChargeMode `json:"DeviceChargeMode"`

	// REQUIRED
	NetworkChargeMode CreateDevicesBodyDeviceChargeNetworkChargeMode `json:"NetworkChargeMode"`
}

type CreateDevicesBodyDeviceChargeMode struct {
	// REQUIRED; 云机计费模式，选项： 1：包年包月 2：按量计费（按天） 3: 按量计费（按月）
	DeviceChargeType int32 `json:"DeviceChargeType"`

	// 订购周期数，后付费模式须为0，会自动校正
	PeriodTerm *int32 `json:"PeriodTerm,omitempty"`

	// 1:手动续费 2:自动续费 3:到期不续费
	RenewType *int32 `json:"RenewType,omitempty"`
}

type CreateDevicesBodyDeviceChargeNetworkChargeMode struct {
	// REQUIRED; 云机对应的带宽计费方式，选项： dailypeak：按日带宽峰值计费 traffic：按实际流量计费 95thpercentile：按带宽95峰计费 bandwidth: 按带宽上线计费
	BandWidthChargeType string `json:"BandWidthChargeType"`
}

type CreateDevicesBodyImageConfig struct {
	// 镜像Id, container订购时表示aosp镜像(必填) bm订购时表示debian镜像
	ImageID *string `json:"ImageId,omitempty"`

	// aosp镜像是否为公共镜像 true（在公共镜像下检索指定的镜像进行重置） false（在当前账号下检索指定的镜像进行重置，默认）
	IsPublicImage *bool `json:"IsPublicImage,omitempty"`
}

type CreateDevicesBodyNetConfig struct {
	// REQUIRED
	ISP CreateDevicesBodyNetConfigISP `json:"ISP"`

	// 带宽，2-100, 不填默认5Mbps
	Bandwidth *int32 `json:"Bandwidth,omitempty"`

	// nat配置，1表示云上，2表示云下，不填由服务决定
	NatID *int32 `json:"NatId,omitempty"`
}

type CreateDevicesQuery struct {
	// REQUIRED; X-iPaaS-Idempotent-Token
	XIPaaSIdempotentToken string `json:"X-iPaaS-Idempotent-Token" query:"X-iPaaS-Idempotent-Token"`
}

type CreateDevicesRes struct {
	// REQUIRED
	ResponseMetadata CreateDevicesResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateDevicesResResult          `json:"Result,omitempty"`
}

type CreateDevicesResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *CreateDevicesResResponseMetadataError `json:"Error,omitempty"`
}

type CreateDevicesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type CreateDevicesResResult struct {
	InstanceNo []*string `json:"InstanceNo,omitempty"`
	OrderNo    *string   `json:"OrderNo,omitempty"`
}

type DeleteContainerImagesBody struct {
	// REQUIRED; 待删除的镜像 ID 列表
	ImageIDList []string `json:"image_id_list"`
}

type DeleteContainerImagesRes struct {
	// REQUIRED
	ResponseMetadata DeleteContainerImagesResResponseMetadata `json:"ResponseMetadata"`
	Result           *DeleteContainerImagesResResult          `json:"Result,omitempty"`
}

type DeleteContainerImagesResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *DeleteContainerImagesResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteContainerImagesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type DeleteContainerImagesResResult struct {
	// 删除失败的镜像ID列表
	FailedList []*DeleteContainerImagesResResultFailedListItem `json:"failed_list,omitempty"`
}

type DeleteContainerImagesResResultFailedListItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type DeleteDevicesBody struct {
	// REQUIRED; 退订配置信息
	DeviceIDList []string `json:"DeviceIdList"`

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`
}

type DeleteDevicesQuery struct {
	// REQUIRED; X-iPaaS-Idempotent-Token
	XIPaaSIdempotentToken string `json:"X-iPaaS-Idempotent-Token" query:"X-iPaaS-Idempotent-Token"`
}

type DeleteDevicesRes struct {
	// REQUIRED
	ResponseMetadata DeleteDevicesResResponseMetadata `json:"ResponseMetadata"`
	Result           *DeleteDevicesResResult          `json:"Result,omitempty"`
}

type DeleteDevicesResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *DeleteDevicesResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteDevicesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type DeleteDevicesResResult struct {
	InstanceNo []*string `json:"InstanceNo,omitempty"`
	OrderNo    *string   `json:"OrderNo,omitempty"`
}

type DetailInstanceQuery struct {
	// REQUIRED; 实例 Id
	InstanceID string `json:"instance_id" query:"instance_id"`
}

type DetailInstanceRes struct {
	// REQUIRED
	ResponseMetadata DetailInstanceResResponseMetadata `json:"ResponseMetadata"`
	Result           *DetailInstanceResResult          `json:"Result,omitempty"`
}

type DetailInstanceResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *DetailInstanceResResponseMetadataError `json:"Error,omitempty"`
}

type DetailInstanceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type DetailInstanceResResult struct {
	AdbKey *DetailInstanceResResultAdbKey `json:"adb_key,omitempty"`

	// 带宽
	Bandwidth     *int32                                `json:"bandwidth,omitempty"`
	Configuration *DetailInstanceResResultConfiguration `json:"configuration,omitempty"`

	// 创建时间, unix 时间戳, 秒级
	CreateAt *int32 `json:"create_at,omitempty"`

	// 机房 ID
	Dc *string `json:"dc,omitempty"`

	// 机房名称
	DcName *string `json:"dc_name,omitempty"`

	// 帧率
	Fps *int32 `json:"fps,omitempty"`

	// 主机ID
	HostID *string `json:"host_id,omitempty"`

	// 镜像ID
	ImageID *string `json:"image_id,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 实例ID
	InstanceID *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string                     `json:"instance_name,omitempty"`
	Isp          *DetailInstanceResResultIsp `json:"isp,omitempty"`

	// 产品ID
	ProductID *string `json:"product_id,omitempty"`

	// 地域
	Region *string `json:"region,omitempty"`

	// 分辨率
	Resolution    *string                               `json:"resolution,omitempty"`
	SecurityGroup *DetailInstanceResResultSecurityGroup `json:"security_group,omitempty"`
	SgBoundSt     *DetailInstanceResResultSgBoundSt     `json:"sg_bound_st,omitempty"`

	// 安全组绑定状态字符串
	SgBoundStStr *string `json:"sg_bound_st_str,omitempty"`

	// 序列号, maybe useless
	Sn     *string                        `json:"sn,omitempty"`
	Status *DetailInstanceResResultStatus `json:"status,omitempty"`

	// 实例状态字符串
	StatusStr *string                     `json:"status_str,omitempty"`
	Tag       *DetailInstanceResResultTag `json:"tag,omitempty"`
}

type DetailInstanceResResultAdbKey struct {
	// 用户权限类型：
	// 1（root） 2（user）
	AuthType *int32 `json:"auth_type,omitempty"`

	// 密钥对绑定的实例数量
	BindInstanceNum *int32 `json:"bind_instance_num,omitempty"`
	CreateAt        *int32 `json:"create_at,omitempty"`

	// 公钥指纹
	Fingerprint *string `json:"fingerprint,omitempty"`
	KeyDesc     *string `json:"key_desc,omitempty"`
	KeyID       *int32  `json:"key_id,omitempty"`
	KeyName     *string `json:"key_name,omitempty"`
	ProductID   *string `json:"product_id,omitempty"`
	PublicKey   *string `json:"public_key,omitempty"`
}

type DetailInstanceResResultConfiguration struct {
	// CPU 核心数
	CPUCore *int32 `json:"cpu_core,omitempty"`

	// 实例资源套餐 ID
	ConfigurationCode *string `json:"configuration_code,omitempty"`

	// 实例资源套餐名称
	ConfigurationName *string `json:"configuration_name,omitempty"`

	// 网络计费名称
	IspCodeName *string `json:"isp_code_name,omitempty"`

	// 网络计费套餐
	IspConfigurationCode *string `json:"isp_configuration_code,omitempty"`

	// 内存，单位MB
	Memory *float32 `json:"memory,omitempty"`
}

type DetailInstanceResResultSecurityGroup struct {
	// BindHostNum *int64 json:"bind_host_num,omitempty" // Deprecated: 请使用BindInstanceNum
	BindInstanceNum *int32 `json:"bind_instance_num,omitempty"`

	// 创建时间，秒级时间戳
	CreateAt *int32 `json:"create_at,omitempty"`

	// 安全组所属业务 ID
	ProductID *string `json:"product_id,omitempty"`

	// 安全组描述
	SecurityGroupDesc *string `json:"security_group_desc,omitempty"`

	// 安全组 ID
	SecurityGroupID *int32 `json:"security_group_id,omitempty"`

	// 安全组名称
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 安全组规则列表
	SecurityRuleList []*DetailInstanceResResultSecurityGroupSecurityRuleListItem `json:"security_rule_list,omitempty"`

	// 更新时间，秒级时间戳
	UpdateAt *int32 `json:"update_at,omitempty"`
}

type DetailInstanceResResultSecurityGroupSecurityRuleListItem struct {
	Expose   *DetailInstanceResResultSecurityGroupSecurityRuleListItemExpose   `json:"expose,omitempty"`
	NatType  *DetailInstanceResResultSecurityGroupSecurityRuleListItemNatType  `json:"nat_type,omitempty"`
	Protocol *DetailInstanceResResultSecurityGroupSecurityRuleListItemProtocol `json:"protocol,omitempty"`

	// 安全组规则ID
	RuleID *int32 `json:"rule_id,omitempty"`

	// 源端口
	SourcePort *int32 `json:"source_port,omitempty"`
}

type DetailInstanceResResultTag struct {
	ProductID          *string `json:"product_id,omitempty"`
	RelatedInstanceNum *int32  `json:"related_instance_num,omitempty"`
	TagDesc            *string `json:"tag_desc,omitempty"`
	TagID              *string `json:"tag_id,omitempty"`
	TagName            *string `json:"tag_name,omitempty"`
}

type DetailSecurityGroupQuery struct {
	// REQUIRED; 安全组所属业务 ID
	ProductID string `json:"product_id" query:"product_id"`

	// REQUIRED; 安全组 ID
	SecurityGroupID int32 `json:"security_group_id" query:"security_group_id"`

	// 是否展示绑定实例数量
	IsShowBoundInsNum *bool `json:"is_show_bound_ins_num,omitempty" query:"is_show_bound_ins_num"`
}

type DetailSecurityGroupRes struct {
	// REQUIRED
	ResponseMetadata DetailSecurityGroupResResponseMetadata `json:"ResponseMetadata"`
	Result           *DetailSecurityGroupResResult          `json:"Result,omitempty"`
}

type DetailSecurityGroupResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *DetailSecurityGroupResResponseMetadataError `json:"Error,omitempty"`
}

type DetailSecurityGroupResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type DetailSecurityGroupResResult struct {
	// BindHostNum *int64 json:"bind_host_num,omitempty" // Deprecated: 请使用BindInstanceNum
	BindInstanceNum *int32 `json:"bind_instance_num,omitempty"`

	// 创建时间，秒级时间戳
	CreateAt *int32 `json:"create_at,omitempty"`

	// 安全组所属业务 ID
	ProductID *string `json:"product_id,omitempty"`

	// 安全组描述
	SecurityGroupDesc *string `json:"security_group_desc,omitempty"`

	// 安全组 ID
	SecurityGroupID *int32 `json:"security_group_id,omitempty"`

	// 安全组名称
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 安全组规则列表
	SecurityRuleList []*DetailSecurityGroupResResultSecurityRuleListItem `json:"security_rule_list,omitempty"`

	// 更新时间，秒级时间戳
	UpdateAt *int32 `json:"update_at,omitempty"`
}

type DetailSecurityGroupResResultSecurityRuleListItem struct {
	Expose   *DetailSecurityGroupResResultSecurityRuleListItemExpose   `json:"expose,omitempty"`
	NatType  *DetailSecurityGroupResResultSecurityRuleListItemNatType  `json:"nat_type,omitempty"`
	Protocol *DetailSecurityGroupResResultSecurityRuleListItemProtocol `json:"protocol,omitempty"`

	// 安全组规则ID
	RuleID *int32 `json:"rule_id,omitempty"`

	// 源端口
	SourcePort *int32 `json:"source_port,omitempty"`
}

type DistributeFileBody struct {
	// REQUIRED; 文件的 MD5 值
	FileMD5 string `json:"file_md5"`

	// REQUIRED; 实例ID
	InstanceID string `json:"instance_id"`

	// REQUIRED; 产品ID
	ProductID               string                                     `json:"product_id"`
	AfterDistributionAction *DistributeFileBodyAfterDistributionAction `json:"after_distribution_action,omitempty"`

	// 在云机中存储文件的路径（默认：/data/file_ds）
	FileDir     *string                        `json:"file_dir,omitempty"`
	FileFormat  *DistributeFileBodyFileFormat  `json:"file_format,omitempty"`
	URLFile     *DistributeFileBodyURLFile     `json:"url_file,omitempty"`
	VolcTosFile *DistributeFileBodyVolcTosFile `json:"volc_tos_file,omitempty"`
}

type DistributeFileBodyAfterDistributionAction struct {
	// REQUIRED
	Action          DistributeFileBodyAfterDistributionActionAction           `json:"action"`
	InstallApkParam *DistributeFileBodyAfterDistributionActionInstallApkParam `json:"install_apk_param,omitempty"`
}

type DistributeFileBodyAfterDistributionActionInstallApkParam struct {
	// 应用安装可选参数列表： 1（ApkInstallAllowTest，允许测试） 2（ApkInstallReplaceExisting，替换现存） 3（ApkInstallGrantAllPerm，获取全部许可） 4（ApkInstallABI，覆盖默认
	// ABI） 5（ApkInstallInternalFlash，内部闪存） 6（ApkInstallAllowDowngrade，允许降级）
	OptionList []*ArrayItemschema `json:"option_list,omitempty"`
}

type DistributeFileBodyURLFile struct {
	// REQUIRED; 保存文件的 URL
	URL string `json:"url"`
}

type DistributeFileBodyVolcTosFile struct {
	// REQUIRED; 火山引擎对象存储中的存储桶名称
	TosBucket string `json:"tos_bucket"`

	// REQUIRED; 火山引擎对象存储中的文件路径
	TosFilePath string `json:"tos_file_path"`

	// 火山引擎对象存储服务地址（地域节点），若为空，则使用默认值：tos-cn-beijing.volces.com 默认 constdef.TosEndpointCNBJOnline
	Endpoint *string `json:"endpoint,omitempty"`

	// 火山引擎对象存储服务区域，若为空，则使用默认值：cn-beijing 默认 constdef.TosRegionBJ
	Region *string `json:"region,omitempty"`
}

type DistributeFileRes struct {
	// REQUIRED
	ResponseMetadata DistributeFileResResponseMetadata `json:"ResponseMetadata"`
	Result           *DistributeFileResResult          `json:"Result,omitempty"`
}

type DistributeFileResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *DistributeFileResResponseMetadataError `json:"Error,omitempty"`
}

type DistributeFileResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type DistributeFileResResult struct {
	// 任务ID
	DistributionTaskID *string `json:"distribution_task_id,omitempty"`
}

type DistributeFileToInstancesBody struct {
	// REQUIRED; 文件的 MD5 值
	FileMD5 string `json:"file_md5"`

	// REQUIRED; 实例ID列表
	InstanceIDs []string `json:"instance_ids"`

	// REQUIRED; 产品ID
	ProductID               string                                                `json:"product_id"`
	AfterDistributionAction *DistributeFileToInstancesBodyAfterDistributionAction `json:"after_distribution_action,omitempty"`

	// 在云机中存储文件的路径（默认：/data/file_ds）
	FileDir     *string                                   `json:"file_dir,omitempty"`
	FileFormat  *DistributeFileToInstancesBodyFileFormat  `json:"file_format,omitempty"`
	URLFile     *DistributeFileToInstancesBodyURLFile     `json:"url_file,omitempty"`
	VolcTosFile *DistributeFileToInstancesBodyVolcTosFile `json:"volc_tos_file,omitempty"`
}

type DistributeFileToInstancesBodyAfterDistributionAction struct {
	// REQUIRED
	Action          DistributeFileToInstancesBodyAfterDistributionActionAction           `json:"action"`
	InstallApkParam *DistributeFileToInstancesBodyAfterDistributionActionInstallApkParam `json:"install_apk_param,omitempty"`
}

type DistributeFileToInstancesBodyAfterDistributionActionInstallApkParam struct {
	// 应用安装可选参数列表： 1（ApkInstallAllowTest，允许测试） 2（ApkInstallReplaceExisting，替换现存） 3（ApkInstallGrantAllPerm，获取全部许可） 4（ApkInstallABI，覆盖默认
	// ABI） 5（ApkInstallInternalFlash，内部闪存） 6（ApkInstallAllowDowngrade，允许降级）
	OptionList []*ArrayItemschema `json:"option_list,omitempty"`
}

type DistributeFileToInstancesBodyURLFile struct {
	// REQUIRED; 保存文件的 URL
	URL string `json:"url"`
}

type DistributeFileToInstancesBodyVolcTosFile struct {
	// REQUIRED; 火山引擎对象存储中的存储桶名称
	TosBucket string `json:"tos_bucket"`

	// REQUIRED; 火山引擎对象存储中的文件路径
	TosFilePath string `json:"tos_file_path"`

	// 火山引擎对象存储服务地址（地域节点），若为空，则使用默认值：tos-cn-beijing.volces.com 默认 constdef.TosEndpointCNBJOnline
	Endpoint *string `json:"endpoint,omitempty"`

	// 火山引擎对象存储服务区域，若为空，则使用默认值：cn-beijing 默认 constdef.TosRegionBJ
	Region *string `json:"region,omitempty"`
}

type DistributeFileToInstancesRes struct {
	// REQUIRED
	ResponseMetadata DistributeFileToInstancesResResponseMetadata `json:"ResponseMetadata"`
	Result           *DistributeFileToInstancesResResult          `json:"Result,omitempty"`
}

type DistributeFileToInstancesResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                             `json:"Version"`
	Error   *DistributeFileToInstancesResResponseMetadataError `json:"Error,omitempty"`
}

type DistributeFileToInstancesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type DistributeFileToInstancesResResult struct {
	// 批任务ID
	DistributionJobID *string `json:"distribution_job_id,omitempty"`

	// 失败记录
	FailedRecords []*DistributeFileToInstancesResResultFailedRecordsItem `json:"failed_records,omitempty"`

	// 任务ID列表
	FileDistributionTaskMap map[string]*string `json:"file_distribution_task_map,omitempty"`
}

type DistributeFileToInstancesResResultFailedRecordsItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type ExecCmdSyncBody struct {
	// REQUIRED; 执行的命令
	Cmd string `json:"cmd"`

	// REQUIRED; 实例ID
	InstanceID string `json:"instance_id"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`
}

type ExecCmdSyncRes struct {
	// REQUIRED
	ResponseMetadata ExecCmdSyncResResponseMetadata `json:"ResponseMetadata"`
	Result           *ExecCmdSyncResResult          `json:"Result,omitempty"`
}

type ExecCmdSyncResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                               `json:"Version"`
	Error   *ExecCmdSyncResResponseMetadataError `json:"Error,omitempty"`
}

type ExecCmdSyncResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ExecCmdSyncResResult struct {
	// 状态码
	Code *int32 `json:"code,omitempty"`

	// 执行结果输出
	Msg *string `json:"msg,omitempty"`
}

type ExportInstanceQuery struct {
	// REQUIRED; 产品 ID
	ProductID string `json:"product_id" query:"product_id"`

	// key ID
	AdbKeyID *int32 `json:"adb_key_id,omitempty" query:"adb_key_id"`

	// 返回数量, 默认 100
	Count *int32 `json:"count,omitempty" query:"count"`

	// 创建时间范围, 开始时间, 秒级时间戳, 开区间
	CreateAfter *int32 `json:"create_after,omitempty" query:"create_after"`

	// 创建时间范围, 结束时间, 秒级时间戳, 闭区间
	CreateBefore *int32 `json:"create_before,omitempty" query:"create_before"`

	// 是否返回详细信息, e.g. tag, security group, key, etc.
	Detail *bool `json:"detail,omitempty" query:"detail"`

	// 云机 ID
	HostID *string `json:"host_id,omitempty" query:"host_id"`

	// 机房
	Idc *string `json:"idc,omitempty" query:"idc"`

	// 实例 ID
	InstanceID *string `json:"instance_id,omitempty" query:"instance_id"`

	// 批量筛选, 实例ID, 逗号分隔 string. 内部注释-勿展示: 为了兼容性，优先级比 InstanceIdList 低, 参见 Rectify()
	InstanceIDIn *string `json:"instance_id_in,omitempty" query:"instance_id_in"`

	// 模糊查询, 实例ID
	InstanceIDLike *string `json:"instance_id_like,omitempty" query:"instance_id_like"`

	// 模糊查询, 实例名称
	InstanceNameLike *string `json:"instance_name_like,omitempty" query:"instance_name_like"`

	// 是否升序, 默认降序
	IsOrderAsc *bool `json:"is_order_asc,omitempty" query:"is_order_asc"`

	// 运营商
	Isp *ExportInstanceQueryIsp `json:"isp,omitempty" query:"isp"`

	// 偏移量, 默认 0
	Offset *int32 `json:"offset,omitempty" query:"offset"`

	// 排序字段, 支持 instance_id, sn
	OrderBy *string `json:"order_by,omitempty" query:"order_by"`

	// 套餐 ID
	PackageID *string `json:"package_id,omitempty" query:"package_id"`

	// 安全组绑定状态
	SgBoundSt *ExportInstanceQuerySgBoundSt `json:"sg_bound_st,omitempty" query:"sg_bound_st"`

	// 实例状态
	Status *int32 `json:"status,omitempty" query:"status"`

	// 批量筛选, 状态, 逗号分隔 int
	StatusIn *string `json:"status_in,omitempty" query:"status_in"`

	// 批量筛选, 标签ID, 逗号分隔 string
	TagIDIn *string `json:"tag_id_in,omitempty" query:"tag_id_in"`
}

type FixInstancesSGBoundBody struct {
	// REQUIRED; 实例 ID 列表
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED; 业务 ID
	ProductID string `json:"product_id"`
}

type FixInstancesSGBoundRes struct {
	// REQUIRED
	ResponseMetadata FixInstancesSGBoundResResponseMetadata `json:"ResponseMetadata"`
	Result           *FixInstancesSGBoundResResult          `json:"Result,omitempty"`
}

type FixInstancesSGBoundResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *FixInstancesSGBoundResResponseMetadataError `json:"Error,omitempty"`
}

type FixInstancesSGBoundResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type FixInstancesSGBoundResResult struct {
	// 修复结果
	Data []*FixInstancesSGBoundResResultDataItem `json:"data,omitempty"`
}

type FixInstancesSGBoundResResultDataItem struct {
	// 修复失败的端口转发条目
	FailedEntries []*FixInstancesSGBoundResResultDataPropertiesItemsItem `json:"failed_entries,omitempty"`

	// 正在修复的端口转发条目
	FixingEntries []*string `json:"fixing_entries,omitempty"`

	// 实例ID
	ID *string `json:"id,omitempty"`

	// 错误信息，如果无错误则为空
	Msg *string `json:"msg,omitempty"`
}

type FixInstancesSGBoundResResultDataPropertiesItemsItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type GetFileDistributionJobDetailBody struct {
	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`

	// 文件分发执行ID, 与 JobId 二选一
	DistributionTaskID *string `json:"distribution_task_id,omitempty"`

	// 任务ID
	JobID *string `json:"job_id,omitempty"`
}

type GetFileDistributionJobDetailRes struct {
	// REQUIRED
	ResponseMetadata GetFileDistributionJobDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetFileDistributionJobDetailResResult          `json:"Result,omitempty"`
}

type GetFileDistributionJobDetailResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                `json:"Version"`
	Error   *GetFileDistributionJobDetailResResponseMetadataError `json:"Error,omitempty"`
}

type GetFileDistributionJobDetailResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type GetFileDistributionJobDetailResResult struct {
	// 批量分发 job id
	JobID     *string                                         `json:"job_id,omitempty"`
	JobStatus *GetFileDistributionJobDetailResResultJobStatus `json:"job_status,omitempty"`

	// 批任务状态(string)
	JobStatusStr *string `json:"job_status_str,omitempty"`

	// 任务列表
	Tasks []*GetFileDistributionJobDetailResResultTasksItem `json:"tasks,omitempty"`
}

type GetFileDistributionJobDetailResResultTasksItem struct {
	// 任务创建时间，秒级时间戳，例：1669975702
	CreateAt *int32 `json:"create_at,omitempty"`

	// 分发任务 id
	DistributionTaskID *string `json:"distribution_task_id,omitempty"`

	// 文件保存在实例中的绝对路径，例：/data/file_ds/test.jpg
	FilePath *string `json:"file_path,omitempty"`

	// 实例ID
	InstanceID *string `json:"instance_id,omitempty"`

	// 批量分发 job id, 非批量分发请求时, 该字段为空
	JobID *string `json:"job_id,omitempty"`

	// 任务执行成功或失败的返回信息
	Msg *string `json:"msg,omitempty"`

	// 任务执行状态信息，枚举值如下： Initial（初始化中） DownloadPending（下载任务等待调度） DownloadQueued（下载任务调度中） DownloadRunning（下载任务执行中） DownloadSucceed（下载成功）
	// DownloadFailed（下载任务失败） DistributePending（分发任务等待调度）
	// DistributeQueued（分发任务待调度） DistributeRunning（分发任务执行中） DistributeSucceed（分发成功） DistributeFailed（分发任务失败） AfterDistributionHookPending:
	// 分发后任务待调度 AfterDistributionHookRunning: 分发后任务执行中
	// AfterDistributionHookSucceed: 分发后任务成功 AfterDistributionHookFailed: 分发后任务失败 UnknownErr（未知错误） AllSucceed（全部执行完成）
	Status *string `json:"status,omitempty"`

	// 任务更新时间，秒级时间戳，例：1669975903
	UpdateAt *int32 `json:"update_at,omitempty"`
}

type GetFileDistributionResultQuery struct {
	// REQUIRED; 文件分发执行ID
	DistributionTaskID string `json:"distribution_task_id" query:"distribution_task_id"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id" query:"product_id"`
}

type GetFileDistributionResultRes struct {
	// REQUIRED
	ResponseMetadata GetFileDistributionResultResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetFileDistributionResultResResult          `json:"Result,omitempty"`
}

type GetFileDistributionResultResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                             `json:"Version"`
	Error   *GetFileDistributionResultResResponseMetadataError `json:"Error,omitempty"`
}

type GetFileDistributionResultResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type GetFileDistributionResultResResult struct {
	// 任务创建时间，秒级时间戳，例：1669975702
	CreateAt *int32 `json:"create_at,omitempty"`

	// 分发任务 id
	DistributionTaskID *string `json:"distribution_task_id,omitempty"`

	// 文件保存在实例中的绝对路径，例：/data/file_ds/test.jpg
	FilePath *string `json:"file_path,omitempty"`

	// 实例ID
	InstanceID *string `json:"instance_id,omitempty"`

	// 批量分发 job id, 非批量分发请求时, 该字段为空
	JobID *string `json:"job_id,omitempty"`

	// 任务执行成功或失败的返回信息
	Msg *string `json:"msg,omitempty"`

	// 任务执行状态信息，枚举值如下： Initial（初始化中） DownloadPending（下载任务等待调度） DownloadQueued（下载任务调度中） DownloadRunning（下载任务执行中） DownloadSucceed（下载成功）
	// DownloadFailed（下载任务失败） DistributePending（分发任务等待调度）
	// DistributeQueued（分发任务待调度） DistributeRunning（分发任务执行中） DistributeSucceed（分发成功） DistributeFailed（分发任务失败） AfterDistributionHookPending:
	// 分发后任务待调度 AfterDistributionHookRunning: 分发后任务执行中
	// AfterDistributionHookSucceed: 分发后任务成功 AfterDistributionHookFailed: 分发后任务失败 UnknownErr（未知错误） AllSucceed（全部执行完成）
	Status *string `json:"status,omitempty"`

	// 任务更新时间，秒级时间戳，例：1669975903
	UpdateAt *int32 `json:"update_at,omitempty"`
}

type GetInfoAfterOrderBody struct {
	// REQUIRED
	OrderNo string `json:"OrderNo"`

	// REQUIRED
	ProductID string `json:"ProductId"`
	Count     *int32 `json:"Count,omitempty"`
}

type GetInfoAfterOrderRes struct {
	// REQUIRED
	ResponseMetadata GetInfoAfterOrderResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetInfoAfterOrderResResult          `json:"Result,omitempty"`
}

type GetInfoAfterOrderResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                     `json:"Version"`
	Error   *GetInfoAfterOrderResResponseMetadataError `json:"Error,omitempty"`
}

type GetInfoAfterOrderResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type GetInfoAfterOrderResResult struct {
	HostInfo []*GetInfoAfterOrderResResultHostInfoItem `json:"HostInfo,omitempty"`
}

type GetInfoAfterOrderResResultHostInfoItem struct {
	HostID      *string   `json:"HostId,omitempty"`
	InstanceIDs []*string `json:"InstanceIds,omitempty"`
	InstanceNo  *string   `json:"InstanceNo,omitempty"`
}

type GetInstancePropertiesBody struct {
	// REQUIRED; 实例ID
	InstanceID string `json:"instance_id"`

	// REQUIRED; 属性名列表
	PropertyNames []string `json:"property_names"`
}

type GetInstancePropertiesRes struct {
	// REQUIRED
	ResponseMetadata GetInstancePropertiesResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetInstancePropertiesResResult          `json:"Result,omitempty"`
}

type GetInstancePropertiesResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *GetInstancePropertiesResResponseMetadataError `json:"Error,omitempty"`
}

type GetInstancePropertiesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type GetInstancePropertiesResResult struct {
	// 属性名和属性值
	Properties []*GetInstancePropertiesResResultPropertiesItem `json:"properties,omitempty"`
}

type GetInstancePropertiesResResultPropertiesItem struct {
	// REQUIRED; 属性名 属性名称长度上限64 不支持输入特殊字符，包括但不限于 +- ;&|(){}[]
	PropertyName string `json:"property_name"`

	// REQUIRED; 属性值 属性值长度上限91 不支持输入特殊字符，包括但不限于 +- ;&|(){}[]
	PropertyValue string `json:"property_value"`
}

type GetInstancePropertyBody struct {
	// REQUIRED; 目标 Id 列表
	ObjectIDList []string `json:"object_id_list"`

	// REQUIRED; 对象的类别，可选枚举值包括： instance（对指定实例为对象进行操作） host（对云机下的所有实例进行操作)
	ObjectType string `json:"object_type"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`

	// REQUIRED; 属性名 属性名称长度上限64 不支持输入特殊字符，包括但不限于 +- ;&|(){}[]
	PropertyName string `json:"property_name"`
}

type GetInstancePropertyRes struct {
	// REQUIRED
	ResponseMetadata GetInstancePropertyResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetInstancePropertyResResult          `json:"Result,omitempty"`
}

type GetInstancePropertyResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *GetInstancePropertyResResponseMetadataError `json:"Error,omitempty"`
}

type GetInstancePropertyResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type GetInstancePropertyResResult struct {
	// 失败的ID列表
	FailedIDList []*GetInstancePropertyResResultFailedIDListItem `json:"failed_id_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 废弃
	PassedIDList []*string `json:"passed_id_list,omitempty"`
}

type GetInstancePropertyResResultFailedIDListItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type GetJobDetailsQuery struct {
	// REQUIRED
	JobID string `json:"job_id" query:"job_id"`

	// REQUIRED
	ProductID string `json:"product_id" query:"product_id"`
}

type GetJobDetailsRes struct {
	// REQUIRED
	ResponseMetadata GetJobDetailsResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetJobDetailsResResult          `json:"Result,omitempty"`
}

type GetJobDetailsResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *GetJobDetailsResResponseMetadataError `json:"Error,omitempty"`
}

type GetJobDetailsResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type GetJobDetailsResResult struct {
	CreateTime *int32                         `json:"create_time,omitempty"`
	JobID      *string                        `json:"job_id,omitempty"`
	Process    *GetJobDetailsResResultProcess `json:"process,omitempty"`
	ProductID  *string                        `json:"product_id,omitempty"`

	// 任务状态：
	// 100：执行中 200：成功 500：全部失败 501：部分失败 502：取消
	Status       *int32                                    `json:"status,omitempty"`
	StatusStr    *string                                   `json:"status_str,omitempty"`
	TaskInfoList []*GetJobDetailsResResultTaskInfoListItem `json:"task_info_list,omitempty"`
	UpdateTime   *int32                                    `json:"update_time,omitempty"`
}

type GetJobDetailsResResultProcess struct {
	FailTaskNum    *int32 `json:"fail_task_num,omitempty"`
	PendingTaskNum *int32 `json:"pending_task_num,omitempty"`
	SuccessTaskNum *int32 `json:"success_task_num,omitempty"`
	TotalTaskNum   *int32 `json:"total_task_num,omitempty"`
}

type GetJobDetailsResResultTaskInfoListItem struct {
	// TaskBrief string json:"task_brief,omitempty"
	CreateTime *int32 `json:"create_time,omitempty"`
	ExpireTime *int32 `json:"expire_time,omitempty"`

	// TaskID string json:"TaskID"VendorID int json:"VendorID"
	GlobalTaskID *string `json:"global_task_id,omitempty"`
	HostID       *string `json:"host_id,omitempty"`
	InstanceID   *string `json:"instance_id,omitempty"`

	// 失败错误信息或者成功执行的返回值
	Msg *string `json:"msg,omitempty"`

	// 子任务状态：
	// 99：等待中 100：执行中 200：成功 500：失败
	TaskStatus    *int32  `json:"task_status,omitempty"`
	TaskStatusStr *string `json:"task_status_str,omitempty"`

	// 任务类型
	TaskType   *string `json:"task_type,omitempty"`
	UpdateTime *int32  `json:"update_time,omitempty"`
}

type ImportContainerImageBody struct {
	// 镜像描述信息，由用户按需添加镜像备注说明，便于业务识别
	ImageAnnotation *string                                `json:"image_annotation,omitempty"`
	ImageFileType   *ImportContainerImageBodyImageFileType `json:"image_file_type,omitempty"`

	// 镜像名称，由用户自定义，用于业务标识，128个英文或中文字符及以内，需以英文或中文开头，可以包含数字、下划线
	ImageName    *string                               `json:"image_name,omitempty"`
	ImageTosInfo *ImportContainerImageBodyImageTosInfo `json:"image_tos_info,omitempty"`
	ImageURLInfo *ImportContainerImageBodyImageURLInfo `json:"image_url_info,omitempty"`
}

type ImportContainerImageBodyImageTosInfo struct {
	// REQUIRED; 镜像文件所在 TOS 的 bucket，例如：cloudphone
	Bucket string `json:"bucket"`

	// REQUIRED; 镜像文件所在 TOS 的 endpoint，例如：tos-cn-beijing-volces.com
	Endpoint string `json:"endpoint"`

	// REQUIRED; 对象存储所在区域： cn-north => 华北 cn-south => 华南 cn-east => 华东 cn-middle => 华中
	Region string `json:"region"`

	// REQUIRED
	SystemImg ImportContainerImageBodyImageTosInfoSystemImg `json:"system_img"`

	// REQUIRED
	VendorImg ImportContainerImageBodyImageTosInfoVendorImg `json:"vendor_img"`
}

type ImportContainerImageBodyImageTosInfoSystemImg struct {
	// REQUIRED; 镜像文件在 TOS 中的存储路径（不能以 / 开头），例如： android-image-build/system.img
	FilePath string `json:"file_path"`

	// 不传入时，不进行Md5校验
	MD5 *string `json:"md5,omitempty"`
}

type ImportContainerImageBodyImageTosInfoVendorImg struct {
	// REQUIRED; 镜像文件在 TOS 中的存储路径（不能以 / 开头），例如： android-image-build/system.img
	FilePath string `json:"file_path"`

	// 不传入时，不进行Md5校验
	MD5 *string `json:"md5,omitempty"`
}

type ImportContainerImageBodyImageURLInfo struct {
	// REQUIRED
	SystemImg ImportContainerImageBodyImageURLInfoSystemImg  `json:"system_img"`
	VendorImg *ImportContainerImageBodyImageURLInfoVendorImg `json:"vendor_img,omitempty"`
}

type ImportContainerImageBodyImageURLInfoSystemImg struct {
	// REQUIRED; 文件的下载地址
	URL string `json:"url"`

	// 文件的MD5值
	MD5 *string `json:"md5,omitempty"`
}

type ImportContainerImageBodyImageURLInfoVendorImg struct {
	// REQUIRED; 文件的下载地址
	URL string `json:"url"`

	// 文件的MD5值
	MD5 *string `json:"md5,omitempty"`
}

type ImportContainerImageRes struct {
	// REQUIRED
	ResponseMetadata ImportContainerImageResResponseMetadata `json:"ResponseMetadata"`
	Result           *ImportContainerImageResResult          `json:"Result,omitempty"`
}

type ImportContainerImageResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *ImportContainerImageResResponseMetadataError `json:"Error,omitempty"`
}

type ImportContainerImageResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ImportContainerImageResResult struct {
	// 镜像ID
	ImageID *string `json:"image_id,omitempty"`
}

type InitializeHostBody struct {
	// REQUIRED; 云机ID列表
	HostIDList []string `json:"HostIdList"`

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`
}

type InitializeHostRes struct {
	// REQUIRED
	ResponseMetadata InitializeHostResResponseMetadata `json:"ResponseMetadata"`
	Result           *InitializeHostResResult          `json:"Result,omitempty"`
}

type InitializeHostResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *InitializeHostResResponseMetadataError `json:"Error,omitempty"`
}

type InitializeHostResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type InitializeHostResResult struct {
	// 失败列表
	FailedList []*InitializeHostResResultFailedListItem `json:"FailedList,omitempty"`
}

type InitializeHostResResultFailedListItem struct {
	// 操作的ID，云机Id 或 实例Id
	ID *string `json:"Id,omitempty"`

	// 信息
	Msg *string `json:"Msg,omitempty"`
}

type InstallApplicationBody struct {
	// REQUIRED; 目标 Id 列表
	ObjectIDList []string `json:"object_id_list"`

	// REQUIRED; 对象的类别，可选枚举值包括： instance（对指定实例为对象进行操作） host（对云机下的所有实例进行操作)
	ObjectType string `json:"object_type"`

	// REQUIRED; 产品ID`
	ProductID string `json:"product_id"`

	// REQUIRED
	TosInfo InstallApplicationBodyTosInfo `json:"tos_info"`

	// 应用安装可选参数列表 1（ApkInstallAllowTest，允许测试） 2（ApkInstallReplaceExisting，替换现存） 3（ApkInstallGrantAllPerm，获取全部许可） 4（ApkInstallABI，覆盖默认
	// ABI） 5（ApkInstallInternalFlash，内部闪存） 6（ApkInstallAllowDowngrade，允许降级）
	OptionList []*ArrayItemschema `json:"option_list,omitempty"`
}

type InstallApplicationBodyTosInfo struct {
	// REQUIRED; 火山引擎对象存储中的存储桶名称
	TosBucket string `json:"tos_bucket"`

	// REQUIRED; 火山引擎对象存储中的文件路径
	TosFilePath string `json:"tos_file_path"`

	// 火山引擎对象存储服务地址（地域节点），若为空，则使用默认值：tos-cn-beijing.volces.com 默认 constdef.TosEndpointCNBJOnline
	Endpoint *string `json:"endpoint,omitempty"`

	// 火山引擎对象存储服务区域，若为空，则使用默认值：cn-beijing 默认 constdef.TosRegionBJ
	Region *string `json:"region,omitempty"`
}

type InstallApplicationRes struct {
	// REQUIRED
	ResponseMetadata InstallApplicationResResponseMetadata `json:"ResponseMetadata"`
	Result           *InstallApplicationResResult          `json:"Result,omitempty"`
}

type InstallApplicationResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *InstallApplicationResResponseMetadataError `json:"Error,omitempty"`
}

type InstallApplicationResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type InstallApplicationResResult struct {
	// 失败的ID列表
	FailedIDList []*InstallApplicationResResultFailedIDListItem `json:"failed_id_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 废弃
	PassedIDList []*string `json:"passed_id_list,omitempty"`
}

type InstallApplicationResResultFailedIDListItem struct {
	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type LatestMetricInstanceQuery struct {
	// REQUIRED; 实例 Id
	InstanceID string `json:"instance_id" query:"instance_id"`
}

type LatestMetricInstanceRes struct {
	// REQUIRED
	ResponseMetadata LatestMetricInstanceResResponseMetadata `json:"ResponseMetadata"`
	Result           *LatestMetricInstanceResResult          `json:"Result,omitempty"`
}

type LatestMetricInstanceResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *LatestMetricInstanceResResponseMetadataError `json:"Error,omitempty"`
}

type LatestMetricInstanceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type LatestMetricInstanceResResult struct {
	// CPU 5分钟内的负载百分比，例如14.18，即 14.18%
	CPULoad5 *float32 `json:"cpu_load5,omitempty"`

	// 空闲内存，单位 byte
	MemFree *int32 `json:"mem_free,omitempty"`

	// 总内存，单位 byte
	MemTotal *int32 `json:"mem_total,omitempty"`

	// 系统 10s 内的总平均接收流量，单位 Kbps
	RxAll *int32 `json:"rx_all,omitempty"`

	// 系统 10s 内的 system 用户平均接收流量，单位 Kbps
	RxSys *int32 `json:"rx_sys,omitempty"`

	// 系统 10s 内的总平均发送流量，单位 Kbps
	TxAll *int32 `json:"tx_all,omitempty"`

	// 系统 10s 内的 system 用户平均发送流量，单位 Kbps
	TxSys *int32 `json:"tx_sys,omitempty"`
}

type ListAdbKeyQuery struct {
	// REQUIRED
	ProductID string `json:"product_id" query:"product_id"`
	Count     *int32 `json:"count,omitempty" query:"count"`
	KeyID     *int32 `json:"key_id,omitempty" query:"key_id"`

	// 密钥对名称
	KeyName *string `json:"key_name,omitempty" query:"key_name"`
	Offset  *int32  `json:"offset,omitempty" query:"offset"`
}

type ListAdbKeyRes struct {
	// REQUIRED
	ResponseMetadata ListAdbKeyResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListAdbKeyResResult          `json:"Result,omitempty"`
}

type ListAdbKeyResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *ListAdbKeyResResponseMetadataError `json:"Error,omitempty"`
}

type ListAdbKeyResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListAdbKeyResResult struct {
	Row   []*ListAdbKeyResResultRowItem `json:"row,omitempty"`
	Total *int32                        `json:"total,omitempty"`
}

type ListAdbKeyResResultRowItem struct {
	AuthType *int32 `json:"auth_type,omitempty"`

	// 密钥对绑定的实例数量
	BindInstanceNum *int32 `json:"bind_instance_num,omitempty"`
	CreateAt        *int32 `json:"create_at,omitempty"`

	// 公钥指纹
	Fingerprint *string `json:"fingerprint,omitempty"`
	KeyDesc     *string `json:"key_desc,omitempty"`
	KeyID       *int32  `json:"key_id,omitempty"`
	KeyName     *string `json:"key_name,omitempty"`
	ProductID   *string `json:"product_id,omitempty"`
}

type ListContainerImagesQuery struct {
	// 单页数量
	Count *int32 `json:"count,omitempty" query:"count"`

	// 默认为false true表示可以查询未发布的公共镜像，此时ImageIDList不能为空
	ExpandScope *bool `json:"expand_scope,omitempty" query:"expand_scope"`

	// 镜像ID列表, 以","符号隔离
	ImageIDList *string `json:"image_id_list,omitempty" query:"image_id_list"`
	ImageName   *string `json:"image_name,omitempty" query:"image_name"`

	// 镜像状态 UNKNOWN（未知） IN_QUEUE（队列中） BUILDING（构建中） BUILT（构建完成） FAILED（构建失败）
	ImageStatus *string `json:"image_status,omitempty" query:"image_status"`

	// 镜像是否为公共镜像
	IsPublicImage *bool `json:"is_public_image,omitempty" query:"is_public_image"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty" query:"offset"`
}

type ListContainerImagesRes struct {
	// REQUIRED
	ResponseMetadata ListContainerImagesResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListContainerImagesResResult          `json:"Result,omitempty"`
}

type ListContainerImagesResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *ListContainerImagesResResponseMetadataError `json:"Error,omitempty"`
}

type ListContainerImagesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListContainerImagesResResult struct {
	Row   []*ListContainerImagesResResultRowItem `json:"row,omitempty"`
	Total *int32                                 `json:"total,omitempty"`
}

type ListContainerImagesResResultRowItem struct {
	// 构建任务创建时间
	CreateAt *int32 `json:"create_at,omitempty"`

	// 镜像摘要
	Digest *string `json:"digest,omitempty"`

	// 镜像地址
	ImageAddr *string `json:"image_addr,omitempty"`

	// 备注信息
	ImageAnnotation *string `json:"image_annotation,omitempty"`

	// 镜像 ID
	ImageID *string `json:"image_id,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像标签
	ImageTag *string `json:"image_tag,omitempty"`

	// 镜像上传时间戳, 秒级
	PushAt *int32 `json:"push_at,omitempty"`

	// 镜像大小，单位 byte
	Size *int32 `json:"size,omitempty"`

	// 镜像状态 string UNKNOWN（未知） IN_QUEUE（队列中） BUILDING（构建中） BUILT（构建完成） FAILED（构建失败）
	Status     *string                                        `json:"status,omitempty"`
	StatusCode *ListContainerImagesResResultRowItemStatusCode `json:"status_code,omitempty"`

	// 构建任务更新时间
	UpdateAt *int32 `json:"update_at,omitempty"`

	// 镜像版本
	Version *string `json:"version,omitempty"`
}

type ListDcCapacityBody struct {
	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`

	// 机房ID
	DcID string `json:"DcId"`

	// 多开的数量，默认为1，单开
	InstanceAmount *int32 `json:"InstanceAmount,omitempty"`
	// 共享存储的大小，单位MB
	InstanceDisk *int32                 `json:"InstanceDisk,omitempty"`
	Isp          *ListDcCapacityBodyIsp `json:"Isp,omitempty"`

	// 套餐code，比如CloudHostARMNode8c12g_daily
	PackageCode *string `json:"PackageCode,omitempty"`

	// pre/ post 预付费还是后付费
	PayMode *string `json:"PayMode,omitempty"`

	// monthly / daily 按月还是按天计费 // monthly / daily 按月还是按天计费
	PayPeriod *string `json:"PayPeriod,omitempty"`
}

type ListDcCapacityRes struct {
	// REQUIRED
	ResponseMetadata ListDcCapacityResResponseMetadata `json:"ResponseMetadata"`
	Result           []*ListDcCapacityResResultItem    `json:"result,omitempty"`
}

type ListDcCapacityResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *ListDcCapacityResResponseMetadataError `json:"Error,omitempty"`
}

type ListDcCapacityResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListDcCapacityResResultItem struct {
	// 当前模式下可订购的数量
	AvailableCount *int32  `json:"AvailableCount,omitempty"`
	DcID           *string `json:"DcId,omitempty"`

	// V2:用于预付或者后付中的估价; 唯一性
	HostVolChargeItemID *string `json:"HostVolChargeItemId,omitempty"`

	// 实例数量，对应单开、双开、四开等
	InstanceAmount *int32                          `json:"InstanceAmount,omitempty"`
	Isp            *ListDcCapacityResResultItemIsp `json:"Isp,omitempty"`

	// "移动", "联通", "电信", "三线", "BGP(多线)"
	IspName *string `json:"IspName,omitempty"`

	// 业务侧套餐ID
	PackageID *string `json:"PackageId,omitempty"`

	// pre 预付费; post 后付费
	PayMode *string `json:"PayMode,omitempty"`

	// 计费周期 [ daily-按天 | monthly-按月]
	PayPeriod *string `json:"PayPeriod,omitempty"`
	ProductID *string `json:"ProductId,omitempty"`
}

type ListHostMetricDataBody struct {
	// REQUIRED; 查询结束的时间戳，闭区间
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 云机ID
	HostID string `json:"HostId"`

	// REQUIRED; 监控指标名称 UpBandwidth -> 上行带宽 DownBandwidth -> 下行带宽 MaxBandwidth -> 最大带宽
	MetricName []ListHostMetricDataBodyMetricNameItem `json:"MetricName"`

	// REQUIRED
	MetricType ListHostMetricDataBodyMetricType `json:"MetricType"`

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`

	// REQUIRED; 查询开始的时间戳，闭区间
	StartTime int32 `json:"StartTime"`

	// 聚合粒度，目前时5min,10min,30min,60min
	SimpleRate *int32 `json:"SimpleRate,omitempty"`
}

type ListHostMetricDataRes struct {
	// REQUIRED
	ResponseMetadata ListHostMetricDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListHostMetricDataResResult          `json:"Result,omitempty"`
}

type ListHostMetricDataResResponseMetadata struct {
	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *ListHostMetricDataResResponseMetadataError `json:"Error,omitempty"`
}

type ListHostMetricDataResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListHostMetricDataResResult struct {
	// 下行带宽的查询结果
	DownBandwidthResult []*ListHostMetricDataResResultDownBandwidthResultItem `json:"DownBandwidthResult,omitempty"`

	// 监控指标查询的结束时间
	EndTime *int32 `json:"EndTime,omitempty"`

	// 最大带宽的查询结果
	MaxBandwidthResult []*ListHostMetricDataResResultMaxBandwidthResultItem `json:"MaxBandwidthResult,omitempty"`

	// 监控指标名称 UpBandwidth -> 上行带宽 DownBandwidth -> 下行带宽 MaxBandwidth -> 最大带宽
	MetricName []*ListHostMetricDataResResultMetricNameItem `json:"MetricName,omitempty"`
	MetricType *ListHostMetricDataResResultMetricType       `json:"MetricType,omitempty"`

	// 聚合粒度，目前时5min,10min,30min,60min
	SimpleRate *int32 `json:"SimpleRate,omitempty"`

	// 监控指标查询的开始时间
	StartTime *int32 `json:"StartTime,omitempty"`

	// 上行带宽的结果
	UpBandwidthResult []*ListHostMetricDataResResultUpBandwidthResultItem `json:"UpBandwidthResult,omitempty"`
}

type ListHostMetricDataResResultDownBandwidthResultItem struct {
	// 时间 example:"2023-07-25T17:40:00+08:00"
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 数值 带宽单位为bps，流量单位 B
	Value *float32 `json:"Value,omitempty"`
}

type ListHostMetricDataResResultMaxBandwidthResultItem struct {
	// 时间 example:"2023-07-25T17:40:00+08:00"
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 数值 带宽单位为bps，流量单位 B
	Value *float32 `json:"Value,omitempty"`
}

type ListHostMetricDataResResultUpBandwidthResultItem struct {
	// 时间 example:"2023-07-25T17:40:00+08:00"
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 数值 带宽单位为bps，流量单位 B
	Value *float32 `json:"Value,omitempty"`
}

type ListHostQuery struct {
	// REQUIRED; 产品 ID
	ProductID string `json:"product_id" query:"product_id"`

	// 分页数量
	Count *int32 `json:"count,omitempty" query:"count"`

	// 机房 ID 列表, split by ','. 内部注释: parse to DcList in function Rectify
	DcIn *string `json:"dc_in,omitempty" query:"dc_in"`

	// 是否返回云机详细参数
	Detail *bool `json:"detail,omitempty" query:"detail"`

	// 云机创建时间晚于, 秒级时间戳, 开区间
	HostCreateAfter *int32 `json:"host_create_after,omitempty" query:"host_create_after"`

	// 云机创建时间早于, 秒级时间戳, 闭区间
	HostCreateBefore *int32 `json:"host_create_before,omitempty" query:"host_create_before"`

	// 云机 ID
	HostID *string `json:"host_id,omitempty" query:"host_id"`

	// 云机 ID 列表, split by ',', 内部注释: parse to HostIdList in function Rectify
	HostIDIn *string `json:"host_id_in,omitempty" query:"host_id_in"`

	// 云机名称模糊查询
	HostNameLike *string `json:"host_name_like,omitempty" query:"host_name_like"`

	// 包含实例 ID 列表. split by ',' 内部注释: parse to InstanceIDList in function Rectify
	InstanceIDIn *string `json:"instance_id_in,omitempty" query:"instance_id_in"`

	// 是否降序
	IsDesc *bool `json:"is_desc,omitempty" query:"is_desc"`

	// 运营商
	Isp *ListHostQueryIsp `json:"isp,omitempty" query:"isp"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty" query:"offset"`

	// 排序字段
	OrderBy   *ListHostQueryOrderBy `json:"order_by,omitempty" query:"order_by"`
	PackageID *string               `json:"package_id,omitempty" query:"package_id"`

	// 云机所在区域
	Region *string `json:"region,omitempty" query:"region"`

	// 云机状态
	Status *int32 `json:"status,omitempty" query:"status"`
}

type ListHostRes struct {

	// REQUIRED
	ResponseMetadata ListHostResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListHostResResult          `json:"Result,omitempty"`
}

type ListHostResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                            `json:"Version"`
	Error   *ListHostResResponseMetadataError `json:"Error,omitempty"`
}

type ListHostResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListHostResResult struct {
	Row   []*ListHostResResultRowItem `json:"row,omitempty"`
	Total *int32                      `json:"total,omitempty"`
}

type ListHostResResultRowItem struct {

	// 平均到每个实例的带宽，例：2 (Mbps)
	Bandwidth     *int32                                 `json:"bandwidth,omitempty"`
	Configuration *ListHostResResultRowItemConfiguration `json:"configuration,omitempty"`

	// 云机创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 机房名称
	DcName *string `json:"dc_name,omitempty"`

	// 过期时间
	ExpireAt *int32 `json:"expire_at,omitempty"`

	// 云机计费项ID
	HostChargeID *string `json:"host_charge_id,omitempty"`

	// 云机创建时间
	HostCreateAt *int32 `json:"host_create_at,omitempty"`

	// 云机 ID
	HostID *string `json:"host_id,omitempty"`

	// 云机名称
	HostName *string `json:"host_name,omitempty"`

	// 付费模式：pre (预付费模式) post(后付费模式)
	HostPayMode *string `json:"host_pay_mode,omitempty"`

	// 包年包月、按天计费、按月计费
	HostPayType *string `json:"host_pay_type,omitempty"`

	// 机房信息
	Idc *string `json:"idc,omitempty"`

	// 单个云机对应的实例数量
	InstanceNum *int32                       `json:"instance_num,omitempty"`
	Isp         *ListHostResResultRowItemIsp `json:"isp,omitempty"`

	// ISP计费项ID
	IspChargeID *string `json:"isp_charge_id,omitempty"`

	// 套餐 ID
	PackageID *string `json:"package_id,omitempty"`

	// 云机资源套餐名称
	PackageName *string `json:"package_name,omitempty"`

	// 云机资源套餐规格
	PackageSpec *string `json:"package_spec,omitempty"`

	// 产品 ID
	ProductID *string `json:"product_id,omitempty"`

	// 机房维度的内网IP
	PublicIP *string `json:"public_ip,omitempty"`

	// 云机所在区域名称： cn-north => 华北 cn-south => 华南 cn-east => 华东 cn-middle => 华中 cn-southwest => 西南
	Region *string                         `json:"region,omitempty"`
	Status *ListHostResResultRowItemStatus `json:"status,omitempty"`

	// 云机更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type ListHostResResultRowItemConfiguration struct {

	// CPU 核心数
	CPUCore *int32 `json:"cpu_core,omitempty"`

	// 实例资源套餐 ID
	ConfigurationCode *string `json:"configuration_code,omitempty"`

	// 实例资源套餐名称
	ConfigurationName *string `json:"configuration_name,omitempty"`

	// 网络计费名称
	IspCodeName *string `json:"isp_code_name,omitempty"`

	// 网络计费套餐
	IspConfigurationCode *string `json:"isp_configuration_code,omitempty"`

	// 内存，单位MB
	Memory *float32 `json:"memory,omitempty"`
}

type ListInstanceMetricDataBody struct {

	// REQUIRED; 查询结束的时间戳，闭区间
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 实例ID
	InstanceID string `json:"InstanceId"`

	// REQUIRED; 监控指标名称 UpBandwidth -> 上行带宽 DownBandwidth -> 下行带宽 MaxBandwidth -> 最大带宽
	MetricName []ListInstanceMetricDataBodyMetricNameItem `json:"MetricName"`

	// REQUIRED
	MetricType ListInstanceMetricDataBodyMetricType `json:"MetricType"`

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`

	// REQUIRED; 查询开始的时间戳，闭区间
	StartTime int32 `json:"StartTime"`

	// 聚合粒度，目前时5min,10min,30min,60min
	SimpleRate *int32 `json:"SimpleRate,omitempty"`
}

type ListInstanceMetricDataRes struct {

	// REQUIRED
	ResponseMetadata ListInstanceMetricDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListInstanceMetricDataResResult          `json:"Result,omitempty"`
}

type ListInstanceMetricDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *ListInstanceMetricDataResResponseMetadataError `json:"Error,omitempty"`
}

type ListInstanceMetricDataResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListInstanceMetricDataResResult struct {

	// 下行带宽的查询结果
	DownBandwidthResult []*ListInstanceMetricDataResResultDownBandwidthResultItem `json:"DownBandwidthResult,omitempty"`

	// 监控指标查询的结束时间
	EndTime *int32 `json:"EndTime,omitempty"`

	// 最大带宽的查询结果
	MaxBandwidthResult []*ListInstanceMetricDataResResultMaxBandwidthResultItem `json:"MaxBandwidthResult,omitempty"`

	// 监控指标名称 UpBandwidth -> 上行带宽 DownBandwidth -> 下行带宽 MaxBandwidth -> 最大带宽
	MetricName []*ListInstanceMetricDataResResultMetricNameItem `json:"MetricName,omitempty"`
	MetricType *ListInstanceMetricDataResResultMetricType       `json:"MetricType,omitempty"`

	// 聚合粒度，目前时5min,10min,30min,60min
	SimpleRate *int32 `json:"SimpleRate,omitempty"`

	// 监控指标查询的开始时间
	StartTime *int32 `json:"StartTime,omitempty"`

	// 上行带宽的结果
	UpBandwidthResult []*ListInstanceMetricDataResResultUpBandwidthResultItem `json:"UpBandwidthResult,omitempty"`
}

type ListInstanceMetricDataResResultDownBandwidthResultItem struct {

	// 时间 example:"2023-07-25T17:40:00+08:00"
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 数值 带宽单位为bps，流量单位 B
	Value *float32 `json:"Value,omitempty"`
}

type ListInstanceMetricDataResResultMaxBandwidthResultItem struct {

	// 时间 example:"2023-07-25T17:40:00+08:00"
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 数值 带宽单位为bps，流量单位 B
	Value *float32 `json:"Value,omitempty"`
}

type ListInstanceMetricDataResResultUpBandwidthResultItem struct {

	// 时间 example:"2023-07-25T17:40:00+08:00"
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 数值 带宽单位为bps，流量单位 B
	Value *float32 `json:"Value,omitempty"`
}

type ListInstanceQuery struct {

	// REQUIRED; 产品 ID
	ProductID string `json:"product_id" query:"product_id"`

	// key ID
	AdbKeyID *int32 `json:"adb_key_id,omitempty" query:"adb_key_id"`

	// 返回数量, 默认 100
	Count *int32 `json:"count,omitempty" query:"count"`

	// 创建时间范围, 开始时间, 秒级时间戳, 开区间
	CreateAfter *int32 `json:"create_after,omitempty" query:"create_after"`

	// 创建时间范围, 结束时间, 秒级时间戳, 闭区间
	CreateBefore *int32 `json:"create_before,omitempty" query:"create_before"`

	// 是否返回详细信息, e.g. tag, security group, key, etc.
	Detail *bool `json:"detail,omitempty" query:"detail"`

	// 云机 ID
	HostID *string `json:"host_id,omitempty" query:"host_id"`

	// 机房
	Idc *string `json:"dc,omitempty" query:"dc"`

	// 实例 ID
	InstanceID *string `json:"instance_id,omitempty" query:"instance_id"`

	// 批量筛选, 实例ID, 逗号分隔 string. 内部注释-勿展示: 为了兼容性，优先级比 InstanceIdList 低, 参见 Rectify()
	InstanceIDIn *string `json:"instance_id_in,omitempty" query:"instance_id_in"`

	// 模糊查询, 实例ID
	InstanceIDLike *string `json:"instance_id_like,omitempty" query:"instance_id_like"`

	// 模糊查询, 实例名称
	InstanceNameLike *string `json:"instance_name_like,omitempty" query:"instance_name_like"`

	// 是否升序, 默认降序
	IsOrderAsc *bool `json:"is_order_asc,omitempty" query:"is_order_asc"`

	// 运营商
	Isp *ListInstanceQueryIsp `json:"isp,omitempty" query:"isp"`

	// 偏移量, 默认 0
	Offset *int32 `json:"offset,omitempty" query:"offset"`

	// 排序字段, 支持 instance_id, sn
	OrderBy *string `json:"order_by,omitempty" query:"order_by"`

	// 套餐 ID
	PackageID *string `json:"package_id,omitempty" query:"package_id"`

	// 安全组绑定状态
	SgBoundSt *ListInstanceQuerySgBoundSt `json:"sg_bound_st,omitempty" query:"sg_bound_st"`

	// 实例状态
	Status *int32 `json:"status,omitempty" query:"status"`

	// 批量筛选, 状态, 逗号分隔 int
	StatusIn *string `json:"status_in,omitempty" query:"status_in"`

	// 批量筛选, 标签ID, 逗号分隔 string
	TagIDIn *string `json:"tag_id_in,omitempty" query:"tag_id_in"`
}

type ListInstanceRes struct {

	// REQUIRED
	ResponseMetadata ListInstanceResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListInstanceResResult          `json:"Result,omitempty"`
}

type ListInstanceResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *ListInstanceResResponseMetadataError `json:"Error,omitempty"`
}

type ListInstanceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListInstanceResResult struct {
	Row   []*ListInstanceResResultRowItem `json:"row,omitempty"`
	Total *int32                          `json:"total,omitempty"`
}

type ListInstanceResResultRowItem struct {
	AdbKey *ListInstanceResResultRowItemAdbKey `json:"adb_key,omitempty"`

	// 带宽
	Bandwidth     *int32                                     `json:"bandwidth,omitempty"`
	Configuration *ListInstanceResResultRowItemConfiguration `json:"configuration,omitempty"`

	// 创建时间, unix 时间戳, 秒级
	CreateAt *int32 `json:"create_at,omitempty"`

	// 机房 ID
	Dc *string `json:"dc,omitempty"`

	// 机房名称
	DcName *string `json:"dc_name,omitempty"`

	// 帧率
	Fps *int32 `json:"fps,omitempty"`

	// 主机ID
	HostID *string `json:"host_id,omitempty"`

	// 镜像ID
	ImageID *string `json:"image_id,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 实例ID
	InstanceID *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string                          `json:"instance_name,omitempty"`
	Isp          *ListInstanceResResultRowItemIsp `json:"isp,omitempty"`

	// 产品ID
	ProductID *string `json:"product_id,omitempty"`

	// 地域
	Region *string `json:"region,omitempty"`

	// 分辨率
	Resolution    *string                                    `json:"resolution,omitempty"`
	SecurityGroup *ListInstanceResResultRowItemSecurityGroup `json:"security_group,omitempty"`
	SgBoundSt     *ListInstanceResResultRowItemSgBoundSt     `json:"sg_bound_st,omitempty"`

	// 安全组绑定状态字符串
	SgBoundStStr *string `json:"sg_bound_st_str,omitempty"`

	// 序列号, maybe useless
	Sn     *string                             `json:"sn,omitempty"`
	Status *ListInstanceResResultRowItemStatus `json:"status,omitempty"`

	// 实例状态字符串
	StatusStr *string                          `json:"status_str,omitempty"`
	Tag       *ListInstanceResResultRowItemTag `json:"tag,omitempty"`
}

type ListInstanceResResultRowItemAdbKey struct {

	// 用户权限类型：
	// 1（root） 2（user）
	AuthType *int32 `json:"auth_type,omitempty"`

	// 密钥对绑定的实例数量
	BindInstanceNum *int32 `json:"bind_instance_num,omitempty"`
	CreateAt        *int32 `json:"create_at,omitempty"`

	// 公钥指纹
	Fingerprint *string `json:"fingerprint,omitempty"`
	KeyDesc     *string `json:"key_desc,omitempty"`
	KeyID       *int32  `json:"key_id,omitempty"`
	KeyName     *string `json:"key_name,omitempty"`
	ProductID   *string `json:"product_id,omitempty"`
	PublicKey   *string `json:"public_key,omitempty"`
}

type ListInstanceResResultRowItemConfiguration struct {

	// CPU 核心数
	CPUCore *int32 `json:"cpu_core,omitempty"`

	// 实例资源套餐 ID
	ConfigurationCode *string `json:"configuration_code,omitempty"`

	// 实例资源套餐名称
	ConfigurationName *string `json:"configuration_name,omitempty"`

	// 网络计费名称
	IspCodeName *string `json:"isp_code_name,omitempty"`

	// 网络计费套餐
	IspConfigurationCode *string `json:"isp_configuration_code,omitempty"`

	// 内存，单位MB
	Memory *float32 `json:"memory,omitempty"`
}

type ListInstanceResResultRowItemSecurityGroup struct {

	// BindHostNum *int64 json:"bind_host_num,omitempty" // Deprecated: 请使用BindInstanceNum
	BindInstanceNum *int32 `json:"bind_instance_num,omitempty"`

	// 创建时间，秒级时间戳
	CreateAt *int32 `json:"create_at,omitempty"`

	// 安全组所属业务 ID
	ProductID *string `json:"product_id,omitempty"`

	// 安全组描述
	SecurityGroupDesc *string `json:"security_group_desc,omitempty"`

	// 安全组 ID
	SecurityGroupID *int32 `json:"security_group_id,omitempty"`

	// 安全组名称
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 安全组规则列表
	SecurityRuleList []*ListInstanceResResultRowPropertiesItemsItem `json:"security_rule_list,omitempty"`

	// 更新时间，秒级时间戳
	UpdateAt *int32 `json:"update_at,omitempty"`
}

type ListInstanceResResultRowItemTag struct {
	ProductID          *string `json:"product_id,omitempty"`
	RelatedInstanceNum *int32  `json:"related_instance_num,omitempty"`
	TagDesc            *string `json:"tag_desc,omitempty"`
	TagID              *string `json:"tag_id,omitempty"`
	TagName            *string `json:"tag_name,omitempty"`
}

type ListInstanceResResultRowPropertiesItemsItem struct {
	Expose   *ListInstanceResResultRowItemSecurityGroupSecurityRuleListItemExpose   `json:"expose,omitempty"`
	NatType  *ListInstanceResResultRowItemSecurityGroupSecurityRuleListItemNatType  `json:"nat_type,omitempty"`
	Protocol *ListInstanceResResultRowItemSecurityGroupSecurityRuleListItemProtocol `json:"protocol,omitempty"`

	// 安全组规则ID
	RuleID *int32 `json:"rule_id,omitempty"`

	// 源端口
	SourcePort *int32 `json:"source_port,omitempty"`
}

type ListPackageBody struct {

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`

	// 套餐code，比如CloudHostARMNode8c12g_daily
	PackageCode *string `json:"PackageCode,omitempty"`

	// 套餐资源类型 k8s
	PackageModel *string `json:"PackageModel,omitempty"`

	// pre/ post 预付费还是后付费
	PayMode *string `json:"PayMode,omitempty"`

	// monthly / daily 按月还是按天计费
	PayPeriod *string `json:"PayPeriod,omitempty"`
}

type ListPackageRes struct {

	// REQUIRED
	ResponseMetadata ListPackageResResponseMetadata `json:"ResponseMetadata"`
	Result           []*ListPackageResResultItem    `json:"Result,omitempty"`
}

type ListPackageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                               `json:"Version"`
	Error   *ListPackageResResponseMetadataError `json:"Error,omitempty"`
}

type ListPackageResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListPackageResResultItem struct {
	CPU *int32 `json:"Cpu,omitempty"`

	// cpu核数
	CPUCore *int32 `json:"CpuCore,omitempty"`

	// V1：火山侧计费配置ID
	ConfigurationCode *string `json:"ConfigurationCode,omitempty"`

	// 套餐帧率
	Fps *int32 `json:"Fps,omitempty"`

	// 实例数量，对应单开、双开、四开等
	InstanceAmount *int32 `json:"InstanceAmount,omitempty"`

	// 内存
	Memory *float32 `json:"Memory,omitempty"`

	// 套餐模式，云原生还是LXC
	Model *string `json:"Model,omitempty"`

	// 套餐描述
	PackageDesc *string `json:"PackageDesc,omitempty"`

	// 业务侧套餐ID
	PackageID *string `json:"PackageId,omitempty"`

	// 套餐名称
	PackageName *string `json:"PackageName,omitempty"`

	// pre 预付费; post 后付费
	PayMode *string `json:"PayMode,omitempty"`

	// 计费周期 [ daily-按天 | monthly-按月]
	Period *string `json:"Period,omitempty"`

	// 分辨率
	ScreenType *string `json:"ScreenType,omitempty"`

	// 磁盘大小
	Size *int32 `json:"Size,omitempty"`

	// V2:用于预付或者后付中的估价; 唯一性
	VolChargeItemID *string `json:"VolChargeItemId,omitempty"`

	// V2:计费项类型
	VolChargeItemType *string `json:"VolChargeItemType,omitempty"`
}

type ListPortMappingQuery struct {

	// REQUIRED; 产品ID
	ProductID string `json:"product_id" query:"product_id"`

	// 返回数量
	Count *int32 `json:"count,omitempty" query:"count"`

	// 实例 ID列表，多个 ID 使用英文逗号分隔, 内部注释: parse to InstanceIDList in function Rectify
	InstanceIDIn *string `json:"instance_id_in,omitempty" query:"instance_id_in"`

	// 端口映射的运营商
	Isp *ListPortMappingQueryIsp `json:"isp,omitempty" query:"isp"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty" query:"offset"`

	// 端口映射ID
	PortMappingID *string `json:"port_mapping_id,omitempty" query:"port_mapping_id"`

	// 端口映射的协议
	ProtocolEnum *ListPortMappingQueryProtocolEnum `json:"protocol_enum,omitempty" query:"protocol_enum"`

	// 安全组ID
	SecurityGroupID *int32 `json:"security_group_id,omitempty" query:"security_group_id"`

	// 安全规则ID
	SecurityRuleID *int32 `json:"security_rule_id,omitempty" query:"security_rule_id"`

	// 端口映射状态
	State *ListPortMappingQueryState `json:"state,omitempty" query:"state"`

	// 端口映射状态列表
	StateIn *string `json:"state_in,omitempty" query:"state_in"`

	// gen by func Rectify
	StatesIn []*string `json:"statesIn,omitempty" query:"statesIn"`
}

type ListPortMappingRes struct {

	// REQUIRED
	ResponseMetadata ListPortMappingResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListPortMappingResResult          `json:"Result,omitempty"`
}

type ListPortMappingResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *ListPortMappingResResponseMetadataError `json:"Error,omitempty"`
}

type ListPortMappingResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListPortMappingResResult struct {
	Row   []*ListPortMappingResResultRowItem `json:"row,omitempty"`
	Total *int32                             `json:"total,omitempty"`
}

type ListPortMappingResResultRowItem struct {

	// 连接云手机的 IP（pubicip 或 proxyip 二选一）
	ConnectIP *string `json:"connect_ip,omitempty"`

	// 连接云手机的端口号（pubicport 或 proxyport 二选一）
	ConnectPort *int32  `json:"connect_port,omitempty"`
	Description *string `json:"description,omitempty"`

	// 实例ID
	InstanceID *string `json:"instance_id,omitempty"`

	// 实例源端口号
	InstancePort *int32                              `json:"instance_port,omitempty"`
	Isp          *ListPortMappingResResultRowItemIsp `json:"isp,omitempty"`

	// 端口映射ID
	PortMappingID *string `json:"port_mapping_id,omitempty"`

	// 产品ID
	ProductID    *string                                      `json:"product_id,omitempty"`
	Protocol     *ListPortMappingResResultRowItemProtocol     `json:"protocol,omitempty"`
	ProtocolEnum *ListPortMappingResResultRowItemProtocolEnum `json:"protocol_enum,omitempty"`

	// 内网代理 IP
	ProxyIP *string `json:"proxy_ip,omitempty"`

	// 内网代理端口
	ProxyPort *int32 `json:"proxy_port,omitempty"`

	// 公网 IP
	PublicIP *string `json:"public_ip,omitempty"`

	// 公网端口号
	PublicPort *int32 `json:"public_port,omitempty"`

	// 安全组ID
	SecurityGroupID *int32 `json:"security_group_id,omitempty"`

	// 安全组规则ID
	SecurityRuleID *int32                                `json:"security_rule_id,omitempty"`
	State          *ListPortMappingResResultRowItemState `json:"state,omitempty"`
}

type ListProductQuery struct {
	AccountID      *int32  `json:"accountId,omitempty" query:"accountId"`
	Count          *int32  `json:"count,omitempty" query:"count"`
	MediaServiceID *int32  `json:"mediaServiceId,omitempty" query:"mediaServiceId"`
	Offset         *int32  `json:"offset,omitempty" query:"offset"`
	ProductID      *string `json:"productId,omitempty" query:"productId"`
	ProductName    *string `json:"productName,omitempty" query:"productName"`
	ProductType    *int32  `json:"productType,omitempty" query:"productType"`
}

type ListProductRes struct {

	// REQUIRED
	ResponseMetadata ListProductResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListProductResResult          `json:"Result,omitempty"`
}

type ListProductResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                               `json:"Version"`
	Error   *ListProductResResponseMetadataError `json:"Error,omitempty"`
}

type ListProductResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListProductResResult struct {
	Row   []*ListProductResResultRowItem `json:"row,omitempty"`
	Total *int32                         `json:"total,omitempty"`
}

type ListProductResResultRowItem struct {

	// 火山引擎主账户ID
	AccountID *int32 `json:"account_id,omitempty"`

	// 审核状态
	AuditStatus *int32 `json:"audit_status,omitempty"`

	// 创建时间
	CreateAt *int32 `json:"create_at,omitempty"`

	// 管理员列表
	OwnerList *string `json:"owner_list,omitempty"`

	// 业务描述
	ProductDesc *string `json:"product_desc,omitempty"`

	// 业务ID
	ProductID *string `json:"product_id,omitempty"`

	// 业务名称
	ProductName *string                                 `json:"product_name,omitempty"`
	ProductType *ListProductResResultRowItemProductType `json:"product_type,omitempty"`

	// 用户列表
	UserList []*ListProductResResultRowPropertiesItemsItem `json:"user_list,omitempty"`

	// 账户名
	UserName *string `json:"user_name,omitempty"`
}

type ListProductResResultRowPropertiesItemsItem struct {
	CreateDate        *string `json:"create_date,omitempty"`
	Description       *string `json:"description,omitempty"`
	DisplayName       *string `json:"display_name,omitempty"`
	Email             *string `json:"email,omitempty"`
	EmailVerify       *bool   `json:"email_verify,omitempty"`
	ID                *int32  `json:"id,omitempty"`
	MobilePhone       *string `json:"mobile_phone,omitempty"`
	MobilePhoneVerify *bool   `json:"mobile_phone_verify,omitempty"`
	Source            *string `json:"source,omitempty"`
	Trn               *string `json:"trn,omitempty"`
	Username          *string `json:"username,omitempty"`
}

type ListSecurityGroupQuery struct {

	// REQUIRED; 安全组所属业务 ID
	ProductID string `json:"product_id" query:"product_id"`

	// 单页数量
	Count *int32 `json:"count,omitempty" query:"count"`

	// 筛选条件，是否返回安全组规则信息
	Detail *bool `json:"detail,omitempty" query:"detail"`

	// 分页参数，偏移量
	Offset *int32 `json:"offset,omitempty" query:"offset"`

	// 安全组ID
	SecurityGroupID *int32 `json:"security_group_id,omitempty" query:"security_group_id"`

	// 安全组名称
	SecurityGroupName *string `json:"security_group_name,omitempty" query:"security_group_name"`
}

type ListSecurityGroupRes struct {

	// REQUIRED
	ResponseMetadata ListSecurityGroupResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListSecurityGroupResResult          `json:"Result,omitempty"`
}

type ListSecurityGroupResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                     `json:"Version"`
	Error   *ListSecurityGroupResResponseMetadataError `json:"Error,omitempty"`
}

type ListSecurityGroupResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListSecurityGroupResResult struct {
	Row   []*ListSecurityGroupResResultRowItem `json:"row,omitempty"`
	Total *int32                               `json:"total,omitempty"`
}

type ListSecurityGroupResResultRowItem struct {

	// BindHostNum *int64 json:"bind_host_num,omitempty" // Deprecated: 请使用BindInstanceNum
	BindInstanceNum *int32 `json:"bind_instance_num,omitempty"`

	// 创建时间，秒级时间戳
	CreateAt *int32 `json:"create_at,omitempty"`

	// 安全组所属业务 ID
	ProductID *string `json:"product_id,omitempty"`

	// 安全组描述
	SecurityGroupDesc *string `json:"security_group_desc,omitempty"`

	// 安全组 ID
	SecurityGroupID *int32 `json:"security_group_id,omitempty"`

	// 安全组名称
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 安全组规则列表
	SecurityRuleList []*ListSecurityGroupResResultRowPropertiesItemsItem `json:"security_rule_list,omitempty"`

	// 更新时间，秒级时间戳
	UpdateAt *int32 `json:"update_at,omitempty"`
}

type ListSecurityGroupResResultRowPropertiesItemsItem struct {
	Expose   *ListSecurityGroupResResultRowItemSecurityRuleListItemExpose   `json:"expose,omitempty"`
	NatType  *ListSecurityGroupResResultRowItemSecurityRuleListItemNatType  `json:"nat_type,omitempty"`
	Protocol *ListSecurityGroupResResultRowItemSecurityRuleListItemProtocol `json:"protocol,omitempty"`

	// 安全组规则ID
	RuleID *int32 `json:"rule_id,omitempty"`

	// 源端口
	SourcePort *int32 `json:"source_port,omitempty"`
}

type ModifyInstanceFpsBody struct {

	// REQUIRED; 帧率，目前仅支持30、60
	Fps int32 `json:"Fps"`

	// REQUIRED; 实例ID
	InstanceID string `json:"InstanceId"`

	// REQUIRED; 产品ID
	ProductID string `json:"ProductId"`
}

type ModifyInstanceFpsRes struct {

	// REQUIRED
	ResponseMetadata ModifyInstanceFpsResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type ModifyInstanceFpsResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                     `json:"Version"`
	Error   *ModifyInstanceFpsResResponseMetadataError `json:"Error,omitempty"`
}

type ModifyInstanceFpsResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ModifyInstanceWindowDisplaySpecBody struct {

	// REQUIRED; 实例ID
	InstanceID string `json:"InstanceId"`

	// REQUIRED; 产品ID
	ProductID        string                                               `json:"ProductId"`
	ResolutionCustom *ModifyInstanceWindowDisplaySpecBodyResolutionCustom `json:"ResolutionCustom,omitempty"`
	ResolutionLevel  *ModifyInstanceWindowDisplaySpecBodyResolutionLevel  `json:"ResolutionLevel,omitempty"`
}

type ModifyInstanceWindowDisplaySpecBodyResolutionCustom struct {

	// REQUIRED; 屏幕显示分辨率X
	DisplayResolutionX int32 `json:"DisplayResolutionX"`

	// REQUIRED; 屏幕显示分辨率Y
	DisplayResolutionY int32 `json:"DisplayResolutionY"`
}

type ModifyInstanceWindowDisplaySpecRes struct {

	// REQUIRED
	ResponseMetadata ModifyInstanceWindowDisplaySpecResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type ModifyInstanceWindowDisplaySpecResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                   `json:"Version"`
	Error   *ModifyInstanceWindowDisplaySpecResResponseMetadataError `json:"Error,omitempty"`
}

type ModifyInstanceWindowDisplaySpecResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type PowerDownInstanceBody struct {

	// REQUIRED
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED
	ProductID string `json:"product_id"`
}

type PowerDownInstanceRes struct {

	// REQUIRED
	ResponseMetadata PowerDownInstanceResResponseMetadata `json:"ResponseMetadata"`
	Result           *PowerDownInstanceResResult          `json:"Result,omitempty"`
}

type PowerDownInstanceResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                     `json:"Version"`
	Error   *PowerDownInstanceResResponseMetadataError `json:"Error,omitempty"`
}

type PowerDownInstanceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type PowerDownInstanceResResult struct {

	// Deprecated: 弃用
	FailIDList []*string `json:"fail_id_list,omitempty"`

	// 失败的记录
	FailedList []*PowerDownInstanceResResultFailedListItem `json:"failed_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 弃用
	SuccessIDList []*string `json:"success_id_list,omitempty"`
}

type PowerDownInstanceResResultFailedListItem struct {

	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type PowerUpInstanceBody struct {

	// REQUIRED; 实例 Id 列表
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED; 产品 Id
	ProductID string `json:"product_id"`
}

type PowerUpInstanceRes struct {

	// REQUIRED
	ResponseMetadata PowerUpInstanceResResponseMetadata `json:"ResponseMetadata"`
	Result           *PowerUpInstanceResResult          `json:"Result,omitempty"`
}

type PowerUpInstanceResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *PowerUpInstanceResResponseMetadataError `json:"Error,omitempty"`
}

type PowerUpInstanceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type PowerUpInstanceResResult struct {

	// Deprecated: 弃用
	FailIDList []*string `json:"fail_id_list,omitempty"`

	// 失败的记录
	FailedList []*PowerUpInstanceResResultFailedListItem `json:"failed_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 弃用
	SuccessIDList []*string `json:"success_id_list,omitempty"`
}

type PowerUpInstanceResResultFailedListItem struct {

	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type PullFileBody struct {

	// REQUIRED; 文件在云机实例中的存储地址
	FilePath string `json:"file_path"`

	// REQUIRED; 实例ID
	InstanceID string `json:"instance_id"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`

	// REQUIRED
	TosInfo PullFileBodyTosInfo `json:"tos_info"`
}

type PullFileBodyTosInfo struct {

	// REQUIRED; 火山引擎对象存储中的存储桶名称
	TosBucket string `json:"tos_bucket"`

	// REQUIRED; 火山引擎对象存储中的文件路径
	TosFilePath string `json:"tos_file_path"`

	// 火山引擎对象存储服务地址（地域节点），若为空，则使用默认值：tos-cn-beijing.volces.com 默认 constdef.TosEndpointCNBJOnline
	Endpoint *string `json:"endpoint,omitempty"`

	// 火山引擎对象存储服务区域，若为空，则使用默认值：cn-beijing 默认 constdef.TosRegionBJ
	Region *string `json:"region,omitempty"`
}

type PullFileRes struct {

	// REQUIRED
	ResponseMetadata PullFileResResponseMetadata `json:"ResponseMetadata"`
	Result           *PullFileResResult          `json:"Result,omitempty"`
}

type PullFileResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                            `json:"Version"`
	Error   *PullFileResResponseMetadataError `json:"Error,omitempty"`
}

type PullFileResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type PullFileResResult struct {

	// 任务ID
	JobID *string `json:"job_id,omitempty"`
}

type PushFileBody struct {

	// REQUIRED; 上传文件至云机实例中的目标目录 上传文件至云机实例中的目标目录，目前仅支持 /data/local/tmp 路径
	DestDirectory string `json:"dest_directory"`

	// REQUIRED; 目标 Id 列表
	ObjectIDList []string `json:"object_id_list"`

	// REQUIRED; 对象的类别，可选枚举值包括： instance（对指定实例为对象进行操作） host（对云机下的所有实例进行操作)
	ObjectType string `json:"object_type"`

	// REQUIRED; 产品ID
	ProductID string               `json:"product_id"`
	TosInfo   *PushFileBodyTosInfo `json:"tos_info,omitempty"`
}

type PushFileBodyTosInfo struct {

	// REQUIRED; 火山引擎对象存储中的存储桶名称
	TosBucket string `json:"tos_bucket"`

	// REQUIRED; 火山引擎对象存储中的文件路径
	TosFilePath string `json:"tos_file_path"`

	// 火山引擎对象存储服务地址（地域节点），若为空，则使用默认值：tos-cn-beijing.volces.com 默认 constdef.TosEndpointCNBJOnline
	Endpoint *string `json:"endpoint,omitempty"`

	// 火山引擎对象存储服务区域，若为空，则使用默认值：cn-beijing 默认 constdef.TosRegionBJ
	Region *string `json:"region,omitempty"`
}

type PushFileRes struct {

	// REQUIRED
	ResponseMetadata PushFileResResponseMetadata `json:"ResponseMetadata"`
	Result           *PushFileResResult          `json:"Result,omitempty"`
}

type PushFileResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                            `json:"Version"`
	Error   *PushFileResResponseMetadataError `json:"Error,omitempty"`
}

type PushFileResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type PushFileResResult struct {

	// 失败的ID列表
	FailedIDList []*PushFileResResultFailedIDListItem `json:"failed_id_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 废弃
	PassedIDList []*string `json:"passed_id_list,omitempty"`
}

type PushFileResResultFailedIDListItem struct {

	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type RebootHostBody struct {

	// REQUIRED; 云机ID列表
	HostIDList []string `json:"HostIdList"`

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`

	// 是否强制重启
	Force *bool `json:"Force,omitempty"`
}

type RebootHostRes struct {

	// REQUIRED
	ResponseMetadata RebootHostResResponseMetadata `json:"ResponseMetadata"`
	Result           *RebootHostResResult          `json:"Result,omitempty"`
}

type RebootHostResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *RebootHostResResponseMetadataError `json:"Error,omitempty"`
}

type RebootHostResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type RebootHostResResult struct {

	// 重启失败的云机列表
	FailIDList []*string `json:"FailIdList,omitempty"`

	// 异步JobId
	JobID *string `json:"JobId,omitempty"`

	// 重启成功的云机列表
	SuccessIDList []*string `json:"SuccessIdList,omitempty"`
}

type RecordScreenBody struct {

	// REQUIRED; 实例ID
	InstanceID string `json:"instance_id"`

	// REQUIRED
	Option RecordScreenBodyOption `json:"option"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`

	// 保存录制文件的名称（文件名称以 .mp4 结尾），当 option 为 start时，为必填项
	FileName *string `json:"file_name,omitempty"`

	// 录屏时长（到期后自动停止），单位: 秒 最大值：10800（3小时） 默认值：180（3分钟）
	TimeLimitSeconds *int32 `json:"time_limit_seconds,omitempty"`
}

type RecordScreenRes struct {

	// REQUIRED
	ResponseMetadata RecordScreenResResponseMetadata `json:"ResponseMetadata"`
	Result           *RecordScreenResResult          `json:"Result,omitempty"`
}

type RecordScreenResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *RecordScreenResResponseMetadataError `json:"Error,omitempty"`
}

type RecordScreenResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type RecordScreenResResult struct {

	// 保存录制文件的地址，例如：/sdcard/mp4/record.mp4
	FilePath *string `json:"file_path,omitempty"`
}

type ResetInstanceToFactoryBody struct {

	// REQUIRED; 目标 Id 列表
	ObjectIDList []string `json:"object_id_list"`

	// REQUIRED; 对象的类别，可选枚举值包括： instance（对指定实例为对象进行操作） host（对云机下的所有实例进行操作)
	ObjectType string `json:"object_type"`

	// REQUIRED; 产品 Id
	ProductID string `json:"product_id"`
}

type ResetInstanceToFactoryRes struct {

	// REQUIRED
	ResponseMetadata ResetInstanceToFactoryResResponseMetadata `json:"ResponseMetadata"`
	Result           *ResetInstanceToFactoryResResult          `json:"Result,omitempty"`
}

type ResetInstanceToFactoryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *ResetInstanceToFactoryResResponseMetadataError `json:"Error,omitempty"`
}

type ResetInstanceToFactoryResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ResetInstanceToFactoryResResult struct {

	// 失���的ID列表
	FailedIDList []*ResetInstanceToFactoryResResultFailedIDListItem `json:"failed_id_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 废弃
	PassedIDList []*string `json:"passed_id_list,omitempty"`
}

type ResetInstanceToFactoryResResultFailedIDListItem struct {

	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type ResetInstancesBody struct {

	// REQUIRED; 镜像Id
	ImageID string `json:"image_id"`

	// REQUIRED; 实例ID列表
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`

	// 镜像是否为公共镜像 true（在公共镜像下检索指定的镜像进行重置） false（在当前账号下检索指定的镜像进行重置，默认）
	IsPublicImage *bool `json:"is_public_image,omitempty"`
}

type ResetInstancesRes struct {

	// REQUIRED
	ResponseMetadata ResetInstancesResResponseMetadata `json:"ResponseMetadata"`
	Result           *ResetInstancesResResult          `json:"Result,omitempty"`
}

type ResetInstancesResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *ResetInstancesResResponseMetadataError `json:"Error,omitempty"`
}

type ResetInstancesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ResetInstancesResResult struct {

	// 失败的记录
	FailedList []*ResetInstancesResResultFailedListItem `json:"failed_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`
}

type ResetInstancesResResultFailedListItem struct {

	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type SetInstanceBandwidthBody struct {

	// REQUIRED; 需要设置的带宽大小，2～300Mbps
	BandwidthMbps int32 `json:"BandwidthMbps"`

	// REQUIRED; 实例列表，单次最大不能超过50个实例
	InstanceIDList []string `json:"InstanceIdList"`

	// REQUIRED; 业务ID
	ProductID string `json:"ProductId"`
}

type SetInstanceBandwidthRes struct {

	// REQUIRED
	ResponseMetadata SetInstanceBandwidthResResponseMetadata `json:"ResponseMetadata"`
	Result           *SetInstanceBandwidthResResult          `json:"Result,omitempty"`
}

type SetInstanceBandwidthResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *SetInstanceBandwidthResResponseMetadataError `json:"Error,omitempty"`
}

type SetInstanceBandwidthResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type SetInstanceBandwidthResResult struct {

	// 失败列表
	FailedList    []*SetInstanceBandwidthResResultFailedListItem `json:"FailedList,omitempty"`
	SuccessIDList []*string                                      `json:"SuccessIdList,omitempty"`
}

type SetInstanceBandwidthResResultFailedListItem struct {

	// 操作的ID，云机Id 或 实例Id
	ID *string `json:"Id,omitempty"`

	// 信息
	Msg *string `json:"Msg,omitempty"`
}

type SetInstancePropertiesBody struct {

	// REQUIRED; 实例ID
	InstanceID string `json:"instance_id"`

	// REQUIRED; 属性名和属性值
	Properties []SetInstancePropertiesBodyPropertiesItem `json:"properties"`
}

type SetInstancePropertiesBodyPropertiesItem struct {

	// REQUIRED; 属性名 属性名称长度上限64 不支持输入特殊字符，包括但不限于 +- ;&|(){}[]
	PropertyName string `json:"property_name"`

	// REQUIRED; 属性值 属性值长度上限91 不支持输入特殊字符，包括但不限于 +- ;&|(){}[]
	PropertyValue string `json:"property_value"`
}

type SetInstancePropertiesRes struct {

	// REQUIRED
	ResponseMetadata SetInstancePropertiesResResponseMetadata `json:"ResponseMetadata"`
	Result           *SetInstancePropertiesResResult          `json:"Result,omitempty"`
}

type SetInstancePropertiesResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *SetInstancePropertiesResResponseMetadataError `json:"Error,omitempty"`
}

type SetInstancePropertiesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type SetInstancePropertiesResResult struct {
	InstanceID *string `json:"instance_id,omitempty"`
}

type UnbindInstanceAdbKeyBody struct {

	// REQUIRED
	InstanceID string `json:"instance_id"`

	// REQUIRED
	ProductID string `json:"product_id"`
}

type UnbindInstanceAdbKeyRes struct {

	// REQUIRED
	ResponseMetadata UnbindInstanceAdbKeyResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                 `json:"Result,omitempty"`
}

type UnbindInstanceAdbKeyResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *UnbindInstanceAdbKeyResResponseMetadataError `json:"Error,omitempty"`
}

type UnbindInstanceAdbKeyResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UnbindInstancesAdbKeyBody struct {

	// REQUIRED
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED
	ProductID string `json:"product_id"`
}

type UnbindInstancesAdbKeyRes struct {

	// REQUIRED
	ResponseMetadata UnbindInstancesAdbKeyResResponseMetadata `json:"ResponseMetadata"`
	Result           *UnbindInstancesAdbKeyResResult          `json:"Result,omitempty"`
}

type UnbindInstancesAdbKeyResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *UnbindInstancesAdbKeyResResponseMetadataError `json:"Error,omitempty"`
}

type UnbindInstancesAdbKeyResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UnbindInstancesAdbKeyResResult struct {
	JobID *string `json:"job_id,omitempty"`
}

type UnbindInstancesSecurityGroupBody struct {

	// REQUIRED; 实例 ID 列表
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED; 业务 ID
	ProductID string `json:"product_id"`
}

type UnbindInstancesSecurityGroupRes struct {

	// REQUIRED
	ResponseMetadata UnbindInstancesSecurityGroupResResponseMetadata `json:"ResponseMetadata"`
	Result           *UnbindInstancesSecurityGroupResResult          `json:"Result,omitempty"`
}

type UnbindInstancesSecurityGroupResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                `json:"Version"`
	Error   *UnbindInstancesSecurityGroupResResponseMetadataError `json:"Error,omitempty"`
}

type UnbindInstancesSecurityGroupResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UnbindInstancesSecurityGroupResResult struct {

	// 异步 JobID
	JobID *string `json:"job_id,omitempty"`
}

type UpdateContainerImageBody struct {

	// REQUIRED
	ImageID         string  `json:"image_id"`
	ImageAnnotation *string `json:"image_annotation,omitempty"`
	ImageName       *string `json:"image_name,omitempty"`
}

type UpdateContainerImageRes struct {

	// REQUIRED
	ResponseMetadata UpdateContainerImageResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateContainerImageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *UpdateContainerImageResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateContainerImageResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UpdateInstancePropertyBody struct {

	// REQUIRED
	InstanceProperty UpdateInstancePropertyBodyInstanceProperty `json:"instance_property"`

	// REQUIRED; 目标 Id 列表
	ObjectIDList []string `json:"object_id_list"`

	// REQUIRED; 对象的类别，可选枚举值包括： instance（对指定实例为对象进行操作） host（对云机下的所有实例进行操作)
	ObjectType string `json:"object_type"`

	// REQUIRED; 产品ID
	ProductID string `json:"product_id"`
}

type UpdateInstancePropertyBodyInstanceProperty struct {

	// REQUIRED; 属性名 属性名称长度上限64 不支持输入特殊字符，包括但不限于 +- ;&|(){}[]
	PropertyName string `json:"property_name"`

	// REQUIRED; 属性值 属性值长度上限91 不支持输入特殊字符，包括但不限于 +- ;&|(){}[]
	PropertyValue string `json:"property_value"`
}

type UpdateInstancePropertyRes struct {

	// REQUIRED
	ResponseMetadata UpdateInstancePropertyResResponseMetadata `json:"ResponseMetadata"`
	Result           *UpdateInstancePropertyResResult          `json:"Result,omitempty"`
}

type UpdateInstancePropertyResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *UpdateInstancePropertyResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateInstancePropertyResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UpdateInstancePropertyResResult struct {

	// 失败的ID列表
	FailedIDList []*UpdateInstancePropertyResResultFailedIDListItem `json:"failed_id_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated: 废弃
	PassedIDList []*string `json:"passed_id_list,omitempty"`
}

type UpdateInstancePropertyResResultFailedIDListItem struct {

	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}

type UpgradeInstancesBody struct {

	// REQUIRED; 镜像Id
	ImageID string `json:"ImageId"`

	// REQUIRED; 实例ID列表
	InstanceIDs []string `json:"InstanceIds"`

	// REQUIRED; 产品ID
	ProductID string `json:"ProductId"`

	// 镜像是否为公共镜像 true（在公共镜像下检索指定的镜像进行重置） false（在当前账号下检索指定的镜像进行重置，默认）
	IsPublicImage *bool `json:"IsPublicImage,omitempty"`
}

type UpgradeInstancesRes struct {

	// REQUIRED
	ResponseMetadata UpgradeInstancesResResponseMetadata `json:"ResponseMetadata"`
	Result           *UpgradeInstancesResResult          `json:"Result,omitempty"`
}

type UpgradeInstancesResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                    `json:"Version"`
	Error   *UpgradeInstancesResResponseMetadataError `json:"Error,omitempty"`
}

type UpgradeInstancesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UpgradeInstancesResResult struct {

	// 失败的记录
	FailedList []*UpgradeInstancesResResultFailedListItem `json:"FailedList,omitempty"`

	// 异步JobId
	JobID *string `json:"JobId,omitempty"`
}

type UpgradeInstancesResResultFailedListItem struct {

	// 操作的ID，云机Id 或 实例Id
	ID *string `json:"Id,omitempty"`

	// 信息
	Msg *string `json:"Msg,omitempty"`
}

type WarmRebootInstanceBody struct {

	// REQUIRED; 实例 Id 列表
	InstanceIDList []string `json:"instance_id_list"`

	// REQUIRED; 产品 Id
	ProductID string `json:"product_id"`
}

type WarmRebootInstanceRes struct {

	// REQUIRED
	ResponseMetadata WarmRebootInstanceResResponseMetadata `json:"ResponseMetadata"`
	Result           *WarmRebootInstanceResResult          `json:"Result,omitempty"`
}

type WarmRebootInstanceResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *WarmRebootInstanceResResponseMetadataError `json:"Error,omitempty"`
}

type WarmRebootInstanceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type WarmRebootInstanceResResult struct {

	// Deprecated
	FailIDList []*string `json:"fail_id_list,omitempty"`

	// 失败的记录
	FailedList []*WarmRebootInstanceResResultFailedListItem `json:"failed_list,omitempty"`

	// 异步JobId
	JobID *string `json:"job_id,omitempty"`

	// Deprecated
	SuccessIDList []*string `json:"success_id_list,omitempty"`
}

type WarmRebootInstanceResResultFailedListItem struct {

	// 对象 ID
	ID *string `json:"id,omitempty"`

	// 信息
	Msg *string `json:"msg,omitempty"`
}
type ExportInstance struct{}
type ExecCmdSync struct{}
type DetailSecurityGroup struct{}
type ListAdbKeyBody struct{}
type UnbindInstancesAdbKey struct{}
type GetFileDistributionJobDetail struct{}
type ExportInstanceRes struct{}
type GetJobDetailsBody struct{}
type RebootHost struct{}
type GetInfoAfterOrderQuery struct{}
type ListInstanceBody struct{}
type InstallApplication struct{}
type UnbindInstanceAdbKey struct{}
type GetFileDistributionResultBody struct{}
type ModifyInstanceWindowDisplaySpec struct{}
type BindInstancesAdbKey struct{}
type ListDcCapacity struct{}
type WarmRebootInstance struct{}
type InitializeHost struct{}
type GetInfoAfterOrder struct{}
type DeleteDevices struct{}
type SetInstanceBandwidthQuery struct{}
type ListProduct struct{}
type PowerDownInstanceQuery struct{}
type DetailInstanceBody struct{}
type ColdRebootInstanceQuery struct{}
type ControlApplication struct{}
type ListAdbKey struct{}
type DistributeFileQuery struct{}
type ListPortMappingBody struct{}
type SetInstanceProperties struct{}
type UpdateContainerImage struct{}
type InstallApplicationQuery struct{}
type RecordScreen struct{}
type ExportInstanceBody struct{}
type UnbindInstanceAdbKeyQuery struct{}
type ModifyInstanceWindowDisplaySpecQuery struct{}
type ListHostMetricDataQuery struct{}
type PowerDownInstance struct{}
type ColdRebootInstance struct{}
type PowerUpInstance struct{}
type UpdateInstancePropertyQuery struct{}
type ListPackageQuery struct{}
type GetFileDistributionJobDetailQuery struct{}
type DistributeFileToInstances struct{}
type AcquireIdempotentToken struct{}
type DetailInstance struct{}
type GetFileDistributionResult struct{}
type ListHostBody struct{}
type BindInstancesSecurityGroup struct{}
type PushFileQuery struct{}
type ImportContainerImage struct{}
type ListProductBody struct{}
type CreateDevices struct{}
type ListPackage struct{}
type ListInstanceMetricData struct{}
type BindInstancesSecurityGroupQuery struct{}
type PullFile struct{}
type ModifyInstanceFpsQuery struct{}
type DeleteContainerImagesQuery struct{}
type FixInstancesSGBound struct{}
type FixInstancesSGBoundQuery struct{}
type ListSecurityGroupBody struct{}
type DeleteContainerImages struct{}
type UpdateInstanceProperty struct{}
type GetJobDetails struct{}
type ListContainerImages struct{}
type RecordScreenQuery struct{}
type DistributeFile struct{}
type AdbCommand struct{}
type DistributeFileToInstancesQuery struct{}
type AdbCommandQuery struct{}
type WarmRebootInstanceQuery struct{}
type ResetInstances struct{}
type GetInstanceProperty struct{}
type AcquireIdempotentTokenQuery struct{}
type UnbindInstancesSecurityGroupQuery struct{}
type InitializeHostQuery struct{}
type ExecCmdSyncQuery struct{}
type BindInstanceAdbKey struct{}
type ImportContainerImageQuery struct{}
type ResetInstanceToFactory struct{}
type ListHostMetricData struct{}
type LatestMetricInstance struct{}
type ListHost struct{}
type ResetInstanceToFactoryQuery struct{}
type ListDcCapacityQuery struct{}
type ListSecurityGroup struct{}
type ModifyInstanceFps struct{}
type ListContainerImagesBody struct{}
type LatestMetricInstanceBody struct{}
type PushFile struct{}
type UpgradeInstancesQuery struct{}
type SetInstanceBandwidth struct{}
type GetInstancePropertyQuery struct{}
type PullFileQuery struct{}
type UnbindInstancesAdbKeyQuery struct{}
type DetailSecurityGroupBody struct{}
type ListInstanceMetricDataQuery struct{}
type PowerUpInstanceQuery struct{}
type GetInstancePropertiesQuery struct{}
type SetInstancePropertiesQuery struct{}
type ResetInstancesQuery struct{}
type UpdateContainerImageQuery struct{}
type UnbindInstancesSecurityGroup struct{}
type UpgradeInstances struct{}
type BindInstancesAdbKeyQuery struct{}
type BindInstanceAdbKeyQuery struct{}
type GetInstanceProperties struct{}
type ControlApplicationQuery struct{}
type ListInstance struct{}
type ListPortMapping struct{}
type RebootHostQuery struct{}
type BindInstanceAdbKeyReq struct {
	*BindInstanceAdbKeyQuery
	*BindInstanceAdbKeyBody
}
type ListInstanceMetricDataReq struct {
	*ListInstanceMetricDataQuery
	*ListInstanceMetricDataBody
}
type PowerDownInstanceReq struct {
	*PowerDownInstanceQuery
	*PowerDownInstanceBody
}
type GetJobDetailsReq struct {
	*GetJobDetailsQuery
	*GetJobDetailsBody
}
type GetFileDistributionJobDetailReq struct {
	*GetFileDistributionJobDetailQuery
	*GetFileDistributionJobDetailBody
}
type GetInfoAfterOrderReq struct {
	*GetInfoAfterOrderQuery
	*GetInfoAfterOrderBody
}
type ListProductReq struct {
	*ListProductQuery
	*ListProductBody
}
type AdbCommandReq struct {
	*AdbCommandQuery
	*AdbCommandBody
}
type SetInstanceBandwidthReq struct {
	*SetInstanceBandwidthQuery
	*SetInstanceBandwidthBody
}
type GetInstancePropertyReq struct {
	*GetInstancePropertyQuery
	*GetInstancePropertyBody
}
type ImportContainerImageReq struct {
	*ImportContainerImageQuery
	*ImportContainerImageBody
}
type PushFileReq struct {
	*PushFileQuery
	*PushFileBody
}
type GetFileDistributionResultReq struct {
	*GetFileDistributionResultQuery
	*GetFileDistributionResultBody
}
type ModifyInstanceWindowDisplaySpecReq struct {
	*ModifyInstanceWindowDisplaySpecQuery
	*ModifyInstanceWindowDisplaySpecBody
}
type PullFileReq struct {
	*PullFileQuery
	*PullFileBody
}
type GetInstancePropertiesReq struct {
	*GetInstancePropertiesQuery
	*GetInstancePropertiesBody
}
type DetailSecurityGroupReq struct {
	*DetailSecurityGroupQuery
	*DetailSecurityGroupBody
}
type RecordScreenReq struct {
	*RecordScreenQuery
	*RecordScreenBody
}
type ExecCmdSyncReq struct {
	*ExecCmdSyncQuery
	*ExecCmdSyncBody
}
type ListHostMetricDataReq struct {
	*ListHostMetricDataQuery
	*ListHostMetricDataBody
}
type UnbindInstancesSecurityGroupReq struct {
	*UnbindInstancesSecurityGroupQuery
	*UnbindInstancesSecurityGroupBody
}
type ResetInstanceToFactoryReq struct {
	*ResetInstanceToFactoryQuery
	*ResetInstanceToFactoryBody
}
type ListAdbKeyReq struct {
	*ListAdbKeyQuery
	*ListAdbKeyBody
}
type ListContainerImagesReq struct {
	*ListContainerImagesQuery
	*ListContainerImagesBody
}
type DeleteContainerImagesReq struct {
	*DeleteContainerImagesQuery
	*DeleteContainerImagesBody
}
type UnbindInstanceAdbKeyReq struct {
	*UnbindInstanceAdbKeyQuery
	*UnbindInstanceAdbKeyBody
}
type DistributeFileReq struct {
	*DistributeFileQuery
	*DistributeFileBody
}
type DeleteDevicesReq struct {
	*DeleteDevicesQuery
	*DeleteDevicesBody
}
type ColdRebootInstanceReq struct {
	*ColdRebootInstanceQuery
	*ColdRebootInstanceBody
}
type WarmRebootInstanceReq struct {
	*WarmRebootInstanceQuery
	*WarmRebootInstanceBody
}
type ListHostReq struct {
	*ListHostQuery
	*ListHostBody
}
type ListSecurityGroupReq struct {
	*ListSecurityGroupQuery
	*ListSecurityGroupBody
}
type ListInstanceReq struct {
	*ListInstanceQuery
	*ListInstanceBody
}
type UpgradeInstancesReq struct {
	*UpgradeInstancesQuery
	*UpgradeInstancesBody
}
type DetailInstanceReq struct {
	*DetailInstanceQuery
	*DetailInstanceBody
}
type ModifyInstanceFpsReq struct {
	*ModifyInstanceFpsQuery
	*ModifyInstanceFpsBody
}
type InitializeHostReq struct {
	*InitializeHostQuery
	*InitializeHostBody
}
type FixInstancesSGBoundReq struct {
	*FixInstancesSGBoundQuery
	*FixInstancesSGBoundBody
}
type BindInstancesSecurityGroupReq struct {
	*BindInstancesSecurityGroupQuery
	*BindInstancesSecurityGroupBody
}
type SetInstancePropertiesReq struct {
	*SetInstancePropertiesQuery
	*SetInstancePropertiesBody
}
type UnbindInstancesAdbKeyReq struct {
	*UnbindInstancesAdbKeyQuery
	*UnbindInstancesAdbKeyBody
}
type PowerUpInstanceReq struct {
	*PowerUpInstanceQuery
	*PowerUpInstanceBody
}
type AcquireIdempotentTokenReq struct {
	*AcquireIdempotentTokenQuery
	*AcquireIdempotentTokenBody
}
type UpdateInstancePropertyReq struct {
	*UpdateInstancePropertyQuery
	*UpdateInstancePropertyBody
}
type ListPackageReq struct {
	*ListPackageQuery
	*ListPackageBody
}
type ResetInstancesReq struct {
	*ResetInstancesQuery
	*ResetInstancesBody
}
type UpdateContainerImageReq struct {
	*UpdateContainerImageQuery
	*UpdateContainerImageBody
}
type CreateDevicesReq struct {
	*CreateDevicesQuery
	*CreateDevicesBody
}
type ListDcCapacityReq struct {
	*ListDcCapacityQuery
	*ListDcCapacityBody
}
type LatestMetricInstanceReq struct {
	*LatestMetricInstanceQuery
	*LatestMetricInstanceBody
}
type InstallApplicationReq struct {
	*InstallApplicationQuery
	*InstallApplicationBody
}
type RebootHostReq struct {
	*RebootHostQuery
	*RebootHostBody
}
type ControlApplicationReq struct {
	*ControlApplicationQuery
	*ControlApplicationBody
}
type DistributeFileToInstancesReq struct {
	*DistributeFileToInstancesQuery
	*DistributeFileToInstancesBody
}
type BindInstancesAdbKeyReq struct {
	*BindInstancesAdbKeyQuery
	*BindInstancesAdbKeyBody
}
type ExportInstanceReq struct {
	*ExportInstanceQuery
	*ExportInstanceBody
}
type ListPortMappingReq struct {
	*ListPortMappingQuery
	*ListPortMappingBody
}

type ReconfigureDevicesPackageBody struct {

	// REQUIRED; 产品ID
	ProductID string `json:"ProductId"`

	// REQUIRED; 配置参数
	ReconfigureParam []ReconfigureDevicesPackageBodyReconfigureParamItem `json:"ReconfigureParam"`
}

type ReconfigureDevicesPackageBodyReconfigureParamItem struct {

	// REQUIRED; 云机 ID
	DeviceID string `json:"DeviceId"`

	// REQUIRED; 实例配置
	InstancesSpec []ReconfigureDevicesPackageBodyReconfigureParamPropertiesItemsItem `json:"InstancesSpec"`

	// REQUIRED; 要切换的目标套餐 ID
	PackageID string `json:"PackageId"`

	// 超时时间, 300-1200s, 默认600s
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty"`
}

type ReconfigureDevicesPackageBodyReconfigureParamPropertiesItemsItem struct {

	// REQUIRED
	BandwidthMbps int32 `json:"BandwidthMbps"`

	// REQUIRED
	DisplayResolutionX int32 `json:"DisplayResolutionX"`

	// REQUIRED
	DisplayResolutionY int32 `json:"DisplayResolutionY"`

	// REQUIRED; 镜像 ID
	ImageID string `json:"ImageId"`

	// REQUIRED; 云机 ID
	InstanceID string `json:"InstanceId"`

	// Dictionary of
	Properties map[string]*string `json:"Properties,omitempty"`
}

type ReconfigureDevicesPackageRes struct {

	// REQUIRED
	ResponseMetadata ReconfigureDevicesPackageResResponseMetadata `json:"ResponseMetadata"`
	Result           *ReconfigureDevicesPackageResResult          `json:"Result,omitempty"`
}

type ReconfigureDevicesPackageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                             `json:"Version"`
	Error   *ReconfigureDevicesPackageResResponseMetadataError `json:"Error,omitempty"`
}

type ReconfigureDevicesPackageResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ReconfigureDevicesPackageResResult struct {

	// 任务ID
	JobID *string `json:"JobId,omitempty"`

	// 配置记录
	Records []*ReconfigureDevicesPackageResResultRecordsItem `json:"Records,omitempty"`
}

type ReconfigureDevicesPackageResResultRecordsItem struct {

	// REQUIRED; 云机 ID
	DeviceID string                                            `json:"DeviceId"`
	Err      *ReconfigureDevicesPackageResResultRecordsItemErr `json:"Err,omitempty"`

	// 任务 ID
	TaskID *string `json:"TaskId,omitempty"`
}

type ReconfigureDevicesPackageResResultRecordsItemErr struct {

	// 错误码, 0 表示成功
	Code *int32 `json:"Code,omitempty"`

	// 错误码字符串
	CodeStr *string `json:"CodeStr,omitempty"`

	// 错误信息
	Msg *string `json:"Msg,omitempty"`
}

type ListTaskInfoResResultRowItemStatus int32

type ListTaskInfoQuery struct {

	// REQUIRED
	ProductID    string  `json:"product_id" query:"product_id"`
	Count        *int32  `json:"count,omitempty" query:"count"`
	CreateAfter  *int32  `json:"create_after,omitempty" query:"create_after"`
	CreateBefore *int32  `json:"create_before,omitempty" query:"create_before"`
	DeviceID     *string `json:"device_id,omitempty" query:"device_id"`
	InstanceID   *string `json:"instance_id,omitempty" query:"instance_id"`

	// request id
	JobID  *string `json:"job_id,omitempty" query:"job_id"`
	Offset *int32  `json:"offset,omitempty" query:"offset"`

	// 逗号分隔
	StatusIn *string `json:"status_in,omitempty" query:"status_in"`

	// task id
	TaskID *string `json:"task_id,omitempty" query:"task_id"`

	// 逗号分隔
	TypeIn       *string `json:"type_in,omitempty" query:"type_in"`
	UpdateAfter  *int32  `json:"update_after,omitempty" query:"update_after"`
	UpdateBefore *int32  `json:"update_before,omitempty" query:"update_before"`
}

type ListTaskInfoRes struct {

	// REQUIRED
	ResponseMetadata ListTaskInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListTaskInfoResResult          `json:"Result,omitempty"`
}

type ListTaskInfoResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *ListTaskInfoResResponseMetadataError `json:"Error,omitempty"`
}

type ListTaskInfoResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	CodeN   *int32  `json:"CodeN,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListTaskInfoResResult struct {
	Row   []*ListTaskInfoResResultRowItem `json:"row,omitempty"`
	Total *int32                          `json:"total,omitempty"`
}

type ListTaskInfoResResultRowItem struct {
	CreateTime *int32                              `json:"create_time,omitempty"`
	DeviceID   *string                             `json:"device_id,omitempty"`
	ExpireTime *int32                              `json:"expire_time,omitempty"`
	InstanceID *string                             `json:"instance_id,omitempty"`
	JobID      *string                             `json:"job_id,omitempty"`
	Msg        *string                             `json:"msg,omitempty"`
	Status     *ListTaskInfoResResultRowItemStatus `json:"status,omitempty"`
	StatusStr  *string                             `json:"status_str,omitempty"`
	TaskID     *string                             `json:"task_id,omitempty"`
	Type       *string                             `json:"type,omitempty"`
	UpdateTime *int32                              `json:"update_time,omitempty"`
}
