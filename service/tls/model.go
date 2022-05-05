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

type CreateProjectRequest struct {
	ProjectName string
	Description string
	Region      string
}

type CreateProjectResponse struct {
	CommonResponse
	ProjectId string `json:"ProjectId"`
}

type DeleteProjectRequest struct {
	ProjectID string `json:"ProjectId"`
}

type DescribeProjectRequest struct {
	ProjectID string `json:"ProjectId"`
}

type DescribeProjectResponse struct {
	CommonResponse
	ProjectInfo
}

type DescribeProjectsRequest struct {
	ProjectName string
	ProjectID   string
	PageNumber  int
	PageSize    int
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
	TopicID string `json:"TopicID"`
}

type DeleteTopicRequest struct {
	TopicID string
}

type ModifyTopicRequest struct {
	TopicID     string  `json:"TopicId"`
	TopicName   *string `json:"TopicName"`
	Ttl         *uint16 `json:"Ttl"`
	Description *string `json:"Description"`
}

type DescribeTopicRequest struct {
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
}

type DescribeTopicsRequest struct {
	ProjectID  string
	PageNumber int
	PageSize   int
	TopicName  string
	TopicID    string
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
	TopicID string
}

type DescribeIndexRequest struct {
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
	TopicID  string          `json:"TopicId"`
	FullText *FullTextInfo   `json:"FullText"`
	KeyValue *[]KeyValueInfo `json:"KeyValue"`
}

type AnalysisResult struct {
	Schema []string                 `json:"schema"`
	Data   []map[string]interface{} `json:"data"`
}

type SearchLogsRequest struct {
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
	} `json:"Shards"`

	Total int `json:"Total"`
}

type PutLogsRequest struct {
	TopicID      string
	HashKey      string
	CompressType string
	LogBody      *pb.LogGroupList
}

type DescribeCursorRequest struct {
	TopicID string
	ShardID int
	From    string
}

type DescribeCursorResponse struct {
	CommonResponse
	Cursor string `json:"cursor"`
}

type ConsumeLogsRequest struct {
	TopicID       string
	ShardID       int
	Cursor        string
	EndCursor     *string
	LogGroupCount *int
	Compression   *string
}

type ConsumeLogsResponse struct {
	CommonResponse
	Cursor string
	Count  int
	Logs   *pb.LogGroupList
}

// FillRequestId 成功返回填充requestId
func (response *CommonResponse) FillRequestId(httpResponse *http.Response) {
	response.RequestID = httpResponse.Header.Get(RequestIDHeader)
}
