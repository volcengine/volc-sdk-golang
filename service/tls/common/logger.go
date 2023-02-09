package common

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogConfig(producerConfig LoggerConfig) log.Logger {
	var logger log.Logger

	if producerConfig.LogFileName == "" {
		if producerConfig.IsJsonType {
			logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
		} else {
			logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		}
	} else {
		if producerConfig.IsJsonType {
			logger = log.NewLogfmtLogger(initLogFlusher(producerConfig))
		} else {
			logger = log.NewJSONLogger(initLogFlusher(producerConfig))
		}
	}
	switch producerConfig.LogLevel {
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "info":
		logger = level.NewFilter(logger, level.AllowInfo())
	case "warn":
		logger = level.NewFilter(logger, level.AllowWarn())
	case "error":
		logger = level.NewFilter(logger, level.AllowError())
	default:
		logger = level.NewFilter(logger, level.AllowInfo())
	}
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	return logger
}

func initLogFlusher(config LoggerConfig) *lumberjack.Logger {
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
