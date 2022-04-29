package tls

import (
	"sync"
	"time"
)

func NewClient(endpoint, accessKeyID, accessKeySecret, securityToken, region string) Client {
	return &LsClient{
		Endpoint:        endpoint,
		accessLock:      &sync.RWMutex{},
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Region:          region,
	}
}

type Client interface {
	ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string)
	SetTimeout(timeout time.Duration)

	PutLogs(request *PutLogsRequest) (response *CommonResponse, err error)
	DescribeCursor(request *GetCursorRequest) (*GetCursorResponse, error)
	ConsumeLogs(request *PullLogsRequest) (*PullLogsResponse, error)

	CreateProject(request *CreateProjectRequest) (*CreateProjectResponse, error)
	DeleteProject(request *DeleteProjectRequest) (*CommonResponse, error)
	DescribeProject(request *GetProjectRequest) (*GetProjectResponse, error)
	DescribeProjects(request *ListProjectRequest) (*ListProjectResponse, error)
	ModifyProject(request *UpdateProjectRequest) (*CommonResponse, error)

	CreateTopic(request *CreateTopicRequest) (*CreateTopicResponse, error)
	DeleteTopic(request *DeleteTopicRequest) (*CommonResponse, error)
	DescribeTopic(request *GetTopicRequest) (*GetTopicResponse, error)
	DescribeTopics(request *ListTopicRequest) (*ListTopicResponse, error)
	ModifyTopic(request *UpdateTopicRequest) (*CommonResponse, error)

	CreateIndex(request *CreateIndexRequest) (*CreateIndexResponse, error)
	DeleteIndex(request *DeleteIndexRequest) (*CommonResponse, error)
	DescribeIndex(request *GetIndexRequest) (*GetIndexResponse, error)
	ModifyIndex(request *UpdateIndexRequest) (*CommonResponse, error)
	SearchLogs(request *SearchIndexRequest) (*SearchIndexResponse, error)

	DescribeShards(request *ListShardRequest) (*ListShardResponse, error)
}
