package producer

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"sync"
	"sync/atomic"
	"time"
)

const (
	TimeoutException = "TimeoutException"

	MaxLogSize       = 1048576
	intMax     int   = 0x7FFFFFFF
	int64Max   int64 = 0x7FFFFFFFFFFFFFFF
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
	logger := logConfig(producerConfig)

	client := NewClient(producerConfig.Endpoint, producerConfig.AccessKeyID, producerConfig.AccessKeySecret, "", producerConfig.Region)
	producerConfig = validateProducerConfig(producerConfig)

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
	producerConfig.MaxBatchSize = validateField(producerConfig.MaxBatchSize, int64(0), int64(1024*1024*5), int64(1024*1024*5)).(int64)
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

func (producer *producer) putToDispatcher(batchLog *BatchLog) error {

	if batchLog.Log.Time == 0 {
		batchLog.Log.Time = time.Now().Unix()
	}

	select {
	case <-producer.closeCh:
		return errors.New("the producer is closed")
	default:
	}

	logSize := GetLogSize(batchLog.Log)
	if logSize > MaxLogSize {
		level.Warn(producer.logger).Log("msg", "failed to put log, log size too large")
		level.Warn(producer.logger).Log("log_content", batchLog.Log.String())
		return fmt.Errorf("the input log size: %d exceeds limitation: %d", logSize, MaxLogSize)
	}

	select {
	case <-producer.closeCh:
		return errors.New("the producer is closed")
	default:
		select {
		case <-producer.closeCh:
			return errors.New("the producer is closed")
		case producer.dispatcher.newLogRecvChan <- batchLog:
		}
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
	close(producer.dispatcher.stopCh)
	producer.dispatcherWaitGroup.Wait()
	close(producer.threadPool.stopCh)
	producer.threadPoolWaitGroup.Wait()
	close(producer.threadPool.sender.stopCh)
	producer.senderWaitGroup.Wait()
	level.Info(producer.logger).Log("msg", "producer close finish")
}

func (producer *producer) ForceClose() {
	close(producer.closeCh)
	close(producer.dispatcher.forceQuitCh)
	producer.dispatcherWaitGroup.Wait()
	close(producer.threadPool.forceQuitCh)
	producer.threadPoolWaitGroup.Wait()
	close(producer.threadPool.sender.stopCh)
	producer.senderWaitGroup.Wait()
	level.Info(producer.logger).Log("msg", "producer close finish")
}

func (producer *producer) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
	producer.cli.ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken)
}
