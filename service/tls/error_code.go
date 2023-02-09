package tls

const (
	// ErrProjectNotExists 日志项目（Project）不存在
	ErrProjectNotExists = "ProjectNotExists"

	// ErrTopicNotExists 指定日志主题不存在
	ErrTopicNotExists = "TopicNotExist"

	// ErrIndexNotExists 指定索引不存在
	ErrIndexNotExists = "IndexNotExists"

	// ErrInvalidParam 请求参数不合法
	ErrInvalidParam = "InvalidArgument"

	// ErrDeserializeFailed 日志反序列化失败
	ErrDeserializeFailed = "DeserializeFailed"

	// ErrInternalServerError 日志服务内部错误
	ErrInternalServerError = "InternalError"

	// ErrProjectAlreadyExists 创建日志项目重复
	ErrProjectAlreadyExists = "ProjectAlreadyExist"

	// ErrTopicAlreadyExist 创建日志主题重复
	ErrTopicAlreadyExist = "TopicAlreadyExist"

	// ErrIndexAlreadyExists 创建日志索引重复
	ErrIndexAlreadyExists = "IndexAlreadyExist"

	// ErrIndexConflict 制定索引状态错误
	ErrIndexConflict = "IndexStateConflict"

	// ErrSearchOutOfRange 日志搜索时指定时间范围越界
	ErrSearchOutOfRange = "SearchOutOfRange"

	// ErrSqlSyntaxError SQL语句语法错误
	ErrSqlSyntaxError = "SqlSyntaxError"

	// ErrSearchSyntaxError 搜索语句语法错误
	ErrSearchSyntaxError = "SearchSyntaxError"

	// ErrSqlResultError SQL检索结果错误
	ErrSqlResultError = "SqlResultError"

	// ProjectQuotaExceed 日志项目配额不足
	ProjectQuotaExceed = "ProjectQuotaExceed"

	// ProjectDeletedHasTopic 删除指定日志项目时，该日志项目仍有topic未被删除
	ProjectDeletedHasTopic = "ProjectDeletedHasTopic"

	// ErrInvalidAccessKeyId AccessKeyID不合法
	ErrInvalidAccessKeyId = "InvalidAccessKeyId"

	// ErrInvalidSecurityToken 授权令牌不合法
	ErrInvalidSecurityToken = "InvalidSecurityToken"

	// ErrAuthorizationQueryParameters 鉴权参数不合法
	ErrAuthorizationQueryParameters = "AuthorizationQueryParametersError"

	// ErrRequestHasExpired 请求超时
	ErrRequestHasExpired = "RequestExpired"

	// ErrSignatureDoesNotMatch 签名不匹配
	ErrSignatureDoesNotMatch = "SignatureDoesNotMatch"

	// ErrRequestTimeTooSkewed 客户端与服务端时间差距太大
	ErrRequestTimeTooSkewed = "RequestTimeTooSkewed"

	// ErrAuthFailed 鉴权失败
	ErrAuthFailed = "AuthFailed"

	// ErrLogMD5CheckFailed 上传日志md5校验失败
	ErrLogMD5CheckFailed = "ErrLogMD5CheckFailed"

	// ErrAccessDenied 提供的AccessKeyID与指定的资源不匹配
	ErrAccessDenied = "AccessDenied"

	// ErrNotSupport 请求方法不支持
	ErrNotSupport = "MethodNotSupport"

	// ErrExceededMaxLogCount 上传日志中Log数量过大
	ErrExceededMaxLogCount = "TooManyLogs"

	// ErrExceededMaxLogSize 上传日志中Log容量太大
	ErrExceededMaxLogSize = "LogSizeTooLarge"

	// ErrExceededMaxLogGroupListSize 上传日志的LogGroupList容量太大
	ErrExceededMaxLogGroupListSize = "LogGroupListSizeTooLarge"

	// ErrExceedQPSLimit QPS过高
	ErrExceedQPSLimit = "ExceedQPSLimit"

	// ErrIndexTypeParam 索引类型不合法
	ErrIndexTypeParam = "InvalidIndexArgument"

	// IndexKeyValueQuotaExceed 索引KeyValue配额不足
	IndexKeyValueQuotaExceed = "IndexKeyValueQuotaExceed"

	// ErrValueTypeConflictAttr 索引类型冲突
	ErrValueTypeConflictAttr = "ValueTypeConflict"

	// ErrSQLFlagConflictAttr SQL Flag冲突
	ErrSQLFlagConflictAttr = "SQLFlagConflict"

	// ErrDelimiterChineseConflict 分隔符与中文配置冲突
	ErrDelimiterChineseConflict = "DelimiterChineseConflict"

	// ErrIndexKeyDuplicate 索引关键字重复
	ErrIndexKeyDuplicate = "IndexKeyDuplicate"

	// ErrIndexKeyValueNULL 索引关键字为空
	ErrIndexKeyValueNULL = "IndexKVNULL"

	// ErrConsumerGroupAlreadyExists 消费者组已存在
	ErrConsumerGroupAlreadyExists = "ConsumerGroupAlreadyExists"

	// ErrConsumerHeartbeatExpired ConsumerGroup心跳过期
	ErrConsumerHeartbeatExpired = "ConsumerHeartbeatExpired"
)
