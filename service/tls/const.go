package tls

const (
	RequestIDHeader  = "x-tls-requestid"
	AgentHeader      = "User-Agent"
	ContentMd5Header = "Content-MD5"
	ServiceName      = "TLS"

	CompressLz4  = "lz4"
	CompressGz   = "gzip"
	CompressNone = "none"

	FullTextIndexKey = "__content__"

	PathCreateProject    = "/CreateProject"
	PathDescribeProject  = "/DescribeProject"
	PathDeleteProject    = "/DeleteProject"
	PathModifyProject    = "/ModifyProject"
	PathDescribeProjects = "/DescribeProjects"

	PathCreateTopic    = "/CreateTopic"
	PathDescribeTopic  = "/DescribeTopic"
	PathDeleteTopic    = "/DeleteTopic"
	PathModifyTopic    = "/ModifyTopic"
	PathDescribeTopics = "/DescribeTopics"

	PathCreateIndex   = "/CreateIndex"
	PathDescribeIndex = "/DescribeIndex"
	PathDeleteIndex   = "/DeleteIndex"
	PathModifyIndex   = "/ModifyIndex"
	PathSearchLogs    = "/SearchLogs"

	PathDescribeShards = "/DescribeShards"

	PathPutLogs            = "/PutLogs"
	PathDescribeCursor     = "/DescribeCursor"
	PathConsumeLogs        = "/ConsumeLogs"
	PathDescribeLogContext = "/DescribeLogContext"

	PathCreateRule               = "/CreateRule"
	PathDeleteRule               = "/DeleteRule"
	PathModifyRule               = "/ModifyRule"
	PathDescribeRule             = "/DescribeRule"
	PathDescribeRules            = "/DescribeRules"
	PathApplyRuleToHostGroups    = "/ApplyRuleToHostGroups"
	PathDeleteRuleFromHostGroups = "/DeleteRuleFromHostGroups"

	PathCreateHostGroup            = "/CreateHostGroup"
	PathDeleteHostGroup            = "/DeleteHostGroup"
	PathModifyHostGroup            = "/ModifyHostGroup"
	PathDescribeHostGroup          = "/DescribeHostGroup"
	PathDescribeHostGroups         = "/DescribeHostGroups"
	PathDescribeHostGroupRules     = "/DescribeHostGroupRules"
	PathModifyHostGroupsAutoUpdate = "/ModifyHostGroupsAutoUpdate"

	PathDeleteHost          = "/DeleteHost"
	PathDescribeHosts       = "/DescribeHosts"
	PathDeleteAbnormalHosts = "/DeleteAbnormalHosts"

	PathCreateAlarmNotifyGroup    = "/CreateAlarmNotifyGroup"
	PathDeleteAlarmNotifyGroup    = "/DeleteAlarmNotifyGroup"
	PathDescribeAlarmNotifyGroups = "/DescribeAlarmNotifyGroups"
	PathModifyAlarmNotifyGroup    = "/ModifyAlarmNotifyGroup"
	PathCreateAlarm               = "/CreateAlarm"
	PathDeleteAlarm               = "/DeleteAlarm"
	PathModifyAlarm               = "/ModifyAlarm"
	PathDescribeAlarms            = "/DescribeAlarms"

	PathCreateDownloadTask    = "/CreateDownloadTask"
	PathDescribeDownloadTasks = "/DescribeDownloadTasks"
	PathDescribeDownloadUrl   = "/DescribeDownloadUrl"

	PathWebTracks = "/WebTracks"

	PathDescribeHistogram   = "/DescribeHistogram"
	PathDescribeHistogramV1 = "/DescribeHistogramV1"

	PathOpenKafkaConsumer     = "/OpenKafkaConsumer"
	PathCloseKafkaConsumer    = "/CloseKafkaConsumer"
	PathDescribeKafkaConsumer = "/DescribeKafkaConsumer"

	PathCreateConsumerGroup    = "/CreateConsumerGroup"
	PathDeleteConsumerGroup    = "/DeleteConsumerGroup"
	PathDescribeConsumerGroups = "/DescribeConsumerGroups"
	PathModifyConsumerGroup    = "/ModifyConsumerGroup"

	PathConsumerHeartbeat = "/ConsumerHeartbeat"

	PathDescribeCheckPoint = "/DescribeCheckPoint"
	PathModifyCheckPoint   = "/ModifyCheckPoint"
	PathResetCheckPoint    = "/ResetCheckPoint"

	PathAddTagsToResource      = "/AddTagsToResource"
	PathRemoveTagsFromResource = "/RemoveTagsFromResource"

	PathCreateETLTask       = "/CreateETLTask"
	PathDeleteETLTask       = "/DeleteETLTask"
	PathDescribeETLTask     = "/DescribeETLTask"
	PathModifyETLTask       = "/ModifyETLTask"
	PathDescribeETLTasks    = "/DescribeETLTasks"
	PathModifyETLTaskStatus = "/ModifyETLTaskStatus"

	PathCreateImportTask    = "/CreateImportTask"
	PathDeleteImportTask    = "/DeleteImportTask"
	PathModifyImportTask    = "/ModifyImportTask"
	PathDescribeImportTask  = "/DescribeImportTask"
	PathDescribeImportTasks = "/DescribeImportTasks"

	PathCreateShipper    = "/CreateShipper"
	PathDeleteShipper    = "/DeleteShipper"
	PathModifyShipper    = "/ModifyShipper"
	PathDescribeShipper  = "/DescribeShipper"
	PathDescribeShippers = "/DescribeShippers"

	PathDescribeSessionAnswer = "/DescribeSessionAnswer"
	PathCreateAppInstance     = "/CreateAppInstance"
	PathDescribeAppInstances  = "/DescribeAppInstances"
	PathDeleteAppInstance     = "/DeleteAppInstance"
	PathCreateAppSceneMeta    = "/CreateAppSceneMeta"
	PathDescribeAppSceneMetas = "/DescribeAppSceneMetas"
	PathModifyAppSceneMeta    = "/ModifyAppSceneMeta"
	PathDeleteAppSceneMeta    = "/DeleteAppSceneMeta"

	HeaderAPIVersion = "x-tls-apiversion"
	APIVersion2      = "0.2.0"
	APIVersion3      = "0.3.0"
)
