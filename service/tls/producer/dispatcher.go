package producer

import (
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type BatchKey struct {
	Topic       string
	Source      string
	ShardHash   string
	FileName    string
	ContextFlow string
	CallBackFun CallBack
}

type BatchLog struct {
	Key BatchKey
	Log *pb.Log
	// state is written by the producer before channel send and consumed only by
	// the dispatcher after receive. The channel handoff is its happens-before
	// edge. The admission-ready bit is the only post-handoff state write; the
	// dispatcher waits for it before any non-atomic state access or clear.
	state uint64
}

const (
	// BatchLog is allocated per log; keep log size and transient flags in one word.
	batchLogCircuitPermitFlag  uint64 = 1 << 63
	batchLogReservedFlag       uint64 = 1 << 62
	batchLogAdmissionReadyFlag uint64 = 1 << 60
	batchLogGenerationShift           = 28
	batchLogGenerationMask     uint64 = ((1 << 32) - 1) << batchLogGenerationShift
	batchLogSizeMask           uint64 = (1 << batchLogGenerationShift) - 1
)

func (batchLog *BatchLog) size() int {
	if batchLog == nil || batchLog.Log == nil {
		return 0
	}
	if size := batchLog.cachedLogSize(); size > 0 {
		return size
	}
	return batchLog.Log.Size()
}

func (batchLog *BatchLog) cachedLogSize() int {
	if batchLog == nil {
		return 0
	}
	return int(batchLog.state & batchLogSizeMask)
}

func (batchLog *BatchLog) setLogSize(size int64) {
	if batchLog == nil || size <= 0 {
		return
	}
	if uint64(size) > batchLogSizeMask {
		return
	}
	flags := batchLog.state &^ batchLogSizeMask
	batchLog.state = flags | (uint64(size) & batchLogSizeMask)
}

func (batchLog *BatchLog) setCircuitPermit(permit bool) {
	if batchLog == nil || !permit {
		return
	}
	batchLog.state |= batchLogCircuitPermitFlag
}

func (batchLog *BatchLog) setReservedBytes(bytes int64) {
	if batchLog == nil {
		return
	}
	if bytes <= 0 {
		batchLog.state &^= batchLogReservedFlag
		return
	}
	batchLog.state |= batchLogReservedFlag
}

func (batchLog *BatchLog) reservedMemoryBytes() int64 {
	if batchLog == nil {
		return 0
	}
	if batchLog.state&batchLogReservedFlag == 0 {
		return 0
	}
	return estimateBatchLogReservationBytes(batchLog, int64(batchLog.cachedLogSize()))
}

func (batchLog *BatchLog) hasCircuitPermit() bool {
	return batchLog != nil && batchLog.state&batchLogCircuitPermitFlag != 0
}

// setMetricsTracked marks this log as admitted during an enabled metrics
// window. The producer writes it before channel handoff; the dispatcher and
// sender then use it to keep terminal counters conservation-safe across a
// later runtime level change.
func (batchLog *BatchLog) setMetricsTracked(generation uint64) {
	if batchLog != nil {
		const maxToken = uint64((1 << 32) - 1)
		token := (generation-1)%maxToken + 1
		batchLog.state = (batchLog.state &^ batchLogGenerationMask) | token<<batchLogGenerationShift
	}
}

func (batchLog *BatchLog) metricsTracked() bool {
	return batchLog != nil && batchLog.metricsGeneration() != 0
}

func (batchLog *BatchLog) metricsGeneration() uint32 {
	if batchLog == nil {
		return 0
	}
	state := atomic.LoadUint64(&batchLog.state)
	return uint32((state & batchLogGenerationMask) >> batchLogGenerationShift)
}

// setAdmissionStatsReady atomically publishes completion of accepted/HWM and
// optional callbacksPending accounting after a successful enqueue.
func (batchLog *BatchLog) setAdmissionStatsReady() {
	if batchLog == nil {
		return
	}
	for {
		old := atomic.LoadUint64(&batchLog.state)
		if old&batchLogAdmissionReadyFlag != 0 ||
			atomic.CompareAndSwapUint64(&batchLog.state, old, old|batchLogAdmissionReadyFlag) {
			return
		}
	}
}

func (batchLog *BatchLog) isAdmissionStatsReady() bool {
	if batchLog == nil {
		return true
	}
	return atomic.LoadUint64(&batchLog.state)&batchLogAdmissionReadyFlag != 0
}

func (batchLog *BatchLog) clearReservation() {
	if batchLog != nil {
		batchLog.state = 0
	}
}

type Dispatcher struct {
	stopCh         chan struct{}
	forceQuitCh    chan struct{}
	runDoneCh      chan struct{}
	newLogRecvChan chan *BatchLog
	retryQueue     *RetryQueue
	lock           *sync.RWMutex
	logGroupData   map[string]*Batch
	producerConfig *Config
	sender         *Sender
	logger         log.Logger
	threadPool     *ThreadPool
	producer       *producer
}

func initDispatcher(config *Config, sender *Sender, logger log.Logger, threadPool *ThreadPool, producer *producer) *Dispatcher {
	return &Dispatcher{
		stopCh:         make(chan struct{}),
		forceQuitCh:    make(chan struct{}),
		runDoneCh:      make(chan struct{}),
		newLogRecvChan: make(chan *BatchLog, config.BatchQueueSize),
		logGroupData:   make(map[string]*Batch),
		retryQueue:     sender.retryQueue,
		lock:           &sync.RWMutex{},
		producerConfig: config,
		sender:         sender,
		logger:         logger,
		threadPool:     threadPool,
		producer:       producer,
	}
}

func (dispatcher *Dispatcher) IsShutDown() bool {
	select {
	case <-dispatcher.stopCh:
		return true
	default:
		return false
	}
}

func (dispatcher *Dispatcher) run(dispatcherWaitGroup *sync.WaitGroup) {
	defer dispatcherWaitGroup.Done()
	defer close(dispatcher.runDoneCh)
	if dispatcher.producer != nil {
		if dispatcher.producer.gaugesTracked() {
			dispatcher.producer.stats.incDispatcherRunning()
		}
		defer func() {
			if dispatcher.producer.gaugesTracked() {
				dispatcher.producer.stats.decDispatcherRunning()
			}
		}()
	}

	for {
		select {
		case newBatchLog := <-dispatcher.newLogRecvChan:
			dispatcher.handleLogs(newBatchLog)
		case <-dispatcher.stopCh:
			// let the background checker to send the rest producerBatches
			return
		case <-dispatcher.forceQuitCh:
			return
		}
	}
}

func (dispatcher *Dispatcher) checkBatches(config *Config) {
	sleepMs := config.LingerTime

	dispatcher.lock.Lock()
	mapCount := len(dispatcher.logGroupData)
	for key, batch := range dispatcher.logGroupData {
		timeInterval := time.Since(batch.createTime)
		if timeInterval >= config.LingerTime {
			level.Debug(dispatcher.logger).Log("msg", "mover goroutine execute sent producerBatch to Sender")
			if !dispatcher.innerSendToServer(key, batch) {
				dispatcher.lock.Unlock()
				return
			}
		} else {
			if sleepMs > timeInterval {
				sleepMs = timeInterval
			}
		}
	}
	dispatcher.lock.Unlock()

	if mapCount == 0 {
		level.Debug(dispatcher.logger).Log("msg", "No data time in map waiting for user configured RemainMs parameter values")
		sleepMs = config.LingerTime
	}

	retryProducerBatchList := dispatcher.retryQueue.getRetryBatch(false)
	if retryProducerBatchList == nil {
		// If there is nothing to send in the retry queue, just wait for the minimum time that was given to me last time.
		for sleepMs > 0 {
			if !dispatcher.sleepWithCancel(sleepMs) {
				return
			}
			sleepMs = 0
		}
	} else {
		for i, producerBatch := range retryProducerBatchList {
			if dispatcher.sendToThreadPool(producerBatch) {
				continue
			}
			// Force-quit: sendToThreadPool returned false. Every remaining
			// retry batch must go through the centralized drop helper so
			// callbacks are unwound and terminal dropped batches are counted.
			if dispatcher.producer != nil {
				for _, batch := range retryProducerBatchList[i:] {
					dispatcher.producer.dropBatchStats(batch)
					dispatcher.producer.releaseBatchResources(batch)
				}
			}
			return
		}
	}
}

func (dispatcher *Dispatcher) sleepWithCancel(duration time.Duration) bool {
	if duration <= 0 {
		return true
	}
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-dispatcher.stopCh:
		return false
	case <-dispatcher.forceQuitCh:
		return false
	case <-timer.C:
		return true
	}
}

func (dispatcher *Dispatcher) checkBatchTime(dispatcherWaitGroup *sync.WaitGroup, config *Config) {
	defer dispatcherWaitGroup.Done()
	if dispatcher.producer != nil {
		if dispatcher.producer.gaugesTracked() {
			dispatcher.producer.stats.incDispatcherRunning()
		}
		defer func() {
			if dispatcher.producer.gaugesTracked() {
				dispatcher.producer.stats.decDispatcherRunning()
			}
		}()
	}

	for {
		select {
		case <-dispatcher.stopCh:
			// The receive loop may already own a BatchLog that is no longer in
			// newLogRecvChan. Wait for it to finish handleLogs before the final
			// channel/map drain, otherwise it can repopulate logGroupData after
			// RetryQueueElegantQuit has cleared it.
			<-dispatcher.runDoneCh
			dispatcher.RetryQueueElegantQuit()
			return
		case <-dispatcher.forceQuitCh:
			return
		default:
			dispatcher.checkBatches(config)
		}
	}
}

func (dispatcher *Dispatcher) RetryQueueElegantQuit() {
	defer close(dispatcher.threadPool.taskChan)

	close(dispatcher.newLogRecvChan)
	for batchLog := range dispatcher.newLogRecvChan {
		dispatcher.handleLogs(batchLog)
	}

	dispatcher.lock.Lock()
	for _, batch := range dispatcher.logGroupData {
		// Record the first enqueue on the producer's monotonic timeline.
		if dispatcher.producer != nil && dispatcher.producer.metricsLevelNow() == MetricsLevelDetailed &&
			batch.latencyAnchorElapsedNanos == 0 {
			batch.latencyAnchorElapsedNanos = dispatcher.producer.stats.elapsedNanos()
			batch.latencyGeneration = dispatcher.producer.currentDetailGeneration()
		}
		// stats: batch leaves aggregate map during close drain → decrement gauges
		if dispatcher.producer != nil && dispatcher.producer.batchGenerationTracked(batch) {
			dispatcher.producer.stats.decAggregateBatches()
			dispatcher.producer.stats.decAggregateGroups(int64(len(batch.logGroupList.GetLogGroups())))
		}
		dispatcher.threadPool.taskChan <- batch
		if dispatcher.producer != nil && dispatcher.producer.metricsEnabled() {
			dispatcher.producer.stats.observeSendQueueDepth(int64(len(dispatcher.threadPool.taskChan)))
		}
	}
	dispatcher.logGroupData = make(map[string]*Batch)
	dispatcher.lock.Unlock()

	// pop all retry batches
	producerBatchList := dispatcher.retryQueue.getRetryBatch(true)
	for _, producerBatch := range producerBatchList {
		// Retry batches were already sealed; only set if somehow missing.
		if dispatcher.producer != nil && dispatcher.producer.metricsLevelNow() == MetricsLevelDetailed &&
			producerBatch.latencyAnchorElapsedNanos == 0 {
			producerBatch.latencyAnchorElapsedNanos = dispatcher.producer.stats.elapsedNanos()
			producerBatch.latencyGeneration = dispatcher.producer.currentDetailGeneration()
		}
		dispatcher.threadPool.taskChan <- producerBatch
		if dispatcher.producer != nil && dispatcher.producer.metricsEnabled() {
			dispatcher.producer.stats.observeSendQueueDepth(int64(len(dispatcher.threadPool.taskChan)))
		}
	}
}

func (dispatcher *Dispatcher) releasePending() {
	if dispatcher == nil || dispatcher.producer == nil {
		return
	}
	for {
		select {
		case batchLog, ok := <-dispatcher.newLogRecvChan:
			if !ok {
				goto releaseMapAndRetry
			}
			// batchLog was accepted at ingress but never aggregated; count it as
			// terminal dropped (force_close reason) and unwind its pending callback
			// so conservation holds during drain, not only after the close catch-all.
			dispatcher.producer.dropBatchLogStats(batchLog)
			dispatcher.producer.releaseBatchLog(batchLog)
		default:
			goto releaseMapAndRetry
		}
	}

releaseMapAndRetry:
	dispatcher.lock.Lock()
	for _, batch := range dispatcher.logGroupData {
		// stats: batch dropped from aggregate map during ForceClose → use the
		// centralized map-drop helper (aggregate gauges + callbacks + dropped batch).
		dispatcher.producer.dropMapBatchStats(batch)
		dispatcher.producer.releaseBatchResources(batch)
	}
	dispatcher.logGroupData = make(map[string]*Batch)
	dispatcher.lock.Unlock()

	retryBatches := dispatcher.retryQueue.getRetryBatch(true)
	for _, batch := range retryBatches {
		// stats: retry batch dropped during ForceClose → use the centralized
		// drop helper (callbacks + dropped batch).
		dispatcher.producer.dropBatchStats(batch)
		dispatcher.producer.releaseBatchResources(batch)
	}
}

func (dispatcher *Dispatcher) handleLogs(batchLog *BatchLog) {
	if batchLog == nil {
		// dispatcher is closed
		return
	}

	// A tracked log is handed to the dispatcher before SendLog can finish its
	// post-enqueue accepted/HWM/callback accounting. Wait for that short atomic
	// sequence so terminal/group facts cannot precede admission facts, including
	// for logs without callbacks.
	// The producer always publishes readiness after a successful tracked handoff,
	// even if the collection generation changes meanwhile. closeCh is
	// NOT used here: Close/ForceClose close closeCh before waitActiveSends, and
	// an active SendLog that has completed the channel handoff will finish its
	// atomic sequence before the close proceeds. Listening on closeCh here would
	// race with that sequence and cause data loss / pending underflow.
	if batchLog.metricsTracked() {
		for !batchLog.isAdmissionStatsReady() {
			runtime.Gosched()
		}
	}

	key := dispatcher.getKeyString(batchLog.Key)
	if generation := batchLog.metricsGeneration(); generation != 0 {
		key += delimiter + uitoa(uint64(generation))
	}
	EnsureLogTime(batchLog.Log, dispatcher.producerConfig.EnableNanosecond)

	dispatcher.lock.Lock()
	defer dispatcher.lock.Unlock()
	producerBatch := dispatcher.getOrCreateProducerBatch(batchLog, key)
	statsTracked := dispatcher.producer.batchLogStatsTracked(batchLog)
	reservedBytes := batchLog.reservedMemoryBytes()
	addSuccess, deltaBytes, newGroup := producerBatch.tryAddLog(batchLog, batchLog.Key.CallBackFun)
	if addSuccess {
		dispatcher.producer.memoryLimiter.movePayload(payloadStageIngress, payloadStageAggregate, reservedBytes)
		dispatcher.producer.addProducerLogGroupSize(deltaBytes)
		// stats: count accepted groups when a new LogGroup is created;
		// track aggregate group gauge. accepted_logs and callbacks_pending
		// are counted at the ingress boundary.
		if statsTracked && newGroup {
			dispatcher.producer.stats.addAcceptedGroups(1)
		}
		if statsTracked && newGroup {
			dispatcher.producer.stats.incAggregateGroups(1)
		}
		if producerBatch.meetSendCondition(dispatcher.producerConfig) {
			dispatcher.innerSendToServer(key, producerBatch)

		}
		return
	}
	// tryAddLog failed on existing batch: send it, then the batchLog is dropped.
	// The log was accepted at ingress; count it as terminal dropped (internal
	// reason) and unwind its pending callback so conservation holds immediately.
	if !dispatcher.innerSendToServer(key, producerBatch) {
		dispatcher.producer.dropBatchLogStats(batchLog)
		dispatcher.producer.releaseBatchLog(batchLog)
		return
	}
	newBatch := newProducerBatch(batchLog, dispatcher.producerConfig)
	dispatcher.logGroupData[key] = newBatch
	// stats: new batch enters aggregate map
	if dispatcher.producer.batchLogStatsTracked(batchLog) {
		dispatcher.producer.stats.incAggregateBatches()
	}
	statsTracked = dispatcher.producer.batchLogStatsTracked(batchLog)
	reservedBytes = batchLog.reservedMemoryBytes()
	addSuccess, deltaBytes, newGroup = newBatch.tryAddLog(batchLog, batchLog.Key.CallBackFun)
	if addSuccess {
		dispatcher.producer.memoryLimiter.movePayload(payloadStageIngress, payloadStageAggregate, reservedBytes)
		dispatcher.producer.addProducerLogGroupSize(deltaBytes)
		// stats: count accepted groups when a new LogGroup is created;
		// track aggregate group gauge. callbacks_pending is counted at ingress.
		if statsTracked && newGroup {
			dispatcher.producer.stats.addAcceptedGroups(1)
		}
		if statsTracked && newGroup {
			dispatcher.producer.stats.incAggregateGroups(1)
		}
	} else {
		// tryAddLog failed on the new batch: the batchLog is dropped. It was
		// accepted at ingress; count it as terminal dropped (internal reason)
		// and unwind its pending callback so conservation holds immediately.
		dispatcher.producer.dropBatchLogStats(batchLog)
		dispatcher.producer.releaseBatchLog(batchLog)
	}
	if newBatch.meetSendCondition(dispatcher.producerConfig) {
		dispatcher.innerSendToServer(key, newBatch)
	}
}

func (dispatcher *Dispatcher) getOrCreateProducerBatch(batchLog *BatchLog, key string) *Batch {
	if producerBatch, ok := dispatcher.logGroupData[key]; ok && producerBatch != nil {
		return producerBatch
	}
	newBatch := newProducerBatch(batchLog, dispatcher.producerConfig)
	dispatcher.logGroupData[key] = newBatch
	// stats: new batch enters aggregate map
	if dispatcher.producer != nil && dispatcher.producer.batchLogStatsTracked(batchLog) {
		dispatcher.producer.stats.incAggregateBatches()
	}
	return newBatch
}

func (dispatcher *Dispatcher) sendToThreadPool(producerBatch *Batch) bool {
	// Record sealed/enqueued time BEFORE publishing to taskChan (avoids data
	// race with sender reading it). First enqueue only; retries/shutdown must
	// not overwrite. The monotonic clock is read only when latency is on.
	if dispatcher.producer != nil && dispatcher.producer.metricsLevelNow() == MetricsLevelDetailed &&
		producerBatch.latencyAnchorElapsedNanos == 0 {
		producerBatch.latencyAnchorElapsedNanos = dispatcher.producer.stats.elapsedNanos()
		producerBatch.latencyGeneration = dispatcher.producer.currentDetailGeneration()
	}
	select {
	case dispatcher.threadPool.taskChan <- producerBatch:
		if dispatcher.producer != nil {
			producerBatch.movePayloadTo(dispatcher.producer.memoryLimiter, payloadStageSend)
		}
		// stats: observe send queue depth after successful enqueue for HWM (metrics=core only)
		if dispatcher.producer != nil && dispatcher.producer.metricsEnabled() {
			dispatcher.producer.stats.observeSendQueueDepth(int64(len(dispatcher.threadPool.taskChan)))
		}
		return true
	case <-dispatcher.forceQuitCh:
		return false
	}
}

func (dispatcher *Dispatcher) innerSendToServer(key string, producerBatch *Batch) bool {
	level.Debug(dispatcher.logger).Log("msg", "Send producerBatch to Sender from dispatcher")

	if !dispatcher.sendToThreadPool(producerBatch) {
		return false
	}

	// stats: batch leaves aggregate map → decrement aggregate gauges
	if dispatcher.producer != nil && dispatcher.producer.batchGenerationTracked(producerBatch) {
		dispatcher.producer.stats.decAggregateBatches()
		dispatcher.producer.stats.decAggregateGroups(int64(len(producerBatch.logGroupList.GetLogGroups())))
	}

	delete(dispatcher.logGroupData, key)
	return true
}

func (dispatcher *Dispatcher) getKeyString(batchKey BatchKey) string {
	var key strings.Builder

	key.WriteString(batchKey.Topic)
	key.WriteString(delimiter)
	key.WriteString(batchKey.ShardHash)
	key.WriteString(delimiter)
	key.WriteString(batchKey.Source)
	key.WriteString(delimiter)
	key.WriteString(batchKey.FileName)
	key.WriteString(delimiter)
	key.WriteString(batchKey.ContextFlow)

	return key.String()
}
