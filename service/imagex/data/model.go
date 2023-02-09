package data

// common

type TSFloatItem struct {
	TimeStamp string  `json:"TimeStamp"`
	Value     float64 `json:"Value"`
}

type TSIntItem struct {
	TimeStamp string `json:"TimeStamp"`
	Value     int64  `json:"Value"`
}

type DataFloatItem struct {
	Data []TSFloatItem `json:"Data"`
}

type DataIntItem struct {
	Data []TSIntItem `json:"Data"`
}

type ServiceDomainDataItem struct {
	ServiceId  string        `json:"ServiceId,omitempty"`
	DomainName string        `json:"DomainName,omitempty"`
	Data       []TSFloatItem `json:"Data"`
}

// query: Query参数; json: JSON参数

// DescribeImageXDomainTrafficData
// Method: GET

type DescribeImageXDomainTrafficDataReq struct {
	ServiceIds  string `query:"ServiceIds"`  // 传入多个用英文逗号分隔
	DomainNames string `query:"DomainNames"` // 传入多个用英文逗号分隔
	GroupBy     string `query:"GroupBy"`     // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持ServiceId,DomainName
	StartTime   string `query:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string `query:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval    string `query:"Interval"`    // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXDomainTrafficDataResp struct {
	TrafficData []ServiceDomainDataItem `json:"TrafficData"`
}

// DescribeImageXDomainBandwidthData
// Method: GET

type DescribeImageXDomainBandwidthDataReq struct {
	ServiceIds  string `query:"ServiceIds"`  // 传入多个用英文逗号分隔
	DomainNames string `query:"DomainNames"` // 传入多个用英文逗号分隔
	GroupBy     string `query:"GroupBy"`     // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持ServiceId,DomainName
	StartTime   string `query:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string `query:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval    string `query:"Interval"`    // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXDomainBandwidthDataResp struct {
	BpsData []ServiceDomainDataItem `json:"BpsData"`
}

// DescribeImageXBucketUsage
// Method: GET

type DescribeImageXBucketUsageReq struct {
	ServiceIds  string `query:"ServiceIds"`  // 传入多个用英文逗号分隔
	BucketNames string `query:"BucketNames"` // 传入多个用英文逗号分隔
	GroupBy     string `query:"GroupBy"`     // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持ServiceId,BucketName
	StartTime   string `query:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string `query:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
}

type DescribeImageXBucketUsageResp struct {
	StorageData []StorageDataItem `json:"StorageData"`
}

type StorageDataItem struct {
	ServiceId  string        `json:"ServiceId,omitempty"`
	BucketName string        `json:"BucketName,omitempty"`
	Data       []TSFloatItem `json:"Data"`
}

// DescribeImageXRequestCntUsage
// Method: GET

type DescribeImageXRequestCntUsageReq struct {
	ServiceIds string `query:"ServiceIds"` // 传入多个用英文逗号分隔
	AdvFeats   string `query:"AdvFeats"`   // 传入多个用英文逗号分隔
	GroupBy    string `query:"GroupBy"`    // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持AdvFeat
	StartTime  string `query:"StartTime"`  // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime    string `query:"EndTime"`    // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval   string `query:"Interval"`   // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXRequestCntUsageResp struct {
	RequestCntData []RequestCntDataItem `json:"RequestCntData"`
}

type RequestCntDataItem struct {
	AdvFeat string        `json:"AdvFeat,omitempty"`
	Data    []TSFloatItem `json:"Data"`
}

// DescribeImageXBaseOpUsage
// Method: GET

type DescribeImageXBaseOpUsageReq struct {
	ServiceIds string `query:"ServiceIds"` // 传入多个用英文逗号分隔
	StartTime  string `query:"StartTime"`  // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime    string `query:"EndTime"`    // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval   string `query:"Interval"`   // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXBaseOpUsageResp struct {
	BaseOpData []DataFloatItem `json:"BaseOpData"`
}

// DescribeImageXCompressUsage
// Method: GET

type DescribeImageXCompressUsageReq struct {
	ServiceIds string `query:"ServiceIds"` // 传入多个用英文逗号分隔
	StartTime  string `query:"StartTime"`  // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime    string `query:"EndTime"`    // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval   string `query:"Interval"`   // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXCompressUsageResp struct {
	CompressData []CompressDataItem `json:"CompressData"`
}

type CompressDataItem struct {
	InSize  []TSFloatItem `json:"InSize"`
	OutSize []TSFloatItem `json:"OutSize"`
}

// DescribeImageXEdgeRequest
// Method: GET

type DescribeImageXEdgeRequestReq struct {
	ServiceIds  string `query:"ServiceIds"`  // 传入多个用英文逗号分隔
	DomainNames string `query:"DomainNames"` // 传入多个用英文逗号分隔
	DataTypes   string `query:"DataTypes"`   // 状态码过滤。传入多个用英文逗号分隔。支持2xx,3xx,4xx,5xx,req_cnt
	GroupBy     string `query:"GroupBy"`     // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持DomainName,DataType
	StartTime   string `query:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string `query:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval    string `query:"Interval"`    // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXEdgeRequestResp struct {
	RequestCnt []EdgeRequestCntItem `json:"RequestCnt"`
}

type EdgeRequestCntItem struct {
	DataType   string      `json:"DataType,omitempty"`
	DomainName string      `json:"DomainName,omitempty"`
	Data       []TSIntItem `json:"Data"`
}

// DescribeImageXMirrorRequestTraffic
// Method: POST

type DescribeImageXMirrorRequestTrafficReq struct {
	ServiceIds  []string `json:"ServiceIds"`  // 服务id过滤
	DomainNames []string `json:"DomainNames"` // 域名过滤
	StartTime   string   `json:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string   `json:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval    string   `json:"Interval"`    // 时间粒度。支持5m,1h,1d。不传默认查询时间段全部数据
	Granularity string   `json:"Granularity"` // 时间粒度。支持5m,1h,1d。不传默认查询时间段全部数据
}

type DescribeImageXMirrorRequestTrafficResp struct {
	Data []TSIntItem `json:"Data"`
}

// DescribeImageXMirrorRequestBandwidth
// Method: POST

type DescribeImageXMirrorRequestBandwidthReq struct {
	ServiceIds  []string `json:"ServiceIds"`  // 服务id过滤
	DomainNames []string `json:"DomainNames"` // 域名过滤
	StartTime   string   `json:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string   `json:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	Granularity string   `json:"Granularity"` // 时间粒度。支持5m,1h,1d。不传默认查询时间段全部数据
	Interval    string   `json:"Interval"`    // 时间粒度。支持5m,1h,1d。不传默认查询时间段全部数据
}

type DescribeImageXMirrorRequestBandwidthResp struct {
	Data []TSFloatItem `json:"Data"`
}

// DescribeImageXMirrorRequestHttpCodeByTime
// Method: POST

type DescribeImageXMirrorRequestHttpCodeByTimeReq struct {
	ServiceIds    []string `json:"ServiceIds"`    // 服务id过滤
	DomainNames   []string `json:"DomainNames"`   // 域名过滤
	AggregateCode string   `json:"AggregateCode"` // 状态码是否聚合。支持true,false。默认为false
	StartTime     string   `json:"StartTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime       string   `json:"EndTime"`       // YYYY-MM-DDThh:mm:ss±hh:mm
	Granularity   string   `json:"Granularity"`   // 时间粒度。支持5m,1h,1d。不传默认查询时间段全部数据
}

type DescribeImageXMirrorRequestHttpCodeByTimeResp struct {
	Data []MirrorHttpCodeDataItem `json:"Data"`
}

type MirrorHttpCodeDataItem struct {
	HttpCode string      `json:"HttpCode"`
	Data     []TSIntItem `json:"Data"`
}

// DescribeImageXHitRateTrafficData
// Method: GET

type DescribeImageXHitRateTrafficDataReq struct {
	ServiceIds  string `query:"ServiceIds"`  // 传入多个用英文逗号分隔
	DomainNames string `query:"DomainNames"` // 传入多个用英文逗号分隔
	GroupBy     string `query:"GroupBy"`     // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持ServiceId,DomainName
	StartTime   string `query:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string `query:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval    string `query:"Interval"`    // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXHitRateTrafficDataResp struct {
	HitRateData []ServiceDomainDataItem `json:"HitRateData"`
}

// DescribeImageXHitRateRequestData
// Method: GET

type DescribeImageXHitRateRequestDataReq struct {
	ServiceIds  string `query:"ServiceIds"`  // 传入多个用英文逗号分隔
	DomainNames string `query:"DomainNames"` // 传入多个用英文逗号分隔
	GroupBy     string `query:"GroupBy"`     // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持ServiceId,DomainName
	StartTime   string `query:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string `query:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval    string `query:"Interval"`    // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXHitRateRequestDataResp struct {
	HitRateData []ServiceDomainDataItem `json:"HitRateData"`
}

// DescribeImageXCDNTopRequestData
// Method: GET

type DescribeImageXCDNTopRequestDataReq struct {
	ServiceIds  string `query:"ServiceIds"`  // 传入多个用英文逗号分隔
	DomainNames string `query:"DomainNames"` // 传入多个用英文逗号分隔
	StartTime   string `query:"StartTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime     string `query:"EndTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	KeyType     string `query:"KeyType"`     // 需要Top的对象
	ValueType   string `query:"ValueType"`   // 单选Traffic/RequestCnt，代表按流量/请求次数排序 (KeyType=Domain时，只能为Traffic)
	IPVersion   string `query:"IPVersion"`   // 单选IPv4/IPv6，不传为不限制； KeyType=Domain时无效
	Country     string `query:"Country"`
	Limit       string `query:"Limit"`  // 分页Limit，默认为0，即返回所有
	Offset      string `query:"Offset"` // 分页Offset，默认为0
}

type DescribeImageXCDNTopRequestDataResp struct {
	Count      int            `json:"Count"`
	TopValue   []TopValueItem `json:"TopValue"`
	TotalValue float64        `json:"TotalValue"`
}

type TopValueItem struct {
	Key   string  `json:"Key"`
	Value float64 `json:"Value"`
}

// DescribeImageXSummary
// Method: GET

type DescribeImageXSummaryReq struct {
	ServiceIds string `query:"ServiceIds"` // 传入多个用英文逗号分隔
	Timestamp  string `query:"Timestamp"`
}

type DescribeImageXSummaryResp struct {
	StorageData      SummaryFloatItem `json:"StorageData"`
	CdnBandwidthData SummaryFloatItem `json:"CdnBandwidthData"`
	CdnTrafficData   SummaryIntItem   `json:"CdnTrafficData"`
	RequestCntData   SummaryIntItem   `json:"RequestCntData"`
	BaseOpData       SummaryIntItem   `json:"BaseOpData"`
}

type SummaryFloatItem struct {
	Value float64 `json:"Value"`
}

type SummaryIntItem struct {
	Value int64 `json:"Value"`
}

// DescribeImageXMirrorRequestHttpCodeOverview
// Method: POST

type DescribeImageXMirrorRequestHttpCodeOverviewReq struct {
	ServiceIds    []string `json:"ServiceIds"`
	DomainNames   []string `json:"DomainNames"`
	StartTime     string   `json:"StartTime"` // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime       string   `json:"EndTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
	AggregateCode string   `json:"AggregateCode"`
	Granularity   string   `json:"Granularity"`
}

type DescribeImageXMirrorRequestHttpCodeOverviewResp struct {
	Data []MirrorHttpCodeRespItem `json:"Data"`
}

type MirrorHttpCodeRespItem struct {
	Domain  string                      `json:"Domain"`
	Count   int64                       `json:"Count"`
	Details []MirrorHttpCodeDetailsItem `json:"Details"`
}

type MirrorHttpCodeDetailsItem struct {
	HttpCode string `json:"HttpCode"`
	Value    int64  `json:"Value"`
}
