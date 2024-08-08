package cdn

type AccountingData struct {
	Metric string
	Values []DataPoint
}

type AccountingDataDetail struct {
	BillingRegion string
	Metrics       []AccountingData
	Name          string
}

type AccountingSummary struct {
	BillingCode string
	Name        string
	TimeStamp   int64
	Value       float64
	Values      []DataPoint
}

type AddCdnCertInfo struct {
	Desc *string `json:",omitempty"`
}

type AddCdnCertificateRequest struct {
	CertInfo      *AddCdnCertInfo `json:",omitempty"`
	CertType      *string         `json:",omitempty"`
	Certificate   Certificate
	CloseSigCheck *bool   `json:",omitempty"`
	EncryType     *string `json:",omitempty"`
	Source        *string `json:",omitempty"`
}

type AddCdnCertificateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           string
}

type AddCdnDomainRequest struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	ConditionalOrigin   *ConditionalOrigin   `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	Domain              string
	DownloadSpeedLimit  *DownloadSpeedLimit `json:",omitempty"`
	FollowRedirect      *bool               `json:",omitempty"`
	HTTPS               *HTTPS              `json:",omitempty"`
	HeaderLogging       *HeaderLog          `json:",omitempty"`
	HttpForcedRedirect  *HttpForcedRedirect `json:",omitempty"`
	IPv6                *IPv6               `json:",omitempty"`
	IpAccessRule        *IpAccessRule       `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit        `json:",omitempty"`
	MassCompression     *MassCompression    `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule   `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginCertCheck     *OriginCertCheck    `json:",omitempty"`
	OriginHost          *string             `json:",omitempty"`
	OriginIPv6          *string             `json:",omitempty"`
	OriginProtocol      *string             `json:",omitempty"`
	OriginRange         *bool               `json:",omitempty"`
	OriginRetry         *OriginRetry        `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Project             *string             `json:",omitempty"`
	Quic                *Quic               `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestBlockRule    *RequestBlockRule   `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResourceTags        []ResourceTag
	ResponseHeader      []ResponseHeaderRule
	ServiceRegion       *string `json:",omitempty"`
	ServiceType         string
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Sparrow             *Sparrow             `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UrlNormalize        *URLNormalize        `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
	WebpAdaptive        *WebpAdaptive        `json:",omitempty"`
}

type AddCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type AddCertificateRequest struct {
	CertType       *string `json:",omitempty"`
	Certificate    string
	CloseSigCheck  *bool   `json:",omitempty"`
	Desc           *string `json:",omitempty"`
	EncryType      *string `json:",omitempty"`
	EncryptionCert *string `json:",omitempty"`
	EncryptionKey  *string `json:",omitempty"`
	PrivateKey     string
	Repeatable     *bool   `json:",omitempty"`
	Source         *string `json:",omitempty"`
}

type AddCertificateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           AddCertificateResult
}

type AddCertificateResult struct {
	CertId string
}

type AddResourceTagsRequest struct {
	ResourceTags []ResourceTag
	Resources    []string
}

type AddResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type AreaAccessRule struct {
	Area     []string
	RuleType *string `json:",omitempty"`
	Switch   *bool   `json:",omitempty"`
}

type AuthCacheAction struct {
	Action     *string `json:",omitempty"`
	CacheKey   []string
	IgnoreCase *bool  `json:",omitempty"`
	Ttl        *int64 `json:",omitempty"`
}

type AuthModeConfig struct {
	BackupRemoteAddr *string `json:",omitempty"`
	MasterRemoteAddr *string `json:",omitempty"`
	PathType         *string `json:",omitempty"`
	PathValue        *string `json:",omitempty"`
	RequestMethod    *string `json:",omitempty"`
}

type AuthRequestHeaderRule struct {
	RequestHeaderComponents *RequestHeaderComponent `json:",omitempty"`
	RequestHeaderInstances  []RequestHeaderInstance
	RequestHost             *string `json:",omitempty"`
}

type AuthResponseConfig struct {
	CacheAction      *AuthCacheAction   `json:",omitempty"`
	ResponseAction   *ResponseAction    `json:",omitempty"`
	StatusCodeAction *StatusCodeAction  `json:",omitempty"`
	TimeOutAction    *AuthTimeoutAction `json:",omitempty"`
}

type AuthTimeoutAction struct {
	Action *string `json:",omitempty"`
	Time   *int64  `json:",omitempty"`
}

type AutoRotate struct {
	Switch *bool `json:",omitempty"`
}

type BandwidthLimit struct {
	BandwidthLimitRule *BandwidthLimitRule `json:",omitempty"`
	Switch             *bool               `json:",omitempty"`
}

type BandwidthLimitAction struct {
	BandwidthThreshold *int64  `json:",omitempty"`
	LimitType          *string `json:",omitempty"`
	SpeedLimitRate     *int64  `json:",omitempty"`
	SpeedLimitRateMax  *int64  `json:",omitempty"`
}

type BandwidthLimitRule struct {
	BandwidthLimitAction *BandwidthLimitAction `json:",omitempty"`
	Condition            *Condition            `json:",omitempty"`
}

type BatchDeployCertRequest struct {
	CertId  string
	CertId2 *string `json:",omitempty"`
	Domain  string
}

type BatchDeployCertResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           BatchDeployCertResult
}

type BatchDeployCertResult struct {
	DeployResult []CertDeployStatus
}

type BatchUpdateCdnConfigRequest struct {
	Aggregate           *bool           `json:",omitempty"`
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	ConditionalOrigin   *ConditionalOrigin   `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	Domains             []string
	DownloadSpeedLimit  *DownloadSpeedLimit `json:",omitempty"`
	FollowRedirect      *bool               `json:",omitempty"`
	HTTPS               *HTTPS              `json:",omitempty"`
	HttpForcedRedirect  *HttpForcedRedirect `json:",omitempty"`
	IPv6                *IPv6               `json:",omitempty"`
	IpAccessRule        *IpAccessRule       `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit        `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule   `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginCertCheck     *OriginCertCheck    `json:",omitempty"`
	OriginHost          *string             `json:",omitempty"`
	OriginIPv6          *string             `json:",omitempty"`
	OriginProtocol      *string             `json:",omitempty"`
	OriginRange         *bool               `json:",omitempty"`
	OriginRetry         *OriginRetry        `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Quic                *Quic               `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestBlockRule    *RequestBlockRule   `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	RewriteHLS          *RewriteHLS          `json:",omitempty"`
	ServiceRegion       *string              `json:",omitempty"`
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UrlNormalize        *URLNormalize        `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
}

type BatchUpdateCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           BatchUpdateCdnConfigResult
}

type BatchUpdateCdnConfigResult struct {
	DeployResult []DeployResult
}

type BlockAction struct {
	Action      *string `json:",omitempty"`
	ErrorPage   *string `json:",omitempty"`
	RedirectUrl *string `json:",omitempty"`
	StatusCode  *string `json:",omitempty"`
}

type BlockRule struct {
	BlockAction *BlockAction `json:",omitempty"`
	Condition   *Condition   `json:",omitempty"`
	RuleName    *string      `json:",omitempty"`
}

type BlockTaskInfo struct {
	BlockReason string
	CreateTime  int64
	Status      string
	TaskID      string
	TaskType    string
	Url         string
}

type BrowserCacheControlRule struct {
	CacheAction *CacheAction `json:",omitempty"`
	Condition   *Condition   `json:",omitempty"`
}

type CacheAction struct {
	Action        *string `json:",omitempty"`
	DefaultPolicy *string `json:",omitempty"`
	IgnoreCase    *bool   `json:",omitempty"`
	Ttl           *int64  `json:",omitempty"`
}

type CacheControlRule struct {
	CacheAction *CacheAction `json:",omitempty"`
	Condition   *Condition   `json:",omitempty"`
}

type CacheHost struct {
	CacheHostRule []CacheHostRule
	Switch        *bool `json:",omitempty"`
}

type CacheHostAction struct {
	CacheHost *string `json:",omitempty"`
}

type CacheHostRule struct {
	CacheHostAction *CacheHostAction `json:",omitempty"`
	Condition       *Condition       `json:",omitempty"`
}

type CacheKeyAction struct {
	CacheKeyComponents []CacheKeyComponent
}

type CacheKeyComponent struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	Object     *string `json:",omitempty"`
	Subobject  *string `json:",omitempty"`
}

type CacheKeyRule struct {
	CacheKeyAction *CacheKeyAction `json:",omitempty"`
	Condition      *Condition      `json:",omitempty"`
}

type CertCheck struct {
	CertInfoList []CertInfo
	Switch       *bool `json:",omitempty"`
}

type CertDeployStatus struct {
	Domain   string
	ErrorMsg string
	Status   string
}

type CertFingerprint struct {
	Sha1   string
	Sha256 string
}

type CertInfo struct {
	CertId        *string      `json:",omitempty"`
	CertName      *string      `json:",omitempty"`
	Certificate   *Certificate `json:",omitempty"`
	Desc          *string      `json:",omitempty"`
	EffectiveTime *int64       `json:",omitempty"`
	EncryType     *string      `json:",omitempty"`
	ExpireTime    *int64       `json:",omitempty"`
	Source        *string      `json:",omitempty"`
}

type Certificate struct {
	Certificate    *string `json:",omitempty"`
	EncryptionCert *string `json:",omitempty"`
	EncryptionKey  *string `json:",omitempty"`
	PrivateKey     *string `json:",omitempty"`
}

type CommonGlobalConfig struct {
	ConfigName *string `json:",omitempty"`
}

type CommonReferType struct {
	IgnoreCase   *bool `json:",omitempty"`
	IgnoreScheme *bool `json:",omitempty"`
	Referers     []string
}

type Compression struct {
	CompressionRules []CompressionRule
	Switch           *bool `json:",omitempty"`
}

type CompressionAction struct {
	CompressionFormat *string `json:",omitempty"`
	CompressionTarget *string `json:",omitempty"`
	CompressionType   []string
	MinFileSizeKB     *int64 `json:",omitempty"`
}

type CompressionRule struct {
	CompressionAction *CompressionAction `json:",omitempty"`
	Condition         *Condition         `json:",omitempty"`
}

type Condition struct {
	ConditionRule []ConditionRule
	Connective    *string `json:",omitempty"`
}

type ConditionRule struct {
	Name     *string `json:",omitempty"`
	Object   *string `json:",omitempty"`
	Operator *string `json:",omitempty"`
	Type     *string `json:",omitempty"`
	Value    *string `json:",omitempty"`
}

type ConditionalOrigin struct {
	OriginRules []OriginRules
	Switch      *bool `json:",omitempty"`
}

type ConditionalOriginAction struct {
	OriginLines []ConditionalOriginLine
}

type ConditionalOriginLine struct {
	Address      *string `json:",omitempty"`
	HttpPort     *string `json:",omitempty"`
	HttpsPort    *string `json:",omitempty"`
	InstanceType *string `json:",omitempty"`
	OriginHost   *string `json:",omitempty"`
}

type ConfiguredDomain struct {
	Domain string
	Type   string
}

type ContentTask struct {
	Area          string
	CreateTime    int64
	Delete        bool
	Isp           string
	Layer         string
	Process       string
	RefreshPrefix bool
	Region        string
	Remark        string
	Status        string
	SubArea       string
	TaskID        string
	TaskType      string
	Url           string
}

type CreateUsageReportRequest struct {
	Aggregate              *string `json:",omitempty"`
	BillingCode            *string `json:",omitempty"`
	BillingRegion          string
	CalculationMethod      *string `json:",omitempty"`
	Domain                 *string `json:",omitempty"`
	EndTime                int64
	ExportType             string
	FreeTimeTrafficCompute *bool   `json:",omitempty"`
	Metric                 *string `json:",omitempty"`
	StartTime              int64
	TaskName               string
	TlsTopic               *string `json:",omitempty"`
}

type CreateUsageReportResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           CreateUsageReportResult
}

type CreateUsageReportResult struct {
	TaskId string
}

type CustomErrorPage struct {
	ErrorPageRule []ErrorPageRule
	Switch        *bool `json:",omitempty"`
}

type CustomVariableInstance struct {
	Operator *string `json:",omitempty"`
	Type     *string `json:",omitempty"`
	Value    *string `json:",omitempty"`
}

type CustomVariableRules struct {
	CustomVariableInstances []CustomVariableInstance
}

type CustomizeAccessAction struct {
	AllowEmpty    *bool `json:",omitempty"`
	ListRules     []string
	RequestHeader *string `json:",omitempty"`
	RuleType      *string `json:",omitempty"`
}

type CustomizeAccessRule struct {
	CustomizeInstances []CustomizeInstance
	Switch             *bool `json:",omitempty"`
}

type CustomizeInstance struct {
	CustomizeRule *CustomizeRule `json:",omitempty"`
}

type CustomizeRule struct {
	AccessAction *CustomizeAccessAction `json:",omitempty"`
	Condition    *Condition             `json:",omitempty"`
}

type DataPoint struct {
	TimeStamp int64
	Value     float64
}

type DeleteCdnCertificateRequest struct {
	CertId string
}

type DeleteCdnCertificateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type DeleteCdnDomainRequest struct {
	Domain string
}

type DeleteCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type DeleteResourceTagsRequest struct {
	ResourceTags []ResourceTag
	Resources    []string
}

type DeleteResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type DeleteUsageReportRequest struct {
	TaskId string
}

type DeleteUsageReportResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DeleteUsageReportResult
}

type DeleteUsageReportResult struct {
	TaskIds []string
}

type DeployResult struct {
	Domain   string
	ErrorMsg string
	Status   string
}

type DescribeAccountingDataRequest struct {
	Aggregate        *string `json:",omitempty"`
	BillingRegion    *string `json:",omitempty"`
	Domain           *string `json:",omitempty"`
	EndTime          int64
	Interval         *int64 `json:",omitempty"`
	InverseDomain    *bool  `json:",omitempty"`
	IsWildcardDomain *bool  `json:",omitempty"`
	Metric           string
	Project          *string `json:",omitempty"`
	Protocol         *string `json:",omitempty"`
	StartTime        int64
}

type DescribeAccountingDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeAccountingDataResult
}

type DescribeAccountingDataResult struct {
	Resources []AccountingDataDetail
}

type DescribeAccountingSummaryRequest struct {
	Aggregate     *string `json:",omitempty"`
	BillingCode   *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        string
	EndTime       int64
	FreeTime      *string `json:",omitempty"`
	InverseDomain *bool   `json:",omitempty"`
	Project       *string `json:",omitempty"`
	StartTime     int64
	TimeZone      *string `json:",omitempty"`
}

type DescribeAccountingSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeAccountingSummaryResult
}

type DescribeAccountingSummaryResult struct {
	Resources []AccountingSummary
}

type DescribeCdnAccessLogRequest struct {
	Domain        string
	EndTime       int64
	PageNum       *int64  `json:",omitempty"`
	PageSize      *int64  `json:",omitempty"`
	ServiceRegion *string `json:",omitempty"`
	StartTime     int64
}

type DescribeCdnAccessLogResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnAccessLogResult
}

type DescribeCdnAccessLogResult struct {
	Domain           string
	DomainLogDetails []DomainLogDetail
	PageNum          int64
	PageSize         int64
	TotalCount       int64
}

type DescribeCdnConfigRequest struct {
	Domain   string
	LockInfo *bool `json:",omitempty"`
}

type DescribeCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnConfigResult
}

type DescribeCdnConfigResult struct {
	DomainConfig  DomainConfig
	FeatureConfig FeatureConfig
}

type DescribeCdnDataDetailRequest struct {
	Domain    string
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Metric    string
	Protocol  *string `json:",omitempty"`
	StartTime int64
	TimeZone  *string `json:",omitempty"`
}

type DescribeCdnDataDetailResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataDetailResult
}

type DescribeCdnDataDetailResult struct {
	DataDetails []NrtDataDetails
	Name        string
}

type DescribeCdnDataRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	InverseDomain       *bool   `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	IsWildcardDomain    *bool   `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
	Protocol            *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	StartTime           int64
	TimeZone            *string `json:",omitempty"`
}

type DescribeCdnDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataResult
}

type DescribeCdnDataResult struct {
	Resources []NrtDataResource
}

type DescribeCdnOriginDataRequest struct {
	Aggregate           *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	InverseDomain       *bool   `json:",omitempty"`
	IsWildcardDomain    *bool   `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
	StartTime           int64
	TimeZone            *string `json:",omitempty"`
}

type DescribeCdnOriginDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnOriginDataResult
}

type DescribeCdnOriginDataResult struct {
	Resources []NrtDataResource
}

type DescribeCdnRegionAndIspRequest struct {
	Area    *string `json:",omitempty"`
	Feature *string `json:",omitempty"`
	Lang    *string `json:",omitempty"`
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
	ServiceInfos []TopInstanceDetail
}

type DescribeCdnUpperIpRequest struct {
	Domain    string
	IpVersion *string `json:",omitempty"`
}

type DescribeCdnUpperIpResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnUpperIpResult
}

type DescribeCdnUpperIpResult struct {
	CdnIpv4 []string
	CdnIpv6 []string
}

type DescribeCertConfigRequest struct {
	CertId    string
	CertId2   *string `json:",omitempty"`
	CertType  *string `json:",omitempty"`
	EncryType *string `json:",omitempty"`
	Status    *string `json:",omitempty"`
}

type DescribeCertConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCertConfigResult
}

type DescribeCertConfigResult struct {
	CertNotConfig       []DomainStatus
	OtherCertConfig     []DomainCertStatus
	SpecifiedCertConfig []DomainCertStatus
}

type DescribeContentBlockTasksRequest struct {
	DomainName *string `json:",omitempty"`
	EndTime    *int64  `json:",omitempty"`
	PageNum    *int64  `json:",omitempty"`
	PageSize   *int64  `json:",omitempty"`
	StartTime  *int64  `json:",omitempty"`
	Status     *string `json:",omitempty"`
	TaskID     *string `json:",omitempty"`
	TaskType   string
	URL        *string `json:",omitempty"`
}

type DescribeContentBlockTasksResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentBlockTasksResult
}

type DescribeContentBlockTasksResult struct {
	Data     []BlockTaskInfo
	PageNum  int64
	PageSize int64
	Total    int64
}

type DescribeContentQuotaResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentQuotaResult
}

type DescribeContentQuotaResult struct {
	BlockLimit         int64
	BlockQuota         int64
	BlockRemain        int64
	PreloadLimit       int64
	PreloadQuota       int64
	PreloadRemain      int64
	RefreshDirLimit    int64
	RefreshDirQuota    int64
	RefreshDirRemain   int64
	RefreshQuota       int64
	RefreshQuotaLimit  int64
	RefreshRegexLimit  int64
	RefreshRegexQuota  int64
	RefreshRegexRemain int64
	RefreshRemain      int64
	UnblockLimit       int64
	UnblockQuota       int64
	UnblockRemain      int64
}

type DescribeContentTasksRequest struct {
	DomainName *string `json:",omitempty"`
	EndTime    *int64  `json:",omitempty"`
	PageNum    *int64  `json:",omitempty"`
	PageSize   *int64  `json:",omitempty"`
	Remark     *string `json:",omitempty"`
	StartTime  *int64  `json:",omitempty"`
	Status     *string `json:",omitempty"`
	TaskID     *string `json:",omitempty"`
	TaskType   string
	Url        *string `json:",omitempty"`
}

type DescribeContentTasksResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentTasksResult
}

type DescribeContentTasksResult struct {
	Data     []ContentTask
	PageNum  int64
	PageSize int64
	Total    int64
}

type DescribeDistrictDataRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Isp       *string `json:",omitempty"`
	Location  *string `json:",omitempty"`
	Metric    string
	Project   *string `json:",omitempty"`
	Protocol  *string `json:",omitempty"`
	Province  *string `json:",omitempty"`
	StartTime int64
}

type DescribeDistrictDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDistrictDataResult
}

type DescribeDistrictDataResult struct {
	MetricDataList []AccountingData
}

type DescribeDistrictIspDataRequest struct {
	Aggregate *string `json:",omitempty"`
	Domain    string
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Metric    string
	Protocol  *string `json:",omitempty"`
	StartTime int64
	TimeZone  *string `json:",omitempty"`
}

type DescribeDistrictIspDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDistrictIspDataResult
}

type DescribeDistrictIspDataResult struct {
	Resources []DomainNrtDetailData
}

type DescribeDistrictRankingRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	Item      string
	Metric    string
	Project   *string `json:",omitempty"`
	StartTime int64
}

type DescribeDistrictRankingResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDistrictRankingResult
}

type DescribeDistrictRankingResult struct {
	Item           string
	TopDataDetails []RankingDataDetail
}

type DescribeDistrictSummaryRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Isp       *string `json:",omitempty"`
	Location  *string `json:",omitempty"`
	Metric    string
	Project   *string `json:",omitempty"`
	Protocol  *string `json:",omitempty"`
	Province  *string `json:",omitempty"`
	StartTime int64
}

type DescribeDistrictSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDistrictSummaryResult
}

type DescribeDistrictSummaryResult struct {
	MetricDataList []MetricSummaryData
}

type DescribeEdgeDataRequest struct {
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Interval      *string `json:",omitempty"`
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeEdgeDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeDataResult
}

type DescribeEdgeDataResult struct {
	MetricDataList []AccountingData
}

type DescribeEdgeNrtDataSummaryRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	InverseDomain       *bool   `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
	Protocol            *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	StartTime           int64
}

type DescribeEdgeNrtDataSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeNrtDataSummaryResult
}

type DescribeEdgeNrtDataSummaryResult struct {
	Resources []NrtDataSummaryResource
}

type DescribeEdgeRankingRequest struct {
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Interval      *string `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeEdgeRankingResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeRankingResult
}

type DescribeEdgeRankingResult struct {
	Item           string
	TopDataDetails []RankingDataDetail
}

type DescribeEdgeStatisticalDataRequest struct {
	Area      *string `json:",omitempty"`
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Metric    string
	Region    *string `json:",omitempty"`
	StartTime int64
}

type DescribeEdgeStatisticalDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeStatisticalDataResult
}

type DescribeEdgeStatisticalDataResult struct {
	Resources []EdgeStatisticalDataResource
}

type DescribeEdgeStatusCodeRankingRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	Item      string
	Metric    string
	Project   *string `json:",omitempty"`
	StartTime int64
}

type DescribeEdgeStatusCodeRankingResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeStatusCodeRankingResult
}

type DescribeEdgeStatusCodeRankingResult struct {
	Item           string
	Metric         string
	TopDataDetails []RankingStatusCodeDetail
}

type DescribeEdgeSummaryRequest struct {
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Interval      *string `json:",omitempty"`
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeEdgeSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeSummaryResult
}

type DescribeEdgeSummaryResult struct {
	MetricDataList []MetricSummaryData
}

type DescribeEdgeTopNrtDataRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Interval      *string `json:",omitempty"`
	InverseDomain *bool   `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeEdgeTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopNrtDataResult
}

type DescribeEdgeTopNrtDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopNrtDataDetail
}

type DescribeEdgeTopStatisticalDataRequest struct {
	Area      *string `json:",omitempty"`
	Domain    string
	EndTime   int64
	Item      string
	Metric    string
	StartTime int64
	UaType    *string `json:",omitempty"`
}

type DescribeEdgeTopStatisticalDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatisticalDataResult
}

type DescribeEdgeTopStatisticalDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []EdgeTopStatisticalDataDetail
}

type DescribeEdgeTopStatusCodeRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	InverseDomain *bool `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeEdgeTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatusCodeResult
}

type DescribeEdgeTopStatusCodeResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopStatusCodeDetail
}

type DescribeIPInfoRequest struct {
	IP string
}

type DescribeIPInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeIPInfoResult
}

type DescribeIPInfoResult struct {
	CdnIp    bool
	IP       string
	ISP      string
	Location string
}

type DescribeIPListInfoRequest struct {
	IpList string
}

type DescribeIPListInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           []IPInfo
}

type DescribeOriginDataRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	Metric    string
	Project   *string `json:",omitempty"`
	StartTime int64
}

type DescribeOriginDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginDataResult
}

type DescribeOriginDataResult struct {
	MetricDataList []AccountingData
}

type DescribeOriginNrtDataSummaryRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	InverseDomain       *bool   `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
	Protocol            *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	StartTime           int64
}

type DescribeOriginNrtDataSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginNrtDataSummaryResult
}

type DescribeOriginNrtDataSummaryResult struct {
	Resources []NrtDataSummaryResource
}

type DescribeOriginRankingRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	Item      string
	Metric    string
	Project   *string `json:",omitempty"`
	StartTime int64
}

type DescribeOriginRankingResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginRankingResult
}

type DescribeOriginRankingResult struct {
	Item           string
	TopDataDetails []RankingDataDetail
}

type DescribeOriginStatusCodeRankingRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	Item      string
	Metric    string
	Project   *string `json:",omitempty"`
	StartTime int64
}

type DescribeOriginStatusCodeRankingResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginStatusCodeRankingResult
}

type DescribeOriginStatusCodeRankingResult struct {
	Item           string
	Metric         string
	TopDataDetails []RankingStatusCodeDetail
}

type DescribeOriginSummaryRequest struct {
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	Metric    string
	Project   *string `json:",omitempty"`
	StartTime int64
}

type DescribeOriginSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginSummaryResult
}

type DescribeOriginSummaryResult struct {
	MetricDataList []MetricSummaryData
}

type DescribeOriginTopNrtDataRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Interval      *string `json:",omitempty"`
	InverseDomain *bool   `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeOriginTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopNrtDataResult
}

type DescribeOriginTopNrtDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopNrtDataDetail
}

type DescribeOriginTopStatusCodeRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	InverseDomain *bool `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeOriginTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopStatusCodeResult
}

type DescribeOriginTopStatusCodeResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopStatusCodeDetail
}

type DescribeStatisticalRankingRequest struct {
	Area      *string `json:",omitempty"`
	Domain    string
	EndTime   int64
	Item      string
	Metric    string
	StartTime int64
	UaType    *string `json:",omitempty"`
}

type DescribeStatisticalRankingResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeStatisticalRankingResult
}

type DescribeStatisticalRankingResult struct {
	Item            string
	Metric          string
	RankingDataList []EdgeTopStatisticalDataDetail
	UaType          string
}

type DescribeUserDataRequest struct {
	Domain    string
	EndTime   int64
	Interval  string
	IpVersion *string `json:",omitempty"`
	Location  *string `json:",omitempty"`
	Province  *string `json:",omitempty"`
	StartTime int64
}

type DescribeUserDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeUserDataResult
}

type DescribeUserDataResult struct {
	MetricDataList []DataPoint
}

type DomainCertStatus struct {
	CerStatus  string
	Domain     string
	DomainLock DomainLock
	Status     string
	Type       string
}

type DomainConfig struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	AutoRotate          *AutoRotate     `json:",omitempty"`
	BackupCname         string
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Cname               *string              `json:",omitempty"`
	Compression         *Compression         `json:",omitempty"`
	CreateTime          *int64               `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	ConditionalOrigin   *ConditionalOrigin   `json:",omitempty"`
	Domain              *string              `json:",omitempty"`
	DownloadSpeedLimit  *DownloadSpeedLimit  `json:",omitempty"`
	FollowRedirect      *bool                `json:",omitempty"`
	HTTPS               *HTTPS               `json:",omitempty"`
	HeaderLogging       *HeaderLog           `json:",omitempty"`
	HttpForcedRedirect  *HttpForcedRedirect  `json:",omitempty"`
	IPv6                *IPv6                `json:",omitempty"`
	IpAccessRule        *IpAccessRule        `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit         `json:",omitempty"`
	LockStatus          *string              `json:",omitempty"`
	MassCompression     *MassCompression     `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule    `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginCertCheck     OriginCertCheck
	OriginHost          *string `json:",omitempty"`
	OriginIPv6          string
	OriginProtocol      string
	OriginRange         *bool               `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Project             *string             `json:",omitempty"`
	Quic                *Quic               `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestBlockRule    RequestBlockRule
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	RewriteHLS          RewriteHLS
	ServiceRegion       *string `json:",omitempty"`
	ServiceType         string
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Sparrow             *Sparrow             `json:",omitempty"`
	Status              *string              `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UpdateTime          *int64               `json:",omitempty"`
	UrlNormalize        URLNormalize
	VideoDrag           *VideoDrag    `json:",omitempty"`
	WebpAdaptive        *WebpAdaptive `json:",omitempty"`
	Websocket           Websocket
}

type DomainLock struct {
	Remark string
	Status string
}

type DomainLogDetail struct {
	EndTime   int64
	LogName   string
	LogPath   string
	LogSize   int64
	StartTime int64
}

type DomainNrtDetailData struct {
	DataDetails []NrtDataDetails
	Name        string
}

type DomainStatus struct {
	Domain     string
	DomainLock DomainLock
	Status     string
	Type       string
}

type DomainSummary struct {
	BackupCname           string
	BackupOrigin          []string
	CacheShared           string
	CacheSharedTargetHost string
	Cname                 string
	ConfigStatus          string
	CreateTime            int64
	Domain                string
	DomainLock            DomainLock
	FeatureConfig         FeatureConfig
	HTTPS                 bool
	IPv6                  bool
	IsConflictDomain      bool
	OriginProtocol        string
	PrimaryOrigin         []string
	Project               string
	ResourceTags          []ResourceTag
	ServiceRegion         string
	ServiceType           string
	SparrowList           []string
	Status                string
	UpdateTime            int64
	Waf                   bool
}

type DownloadSpeedLimit struct {
	DownloadSpeedLimitRules []DownloadSpeedLimitRule
	Switch                  *bool `json:",omitempty"`
}

type DownloadSpeedLimitAction struct {
	SpeedLimitRate      *int64          `json:",omitempty"`
	SpeedLimitRateAfter *int64          `json:",omitempty"`
	SpeedLimitTime      *SpeedLimitTime `json:",omitempty"`
}

type DownloadSpeedLimitRule struct {
	Condition                *Condition                `json:",omitempty"`
	DownloadSpeedLimitAction *DownloadSpeedLimitAction `json:",omitempty"`
}

type EdgeStatisticalDataResource struct {
	Metrics []AccountingData
	Name    string
}

type EdgeTopStatisticalDataDetail struct {
	ItemKey   string
	ItemKeyCN string
	Value     float64
}

type ErrorObj struct {
	CodeN   int64
	Code    string
	Message string
}

type ErrorPageAction struct {
	Action       *string `json:",omitempty"`
	RedirectCode *string `json:",omitempty"`
	RedirectUrl  *string `json:",omitempty"`
	StatusCode   *string `json:",omitempty"`
}

type ErrorPageRule struct {
	ErrorPageAction *ErrorPageAction `json:",omitempty"`
}

type FeatureConfig struct {
	OriginV2 bool
}

type ForcedRedirect struct {
	EnableForcedRedirect *bool   `json:",omitempty"`
	StatusCode           *string `json:",omitempty"`
}

type HTTPS struct {
	CertCheck      *CertCheck `json:",omitempty"`
	CertInfo       *CertInfo  `json:",omitempty"`
	CertInfoList   []CertInfo
	DisableHttp    *bool           `json:",omitempty"`
	ForcedRedirect *ForcedRedirect `json:",omitempty"`
	HTTP2          *bool           `json:",omitempty"`
	Hsts           *Hsts           `json:",omitempty"`
	OCSP           *bool           `json:",omitempty"`
	Switch         *bool           `json:",omitempty"`
	TlsVersion     []string
}

type HeaderLog struct {
	HeaderLogging *string `json:",omitempty"`
	Switch        *bool   `json:",omitempty"`
}

type Hsts struct {
	Subdomain *string `json:",omitempty"`
	Switch    *bool   `json:",omitempty"`
	Ttl       *int64  `json:",omitempty"`
}

type HttpForcedRedirect struct {
	EnableForcedRedirect *bool   `json:",omitempty"`
	StatusCode           *string `json:",omitempty"`
}

type IPInfo struct {
	CdnIp    bool
	IP       string
	ISP      string
	Location string
}

type IPv6 struct {
	Switch *bool `json:",omitempty"`
}

type IpAccessRule struct {
	Ip           []string
	RuleType     *string             `json:",omitempty"`
	SharedConfig *CommonGlobalConfig `json:",omitempty"`
	Switch       *bool               `json:",omitempty"`
}

type IpFreqLimit struct {
	IpFreqLimitRules []IpFreqLimitRule
	Switch           *bool `json:",omitempty"`
}

type IpFreqLimitAction struct {
	Action        *string `json:",omitempty"`
	FreqLimitRate *int64  `json:",omitempty"`
	StatusCode    *string `json:",omitempty"`
}

type IpFreqLimitRule struct {
	Condition         *Condition         `json:",omitempty"`
	IpFreqLimitAction *IpFreqLimitAction `json:",omitempty"`
}

type ListCdnCertInfoRequest struct {
	CertId           *string `json:",omitempty"`
	CertType         *string `json:",omitempty"`
	Configured       *bool   `json:",omitempty"`
	ConfiguredDomain *string `json:",omitempty"`
	DnsName          *string `json:",omitempty"`
	EncryType        *string `json:",omitempty"`
	PageNum          *int64  `json:",omitempty"`
	PageSize         *int64  `json:",omitempty"`
	Source           *string `json:",omitempty"`
	Status           *string `json:",omitempty"`
}

type ListCdnCertInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListCdnCertInfoResult
}

type ListCdnCertInfoResult struct {
	CertInfo      []ListCertInfo
	ExpiringCount int64
	PageNum       int64
	PageSize      int64
	Total         int64
}

type ListCdnDomainsRequest struct {
	Domain         *string `json:",omitempty"`
	ExactMatch     *bool   `json:",omitempty"`
	FeatureConfig  *bool   `json:",omitempty"`
	HTTPS          *bool   `json:",omitempty"`
	IPv6           *bool   `json:",omitempty"`
	NeedSparrows   *bool   `json:",omitempty"`
	OriginProtocol *string `json:",omitempty"`
	PageNum        *int64  `json:",omitempty"`
	PageSize       *int64  `json:",omitempty"`
	PrimaryOrigin  *string `json:",omitempty"`
	Project        *string `json:",omitempty"`
	ResourceTags   []string
	ServiceRegion  *string `json:",omitempty"`
	ServiceType    *string `json:",omitempty"`
	Status         *string `json:",omitempty"`
	TagConnective  *string `json:",omitempty"`
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

type ListCertInfo struct {
	CertFingerprint        CertFingerprint
	CertId                 string
	CertName               string
	CertType               string
	ConfiguredDomain       string
	ConfiguredDomainDetail []ConfiguredDomain
	Desc                   string
	DnsName                string
	EffectiveTime          int64
	EncryType              string
	ExpireTime             int64
	Source                 string
	Status                 string
}

type ListCertInfoRequest struct {
	CertId           *string   `json:",omitempty"`
	CertType         *string   `json:",omitempty"`
	ConfiguredDomain *string   `json:",omitempty"`
	EncryType        *string   `json:",omitempty"`
	FuzzyMatch       *bool     `json:",omitempty"`
	Name             *string   `json:",omitempty"`
	PageNum          *int64    `json:",omitempty"`
	PageSize         *int64    `json:",omitempty"`
	SetPagination    *bool     `json:",omitempty"`
	SortRule         *SortRule `json:",omitempty"`
	Source           string
	Status           *string `json:",omitempty"`
}

type ListCertInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListCertInfoResult
}

type ListCertInfoResult struct {
	CertInfo      []ListCertInfo
	ExpiringCount int64
	PageNum       int64
	PageSize      int64
	Total         int64
}

type ListResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListResourceTagsResult
}

type ListResourceTagsResult struct {
	ResourceTags []ResourceTag
}

type ListUsageReportsRequest struct {
	ExportType *string `json:",omitempty"`
	PageNum    *int64  `json:",omitempty"`
	PageSize   *int64  `json:",omitempty"`
	Status     *int64  `json:",omitempty"`
	TaskName   *string `json:",omitempty"`
}

type ListUsageReportsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListUsageReportsResult
}

type ListUsageReportsResult struct {
	PageNum             int64
	PageSize            int64
	Total               int64
	UsageReportsDetails []UsageReportsDetail
}

type MassCompression struct {
	MassCompressionRule []MassCompressionRule
	Switch              *bool `json:",omitempty"`
}

type MassCompressionAction struct {
	CompressionRatio *int64  `json:",omitempty"`
	CompressionType  *string `json:",omitempty"`
}

type MassCompressionRule struct {
	CompressionAction *MassCompressionAction `json:",omitempty"`
}

type MethodDeniedRule struct {
	Methods *string `json:",omitempty"`
	Switch  *bool   `json:",omitempty"`
}

type MetricSummaryData struct {
	Metric string
	Value  float64
}

type MetricTimestampValue struct {
	Metric string
	Values []TimestampValue
}

type MetricValue struct {
	Metric string
	Value  float64
}

type NamePair struct {
	Code string
	Name string
}

type NegativeCache struct {
	Condition         *Condition           `json:",omitempty"`
	NegativeCacheRule *NegativeCacheAction `json:",omitempty"`
}

type NegativeCacheAction struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	StatusCode *string `json:",omitempty"`
	Ttl        *int64  `json:",omitempty"`
}

type NrtDataDetails struct {
	Isp     string
	Metrics []MetricTimestampValue
	Region  string
}

type NrtDataResource struct {
	BillingRegion string
	Isp           string
	Metrics       []MetricTimestampValue
	Name          string
	Region        string
}

type NrtDataSummaryResource struct {
	BillingRegion string
	Metrics       []MetricValue
	Name          string
}

type OriginAccessRule struct {
	AllowEmpty *bool `json:",omitempty"`
	IgnoreCase *bool `json:",omitempty"`
	Origins    []string
	RuleType   *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
}

type OriginAction struct {
	OriginLines []OriginLine
}

type OriginArgAction struct {
	OriginArgComponents []OriginArgComponents
}

type OriginArgComponents struct {
	Action    *string `json:",omitempty"`
	Object    *string `json:",omitempty"`
	Subobject *string `json:",omitempty"`
}

type OriginArgRule struct {
	Condition       *Condition       `json:",omitempty"`
	OriginArgAction *OriginArgAction `json:",omitempty"`
}

type OriginCertCheck struct {
	Switch *bool `json:",omitempty"`
}

type OriginLine struct {
	Address             *string            `json:",omitempty"`
	BucketName          *string            `json:",omitempty"`
	HttpPort            *string            `json:",omitempty"`
	HttpsPort           *string            `json:",omitempty"`
	InstanceType        *string            `json:",omitempty"`
	OriginHost          *string            `json:",omitempty"`
	OriginType          *string            `json:",omitempty"`
	PrivateBucketAccess *bool              `json:",omitempty"`
	PrivateBucketAuth   *PrivateBucketAuth `json:",omitempty"`
	Region              *string            `json:",omitempty"`
	SignedOriginAuth    *SignedOriginAuth  `json:",omitempty"`
	Weight              *string            `json:",omitempty"`
}

type OriginRetry struct {
	StatusCode *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
}

type OriginRewrite struct {
	OriginRewriteRule []OriginRewriteRule
	Switch            *bool `json:",omitempty"`
}

type OriginRewriteAction struct {
	RewriteType *string `json:",omitempty"`
	SourcePath  *string `json:",omitempty"`
	TargetPath  *string `json:",omitempty"`
}

type OriginRewriteRule struct {
	Condition           *Condition           `json:",omitempty"`
	OriginRewriteAction *OriginRewriteAction `json:",omitempty"`
}

type OriginRule struct {
	Condition    *Condition    `json:",omitempty"`
	OriginAction *OriginAction `json:",omitempty"`
}

type OriginRules struct {
	Actions   *ConditionalOriginAction `json:",omitempty"`
	Condition *RecursionCondition      `json:",omitempty"`
}

type OriginSni struct {
	SniDomain *string `json:",omitempty"`
	Switch    *bool   `json:",omitempty"`
}

type PageOptimization struct {
	OptimizationType []string
	Switch           *bool `json:",omitempty"`
}

type PreloadHeader struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

type PrivateBucketAuth struct {
	AuthType           *string             `json:",omitempty"`
	Switch             *bool               `json:",omitempty"`
	TosAuthInformation *TosAuthInformation `json:",omitempty"`
}

type QueryStringComponents struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type QueryStringInstance struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type QueryStringRule struct {
	QueryStringComponents *QueryStringComponents `json:",omitempty"`
	QueryStringInstances  []QueryStringInstance
}

type Quic struct {
	Switch *bool `json:",omitempty"`
}

type RankingDataDetail struct {
	Metric       string
	ValueDetails []RankingValueDetail
}

type RankingStatusCodeDetail struct {
	ItemKey        string
	Status2xx      float64
	Status2xxRatio float64
	Status3xx      float64
	Status3xxRatio float64
	Status4xx      float64
	Status4xxRatio float64
	Status5xx      float64
	Status5xxRatio float64
}

type RankingValueDetail struct {
	ItemKey   string
	Ratio     float64
	Timestamp int64
	Value     float64
}

type RecursionCondition struct {
	ConditionGroups []SubRecursionCondition
	Connective      *string `json:",omitempty"`
	IsGroup         *bool   `json:",omitempty"`
}

type RecursionConditionRule struct {
	Object   *string `json:",omitempty"`
	Operator *string `json:",omitempty"`
	Value    []string
}

type RedirectionAction struct {
	RedirectCode          *string                `json:",omitempty"`
	SourcePath            *string                `json:",omitempty"`
	TargetHost            *string                `json:",omitempty"`
	TargetPath            *string                `json:",omitempty"`
	TargetProtocol        *string                `json:",omitempty"`
	TargetQueryComponents *TargetQueryComponents `json:",omitempty"`
}

type RedirectionRewrite struct {
	RedirectionRule []RedirectionRule
	Switch          *bool `json:",omitempty"`
}

type RedirectionRule struct {
	RedirectionAction *RedirectionAction `json:",omitempty"`
}

type RefererAccessRule struct {
	AllowEmpty   *bool `json:",omitempty"`
	Referers     []string
	ReferersType *ReferersType       `json:",omitempty"`
	RuleType     *string             `json:",omitempty"`
	SharedConfig *CommonGlobalConfig `json:",omitempty"`
	Switch       *bool               `json:",omitempty"`
}

type RefererType struct {
	Referers []string
}

type ReferersType struct {
	CommonType  *CommonReferType `json:",omitempty"`
	RegularType *RefererType     `json:",omitempty"`
}

type RemoteAuth struct {
	RemoteAuthRules []RemoteAuthRule
	Switch          *bool `json:",omitempty"`
}

type RemoteAuthRule struct {
	Condition            *Condition            `json:",omitempty"`
	RemoteAuthRuleAction *RemoteAuthRuleAction `json:",omitempty"`
}

type RemoteAuthRuleAction struct {
	AuthModeConfig     *AuthModeConfig        `json:",omitempty"`
	AuthResponseConfig *AuthResponseConfig    `json:",omitempty"`
	QueryStringRules   *QueryStringRule       `json:",omitempty"`
	RequestBodyRules   *string                `json:",omitempty"`
	RequestHeaderRules *AuthRequestHeaderRule `json:",omitempty"`
}

type RequestBlockRule struct {
	BlockRule []BlockRule
	Switch    *bool `json:",omitempty"`
}

type RequestHeaderAction struct {
	RequestHeaderInstances []RequestHeaderInstance
}

type RequestHeaderComponent struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type RequestHeaderInstance struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type RequestHeaderRule struct {
	Condition           *Condition           `json:",omitempty"`
	RequestHeaderAction *RequestHeaderAction `json:",omitempty"`
}

type ResourceTag struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

type ResponseAction struct {
	StatusCode *string `json:",omitempty"`
}

type ResponseHeaderAction struct {
	ResponseHeaderInstances []ResponseHeaderInstance
}

type ResponseHeaderInstance struct {
	AccessOriginControl *bool   `json:",omitempty"`
	Action              *string `json:",omitempty"`
	Key                 *string `json:",omitempty"`
	Value               *string `json:",omitempty"`
	ValueType           *string `json:",omitempty"`
}

type ResponseHeaderRule struct {
	Condition            *Condition            `json:",omitempty"`
	ResponseHeaderAction *ResponseHeaderAction `json:",omitempty"`
}

type ResponseMetadata struct {
	RequestId string
	Service   *string   `json:",omitempty"`
	Region    *string   `json:",omitempty"`
	Action    *string   `json:",omitempty"`
	Version   *string   `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}

type RewriteHLS struct {
	SignName *string `json:",omitempty"`
	Switch   *bool   `json:",omitempty"`
}

type RewriteM3u8Rule struct {
	DeleteParam      *bool `json:",omitempty"`
	KeepM3u8Param    *bool `json:",omitempty"`
	TransferEncoding *bool `json:",omitempty"`
}

type SharedCname struct {
	Cname  *string `json:",omitempty"`
	Switch *bool   `json:",omitempty"`
}

type SignedOriginAuth struct {
	SignedOriginAuthRules []SignedOriginAuthRule
	Switch                *bool `json:",omitempty"`
}

type SignedOriginAuthAction struct {
	Duration        *int64  `json:",omitempty"`
	KeyName         *string `json:",omitempty"`
	MasterAccessKey *string `json:",omitempty"`
	MasterSecretKey *string `json:",omitempty"`
	OriginAuthType  *string `json:",omitempty"`
	SignName        *string `json:",omitempty"`
	TimeFormat      *string `json:",omitempty"`
	TimeName        *string `json:",omitempty"`
}

type SignedOriginAuthRule struct {
	Condition              *Condition              `json:",omitempty"`
	SignedOriginAuthAction *SignedOriginAuthAction `json:",omitempty"`
}

type SignedUrlAuth struct {
	SignedUrlAuthRules []SignedUrlAuthRule
	Switch             *bool `json:",omitempty"`
}

type SignedUrlAuthAction struct {
	AuthAlgorithm       *string              `json:",omitempty"`
	BackupSecretKey     *string              `json:",omitempty"`
	CustomVariableRules *CustomVariableRules `json:",omitempty"`
	Duration            *int64               `json:",omitempty"`
	KeepOriginArg       *bool                `json:",omitempty"`
	MasterSecretKey     *string              `json:",omitempty"`
	MpdVarExpand        *bool                `json:",omitempty"`
	RewriteM3u8         *bool                `json:",omitempty"`
	RewriteM3u8Rule     *RewriteM3u8Rule     `json:",omitempty"`
	RewriteMpd          *bool                `json:",omitempty"`
	SignName            *string              `json:",omitempty"`
	SignatureRule       []string
	TimeFormat          *string `json:",omitempty"`
	TimeName            *string `json:",omitempty"`
	URLAuthType         *string `json:",omitempty"`
}

type SignedUrlAuthRule struct {
	Condition           *Condition           `json:",omitempty"`
	SignedUrlAuthAction *SignedUrlAuthAction `json:",omitempty"`
}

type SignedUrlAuthRules struct {
	Duration *int64 `json:",omitempty"`
}

type SortRule struct {
	Asc     *bool   `json:",omitempty"`
	OrderBy *string `json:",omitempty"`
}

type Sparrow struct {
	SparrowRules []SparrowRule
	Switch       *bool `json:",omitempty"`
}

type SparrowAction struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	SparrowID  *string `json:",omitempty"`
}

type SparrowRule struct {
	Condition     *Condition     `json:",omitempty"`
	SparrowAction *SparrowAction `json:",omitempty"`
}

type SpeedLimitTime struct {
	BeginTime *string `json:",omitempty"`
	DayWeek   *string `json:",omitempty"`
	EndTime   *string `json:",omitempty"`
}

type StartCdnDomainRequest struct {
	Domain string
}

type StartCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type StatusCodeAction struct {
	DefaultAction *string `json:",omitempty"`
	FailCode      *string `json:",omitempty"`
	SuccessCode   *string `json:",omitempty"`
}

type StopCdnDomainRequest struct {
	Domain string
}

type StopCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type SubRecursionCondition struct {
	Condition *RecursionConditionRule `json:",omitempty"`
	IsGroup   *bool                   `json:",omitempty"`
}

type SubmitBlockTaskRequest struct {
	Type *string `json:",omitempty"`
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
	Area                   *string `json:",omitempty"`
	ConcurrentLimit        *int64  `json:",omitempty"`
	Deduplicate            *bool   `json:",omitempty"`
	Isp                    *string `json:",omitempty"`
	Layer                  *string `json:",omitempty"`
	Region                 *string `json:",omitempty"`
	RequestHeaderInstances []PreloadHeader
	SubArea                *string `json:",omitempty"`
	Urls                   string
}

type SubmitPreloadTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitPreloadTaskResult
}

type SubmitPreloadTaskResult struct {
	CommitNum *int64 `json:",omitempty"`
	TaskID    string
}

type SubmitRefreshTaskRequest struct {
	Delete *bool   `json:",omitempty"`
	Prefix *bool   `json:",omitempty"`
	Type   *string `json:",omitempty"`
	Urls   string
}

type SubmitRefreshTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitRefreshTaskResult
}

type SubmitRefreshTaskResult struct {
	TaskID string
}

type SubmitUnblockTaskRequest struct {
	Type *string `json:",omitempty"`
	Urls string
}

type SubmitUnblockTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitUnblockTaskResult
}

type SubmitUnblockTaskResult struct {
	TaskID string
}

type TargetQueryComponents struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type Timeout struct {
	Switch       *bool `json:",omitempty"`
	TimeoutRules []TimeoutRule
}

type TimeoutAction struct {
	HttpTimeout *int64 `json:",omitempty"`
	TcpTimeout  *int64 `json:",omitempty"`
}

type TimeoutRule struct {
	Condition     *Condition     `json:",omitempty"`
	TimeoutAction *TimeoutAction `json:",omitempty"`
}

type TimestampValue struct {
	Timestamp int64
	Value     float64
}

type TopInstanceDetail struct {
	BeginTime        string
	BillingCode      string
	BillingCycle     string
	BillingData      string
	BillingDesc      string
	CreateTime       string
	FreeTimePeriods  []int64
	InstanceCategory string
	InstanceType     string
	MetricType       string
	ServiceRegion    string
	StartTime        string
	Status           string
}

type TopNrtDataDetail struct {
	Bandwidth                float64
	BandwidthPeakTime        int64
	BsBandwidth              float64
	BsBandwidthPeakTime      int64
	BsFlux                   float64
	BsFluxRatio              float64
	DynamicRequest           int64
	DynamicRequestRatio      float64
	Flux                     float64
	FluxRatio                float64
	InboundBandwidth         float64
	InboundBandwidthPeakTime int64
	InboundFlux              float64
	InboundFluxRatio         float64
	ItemKey                  string
	ItemKeyCN                string
	PV                       float64
	PVRatio                  float64
	Quic                     int64
	StaticRequest            int64
	StaticRequestRatio       float64
}

type TopStatusCodeDetail struct {
	Status2xx      float64 `json:"2xx"`
	Status2xxRatio float64 `json:"2xxRatio"`
	Status3xx      float64 `json:"3xx"`
	Status3xxRatio float64 `json:"3xxRatio"`
	Status4xx      float64 `json:"4xx"`
	Status4xxRatio float64 `json:"4xxRatio"`
	Status5xx      float64 `json:"5xx"`
	Status5xxRatio float64 `json:"5xxRatio"`
	ItemKey        string
}

type TosAuthInformation struct {
	AccessKeyId       *string `json:",omitempty"`
	AccessKeySecret   *string `json:",omitempty"`
	AccountKey        *string `json:",omitempty"`
	RoleAccountId     *string `json:",omitempty"`
	RoleName          *string `json:",omitempty"`
	RolePassAccountId *string `json:",omitempty"`
	RolePassName      *string `json:",omitempty"`
}

type URLNormalize struct {
	NormalizeObject []string
	Switch          *bool `json:",omitempty"`
}

type UpdateCdnConfigRequest struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	AutoRotate          *AutoRotate     `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	ConditionalOrigin   *ConditionalOrigin   `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	Domain              *string              `json:",omitempty"`
	DownloadSpeedLimit  *DownloadSpeedLimit  `json:",omitempty"`
	FollowRedirect      *bool                `json:",omitempty"`
	HTTPS               *HTTPS               `json:",omitempty"`
	HeaderLogging       *HeaderLog           `json:",omitempty"`
	HttpForcedRedirect  *HttpForcedRedirect  `json:",omitempty"`
	IPv6                *IPv6                `json:",omitempty"`
	IpAccessRule        *IpAccessRule        `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit         `json:",omitempty"`
	MassCompression     *MassCompression     `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule    `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginCertCheck     *OriginCertCheck    `json:",omitempty"`
	OriginHost          *string             `json:",omitempty"`
	OriginIPv6          *string             `json:",omitempty"`
	OriginProtocol      *string             `json:",omitempty"`
	OriginRange         *bool               `json:",omitempty"`
	OriginRetry         *OriginRetry        `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Quic                *Quic               `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestBlockRule    *RequestBlockRule   `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	RewriteHLS          *RewriteHLS          `json:",omitempty"`
	ServiceRegion       *string              `json:",omitempty"`
	ServiceType         *string              `json:",omitempty"`
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Sparrow             *Sparrow             `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UrlNormalize        *URLNormalize        `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
	WebpAdaptive        *WebpAdaptive        `json:",omitempty"`
	Websocket           *Websocket           `json:",omitempty"`
}

type UpdateCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UpdateResourceTagsRequest struct {
	ResourceTags []ResourceTag
	Resources    []string
}

type UpdateResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UriParamSup struct {
	JoinSymbol  *string `json:",omitempty"`
	SplitSymbol *string `json:",omitempty"`
	StartLevel  *int64  `json:",omitempty"`
	TermLevel   *int64  `json:",omitempty"`
}

type UrlDecrypt struct {
	SignedUrlAuthRules *SignedUrlAuthRules `json:",omitempty"`
	Switch             *bool               `json:",omitempty"`
}

type UsageReportsDetail struct {
	BillingCode       string
	BillingRegion     string
	CalculationMethod string
	CreateTime        int64
	DownLoadUrl       string
	EndTime           int64
	ExportType        string
	Metric            string
	StartTime         int64
	Status            int64
	TaskId            string
	TaskName          string
	TimeZone          string
}

type UserAgentAccessRule struct {
	AllowEmpty *bool   `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	RuleType   *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
	UserAgent  []string
}

type VideoDrag struct {
	Switch *bool `json:",omitempty"`
}

type WebpAdaptive struct {
	Switch *bool `json:",omitempty"`
}

type Websocket struct {
	Switch  *bool  `json:",omitempty"`
	Timeout *int64 `json:",omitempty"`
}
