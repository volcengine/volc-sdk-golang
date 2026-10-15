package consumer

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

// consumerStatsState holds all pre-aggregated counters for a consumer instance.
// All fields are pre-aggregated and atomic so GetStats does not acquire
// business locks or traverse queues/maps. No background goroutine is started
// by this type.
type consumerStatsState struct {
	detailedOnce              sync.Once
	runtimeMetricsLevel       int32
	collectionGeneration      uint64
	detailGeneration          uint64
	collectionStartedAtUnixMs int64

	// nowFunc is the per-instance clock seam for latency instrumentation.
	// Defaults to time.Now. Overridable in tests; each consumer has its own
	// copy so parallel tests / multiple consumers never interfere.
	// Only invoked when histogramProfile != Off (gated at call sites), so the
	// off hot path never dereferences it — no function call/allocation.
	nowFunc func() time.Time

	instanceID       string
	startTimeUnixMs  int64
	stopTimeUnixMs   int64 // set once at close; freezes CapturedAtUnixMs/UptimeMs for final snapshots
	snapshotSequence uint64

	// latency histograms (nil when histogramProfile == Off)
	fetchCallDuration      *histogramState
	dispatchDelay          *histogramState
	processDuration        *histogramState
	checkpointCallDuration *histogramState

	// cumulative counters (atomic int64)
	fetchCalls             int64
	fetchFailures          int64
	fetchedLogGroups       int64
	fetchedLogs            int64
	processAttempts        int64
	processFailures        int64
	processedSuccessGroups int64
	processedSuccessLogs   int64
	heartbeatCalls         int64
	heartbeatFailures      int64
	checkpointCalls        int64
	checkpointFailures     int64
	backoffEvents          int64
	rebalanceEvents        int64
	lastSuccessTimeUnixMs  int64
	lastErrorTimeUnixMs    int64

	// in-flight (atomic int32)
	fetchInFlight   int32
	processInFlight int32

	// shard/batch pre-aggregates (atomic int64), updated at mutation boundaries
	// so GetStats never traverses workerMap (O(1), no shard-map traversal).
	assignedShards int64
	heldBatchRefs  int64
	readyToProcess int64

	// worker role running counts (atomic int32)
	mainLoopRunning   int32
	heartbeatRunning  int32
	checkpointRunning int32
}

func newConsumerStatsState(level MetricsLevel) *consumerStatsState {
	now := nowMs()
	state := &consumerStatsState{
		runtimeMetricsLevel:  int32(level),
		collectionGeneration: 1,
		detailGeneration:     1,
		nowFunc:              time.Now,
		instanceID:           newObservabilityInstanceID("consumer"),
		startTimeUnixMs:      now,
	}
	if level == MetricsLevelDetailed {
		state.initDetailedHistograms()
	}
	if level != MetricsLevelOff {
		state.collectionStartedAtUnixMs = now
	}
	return state
}

func (s *consumerStatsState) collectionGenerationNow() uint64 {
	if s == nil {
		return 0
	}
	return atomic.LoadUint64(&s.collectionGeneration)
}

func (s *consumerStatsState) collectionStartedAtNow() int64 {
	if s == nil {
		return 0
	}
	return atomic.LoadInt64(&s.collectionStartedAtUnixMs)
}

func (s *consumerStatsState) metricsLevelNow() MetricsLevel {
	if s == nil {
		return MetricsLevelOff
	}
	return MetricsLevel(atomic.LoadInt32(&s.runtimeMetricsLevel))
}

func (s *consumerStatsState) isEnabled() bool {
	return s != nil && s.metricsLevelNow() != MetricsLevelOff
}

func (s *consumerStatsState) currentCollectionGeneration() uint64 {
	if !s.isEnabled() {
		return 0
	}
	return atomic.LoadUint64(&s.collectionGeneration)
}

func (s *consumerStatsState) generationEnabled(generation uint64) bool {
	return generation != 0 && s.isEnabled() &&
		atomic.LoadUint64(&s.collectionGeneration) == generation
}

// gaugesTracked is false at Off: collection is stopped, not merely hidden.
func (s *consumerStatsState) gaugesTracked() bool {
	return s.isEnabled()
}

// isLatencyEnabled reports whether Detailed collection is currently active.
func (s *consumerStatsState) isLatencyEnabled() bool {
	return s.metricsLevelNow() == MetricsLevelDetailed
}

func (s *consumerStatsState) currentDetailedGeneration() uint64 {
	if !s.isLatencyEnabled() {
		return 0
	}
	return atomic.LoadUint64(&s.detailGeneration)
}

func (s *consumerStatsState) detailedGenerationEnabled(generation uint64) bool {
	return generation != 0 && s.metricsLevelNow() == MetricsLevelDetailed &&
		atomic.LoadUint64(&s.detailGeneration) == generation
}

func (s *consumerStatsState) initDetailedHistograms() {
	if s == nil {
		return
	}
	s.detailedOnce.Do(func() {
		s.fetchCallDuration = newHistogramState(histogramProfileFastV1)
		s.dispatchDelay = newHistogramState(histogramProfileFastV1)
		s.processDuration = newHistogramState(histogramProfileSlowV1)
		s.checkpointCallDuration = newHistogramState(histogramProfileSlowV1)
	})
}

// --- operation-boundary counter mutators (no-op when disabled) ---

func (s *consumerStatsState) beginFetchCall() uint64 {
	generation := s.currentCollectionGeneration()
	if generation != 0 {
		atomic.AddInt32(&s.fetchInFlight, 1)
	}
	return generation
}

func (s *consumerStatsState) finishFetchCall(generation uint64, failed bool) {
	if !s.generationEnabled(generation) {
		return
	}
	decrementPositiveInt32(&s.fetchInFlight)
	atomic.AddInt64(&s.fetchCalls, 1)
	if failed {
		atomic.AddInt64(&s.fetchFailures, 1)
		s.markError()
	}
}

func (s *consumerStatsState) addFetchedFor(generation uint64, groups, logs int64) {
	if s.generationEnabled(generation) {
		atomic.AddInt64(&s.fetchedLogGroups, groups)
		atomic.AddInt64(&s.fetchedLogs, logs)
		s.markSuccess()
	}
}

func (s *consumerStatsState) markSuccess() {
	if s.isEnabled() {
		atomic.StoreInt64(&s.lastSuccessTimeUnixMs, nowMs())
	}
}

func (s *consumerStatsState) markError() {
	if s.isEnabled() {
		atomic.StoreInt64(&s.lastErrorTimeUnixMs, nowMs())
	}
}

func (s *consumerStatsState) beginProcessAttempt() uint64 {
	generation := s.currentCollectionGeneration()
	if generation != 0 {
		atomic.AddInt32(&s.processInFlight, 1)
	}
	return generation
}

func (s *consumerStatsState) finishProcessAttempt(generation uint64, groups, logs int64, succeeded bool) {
	if !s.generationEnabled(generation) {
		return
	}
	decrementPositiveInt32(&s.processInFlight)
	atomic.AddInt64(&s.processAttempts, 1)
	if succeeded {
		atomic.AddInt64(&s.processedSuccessGroups, groups)
		atomic.AddInt64(&s.processedSuccessLogs, logs)
		s.markSuccess()
		return
	}
	atomic.AddInt64(&s.processFailures, 1)
	s.markError()
}

func (s *consumerStatsState) beginCheckpointCall() uint64 {
	return s.currentCollectionGeneration()
}

func (s *consumerStatsState) finishCheckpointCall(generation uint64, failed bool) {
	if !s.generationEnabled(generation) {
		return
	}
	atomic.AddInt64(&s.checkpointCalls, 1)
	if failed {
		atomic.AddInt64(&s.checkpointFailures, 1)
		s.markError()
		return
	}
	s.markSuccess()
}

func (s *consumerStatsState) beginHeartbeatCall() uint64 {
	return s.currentCollectionGeneration()
}

func (s *consumerStatsState) finishHeartbeatCall(generation uint64, failed bool) {
	if !s.generationEnabled(generation) {
		return
	}
	atomic.AddInt64(&s.heartbeatCalls, 1)
	if failed {
		atomic.AddInt64(&s.heartbeatFailures, 1)
		s.markError()
		return
	}
	s.markSuccess()
}

func (s *consumerStatsState) incBackoffEvents() {
	if s.isEnabled() {
		atomic.AddInt64(&s.backoffEvents, 1)
	}
}

func (s *consumerStatsState) incRebalanceEvents() {
	if s.isEnabled() {
		atomic.AddInt64(&s.rebalanceEvents, 1)
	}
}

// --- shard/batch pre-aggregates (updated at mutation boundaries) ---

func (s *consumerStatsState) setAssignedShards(n int) {
	if s.gaugesTracked() {
		atomic.StoreInt64(&s.assignedShards, int64(n))
	}
}

func (s *consumerStatsState) incHeldBatchRefs() {
	if s.gaugesTracked() {
		atomic.AddInt64(&s.heldBatchRefs, 1)
	}
}

func (s *consumerStatsState) decHeldBatchRefs() {
	if s.gaugesTracked() {
		decrementPositiveInt64(&s.heldBatchRefs, 1)
	}
}

func (s *consumerStatsState) incReadyToProcess() {
	if s.gaugesTracked() {
		atomic.AddInt64(&s.readyToProcess, 1)
	}
}

func (s *consumerStatsState) decReadyToProcess() {
	if s.gaugesTracked() {
		decrementPositiveInt64(&s.readyToProcess, 1)
	}
}

// Worker role counts are per-instance lifecycle facts. Keep them updated at
// goroutine entry/exit even when collection is Off; collection generations do
// not reset these counts.

func (s *consumerStatsState) incMainLoopRunning() {
	if s != nil {
		atomic.AddInt32(&s.mainLoopRunning, 1)
	}
}
func (s *consumerStatsState) decMainLoopRunning() {
	if s != nil {
		decrementPositiveInt32(&s.mainLoopRunning)
	}
}
func (s *consumerStatsState) incHeartbeatRunning() {
	if s != nil {
		atomic.AddInt32(&s.heartbeatRunning, 1)
	}
}
func (s *consumerStatsState) decHeartbeatRunning() {
	if s != nil {
		decrementPositiveInt32(&s.heartbeatRunning)
	}
}
func (s *consumerStatsState) incCheckpointRunning() {
	if s != nil {
		atomic.AddInt32(&s.checkpointRunning, 1)
	}
}
func (s *consumerStatsState) decCheckpointRunning() {
	if s != nil {
		decrementPositiveInt32(&s.checkpointRunning)
	}
}

// --- snapshot ---

// snapshotSequence uses the high bit (1<<63) as the frozen flag. The low 63
// bits hold the sequence number. Once frozen, nextSnapshotSeq returns the
// masked final value without incrementing, and closed reads return the same
// value. This mirrors the producer freeze protocol (see producer.statsState).
const (
	snapshotSeqFrozenBit uint64 = 1 << 63
	snapshotSeqMask      uint64 = snapshotSeqFrozenBit - 1
)

// nextSnapshotSeq atomically increments the sequence only when not frozen.
// If frozen, returns the masked final value (without the frozen bit).
// This makes the freeze protocol safe against concurrent GetStats/Stop:
// a GetStats that observed not-closed cannot increment after the freeze,
// and the frozen high bit is never exposed to the user.
func (s *consumerStatsState) nextSnapshotSeq() uint64 {
	for {
		old := atomic.LoadUint64(&s.snapshotSequence)
		if old&snapshotSeqFrozenBit != 0 {
			return old & snapshotSeqMask
		}
		if atomic.CompareAndSwapUint64(&s.snapshotSequence, old, old+1) {
			return old + 1
		}
	}
}

// frozenSnapshotSeq returns the masked final sequence value. Safe to call
// only when the consumer is known to be closed.
func (s *consumerStatsState) frozenSnapshotSeq() uint64 {
	return atomic.LoadUint64(&s.snapshotSequence) & snapshotSeqMask
}

// freezeOnClose captures the close-time clock and marks the snapshot sequence
// as frozen so subsequent GetStats calls return an immutable final snapshot.
// Called once, after all background work completes and before the consumer's
// closed flag is set, so any GetStats observing closed also observes the
// frozen clock. Idempotent: stopTime is stored only on the first call (the
// one that actually sets the frozen bit), so concurrent/repeated Stop calls
// cannot rewrite the final clock.
func (s *consumerStatsState) freezeOnClose() {
	for {
		old := atomic.LoadUint64(&s.snapshotSequence)
		if old&snapshotSeqFrozenBit != 0 {
			return
		}
		if atomic.CompareAndSwapUint64(&s.snapshotSequence, old, old|snapshotSeqFrozenBit) {
			// Only the caller that wins the CAS stores the final clock, so a
			// second concurrent Stop cannot overwrite it.
			atomic.StoreInt64(&s.stopTimeUnixMs, nowMs())
			return
		}
	}
}

// resumeForStart thaws a snapshot frozen by the previous lifecycle
// generation. The caller must hold the consumer metrics write lock: public
// GetStats holds the matching read lock, so it cannot observe a closed/frozen
// and running/unfrozen transition half-way through. Counters, InstanceID, and
// the low sequence value are intentionally retained across generations.
func (s *consumerStatsState) resumeForStart() {
	if s == nil {
		return
	}
	for {
		old := atomic.LoadUint64(&s.snapshotSequence)
		if old&snapshotSeqFrozenBit == 0 {
			break
		}
		if atomic.CompareAndSwapUint64(&s.snapshotSequence, old, old&snapshotSeqMask) {
			break
		}
	}
	atomic.StoreInt64(&s.stopTimeUnixMs, 0)
	if s.isEnabled() {
		// A restart begins a new collection baseline while retaining the
		// instance identity and cumulative counters. Reset only generation-scoped
		// gauges/ages and publish the new collection window under the caller's
		// metrics write lock.
		atomic.AddUint64(&s.collectionGeneration, 1)
		s.resetCollectionGauges()
		atomic.StoreInt64(&s.assignedShards, 0)
		atomic.StoreInt64(&s.collectionStartedAtUnixMs, nowMs())
	}
}

func (s *consumerStatsState) snapshot(c *consumer) consumerMeasurements {
	// Read the lifecycle state ONCE so state / consistency / clock / sequence
	// are mutually consistent within a single snapshot. A snapshot must never
	// mix stateClosed with a live clock or weak consistency.
	state := c.lifecycleState()
	closed := state == stateClosed

	// Sequence: frozen at close time; continues to increment while running or
	// closing (writers may still be active). The frozen value is read atomically
	// via the masked high-bit protocol; nextSnapshotSeq never exposes the bit.
	var seq uint64
	if closed {
		seq = s.frozenSnapshotSeq()
	} else {
		seq = s.nextSnapshotSeq()
	}

	// Clock: when closed, use the frozen close-time clock so all time-derived
	// fields (CapturedAtUnixMs, UptimeMs) are immutable across snapshots.
	// When not closed, use the current wall clock.
	var now int64
	if closed {
		now = atomic.LoadInt64(&s.stopTimeUnixMs)
	} else {
		now = nowMs()
	}

	// Consistency: only truly closed snapshots are final. Running/closing
	// snapshots are weak because writers may still be updating counters.
	consistency := consistencyWeak
	if closed {
		consistency = consistencyFinal
	}

	st := consumerMeasurements{
		InstanceID:       s.instanceID,
		CapturedAtUnixMs: now,
		SDKName:          base.SDKName,
		SDKLanguage:      "go",
		SDKVersion:       base.SDKVersion,
		ComponentType:    "consumer",
		State:            state,
		Consistency:      consistency,
		MetricsLevel:     s.metricsLevelNow(),
		MetricsDisabled:  !s.isEnabled(),
		UptimeMs:         now - s.startTimeUnixMs,
		SnapshotSequence: seq,
	}

	if s.metricsLevelNow() == MetricsLevelDetailed {
		st.FetchCallDurationMs = s.fetchCallDuration.snapshot()
		st.DispatchDelayMs = s.dispatchDelay.snapshot()
		st.ProcessDurationMs = s.processDuration.snapshot()
		st.CheckpointCallDurationMs = s.checkpointCallDuration.snapshot()
	}

	if !s.isEnabled() {
		// Off: return only lifecycle/schema fields; business counters absent.
		return st
	}

	// Worker roles (declared roles only)
	st.WorkerRoles = []workerRoleMeasurements{
		{Role: roleMainLoop, Expected: 1, Running: uint32(atomic.LoadInt32(&s.mainLoopRunning)), InFlight: 0},
		{Role: roleHeartbeat, Expected: 1, Running: uint32(atomic.LoadInt32(&s.heartbeatRunning)), InFlight: 0},
		{Role: roleCheckpoint, Expected: 1, Running: uint32(atomic.LoadInt32(&s.checkpointRunning)), InFlight: 0},
	}

	// Assignment & shard workers (read from pre-aggregated atomics; no map traversal)
	assigned := uint32(atomic.LoadInt64(&s.assignedShards))
	st.AssignedShards = &assigned

	// Held batch references / ready-to-process (pre-aggregated atomics)
	held := uint64(atomic.LoadInt64(&s.heldBatchRefs))
	ready := uint64(atomic.LoadInt64(&s.readyToProcess))
	st.HeldBatchReferences = &held
	st.ReadyToProcess = &ready

	// In-flight
	fif := uint32(atomic.LoadInt32(&s.fetchInFlight))
	pif := uint32(atomic.LoadInt32(&s.processInFlight))
	st.FetchInFlight = &fif
	st.ProcessInFlight = &pif

	// Checkpoint pending (pre-aggregated atomics; no map traversal or lock)
	if cp := c.checkpointSnapshot(); cp != nil {
		st.CheckpointPending = cp
	}

	// Cumulative counters
	v := func(n int64) *uint64 { r := uint64(n); return &r }
	st.FetchCalls = v(atomic.LoadInt64(&s.fetchCalls))
	st.FetchFailures = v(atomic.LoadInt64(&s.fetchFailures))
	st.FetchedLogGroups = v(atomic.LoadInt64(&s.fetchedLogGroups))
	st.FetchedLogs = v(atomic.LoadInt64(&s.fetchedLogs))
	st.ProcessAttempts = v(atomic.LoadInt64(&s.processAttempts))
	st.ProcessFailures = v(atomic.LoadInt64(&s.processFailures))
	st.ProcessedSuccessGroups = v(atomic.LoadInt64(&s.processedSuccessGroups))
	st.ProcessedSuccessLogs = v(atomic.LoadInt64(&s.processedSuccessLogs))
	st.HeartbeatCalls = v(atomic.LoadInt64(&s.heartbeatCalls))
	st.HeartbeatFailures = v(atomic.LoadInt64(&s.heartbeatFailures))
	st.CheckpointCalls = v(atomic.LoadInt64(&s.checkpointCalls))
	st.CheckpointFailures = v(atomic.LoadInt64(&s.checkpointFailures))
	st.BackoffEvents = v(atomic.LoadInt64(&s.backoffEvents))
	st.RebalanceEvents = v(atomic.LoadInt64(&s.rebalanceEvents))
	if timestamp := atomic.LoadInt64(&s.lastSuccessTimeUnixMs); timestamp > 0 && now >= timestamp {
		age := uint64(now - timestamp)
		st.LastSuccessAgeMs = &age
	}
	if timestamp := atomic.LoadInt64(&s.lastErrorTimeUnixMs); timestamp > 0 && now >= timestamp {
		age := uint64(now - timestamp)
		st.LastErrorAgeMs = &age
	}

	return st
}

func decrementPositiveInt32(value *int32) {
	for {
		current := atomic.LoadInt32(value)
		if current <= 0 || atomic.CompareAndSwapInt32(value, current, current-1) {
			return
		}
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

func (s *consumerStatsState) resetCollectionGauges() {
	atomic.StoreInt32(&s.fetchInFlight, 0)
	atomic.StoreInt32(&s.processInFlight, 0)
	atomic.StoreInt64(&s.heldBatchRefs, 0)
	atomic.StoreInt64(&s.readyToProcess, 0)
	atomic.StoreInt64(&s.lastSuccessTimeUnixMs, 0)
	atomic.StoreInt64(&s.lastErrorTimeUnixMs, 0)
}
