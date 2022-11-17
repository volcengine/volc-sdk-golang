package innerlogger

const (
	EnvLogLevel      = "LOG_SERVICE_SDK_LOG_LEVEL"
	EnvLogFileName   = "LOG_SERVICE_SDK_LOG_FILE_NAME"
	EnvIsJsonType    = "LOG_SERVICE_SDK_IS_JSON_TYPE"
	EnvLogMaxSize    = "LOG_SERVICE_SDK_LOG_MAX_SIZE"
	EnvLogMaxBackups = "LOG_SERVICE_SDK_LOG_MAX_BACKUPS"
	EnvLogCompress   = "LOG_SERVICE_SDK_LOG_COMPRESS"
)

const (
	LogLevelDebug = 0
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)
