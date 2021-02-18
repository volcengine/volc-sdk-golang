package vod

type PublicErrorCode string

// General
const (
	success       PublicErrorCode = "Success"       // 请求成功，对应200
	internalError PublicErrorCode = "InternalError" // 内部错误，对应500
)

// Invalid
const (
	invalidAction    PublicErrorCode = "InvalidAction"    // 接口不存在
	invalidVersion   PublicErrorCode = "InvalidVersion"   // 版本不存在
	invalidParameter PublicErrorCode = "InvalidParameter" // 非法参数
	invalidDomain    PublicErrorCode = "InvalidDomain"    // 非法域名
)

// Missing
const (
	missingParameter PublicErrorCode = "MissingParameter" // 缺少参数
)

// Unsupported
const (
	unsupportedProtocol  PublicErrorCode = "UnsupportedProtocol"  // 请求协议错误
	unsupportedRegion    PublicErrorCode = "UnsupportedRegion"    // 地域不支持
	unsupportedOperation PublicErrorCode = "UnsupportedOperation" // 操作不支持
)

// LimitExceeded
const (
	requestLimitExceeded PublicErrorCode = "RequestLimitExceeded" // 请求被限流
)

// Forbidden
const (
	requestForbidden PublicErrorCode = "RequestForbidden" // 请求被禁止，对应403
)

// Denied
const (
	operationDenied PublicErrorCode = "OperationDenied" // 操作被拒绝、未开通
)

// NotFound
const (
	resourceNotFound PublicErrorCode = "ResourceNotFound" // 资源不存在
)

// InUse
const (
	resourceInUse PublicErrorCode = "ResourceInUse" // 资源被占用
)

// Failed
const (
	failedOperation PublicErrorCode = "FailedOperation" // 操作失败
)

// Unauthorized
const (
	unauthorized PublicErrorCode = "Unauthorized" // 操作未授权，对应401
)

// Unavailable
const (
	serviceUnavailable  PublicErrorCode = "ServiceUnavailable"  // 服务不可用
	resourceUnavailable PublicErrorCode = "ResourceUnavailable" // 资源不可用
)

// Insufficient
const (
	resourceInsufficient PublicErrorCode = "ResourceInsufficient" // 资源不足
	quotaInsufficient    PublicErrorCode = "QuotaInsufficient"    // 存储空间不足
)

// Timeout
const (
	serviceTimeout PublicErrorCode = "ServiceTimeout" // 服务超时
)

// Busy
const (
	busy PublicErrorCode = "Busy" // 服务器忙
)
