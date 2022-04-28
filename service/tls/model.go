package tls

import (
	"database/sql/driver"
	"encoding/json"
	"net/http"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type CommonResponse struct {
	RequestID string `json:"requestID"`
}

type CreateProjectRequest struct {
	ProjectName string
	Description string
	Region      string
}

type CreateProjectResponse struct {
	CommonResponse
	ProjectId string `json:"project_id"`
}

type DeleteProjectRequest struct {
	ProjectID string
}

type GetProjectRequest struct {
	ProjectID string
}

type GetProjectResponse struct {
	CommonResponse
	ProjectID       string `json:"project_id"`
	ProjectName     string `json:"project_name"`
	Description     string `json:"description"`
	CreateTimestamp string `json:"create_timestamp"`
	TopicCount      int64  `json:"topic_count"`
	InnerNetDomain  string `json:"inner_net_domain"`
}

type ListProjectRequest struct {
	ProjectName string
	ProjectID   string
	Offset      int
	Limit       int
}

type ListProjectResponse struct {
	CommonResponse
	Projects []*struct {
		ProjectID       string `json:"project_id"`
		ProjectName     string `json:"project_name"`
		Description     string `json:"description"`
		CreateTimestamp string `json:"create_timestamp"`
		TopicCount      int64  `json:"topic_count"`
		InnerNetDomain  string `json:"inner_net_domain"`
	} `json:"projects"`

	Total int64 `json:"total"`
}

type UpdateProjectRequest struct {
	ProjectID   string
	ProjectName *string
	Description *string
}

type CreateTopicRequest struct {
	ProjectID   string
	TopicName   string
	Ttl         uint16
	Description string
	ShardCount  int
}

type CreateTopicResponse struct {
	CommonResponse
	TopicID string `json:"topic_id"`
}

type DeleteTopicRequest struct {
	TopicID string
}

type UpdateTopicRequest struct {
	TopicID     string  `json:"topic_id"`
	TopicName   *string `json:"topic_name"`
	Ttl         *uint16 `json:"ttl"`
	Description *string `json:"description"`
}

type GetTopicRequest struct {
	TopicID string
}

type GetTopicResponse struct {
	CommonResponse
	TopicName       string `json:"topic_name"`
	ProjectID       string `json:"project_id"`
	TopicID         string `json:"topic_id"`
	Ttl             uint16 `json:"ttl"`
	CreateTimestamp string `json:"create_timestamp"`
	ModifyTimestamp string `json:"modify_timestamp"`
	ShardCount      int32  `json:"shard_count"`
	Description     string `json:"description"`
}

type ListTopicRequest struct {
	ProjectID string
	Offset    int
	Limit     int
	TopicName string
	TopicID   string
}

type Topic struct {
	TopicName       string `json:"topic_name"`
	ProjectID       string `json:"project_id"`
	TopicID         string `json:"topic_id"`
	Ttl             uint16 `json:"ttl"`
	CreateTimestamp string `json:"create_timestamp"`
	ModifyTimestamp string `json:"modify_timestamp"`
	ShardCount      int32  `json:"shard_count"`
	Description     string `json:"description"`
}

type ListTopicResponse struct {
	CommonResponse
	Topics []*Topic `json:"topics"`
	Total  int      `json:"total"`
}

type Value struct {
	ValueType      string `json:"value_type"`
	Delimiter      string `json:"delimiter"`
	CasSensitive   bool   `json:"cas_sensitive"`
	IncludeChinese bool   `json:"include_chinese"`
	SQLFlag        bool   `json:"sql_flag"`
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
	TopicID        string
	FulltextIndex  bool
	CasSensitive   bool
	IncludeChinese bool
	Delimiter      string
	KeyValueIndex  bool
	KeyValueList   KeyValueList
}

type CreateIndexResponse struct {
	CommonResponse
	TopicID string `json:"topic_id"`
}

type DeleteIndexRequest struct {
	TopicID string
}

type GetIndexRequest struct {
	TopicID string
}

type GetIndexResponse struct {
	CommonResponse
	TopicID         string       `json:"topic_id"`
	FulltextIndex   bool         `json:"fulltext_index"`
	CasSensitive    bool         `json:"cas_sensitive"`
	Delimiter       string       `json:"delimiter"`
	IncludeChinese  bool         `json:"include_chinese"`
	CreateTimestamp string       `json:"create_timestamp"`
	UpdateTimestamp string       `json:"update_timestamp"`
	KeyValueIndex   bool         `json:"key_value_index"`
	KeyValueList    KeyValueList `json:"key_value_list"`
}

type UpdateIndexRequest struct {
	TopicID        string
	FulltextIndex  bool
	CasSensitive   bool
	IncludeChinese bool
	Delimiter      string
	KeyValueIndex  bool
	KeyValueList   KeyValueList
}

type AnalysisResult struct {
	Schema []string                 `json:"schema"`
	Data   []map[string]interface{} `json:"data"`
}

type SearchIndexRequest struct {
	TopicID   string `json:"topic_id"`
	Query     string `json:"query"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Limit     int    `json:"limit"`
	HighLight bool   `json:"high_light"`
	Context   string `json:"context"`
	Sort      string `json:"sort"`
}

type SearchIndexResponse struct {
	CommonResponse
	Status         string                   `json:"result_status"`
	HitCount       int                      `json:"hit_count"`
	Context        string                   `json:"context"`
	ListOver       bool                     `json:"list_over"`
	Analysis       bool                     `json:"analysis"`
	Logs           []map[string]interface{} `json:"logs"`
	HighLight      []string                 `json:"high_light,omitempty"`
	AnalysisResult *AnalysisResult          `json:"analysis_result"`
}

type ListShardRequest struct {
	TopicID string
	Offset  int
	Limit   int
}

type ListShardResponse struct {
	CommonResponse
	Shards []*struct {
		TopicID           string `json:"topic_id"`
		ShardID           int32  `json:"shard_id"`
		InclusiveBeginKey string `json:"inclusive_begin_key"`
		ExclusiveEndKey   string `json:"exclusive_end_key"`
		Status            string `json:"status"`
		ModifyTimestamp   string `json:"modify_timestamp"`
	} `json:"shards"`

	Total int `json:"total"`
}

type PutLogsRequest struct {
	TopicID      string
	HashKey      string
	CompressType string
	LogBody      *pb.LogGroupList
}

type GetCursorRequest struct {
	TopicID string
	ShardID int
	From    string
}

type GetCursorResponse struct {
	CommonResponse
	Cursor string `json:"cursor"`
}

type PullLogsRequest struct {
	TopicID     string
	ShardID     int
	Cursor      string
	EndCursor   *string
	Count       *int
	Compression *string
}

type PullLogsResponse struct {
	CommonResponse
	Cursor string
	Count  int
	Logs   *pb.LogGroupList
}

// FillRequestId 成功返回填充requestId
func (response *CommonResponse) FillRequestId(httpResponse *http.Response) {
	response.RequestID = httpResponse.Header.Get(RequestIDHeader)
}
