package wtn_v20231101

type ListRealTimePublicStreamInfoBody struct {

	// REQUIRED
	AppID string `json:"AppId"`

	// REQUIRED
	EndTime string `json:"EndTime"`

	// REQUIRED
	Indicator []string `json:"Indicator"`

	// REQUIRED
	StartTime string  `json:"StartTime"`
	StreamID  *string `json:"StreamId,omitempty"`
}

type ListRealTimePublicStreamInfoRes struct {

	// REQUIRED
	ResponseMetadata ListRealTimePublicStreamInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRealTimePublicStreamInfoResResult          `json:"Result,omitempty"`
}

type ListRealTimePublicStreamInfoResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	Error *ListRealTimePublicStreamInfoResResponseMetadataError `json:"Error,omitempty"`
}

type ListRealTimePublicStreamInfoResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message     string  `json:"Message"`
	CodeN       *int32  `json:"CodeN,omitempty"`
	Description *string `json:"Description,omitempty"`
}

type ListRealTimePublicStreamInfoResResult struct {
	Indicators []*ListRealTimePublicStreamInfoResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListRealTimePublicStreamInfoResResultIndicatorsItemOverview struct {

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

type ListRealTimePublicStreamInfoResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListRealTimePublicStreamInfoResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                                      `json:"Name,omitempty"`
	Overview *ListRealTimePublicStreamInfoResResultIndicatorsItemOverview `json:"overview,omitempty"`

	// 指标单位
	Unit *string `json:"Unit,omitempty"`
}

type ListRealTimePublicStreamInfoResResultIndicatorsPropertiesItemsItem struct {
	StreamID *string `json:"StreamId,omitempty"`

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}
