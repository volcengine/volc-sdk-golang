package producer

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

const (
	TimeoutException = "TimeoutException"

	intMax   int   = 0x7FFFFFFF
	int64Max int64 = 0x7FFFFFFFFFFFFFFF
)

type producer struct {
	closeCh              chan struct{}
	cli                  Client
	config               *Config
	dispatcher           *Dispatcher
	threadPool           *ThreadPool
	senderWaitGroup      *sync.WaitGroup
	threadPoolWaitGroup  *sync.WaitGroup
	dispatcherWaitGroup  *sync.WaitGroup
	shardCount           int
	logger               log.Logger
	producerLogGroupSize int64
}

func newProducer(producerConfig *Config) *producer {
	logger := common.LogConfig(producerConfig.LoggerConfig)
	client := NewClient(producerConfig.Endpoint, producerConfig.AccessKeyID, producerConfig.AccessKeySecret,
		producerConfig.SecurityToken, producerConfig.Region)

	producerConfig = validateProducerConfig(producerConfig)
	if producerConfig.MaxBatchCount > 10000 {
		level.Warn(logger).Log("msg", "MaxBatchCount is adjusted to 10000 (the maximum allowed value)")
		producerConfig.MaxBatchCount = 10000
	}
	if producerConfig.MaxBatchSize > 5*1024*1024 {
		level.Warn(logger).Log("msg", "MaxBatchSize is adjusted to 5 MB (the maximum allowed value)")
		producerConfig.MaxBatchSize = 5 * 1024 * 1024
	}

	errorStatusMap := func() map[int]struct{} {
		errorCodeMap := make(map[int]struct{})
		for _, v := range producerConfig.NoRetryStatusCodeList {
			errorCodeMap[v] = struct{}{}
		}

		return errorCodeMap
	}()

	producer := &producer{
		cli:        client,
		config:     producerConfig,
		shardCount: producerConfig.ShardCount,
	}

	sender := initSender(producer.cli, newRetryQueue(), producerConfig.MaxSenderCount, logger, errorStatusMap, producer)
	threadPool := initThreadPool(sender, logger)
	dispatcher := initDispatcher(producerConfig, sender, logger, threadPool, producer)

	producer.dispatcher = dispatcher
	producer.threadPool = threadPool
	producer.senderWaitGroup = &sync.WaitGroup{}
	producer.threadPoolWaitGroup = &sync.WaitGroup{}
	producer.dispatcherWaitGroup = &sync.WaitGroup{}
	producer.logger = logger
	producer.closeCh = make(chan struct{})

	return producer
}

func validateProducerConfig(producerConfig *Config) *Config {
	producerConfig.MaxReservedAttempts = validateField(producerConfig.MaxReservedAttempts, 0, intMax, 11).(int)
	producerConfig.MaxBatchCount = validateField(producerConfig.MaxBatchCount, 0, 40960, 40960).(int)
	producerConfig.MaxBatchSize = validateField(producerConfig.MaxBatchSize, int64(0), int64(1024*1024*10), int64(1024*1024*5)).(int64)
	producerConfig.MaxSenderCount = validateField(producerConfig.MaxSenderCount, int64(0), int64Max, int64(50)).(int64)
	producerConfig.BaseRetryBackoffMs = validateField(producerConfig.BaseRetryBackoffMs, int64(0), int64Max, int64(100)).(int64)
	producerConfig.MaxRetryBackoffMs = validateField(producerConfig.MaxRetryBackoffMs, int64(0), int64Max, int64(100)).(int64)
	producerConfig.TotalSizeLnBytes = validateField(producerConfig.TotalSizeLnBytes, int64(0), int64Max, int64(100*1024*1024)).(int64)
	producerConfig.LingerTime = validateField(producerConfig.LingerTime, 100*time.Millisecond, time.Duration(int64Max), 2000*time.Millisecond).(time.Duration)

	return producerConfig
}

func validateField(val, min, max, defaultVal interface{}) interface{} {
	switch val.(type) {
	case int64:
		convertedVal := val.(int64)
		if convertedVal <= min.(int64) || convertedVal > max.(int64) {
			return defaultVal
		}

		return val
	case int:
		convertedVal := val.(int)
		if convertedVal <= min.(int) || convertedVal > max.(int) {
			return defaultVal
		}

		return val
	case time.Duration:
		convertedVal := val.(time.Duration)
		if convertedVal <= min.(time.Duration) || convertedVal > max.(time.Duration) {
			return defaultVal
		}

		return val
	default:
		return val
	}
}

func (producer *producer) SendLog(shardHash, topic, source string, filename string, log *pb.Log, callBack CallBack) error {
	err := producer.waitTime()
	if err != nil {
		return err
	}

	batchLog := &BatchLog{
		Key: BatchKey{
			Topic:       topic,
			Source:      source,
			ShardHash:   shardHash,
			FileName:    filename,
			CallBackFun: callBack,
		},
		Log: log,
	}

	return producer.putToDispatcher(batchLog)
}

func (producer *producer) SendLogs(shardHash, topic, source, filename string, logs *pb.LogGroup, callBack CallBack) error {
	err := producer.waitTime()
	if err != nil {
		return err
	}

	for _, mlog := range logs.Logs {
		batchLog := &BatchLog{
			Key: BatchKey{
				Topic:       topic,
				Source:      source,
				ShardHash:   shardHash,
				FileName:    filename,
				CallBackFun: callBack,
			},
			Log: mlog,
		}

		if err := producer.putToDispatcher(batchLog); err != nil {
			return err
		}
	}

	return nil
}

func (producer *producer) waitTime() error {
	if producer.config.MaxBlockSec > 0 {
		for i := 0; i < producer.config.MaxBlockSec; i++ {
			if atomic.LoadInt64(&producer.producerLogGroupSize) <= producer.config.TotalSizeLnBytes {
				return nil
			}

			time.Sleep(time.Second)
		}

		return errors.New(TimeoutException)
	}

	if producer.config.MaxBlockSec == 0 {
		if atomic.LoadInt64(&producer.producerLogGroupSize) > producer.config.TotalSizeLnBytes {
			return errors.New(TimeoutException)
		}

		return nil
	}

	if producer.config.MaxBlockSec < 0 {
		for {
			if atomic.LoadInt64(&producer.producerLogGroupSize) <= producer.config.TotalSizeLnBytes {
				return nil
			}

			time.Sleep(time.Second)
		}
	}

	return nil
}

func (producer *producer) putToDispatcher(batchLog *BatchLog) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("the producer is closed")
		}
	}()

	select {
	case <-producer.closeCh:
		err = errors.New("the producer is closed")
		return err
	default:
	}

	if batchLog.Log.Time == 0 {
		batchLog.Log.Time = time.Now().UnixNano() / time.Millisecond.Nanoseconds()
	}

	select {
	case <-producer.closeCh:
		err = errors.New("the producer is closed")
		return err
	case producer.dispatcher.newLogRecvChan <- batchLog:
	}

	return nil
}

func (producer *producer) Start() {
	producer.threadPoolWaitGroup.Add(1)
	go producer.threadPool.start(producer.senderWaitGroup, producer.threadPoolWaitGroup)

	producer.dispatcherWaitGroup.Add(2)
	go producer.dispatcher.checkBatchTime(producer.dispatcherWaitGroup, producer.config)
	go producer.dispatcher.run(producer.dispatcherWaitGroup)

	level.Info(producer.logger).Log("msg", "producer start")
}

func (producer *producer) Close() {
	close(producer.closeCh)
	close(producer.threadPool.sender.stopCh)

	close(producer.dispatcher.stopCh)
	producer.dispatcherWaitGroup.Wait()

	close(producer.threadPool.stopCh)
	producer.threadPoolWaitGroup.Wait()
	producer.senderWaitGroup.Wait()

	level.Info(producer.logger).Log("msg", "producer close finish")
}

func (producer *producer) ForceClose() {
	close(producer.closeCh)
	close(producer.threadPool.sender.stopCh)

	close(producer.dispatcher.forceQuitCh)
	producer.dispatcherWaitGroup.Wait()

	close(producer.threadPool.forceQuitCh)
	producer.threadPoolWaitGroup.Wait()
	producer.senderWaitGroup.Wait()

	level.Info(producer.logger).Log("msg", "producer close finish")
}

func (producer *producer) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
	producer.cli.ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken)
}
