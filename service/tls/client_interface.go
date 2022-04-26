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

	PutLogs(request *PutLogsRequest) (err error)
	GetCursor(request *GetCursorRequest) (*GetCursorResponse, error)
	PullLogs(request *PullLogsRequest) (*PullLogsResponse, error)

	CreateProject(request *CreateProjectRequest) (*CreateProjectResponse, error)
	DeleteProject(request *DeleteProjectRequest) error
	GetProject(request *GetProjectRequest) (*GetProjectResponse, error)
	ListProject(request *ListProjectRequest) (*ListProjectResponse, error)
	UpdateProject(request *UpdateProjectRequest) error

	CreateTopic(request *CreateTopicRequest) (*CreateTopicResponse, error)
	DeleteTopic(request *DeleteTopicRequest) error
	GetTopic(request *GetTopicRequest) (*GetTopicResponse, error)
	ListTopic(request *ListTopicRequest) (*ListTopicResponse, error)
	UpdateTopic(request *UpdateTopicRequest) error

	CreateIndex(request *CreateIndexRequest) (*CreateIndexResponse, error)
	DeleteIndex(request *DeleteIndexRequest) error
	GetIndex(request *GetIndexRequest) (*GetIndexResponse, error)
	UpdateIndex(request *UpdateIndexRequest) error
	SearchIndex(request *SearchIndexRequest) (*SearchIndexResponse, error)

	ListShard(request *ListShardRequest) (*ListShardResponse, error)
}
