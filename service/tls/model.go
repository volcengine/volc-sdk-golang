package tls

import (
	"database/sql/driver"
	"encoding/json"
	"net/http"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type CommonResponse struct {
	RequestID string `json:"RequestId"`
}

type CommonRequest struct {
	Headers map[string]string `json:"-"`
}

type CreateProjectRequest struct {
	CommonRequest
	ProjectName string
	Description string
	Region      string
}

type CreateProjectResponse struct {
	CommonResponse
	ProjectID string `json:"ProjectId"`
}

type DeleteProjectRequest struct {
	CommonRequest
	ProjectID string `json:"ProjectId"`
}

type DescribeProjectRequest struct {
	CommonRequest
	ProjectID string `json:"ProjectId"`
}

type DescribeProjectResponse struct {
	CommonResponse
	ProjectInfo
}

type DescribeProjectsRequest struct {
	CommonRequest
	ProjectName string
	ProjectID   string
	PageNumber  int
	PageSize    int
	IsFullName  bool
}

type ProjectInfo struct {
	ProjectID       string `json:"ProjectId"`
	ProjectName     string `json:"ProjectName"`
	Description     string `json:"Description"`
	CreateTimestamp string `json:"CreateTime"`
	TopicCount      int64  `json:"TopicCount"`
	InnerNetDomain  string `json:"InnerNetDomain"`
}

type DescribeProjectsResponse struct {
	CommonResponse
	Projects []ProjectInfo `json:"Projects"`

	Total int64 `json:"Total"`
}

type ModifyProjectRequest struct {
	CommonRequest
	ProjectID   string
	ProjectName *string
	Description *string
}

type CreateTopicRequest struct {
	CommonRequest
	ProjectID      string
	TopicName      string
	Ttl            uint16
	Description    string
	ShardCount     int
	MaxSplitShard  *int32 `json:",omitempty"`
	AutoSplit      bool
	EnableTracking *bool `json:",omitempty"`
}

type CreateTopicResponse struct {
	CommonResponse
	TopicID string `json:"TopicID"`
}

type DeleteTopicRequest struct {
	CommonRequest
	TopicID string
}

type ModifyTopicRequest struct {
	CommonRequest
	TopicID        string  `json:"TopicId"`
	TopicName      *string `json:"TopicName"`
	Ttl            *uint16 `json:"Ttl"`
	Description    *string `json:"Description"`
	MaxSplitShard  *int32  `json:",omitempty"`
	AutoSplit      *bool   `json:",omitempty"`
	EnableTracking *bool   `json:",omitempty"`
}

type DescribeTopicRequest struct {
	CommonRequest
	TopicID string
}

type DescribeTopicResponse struct {
	CommonResponse
	TopicName       string `json:"TopicName"`
	ProjectID       string `json:"ProjectId"`
	TopicID         string `json:"TopicId"`
	Ttl             uint16 `json:"Ttl"`
	CreateTimestamp string `json:"CreateTime"`
	ModifyTimestamp string `json:"ModifyTime"`
	ShardCount      int32  `json:"ShardCount"`
	Description     string `json:"Description"`
	MaxSplitShard   int32  `json:"MaxSplitShard"`
	AutoSplit       bool   `json:"AutoSplit"`
	EnableTracking  bool   `json:"EnableTracking"`
}

type DescribeTopicsRequest struct {
	CommonRequest
	ProjectID  string
	PageNumber int
	PageSize   int
	TopicName  string
	TopicID    string
	IsFullName bool
}

type Topic struct {
	TopicName       string `json:"TopicName"`
	ProjectID       string `json:"ProjectId"`
	TopicID         string `json:"TopicId"`
	Ttl             uint16 `json:"Ttl"`
	CreateTimestamp string `json:"CreateTime"`
	ModifyTimestamp string `json:"ModifyTime"`
	ShardCount      int32  `json:"ShardCount"`
	Description     string `json:"Description"`
	MaxSplitShard   int32  `json:"MaxSplitShard"`
	AutoSplit       bool   `json:"AutoSplit"`
	EnableTracking  bool   `json:"EnableTracking"`
}

type DescribeTopicsResponse struct {
	CommonResponse
	Topics []*Topic `json:"Topics"`
	Total  int      `json:"Total"`
}

type Value struct {
	ValueType      string `json:"ValueType"`
	Delimiter      string `json:"Delimiter"`
	CasSensitive   bool   `json:"CaseSensitive"`
	IncludeChinese bool   `json:"IncludeChinese"`
	SQLFlag        bool   `json:"SqlFlag"`
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
	TopicID  string          `json:"TopicId"`
	FullText *FullTextInfo   `json:"FullText"`
	KeyValue *[]KeyValueInfo `json:"KeyValue"`
}

type FullTextInfo struct {
	Delimiter      string `json:"Delimiter"`
	CaseSensitive  bool   `json:"CaseSensitive"`
	IncludeChinese bool   `json:"IncludeChinese"`
}

type KeyValueInfo struct {
	Key   string
	Value Value
}

type CreateIndexResponse struct {
	CommonResponse
	TopicID string `json:"TopicId"`
}

type DeleteIndexRequest struct {
	CommonRequest
	TopicID string
}

type DescribeIndexRequest struct {
	CommonRequest
	TopicID string
}

type DescribeIndexResponse struct {
	CommonResponse
	TopicID    string          `json:"TopicId"`
	FullText   *FullTextInfo   `json:"FullText"`
	KeyValue   *[]KeyValueInfo `json:"KeyValue"`
	CreateTime string          `json:"CreateTime"`
	ModifyTime string          `json:"ModifyTime"`
}

type ModifyIndexRequest struct {
	CommonRequest
	TopicID  string          `json:"TopicId"`
	FullText *FullTextInfo   `json:"FullText"`
	KeyValue *[]KeyValueInfo `json:"KeyValue"`
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

type DescribeShardsResponse struct {
	CommonResponse
	Shards []*struct {
		TopicID           string `json:"TopicId"`
		ShardID           int32  `json:"ShardId"`
		InclusiveBeginKey string `json:"InclusiveBeginKey"`
		ExclusiveEndKey   string `json:"ExclusiveEndKey"`
		Status            string `json:"Status"`
		ModifyTimestamp   string `json:"ModifyTime"`
		StopWriteTime     string `json:"StopWriteTime"`
	} `json:"Shards"`

	Total int `json:"Total"`
}

type PutLogsRequest struct {
	CommonRequest
	TopicID      string
	HashKey      string
	CompressType string
	LogBody      *pb.LogGroupList
}

type DescribeCursorRequest struct {
	CommonRequest
	TopicID string
	ShardID int
	From    string
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
	EndCursor     *string
	LogGroupCount *int
	Compression   *string
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

type DescribeLogContextResponse struct {
	CommonResponse
	LogContextInfos []map[string]interface{} `json:"LogContextInfos"`
	PrevOver        bool                     `json:"PrevOver"`
	NextOver        bool                     `json:"NextOver"`
}

type CreateRuleRequest struct {
	CommonRequest
	// 采集配置所属于的日志主题ID，即TopicID。
	TopicID string `json:"TopicId" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// 采集配置的名称:
	// - 在一个Topic中唯一;
	// - 长度限制为3-63个字符，只能包含小写字母、数字和短划线（-），必须以小写字母或者数字开头和结尾。
	RuleName string `json:"RuleName" binding:"required" example:"testName"`
	// 采集路径列表:
	// - 最多能够创建10个采集路径;
	// - 采集路径必须是一个绝对路径;
	// 当InputType=0时:
	// - 采集路径支持完整匹配和通配符模式匹配，通配符只支持"**"、"*"、"?"，但是最多只能配置1个"**"通配符.
	// 当InputType=1时:
	// - 不允许配置采集路径列表
	// 当InputType=2时:
	// - 采集路径支持完整匹配和通配符模式匹配，通配符只支持“*”、“?”，不支持“**”。
	Paths *[]string `json:"Paths" binding:"required" example:"[\"/data/nginx/log/**/access.log\"]"`
	// 采集的日志类型:
	// - minimalist_log代表极简日志；
	// - delimiter_log代表分隔符格式日志；
	// - json_log代表json格式日志；
	// - multiline_log代表多行日志；
	// - fullregex_log代表完整正则日志；
	// - 当创建采集配置时，如果用户没有配置采集的日志类型，那么采集的日志类型默认值是minimalist_log。
	LogType *string `json:"LogType" enums:"minimalist_log,delimiter_log,json_log,multiline_log,fullregex_log" example:"delimiter_log"`
	// 提取规则: 如果配置非minimalist_log或者非json_log的采集的日志类型，那么必须同时配置提取规则。
	ExtractRule *ExtractRule `json:"ExtractRule"`
	// 采集黑名单列表: 最多能够创建10个采集黑名单。
	// 当InputType=0时:
	// - 当Type是Path时，Value表示一个目录。支持完整匹配和通配符模式匹配，通配符只支持“*”、“?”，不支持“**”通配符。
	// - 当Type是File时，Value表示一个文件名称。支持完整匹配和通配符模式匹配，通配符只支持“**”、“*”、“?”，但是最多只能配置1个“**”通配符。
	// 当InputType=1时:
	// - 不允许配置采集黑名单列表。
	// 当InputType=2时：
	// - 当Type是Path时，Value表示一个目录。仅支持完整匹配，不支持通配符模式匹配。
	// - 当Type是File时，Value表示一个文件名称。目录部分仅支持完整匹配，不支持通配符模式匹配。文件名称部分支持完整匹配和通配符模式匹配，通配符只支持“*”、“?”。
	ExcludePaths *[]ExcludePath `json:"ExcludePaths"`
	// 用户自定义的采集规则。
	UserDefineRule *UserDefineRule `json:"UserDefineRule"`
	// 日志样例：日志样例的长度必须不大于3000个字符
	LogSample *string `json:"LogSample" example:"xxxxxx"`
	// 采集类型：
	// - 0：宿主机日志文件
	// - 1：k8s容器标准输出
	// - 2：k8s容器内日志文件
	// 当创建采集配置时，如果用户没有配置采集的类型，那么采集的类型默认值是0。
	InputType *int `json:"InputType"`
	// 容器采集规则
	ContainerRule *ContainerRule `json:"ContainerRule"`
}

type ExtractRule struct {
	// 分隔符;\n当且仅当采集的日志类型为delimiter_log时有效
	Delimiter string `json:"Delimiter,omitempty" example:"#"`
	// 第一行日志需要匹配的regex;\n当且仅当采集的日志类型为multiline_log时有效;\n必须是合法的正则表达式;
	BeginRegex string `json:"BeginRegex,omitempty" example:"\\[(\\d+-\\d+-\\w+:\\d+:\\d+,\\d+)\\]\\s\\[(\\w+)\\]\\s(.*)"`
	// 整条日志需要匹配的regex;当且仅当采集的日志类型为fullregex_log时有效;必须是合法的正则表达式。
	LogRegex string `json:"LogRegex,omitempty" example:"\\[(\\d+-\\d+-\\w+:\\d+:\\d+,\\d+)\\]\\s\\[(\\w+)\\]\\s(.*)"`
	// 用户自定义的每个字段的名字key列表;当且仅当采集的日志类型为delimiter_log或fullregex_log时有效。
	// - 最多能够配置100个名字key。
	// - 命名规则同「索引配置」的key名称。
	// - 当采集的日志类型为delimiter_log时, 不能配置非空、重复的名字key，不能全部配置为空的名字key。
	// - 当采集的日志类型为fullregex_log时，不能配置重复的名字key， 不能配置空的名字key。
	Keys []string `json:"Keys,omitempty"`
	// 时间字段的名字key。
	// - TimeKey和TimeFormat必须成对出现。
	// - 命名规则同「索引配置」的key名称。
	TimeKey string `json:"TimeKey,omitempty" example:"time"`
	// 时间字段的格式。
	// - TimeKey和TimeFormat必须成对出现。
	// - 参考c语言的strftime函数对于format参数的说明：https://man7.org/linux/man-pages/man3/strftime.3.html
	TimeFormat string `json:"TimeFormat,omitempty" example:"%Y-%m-%dT%H:%M:%S,%f"`
	// 过滤规则列表。
	// - 当采集的日志类型为minimalist_log或multiline_log时, 最多能够配置1条过滤规则，并且过滤字段的名字key必须为__content__。
	// - 当采集的日志类型为delimiter_log、json_log或fullregex_log时, 最多能够配置5条过滤规则，并且过滤字段的名字key不能重复、不能为空，命名规则同「索引配置」的key名称。过滤字段的日志内容需要匹配的regex必须是合法的正则表达式，并且长度限制是256个字符。
	FilterKeyRegex []FilterKeyRegex `json:"FilterKeyRegex,omitempty"`
	// 是否上传解析失败的日志。
	// - UnMatchUpLoadSwitch=true和UnMatchLogKey必须成对出现。
	// - true表示上传解析失败的日志，false表示不上传解析失败的日志。
	UnMatchUpLoadSwitch bool `json:"UnMatchUpLoadSwitch,omitempty" example:"true"`
	// 当上传解析失败的日志时，解析失败的日志的key名称。
	// - UnMatchUpLoadSwitch=true和UnMatchLogKey必须成对出现。
	// - 命名规则同「索引配置」的key名称。
	UnMatchLogKey string `json:"UnMatchLogKey,omitempty" example:"LogParseFailed"`
	// 根据日志模板自动提取日志内容。
	LogTemplate LogTemplate `json:"LogTemplate,omitempty"`
}

type FilterKeyRegex struct {
	// 过滤字段的名字key。
	Key string `json:"Key,omitempty" example:"__content__"`
	// 过滤字段的日志内容需要匹配的regex。
	Regex string `json:"Regex,omitempty" example:".*ERROR.*"`
}

type LogTemplate struct {
	// 日志模板的类型。
	// - Nginx：Nginx类型的日志模板；
	Type string `json:"Type,omitempty" example:"Nginx"`
	// 日志模板的格式。
	Format string `json:"Format,omitempty" example:"log_format main  '$remote_addr - $remote_user [$time_local] \"$request\" $request_time $request_length $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\"';"`
}

type ExcludePath struct {
	// 类型，File或Path。
	Type string `json:"Type,omitempty"`
	// 当Type是Path时，Value表示一个目录。
	// - Value必须是一个绝对路径；支持完整匹配和通配符模式匹配，通配符只支持“*”、“?”，不支持“**”通配符。
	// 当Type是File时，Value表示一个文件名称。
	// - Value必须是一个绝对路径；  - 支持完整匹配和通配符模式匹配，通配符只支持“**”、“*”、“?”，但是最多只能配置1个“**”通配符。
	Value string `json:"Value,omitempty"`
}

type UserDefineRule struct {
	// 解析路径的规则。
	ParsePathRule *ParsePathRule `json:"ParsePathRule,omitempty"`
	// 根据HashKey路由日志分区。
	ShardHashKey *ShardHashKey `json:"ShardHashKey,omitempty"`
	// 是否上传原始日志。
	// - true：上传原始日志。
	// - false：不上传原始日志。
	EnableRawLog bool `json:"EnableRawLog,omitempty"`
	// 添加常量字段
	// - 最多能够添加5个常量字段
	// - 键不能重复、不能为空，最多包含128个字符：字母（a-z、A-Z），数字（0-9 ），中划线（-），下划线（_），点（.），斜线（/），但是不能以下划线（_）开头
	// - 值不能为空，长度最多为512KB
	Fields map[string]string `json:"Fields,omitempty"`
}

type ShardHashKey struct {
	// 路由日志分区的HashKey。
	// - 长度必须不大于32位；
	// - 每个字符必须都是16进制（0~9、a~f、A~F）；
	// - 不能是ffffffffffffffffffffffffffffffff，因为取值范围是：[00000000000000000000000000000000,ffffffffffffffffffffffffffffffff）；
	HashKey string `json:"HashKey,omitempty"`
}

type ParsePathRule struct {
	// 路径样例。
	//- 路径样例必须是一个绝对路径。
	//- 路径样例不能包含通配符*、?、**。
	PathSample string `json:"PathSample,omitempty"`
	//路径样例需要匹配的regex。
	//- 必须是合法的正则表达式。
	Regex string `json:"Regex,omitempty"`
	//解析出来的路径字段的名字key列表。
	//- 最多配置100个名字key
	//- 不能配置重复的名字key。
	//- 不能配置空的名字key。
	//- 命名规则同「索引配置」的key名称。
	Keys []string `json:"Keys,omitempty"`
}

type ContainerRule struct {
	// stdout-采集容器标准输出stdout。
	// stderr-采集容器标准错stderr。
	// all-采集容器标准输出stdout和容器标准错stderr。
	// 默认值是all
	Stream string `json:"Stream,omitempty" enums:"stdout,stderr,all"`
	// 通过容器名称指定待采集的容器，支持正则匹配。例如设置为"ContainerNameRegex":"^(container-test)$"，表示匹配所有名为container-test的容器。
	ContainerNameRegex string `json:"ContainerNameRegex,omitempty"`
	// 通过容器Label白名单指定待采集的容器。
	// 如果要设置容器Label白名单，则Key必填。
	// - 如果Value不为空，则只采集在容器Label中包含Key=Value的容器。
	// - 如果Value为空，则采集所有在容器Label中包含Key的容器。
	// 说明：
	// - 多个键值对之间为或关系，即只要容器Label满足任一键值对即可被采集。
	// - Value默认为字符串匹配，即只有Value和容器Label完全相同才会匹配。如果该值以^开头并且以$结尾，则为正则匹配，例如：Value配置为^(kube-system|istio-system)$，可同时匹配kube-system和istio-system。
	// - 请勿设置相同的Key，如果重名只生效一个。
	IncludeContainerLabelRegex map[string]string `json:"IncludeContainerLabelRegex,omitempty"`
	// 通过容器Label黑名单排除不采集的容器。
	// 如果要设置容器Label黑名单，则Key必填。
	// - 如果Value不为空，则只排除在容器Label中包含Key=Value的容器。
	// - 如果Value为空，则排除所有在容器Label中包含Key的容器。
	// 说明：
	// - 多个键值对之间为或关系，即只要容器Label满足任一键值对即不被采集。
	// - Value默认为字符串匹配，即只有Value和容器Label完全相同才会匹配。如果该值以^开头并且以$结尾，则为正则匹配，例如：Value配置为^(kube-system|istio-system)$，可同时匹配kube-system和istio-system。
	// - 请勿设置相同的Key，如果重名只生效一个。
	ExcludeContainerLabelRegex map[string]string `json:"ExcludeContainerLabelRegex,omitempty"`
	// 通过容器环境变量白名单指定待采集的容器。
	// 如果要设置容器环境变量白名单，则Key必填。
	// - 如果Value不为空，则只采集在容器环境变量中包含Key=Value的容器。
	// - 如果Value为空，则采集所有在容器环境变量中包含Key的容器。
	// 说明：
	// - 多个键值对之间为或关系，即只要容器环境变量满足任一键值对即可被采集。
	// - Value默认为字符串匹配，即只有Value和容器环境变量完全相同才会匹配。如果该值以^开头并且以$结尾，则为正则匹配，例如：Value配置为^(kube-system|istio-system)$，可同时匹配kube-system和istio-system。
	IncludeContainerEnvRegex map[string]string `json:"IncludeContainerEnvRegex,omitempty"`
	// 通过容器环境变量黑名单排除不采集的容器。
	// 如果要容器设置环境变量黑名单，则Key必填。
	// - 如果Value不为空，则只排除在容器环境变量中包含Key=Value的容器。
	// - 如果Value为空，则排除所有在容器环境变量中包含Key的容器。
	// 说明：
	// - 多个键值对之间为或关系，即只要容器环境变量满足任一键值对即不被采集。
	// - Value默认为字符串匹配，即只有Value和容器环境变量完全相同才会匹配。如果该值以^开头并且以$结尾，则为正则匹配，例如：Value配置为^(kube-system|istio-system)$，可同时匹配kube-system和istio-system。
	ExcludeContainerEnvRegex map[string]string `json:"ExcludeContainerEnvRegex,omitempty"`
	// 设置环境变量日志标签后，日志服务将在日志中新增环境变量相关字段。例如设置EnvKey为VERSION，设置EnvValue为env_version，当容器中包含环境变量VERSION=v1.0.0时，会将该信息添加到日志中，即添加字段__tag__env_version__: v1.0.0。
	EnvTag map[string]string `json:"EnvTag,omitempty"`
	// Kubernetes容器的采集规则
	KubernetesRule KubernetesRule `json:"KubernetesRule,omitempty"`
}

type KubernetesRule struct {
	// 通过Namespace名称指定采集的容器，支持正则匹配。例如设置为"NamespaceNameRegex":"^(default|nginx)$"，表示匹配nginx命名空间、default命名空间下的所有容器。
	NamespaceNameRegex string `json:"NamespaceNameRegex,omitempty"`
	// 通过工作负载的类型指定采集的容器。例如设置为"WorkloadType":"Deployment"，表示匹配Deployment类型工作负载下的所有容器。
	// - Deployment: 无状态
	// - StatefulSet: 有状态
	// - DaemonSet: 守护进程
	// - Job: 任务
	// - CronJob: 定时任务
	WorkloadType string `json:"WorkloadType,omitempty" enums:"Deployment,StatefulSet,DaemonSet,Job,CronJob"`
	// 通过工作负载名称指定采集的容器，支持正则匹配。例如设置为"WorkloadNameRegex":"^(nginx-deployment.*)$"，表示匹配以nginx-deployment开头的工作负载下的所有容器。
	WorkloadNameRegex string `json:"WorkloadNameRegex,omitempty"`
	// 通过Pod Label白名单指定待采集的容器。如果您要设置Pod Label白名单，那么Key必填，Value可选填。
	// - 如果Value为空，则Pod Label中包含Key的Pod下的容器都匹配。
	// - 如果Value不为空，则Pod Label中包含Key=Value的Pod下的容器才匹配。
	// - Value默认为字符串匹配，即只有Value和Pod Label的值完全相同才会匹配。如果该值以^开头并且以$结尾，则为正则匹配。例如设置Key为app，设置Value为^(test1|test2)$，表示匹配Pod Label中包含app:test1、app:test2的Pod下的容器。
	// 多个白名单之间为或关系，即只要Pod的Label满足任一白名单即可匹配。
	PodNameRegex string `json:"PodNameRegex,omitempty"`
	// 通过Pod Label黑名单排除不采集的容器。如果您要设置Pod Label黑名单，那么Key必填，Value可选填。
	// - 如果Value为空，则Pod Label中包含Key的Pod下的容器都被排除。
	// - 如果Value不为空，则Pod Label中包含Key=Value的Pod下的容器才会被排除。
	// - Value默认为字符串匹配，即只有Value和Pod Label的值完全相同才会匹配。如果该值以^开头并且以$结尾，则为正则匹配。例如设置Key为app，设置Value为^(test1|test2)$，表示匹配Pod Label中包含app:test1、app:test2的Pod下的容器。
	// 多个黑名单之间为或关系，即只要容器Label满足任一黑名单对即可被排除。
	IncludePodLabelRegex map[string]string `json:"IncludePodLabelRegex,omitempty"`
	// 通过Pod名称指定待采集的容器，支持正则匹配。例如设置为"PodNameRegex":"^(nginx-log-demo.*)$"，表示匹配以nginx-log-demo开头的Pod下的所有容器。
	ExcludePodLabelRegex map[string]string `json:"ExcludePodLabelRegex,omitempty"`
	// 设置Kubernetes Label日志标签后，日志服务将在日志中新增Kubernetes Label相关字段。例如设置LabelKey为app，设置LabelValue为k8s_label_app，当Kubernetes中包含Label app=serviceA时，会将该信息添加到日志中，即添加字段__tag__k8s_label_app__: serviceA。
	LabelTag map[string]string `json:"LabelTag,omitempty"`
}

type CreateRuleResponse struct {
	CommonResponse
	// 采集配置的ID
	RuleID string `json:"RuleId" example:"faf6d529-e75e-457f-a23a-9c4203a6dff3xx"`
}

type DeleteRuleRequest struct {
	CommonRequest
	// 采集配置的ID。
	RuleID string `json:"RuleId" binding:"required" example:"faf6d529-e75e-457f-a23a-9c4203a6dff3xx"`
}

type ModifyRuleRequest struct {
	CommonRequest
	// 采集配置的ID。
	RuleID string `json:"RuleId" binding:"required" example:"faf6d529-e75e-457f-a23a-9c4203a6dff3xx"`
	// 采集配置的名称:
	// - 在一个Topic中唯一;
	// - 长度限制为3-63个字符，只能包含小写字母、数字和短划线（-），必须以小写字母或者数字开头和结尾。
	RuleName *string `json:"RuleName" example:"testName"`
	// 采集路径列表:
	// - 最多能够创建10个采集路径;
	// - 采集路径必须是一个绝对路径;
	// 当InputType=0时:
	// - 采集路径支持完整匹配和通配符模式匹配，通配符只支持"**"、"*"、"?"，但是最多只能配置1个"**"通配符.
	// 当InputType=1时:
	// - 不允许配置采集路径列表
	// 当InputType=2时:
	// - 采集路径支持完整匹配和通配符模式匹配，通配符只支持“*”、“?”，不支持“**”。
	Paths *[]string `json:"Paths" example:"[\"/data/nginx/log/**/access.log\"]"`
	// 采集的日志类型:
	// - minimalist_log代表极简日志；
	// - delimiter_log代表分隔符格式日志；
	// - json_log代表json格式日志；
	// - multiline_log代表多行日志；
	// - fullregex_log代表完整正则日志；
	// 只要修改提取规则，无论是否修改采集的日志类型，采集的日志类型必须配置。
	LogType *string `json:"LogType" enums:"minimalist_log,delimiter_log,json_log,multiline_log,fullregex_log" example:"delimiter_log"`
	// 提取规则: 如果配置非minimalist_log或者非json_log的采集的日志类型，那么必须同时配置提取规则。
	ExtractRule *ExtractRule `json:"ExtractRule"`
	// 采集黑名单列表: 最多能够创建10个采集黑名单。
	// 当InputType=0时：
	// - 当Type是Path时，Value表示一个目录。支持完整匹配和通配符模式匹配，通配符只支持“*”、“?”，不支持“**”通配符。
	// - 当Type是File时，Value表示一个文件名称。支持完整匹配和通配符模式匹配，通配符只支持“**”、“*”、“?”，但是最多只能配置1个“**”通配符。
	// 当InputType=1时：
	// - 不允许配置采集黑名单列表。
	// 当InputType=2时：
	// - 当Type是Path时，Value表示一个目录。仅支持完整匹配，不支持通配符模式匹配。
	// - 当Type是File时，Value表示一个文件名称。目录部分仅支持完整匹配，不支持通配符模式匹配。文件名称部分支持完整匹配和通配符模式匹配，通配符只支持“*”、“?”。
	ExcludePaths *[]ExcludePath `json:"ExcludePaths"`
	// 用户自定义的采集规则。
	UserDefineRule *UserDefineRule `json:"UserDefineRule"`
	// 日志样例: 日志样例的长度必须不大于3000个字符
	LogSample *string `json:"LogSample" example:"xxxxxx"`
	// 采集类型：
	// - 0：宿主机日志文件
	// - 1：k8s容器标准输出
	// - 2：k8s容器内日志文件
	// 只要修改容器采集规则，无论是否修改采集的类型，采集的类型必须配置。
	InputType *int `json:"InputType"`
	// 容器采集规则
	ContainerRule *ContainerRule `json:"ContainerRule"`
}

type DescribeRuleRequest struct {
	CommonRequest
	RuleID string `json:"RuleId"`
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

type RuleInfo struct {
	TopicID        string         `json:"TopicId"`
	TopicName      string         `json:"TopicName"`
	RuleID         string         `json:"RuleId"`
	RuleName       string         `json:"RuleName"`
	Paths          []string       `json:"Paths"`
	LogType        string         `json:"LogType" enums:"minimalist_log,delimiter_log,json_log,multiline_log,fullregex_log"`
	ExtractRule    ExtractRule    `json:"ExtractRule"`
	ExcludePaths   []ExcludePath  `json:"ExcludePaths"`
	UserDefineRule UserDefineRule `json:"UserDefineRule"`
	LogSample      string         `json:"LogSample"`
	InputType      int            `json:"InputType"`
	ContainerRule  ContainerRule  `json:"ContainerRule"`
	CreateTime     string         `json:"CreateTime"`
	ModifyTime     string         `json:"ModifyTime"`
}

type HostGroupInfo struct {
	HostGroupID                   string `json:"HostGroupId"`
	HostGroupName                 string `json:"HostGroupName"`
	HostGroupType                 string `json:"HostGroupType" enums:"IP,Label"`
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
}

type DescribeRulesRequest struct {
	CommonRequest
	//必选
	ProjectID string `json:"ProjectId"`
	// 分页查询时的页码。默认为1，即从第一页数据开始返回。

	//可选
	PageNumber int `json:"PageNumber"`
	// 分页大小。默认为20，最大为100。
	PageSize int `json:"PageSize"`
	// 采集配置的ID关键词，支持模糊搜索。
	TopicID *string `json:"TopicId"`
	// 日志主题的名称关键词，支持模糊搜索。
	TopicName *string `json:"TopicName"`
	// 日志主题的ID关键词，支持模糊搜索。
	RuleID *string `json:"RuleId"`
	// 采集配置的名称关键词，支持模糊搜索。
	RuleName *string `json:"RuleName"`
}

type CreateHostGroupRequest struct {
	CommonRequest
	// 机器组的名称。
	HostGroupName string `json:"HostGroupName" binding:"required" example:"mgn1"`
	// 机器组的类型: IP表示机器IP; Label表示机器标识。
	HostGroupType string `json:"HostGroupType" binding:"required" example:"IP" enums:"IP,Label"`
	// 机器标识符。
	HostIdentifier *string `json:"HostIdentifier" example:"label"`
	// 机器IP列表。
	HostIPList *[]string `json:"HostIpList" example:"[\"192.168.0.1\", \"127.0.0.1\"]"`
	// 机器组服务器中安装的LogCollector是否开启自动升级功能。
	AutoUpdate *bool `json:",omitempty"`
	// LogCollector自动升级的开始时间。
	UpdateStartTime *string `json:",omitempty" example:"00:00"`
	// LogCollector自动升级的结束时间。
	UpdateEndTime *string `json:",omitempty" example:"02:00"`
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
	// 机器组的Id。
	HostGroupID string `json:"HostGroupId" binding:"required" example:"c7e0e442-19bf-4fb3-b547-5992fb8b****"`
}

type ApplyRuleToHostGroupsRequest struct {
	CommonRequest
	// 采集配置的ID。
	RuleID string `json:"RuleId" binding:"required" example:"faf6d529-e75e-457f-a23a-9c4203a6dff3xx"`
	// 机器组的ID列表。
	HostGroupIDs []string `json:"HostGroupIds" binding:"required" example:"[\"bcb20377-7e48-4e90-968b-922eb259****\",\"6622a14b-770f-4171-a385-3b68486f****\"]"`
}

type DeleteRuleFromHostGroupsRequest struct {
	CommonRequest
	// 采集配置的ID。
	RuleID string `json:"RuleId" binding:"required" example:"faf6d529-e75e-457f-a23a-9c4203a6dff3xx"`
	// 机器组的ID列表。
	HostGroupIDs []string `json:"HostGroupIds" binding:"required" example:"[\"bcb20377-7e48-4e90-968b-922eb259****\",\"6622a14b-770f-4171-a385-3b68486f****\"]"`
}

type ModifyHostGroupRequest struct {
	CommonRequest
	// 机器组Id。
	HostGroupID string `json:"HostGroupId" binding:"required" example:"0fdaa6b6-3c9f-424c-8664-fc0d222c****"`
	// 机器组的名称。
	HostGroupName *string `json:"HostGroupName" example:"mgn1"`
	// 机器组的类型：IP表示机器IP；Label表示机器标识。
	HostGroupType *string `json:"HostGroupType" example:"IP" enums:"IP,Label"`
	// 机器IP列表。
	HostIPList *[]string `json:"HostIpList" example:"[\"192.168.0.1\", \"127.0.0.1\"]"`
	// 机器标识符
	HostIdentifier *string `json:"HostIdentifier" example:"label"`
	// 机器组服务器中安装的LogCollector是否开启自动升级功能。
	AutoUpdate *bool `json:",omitempty"`
	// LogCollector自动升级的开始时间。
	UpdateStartTime *string `json:",omitempty" example:"00:00"`
	// LogCollector自动升级的结束时间。
	UpdateEndTime *string `json:",omitempty" example:"02:00"`
}

type DescribeHostGroupRequest struct {
	CommonRequest
	HostGroupID string
}

type DescribeHostGroupResponse struct {
	CommonResponse
	HostGroupHostsRulesInfo *HostGroupHostsRulesInfo `json:"HostGroupHostsRulesInfo"`
}

// HostGroupHostsRulesInfo 展示层对象
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
	HostGroupID    *string
	HostGroupName  *string
	HostIdentifier *string `json:",omitempty"`
}

type DescribeHostGroupsResponse struct {
	CommonResponse
	Total                    int64                      `json:"Total"`
	HostGroupHostsRulesInfos []*HostGroupHostsRulesInfo `json:"HostGroupHostsRulesInfos"`
}

type DescribeHostsRequest struct {
	CommonRequest
	HostGroupID     string
	PageNumber      int
	PageSize        int
	IP              *string
	HeartbeatStatus *int
}

type DescribeHostsResponse struct {
	CommonResponse
	Total     int64       `json:"Total"`
	HostInfos []*HostInfo `json:"HostInfos"`
}

type DeleteHostRequest struct {
	CommonRequest
	// 机器组的Id。
	HostGroupID string `json:"HostGroupId" binding:"required" example:"c7e0e442-19bf-4fb3-b547-5992fb8b****"`
	// 机器的IP。
	IP string `json:"Ip" binding:"required" example:"1.1.1.1"`
}

type DescribeHostGroupRulesRequest struct {
	CommonRequest
	HostGroupID string
	PageNumber  int
	PageSize    int
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

type ModifyHostGroupsAutoUpdateResponse struct {
	CommonResponse
}

type CreateAlarmRequest struct {
	CommonRequest
	// 告警策略名称。
	// - 同一个日志项目下，告警策略名称不可重复。
	// - 只能包括小写字母、数字、-
	// - 必须以小写字母或者数字开头和结尾。
	// - 长度为3-63字符。
	AlarmName string `json:"AlarmName" binding:"required" example:"test-alarm"`
	// 告警策略所属的ProjectId
	ProjectID string `json:"ProjectId" binding:"required"`
	// 是否开启告警策略。默认值为true
	Status *bool `json:"Status" example:"true"`
	// 查询或者分析命令，支持1-3条
	QueryRequest QueryRequests `json:"QueryRequest" binding:"required" `
	// 查询或分析请求的执行周期
	RequestCycle RequestCycle `json:"RequestCycle" binding:"required"`
	// 告警触发条件，参考告警条件表达式：https://bytedance.feishu.cn/wiki/wikcnxhR5kkc31qgQGdMiOPuFle
	Condition string `json:"Condition" binding:"required"`
	// 持续周期。持续满足触发条件TriggerPeriod个周期后，再进行告警；最小值为1，最大值为10，默认为1。
	TriggerPeriod int `json:"TriggerPeriod" binding:"required"`
	// 告警重复的周期。单位是分钟。取值范围是0~1440。
	AlarmPeriod int `json:"AlarmPeriod" binding:"required"`
	// 告警对应的通知列表
	AlarmNotifyGroup []string `json:"AlarmNotifyGroup" binding:"required"`
	// 用户自定义告警内容
	UserDefineMsg *string `json:"UserDefineMsg"`
}

type QueryRequests []QueryRequest

type RequestCycle struct {
	// 可选值：Period-周期执行；Fixed-定期执行。
	Type string `json:"Type" binding:"required" enums:"Period,Fixed"`
	// 执行的周期，或者定制执行的时间节点。单位为分钟，取值范围为1~1440。
	Time int `json:"Time"`
}

type QueryRequest struct {
	CommonRequest
	// 查询语句，支持的最大长度为1024。
	Query string `json:"Query" gorm:"Query" binding:"required"`
	// 告警对象序号；从1开始递增。
	Number uint8 `json:"Number" gorm:"Number" binding:"required"`
	// 告警策略执行的topicId。
	TopicID string `json:"TopicId" gorm:"-" binding:"required"`
	// 告警策略执行的TopicName。
	TopicName string `json:"TopicName,omitempty" gorm:"-"`
	// 查询范围起始时间相对当前的历史时间，单位为分钟，取值为非正，最大值为0，最小值为-1440。
	StartTimeOffset int `json:"StartTimeOffset" gorm:"StartTimeOffset" binding:"required"`
	// 查询范围终止时间相对当前的历史时间，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。
	EndTimeOffset int `json:"EndTimeOffset" gorm:"EndTimeOffset" binding:"required"`
}

type CreateAlarmResponse struct {
	CommonResponse
	AlarmID string `json:"AlarmId"`
}

type DeleteAlarmRequest struct {
	CommonRequest
	// 创建时返回的告警策略id
	AlarmID string `json:"AlarmId" binding:"required"`
}

type ModifyAlarmRequest struct {
	CommonRequest
	// 告警策略Id
	AlarmID string `json:"AlarmId" binding:"required"`
	// 告警策略名称。
	// - 同一个日志项目下，告警策略名称不可重复。
	// - 只能包括小写字母、数字、-
	// - 必须以小写字母或者数字开头和结尾。
	// - 长度为3-63字符。
	AlarmName *string `json:"AlarmName"`
	// 是否开启告警策略。默认值为true
	Status *bool `json:"Status" example:"true"`
	// 查询或者分析命令，支持1-3条
	QueryRequest *QueryRequests `json:"QueryRequest"`
	// 查询或分析请求的执行周期
	RequestCycle *RequestCycle `json:"RequestCycle"`
	// 告警触发条件
	Condition *string `json:"Condition"`
	// 持续周期。持续满足触发条件TriggerPeriod个周期后，再进行告警；最小值为1，最大值为10，默认为1。
	TriggerPeriod *int `json:"TriggerPeriod"`
	// 告警重复的周期。单位是分钟。取值范围是10~1440。
	// 一个手机号码每天最多收到50条告警短信，且与火山引擎其他消息通知服务共用短信额度
	AlarmPeriod *int `json:"AlarmPeriod"`
	// 告警对应的通知列表
	AlarmNotifyGroup *[]string `json:"AlarmNotifyGroup"`
	// 用户自定义告警内容
	UserDefineMsg *string `json:"UserDefineMsg"`
}

type DescribeAlarmsRequest struct {
	CommonRequest
	ProjectID     string
	TopicID       *string
	TopicName     *string
	AlarmName     *string
	AlarmPolicyID *string
	Status        *bool
	PageNumber    int
	PageSize      int
}

type DescribeAlarmsResponse struct {
	CommonResponse
	Total         int         `json:"Total"`
	AlarmPolicies []QueryResp `json:"Alarms"`
}

type QueryResp struct {
	AlarmID          string             `json:"AlarmId"`
	AlarmName        string             `json:"AlarmName"`
	ProjectID        string             `json:"ProjectId"`
	Status           bool               `json:"Status"`
	QueryRequest     []QueryRequest     `json:"QueryRequest"`
	RequestCycle     RequestCycle       `json:"RequestCycle"`
	Condition        string             `json:"Condition"`
	TriggerPeriod    int                `json:"TriggerPeriod"`
	AlarmPeriod      int                `json:"AlarmPeriod"`
	AlarmNotifyGroup []NotifyGroupsInfo `json:"AlarmNotifyGroup"`
	UserDefineMsg    string             `json:"UserDefineMsg"`
	CreateTimestamp  string             `json:"CreateTime"`
	ModifyTimestamp  string             `json:"ModifyTime"`
}

type NotifyGroupsInfo struct {
	GroupName       string      `json:"AlarmNotifyGroupName"`
	NotifyGroupID   string      `json:"AlarmNotifyGroupId"`
	NoticeType      NoticeTypes `json:"NotifyType"`
	Receivers       Receivers   `json:"Receivers"`
	CreateTimestamp string      `json:"CreateTime"`
	ModifyTimestamp string      `json:"ModifyTime"`
}

type NoticeTypes []NoticeType

type NoticeType string

type Receivers []Receiver

type Receiver struct {
	// 接受者类型。可选值：User-用户ID;暂不支持其余接收者类型。
	ReceiverType ReveiverType `json:"ReceiverType" binding:"required"`
	// 接收者的名字。这里前端通过list所有用户选择用户，不涉及该限制长度1~64，仅支持英文、数字、下划线、"."、"-"、"@"
	ReceiverNames []string `json:"ReceiverNames" binding:"required"`
	// 通知接收渠道。
	ReceiverChannels []ReceiverChannel `json:"ReceiverChannels" binding:"required" enums:"Email,Sms"`
	// 允许接收信息的开始时间。
	StartTime string `json:"StartTime" binding:"required"`
	// 允许接收信息的开始时间。
	EndTime string `json:"EndTime" binding:"required"`
	// 如果选择了飞书渠道，则必填
	Webhook string `json:",omitempty"`
}

type ReveiverType string

type ReceiverChannel string

type CreateAlarmNotifyGroupRequest struct {
	CommonRequest
	// 告警通知组名称。
	// - 同一个账户下，通知组名称不可重复。
	// - 只能包括小写字母、数字、-。
	// - 必须以小写字母或者数字开头和结尾。
	// - 长度为3-63字符。
	GroupName string `json:"AlarmNotifyGroupName" binding:"required" example:"test-alarm-nofify"`
	// 告警通知的类型。可选值，选择一个或者多个：Trigger-告警触发;Recovery-告警恢复。
	NoticeType NoticeTypes `json:"NotifyType" binding:"required"`
	// 告警对应的通知列表，最少1一个，最大支持10个。
	Receivers Receivers `json:"Receivers" binding:"required"`
}

type CreateAlarmNotifyGroupResponse struct {
	CommonResponse
	NotifyGroupID string `json:"AlarmNotifyGroupId"`
}

type DeleteAlarmNotifyGroupRequest struct {
	CommonRequest
	// 创建时返回的NotifyGroupId
	AlarmNotifyGroupID string `json:"AlarmNotifyGroupId" binding:"required"`
}

type DeleteAlarmNotifyGroupResponse struct {
}

type ModifyAlarmNotifyGroupRequest struct {
	CommonRequest
	// 通知组id
	AlarmNotifyGroupID string `json:"AlarmNotifyGroupId" binding:"required"`
	// 告警通知组名称。
	// - 同一个账户下，通知组名称不可重复。
	// - 只能包括小写字母、数字、-。
	// - 必须以小写字母或者数字开头和结尾。
	// - 长度为3-63字符。
	AlarmNotifyGroupName *string `json:"AlarmNotifyGroupName"`
	// 告警通知的类型。可选值，选择一个或者多个：Trigger-告警触发;Recovery-告警恢复。
	NoticeType *NoticeTypes `json:"NotifyType" enums:"Trigger,Recovery"`
	// 告警对应的通知列表。
	Receivers *Receivers `json:"Receivers"`
}

type DescribeAlarmNotifyGroupsRequest struct {
	CommonRequest
	GroupName     *string
	NotifyGroupID *string
	UserName      *string
	PageNumber    int
	PageSize      int
}

type DescribeAlarmNotifyGroupsResponse struct {
	CommonResponse
	Total             int64               `json:"Total"`
	AlarmNotifyGroups []*NotifyGroupsInfo `json:"AlarmNotifyGroups"`
}

type CreateDownloadTaskRequest struct {
	CommonRequest
	TopicID     string `json:"TopicId"`
	TaskName    string
	Query       string
	StartTime   int64
	EndTime     int64
	Compression string
	DataFormat  string
	Limit       int
	Sort        string
}
type CreateDownloadTaskResponse struct {
	CommonResponse
	TaskId string
}

type DescribeDownloadTasksRequest struct {
	CommonRequest
	TopicID    string `json:"-"`
	PageNumber *int   `json:"-"`
	PageSize   *int   `json:"-"`
}
type DownloadTaskResp struct {
	TaskId      string
	TaskName    string
	TopicId     string
	Query       string
	StartTime   string
	EndTime     string
	LogCount    int64
	LogSize     int64
	Compression string
	DataFormat  string
	TaskStatus  string
	CreateTime  string
}
type DescribeDownloadTasksResponse struct {
	CommonResponse
	Tasks []*DownloadTaskResp
	Total int64
}

type DescribeDownloadUrlRequest struct {
	CommonRequest
	TaskId string `json:"-"`
}
type DescribeDownloadUrlResponse struct {
	CommonResponse
	DownloadUrl string
}

type WebTracksRequest struct {
	CommonRequest
	TopicID      string `json:"-"`
	ProjectID    string `json:"-"`
	CompressType string `json:"-"`
	Logs         []map[string]string
	Source       string `json:"Source,omitempty"`
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
type DescribeHistogramResponse struct {
	CommonResponse
	HistogramInfos []HistogramInfo `json:"Histogram"`
	Interval       int64           `json:"Interval"`
	TotalCount     int64           `json:"TotalCount"`
	ResultStatus   string          `json:"ResultStatus"`
}

type OpenKafkaConsumerRequest struct {
	CommonRequest
	TopicID string `json:"TopicId"`
}
type OpenKafkaConsumerResponse struct {
	CommonResponse
}
type CloseKafkaConsumerRequest struct {
	CommonRequest
	TopicID string `json:"TopicId"`
}
type CloseKafkaConsumerResponse struct {
	CommonResponse
}
type DescribeKafkaConsumerRequest struct {
	CommonRequest
	TopicID string `json:"-"`
}
type DescribeKafkaConsumerResponse struct {
	CommonResponse
	AllowConsume bool
	ConsumeTopic string
}

type CreateConsumerGroupRequest struct {
	// ProjectID ConsumerGroup所绑定ProjectID
	ProjectID string `json:"ProjectID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// TopicID ConsumerGroup所要消费的TopicID
	TopicIDList []string `json:"TopicIDList" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// ConsumerGroupName 名称，在一个topic下唯一
	ConsumerGroupName string `json:"ConsumerGroupName" binding:"required" example:"consumerGroupA"`
	// TTL ConsumerGroup的心跳存活时间
	HeartbeatTTL int `json:"HeartbeatTTL" binding:"required" example:"3"`
	// OrderedConsume 消费是否严格保序，与shard分裂相关
	OrderedConsume bool `json:"OrderedConsume" binding:"required" example:"false"`
}

type CreateConsumerGroupResponse struct {
	CommonResponse
}

type DeleteConsumerGroupRequest struct {
	// ProjectID ConsumerGroup所绑定project
	ProjectID string `json:"ProjectID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// ConsumerGroupName 名称，在一个topic下唯一
	ConsumerGroupName string `json:"ConsumerGroupName" binding:"required" example:"consumerGroupA"`
}

type DeleteConsumerGroupResponse struct {
	CommonResponse
}

type DescribeConsumerGroupsRequest struct {
	// ProjectID ConsumerGroup所绑定ProjectID
	ProjectID string `json:"ProjectID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`

	PageNumber int
	PageSize   int
}

type ConsumerGroupResp struct {
	ProjectID         string `json:"ProjectID"`
	ConsumerGroupName string `json:"ConsumerGroupName"`
	HeartbeatTTL      int    `json:"HeartbeatTTL"`
	OrderedConsume    bool   `json:"OrderedConsume"`
}

type DescribeConsumerGroupsResponse struct {
	CommonResponse
	ConsumerGroups []*ConsumerGroupResp `json:"ConsumerGroups"`
}

type ModifyConsumerGroupRequest struct {
	// ProjectID ConsumerGroup所绑定ProjectID
	ProjectID string `json:"ProjectID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// TopicID ConsumerGroup所要消费的TopicID
	TopicIDList *[]string `json:"TopicIDList" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// ConsumerGroupName 名称，在一个topic下唯一
	ConsumerGroupName string `json:"ConsumerGroupName" binding:"required" example:"consumerGroupA"`
	// TTL ConsumerGroup的心跳存活时间
	HeartbeatTTL *int `json:"HeartbeatTTL" example:"3"`
	// OrderedConsume 消费是否严格保序，与shard分裂相关
	OrderedConsume *bool `json:"OrderedConsume" example:"false"`
}

type ModifyConsumerGroupResponse struct {
	CommonResponse
}

type ConsumerHeartbeatRequest struct {
	// ProjectID ConsumerGroup所绑定ProjectID
	ProjectID string `json:"ProjectID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// ConsumerGroupName 名称，在一个topic下唯一
	ConsumerGroupName string `json:"ConsumerGroupName" binding:"required" example:"consumerGroupA"`
	// ConsumerName Consumer名称，在一个consumerGroup中唯一
	ConsumerName string `json:"ConsumerName" binding:"required" example:"consumerA"`
}

type ConsumeShard struct {
	TopicID string `yaml:"TopicID"`
	ShardID int    `yaml:"ShardID"`
}

type ConsumerHeartbeatResponse struct {
	CommonResponse
	Shards []*ConsumeShard `json:"Shards"`
}

type DescribeCheckPointRequest struct {
	// ProjectID consumerGroup所绑定的projectID
	ProjectID string `json:"ProjectID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// TopicID ConsumerGroup所绑定topicID
	TopicID string `json:"TopicID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// ConsumerGroupName 名称，在一个topic下唯一
	ConsumerGroupName string `json:"ConsumerGroupName" example:"consumerGroupA"`
	// ShardID 需要获取checkpoint的shardID
	ShardID int `json:"ShardID" binding:"required" example:"1"`
}

type DescribeCheckPointResponse struct {
	CommonResponse
	ShardID    int32  `json:"ShardID"`
	Checkpoint string `json:"Checkpoint"`
	UpdateTime int    `json:"UpdateTime"`
	Consumer   string `json:"Consumer"`
}

type ModifyCheckPointRequest struct {
	// ProjectID consumerGroup所绑定的projectID
	ProjectID string `json:"ProjectID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// TopicID ConsumerGroup所绑定topicID
	TopicID string `json:"TopicID" binding:"required" example:"4a9bd4bd-53f1-43ff-b88a-64ee1be5****"`
	// ConsumerGroupName 名称，在一个topic下唯一
	ConsumerGroupName string `json:"ConsumerGroupName" example:"consumerGroupA"`
	// ShardID 需要获取checkpoint的shardID
	ShardID int `json:"ShardID" binding:"required" example:"1"`
	// Checkpoint 想要设置的checkpoint
	Checkpoint string `json:"Checkpoint" binding:"required" example:"MTUyNDE1NTM3OTM3MzkwODQ5Ng=="`
}

type ModifyCheckPointResponse struct {
	CommonResponse
}

// FillRequestId 成功返回填充requestId
func (response *CommonResponse) FillRequestId(httpResponse *http.Response) {
	response.RequestID = httpResponse.Header.Get(RequestIDHeader)
}
