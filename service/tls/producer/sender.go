package producer

import (
	"errors"
	"math/rand"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	. "github.com/volcengine/volc-sdk-golang/service/tls"
)

type CallBack interface {
	Success(result *Result)
	Fail(result *Result)
}

type Sender struct {
	stopCh               chan struct{}
	maxSender            chan int
	client               Client
	retryQueue           *RetryQueue
	logger               log.Logger
	noRetryStatusCodeMap map[int]struct{}
	producer             *producer
}

func initSender(client Client, retryQueue *RetryQueue, maxSenderCount int64, logger log.Logger, errorStatusMap map[int]struct{}, producer *producer) *Sender {
	sender := &Sender{
		stopCh:               make(chan struct{}),
		client:               client,
		retryQueue:           retryQueue,
		maxSender:            make(chan int, maxSenderCount),
		logger:               logger,
		noRetryStatusCodeMap: errorStatusMap,
		producer:             producer,
	}
	// Pre-fill the maxSender channel with stable indices 0..N-1.
	// Each sender goroutine acquires one index (exclusive from acquire to
	// release) and uses it as the callback shard index, eliminating true
	// sharing on the callback counters.
	for i := int64(0); i < maxSenderCount; i++ {
		sender.maxSender <- int(i)
	}
	return sender
}

func (sender *Sender) IsShutDown() bool {
	select {
	case <-sender.stopCh:
		return true
	default:
		return false
	}
}

func (sender *Sender) sendToServer(batch *Batch, senderIdx int) {
	level.Debug(sender.logger).Log("msg", "sender send data to server")
	var err error
	if sender.producer == nil {
		panic("producer: sender has no producer")
	}
	if batch == nil || batch.logGroupList == nil {
		panic("producer: sender has no batch payload")
	}
	producer := sender.producer
	stats := producer.stats

	// batch_queue_delay: time from batch sealed/enqueued to FIRST send attempt
	// only (not retries). Uses a compact monotonic timestamp set once before
	// taskChan publish; zero means not enqueued via the normal path.
	// — do NOT fall back to createTime (would conflate batching/linger time).
	// Gated by MetricsLevelDetailed. One observation per batch, not per log.
	if stats != nil && producer.detailGenerationEnabled(batch.latencyGeneration) &&
		producer.batchGenerationTracked(batch) && batch.attemptCount == 0 &&
		batch.latencyAnchorElapsedNanos != 0 {
		firstSendElapsed := stats.elapsedNanos()
		delayNanos := firstSendElapsed - batch.latencyAnchorElapsedNanos
		if delayNanos < 0 {
			delayNanos = 0
		}
		stats.batchQueueDelay.observe(float64(delayNanos)/float64(time.Millisecond), senderIdx)
		// The same compact field now anchors batch_delivery at the exact first
		// send attempt, matching the public latency contract.
		batch.latencyAnchorElapsedNanos = firstSendElapsed
	}

	if producer.failureController != nil && batch.getCircuitPermitCount() == 0 {
		allowed, permit := producer.failureController.allowSend()
		if !allowed {
			sender.handleFailure(batch, newCircuitOpenError(), senderIdx, false)
			return
		}
		if permit {
			batch.addCircuitPermit()
		}
	}

	actual := int64(batch.logGroupList.Size())
	delta := batch.setTotalDataSize(actual)
	if delta != 0 {
		producer.addProducerLogGroupSize(delta)
	}

	putLogsReq := &PutLogsRequest{
		TopicID:          batch.topic,
		CompressType:     "lz4",
		LogBody:          batch.logGroupList,
		EnableNanosecond: producer.config.EnableNanosecond,
	}
	logCount := batch.getLogCount()
	groupCount := len(batch.logGroupList.GetLogGroups())
	earliestLogTime := batch.getMinLogTime()
	latestLogTime := batch.getMaxLogTime()
	putLogsReq.LogCount = &logCount
	putLogsReq.EarliestLogTime = &earliestLogTime
	putLogsReq.LatestLogTime = &latestLogTime

	if batch.shardHash != nil {
		putLogsReq.HashKey = *batch.shardHash
	}

	sendBufferBytes := estimatePutLogsBufferBytes(batch.getTotalDataSize())
	if err = producer.acquireTemporaryBytes(sendBufferBytes); err != nil {
		sender.handleFailure(batch, err, senderIdx)
		return
	}
	temporaryReleased := false
	defer func() {
		if !temporaryReleased {
			producer.releaseTemporaryBytes(sendBufferBytes)
		}
	}()

	// send_call_duration: time for one PutLogs invocation. Producer-level retry
	// attempts are distinct calls. The call facts use their own collection token:
	// admission may belong to an older generation, but a call that starts and
	// finishes in the current generation is still a real send-call fact.
	sendGeneration := producer.beginSendCallObservation()
	sendLatencyGeneration := producer.currentDetailGeneration()
	latencyOn := sendLatencyGeneration != 0
	gaugesOn := sendGeneration != 0

	// sendStart is captured inside the closure, immediately before the exact
	// PutLogs invocation, so send_call_duration measures the real call duration.
	var sendStart time.Time
	// stats: send_calls_in_flight wraps ONLY the underlying PutLogs invocation,
	// not callback handling or retry logic. A deferred decrement inside the
	// narrow closure ensures the gauge unwinds even if PutLogs panics. On panic,
	// the batch is recorded as terminal-dropped (internal reason) so conservation
	// holds, then the panic is re-raised for the existing outer recover. The
	// metrics-Off path remains a direct call with no atomic operations.
	var resp *CommonResponse
	if gaugesOn || latencyOn {
		func() {
			defer func() {
				if r := recover(); r != nil {
					if latencyOn && stats != nil && producer.detailGenerationEnabled(sendLatencyGeneration) {
						stats.sendCallDuration.observeSince(sendStart, senderIdx)
					}
					producer.finishSendCallObservation(sendGeneration, true, int64(logCount), int64(groupCount))
					sender.recordPutLogsPanic(batch)
					panic(r)
				}
			}()
			if latencyOn {
				sendStart = time.Now()
			}
			resp, err = sender.client.PutLogs(putLogsReq)
		}()
	} else {
		resp, err = sender.client.PutLogs(putLogsReq)
	}
	if latencyOn && stats != nil && producer.detailGenerationEnabled(sendLatencyGeneration) {
		stats.sendCallDuration.observeSince(sendStart, senderIdx)
	}
	producer.finishSendCallObservation(sendGeneration, err != nil, int64(logCount), int64(groupCount))
	producer.releaseTemporaryBytes(sendBufferBytes)
	temporaryReleased = true

	if err == nil {
		sender.handleSuccess(batch, resp, senderIdx)
		return
	}

	sender.handleFailure(batch, err, senderIdx)
}

// recordPutLogsPanic records terminal accounting for a batch whose PutLogs call
// panicked. Called before re-panicking so conservation holds: the send is
// counted as terminal-dropped and pending callbacks are unwound exactly once.
// Send-call facts are recorded by the operation-generation wrapper above.
func (sender *Sender) recordPutLogsPanic(batch *Batch) {
	if sender.producer == nil || sender.producer.stats == nil ||
		batch == nil {
		return
	}
	s := sender.producer.stats
	if logCount, groupCount, tracked := sender.producer.batchStatsCounts(batch); tracked {
		s.addTerminalDroppedLogs(logCount)
		s.addTerminalDroppedGroups(groupCount)
	}
	if sender.producer.batchGenerationTracked(batch) {
		for range batch.callBackList {
			s.decCallbacksPending()
		}
	}
}

func (sender *Sender) handleSuccess(batch *Batch, putLogsResp *CommonResponse, senderIdx int) {
	level.Debug(sender.logger).Log("msg", "sendToServer succeeded,Execute successful callback function")
	if sender.producer == nil {
		panic("producer: sender has no producer")
	}
	producer := sender.producer
	stats := producer.stats
	if producer.failureController != nil {
		producer.failureController.afterSend(classifiedFailure{}, true, batch.takeCircuitPermitCount())
	}

	batch.result.SuccessFlag = true
	if batch.attemptCount < batch.maxReservedAttempts {
		batch.result.Attempts = append(batch.result.Attempts,
			newAttempt(true, putLogsResp.RequestID, "", "", GetTimeMs(time.Now().UnixNano())))
	}

	producer.releaseBatchResources(batch)

	// batch_delivery: time from first send attempt to terminal state (success).
	// One observation per terminal delivery, NOT per-log.
	if stats != nil && producer.detailGenerationEnabled(batch.latencyGeneration) &&
		producer.batchGenerationTracked(batch) && batch.latencyAnchorElapsedNanos != 0 {
		deliveryNanos := stats.elapsedNanos() - batch.latencyAnchorElapsedNanos
		if deliveryNanos < 0 {
			deliveryNanos = 0
		}
		stats.batchDelivery.observe(float64(deliveryNanos)/float64(time.Millisecond), senderIdx)
	}

	// stats: terminal success (logs, groups, batches units)
	if logCount, groupCount, tracked := producer.batchStatsCounts(batch); tracked && stats != nil {
		stats.addTerminalSuccessLogs(logCount)
		stats.addTerminalSuccessGroups(groupCount)
		stats.setLastSuccessTime(time.Now().UnixNano())
	}

	// callback_duration: time for each user callback invocation. One per
	// callback, not per log. Its interval is independently generation-gated;
	// callback pending/in-flight facts remain tied to the admission generation.
	callbackLatencyGeneration := producer.currentDetailGeneration()
	cbLatencyOn := stats != nil && callbackLatencyGeneration != 0
	gaugesOn := stats != nil && producer.batchGenerationTracked(batch)

	// Panic-safety: if a user callback panics, the existing outer recover in
	// thread_pool catches it and the loop does not continue. This defer ensures
	// the in-flight gauge always decrements and any callbacks skipped by the
	// panic no longer remain pending. No-op on the normal completion path;
	// metrics-Off path is untouched (no atomics).
	processed := 0
	inFlight := false
	defer func() {
		if !gaugesOn || stats == nil {
			return
		}
		if inFlight && producer.batchGenerationTracked(batch) {
			stats.decCallbacksInFlight()
		}
		remaining := len(batch.callBackList) - processed
		if inFlight {
			remaining-- // the in-flight callback already had its pending decremented
		}
		if producer.batchGenerationTracked(batch) {
			for j := 0; j < remaining; j++ {
				stats.decCallbacksPending()
			}
		}
	}()

	for _, callBack := range batch.callBackList {
		// stats: callback leaves pending and starts executing
		if gaugesOn && stats != nil && producer.batchGenerationTracked(batch) {
			stats.decCallbacksPending()
			stats.incCallbacksInFlight()
			inFlight = true
		}
		var cbStart time.Time
		if cbLatencyOn {
			cbStart = time.Now()
		}
		callBack.Success(batch.result)
		if cbLatencyOn && stats != nil && producer.detailGenerationEnabled(callbackLatencyGeneration) {
			stats.callbackDuration.observeSince(cbStart, senderIdx)
		}
		if gaugesOn && stats != nil && producer.batchGenerationTracked(batch) {
			stats.decCallbacksInFlight()
			inFlight = false
		}
		processed++
	}

	batch.retryBackoffMs = 0
}

func (sender *Sender) handleFailure(batch *Batch, err error, senderIdx int, recordCircuitResult ...bool) {
	_ = level.Info(sender.logger).Log("msg", "sendToServer failed", "error", err)

	failure := classifyFailure(err, sender.noRetryStatusCodeMap)
	recordResult := true
	if len(recordCircuitResult) > 0 {
		recordResult = recordCircuitResult[0]
	}
	if recordResult && sender.producer != nil && sender.producer.failureController != nil {
		sender.producer.failureController.afterSend(failure, false, batch.takeCircuitPermitCount())
	}
	noNeedRetry := batch.attemptCount >= batch.maxRetryTimes

	if sender.IsShutDown() || !failure.Retryable || noNeedRetry {
		sender.addErrorMessageToBatchAttempt(batch, err, false)
		sender.FailedCallback(batch, senderIdx)
		return
	}

	_ = level.Debug(sender.logger).Log("msg", "Submit to the retry queue after meeting the retry criteria")

	sender.addErrorMessageToBatchAttempt(batch, err, true)

	if batch.attemptCount == 1 {
		batch.retryBackoffMs += batch.baseRetryBackoffMs
	} else {
		increaseBackoffMs := rand.Float64() * float64(batch.baseIncreaseRetryBackoffMs)
		batch.retryBackoffMs += int64(increaseBackoffMs)
	}

	if batch.retryBackoffMs > batch.maxRetryIntervalInMs {
		batch.retryBackoffMs = batch.maxRetryIntervalInMs
	}

	batch.nextRetryMs = GetTimeMs(time.Now().UnixNano()) + batch.retryBackoffMs

	if sender.producer != nil {
		batch.movePayloadTo(sender.producer.memoryLimiter, payloadStageRetry)
	}
	sender.retryQueue.addToRetryQueue(batch, sender.logger)

	// stats: retry scheduled — only after actually entering retry queue
	// (retry queue depth is tracked authoritatively inside RetryQueue)
	if sender.producer != nil && sender.producer.stats != nil && sender.producer.metricsEnabled() {
		sender.producer.stats.addRetriesScheduled(1)
	}
}

func (sender *Sender) FailedCallback(batch *Batch, senderIdx int) {
	level.Info(sender.logger).Log("msg", "sendToServer failed,Execute failed callback function")
	if sender.producer == nil {
		panic("producer: sender has no producer")
	}
	producer := sender.producer
	stats := producer.stats
	producer.releaseBatchResources(batch)

	// batch_delivery: time from first send attempt to terminal state (failure).
	// One observation per terminal delivery, NOT per-log.
	if stats != nil && producer.detailGenerationEnabled(batch.latencyGeneration) &&
		producer.batchGenerationTracked(batch) && batch.latencyAnchorElapsedNanos != 0 {
		deliveryNanos := stats.elapsedNanos() - batch.latencyAnchorElapsedNanos
		if deliveryNanos < 0 {
			deliveryNanos = 0
		}
		stats.batchDelivery.observe(float64(deliveryNanos)/float64(time.Millisecond), senderIdx)
	}

	// stats: terminal failure (logs, groups, batches units)
	if logCount, groupCount, tracked := producer.batchStatsCounts(batch); tracked && stats != nil {
		stats.addTerminalFailureLogs(logCount)
		stats.addTerminalFailureGroups(groupCount)
		stats.setLastErrorTime(time.Now().UnixNano())
	}

	// callback_duration: time for each failed callback invocation.
	callbackLatencyGeneration := producer.currentDetailGeneration()
	cbLatencyOn := stats != nil && callbackLatencyGeneration != 0
	gaugesOn := stats != nil && producer.batchGenerationTracked(batch)

	// Panic-safety: mirror handleSuccess — unwind in-flight and remaining
	// pending callbacks if a user callback panics. No-op on normal completion;
	// metrics-Off path untouched.
	processed := 0
	inFlight := false
	defer func() {
		if !gaugesOn || stats == nil {
			return
		}
		if inFlight && producer.batchGenerationTracked(batch) {
			stats.decCallbacksInFlight()
		}
		remaining := len(batch.callBackList) - processed
		if inFlight {
			remaining--
		}
		if producer.batchGenerationTracked(batch) {
			for j := 0; j < remaining; j++ {
				stats.decCallbacksPending()
			}
		}
	}()

	for _, callBack := range batch.callBackList {
		// stats: callback leaves pending and starts executing
		if gaugesOn && stats != nil && producer.batchGenerationTracked(batch) {
			stats.decCallbacksPending()
			stats.incCallbacksInFlight()
			inFlight = true
		}
		var cbStart time.Time
		if cbLatencyOn {
			cbStart = time.Now()
		}
		callBack.Fail(batch.result)
		if cbLatencyOn && stats != nil && producer.detailGenerationEnabled(callbackLatencyGeneration) {
			stats.callbackDuration.observeSince(cbStart, senderIdx)
		}
		if gaugesOn && stats != nil && producer.batchGenerationTracked(batch) {
			stats.decCallbacksInFlight()
			inFlight = false
		}
		processed++
	}
}

func (sender *Sender) addErrorMessageToBatchAttempt(producerBatch *Batch, err error, retryInfo bool) {
	if producerBatch.attemptCount < producerBatch.maxReservedAttempts {
		var tlsError *Error
		ok := errors.As(err, &tlsError)
		var attempt *Attempt
		if ok {
			if retryInfo {
				level.Info(sender.logger).Log("msg", "sendToServer failed,start retrying", "retry times", producerBatch.attemptCount, "requestId", tlsError.RequestID, "error code", tlsError.Code, "error message", tlsError.Message)
			}
			attempt = newAttempt(false, tlsError.RequestID, tlsError.Code, tlsError.Message, GetTimeMs(time.Now().UnixNano()))
		} else {
			level.Error(sender.logger).Log("msg", "putLogs internal err", err.Error())
			attempt = newAttempt(false, "", "", err.Error(), GetTimeMs(time.Now().UnixNano()))
		}
		producerBatch.result.Attempts = append(producerBatch.result.Attempts, attempt)
	}

	producerBatch.result.SuccessFlag = false
	producerBatch.attemptCount += 1
}
