package tls

import (
	"context"
	"net/http"
	"sync"
	"time"
)

func NewClient(endpoint, accessKeyID, accessKeySecret, securityToken, region string) Client {
	apiVersion := APIVersion3
	return &LsClient{
		Client:          defaultHttpClient,
		Endpoint:        endpoint,
		accessLock:      &sync.RWMutex{},
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Region:          region,
		APIVersion:      apiVersion,
	}
}

type Client interface {
	GetHttpClient() *http.Client
	SetHttpClient(client *http.Client) error
	ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string)
	SetTimeout(timeout time.Duration)
	SetAPIVersion(version string)
	SetCustomUserAgent(customUserAgent string)

	PutLogs(request *PutLogsRequest) (response *CommonResponse, err error)
	PutLogsV2(request *PutLogsV2Request) (response *CommonResponse, err error)
	DescribeCursor(request *DescribeCursorRequest) (*DescribeCursorResponse, error)
	ConsumeLogs(request *ConsumeLogsRequest) (*ConsumeLogsResponse, error)
	DescribeLogContext(request *DescribeLogContextRequest) (*DescribeLogContextResponse, error)

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
	SearchLogsV2(request *SearchLogsRequest) (*SearchLogsResponse, error)

	DescribeShards(request *DescribeShardsRequest) (*DescribeShardsResponse, error)
	ManualShardSplit(ctx context.Context, request *ManualShardSplitRequest) (*ManualShardSplitResponse, error)

	CreateRule(request *CreateRuleRequest) (*CreateRuleResponse, error)
	DeleteRule(request *DeleteRuleRequest) (*CommonResponse, error)
	ModifyRule(request *ModifyRuleRequest) (*CommonResponse, error)
	// Deprecated: Please use DescribeRuleV2 instead.
	DescribeRule(request *DescribeRuleRequest) (*DescribeRuleResponse, error)
	DescribeRules(request *DescribeRulesRequest) (*DescribeRulesResponse, error)
	ApplyRuleToHostGroups(request *ApplyRuleToHostGroupsRequest) (*CommonResponse, error)
	DeleteRuleFromHostGroups(request *DeleteRuleFromHostGroupsRequest) (*CommonResponse, error)
	DescribeRuleV2(request *DescribeRuleRequestV2) (*DescribeRuleResponseV2, error)
	DescribeBoundHostGroups(request *DescribeBoundHostGroupsRequest) (*DescribeBoundHostGroupsResponse, error)

	CreateHostGroup(request *CreateHostGroupRequest) (*CreateHostGroupResponse, error)
	DeleteHostGroup(request *DeleteHostGroupRequest) (*CommonResponse, error)
	ModifyHostGroup(request *ModifyHostGroupRequest) (*CommonResponse, error)
	// Deprecated: Please use DescribeHostGroupV2 instead.
	DescribeHostGroup(request *DescribeHostGroupRequest) (*DescribeHostGroupResponse, error)
	// Deprecated: Please use DescribeHostGroupsV2 instead.
	DescribeHostGroups(request *DescribeHostGroupsRequest) (*DescribeHostGroupsResponse, error)
	DescribeHosts(request *DescribeHostsRequest) (*DescribeHostsResponse, error)
	DeleteHost(request *DeleteHostRequest) (*CommonResponse, error)
	DescribeHostGroupRules(request *DescribeHostGroupRulesRequest) (*DescribeHostGroupRulesResponse, error)
	ModifyHostGroupsAutoUpdate(request *ModifyHostGroupsAutoUpdateRequest) (*ModifyHostGroupsAutoUpdateResponse, error)
	DeleteAbnormalHosts(request *DeleteAbnormalHostsRequest) (*CommonResponse, error)
	DescribeHostGroupV2(request *DescribeHostGroupRequestV2) (*DescribeHostGroupResponseV2, error)
	DescribeHostGroupsV2(request *DescribeHostGroupsRequestV2) (*DescribeHostGroupsResponseV2, error)

	CreateAlarm(request *CreateAlarmRequest) (*CreateAlarmResponse, error)
	DeleteAlarm(request *DeleteAlarmRequest) (*CommonResponse, error)
	DeleteAlarmContentTemplate(request *DeleteAlarmContentTemplateRequest) (*CommonResponse, error)
	ModifyAlarm(request *ModifyAlarmRequest) (*CommonResponse, error)
	ModifyAlarmWebhookIntegration(request *ModifyAlarmWebhookIntegrationReq) (*CommonResponse, error)
	DescribeAlarms(request *DescribeAlarmsRequest) (*DescribeAlarmsResponse, error)
	DescribeAlarmWebhookIntegrations(request *DescribeAlarmWebhookIntegrationsRequest) (*DescribeAlarmWebhookIntegrationsResponse, error)
	DescribeAlarmContentTemplates(request *DescribeAlarmContentTemplatesRequest) (*DescribeAlarmContentTemplatesResponse, error)
	CreateAlarmWebhookIntegration(request *CreateAlarmWebhookIntegrationRequest) (*CreateAlarmWebhookIntegrationResponse, error)
	CreateAlarmNotifyGroup(request *CreateAlarmNotifyGroupRequest) (*CreateAlarmNotifyGroupResponse, error)
	DeleteAlarmNotifyGroup(request *DeleteAlarmNotifyGroupRequest) (*CommonResponse, error)
	ModifyAlarmNotifyGroup(request *ModifyAlarmNotifyGroupRequest) (*CommonResponse, error)
	DescribeAlarmNotifyGroups(request *DescribeAlarmNotifyGroupsRequest) (*DescribeAlarmNotifyGroupsResponse, error)
	CreateAlarmContentTemplate(request *CreateAlarmContentTemplateRequest) (*CreateAlarmContentTemplateResponse, error)

	DeleteAlarmWebhookIntegration(request *DeleteAlarmWebhookIntegrationRequest) (*CommonResponse, error)
	ModifyAlarmContentTemplate(request *ModifyAlarmContentTemplateRequest) (*ModifyAlarmContentTemplateResponse, error)

	CreateDownloadTask(request *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error)
	DescribeDownloadTasks(request *DescribeDownloadTasksRequest) (*DescribeDownloadTasksResponse, error)
	DescribeDownloadUrl(request *DescribeDownloadUrlRequest) (*DescribeDownloadUrlResponse, error)
	CancelDownloadTask(request *CancelDownloadTaskRequest) (*CancelDownloadTaskResponse, error)

	WebTracks(request *WebTracksRequest) (*WebTracksResponse, error)

	OpenKafkaConsumer(request *OpenKafkaConsumerRequest) (*OpenKafkaConsumerResponse, error)
	CloseKafkaConsumer(request *CloseKafkaConsumerRequest) (*CloseKafkaConsumerResponse, error)
	DescribeKafkaConsumer(request *DescribeKafkaConsumerRequest) (*DescribeKafkaConsumerResponse, error)

	// Deprecated: use DescribeHistogramV1 instead
	DescribeHistogram(request *DescribeHistogramRequest) (*DescribeHistogramResponse, error)
	DescribeHistogramV1(request *DescribeHistogramV1Request) (*DescribeHistogramV1Response, error)

	CreateConsumerGroup(request *CreateConsumerGroupRequest) (*CreateConsumerGroupResponse, error)
	DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (*CommonResponse, error)
	DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (*DescribeConsumerGroupsResponse, error)
	ModifyConsumerGroup(request *ModifyConsumerGroupRequest) (*CommonResponse, error)

	ConsumerHeartbeat(request *ConsumerHeartbeatRequest) (*ConsumerHeartbeatResponse, error)
	DescribeCheckPoint(request *DescribeCheckPointRequest) (*DescribeCheckPointResponse, error)
	ModifyCheckPoint(request *ModifyCheckPointRequest) (*CommonResponse, error)
	ResetCheckPoint(request *ResetCheckPointRequest) (*CommonResponse, error)

	AddTagsToResource(request *AddTagsToResourceRequest) (*CommonResponse, error)
	RemoveTagsFromResource(request *RemoveTagsFromResourceRequest) (*CommonResponse, error)
	TagResources(request *TagResourcesRequest) (*TagResourcesResponse, error)
	ListTagsForResources(request *ListTagsForResourcesRequest) (*ListTagsForResourcesResponse, error)
	UntagResources(request *UntagResourcesRequest) (*CommonResponse, error)

	CreateETLTask(request *CreateETLTaskRequest) (*CreateETLTaskResponse, error)
	DeleteETLTask(request *DeleteETLTaskRequest) (*CommonResponse, error)
	ModifyETLTask(request *ModifyETLTaskRequest) (*CommonResponse, error)
	DescribeETLTask(request *DescribeETLTaskRequest) (*DescribeETLTaskResponse, error)
	DescribeETLTasks(request *DescribeETLTasksRequest) (*DescribeETLTasksResponse, error)
	ModifyETLTaskStatus(request *ModifyETLTaskStatusRequest) (*CommonResponse, error)

	CreateImportTask(request *CreateImportTaskRequest) (*CreateImportTaskResponse, error)
	DeleteImportTask(request *DeleteImportTaskRequest) (*DeleteImportTaskResponse, error)
	ModifyImportTask(request *ModifyImportTaskRequest) (*ModifyImportTaskResponse, error)
	DescribeImportTask(request *DescribeImportTaskRequest) (*DescribeImportTaskResponse, error)
	DescribeImportTasks(request *DescribeImportTasksRequest) (*DescribeImportTasksResponse, error)

	CreateShipper(request *CreateShipperRequest) (*CreateShipperResponse, error)
	DeleteShipper(request *DeleteShipperRequest) (*DeleteShipperResponse, error)
	ModifyShipper(request *ModifyShipperRequest) (*ModifyShipperResponse, error)
	DescribeShipper(request *DescribeShipperRequest) (*DescribeShipperResponse, error)
	DescribeShippers(request *DescribeShippersRequest) (*DescribeShippersResponse, error)

	// AI应用实例管理
	CreateAppInstance(request *CreateAppInstanceReq) (*CreateAppInstanceResp, error)
	DescribeAppInstances(request *DescribeAppInstancesReq) (*DescribeAppInstancesResp, error)
	DeleteAppInstance(request *DeleteAppInstanceReq) (*DeleteAppInstanceResp, error)

	// AI场景元数据管理
	CreateAppSceneMeta(request *CreateAppSceneMetaReq) (*CreateAppSceneMetaResp, error)
	DescribeAppSceneMetas(request *DescribeAppSceneMetasReq) (*DescribeAppSceneMetasResp, error)
	DeleteAppSceneMeta(request *DeleteAppSceneMetaReq) (*DeleteAppSceneMetaResp, error)
	ModifyAppSceneMeta(request *ModifyAppSceneMetaReq) (*ModifyAppSceneMetaResp, error)

	// AI对话接口
	DescribeSessionAnswer(request *DescribeSessionAnswerReq) (reader *CopilotSSEReader, err error)

	// Trace 相关接口
	DeleteTraceInstance(request *DeleteTraceInstanceRequest) (*DeleteTraceInstanceResponse, error)
	DescribeTraceInstance(request *DescribeTraceInstanceRequest) (*DescribeTraceInstanceResponse, error)
	DescribeTrace(request *DescribeTraceRequest) (*DescribeTraceResponse, error)
	DescribeTraceInstances(request *DescribeTraceInstancesRequest) (*DescribeTraceInstancesResponse, error)
	CreateTraceInstance(request *CreateTraceInstanceRequest) (*CreateTraceInstanceResponse, error)
	SearchTraces(request *SearchTracesRequest) (*SearchTracesResponse, error)

	// 账号管理相关接口
	GetAccountStatus(request *GetAccountStatusRequest) (*GetAccountStatusResponse, error)
	ActiveTlsAccount(request *ActiveTlsAccountRequest) (*ActiveTlsAccountResponse, error)
	// 定时SQL任务接口
	CreateScheduleSqlTask(request *CreateScheduleSqlTaskRequest) (*CreateScheduleSqlTaskResponse, error)
	DeleteScheduleSqlTask(request *DeleteScheduleSqlTaskRequest) (*DeleteScheduleSqlTaskResponse, error)
	DescribeScheduleSqlTask(request *DescribeScheduleSqlTaskRequest) (*DescribeScheduleSqlTaskResponse, error)
	DescribeScheduleSqlTasks(request *DescribeScheduleSqlTasksRequest) (*DescribeScheduleSqlTasksResponse, error)
	ModifyScheduleSqlTask(request *ModifyScheduleSqlTaskRequest) (*ModifyScheduleSqlTaskResponse, error)
}
