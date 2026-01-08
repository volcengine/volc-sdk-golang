package tls

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type CommonResponse struct {
	RequestID string `json:"RequestId"`
}

type CommonRequest struct {
	Headers map[string]string `json:"-"`
}

type TagInfo struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type CreateProjectRequest struct {
	CommonRequest
	ProjectName    string    `json:","`
	Description    string    `json:","`
	Region         string    `json:","`
	IamProjectName *string   `json:",omitempty"`
	Tags           []TagInfo `json:",omitempty"`
}

func (v *CreateProjectRequest) CheckValidation() error {
	if len(v.ProjectName) <= 0 {
		return errors.New("Invalid argument, empty ProjectName")
	}
	if len(v.Region) <= 0 {
		return errors.New("Invalid argument, empty Region")
	}
	return nil
}

type CreateProjectResponse struct {
	CommonResponse
	ProjectID string `json:"ProjectId"`
}

type DeleteProjectRequest struct {
	CommonRequest
	ProjectID string `json:"ProjectId"`
}

func (v *DeleteProjectRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	return nil
}

type DescribeProjectRequest struct {
	CommonRequest
	ProjectID string `json:"ProjectId"`
}

func (v *DescribeProjectRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	return nil
}

type DescribeProjectResponse struct {
	CommonResponse
	ProjectInfo
}

type DescribeProjectsRequest struct {
	CommonRequest
	ProjectName    string
	ProjectID      string
	PageNumber     int
	PageSize       int
	IsFullName     bool
	IamProjectName *string
	Tags           []TagInfo
}

func (v *DescribeProjectsRequest) CheckValidation() error {
	return nil
}

type ProjectInfo struct {
	ProjectID       string    `json:"ProjectId"`
	ProjectName     string    `json:"ProjectName"`
	Description     string    `json:"Description"`
	CreateTimestamp string    `json:"CreateTime"`
	TopicCount      int64     `json:"TopicCount"`
	InnerNetDomain  string    `json:"InnerNetDomain"`
	IamProjectName  string    `json:"IamProjectName"`
	Tags            []TagInfo `json:"Tags"`
}

type DescribeProjectsResponse struct {
	CommonResponse
	Projects []ProjectInfo `json:"Projects"`
	Total    int64         `json:"Total"`
}

type ModifyProjectRequest struct {
	CommonRequest
	ProjectID   string
	ProjectName *string `json:",omitempty"`
	Description *string `json:",omitempty"`
}

func (v *ModifyProjectRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	return nil
}

type EncryptUserCmkConf struct {
	UserCmkID string `json:"user_cmk_id"`
	Trn       string `json:"trn"`
	RegionID  string `json:"region_id"`
}

type EncryptConf struct {
	Enable             bool                `json:"enable"`
	EncryptType        string              `json:"encrypt_type"`
	EncryptUserCmkConf *EncryptUserCmkConf `json:"user_cmk_info"`
}

type CreateTopicRequest struct {
	CommonRequest
	ProjectID      string       `json:"ProjectId"`
	TopicName      string       `json:","`
	Ttl            uint16       `json:","`
	Description    string       `json:","`
	ShardCount     int          `json:","`
	MaxSplitShard  *int32       `json:",omitempty"`
	AutoSplit      bool         `json:","`
	EnableTracking *bool        `json:",omitempty"`
	TimeKey        *string      `json:",omitempty"`
	TimeFormat     *string      `json:",omitempty"`
	Tags           []TagInfo    `json:",omitempty"`
	LogPublicIP    *bool        `json:",omitempty"`
	EnableHotTtl   *bool        `json:",omitempty"`
	HotTtl         *int32       `json:",omitempty"`
	ColdTtl        *int32       `json:",omitempty"`
	ArchiveTtl     *int32       `json:",omitempty"`
	EncryptConf    *EncryptConf `json:",omitempty"`
}

func (v *CreateTopicRequest) CheckValidation() error {
	if len(v.TopicName) <= 0 {
		return errors.New("Invalid argument, empty TopicName")
	}
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	if v.Ttl <= 0 {
		return errors.New("Invalid Ttl, must be bigger than 0")
	}
	if v.ShardCount <= 0 {
		return errors.New("Invalid ShardCount, must be bigger than 0")
	}
	return nil
}

type CreateTopicResponse struct {
	CommonResponse
	TopicID string `json:"TopicID"`
}

type DeleteTopicRequest struct {
	CommonRequest
	TopicID string
}

func (v *DeleteTopicRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type ModifyTopicRequest struct {
	CommonRequest
	TopicID        string       `json:"TopicId"`
	TopicName      *string      `json:",omitempty"`
	Ttl            *uint16      `json:",omitempty"`
	Description    *string      `json:",omitempty"`
	MaxSplitShard  *int32       `json:",omitempty"`
	AutoSplit      *bool        `json:",omitempty"`
	EnableTracking *bool        `json:",omitempty"`
	TimeKey        *string      `json:",omitempty"`
	TimeFormat     *string      `json:",omitempty"`
	LogPublicIP    *bool        `json:",omitempty"`
	EnableHotTtl   *bool        `json:",omitempty"`
	HotTtl         *int32       `json:",omitempty"`
	ColdTtl        *int32       `json:",omitempty"`
	ArchiveTtl     *int32       `json:",omitempty"`
	EncryptConf    *EncryptConf `json:",omitempty"`
}

func (v *ModifyTopicRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type DescribeTopicRequest struct {
	CommonRequest
	TopicID string
}

func (v *DescribeTopicRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type DescribeTopicResponse struct {
	CommonResponse
	TopicName       string      `json:"TopicName"`
	ProjectID       string      `json:"ProjectId"`
	TopicID         string      `json:"TopicId"`
	Ttl             uint16      `json:"Ttl"`
	CreateTimestamp string      `json:"CreateTime"`
	ModifyTimestamp string      `json:"ModifyTime"`
	ShardCount      int32       `json:"ShardCount"`
	Description     string      `json:"Description"`
	MaxSplitShard   int32       `json:"MaxSplitShard"`
	AutoSplit       bool        `json:"AutoSplit"`
	EnableTracking  bool        `json:"EnableTracking"`
	TimeKey         string      `json:"TimeKey"`
	TimeFormat      string      `json:"TimeFormat"`
	Tags            []TagInfo   `json:"Tags"`
	LogPublicIP     bool        `json:"LogPublicIP"`
	EnableHotTtl    bool        `json:"EnableHotTtl"`
	HotTtl          int32       `json:"HotTtl"`
	ColdTtl         int32       `json:"ColdTtl"`
	ArchiveTtl      int32       `json:"ArchiveTtl"`
	EncryptConf     EncryptConf `json:",omitempty"`
}

type DescribeTopicsRequest struct {
	CommonRequest
	ProjectID   string
	ProjectName string
	PageNumber  int
	PageSize    int
	TopicName   string
	TopicID     string
	IsFullName  bool
	Tags        []TagInfo
}

func (v *DescribeTopicsRequest) CheckValidation() error {
	return nil
}

type Topic struct {
	TopicName       string      `json:"TopicName"`
	ProjectID       string      `json:"ProjectId"`
	ProjectName     string      `json:"ProjectName"`
	TopicID         string      `json:"TopicId"`
	Ttl             uint16      `json:"Ttl"`
	CreateTimestamp string      `json:"CreateTime"`
	ModifyTimestamp string      `json:"ModifyTime"`
	ShardCount      int32       `json:"ShardCount"`
	Description     string      `json:"Description"`
	MaxSplitShard   int32       `json:"MaxSplitShard"`
	AutoSplit       bool        `json:"AutoSplit"`
	EnableTracking  bool        `json:"EnableTracking"`
	TimeKey         string      `json:"TimeKey"`
	TimeFormat      string      `json:"TimeFormat"`
	Tags            []TagInfo   `json:"Tags"`
	LogPublicIP     bool        `json:"LogPublicIP"`
	EnableHotTtl    bool        `json:"EnableHotTtl"`
	HotTtl          int32       `json:"HotTtl"`
	ColdTtl         int32       `json:"ColdTtl"`
	ArchiveTtl      int32       `json:"ArchiveTtl"`
	EncryptConf     EncryptConf `json:",omitempty"`
}

type DescribeTopicsResponse struct {
	CommonResponse
	Topics []*Topic `json:"Topics"`
	Total  int      `json:"Total"`
}

type Value struct {
	ValueType      string         `json:"ValueType"`
	Delimiter      string         `json:"Delimiter"`
	CaseSensitive  bool           `json:"CaseSensitive"`
	IncludeChinese bool           `json:"IncludeChinese"`
	SQLFlag        bool           `json:"SqlFlag"`
	JsonKeys       []KeyValueInfo `json:"JsonKeys"`
	IndexAll       bool           `json:"IndexAll"`
	AutoIndexFlag  *bool          `json:"AutoIndexFlag,omitempty"`
	IndexSQLAll    *bool          `json:"IndexSQLAll,omitempty"`
}

type KeyValueParam struct {
	Key   string `json:"key"`
	Value Value  `json:"value"`
}

type KeyValueList []KeyValueParam

func (c KeyValueList) Value() (driver.Value, error) {
	b, err := json.Marshal(c)

	return string(b), err
}

type CreateIndexRequest struct {
	CommonRequest
	TopicID           string          `json:"TopicId"`
	FullText          *FullTextInfo   `json:",omitempty"`
	KeyValue          *[]KeyValueInfo `json:",omitempty"`
	UserInnerKeyValue *[]KeyValueInfo `json:",omitempty"`
	MaxTextLen        *int32          `json:",omitempty"`
	EnableAutoIndex   *bool           `json:",omitempty"`
}

func (v *CreateIndexRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type FullTextInfo struct {
	Delimiter      string `json:"Delimiter"`
	CaseSensitive  bool   `json:"CaseSensitive"`
	IncludeChinese bool   `json:"IncludeChinese"`
}

type KeyValueInfo struct {
	Key   string `json:"Key"`
	Value Value  `json:"Value"`
}

type CreateIndexResponse struct {
	CommonResponse
	TopicID string `json:"TopicId"`
}

type DeleteIndexRequest struct {
	CommonRequest
	TopicID string
}

func (v *DeleteIndexRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type DescribeIndexRequest struct {
	CommonRequest
	TopicID string
}

func (v *DescribeIndexRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type DescribeIndexResponse struct {
	CommonResponse
	TopicID           string          `json:"TopicId"`
	FullText          *FullTextInfo   `json:"FullText"`
	KeyValue          *[]KeyValueInfo `json:"KeyValue"`
	UserInnerKeyValue *[]KeyValueInfo `json:"UserInnerKeyValue"`
	CreateTime        string          `json:"CreateTime"`
	ModifyTime        string          `json:"ModifyTime"`
	MaxTextLen        int32           `json:"MaxTextLen,omitempty"`
	EnableAutoIndex   bool            `json:"EnableAutoIndex,omitempty"`
}

type ModifyIndexRequest struct {
	CommonRequest
	TopicID           string          `json:"TopicId"`
	FullText          *FullTextInfo   `json:",omitempty"`
	KeyValue          *[]KeyValueInfo `json:",omitempty"`
	UserInnerKeyValue *[]KeyValueInfo `json:",omitempty"`
	MaxTextLen        *int32          `json:",omitempty"`
	EnableAutoIndex   *bool           `json:",omitempty"`
}

func (v *ModifyIndexRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type AnalysisResult struct {
	Schema []string                 `json:"Schema"`
	Type   map[string]string        `json:"Type"`
	Data   []map[string]interface{} `json:"Data"`
}

type SearchLogsRequest struct {
	CommonRequest
	TopicID   string `json:"TopicId"`
	Query     string `json:"Query"`
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Limit     int    `json:"Limit"`
	HighLight bool   `json:"HighLight"`
	Context   string `json:"Context"`
	Sort      string `json:"Sort"`
}

func (v *SearchLogsRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if v.EndTime < v.StartTime {
		return errors.New("Invalid argument EndTime < StartTime")
	}
	return nil
}

type SearchLogsResponse struct {
	CommonResponse
	Status         string                   `json:"ResultStatus"`
	Analysis       bool                     `json:"Analysis"`
	ListOver       bool                     `json:"ListOver"`
	HitCount       int                      `json:"HitCount"`
	Count          int                      `json:"Count"`
	Limit          int                      `json:"Limit"`
	Logs           []map[string]interface{} `json:"Logs"`
	AnalysisResult *AnalysisResult          `json:"AnalysisResult"`
	Context        string                   `json:"Context"`
	HighLight      []string                 `json:"HighLight,omitempty"`
}

type DescribeShardsRequest struct {
	CommonRequest
	TopicID    string
	PageNumber int
	PageSize   int
}

func (v *DescribeShardsRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type ShardInfo struct {
	TopicID           string `json:"TopicId"`
	ShardID           int32  `json:"ShardId"`
	InclusiveBeginKey string `json:"InclusiveBeginKey"`
	ExclusiveEndKey   string `json:"ExclusiveEndKey"`
	Status            string `json:"Status"`
	ModifyTimestamp   string `json:"ModifyTime"`
	StopWriteTime     string `json:"StopWriteTime"`
}

type DescribeShardsResponse struct {
	CommonResponse
	Shards []*ShardInfo `json:"Shards"`
	Total  int          `json:"Total"`
}

type ManualShardSplitRequest struct {
	CommonRequest
	TopicID string `json:"TopicId"`
	ShardID int    `json:"ShardId"`
	Number  int    `json:"Number"`
}

func (v *ManualShardSplitRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}

	switch v.Number {
	case 2, 4, 8, 16:
		// valid
	default:
		return errors.New("Invalid argument, Number must be one of 2, 4, 8, 16")
	}
	return nil
}

type ManualShardSplitResponse struct {
	CommonResponse
	Shards []*ShardInfo `json:"Shards"`
}

type PutLogsRequest struct {
	CommonRequest
	TopicID      string
	HashKey      string
	CompressType string
	LogBody      *pb.LogGroupList
}

func (v *PutLogsRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if v.LogBody == nil {
		return errors.New("Invalid argument, empty LogBody")
	}

	return nil
}

type DescribeCursorRequest struct {
	CommonRequest
	TopicID string
	ShardID int
	From    string
}

func (v *DescribeCursorRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if len(v.From) <= 0 {
		return errors.New("Invalid argument, empty From")
	}
	// ShardId is okay to be set 0, ignore
	return nil
}

type DescribeCursorResponse struct {
	CommonResponse
	Cursor string `json:"cursor"`
}

type ConsumeLogsRequest struct {
	CommonRequest
	TopicID       string
	ShardID       int
	Cursor        string
	Original      bool
	EndCursor     *string `json:",omitempty"`
	LogGroupCount *int    `json:",omitempty"`
	Compression   *string `json:",omitempty"`

	ConsumerGroupName *string `json:",omitempty"`
	ConsumerName      *string `json:",omitempty"`
}

func (v *ConsumeLogsRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type ConsumeLogsResponse struct {
	CommonResponse
	Cursor string //X-Tls-Cursor
	Count  int    //X-Tls-Count
	Logs   *pb.LogGroupList
}

type DescribeLogContextRequest struct {
	CommonRequest
	TopicId       string `json:"TopicId"`
	ContextFlow   string `json:"ContextFlow"`
	PackageOffset int64  `json:"PackageOffset"`
	Source        string `json:"Source"`
	PrevLogs      *int64 `json:",omitempty"`
	NextLogs      *int64 `json:",omitempty"`
}

func (v *DescribeLogContextRequest) CheckValidation() error {
	if len(v.TopicId) <= 0 {
		return errors.New("Invalid argument, empty TopicId")
	}
	if len(v.ContextFlow) <= 0 {
		return errors.New("Invalid arguemnt, empty ContextFlow")
	}
	// PackageOffset and Source is okay to be set 0, ignore.
	return nil
}

type DescribeLogContextResponse struct {
	CommonResponse
	LogContextInfos []map[string]interface{} `json:"LogContextInfos"`
	PrevOver        bool                     `json:"PrevOver"`
	NextOver        bool                     `json:"NextOver"`
}

type CreateRuleRequest struct {
	CommonRequest
	TopicID        string          `json:"TopicId"`
	RuleName       string          `json:"RuleName"`
	Paths          *[]string       `json:"Paths"`
	LogType        *string         `json:"LogType"`
	ExtractRule    *ExtractRule    `json:"ExtractRule"`
	ExcludePaths   *[]ExcludePath  `json:"ExcludePaths"`
	UserDefineRule *UserDefineRule `json:"UserDefineRule"`
	LogSample      *string         `json:"LogSample"`
	InputType      *int            `json:"InputType"`
	ContainerRule  *ContainerRule  `json:"ContainerRule"`
}

func (v *CreateRuleRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if len(v.RuleName) <= 0 {
		return errors.New("Invalid argument, empty RuleName")
	}
	return nil
}

type ExtractRule struct {
	Delimiter           string           `json:"Delimiter,omitempty"`
	BeginRegex          string           `json:"BeginRegex,omitempty"`
	LogRegex            string           `json:"LogRegex,omitempty"`
	Keys                []string         `json:"Keys,omitempty"`
	TimeKey             string           `json:"TimeKey,omitempty"`
	TimeFormat          string           `json:"TimeFormat,omitempty"`
	FilterKeyRegex      []FilterKeyRegex `json:"FilterKeyRegex,omitempty"`
	UnMatchUpLoadSwitch bool             `json:"UnMatchUpLoadSwitch"`
	UnMatchLogKey       string           `json:"UnMatchLogKey,omitempty"`
	LogTemplate         LogTemplate      `json:"LogTemplate,omitempty"`
	Quote               string           `json:"Quote,omitempty"`
	TimeZone            string           `json:"TimeZone,omitempty"`
	TimeExtractRegex    string           `json:"TimeExtractRegex,omitempty"`
	TimeSample          string           `json:"TimeSample,omitempty"`
	EnableNanosecond    *bool            `json:"EnableNanosecond,omitempty"`
}

type FilterKeyRegex struct {
	Key   string `json:"Key,omitempty"`
	Regex string `json:"Regex,omitempty"`
}

type LogTemplate struct {
	Type   string `json:"Type,omitempty"`
	Format string `json:"Format,omitempty"`
}

type ExcludePath struct {
	Type  string `json:"Type,omitempty"`
	Value string `json:"Value,omitempty"`
}

type UserDefineRule struct {
	ParsePathRule        *ParsePathRule    `json:"ParsePathRule,omitempty"`
	ShardHashKey         *ShardHashKey     `json:"ShardHashKey,omitempty"`
	EnableRawLog         bool              `json:"EnableRawLog,omitempty"`
	Fields               map[string]string `json:"Fields,omitempty"`
	Plugin               *Plugin           `json:"Plugin,omitempty"`
	Advanced             *Advanced         `json:"Advanced,omitempty"`
	TailFiles            bool              `json:"TailFiles,omitempty"`
	RawLogKey            string            `json:"RawLogKey,omitempty"`
	HostnameKey          string            `json:"HostnameKey,omitempty"`
	EnableHostname       *bool             `json:"EnableHostname,omitempty"`
	HostGroupLabelKey    string            `json:"HostGroupLabelKey,omitempty"`
	EnableHostGroupLabel *bool             `json:"EnableHostGroupLabel,omitempty"`
	TailSizeKb           *int64            `json:"TailSizeKb,omitempty"`
	IgnoreOlder          *int64            `json:"IgnoreOlder,omitempty"`
	MultiCollectsType    string            `json:"MultiCollectsType,omitempty"`
}

type ShardHashKey struct {
	HashKey string `json:"HashKey,omitempty"`
}

type ParsePathRule struct {
	PathSample string   `json:"PathSample,omitempty"`
	Regex      string   `json:"Regex,omitempty"`
	Keys       []string `json:"Keys,omitempty"`
}

type Plugin struct {
	Processors []map[string]interface{} `json:"processors,omitempty"`
}

type Advanced struct {
	CloseInactive              *int  `json:"CloseInactive,omitempty"`
	CloseRemoved               *bool `json:"CloseRemoved,omitempty"`
	CloseRenamed               *bool `json:"CloseRenamed,omitempty"`
	CloseEOF                   *bool `json:"CloseEOF,omitempty"`
	CloseTimeout               *int  `json:"CloseTimeout,omitempty"`
	NoLineTerminatorEOFMaxTime *int  `json:"NoLineTerminatorEOFMaxTime,omitempty"`
}

type ContainerRule struct {
	Stream                     string            `json:"Stream,omitempty"`
	ContainerNameRegex         string            `json:"ContainerNameRegex,omitempty"`
	IncludeContainerLabelRegex map[string]string `json:"IncludeContainerLabelRegex,omitempty"`
	ExcludeContainerLabelRegex map[string]string `json:"ExcludeContainerLabelRegex,omitempty"`
	IncludeContainerEnvRegex   map[string]string `json:"IncludeContainerEnvRegex,omitempty"`
	ExcludeContainerEnvRegex   map[string]string `json:"ExcludeContainerEnvRegex,omitempty"`
	EnvTag                     map[string]string `json:"EnvTag,omitempty"`
	KubernetesRule             KubernetesRule    `json:"KubernetesRule,omitempty"`
}

type KubernetesRule struct {
	NamespaceNameRegex        string            `json:"NamespaceNameRegex,omitempty"`
	WorkloadType              string            `json:"WorkloadType,omitempty"`
	WorkloadNameRegex         string            `json:"WorkloadNameRegex,omitempty"`
	PodNameRegex              string            `json:"PodNameRegex,omitempty"`
	IncludePodLabelRegex      map[string]string `json:"IncludePodLabelRegex,omitempty"`
	ExcludePodLabelRegex      map[string]string `json:"ExcludePodLabelRegex,omitempty"`
	LabelTag                  map[string]string `json:"LabelTag,omitempty"`
	AnnotationTag             map[string]string `json:"AnnotationTag,omitempty"`
	IncludePodAnnotationRegex map[string]string `json:"IncludePodAnnotationRegex,omitempty"`
	ExcludePodAnnotationRegex map[string]string `json:"ExcludePodAnnotationRegex,omitempty"`
	EnableAllLabelTag         *bool             `json:"EnableAllLabelTag,omitempty"`
}

type CreateRuleResponse struct {
	CommonResponse
	RuleID string `json:"RuleId"`
}

type DeleteRuleRequest struct {
	CommonRequest
	RuleID string `json:"RuleId"`
}

func (v *DeleteRuleRequest) CheckValidation() error {
	if len(v.RuleID) <= 0 {
		return errors.New("Invalid argument, empty RuleID")
	}
	return nil
}

type ModifyRuleRequest struct {
	CommonRequest
	RuleID         string          `json:"RuleId,omitempty"`
	RuleName       *string         `json:"RuleName,omitempty"`
	Paths          *[]string       `json:"Paths,omitempty"`
	LogType        *string         `json:"LogType,omitempty"`
	ExtractRule    *ExtractRule    `json:"ExtractRule,omitempty"`
	ExcludePaths   *[]ExcludePath  `json:"ExcludePaths,omitempty"`
	UserDefineRule *UserDefineRule `json:"UserDefineRule,omitempty"`
	LogSample      *string         `json:"LogSample,omitempty"`
	InputType      *int            `json:"InputType,omitempty"`
	ContainerRule  *ContainerRule  `json:"ContainerRule,omitempty"`
	Pause          *int            `json:"Pause,omitempty"`
}

func (v *ModifyRuleRequest) CheckValidation() error {
	if len(v.RuleID) <= 0 {
		return errors.New("Invalid argument, empty RuleID")
	}
	return nil
}

type DescribeRuleRequest struct {
	CommonRequest
	RuleID string `json:"RuleId"`
}

func (v *DescribeRuleRequest) CheckValidation() error {
	if len(v.RuleID) <= 0 {
		return errors.New("Invalid argument, empty RuleID")
	}
	return nil
}

type DescribeRuleResponse struct {
	CommonResponse
	ProjectID      string           `json:"ProjectId"`
	ProjectName    string           `json:"ProjectName"`
	TopicID        string           `json:"TopicId"`
	TopicName      string           `json:"TopicName"`
	RuleInfo       *RuleInfo        `json:"RuleInfo"`
	HostGroupInfos []*HostGroupInfo `json:"HostGroupInfos"`
}

type DescribeRuleRequestV2 struct {
	CommonRequest
	RuleID string `json:"RuleId"`
}

func (v *DescribeRuleRequestV2) CheckValidation() error {
	if len(v.RuleID) <= 0 {
		return errors.New("Invalid argument, empty RuleID")
	}
	return nil
}

type DescribeRuleResponseV2 struct {
	CommonResponse
	ProjectID   string    `json:"ProjectId"`
	ProjectName string    `json:"ProjectName"`
	TopicID     string    `json:"TopicId"`
	TopicName   string    `json:"TopicName"`
	RuleInfo    *RuleInfo `json:"RuleInfo"`
}

type DescribeBoundHostGroupsRequest struct {
	CommonRequest
	RuleID     string
	PageNumber int
	PageSize   int
}

func (v *DescribeBoundHostGroupsRequest) CheckValidation() error {
	if len(v.RuleID) <= 0 {
		return errors.New("Invalid argument, empty RuleID")
	}
	return nil
}

type DescribeBoundHostGroupsResponse struct {
	CommonResponse
	Total          int64            `json:"Total"`
	HostGroupInfos []*HostGroupInfo `json:"HostGroupInfos"`
}

type RuleInfo struct {
	TopicID        string         `json:"TopicId"`
	TopicName      string         `json:"TopicName"`
	RuleID         string         `json:"RuleId"`
	RuleName       string         `json:"RuleName"`
	Paths          []string       `json:"Paths"`
	LogType        string         `json:"LogType"`
	ExtractRule    ExtractRule    `json:"ExtractRule"`
	ExcludePaths   []ExcludePath  `json:"ExcludePaths"`
	UserDefineRule UserDefineRule `json:"UserDefineRule"`
	LogSample      string         `json:"LogSample"`
	InputType      int            `json:"InputType"`
	ContainerRule  ContainerRule  `json:"ContainerRule"`
	CreateTime     string         `json:"CreateTime"`
	ModifyTime     string         `json:"ModifyTime"`
	Pause          int            `json:"Pause"`
}

type HostGroupInfo struct {
	HostGroupID                   string `json:"HostGroupId"`
	HostGroupName                 string `json:"HostGroupName"`
	HostGroupType                 string `json:"HostGroupType"`
	HostIdentifier                string `json:"HostIdentifier"`
	HostCount                     int    `json:"HostCount"`
	NormalHeartbeatStatusNumber   int    `json:"NormalHeartbeatStatusCount"`
	AbnormalHeartbeatStatusNumber int    `json:"AbnormalHeartbeatStatusCount"`
	RuleCount                     int    `json:"RuleCount"`
	CreateTime                    string `json:"CreateTime"`
	ModifyTime                    string `json:"ModifyTime"`
	AutoUpdate                    bool   `json:"AutoUpdate"`
	UpdateStartTime               string `json:"UpdateStartTime"`
	UpdateEndTime                 string `json:"UpdateEndTime"`
	AgentLatestVersion            string `json:"AgentLatestVersion"`
	ServiceLogging                bool   `json:"ServiceLogging"`
	IamProjectName                string `json:"IamProjectName"`
}

type DescribeRulesRequest struct {
	CommonRequest
	ProjectID      string  `json:"ProjectId"`
	PageNumber     int     `json:"PageNumber"`
	PageSize       int     `json:"PageSize"`
	TopicID        *string `json:"TopicId,omitempty"`
	TopicName      *string `json:"TopicName,omitempty"`
	RuleID         *string `json:"RuleId,omitempty"`
	RuleName       *string `json:"RuleName,omitempty"`
	ProjectName    *string `json:"ProjectName,omitempty"`
	IamProjectName *string `json:"IamProjectName,omitempty"`
	LogType        *string `json:"LogType,omitempty"`
	Pause          *int    `json:"Pause,omitempty"`
}

func (v *DescribeRulesRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	return nil
}

type CreateHostGroupRequest struct {
	CommonRequest
	HostGroupName   string    `json:"HostGroupName"`
	HostGroupType   string    `json:"HostGroupType"`
	HostIdentifier  *string   `json:"HostIdentifier"`
	HostIPList      *[]string `json:"HostIpList"`
	AutoUpdate      *bool     `json:",omitempty"`
	UpdateStartTime *string   `json:",omitempty"`
	UpdateEndTime   *string   `json:",omitempty"`
	ServiceLogging  *bool     `json:",omitempty"`
	IamProjectName  *string   `json:",omitempty"`
}

func (v *CreateHostGroupRequest) CheckValidation() error {
	if len(v.HostGroupName) <= 0 {
		return errors.New("Invalid argument, empty HostGroupName")
	}
	if len(v.HostGroupType) <= 0 {
		return errors.New("Invalid argument, empty HostGroupType")
	}
	return nil
}

type CreateHostGroupResponse struct {
	CommonResponse
	HostGroupID string `json:"HostGroupId"`
}

type DescribeRulesResponse struct {
	CommonResponse
	Total     int64       `json:"Total"`
	RuleInfos []*RuleInfo `json:"RuleInfos"`
}

type DeleteHostGroupRequest struct {
	CommonRequest
	HostGroupID string `json:"HostGroupId"`
}

func (v *DeleteHostGroupRequest) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty HostGroupID")
	}
	return nil
}

type ApplyRuleToHostGroupsRequest struct {
	CommonRequest
	RuleID       string   `json:"RuleId"`
	HostGroupIDs []string `json:"HostGroupIds"`
}

func (v *ApplyRuleToHostGroupsRequest) CheckValidation() error {
	if len(v.RuleID) <= 0 {
		return errors.New("Invalid argument, empty RuleID")
	}
	if len(v.HostGroupIDs) <= 0 {
		return errors.New("Invalid argument, empty HostGroupIDs")
	}
	return nil
}

type DeleteRuleFromHostGroupsRequest struct {
	CommonRequest
	RuleID       string   `json:"RuleId"`
	HostGroupIDs []string `json:"HostGroupIds"`
}

func (v *DeleteRuleFromHostGroupsRequest) CheckValidation() error {
	if len(v.RuleID) <= 0 {
		return errors.New("Invalid argument, empty RuleID")
	}
	if len(v.HostGroupIDs) <= 0 {
		return errors.New("Invalid argument, empty HostGroupIDs")
	}
	return nil
}

type ModifyHostGroupRequest struct {
	CommonRequest
	HostGroupID     string    `json:"HostGroupId"`
	HostGroupName   *string   `json:"HostGroupName,omitempty"`
	HostGroupType   *string   `json:"HostGroupType,omitempty"`
	HostIPList      *[]string `json:"HostIpList,omitempty"`
	HostIdentifier  *string   `json:"HostIdentifier,omitempty"`
	AutoUpdate      *bool     `json:",omitempty"`
	UpdateStartTime *string   `json:",omitempty"`
	UpdateEndTime   *string   `json:",omitempty"`
	ServiceLogging  *bool     `json:",omitempty"`
}

func (v *ModifyHostGroupRequest) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty HostGroupID")
	}
	return nil
}

type DescribeHostGroupRequest struct {
	CommonRequest
	HostGroupID string
}

func (v *DescribeHostGroupRequest) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty HostGroupID")
	}
	return nil
}

type DescribeHostGroupResponse struct {
	CommonResponse
	HostGroupHostsRulesInfo *HostGroupHostsRulesInfo `json:"HostGroupHostsRulesInfo"`
}

type DescribeHostGroupRequestV2 struct {
	CommonRequest
	HostGroupID string
}

func (v *DescribeHostGroupRequestV2) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty HostGroupID")
	}
	return nil
}

type DescribeHostGroupResponseV2 struct {
	CommonResponse
	HostGroupHostsRulesInfo *HostGroupHostsRulesInfoV2 `json:"HostGroupHostsRulesInfo"`
}

type HostGroupHostsRulesInfo struct {
	HostGroupInfo *HostGroupInfo `json:"HostGroupInfo"`
	HostInfos     []*HostInfo    `json:"HostInfos"`
	RuleInfos     []*RuleInfo    `json:"RuleInfos"`
}

type HostInfo struct {
	IP                  string `json:"Ip"`
	LogCollectorVersion string `json:"LogCollectorVersion"`
	HeartbeatStatus     int    `json:"HeartbeatStatus"`
}

type DescribeHostGroupsRequest struct {
	CommonRequest
	PageNumber     int
	PageSize       int
	HostGroupID    *string `json:",omitempty"`
	HostGroupName  *string `json:",omitempty"`
	HostIdentifier *string `json:",omitempty"`
	AutoUpdate     *bool   `json:",omitempty"`
	ServiceLogging *bool   `json:",omitempty"`
	IamProjectName *string `json:",omitempty"`
}

func (v *DescribeHostGroupsRequest) CheckValidation() error {
	return nil
}

type DescribeHostGroupsResponse struct {
	CommonResponse
	Total                    int64                      `json:"Total"`
	HostGroupHostsRulesInfos []*HostGroupHostsRulesInfo `json:"HostGroupHostsRulesInfos"`
}

type DescribeHostGroupsRequestV2 struct {
	CommonRequest
	PageNumber     int
	PageSize       int
	HostGroupID    *string `json:",omitempty"`
	HostGroupName  *string `json:",omitempty"`
	HostIdentifier *string `json:",omitempty"`
	AutoUpdate     *bool   `json:",omitempty"`
	ServiceLogging *bool   `json:",omitempty"`
	IamProjectName *string `json:",omitempty"`
}

func (v *DescribeHostGroupsRequestV2) CheckValidation() error {
	return nil
}

type DescribeHostGroupsResponseV2 struct {
	CommonResponse
	Total                    int64                        `json:"Total"`
	HostGroupHostsRulesInfos []*HostGroupHostsRulesInfoV2 `json:"HostGroupHostsRulesInfos"`
}

// HostGroupHostsRulesInfoV2 展示层对象
type HostGroupHostsRulesInfoV2 struct {
	HostGroupInfo *HostGroupInfo `json:"HostGroupInfo"`
}

type DescribeHostsRequest struct {
	CommonRequest
	HostGroupID     string
	PageNumber      int
	PageSize        int
	IP              *string
	HeartbeatStatus *int
}

func (v *DescribeHostsRequest) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty HostGroupID")
	}
	return nil
}

type DescribeHostsResponse struct {
	CommonResponse
	Total     int64       `json:"Total"`
	HostInfos []*HostInfo `json:"HostInfos"`
}

type DeleteHostRequest struct {
	CommonRequest
	HostGroupID string `json:"HostGroupId"`
	IP          string `json:"Ip"`
}

func (v *DeleteHostRequest) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty Host")
	}
	if len(v.IP) <= 0 {
		return errors.New("Invalid arguemnt, empty IP")
	}
	return nil
}

type DescribeHostGroupRulesRequest struct {
	CommonRequest
	HostGroupID string
	PageNumber  int
	PageSize    int
}

func (v *DescribeHostGroupRulesRequest) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty HostGroupID")
	}
	return nil
}

type DescribeHostGroupRulesResponse struct {
	CommonResponse
	Total     int64       `json:"Total"`
	RuleInfos []*RuleInfo `json:"RuleInfos"`
}

type ModifyHostGroupsAutoUpdateRequest struct {
	CommonRequest
	HostGroupIds    []string
	AutoUpdate      *bool   `json:",omitempty"`
	UpdateStartTime *string `json:",omitempty"`
	UpdateEndTime   *string `json:",omitempty"`
}

func (v *ModifyHostGroupsAutoUpdateRequest) CheckValidation() error {
	if len(v.HostGroupIds) <= 0 {
		return errors.New("Invalid argument, empty HostGroupIds")
	}
	return nil
}

type ModifyHostGroupsAutoUpdateResponse struct {
	CommonResponse
}

type DeleteAbnormalHostsRequest struct {
	CommonRequest
	HostGroupID string `json:"HostGroupId"`
}

func (v *DeleteAbnormalHostsRequest) CheckValidation() error {
	if len(v.HostGroupID) <= 0 {
		return errors.New("Invalid argument, empty HostGroupId")
	}

	return nil
}

type AlarmPeriodSetting struct {
	Sms            int `json:"SMS"`
	Phone          int `json:"Phone"`
	Email          int `json:"Email"`
	GeneralWebhook int `json:"GeneralWebhook"`
}

type JoinConfig struct {
	Condition        string `json:"Condition"`
	SetOperationType string `json:"SetOperationType"`
}

type TriggerCondition struct {
	Severity       string `json:"Severity"`
	Condition      string `json:"Condition"`
	CountCondition string `json:"CountCondition"`
	NoData         *bool  `json:"NoData,omitempty"`
}

type CreateAlarmRequest struct {
	CommonRequest
	AlarmName          string              `json:"AlarmName"`
	ProjectID          string              `json:"ProjectId"`
	Status             *bool               `json:"Status,omitempty"`
	QueryRequest       QueryRequests       `json:"QueryRequest"`
	RequestCycle       RequestCycle        `json:"RequestCycle"`
	Condition          string              `json:"Condition"`
	TriggerPeriod      int                 `json:"TriggerPeriod"`
	AlarmPeriod        int                 `json:"AlarmPeriod"`
	AlarmNotifyGroup   []string            `json:"AlarmNotifyGroup"`
	UserDefineMsg      *string             `json:"UserDefineMsg,omitempty"`
	Severity           *string             `json:"Severity,omitempty"`
	AlarmPeriodDetail  *AlarmPeriodSetting `json:"AlarmPeriodDetail,omitempty"`
	JoinConfigurations []JoinConfig        `json:"JoinConfigurations,omitempty"`
	TriggerConditions  []TriggerCondition  `json:"TriggerConditions,omitempty"`
	SendResolved       *bool               `json:"SendResolved,omitempty"`
}

func (v *CreateAlarmRequest) CheckValidation() error {
	if len(v.AlarmName) <= 0 {
		return errors.New("Invalid argument, empty AlarmName")
	}
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	if len(v.QueryRequest) <= 0 {
		return errors.New("Invalid argument, empty QueryRequest")
	}
	if len(v.RequestCycle.Type) <= 0 {
		return errors.New("Invalid argument, empty RequestCycle.Type")
	}
	if v.RequestCycle.Time <= 0 {
		return errors.New("Invalid argument, RequestCycle.Time <= 0")
	}
	if v.AlarmPeriod <= 0 {
		return errors.New("Invalid argument, empty AlarmPeriod")
	}
	if len(v.AlarmNotifyGroup) <= 0 {
		return errors.New("Invalid argument, empty AlarmNotifyGroup")
	}
	if len(v.Condition) <= 0 {
		return errors.New("Invalid argument, empty Condition")
	}
	return nil
}

type QueryRequests []QueryRequest

type RequestCycle struct {
	Type         string  `json:"Type"`
	Time         int     `json:"Time"`
	CronTab      string  `json:"CronTab"`
	CronTimeZone *string `json:"CronTimeZone,omitempty"`
}

type QueryRequest struct {
	CommonRequest
	Query               string `json:"Query"`
	Number              uint8  `json:"Number"`
	TopicID             string `json:"TopicId"`
	TopicName           string `json:"TopicName,omitempty"`
	StartTimeOffset     int    `json:"StartTimeOffset"`
	EndTimeOffset       int    `json:"EndTimeOffset"`
	TimeSpanType        string `json:"TimeSpanType,omitempty"`
	TruncatedTime       string `json:"TruncatedTime,omitempty"`
	EndTimeOffsetUnit   string `json:"EndTimeOffsetUnit,omitempty"`
	StartTimeOffsetUnit string `json:"StartTimeOffsetUnit,omitempty"`
}

type CreateAlarmResponse struct {
	CommonResponse
	AlarmID string `json:"AlarmId"`
}

type DeleteAlarmRequest struct {
	CommonRequest
	AlarmID string `json:"AlarmId"`
}

func (v *DeleteAlarmRequest) CheckValidation() error {
	if len(v.AlarmID) <= 0 {
		return errors.New("Invalid argument, empty AlarmID")
	}
	return nil
}

type DeleteAlarmWebhookIntegrationRequest struct {
	CommonRequest
	WebhookID string `json:"WebhookID"`
}

func (v *DeleteAlarmWebhookIntegrationRequest) CheckValidation() error {
	if len(v.WebhookID) <= 0 {
		return errors.New("Invalid argument, empty WebhookID")
	}
	return nil
}

type DeleteAlarmContentTemplateRequest struct {
	CommonRequest
	AlarmContentTemplateID string `json:"AlarmContentTemplateId"`
}

func (v *DeleteAlarmContentTemplateRequest) CheckValidation() error {
	if len(v.AlarmContentTemplateID) <= 0 {
		return errors.New("Invalid argument, empty AlarmContentTemplateID")
	}
	return nil
}

type ModifyAlarmRequest struct {
	CommonRequest
	AlarmID            string              `json:"AlarmId"`
	AlarmName          *string             `json:"AlarmName"`
	Status             *bool               `json:"Status"`
	QueryRequest       *QueryRequests      `json:"QueryRequest"`
	RequestCycle       *RequestCycle       `json:"RequestCycle"`
	Condition          *string             `json:"Condition"`
	TriggerPeriod      *int                `json:"TriggerPeriod"`
	AlarmPeriod        *int                `json:"AlarmPeriod"`
	AlarmNotifyGroup   *[]string           `json:"AlarmNotifyGroup"`
	UserDefineMsg      *string             `json:"UserDefineMsg"`
	Severity           *string             `json:"Severity,omitempty"`
	AlarmPeriodDetail  *AlarmPeriodSetting `json:"AlarmPeriodDetail,omitempty"`
	JoinConfigurations []JoinConfig        `json:"JoinConfigurations,omitempty"`
	TriggerConditions  []TriggerCondition  `json:"TriggerConditions,omitempty"`
	SendResolved       *bool               `json:"SendResolved,omitempty"`
}

func (v *ModifyAlarmRequest) CheckValidation() error {
	if len(v.AlarmID) <= 0 {
		return errors.New("Invalid argument, empty AlarmID")
	}
	return nil
}

type ModifyAlarmWebhookIntegrationReq struct {
	CommonRequest
	WebhookID      string                   `json:"WebhookID"`
	WebhookName    string                   `json:"WebhookName"`
	WebhookType    string                   `json:"WebhookType"`
	WebhookUrl     string                   `json:"WebhookUrl"`
	WebhookMethod  *string                  `json:"WebhookMethod,omitempty"`
	WebhookSecret  *string                  `json:"WebhookSecret,omitempty"`
	WebhookHeaders []GeneralWebhookHeaderKV `json:"WebhookHeaders,omitempty"`
}

func (v *ModifyAlarmWebhookIntegrationReq) CheckValidation() error {
	if len(v.WebhookID) <= 0 {
		return errors.New("Invalid argument, empty WebhookID")
	}
	if len(v.WebhookName) <= 0 {
		return errors.New("Invalid argument, empty WebhookName")
	}
	if len(v.WebhookType) <= 0 {
		return errors.New("Invalid argument, empty WebhookType")
	}
	if len(v.WebhookUrl) <= 0 {
		return errors.New("Invalid argument, empty WebhookUrl")
	}
	return nil
}

type DescribeAlarmsRequest struct {
	CommonRequest
	ProjectID      string
	ProjectName    *string
	TopicID        *string
	TopicName      *string
	AlarmName      *string
	AlarmPolicyID  *string
	Status         *bool
	AlarmDisabled  *bool
	Severity       *string
	IamProjectName *string
	PageNumber     int
	PageSize       int
}

func (v *DescribeAlarmsRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	return nil
}

type DescribeAlarmsResponse struct {
	CommonResponse
	Total         int         `json:"Total"`
	AlarmPolicies []QueryResp `json:"Alarms"`
}

type QueryResp struct {
	AlarmID            string             `json:"AlarmId"`
	AlarmName          string             `json:"AlarmName"`
	ProjectID          string             `json:"ProjectId"`
	Status             bool               `json:"Status"`
	QueryRequest       []QueryRequest     `json:"QueryRequest"`
	RequestCycle       RequestCycle       `json:"RequestCycle"`
	Condition          string             `json:"Condition"`
	TriggerPeriod      int                `json:"TriggerPeriod"`
	AlarmPeriod        int                `json:"AlarmPeriod"`
	AlarmNotifyGroup   []NotifyGroupsInfo `json:"AlarmNotifyGroup"`
	UserDefineMsg      string             `json:"UserDefineMsg"`
	CreateTimestamp    string             `json:"CreateTime"`
	ModifyTimestamp    string             `json:"ModifyTime"`
	Severity           string             `json:"Severity"`
	AlarmPeriodDetail  AlarmPeriodSetting `json:"AlarmPeriodDetail"`
	JoinConfigurations []JoinConfig       `json:"JoinConfigurations"`
	TriggerConditions  []TriggerCondition `json:"TriggerConditions"`
	SendResolved       *bool              `json:"SendResolved,omitempty"`
}

type NotifyGroupsInfo struct {
	GroupName       string       `json:"AlarmNotifyGroupName"`
	NotifyGroupID   string       `json:"AlarmNotifyGroupId"`
	NoticeType      NoticeTypes  `json:"NotifyType"`
	Receivers       Receivers    `json:"Receivers"`
	CreateTimestamp string       `json:"CreateTime"`
	ModifyTimestamp string       `json:"ModifyTime"`
	IamProjectName  string       `json:"IamProjectName"`
	NoticeRules     []NoticeRule `json:"NoticeRules,omitempty"`
}

type NoticeTypes []NoticeType

type NoticeType string

type Receivers []Receiver

type Receiver struct {
	ReceiverType                ReveiverType             `json:"ReceiverType"`
	ReceiverNames               []string                 `json:"ReceiverNames"`
	ReceiverChannels            []ReceiverChannel        `json:"ReceiverChannels"`
	StartTime                   string                   `json:"StartTime"`
	EndTime                     string                   `json:"EndTime"`
	Webhook                     string                   `json:",omitempty"`
	GeneralWebhookUrl           *string                  `json:"GeneralWebhookUrl,omitempty"`
	GeneralWebhookBody          *string                  `json:"GeneralWebhookBody,omitempty"`
	AlarmWebhookAtUsers         []string                 `json:"AlarmWebhookAtUsers,omitempty"`
	AlarmWebhookIsAtAll         *bool                    `json:"AlarmWebhookIsAtAll,omitempty"`
	AlarmWebhookAtGroups        []string                 `json:"AlarmWebhookAtGroups,omitempty"`
	GeneralWebhookMethod        *string                  `json:"GeneralWebhookMethod,omitempty"`
	GeneralWebhookHeaders       []GeneralWebhookHeaderKV `json:"GeneralWebhookHeaders,omitempty"`
	AlarmContentTemplateId      *string                  `json:"AlarmContentTemplateId,omitempty"`
	AlarmWebhookIntegrationId   *string                  `json:"AlarmWebhookIntegrationId,omitempty"`
	AlarmWebhookIntegrationName *string                  `json:"AlarmWebhookIntegrationName,omitempty"`
}

type ReveiverType string

type ReceiverChannel string

type NoticeRules []NoticeRule

type NoticeRule struct {
	HasNext       *bool      `json:"HasNext,omitempty"`
	RuleNode      *RuleNode  `json:"RuleNode,omitempty"`
	HasEndNode    *bool      `json:"HasEndNode,omitempty"`
	ReceiverInfos *Receivers `json:"ReceiverInfos,omitempty"`
}

type RuleNode struct {
	Type     string     `json:"Type,omitempty"`
	Value    []string   `json:"Value,omitempty"`
	Children []RuleNode `json:"Children,omitempty"`
}

type GeneralWebhookHeaderKV struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type CreateAlarmWebhookIntegrationRequest struct {
	CommonRequest
	WebhookHeaders *[]GeneralWebhookHeaderKV `json:"WebhookHeaders,omitempty"`
	WebhookMethod  *string                   `json:"WebhookMethod,omitempty"`
	WebhookName    string                    `json:"WebhookName"`
	WebhookSecret  *string                   `json:"WebhookSecret,omitempty"`
	WebhookType    string                    `json:"WebhookType"`
	WebhookUrl     string                    `json:"WebhookUrl"`
}

func (v *CreateAlarmWebhookIntegrationRequest) CheckValidation() error {
	if len(v.WebhookName) <= 0 {
		return errors.New("Invalid argument, empty WebhookName")
	}
	if len(v.WebhookType) <= 0 {
		return errors.New("Invalid argument, empty WebhookType")
	}
	if len(v.WebhookUrl) <= 0 {
		return errors.New("Invalid argument, empty WebhookUrl")
	}
	if v.WebhookType == WebhookTypeGeneralWebhook {
		if v.WebhookHeaders == nil || len(*v.WebhookHeaders) == 0 {
			return errors.New("Invalid argument, empty WebhookHeaders")
		}
	}
	return nil
}

type CreateAlarmWebhookIntegrationResponse struct {
	CommonResponse
	AlarmWebhookIntegrationID string `json:"AlarmWebhookIntegrationId"`
}

type CreateAlarmNotifyGroupRequest struct {
	CommonRequest
	GroupName      string        `json:"AlarmNotifyGroupName"`
	NoticeType     NoticeTypes   `json:"NotifyType"`
	Receivers      Receivers     `json:"Receivers"`
	IamProjectName *string       `json:"IamProjectName,omitempty"`
	NoticeRules    *[]NoticeRule `json:"NoticeRules,omitempty"`
}

func (v *CreateAlarmNotifyGroupRequest) CheckValidation() error {
	if len(v.GroupName) <= 0 {
		return errors.New("Invalid argument, empty GroupName")
	}

	// Check mutual exclusivity: if NoticeRules is configured, NotifyType and Receivers must be empty
	if v.NoticeRules != nil && len(*v.NoticeRules) > 0 {
		if len(v.NoticeType) > 0 || len(v.Receivers) > 0 {
			return errors.New("Invalid argument, if NoticeRules is configured, NotifyType and Receivers must be empty")
		}
	} else {
		// If NoticeRules is empty, NotifyType and Receivers must be configured
		if len(v.NoticeType) <= 0 {
			return errors.New("Invalid argument, empty NotifyType")
		}
		if len(v.Receivers) <= 0 {
			return errors.New("Invalid argument, empty Receivers")
		}
	}
	return nil
}

type CreateAlarmNotifyGroupResponse struct {
	CommonResponse
	NotifyGroupID string `json:"AlarmNotifyGroupId"`
}

type DeleteAlarmNotifyGroupRequest struct {
	CommonRequest
	AlarmNotifyGroupID string `json:"AlarmNotifyGroupId"`
}

func (v *DeleteAlarmNotifyGroupRequest) CheckValidation() error {
	if len(v.AlarmNotifyGroupID) <= 0 {
		return errors.New("AlarmNotifyGroupID")
	}
	return nil
}

type DeleteAlarmNotifyGroupResponse struct {
}

type ModifyAlarmNotifyGroupRequest struct {
	CommonRequest
	AlarmNotifyGroupID   string       `json:"AlarmNotifyGroupId"`
	AlarmNotifyGroupName *string      `json:"AlarmNotifyGroupName,omitempty"`
	NoticeType           *NoticeTypes `json:"NotifyType,omitempty"`
	Receivers            *Receivers   `json:"Receivers,omitempty"`
	NoticeRules          *NoticeRules `json:"NoticeRules,omitempty"`
}

func (v *ModifyAlarmNotifyGroupRequest) CheckValidation() error {
	if len(v.AlarmNotifyGroupID) <= 0 {
		return errors.New("Invalid argument, empty AlarmNotifyGroupID")
	}
	return nil
}

type DescribeAlarmNotifyGroupsRequest struct {
	CommonRequest
	GroupName      *string
	NotifyGroupID  *string
	UserName       *string
	PageNumber     int
	PageSize       int
	IamProjectName *string
}

func (v *DescribeAlarmNotifyGroupsRequest) CheckValidation() error {
	return nil
}

type DescribeAlarmNotifyGroupsResponse struct {
	CommonResponse
	Total             int64               `json:"Total"`
	AlarmNotifyGroups []*NotifyGroupsInfo `json:"AlarmNotifyGroups"`
}

// DescribeAlarmWebhookIntegrations
type DescribeAlarmWebhookIntegrationsRequest struct {
	CommonRequest
	WebhookID   *string
	WebhookName *string
	WebhookType *string
	PageNumber  int
	PageSize    int
}

func (v *DescribeAlarmWebhookIntegrationsRequest) CheckValidation() error {
	return nil
}

type WebhookIntegrationInfo struct {
	WebhookID      string                   `json:"WebhookID"`
	CreateTime     string                   `json:"CreateTime"`
	ModifyTime     string                   `json:"ModifyTime"`
	WebhookUrl     string                   `json:"WebhookUrl"`
	WebhookName    string                   `json:"WebhookName"`
	WebhookType    string                   `json:"WebhookType"`
	WebhookMethod  *string                  `json:"WebhookMethod,omitempty"`
	WebhookSecret  *string                  `json:"WebhookSecret,omitempty"`
	WebhookHeaders []GeneralWebhookHeaderKV `json:"WebhookHeaders,omitempty"`
}

type DescribeAlarmWebhookIntegrationsResponse struct {
	CommonResponse
	Total               int                      `json:"Total"`
	WebhookIntegrations []WebhookIntegrationInfo `json:"WebhookIntegrations"`
}

type CreateDownloadTaskRequest struct {
	CommonRequest
	TopicID         string `json:"TopicId"`
	TaskName        string
	Query           string
	StartTime       int64
	EndTime         int64
	Compression     string
	DataFormat      string
	Limit           int
	Sort            string
	TaskType        int                          `json:"TaskType"`
	AllowIncomplete *bool                        `json:"AllowIncomplete,omitempty"`
	LogContextInfos *DownloadTaskLogContextInfos `json:"LogContextInfos,omitempty"`
}

type DownloadTaskLogContextInfos struct {
	Source        *string `json:"Source"`
	ContextFlow   *string `json:"ContextFlow"`
	PackageOffset *int64  `json:"PackageOffset"`
}

func (v *CreateDownloadTaskRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if len(v.TaskName) <= 0 {
		return errors.New("Invalid argument, empty TaskName")
	}
	if v.EndTime < v.StartTime {
		return errors.New("Invalid argument, EndTime < StartTime")
	}
	if len(v.DataFormat) <= 0 {
		return errors.New("Invalid argument, empty DataFormat")
	}
	if len(v.Sort) <= 0 {
		return errors.New("Invalid argument, empty Sort")
	}
	if v.Limit <= 0 {
		return errors.New("Invalid argument, empty Limit")
	}
	if len(v.Compression) <= 0 {
		return errors.New("Invalid argument, empty Compression")
	}
	return nil
}

type CreateDownloadTaskResponse struct {
	CommonResponse
	TaskId string
}

type DescribeDownloadTasksRequest struct {
	CommonRequest
	TopicID    string  `json:"-"`
	PageNumber *int    `json:"-"`
	PageSize   *int    `json:"-"`
	TaskName   *string `json:"-"`
}

func (v *DescribeDownloadTasksRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("TopicID")
	}
	return nil
}

type DownloadTaskResp struct {
	TaskId          string                       `json:"TaskId"`
	TaskName        string                       `json:"TaskName"`
	TopicId         string                       `json:"TopicId"`
	Query           string                       `json:"Query"`
	StartTime       string                       `json:"StartTime"`
	EndTime         string                       `json:"EndTime"`
	LogCount        int64                        `json:"LogCount"`
	LogSize         int64                        `json:"LogSize"`
	Compression     string                       `json:"Compression"`
	DataFormat      string                       `json:"DataFormat"`
	TaskStatus      string                       `json:"TaskStatus"`
	CreateTime      string                       `json:"CreateTime"`
	TaskType        int                          `json:"TaskType"`
	AllowIncomplete bool                         `json:"AllowIncomplete"`
	LogContextInfos *DownloadTaskLogContextInfos `json:"LogContextInfos,omitempty"`
}
type DescribeDownloadTasksResponse struct {
	CommonResponse
	Tasks []*DownloadTaskResp `json:"Tasks"`
	Total int64               `json:"Total"`
}

type DescribeDownloadUrlRequest struct {
	CommonRequest
	TaskId string `json:"-"`
}

func (v *DescribeDownloadUrlRequest) CheckValidation() error {
	if len(v.TaskId) <= 0 {
		return errors.New("Invalid argument, empty TaskId")
	}
	return nil
}

type DescribeDownloadUrlResponse struct {
	CommonResponse
	DownloadUrl string
}

type CancelDownloadTaskRequest struct {
	CommonRequest
	TaskId string `json:"TaskId"`
}

func (v *CancelDownloadTaskRequest) CheckValidation() error {
	if len(v.TaskId) <= 0 {
		return errors.New("Invalid argument, empty TaskId")
	}
	return nil
}

type CancelDownloadTaskResponse struct {
	CommonResponse
}

type WebTracksRequest struct {
	CommonRequest
	TopicID      string `json:"-"`
	ProjectID    string `json:"-"`
	CompressType string `json:"-"`
	Logs         []map[string]string
	Source       string `json:"Source,omitempty"`
}

func (v *WebTracksRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if len(v.Logs) > 10000 {
		return errors.New("Invalid argument, log count more than 10000 in a single logGroup")
	}
	return nil
}

type WebTracksResponse struct {
	CommonResponse
}

type HistogramInfo struct {
	Count int64 `json:"Count"`
	Time  int64 `json:"Time"`
}
type DescribeHistogramRequest struct {
	CommonRequest
	TopicID   string `json:"TopicId"`
	Query     string `json:","`
	StartTime int64  `json:","`
	EndTime   int64  `json:","`
	Interval  *int64 `json:",omitempty"`
}

func (v *DescribeHistogramRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if v.EndTime < v.StartTime {
		return errors.New("Invalid argument, EndTime < StartTime")
	}
	// Query is okay to be empty, ignore.
	return nil
}

type DescribeHistogramResponse struct {
	CommonResponse
	HistogramInfos []HistogramInfo `json:"Histogram"`
	Interval       int64           `json:"Interval"`
	TotalCount     int64           `json:"TotalCount"`
	ResultStatus   string          `json:"ResultStatus"`
}

type HistogramInfoV1 struct {
	Count        int64  `json:"Count"`
	EndTime      int64  `json:"EndTime"`
	StartTime    int64  `json:"StartTime"`
	ResultStatus string `json:"ResultStatus"`
}

type DescribeHistogramV1Request struct {
	CommonRequest
	TopicID   string `json:"TopicId"`
	Query     string `json:"Query"`
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Interval  int64  `json:"Interval,omitempty"`
}

func (v *DescribeHistogramV1Request) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if v.EndTime < v.StartTime {
		return errors.New("Invalid argument, EndTime < StartTime")
	}
	// Query is okay to be empty, ignore.
	return nil
}

type DescribeHistogramV1Response struct {
	CommonResponse
	HistogramInfos []HistogramInfoV1 `json:"Histogram"`
	ResultStatus   string            `json:"ResultStatus"`
	TotalCount     int64             `json:"TotalCount"`
}

type OpenKafkaConsumerRequest struct {
	CommonRequest
	TopicID string `json:"TopicId"`
}

func (v *OpenKafkaConsumerRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type OpenKafkaConsumerResponse struct {
	CommonResponse
}
type CloseKafkaConsumerRequest struct {
	CommonRequest
	TopicID string `json:"TopicId"`
}

func (v *CloseKafkaConsumerRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type CloseKafkaConsumerResponse struct {
	CommonResponse
}
type DescribeKafkaConsumerRequest struct {
	CommonRequest
	TopicID string `json:"-"`
}

func (v *DescribeKafkaConsumerRequest) CheckValidation() error {
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	return nil
}

type DescribeKafkaConsumerResponse struct {
	CommonResponse
	AllowConsume bool
	ConsumeTopic string
}

type CreateConsumerGroupRequest struct {
	ProjectID         string   `json:"ProjectID"`
	TopicIDList       []string `json:"TopicIDList"`
	ConsumerGroupName string   `json:"ConsumerGroupName"`
	HeartbeatTTL      int      `json:"HeartbeatTTL"`
	OrderedConsume    bool     `json:"OrderedConsume"`
}

func (v *CreateConsumerGroupRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	if len(v.TopicIDList) <= 0 {
		return errors.New("Invalid argument, empty TopicIDList")
	}
	if len(v.ConsumerGroupName) <= 0 {
		return errors.New("Invalid argument, empty ConsumerGroupName")
	}
	return nil
}

type CreateConsumerGroupResponse struct {
	CommonResponse
}

type DeleteConsumerGroupRequest struct {
	ProjectID         string `json:"ProjectID"`
	ConsumerGroupName string `json:"ConsumerGroupName"`
}

type DeleteConsumerGroupResponse struct {
	CommonResponse
}

type DescribeConsumerGroupsRequest struct {
	ProjectID         string  `json:"ProjectID"`
	ProjectName       string  `json:"ProjectName"`
	ConsumerGroupName string  `json:"ConsumerGroupName"`
	TopicID           string  `json:"TopicID"`
	TopicName         *string `json:"TopicName,omitempty"`
	IamProjectName    *string `json:"IamProjectName,omitempty"`
	PageNumber        int
	PageSize          int
}

type ConsumerGroupResp struct {
	ProjectID         string   `json:"ProjectID"`
	ProjectName       string   `json:"ProjectName"`
	TopicIDList       []string `json:"topic_id"`
	ConsumerGroupName string   `json:"ConsumerGroupName"`
	HeartbeatTTL      int      `json:"HeartbeatTTL"`
	OrderedConsume    bool     `json:"OrderedConsume"`
}

type DescribeConsumerGroupsResponse struct {
	CommonResponse
	ConsumerGroups []*ConsumerGroupResp `json:"ConsumerGroups"`
	Total          int                  `json:"Total"`
	DashboardID    string               `json:"DashboardId"`
}

type ModifyConsumerGroupRequest struct {
	ProjectID         string    `json:"ProjectID"`
	TopicIDList       *[]string `json:"TopicIDList"`
	ConsumerGroupName string    `json:"ConsumerGroupName"`
	HeartbeatTTL      *int      `json:"HeartbeatTTL"`
	OrderedConsume    *bool     `json:"OrderedConsume"`
}

type ModifyConsumerGroupResponse struct {
	CommonResponse
}

type ConsumerHeartbeatRequest struct {
	ProjectID         string `json:"ProjectID"`
	ConsumerGroupName string `json:"ConsumerGroupName"`
	ConsumerName      string `json:"ConsumerName"`
}

type ConsumeShard struct {
	TopicID string `yaml:"TopicID" json:"TopicID"`
	ShardID int    `yaml:"ShardID" json:"ShardID"`
}

type ConsumerHeartbeatResponse struct {
	CommonResponse
	Shards []*ConsumeShard `json:"Shards"`
}

type DescribeCheckPointRequest struct {
	ProjectID         string `json:"ProjectID"`
	TopicID           string `json:"TopicID"`
	ConsumerGroupName string `json:"ConsumerGroupName"`
	ShardID           int    `json:"ShardID"`
}

type DescribeCheckPointResponse struct {
	CommonResponse
	ShardID    int32  `json:"ShardID"`
	Checkpoint string `json:"Checkpoint"`
	UpdateTime int    `json:"UpdateTime"`
	Consumer   string `json:"Consumer"`
}

type ModifyCheckPointRequest struct {
	ProjectID         string `json:"ProjectID"`
	TopicID           string `json:"TopicID"`
	ConsumerGroupName string `json:"ConsumerGroupName"`
	ShardID           int    `json:"ShardID"`
	Checkpoint        string `json:"Checkpoint"`
}

type ModifyCheckPointResponse struct {
	CommonResponse
}

type ResetCheckPointRequest struct {
	ProjectID         string `json:"ProjectID"`
	ConsumerGroupName string `json:"ConsumerGroupName"`
	Position          string `json:"Position"`
}

func (v *ResetCheckPointRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	if len(v.ConsumerGroupName) <= 0 {
		return errors.New("Invalid argument, empty ConsumerGroupName")
	}
	if len(v.Position) <= 0 {
		return errors.New("Invalid argument, empty Position")
	}

	return nil
}

type AddTagsToResourceRequest struct {
	CommonRequest
	ResourceType  string    `json:"ResourceType"`
	ResourcesList []string  `json:"ResourcesList"`
	Tags          []TagInfo `json:"Tags"`
}

func (v *AddTagsToResourceRequest) CheckValidation() error {
	if len(v.ResourceType) <= 0 {
		return errors.New("Invalid argument, empty ResourceType")
	}
	if v.ResourcesList == nil || len(v.ResourcesList) == 0 {
		return errors.New("Invalid argument, empty ResourceList")
	}
	if v.Tags == nil || len(v.Tags) == 0 {
		return errors.New("Invalid argument, empty Tags")
	}

	return nil
}

type RemoveTagsFromResourceRequest struct {
	CommonRequest
	ResourceType  string   `json:"ResourceType"`
	ResourcesList []string `json:"ResourcesList"`
	TagKeyList    []string `json:"TagKeyList"`
}

func (v *RemoveTagsFromResourceRequest) CheckValidation() error {
	if len(v.ResourceType) <= 0 {
		return errors.New("Invalid argument, empty ResourceType")
	}
	if v.ResourcesList == nil || len(v.ResourcesList) == 0 {
		return errors.New("Invalid argument, empty ResourceList")
	}
	if v.TagKeyList == nil || len(v.TagKeyList) == 0 {
		return errors.New("Invalid argument, empty TagKeyList")
	}

	return nil
}

type UntagResourcesRequest struct {
	CommonRequest
	ResourceType  string   `json:"ResourceType"`
	ResourcesList []string `json:"ResourcesIds"`
	TagKeyList    []string `json:"TagKeys"`
}

func (v *UntagResourcesRequest) CheckValidation() error {
	if len(v.ResourceType) <= 0 {
		return errors.New("Invalid argument, empty ResourceType")
	}
	if v.ResourcesList == nil || len(v.ResourcesList) == 0 {
		return errors.New("Invalid argument, empty ResourceList")
	}
	if v.TagKeyList == nil || len(v.TagKeyList) == 0 {
		return errors.New("Invalid argument, empty TagKeyList")
	}

	return nil
}

type TagResourcesRequest struct {
	CommonRequest
	ResourceType string    `json:"ResourceType"`
	ResourcesIds []string  `json:"ResourcesIds"`
	Tags         []TagInfo `json:"Tags"`
}

func (v *TagResourcesRequest) CheckValidation() error {
	if len(v.ResourceType) <= 0 {
		return errors.New("Invalid argument, empty ResourceType")
	}
	if v.ResourcesIds == nil || len(v.ResourcesIds) == 0 {
		return errors.New("Invalid argument, empty ResourcesIds")
	}
	if len(v.ResourcesIds) > 50 {
		return errors.New("Invalid argument, ResourcesIds count exceeds 50")
	}
	if v.Tags == nil || len(v.Tags) == 0 {
		return errors.New("Invalid argument, empty Tags")
	}
	return nil
}

type TagResourcesResponse struct {
	CommonResponse
}

func (response *CommonResponse) FillRequestId(httpResponse *http.Response) {
	response.RequestID = httpResponse.Header.Get(RequestIDHeader)
}

type ModifyScheduleSqlTaskRequest struct {
	CommonRequest
	TaskId            string        `json:"TaskId"`
	TaskName          *string       `json:"TaskName,omitempty"`
	Description       *string       `json:"Description,omitempty"`
	DestRegion        *string       `json:"DestRegion,omitempty"`
	DestTopicID       *string       `json:"DestTopicID,omitempty"`
	Status            *int          `json:"Status,omitempty"`
	ProcessSqlDelay   *int          `json:"ProcessSqlDelay,omitempty"`
	ProcessTimeWindow *string       `json:"ProcessTimeWindow,omitempty"`
	Query             *string       `json:"Query,omitempty"`
	RequestCycle      *RequestCycle `json:"RequestCycle,omitempty"`
}

func (v *ModifyScheduleSqlTaskRequest) CheckValidation() error {
	if len(v.TaskId) <= 0 {
		return errors.New("Invalid argument, empty TaskId")
	}
	return nil
}

type ModifyScheduleSqlTaskResponse struct {
	CommonResponse
}

type LogContent struct {
	Key   string
	Value string
}
type Log struct {
	Contents []LogContent
	Time     int64 // 可以不填, 默认使用当前时间
}
type PutLogsV2Request struct {
	CommonRequest
	TopicID      string
	HashKey      string
	CompressType string
	Source       string
	FileName     string
	Logs         []Log
}

type BackendConfig struct {
	TTL                uint16 `json:"Ttl,omitempty"`
	MaxSplitPartitions int32  `json:"MaxSplitPartitions,omitempty"`
	AutoSplit          bool   `json:"AutoSplit,omitempty"`
	EnableHotTTL       bool   `json:"EnableHotTtl,omitempty"`
	HotTTL             uint16 `json:"HotTtl,omitempty"`
	ColdTTL            uint16 `json:"ColdTtl,omitempty"`
	ArchiveTTL         uint16 `json:"ArchiveTtl,omitempty"`
}

type ModifyTraceInstanceRequest struct {
	CommonRequest
	TraceInstanceId string         `json:"TraceInstanceId"`
	Description     *string        `json:"Description,omitempty"`
	BackendConfig   *BackendConfig `json:"BackendConfig,omitempty"`
}

func (v *ModifyTraceInstanceRequest) CheckValidation() error {
	if len(v.TraceInstanceId) <= 0 {
		return errors.New("Invalid argument, empty TraceInstanceId")
	}
	return nil
}

type ModifyTraceInstanceResponse struct {
	CommonResponse
}

type TargetResource struct {
	Alias   string  `json:"Alias"`
	TopicID string  `json:"TopicId"`
	Region  string  `json:"Region"`
	RoleTrn *string `json:"RoleTrn,omitempty"`
}

type CreateETLTaskRequest struct {
	CommonRequest
	DSLType         string
	Description     *string `json:",omitempty"`
	Enable          bool
	FromTime        *int64 `json:",omitempty"`
	Name            string
	Script          string
	SourceTopicId   string
	TargetResources []TargetResource
	TaskType        string
	ToTime          *int64 `json:",omitempty"`
}

type CreateETLTaskResponse struct {
	CommonResponse
	TaskID string `json:"TaskId"`
}

type DeleteETLTaskRequest struct {
	CommonRequest
	TaskID string `json:"TaskId"`
}

type ModifyETLTaskRequest struct {
	CommonRequest
	TaskID          string           `json:"TaskId"`
	Description     *string          `json:",omitempty"`
	Name            *string          `json:",omitempty"`
	Script          *string          `json:",omitempty"`
	TargetResources []TargetResource `json:"TargetResources,omitempty"`
}

type DescribeETLTaskRequest struct {
	CommonRequest
	TaskID string `json:"TaskId"`
}

type TargetResourcesResp struct {
	Alias       string
	TopicID     string `json:"TopicId"`
	TopicName   string
	ProjectName string
	ProjectID   string `json:"ProjectId"`
	Region      string
	RoleTrn     string
}

type ETLTaskResponse struct {
	CreateTime      string
	DSLType         string
	TaskID          string `json:"TaskId"`
	Description     string
	ETLStatus       string
	Enable          bool
	FromTime        *int64
	LastEnableTime  string
	ModifyTime      string
	Name            string
	ProjectID       string `json:"ProjectId"`
	ProjectName     string
	Script          string
	SourceTopicID   string `json:"SourceTopicId"`
	SourceTopicName string
	TargetResources []TargetResourcesResp
	TaskType        string
	ToTime          *int64
}

type DescribeETLTaskResponse struct {
	CommonResponse
	ETLTaskResponse
}

type DescribeETLTasksRequest struct {
	CommonRequest
	ProjectID       string `json:"ProjectId"`
	ProjectName     string
	IamProjectName  string
	SourceTopicID   string `json:"SourceTopicId"`
	SourceTopicName string
	TaskID          string `json:"TaskId"`
	TaskName        string
	Status          string
	PageNumber      int
	PageSize        int
}

type DescribeETLTasksResponse struct {
	CommonResponse
	Tasks []ETLTaskResponse
	Total int
}

type ModifyETLTaskStatusRequest struct {
	CommonRequest
	TaskID string `json:"TaskId"`
	Enable bool   `json:"Enable"`
}

type TosSourceInfo struct {
	Bucket       string `json:"bucket,omitempty"`
	Prefix       string `json:"prefix,omitempty"`
	Region       string `json:"region,omitempty"`
	CompressType string `json:"compress_type,omitempty"`
}

type KafkaSourceInfo struct {
	Host              string `json:"host,omitempty"`
	Group             string `json:"group,omitempty"`
	Topic             string `json:"topic,omitempty"`
	Encode            string `json:"encode,omitempty"`
	Password          string `json:"password,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Username          string `json:"username,omitempty"`
	Mechanism         string `json:"mechanism,omitempty"`
	InstanceID        string `json:"instance_id,omitempty"`
	IsNeedAuth        bool   `json:"is_need_auth,omitempty"`
	InitialOffset     int    `json:"initial_offset,omitempty"`
	TimeSourceDefault int    `json:"time_source_default,omitempty"`
}

type ImportSourceInfo struct {
	*TosSourceInfo   `json:"TosSourceInfo,omitempty"`
	*KafkaSourceInfo `json:"KafkaSourceInfo,omitempty"`
}

type ImportExtractRule struct {
	TimeZone         string `json:"TimeZone,omitempty"`
	SkipLineCount    int    `json:"SkipLineCount,omitempty"`
	TimeExtractRegex string `json:"TimeExtractRegex,omitempty"`
	*ExtractRule
}

type TargetInfo struct {
	Region         string             `json:"Region"`
	LogType        string             `json:"LogType"`
	LogSample      string             `json:"LogSample"`
	ExtractRule    *ImportExtractRule `json:"ExtractRule,omitempty"`
	UserDefineRule *UserDefineRule    `json:"UserDefineRule,omitempty"`
}

type CreateImportTaskRequest struct {
	CommonRequest
	ProjectID        string            `json:"ProjectID"`
	TopicID          string            `json:"TopicID"`
	TaskName         string            `json:"TaskName"`
	SourceType       string            `json:"SourceType"`
	Description      string            `json:"Description,omitempty"`
	ImportSourceInfo *ImportSourceInfo `json:"ImportSourceInfo"`
	TargetInfo       *TargetInfo       `json:"TargetInfo"`
}

func (v *CreateImportTaskRequest) CheckValidation() error {
	return nil
}

type CreateImportTaskResponse struct {
	CommonResponse
	TaskID string `json:"TaskId"`
}

type DeleteImportTaskRequest struct {
	CommonRequest
	TaskID string `json:"TaskId"`
}

type DeleteImportTaskResponse struct {
	CommonResponse
}

type ModifyImportTaskRequest struct {
	CommonRequest
	ProjectID        string            `json:"ProjectID,omitempty"`
	TopicID          string            `json:"TopicID"`
	TaskID           string            `json:"TaskId"`
	TaskName         string            `json:"TaskName"`
	Status           int               `json:"Status"`
	SourceType       string            `json:"SourceType"`
	Description      *string           `json:"Description,omitempty"`
	ImportSourceInfo *ImportSourceInfo `json:"ImportSourceInfo"`
	TargetInfo       *TargetInfo       `json:"TargetInfo"`
}

type ModifyImportTaskResponse struct {
	CommonResponse
}

type DescribeImportTaskRequest struct {
	CommonRequest
	TaskID string `json:"TaskId,omitempty"`
}

type TaskStatistics struct {
	Total            int    `json:"Total"`
	Failed           int    `json:"Failed"`
	Skipped          int    `json:"Skipped"`
	NotExist         int    `json:"NotExist"`
	BytesTotal       int    `json:"BytesTotal"`
	TaskStatus       string `json:"TaskStatus"`
	Transferred      int    `json:"Transferred"`
	BytesTransferred int    `json:"BytesTransferred"`
}

type ImportTaskInfo struct {
	TaskID           string            `json:"TaskId"`
	Status           int               `json:"status"`
	TopicID          string            `json:"TopicId"`
	TaskName         string            `json:"TaskName"`
	ProjectID        string            `json:"ProjectId"`
	TopicName        string            `json:"TopicName"`
	CreateTime       string            `json:"CreateTime"`
	SourceType       string            `json:"SourceType"`
	Description      string            `json:"Description"`
	ProjectName      string            `json:"ProjectName"`
	TargetInfo       *TargetInfo       `json:"TargetInfo"`
	TaskStatistics   *TaskStatistics   `json:"TaskStatistics"`
	ImportSourceInfo *ImportSourceInfo `json:"ImportSourceInfo"`
}

type DescribeImportTaskResponse struct {
	CommonResponse
	TaskInfo *ImportTaskInfo `json:"TaskInfo"`
}

type DescribeImportTasksRequest struct {
	CommonRequest
	TaskID         string `json:"TaskId,omitempty"`
	TaskName       string `json:"TaskName,omitempty"`
	ProjectID      string `json:"ProjectId,omitempty"`
	ProjectName    string `json:"ProjectName,omitempty"`
	IamProjectName string `json:"IamProjectName,omitempty"`
	TopicID        string `json:"TopicId,omitempty"`
	TopicName      string `json:"TopicName,omitempty"`
	PageNumber     int    `json:"PageNumber,omitempty"`
	PageSize       int    `json:"PageSize,omitempty"`
	SourceType     string `json:"SourceType,omitempty"`
	Status         string `json:"Status,omitempty"`
}

type DescribeImportTasksResponse struct {
	CommonResponse
	TaskInfo []*ImportTaskInfo `json:"TaskInfo"`
	Total    int               `json:"Total"`
}

type ActiveTlsAccountRequest struct {
	CommonRequest
}

func (v *ActiveTlsAccountRequest) CheckValidation() error {
	return nil
}

type ActiveTlsAccountResponse struct {
	CommonResponse
}

type DescribeTraceInstanceRequest struct {
	CommonRequest
	TraceInstanceId string `json:"TraceInstanceId"`
}

func (v *DescribeTraceInstanceRequest) CheckValidation() error {
	if len(v.TraceInstanceId) <= 0 {
		return errors.New("Invalid argument, empty TraceInstanceId")
	}
	return nil
}

type DescribeTraceInstanceResponse struct {
	CommonResponse
	CreateTime               string              `json:"CreateTime"`
	DependencyTopicId        string              `json:"DependencyTopicId"`
	DependencyTopicTopicName string              `json:"DependencyTopicTopicName"`
	Description              string              `json:"Description"`
	ModifyTime               string              `json:"ModifyTime"`
	ProjectId                string              `json:"ProjectId"`
	ProjectName              string              `json:"ProjectName"`
	TraceInstanceId          string              `json:"TraceInstanceId"`
	TraceInstanceName        string              `json:"TraceInstanceName"`
	TraceInstanceStatus      TraceInstanceStatus `json:"TraceInstanceStatus"`
	TraceTopicId             string              `json:"TraceTopicId"`
	TraceTopicName           string              `json:"TraceTopicName"`
}

// SearchTraces related models translated from NodeJS SDK

type TraceQueryParameters struct {
	Asc           *bool             `json:"Asc,omitempty"`
	Kind          *string           `json:"Kind,omitempty"`
	Limit         *int              `json:"Limit,omitempty"`
	Order         *string           `json:"Order,omitempty"`
	Offset        *int              `json:"Offset,omitempty"`
	TraceId       *string           `json:"TraceId,omitempty"`
	Attributes    map[string]string `json:"Attributes,omitempty"`
	StatusCode    *string           `json:"StatusCode,omitempty"`
	DurationMax   *int              `json:"DurationMax,omitempty"`
	DurationMin   *int              `json:"DurationMin,omitempty"`
	ServiceName   *string           `json:"ServiceName,omitempty"`
	StartTimeMax  *int              `json:"StartTimeMax,omitempty"`
	StartTimeMin  *int              `json:"StartTimeMin,omitempty"`
	OperationName *string           `json:"OperationName,omitempty"`
}

type TraceInfo struct {
	EndTime       int               `json:"EndTime"`
	TraceId       string            `json:"TraceId"`
	Duration      int               `json:"Duration"`
	StartTime     int               `json:"StartTime"`
	Attributes    map[string]string `json:"Attributes"`
	StatusCode    string            `json:"StatusCode"`
	ServiceName   string            `json:"ServiceName"`
	OperationName string            `json:"OperationName"`
}

type SearchTracesRequest struct {
	CommonRequest
	Query           TraceQueryParameters `json:"Query"`
	TraceInstanceId string               `json:"TraceInstanceId"`
}

func (v *SearchTracesRequest) CheckValidation() error {
	if len(v.TraceInstanceId) <= 0 {
		return errors.New("Invalid argument, empty TraceInstanceId")
	}
	return nil
}

type SearchTracesResponse struct {
	CommonResponse
	Total      int         `json:"Total"`
	TraceInfos []TraceInfo `json:"TraceInfos"`
}

type CsvInfo struct {
	Keys            []string `json:"Keys"`
	Delimiter       string   `json:"Delimiter"`
	EscapeChar      string   `json:"EscapeChar"`
	PrintHeader     bool     `json:"PrintHeader"`
	NonFieldContent string   `json:"NonFieldContent"`
}

type JsonInfo struct {
	Keys   []string `json:"Keys,omitempty"`
	Enable bool     `json:"Enable"`
	Escape bool     `json:"Escape,omitempty"`
}

type ContentInfo struct {
	Format   string    `json:"Format,omitempty"`
	CsvInfo  *CsvInfo  `json:"CsvInfo,omitempty"`
	JsonInfo *JsonInfo `json:"JsonInfo,omitempty"`
}

type TosShipperInfo struct {
	Bucket          string `json:"Bucket,omitempty"`
	Prefix          string `json:"Prefix,omitempty"`
	MaxSize         int    `json:"MaxSize,omitempty"`
	Compress        string `json:"Compress,omitempty"`
	Interval        int    `json:"Interval,omitempty"`
	PartitionFormat string `json:"PartitionFormat,omitempty"`
}

type KafkaShipperInfo struct {
	Instance   string `json:"Instance"`
	Compress   string `json:"Compress"`
	KafkaTopic string `json:"KafkaTopic"`
	StartTime  int    `json:"StartTime,omitempty"`
	EndTime    int    `json:"EndTime,omitempty"`
}

type CreateShipperRequest struct {
	CommonRequest
	TopicId          string            `json:"TopicId"`
	ShipperName      string            `json:"ShipperName"`
	ShipperType      string            `json:"ShipperType"`
	ShipperStartTime int               `json:"ShipperStartTime,omitempty"`
	ShipperEndTime   int               `json:"ShipperEndTime,omitempty"`
	RoleTrn          string            `json:"RoleTrn,omitempty"`
	TosShipperInfo   *TosShipperInfo   `json:"TosShipperInfo,omitempty"`
	KafkaShipperInfo *KafkaShipperInfo `json:"KafkaShipperInfo,omitempty"`
	ContentInfo      *ContentInfo      `json:"ContentInfo"`
}

func (v *CreateShipperRequest) CheckValidation() error {
	return nil
}

type CreateShipperResponse struct {
	CommonResponse
	ShipperId string `json:"ShipperId"`
}

type DeleteShipperRequest struct {
	CommonRequest
	ShipperId string `json:"ShipperId"`
}

type DeleteShipperResponse struct {
	CommonResponse
}

type ModifyShipperRequest struct {
	CommonRequest
	ShipperId        string            `json:"ShipperId"`
	ShipperName      string            `json:"ShipperName,omitempty"`
	ShipperType      string            `json:"ShipperType"`
	Status           *bool             `json:"Status,omitempty"`
	RoleTrn          string            `json:"RoleTrn,omitempty"`
	ContentInfo      *ContentInfo      `json:"ContentInfo,omitempty"`
	TosShipperInfo   *TosShipperInfo   `json:"TosShipperInfo,omitempty"`
	KafkaShipperInfo *KafkaShipperInfo `json:"KafkaShipperInfo,omitempty"`
}

func (v *ModifyShipperRequest) CheckValidation() error {
	return nil
}

type ModifyShipperResponse struct {
	CommonResponse
}

type DescribeShipperRequest struct {
	CommonRequest
	ShipperId string `json:"ShipperId"`
}

type DescribeShipperResponse struct {
	CommonResponse
	ProjectId        string            `json:"ProjectId,omitempty"`
	ProjectName      string            `json:"ProjectName,omitempty"`
	TopicId          string            `json:"TopicId,omitempty"`
	TopicName        string            `json:"TopicName,omitempty"`
	ShipperId        string            `json:"ShipperId,omitempty"`
	ShipperName      string            `json:"ShipperName,omitempty"`
	ShipperType      string            `json:"ShipperType,omitempty"`
	DashboardId      string            `json:"DashboardId,omitempty"`
	CreateTime       string            `json:"CreateTime,omitempty"`
	ModifyTime       string            `json:"ModifyTime,omitempty"`
	ShipperStartTime int               `json:"ShipperStartTime,omitempty"`
	ShipperEndTime   int               `json:"ShipperEndTime,omitempty"`
	Status           bool              `json:"Status,omitempty"`
	RoleTrn          string            `json:"RoleTrn,omitempty"`
	ContentInfo      *ContentInfo      `json:"ContentInfo,omitempty"`
	TosShipperInfo   *TosShipperInfo   `json:"TosShipperInfo,omitempty"`
	KafkaShipperInfo *KafkaShipperInfo `json:"KafkaShipperInfo,omitempty"`
}

type DescribeShippersRequest struct {
	CommonRequest
	IamProjectName string `json:"IamProjectName,omitempty"`
	ProjectId      string `json:"ProjectId,omitempty"`
	ProjectName    string `json:"ProjectName,omitempty"`
	TopicId        string `json:"TopicId,omitempty"`
	TopicName      string `json:"TopicName,omitempty"`
	ShipperId      string `json:"ShipperId,omitempty"`
	ShipperName    string `json:"ShipperName,omitempty"`
	ShipperType    string `json:"ShipperType,omitempty"`
	PageNumber     int    `json:"PageNumber,omitempty"`
	PageSize       int    `json:"PageSize,omitempty"`
}

type DescribeShippersResponse struct {
	CommonResponse
	Shippers []*DescribeShipperResponse `json:"Shippers"`
	Total    int                        `json:"Total"`
}

type Intent int32

const (
	IntentText2Tls               Intent = 0
	IntentTls2Text               Intent = 1
	IntentDocChat                Intent = 2
	IntentSyntaxErrorFix         Intent = 3
	IntentSyntaxErrorExplanation Intent = 4
	IntentUnknown                Intent = 20
)

type IntentInfo struct {
	Name   string `json:"Name,omitempty"`
	Reason string `json:"Reason,omitempty"`
	Type   Intent `json:"Type,omitempty"`
}

type DescribeSessionAnswerReq struct {
	CommonRequest
	// 用户创建的Ai assistant实例id
	InstanceId string `json:"InstanceId" binding:"required"`
	// TopicId
	TopicId string `json:"TopicId" binding:"required"`
	// 对话id
	SessionId string `json:"SessionId" binding:"required"`
	// 用户问题
	Question string `json:"Question" binding:"required"`
	// QuestionId 问题的id,重新回答时使用
	QuestionId string `json:"QuestionId,omitempty"`
	// 用户意图
	Intent *Intent `json:"Intent,omitempty"`
}

const (
	MaxQuestionLength = 2000
)

func (v *DescribeSessionAnswerReq) CheckValidation() error {
	if len(v.InstanceId) == 0 {
		return errors.New("invalid argument, empty InstanceId")
	}
	if len(v.TopicId) == 0 {
		return errors.New("invalid argument, empty TopicId")
	}
	if len(v.SessionId) == 0 {
		return errors.New("invalid argument, empty SessionId")
	}
	if len(v.Question) == 0 || len(v.Question) > MaxQuestionLength {
		return errors.New("invalid argument Question")
	}

	return nil
}

type SessionMessageType string

const (
	CopilotProgress SessionMessageType = "progress"
	CopilotMessage  SessionMessageType = "message"
	CopilotError    SessionMessageType = "error"
)

type SessionResponseMessage struct {
	// 问题id
	QuestionId string `json:"QuestionId"`
	SessionId  string `json:"SessionId"`
	MessageId  string `json:"MessageId"`
	Answer     string `json:"Answer"`
	// 是否通过校验,未通过前端需处理
	PassDetect bool `json:"PassDetect"`
}

type AgentRspMsgType int32

const (
	// AgentRspMsgTypeIntentRecognition 意图识别
	AgentRspMsgTypeIntentRecognition AgentRspMsgType = 0
	// AgentRspMsgTypeToolCalling 工具调用
	AgentRspMsgTypeToolCalling AgentRspMsgType = 1
	// AgentRspMsgTypeInference 实际回答
	AgentRspMsgTypeInference AgentRspMsgType = 2
	// AgentRspMsgTypeQuestionsSuggestions 问题建议
	AgentRspMsgTypeQuestionsSuggestions AgentRspMsgType = 3
	// AgentRspMsgTypeRetrieval 语料召回
	AgentRspMsgTypeRetrieval AgentRspMsgType = 4
	// AgentRspMsgTypeReasoning 思考过程
	AgentRspMsgTypeReasoning AgentRspMsgType = 5
)

type ToolCallingInfo struct {
	Name string `json:"Name,omitempty"`
}

type KnowledgeRetrieval struct {
	// 参考的文档
	Documents []string `json:"Documents,omitempty"`
}

type ModelAnswer struct {
	Answer     string `json:"Answer,omitempty"`
	PassDetect bool   `json:"PassDetect,omitempty"`
}

type Stage struct {
	NodeName    string `json:"NodeName,omitempty"`
	NodeContent string `json:"NodeContent,omitempty"`
}

type SSEMessage struct {
	Event string
	Data  DescribeSessionAnswerResp
}

type CopilotAnswer struct {
	QuestionId string `json:"QuestionId"`
	SessionId  string `json:"SessionId"`
	// 模型回答的messageId
	MessageId string `json:"MessageId"`

	// 如果SessionMessageType为Message，此字段用于指示消息类型
	RspMsgType AgentRspMsgType `json:"RspMsgType,omitempty"`
	// 模型回答部分
	ModelAnswer *ModelAnswer `json:"ModelAnswer,omitempty"`
	// 意图识别细节
	IntentInfo *IntentInfo `json:"IntentInfo,omitempty"`
	// 工具调用细节
	Tools *ToolCallingInfo `json:"Tools,omitempty"`
	// 知识召回细节
	RetrievalInfo *KnowledgeRetrieval `json:"RetrievalInfo,omitempty"`
	// 关联问题建议
	Suggestions []string `json:"Suggestions,omitempty"`

	// 深度思考部分
	ReasoningContent *ModelAnswer `json:"ReasoningContent,omitempty"`
	// 上屏展示的工作流推进进度
	StageInfo *Stage `json:"FlowStage,omitempty"`
}

type DescribeSessionAnswerResp struct {
	// SSE报文类型
	ConversationMessageType SessionMessageType `json:"ConversationMessageType" enums:"progress,message,error"`
	// 模型回答的消息内容
	Message *SessionResponseMessage `json:"Message,omitempty"`
	// 如果SessionMessageType为Message，此字段用于指示消息类型
	RspMsgType AgentRspMsgType `json:"RspMsgType,omitempty"`
	// 意图识别细节
	IntentInfo *IntentInfo `json:"IntentInfo,omitempty"`
	// 工具调用细节
	Tools *ToolCallingInfo `json:"Tools,omitempty"`
	// 知识召回细节
	RetrievalInfo *KnowledgeRetrieval `json:"RetrievalInfo,omitempty"`
	// 关联问题建议
	Suggestions []string `json:"Suggestions,omitempty"`

	// 深度思考部分
	ReasoningContent *ModelAnswer `json:"ReasoningContent,omitempty"`
	// 上屏展示的工作流推进进度
	StageInfo *Stage `json:"FlowStage,omitempty"`

	QuestionId string `json:"QuestionId"`
	SessionId  string `json:"SessionId"`
	// 模型回答的messageId
	MessageId string `json:"MessageId"`
}

// DingTalkContentTemplateInfo 钉钉通知内容模版
type DingTalkContentTemplateInfo struct {
	// 告警通知内容的主题
	Title *string `json:"Title,omitempty"`
	// 告警通知中固定内容的语言，可选值为 zh-CN、en-US
	Locale *string `json:"Locale,omitempty"`
	// 告警通知内容，支持普通文本格式，支持插入内容变量、内容函数等。变量渲染后的通知内容长度最长为 8 KB，超过限制长度后会被截断。正文留空，表示使用默认内容
	Content *string `json:"Content,omitempty"`
}

// EmailContentTemplateInfo 邮件通知内容模版
type EmailContentTemplateInfo struct {
	// 告警通知中固定内容的语言，可选值为 zh-CN、en-US
	Locale *string `json:"Locale,omitempty"`
	// 告警通知内容，支持普通文本格式，支持插入内容变量、内容函数等。变量渲染后的通知内容长度最长为 8 KB，超过限制长度后会被截断。正文留空，表示使用默认内容
	Content *string `json:"Content,omitempty"`
	// 邮件通知的主题
	Subject *string `json:"Subject,omitempty"`
}

// LarkContentTemplateInfo 飞书通知内容模版
type LarkContentTemplateInfo struct {
	// 告警通知内容的主题
	Title *string `json:"Title,omitempty"`
	// 告警通知中固定内容的语言，可选值为 zh-CN、en-US
	Locale *string `json:"Locale,omitempty"`
	// 告警通知内容，支持普通文本格式，支持插入内容变量、内容函数等。变量渲染后的通知内容长度最长为 8 KB，超过限制长度后会被截断。正文留空，表示使用默认内容
	Content *string `json:"Content,omitempty"`
}

// SmsContentTemplateInfo 短信通知内容模版
type SmsContentTemplateInfo struct {
	// 告警通知中固定内容的语言，可选值为 zh-CN、en-US
	Locale *string `json:"Locale,omitempty"`
	// 告警通知内容，支持普通文本格式，支持插入内容变量、内容函数等。建议将变量渲染后的通知内容长度控制在 256 个字符以内，超过限制长度后可能会被截断。正文留空，表示使用默认内容
	Content *string `json:"Content,omitempty"`
}

// VmsContentTemplateInfo 语音通知内容模版
type VmsContentTemplateInfo struct {
	// 告警通知中固定内容的语言，可选值为 zh-CN、en-US
	Locale *string `json:"Locale,omitempty"`
	// 告警通知内容，支持普通文本格式，支持插入内容变量、内容函数等。建议将变量渲染后的通知内容长度控制在 256 个字符以内，超过限制长度后可能会被截断。正文留空，表示使用默认内容
	Content *string `json:"Content,omitempty"`
}

// WeChatContentTemplateInfo 企业微信通知内容模版
type WeChatContentTemplateInfo struct {
	// 告警通知内容的主题
	Title *string `json:"Title,omitempty"`
	// 告警通知中固定内容的语言，可选值为 zh-CN、en-US
	Locale *string `json:"Locale,omitempty"`
	// 告警通知内容，支持普通文本格式，支持插入内容变量、内容函数等。变量渲染后的通知内容长度最长为 8 KB，超过限制长度后会被截断。正文留空，表示使用默认内容
	Content *string `json:"Content,omitempty"`
}

// WebhookContentTemplateInfo 自定义的 Webhook 告警通知内容模版
type WebhookContentTemplateInfo struct {
	// 告警通知内容，通常为 JSON 格式，支持插入内容变量、内容函数等。变量渲染后的通知内容长度最长为 16 KB，超过限制长度后会被截断。正文留空，表示使用默认内容
	Content *string `json:"Content,omitempty"`
}

// CreateAlarmContentTemplateRequest 创建告警内容模版请求
type CreateAlarmContentTemplateRequest struct {
	CommonRequest
	// 告警通知内容模版的名称。命名规则请参考资源命名规则
	AlarmContentTemplateName string `json:"AlarmContentTemplateName"`
	// 钉钉通知内容模版
	DingTalk *DingTalkContentTemplateInfo `json:"DingTalk,omitempty"`
	// 邮件通知内容模版
	Email *EmailContentTemplateInfo `json:"Email,omitempty"`
	// 飞书通知内容模版
	Lark *LarkContentTemplateInfo `json:"Lark,omitempty"`
	// 是否需要校验内容模版
	NeedValidContent *bool `json:"NeedValidContent,omitempty"`
	// 短信通知内容模版
	Sms *SmsContentTemplateInfo `json:"Sms,omitempty"`
	// 语音通知内容模版
	Vms *VmsContentTemplateInfo `json:"Vms,omitempty"`
	// 企业微信通知内容模版
	WeChat *WeChatContentTemplateInfo `json:"WeChat,omitempty"`
	// 自定义的 Webhook 告警通知内容模版
	Webhook *WebhookContentTemplateInfo `json:"Webhook,omitempty"`
}

// CheckValidation 校验创建告警内容模版请求
func (v *CreateAlarmContentTemplateRequest) CheckValidation() error {
	if len(v.AlarmContentTemplateName) <= 0 {
		return errors.New("Invalid argument, empty AlarmContentTemplateName")
	}
	return nil
}

// CreateAlarmContentTemplateResponse 创建告警内容模版响应
type CreateAlarmContentTemplateResponse struct {
	CommonResponse
	// 告警通知模版 ID
	AlarmContentTemplateID string `json:"AlarmContentTemplateId"`
}

type CreateAppInstanceReq struct {
	CommonRequest
	// 实例类型，如ai助手等
	InstanceType AppInstanceType `json:"InstanceType"  binding:"required" enums:"anomaly_analysis,ai_assistant,other"`
	// 实例名称，当InstanceType为ai_assistant时，传入账户ID
	InstanceName string `json:"InstanceName"  binding:"required"`
	// 实例描述
	Description *string `json:"Description"`
}

func (c *CreateAppInstanceReq) CheckValidation() error {
	if len(c.InstanceName) == 0 {
		return errors.New("invalid argument, empty InstanceName")
	}

	return c.InstanceType.Validate()
}

type AppInstanceType string

const (
	AppInstanceTypeAiAssistant AppInstanceType = "ai_assistant"
)

func (a *AppInstanceType) Validate() error {
	switch *a {
	case AppInstanceTypeAiAssistant:
		return nil
	}
	return errors.New("invalid AppInstanceType")
}

type CreateAppInstanceResp struct {
	CommonResponse
	InstanceID string `json:"InstanceID"`
}

type DescribeAppInstancesReq struct {
	CommonRequest
	PageNumber   int              `json:"PageNumber"`
	PageSize     int              `json:"PageSize"`
	InstanceName *string          `json:"InstanceName"`
	InstanceType *AppInstanceType `json:"InstanceType"`
}

func (d *DescribeAppInstancesReq) CheckValidation() error {
	if d.PageNumber <= 0 {
		return errors.New("invalid argument, PageNumber")
	}
	if d.PageSize <= 0 {
		return errors.New("invalid argument, PageSize")
	}
	return nil
}

type DescribeAppInstancesResp struct {
	CommonResponse
	InstanceInfo []*InstanceInfo `json:"InstanceInfo"`
	Total        int64           `json:"Total"`
}

type InstanceInfo struct {
	// 实例Id
	InstanceId string `json:"InstanceId" required:"true"`
	// 实例名称
	InstanceName string `json:"InstanceName" required:"true"`
	// 实例类型
	InstanceType AppInstanceType `json:"InstanceType" required:"true"`
	// 描述
	Description *string `json:"Description"`
	// 创建时间
	CreateTime string `json:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime string `json:"UpdateTime,omitempty"`
}

type DeleteAppInstanceReq struct {
	CommonRequest
	InstanceId string `json:"InstanceId"  binding:"required"`
}

func (d *DeleteAppInstanceReq) CheckValidation() error {
	if d.InstanceId == "" {
		return errors.New("invalid argument, InstanceId")
	}
	return nil
}

type DeleteAppInstanceResp struct {
	CommonResponse
}

type APPMetaType string

var (
	AppMetaTypeAiAssistantSession            APPMetaType = "tls.app.ai_assistant.session"             // AI助手会话
	AppMetaTypeAiAssistantHistoryMessage     APPMetaType = "tls.app.ai_assistant.history_message"     // AI助手会话历史消息
	AppMetaTypeAiAssistantText2SqlSuggestion APPMetaType = "tls.app.ai_assistant.text2sql_suggestion" // AI助手文本转SQL建议
	AppMetaTypeAiAssistantText2SqlFeedBack   APPMetaType = "tls.app.ai_assistant.feed_back"           // AI助手文本转SQL 反馈
)

type CreateAppSceneMetaReq struct {
	CommonRequest
	// 应用实例id
	InstanceId string `json:"InstanceId"  binding:"required"`
	// 创建的App对应的Meta信息
	CreateAPPMetaType APPMetaType `json:"CreateAPPMetaType" binding:"required" enums:"tls.app.ai_assistant.session"`
	// topicId
	Id string `json:"Id"`
}

func (d *CreateAppSceneMetaReq) CheckValidation() error {
	if d.InstanceId == "" {
		return errors.New("invalid argument, InstanceId")
	}

	return nil
}

type CreateAppSceneMetaResp struct {
	CommonResponse
	Id string `json:"Id"`
}

type DescribeAppSceneMetasReq struct {
	CommonRequest
	InstanceId          string      `json:"InstanceId"`
	Id                  *string     `json:"Id"`
	DescribeAPPMetaType APPMetaType `json:"DescribeAPPMetaType" binding:"required" enums:"tls.app.ai_assistant.session,tls.app.ai_assistant.history_message,tls.app.ai_assistant.text2sql_suggestion"`
	PageNumber          int         `json:"PageNumber"`
	PageSize            int         `json:"PageSize"`
	PageContext         *string     `json:"PageContext"`
}

func (d *DescribeAppSceneMetasReq) CheckValidation() error {
	if d.InstanceId == "" {
		return errors.New("invalid argument, InstanceId")
	}
	if d.PageNumber < 0 {
		return errors.New("invalid argument, PageNumber")
	}
	if d.PageSize < 0 {
		return errors.New("invalid argument, PageSize")
	}
	return nil
}

type DescribeSessionMeta struct {
	SessionId  string `json:"SessionId"`
	Title      string `json:"Title"`
	CreateTime string `json:"CreateTime"`
	UpdateTime string `json:"UpdateTime"`
}
type DescribeSessionMessage struct {
	// 消息Id
	MessageId string `json:"MessageId"`
	// 创建消息时间
	CreatedTimeStamp string `json:"CreatedTimeStamp"`
	// 消息类型
	SessionMessageType string `json:"SessionMessageType" enums:"User,Assistant,System"`
	// 消息状态
	MsgStatus string `json:"MsgStatus" enums:"InProcess,Finished"`
	// 消息内容
	Messages []string `json:"Messages"`
	// 是否通过敏感词
	PassDetect bool `json:"PassDetect"`
	// 推理内容
	ReasoningContent string `json:"ReasoningContent"`
}

type DescribeAppSceneMetasRes struct {
	// 查询会话列表
	DescribeSessionMeta *DescribeSessionMeta `json:"DescribeSessionMeta,omitempty"`
	// 查询历史会话消息
	DescribeSessionMessage []*DescribeSessionMessage `json:"DescribeSessionMessage,omitempty"`
	// 查询会话建议
	DescribeSessionSuggestion string `json:"DescribeSessionSuggestion,omitempty"`
}

type DescribeAppSceneMetasResp struct {
	CommonResponse
	PageContext string                      `json:"PageContext,omitempty"`
	Total       int64                       `json:"Total"`
	Items       []*DescribeAppSceneMetasRes `json:"Items"`
}

type DeleteAppSceneMetaReq struct {
	CommonRequest
	InstanceId string `json:"InstanceId"  binding:"required"`
	MetaId     string `json:"MetaId"`
	// 场景类型
	DeleteAPPMetaType APPMetaType `json:"DeleteAPPMetaType" binding:"required" enums:"tls.app.ai_assistant.session"`
}

func (d *DeleteAppSceneMetaReq) CheckValidation() error {
	if d.InstanceId == "" {
		return errors.New("invalid argument, InstanceId")
	}
	if d.MetaId == "" {
		return errors.New("invalid argument, MetaId")
	}
	if d.DeleteAPPMetaType == AppMetaTypeAiAssistantSession {
		_, err := strconv.ParseInt(d.MetaId, 10, 64)
		if err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppSceneMetaResp struct {
	CommonResponse
}

type ModifyAppSceneMetaReq struct {
	CommonRequest
	// Ai应用实例的Id
	InstanceId string `json:"InstanceId"  binding:"required"`
	// 修改的App类型
	ModifyAPPMetaType APPMetaType `json:"ModifyAPPMetaType" binding:"required" enums:"tls.app.ai_assistant.feed_back"`
	// topicId
	Id     string     `json:"Id"`
	Record MetaRecord `json:"Record"  binding:"required"`
}

func (m *ModifyAppSceneMetaReq) CheckValidation() error {
	if m.InstanceId == "" {
		return errors.New("invalid argument, InstanceId")
	}
	return nil
}

type MetaRecord struct {
	// 反馈信息
	FeedBackMeta *FeedBackMeta `json:"FeedBackMeta,omitempty"`
}

type Feature int32

const (
	FeatureText2Sql         Feature = 0
	FeatureDoc              Feature = 1
	FeatureErrorFix         Feature = 3
	FeatureErrorExplanation Feature = 4
)

const (
	FeedBackTypePositive = "Positive"
	FeedBackTypeNegative = "Negative"
)

type FeedBackMeta struct {
	// 对话id
	SessionId string `json:"SessionId"`
	// 消息id
	MessageId string `json:"MessageId"`
	// 反馈类型
	FeedBackType string `json:"FeedBackType" enums:"Positive,Negative"`
	// Feature类型
	AiAssistantFeature Feature `json:"AiAssistantFeature"`
}

type GetAccountStatusRequest struct {
	CommonRequest
}

func (v *GetAccountStatusRequest) CheckValidation() error {
	return nil
}

type GetAccountStatusResponse struct {
	CommonResponse
	ArchVersion string `json:"ArchVersion"`
	Status      string `json:"Status"`
}

type CreateScheduleSqlTaskRequest struct {
	CommonRequest
	// 定时 SQL 分析任务名称。命名规则请参考资源命名规则
	TaskName string `json:"TaskName"`
	// 定时 SQL 分析任务的类型。
	// 0 表示“日志到日志”分析任务。
	// 1 表示“日志到指标”分析任务。
	TaskType int `json:"TaskType"`
	// 待进行定时 SQL 分析的原始日志所在的日志主题 ID。仅支持当前地域的日志主题
	TopicID string `json:"TopicID"`
	// 目标日志主题所属地域。默认为当前地域
	DestRegion *string `json:"DestRegion,omitempty"`
	// 用于存储定时 SQL 分析结果数据的目标日志主题ID
	DestTopicID string `json:"DestTopicID"`
	// 调度定时 SQL 分析任务的开始时间，即创建第一个实例的时间。格式为秒级时间戳
	ProcessStartTime int `json:"ProcessStartTime"`
	// 调度定时 SQL 分析任务的结束时间，格式为秒级时间戳。如果不配置，表示持续运行定时 SQL 分析任务
	ProcessEndTime *int `json:"ProcessEndTime,omitempty"`
	// SQL 时间窗口，即定时 SQL 分析任务运行时，日志检索与分析的时间范围，左闭右开格式。最大 24 小时，最小 1 分钟
	ProcessTimeWindow string `json:"ProcessTimeWindow"`
	// 定时 SQL 分析任务定期执行的检索与分析语句，应符合日志服务的检索与分析语法
	Query string `json:"Query"`
	// 定时 SQL 分析任务的调度周期。调度周期决定每个实例的调度时间
	RequestCycle RequestCycle `json:"RequestCycle"`
	// 完成任务配置后是否立即启动定时 SQL 分析任务。0：关闭任务，后续需手动启动任务；1：立即启动
	Status int `json:"Status"`
	// 每次调度的延迟时间。取值范围为 0～120，单位为秒。如果不配置，则表示 0，即无延时
	ProcessSqlDelay *int `json:"ProcessSqlDelay,omitempty"`
	// 定时 SQL 分析任务的简单描述。不支持 `<>`、`'`、`\`、`\\`；长度范围为 0～64 个字符
	Description *string `json:"Description,omitempty"`
}

func (v *CreateScheduleSqlTaskRequest) CheckValidation() error {
	if len(v.TaskName) <= 0 {
		return errors.New("Invalid argument, empty TaskName")
	}
	if len(v.TopicID) <= 0 {
		return errors.New("Invalid argument, empty TopicID")
	}
	if len(v.DestTopicID) <= 0 {
		return errors.New("Invalid argument, empty DestTopicID")
	}
	if v.ProcessStartTime <= 0 {
		return errors.New("Invalid argument, ProcessStartTime must be greater than 0")
	}
	if len(v.ProcessTimeWindow) <= 0 {
		return errors.New("Invalid argument, empty ProcessTimeWindow")
	}
	if len(v.Query) <= 0 {
		return errors.New("Invalid argument, empty Query")
	}
	if v.RequestCycle.Time <= 0 {
		return errors.New("Invalid argument, RequestCycle.Time must be greater than 0")
	}
	if len(v.RequestCycle.Type) <= 0 {
		return errors.New("Invalid argument, empty RequestCycle.Type")
	}
	if v.Status != 0 && v.Status != 1 {
		return errors.New("Invalid argument, Status must be 0 or 1")
	}
	return nil
}

type CreateScheduleSqlTaskResponse struct {
	CommonResponse
	// 定时 SQL 分析任务 ID
	TaskID string `json:"TaskId"`
}

type ModifyAppSceneMetaResp struct {
	CommonResponse
}

type DeleteScheduleSqlTaskRequest struct {
	CommonRequest
	TaskID string `json:"TaskId"`
}

type DescribeScheduleSqlTaskRequest struct {
	CommonRequest
	TaskID string `json:"TaskId"`
}

func (v *DeleteScheduleSqlTaskRequest) CheckValidation() error {
	if len(v.TaskID) <= 0 {
		return errors.New("Invalid argument, empty TaskID")
	}
	return nil
}

func (v *DescribeScheduleSqlTaskRequest) CheckValidation() error {
	if len(v.TaskID) <= 0 {
		return errors.New("Invalid argument, empty TaskID")
	}
	return nil
}

type DeleteScheduleSqlTaskResponse struct {
	CommonResponse
}

type TraceInstanceStatus string

const (
	TraceInstanceStatusCreating TraceInstanceStatus = "CREATING"
	TraceInstanceStatusCreated  TraceInstanceStatus = "CREATED"
	TraceInstanceStatusDeleting TraceInstanceStatus = "DELETING"
)

type DescribeTraceInstancesRequest struct {
	CommonRequest
	PageNumber        *int    `json:",omitempty"`
	PageSize          *int    `json:",omitempty"`
	TraceInstanceName *string `json:",omitempty"`
	TraceInstanceID   *string `json:",omitempty"`
	ProjectID         *string `json:",omitempty"`
	ProjectName       *string `json:",omitempty"`
	Status            *string `json:",omitempty"`
	IamProjectName    *string `json:",omitempty"`
}

func (v *DescribeTraceInstancesRequest) CheckValidation() error {
	return nil
}

type DescribeTraceInstancesResponse struct {
	CommonResponse
	Total          int64            `json:"Total"`
	TraceInstances []*TraceInstance `json:"TraceInstances"`
}

type TraceInstance struct {
	CreateTime               string              `json:"CreateTime"`
	DependencyTopicID        string              `json:"DependencyTopicId"`
	DependencyTopicTopicName string              `json:"DependencyTopicTopicName"`
	Description              string              `json:"Description"`
	ModifyTime               string              `json:"ModifyTime"`
	ProjectID                string              `json:"ProjectId"`
	ProjectName              string              `json:"ProjectName"`
	TraceInstanceID          string              `json:"TraceInstanceId"`
	TraceInstanceName        string              `json:"TraceInstanceName"`
	TraceInstanceStatus      TraceInstanceStatus `json:"TraceInstanceStatus"`
	TraceTopicID             string              `json:"TraceTopicId"`
	TraceTopicName           string              `json:"TraceTopicName"`
}

// Trace 相关结构体
type CreateTraceInstanceRequest struct {
	CommonRequest
	ProjectID         string `json:"ProjectId"`
	TraceInstanceName string `json:"TraceInstanceName"`
	Description       string `json:"Description,omitempty"`
}

func (v *CreateTraceInstanceRequest) CheckValidation() error {
	if len(v.ProjectID) <= 0 {
		return errors.New("Invalid argument, empty ProjectID")
	}
	if len(v.TraceInstanceName) <= 0 {
		return errors.New("Invalid argument, empty TraceInstanceName")
	}
	return nil
}

type CreateTraceInstanceResponse struct {
	CommonResponse
	TraceInstanceID string `json:"TraceInstanceId"`
}

type DeleteTraceInstanceRequest struct {
	CommonRequest
	TraceInstanceId string `json:"TraceInstanceId"`
}

func (v *DeleteTraceInstanceRequest) CheckValidation() error {
	if len(v.TraceInstanceId) <= 0 {
		return errors.New("Invalid argument, empty TraceInstanceId")
	}
	return nil
}

type DeleteTraceInstanceResponse struct {
	CommonResponse
}

type DescribeScheduleSqlTaskResponse struct {
	CommonResponse
	TaskID            string        `json:"TaskId,omitempty"`
	TaskName          string        `json:"TaskName,omitempty"`
	Description       string        `json:"Description,omitempty"`
	SourceProjectID   string        `json:"SourceProjectID,omitempty"`
	SourceProjectName string        `json:"SourceProjectName,omitempty"`
	SourceTopicID     string        `json:"SourceTopicID,omitempty"`
	SourceTopicName   string        `json:"SourceTopicName,omitempty"`
	DestRegion        string        `json:"DestRegion,omitempty"`
	DestProjectID     string        `json:"DestProjectID,omitempty"`
	DestTopicID       string        `json:"DestTopicID,omitempty"`
	DestTopicName     string        `json:"DestTopicName,omitempty"`
	Status            *int          `json:"Status,omitempty"`
	ProcessStartTime  *int64        `json:"ProcessStartTime,omitempty"`
	ProcessEndTime    *int64        `json:"ProcessEndTime,omitempty"`
	ProcessSqlDelay   *int          `json:"ProcessSqlDelay,omitempty"`
	ProcessTimeWindow string        `json:"ProcessTimeWindow,omitempty"`
	Query             string        `json:"Query,omitempty"`
	RequestCycle      *RequestCycle `json:"RequestCycle,omitempty"`
	CreateTimeStamp   *int64        `json:"CreateTimeStamp,omitempty"`
	ModifyTimeStamp   *int64        `json:"ModifyTimeStamp,omitempty"`
}

// DescribeScheduleSqlTasks 相关结构体
type DescribeScheduleSqlTasksRequest struct {
	CommonRequest
	ProjectId       *string `json:"ProjectId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty"`
	IamProjectName  *string `json:"IamProjectName,omitempty"`
	TopicId         *string `json:"TopicId,omitempty"`
	SourceTopicName *string `json:"SourceTopicName,omitempty"`
	TaskId          *string `json:"TaskId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty"`
	Status          *string `json:"Status,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty"`
}

func (v *DescribeScheduleSqlTasksRequest) CheckValidation() error {
	return nil
}

type DescribeScheduleSqlTasksResponse struct {
	CommonResponse
	Tasks []DescribeScheduleSqlTaskResp `json:"Tasks"`
	Total int                           `json:"Total"`
}

type DescribeScheduleSqlTaskResp struct {
	Query             string       `json:"Query"`
	Status            int          `json:"Status"`
	TaskId            string       `json:"TaskId"`
	TaskName          string       `json:"TaskName"`
	DestRegion        string       `json:"DestRegion"`
	Description       string       `json:"Description"`
	DestTopicID       string       `json:"DestTopicID"`
	RequestCycle      RequestCycle `json:"RequestCycle"`
	DestTopicName     string       `json:"DestTopicName"`
	ProcessEndTime    int          `json:"ProcessEndTime"`
	CreateTimeStamp   int          `json:"CreateTimeStamp"`
	ModifyTimeStamp   int          `json:"ModifyTimeStamp"`
	ProcessSqlDelay   int          `json:"ProcessSqlDelay"`
	SourceProjectID   string       `json:"SourceProjectID"`
	SourceProjectName string       `json:"SourceProjectName"`
	ProcessTimeWindow string       `json:"ProcessTimeWindow"`
	SourceTopicName   string       `json:"SourceTopicName"`
	SourceTopicID     string       `json:"SourceTopicID"`
	ProcessStartTime  int          `json:"ProcessStartTime"`
}
type FilterTag struct {
	Key    string   `json:"Key,omitempty"`
	Values []string `json:"Values,omitempty"`
}

type ResourceTag struct {
	TagKey       string `json:"TagKey"`
	TagValue     string `json:"TagValue"`
	ResourceID   string `json:"ResourceId"`
	ResourceType string `json:"ResourceType"`
}

type ListTagsForResourcesRequest struct {
	CommonRequest
	MaxResults   int          `json:"MaxResults,omitempty"`
	NextToken    string       `json:"NextToken,omitempty"`
	ResourceType string       `json:"ResourceType"`
	ResourcesIds []string     `json:"ResourcesIds,omitempty"`
	TagFilters   []*FilterTag `json:"TagFilters,omitempty"`
}

func (v *ListTagsForResourcesRequest) CheckValidation() error {
	if len(v.ResourceType) <= 0 {
		return errors.New("Invalid argument, empty ResourceType")
	}
	return nil
}

type ListTagsForResourcesResponse struct {
	CommonResponse
	ResourceTags []ResourceTag `json:"ResourceTags"`
	NextToken    string        `json:"NextToken"`
}

// DescribeTrace
type DescribeTraceRequest struct {
	CommonRequest
	TraceId         string `json:"TraceId"`
	TraceInstanceId string `json:"TraceInstanceId"`
}

func (v *DescribeTraceRequest) CheckValidation() error {
	if len(v.TraceId) <= 0 {
		return errors.New("Invalid argument, empty TraceId")
	}
	if len(v.TraceInstanceId) <= 0 {
		return errors.New("Invalid argument, empty TraceInstanceId")
	}
	return nil
}

// Alarm Content Template related structures

type ModifyAlarmContentTemplateRequest struct {
	CommonRequest
	AlarmContentTemplateID   string                       `json:"AlarmContentTemplateId"`
	AlarmContentTemplateName string                       `json:"AlarmContentTemplateName"`
	DingTalk                 *DingTalkContentTemplateInfo `json:"DingTalk,omitempty"`
	Email                    *EmailContentTemplateInfo    `json:"Email,omitempty"`
	Lark                     *LarkContentTemplateInfo     `json:"Lark,omitempty"`
	NeedValidContent         *bool                        `json:"NeedValidContent,omitempty"`
	Sms                      *SmsContentTemplateInfo      `json:"Sms,omitempty"`
	Vms                      *VmsContentTemplateInfo      `json:"Vms,omitempty"`
	WeChat                   *WeChatContentTemplateInfo   `json:"WeChat,omitempty"`
	Webhook                  *WebhookContentTemplateInfo  `json:"Webhook,omitempty"`
}

func (v *ModifyAlarmContentTemplateRequest) CheckValidation() error {
	if len(v.AlarmContentTemplateID) <= 0 {
		return errors.New("Invalid argument, empty AlarmContentTemplateID")
	}
	if len(v.AlarmContentTemplateName) <= 0 {
		return errors.New("Invalid argument, empty AlarmContentTemplateName")
	}
	return nil
}

// DescribeAlarmContentTemplates 相关结构体
type DescribeAlarmContentTemplatesRequest struct {
	CommonRequest
	AlarmContentTemplateName *string
	AlarmContentTemplateID   *string
	OrderField               *string
	ASC                      *bool
	PageNumber               int
	PageSize                 int
}

func (v *DescribeAlarmContentTemplatesRequest) CheckValidation() error {
	return nil
}

type DescribeAlarmContentTemplatesResponse struct {
	CommonResponse
	AlarmContentTemplates []ContentTemplateInfo `json:"AlarmContentTemplates"`
	Total                 int                   `json:"Total"`
}

type ContentTemplateInfo struct {
	Sms                      *SmsContentTemplateInfo      `json:"Sms,omitempty"`
	Vms                      *VmsContentTemplateInfo      `json:"Vms,omitempty"`
	Lark                     *LarkContentTemplateInfo     `json:"Lark,omitempty"`
	Email                    *EmailContentTemplateInfo    `json:"Email,omitempty"`
	WeChat                   *WeChatContentTemplateInfo   `json:"WeChat,omitempty"`
	Webhook                  *WebhookContentTemplateInfo  `json:"Webhook,omitempty"`
	DingTalk                 *DingTalkContentTemplateInfo `json:"DingTalk,omitempty"`
	IsDefault                bool                         `json:"IsDefault"`
	CreateTime               string                       `json:"CreateTime"`
	ModifyTime               string                       `json:"ModifyTime"`
	AlarmContentTemplateID   string                       `json:"AlarmContentTemplateId"`
	AlarmContentTemplateName string                       `json:"AlarmContentTemplateName"`
}

type DescribeTraceResponse struct {
	CommonResponse
	Trace Trace `json:"Trace"`
}

type Trace struct {
	Spans []Span `json:"Spans"`
}

type Span struct {
	Kind                   string                 `json:"Kind"`
	Name                   string                 `json:"Name"`
	Links                  []SpanLink             `json:"Links,omitempty"`
	Events                 []SpanEvent            `json:"Events,omitempty"`
	SpanId                 string                 `json:"SpanId"`
	Status                 Status                 `json:"Status"`
	EndTime                int                    `json:"EndTime"`
	TraceId                string                 `json:"TraceId"`
	Resource               Resource               `json:"Resource"`
	StartTime              int                    `json:"StartTime"`
	Attributes             []KeyValue             `json:"Attributes,omitempty"`
	TraceState             string                 `json:"TraceState,omitempty"`
	ParentSpanId           string                 `json:"ParentSpanId,omitempty"`
	InstrumentationLibrary InstrumentationLibrary `json:"InstrumentationLibrary,omitempty"`
}

type SpanLink struct {
	SpanId     string     `json:"SpanId"`
	TraceId    string     `json:"TraceId"`
	Attributes []KeyValue `json:"Attributes,omitempty"`
	TraceState string     `json:"TraceState,omitempty"`
}

type SpanEvent struct {
	Name       string     `json:"Name"`
	Timestamp  int        `json:"Timestamp"`
	Attributes []KeyValue `json:"Attributes,omitempty"`
}

type Status struct {
	Code    string  `json:"Code"`
	Message *string `json:"Message,omitempty"`
}

type Resource struct {
	Attributes []KeyValue `json:"Attributes"`
}

type KeyValue struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type InstrumentationLibrary struct {
	Name    string  `json:"Name"`
	Version *string `json:"Version,omitempty"`
}

type ModifyAlarmContentTemplateResponse struct {
	CommonResponse
}
