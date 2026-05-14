package producer

import (
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/common"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

const (
	TimeoutException = "TimeoutException"

	intMax   int   = 0x7FFFFFFF
	int64Max int64 = 0x7FFFFFFFFFFFFFFF
)

var ErrMemoryTimeout = errors.New(TimeoutException)

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
	sizeNotifier         *broadcastNotifier
	memoryLimiter        *memoryLimiter
	memoryBlockTimeout   time.Duration
	failureController    *failureController
	closing              int32
	activeSends          int64
	once                 sync.Once
}

func newProducer(producerConfig *Config) *producer {
	var logger log.Logger
	if producerConfig.Logger != nil {
		logger = *producerConfig.Logger
	} else {
		logger = common.LogConfig(producerConfig.LoggerConfig)
	}
	client := NewClientWithAPIKey(producerConfig.Endpoint, producerConfig.Region, producerConfig.APIKey,
		producerConfig.AccessKeyID, producerConfig.AccessKeySecret, producerConfig.SecurityToken)

	originalMaxProducerMemoryBytes := producerConfig.MaxProducerMemoryBytes
	producerConfig = validateProducerConfig(producerConfig)
	if originalMaxProducerMemoryBytes != 0 && producerConfig.MaxProducerMemoryBytes != originalMaxProducerMemoryBytes {
		level.Warn(logger).Log("msg", "MaxProducerMemoryBytes adjusted", "configured", originalMaxProducerMemoryBytes, "adjusted", producerConfig.MaxProducerMemoryBytes, "TotalSizeLnBytes", producerConfig.TotalSizeLnBytes)
	}
	configureClientForProducer(client, producerConfig)
	if producerConfig.MaxBatchCount > MaxBatchCount {
		level.Warn(logger).Log("msg", "MaxBatchCount is adjusted to "+strconv.FormatInt(int64(MaxBatchCount), 10))
		producerConfig.MaxBatchCount = MaxBatchCount
	}
	if producerConfig.MaxBatchSize > MaxBatchSize {
		level.Warn(logger).Log("msg", "MaxBatchSize is adjusted to "+strconv.FormatInt(MaxBatchSize, 10))
		producerConfig.MaxBatchSize = MaxBatchSize
	}

	errorStatusMap := func() map[int]struct{} {
		errorCodeMap := make(map[int]struct{})
		for _, v := range producerConfig.NoRetryStatusCodeList {
			errorCodeMap[v] = struct{}{}
		}

		return errorCodeMap
	}()

	producer := &producer{
		cli:                client,
		config:             producerConfig,
		shardCount:         producerConfig.ShardCount,
		sizeNotifier:       newBroadcastNotifier(),
		memoryLimiter:      newMemoryLimiter(producerConfig.MaxProducerMemoryBytes, derivePayloadMemoryLimit(producerConfig)),
		memoryBlockTimeout: memoryBlockTimeout(producerConfig),
		failureController:  newFailureController(producerConfig),
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
	producerConfig.MaxBatchCount = validateField(producerConfig.MaxBatchCount, 0, MaxBatchCount, 4096).(int)
	producerConfig.MaxBatchSize = validateField(producerConfig.MaxBatchSize, int64(0), int64(1024*1024*10), int64(512*1024)).(int64)
	producerConfig.MaxSenderCount = validateField(producerConfig.MaxSenderCount, int64(0), int64Max, int64(50)).(int64)
	producerConfig.BaseRetryBackoffMs = validateField(producerConfig.BaseRetryBackoffMs, int64(0), int64Max, int64(1000)).(int64)
	producerConfig.MaxRetryBackoffMs = validateField(producerConfig.MaxRetryBackoffMs, int64(0), int64Max, int64(10*1000)).(int64)
	if producerConfig.Retries < 0 {
		producerConfig.Retries = 0
	}
	defaultTotalSizeBytes := int64(100 * 1024 * 1024)
	producerConfig.TotalSizeLnBytes = validateField(producerConfig.TotalSizeLnBytes, int64(0), int64Max, defaultTotalSizeBytes).(int64)
	defaultMemoryBytes := defaultProducerMemoryBytes(producerConfig.TotalSizeLnBytes)
	producerConfig.MaxProducerMemoryBytes = validateField(producerConfig.MaxProducerMemoryBytes, int64(0), int64Max, defaultMemoryBytes).(int64)
	if producerConfig.MaxProducerMemoryBytes < producerConfig.TotalSizeLnBytes {
		producerConfig.MaxProducerMemoryBytes = defaultMemoryBytes
	}
	producerConfig.LingerTime = validateField(producerConfig.LingerTime, 100*time.Millisecond, time.Duration(int64Max), 2000*time.Millisecond).(time.Duration)
	if producerConfig.BatchQueueSize == 0 {
		producerConfig.BatchQueueSize = 1000000
	} else if producerConfig.BatchQueueSize < 100 {
		producerConfig.BatchQueueSize = 100
	} else if producerConfig.BatchQueueSize > intMax {
		producerConfig.BatchQueueSize = intMax
	}
	if producerConfig.RetryMode != RetryModeLegacyDoubleRetry && producerConfig.RetryMode != RetryModeProducerManaged {
		producerConfig.RetryMode = RetryModeLegacyDoubleRetry
	}
	if producerConfig.RetryMode == RetryModeProducerManaged && producerConfig.RequestTimeout <= 0 {
		producerConfig.RequestTimeout = 10 * time.Second
	}
	if producerConfig.FailurePolicy != FailurePolicyRetryThenCallback &&
		producerConfig.FailurePolicy != FailurePolicyFailFast &&
		producerConfig.FailurePolicy != FailurePolicyDropWithCallback {
		producerConfig.FailurePolicy = FailurePolicyRetryThenCallback
	}
	producerConfig.CircuitBreaker = normalizeCircuitBreakerConfig(producerConfig.CircuitBreaker)

	return producerConfig
}

func configureClientForProducer(client Client, config *Config) {
	if client == nil || config == nil {
		return
	}
	if config.RequestTimeout > 0 {
		if httpClient := client.GetHttpClient(); httpClient != nil {
			cloned := *httpClient
			cloned.Timeout = config.RequestTimeout
			_ = client.SetHttpClient(&cloned)
		} else {
			client.SetTimeout(config.RequestTimeout)
		}
	}
	if config.ClientRetryPolicy != nil {
		client.SetRetryPolicy(*config.ClientRetryPolicy)
	}
	if config.RetryMode == RetryModeProducerManaged {
		policy := client.GetRetryPolicy()
		policy.MaxAttempts = 1
		client.SetRetryPolicy(policy)
	}
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
	if handled, err := producer.handleCircuitOpenBeforeEnqueue(callBack, 1); handled {
		return err
	}
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
	if logs != nil {
		if handled, err := producer.handleCircuitOpenBeforeEnqueue(callBack, len(logs.Logs)); handled {
			return err
		}
	}
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
				ContextFlow: logs.ContextFlow,
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

func (producer *producer) handleCircuitOpenBeforeEnqueue(callBack CallBack, callbackCount int) (bool, error) {
	if producer == nil || producer.failureController == nil || !producer.failureController.rejectNewData() {
		return false, nil
	}
	return producer.rejectForCircuitOpen(callBack, callbackCount)
}

func (producer *producer) rejectForCircuitOpen(callBack CallBack, callbackCount int) (bool, error) {
	switch producer.config.FailurePolicy {
	case FailurePolicyDropWithCallback:
		for i := 0; i < callbackCount; i++ {
			if callBack != nil {
				callBack.Fail(newCircuitOpenResult())
			}
		}
		return true, nil
	case FailurePolicyFailFast:
		return true, newCircuitOpenError()
	default:
		return false, nil
	}
}

func newCircuitOpenResult() *Result {
	result := newResult()
	result.Attempts = append(result.Attempts, newAttempt(false, "", CircuitOpenException, CircuitOpenException, GetTimeMs(time.Now().UnixNano())))
	return result
}

func (producer *producer) waitTime() error {
	if producer.config.MaxBlockSec > 0 {
		deadline := time.Now().Add(time.Duration(producer.config.MaxBlockSec) * time.Second)
		waited := false
		for {
			if producer.currentProducerLogGroupSize() <= producer.config.TotalSizeLnBytes {
				if waited {
					level.Warn(producer.logger).Log("msg", "wait for produce memory")
				}
				return nil
			}
			notifyCh := producer.sizeNotifier.beginWait()
			if producer.currentProducerLogGroupSize() <= producer.config.TotalSizeLnBytes {
				producer.sizeNotifier.endWait()
				if waited {
					level.Warn(producer.logger).Log("msg", "wait for produce memory")
				}
				return nil
			}
			now := time.Now()
			if !now.Before(deadline) {
				producer.sizeNotifier.endWait()
				return ErrMemoryTimeout
			}
			sleepFor := time.Second
			if remaining := deadline.Sub(now); remaining < sleepFor {
				sleepFor = remaining
			}
			if !waitBroadcast(notifyCh, sleepFor, producer.closeCh) {
				producer.sizeNotifier.endWait()
				return errors.New("the producer is closed")
			}
			producer.sizeNotifier.endWait()
			waited = true
		}
	}

	if producer.config.MaxBlockSec == 0 {
		if producer.currentProducerLogGroupSize() > producer.config.TotalSizeLnBytes {
			return ErrMemoryTimeout
		}

		return nil
	}

	if producer.config.MaxBlockSec < 0 {
		for {
			if producer.currentProducerLogGroupSize() <= producer.config.TotalSizeLnBytes {
				return nil
			}
			notifyCh := producer.sizeNotifier.beginWait()
			if producer.currentProducerLogGroupSize() <= producer.config.TotalSizeLnBytes {
				producer.sizeNotifier.endWait()
				return nil
			}
			if !waitBroadcast(notifyCh, 0, producer.closeCh) {
				producer.sizeNotifier.endWait()
				return errors.New("the producer is closed")
			}
			producer.sizeNotifier.endWait()
		}
	}

	return nil
}

func memoryBlockTimeout(config *Config) time.Duration {
	if config == nil {
		return 0
	}
	if config.MaxBlockSec > 0 {
		return time.Duration(config.MaxBlockSec) * time.Second
	}
	if config.MaxBlockSec < 0 {
		return -1
	}
	return 0
}

func (producer *producer) reserveBatchLog(batchLog *BatchLog, bytes int64) error {
	if producer == nil || producer.memoryLimiter == nil || batchLog == nil || bytes <= 0 {
		return nil
	}
	if producer.memoryBlockTimeout == 0 {
		if err := producer.memoryLimiter.acquirePayload(bytes, 0); err != nil {
			return err
		}
		batchLog.setReservedBytes(bytes)
		return nil
	}
	if err := producer.memoryLimiter.acquirePayloadWithCancel(bytes, producer.memoryBlockTimeout, producer.closeCh); err != nil {
		return err
	}
	batchLog.setReservedBytes(bytes)
	return nil
}

func (producer *producer) releaseBatchLog(batchLog *BatchLog) {
	if producer == nil || batchLog == nil {
		return
	}
	if batchLog.hasCircuitPermit() && producer.failureController != nil {
		producer.failureController.releasePermit()
	}
	if producer.memoryLimiter != nil {
		producer.memoryLimiter.release(batchLog.reservedMemoryBytes())
	}
	batchLog.clearReservation()
}

func (producer *producer) releaseBatchResources(batch *Batch) {
	if producer == nil || batch == nil {
		return
	}
	if producer.failureController != nil {
		producer.failureController.releasePermits(batch.takeCircuitPermitCount())
	}
	batch.releaseReservedBytes(producer.memoryLimiter)
	producer.addProducerLogGroupSize(batch.setTotalDataSize(0))
}

func (producer *producer) currentProducerLogGroupSize() int64 {
	if producer == nil {
		return 0
	}
	return atomic.LoadInt64(&producer.producerLogGroupSize)
}

func (producer *producer) addProducerLogGroupSize(delta int64) int64 {
	if producer == nil || delta == 0 {
		return producer.currentProducerLogGroupSize()
	}
	next := atomic.AddInt64(&producer.producerLogGroupSize, delta)
	if delta < 0 {
		producer.sizeNotifier.notify()
	}
	return next
}

func (producer *producer) acquireTemporaryBytes(bytes int64) error {
	if producer == nil || producer.memoryLimiter == nil || bytes <= 0 {
		return nil
	}
	return producer.memoryLimiter.acquireTemporaryWithCancel(bytes, 0, producer.closeCh)
}

func (producer *producer) releaseTemporaryBytes(bytes int64) {
	if producer == nil || producer.memoryLimiter == nil || bytes <= 0 {
		return
	}
	producer.memoryLimiter.release(bytes)
}

func (producer *producer) beginSend() bool {
	atomic.AddInt64(&producer.activeSends, 1)
	if atomic.LoadInt32(&producer.closing) != 0 {
		atomic.AddInt64(&producer.activeSends, -1)
		return false
	}
	return true
}

func (producer *producer) endSend() {
	atomic.AddInt64(&producer.activeSends, -1)
}

func (producer *producer) waitActiveSends() {
	for atomic.LoadInt64(&producer.activeSends) != 0 {
		time.Sleep(time.Millisecond)
	}
}

func (producer *producer) putToDispatcher(batchLog *BatchLog) (err error) {
	if batchLog.Log.Time == 0 {
		batchLog.Log.Time = time.Now().UnixNano() / time.Millisecond.Nanoseconds()
	}

	logSize := int64(GetLogSize(batchLog.Log))
	if logSize > MaxBatchSize {
		err = errors.New("the log size " + strconv.FormatInt(logSize, 10) + " is larger than the max batch size " + strconv.FormatInt(MaxBatchSize, 10))
		return err
	}

	if logSize > producer.config.TotalSizeLnBytes {
		err = errors.New("the log size " + strconv.FormatInt(logSize, 10) + "is larger than the total size " + strconv.FormatInt(producer.config.TotalSizeLnBytes, 10))
		return err
	}
	batchLog.setLogSize(logSize)

	if !producer.beginSend() {
		err = errors.New("the producer is closed")
		return err
	}
	defer func() {
		producer.endSend()
		if r := recover(); r != nil {
			producer.releaseBatchLog(batchLog)
			err = errors.New("the producer is closed")
		}
	}()

	allowed, permit := true, false
	if producer.failureController != nil {
		allowed, permit = producer.failureController.allowEnqueue()
	}
	if !allowed {
		handled, circuitErr := producer.rejectForCircuitOpen(batchLog.Key.CallBackFun, 1)
		if handled {
			err = circuitErr
			return err
		}
	}
	batchLog.setCircuitPermit(permit)

	if err = producer.reserveBatchLog(batchLog, estimateBatchLogReservationBytes(batchLog, logSize)); err != nil {
		producer.releaseBatchLog(batchLog)
		return err
	}

	select {
	case <-producer.closeCh:
		err = errors.New("the producer is closed")
		producer.releaseBatchLog(batchLog)
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
	producer.once.Do(func() {
		atomic.StoreInt32(&producer.closing, 1)
		close(producer.closeCh)
		producer.waitActiveSends()

		close(producer.threadPool.sender.stopCh)

		close(producer.dispatcher.stopCh)
		producer.dispatcherWaitGroup.Wait()

		close(producer.threadPool.stopCh)
		producer.threadPoolWaitGroup.Wait()
		producer.senderWaitGroup.Wait()

		level.Info(producer.logger).Log("msg", "producer close finish")
	})
}

func (producer *producer) ForceClose() {
	producer.once.Do(func() {
		atomic.StoreInt32(&producer.closing, 1)
		close(producer.closeCh)
		producer.waitActiveSends()

		close(producer.threadPool.sender.stopCh)

		close(producer.dispatcher.forceQuitCh)
		producer.dispatcherWaitGroup.Wait()
		producer.dispatcher.releasePending()

		close(producer.threadPool.forceQuitCh)
		producer.threadPoolWaitGroup.Wait()
		producer.threadPool.releasePending()
		producer.senderWaitGroup.Wait()
		producer.memoryLimiter.reset()

		level.Info(producer.logger).Log("msg", "producer force close finish")
	})
}

func (producer *producer) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
	producer.cli.ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken)
}

func (producer *producer) SetAPIKey(apiKey string) {
	producer.cli.SetAPIKey(apiKey)
}
