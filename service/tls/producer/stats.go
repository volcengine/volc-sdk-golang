package producer

import (
	"sync/atomic"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

// histogramProfile controls which histogram buckets are used.
type histogramProfile int

const (
	histogramProfileOff histogramProfile = iota
	histogramProfileFastV1
	histogramProfileSlowV1
)

// producerState is the lifecycle state of the producer.
type producerState string

const (
	stateNew     producerState = "new"
	stateRunning producerState = "running"
	stateClosing producerState = "closing"
	stateClosed  producerState = "closed"
)

// consistencyLevel indicates snapshot consistency.
type consistencyLevel string

const (
	consistencyWeak  consistencyLevel = "weak"
	consistencyFinal consistencyLevel = "final"
)

// workerRole identifies a producer worker role for role stats.
type workerRole string

const (
	roleDispatcher workerRole = "dispatcher"
	roleSender     workerRole = "sender"
)

// queueMeasurements is the private source for the public queue contract.
type queueMeasurements struct {
	CurrentItems       uint64
	CapacityItems      uint64
	HighWatermarkItems uint64
	Bounded            bool
}

// workerRoleMeasurements reports one bounded worker-role fact.
type workerRoleMeasurements struct {
	Role     workerRole
	Expected uint32
	Running  uint32
	InFlight uint32
}

// memoryMeasurements is the private managed-memory accounting source.
type memoryMeasurements struct {
	AccountedLimitBytes           uint64
	PayloadAdmissionLimitBytes    uint64
	AccountedUsedBytes            uint64
	PayloadUsedBytes              uint64
	TemporaryUsedBytes            uint64
	AccountingRejectedTotal       uint64
	PayloadAdmissionRejectedTotal uint64
	RuntimeBufferRejectedTotal    uint64
	AccountedHighWatermarkBytes   uint64
}

type memoryGateMeasurements struct {
	LimitBytes         uint64
	UsedBytes          uint64
	HighWatermarkBytes uint64
	RejectedTotal      uint64
}

// producerMeasurements is the private, pre-aggregated input used to project
// the public view-specific Contract snapshot.
type producerMeasurements struct {
	// Common snapshot facts.
	InstanceID       string
	CapturedAtUnixMs int64
	SDKName          string
	SDKLanguage      string
	SDKVersion       string
	ComponentType    string
	State            producerState
	Consistency      consistencyLevel
	MetricsLevel     MetricsLevel
	MetricsDisabled  bool
	UptimeMs         int64
	SnapshotSequence uint64
	LastSuccessAgeMs *uint64
	LastErrorAgeMs   *uint64

	// Managed memory.
	BatchInflightMemoryGate *memoryGateMeasurements
	Memory                  *memoryMeasurements

	// Worker roles (only declared roles are output).
	WorkerRoles []workerRoleMeasurements

	// Producer core counters.
	AcceptedLogs   *uint64
	AcceptedGroups *uint64

	TerminalSuccessLogs *uint64
	TerminalFailureLogs *uint64
	TerminalDroppedLogs *uint64

	TerminalSuccessGroups *uint64
	TerminalFailureGroups *uint64
	TerminalDroppedGroups *uint64

	RejectedInvalidArgument    *uint64
	RejectedMemoryFull         *uint64
	RejectedBatchInflightLimit *uint64
	RejectedManagedAdmission   *uint64
	RejectedClosed             *uint64
	RejectedCircuitOpen        *uint64

	// Queues and in-flight facts.
	IngressQueue      *queueMeasurements
	AggregateBatches  *uint64
	AggregateGroups   *uint64
	SendQueue         *queueMeasurements
	SendCallsInFlight *uint32
	RetryQueue        *queueMeasurements
	CallbacksPending  *uint64
	CallbacksInFlight *uint32

	// Send-call and terminal counters.
	SendCallsSuccess      *uint64
	SendCallsFailure      *uint64
	SendCallPayloadLogs   *uint64
	SendCallPayloadGroups *uint64

	RetriesScheduled *uint64

	// Detailed latency histograms. Nil below MetricsLevelDetailed and updated
	// only at batch/request/callback boundaries, never per log.
	BatchQueueDelayMs       *HistogramSnapshot
	SendCallDurationMs      *HistogramSnapshot
	BatchDeliveryDurationMs *HistogramSnapshot
	CallbackDurationMs      *HistogramSnapshot
}

// collectMeasurements returns the pre-aggregated internal measurements used
// to project the public view-specific snapshot.
// It is O(1), performs no I/O, invokes no callbacks, and does not traverse
// business queues/maps. Safe to call from any goroutine at any lifecycle state.
func (producer *producer) collectMeasurements() producerMeasurements {
	if producer == nil || producer.stats == nil {
		return producerMeasurements{MetricsDisabled: true}
	}

	s := producer.stats

	// Determine lifecycle state precisely
	state := stateNew
	if atomic.LoadInt32(&producer.started) != 0 {
		state = stateRunning
	}
	closing := atomic.LoadInt32(&producer.closing) != 0
	closed := atomic.LoadInt32(&producer.closed) != 0
	if closed {
		// True closure: background drain and callback work are complete.
		state = stateClosed
	} else if closing {
		// Close has begun but background drain/callbacks may still be active.
		state = stateClosing
	}

	// consistency: only truly closed snapshots are final. Closing snapshots
	// are weak because writers (drain/callbacks) may still be updating counters.
	consistency := consistencyWeak
	if state == stateClosed {
		consistency = consistencyFinal
	}

	// Clock: when closed, use the frozen close-time clock so all time-derived
	// fields (CapturedAtUnixMs, UptimeMs, ages) are immutable across snapshots.
	// When not closed, use the current wall clock.
	now := time.Now()
	if state == stateClosed {
		now = s.closeTime
	}
	nowMs := unixMillis(now)
	nowNs := now.UnixNano()

	// snapshot_sequence: frozen at close time; continues to increment while
	// running or closing (writers still active). The frozen value is read
	// atomically via the masked high-bit protocol.
	var seq uint64
	if state == stateClosed {
		seq = s.frozenSnapshotSeq()
	} else {
		seq = s.nextSnapshotSeq()
	}

	metricsLevel := producer.metricsLevelNow()
	stats := producerMeasurements{
		InstanceID:       producer.instanceID,
		CapturedAtUnixMs: nowMs,
		SDKName:          base.SDKName,
		SDKLanguage:      "go",
		SDKVersion:       base.SDKVersion,
		ComponentType:    "producer",
		State:            state,
		Consistency:      consistency,
		MetricsLevel:     metricsLevel,
		MetricsDisabled:  metricsLevel == MetricsLevelOff,
		UptimeMs:         s.uptime(now),
		SnapshotSequence: seq,
	}

	if metricsLevel == MetricsLevelDetailed {
		stats.BatchQueueDelayMs = s.batchQueueDelay.snapshot()
		stats.SendCallDurationMs = s.sendCallDuration.snapshot()
		stats.BatchDeliveryDurationMs = s.batchDelivery.snapshot()
		stats.CallbackDurationMs = s.callbackDuration.snapshot()
	}

	// When collection is Off, only lifecycle/schema/capability facts are populated.
	if metricsLevel == MetricsLevelOff {
		return stats
	}

	// --- metrics on: populate all business counters ---

	// Read terminal outcomes before accepted totals so a weak concurrent
	// snapshot cannot observe terminal work without its earlier admission.
	termSuccLogs := atomic.LoadInt64(&s.terminalSuccessLogs)
	termFailLogs := atomic.LoadInt64(&s.terminalFailureLogs)
	termDropLogs := atomic.LoadInt64(&s.terminalDroppedLogs)
	termSuccGroups := atomic.LoadInt64(&s.terminalSuccessGroups)
	termFailGroups := atomic.LoadInt64(&s.terminalFailureGroups)
	termDropGroups := atomic.LoadInt64(&s.terminalDroppedGroups)
	acceptedLogs := atomic.LoadInt64(&s.acceptedLogs)
	acceptedGroups := atomic.LoadInt64(&s.acceptedGroups)

	stats.AcceptedLogs = uint64Ptr(uint64(acceptedLogs))
	stats.AcceptedGroups = uint64Ptr(uint64(acceptedGroups))

	stats.TerminalSuccessLogs = uint64Ptr(uint64(termSuccLogs))
	stats.TerminalFailureLogs = uint64Ptr(uint64(termFailLogs))
	stats.TerminalDroppedLogs = uint64Ptr(uint64(termDropLogs))

	stats.TerminalSuccessGroups = uint64Ptr(uint64(termSuccGroups))
	stats.TerminalFailureGroups = uint64Ptr(uint64(termFailGroups))
	stats.TerminalDroppedGroups = uint64Ptr(uint64(termDropGroups))

	// rejected by reason
	stats.RejectedInvalidArgument = uint64Ptr(uint64(atomic.LoadInt64(&s.rejectedInvalidArgument)))
	batchRejected := atomic.LoadInt64(&s.rejectedBatchInflightLimit)
	managedRejected := atomic.LoadInt64(&s.rejectedManagedAdmission)
	stats.RejectedMemoryFull = uint64Ptr(uint64(batchRejected + managedRejected))
	stats.RejectedBatchInflightLimit = uint64Ptr(uint64(batchRejected))
	stats.RejectedManagedAdmission = uint64Ptr(uint64(managedRejected))
	stats.RejectedClosed = uint64Ptr(uint64(atomic.LoadInt64(&s.rejectedClosed)))
	stats.RejectedCircuitOpen = uint64Ptr(uint64(atomic.LoadInt64(&s.rejectedCircuitOpen)))

	// aggregate
	aggBatches := atomic.LoadInt64(&s.aggregateBatches)
	stats.AggregateBatches = uint64Ptr(uint64(aggBatches))
	stats.AggregateGroups = uint64Ptr(uint64(atomic.LoadInt64(&s.aggregateGroups)))

	// send_queue is the physical sender task queue depth (len/cap of taskChan).
	stats.SendQueue = &queueMeasurements{
		CurrentItems:       uint64(producer.threadPool.sendQueueLen()),
		CapacityItems:      uint64(producer.threadPool.sendQueueCap()),
		HighWatermarkItems: uint64(atomic.LoadInt64(&s.sendQueueDepthMax)),
		Bounded:            true,
	}

	// retry_queue is an atomic observer projection maintained at heap mutations.
	retryDepth := 0
	retryDepthMax := int64(0)
	if producer.dispatcher != nil && producer.dispatcher.retryQueue != nil {
		retryDepth, retryDepthMax = producer.dispatcher.retryQueue.snapshot()
	}
	stats.RetryQueue = &queueMeasurements{
		CurrentItems:       uint64(retryDepth),
		HighWatermarkItems: uint64(retryDepthMax),
		Bounded:            false,
	}

	// ingress_queue is the physical depth of newLogRecvChan.
	ingressDepth := 0
	ingressCap := 0
	if producer.dispatcher != nil && producer.dispatcher.newLogRecvChan != nil {
		ingressDepth = len(producer.dispatcher.newLogRecvChan)
		ingressCap = cap(producer.dispatcher.newLogRecvChan)
	}
	stats.IngressQueue = &queueMeasurements{
		CurrentItems:       uint64(ingressDepth),
		CapacityItems:      uint64(ingressCap),
		HighWatermarkItems: uint64(atomic.LoadInt64(&s.ingressQueueDepthMax)),
		Bounded:            true,
	}

	// Worker roles report actual alive goroutine counts.
	// dispatcher: 2 long-lived goroutines (run + checkBatchTime)
	// sender: dynamic goroutines (0..MaxSenderCount) spawned per batch in handleBatch
	if atomic.LoadInt32(&producer.started) != 0 {
		stats.WorkerRoles = []workerRoleMeasurements{
			{Role: roleDispatcher, Expected: 2, Running: uint32(atomic.LoadInt32(&s.dispatcherRunning)), InFlight: 0},
			{Role: roleSender, Expected: uint32(producer.config.MaxSenderCount), Running: uint32(atomic.LoadInt32(&s.senderRunning)), InFlight: uint32(atomic.LoadInt64(&s.sendCallsInFlight))},
		}
	}

	// send calls
	stats.SendCallsInFlight = uint32Ptr(uint32(atomic.LoadInt64(&s.sendCallsInFlight)))
	stats.SendCallsSuccess = uint64Ptr(uint64(atomic.LoadInt64(&s.sendCallsSuccess)))
	stats.SendCallsFailure = uint64Ptr(uint64(atomic.LoadInt64(&s.sendCallsFailure)))
	stats.SendCallPayloadLogs = uint64Ptr(uint64(atomic.LoadInt64(&s.sendCallPayloadLogs)))
	stats.SendCallPayloadGroups = uint64Ptr(uint64(atomic.LoadInt64(&s.sendCallPayloadGroups)))

	// callbacks
	stats.CallbacksPending = uint64Ptr(uint64(atomic.LoadInt64(&s.callbacksPending)))
	stats.CallbacksInFlight = uint32Ptr(uint32(atomic.LoadInt64(&s.callbacksInFlight)))
	// retries
	stats.RetriesScheduled = uint64Ptr(uint64(atomic.LoadInt64(&s.retriesScheduled)))

	// memory (read from memoryLimiter at snapshot time)
	batchInflightUsed := producer.currentProducerLogGroupSize()
	if batchInflightUsed < 0 {
		batchInflightUsed = 0
	}
	stats.BatchInflightMemoryGate = &memoryGateMeasurements{
		LimitBytes:         uint64(producer.config.TotalSizeLnBytes),
		UsedBytes:          uint64(batchInflightUsed),
		HighWatermarkBytes: uint64(atomic.LoadInt64(&s.batchInflightMemoryHighWatermark)),
		RejectedTotal:      uint64(batchRejected),
	}
	if producer.memoryLimiter != nil {
		used := producer.memoryLimiter.Used()
		payloadRejected := producer.memoryLimiter.PayloadRejected()
		temporaryRejected := producer.memoryLimiter.TemporaryRejected()
		stats.Memory = &memoryMeasurements{
			AccountedLimitBytes:           uint64(producer.memoryLimiter.limit),
			PayloadAdmissionLimitBytes:    uint64(producer.memoryLimiter.payloadLimit),
			AccountedUsedBytes:            uint64(used),
			PayloadUsedBytes:              uint64(producer.memoryLimiter.PayloadUsed()),
			TemporaryUsedBytes:            uint64(producer.memoryLimiter.TemporaryUsed()),
			AccountingRejectedTotal:       uint64(payloadRejected + temporaryRejected),
			PayloadAdmissionRejectedTotal: uint64(payloadRejected),
			RuntimeBufferRejectedTotal:    uint64(temporaryRejected),
			AccountedHighWatermarkBytes:   uint64(producer.memoryLimiter.HighWatermark()),
		}
	}

	// last success/error age
	if t := atomic.LoadInt64(&s.lastSuccessTime); t > 0 {
		age := uint64((nowNs - t) / 1e6)
		stats.LastSuccessAgeMs = &age
	}
	if t := atomic.LoadInt64(&s.lastErrorTime); t > 0 {
		age := uint64((nowNs - t) / 1e6)
		stats.LastErrorAgeMs = &age
	}

	return stats
}

func unixMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// uitoa converts uint64 to decimal string without strconv import overhead.
func uitoa(n uint64) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}

func uint64Ptr(v uint64) *uint64 { return &v }
func uint32Ptr(v uint32) *uint32 { return &v }
