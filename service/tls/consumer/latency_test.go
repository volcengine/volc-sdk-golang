package consumer

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// Consumer latency histogram contract tests.
// Deterministic: directly calls logConsumer.fetchData/consume (no sleep).
// Covers all four latency metrics, Core/Detailed gating, JSON, nil,
// concurrent, off-path clock=0, panic, stale dispatch, checkpoint loop.

func newLatencyConsumer(t *testing.T, level MetricsLevel) (*consumer, *mockConsumerClient) {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "latency-project"
	conf.TopicIDList = []string{"latency-topic"}
	conf.ConsumerGroupName = "latency-group"
	conf.ConsumerName = "latency-consumer"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = level
	conf.DataFetchIntervalInMillisecond = 10
	conf.HeartbeatIntervalInSecond = 1
	conf.FlushCheckpointIntervalSecond = 1

	mc := &mockConsumerClient{}
	c, err := newTestConsumerWithClient(context.Background(), conf, mc, func(topicID string, shardID int, l *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newTestConsumerWithClient: %v", err)
	}
	return c, mc
}

// newWorker creates a logConsumer directly for deterministic testing
// (no Start/sleep needed). The worker is initialized with a checkpoint.
func newWorker(t *testing.T, c *consumer) *logConsumer {
	t.Helper()
	shard := &tls.ConsumeShard{TopicID: "latency-topic", ShardID: 0}
	lc := c.newLogConsumer(shard)
	if err := lc.init(); err != nil {
		t.Fatalf("worker init: %v", err)
	}
	return lc
}

func TestLatencyHistogramsNilWhenOff(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelCore)
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if st.FetchCallDurationMs != nil || st.DispatchDelayMs != nil ||
		st.ProcessDurationMs != nil || st.CheckpointCallDurationMs != nil {
		t.Errorf("all histograms must be nil when latency off")
	}
}

func TestLatencyHistogramsPresentWhenOn(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if st.FetchCallDurationMs == nil || st.DispatchDelayMs == nil ||
		st.ProcessDurationMs == nil || st.CheckpointCallDurationMs == nil {
		t.Fatalf("all histograms must be non-nil when latency on")
	}
}

func TestFetchCallCountScalesWithConsumeLogs(t *testing.T) {
	c, mc := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)

	lc.fetchData()
	lc.fetchData()
	lc.fetchData()

	st := c.collectMeasurements()
	if got, want := st.FetchCallDurationMs.Count, uint64(atomic.LoadInt64(&mc.consumeLogsCount)); got != want {
		t.Errorf("fetch_call count=%d, want %d (one per ConsumeLogs call)", got, want)
	}
}

func TestFetchCallBucketCountFastV1(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if got, want := len(st.FetchCallDurationMs.CumulativeBucketCounts), 14; got != want {
		t.Errorf("fast_v1 bucket count=%d, want %d", got, want)
	}
}

func TestDetailedUsesFastBucketsForFetchCalls(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if got, want := len(st.FetchCallDurationMs.CumulativeBucketCounts), 14; got != want {
		t.Errorf("fast_v1 bucket count=%d, want %d", got, want)
	}
}

func TestDetailedEnablesLatencyHistograms(t *testing.T) {
	c, mc := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if st.MetricsDisabled || st.MetricsLevel != MetricsLevelDetailed {
		t.Errorf("expected Detailed collection, got level=%v disabled=%v", st.MetricsLevel, st.MetricsDisabled)
	}
	if st.FetchCallDurationMs == nil {
		t.Fatalf("fetch_call histogram must be present in Detailed")
	}
	if got, want := st.FetchCallDurationMs.Count, uint64(atomic.LoadInt64(&mc.consumeLogsCount)); got != want {
		t.Errorf("fetch_call count=%d, want %d", got, want)
	}
}

func TestCoreDoesNotCollectLatency(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelCore)
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if st.MetricsDisabled {
		t.Errorf("MetricsDisabled should be false when metrics_level=core")
	}
	if st.FetchCallDurationMs != nil || st.DispatchDelayMs != nil ||
		st.ProcessDurationMs != nil || st.CheckpointCallDurationMs != nil {
		t.Errorf("all histograms must be nil when latency=off even if metrics=on")
	}
}

func TestJSONMarshalOmitsHistogramsWhenMetricsOff(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelOff)
	st, getErr := c.GetStats(StatsOptions{View: StatsViewDebug})
	if getErr != nil {
		t.Fatalf("GetStats: %v", getErr)
	}
	data, err := json.Marshal(st)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if _, ok := m["latency_histograms"]; ok {
		t.Error("Off snapshot must omit latency_histograms")
	}
}

func TestNilReceiverGetStats(t *testing.T) {
	var c *consumer
	st := c.collectMeasurements()
	if !st.MetricsDisabled {
		t.Errorf("nil receiver GetStats should return MetricsDisabled=true")
	}
}

func TestConcurrentGetStatsSafe(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = c.collectMeasurements()
		}()
	}
	wg.Wait()
}

func TestOffPathClockCallsZero(t *testing.T) {
	// When latency=off, per-instance nowFunc must NOT be called in the hot path.
	c, _ := newLatencyConsumer(t, MetricsLevelCore)
	lc := newWorker(t, c)

	var clockCalls int64
	c.stats.nowFunc = func() time.Time {
		atomic.AddInt64(&clockCalls, 1)
		return time.Now()
	}

	lc.fetchData()
	lc.consume()

	if got := atomic.LoadInt64(&clockCalls); got != 0 {
		t.Errorf("off-path clock calls=%d, want 0 (latency off must not call clock)", got)
	}
}

func TestProcessDurationCountMatchesCallbackInvocations(t *testing.T) {
	var callbackCount int64
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	c.consumeFunc = func(topicID string, shardID int, l *pb.LogGroupList) {
		atomic.AddInt64(&callbackCount, 1)
	}
	lc := newWorker(t, c)
	lc.fetchData()
	lc.consume()
	lc.fetchData()
	lc.consume()

	st := c.collectMeasurements()
	if got, want := st.ProcessDurationMs.Count, uint64(atomic.LoadInt64(&callbackCount)); got != want {
		t.Errorf("process count=%d, want %d (one per callback invocation)", got, want)
	}
}

func TestProcessDurationObservedOnPanic(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	c.consumeFunc = func(topicID string, shardID int, l *pb.LogGroupList) {
		panic("test panic")
	}
	lc := newWorker(t, c)
	lc.fetchData()

	// consume runs in a goroutine with recovery (as in real consumer); wait for it
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer func() { recover() }() // consumeFunc panics; simulate GoWithRecovery
		lc.consume()
	}()
	<-done

	st := c.collectMeasurements()
	if st.ProcessDurationMs.Count != 1 {
		t.Errorf("process count=%d, want 1 (must observe even on panic)", st.ProcessDurationMs.Count)
	}
}

func TestDispatchDelayCountScalesWithBatches(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()
	lc.consume()
	lc.fetchData()
	lc.consume()

	st := c.collectMeasurements()
	if st.DispatchDelayMs.Count != 2 {
		t.Errorf("dispatch count=%d, want 2 (one per non-empty batch)", st.DispatchDelayMs.Count)
	}
}

func TestDispatchDelayNotCountedForEmptyBatch(t *testing.T) {
	c, mc := newLatencyConsumer(t, MetricsLevelDetailed)
	mc.emptyResponse = true
	lc := newWorker(t, c)
	lc.fetchData()
	lc.consume()

	st := c.collectMeasurements()
	if st.DispatchDelayMs.Count != 0 {
		t.Errorf("dispatch count=%d, want 0 (empty batch should not count)", st.DispatchDelayMs.Count)
	}
}

func TestStaleDispatchTimestampNotCarriedToNextBatch(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)

	// First batch: fetch + consume (dispatch observed, timestamp cleared)
	lc.fetchData()
	lc.consume()
	// dispatchStartTime should be zero after consume
	if !lc.dispatchStartTime.IsZero() {
		t.Fatalf("dispatchStartTime not cleared after consume: %v", lc.dispatchStartTime)
	}

	// Second batch: fetch sets new timestamp
	lc.fetchData()
	if lc.dispatchStartTime.IsZero() {
		t.Fatalf("dispatchStartTime not set on second fetch")
	}
}

func TestCheckpointCallCountScalesWithModifyCheckPoint(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()
	lc.consume() // enqueues checkpoint

	// Trigger checkpoint upload directly
	c.checkpoint.uploadCheckpoint()

	st := c.collectMeasurements()
	if st.CheckpointCallDurationMs.Count < 1 {
		t.Errorf("checkpoint count=%d, want >=1 (one per ModifyCheckPoint call)", st.CheckpointCallDurationMs.Count)
	}
}

func TestCheckpointCallCountMultipleShards(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	// Add multiple checkpoints (simulate multiple shards)
	c.checkpoint.addCheckpoint(&checkpointInfo{shardInfo: &tls.ConsumeShard{TopicID: "t", ShardID: 0}, checkpoint: "c0"})
	c.checkpoint.addCheckpoint(&checkpointInfo{shardInfo: &tls.ConsumeShard{TopicID: "t", ShardID: 1}, checkpoint: "c1"})
	c.checkpoint.addCheckpoint(&checkpointInfo{shardInfo: &tls.ConsumeShard{TopicID: "t", ShardID: 2}, checkpoint: "c2"})

	c.checkpoint.uploadCheckpoint()

	st := c.collectMeasurements()
	if st.CheckpointCallDurationMs.Count != 3 {
		t.Errorf("checkpoint count=%d, want 3 (one per ModifyCheckPoint in loop)", st.CheckpointCallDurationMs.Count)
	}
}

func TestCoreDoesNotAllocateHistogram(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelCore)
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if st.FetchCallDurationMs != nil {
		t.Errorf("histogram must be nil in Core")
	}
}

// --- Additional boundary tests ---

func TestFetchCallCountedOnFailure(t *testing.T) {
	c, mc := newLatencyConsumer(t, MetricsLevelDetailed)
	mc.consumeLogsErr = &tls.Error{HTTPCode: 500, Code: "InternalError", Message: "server error"}
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if got, want := st.FetchCallDurationMs.Count, uint64(1); got != want {
		t.Errorf("fetch_call count=%d, want %d (must count even on failure)", got, want)
	}
}

func TestFetchCallCountedOnEmptyResponse(t *testing.T) {
	c, mc := newLatencyConsumer(t, MetricsLevelDetailed)
	mc.emptyResponse = true
	lc := newWorker(t, c)
	lc.fetchData()

	st := c.collectMeasurements()
	if got, want := st.FetchCallDurationMs.Count, uint64(atomic.LoadInt64(&mc.consumeLogsCount)); got != want {
		t.Errorf("fetch_call count=%d, want %d (must count even on empty response)", got, want)
	}
}

func TestDispatchDelayCountedOnErrInvalidContent(t *testing.T) {
	c, mc := newLatencyConsumer(t, MetricsLevelDetailed)
	mc.consumeLogsInvalidContent = true
	lc := newWorker(t, c)
	lc.fetchData() // returns ErrInvalidContent but sets currLogGroupList
	lc.consume()   // should still dispatch and count

	st := c.collectMeasurements()
	if st.DispatchDelayMs.Count != 1 {
		t.Errorf("dispatch count=%d, want 1 (ErrInvalidContent path must still dispatch)", st.DispatchDelayMs.Count)
	}
	if st.ProcessDurationMs.Count != 1 {
		t.Errorf("process count=%d, want 1 (ErrInvalidContent path must call callback)", st.ProcessDurationMs.Count)
	}
}

func TestCheckpointCallCountedOnFailure(t *testing.T) {
	c, mc := newLatencyConsumer(t, MetricsLevelDetailed)
	mc.modifyCheckPointErr = &tls.Error{HTTPCode: 500, Code: "InternalError", Message: "checkpoint fail"}
	lc := newWorker(t, c)
	lc.fetchData()
	lc.consume() // enqueues checkpoint

	c.checkpoint.uploadCheckpoint()

	st := c.collectMeasurements()
	if st.CheckpointCallDurationMs.Count != 1 {
		t.Errorf("checkpoint count=%d, want 1 (must count even on failure)", st.CheckpointCallDurationMs.Count)
	}
}

func TestConcurrentObserveAndGetStatsSafe(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)

	var wg sync.WaitGroup
	// observer: continuously fetch+consume
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			lc.fetchData()
			lc.consume()
		}
	}()
	// reader: continuously GetStats
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			_ = c.collectMeasurements()
		}
	}()
	wg.Wait()
}

func TestStopHistoryStable(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()
	lc.consume()

	st1 := c.collectMeasurements()
	fetchCount1 := st1.FetchCallDurationMs.Count

	// Simulate stop: no new observations after this point
	st2 := c.collectMeasurements()
	if st2.FetchCallDurationMs.Count != fetchCount1 {
		t.Errorf("fetch count changed between GetStats calls: %d -> %d (history must be stable)",
			fetchCount1, st2.FetchCallDurationMs.Count)
	}
}

func TestStopConcurrentGetStatsSafe(t *testing.T) {
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	lc := newWorker(t, c)
	lc.fetchData()

	// Concurrent GetStats while "stopping" (no new observations)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = c.collectMeasurements()
		}()
	}
	wg.Wait()
}

func TestPerInstanceClockIsolation(t *testing.T) {
	// Two consumers with different clock overrides must not interfere.
	c1, _ := newLatencyConsumer(t, MetricsLevelDetailed)
	c2, _ := newLatencyConsumer(t, MetricsLevelDetailed)

	var c1Calls, c2Calls int64
	c1.stats.nowFunc = func() time.Time {
		atomic.AddInt64(&c1Calls, 1)
		return time.Now()
	}
	c2.stats.nowFunc = func() time.Time {
		atomic.AddInt64(&c2Calls, 1)
		return time.Now()
	}

	lc1 := newWorker(t, c1)
	lc2 := newWorker(t, c2)
	lc1.fetchData()
	lc2.fetchData()

	if c1Calls == 0 || c2Calls == 0 {
		t.Errorf("c1Calls=%d c2Calls=%d, both must be >0 (per-instance clock must work independently)", c1Calls, c2Calls)
	}
}

func TestClockSeamControlsBothStartAndEnd(t *testing.T) {
	// Mutation-proof test: observeSince must use the per-instance fake clock
	// for BOTH start and end. If it falls back to real time.Since(start),
	// the fake clock cannot control the measured duration.
	h := newHistogramState(histogramProfileFastV1)
	if h == nil {
		t.Fatal("histogram must be non-nil for fast_v1")
	}

	// Fake clock returns a fixed base time; second call advances by exactly 100ms.
	base := time.Unix(0, 0)
	callCount := 0
	fakeClock := func() time.Time {
		callCount++
		if callCount == 1 {
			return base
		}
		return base.Add(100 * time.Millisecond)
	}

	start := fakeClock()             // callCount=1, returns base
	h.observeSince(start, fakeClock) // callCount=2, returns base+100ms

	snap := h.snapshot()
	if snap.Count != 1 {
		t.Fatalf("count=%d, want 1", snap.Count)
	}
	// If observeSince used real time.Since(start), SumMs would be ~realNow-1970 (huge).
	// With the fake clock, it must be exactly ~100ms.
	if snap.SumMs < 99.0 || snap.SumMs > 101.0 {
		t.Errorf("SumMs=%.2f, want ~100 (observeSince must use fake clock for both start and end, not real time.Since)",
			snap.SumMs)
	}
}

func TestProcessDurationObservedBeforePostCallbackWork(t *testing.T) {
	// Prove process_duration is observed before terminal accounting/checkpoint work.
	// Strategy: hold checkpoint.mapLock so addCheckpoint blocks after callback
	// returns. If observation is scoped to consumeFunc (correct), count=1 even
	// though addCheckpoint is blocked. If observation moved after addCheckpoint
	// (bug), count=0.
	c, _ := newLatencyConsumer(t, MetricsLevelDetailed)

	var callbackMu sync.Mutex
	callbackReturned := make(chan struct{})
	c.consumeFunc = func(topicID string, shardID int, l *pb.LogGroupList) {
		callbackMu.Lock()
		defer callbackMu.Unlock()
		close(callbackReturned)
	}
	lc := newWorker(t, c)
	lc.fetchData()

	// Hold mapLock so addCheckpoint blocks after callback returns.
	c.checkpoint.mapLock.Lock()
	defer c.checkpoint.mapLock.Unlock()

	go lc.consume()

	<-callbackReturned
	// Acquire callbackMu to guarantee callback has fully returned and
	// consumeWithTiming's defer (process_duration observation) has executed.
	callbackMu.Lock()
	callbackMu.Unlock()

	// Read the process histogram directly to isolate callback timing. GetStats
	// itself uses pre-aggregated checkpoint atomics and does not take mapLock.
	snap := c.stats.processDuration.snapshot()
	if snap == nil || snap.Count != 1 {
		count := uint64(0)
		if snap != nil {
			count = snap.Count
		}
		t.Errorf("process count=%d, want 1 (must be observed before addCheckpoint blocks)",
			count)
	}
	stats, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if stats.Operational == nil || stats.Operational.ProcessInFlight != 0 ||
		stats.Operational.Counters.ProcessAttemptsSuccessTotal != 1 {
		t.Fatalf("process outcome must finish at callback return, before checkpoint bookkeeping: %+v", stats.Operational)
	}
}
