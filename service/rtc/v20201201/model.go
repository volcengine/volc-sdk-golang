package rtc_v20201201

type AddBusinessIDBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 业务 ID，同一个 AppId 下不允许重复。长度不超过 64 个字符，取值范围为：数字、字母、_、-、@。
	BusinessID string `json:"BusinessId"`

	// REQUIRED; 业务名称，同一个 AppId 下不允许重复。长度不超过 24 个字符，取值范围为：数字、中文字符、字母和下划线。
	Remarks string `json:"Remarks"`
}

type AddBusinessIDRes struct {

	// REQUIRED
	ResponseMetadata AddBusinessIDResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result string `json:"Result"`
}

type AddBusinessIDResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *AddBusinessIDResResponseMetadataError `json:"Error,omitempty"`
}

// AddBusinessIDResResponseMetadataError - 仅在请求失败时返回。
type AddBusinessIDResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type BanRoomUserBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 指定房间 ID
	RoomID string `json:"RoomId"`

	// 封禁时长，单位为秒，取值范围为[60,259200]。 若传入值为空或 0表示允许用户重新进房。 若传入值大于0，且小于60，自动修改为60。 若传入值大于259200，自动修改为259200。
	ForbiddenInterval *int32 `json:"ForbiddenInterval,omitempty"`

	// 希望封禁用户的 ID
	UserID *string `json:"UserId,omitempty"`
}

type BanRoomUserRes struct {

	// REQUIRED
	ResponseMetadata BanRoomUserResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BanRoomUserResResult `json:"Result"`
}

type BanRoomUserResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error BanRoomUserResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type BanRoomUserResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

type BanRoomUserResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 "success", 失败时为空。
	Message string `json:"message"`
}

type BanUserStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 用于校验当前账号是否具有封禁权限的 Token，生成方式与加入房间时的 Token 生成方式一致
	Token string `json:"Token"`

	// REQUIRED; 你希望封禁音/视频流的用户的 ID
	UserID string `json:"UserId"`

	// 是否封禁音频流，置为 true 时，表示封禁音频流
	Audio *bool `json:"Audio,omitempty"`

	// 是否封禁视频流，置为 true 时，表示封禁视频流
	Video *bool `json:"Video,omitempty"`
}

type BanUserStreamRes struct {

	// REQUIRED
	ResponseMetadata BanUserStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *BanUserStreamResResult          `json:"Result,omitempty"`
}

type BanUserStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error BanUserStreamResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type BanUserStreamResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

type BanUserStreamResResult struct {

	// REQUIRED
	Message string `json:"message"`
}

type BatchSendRoomUnicastBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 是否为二进制消息。
	Binary bool `json:"Binary"`

	// REQUIRED; 业务服务端的唯一标识。 命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。 在一个 AppID 下，不能和真实用户用于实时消息通信的 user_ID 重复； 建议使用固定的 ID 的发送消息。
	From string `json:"From"`

	// REQUIRED; 点对点消息 。如果是二进制消息，需进行 base64 编码
	Message string `json:"Message"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 消息接收者的 user_ID
	To []string `json:"To"`
}

type BatchSendRoomUnicastRes struct {

	// REQUIRED
	ResponseMetadata BatchSendRoomUnicastResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 返回值 Result 仅在请求成功时返回消息 success ，表示服务端成功接收到消息，失败时为空。
	Result BatchSendRoomUnicastResResult `json:"Result"`
}

type BatchSendRoomUnicastResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *BatchSendRoomUnicastResResponseMetadataError `json:"Error,omitempty"`
}

// BatchSendRoomUnicastResResponseMetadataError - 仅在请求失败时返回。
type BatchSendRoomUnicastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// BatchSendRoomUnicastResResult - 返回值 Result 仅在请求成功时返回消息 success ，表示服务端成功接收到消息，失败时为空。
type BatchSendRoomUnicastResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 Success，失败时为空。
	Message string `json:"Message"`
}

type CreateAppBody struct {

	// REQUIRED; * 应用名称
	// * 命名规则：字符串中只能包含中文字符、英文大小写字符、数字和下划线；长度不能超过24个字符。
	AppName string `json:"AppName"`
}

type CreateAppRes struct {

	// REQUIRED
	ResponseMetadata CreateAppResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateAppResResult `json:"Result"`
}

type CreateAppResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *CreateAppResResponseMetadataError `json:"Error,omitempty"`
}

// CreateAppResResponseMetadataError - 仅在请求失败时返回。
type CreateAppResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type CreateAppResResult struct {

	// REQUIRED; 应用 ID
	AppID string `json:"AppId"`

	// REQUIRED; 主 AppKey，用于生成 Token
	AppKey string `json:"AppKey"`

	// REQUIRED; 创建时间
	CreateDate string `json:"CreateDate"`

	// REQUIRED; 计费实例 ID
	InstanceID string `json:"InstanceId"`

	// REQUIRED; * 服务状态。枚举值为：
	// * 0: 创建中——指 AppID 在初始化
	// * 1: 运行中——指当前 AppID 为正常服务状态
	// * 6: 欠费关停
	// * 98: 已停用——指调用服务端停用接口，当前 AppID 被设置为不可用状态
	InstanceStatus int32 `json:"InstanceStatus"`

	// REQUIRED; 应用名称
	Name string `json:"Name"`

	// REQUIRED; 应用所属的账号 ID
	Owner string `json:"Owner"`

	// REQUIRED; 副 AppKey，启用后可用于生成 Token
	SecondaryAppKey string `json:"SecondaryAppKey"`
}

type CreateByteplusAppRes struct {

	// REQUIRED
	ResponseMetadata CreateByteplusAppResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateByteplusAppResResult `json:"Result"`
}

type CreateByteplusAppResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *CreateByteplusAppResResponseMetadataError `json:"Error,omitempty"`
}

// CreateByteplusAppResResponseMetadataError - 仅在请求失败时返回。
type CreateByteplusAppResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type CreateByteplusAppResResult struct {

	// REQUIRED
	AppID string `json:"AppId"`

	// REQUIRED
	AppKey string `json:"AppKey"`

	// REQUIRED
	CreateDate string `json:"CreateDate"`

	// REQUIRED
	Enabled bool `json:"Enabled"`

	// REQUIRED
	InstanceID string `json:"InstanceId"`

	// REQUIRED
	InstanceStatus int32 `json:"InstanceStatus"`

	// REQUIRED
	LicenseStatus int32 `json:"LicenseStatus"`

	// REQUIRED
	LicenseType int32 `json:"LicenseType"`

	// REQUIRED
	Name string `json:"Name"`

	// REQUIRED
	Owner string `json:"Owner"`

	// REQUIRED
	SecondaryAppKey string `json:"SecondaryAppKey"`

	// REQUIRED
	Status int32 `json:"Status"`
}

type CreateCallbackBody struct {

	// REQUIRED
	AppID string `json:"AppID"`

	// REQUIRED
	BusinessID string `json:"BusinessId"`

	// REQUIRED
	EventList []string `json:"EventList"`

	// REQUIRED
	SecretKey string `json:"SecretKey"`

	// REQUIRED
	URL string `json:"URL"`
}

type CreateCallbackRes struct {

	// REQUIRED
	ResponseMetadata CreateCallbackResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result string `json:"Result"`
}

type CreateCallbackResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *CreateCallbackResResponseMetadataError `json:"Error,omitempty"`
}

// CreateCallbackResResponseMetadataError - 仅在请求失败时返回。
type CreateCallbackResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type CreateFailRecoveryPolicyBody struct {

	// REQUIRED; 应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 故障线路
	FromVendor string `json:"FromVendor"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 切换至线路
	ToVendor string `json:"ToVendor"`
}

type CreateFailRecoveryPolicyRes struct {

	// REQUIRED
	ResponseMetadata CreateFailRecoveryPolicyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateFailRecoveryPolicyResResult `json:"Result,omitempty"`
}

type CreateFailRecoveryPolicyResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *CreateFailRecoveryPolicyResResponseMetadataError `json:"Error,omitempty"`
}

// CreateFailRecoveryPolicyResResponseMetadataError - 仅在请求失败时返回。
type CreateFailRecoveryPolicyResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// CreateFailRecoveryPolicyResResult - 视请求的接口而定
type CreateFailRecoveryPolicyResResult struct {

	// REQUIRED; 创建的策略的详细信息
	Policy CreateFailRecoveryPolicyResResultPolicy `json:"Policy"`
}

// CreateFailRecoveryPolicyResResultPolicy - 创建的策略的详细信息
type CreateFailRecoveryPolicyResResultPolicy struct {

	// REQUIRED; RTC应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 创建时间（毫秒）
	CreateTime int32 `json:"CreateTime"`

	// REQUIRED; 故障线路
	FromVendor string `json:"FromVendor"`

	// REQUIRED; 数据条目的ID。后面删除和修改都需要指定该ID
	ID int32 `json:"ID"`

	// REQUIRED; 0~100的整数，故障线路恢复的百分比
	RecoveryPercentage int32 `json:"RecoveryPercentage"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 切换至线路
	ToVendor string `json:"ToVendor"`

	// REQUIRED; 更新时间（毫秒）
	UpdateTime string `json:"UpdateTime"`
}

type CreateVendorPolicyBody struct {

	// REQUIRED; RTC应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 业务标识
	BusinessID string `json:"BusinessID"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 【json类型】各线路放量比例，例： {"agora": 50, "trtc": 50}
	VendorProportion string `json:"VendorProportion"`
}

type CreateVendorPolicyRes struct {

	// REQUIRED
	ResponseMetadata CreateVendorPolicyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateVendorPolicyResResult `json:"Result,omitempty"`
}

type CreateVendorPolicyResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *CreateVendorPolicyResResponseMetadataError `json:"Error,omitempty"`
}

// CreateVendorPolicyResResponseMetadataError - 仅在请求失败时返回。
type CreateVendorPolicyResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// CreateVendorPolicyResResult - 视请求的接口而定
type CreateVendorPolicyResResult struct {

	// REQUIRED; 创建的策略详情
	Policy CreateVendorPolicyResResultPolicy `json:"Policy"`
}

// CreateVendorPolicyResResultPolicy - 创建的策略详情
type CreateVendorPolicyResResultPolicy struct {

	// REQUIRED; 应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 应用标识
	BusinessID string `json:"BusinessID"`

	// REQUIRED; 创建时间（毫秒）
	CreateTime int32 `json:"CreateTime"`

	// REQUIRED; 数据条目的ID。后面删除和修改都需要指定该ID
	ID int32 `json:"ID"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 更新时间（毫秒）
	UpdateTime int32 `json:"UpdateTime"`

	// REQUIRED; 各线路放量比例
	VendorProportion string `json:"VendorProportion"`
}

type DeleteBusinessIDBody struct {

	// REQUIRED; 应用的唯一标识
	AppID string `json:"AppId"`

	// REQUIRED; 待删除的业务标识对应的业务 ID
	BusinessID string `json:"BusinessId"`
}

type DeleteBusinessIDRes struct {

	// REQUIRED
	ResponseMetadata DeleteBusinessIDResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result string `json:"Result"`
}

type DeleteBusinessIDResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *DeleteBusinessIDResResponseMetadataError `json:"Error,omitempty"`
}

// DeleteBusinessIDResResponseMetadataError - 仅在请求失败时返回。
type DeleteBusinessIDResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type DeleteCallbackQuery struct {

	// REQUIRED
	AppID      string  `json:"AppId" query:"AppId"`
	BusinessID *string `json:"BusinessId,omitempty" query:"BusinessId"`
}

type DeleteCallbackRes struct {

	// REQUIRED
	ResponseMetadata DeleteCallbackResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 结果
	Result string `json:"Result"`
}

type DeleteCallbackResResponseMetadata struct {

	// REQUIRED; 接口名称
	Action string `json:"Action"`

	// REQUIRED; 区域
	Region string `json:"Region"`

	// REQUIRED; 请求ID
	RequestID string `json:"RequestId"`

	// REQUIRED; 服务Provider
	Service string `json:"Service"`

	// REQUIRED; 接口版本
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *DeleteCallbackResResponseMetadataError `json:"Error,omitempty"`
}

// DeleteCallbackResResponseMetadataError - 仅在请求失败时返回。
type DeleteCallbackResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type DeleteFailRecoveryPolicyBody struct {

	// REQUIRED
	ID int32 `json:"ID"`
}

type DeleteFailRecoveryPolicyRes struct {

	// REQUIRED
	ResponseMetadata DeleteFailRecoveryPolicyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteFailRecoveryPolicyResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *DeleteFailRecoveryPolicyResResponseMetadataError `json:"Error,omitempty"`
}

// DeleteFailRecoveryPolicyResResponseMetadataError - 仅在请求失败时返回。
type DeleteFailRecoveryPolicyResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type DismissRoomBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`
}

type DismissRoomRes struct {

	// REQUIRED
	ResponseMetadata DismissRoomResResponseMetadata `json:"ResponseMetadata"`
	Result           *DismissRoomResResult          `json:"Result,omitempty"`
}

type DismissRoomResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error DismissRoomResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DismissRoomResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

type DismissRoomResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 "success", 失败时为空。
	Message string `json:"message"`
}

type Event20231101 struct {

	// 原因，只有error事件有
	Cause *string `json:"Cause,omitempty"`

	// 事件描述
	Description *string `json:"Description,omitempty"`

	// 原始日志信息，这一期不做在前端展示
	Detail *string `json:"Detail,omitempty"`

	// 展示名
	DisplayName *string `json:"DisplayName,omitempty"`

	// 事件级别 Info、Warning、Error
	Level *int64 `json:"Level,omitempty"`

	// 事件名，英文
	Name *string `json:"Name,omitempty"`

	// 子事件，构成该聚合事件的事件，这一期不做在前端展示
	SubEvents []*Event20231101SubEventsItem `json:"SubEvents,omitempty"`

	// 标签
	Tags []*string `json:"Tags,omitempty"`

	// 事件时间，若为聚合类事件，取最大时间
	Time *int64 `json:"Time,omitempty"`
}

type Event20231101SubEventsItem struct {

	// 原因，只有error事件有
	Cause *string `json:"Cause,omitempty"`

	// 事件描述
	Description *string `json:"Description,omitempty"`

	// 原始日志信息，这一期不做在前端展示
	Detail *string `json:"Detail,omitempty"`

	// 展示名
	DisplayName *string `json:"DisplayName,omitempty"`

	// 事件级别 Info、Warning、Error
	Level *int64 `json:"Level,omitempty"`

	// 事件名，英文
	Name *string `json:"Name,omitempty"`

	// 子事件，构成该聚合事件的事件，这一期不做在前端展示
	SubEvents []*Event20231101 `json:"SubEvents,omitempty"`

	// 标签
	Tags []*string `json:"Tags,omitempty"`

	// 事件时间，若为聚合类事件，取最大时间
	Time *int64 `json:"Time,omitempty"`
}

type GetAllBusinessIDBody struct {

	// REQUIRED; 应用的唯一标识
	AppID string `json:"AppID"`

	// 偏移量，单位为条，默认值为 0。
	Limit *string `json:"Limit,omitempty"`

	// 返回的业务标识上限。最大值为 100，默认值为 100。
	Offset *string `json:"Offset,omitempty"`
}

type GetAllBusinessIDRes struct {

	// REQUIRED
	ResponseMetadata GetAllBusinessIDResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetAllBusinessIDResResult `json:"Result"`
}

type GetAllBusinessIDResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetAllBusinessIDResResponseMetadataError `json:"Error,omitempty"`
}

// GetAllBusinessIDResResponseMetadataError - 仅在请求失败时返回。
type GetAllBusinessIDResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetAllBusinessIDResResult struct {

	// REQUIRED; 业务标识相关信息
	BusinessRecords []GetAllBusinessIDResResultBusinessRecordsItem `json:"BusinessRecords"`

	// REQUIRED; 该 AppId 下业务标识总数
	TotalSize int32 `json:"TotalSize"`
}

type GetAllBusinessIDResResultBusinessRecordsItem struct {

	// REQUIRED; 应用的唯一标识
	AppID string `json:"AppId"`

	// REQUIRED; 业务 ID
	BusinessID string `json:"BusinessId"`

	// REQUIRED; 该业务标识创建时间戳，格式为 RFC3339，单位为秒
	CreatedTime string `json:"CreatedTime"`

	// REQUIRED; 业务名称
	Remarks string `json:"Remarks"`
}

type GetCallbackQuery struct {

	// REQUIRED
	AppID string `json:"AppId" query:"AppId"`
}

type GetCallbackRes struct {

	// REQUIRED
	ResponseMetadata GetCallbackResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetCallbackResResult `json:"Result"`
}

type GetCallbackResResponseMetadata struct {

	// REQUIRED; 接口名称
	Action string `json:"Action"`

	// REQUIRED; 区域
	Region string `json:"Region"`

	// REQUIRED; 请求ID
	RequestID string `json:"RequestId"`

	// REQUIRED; 服务Provider
	Service string `json:"Service"`

	// REQUIRED; 接口版本
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetCallbackResResponseMetadataError `json:"Error,omitempty"`
}

// GetCallbackResResponseMetadataError - 仅在请求失败时返回。
type GetCallbackResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetCallbackResResult struct {

	// REQUIRED
	AccountID string `json:"AccountId"`

	// REQUIRED
	AppID string `json:"AppId"`

	// REQUIRED
	BusinessID string `json:"BusinessId"`

	// REQUIRED
	Callbacks []GetCallbackResResultCallbacksItem `json:"Callbacks"`

	// REQUIRED
	CreatedAt string `json:"CreatedAt"`

	// REQUIRED
	EventList string `json:"EventList"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	SecretKey string `json:"SecretKey"`

	// REQUIRED
	URL string `json:"URL"`

	// REQUIRED
	UpdatedAt string `json:"UpdatedAt"`
}

type GetCallbackResResultCallbacksItem struct {
	AccountID  *string   `json:"AccountId,omitempty"`
	AppID      *string   `json:"AppId,omitempty"`
	BusinessID *string   `json:"BusinessId,omitempty"`
	CreatedAt  *string   `json:"CreatedAt,omitempty"`
	EventList  []*string `json:"EventList,omitempty"`
	Region     *string   `json:"Region,omitempty"`
	SecretKey  *string   `json:"SecretKey,omitempty"`
	URL        *string   `json:"URL,omitempty"`
	UpdatedAt  *string   `json:"UpdatedAt,omitempty"`
}

type GetFailRecoveryPoliciesBody struct {

	// REQUIRED; RTC应用id
	AppID string `json:"AppID"`

	// REQUIRED; 故障线路
	FromVendor string `json:"FromVendor"`

	// REQUIRED; 排序方式，可从"createTime", "updateTime", "priority"中取值，默认为"updateTime"
	OrderBy string `json:"OrderBy"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 目标线路
	ToVendor string `json:"ToVendor"`
}

type GetFailRecoveryPoliciesRes struct {

	// REQUIRED
	ResponseMetadata GetFailRecoveryPoliciesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetFailRecoveryPoliciesResResult `json:"Result,omitempty"`
}

type GetFailRecoveryPoliciesResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetFailRecoveryPoliciesResResponseMetadataError `json:"Error,omitempty"`
}

// GetFailRecoveryPoliciesResResponseMetadataError - 仅在请求失败时返回。
type GetFailRecoveryPoliciesResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// GetFailRecoveryPoliciesResResult - 视请求的接口而定
type GetFailRecoveryPoliciesResResult struct {

	// REQUIRED; 切线策略列表
	Policies []GetFailRecoveryPoliciesResResultPoliciesItem `json:"Policies"`
}

type GetFailRecoveryPoliciesResResultPoliciesItem struct {

	// REQUIRED; RTC应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 创建时间（毫秒）
	CreateTime int32 `json:"CreateTime"`

	// REQUIRED; 故障线路
	FromVendor string `json:"FromVendor"`

	// REQUIRED; 数据条目的ID。后面删除和修改都需要指定该ID
	ID int32 `json:"ID"`

	// REQUIRED; 0~100的整数，故障线路恢复的百分比
	RecoveryPercentage int32 `json:"RecoveryPercentage"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 切换至线路
	ToVendor string `json:"ToVendor"`

	// REQUIRED; 更新时间（毫秒）
	UpdateTime string `json:"UpdateTime"`
}

type GetPushMixedStreamToCDNTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 要查询的转推直播任务 ID。通过服务端发起时，该值为调用 OpenAPI 时传入的 TaskId。通过客户端 SDK 发起时，TaskId 是按照 userId@@taskId 格式拼接而成的字符串；当传入的 taskId 为空时，这里的
	// TaskId 为 userId。
	// TaskId和UserId均为非必填参数，但是你需要至少填一个参数以保证可以正常发起请求。
	TaskID *string `json:"TaskId,omitempty" query:"TaskId"`

	// 客户端发起转推任务的用户 ID。
	// 你在客户端发起多个任务，当使用该接口进行查询时：
	// * 查询第一个任务时，UserId 可以传入发起转推任务的用户 UserId，TaskId 传入空字符串；或直接将该用户的 UserId 传入 TaskId。
	// * 查询第二个以上任务时，UserId 和 TaskId 为发起转推任务的用户 UserId 和 TaskId。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type GetPushMixedStreamToCDNTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPushMixedStreamToCDNTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetPushMixedStreamToCDNTaskResResult `json:"Result"`
}

type GetPushMixedStreamToCDNTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetPushMixedStreamToCDNTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResponseMetadataError - 仅在请求失败时返回。
type GetPushMixedStreamToCDNTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResult struct {

	// 合流转推任务信息
	PushMixedStreamToCDNTask *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask `json:"PushMixedStreamToCDNTask,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask - 合流转推任务信息
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask struct {

	// 控制选项
	Control *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControl `json:"Control,omitempty"`
	Encode  *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskEncode  `json:"Encode,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime        *int64                                                                      `json:"EndTime,omitempty"`
	ExcludeStreams *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreams `json:"ExcludeStreams,omitempty"`
	Layout         *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayout         `json:"Layout,omitempty"`

	// 推流状态
	// * 0：运行中，未获取到任务状态，建议稍后重新查询
	// * 1：未开始推流
	// * 2：首次连接 CDN 服务
	// * 3：正在重连 CDN 服务
	// * 4：连接 CDN 服务成功，正在尝试推流。
	// * 5：连接 CDN 服务成功，推流成功
	// * 6：已停止推流。
	// 仅当Status=2 时，PushStreamState 有实际意义；当Status=3 时，PushStreamState=6; status 为其他值时，PushStreamState 均为0。
	PushStreamState *int32 `json:"PushStreamState,omitempty"`

	// 推流 CDN 地址。仅支持 RTMP 协议。
	PushURL *string `json:"PushURL,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0：未知异常状态
	// * 1：未开始
	// * 2： 运行中
	// * 3： 已结束
	// * 4： 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：客户端用户主动退房/调用停止转推接口
	StopReason    *string                                                                    `json:"StopReason,omitempty"`
	TargetStreams *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreams `json:"TargetStreams,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControl - 控制选项
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControl struct {

	// 选择补帧模式。默认值为0，可以取0和1。0为补最后一帧，1为补黑帧。值不合法时会自动调整为默认值。
	// 自动布局模式下，没有补帧的逻辑。
	// 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系:
	// 你可以在 Region 中传入 Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// （仅对录制有效）最大录制时长，单位为秒。默认值为 0。0 表示不限制录制时长。
	MaxRecordTime *int32 `json:"MaxRecordTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                                                            `json:"PushStreamMode,omitempty"`
	SpatialConfig  *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfig `json:"SpatialConfig,omitempty"`

	// (仅对转推直播有效）有效说话音量大小。取值范围为[0,255]，默认值为0。
	TalkVolume *int32 `json:"TalkVolume,omitempty"`

	// (仅对转推直播有效）用户说话音量的回调间隔，单位为秒，取值范围为[1,∞]，默认值为 2。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfig struct {

	// 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward = [1, 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
	AudienceSpatialOrientation *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。
	// * false：关闭。
	// * true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfigAudienceSpatialOrientation - 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward
// = [1, 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfigAudienceSpatialOrientation struct {

	// 头顶朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 前方朝向
	Up []*float32 `json:"Up,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
	// 当AudioProfile=1且 AudioChannels = 1 时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	// 当AudioProfile=1且 AudioChannels = 2 时：
	// * 若输入参数取值范围为 [32,128]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [128,192]，编码码率固定为128。
	// 当AudioProfile=2时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。
	// * 1：单声道
	// * 2：双声道。 默认值为 2。值不合法时，自动调整为默认值。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。默认值为 0，此时使用 aac 编码协议。目前只支持 aac。值不合法时，自动调整为默认值。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型，在使用 aac 编码时生效。取值范围为 {0, 1, 2}。- 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。取 2 时，只支持输出流音频声道数为双声道。 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。默认值 48000，取值为 [32000,44100,48000]，单位是 Hz。值不合法时，自动调整为默认值。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 输出视频码率。取值范围 [1,10000]，单位为 Kbps，默认值为自适应。值不合法时，自动调整为默认值。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。默认值为 0，可以取 0 或 1。取 0 时使用 H.264,取 1 时使用 ByteVC1 编码器。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出视频帧率。默认为 15，取值范围为 [1,60]。值不合法时，自动调整为默认值。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。默认为 4，取值范围为 [1,5]，单位为秒。值不合法时，自动调整为默认值。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 输出画面的高度，范围为[2, 1920]，必须是偶数，默认值为480。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面高度以 canvas 中的 Height 为主。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 输出画面的宽度。默认值为640，范围为 [2, 1920]，必须是偶数。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面宽度以 canvas 中的 Width 为主。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayout struct {
	CustomLayout *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                                                             `json:"LayoutMode,omitempty"`
	MainVideoStream *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutMainVideoStream `json:"MainVideoStream,omitempty"`

	// 在垂直布局模式下生效，指定主画面流的属性。如果此参数为空，则主画面为随机的一路流。该参数已废弃，当前版本 MainVideoStreamIndex 依然可用，但我们强烈建议你使用 MainVideoStream 参数。如果你同时指定了 MainVideoStream
	// 和 MainVideoStreamIndex 的值，此时只有 MainVideoStream 生效。
	MainVideoStreamIndex *int32 `json:"MainVideoStreamIndex,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayout struct {
	Canvas *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即
	// Regions.Region.StreamIndex 要与
	// TargetStreams.StreamList.Stream.Index 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，用十六进制颜色码（HEX）表示。例如，#FFFFFF 表示纯白，#000000 表示纯黑。默认值为 #000000。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 值不合法时，自动调整为默认值。
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushSingleStreamToCDNTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 要查询的转推直播任务 ID。通过服务端发起时，该值为调用 OpenAPI 时传入的 TaskId。通过客户端 SDK 发起时，TaskId 是按照 userId@@taskId 格式拼接而成的字符串；当传入的 taskId 为空时，这里的
	// TaskId 为 userId。
	// TaskId和UserId均为非必填参数，但是你需要至少填一个参数以保证可以正常发起请求。
	TaskID *string `json:"TaskId,omitempty" query:"TaskId"`

	// 客户端发起转推任务的用户 ID。
	// 你在客户端发起多个任务，当使用该接口进行查询时：
	// * 查询第一个任务时，UserId 可以传入发起转推任务的用户 UserId，TaskId 传入空字符串；或直接将该用户的 UserId 传入 TaskId。
	// * 查询第二个以上任务时，UserId 和 TaskId 为发起转推任务的用户 UserId 和 TaskId。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type GetPushSingleStreamToCDNTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPushSingleStreamToCDNTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetPushSingleStreamToCDNTaskResResult          `json:"Result,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetPushSingleStreamToCDNTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetPushSingleStreamToCDNTaskResResponseMetadataError - 仅在请求失败时返回。
type GetPushSingleStreamToCDNTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResult struct {
	PushSingleStreamToCDNTask *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask `json:"PushSingleStreamToCDNTask,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask struct {
	Control *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskControl `json:"Control,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 推流状态
	// * 0：未获取到任务状态。建议稍后重新查询
	// * 1：未开始推流
	// * 2：首次连接 CDN 服务
	// * 3：正在重连 CDN 服务
	// * 4：连接 CDN 服务成功，正在尝试推流。
	// * 5：连接 CDN 服务成功，推流成功
	// * 6：已停止推流。
	// 仅当Status=2 时，PushStreamState 有实际意义；当Status=3 时，PushStreamState=6; status 为其他值时，PushStreamState 均为0。
	PushStreamState *int32 `json:"PushStreamState,omitempty"`

	// 推流地址。目前仅支持 rtmp 协议
	PushURL *string `json:"PushURL,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。- 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：客户端用户主动退房/调用停止转推接口
	StopReason *string                                                               `json:"StopReason,omitempty"`
	Stream     *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskStream `json:"Stream,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskControl struct {

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为 0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskStream struct {

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 用户Id，表示这个流所属的用户。
	UserID *string `json:"UserId,omitempty"`
}

type GetRecordTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的云端录制任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata GetRecordTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetRecordTaskResResult          `json:"Result,omitempty"`
}

type GetRecordTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetRecordTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetRecordTaskResResponseMetadataError - 仅在请求失败时返回。
type GetRecordTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetRecordTaskResResult struct {

	// REQUIRED; 录制任务信息
	RecordTask GetRecordTaskResResultRecordTask `json:"RecordTask"`
}

// GetRecordTaskResResultRecordTask - 录制任务信息
type GetRecordTaskResResultRecordTask struct {

	// 任务结束的时间，为 Unix 时间戳，单位毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 录制生成的文件列表。
	RecordFileList []*GetRecordTaskResResultRecordTaskRecordFileListItem `json:"RecordFileList,omitempty"`

	// 任务开始的时间，为 Unix 时间戳，单位毫秒。
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态:
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int32 `json:"Status,omitempty"`

	// 任务停止的原因：空：表示任务未结束UnknownStopReason：未知停止原因StopByAPI：用户主动通过 API 停止StartTaskFailed：任务启动失败IdleTimeOut：超过了最大空闲时间
	StopReason *string `json:"StopReason,omitempty"`
}

type GetRecordTaskResResultRecordTaskRecordFileListItem struct {

	// REQUIRED; 当前录制文件创建的时间，为 Unix 时间戳，单位毫秒。
	StartTime int64 `json:"StartTime"`

	// 音频录制编码器
	AudioCodec *string `json:"AudioCodec,omitempty"`

	// 文件时长，单位毫秒。
	Duration *int64 `json:"Duration,omitempty"`

	// 文件在对象存储平台中的完整路径，如abc/efg/123.mp4。仅在你选择配置存储到对象存储平台时，此参数有效。
	ObjectKey *string `json:"ObjectKey,omitempty"`

	// 文件大小，单位字节。
	Size *int64 `json:"Size,omitempty"`

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetRecordTaskResResultRecordTaskRecordFileListPropertiesItemsItem `json:"StreamList,omitempty"`

	// 文件在点播平台的唯一标识。你可以根据 vid 可以在点播平台上找到对应的文件。仅在你选择配置存储到 Vod 平台时，此参数有效。
	Vid *string `json:"Vid,omitempty"`

	// 视频录制编码协议
	VideoCodec *string `json:"VideoCodec,omitempty"`

	// 录制视频高度，单位像素。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 录制视频宽度，单位像素。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type GetRecordTaskResResultRecordTaskRecordFileListPropertiesItemsItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetResourcePackNumRes struct {

	// REQUIRED
	ResponseMetadata GetResourcePackNumResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result int32 `json:"Result"`
}

type GetResourcePackNumResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetResourcePackNumResResponseMetadataError `json:"Error,omitempty"`
}

// GetResourcePackNumResResponseMetadataError - 仅在请求失败时返回。
type GetResourcePackNumResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetWebCastTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的云录屏任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetWebCastTaskRes struct {

	// REQUIRED
	ResponseMetadata GetWebCastTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetWebCastTaskResResult          `json:"Result,omitempty"`
}

type GetWebCastTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetWebCastTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetWebCastTaskResResponseMetadataError - 仅在请求失败时返回。
type GetWebCastTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetWebCastTaskResResult struct {

	// 云录屏任务信息
	WebCastTask *GetWebCastTaskResResultWebCastTask `json:"WebCastTask,omitempty"`
}

// GetWebCastTaskResResultWebCastTask - 云录屏任务信息
type GetWebCastTaskResResultWebCastTask struct {

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime           *int64                                               `json:"EndTime,omitempty"`
	EventNotifyConfig *GetWebCastTaskResResultWebCastTaskEventNotifyConfig `json:"EventNotifyConfig,omitempty"`

	// 最大运行时间，超过此时间后，任务自动终止。单位为秒。取值范围为[10,86400]，默认值为 86400。不填时自动调整为默认值。
	MaxRunningTime *int32                                           `json:"MaxRunningTime,omitempty"`
	MonitorConfig  *GetWebCastTaskResResultWebCastTaskMonitorConfig `json:"MonitorConfig,omitempty"`

	// 需要转推的网页地址，可以携带自定义的 queryParams 来鉴权等，总体长度不超过 1024。
	SourceURL *string `json:"SourceURL,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * ExceedMaxRunningTime：任务超过最大运行时间
	StopReason *string `json:"StopReason,omitempty"`

	// 推送网页音视频内容的用户对应的 UserId
	UserID *string `json:"UserId,omitempty"`

	// 输出的视频参数，最多支持2路，以大小流的方式支持接收端按需订阅，将以最大的视频流分辨率作为网页渲染分辨率，为空时按默认值填充一路
	VideoSolutions []*GetWebCastTaskResResultWebCastTaskVideoSolutionsItem `json:"VideoSolutions,omitempty"`
}

type GetWebCastTaskResResultWebCastTaskEventNotifyConfig struct {

	// 是否启用页面主动事件通知, 默认值为false。
	// * false：页面在打开后就会开始采集，在收到 StopWebCast openAPI 请求后结束采集。
	// * true：在页面中注入两个 JS 函数：onWebcastStart()和 onWebcastEnd()。
	// 当页面判断资源加载完成之后调用onWebcastStart()，控制程序才会开始进行页面内容的采集。当页面判断本次任务内容已完成时调用onWebcastEnd() 通知控制程序结束本次任务。StopWebCast openAPI 效果不变，业务可提前结束任务。其他页面内容、JS
	// 线程的检测（若启用），将在收到 onWebcastStart()事件后才开始。
	// 当启用页面主动事件通知后，你可以参考以下示例代码来通知采集开始。
	// <script>
	// if (ready() && typeof onWebcastStart === 'function')
	// onWebcastStart();
	// </script>
	EnableEventNotify *bool `json:"EnableEventNotify,omitempty"`

	// 启用页面主动事件通知后，等待开始事件的超时时间。取值范围为[0,60]，单位为秒。默认值为0，表示不启用。仅当 EnableEventNotify 为 true 时，此参数有效。
	// * 当在超时时间内收到开始事件，采集功能正常运行，用户将收到 Status=1的回调。
	// * 当超时时间内未收到开始事件，将进行刷新，等待时间被重置，再次发生超时后将进行任务重调度。刷新时将回调 Status=4，Reason=" StartEventTimeout"。重调度时将回调 Status=5，Reason="StartEventTimeout"。
	StartTimeout *int32 `json:"StartTimeout,omitempty"`
}

type GetWebCastTaskResResultWebCastTaskMonitorConfig struct {

	// 对页面是否白屏的检测间隔。取值范围为[2,30]，单位为秒。默认值为0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageBlank" 。
	// * 再次出现连续两次检测命中时将进行任务重调度，并回调Status=5，Reason="PageBlank"。
	// 注意：页面全白可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	BlankCheckInterval *int32 `json:"BlankCheckInterval,omitempty"`

	// 对页面 JS 线程是否崩溃/卡死的检测间隔。 取值范围为[0,30]，单位为秒。默认值为0，表示不启用。
	// 当出现检测命中时将进行任务重调度，并回调 Status=5，Reason="PageCrash"。
	CrashCheckInterval *int32 `json:"CrashCheckInterval,omitempty"`

	// 对页面内容是否无变化的检测间隔。取值范围为[2,30]，单位为秒。默认值为0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageFreeze"。
	// * 再次出现连续两次检测命中时，将进行任务重调度，并回调Status=5，Reason="PageFreeze"。
	// 注意：页面无变化可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	FreezeCheckInterval *int32 `json:"FreezeCheckInterval,omitempty"`
}

type GetWebCastTaskResResultWebCastTaskVideoSolutionsItem struct {

	// 最大发送码率，取值范围为[0,10000]，单位为 Kbps，默认值 0，为 0 时表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，单位为 帧/秒，范围为[1,60]，默认值为 15。帧率和码率设置建议参照视频发布参数对照表 [70122#param]以获取最佳体验。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 视频高度，单位为像素，范围为 [50,1080]，默认值为 720。必须是偶数，值为奇数时自动调整为偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频宽度，单位为像素，范围为 [50,1920]，默认值为 1280。必须是偶数，值为奇数时自动调整为偶数。
	Width *int32 `json:"Width,omitempty"`
}

type IndicatorTag struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                      `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTagSubTagsItem `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTag202201 struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                            `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag202201SubTagsItem `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTag202201SubTagsItem struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                 `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag202201 `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTagSubTagsItem struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool           `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type KickUserBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 你希望移出用户的 ID
	UserID string `json:"UserId"`
}

type KickUserRes struct {

	// REQUIRED
	ResponseMetadata KickUserResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *KickUserResResult `json:"Result,omitempty"`
}

type KickUserResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error KickUserResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type KickUserResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

// KickUserResResult - 视请求的接口而定
type KickUserResResult struct {

	// REQUIRED
	Message string `json:"message"`
}

type ListAppsBody struct {

	// 已创建的应用的 ID
	AppID *string `json:"AppId,omitempty"`

	// 返回条目数的上限，默认为 50，最大为 100
	Limit *string `json:"Limit,omitempty"`

	// * 偏移量，单位为条；
	// * 因单次查询返回结果数量有限，如果需要大量查询，需要借助偏移量实现分页查询；
	Offset *string `json:"Offset,omitempty"`

	// * 默认按照 App 创建时间升序返回；
	// * 设置为 1 时，降序返回；
	Reverse *int32 `json:"Reverse,omitempty"`
}

type ListAppsRes struct {

	// REQUIRED
	ResponseMetadata ListAppsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 查询的指定 AppId 的信息在 Result 中显示。
	Result ListAppsResResult `json:"Result"`
}

type ListAppsResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListAppsResResponseMetadataError `json:"Error,omitempty"`
}

// ListAppsResResponseMetadataError - 仅在请求失败时返回。
type ListAppsResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// ListAppsResResult - 查询的指定 AppId 的信息在 Result 中显示。
type ListAppsResResult struct {

	// REQUIRED
	AppList []ListAppsResResultAppListItem `json:"AppList"`

	// REQUIRED; 返回条目数的限制
	Limit int32 `json:"Limit"`

	// REQUIRED; 偏移量，单位为条
	Offset int32 `json:"Offset"`

	// REQUIRED; 当前条件查询到的 App 数量
	Total int32 `json:"Total"`
}

type ListAppsResResultAppListItem struct {

	// 应用 ID
	AppID *string `json:"AppId,omitempty"`

	// 主 AppKey，用于生成 Token
	AppKey *string `json:"AppKey,omitempty"`

	// 应用名称
	AppName *string `json:"AppName,omitempty"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitempty"`

	// 副 AppKey，启用后可用于生成 Token
	SecondaryAppKey *string `json:"SecondaryAppKey,omitempty"`

	// 服务状态。枚举值为：
	// * 0: 创建中——指 AppID 在初始化
	// * 1: 运行中——指当前 AppID 为正常服务状态
	// * 6: 欠费关停
	// * 98: 已停用——指调用服务端停用接口，当前 AppID 被设置为不可用状态
	Status *int32 `json:"Status,omitempty"`
}

type ListCallDetailQuery struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间结束时间，格式为RFC3339，单位秒
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 房间开始时间，格式为RFC3339，单位秒
	StartTime string `json:"StartTime" query:"StartTime"`

	// REQUIRED; 查询的用户ID列表，多个用户ID用逗号隔开，最多可以指定 20 个。值不合法时默认剔除。
	UserID string `json:"UserId" query:"UserId"`

	// 获取数据的传输方向。值可取0，1，2。
	// * 0:上下行
	// * 1:上行
	// * 2: 下行。 默认值为0。不传时自动调整为默认值。
	Direction *string `json:"Direction,omitempty" query:"Direction"`
}

type ListCallDetailRes struct {

	// REQUIRED
	ResponseMetadata ListCallDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListCallDetailResResult          `json:"Result,omitempty"`
}

type ListCallDetailResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListCallDetailResResponseMetadataError `json:"Error,omitempty"`
}

// ListCallDetailResResponseMetadataError - 仅在请求失败时返回。
type ListCallDetailResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListCallDetailResResult struct {

	// REQUIRED; 用户通话质量指标信息，仅在请求成功时返回。
	UserCallDetail []ListCallDetailResResultUserCallDetailItem `json:"UserCallDetail"`
}

type ListCallDetailResResultUserCallDetailItem struct {

	// REQUIRED; 通话质量指标数据详情。本接口支持的通话质量指标参看指标列表 [1167931#list]。
	CallDetail []ListCallDetailResResultUserCallDetailPropertiesItemsItem `json:"CallDetail"`

	// REQUIRED; 获取数据的传输方向。
	Direction string `json:"Direction"`

	// REQUIRED; 对端用户 Id，只有UserId用户处于下行时存在
	PeerID string `json:"PeerId"`

	// REQUIRED; 查询房间 Id
	RoomID string `json:"RoomId"`

	// REQUIRED; 查询用户 Id
	UserID string `json:"UserId"`
}

// ListCallDetailResResultUserCallDetailPropertiesItemsDataItem - 指标数据详情
type ListCallDetailResResultUserCallDetailPropertiesItemsDataItem struct {

	// REQUIRED; 该指标数据点的 Unix 时间戳，单位为毫秒。
	Time int64 `json:"Time"`

	// REQUIRED; 该指标数据点的数据值
	Value float32 `json:"Value"`
}

// ListCallDetailResResultUserCallDetailPropertiesItemsItem - 通话质量指标数据详情。本接口支持的通话质量指标参看指标列表 [1167931#list]。
type ListCallDetailResResultUserCallDetailPropertiesItemsItem struct {

	// REQUIRED; 指标数据详情
	Data []ListCallDetailResResultUserCallDetailPropertiesItemsDataItem `json:"Data"`

	// REQUIRED; 该指标中文名称
	Desc string `json:"Desc"`

	// REQUIRED; 该项指标字段名称
	Name string `json:"Name"`

	// REQUIRED; 指标单位
	Unit string `json:"Unit"`
}

type ListDetectionTaskQuery struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 用户 ID
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type ListDetectionTaskRes struct {

	// REQUIRED
	ResponseMetadata ListDetectionTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListDetectionTaskResResult `json:"Result,omitempty"`
}

type ListDetectionTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListDetectionTaskResResponseMetadataError `json:"Error,omitempty"`
}

// ListDetectionTaskResResponseMetadataError - 仅在请求失败时返回。
type ListDetectionTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// ListDetectionTaskResResult - 视请求的接口而定
type ListDetectionTaskResResult struct {

	// REQUIRED; 任务状态信息。如果从未对指定的用户发起审核，查询其审核状态时，返回的 EventData 值为空列表。
	EventData []ListDetectionTaskResResultEventDataItem `json:"EventData"`

	// REQUIRED; 仅在请求成功时返回success。
	Message string `json:"Message"`
}

type ListDetectionTaskResResultEventDataItem struct {

	// REQUIRED; 用户审核发起时间戳，unix 时间，单位为秒。
	CreateTime int32 `json:"CreateTime"`

	// REQUIRED; 进行审核的内容类型： 1：视频截图； 2：音频切片； 3：视频截图+音频切片（默认值）
	MediaType int32 `json:"MediaType"`

	// REQUIRED; 审核房间 ID
	RoomID string `json:"RoomId"`

	// REQUIRED; 审核状态： 1：运行中： 2：已完成；
	Status int32 `json:"Status"`

	// REQUIRED; 若开始审核接口填入 UserId，此处返回填入 UserId。若开始审核接口未填入 UserId，此处返回房间内触发过审核任务的用户的 UserId。
	UserID string `json:"UserId"`
}

type ListHotMusicBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// 自定义热歌榜ID 列表
	CustomHotListID []*string `json:"CustomHotListId,omitempty"`

	// 热歌榜类型。
	// * [1]：火山内容中心热歌榜。
	// * [2]：项目热歌榜。
	// * [1,2]：全部热歌榜。 默认值为[1,2]。
	HotTypes []*int32 `json:"HotTypes,omitempty"`
}

type ListHotMusicRes struct {

	// REQUIRED
	ResponseMetadata ListHotMusicResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListHotMusicResResult          `json:"Result,omitempty"`
}

type ListHotMusicResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListHotMusicResResponseMetadataError `json:"Error,omitempty"`
}

// ListHotMusicResResponseMetadataError - 仅在请求失败时返回。
type ListHotMusicResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListHotMusicResResult struct {

	// 热歌榜单详情
	HotList []*ListHotMusicResResultHotListItem `json:"HotList,omitempty"`
}

type ListHotMusicResResultHotListItem struct {

	// REQUIRED; 自定义热歌榜 ID
	CustomHotListID string `json:"CustomHotListId"`

	// REQUIRED; 自定义热歌榜名
	CustomHotListName string `json:"CustomHotListName"`

	// REQUIRED; 热歌榜单名称
	HotName string `json:"HotName"`

	// REQUIRED; 歌曲详情
	HotSong []ListHotMusicResResultHotListPropertiesItemsItem `json:"HotSong"`

	// REQUIRED; 热歌榜单 ID
	HotType int32 `json:"HotType"`
}

type ListHotMusicResResultHotListPropertiesItemsItem struct {

	// 歌曲是否支持伴唱原唱切换。
	// * 0： 不支持。
	// * 1： 声道切换
	// * 2： 音轨切换
	BgmType *int32 `json:"BgmType,omitempty"`

	// 歌曲时长，单位为秒
	Duration *int32 `json:"Duration,omitempty"`

	// 歌词类型
	// * [0]： krc
	// * [1]： lrc
	// * [0,1]： 两种歌词均有。 若为空，则表示为无歌词。
	LyricType []*int32 `json:"LyricType,omitempty"`

	// 歌曲是否支持打分。
	// * 0： 不支持
	// * 1： 支持
	PitchType *int32 `json:"PitchType,omitempty"`

	// 歌曲封面地址
	PosterURL *string `json:"PosterUrl,omitempty"`

	// 歌曲排名
	Rank *int32 `json:"Rank,omitempty"`

	// 歌曲是否支持录制。
	// * 0： 不支持
	// * 1： 支持
	RecordType *int32 `json:"RecordType,omitempty"`

	// 歌曲授权范围。
	// * 0：中国大陆
	// * 1：全球
	RegionType *int32 `json:"RegionType,omitempty"`

	// 场景类型：
	// * 0：所有。
	// * 1：BGM
	// * 2：k 歌
	SceneType *int32 `json:"SceneType,omitempty"`

	// 歌曲高潮部分
	Segment *string `json:"Segment,omitempty"`

	// 歌曲歌手名
	Singer *string `json:"Singer,omitempty"`

	// 歌曲 ID
	SongID *string `json:"SongId,omitempty"`

	// 歌曲名称
	Songname *string `json:"Songname,omitempty"`

	// 供应商 ID
	VendorID *int32 `json:"VendorId,omitempty"`

	// 供应商名称
	VendorName *string `json:"VendorName,omitempty"`
}

type ListMusicsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// 过滤选项：
	// * 1：过滤没有歌词的歌曲
	// * 2：过滤不支持打分的歌曲
	// * 3：过滤不支持伴唱切换的歌曲
	// * 4：过滤没有高潮片段的歌曲
	Filters []*int32 `json:"Filters,omitempty"`

	// 歌曲更新或新增时间戳，Unix 时间，单位为秒。表示筛选更新或新增时间大于等于 LastUpdateTime 的有效歌曲
	LastUpdateTime *int32 `json:"LastUpdateTime,omitempty"`

	// 分页序号，默认值为1
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页歌曲数量，取值范围为 (0,100]，默认值为 10
	PageSize *int32 `json:"PageSize,omitempty"`

	// 查询歌曲 ID，最多可以指定 200 个
	SongIDs []*int32 `json:"SongIds,omitempty"`
}

type ListMusicsRes struct {

	// REQUIRED
	ResponseMetadata ListMusicsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListMusicsResResult `json:"Result"`
}

type ListMusicsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListMusicsResResponseMetadataError `json:"Error,omitempty"`
}

// ListMusicsResResponseMetadataError - 仅在请求失败时返回。
type ListMusicsResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListMusicsResResult struct {

	// REQUIRED; 返回歌曲详情
	List []ListMusicsResResultListItem `json:"List"`

	// REQUIRED; 返回歌曲总数
	Total int32 `json:"Total"`
}

type ListMusicsResResultListItem struct {

	// 歌曲是否支持伴唱原唱切换。
	// * 0： 不支持。
	// * 1： 声道切换
	// * 2： 音轨切换
	BgmType *int32 `json:"BgmType,omitempty"`

	// 歌曲时长，单位为秒
	Duration *int32 `json:"Duration,omitempty"`

	// 歌词类型
	// * 0： krc
	// * 1： lrc
	// * [0,1]： 两种歌词均有。 若为空，则表示为无歌词。
	LyricType []*int32                             `json:"LyricType,omitempty"`
	LyricURL  *ListMusicsResResultListItemLyricURL `json:"LyricUrl,omitempty"`

	// 歌曲是否支持打分。
	// * 0： 不支持
	// * 1： 支持
	PitchType *int32 `json:"PitchType,omitempty"`

	// 歌曲封面地址
	PosterURL *string `json:"PosterUrl,omitempty"`

	// 歌曲是否支持录制。
	// * 0： 不支持
	// * 1： 支持
	RecordType *int32 `json:"RecordType,omitempty"`

	// 歌曲高潮部分
	Segment *string `json:"Segment,omitempty"`

	// 歌曲歌手名
	Singer *string `json:"Singer,omitempty"`

	// 歌曲 ID
	SongID *string `json:"SongId,omitempty"`

	// 歌曲名称
	Songname *string `json:"Songname,omitempty"`

	// 歌曲更新时间戳，unix 时间，单位为秒
	UpdateAt *int32 `json:"UpdateAt,omitempty"`

	// 供应商 ID
	VendorID *int32 `json:"VendorId,omitempty"`

	// 供应商名称
	VendorName *string `json:"VendorName,omitempty"`
}

type ListMusicsResResultListItemLyricURL struct {

	// krc 歌词地址
	KrcURL *string `json:"KrcUrl,omitempty"`

	// lrc 歌词地址
	LrcURL *string `json:"LrcUrl,omitempty"`

	// midi 文件地址
	MidiURL *string `json:"MidiUrl,omitempty"`
}

type ListOperationDataBody struct {

	// REQUIRED; 返回聚合时间的粒度，支持设为以下值：
	// * 1d：粒度为 1 天
	// * 1h：粒度为 1 小时
	AggregateGranularity string `json:"AggregateGranularity"`

	// REQUIRED; 你的音视频应用的唯一标志。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。Indicator 可选值，参看离线运营数据相关 indicator [1167931#indicator]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`
}

type ListOperationDataRes struct {

	// REQUIRED
	ResponseMetadata ListOperationDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListOperationDataResResult          `json:"Result,omitempty"`
}

type ListOperationDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListOperationDataResResponseMetadataError `json:"Error,omitempty"`
}

// ListOperationDataResResponseMetadataError - 仅在请求失败时返回。
type ListOperationDataResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListOperationDataResResult struct {

	// 返回聚合时间的粒度
	AggregateGranularity *string `json:"AggregateGranularity,omitempty"`

	// 具体指标数据
	Indicators []*ListOperationDataResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListOperationDataResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListOperationDataResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                           `json:"Name,omitempty"`
	Overview *ListOperationDataResResultIndicatorsItemOverview `json:"overview,omitempty"`
	Unit     *string                                           `json:"Unit,omitempty"`
}

type ListOperationDataResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListOperationDataResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListOperationDistributionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志。
	AppID string `json:"AppId"`

	// REQUIRED; 查询维度。目前仅支持查询:
	// * 一级行政区（包括港澳台地区）
	// * 国家
	Dimension string `json:"Dimension"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。
	// 目前仅支持查询 user_num，即此 AppId 在所选日期的通话总人数。通话人数按用户 id 去重。
	Indicator string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`
}

type ListOperationDistributionRes struct {

	// REQUIRED
	ResponseMetadata ListOperationDistributionResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListOperationDistributionResResult          `json:"Result,omitempty"`
}

type ListOperationDistributionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListOperationDistributionResResponseMetadataError `json:"Error,omitempty"`
}

// ListOperationDistributionResResponseMetadataError - 仅在请求失败时返回。
type ListOperationDistributionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListOperationDistributionResResult struct {

	// 具体的指标值以及对应时间
	Data []*ListOperationDistributionResResultDataItem `json:"Data,omitempty"`

	// 查询维度
	Dimension *string `json:"Dimension,omitempty"`

	// 指标名称
	Name *string `json:"Name,omitempty"`
}

type ListOperationDistributionResResultDataItem struct {

	// 维度名
	DistributionName *string `json:"DistributionName,omitempty"`

	// 占比
	Proportion *float64 `json:"Proportion,omitempty"`

	// 指标值
	Value *float64 `json:"Value,omitempty"`
}

type ListQualityBody struct {

	// REQUIRED; 返回聚合时间的粒度，支持设为以下值：
	// * 1d：粒度为 1 天
	// * 5min：粒度为 5 分钟
	AggregateGranularity string `json:"AggregateGranularity"`

	// REQUIRED; 你的音视频应用的唯一标志。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。Indicator 可选值，参看质量数据相关 indicator [1167931#listquality]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`

	// 网络类型，枚举值为：2g、3g、4g、5g、wifi
	// ProductType = web 时该参数不生效
	Access []*string `json:"Access,omitempty"`

	// 枚举值为：android、iOS、linux、mac、windows
	// ProductType = web 时该参数不生效
	OS []*string `json:"OS,omitempty"`

	// 要查询的产品类型，枚举值为 native 和 web。
	// * native：指 android、iOS、linux、mac或 windows 平台的 veRTC SDK。
	// * web：指 Web 平台的veRTC SDK。 默认值为native
	ProductType *string `json:"ProductType,omitempty"`

	// 房间 ID。如果不填，代表查询该 AppId 的整体离线指标。
	RoomID *string `json:"RoomId,omitempty"`

	// 查询的用户 ID 列表，最多可以指定 20 个。值不合法时默认剔除。此字段仅在 RoomId 不为空时生效。
	UserID []*string `json:"UserId,omitempty"`
}

type ListQualityDistributionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志。
	AppID string `json:"AppId"`

	// REQUIRED; 查询维度，一次仅支持查询一个维度。
	// 支持设为以下值：
	// * Province：一级行政区（包括港澳台地区）
	// * Country：国家
	// * OS：用户设备平台 包括：android、ios、linux、mac、windows
	// * Access：用户网络类型 包括：2g、3g、4g、5g、wifi
	Dimension string `json:"Dimension"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称，一次仅支持查询一个指标。Indicator 可选值，参看质量数据相关 indicator [1167931#listquality]。
	Indicator string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`

	// 要查询的数据所属设备端，支持设为以下值：
	// * native：指 Android、iOS、Linux、Mac、Windows 端。
	// * web：指 Web 端。 默认值为 native。
	Platform    *string `json:"Platform,omitempty"`
	ProductType *string `json:"ProductType,omitempty"`
}

type ListQualityDistributionRes struct {

	// REQUIRED
	ResponseMetadata ListQualityDistributionResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListQualityDistributionResResult          `json:"Result,omitempty"`
}

type ListQualityDistributionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListQualityDistributionResResponseMetadataError `json:"Error,omitempty"`
}

// ListQualityDistributionResResponseMetadataError - 仅在请求失败时返回。
type ListQualityDistributionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListQualityDistributionResResult struct {

	// 具体指标数据。
	Indicators []*ListQualityDistributionResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListQualityDistributionResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListQualityDistributionResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name *string `json:"Name,omitempty"`
}

type ListQualityDistributionResResultIndicatorsPropertiesItemsItem struct {

	// 具体的维度信息。
	Dimension *string `json:"Dimension,omitempty"`

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标数值，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListQualityRes struct {

	// REQUIRED
	ResponseMetadata ListQualityResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListQualityResResult          `json:"Result,omitempty"`
}

type ListQualityResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListQualityResResponseMetadataError `json:"Error,omitempty"`
}

// ListQualityResResponseMetadataError - 仅在请求失败时返回。
type ListQualityResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListQualityResResult struct {

	// 返回数据的粒度
	AggregateGranularity *string `json:"AggregateGranularity,omitempty"`

	// 具体指标数据
	Indicators []*ListQualityResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListQualityResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListQualityResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                     `json:"Name,omitempty"`
	Overview *ListQualityResResultIndicatorsItemOverview `json:"overview,omitempty"`
	Unit     *string                                     `json:"Unit,omitempty"`
}

type ListQualityResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListQualityResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRealTimeOperationDataBody struct {

	// REQUIRED; 你的音视频应用的唯一标志。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。 Indicator 可选值，参看实时运营数据相关 indicator [1167931#realtime]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`
}

type ListRealTimeOperationDataRes struct {

	// REQUIRED
	ResponseMetadata ListRealTimeOperationDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRealTimeOperationDataResResult          `json:"Result,omitempty"`
}

type ListRealTimeOperationDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRealTimeOperationDataResResponseMetadataError `json:"Error,omitempty"`
}

// ListRealTimeOperationDataResResponseMetadataError - 仅在请求失败时返回。
type ListRealTimeOperationDataResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRealTimeOperationDataResResult struct {

	// 具体指标数据
	Indicators []*ListRealTimeOperationDataResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListRealTimeOperationDataResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListRealTimeOperationDataResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                                   `json:"Name,omitempty"`
	Overview *ListRealTimeOperationDataResResultIndicatorsItemOverview `json:"overview,omitempty"`
	Unit     *string                                                   `json:"Unit,omitempty"`
}

type ListRealTimeOperationDataResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListRealTimeOperationDataResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRealTimeQualityBody struct {

	// REQUIRED; 你的音视频应用的唯一标志。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。Indicator 可选值，参看质量数据相关 indicator [1167931#listquality]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`

	// 要查询的数据所属设备端，支持设为以下值：
	// * native：指 Android、iOS、Linux、Mac、Windows 端。
	// * web：指 Web 端。
	// 默认值为native。
	ProductType *string `json:"ProductType,omitempty"`

	// 房间 ID，是房间的唯一标志。如果不填，代表查询该 AppId 的整体实时指标。
	RoomID *string `json:"RoomId,omitempty"`

	// 查询的用户 ID 列表，最多可以指定 20 个。值不合法时默认剔除。此字段仅在 RoomID 不为空时生效。
	UserID []*string `json:"UserId,omitempty"`
}

type ListRealTimeQualityDistributionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志。
	AppID string `json:"AppId"`

	// REQUIRED; 查询的维度，一次仅支持查询一个维度。
	// 支持设为以下值：
	// * Province： 一级行政区（包括港澳台地区）
	// * Country：国家
	// * OS：用户设备平台 包括：android、ios、linux、mac、windows
	// * Access：用户网络类型 包括：2g、3g、4g、5g、wifi
	Dimension string `json:"Dimension"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。一次仅支持查询一个指标。Indicator 可选值，参看质量数据相关 indicator [1167931#listquality]。
	Indicator string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime   string  `json:"StartTime"`
	Platform    *string `json:"Platform,omitempty"`
	ProductType *string `json:"ProductType,omitempty"`
}

type ListRealTimeQualityDistributionRes struct {

	// REQUIRED
	ResponseMetadata ListRealTimeQualityDistributionResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRealTimeQualityDistributionResResult          `json:"Result,omitempty"`
}

type ListRealTimeQualityDistributionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRealTimeQualityDistributionResResponseMetadataError `json:"Error,omitempty"`
}

// ListRealTimeQualityDistributionResResponseMetadataError - 仅在请求失败时返回。
type ListRealTimeQualityDistributionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRealTimeQualityDistributionResResult struct {

	// 具体指标数据
	Indicators []*ListRealTimeQualityDistributionResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListRealTimeQualityDistributionResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListRealTimeQualityDistributionResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name *string `json:"Name,omitempty"`
}

type ListRealTimeQualityDistributionResResultIndicatorsPropertiesItemsItem struct {

	// 具体的维度信息。
	Dimension *string `json:"Dimension,omitempty"`

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标数值，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRealTimeQualityRes struct {

	// REQUIRED
	ResponseMetadata ListRealTimeQualityResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRealTimeQualityResResult          `json:"Result,omitempty"`
}

type ListRealTimeQualityResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRealTimeQualityResResponseMetadataError `json:"Error,omitempty"`
}

// ListRealTimeQualityResResponseMetadataError - 仅在请求失败时返回。
type ListRealTimeQualityResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRealTimeQualityResResult struct {

	// 具体指标数据
	Indicators []*ListRealTimeQualityResResultIndicatorsItem `json:"Indicators,omitempty"`

	// 房间 ID
	RoomID *string `json:"RoomId,omitempty"`

	// 用户 ID
	UserID []*string `json:"UserId,omitempty"`
}

type ListRealTimeQualityResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListRealTimeQualityResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                             `json:"Name,omitempty"`
	Overview *ListRealTimeQualityResResultIndicatorsItemOverview `json:"overview,omitempty"`
	Unit     *string                                             `json:"Unit,omitempty"`
}

type ListRealTimeQualityResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListRealTimeQualityResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRelayStreamQuery struct {

	// REQUIRED; 应用的唯一标志。
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 要查询房间的 ID
	RoomID string `json:"RoomId" query:"RoomId"`

	// 页大小，取值范围为[1,100]，默认值为 10
	Limit *int32 `json:"Limit,omitempty" query:"Limit"`

	// 起始位置，取值范围为[0,9999]，默认值为 0。
	Offset *int32 `json:"Offset,omitempty" query:"Offset"`

	// 接口调用时间顺序。 0：倒序。1：正序。默认值为0
	Order *int32 `json:"Order,omitempty" query:"Order"`
}

type ListRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata ListRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRelayStreamResResult          `json:"Result,omitempty"`
}

type ListRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// ListRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type ListRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRelayStreamResResult struct {

	// REQUIRED; 页大小
	Limit int32 `json:"Limit"`

	// REQUIRED; 下一页索引
	Offset int32 `json:"Offset"`

	// REQUIRED; 任务列表
	Task []ListRelayStreamResResultTaskItem `json:"Task"`

	// REQUIRED; 当前页大小
	Total int32 `json:"Total"`
}

// ListRelayStreamResResultTaskItem - 任务列表
type ListRelayStreamResResultTaskItem struct {

	// 应用的唯一标志
	AppID *string `json:"AppId,omitempty"`

	// 任务开始时间，Unix 时间戳，单位为秒
	CreateTime *int32 `json:"CreateTime,omitempty"`

	// 发送帧率，值的范围为[1，60]，默认值为15，转码时生效。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 媒体类型。0：音视频 1：音频2：视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 房间 ID，是房间的唯一标志
	RoomID *string `json:"RoomId,omitempty"`

	// 任务当前状态。 1：任务启动中。2：任务运行中。3：任务已结束
	Status *int32 `json:"Status,omitempty"`

	// 流处理模式。0：转码1：转封装
	StreamMode *int32 `json:"StreamMode,omitempty"`

	// 在线媒体流地址
	StreamURL *string `json:"StreamUrl,omitempty"`

	// 任务 ID
	TaskID *string `json:"TaskId,omitempty"`

	// 在线媒体流对应的 UserId
	UserID *string `json:"UserId,omitempty"`

	// 视频高度，单位为 px，范围为 [16, 1920]。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 视频宽度。单位为 px，范围为 [16, 1920]。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type ListResourcePackagesRes struct {

	// REQUIRED
	ResponseMetadata ListResourcePackagesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result []ListResourcePackagesResResultItem `json:"Result"`
}

type ListResourcePackagesResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListResourcePackagesResResponseMetadataError `json:"Error,omitempty"`
}

// ListResourcePackagesResResponseMetadataError - 仅在请求失败时返回。
type ListResourcePackagesResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListResourcePackagesResResultItem struct {

	// REQUIRED
	ConfigurationCode string `json:"ConfigurationCode"`

	// REQUIRED
	PackageType int32 `json:"PackageType"`

	// REQUIRED
	Price int32 `json:"Price"`

	// REQUIRED
	ZhName             string   `json:"ZhName"`
	EffectiveTimeType  *int32   `json:"EffectiveTimeType,omitempty"`
	ValidityOptionList []*int32 `json:"ValidityOptionList,omitempty"`
	ValidityPeriod     *int32   `json:"ValidityPeriod,omitempty"`
}

type ListResourcePackagesV2Res struct {

	// REQUIRED
	ResponseMetadata ListResourcePackagesV2ResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result []ListResourcePackagesV2ResResultItem `json:"Result"`
}

type ListResourcePackagesV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListResourcePackagesV2ResResponseMetadataError `json:"Error,omitempty"`
}

// ListResourcePackagesV2ResResponseMetadataError - 仅在请求失败时返回。
type ListResourcePackagesV2ResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListResourcePackagesV2ResResultItem struct {
	PackageConfigs []*ListResourcePackagesV2ResResultPropertiesItemsItem `json:"PackageConfigs,omitempty"`

	// 资源包类型，1: 音视频通话资源包 2: 云端录制资源包
	ResourceType *int32 `json:"ResourceType,omitempty"`
}

type ListResourcePackagesV2ResResultPropertiesItemsItem struct {

	// REQUIRED
	ConfigurationCode string `json:"ConfigurationCode"`

	// REQUIRED
	EffectiveTimeType int32 `json:"EffectiveTimeType"`

	// REQUIRED
	PackageType int32 `json:"PackageType"`

	// REQUIRED
	Price float32 `json:"Price"`

	// REQUIRED
	ValidityOptionList string `json:"ValidityOptionList"`

	// REQUIRED
	ValidityPeriod int32 `json:"ValidityPeriod"`

	// REQUIRED
	ZhName string `json:"ZhName"`
}

type ListRoomInfoQuery struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位秒
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位秒
	StartTime string `json:"StartTime" query:"StartTime"`

	// 分页序号，默认值为1
	PageNum *int32 `json:"PageNum,omitempty" query:"PageNum"`

	// 每页房间数，最大不能超过 100。默认为 20。如果指定值超过 100，每页的房间数为 100。
	PageSize *int32 `json:"PageSize,omitempty" query:"PageSize"`

	// 房间 ID，是房间的唯一标志
	RoomID *string `json:"RoomId,omitempty" query:"RoomId"`
}

type ListRoomInfoRes struct {

	// REQUIRED
	ResponseMetadata ListRoomInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRoomInfoResResult          `json:"Result,omitempty"`
}

type ListRoomInfoResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRoomInfoResResponseMetadataError `json:"Error,omitempty"`
}

// ListRoomInfoResResponseMetadataError - 仅在请求失败时返回。
type ListRoomInfoResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRoomInfoResResult struct {

	// 是否还有房间未列出。若为true，表示符合条件房间过多。若要查询房间未列出，请修改查询条件。
	HasMore *bool `json:"HasMore,omitempty"`

	// 分页序号
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页房间数
	PageSize *int32 `json:"PageSize,omitempty"`

	// 房间信息
	RoomList []*ListRoomInfoResResultRoomListItem `json:"RoomList,omitempty"`

	// 查询到的房间总数。若同一 RoomId 下有多次通话，记为多个房间。
	Total *int32 `json:"Total,omitempty"`
}

type ListRoomInfoResResultRoomListItem struct {

	// 当前仍在房间用户数
	ActiveUserNum *int32 `json:"ActiveUserNum,omitempty"`

	// 通话 Id，是通话的唯一标识
	CallID *string `json:"CallId,omitempty"`

	// 通话创建时间，格式为 RFC3339，单位秒
	CreatedTime *string `json:"CreatedTime,omitempty"`

	// 通话结束时间，格式为 RFC3339，单位秒，若查询时还未结束，则返回空值。
	DestroyTime *string `json:"DestroyTime,omitempty"`

	// 通话是否结束
	IsFinished *bool `json:"IsFinished,omitempty"`

	// 返回房间 Id
	RoomID *string `json:"RoomId,omitempty"`
}

type ListUsagesBody struct {

	// 你的音视频应用的唯一标志
	AppID *string `json:"AppId,omitempty"`

	// 查询数据的结束时间，格式要求符合 RFC3339 规范。起止时间的跨度最多为30天
	EndTime *string `json:"EndTime,omitempty"`

	// 查询数据的开始时间，格式要求符合 RFC3339 规范
	StartTime *string `json:"StartTime,omitempty"`
}

type ListUsagesRes struct {

	// REQUIRED
	ResponseMetadata ListUsagesResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListUsagesResResult          `json:"Result,omitempty"`
}

type ListUsagesResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListUsagesResResponseMetadataError `json:"Error,omitempty"`
}

// ListUsagesResResponseMetadataError - 仅在请求失败时返回。
type ListUsagesResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListUsagesResResult struct {

	// 具体的通话时长数据
	Usages []*ListUsagesResResultUsagesItem `json:"Usages,omitempty"`
}

type ListUsagesResResultUsagesItem struct {

	// 计价档位为纯音频的通话时长，单位：分钟
	Audio *int64 `json:"Audio,omitempty"`

	// 指标聚合的时间，格式为 RFC3339 规范。这个参数的值，指通话时长对应时间段的开始时刻。以示例值为例，意味着对应的时间段为 北京时间2021年7月24日0点起的1天。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 计价档位为 1080P 视频的通话时长，单位：分钟
	VideoHD *int64 `json:"VideoHD,omitempty"`

	// 计价档位为 720P 视频的通话时长，单位：分钟
	VideoSD *int64 `json:"VideoSD,omitempty"`

	// 计价档位为 360P 视频的通话时长，单位：分钟
	VideoSDM *int64 `json:"VideoSDM,omitempty"`
}

type ListUserInfoQuery struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒
	StartTime string `json:"StartTime" query:"StartTime"`

	// 是否排除 Linux 用户。默认为 true，表示排除 Linux 用户
	ExcludeServerUser *bool `json:"ExcludeServerUser,omitempty" query:"ExcludeServerUser"`

	// 分页序号，默认值为1
	PageNum *int32 `json:"PageNum,omitempty" query:"PageNum"`

	// 每页用户数，最大不能超过 100。默认为 20。如果指定值超过 100，每页的用户数为 100。
	PageSize *int32 `json:"PageSize,omitempty" query:"PageSize"`

	// 查询的用户 ID 列表，多个用户 ID 用逗号隔开，最多可以指定 10 个。值不合法时默认剔除。为空时，查询房间内全部用户信息。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type ListUserInfoRes struct {

	// REQUIRED
	ResponseMetadata ListUserInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListUserInfoResResult          `json:"Result,omitempty"`
}

type ListUserInfoResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListUserInfoResResponseMetadataError `json:"Error,omitempty"`
}

// ListUserInfoResResponseMetadataError - 仅在请求失败时返回。
type ListUserInfoResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListUserInfoResResult struct {

	// 分页序号
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页用户数，若同一用户有多次通话记为多个用户
	PageSize *int32 `json:"PageSize,omitempty"`

	// 查询到的用户总数，若同一用户有多次通话记为多个用户
	Total *int32 `json:"Total,omitempty"`

	// 用户信息
	UserList []*ListUserInfoResResultUserListItem `json:"UserList,omitempty"`
}

type ListUserInfoResResultUserListItem struct {

	// CallId下用户最后一次进房对应的网络类型。枚举值：2g、3g、4g、5g、wifi、unknown。
	Access *string `json:"Access,omitempty"`

	// 通话 Id，是通话的唯一标识。
	CallID *string `json:"CallId,omitempty"`

	// CallId下用户第一次进入通话时间,格式为 RFC3339，单位秒
	CreatedAt *string `json:"CreatedAt,omitempty"`

	// CallId下用户最后一次进房对应的设备型号
	DeviceType *string `json:"DeviceType,omitempty"`

	// CallId下用户从第一次进房到最后一次离开房间的时间范围内真实在线时长（多次进出房时间间隔累加），单位为秒
	Duration *int32 `json:"Duration,omitempty"`

	// CallId下用户是否离开房间
	Finished *bool `json:"Finished,omitempty"`

	// CallId下用户最后一次退出通话时间，格式为 RFC3339，单位秒。如果此时用户在线，返回为空。
	LeaveAt *string `json:"LeaveAt,omitempty"`

	// CallId下用户最后一次进房对应的设备平台。枚举值：android、ios、linux、mac、windows、web、unknown
	OS *string `json:"OS,omitempty"`

	// CallId下用户是否发布过流
	Pub *bool `json:"Pub,omitempty"`

	// CallId下用户在通话内全部进退房记录
	Record []*ListUserInfoResResultUserListPropertiesItemsItem `json:"Record,omitempty"`

	// 查询房间 Id
	RoomID *string `json:"RoomId,omitempty"`

	// CallId下用户最后一次进房使用 sdk 版本号
	SdkVersion *string `json:"SdkVersion,omitempty"`

	// 查询用户 Id
	UserID *string `json:"UserId,omitempty"`
}

type ListUserInfoResResultUserListPropertiesItemsItem struct {

	// 用户进房/退房时间，格式为 RFC3339，单位秒。
	Time *string `json:"Time,omitempty"`

	// 用户进房/退房类型，取值为 join_room或 leave_room。
	Type *string `json:"Type,omitempty"`
}

type ModifyAppStatusBody struct {

	// REQUIRED; 需要停用/启用的 AppId
	AppID string `json:"AppId"`

	// REQUIRED; * 1：将 App 状态设置为启用；
	// * 2：将 App 状态设置为停用；
	Status int32 `json:"Status"`
}

type ModifyAppStatusRes struct {

	// REQUIRED
	ResponseMetadata ModifyAppStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 返回值Result仅在请求成功时返回 "success",失败时为空。
	Result string `json:"Result"`
}

type ModifyAppStatusResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ModifyAppStatusResResponseMetadataError `json:"Error,omitempty"`
}

// ModifyAppStatusResResponseMetadataError - 仅在请求失败时返回。
type ModifyAppStatusResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ModifyBusinessRemarksBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 业务ID
	BusinessID string `json:"BusinessId"`

	// REQUIRED; 新业务名称，同一个 App 下不允许重复。长度不超过 24 个字符，取值范围为：数字、中文字符、字母和下划线
	Remarks string `json:"Remarks"`
}

type ModifyBusinessRemarksRes struct {

	// REQUIRED
	ResponseMetadata ModifyBusinessRemarksResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result string `json:"Result"`
}

type ModifyBusinessRemarksResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ModifyBusinessRemarksResResponseMetadataError `json:"Error,omitempty"`
}

// ModifyBusinessRemarksResResponseMetadataError - 仅在请求失败时返回。
type ModifyBusinessRemarksResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type SearchMusicsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// 过滤选项：
	// * 1：过滤没有歌词的歌曲
	// * 2：过滤不支持打分的歌曲
	// * 3：过滤不支持伴唱切换的歌曲
	// * 4：过滤没有高潮片段的歌曲
	Filters []*int32 `json:"Filters,omitempty"`

	// 搜索关键字。匹配优先级为：精准歌曲名匹配 -> 精准歌手名匹配 -> 歌曲名模糊匹配
	KeyWord *string `json:"KeyWord,omitempty"`

	// 分页序号，默认值为1
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页歌曲数量，取值范围为（0,100]，默认值为10
	PageSize *int32 `json:"PageSize,omitempty"`
}

type SearchMusicsRes struct {

	// REQUIRED
	ResponseMetadata SearchMusicsResResponseMetadata `json:"ResponseMetadata"`
	Result           *SearchMusicsResResult          `json:"Result,omitempty"`
}

type SearchMusicsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *SearchMusicsResResponseMetadataError `json:"Error,omitempty"`
}

// SearchMusicsResResponseMetadataError - 仅在请求失败时返回。
type SearchMusicsResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type SearchMusicsResResult struct {

	// 返回歌曲详细信息
	List []*SearchMusicsResResultListItem `json:"List,omitempty"`

	// 按照搜索条件返回的歌曲总数
	Total *int32 `json:"Total,omitempty"`
}

type SearchMusicsResResultListItem struct {

	// 歌曲是否支持伴唱原唱切换。
	// * 0： 不支持。
	// * 1： 声道切换
	// * 2： 音轨切换
	BgmType *int32 `json:"BgmType,omitempty"`

	// 歌曲时长，单位为秒
	Duration *int32 `json:"Duration,omitempty"`

	// 歌词类型
	// * 0： krc
	// * 1： lrc
	// * [0,1]： 两种歌词均有。 若为空，则表示为无歌词。
	LyricType []*int32                               `json:"LyricType,omitempty"`
	LyricURL  *SearchMusicsResResultListItemLyricURL `json:"LyricUrl,omitempty"`

	// 歌曲是否支持打分。
	// * 0： 不支持
	// * 1： 支持
	PitchType *int32 `json:"PitchType,omitempty"`

	// 歌曲封面地址
	PosterURL *string `json:"PosterUrl,omitempty"`

	// 歌曲是否支持录制。
	// * 0： 不支持
	// * 1： 支持
	RecordType *int32 `json:"RecordType,omitempty"`

	// 歌曲高潮部分
	Segment *string `json:"Segment,omitempty"`

	// 歌曲歌手名
	Singer *string `json:"Singer,omitempty"`

	// 歌曲 ID
	SongID *string `json:"SongId,omitempty"`

	// 歌曲名称
	Songname *string `json:"Songname,omitempty"`

	// 歌曲更新时间戳，unix 时间，单位为秒
	UpdateAt *int32 `json:"UpdateAt,omitempty"`

	// 供应商 ID
	VendorID *int32 `json:"VendorId,omitempty"`

	// 供应商名称
	VendorName *string `json:"VendorName,omitempty"`
}

type SearchMusicsResResultListItemLyricURL struct {

	// krc 歌词地址
	KrcURL *string `json:"KrcUrl,omitempty"`

	// lrc 歌词地址
	LrcURL *string `json:"LrcUrl,omitempty"`

	// midi 文件地址
	MidiURL *string `json:"MidiUrl,omitempty"`
}

type SendBroadcastBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; * 字段为 true，发送二进制消息；
	// * 字段为 false，发送文本消息。
	Binary bool `json:"Binary"`

	// REQUIRED; 业务服务端的唯一标识。 命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。 在一个 AppID 下，不能和真实用户用于实时消息通信的 user_ID 重复； 建议使用固定的 ID 的发送消息。
	From string `json:"From"`

	// REQUIRED; 广播消息内容。如果是二进制消息，需进行 base64 编码
	Message string `json:"Message"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`
}

type SendBroadcastRes struct {

	// REQUIRED
	ResponseMetadata SendBroadcastResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result SendBroadcastResResult `json:"Result"`
}

type SendBroadcastResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *SendBroadcastResResponseMetadataError `json:"Error,omitempty"`
}

// SendBroadcastResResponseMetadataError - 仅在请求失败时返回。
type SendBroadcastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type SendBroadcastResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 Success，失败时为空。
	Message string `json:"Message"`
}

type SendRoomUnicastBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; * 字段为 true，发送二进制消息；
	// * 字段为 false，发送文本消息。
	Binary bool `json:"Binary"`

	// REQUIRED; 业务服务端的唯一标识。 命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。 在一个 AppID 下，不能和真实用户用于实时消息通信的 user_ID 重复； 建议使用固定的 ID 的发送消息。
	From string `json:"From"`

	// REQUIRED; 点对点消息内容。如果是二进制消息，需进行 base64 编码
	Message string `json:"Message"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 消息接收用户调用 login 接口登录时设置的 ID，可用于接收房间内消息
	To string `json:"To"`
}

type SendRoomUnicastRes struct {

	// REQUIRED
	ResponseMetadata SendRoomUnicastResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 仅在请求成功时返回消息 "success"，表示服务端成功接收到消息，失败时为空。
	Result SendRoomUnicastResResult `json:"Result"`
}

type SendRoomUnicastResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *SendRoomUnicastResResponseMetadataError `json:"Error,omitempty"`
}

// SendRoomUnicastResResponseMetadataError - 仅在请求失败时返回。
type SendRoomUnicastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// SendRoomUnicastResResult - 仅在请求成功时返回消息 "success"，表示服务端成功接收到消息，失败时为空。
type SendRoomUnicastResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 Success，失败时为空。
	Message string `json:"Message"`
}

type SendUnicastBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; * 字段为 true，发送二进制消息；
	// * 字段为 false，发送文本消息。
	Binary bool `json:"Binary"`

	// REQUIRED; 业务服务端的唯一标识。命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。 在一个 AppID 下，不能和真实用户用于实时消息通信的 user_ID 重复； 建议使用固定的 ID 的发送消息。
	From string `json:"From"`

	// REQUIRED; 点对点消息内容。如果是二进制消息，需进行 base64 编码
	Message string `json:"Message"`

	// REQUIRED; 消息接收用户调用 login [70080#login] 接口登录时设置的 ID，可用于接收房间外消息
	To string `json:"To"`
}

type SendUnicastRes struct {

	// REQUIRED
	ResponseMetadata SendUnicastResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 仅在请求成功时返回消息 "success"，表示服务端成功接收到消息，失败时为空。
	Result SendUnicastResResult `json:"Result"`
}

type SendUnicastResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *SendUnicastResResponseMetadataError `json:"Error,omitempty"`
}

// SendUnicastResResponseMetadataError - 仅在请求失败时返回。
type SendUnicastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// SendUnicastResResult - 仅在请求成功时返回消息 "success"，表示服务端成功接收到消息，失败时为空。
type SendUnicastResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 Success，失败时为空。
	Message string `json:"Message"`
}

type StartDetectionBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// 回调地址 开启审核后，如果可能存在违规信息，此地址会收到违规信息回调。如果地址无效或为空，审核会继续，但不会收到违规信息的回调结果。
	Callback *string `json:"Callback,omitempty"`

	// 回调种类。
	// * 0：违规回调
	// * 1：全部回调。 默认值为 0。
	CallbackType *int32 `json:"CallbackType,omitempty"`

	// 每段音频切片的时长，范围为[1000，600000]。单位 ms。默认值为20000。该参数不建议设置过大，如果设置过大, 会出现审核延迟的问题，且造成最后一段切片计费误差向上取整偏大。
	Duration *int32 `json:"Duration,omitempty"`

	// 任务最大空闲超时时间。如果指定的用户停止推流或素材间隔过长，导致素材接收不到，那么审核任务会在空闲时间超过设定值后自动停止。值的范围为[1，10800]，单位为秒。默认值为180。
	IdleSec *int32 `json:"IdleSec,omitempty"`

	// 相邻截图之间的间隔时间，范围为[100，600000]。单位 ms。默认值为2000。
	Interval *int32 `json:"Interval,omitempty"`

	// 进行审核的内容类型：
	// * 1：视频截图；
	// * 2：音频切片；
	// * 3：视频截图+音频切片（默认值）
	// > 视频截图：审核过程中，RTC 会按照设定的时间间隔，进行周期性截图，并对截图进行审核。若出现违规信息，会返回审核结果。 音频切片：审核过程中，RTC 会按设定的音频切片时长，保存每段音频切片，并对切片进行审核。若出现违规信息，会返回审核结果。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 用户 ID，要审核的用户ID。若不填，则审核房间内所有推流用户。最多可以审核 17 路流。
	// * 如果先发起单流审核 再发起房间级审核，会发起房间级审核并停止单流审核；
	// * 如果先发起房间级审核，再发起单流审核，会引发错误，提示：已有审核进行中；
	// * 如果先发起单流音频审核和单流视频审核，再发起房间音频审核，会合并单流音频审核到房间音频审核，单流视频审核无变化；
	// * 如果先发起单流音频审核和单流视频审核，再发起房间音频审核，之后又发起房间音视频审核，会合并所有审核到房间音视频审核。
	UserID *string `json:"UserId,omitempty"`
}

type StartDetectionRes struct {

	// REQUIRED
	ResponseMetadata StartDetectionResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result StartDetectionResResult `json:"Result"`
}

type StartDetectionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartDetectionResResponseMetadataError `json:"Error,omitempty"`
}

// StartDetectionResResponseMetadataError - 仅在请求失败时返回。
type StartDetectionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartDetectionResResult struct {

	// REQUIRED; 仅在请求成功时返回"success",失败时为空
	Message string `json:"Message"`
}

type StartPushMixedStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 推流 CDN 地址。仅支持 RTMP 协议。
	PushURL string `json:"PushURL"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 合流转推任务 ID。你必须对每个合流转推任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 控制选项
	Control        *StartPushMixedStreamToCDNBodyControl        `json:"Control,omitempty"`
	Encode         *StartPushMixedStreamToCDNBodyEncode         `json:"Encode,omitempty"`
	ExcludeStreams *StartPushMixedStreamToCDNBodyExcludeStreams `json:"ExcludeStreams,omitempty"`
	Layout         *StartPushMixedStreamToCDNBodyLayout         `json:"Layout,omitempty"`
	TargetStreams  *StartPushMixedStreamToCDNBodyTargetStreams  `json:"TargetStreams,omitempty"`
}

// StartPushMixedStreamToCDNBodyControl - 控制选项
type StartPushMixedStreamToCDNBodyControl struct {

	// 选择补帧模式。默认值为0，可以取0和1。0为补最后一帧，1为补黑帧。值不合法时会自动调整为默认值。
	// 自动布局模式下，没有补帧的逻辑。
	// 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系:
	// 你可以在 Region 中传入 Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// （仅对录制有效）最大录制时长，单位为秒。默认值为 0。0 表示不限制录制时长。
	MaxRecordTime *int32 `json:"MaxRecordTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                             `json:"PushStreamMode,omitempty"`
	SpatialConfig  *StartPushMixedStreamToCDNBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`

	// (仅对转推直播有效）有效说话音量大小。取值范围为[0,255]，默认值为0。
	TalkVolume *int32 `json:"TalkVolume,omitempty"`

	// (仅对转推直播有效）用户说话音量的回调间隔，单位为秒，取值范围为[1,∞]，默认值为 2。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type StartPushMixedStreamToCDNBodyControlSpatialConfig struct {

	// 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward = [1, 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
	AudienceSpatialOrientation *StartPushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。
	// * false：关闭。
	// * true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

// StartPushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation - 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward
// = [1, 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
type StartPushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 头顶朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 前方朝向
	Up []*float32 `json:"Up,omitempty"`
}

type StartPushMixedStreamToCDNBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
	// 当AudioProfile=1且 AudioChannels = 1 时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	// 当AudioProfile=1且 AudioChannels = 2 时：
	// * 若输入参数取值范围为 [32,128]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [128,192]，编码码率固定为128。
	// 当AudioProfile=2时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。
	// * 1：单声道
	// * 2：双声道。 默认值为 2。值不合法时，自动调整为默认值。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。默认值为 0，此时使用 aac 编码协议。目前只支持 aac。值不合法时，自动调整为默认值。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型，在使用 aac 编码时生效。取值范围为 {0, 1, 2}。- 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。取 2 时，只支持输出流音频声道数为双声道。 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。默认值 48000，取值为 [32000,44100,48000]，单位是 Hz。值不合法时，自动调整为默认值。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 输出视频码率。取值范围 [1,10000]，单位为 Kbps，默认值为自适应。值不合法时，自动调整为默认值。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。默认值为 0，可以取 0 或 1。取 0 时使用 H.264,取 1 时使用 ByteVC1 编码器。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出视频帧率。默认为 15，取值范围为 [1,60]。值不合法时，自动调整为默认值。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。默认为 4，取值范围为 [1,5]，单位为秒。值不合法时，自动调整为默认值。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 输出画面的高度，范围为[2, 1920]，必须是偶数，默认值为480。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面高度以 canvas 中的 Height 为主。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 输出画面的宽度。默认值为640，范围为 [2, 1920]，必须是偶数。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面宽度以 canvas 中的 Width 为主。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type StartPushMixedStreamToCDNBodyExcludeStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartPushMixedStreamToCDNBodyExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartPushMixedStreamToCDNBodyExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayout struct {
	CustomLayout *StartPushMixedStreamToCDNBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                              `json:"LayoutMode,omitempty"`
	MainVideoStream *StartPushMixedStreamToCDNBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`

	// 在垂直布局模式下生效，指定主画面流的属性。如果此参数为空，则主画面为随机的一路流。该参数已废弃，当前版本 MainVideoStreamIndex 依然可用，但我们强烈建议你使用 MainVideoStream 参数。如果你同时指定了 MainVideoStream
	// 和 MainVideoStreamIndex 的值，此时只有 MainVideoStream 生效。
	MainVideoStreamIndex *int32 `json:"MainVideoStreamIndex,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayout struct {
	Canvas *StartPushMixedStreamToCDNBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即
	// Regions.Region.StreamIndex 要与
	// TargetStreams.StreamList.Stream.Index 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，用十六进制颜色码（HEX）表示。例如，#FFFFFF 表示纯白，#000000 表示纯黑。默认值为 #000000。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 值不合法时，自动调整为默认值。
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNBodyTargetStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartPushMixedStreamToCDNBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartPushMixedStreamToCDNBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StartPushMixedStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                      `json:"Result,omitempty"`
}

type StartPushMixedStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartPushMixedStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushMixedStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StartPushMixedStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartPushPublicStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 公共流 ID。你必须对每路公共流，设定 PublicStreamId 时命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	PublicStreamID string `json:"PublicStreamId"`

	// REQUIRED; 为公共流指定单路或多路媒体流及对应参数，Stream 数组。最高支持 17 路。
	TargetStreams []StartPushPublicStreamBodyTargetStreamsItem `json:"TargetStreams"`

	// 你创建的业务标识
	BusinessID *string                           `json:"BusinessId,omitempty"`
	Control    *StartPushPublicStreamBodyControl `json:"Control,omitempty"`
	Encode     *StartPushPublicStreamBodyEncode  `json:"Encode,omitempty"`

	// 你可以通过本参数排除掉不需要包含在公共流中的音视频流，即黑名单。参数默认为空。黑名单中的流不得超过 17 路。此参数中的 stream 不应和 TargetStreams 中重复。
	ExcludeStreams []*StartPushPublicStreamBodyExcludeStreamsItem `json:"ExcludeStreams,omitempty"`
	Layout         *StartPushPublicStreamBodyLayout               `json:"Layout,omitempty"`

	// 公共流处理模式。0：转码。1：转封装。
	// 当 TranscodeMode=1 时，
	// * TargetStreams 只能指定一路流，且该路流的 UserId不能为空，需为对应房间用户的 UserId。
	// * ExcludeStreams 必须为空。
	// * Encode.VideoConfig 设置不生效。
	// * Layout 设置不生效。
	TranscodeMode *int32 `json:"TranscodeMode,omitempty"`
}

type StartPushPublicStreamBodyControl struct {

	// 插入公共流的自定义信息，可用于随流信息同步，长度不超过 4 kB。 数据会添加到当前视频帧开始的连续 30 个视频帧中。 只在调用UpdatePublicStreamParam时有效。
	DataMsg *string `json:"DataMsg,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为[10, 86400]，默认值为180。只在调用StartPushPublicStream时有效。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 房间用户发布状态回调间隔，仅在纯音频时触发。单位为毫秒，默认值为 2000，取值范围为 [1000,2147483647]。 值不合法时自动调整为默认值。
	StreamPublishStatsInterval *int32 `json:"StreamPublishStatsInterval,omitempty"`

	// 是否开启房间用户发布状态回调，仅在纯音频时触发。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户发布状态回调。
	// * false：不开启房间用户发布状态回调。 默认值为false。
	StreamPublishStatsMode *bool `json:"StreamPublishStatsMode,omitempty"`

	// 房间用户采集状态回调间隔，仅在纯音频时触发。单位为毫秒，默认值为2000，取值范围为[1000,2147483647]。 值不合法时自动调整为默认值。
	UserCaptureStatsInterval *int32 `json:"UserCaptureStatsInterval,omitempty"`

	// 是否开启房间用户采集状态回调，仅在纯音频时触发。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户采集状态回调。
	// * false：不开启房间用户采集状态回调。 默认值为false。
	UserCaptureStatsMode *bool `json:"UserCaptureStatsMode,omitempty"`

	// 音量指示的回调间隔。单位为毫秒，最小值为100，默认值为2000。 值不合法时自动调整为默认值。
	// VideoConfig.FrameRate大于 10 fps 时，回调间隔才能达到 100ms。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// 是否开启音量指示模式。
	// * true：开启音量提示。
	// * false：不开启音量提示。 默认值为false。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type StartPushPublicStreamBodyEncode struct {

	// 视频编码配置
	VideoConfig *StartPushPublicStreamBodyEncodeVideoConfig `json:"VideoConfig,omitempty"`
}

// StartPushPublicStreamBodyEncodeVideoConfig - 视频编码配置
type StartPushPublicStreamBodyEncodeVideoConfig struct {

	// 最高输出视频码率。取值范围[**0,10000**]，单位为 Kbps，默认值0，为0时表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 输出视频帧率。默认为15，取值范围为 [1,60]。单位 fps
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 输出画面的高度，默认值为 480。范围为 [16, 1920]，必须是偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频编码协议。可取值为0或5，默认值为0。
	// * 0：H.264。
	// * 5：VP8。 如果选择 VP8 格式，请先联系火山技术支持配置。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出画面的宽度。默认值为 640，范围为 [16, 1920]，必须是偶数。
	Width *int32 `json:"Width,omitempty"`
}

type StartPushPublicStreamBodyExcludeStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。默认值为 0 。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。默认值为 0。
	// * 0：媒体设备采集到的音视频
	// * 1：屏幕流
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。 UserId 为空时，表示订阅房间内所有流。 UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type StartPushPublicStreamBodyLayout struct {

	// REQUIRED; 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [https://www.volcengine.com/docs/6348/115995#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [https://www.volcengine.com/docs/6348/115995#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [https://www.volcengine.com/docs/6348/115995#horizontal]
	LayoutMode     int32                                          `json:"LayoutMode"`
	CustomLayout   *StartPushPublicStreamBodyLayoutCustomLayout   `json:"CustomLayout,omitempty"`
	VerticalLayout *StartPushPublicStreamBodyLayoutVerticalLayout `json:"VerticalLayout,omitempty"`
}

type StartPushPublicStreamBodyLayoutCustomLayout struct {

	// REQUIRED; 自定义布局下，多路视频配置
	Regions []StartPushPublicStreamBodyLayoutCustomLayoutRegionsItem `json:"Regions"`

	// 背整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。值不合法时，自动调整为默认值。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundColor *string `json:"BackgroundColor,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 选择补帧模式。默认值为0，可以取0和1。0为补最后一帧，1为补黑帧。值不合法时会自动调整为默认值。
	// 自动布局模式下，没有补帧的逻辑。
	// 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系: 你可以在 Region 中传入 Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`
}

type StartPushPublicStreamBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流标识。
	// StreamIndex 即 Stream.Index，用来指定布局设置将被应用到哪路流。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为(0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 占位图片的 url
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的纵坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：- 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。值不合法时，自动调整为默认值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartPushPublicStreamBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 当画面有重叠时，使用此参数设置指定画面的图层顺序，取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartPushPublicStreamBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartPushPublicStreamBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartPushPublicStreamBodyLayoutVerticalLayout struct {
	MainStream *StartPushPublicStreamBodyLayoutVerticalLayoutMainStream `json:"MainStream,omitempty"`
}

type StartPushPublicStreamBodyLayoutVerticalLayoutMainStream struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。默认值为 0 。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。默认值为 0。
	// * 0：媒体设备采集到的音视频
	// * 1：屏幕流
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。 UserId 为空时，表示订阅房间内所有流。 UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type StartPushPublicStreamBodyTargetStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。默认值为 0 。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。默认值为 0。
	// * 0：媒体设备采集到的音视频
	// * 1：屏幕流
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。 UserId 为空时，表示订阅房间内所有流。 UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type StartPushPublicStreamRes struct {

	// REQUIRED
	ResponseMetadata StartPushPublicStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                  `json:"Result,omitempty"`
}

type StartPushPublicStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartPushPublicStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushPublicStreamResResponseMetadataError - 仅在请求失败时返回。
type StartPushPublicStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartPushSingleStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED
	Control StartPushSingleStreamToCDNBodyControl `json:"Control"`

	// REQUIRED; 推流地址。目前仅支持 rtmp 协议
	PushURL string `json:"PushURL"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED
	Stream StartPushSingleStreamToCDNBodyStream `json:"Stream"`

	// REQUIRED; 转推任务 ID。你必须对每个转推任务设定 TaskId，且在后续进行任务结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StartPushSingleStreamToCDNBodyControl struct {

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为 0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type StartPushSingleStreamToCDNBodyStream struct {

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 用户Id，表示这个流所属的用户。
	UserID *string `json:"UserId,omitempty"`
}

type StartPushSingleStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StartPushSingleStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                       `json:"Result,omitempty"`
}

type StartPushSingleStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartPushSingleStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushSingleStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StartPushSingleStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// REQUIRED; 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod] 服务的配置信息，用于存储录制结果。
	Vod StartRecordBodyVod `json:"Vod"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 控制选项
	Control        *StartRecordBodyControl        `json:"Control,omitempty"`
	Encode         *StartRecordBodyEncode         `json:"Encode,omitempty"`
	ExcludeStreams *StartRecordBodyExcludeStreams `json:"ExcludeStreams,omitempty"`
	Layout         *StartRecordBodyLayout         `json:"Layout,omitempty"`

	// 自定义录制文件名模式，具体参看自定义录制文件名 [106873]。
	RecordFileNamePattern *string `json:"RecordFileNamePattern,omitempty"`

	// 使用此参数设定录制模式：单流/合流录制。0 表示合流录制，1 表示单流录制。默认值为 0。
	RecordMode    *int32                        `json:"RecordMode,omitempty"`
	TargetStreams *StartRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

// StartRecordBodyControl - 控制选项
type StartRecordBodyControl struct {

	// 选择补帧模式。默认值为0，可以取0和1。0为补最后一帧，1为补黑帧。值不合法时会自动调整为默认值。
	// 自动布局模式下，没有补帧的逻辑。
	// 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系:
	// 你可以在 Region 中传入 Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// （仅对录制有效）最大录制时长，单位为秒。默认值为 0。0 表示不限制录制时长。
	MaxRecordTime *int32 `json:"MaxRecordTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                               `json:"PushStreamMode,omitempty"`
	SpatialConfig  *StartRecordBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`

	// (仅对转推直播有效）有效说话音量大小。取值范围为[0,255]，默认值为0。
	TalkVolume *int32 `json:"TalkVolume,omitempty"`

	// (仅对转推直播有效）用户说话音量的回调间隔，单位为秒，取值范围为[1,∞]，默认值为 2。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type StartRecordBodyControlSpatialConfig struct {

	// 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward = [1, 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
	AudienceSpatialOrientation *StartRecordBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。
	// * false：关闭。
	// * true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

// StartRecordBodyControlSpatialConfigAudienceSpatialOrientation - 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward = [1,
// 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
type StartRecordBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 头顶朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 前方朝向
	Up []*float32 `json:"Up,omitempty"`
}

type StartRecordBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
	// 当AudioProfile=1且 AudioChannels = 1 时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	// 当AudioProfile=1且 AudioChannels = 2 时：
	// * 若输入参数取值范围为 [32,128]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [128,192]，编码码率固定为128。
	// 当AudioProfile=2时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。
	// * 1：单声道
	// * 2：双声道。 默认值为 2。值不合法时，自动调整为默认值。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。默认值为 0，此时使用 aac 编码协议。目前只支持 aac。值不合法时，自动调整为默认值。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型，在使用 aac 编码时生效。取值范围为 {0, 1, 2}。- 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。取 2 时，只支持输出流音频声道数为双声道。 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。默认值 48000，取值为 [32000,44100,48000]，单位是 Hz。值不合法时，自动调整为默认值。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 输出视频码率。取值范围 [1,10000]，单位为 Kbps，默认值为自适应。值不合法时，自动调整为默认值。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。默认值为 0，可以取 0 或 1。取 0 时使用 H.264,取 1 时使用 ByteVC1 编码器。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出视频帧率。默认为 15，取值范围为 [1,60]。值不合法时，自动调整为默认值。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。默认为 4，取值范围为 [1,5]，单位为秒。值不合法时，自动调整为默认值。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 输出画面的高度，范围为[2, 1920]，必须是偶数，默认值为480。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面高度以 canvas 中的 Height 为主。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 输出画面的宽度。默认值为640，范围为 [2, 1920]，必须是偶数。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面宽度以 canvas 中的 Width 为主。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type StartRecordBodyExcludeStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartRecordBodyExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordBodyLayout struct {
	CustomLayout *StartRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                `json:"LayoutMode,omitempty"`
	MainVideoStream *StartRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`

	// 在垂直布局模式下生效，指定主画面流的属性。如果此参数为空，则主画面为随机的一路流。该参数已废弃，当前版本 MainVideoStreamIndex 依然可用，但我们强烈建议你使用 MainVideoStream 参数。如果你同时指定了 MainVideoStream
	// 和 MainVideoStreamIndex 的值，此时只有 MainVideoStream 生效。
	MainVideoStreamIndex *int32 `json:"MainVideoStreamIndex,omitempty"`
}

type StartRecordBodyLayoutCustomLayout struct {
	Canvas *StartRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即
	// Regions.Region.StreamIndex 要与
	// TargetStreams.StreamList.Stream.Index 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*StartRecordBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type StartRecordBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，用十六进制颜色码（HEX）表示。例如，#FFFFFF 表示纯白，#000000 表示纯黑。默认值为 #000000。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type StartRecordBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 值不合法时，自动调整为默认值。
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartRecordBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordBodyTargetStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

// StartRecordBodyVod - 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod] 服务的配置信息，用于存储录制结果。
type StartRecordBodyVod struct {

	// 平台账号ID，例如：200000000。查看路径：控制台->账号管理->主账号信息。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID *string `json:"AccountId,omitempty"`

	// 用于存储录制结果的点播空间名。
	Space *string `json:"Space,omitempty"`
}

type StartRecordRes struct {

	// REQUIRED
	ResponseMetadata StartRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
}

type StartRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StartRecordResResponseMetadataError - 仅在请求失败时返回。
type StartRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartRelayStreamBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED
	Control StartRelayStreamBodyControl `json:"Control"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// REQUIRED; 客户端与业务服务端进行通讯时用于身份认证的 token 值
	Token string `json:"Token"`

	// REQUIRED; 在线媒体流对应的的 UserId
	UserID string `json:"UserId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 任务的空闲超时时间。超过此时间后，任务自动终止。值的范围为 [5, 600] ，单位为秒。默认值为300。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`
}

type StartRelayStreamBodyControl struct {

	// REQUIRED; 在线流媒体地址。媒体格式应为：hls、rtmp、mp4、flv、dash、或 ts。如源流为海外，建议联系技术支持，以保障最佳体验。
	StreamURL string `json:"StreamUrl"`

	// 最大发送码率，单位为 Kbps,不填则不限制，转码时生效。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，值的范围为[1，30]，默认值为15，转码时生效。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 是否循环播放，仅对源流为点播流时生效。
	Loop *bool `json:"Loop,omitempty"`

	// 媒体类型。
	// * 0：音视频
	// * 1：音频。采用此选项时，必须是 AAC 或 Opus 编码。
	// * 2：视频 默认值为0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 任务起始时间戳，用于定时播放，Unix时间，单位为秒。默认为 0，表示立即启动。此参数仅对 StartRelayStream接口生效。
	StartTimeStamp *int32 `json:"StartTimeStamp,omitempty"`

	// 流处理模式。
	// * 0：转码。采用此选项时，原视频编码方式必须是 H.264 或 ByteVC1。
	// * 1：转封装。采用此选项时，原视频编码方式必须是 H.264。转封装时，源流的视频关键帧间隔若过大，会影响 RTC 体验，建议 1s，但最大不超过 5s。 默认值为0。
	StreamMode *int32 `json:"StreamMode,omitempty"`

	// 视频高度，转码时必填。单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 视频宽度。转码时必填，单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type StartRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata StartRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                             `json:"Result,omitempty"`
}

type StartRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StartRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type StartRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartSegmentBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 截图任务 ID。你可以自行设定TaskId以区分不同的切片任务，且在后续更新和结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// REQUIRED
	TosConfig StartSegmentBodyTosConfig `json:"TosConfig"`

	// 业务标识。添加 BusinessId 后，你可以在控制台上查看根据 BusinessId 划分的质量和用量数据。
	BusinessID *string `json:"BusinessId,omitempty"`

	// 切片高级功能
	Control *StartSegmentBodyControl `json:"Control,omitempty"`

	// 每个音频切片的时长。单位为秒。范围为[1, 600]，默认值为20秒。值不合法时，自动调整为默认值。
	Duration *int32 `json:"Duration,omitempty"`

	// 是否在开启音视频切片时，立刻开始切片。 true 为在开启音视频切片时，立刻开启切片；false 时，在开启音视频切片时，不开始切片，热启动任务。 默认值为true。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。 切片文件名由 Identifier + UserId + 时间戳 + 序号组成。默认情况下 Identifier 为 随机生成的 UUID。在自定义文件名时，Identifier 的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。
	// 在自定义文件名时，你需确保 Identifier 命名全局唯一，否则在 TOS 平台会因文件名重复被覆盖。
	Identifier *string `json:"Identifier,omitempty"`

	// 任务最大的空闲超时时间。如果切片任务订阅的所有流都已停止发布，那么任务会在空闲时间超过设定值后自动停止。值的范围为[1, 86400] ，单位为秒。默认值为 180秒。值不合法时，自动调整为默认值。
	MaxIdleTime   *int32                         `json:"MaxIdleTime,omitempty"`
	TargetStreams *StartSegmentBodyTargetStreams `json:"TargetStreams,omitempty"`
}

// StartSegmentBodyControl - 切片高级功能
type StartSegmentBodyControl struct {

	// Align 是音频切片对齐功能的开关，默认为 False。你可以使用音频切片对齐功能，对齐各个用户音频切片的开始和结束时刻。
	// * 当 Align=False 时，关闭音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。如果用户在切片的周期中，有部分时间未发布音频，返回的音频切片时长会小于切片周期。各个用户音频切片开始时间不一定一致。
	// * 当 Align=True 时，开启音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。切片长度和切片周期相等，且各个用户音频切片开始的时间一致。如果用户在切片的周期中，有部分时间未发布音频，切片长度不变，这段时间呈现静音帧。如果用户在某个切片周期中始终没有发布音频，则不生成音频切片。
	Align *bool `json:"Align,omitempty"`

	// 是否忽略静音切片。默认值为 false。
	IgnoreSilence *bool `json:"IgnoreSilence,omitempty"`

	// Mixed 是音频切片合流功能的开关。默认为 False。
	// * 当 Mixed=False 时，只会对 TargetStreams 中指定的音频流分别切片。
	//
	//
	// * 当 Mixed=True 时，除了会对 TargetStreams 中指定的音频流分别切片，还会对指定的音频流进行混音，生成合流切片，合流切片对应的用户名为 mixed。此时，任务创建后，不管是否有人上麦，会持续回调混音切片。
	//
	// 不同平台的回调参看：
	//
	//
	// 操作 ANDROID API IOS API WINDOWS API
	// 本地麦克风录制和远端所有用户混音后的音频数据回调 onMixedAudioFrame [70081#onmixedaudioframe] onMixedAudioFrame: [70087#onmixedaudioframe] onMixedAudioFrame
	// [70096#onmixedaudioframe]
	Mixed *bool `json:"Mixed,omitempty"`

	// 冗余切片时长，单位为秒。 当前服务端按照传入的Duration值进行固定时长切片,可能存在敏感词被截断，未被识别的情况。此时你可以添加冗余切片，即上一段切片的末尾指定时长，来避免此情况，此时切片的时长变为RedundantDuration
	// + Duration。 例如：当 Duration = 20，redundantDuration = 3 时，最终输出的前三个文件时长为：[0-20],
	// [17-40], [37-60]。
	RedundantDuration *int32 `json:"RedundantDuration,omitempty"`
}

type StartSegmentBodyTargetStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartSegmentBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartSegmentBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartSegmentBodyTosConfig struct {

	// REQUIRED; 存储桶名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; Bucket 对应的访问域名。参看 地域和访问域名 [https://www.volcengine.com/docs/6349/107356]。
	// 该 API 下 EndPoint 只支持 S3 Endpoint 外网域名。
	EndPoint string `json:"EndPoint"`

	// REQUIRED; Bucket 所在地域对应的 Region ID。参看 地域和访问域名 [https://www.volcengine.com/docs/6349/107356]。
	Region string `json:"Region"`

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	UserAccountID string `json:"UserAccountId"`
}

type StartSegmentRes struct {

	// REQUIRED
	ResponseMetadata StartSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type StartSegmentResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// StartSegmentResResponseMetadataError - 仅在请求失败时返回。
type StartSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartSnapshotBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 截图任务 ID。你可以自行设定TaskId以区分不同的切片任务，且在后续进行任务更新和结束时也须使用该 TaskId。关于 TaskId 的详细说明，参看备注 1 。 关于 TaskId 及以上 Id 字段的命名规则，参看
	// ID [69835#____id]。
	TaskID string `json:"TaskId"`

	// REQUIRED
	TosConfig StartSnapshotBodyTosConfig `json:"TosConfig"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 图片的相关配置：图片格式，尺寸和截图间隔时间。
	ImageConfig *StartSnapshotBodyImageConfig `json:"ImageConfig,omitempty"`

	// 任务最大的空闲超时时间。如果抽帧截图任务订阅的所有流都已停止发布，那么任务会在空闲时间超过设定值后自动停止。值的范围为[1, 86400] ，单位为秒。默认值为 180秒。值不合法时，自动调整为默认值。
	MaxIdleTime   *int32                          `json:"MaxIdleTime,omitempty"`
	TargetStreams *StartSnapshotBodyTargetStreams `json:"TargetStreams,omitempty"`
}

// StartSnapshotBodyImageConfig - 图片的相关配置：图片格式，尺寸和截图间隔时间。
type StartSnapshotBodyImageConfig struct {

	// 图片的格式。值可取0或1，默认为0。选择0时，图片格式为 .jpeg；选择1时，图片格式为 .png。默认值为0。值不合法时，自动调整为默认值。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，单位为 px，取值范围为[0, 1080]，默认值为0，此时，和视频流的实际高度相同。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，单位为 秒，取值范围为[1, 600]，默认值为2。值不合法时，自动调整为默认值。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，单位为 px，取值范围为[0, 1920]。默认值为0，此时，和视频流的实际宽度相同。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type StartSnapshotBodyTargetStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartSnapshotBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartSnapshotBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartSnapshotBodyTosConfig struct {

	// REQUIRED; 存储桶名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; Bucket 对应的访问域名。参看 地域和访问域名 [https://www.volcengine.com/docs/6349/107356]。
	// 该 API 下 EndPoint 只支持 S3 Endpoint 外网域名。
	EndPoint string `json:"EndPoint"`

	// REQUIRED; Bucket 所在地域对应的 Region ID。参看 地域和访问域名 [https://www.volcengine.com/docs/6349/107356]。
	Region string `json:"Region"`

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	UserAccountID string `json:"UserAccountId"`
}

type StartSnapshotRes struct {

	// REQUIRED
	ResponseMetadata StartSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                          `json:"Result,omitempty"`
}

type StartSnapshotResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// StartSnapshotResResponseMetadataError - 仅在请求失败时返回。
type StartSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartWBRecordBody struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"AppId"`

	// REQUIRED; 需要录制的白板房间 ID，同一个 appId 中，为每个房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制任务 ID。你可以自行设定 TaskId 以区分不同的白板录制任务。 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// REQUIRED; 任务发起方的用户 ID。不能与房间中其他用户的 ID 重复，否则先进房的用户会被移出房间。
	UserID string `json:"UserId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 自定义 UI 的 Web 页面地址。 不填表示使用默认白板页面。
	SourceURL *string `json:"SourceURL,omitempty"`
}

type StartWBRecordRes struct {

	// REQUIRED
	ResponseMetadata StartWBRecordResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StartWBRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartWBRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StartWBRecordResResponseMetadataError - 仅在请求失败时返回。
type StartWBRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartWebcastBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 向指定 RTC 房间推送网页音视频内容，房间 ID 是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 需要转推的网页地址，可以携带自定义的 queryParams 来鉴权等，总体长度不超过 1024。
	SourceURL string `json:"SourceURL"`

	// REQUIRED; 任务 ID。你必须对每个云录屏任务设定 TaskId，且在后续结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// REQUIRED; 推送网页音视频内容的用户对应的 UserId。不能与房间中其他用户的 ID 重复，否则先进房的用户会被移出房间。 建议添加有规律的前缀，避免重复。例如， webcast_。
	UserID string `json:"UserId"`

	// 业务标识
	BusinessID        *string                            `json:"BusinessId,omitempty"`
	EventNotifyConfig *StartWebcastBodyEventNotifyConfig `json:"EventNotifyConfig,omitempty"`

	// 最大运行时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10,86400]，默认值为 86400。不填时自动调整为默认值。
	MaxRunningTime *int32                         `json:"MaxRunningTime,omitempty"`
	MonitorConfig  *StartWebcastBodyMonitorConfig `json:"MonitorConfig,omitempty"`

	// 输出流的视频参数。网页渲染分辨率为输入最大视频流的分辨率。该值为空时按默认值填充一路。
	// * 当设置一路流时，不会输出多路流，输出流参数与设置值相同。
	// * 当设置两路流时， RTC 会启动推送多路流 [70139]功能，设置的两路分辨率相当于分辨率等级划分中的大流和中流的上限。
	VideoSolutions []*StartWebcastBodyVideoSolutionsItem `json:"VideoSolutions,omitempty"`
}

type StartWebcastBodyEventNotifyConfig struct {

	// 是否启用页面主动事件通知, 默认值为false。
	// * false：页面在打开后就会开始采集，在收到 StopWebCast openAPI 请求后结束采集。
	// * true：在页面中注入两个 JS 函数：onWebcastStart()和 onWebcastEnd()。
	// 当页面判断资源加载完成之后调用onWebcastStart()，控制程序才会开始进行页面内容的采集。当页面判断本次任务内容已完成时调用onWebcastEnd() 通知控制程序结束本次任务。StopWebCast openAPI 效果不变，业务可提前结束任务。其他页面内容、JS
	// 线程的检测（若启用），将在收到 onWebcastStart()事件后才开始。
	// 当启用页面主动事件通知后，你可以参考以下示例代码来通知采集开始。
	// <script>
	// if (ready() && typeof onWebcastStart === 'function')
	// onWebcastStart();
	// </script>
	EnableEventNotify *bool `json:"EnableEventNotify,omitempty"`

	// 启用页面主动事件通知后，等待开始事件的超时时间。取值范围为[0,60]，单位为秒。默认值为0，表示不启用。仅当 EnableEventNotify 为 true 时，此参数有效。
	// * 当在超时时间内收到开始事件，采集功能正常运行，用户将收到 Status=1的回调。
	// * 当超时时间内未收到开始事件，将进行刷新，等待时间被重置，再次发生超时后将进行任务重调度。刷新时将回调 Status=4，Reason=" StartEventTimeout"。重调度时将回调 Status=5，Reason="StartEventTimeout"。
	StartTimeout *int32 `json:"StartTimeout,omitempty"`
}

type StartWebcastBodyMonitorConfig struct {

	// 对页面是否白屏的检测间隔。取值范围为[2,30]，单位为秒。默认值为0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageBlank" 。
	// * 再次出现连续两次检测命中时将进行任务重调度，并回调Status=5，Reason="PageBlank"。
	// 注意：页面全白可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	BlankCheckInterval *int32 `json:"BlankCheckInterval,omitempty"`

	// 对页面 JS 线程是否崩溃/卡死的检测间隔。 取值范围为[0,30]，单位为秒。默认值为0，表示不启用。
	// 当出现检测命中时将进行任务重调度，并回调 Status=5，Reason="PageCrash"。
	CrashCheckInterval *int32 `json:"CrashCheckInterval,omitempty"`

	// 对页面内容是否无变化的检测间隔。取值范围为[2,30]，单位为秒。默认值为0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageFreeze"。
	// * 再次出现连续两次检测命中时，将进行任务重调度，并回调Status=5，Reason="PageFreeze"。
	// 注意：页面无变化可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	FreezeCheckInterval *int32 `json:"FreezeCheckInterval,omitempty"`
}

type StartWebcastBodyVideoSolutionsItem struct {

	// 最大发送码率，取值范围为[0,10000]，单位为 Kbps，默认值 0，为 0 时表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，单位为 帧/秒，范围为[1,60]，默认值为 15。帧率和码率设置建议参照视频发布参数对照表 [70122#param]以获取最佳体验。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 视频高度，单位为像素，范围为 [50,1080]，默认值为 720。必须是偶数，值为奇数时自动调整为偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频宽度，单位为像素，范围为 [50,1920]，默认值为 1280。必须是偶数，值为奇数时自动调整为偶数。
	Width *int32 `json:"Width,omitempty"`
}

type StartWebcastRes struct {

	// REQUIRED
	ResponseMetadata StartWebcastResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StartWebcastResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartWebcastResResponseMetadataError `json:"Error,omitempty"`
}

// StartWebcastResResponseMetadataError - 仅在请求失败时返回。
type StartWebcastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopDetectionBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// 用户 ID
	UserID *string `json:"UserId,omitempty"`
}

type StopDetectionRes struct {

	// REQUIRED
	ResponseMetadata StopDetectionResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result StopDetectionResResult `json:"Result"`
}

type StopDetectionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopDetectionResResponseMetadataError `json:"Error,omitempty"`
}

// StopDetectionResResponseMetadataError - 仅在请求失败时返回。
type StopDetectionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopDetectionResResult struct {

	// REQUIRED; 仅在请求成功时返回"success",失败时为空。
	Message string `json:"Message"`
}

type StopPushPublicStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 公共流 ID。
	PublicStreamID string `json:"PublicStreamId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopPushPublicStreamRes struct {

	// REQUIRED
	ResponseMetadata StopPushPublicStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                 `json:"Result,omitempty"`
}

type StopPushPublicStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopPushPublicStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StopPushPublicStreamResResponseMetadataError - 仅在请求失败时返回。
type StopPushPublicStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopPushStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 转推任务 ID。你必须对每个转推任务设定 TaskId，且在进行任务结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type StopPushStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StopPushStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                `json:"Result,omitempty"`
}

type StopPushStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopPushStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StopPushStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StopPushStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在进行结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopRecordRes struct {

	// REQUIRED
	ResponseMetadata StopRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                       `json:"Result,omitempty"`
}

type StopRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StopRecordResResponseMetadataError - 仅在请求失败时返回。
type StopRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopRelayStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个任务设定 TaskId，且在进行结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type StopRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata StopRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                            `json:"Result,omitempty"`
}

type StopRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StopRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type StopRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopSegmentBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 你需要关闭的音频切片任务的 ID。
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopSegmentRes struct {

	// REQUIRED
	ResponseMetadata StopSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
}

type StopSegmentResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// StopSegmentResResponseMetadataError - 仅在请求失败时返回。
type StopSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopSnapshotBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopSnapshotRes struct {

	// REQUIRED
	ResponseMetadata StopSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type StopSnapshotResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// StopSnapshotResResponseMetadataError - 仅在请求失败时返回。
type StopSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopWBRecordBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，同一个 appId 中，每个房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制任务 ID。调用 StartWBRecord 时使用的任务 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 调用接口的用户 ID
	UserID string `json:"UserId"`
}

type StopWBRecordRes struct {

	// REQUIRED
	ResponseMetadata StopWBRecordResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopWBRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopWBRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StopWBRecordResResponseMetadataError - 仅在请求失败时返回。
type StopWBRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopWebcastBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个云录屏任务设定 TaskId，且在结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type StopWebcastRes struct {

	// REQUIRED
	ResponseMetadata StopWebcastResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
}

type StopWebcastResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopWebcastResResponseMetadataError `json:"Error,omitempty"`
}

// StopWebcastResResponseMetadataError - 仅在请求失败时返回。
type StopWebcastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UnbanUserStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 用于校验当前账号是否具有解封权限的 Token，生成方式与加入房间时的Token生成方式一致
	Token string `json:"Token"`

	// REQUIRED; 需要被解封音/视频流的用户的 ID
	UserID string `json:"UserId"`

	// 是否解封音频流，置为 true 时，表示解封音频流
	Audio *bool `json:"Audio,omitempty"`

	// 是否解封视频流，置为 true 时，表示解封视频流
	Video *bool `json:"Video,omitempty"`
}

type UnbanUserStreamRes struct {

	// REQUIRED
	ResponseMetadata UnbanUserStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *UnbanUserStreamResResult          `json:"Result,omitempty"`
}

type UnbanUserStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error UnbanUserStreamResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UnbanUserStreamResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

type UnbanUserStreamResResult struct {

	// REQUIRED
	Message string `json:"message"`
}

type UpdateBanRoomUserRuleBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 指定房间 ID
	RoomID string `json:"RoomId"`

	// 封禁时长，单位为秒，取值范围为[60,259290]。 若传入值为空或 0表示允许用户重新进房。 若传入值大于0，且小于60，自动修改为60。 若传入值大于259290，自动修改为259290。
	ForbiddenInterval *int32 `json:"ForbiddenInterval,omitempty"`

	// 希望封禁用户的 ID
	UserID *string `json:"UserId,omitempty"`
}

type UpdateBanRoomUserRuleRes struct {

	// REQUIRED
	ResponseMetadata UpdateBanRoomUserRuleResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result UpdateBanRoomUserRuleResResult `json:"Result"`
}

type UpdateBanRoomUserRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error UpdateBanRoomUserRuleResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateBanRoomUserRuleResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

type UpdateBanRoomUserRuleResResult struct {

	// REQUIRED
	Message string `json:"message"`
}

type UpdateCallbackBody struct {

	// REQUIRED
	AppID string `json:"AppID"`

	// REQUIRED
	BusinessID string `json:"BusinessID"`

	// REQUIRED
	EventList []string `json:"EventList"`

	// REQUIRED
	SecretKey string `json:"SecretKey"`

	// REQUIRED
	URL string `json:"URL"`
}

type UpdateCallbackRes struct {

	// REQUIRED
	ResponseMetadata UpdateCallbackResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result string `json:"Result"`
}

type UpdateCallbackResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateCallbackResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateCallbackResResponseMetadataError - 仅在请求失败时返回。
type UpdateCallbackResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateFailRecoveryPolicyBody struct {

	// REQUIRED; Policy ID，用以指定一条FailRecoveryPolicy
	ID int32 `json:"ID"`

	// REQUIRED; 0~100的整数，故障线路恢复比例
	RecoveryPercentage int32 `json:"RecoveryPercentage"`
}

type UpdateFailRecoveryPolicyRes struct {

	// REQUIRED
	ResponseMetadata UpdateFailRecoveryPolicyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *UpdateFailRecoveryPolicyResResult `json:"Result,omitempty"`
}

type UpdateFailRecoveryPolicyResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateFailRecoveryPolicyResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateFailRecoveryPolicyResResponseMetadataError - 仅在请求失败时返回。
type UpdateFailRecoveryPolicyResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// UpdateFailRecoveryPolicyResResult - 视请求的接口而定
type UpdateFailRecoveryPolicyResResult struct {

	// REQUIRED; 更新后的策略的详细信息
	Policy UpdateFailRecoveryPolicyResResultPolicy `json:"Policy"`
}

// UpdateFailRecoveryPolicyResResultPolicy - 更新后的策略的详细信息
type UpdateFailRecoveryPolicyResResultPolicy struct {

	// REQUIRED; RTC应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 创建时间（毫秒）
	CreateTime int32 `json:"CreateTime"`

	// REQUIRED; 故障线路
	FromVendor string `json:"FromVendor"`

	// REQUIRED; 数据条目的ID。后面删除和修改都需要指定该ID
	ID int32 `json:"ID"`

	// REQUIRED; 0~100的整数，故障线路恢复的百分比
	RecoveryPercentage int32 `json:"RecoveryPercentage"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 切换至线路
	ToVendor string `json:"ToVendor"`

	// REQUIRED; 更新时间（毫秒）
	UpdateTime string `json:"UpdateTime"`
}

type UpdatePublicStreamParamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 公共流 ID。
	PublicStreamID string `json:"PublicStreamId"`

	// REQUIRED; 为公共流指定单路或多路媒体流及对应参数，Stream 数组。最高支持 17 路。
	TargetStreams []UpdatePublicStreamParamBodyTargetStreamsItem `json:"TargetStreams"`

	// 业务标识
	BusinessID *string                             `json:"BusinessId,omitempty"`
	Control    *UpdatePublicStreamParamBodyControl `json:"Control,omitempty"`
	Encode     *UpdatePublicStreamParamBodyEncode  `json:"Encode,omitempty"`

	// 你可以通过本参数排除掉不需要包含在公共流中的音视频流，即黑名单。参数默认为空。黑名单中的流不得超过 17 路。此参数中的 stream 不应和 TargetStreams 中重复。
	ExcludeStreams []*UpdatePublicStreamParamBodyExcludeStreamsItem `json:"ExcludeStreams,omitempty"`
	Layout         *UpdatePublicStreamParamBodyLayout               `json:"Layout,omitempty"`

	// 公共流处理模式。0：转码。1：转封装。
	// 当 TranscodeMode=1 时，
	// * TargetStreams 只能指定一路流，且该路流的 UserId不能为空，需为对应房间用户的 UserId。
	// * ExcludeStreams 必须为空。
	// * Encode.VideoConfig 设置不生效。
	// * Layout 设置不生效。
	TranscodeMode *int32 `json:"TranscodeMode,omitempty"`
}

type UpdatePublicStreamParamBodyControl struct {

	// 插入公共流的自定义信息，可用于随流信息同步，长度不超过 4 kB。 数据会添加到当前视频帧开始的连续 30 个视频帧中。 只在调用UpdatePublicStreamParam时有效。
	DataMsg *string `json:"DataMsg,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为[10, 86400]，默认值为180。只在调用StartPushPublicStream时有效。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 房间用户发布状态回调间隔，仅在纯音频时触发。单位为毫秒，默认值为 2000，取值范围为 [1000,2147483647]。 值不合法时自动调整为默认值。
	StreamPublishStatsInterval *int32 `json:"StreamPublishStatsInterval,omitempty"`

	// 是否开启房间用户发布状态回调，仅在纯音频时触发。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户发布状态回调。
	// * false：不开启房间用户发布状态回调。 默认值为false。
	StreamPublishStatsMode *bool `json:"StreamPublishStatsMode,omitempty"`

	// 房间用户采集状态回调间隔，仅在纯音频时触发。单位为毫秒，默认值为2000，取值范围为[1000,2147483647]。 值不合法时自动调整为默认值。
	UserCaptureStatsInterval *int32 `json:"UserCaptureStatsInterval,omitempty"`

	// 是否开启房间用户采集状态回调，仅在纯音频时触发。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户采集状态回调。
	// * false：不开启房间用户采集状态回调。 默认值为false。
	UserCaptureStatsMode *bool `json:"UserCaptureStatsMode,omitempty"`

	// 音量指示的回调间隔。单位为毫秒，最小值为100，默认值为2000。 值不合法时自动调整为默认值。
	// VideoConfig.FrameRate大于 10 fps 时，回调间隔才能达到 100ms。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// 是否开启音量指示模式。
	// * true：开启音量提示。
	// * false：不开启音量提示。 默认值为false。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type UpdatePublicStreamParamBodyEncode struct {

	// 视频编码配置
	VideoConfig *UpdatePublicStreamParamBodyEncodeVideoConfig `json:"VideoConfig,omitempty"`
}

// UpdatePublicStreamParamBodyEncodeVideoConfig - 视频编码配置
type UpdatePublicStreamParamBodyEncodeVideoConfig struct {

	// 最高输出视频码率。取值范围[**0,10000**]，单位为 Kbps，默认值0，为0时表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 输出视频帧率。默认为15，取值范围为 [1,60]。单位 fps
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 输出画面的高度，默认值为 480。范围为 [16, 1920]，必须是偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频编码协议。可取值为0或5，默认值为0。
	// * 0：H.264。
	// * 5：VP8。 如果选择 VP8 格式，请先联系火山技术支持配置。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出画面的宽度。默认值为 640，范围为 [16, 1920]，必须是偶数。
	Width *int32 `json:"Width,omitempty"`
}

type UpdatePublicStreamParamBodyExcludeStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。默认值为 0 。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。默认值为 0。
	// * 0：媒体设备采集到的音视频
	// * 1：屏幕流
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。 UserId 为空时，表示订阅房间内所有流。 UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type UpdatePublicStreamParamBodyLayout struct {

	// REQUIRED; 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [https://www.volcengine.com/docs/6348/115995#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [https://www.volcengine.com/docs/6348/115995#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [https://www.volcengine.com/docs/6348/115995#horizontal]
	LayoutMode     int32                                            `json:"LayoutMode"`
	CustomLayout   *UpdatePublicStreamParamBodyLayoutCustomLayout   `json:"CustomLayout,omitempty"`
	VerticalLayout *UpdatePublicStreamParamBodyLayoutVerticalLayout `json:"VerticalLayout,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutCustomLayout struct {

	// REQUIRED; 自定义布局下，多路视频配置
	Regions []UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItem `json:"Regions"`

	// 背整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。值不合法时，自动调整为默认值。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundColor *string `json:"BackgroundColor,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 选择补帧模式。默认值为0，可以取0和1。0为补最后一帧，1为补黑帧。值不合法时会自动调整为默认值。
	// 自动布局模式下，没有补帧的逻辑。
	// 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系: 你可以在 Region 中传入 Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流标识。
	// StreamIndex 即 Stream.Index，用来指定布局设置将被应用到哪路流。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为(0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 占位图片的 url
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的纵坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：- 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。值不合法时，自动调整为默认值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 当画面有重叠时，使用此参数设置指定画面的图层顺序，取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutVerticalLayout struct {
	MainStream *UpdatePublicStreamParamBodyLayoutVerticalLayoutMainStream `json:"MainStream,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutVerticalLayoutMainStream struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。默认值为 0 。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。默认值为 0。
	// * 0：媒体设备采集到的音视频
	// * 1：屏幕流
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。 UserId 为空时，表示订阅房间内所有流。 UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type UpdatePublicStreamParamBodyTargetStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。默认值为 0 。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。默认值为 0。
	// * 0：媒体设备采集到的音视频
	// * 1：屏幕流
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。 UserId 为空时，表示订阅房间内所有流。 UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type UpdatePublicStreamParamRes struct {

	// REQUIRED
	ResponseMetadata UpdatePublicStreamParamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                    `json:"Result,omitempty"`
}

type UpdatePublicStreamParamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdatePublicStreamParamResResponseMetadataError `json:"Error,omitempty"`
}

// UpdatePublicStreamParamResResponseMetadataError - 仅在请求失败时返回。
type UpdatePublicStreamParamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdatePushMixedStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 合流转推任务 ID。你必须对每个合流转推任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 控制选项
	Control *UpdatePushMixedStreamToCDNBodyControl `json:"Control,omitempty"`
	Encode  *UpdatePushMixedStreamToCDNBodyEncode  `json:"Encode,omitempty"`

	// 是否更新部分参数。
	// * False：否。
	// * True：是。
	// 默认值为 False。
	// 开启部分更新后，必须按照参数层级传入，且数组类参数需要传入该数组中所有参数。
	IsUpdatePartialParam *bool                                 `json:"IsUpdatePartialParam,omitempty"`
	Layout               *UpdatePushMixedStreamToCDNBodyLayout `json:"Layout,omitempty"`

	// 更新请求序列号。填写该参数后，服务端会对请求进行校验，请确保最后一次更新请求的序列号大于前一次请求的序列号。
	// 建议更新部分参数场景下传入此参数，以确保服务端按照最新请求进行更新。
	SequenceNumber *int32                                       `json:"SequenceNumber,omitempty"`
	TargetStreams  *UpdatePushMixedStreamToCDNBodyTargetStreams `json:"TargetStreams,omitempty"`
}

// UpdatePushMixedStreamToCDNBodyControl - 控制选项
type UpdatePushMixedStreamToCDNBodyControl struct {

	// 选择补帧模式。默认值为0，可以取0和1。0为补最后一帧，1为补黑帧。值不合法时会自动调整为默认值。
	// 自动布局模式下，没有补帧的逻辑。
	// 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系:
	// 你可以在 Region 中传入 Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// （仅对录制有效）最大录制时长，单位为秒。默认值为 0。0 表示不限制录制时长。
	MaxRecordTime *int32 `json:"MaxRecordTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                              `json:"PushStreamMode,omitempty"`
	SpatialConfig  *UpdatePushMixedStreamToCDNBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`

	// (仅对转推直播有效）有效说话音量大小。取值范围为[0,255]，默认值为0。
	TalkVolume *int32 `json:"TalkVolume,omitempty"`

	// (仅对转推直播有效）用户说话音量的回调间隔，单位为秒，取值范围为[1,∞]，默认值为 2。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControlSpatialConfig struct {

	// 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward = [1, 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
	AudienceSpatialOrientation *UpdatePushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。
	// * false：关闭。
	// * true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

// UpdatePushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation - 设置观众朝向。各个向量两两垂直，如果传入的值没有保证两两垂直，自动赋予默认值。默认值为：forward
// = [1, 0, 0], right = [0, 1, 0], up = [0, 0, 1]。
type UpdatePushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 头顶朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 前方朝向
	Up []*float32 `json:"Up,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
	// 当AudioProfile=1且 AudioChannels = 1 时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	// 当AudioProfile=1且 AudioChannels = 2 时：
	// * 若输入参数取值范围为 [32,128]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [128,192]，编码码率固定为128。
	// 当AudioProfile=2时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。
	// * 1：单声道
	// * 2：双声道。 默认值为 2。值不合法时，自动调整为默认值。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。默认值为 0，此时使用 aac 编码协议。目前只支持 aac。值不合法时，自动调整为默认值。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型，在使用 aac 编码时生效。取值范围为 {0, 1, 2}。- 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。取 2 时，只支持输出流音频声道数为双声道。 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。默认值 48000，取值为 [32000,44100,48000]，单位是 Hz。值不合法时，自动调整为默认值。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 输出视频码率。取值范围 [1,10000]，单位为 Kbps，默认值为自适应。值不合法时，自动调整为默认值。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。默认值为 0，可以取 0 或 1。取 0 时使用 H.264,取 1 时使用 ByteVC1 编码器。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出视频帧率。默认为 15，取值范围为 [1,60]。值不合法时，自动调整为默认值。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。默认为 4，取值范围为 [1,5]，单位为秒。值不合法时，自动调整为默认值。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 输出画面的高度，范围为[2, 1920]，必须是偶数，默认值为480。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面高度以 canvas 中的 Height 为主。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 输出画面的宽度。默认值为640，范围为 [2, 1920]，必须是偶数。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面宽度以 canvas 中的 Width 为主。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayout struct {
	CustomLayout *UpdatePushMixedStreamToCDNBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                               `json:"LayoutMode,omitempty"`
	MainVideoStream *UpdatePushMixedStreamToCDNBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`

	// 在垂直布局模式下生效，指定主画面流的属性。如果此参数为空，则主画面为随机的一路流。该参数已废弃，当前版本 MainVideoStreamIndex 依然可用，但我们强烈建议你使用 MainVideoStream 参数。如果你同时指定了 MainVideoStream
	// 和 MainVideoStreamIndex 的值，此时只有 MainVideoStream 生效。
	MainVideoStreamIndex *int32 `json:"MainVideoStreamIndex,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayout struct {
	Canvas *UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即
	// Regions.Region.StreamIndex 要与
	// TargetStreams.StreamList.Stream.Index 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，用十六进制颜色码（HEX）表示。例如，#FFFFFF 表示纯白，#000000 表示纯黑。默认值为 #000000。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 值不合法时，自动调整为默认值。
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyTargetStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*UpdatePushMixedStreamToCDNBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdatePushMixedStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata UpdatePushMixedStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                       `json:"Result,omitempty"`
}

type UpdatePushMixedStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdatePushMixedStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// UpdatePushMixedStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type UpdatePushMixedStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID    *string                        `json:"BusinessId,omitempty"`
	Layout        *UpdateRecordBodyLayout        `json:"Layout,omitempty"`
	TargetStreams *UpdateRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type UpdateRecordBodyLayout struct {
	CustomLayout *UpdateRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                 `json:"LayoutMode,omitempty"`
	MainVideoStream *UpdateRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`

	// 在垂直布局模式下生效，指定主画面流的属性。如果此参数为空，则主画面为随机的一路流。该参数已废弃，当前版本 MainVideoStreamIndex 依然可用，但我们强烈建议你使用 MainVideoStream 参数。如果你同时指定了 MainVideoStream
	// 和 MainVideoStreamIndex 的值，此时只有 MainVideoStream 生效。
	MainVideoStreamIndex *int32 `json:"MainVideoStreamIndex,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayout struct {
	Canvas *UpdateRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即
	// Regions.Region.StreamIndex 要与
	// TargetStreams.StreamList.Stream.Index 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*UpdateRecordBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，用十六进制颜色码（HEX）表示。例如，#FFFFFF 表示纯白，#000000 表示纯黑。默认值为 #000000。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 值不合法时，自动调整为默认值。
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdateRecordBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdateRecordBodyTargetStreams struct {

	// 由 Stream 组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*UpdateRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type UpdateRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdateRecordRes struct {

	// REQUIRED
	ResponseMetadata UpdateRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type UpdateRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateRecordResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateRecordResResponseMetadataError - 仅在请求失败时返回。
type UpdateRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateRelayStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED
	Control UpdateRelayStreamBodyControl `json:"Control"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个任务设定 TaskId，且在进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type UpdateRelayStreamBodyControl struct {

	// REQUIRED; 在线流媒体地址。媒体格式应为：hls、rtmp、mp4、flv、dash、或 ts。如源流为海外，建议联系技术支持，以保障最佳体验。
	StreamURL string `json:"StreamUrl"`

	// 最大发送码率，单位为 Kbps,不填则不限制，转码时生效。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，值的范围为[1，30]，默认值为15，转码时生效。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 是否循环播放，仅对源流为点播流时生效。
	Loop *bool `json:"Loop,omitempty"`

	// 媒体类型。
	// * 0：音视频
	// * 1：音频。采用此选项时，必须是 AAC 或 Opus 编码。
	// * 2：视频 默认值为0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 任务起始时间戳，用于定时播放，Unix时间，单位为秒。默认为 0，表示立即启动。此参数仅对 StartRelayStream接口生效。
	StartTimeStamp *int32 `json:"StartTimeStamp,omitempty"`

	// 流处理模式。
	// * 0：转码。采用此选项时，原视频编码方式必须是 H.264 或 ByteVC1。
	// * 1：转封装。采用此选项时，原视频编码方式必须是 H.264。转封装时，源流的视频关键帧间隔若过大，会影响 RTC 体验，建议 1s，但最大不超过 5s。 默认值为0。
	StreamMode *int32 `json:"StreamMode,omitempty"`

	// 视频高度，转码时必填。单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 视频宽度。转码时必填，单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type UpdateRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata UpdateRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                              `json:"Result,omitempty"`
}

type UpdateRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type UpdateRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateSegmentBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 切片任务 ID。你必须对每个切片任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。关于 TaskId 的详细说明，参看TaskId说明 。 关于 TaskId 及以上 Id 字段的命名规则，参看 ID
	// [69835#____id]。
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 音频切片的时长。单位为秒。范围为 [1, 600]，默认值为 20秒。值不合法时，自动调整为默认值。 更新该字段后，计时器会重新启动，当前切片立即停止，生成一个新切片。
	Duration *int32 `json:"Duration,omitempty"`

	// 是否在开启音视切片时，立刻开始切片。选择 True 时，开启切片；选择 False 时，不切片。默认值 true。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。切片名设置详情，参看切片名设置规则
	Identifier *string `json:"Identifier,omitempty"`
}

type UpdateSegmentRes struct {

	// REQUIRED
	ResponseMetadata UpdateSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                          `json:"Result,omitempty"`
}

type UpdateSegmentResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateSegmentResResponseMetadataError - 仅在请求失败时返回。
type UpdateSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateSnapshotBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 截图任务 ID。你可以自行设定TaskId以区分不同的切片任务，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID  *string                        `json:"BusinessId,omitempty"`
	ImageConfig *UpdateSnapshotBodyImageConfig `json:"ImageConfig,omitempty"`
}

type UpdateSnapshotBodyImageConfig struct {

	// 图片的格式。值可取0或1，默认为0。选择0时，图片格式为 JEPG；选择1时，图片格式为 PNG。默认值为0。值不合法时，自动调整为默认值。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，单位为像素，取值范围为[0, 1920]，默认值为0，此时，和视频流的实际高度相同。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，单位为秒，取值范围为[1, 600]，默认值为2。值不合法时，自动调整为默认值。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，单位为像素，取值范围为[0, 1920]。默认值为0，此时，和视频流的实际宽度相同。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type UpdateSnapshotRes struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSnapshotResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateSnapshotResResponseMetadataError - 仅在请求失败时返回。
type UpdateSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateVendorPolicyBody struct {

	// REQUIRED; Policy ID，用以指定一条VendorPolicy
	ID int32 `json:"ID"`

	// REQUIRED; 【Json类型】各线路放量比例，例： {"agora": 50, "trtc": 50}
	VendorProportion string `json:"VendorProportion"`
}

type UpdateVendorPolicyRes struct {

	// REQUIRED
	ResponseMetadata UpdateVendorPolicyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *UpdateVendorPolicyResResult `json:"Result,omitempty"`
}

type UpdateVendorPolicyResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateVendorPolicyResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateVendorPolicyResResponseMetadataError - 仅在请求失败时返回。
type UpdateVendorPolicyResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// UpdateVendorPolicyResResult - 视请求的接口而定
type UpdateVendorPolicyResResult struct {

	// REQUIRED; 更新后的策略详情
	Policy UpdateVendorPolicyResResultPolicy `json:"Policy"`
}

// UpdateVendorPolicyResResultPolicy - 更新后的策略详情
type UpdateVendorPolicyResResultPolicy struct {

	// REQUIRED; 应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 应用标识
	BusinessID string `json:"BusinessID"`

	// REQUIRED; 创建时间（毫秒）
	CreateTime int32 `json:"CreateTime"`

	// REQUIRED; 数据条目的ID。后面删除和修改都需要指定该ID
	ID int32 `json:"ID"`

	// REQUIRED; 房间ID
	RoomID string `json:"RoomID"`

	// REQUIRED; 更新时间（毫秒）
	UpdateTime int32 `json:"UpdateTime"`

	// REQUIRED; 各线路放量比例
	VendorProportion string `json:"VendorProportion"`
}

type WbTranscodeCreateBody struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"app_id"`

	// REQUIRED; 用户 ID
	Operator string `json:"operator"`

	// REQUIRED; 需要转换为图片的文档链接地址。每次调用接口只能请求处理一份文档。
	Resource string `json:"resource"`

	// 转码成功后返回的 URL 预签名有效期，单位为秒，取值范围为：[0,604800]。默认值为 0。0表示无限期。为 0 时，你需要在 TOS 服务的 bucket 为公共读。
	PreSignDuration *int32 `json:"pre_sign_duration,omitempty"`

	// 静态转码的转码优先级
	// * 0: 非实时转码
	// * 1: 实时转码 默认值为 0。
	Priority *int32 `json:"priority,omitempty"`

	// 动态转码文件设置。动态转码必填，静态转码无需填写
	ResourceAttr *WbTranscodeCreateBodyResourceAttr `json:"resource_attr,omitempty"`

	// 对象存储属性。使用火山引擎的对象存储服务，且本次传入的参数与控制台设置的属性有差异，则以传入参数为准。
	// * 使用第三方的对象存储服务：必填
	// * 使用火山引擎的对象存储服务，并已经在控制台设置了图片存储属性，则此参数选填
	StorageConfig *WbTranscodeCreateBodyStorageConfig `json:"storage_config,omitempty"`

	// 转码参数设置
	TranscodeConfig *WbTranscodeCreateBodyTranscodeConfig `json:"transcode_config,omitempty"`

	// 转码类型
	// * 0: 静态转码
	// * 1: 动态转码 默认值为 0。
	TranscodeMode *int32 `json:"transcode_mode,omitempty"`
}

// WbTranscodeCreateBodyResourceAttr - 动态转码文件设置。动态转码必填，静态转码无需填写
type WbTranscodeCreateBodyResourceAttr struct {

	// REQUIRED; 文件名
	FileName string `json:"file_name"`

	// REQUIRED; 文件大小，单位：byte
	Size int32 `json:"size"`
}

// WbTranscodeCreateBodyStorageConfig - 对象存储属性。使用火山引擎的对象存储服务，且本次传入的参数与控制台设置的属性有差异，则以传入参数为准。
// * 使用第三方的对象存储服务：必填
// * 使用火山引擎的对象存储服务，并已经在控制台设置了图片存储属性，则此参数选填
type WbTranscodeCreateBodyStorageConfig struct {

	// REQUIRED; 存储类型 【默认值】0: Tos1: 第三方对象存储接口，支持阿里云和亚马逊
	Type int32 `json:"type"`

	// 第三方对象存储服务参数设置
	CustomConfig *WbTranscodeCreateBodyStorageConfigCustomConfig `json:"custom_config,omitempty"`

	// 火山引擎的对象存储服务参数设置
	TosConfig *WbTranscodeCreateBodyStorageConfigTosConfig `json:"tos_config,omitempty"`
}

// WbTranscodeCreateBodyStorageConfigCustomConfig - 第三方对象存储服务参数设置
type WbTranscodeCreateBodyStorageConfigCustomConfig struct {

	// REQUIRED; Access Key
	AccessKey string `json:"access_key"`

	// REQUIRED; 桶名称
	Bucket string `json:"bucket"`

	// REQUIRED; 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1217553]
	Region int32 `json:"region"`

	// REQUIRED; Secret Key AK/SK建议只开通写权限，关闭读权限。
	SecretKey string `json:"secret_key"`

	// REQUIRED; 第三方存储供应商0: Amazon 亚马逊1: Alicloud 阿里云
	Vendor int32 `json:"vendor"`
}

// WbTranscodeCreateBodyStorageConfigTosConfig - 火山引擎的对象存储服务参数设置
type WbTranscodeCreateBodyStorageConfigTosConfig struct {

	// REQUIRED; Bucket 所属的火山引擎账号 ID。在登录火山引擎后，可在头像的悬浮菜单中找到账号 ID。[https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_0819c44c6aadff358a7dfc52c5daab57.png]
	AccountID string `json:"account_id"`

	// REQUIRED; 桶名称。登录TOS 控制台 [https://console.volcengine.com/tos/bucket]开通和获取。
	Bucket string `json:"bucket"`

	// REQUIRED; * 0： cn-beijing，华北 2（北京）
	// * 1：cn-guangzhou，华南 1（广州）
	// * 2：cn-shanghai，华东 2（上海）
	Region int32 `json:"region"`
}

// WbTranscodeCreateBodyTranscodeConfig - 转码参数设置
type WbTranscodeCreateBodyTranscodeConfig struct {

	// REQUIRED; 输入文件类型1: ppt2: pptx3: doc4: docx5: pdf
	InputFormat int32 `json:"input_format"`

	// REQUIRED; 输出文件类型1: png2: jpg/jpeg
	OutputFormat int32 `json:"output_format"`

	// REQUIRED; 转码后的页面高度，单位为像素
	OutputHeight int32 `json:"output_height"`

	// REQUIRED; 转码后的页面宽度，单位为像素
	OutputWidth int32 `json:"output_width"`

	// 是否按照指定分辨率拉伸页面 默认为false，按照文件的原始宽高比适配指定分辨率。
	ForceUseResolution *bool `json:"force_use_resolution,omitempty"`

	// 是否生成缩略图，默认为false
	Thumbnail *bool `json:"thumbnail,omitempty"`

	// 缩略图分辨率高，默认为180
	ThumbnailHeight *int32 `json:"thumbnail_height,omitempty"`

	// 缩略图分辨率宽，默认为320
	ThumbnailWidth *int32 `json:"thumbnail_width,omitempty"`
}

type WbTranscodeCreateRes struct {

	// REQUIRED
	ResponseMetadata WbTranscodeCreateResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *WbTranscodeCreateResResult `json:"Result,omitempty"`
}

type WbTranscodeCreateResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *WbTranscodeCreateResResponseMetadataError `json:"Error,omitempty"`
}

// WbTranscodeCreateResResponseMetadataError - 仅在请求失败时返回。
type WbTranscodeCreateResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// WbTranscodeCreateResResult - 视请求的接口而定
type WbTranscodeCreateResResult struct {

	// REQUIRED; task_id 为查询任务进度和获取转码后链接的必填参数
	TaskID string `json:"task_id"`
}

type WbTranscodeGetQuery struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"app_id" query:"app_id"`

	// REQUIRED; 成功调用WbTranscodeQuery后返回的任务ID
	TaskID string `json:"task_id" query:"task_id"`
}

type WbTranscodeGetRes struct {

	// REQUIRED
	ResponseMetadata WbTranscodeGetResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result WbTranscodeGetResResult `json:"Result"`
}

type WbTranscodeGetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *WbTranscodeGetResResponseMetadataError `json:"Error,omitempty"`
}

// WbTranscodeGetResResponseMetadataError - 仅在请求失败时返回。
type WbTranscodeGetResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type WbTranscodeGetResResult struct {

	// REQUIRED; 动态转码结果FileId，使用动态转码结果时在SDK接口传入
	FileID string `json:"file_id"`

	// REQUIRED; 转码文件名
	FileName string `json:"file_name"`

	// REQUIRED; 转码后的页面高度，单位为像素
	Height int32 `json:"height"`

	// REQUIRED; 静态转码图片详情
	Images []WbTranscodeGetResResultImagesItem `json:"images"`

	// REQUIRED; 转码类型
	// * 0: 静态转码
	// * 1: 动态转码
	TranscodeMode int32 `json:"transcode_mode"`

	// REQUIRED; 转码后的页面宽度，单位为像素
	Width int32 `json:"width"`
}

type WbTranscodeGetResResultImagesItem struct {

	// REQUIRED; 转码图片 URL
	Img string `json:"img"`

	// REQUIRED; 页码 ID
	PageID int32 `json:"page_id"`

	// REQUIRED; 缩略图 URL
	ThumbnailURL string `json:"thumbnail_url"`
}

type WbTranscodeQueryQuery struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"app_id" query:"app_id"`

	// REQUIRED; 成功调用WbTranscodeQuery后返回的任务ID
	TaskID string `json:"task_id" query:"task_id"`
}

type WbTranscodeQueryRes struct {

	// REQUIRED
	ResponseMetadata WbTranscodeQueryResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result WbTranscodeQueryResResult `json:"Result"`
}

type WbTranscodeQueryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *WbTranscodeQueryResResponseMetadataError `json:"Error,omitempty"`
}

// WbTranscodeQueryResResponseMetadataError - 仅在请求失败时返回。
type WbTranscodeQueryResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type WbTranscodeQueryResResult struct {

	// REQUIRED; 转码任务状态
	// * 0：未开始
	// * 1：进行中
	// * 2：已完成
	// * 3：失败
	Status int32 `json:"status"`

	// 错误码
	ErrCode *int32 `json:"err_code,omitempty"`

	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty"`

	// 静态转码进度。100表示已完成。 动态转码恒定为0。
	Progress *int32 `json:"progress,omitempty"`
}
type AddBusinessID struct{}
type AddBusinessIDQuery struct{}
type BanRoomUser struct{}
type BanRoomUserQuery struct{}
type BanUserStream struct{}
type BanUserStreamQuery struct{}
type BatchSendRoomUnicast struct{}
type BatchSendRoomUnicastQuery struct{}
type CreateApp struct{}
type CreateAppQuery struct{}
type CreateByteplusApp struct{}
type CreateByteplusAppBody struct{}
type CreateByteplusAppQuery struct{}
type CreateCallback struct{}
type CreateCallbackQuery struct{}
type CreateFailRecoveryPolicy struct{}
type CreateFailRecoveryPolicyQuery struct{}
type CreateVendorPolicy struct{}
type CreateVendorPolicyQuery struct{}
type DeleteBusinessID struct{}
type DeleteBusinessIDQuery struct{}
type DeleteCallback struct{}
type DeleteCallbackBody struct{}
type DeleteFailRecoveryPolicy struct{}
type DeleteFailRecoveryPolicyQuery struct{}
type DismissRoom struct{}
type DismissRoomQuery struct{}
type GetAllBusinessID struct{}
type GetAllBusinessIDQuery struct{}
type GetCallback struct{}
type GetCallbackBody struct{}
type GetFailRecoveryPolicies struct{}
type GetFailRecoveryPoliciesQuery struct{}
type GetPushMixedStreamToCDNTask struct{}
type GetPushMixedStreamToCDNTaskBody struct{}
type GetPushSingleStreamToCDNTask struct{}
type GetPushSingleStreamToCDNTaskBody struct{}
type GetRecordTask struct{}
type GetRecordTaskBody struct{}
type GetResourcePackNum struct{}
type GetResourcePackNumBody struct{}
type GetResourcePackNumQuery struct{}
type GetWebCastTask struct{}
type GetWebCastTaskBody struct{}
type KickUser struct{}
type KickUserQuery struct{}
type ListApps struct{}
type ListAppsQuery struct{}
type ListCallDetail struct{}
type ListCallDetailBody struct{}
type ListDetectionTask struct{}
type ListDetectionTaskBody struct{}
type ListHotMusic struct{}
type ListHotMusicQuery struct{}
type ListMusics struct{}
type ListMusicsQuery struct{}
type ListOperationData struct{}
type ListOperationDataQuery struct{}
type ListOperationDistribution struct{}
type ListOperationDistributionQuery struct{}
type ListQuality struct{}
type ListQualityDistribution struct{}
type ListQualityDistributionQuery struct{}
type ListQualityQuery struct{}
type ListRealTimeOperationData struct{}
type ListRealTimeOperationDataQuery struct{}
type ListRealTimeQuality struct{}
type ListRealTimeQualityDistribution struct{}
type ListRealTimeQualityDistributionQuery struct{}
type ListRealTimeQualityQuery struct{}
type ListRelayStream struct{}
type ListRelayStreamBody struct{}
type ListResourcePackages struct{}
type ListResourcePackagesBody struct{}
type ListResourcePackagesQuery struct{}
type ListResourcePackagesV2 struct{}
type ListResourcePackagesV2Body struct{}
type ListResourcePackagesV2Query struct{}
type ListRoomInfo struct{}
type ListRoomInfoBody struct{}
type ListUsages struct{}
type ListUsagesQuery struct{}
type ListUserInfo struct{}
type ListUserInfoBody struct{}
type ModifyAppStatus struct{}
type ModifyAppStatusQuery struct{}
type ModifyBusinessRemarks struct{}
type ModifyBusinessRemarksQuery struct{}
type SearchMusics struct{}
type SearchMusicsQuery struct{}
type SendBroadcast struct{}
type SendBroadcastQuery struct{}
type SendRoomUnicast struct{}
type SendRoomUnicastQuery struct{}
type SendUnicast struct{}
type SendUnicastQuery struct{}
type StartDetection struct{}
type StartDetectionQuery struct{}
type StartPushMixedStreamToCDN struct{}
type StartPushMixedStreamToCDNQuery struct{}
type StartPushPublicStream struct{}
type StartPushPublicStreamQuery struct{}
type StartPushSingleStreamToCDN struct{}
type StartPushSingleStreamToCDNQuery struct{}
type StartRecord struct{}
type StartRecordQuery struct{}
type StartRelayStream struct{}
type StartRelayStreamQuery struct{}
type StartSegment struct{}
type StartSegmentQuery struct{}
type StartSnapshot struct{}
type StartSnapshotQuery struct{}
type StartWBRecord struct{}
type StartWBRecordQuery struct{}
type StartWebcast struct{}
type StartWebcastQuery struct{}
type StopDetection struct{}
type StopDetectionQuery struct{}
type StopPushPublicStream struct{}
type StopPushPublicStreamQuery struct{}
type StopPushStreamToCDN struct{}
type StopPushStreamToCDNQuery struct{}
type StopRecord struct{}
type StopRecordQuery struct{}
type StopRelayStream struct{}
type StopRelayStreamQuery struct{}
type StopSegment struct{}
type StopSegmentQuery struct{}
type StopSnapshot struct{}
type StopSnapshotQuery struct{}
type StopWBRecord struct{}
type StopWBRecordQuery struct{}
type StopWebcast struct{}
type StopWebcastQuery struct{}
type UnbanUserStream struct{}
type UnbanUserStreamQuery struct{}
type UpdateBanRoomUserRule struct{}
type UpdateBanRoomUserRuleQuery struct{}
type UpdateCallback struct{}
type UpdateCallbackQuery struct{}
type UpdateFailRecoveryPolicy struct{}
type UpdateFailRecoveryPolicyQuery struct{}
type UpdatePublicStreamParam struct{}
type UpdatePublicStreamParamQuery struct{}
type UpdatePushMixedStreamToCDN struct{}
type UpdatePushMixedStreamToCDNQuery struct{}
type UpdateRecord struct{}
type UpdateRecordQuery struct{}
type UpdateRelayStream struct{}
type UpdateRelayStreamQuery struct{}
type UpdateSegment struct{}
type UpdateSegmentQuery struct{}
type UpdateSnapshot struct{}
type UpdateSnapshotQuery struct{}
type UpdateVendorPolicy struct{}
type UpdateVendorPolicyQuery struct{}
type WbTranscodeCreate struct{}
type WbTranscodeCreateQuery struct{}
type WbTranscodeGet struct{}
type WbTranscodeGetBody struct{}
type WbTranscodeQueryBody struct{}
type AddBusinessIDReq struct {
	*AddBusinessIDQuery
	*AddBusinessIDBody
}
type BanRoomUserReq struct {
	*BanRoomUserQuery
	*BanRoomUserBody
}
type BanUserStreamReq struct {
	*BanUserStreamQuery
	*BanUserStreamBody
}
type BatchSendRoomUnicastReq struct {
	*BatchSendRoomUnicastQuery
	*BatchSendRoomUnicastBody
}
type CreateAppReq struct {
	*CreateAppQuery
	*CreateAppBody
}
type CreateByteplusAppReq struct {
	*CreateByteplusAppQuery
	*CreateByteplusAppBody
}
type CreateCallbackReq struct {
	*CreateCallbackQuery
	*CreateCallbackBody
}
type CreateFailRecoveryPolicyReq struct {
	*CreateFailRecoveryPolicyQuery
	*CreateFailRecoveryPolicyBody
}
type CreateVendorPolicyReq struct {
	*CreateVendorPolicyQuery
	*CreateVendorPolicyBody
}
type DeleteBusinessIDReq struct {
	*DeleteBusinessIDQuery
	*DeleteBusinessIDBody
}
type DeleteCallbackReq struct {
	*DeleteCallbackQuery
	*DeleteCallbackBody
}
type DeleteFailRecoveryPolicyReq struct {
	*DeleteFailRecoveryPolicyQuery
	*DeleteFailRecoveryPolicyBody
}
type DismissRoomReq struct {
	*DismissRoomQuery
	*DismissRoomBody
}
type GetAllBusinessIDReq struct {
	*GetAllBusinessIDQuery
	*GetAllBusinessIDBody
}
type GetCallbackReq struct {
	*GetCallbackQuery
	*GetCallbackBody
}
type GetFailRecoveryPoliciesReq struct {
	*GetFailRecoveryPoliciesQuery
	*GetFailRecoveryPoliciesBody
}
type GetPushMixedStreamToCDNTaskReq struct {
	*GetPushMixedStreamToCDNTaskQuery
	*GetPushMixedStreamToCDNTaskBody
}
type GetPushSingleStreamToCDNTaskReq struct {
	*GetPushSingleStreamToCDNTaskQuery
	*GetPushSingleStreamToCDNTaskBody
}
type GetRecordTaskReq struct {
	*GetRecordTaskQuery
	*GetRecordTaskBody
}
type GetResourcePackNumReq struct {
	*GetResourcePackNumQuery
	*GetResourcePackNumBody
}
type GetWebCastTaskReq struct {
	*GetWebCastTaskQuery
	*GetWebCastTaskBody
}
type KickUserReq struct {
	*KickUserQuery
	*KickUserBody
}
type ListAppsReq struct {
	*ListAppsQuery
	*ListAppsBody
}
type ListCallDetailReq struct {
	*ListCallDetailQuery
	*ListCallDetailBody
}
type ListDetectionTaskReq struct {
	*ListDetectionTaskQuery
	*ListDetectionTaskBody
}
type ListHotMusicReq struct {
	*ListHotMusicQuery
	*ListHotMusicBody
}
type ListMusicsReq struct {
	*ListMusicsQuery
	*ListMusicsBody
}
type ListOperationDataReq struct {
	*ListOperationDataQuery
	*ListOperationDataBody
}
type ListOperationDistributionReq struct {
	*ListOperationDistributionQuery
	*ListOperationDistributionBody
}
type ListQualityReq struct {
	*ListQualityQuery
	*ListQualityBody
}
type ListQualityDistributionReq struct {
	*ListQualityDistributionQuery
	*ListQualityDistributionBody
}
type ListRealTimeOperationDataReq struct {
	*ListRealTimeOperationDataQuery
	*ListRealTimeOperationDataBody
}
type ListRealTimeQualityReq struct {
	*ListRealTimeQualityQuery
	*ListRealTimeQualityBody
}
type ListRealTimeQualityDistributionReq struct {
	*ListRealTimeQualityDistributionQuery
	*ListRealTimeQualityDistributionBody
}
type ListRelayStreamReq struct {
	*ListRelayStreamQuery
	*ListRelayStreamBody
}
type ListResourcePackagesReq struct {
	*ListResourcePackagesQuery
	*ListResourcePackagesBody
}
type ListResourcePackagesV2Req struct {
	*ListResourcePackagesV2Query
	*ListResourcePackagesV2Body
}
type ListRoomInfoReq struct {
	*ListRoomInfoQuery
	*ListRoomInfoBody
}
type ListUsagesReq struct {
	*ListUsagesQuery
	*ListUsagesBody
}
type ListUserInfoReq struct {
	*ListUserInfoQuery
	*ListUserInfoBody
}
type ModifyAppStatusReq struct {
	*ModifyAppStatusQuery
	*ModifyAppStatusBody
}
type ModifyBusinessRemarksReq struct {
	*ModifyBusinessRemarksQuery
	*ModifyBusinessRemarksBody
}
type SearchMusicsReq struct {
	*SearchMusicsQuery
	*SearchMusicsBody
}
type SendBroadcastReq struct {
	*SendBroadcastQuery
	*SendBroadcastBody
}
type SendRoomUnicastReq struct {
	*SendRoomUnicastQuery
	*SendRoomUnicastBody
}
type SendUnicastReq struct {
	*SendUnicastQuery
	*SendUnicastBody
}
type StartDetectionReq struct {
	*StartDetectionQuery
	*StartDetectionBody
}
type StartPushMixedStreamToCDNReq struct {
	*StartPushMixedStreamToCDNQuery
	*StartPushMixedStreamToCDNBody
}
type StartPushPublicStreamReq struct {
	*StartPushPublicStreamQuery
	*StartPushPublicStreamBody
}
type StartPushSingleStreamToCDNReq struct {
	*StartPushSingleStreamToCDNQuery
	*StartPushSingleStreamToCDNBody
}
type StartRecordReq struct {
	*StartRecordQuery
	*StartRecordBody
}
type StartRelayStreamReq struct {
	*StartRelayStreamQuery
	*StartRelayStreamBody
}
type StartSegmentReq struct {
	*StartSegmentQuery
	*StartSegmentBody
}
type StartSnapshotReq struct {
	*StartSnapshotQuery
	*StartSnapshotBody
}
type StartWBRecordReq struct {
	*StartWBRecordQuery
	*StartWBRecordBody
}
type StartWebcastReq struct {
	*StartWebcastQuery
	*StartWebcastBody
}
type StopDetectionReq struct {
	*StopDetectionQuery
	*StopDetectionBody
}
type StopPushPublicStreamReq struct {
	*StopPushPublicStreamQuery
	*StopPushPublicStreamBody
}
type StopPushStreamToCDNReq struct {
	*StopPushStreamToCDNQuery
	*StopPushStreamToCDNBody
}
type StopRecordReq struct {
	*StopRecordQuery
	*StopRecordBody
}
type StopRelayStreamReq struct {
	*StopRelayStreamQuery
	*StopRelayStreamBody
}
type StopSegmentReq struct {
	*StopSegmentQuery
	*StopSegmentBody
}
type StopSnapshotReq struct {
	*StopSnapshotQuery
	*StopSnapshotBody
}
type StopWBRecordReq struct {
	*StopWBRecordQuery
	*StopWBRecordBody
}
type StopWebcastReq struct {
	*StopWebcastQuery
	*StopWebcastBody
}
type UnbanUserStreamReq struct {
	*UnbanUserStreamQuery
	*UnbanUserStreamBody
}
type UpdateBanRoomUserRuleReq struct {
	*UpdateBanRoomUserRuleQuery
	*UpdateBanRoomUserRuleBody
}
type UpdateCallbackReq struct {
	*UpdateCallbackQuery
	*UpdateCallbackBody
}
type UpdateFailRecoveryPolicyReq struct {
	*UpdateFailRecoveryPolicyQuery
	*UpdateFailRecoveryPolicyBody
}
type UpdatePublicStreamParamReq struct {
	*UpdatePublicStreamParamQuery
	*UpdatePublicStreamParamBody
}
type UpdatePushMixedStreamToCDNReq struct {
	*UpdatePushMixedStreamToCDNQuery
	*UpdatePushMixedStreamToCDNBody
}
type UpdateRecordReq struct {
	*UpdateRecordQuery
	*UpdateRecordBody
}
type UpdateRelayStreamReq struct {
	*UpdateRelayStreamQuery
	*UpdateRelayStreamBody
}
type UpdateSegmentReq struct {
	*UpdateSegmentQuery
	*UpdateSegmentBody
}
type UpdateSnapshotReq struct {
	*UpdateSnapshotQuery
	*UpdateSnapshotBody
}
type UpdateVendorPolicyReq struct {
	*UpdateVendorPolicyQuery
	*UpdateVendorPolicyBody
}
type WbTranscodeCreateReq struct {
	*WbTranscodeCreateQuery
	*WbTranscodeCreateBody
}
type WbTranscodeGetReq struct {
	*WbTranscodeGetQuery
	*WbTranscodeGetBody
}
type WbTranscodeQueryReq struct {
	*WbTranscodeQueryQuery
	*WbTranscodeQueryBody
}
