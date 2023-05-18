package imagex

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

type DomainDataItem struct {
	DomainName string        `json:"DomainName,omitempty"`
	Data       []TSFloatItem `json:"Data"`
}

type ExtraDim struct {
	Dim  string   `json:"Dim"`
	Vals []string `json:"Vals"`
}

type CurveDataIntItem struct {
	Value     int64  `json:"Value"`
	Count     int64  `json:"Count"`
	Timestamp string `json:"Timestamp"`
}

type CurveDataFloatItem struct {
	Value     float64 `json:"Value"`
	Count     int64   `json:"Count"`
	Timestamp string  `json:"Timestamp"`
}

type CurveIntItem struct {
	Type  string             `json:"Type"`
	Count int64              `json:"Count"`
	Data  []CurveDataIntItem `json:"Data"`
}

type CurveFloatItem struct {
	Type  string               `json:"Type"`
	Count int64                `json:"Count"`
	Data  []CurveDataFloatItem `json:"Data"`
}

type ErrorCodeAllDetailsItem struct {
	ErrorCode     string `json:"ErrorCode"`
	ErrorCodeDesc string `json:"ErrorCodeDesc"`
	Value         int64  `json:"Value"`
}

type ErrorCodeAllItem struct {
	ErrorCode     string                    `json:"ErrorCode,omitempty"`
	ErrorCodeDesc string                    `json:"ErrorCodeDesc,omitempty"`
	Isp           string                    `json:"Isp,omitempty"`
	Domain        string                    `json:"Domain,omitempty"`
	Region        string                    `json:"Region,omitempty"`
	RegionType    string                    `json:"RegionType,omitempty"`
	Value         int64                     `json:"Value"`
	Details       []ErrorCodeAllDetailsItem `json:"Details,omitempty"`
}

type ErrorCodeItem struct {
	ErrorCode string         `json:"ErrorCode"`
	Count     int64          `json:"Count"`
	Data      []CurveIntItem `json:"Data"`
}

type AllItem struct {
	Domain     string  `json:"Domain,omitempty"`
	Isp        string  `json:"Isp,omitempty"`
	Region     string  `json:"Region,omitempty"`
	RegionType string  `json:"RegionType,omitempty"`
	Value      float64 `json:"Value"`
	Count      int64   `json:"Count"`
}

type ClientTopItem struct {
	Url   string `json:"Url"`
	Value int64  `json:"Value"`
	Count int64  `json:"Count"`
}

// query: Query参数; json: JSON参数

// DescribeImageXDomainTrafficData
// Method: GET

type DescribeImageXDomainTrafficDataReq struct {
	ServiceIds    string `query:"ServiceIds"`    // 传入多个用英文逗号分隔
	DomainNames   string `query:"DomainNames"`   // 传入多个用英文逗号分隔
	BillingRegion string `query:"BillingRegion"` // 过滤计费区域
	GroupBy       string `query:"GroupBy"`       // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持ServiceId,DomainName
	StartTime     string `query:"StartTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime       string `query:"EndTime"`       // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval      string `query:"Interval"`      // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXDomainTrafficDataResp struct {
	TrafficData []ServiceDomainDataItem `json:"TrafficData"`
}

// DescribeImageXDomainBandwidthData
// Method: GET

type DescribeImageXDomainBandwidthDataReq struct {
	ServiceIds    string `query:"ServiceIds"`    // 传入多个用英文逗号分隔
	DomainNames   string `query:"DomainNames"`   // 传入多个用英文逗号分隔
	BillingRegion string `query:"BillingRegion"` // 过滤计费区域
	GroupBy       string `query:"GroupBy"`       // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持ServiceId,DomainName
	StartTime     string `query:"StartTime"`     // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime       string `query:"EndTime"`       // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval      string `query:"Interval"`      // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
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
	ServiceIds   string `query:"ServiceIds"`   // 传入多个用英文逗号分隔
	DomainNames  string `query:"DomainNames"`  // 传入多个用英文逗号分隔
	Regions      string `query:"Regions"`      // cdn区域。支持参数：中国大陆、亚太一区、亚太二区、亚太三区、欧洲区、北美区、南美区、中东区。传入多个用英文逗号分割。不传表示不过滤。
	UserCountry  string `query:"UserCountry"`  // 客户端国家。支持参数：中国、海外、美国等。传入多个用英文逗号分割。不传表示不过滤。
	UserProvince string `query:"UserProvince"` // 客户端省份。传入多个用英文逗号分割。不传表示不过滤。
	Protocols    string `query:"Protocols"`    // 过滤网络协议。支持参数：HTTP、HTTPS。传入多个用英文逗号分割。不传为不过滤。
	Isp          string `query:"Isp"`          // 过滤运营商。传入多个用英文逗号分割。不传为不过滤。
	DataTypes    string `query:"DataTypes"`    // 状态码过滤。传入多个用英文逗号分隔。支持2xx,3xx,4xx,5xx,req_cnt
	GroupBy      string `query:"GroupBy"`      // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持DomainName,DataType
	StartTime    string `query:"StartTime"`    // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime      string `query:"EndTime"`      // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval     string `query:"Interval"`     // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
	DetailedCode string `query:"DetailedCode"` // 是否拆分状态码。不传默认为false，表示不拆分。
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

// DescribeImageXEdgeRequestBandwidth
// Method: GET

type DescribeImageXEdgeRequestBandwidthReq struct {
	ServiceIds   string `query:"ServiceIds"`   // 传入多个用英文逗号分隔
	DomainNames  string `query:"DomainNames"`  // 传入多个用英文逗号分隔
	Regions      string `query:"Regions"`      // cdn区域。支持参数：中国大陆、亚太一区、亚太二区、亚太三区、欧洲区、北美区、南美区、中东区。传入多个用英文逗号分割。不传表示不过滤。
	UserCountry  string `query:"UserCountry"`  // 客户端国家。支持参数：中国、海外、美国等。传入多个用英文逗号分割。不传表示不过滤。
	UserProvince string `query:"UserProvince"` // 客户端省份。传入多个用英文逗号分割。不传表示不过滤。
	Protocols    string `query:"Protocols"`    // 过滤网络协议。支持参数：HTTP、HTTPS。传入多个用英文逗号分割。不传为不过滤。
	Isp          string `query:"Isp"`          // 过滤运营商。传入多个用英文逗号分割。不传为不过滤。
	GroupBy      string `query:"GroupBy"`      // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持DomainName
	StartTime    string `query:"StartTime"`    // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime      string `query:"EndTime"`      // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval     string `query:"Interval"`     // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXEdgeRequestBandwidthResp struct {
	BpsData []DomainDataItem `json:"BpsData"`
}

// DescribeImageXEdgeRequestTraffic
// Method: GET

type DescribeImageXEdgeRequestTrafficReq struct {
	ServiceIds   string `query:"ServiceIds"`   // 传入多个用英文逗号分隔
	DomainNames  string `query:"DomainNames"`  // 传入多个用英文逗号分隔
	Regions      string `query:"Regions"`      // cdn区域。支持参数：中国大陆、亚太一区、亚太二区、亚太三区、欧洲区、北美区、南美区、中东区。传入多个用英文逗号分割。不传表示不过滤。
	UserCountry  string `query:"UserCountry"`  // 客户端国家。支持参数：中国、海外、美国等。传入多个用英文逗号分割。不传表示不过滤。
	UserProvince string `query:"UserProvince"` // 客户端省份。传入多个用英文逗号分割。不传表示不过滤。
	Protocols    string `query:"Protocols"`    // 过滤网络协议。支持参数：HTTP、HTTPS。传入多个用英文逗号分割。不传为不过滤。
	Isp          string `query:"Isp"`          // 过滤运营商。传入多个用英文逗号分割。不传为不过滤。
	GroupBy      string `query:"GroupBy"`      // 维度拆分的维度值。不传表示不拆分维度。传入多个用英文逗号分隔。支持DomainName
	StartTime    string `query:"StartTime"`    // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime      string `query:"EndTime"`      // YYYY-MM-DDThh:mm:ss±hh:mm
	Interval     string `query:"Interval"`     // 时间粒度，单位秒。支持300,3600,86400。不传默认查询时间段全部数据
}

type DescribeImageXEdgeRequestTrafficResp struct {
	TrafficData []DomainDataItem `json:"TrafficData"`
}

// DescribeImageXEdgeRequestRegions
// Method: GET

type DescribeImageXEdgeRequestRegionsReq struct {
	StartTime string `query:"StartTime"` // YYYY-MM-DDThh:mm:ss±hh:mm
	EndTime   string `query:"EndTime"`   // YYYY-MM-DDThh:mm:ss±hh:mm
}

type DescribeImageXEdgeRequestRegionsResp struct {
	UserProvince []string `json:"UserProvince"`
	UserCountry  []string `json:"UserCountry"`
}

// DescribeImageXServiceQuality
// Method: GET

type DescribeImageXServiceQualityReq struct {
	Region string `query:"Region"`
}

type DescribeImageXServiceQualityResp struct {
	FailureRate float64            `json:"FailureRate"`
	UploadData  UploadOverviewResp `json:"UploadData"`
	CdnData     CdnOverviewResp    `json:"CdnData"`
	ClientData  ClientOverviewResp `json:"ClientData"`
}

type UploadOverviewResp struct {
	SuccessRate float64 `json:"SuccessRate"`
	UploadCount int     `json:"UploadCount"`
	AvgSize     float64 `json:"AvgSize"`
	AvgDuration float64 `json:"AvgDuration"`
	TotalCount  int     `json:"TotalCount"`
}

type CdnOverviewResp struct {
	SuccessRate float64 `json:"SuccessRate"`
	AvgDuration float64 `json:"AvgDuration"`
	TotalCount  int     `json:"TotalCount"`
}

type ClientOverviewResp struct {
	FailureRate       float64 `json:"FailureRate"`
	SuccessRate       float64 `json:"SuccessRate"`
	AvgDecodeDuration float64 `json:"AvgDecodeDuration"`
	AvgQueueDuration  float64 `json:"AvgQueueDuration"`
	AvgLoadDuration   float64 `json:"AvgLoadDuration"`
	TotalCount        int     `json:"TotalCount"`
}

// GetImageXQueryApps
// Method: GET

type GetImageXQueryAppsReq struct {
	Source string `query:"Source"`
}

type GetImageXQueryAppsResp struct {
	Apps []App `json:"Apps"`
}

type App struct {
	AppId       string `json:"AppId"`
	AppName     string `json:"AppName"`
	AppRegion   string `json:"AppRegion"`
	PackageName string `json:"PackageName"`
	BundleId    string `json:"BundleId"`
}

// GetImageXQueryRegions
// Method: GET

type GetImageXQueryRegionsReq struct {
	Source string `query:"Source"`
	Appid  string `query:"Appid"`
	OS     string `query:"OS"`
}

type GetImageXQueryRegionsResp struct {
	Province []string `json:"Province"`
	Country  []string `json:"Country"`
}

// GetImageXQueryDims
// Method: GET

type GetImageXQueryDimsReq struct {
	Source string `query:"Source"`
	Appid  string `query:"Appid"`
	OS     string `query:"OS"`
}

type GetImageXQueryDimsResp struct {
	Dim []string `json:"Dim"`
}

// GetImageXQueryVals
// Method: GET

type GetImageXQueryValsReq struct {
	Dim    string `query:"Dim"`
	Source string `query:"Source"`
	Appid  string `query:"Appid"`
	OS     string `query:"OS"`
}

type GetImageXQueryValsResp struct {
	DimVal []string `json:"DimVal"`
}

// DescribeImageXUploadSuccessRateByTime
// Method: POST

type DescribeImageXUploadSuccessRateByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	UploadType  int        `json:"UploadType"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXUploadSuccessRateByTimeResp struct {
	SuccessRateData []CurveFloatItem `json:"SuccessRateData"`
}

// DescribeImageXUploadErrorCodeAll
// Method: POST

type DescribeImageXUploadErrorCodeAllReq struct {
	Appid      string     `json:"Appid"`
	AppVer     []string   `json:"AppVer"`
	OS         string     `json:"OS"`
	SdkVer     []string   `json:"SdkVer"`
	Country    string     `json:"Country"`
	Province   string     `json:"Province"`
	Isp        []string   `json:"Isp"`
	ExtraDims  []ExtraDim `json:"ExtraDims"`
	UploadType int        `json:"UploadType"`
	StartTime  string     `json:"StartTime"`
	EndTime    string     `json:"EndTime"`
	GroupBy    string     `json:"GroupBy"`
	OrderBy    string     `json:"OrderBy"`
	OrderAsc   bool       `json:"OrderAsc"`
}

type DescribeImageXUploadErrorCodeAllResp struct {
	ErrorCodeData []ErrorCodeAllItem `json:"ErrorCodeData"`
}

// DescribeImageXUploadErrorCodeByTime
// Method: POST

type DescribeImageXUploadErrorCodeByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	UploadType  int        `json:"UploadType"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXUploadErrorCodeByTimeResp struct {
	ErrorCodeData []ErrorCodeItem `json:"ErrorCodeData"`
}

// DescribeImageXUploadCountByTime
// Method: POST

type DescribeImageXUploadCountByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	UploadType  int        `json:"UploadType"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXUploadCountByTimeResp struct {
	UploadCountData []CurveIntItem `json:"UploadCountData"`
}

// DescribeImageXUploadFileSize
// Method: POST

type DescribeImageXUploadFileSizeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	UploadType  int        `json:"UploadType"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXUploadFileSizeResp struct {
	FSizeData []CurveFloatItem `json:"FSizeData"`
}

// DescribeImageXUploadSpeed
// Method: POST

type DescribeImageXUploadSpeedReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	UploadType  int        `json:"UploadType"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXUploadSpeedResp struct {
	SpeedData []CurveFloatItem `json:"SpeedData"`
}

// DescribeImageXUploadDuration
// Method: POST

type DescribeImageXUploadDurationReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	UploadType  int        `json:"UploadType"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXUploadDurationResp struct {
	DurationData []CurveFloatItem `json:"DurationData"`
}

// DescribeImageXUploadSegmentSpeedByTime
// Method: POST

type DescribeImageXUploadSegmentSpeedByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	UploadType  int        `json:"UploadType"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXUploadSegmentSpeedByTimeResp struct {
	SpeedData []CurveFloatItem `json:"SpeedData"`
}

// DescribeImageXCdnSuccessRateByTime
// Method: POST

type DescribeImageXCdnSuccessRateByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXCdnSuccessRateByTimeResp struct {
	SuccessRateData []CurveFloatItem `json:"SuccessRateData"`
}

// DescribeImageXCdnSuccessRateAll
// Method: POST

type DescribeImageXCdnSuccessRateAllReq struct {
	Appid     string     `json:"Appid"`
	AppVer    []string   `json:"AppVer"`
	OS        string     `json:"OS"`
	SdkVer    []string   `json:"SdkVer"`
	Country   string     `json:"Country"`
	Province  string     `json:"Province"`
	Isp       []string   `json:"Isp"`
	Domain    []string   `json:"Domain"`
	ImageType []string   `json:"ImageType"`
	ExtraDims []ExtraDim `json:"ExtraDims"`
	StartTime string     `json:"StartTime"`
	EndTime   string     `json:"EndTime"`
	GroupBy   string     `json:"GroupBy"`
	OrderBy   string     `json:"OrderBy"`
	OrderAsc  bool       `json:"OrderAsc"`
}

type DescribeImageXCdnSuccessRateAllResp struct {
	SuccessRateData []AllItem `json:"SuccessRateData"`
}

// DescribeImageXCdnErrorCodeByTime
// Method: POST

type DescribeImageXCdnErrorCodeByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXCdnErrorCodeByTimeResp struct {
	ErrorCodeData []ErrorCodeItem `json:"ErrorCodeData"`
}

// DescribeImageXCdnErrorCodeAll
// Method: POST

type DescribeImageXCdnErrorCodeAllReq struct {
	Appid     string     `json:"Appid"`
	AppVer    []string   `json:"AppVer"`
	OS        string     `json:"OS"`
	SdkVer    []string   `json:"SdkVer"`
	Country   string     `json:"Country"`
	Province  string     `json:"Province"`
	Isp       []string   `json:"Isp"`
	Domain    []string   `json:"Domain"`
	ImageType []string   `json:"ImageType"`
	ExtraDims []ExtraDim `json:"ExtraDims"`
	StartTime string     `json:"StartTime"`
	EndTime   string     `json:"EndTime"`
	GroupBy   string     `json:"GroupBy"`
	OrderBy   string     `json:"OrderBy"`
	OrderAsc  bool       `json:"OrderAsc"`
}

type DescribeImageXCdnErrorCodeAllResp struct {
	ErrorCodeData []ErrorCodeAllItem `json:"ErrorCodeData"`
}

// DescribeImageXCdnDurationDetailByTime
// Method: POST

type DescribeImageXCdnDurationDetailByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	Phase       string     `json:"Phase"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXCdnDurationDetailByTimeResp struct {
	DurationData []CurveFloatItem `json:"DurationData"`
}

// DescribeImageXCdnDurationAll
// Method: POST

type DescribeImageXCdnDurationAllReq struct {
	Appid     string     `json:"Appid"`
	AppVer    []string   `json:"AppVer"`
	OS        string     `json:"OS"`
	SdkVer    []string   `json:"SdkVer"`
	Country   string     `json:"Country"`
	Province  string     `json:"Province"`
	Isp       []string   `json:"Isp"`
	Domain    []string   `json:"Domain"`
	ImageType []string   `json:"ImageType"`
	ExtraDims []ExtraDim `json:"ExtraDims"`
	StartTime string     `json:"StartTime"`
	EndTime   string     `json:"EndTime"`
	GroupBy   string     `json:"GroupBy"`
	OrderBy   string     `json:"OrderBy"`
	OrderAsc  bool       `json:"OrderAsc"`
}

type DescribeImageXCdnDurationAllResp struct {
	DurationData []AllItem `json:"DurationData"`
}

// DescribeImageXCdnReuseRateByTime
// Method: POST

type DescribeImageXCdnReuseRateByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXCdnReuseRateByTimeResp struct {
	ReuseRateData []CurveFloatItem `json:"ReuseRateData"`
}

// DescribeImageXCdnReuseRateAll
// Method: POST

type DescribeImageXCdnReuseRateAllReq struct {
	Appid     string     `json:"Appid"`
	AppVer    []string   `json:"AppVer"`
	OS        string     `json:"OS"`
	SdkVer    []string   `json:"SdkVer"`
	Country   string     `json:"Country"`
	Province  string     `json:"Province"`
	Isp       []string   `json:"Isp"`
	Domain    []string   `json:"Domain"`
	ImageType []string   `json:"ImageType"`
	ExtraDims []ExtraDim `json:"ExtraDims"`
	StartTime string     `json:"StartTime"`
	EndTime   string     `json:"EndTime"`
	GroupBy   string     `json:"GroupBy"`
	OrderBy   string     `json:"OrderBy"`
	OrderAsc  bool       `json:"OrderAsc"`
}

type DescribeImageXCdnReuseRateAllResp struct {
	ReuseRateData []AllItem `json:"ReuseRateData"`
}

// DescribeImageXCdnProtocolRateByTime
// Method: POST

type DescribeImageXCdnProtocolRateByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXCdnProtocolRateByTimeResp struct {
	ProtocolRateData []CurveFloatItem `json:"ProtocolRateData"`
}

// DescribeImageXClientErrorCodeAll
// Method: POST

type DescribeImageXClientErrorCodeAllReq struct {
	Appid     string     `json:"Appid"`
	AppVer    []string   `json:"AppVer"`
	OS        string     `json:"OS"`
	SdkVer    []string   `json:"SdkVer"`
	Country   string     `json:"Country"`
	Province  string     `json:"Province"`
	Isp       []string   `json:"Isp"`
	Domain    []string   `json:"Domain"`
	ImageType []string   `json:"ImageType"`
	ExtraDims []ExtraDim `json:"ExtraDims"`
	StartTime string     `json:"StartTime"`
	EndTime   string     `json:"EndTime"`
	GroupBy   string     `json:"GroupBy"`
	OrderBy   string     `json:"OrderBy"`
	OrderAsc  bool       `json:"OrderAsc"`
}

type DescribeImageXClientErrorCodeAllResp struct {
	ErrorCodeData []ErrorCodeAllItem `json:"ErrorCodeData"`
}

// DescribeImageXClientErrorCodeByTime
// Method: POST

type DescribeImageXClientErrorCodeByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientErrorCodeByTimeResp struct {
	ErrorCodeData []ErrorCodeItem `json:"ErrorCodeData"`
}

// DescribeImageXClientDecodeSuccessRateByTime
// Method: POST

type DescribeImageXClientDecodeSuccessRateByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientDecodeSuccessRateByTimeResp struct {
	SuccessRateData []CurveFloatItem `json:"SuccessRateData"`
}

// DescribeImageXClientDecodeDurationByTime
// Method: POST

type DescribeImageXClientDecodeDurationByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientDecodeDurationByTimeResp struct {
	DurationData []CurveFloatItem `json:"DurationData"`
}

// DescribeImageXClientQueueDurationByTime
// Method: POST

type DescribeImageXClientQueueDurationByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientQueueDurationByTimeResp struct {
	DurationData []CurveFloatItem `json:"DurationData"`
}

// DescribeImageXClientLoadDurationAll
// Method: POST

type DescribeImageXClientLoadDurationAllReq struct {
	Appid     string     `json:"Appid"`
	AppVer    []string   `json:"AppVer"`
	OS        string     `json:"OS"`
	SdkVer    []string   `json:"SdkVer"`
	Country   string     `json:"Country"`
	Province  string     `json:"Province"`
	Isp       []string   `json:"Isp"`
	Domain    []string   `json:"Domain"`
	ImageType []string   `json:"ImageType"`
	ExtraDims []ExtraDim `json:"ExtraDims"`
	StartTime string     `json:"StartTime"`
	EndTime   string     `json:"EndTime"`
	GroupBy   string     `json:"GroupBy"`
	OrderBy   string     `json:"OrderBy"`
	OrderAsc  bool       `json:"OrderAsc"`
}

type DescribeImageXClientLoadDurationAllResp struct {
	DurationData []AllItem `json:"DurationData"`
}

// DescribeImageXClientLoadDuration
// Method: POST

type DescribeImageXClientLoadDurationReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientLoadDurationResp struct {
	DurationData []CurveFloatItem `json:"DurationData"`
}

// DescribeImageXClientFailureRate
// Method: POST

type DescribeImageXClientFailureRateReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientFailureRateResp struct {
	FailureRateData []CurveFloatItem `json:"FailureRateData"`
}

// DescribeImageXClientSdkVerByTime
// Method: POST

type DescribeImageXClientSdkVerByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientSdkVerByTimeResp struct {
	SdkVerData []SdkVerItem `json:"SdkVerData"`
}

type SdkVerItem struct {
	SdkVer string             `json:"SdkVer"`
	Count  int64              `json:"Count"`
	Data   []CurveDataIntItem `json:"Data"`
}

// DescribeImageXClientFileSize
// Method: POST

type DescribeImageXClientFileSizeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientFileSizeResp struct {
	FSizeData []CurveFloatItem `json:"FSizeData"`
}

// DescribeImageXClientTopFileSize
// Method: POST

type DescribeImageXClientTopFileSizeReq struct {
	Appid     string `json:"Appid"`
	OS        string `json:"OS"`
	StartTime string `json:"StartTime"`
	EndTime   string `json:"EndTime"`
	Top       int    `json:"Top"`
}

type DescribeImageXClientTopFileSizeResp struct {
	TopUrlData []ClientTopItem `json:"TopUrlData"`
}

// DescribeImageXClientCountByTime
// Method: POST

type DescribeImageXClientCountByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientCountByTimeResp struct {
	ClientCountData []CurveIntItem `json:"ClientCountData"`
}

// DescribeImageXClientScoreByTime
// Method: POST

type DescribeImageXClientScoreByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	ScoreType   string     `json:"ScoreType"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientScoreByTimeResp struct {
	ScoreData []CurveFloatItem `json:"ScoreData"`
}

// DescribeImageXClientDemotionRateByTime
// Method: POST

type DescribeImageXClientDemotionRateByTimeReq struct {
	Appid        string     `json:"Appid"`
	AppVer       []string   `json:"AppVer"`
	OS           string     `json:"OS"`
	SdkVer       []string   `json:"SdkVer"`
	Country      string     `json:"Country"`
	Province     string     `json:"Province"`
	Isp          []string   `json:"Isp"`
	Domain       []string   `json:"Domain"`
	ImageType    []string   `json:"ImageType"`
	ExtraDims    []ExtraDim `json:"ExtraDims"`
	GroupBy      string     `json:"GroupBy"`
	DemotionType string     `json:"DemotionType"`
	StartTime    string     `json:"StartTime"`
	EndTime      string     `json:"EndTime"`
	Granularity  string     `json:"Granularity"`
}

type DescribeImageXClientDemotionRateByTimeResp struct {
	DemotionRateData []CurveFloatItem `json:"DemotionRateData"`
}

// DescribeImageXClientTopDemotionURL
// Method: POST

type DescribeImageXClientTopDemotionURLReq struct {
	Appid        string `json:"Appid"`
	OS           string `json:"OS"`
	DemotionType string `json:"DemotionType"`
	StartTime    string `json:"StartTime"`
	EndTime      string `json:"EndTime"`
	Top          int    `json:"Top"`
}

type DescribeImageXClientTopDemotionURLResp struct {
	TopUrlData []ClientTopItem `json:"TopUrlData"`
}

// DescribeImageXClientQualityRateByTime
// Method: POST

type DescribeImageXClientQualityRateByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	Country     string     `json:"Country"`
	Province    string     `json:"Province"`
	Isp         []string   `json:"Isp"`
	Domain      []string   `json:"Domain"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	QualityType string     `json:"QualityType"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXClientQualityRateByTimeResp struct {
	QualityRateData []CurveFloatItem `json:"QualityRateData"`
}

// DescribeImageXClientTopQualityURL
// Method: POST

type DescribeImageXClientTopQualityURLReq struct {
	Appid       string `json:"Appid"`
	OS          string `json:"OS"`
	QualityType string `json:"QualityType"`
	StartTime   string `json:"StartTime"`
	EndTime     string `json:"EndTime"`
	Top         int    `json:"Top"`
}

type DescribeImageXClientTopQualityURLResp struct {
	TopUrlData []ClientTopItem `json:"TopUrlData"`
}

// DescribeImageXSensibleCountByTime
// Method: POST

type DescribeImageXSensibleCountByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXSensibleCountByTimeResp struct {
	SensibleCountData []CurveIntItem `json:"SensibleCountData"`
}

// DescribeImageXSensibleCacheHitRateByTime
// Method: POST

type DescribeImageXSensibleCacheHitRateByTimeReq struct {
	Appid       string     `json:"Appid"`
	AppVer      []string   `json:"AppVer"`
	OS          string     `json:"OS"`
	SdkVer      []string   `json:"SdkVer"`
	ImageType   []string   `json:"ImageType"`
	ExtraDims   []ExtraDim `json:"ExtraDims"`
	GroupBy     string     `json:"GroupBy"`
	Type        string     `json:"Type"`
	StartTime   string     `json:"StartTime"`
	EndTime     string     `json:"EndTime"`
	Granularity string     `json:"Granularity"`
}

type DescribeImageXSensibleCacheHitRateByTimeResp struct {
	CacheHitData []CurveFloatItem `json:"CacheHitData"`
}

// DescribeImageXSensibleTopSizeURL
// Method: POST

type DescribeImageXSensibleTopSizeURLReq struct {
	Appid      string     `json:"Appid"`
	AppVer     []string   `json:"AppVer"`
	OS         string     `json:"OS"`
	SdkVer     []string   `json:"SdkVer"`
	ImageType  []string   `json:"ImageType"`
	ExtraDims  []ExtraDim `json:"ExtraDims"`
	StartTime  string     `json:"StartTime"`
	EndTime    string     `json:"EndTime"`
	Top        int        `json:"Top"`
	OrderByIdx int        `json:"OrderByIdx"`
}

type DescribeImageXSensibleTopSizeURLResp struct {
	TopUrlData []SensibleSizeItem `json:"TopUrlData"`
}

type SensibleSizeItem struct {
	ActivityViewTree string `json:"ActivityViewTree"`
	Count            int64  `json:"Count"`
	Value            int64  `json:"Value"`
	URL              string `json:"URL"`
	BizTag           string `json:"BizTag"`
}

// DescribeImageXSensibleTopRamURL
// Method: POST

type DescribeImageXSensibleTopRamURLReq struct {
	Appid      string     `json:"Appid"`
	AppVer     []string   `json:"AppVer"`
	OS         string     `json:"OS"`
	SdkVer     []string   `json:"SdkVer"`
	ImageType  []string   `json:"ImageType"`
	ExtraDims  []ExtraDim `json:"ExtraDims"`
	StartTime  string     `json:"StartTime"`
	EndTime    string     `json:"EndTime"`
	Top        int        `json:"Top"`
	OrderByIdx int        `json:"OrderByIdx"`
}

type DescribeImageXSensibleTopRamURLResp struct {
	TopUrlData []SensibleRamItem `json:"TopUrlData"`
}

type SensibleRamItem struct {
	ActivityViewTree string `json:"ActivityViewTree"`
	Count            int64  `json:"Count"`
	RamSize          int64  `json:"RamSize"`
	URL              string `json:"URL"`
	BizTag           string `json:"BizTag"`
}

// DescribeImageXSensibleTopResolutionURL
// Method: POST

type DescribeImageXSensibleTopResolutionURLReq struct {
	Appid      string     `json:"Appid"`
	AppVer     []string   `json:"AppVer"`
	OS         string     `json:"OS"`
	SdkVer     []string   `json:"SdkVer"`
	ImageType  []string   `json:"ImageType"`
	ExtraDims  []ExtraDim `json:"ExtraDims"`
	StartTime  string     `json:"StartTime"`
	EndTime    string     `json:"EndTime"`
	Top        int        `json:"Top"`
	OrderByIdx int        `json:"OrderByIdx"`
}

type DescribeImageXSensibleTopResolutionURLResp struct {
	TopUrlData []SensibleResolutionItem `json:"TopUrlData"`
}

type SensibleResolutionItem struct {
	ActivityViewTree string `json:"ActivityViewTree"`
	Count            int64  `json:"Count"`
	URL              string `json:"URL"`
	BizTag           string `json:"BizTag"`
	ImageHeight      int64  `json:"ImageHeight"`
	ImageWidth       int64  `json:"ImageWidth"`
	ViewHeight       int64  `json:"ViewHeight"`
	ViewWidth        int64  `json:"ViewWidth"`
}

// DescribeImageXSensibleTopUnknownURL
// Method: POST

type DescribeImageXSensibleTopUnknownURLReq struct {
	Appid      string     `json:"Appid"`
	AppVer     []string   `json:"AppVer"`
	OS         string     `json:"OS"`
	SdkVer     []string   `json:"SdkVer"`
	ImageType  []string   `json:"ImageType"`
	ExtraDims  []ExtraDim `json:"ExtraDims"`
	StartTime  string     `json:"StartTime"`
	EndTime    string     `json:"EndTime"`
	Top        int        `json:"Top"`
	OrderByIdx int        `json:"OrderByIdx"`
}

type DescribeImageXSensibleTopUnknownURLResp struct {
	TopUrlData []SensibleUnknownItem `json:"TopUrlData"`
}

type SensibleUnknownItem struct {
	ActivityViewTree string `json:"ActivityViewTree"`
	Count            int64  `json:"Count"`
	URL              string `json:"URL"`
	BizTag           string `json:"BizTag"`
	ImageHeight      int64  `json:"ImageHeight"`
	ImageWidth       int64  `json:"ImageWidth"`
	ViewHeight       int64  `json:"ViewHeight"`
	ViewWidth        int64  `json:"ViewWidth"`
	FileSize         int64  `json:"FileSize"`
	RamSize          int64  `json:"RamSize"`
}
