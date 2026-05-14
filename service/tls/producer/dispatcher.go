package producer

import (
	"strings"
	"sync"
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
	// state is written by the producer goroutine before channel send and by the dispatcher after receive.
	// That channel handoff provides the happens-before edge. It is intentionally non-atomic;
	// adding multiple consumers requires a new ownership model.
	state uint64
}

const (
	// BatchLog is allocated per log; keep log size and transient flags in one word.
	batchLogCircuitPermitFlag uint64 = 1 << 63
	batchLogReservedFlag      uint64 = 1 << 62
	batchLogSizeMask          uint64 = batchLogReservedFlag - 1
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

func (batchLog *BatchLog) clearReservation() {
	if batchLog != nil {
		batchLog.state = 0
	}
}

type Dispatcher struct {
	stopCh         chan struct{}
	forceQuitCh    chan struct{}
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
			if dispatcher.producer != nil {
				for _, batch := range retryProducerBatchList[i:] {
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

	for {
		select {
		case <-dispatcher.stopCh:
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
		dispatcher.threadPool.taskChan <- batch
	}
	dispatcher.logGroupData = make(map[string]*Batch)
	dispatcher.lock.Unlock()

	// pop all retry batches
	producerBatchList := dispatcher.retryQueue.getRetryBatch(true)
	for _, producerBatch := range producerBatchList {
		dispatcher.threadPool.taskChan <- producerBatch
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
			dispatcher.producer.releaseBatchLog(batchLog)
		default:
			goto releaseMapAndRetry
		}
	}

releaseMapAndRetry:
	dispatcher.lock.Lock()
	for _, batch := range dispatcher.logGroupData {
		dispatcher.producer.releaseBatchResources(batch)
	}
	dispatcher.logGroupData = make(map[string]*Batch)
	dispatcher.lock.Unlock()

	for _, batch := range dispatcher.retryQueue.getRetryBatch(true) {
		dispatcher.producer.releaseBatchResources(batch)
	}
}

func (dispatcher *Dispatcher) handleLogs(batchLog *BatchLog) {
	if batchLog == nil {
		// dispatcher is closed
		return
	}

	key := dispatcher.getKeyString(batchLog.Key)
	EnsureLogTime(batchLog.Log, dispatcher.producerConfig.EnableNanosecond)

	dispatcher.lock.Lock()
	defer dispatcher.lock.Unlock()
	producerBatch := dispatcher.getOrCreateProducerBatch(batchLog, key)
	addSuccess, deltaBytes := producerBatch.tryAddLog(batchLog, batchLog.Key.CallBackFun)
	if addSuccess {
		dispatcher.producer.addProducerLogGroupSize(deltaBytes)
		if producerBatch.meetSendCondition(dispatcher.producerConfig) {
			dispatcher.innerSendToServer(key, producerBatch)

		}
		return
	}
	if !dispatcher.innerSendToServer(key, producerBatch) {
		dispatcher.producer.releaseBatchLog(batchLog)
		return
	}
	newBatch := newProducerBatch(batchLog, dispatcher.producerConfig)
	dispatcher.logGroupData[key] = newBatch
	addSuccess, deltaBytes = newBatch.tryAddLog(batchLog, batchLog.Key.CallBackFun)
	if addSuccess {
		dispatcher.producer.addProducerLogGroupSize(deltaBytes)
	} else {
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
	return newBatch
}

func (dispatcher *Dispatcher) sendToThreadPool(producerBatch *Batch) bool {
	select {
	case dispatcher.threadPool.taskChan <- producerBatch:
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
