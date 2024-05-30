package rtc_v20230801

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

type GetRoomOnlineUsersQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 指定房间 ID
	RoomID string `json:"RoomId" query:"RoomId"`
}

type GetRoomOnlineUsersRes struct {

	// REQUIRED
	ResponseMetadata GetRoomOnlineUsersResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetRoomOnlineUsersResResult          `json:"Result,omitempty"`
}

type GetRoomOnlineUsersResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error GetRoomOnlineUsersResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type GetRoomOnlineUsersResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

type GetRoomOnlineUsersResResult struct {

	// 不可见用户列表
	InvisibleUserList []*string `json:"InvisibleUserList,omitempty"`

	// 查询的房间是否存在。
	// * true：存在。
	// * false：不存在。
	// 当 RoomExists 的值为 false 时，不会返回其他字段。
	RoomExists *bool `json:"RoomExists,omitempty"`

	// 查询到的不可见用户总数。不可见用户最多返回 10000 名。
	TotalInvisibleUser *int32 `json:"TotalInvisibleUser,omitempty"`

	// 查询到的用户总数
	TotalUser *int32 `json:"TotalUser,omitempty"`

	// 可见用户列表
	VisibleUserList []*string `json:"VisibleUserList,omitempty"`
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
type GetRoomOnlineUsers struct{}
type GetRoomOnlineUsersBody struct{}
type GetRoomOnlineUsersReq struct {
	*GetRoomOnlineUsersQuery
	*GetRoomOnlineUsersBody
}
