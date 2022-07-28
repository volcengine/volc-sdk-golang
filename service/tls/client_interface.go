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
	DescribeCursor(request *DescribeCursorRequest) (*DescribeCursorResponse, error)
	ConsumeLogs(request *ConsumeLogsRequest) (*ConsumeLogsResponse, error)

	CreateProject(request *CreateProjectRequest) (*CreateProjectResponse, error)
	DeleteProject(request *DeleteProjectRequest) (*CommonResponse, error)
	DescribeProject(request *DescribeProjectRequest) (*DescribeProjectResponse, error)
	DescribeProjects(request *DescribeProjectsRequest) (*DescribeProjectsResponse, error)
	ModifyProject(request *ModifyProjectRequest) (*CommonResponse, error)

	CreateTopic(request *CreateTopicRequest) (*CreateTopicResponse, error)
	DeleteTopic(request *DeleteTopicRequest) (*CommonResponse, error)
	DescribeTopic(request *DescribeTopicRequest) (*DescribeTopicResponse, error)
	DescribeTopics(request *DescribeTopicsRequest) (*DescribeTopicsResponse, error)
	ModifyTopic(request *ModifyTopicRequest) (*CommonResponse, error)

	CreateIndex(request *CreateIndexRequest) (*CreateIndexResponse, error)
	DeleteIndex(request *DeleteIndexRequest) (*CommonResponse, error)
	DescribeIndex(request *DescribeIndexRequest) (*DescribeIndexResponse, error)
	ModifyIndex(request *ModifyIndexRequest) (*CommonResponse, error)
	SearchLogs(request *SearchLogsRequest) (*SearchLogsResponse, error)

	DescribeShards(request *DescribeShardsRequest) (*DescribeShardsResponse, error)

	CreateRule(request *CreateRuleRequest) (*CreateRuleResponse, error)
	DeleteRule(request *DeleteRuleRequest) (*CommonResponse, error)
	ModifyRule(request *ModifyRuleRequest) (*CommonResponse, error)
	DescribeRule(request *DescribeRuleRequest) (*DescribeRuleResponse, error)
	DescribeRules(request *DescribeRulesRequest) (*DescribeRulesResponse, error)
	ApplyRuleToHostGroups(request *ApplyRuleToHostGroupsRequest) (*CommonResponse, error)
	DeleteRuleFromHostGroups(request *DeleteRuleFromHostGroupsRequest) (*CommonResponse, error)

	CreateHostGroup(request *CreateHostGroupRequest) (*CreateHostGroupResponse, error)
	DeleteHostGroup(request *DeleteHostGroupRequest) (*CommonResponse, error)
	ModifyHostGroup(request *ModifyHostGroupRequest) (*CommonResponse, error)
	DescribeHostGroup(request *DescribeHostGroupRequest) (*DescribeHostGroupResponse, error)
	DescribeHostGroups(request *DescribeHostGroupsRequest) (*DescribeHostGroupsResponse, error)
	DescribeHosts(request *DescribeHostsRequest) (*DescribeHostsResponse, error)
	DeleteHost(request *DeleteHostRequest) (*CommonResponse, error)
	DescribeHostGroupRules(request *DescribeHostGroupRulesRequest) (*DescribeHostGroupRulesResponse, error)

	CreateAlarm(request *CreateAlarmRequest) (*CreateAlarmResponse, error)
	DeleteAlarm(request *DeleteAlarmRequest) (*CommonResponse, error)
	ModifyAlarm(request *ModifyAlarmRequest) (*CommonResponse, error)
	DescribeAlarms(request *DescribeAlarmsRequest) (*DescribeAlarmsResponse, error)
	CreateAlarmNotifyGroup(request *CreateAlarmNotifyGroupRequest) (*CreateAlarmNotifyGroupResponse, error)
	DeleteAlarmNotifyGroup(request *DeleteAlarmNotifyGroupRequest) (*CommonResponse, error)
	ModifyAlarmNotifyGroup(request *ModifyAlarmNotifyGroupRequest) (*CommonResponse, error)
	DescribeAlarmNotifyGroups(request *DescribeAlarmNotifyGroupsRequest) (*DescribeAlarmNotifyGroupsResponse, error)
}
