package producer

import (
	"github.com/go-kit/kit/log"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/common"
)

const delimiter = "|"

var MaxBatchSize int64 = 8 * 1024 * 1024
var MaxBatchCount = 10000

type Config struct {
	TotalSizeLnBytes      int64
	MaxSenderCount        int64
	MaxBlockSec           int
	MaxBatchSize          int64
	MaxBatchCount         int
	LingerTime            time.Duration
	Retries               int
	MaxReservedAttempts   int
	BaseRetryBackoffMs    int64
	MaxRetryBackoffMs     int64
	AdjustShardHashFlag   bool
	ShardCount            int
	NoRetryStatusCodeList []int

	common.LoggerConfig
	common.ClientConfig
	Logger *log.Logger
}

func GetDefaultProducerConfig() *Config {
	return &Config{
		LoggerConfig: common.LoggerConfig{
			LogLevel:      "info",
			LogFileName:   "",
			IsJsonType:    false,
			LogMaxSize:    10,
			LogMaxBackups: 10,
			LogCompress:   false,
		},
		TotalSizeLnBytes:      100 * 1024 * 1024,
		MaxSenderCount:        50,
		MaxBlockSec:           60,
		MaxBatchSize:          512 * 1024,
		LingerTime:            2000 * time.Millisecond,
		Retries:               10,
		MaxReservedAttempts:   11,
		BaseRetryBackoffMs:    1000,
		MaxRetryBackoffMs:     10 * 1000,
		AdjustShardHashFlag:   true,
		ShardCount:            2,
		MaxBatchCount:         4096,
		NoRetryStatusCodeList: []int{400, 404},
	}
}
