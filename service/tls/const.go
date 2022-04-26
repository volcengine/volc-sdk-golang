package tls

const (
	defaultLogUserAgent = "tls-golang-sdk-v0.1.0"
	RequestIDHeader     = "x-tls-requestid"
	AgentHeader         = "User-Agent"
	ContentMd5Header    = "Content-MD5"
	ServiceName         = "TLS"

	CompressLz4  = "lz4"
	CompressNone = "none"

	ProjectUrl     = "/project"
	ListProjectUrl = "/projects"
	TopicUrl       = "/topic"
	ListTopicUrl   = "/topics"
	IndexUrl       = "/index"
	ListShardsUrl  = "/shards"
	SearchUrl      = "/search"
	PutLogsUrl     = "/put_logs"
	GetCursorUrl   = "/GetCursor"
	PullLogsUrl    = "/PullLogs"
)
