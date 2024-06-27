package rtc_v20230720

type BatchSendRoomUnicastBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 消息类型。
	// * true：二进制消息。false：文本消息。
	Binary bool `json:"Binary"`

	// REQUIRED; 业务服务端的唯一标识。 命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。 在一个 AppID 下，不能和真实用户用于实时消息通信的 user_ID 重复。 建议使用固定的 ID 的发送消息。
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// BatchSendRoomUnicastResResult - 返回值 Result 仅在请求成功时返回消息 success ，表示服务端成功接收到消息，失败时为空。
type BatchSendRoomUnicastResResult struct {

	// REQUIRED; 请求成功时返回消息Success，表示服务端成功接收到消息，失败时返回具体错误信息。
	Message string `json:"Message"`
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

type SendBroadcastBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 消息类型。
	// * true：二进制消息。false：文本消息。
	Binary bool `json:"Binary"`

	// REQUIRED; 业务服务端的唯一标识； 命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。 在一个 AppID 下，不能和真实用户用于实时消息通信的 user_ID 重复； 建议使用固定的 ID 的发送消息。
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type SendBroadcastResResult struct {

	// REQUIRED; 请求成功时返回消息Success，表示服务端成功接收到消息，失败时返回具体错误信息。
	Message string `json:"Message"`
}

type SendRoomUnicastBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 消息类型。
	// * true：二进制消息。false：文本消息。
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// SendRoomUnicastResResult - 仅在请求成功时返回消息 "success"，表示服务端成功接收到消息，失败时为空。
type SendRoomUnicastResResult struct {

	// REQUIRED; 请求成功时返回消息Success，表示服务端成功接收到消息，失败时返回具体错误信息。
	Message string `json:"Message"`
}

type SendUnicastBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; * 字段为 true，发送二进制消息；
	// * 字段为 false，发送文本消息。
	Binary bool `json:"Binary"`

	// REQUIRED; 业务服务端的唯一标识。 命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。 在一个 AppID 下，不能和真实用户用于实时消息通信的 user_ID 重复； 建议使用固定的 ID 的发送消息。
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// SendUnicastResResult - 仅在请求成功时返回消息 "success"，表示服务端成功接收到消息，失败时为空。
type SendUnicastResResult struct {

	// REQUIRED; 请求成功时返回消息Success，表示服务端成功接收到消息，失败时返回具体错误信息。
	Message string `json:"Message"`
}
type BatchSendRoomUnicast struct{}
type BatchSendRoomUnicastQuery struct{}
type SendBroadcast struct{}
type SendBroadcastQuery struct{}
type SendRoomUnicast struct{}
type SendRoomUnicastQuery struct{}
type SendUnicast struct{}
type SendUnicastQuery struct{}
type BatchSendRoomUnicastReq struct {
	*BatchSendRoomUnicastQuery
	*BatchSendRoomUnicastBody
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
