package tls

const (
	// 指定日志项目不存在
	PROJECT_NOT_EXIST = "ProjectNotExists"

	// 指定日志主题不存在
	TOPIC_NOT_EXIST = "TopicNotExist"

	// 指定索引不存在
	INDEX_NOT_EXIST = "IndexNotExists"

	// 请求参数不合法
	INVALID_ARGUMENT = "InvalidArgument"

	// 日志反序列化失败
	DESERIALIZATION_FAILED = "DeserializeFailed"

	// 日志服务内部错误
	INTERNAL_SERVER_ERROR = "ErrInternalServerError"

	// 创建日志项目重复
	PROJECT_DUPLICATE = "ProjectAlreadyExist"

	// 创建日志主题重复
	TOPIC_DUPLICATE = "TopicAlreadyExist"

	// 创建日志索引重复
	INDEX_DUPLICATE = "IndexAlreadyExist"

	// 制定索引状态错误
	INDEX_STATUS_CONFLICT = "IndexStateConflict"

	// 日志搜索时指定时间范围越界
	SEARCH_OUT_OF_RANGE = "SearchOutOfRange"

	// SQL语句语法错误
	SQL_SYNTAX_ERROR = "SqlSyntaxError"

	// 搜索语句语法错误
	SEARCH_SYNTAX_ERROR = "SearchSyntaxError"

	// SQL检索结果错误
	SQL_RESULT_ERROR = "SqlResultError"

	// 日志项目配额不足
	EXCEED_PROJECT_QUOTA = "ProjectQuotaExceed"

	// 删除指定日志项目时，该日志项目仍有topic未被删除
	DELETE_PROJECT_HAS_TOPIC = "ProjectDeletedHasTopic"

	// AccessKeyID不合法
	INVALID_ACCESSKEYID = "InvalidAccessKeyId"

	// 授权令牌不合法
	INVALID_SECURITY_TOKEN = "InvalidSecurityToken"

	// 鉴权参数不合法
	AUTH_INVALID_ARGUMENT = "AuthorizationQueryParametersError"

	// 请求超时
	REQUEST_EXPIRED = "RequestExpired"

	// 签名不匹配
	SIGNATURE_NOT_MATCH = "SignatureDoesNotMatch"

	// 客户端与服务端时间差距太大
	REQUEST_TIME_SKEWED = "RequestTimeTooSkewed"

	// 鉴权失败
	AUTH_FAILED = "AuthFailed"

	// 上传日志md5校验失败
	LOGS_MD5_CHECK_FAILED = "ErrLogMD5CheckFailed"

	// 提供的AccessKeyID与指定的资源不匹配
	ACCESS_DENIED = "AccessDenied"

	// 请求方法不支持
	METHOD_NOT_SUPPORT = "MethodNotSupport"

	// 上传日志中Log数量过大
	EXCEED_MAX_LOG_COUNT = "TooManyLogs"

	// 上传日志中Log容量太大
	EXCEED_MAX_LOG_SIZE = "LogSizeTooLarge"

	// 上传日志的LogGroupList容量太大
	EXCEED_MAX_LOGGROUPLIST_SIZE = "LogGroupListSizeTooLarge"

	// QPS过高
	EXCEED_QPS_LIMIT = "ExceedQPSLimit"

	// 索引类型不合法
	INVALID_TYPE_PARAM = "InvalidIndexArgument"

	// 索引KeyValue配额不足
	EXCEED_KEY_VALUE_QUOTA = "IndexKeyValueQuotaExceed"

	// 索引类型冲突
	VALUE_TYPE_CONFLICT = "ValueTypeConflict"

	// SQL Flag冲突
	SQL_FLAG_CONFLICT = "SQLFlagConflict"

	// 分隔符与中文配置冲突
	DELIMITER_CHINESE_CONFLICT = "DelimiterChineseConflict"

	// 索引关键字重复
	INDEX_KEY_DUPLICATE = "IndexKeyDuplicate"

	// 索引关键字为空
	INDEX_KEY_VALUE_NULL = "IndexKVNULL"
)
