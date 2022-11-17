package innerlogger

import (
	"github.com/go-kit/kit/log"
)

type LoggerConfig struct {
	LogLevel      string
	LogFileName   string
	IsJsonType    bool
	LogMaxSize    int
	LogMaxBackups int
	LogCompress   bool
}

type TLSSDKLogger struct {
	log.Logger
	loggerConfig LoggerConfig
}
