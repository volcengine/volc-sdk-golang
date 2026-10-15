package producer

import (
	"sync"
	"sync/atomic"
	"time"
)

// statsState holds all pre-aggregated Producer counters and gauges.
// All fields are atomic so GetStats() can read them without locking and without
// traversing business queues/maps. No background goroutine is associated with
// this struct; it is only updated at instrumentation boundaries.
type statsState struct {
	detailedOnce sync.Once
	// lifecycle
	startTime   time.Time
	snapshotSeq uint64 // high bit (1<<63) is the frozen flag; low 63 bits are the sequence

	// closeTime is set once when the producer fully closes (after all
	// background drain and callback work complete). Zero until then.
	// It freezes all time-derived snapshot fields (CapturedAtUnixMs,
	// UptimeMs, LastSuccessAgeMs, LastErrorAgeMs) so closed snapshots are
	// immutable. Written before the snapshotSeq frozen bit is CAS-set, so any
	// GetStats observing the frozen bit also observes a valid closeTime.
	closeTime time.Time

	// accepted / rejected (enqueue boundary)
	acceptedLogs   int64
	acceptedGroups int64

	// terminal outcomes by log unit (send call / callback boundary)
	terminalSuccessLogs int64
	terminalFailureLogs int64
	terminalDroppedLogs int64

	// terminal outcomes by group unit
	terminalSuccessGroups int64
	terminalFailureGroups int64
	terminalDroppedGroups int64

	// rejected by reason
	rejectedInvalidArgument    int64
	rejectedBatchInflightLimit int64
	rejectedManagedAdmission   int64
	rejectedClosed             int64
	rejectedCircuitOpen        int64

	// send call counters
	sendCallsSuccess int64
	sendCallsFailure int64

	// send call payload totals
	sendCallPayloadLogs   int64
	sendCallPayloadGroups int64

	// in-flight / queue gauges (updated at batch add/remove boundaries)
	aggregateBatches  int64
	aggregateGroups   int64
	sendCallsInFlight int64
	callbacksPending  int64
	callbacksInFlight int64

	// retry scheduled count (cumulative); depth tracked authoritatively inside RetryQueue
	retriesScheduled int64

	// send queue high watermark (observed len(taskChan) after enqueue, metrics=core only)
	sendQueueDepthMax int64

	// ingress queue high watermark (observed len(newLogRecvChan) after enqueue, metrics=core only)
	ingressQueueDepthMax int64

	// logical protobuf bytes held by aggregate/send/retry batches in the
	// current collection generation.
	batchInflightMemoryHighWatermark int64

	// worker role running counts (alive goroutines), updated at goroutine start/exit
	dispatcherRunning int32
	senderRunning     int32

	// last success/error timestamps (monotonic ns, for age derivation)
	lastSuccessTime int64
	lastErrorTime   int64

	// Latency histograms are nil below MetricsLevelDetailed:
	// off must not allocate histograms or record timing).
	batchQueueDelay  *histogramState
	sendCallDuration *histogramState
	batchDelivery    *histogramState
	callbackDuration *histogramState
}

func newStatsState(level MetricsLevel, maxSenderCount int64) *statsState {
	s := &statsState{
		startTime: time.Now(),
	}
	if level == MetricsLevelDetailed {
		s.initDetailedHistograms(maxSenderCount)
	}
	return s
}

func (s *statsState) initDetailedHistograms(maxSenderCount int64) {
	if s == nil {
		return
	}
	s.detailedOnce.Do(func() {
		shards := histogramShardCount(maxSenderCount)
		s.batchQueueDelay = newHistogramState(histogramProfileFastV1, shards)
		s.sendCallDuration = newHistogramState(histogramProfileFastV1, shards)
		s.batchDelivery = newHistogramState(histogramProfileSlowV1, shards)
		s.callbackDuration = newHistogramState(histogramProfileSlowV1, shards)
	})
}

// uptime returns the monotonic uptime in milliseconds.
func (s *statsState) uptime(now time.Time) int64 {
	return now.Sub(s.startTime).Milliseconds()
}

// elapsedNanos returns a compact monotonic timestamp relative to startTime.
// A positive value is reserved for a recorded timestamp; zero means unset.
// Keeping this as int64 avoids adding a second 24-byte time.Time to every
// producer Batch when latency metrics are disabled.
func (s *statsState) elapsedNanos() int64 {
	elapsed := time.Since(s.startTime).Nanoseconds()
	if elapsed <= 0 {
		return 1
	}
	return elapsed
}

// --- atomic helpers: accepted ---
func (s *statsState) addAcceptedLogs(n int64)   { atomic.AddInt64(&s.acceptedLogs, n) }
func (s *statsState) addAcceptedGroups(n int64) { atomic.AddInt64(&s.acceptedGroups, n) }

// --- atomic helpers: rejected ---
func (s *statsState) addRejectedBatchInflightLimit(n int64) {
	atomic.AddInt64(&s.rejectedBatchInflightLimit, n)
}
func (s *statsState) addRejectedManagedAdmission(n int64) {
	atomic.AddInt64(&s.rejectedManagedAdmission, n)
}

// --- atomic helpers: terminal logs ---
func (s *statsState) addTerminalSuccessLogs(n int64) { atomic.AddInt64(&s.terminalSuccessLogs, n) }
func (s *statsState) addTerminalFailureLogs(n int64) { atomic.AddInt64(&s.terminalFailureLogs, n) }
func (s *statsState) addTerminalDroppedLogs(n int64) { atomic.AddInt64(&s.terminalDroppedLogs, n) }

// --- atomic helpers: terminal groups ---
func (s *statsState) addTerminalSuccessGroups(n int64) { atomic.AddInt64(&s.terminalSuccessGroups, n) }
func (s *statsState) addTerminalFailureGroups(n int64) { atomic.AddInt64(&s.terminalFailureGroups, n) }
func (s *statsState) addTerminalDroppedGroups(n int64) { atomic.AddInt64(&s.terminalDroppedGroups, n) }

// --- atomic helpers: send calls ---
func (s *statsState) addSendCallsSuccess(n int64) { atomic.AddInt64(&s.sendCallsSuccess, n) }
func (s *statsState) addSendCallsFailure(n int64) { atomic.AddInt64(&s.sendCallsFailure, n) }

// --- atomic helpers: send call payload ---
func (s *statsState) addSendCallPayloadLogs(n int64)   { atomic.AddInt64(&s.sendCallPayloadLogs, n) }
func (s *statsState) addSendCallPayloadGroups(n int64) { atomic.AddInt64(&s.sendCallPayloadGroups, n) }

// --- atomic helpers: retry ---
func (s *statsState) addRetriesScheduled(n int64) { atomic.AddInt64(&s.retriesScheduled, n) }

// --- atomic helpers: send queue high watermark (observed len after enqueue) ---
func (s *statsState) observeSendQueueDepth(v int64) {
	for {
		old := atomic.LoadInt64(&s.sendQueueDepthMax)
		if v <= old || atomic.CompareAndSwapInt64(&s.sendQueueDepthMax, old, v) {
			break
		}
	}
}

// --- atomic helpers: ingress queue high watermark (observed len after enqueue) ---
// observeIngressQueueDepth updates the HWM only when v exceeds the current max.
// The CAS loop runs at most once per call: if v <= old it breaks without CAS;
// otherwise a single CAS attempts the update (retry only on lost race). Called
// from the producer at successful enqueue with the true post-enqueue depth.
func (s *statsState) observeIngressQueueDepth(v int64) {
	for {
		old := atomic.LoadInt64(&s.ingressQueueDepthMax)
		if v <= old || atomic.CompareAndSwapInt64(&s.ingressQueueDepthMax, old, v) {
			break
		}
	}
}

func (s *statsState) observeBatchInflightMemory(v int64) {
	updatePositiveHighWatermark(&s.batchInflightMemoryHighWatermark, v)
}

// --- atomic helpers: worker role running counts (alive goroutines) ---
func (s *statsState) incDispatcherRunning() { atomic.AddInt32(&s.dispatcherRunning, 1) }
func (s *statsState) decDispatcherRunning() { atomic.AddInt32(&s.dispatcherRunning, -1) }
func (s *statsState) incSenderRunning()     { atomic.AddInt32(&s.senderRunning, 1) }
func (s *statsState) decSenderRunning()     { atomic.AddInt32(&s.senderRunning, -1) }

// --- atomic helpers: in-flight ---
func (s *statsState) incSendCallsInFlight() {
	atomic.AddInt64(&s.sendCallsInFlight, 1)
}
func (s *statsState) decSendCallsInFlight() { decrementPositiveInt64(&s.sendCallsInFlight, 1) }

func (s *statsState) incAggregateBatches()       { atomic.AddInt64(&s.aggregateBatches, 1) }
func (s *statsState) decAggregateBatches()       { decrementPositiveInt64(&s.aggregateBatches, 1) }
func (s *statsState) incAggregateGroups(n int64) { atomic.AddInt64(&s.aggregateGroups, n) }
func (s *statsState) decAggregateGroups(n int64) { decrementPositiveInt64(&s.aggregateGroups, n) }

func (s *statsState) incCallbacksPending()  { atomic.AddInt64(&s.callbacksPending, 1) }
func (s *statsState) decCallbacksPending()  { decrementPositiveInt64(&s.callbacksPending, 1) }
func (s *statsState) incCallbacksInFlight() { atomic.AddInt64(&s.callbacksInFlight, 1) }
func (s *statsState) decCallbacksInFlight() { decrementPositiveInt64(&s.callbacksInFlight, 1) }

func (s *statsState) resetCollectionGauges(running bool) {
	if s == nil {
		return
	}
	atomic.StoreInt64(&s.aggregateBatches, 0)
	atomic.StoreInt64(&s.aggregateGroups, 0)
	atomic.StoreInt64(&s.sendCallsInFlight, 0)
	atomic.StoreInt64(&s.callbacksPending, 0)
	atomic.StoreInt64(&s.callbacksInFlight, 0)
	atomic.StoreInt64(&s.sendQueueDepthMax, 0)
	atomic.StoreInt64(&s.ingressQueueDepthMax, 0)
	atomic.StoreInt64(&s.batchInflightMemoryHighWatermark, 0)
	atomic.StoreInt64(&s.lastSuccessTime, 0)
	atomic.StoreInt64(&s.lastErrorTime, 0)
	atomic.StoreInt32(&s.senderRunning, 0)
	if running {
		atomic.StoreInt32(&s.dispatcherRunning, 2)
	} else {
		atomic.StoreInt32(&s.dispatcherRunning, 0)
	}
}

func decrementPositiveInt64(value *int64, amount int64) {
	if amount <= 0 {
		return
	}
	for {
		current := atomic.LoadInt64(value)
		next := current - amount
		if next < 0 {
			next = 0
		}
		if atomic.CompareAndSwapInt64(value, current, next) {
			return
		}
	}
}

// --- atomic helpers: timestamps ---
func (s *statsState) setLastSuccessTime(t int64) { atomic.StoreInt64(&s.lastSuccessTime, t) }
func (s *statsState) setLastErrorTime(t int64)   { atomic.StoreInt64(&s.lastErrorTime, t) }

// snapshotSeq uses the high bit (1<<63) as the frozen flag. The low 63 bits
// hold the sequence number. Once frozen, nextSnapshotSeq returns the masked
// final value without incrementing, and closed reads return the same value.
const (
	snapshotSeqFrozenBit uint64 = 1 << 63
	snapshotSeqMask      uint64 = snapshotSeqFrozenBit - 1
)

// nextSnapshotSeq atomically increments the sequence only when not frozen.
// If frozen, returns the masked final value (without the frozen bit).
// This makes the freeze protocol safe against concurrent GetStats/Close:
// a GetStats that observed not-closed cannot increment after the freeze.
func (s *statsState) nextSnapshotSeq() uint64 {
	for {
		old := atomic.LoadUint64(&s.snapshotSeq)
		if old&snapshotSeqFrozenBit != 0 {
			return old & snapshotSeqMask
		}
		if atomic.CompareAndSwapUint64(&s.snapshotSeq, old, old+1) {
			return old + 1
		}
	}
}

// freezeCloseTime captures the close-time clock and atomically freezes the
// snapshot sequence by CAS-setting the frozen bit. The current sequence value
// (without the bit) becomes the final sequence number. Called once, after all
// background work completes and before the producer.closed flag is set.
func (s *statsState) freezeCloseTime() {
	s.closeTime = time.Now()
	for {
		old := atomic.LoadUint64(&s.snapshotSeq)
		if old&snapshotSeqFrozenBit != 0 {
			return
		}
		if atomic.CompareAndSwapUint64(&s.snapshotSeq, old, old|snapshotSeqFrozenBit) {
			return
		}
	}
}

// frozenSnapshotSeq returns the masked final sequence value. Safe to call
// only when the producer is known to be closed.
func (s *statsState) frozenSnapshotSeq() uint64 {
	return atomic.LoadUint64(&s.snapshotSeq) & snapshotSeqMask
}
