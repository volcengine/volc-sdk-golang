package cdn

type AddCdnDomainRequest struct {
	Domain             string
	ServiceType        string
	Origin             []OriginRule
	OriginProtocol     string
	Project            *string `json:",omitempty"`
	ResourceTags       []ResourceTagEntry
	OriginHost         *string `json:",omitempty"`
	OriginRange        *bool   `json:",omitempty"`
	FollowRedirect     *bool   `json:",omitempty"`
	Cache              []CacheControlRule
	CacheKey           []CacheKeyGenerationRule
	NegativeCache      []NegativeCacheRule
	IpAccessRule       *IpAccessRule       `json:",omitempty"`
	RefererAccessRule  *RefererAccessRule  `json:",omitempty"`
	OriginAccessRule   *OriginAccessRule   `json:",omitempty"`
	SignedUrlAuth      *SignedUrlAuth      `json:",omitempty"`
	BandwidthLimit     *BandwidthLimit     `json:",omitempty"`
	IpFreqLimit        *IpFreqLimit        `json:",omitempty"`
	IpSpeedLimit       *IpSpeedLimit       `json:",omitempty"`
	DownloadSpeedLimit *DownloadSpeedLimit `json:",omitempty"`
	ResponseHeader     []ResponseHeaderRule
	RequestHeader      []RequestHeaderRule
	Compression        *Compression `json:",omitempty"`
	HTTPS              *HTTPS       `json:",omitempty"`
}
type AddCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type AddResourceTagsRequest struct {
	Resources    []string
	ResourceTags []ResourceTagEntry
}
type AddResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type BandwidthLimit struct {
	Switch             bool
	BandwidthLimitRule *BandwidthLimitRule `json:",omitempty"`
}
type BandwidthLimitAction struct {
	BandwidthThreshold int64
	LimitType          string
	SpeedLimitRate     int64
}
type BandwidthLimitRule struct {
	Condition            *Condition            `json:",omitempty"`
	BandwidthLimitAction *BandwidthLimitAction `json:",omitempty"`
}
type CacheAction struct {
	Action string
	Ttl    int64
}
type CacheControlRule struct {
	Condition   *Condition   `json:",omitempty"`
	CacheAction *CacheAction `json:",omitempty"`
}
type CacheKeyAction struct {
	CacheKeyComponents []CacheKeyComponent
}
type CacheKeyComponent struct {
	Object     string
	Action     string
	Subobject  string
	IgnoreCase bool
}
type CacheKeyGenerationRule struct {
	Condition      *Condition      `json:",omitempty"`
	CacheKeyAction *CacheKeyAction `json:",omitempty"`
}
type CertInfo struct {
	CertId string
}
type Compression struct {
	Switch           bool
	CompressionRules []CompressionRule
}
type CompressionAction struct {
	CompressionType   []string
	CompressionTarget string
}
type CompressionRule struct {
	Condition         *Condition         `json:",omitempty"`
	CompressionAction *CompressionAction `json:",omitempty"`
}
type Condition struct {
	Connective    string
	ConditionRule []ConditionRule
}
type ConditionRule struct {
	Type     string
	Object   string
	Operator string
	Value    string
}
type ContentTask struct {
	Url        string
	Status     string
	TaskType   string
	CreateTime int64
	TaskID     string
}
type DataDetail struct {
	Isp     string
	Region  string
	Metrics []MetricStatData
}
type DeleteCdnDomainRequest struct {
	Domain string
}
type DeleteCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type DeleteResourceTagsRequest struct {
	Resources    []string
	ResourceTags []ResourceTagEntry
}
type DeleteResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type DescribeAccountingDataRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    *string `json:",omitempty"`
}
type DescribeAccountingDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeAccountingDataResult
}
type DescribeAccountingDataResult struct {
	Resources []ResourceStatData
}
type DescribeCdnAccessLogRequest struct {
	Domain    string
	StartTime int64
	EndTime   int64
	PageNum   *int64 `json:",omitempty"`
	PageSize  *int64 `json:",omitempty"`
}
type DescribeCdnAccessLogResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnAccessLogResult
}
type DescribeCdnAccessLogResult struct {
	Domain           string
	PageSize         int64
	PageNum          int64
	TotalCount       int64
	DomainLogDetails []DomainLogDetail
}
type DescribeCdnConfigRequest struct {
	Domain string
}
type DescribeCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnConfigResult
}
type DescribeCdnConfigResult struct {
	DomainConfig DomainConfig
}
type DescribeCdnDataDetailRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    string
	Interval  *string `json:",omitempty"`
	Protocol  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
}
type DescribeCdnDataDetailResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataDetailResult
}
type DescribeCdnDataDetailResult struct {
	Domain      string
	DataDetails []DataDetail
}
type DescribeCdnDataRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    *string `json:",omitempty"`
	Interval  *string `json:",omitempty"`
	Area      *string `json:",omitempty"`
	Region    *string `json:",omitempty"`
	Isp       *string `json:",omitempty"`
	Protocol  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Aggregate *string `json:",omitempty"`
}
type DescribeCdnDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataResult
}
type DescribeCdnDataResult struct {
	Resources []ResourceStatData
}
type DescribeCdnDomainTopDataRequest struct {
	Domain      string
	StartTime   int64
	EndTime     int64
	Item        string
	Metric      string
	Area        *string `json:",omitempty"`
	ProjectName *string `json:",omitempty"`
}
type DescribeCdnDomainTopDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDomainTopDataResult
}
type DescribeCdnDomainTopDataResult struct {
	Domain         string
	TopDataDetails []TopDataDetailDeprecated
}
type DescribeCdnOriginDataRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    *string `json:",omitempty"`
	Interval  *string `json:",omitempty"`
	Aggregate *string `json:",omitempty"`
}
type DescribeCdnOriginDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnOriginDataResult
}
type DescribeCdnOriginDataResult struct {
	Resources []ResourceStatData
}
type DescribeCdnRegionAndIspRequest struct {
	Area *string `json:",omitempty"`
}
type DescribeCdnRegionAndIspResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnRegionAndIspResult
}
type DescribeCdnRegionAndIspResult struct {
	Isps    []NamePair
	Regions []NamePair
}
type DescribeCdnServiceResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnServiceResult
}
type DescribeCdnServiceResult struct {
	ServiceInfos []ServiceInformation
}
type DescribeCdnUpperIpRequest struct {
	Domain    string
	IpVersion string
}
type DescribeCdnUpperIpResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnUpperIpResult
}
type DescribeCdnUpperIpResult struct {
	CdnIpv4 []string
	CdnIpv6 []string
}
type DescribeContentBlockTasksRequest struct {
	Url       *string `json:",omitempty"`
	Domain    *string `json:",omitempty"`
	TaskID    *string `json:",omitempty"`
	TaskType  string
	Status    *string `json:",omitempty"`
	StartTime int64
	EndTime   int64
	PageNum   *int64 `json:",omitempty"`
	PageSize  *int64 `json:",omitempty"`
}
type DescribeContentBlockTasksResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentBlockTasksResult
}
type DescribeContentBlockTasksResult struct {
	Total    int64
	PageNum  int64
	PageSize int64
	Data     []ContentTask
}
type DescribeContentQuotaResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentQuotaResult
}
type DescribeContentQuotaResult struct {
	RefreshQuota     int64
	RefreshRemain    int64
	PreloadQuota     int64
	PreloadRemain    int64
	RefreshDirQuota  int64
	RefreshDirRemain int64
}
type DescribeContentTasksRequest struct {
	Url        *string `json:",omitempty"`
	DomainName *string `json:",omitempty"`
	TaskID     *string `json:",omitempty"`
	TaskType   string
	Status     *string `json:",omitempty"`
	StartTime  int64
	EndTime    int64
	PageNum    *int64 `json:",omitempty"`
	PageSize   *int64 `json:",omitempty"`
}
type DescribeContentTasksResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentTasksResult
}
type DescribeContentTasksResult struct {
	Total    int64
	PageNum  int64
	PageSize int64
	Data     []ContentTask
}
type DescribeEdgeNrtDataSummaryRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    *string `json:",omitempty"`
	Interval  *string `json:",omitempty"`
	Area      *string `json:",omitempty"`
	Region    *string `json:",omitempty"`
	Isp       *string `json:",omitempty"`
	Protocol  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Aggregate *string `json:",omitempty"`
}
type DescribeEdgeNrtDataSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeNrtDataSummaryResult
}
type DescribeEdgeNrtDataSummaryResult struct {
	Resources []ResourceSummary
}
type DescribeEdgeStatisticalDataRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    string
	Interval  *string `json:",omitempty"`
	Area      *string `json:",omitempty"`
	Region    *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
}
type DescribeEdgeStatisticalDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeStatisticalDataResult
}
type DescribeEdgeStatisticalDataResult struct {
	Resources []ResourceStatData
}
type DescribeEdgeTopNrtDataRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    *string `json:",omitempty"`
	Item      string
	Area      *string `json:",omitempty"`
	Interval  *string `json:",omitempty"`
}
type DescribeEdgeTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopNrtDataResult
}
type DescribeEdgeTopNrtDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopDetail
}
type DescribeEdgeTopStatisticalDataRequest struct {
	StartTime int64
	EndTime   int64
	Metric    *string `json:",omitempty"`
	Domain    string
	Item      string
	Area      *string `json:",omitempty"`
}
type DescribeEdgeTopStatisticalDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatisticalDataResult
}
type DescribeEdgeTopStatisticalDataResult struct {
	Domain         string
	TopDataDetails []TopDataDetail
}
type DescribeEdgeTopStatusCodeRequest struct {
	StartTime int64
	EndTime   int64
	Domain    *string `json:",omitempty"`
	Item      string
	Metric    string
	Area      *string `json:",omitempty"`
}
type DescribeEdgeTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatusCodeResult
}
type DescribeEdgeTopStatusCodeResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopStatusDetail
}
type DescribeIPInfoRequest struct {
	IP string
}
type DescribeIPInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeIPInfoResult
}
type DescribeIPInfoResult struct {
	IP       string
	Location string
	ISP      string
	CdnIp    bool
}
type DescribeOriginNrtDataSummaryRequest struct {
	StartTime int64
	EndTime   int64
	Metric    string
	Domain    *string `json:",omitempty"`
	Interval  *string `json:",omitempty"`
	Aggregate *string `json:",omitempty"`
}
type DescribeOriginNrtDataSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginNrtDataSummaryResult
}
type DescribeOriginNrtDataSummaryResult struct {
	Resources []ResourceSummary
}
type DescribeOriginTopNrtDataRequest struct {
	Domain    *string `json:",omitempty"`
	StartTime *int64  `json:",omitempty"`
	EndTime   *int64  `json:",omitempty"`
	Item      string
	Interval  *string `json:",omitempty"`
	Metric    string
}
type DescribeOriginTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopNrtDataResult
}
type DescribeOriginTopNrtDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopDetail
}
type DescribeOriginTopStatusCodeRequest struct {
	StartTime int64
	EndTime   int64
	Domain    *string `json:",omitempty"`
	Item      string
	Metric    string
}
type DescribeOriginTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopStatusCodeResult
}
type DescribeOriginTopStatusCodeResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopStatusDetail
}
type DomainConfig struct {
	Cname              string
	Domain             string
	Project            string
	ServiceType        string
	ServiceRegion      string
	UpdateTime         int64
	CreateTime         int64
	Status             string
	OriginHost         string
	OriginProtocol     string
	OriginRange        bool
	FollowRedirect     bool
	Origin             []OriginRule
	HTTPS              HTTPS
	IpAccessRule       IpAccessRule
	RefererAccessRule  RefererAccessRule
	OriginAccessRule   OriginAccessRule
	SignedUrlAuth      SignedUrlAuth
	BandwidthLimit     BandwidthLimit
	IpFreqLimit        IpFreqLimit
	IpSpeedLimit       IpSpeedLimit
	Cache              []CacheControlRule
	CacheKey           []CacheKeyGenerationRule
	NegativeCache      []NegativeCacheRule
	ResponseHeader     []ResponseHeaderRule
	RequestHeader      []RequestHeaderRule
	Compression        Compression
	DownloadSpeedLimit DownloadSpeedLimit
}
type DomainLogDetail struct {
	StartTime int64
	EndTime   int64
	LogName   string
	LogPath   string
	LogSize   int64
}
type DomainSummary struct {
	Domain        string
	Status        string
	Cname         string
	ServiceRegion string
	ServiceType   string
	CreateTime    int64
	UpdateTime    int64
	Resources     []ResourceTagEntry
}
type DownloadSpeedLimit struct {
	Switch                  *bool `json:",omitempty"`
	DownloadSpeedLimitRules []DownloadSpeedLimitRule
}
type DownloadSpeedLimitAction struct {
	SpeedLimitRate              int64
	SpeedLimitRateAfter         int64
	DownloadSpeedLimitRate      int64
	DownloadSpeedLimitRateAfter int64
}
type DownloadSpeedLimitRule struct {
	Condition                *Condition                `json:",omitempty"`
	DownloadSpeedLimitAction *DownloadSpeedLimitAction `json:",omitempty"`
}
type EmptyResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type ErrorObj struct {
	CodeN   int64
	Code    string
	Message string
}
type ForcedRedirect struct {
	EnableForcedRedirect bool
	StatusCode           string
}
type HTTPS struct {
	Switch         bool
	HTTP2          bool
	DisableHttp    bool
	TlsVersion     []string
	CertInfo       *CertInfo       `json:",omitempty"`
	ForcedRedirect *ForcedRedirect `json:",omitempty"`
}
type IpAccessRule struct {
	Switch   bool
	RuleType string
	Ip       []string
}
type IpFreqLimit struct {
	Switch           bool
	IpFreqLimitRules []IpFreqLimitRule
}
type IpFreqLimitAction struct {
	Action        string
	FreqLimitRate int64
	StateCode     string
}
type IpFreqLimitRule struct {
	Condition         *Condition         `json:",omitempty"`
	IpFreqLimitAction *IpFreqLimitAction `json:",omitempty"`
}
type IpSpeedLimit struct {
	Switch            bool
	IpSpeedLimitRules []IpSpeedLimitRule
}
type IpSpeedLimitAction struct {
	SpeedLimitRate int64
}
type IpSpeedLimitRule struct {
	Condition          *Condition          `json:",omitempty"`
	IpSpeedLimitAction *IpSpeedLimitAction `json:",omitempty"`
}
type ListCdnDomainsRequest struct {
	Domain       *string `json:",omitempty"`
	ServiceType  *string `json:",omitempty"`
	ResourceTags []ResourceTagEntry
	Status       *string `json:",omitempty"`
	PageNum      *int64  `json:",omitempty"`
	PageSize     *int64  `json:",omitempty"`
}
type ListCdnDomainsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListCdnDomainsResult
}
type ListCdnDomainsResult struct {
	Data     []DomainSummary
	PageNum  int64
	PageSize int64
	Total    int64
}
type ListResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListResourceTagsResult
}
type ListResourceTagsResult struct {
	ResourceTags []ResourceTagEntry
}
type MetricStatData struct {
	Metric string
	Values []TimeSeriesData
}
type MetricSummary struct {
	Metric string
	Value  float64
}
type NamePair struct {
	Code string
	Name string
}
type NegativeCacheAction struct {
	StatusCode string
	Action     string
	Ttl        int64
}
type NegativeCacheRule struct {
	Condition         *Condition `json:",omitempty"`
	NegativeCacheRule NegativeCacheAction
}
type OriginAccessRule struct {
	Switch     bool
	RuleType   string
	Origins    []string
	AllowEmpty bool
}
type OriginAction struct {
	OriginLines []OriginLine
}
type OriginLine struct {
	OriginType          string
	InstanceType        string
	Address             string
	HttpPort            string
	HttpsPort           string
	Weight              string
	PrivateBucketAccess *bool `json:",omitempty"`
}
type OriginRule struct {
	Condition    *Condition `json:",omitempty"`
	OriginAction OriginAction
}
type RefererAccessRule struct {
	Switch     bool
	RuleType   string
	Referers   []string
	AllowEmpty bool
}
type RequestHeaderAction struct {
	RequestHeaderInstances []RequestHeaderInstance
}
type RequestHeaderInstance struct {
	Action    string
	Key       string
	Value     string
	ValueType string
}
type RequestHeaderRule struct {
	Condition           *Condition           `json:",omitempty"`
	RequestHeaderAction *RequestHeaderAction `json:",omitempty"`
}
type ResourceStatData struct {
	Name    string
	Metrics []MetricStatData
}
type ResourceSummary struct {
	Name    string
	Metrics []MetricSummary
}
type ResourceTagEntry struct {
	Key   string
	Value string
}
type ResponseHeaderAction struct {
	ResponseHeaderInstances []ResponseHeaderInstance
}
type ResponseHeaderInstance struct {
	Action    string
	Key       string
	Value     string
	ValueType string
}
type ResponseHeaderRule struct {
	Condition            *Condition            `json:",omitempty"`
	ResponseHeaderAction *ResponseHeaderAction `json:",omitempty"`
}
type ResponseMetadata struct {
	RequestID string
	Service   *string   `json:",omitempty"`
	Region    *string   `json:",omitempty"`
	Action    *string   `json:",omitempty"`
	Version   *string   `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}
type ServiceInformation struct {
	Status       string
	CreateTime   string
	StartTime    string
	BillingCycle string
	BillingDesc  string
	BillingCode  string
	BillingData  string
	InstanceType string
}
type SignedUrlAuth struct {
	Switch             bool
	SignedUrlAuthRules []SignedUrlAuthRule
}
type SignedUrlAuthAction struct {
	URLAuthType     string
	MasterSecretKey string
	BackupSecretKey string
	SignName        string
	TimeName        string
	Duration        int64
	TimeFormat      string
	SignatureRule   []string
}
type SignedUrlAuthRule struct {
	Condition           *Condition           `json:",omitempty"`
	SignedUrlAuthAction *SignedUrlAuthAction `json:",omitempty"`
}
type StartCdnDomainRequest struct {
	Domain string
}
type StartCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type StopCdnDomainRequest struct {
	Domain string
}
type StopCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type SubmitBlockTaskRequest struct {
	Urls string
}
type SubmitBlockTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitBlockTaskResult
}
type SubmitBlockTaskResult struct {
	TaskID string
}
type SubmitPreloadTaskRequest struct {
	Urls string
}
type SubmitPreloadTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitPreloadTaskResult
}
type SubmitPreloadTaskResult struct {
	TaskID string
}
type SubmitRefreshTaskRequest struct {
	Type *string `json:",omitempty"`
	Urls string
}
type SubmitRefreshTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitRefreshTaskResult
}
type SubmitRefreshTaskResult struct {
	TaskID string
}
type SubmitUnblockTaskRequest struct {
	Urls string
}
type SubmitUnblockTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitUnblockTaskResult
}
type SubmitUnblockTaskResult struct {
	TaskID string
}
type TimeSeriesData struct {
	Timestamp int64
	Value     float64
}
type TopDataDetail struct {
	ItemKey   string
	ItemKeyCN string
	Value     float64
}
type TopDataDetailDeprecated struct {
	Item      string
	PV        int64
	PVRatio   float64
	Flux      int64
	FluxRatio float64
	Bandwidth float64
}
type TopDetail struct {
	ItemKey           string
	ItemKeyCN         string
	Bandwidth         float64
	BandwidthPeakTime int64
	Flux              float64
	FluxRatio         float64
	PV                int64
	PVRatio           int64
}
type TopStatusDetail struct {
	ItemKey        string
	Status2xx      float64 `json:"2xx"`
	Status2xxRatio float64 `json:"2xxRatio"`
	Status3xx      float64 `json:"3xx"`
	Status3xxRatio float64 `json:"3xxRatio"`
	Status4xx      float64 `json:"4xx"`
	Status4xxRatio float64 `json:"4xxRatio"`
	Status5xx      float64 `json:"5xx"`
	Status5xxRatio float64 `json:"5xxRatio"`
}
type UpdateCdnConfigRequest struct {
	Domain             string
	ResourceTags       []ResourceTagEntry
	OriginHost         *string `json:",omitempty"`
	OriginProtocol     *string `json:",omitempty"`
	OriginRange        *bool   `json:",omitempty"`
	FollowRedirect     *bool   `json:",omitempty"`
	Origin             []OriginRule
	Cache              []CacheControlRule
	CacheKey           []CacheKeyGenerationRule
	NegativeCache      []NegativeCacheRule
	IpAccessRule       *IpAccessRule       `json:",omitempty"`
	RefererAccessRule  *RefererAccessRule  `json:",omitempty"`
	OriginAccessRule   *OriginAccessRule   `json:",omitempty"`
	SignedUrlAuth      *SignedUrlAuth      `json:",omitempty"`
	BandwidthLimit     *BandwidthLimit     `json:",omitempty"`
	IpFreqLimit        *IpFreqLimit        `json:",omitempty"`
	IpSpeedLimit       *IpSpeedLimit       `json:",omitempty"`
	DownloadSpeedLimit *DownloadSpeedLimit `json:",omitempty"`
	ResponseHeader     []ResponseHeaderRule
	RequestHeader      []RequestHeaderRule
	Compression        *Compression `json:",omitempty"`
	HTTPS              *HTTPS       `json:",omitempty"`
}
type UpdateCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
type UpdateResourceTagsRequest struct {
	Resources    []string
	ResourceTags []ResourceTagEntry
}
type UpdateResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
