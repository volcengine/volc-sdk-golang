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
var errProducerClosed = errors.New("the producer is closed")

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
	closed               int32
	activeSends          int64
	activeAdmissions     int64
	once                 sync.Once

	// Instance observability state.
	stats                     *statsState
	instanceID                string
	runtimeMetricsLevel       int32
	metricsMu                 sync.RWMutex
	collectionGeneration      uint64
	detailGeneration          uint64
	collectionStartedAtUnixMs int64
	started                   int32 // 0=new, 1=running (set in Start())
}

func newProducer(producerConfig *Config) *producer {
	client := NewClientWithAPIKey(producerConfig.Endpoint, producerConfig.Region, producerConfig.APIKey,
		producerConfig.AccessKeyID, producerConfig.AccessKeySecret, producerConfig.SecurityToken)
	var logger log.Logger
	if producerConfig.Logger != nil {
		logger = *producerConfig.Logger
	} else {
		logger = common.LogConfig(producerConfig.LoggerConfig)
	}

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
		cli:                  client,
		config:               producerConfig,
		shardCount:           producerConfig.ShardCount,
		sizeNotifier:         newBroadcastNotifier(),
		memoryLimiter:        newMemoryLimiter(producerConfig.MaxProducerMemoryBytes, derivePayloadMemoryLimit(producerConfig)),
		memoryBlockTimeout:   memoryBlockTimeout(producerConfig),
		failureController:    newFailureController(producerConfig),
		stats:                newStatsState(producerConfig.MetricsLevel, producerConfig.MaxSenderCount),
		instanceID:           newObservabilityInstanceID("producer"),
		runtimeMetricsLevel:  int32(producerConfig.MetricsLevel),
		collectionGeneration: 1,
		detailGeneration:     1,
		started:              0, // 0=new, set to 1 in Start()
	}
	if producerConfig.MetricsLevel != MetricsLevelOff {
		producer.collectionStartedAtUnixMs = unixMillis(time.Now())
	}
	producer.memoryLimiter.setMetricsEnabled(producerConfig.MetricsLevel != MetricsLevelOff)

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
	if !validMetricsLevel(producerConfig.MetricsLevel) {
		producerConfig.MetricsLevel = MetricsLevelOff
	}
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
	if !producer.beginAdmission() {
		return errProducerClosed
	}
	defer producer.endAdmission()
	err := producer.waitTime()
	if err != nil {
		producer.recordWaitTimeRejection(err)
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
	if !producer.beginAdmission() {
		return errProducerClosed
	}
	defer producer.endAdmission()
	err := producer.waitTime()
	if err != nil {
		producer.recordWaitTimeRejection(err)
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
				return errProducerClosed
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
				return errProducerClosed
			}
			producer.sizeNotifier.endWait()
		}
	}

	return nil
}

// recordWaitTimeRejection accounts for a failed public admission decision at
// the legacy TotalSizeLnBytes gate. SendLogs is still one decision here: the
// group has not been split into per-log admissions yet.
func (producer *producer) recordWaitTimeRejection(err error) {
	if !producer.metricsEnabled() {
		return
	}
	if errors.Is(err, ErrMemoryTimeout) {
		producer.stats.addRejectedBatchInflightLimit(1)
	} else if errors.Is(err, errProducerClosed) {
		atomic.AddInt64(&producer.stats.rejectedClosed, 1)
	}
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

func (producer *producer) reserveBatchLog(batchLog *BatchLog, bytes int64) (bool, bool, error) {
	if producer == nil || producer.memoryLimiter == nil || batchLog == nil || bytes <= 0 {
		return false, false, nil
	}
	if producer.memoryBlockTimeout == 0 {
		rejected, observed, err := producer.memoryLimiter.acquirePayloadOutcome(bytes, 0)
		if err != nil {
			return rejected, observed, err
		}
		batchLog.setReservedBytes(bytes)
		return false, false, nil
	}
	rejected, observed, err := producer.memoryLimiter.acquirePayloadWithCancelOutcome(bytes, producer.memoryBlockTimeout, producer.closeCh)
	if err != nil {
		return rejected, observed, err
	}
	batchLog.setReservedBytes(bytes)
	return false, false, nil
}

func (producer *producer) releaseBatchLog(batchLog *BatchLog) {
	if producer == nil || batchLog == nil {
		return
	}
	if batchLog.hasCircuitPermit() && producer.failureController != nil {
		producer.failureController.releasePermit()
	}
	if producer.memoryLimiter != nil {
		producer.memoryLimiter.releasePayloadFrom(payloadStageIngress, batchLog.reservedMemoryBytes())
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
	if delta > 0 && producer.metricsEnabled() {
		producer.stats.observeBatchInflightMemory(next)
	}
	if delta < 0 {
		producer.sizeNotifier.notify()
	}
	return next
}

func (producer *producer) acquireTemporaryBytes(bytes int64) error {
	if producer == nil || producer.memoryLimiter == nil || bytes <= 0 {
		return nil
	}
	// closeCh rejects new ingress as soon as either Close or ForceClose starts.
	// Pending batches must still acquire their send buffer during graceful
	// Close; only ForceClose closes dispatcher.forceQuitCh and may cancel them.
	// A nil channel is intentional for minimal unit-test producers and never
	// cancels.
	var forceQuitCh <-chan struct{}
	if producer.dispatcher != nil {
		forceQuitCh = producer.dispatcher.forceQuitCh
	}
	return producer.memoryLimiter.acquireTemporaryWithCancel(bytes, 0, forceQuitCh)
}

func (producer *producer) releaseTemporaryBytes(bytes int64) {
	if producer == nil || producer.memoryLimiter == nil || bytes <= 0 {
		return
	}
	producer.memoryLimiter.releaseTemporary(bytes)
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

func (producer *producer) beginAdmission() bool {
	atomic.AddInt64(&producer.activeAdmissions, 1)
	if atomic.LoadInt32(&producer.closing) != 0 {
		atomic.AddInt64(&producer.activeAdmissions, -1)
		return false
	}
	return true
}

func (producer *producer) endAdmission() {
	atomic.AddInt64(&producer.activeAdmissions, -1)
}

func (producer *producer) waitActiveSends() {
	for atomic.LoadInt64(&producer.activeSends) != 0 || atomic.LoadInt64(&producer.activeAdmissions) != 0 {
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
		if producer.metricsEnabled() {
			atomic.AddInt64(&producer.stats.rejectedInvalidArgument, 1)
		}
		return err
	}

	if logSize > producer.config.TotalSizeLnBytes {
		err = errors.New("the log size " + strconv.FormatInt(logSize, 10) + "is larger than the total size " + strconv.FormatInt(producer.config.TotalSizeLnBytes, 10))
		if producer.metricsEnabled() {
			atomic.AddInt64(&producer.stats.rejectedInvalidArgument, 1)
		}
		return err
	}
	batchLog.setLogSize(logSize)

	if !producer.beginSend() {
		err = errProducerClosed
		if producer.metricsEnabled() {
			atomic.AddInt64(&producer.stats.rejectedClosed, 1)
		}
		return err
	}
	defer func() {
		producer.endSend()
		if r := recover(); r != nil {
			producer.releaseBatchLog(batchLog)
			err = errProducerClosed
			if producer.metricsEnabled() {
				atomic.AddInt64(&producer.stats.rejectedClosed, 1)
			}
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
			if producer.metricsEnabled() {
				atomic.AddInt64(&producer.stats.rejectedCircuitOpen, 1)
			}
			return err
		}
	}
	batchLog.setCircuitPermit(permit)

	memoryRejected, rejectionObserved := false, false
	if memoryRejected, rejectionObserved, err = producer.reserveBatchLog(batchLog, estimateBatchLogReservationBytes(batchLog, logSize)); err != nil {
		producer.releaseBatchLog(batchLog)
		if memoryRejected {
			if rejectionObserved {
				producer.stats.addRejectedManagedAdmission(1)
			}
		} else {
			if producer.metricsEnabled() {
				atomic.AddInt64(&producer.stats.rejectedClosed, 1)
			}
		}
		return err
	}

	metricsEnabled := producer.metricsEnabled()
	var metricsGeneration uint32
	if metricsEnabled {
		// Bind terminal accounting to this admission window before the channel
		// handoff. A later Off/On change must not make terminal counters exceed
		// accepted counters for the visible interval.
		batchLog.setMetricsTracked(producer.collectionGenerationNow())
		metricsGeneration = batchLog.metricsGeneration()
	}

	select {
	case <-producer.closeCh:
		err = errProducerClosed
		producer.releaseBatchLog(batchLog)
		if producer.metricsEnabled() {
			atomic.AddInt64(&producer.stats.rejectedClosed, 1)
		}
		return err
	case producer.dispatcher.newLogRecvChan <- batchLog:
		// accepted: log enqueued to dispatcher (skip atomic ops when metrics off)
		// accepted_logs counted at the ingress boundary: a log is "accepted" once
		// SendLog successfully enqueues it. This event boundary is immutable;
		// weak consistency does not allow redefining it to aggregation time.
		// accepted_groups is counted in dispatcher when a new group is created.
		// callbacks_pending is counted here so SendLog success implies pending>=1
		// before the dispatcher aggregates the log.
		// ingress HWM observed at post-enqueue depth (the true peak), only CAS
		// when v exceeds the current HWM.
		trackedNow := metricsEnabled && producer.metricsEnabled() &&
			metricsGeneration == producer.collectionGenerationToken()
		if trackedNow {
			producer.stats.addAcceptedLogs(1)
			producer.stats.observeIngressQueueDepth(int64(len(producer.dispatcher.newLogRecvChan)))
		}
		if batchLog.Key.CallBackFun != nil {
			if trackedNow {
				producer.stats.incCallbacksPending()
			}
		}
		if metricsGeneration != 0 {
			// Always release a dispatcher waiting on this tracked handoff. The
			// current generation may have changed since the send completed.
			batchLog.setAdmissionStatsReady()
		}
	}

	return nil
}

func (producer *producer) Start() {
	atomic.StoreInt32(&producer.started, 1)
	// Enable retry-queue HWM tracking for the current level. Runtime changes use
	// the same mutex-protected queue switch in SetMetricsLevel.
	if producer.dispatcher != nil && producer.dispatcher.retryQueue != nil {
		producer.dispatcher.retryQueue.setMetricsEnabled(producer.metricsEnabled())
	}
	producer.threadPoolWaitGroup.Add(1)
	go producer.threadPool.start(producer.senderWaitGroup, producer.threadPoolWaitGroup)

	producer.dispatcherWaitGroup.Add(2)
	go producer.dispatcher.checkBatchTime(producer.dispatcherWaitGroup, producer.config)
	go producer.dispatcher.run(producer.dispatcherWaitGroup)

	level.Info(producer.logger).Log("msg", "producer start")
}

func (producer *producer) Close() {
	producer.once.Do(func() {
		// Linearize Close against SetMetricsLevel. Once closing is visible, the
		// setter cannot publish a new level after the final snapshot freezes.
		producer.metricsMu.Lock()
		atomic.StoreInt32(&producer.closing, 1)
		producer.metricsMu.Unlock()
		close(producer.closeCh)
		producer.waitActiveSends()

		close(producer.threadPool.sender.stopCh)

		close(producer.dispatcher.stopCh)
		producer.dispatcherWaitGroup.Wait()

		close(producer.threadPool.stopCh)
		producer.threadPoolWaitGroup.Wait()
		producer.senderWaitGroup.Wait()

		// Freeze the close-time clock and final sequence before marking closed,
		// so all subsequent closed snapshots are immutable.
		if producer.stats != nil {
			producer.stats.freezeCloseTime()
		}
		atomic.StoreInt32(&producer.closed, 1)
		level.Info(producer.logger).Log("msg", "producer close finish")
	})
}

func (producer *producer) ForceClose() {
	producer.once.Do(func() {
		// See Close: only the transition is serialized; the blocking drain is
		// intentionally outside metricsMu.
		producer.metricsMu.Lock()
		atomic.StoreInt32(&producer.closing, 1)
		producer.metricsMu.Unlock()
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

		// Freeze the close-time clock and final sequence before marking closed,
		// so all subsequent closed snapshots are immutable.
		if producer.stats != nil {
			producer.stats.freezeCloseTime()
		}
		atomic.StoreInt32(&producer.closed, 1)
		level.Info(producer.logger).Log("msg", "producer force close finish")
	})
}

func (producer *producer) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
	producer.cli.ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken)
}

func (producer *producer) SetAPIKey(apiKey string) {
	producer.cli.SetAPIKey(apiKey)
}

// batchStatsCounts returns only work admitted in the current collection
// generation. Work crossing an Off/On boundary cannot contaminate the new
// generation.
func (producer *producer) batchStatsCounts(batch *Batch) (logs, groups int64, tracked bool) {
	if producer == nil || producer.stats == nil || batch == nil {
		return 0, 0, false
	}
	if !producer.metricsEnabled() || batch.metricsTracking == nil ||
		batch.metricsTracking.generation != producer.collectionGenerationToken() {
		return 0, 0, false
	}
	logs, groups = batch.statsCounts()
	return logs, groups, logs > 0
}

func (producer *producer) batchLogStatsTracked(batchLog *BatchLog) bool {
	if producer == nil || producer.stats == nil || batchLog == nil {
		return false
	}
	return producer.metricsEnabled() && batchLog.metricsGeneration() == producer.collectionGenerationToken()
}

func (producer *producer) batchGenerationTracked(batch *Batch) bool {
	return producer != nil && producer.metricsEnabled() && batch != nil &&
		batch.metricsTracking != nil &&
		batch.metricsTracking.generation == producer.collectionGenerationToken()
}

// dropBatchStats records terminal log/group outcomes and unwinds pending
// callbacks for a batch that is force-dropped before a server result. It is the
// single source of truth for force-drop accounting across map, retry, taskChan,
// and handleBatch force-quit paths. No-op when collection is off or the batch
// belongs to another collection generation.
func (producer *producer) dropBatchStats(batch *Batch) {
	if producer == nil || producer.stats == nil || batch == nil {
		return
	}
	logs, groups, tracked := producer.batchStatsCounts(batch)
	if !tracked {
		return
	}
	producer.stats.addTerminalDroppedLogs(logs)
	producer.stats.addTerminalDroppedGroups(groups)
	for range batch.callBackList {
		producer.stats.decCallbacksPending()
	}
}

// dropMapBatchStats unwinds aggregate gauges in addition to the force-drop
// accounting. Used only for batches dropped from the dispatcher aggregate map.
func (producer *producer) dropMapBatchStats(batch *Batch) {
	if producer == nil || producer.stats == nil || batch == nil {
		return
	}
	if producer.batchGenerationTracked(batch) {
		producer.stats.decAggregateBatches()
		producer.stats.decAggregateGroups(int64(len(batch.logGroupList.GetLogGroups())))
	}
	producer.dropBatchStats(batch)
}

// dropBatchLogStats records a BatchLog that was accepted at the ingress
// boundary but never aggregated into a batch (e.g. tryAddLog failure,
// newLogRecvChan drain during ForceClose) as terminal-dropped.
// The log was counted as 1 accepted log at ingress; it never created a group
// (groups are created at aggregation), so terminalDroppedGroups is unchanged.
// This transfers this single log from "owned" to "terminal dropped", reducing
// owned by 1. Global conservation (accepted = terminal + owned) holds for the
// producer as a whole once all in-flight logs settle; this function does not by
// itself guarantee owned=0 globally.
// Also unwinds the pending callback if the log carried one. No-op when off.
func (producer *producer) dropBatchLogStats(batchLog *BatchLog) {
	if producer == nil || producer.stats == nil || batchLog == nil {
		return
	}
	s := producer.stats
	// The log was accepted at ingress; count it as terminal dropped so it does
	// not remain "owned" forever. 1 log, 0 groups (no group was created).
	if producer.batchLogStatsTracked(batchLog) {
		s.addTerminalDroppedLogs(1)
	}
	if producer.batchLogStatsTracked(batchLog) && batchLog.Key.CallBackFun != nil {
		s.decCallbacksPending()
	}
}
