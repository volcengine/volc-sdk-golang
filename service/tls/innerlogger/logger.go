package innerlogger

import (
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gopkg.in/natefinch/lumberjack.v2"
)

var DefaultLogger TLSSDKLogger = initDefaultLogger()

func (l *TLSSDKLogger) GetLogLevel() int {
	switch l.loggerConfig.LogLevel {
	case "debug":
		return LogLevelDebug
	case "info":
		return LogLevelInfo
	case "warn":
		return LogLevelWarn
	case "error":
		return LogLevelError
	default:
		return LogLevelInfo
	}
}

func NewLogger(config *LoggerConfig) TLSSDKLogger {
	var logger TLSSDKLogger
	logger.loggerConfig = *config

	if config.LogFileName == "" {
		if config.IsJsonType {
			logger.Logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
		} else {
			logger.Logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		}
	} else {
		if config.IsJsonType {
			logger.Logger = log.NewJSONLogger(initLogFlusher(config))
		} else {
			logger.Logger = log.NewLogfmtLogger(initLogFlusher(config))
		}
	}
	switch config.LogLevel {
	case "debug":
		logger.Logger = level.NewFilter(logger.Logger, level.AllowDebug())
	case "info":
		logger.Logger = level.NewFilter(logger.Logger, level.AllowInfo())
	case "warn":
		logger.Logger = level.NewFilter(logger.Logger, level.AllowWarn())
	case "error":
		logger.Logger = level.NewFilter(logger.Logger, level.AllowError())
	default:
		logger.Logger = level.NewFilter(logger.Logger, level.AllowInfo())
	}
	logger.Logger = log.With(logger.Logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	return logger
}

func initLogFlusher(config *LoggerConfig) *lumberjack.Logger {
	if config.LogMaxSize == 0 {
		config.LogMaxSize = 10
	}
	if config.LogMaxBackups == 0 {
		config.LogMaxBackups = 10
	}
	return &lumberjack.Logger{
		Filename:   config.LogFileName,
		MaxSize:    config.LogMaxSize,
		MaxBackups: config.LogMaxBackups,
		Compress:   config.LogCompress,
	}
}

func initDefaultLogger() TLSSDKLogger {
	var intCastOrDefault = func(s string, defaultValue int) int {
		var (
			err error
			v   int64
		)
		v, err = strconv.ParseInt(os.Getenv(s), 10, 32)
		if err != nil {
			return defaultValue
		}
		return int(v)
	}
	var boolCastOrDefaul = func(s string, defaultValue bool) bool {
		switch strings.ToLower(s) {
		case "true":
			return true
		case "false":
			return false
		default:
			return defaultValue
		}
	}

	defaultLoggerConfig := &LoggerConfig{
		LogLevel:      strings.ToLower(os.Getenv(EnvLogLevel)),
		LogFileName:   os.Getenv(EnvLogFileName),
		IsJsonType:    boolCastOrDefaul(os.Getenv(EnvIsJsonType), false),
		LogMaxSize:    intCastOrDefault(os.Getenv(EnvLogMaxSize), 10),
		LogMaxBackups: intCastOrDefault(os.Getenv(EnvLogMaxBackups), 10),
		LogCompress:   boolCastOrDefaul(os.Getenv(EnvLogCompress), false),
	}

	return NewLogger(defaultLoggerConfig)
}
