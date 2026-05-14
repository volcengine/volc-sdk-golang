package producer

import (
	"time"

	"github.com/go-kit/kit/log"

	tls "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
)

const delimiter = "|"

var MaxBatchSize int64 = 9*1024*1024 + 512*1024
var MaxBatchCount = 32768
var MaxLogGroupCount = 10000

type RetryMode int

const (
	RetryModeLegacyDoubleRetry RetryMode = iota
	RetryModeProducerManaged
)

type FailurePolicy int

const (
	FailurePolicyRetryThenCallback FailurePolicy = iota
	FailurePolicyFailFast
	FailurePolicyDropWithCallback
)

type CircuitBreakerConfig struct {
	Enabled             bool
	MinRequests         int
	FailureRatio        float64
	ConsecutiveFailures int
	OpenDuration        time.Duration
	HalfOpenMaxRequests int
}

type Config struct {
	TotalSizeLnBytes       int64
	MaxProducerMemoryBytes int64
	MaxSenderCount         int64
	MaxBlockSec            int
	BatchQueueSize         int
	MaxBatchSize           int64
	MaxBatchCount          int
	LingerTime             time.Duration
	Retries                int
	MaxReservedAttempts    int
	BaseRetryBackoffMs     int64
	MaxRetryBackoffMs      int64
	AdjustShardHashFlag    bool
	ShardCount             int
	NoRetryStatusCodeList  []int
	EnableNanosecond       bool
	RetryMode              RetryMode
	RequestTimeout         time.Duration
	ClientRetryPolicy      *tls.RetryPolicy
	FailurePolicy          FailurePolicy
	CircuitBreaker         CircuitBreakerConfig

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
		TotalSizeLnBytes:       100 * 1024 * 1024,
		MaxProducerMemoryBytes: 200 * 1024 * 1024,
		MaxSenderCount:         50,
		MaxBlockSec:            60,
		BatchQueueSize:         1000000,
		MaxBatchSize:           512 * 1024,
		LingerTime:             2000 * time.Millisecond,
		Retries:                10,
		MaxReservedAttempts:    11,
		BaseRetryBackoffMs:     1000,
		MaxRetryBackoffMs:      10 * 1000,
		AdjustShardHashFlag:    true,
		ShardCount:             2,
		MaxBatchCount:          4096,
		NoRetryStatusCodeList:  []int{400, 403, 404},
		RetryMode:              RetryModeLegacyDoubleRetry,
		FailurePolicy:          FailurePolicyRetryThenCallback,
	}
}
