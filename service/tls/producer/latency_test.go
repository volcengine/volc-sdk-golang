package producer

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// Producer latency histogram contract tests.
// Uses deterministic batch construction and direct sendToServer calls so the
// assertions do not depend on batching timing or sleeps.

func newLatencyTestProducer(t *testing.T, level MetricsLevel) (*producer, *mockProducerClient) {
	t.Helper()
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://127.0.0.1:0"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.MetricsLevel = level
	cfg.MaxBatchCount = 100
	cfg.MaxSenderCount = 2
	cfg.LingerTime = 50 * time.Millisecond

	mc := &mockProducerClient{}
	p := newTestProducerWithClient(cfg, mc)
	p.Start()
	t.Cleanup(func() { p.ForceClose() })
	return p, mc
}

// newBatchWithLogs creates a Batch with n logs, all under the same key,
// so they go into one PutLogs request when sent.
func newBatchWithLogs(t *testing.T, p *producer, n int) *Batch {
	t.Helper()
	cfg := p.config
	bl := &BatchLog{
		Key: BatchKey{Topic: "topic", Source: "src", FileName: "file"},
		Log: &pb.Log{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}},
	}
	if p.metricsEnabled() {
		bl.setMetricsTracked(p.collectionGenerationNow())
	}
	batch := initProducerBatch(bl, cfg) // adds first log + callback
	for i := 1; i < n; i++ {
		bl2 := &BatchLog{
			Key: BatchKey{Topic: "topic", Source: "src", FileName: "file"},
			Log: &pb.Log{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}},
		}
		if p.metricsEnabled() {
			bl2.setMetricsTracked(p.collectionGenerationNow())
		}
		batch.tryAddLog(bl2, nil)
	}
	// Simulate just-enqueued on the producer's monotonic timeline.
	batch.latencyAnchorElapsedNanos = p.stats.elapsedNanos()
	batch.latencyGeneration = p.currentDetailGeneration()
	return batch
}

func TestHistogramsNilWhenLatencyOff(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelCore)
	st := p.collectMeasurements()
	if st.BatchQueueDelayMs != nil || st.SendCallDurationMs != nil ||
		st.BatchDeliveryDurationMs != nil || st.CallbackDurationMs != nil {
		t.Errorf("all histograms must be nil when latency off")
	}
}

func TestHistogramsPresentWhenLatencyOn(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelDetailed)
	st := p.collectMeasurements()
	if st.BatchQueueDelayMs == nil || st.SendCallDurationMs == nil ||
		st.BatchDeliveryDurationMs == nil || st.CallbackDurationMs == nil {
		t.Fatalf("all histograms must be non-nil when latency on")
	}
}

func TestHistogramBucketCountFastV1(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelDetailed)
	st := p.collectMeasurements()
	if got, want := len(st.SendCallDurationMs.CumulativeBucketCounts), 14; got != want {
		t.Errorf("fast_v1 bucket count=%d, want %d", got, want)
	}
	if st.SendCallDurationMs.BucketProfile != "fast_v1" {
		t.Errorf("profile=%s, want fast_v1", st.SendCallDurationMs.BucketProfile)
	}
}

func TestDetailedUsesFastBucketsForSendCalls(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelDetailed)
	st := p.collectMeasurements()
	if got, want := len(st.SendCallDurationMs.CumulativeBucketCounts), 14; got != want {
		t.Errorf("fast_v1 bucket count=%d, want %d", got, want)
	}
}

// Per-log independence: one batch with 3 logs -> one PutLogs call with 3 logs.
// send_call count must == putLogsCount (1) and totalLogs (3) > putLogsCount (1),
// proving observations are per-request not per-log.
func TestSendCallCountScalesWithRequestsNotLogs(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)

	batch := newBatchWithLogs(t, p, 3)
	p.threadPool.sender.sendToServer(batch, 0)

	if mc.putLogsCount != 1 {
		t.Fatalf("putLogsCount=%d, want 1 (one batch = one request)", mc.putLogsCount)
	}
	if mc.totalLogs != 3 {
		t.Fatalf("totalLogs=%d, want 3 (3 logs in one request)", mc.totalLogs)
	}
	if mc.totalLogs <= mc.putLogsCount {
		t.Fatalf("cannot prove per-log independence: totalLogs=%d <= putLogsCount=%d",
			mc.totalLogs, mc.putLogsCount)
	}
	st := p.collectMeasurements()
	if got, want := st.SendCallDurationMs.Count, uint64(mc.putLogsCount); got != want {
		t.Errorf("send_call count=%d, want %d (one per PutLogs request, not per-log)", got, want)
	}
}

// batch_queue_delay: one observation per batch, first send attempt only.
func TestBatchQueueDelayCountsFirstAttemptOnly(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)

	batch := newBatchWithLogs(t, p, 2)
	p.threadPool.sender.sendToServer(batch, 0)

	st := p.collectMeasurements()
	if got, want := st.BatchQueueDelayMs.Count, uint64(1); got != want {
		t.Errorf("batch_queue_delay count=%d, want %d (one per batch first-attempt)", got, want)
	}
	if got, want := st.SendCallDurationMs.Count, uint64(mc.putLogsCount); got != want {
		t.Errorf("send_call count=%d, want %d", got, want)
	}
}

// Retry: first PutLogs fails with retryable error (HTTP 500), batch enters
// retry queue (handleFailure → addToRetryQueue → attemptCount++), then is
// retrieved via getRetryBatch and retried successfully.
// Deterministic mechanism: getRetryBatch(true) only bypasses the wall-clock
// due-time gate; the real retryable-failure → handleFailure → enqueue →
// attemptCount++ path is NOT bypassed. No sleep, no manual double-call.
// Assertions: queue size 1→0, same batch retrieved, PutLogs=2,
// send_call count=2 (every attempt), batch_queue_delay count=1 (first attempt only).
func TestRetrySendCallCountsAllAttemptsQueueDelayCountsFirstOnly(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)
	mc.failCount = 1                // first PutLogs fails
	mc.retryableFail = true         // fail with retryable HTTP 500 so batch enters retry queue
	p.config.BaseRetryBackoffMs = 0 // no backoff (deterministic, no sleep)

	batch := newBatchWithLogs(t, p, 1)
	p.threadPool.sender.sendToServer(batch, 0) // fails → handleFailure → addToRetryQueue

	// Queue should have exactly 1 batch (the failed one)
	if qs := p.threadPool.sender.retryQueue.Len(); qs != 1 {
		t.Fatalf("retry queue size=%d before retrieval, want 1", qs)
	}

	// Retrieve the batch from the retry queue (real retry path, no sleep).
	// stopFlag=true bypasses only the wall-clock due-time gate.
	retried := p.threadPool.sender.retryQueue.getRetryBatch(true)
	if len(retried) != 1 {
		t.Fatalf("expected 1 batch from retry queue, got %d", len(retried))
	}
	if retried[0] != batch {
		t.Fatalf("retrieved batch is not the same batch that was enqueued")
	}
	if retried[0].attemptCount != 1 {
		t.Fatalf("attemptCount=%d after first failure, want 1 (real retry path incremented it)", retried[0].attemptCount)
	}
	// Queue should now be empty (0)
	if qs := p.threadPool.sender.retryQueue.Len(); qs != 0 {
		t.Fatalf("retry queue size=%d after retrieval, want 0", qs)
	}

	// Second attempt: succeeds
	p.threadPool.sender.sendToServer(retried[0], 0)

	if mc.putLogsCount != 2 {
		t.Fatalf("putLogsCount=%d, want 2 (one fail + one retry)", mc.putLogsCount)
	}
	st := p.collectMeasurements()
	if got, want := st.SendCallDurationMs.Count, uint64(2); got != want {
		t.Errorf("send_call count=%d, want %d (one per PutLogs attempt incl. retry)", got, want)
	}
	if got, want := st.BatchQueueDelayMs.Count, uint64(1); got != want {
		t.Errorf("batch_queue_delay count=%d, want %d (first attempt only, attemptCount==0 guard)", got, want)
	}
}

// Detailed is the only level that allocates and records latency histograms.
func TestDetailedEnablesLatencyHistograms(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)

	batch := newBatchWithLogs(t, p, 2)
	p.threadPool.sender.sendToServer(batch, 0)

	st := p.collectMeasurements()
	if st.MetricsDisabled || st.MetricsLevel != MetricsLevelDetailed {
		t.Errorf("expected Detailed collection, got level=%v disabled=%v", st.MetricsLevel, st.MetricsDisabled)
	}
	if st.SendCallDurationMs == nil {
		t.Fatalf("send_call histogram must be present in Detailed")
	}
	if got, want := st.SendCallDurationMs.Count, uint64(mc.putLogsCount); got != want {
		t.Errorf("send_call count=%d, want %d", got, want)
	}
}

func TestCoreDoesNotCollectLatency(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelCore)

	batch := newBatchWithLogs(t, p, 2)
	p.threadPool.sender.sendToServer(batch, 0)

	st := p.collectMeasurements()
	if st.MetricsDisabled {
		t.Errorf("MetricsDisabled should be false when metrics_level=core")
	}
	if st.SendCallDurationMs != nil || st.BatchQueueDelayMs != nil ||
		st.BatchDeliveryDurationMs != nil || st.CallbackDurationMs != nil {
		t.Errorf("all histograms must be nil when latency=off even if metrics=on")
	}
}

// batch_delivery: success terminal state -> one observation per batch.
func TestBatchDeliveryCountsOnSuccess(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)

	batch := newBatchWithLogs(t, p, 2)
	p.threadPool.sender.sendToServer(batch, 0) // success

	if mc.putLogsCount != 1 {
		t.Fatalf("putLogsCount=%d, want 1", mc.putLogsCount)
	}
	st := p.collectMeasurements()
	if got, want := st.BatchDeliveryDurationMs.Count, uint64(1); got != want {
		t.Errorf("batch_delivery count=%d, want %d (one per terminal success)", got, want)
	}
}

// batch_delivery: final failure terminal state -> one observation per batch.
func TestBatchDeliveryCountsOnFinalFailure(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)
	mc.failCount = 1 // first (and only) attempt fails -> FailedCallback

	batch := newBatchWithLogs(t, p, 1)
	p.threadPool.sender.sendToServer(batch, 0) // fails -> terminal failure (no retry configured)

	st := p.collectMeasurements()
	if got, want := st.BatchDeliveryDurationMs.Count, uint64(1); got != want {
		t.Errorf("batch_delivery count=%d, want %d (one per terminal failure)", got, want)
	}
}

// callback_duration: success callback -> one observation per callback invocation.
func TestCallbackDurationCountsOnSuccessCallback(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)

	cb := &testCallback{}
	bl := &BatchLog{
		Key: BatchKey{Topic: "topic", Source: "src", FileName: "file", CallBackFun: cb},
		Log: &pb.Log{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}},
	}
	bl.setMetricsTracked(p.collectionGenerationNow())
	batch := initProducerBatch(bl, p.config)
	p.threadPool.sender.sendToServer(batch, 0) // success -> callback invoked

	if mc.putLogsCount != 1 {
		t.Fatalf("putLogsCount=%d, want 1", mc.putLogsCount)
	}
	if cb.successCount != 1 {
		t.Fatalf("callback successCount=%d, want 1", cb.successCount)
	}
	st := p.collectMeasurements()
	if got, want := st.CallbackDurationMs.Count, uint64(cb.successCount); got != want {
		t.Errorf("callback count=%d, want %d (one per callback invocation)", got, want)
	}
}

// callback_duration: failure callback -> one observation per callback invocation.
func TestCallbackDurationCountsOnFailureCallback(t *testing.T) {
	p, mc := newLatencyTestProducer(t, MetricsLevelDetailed)
	mc.failCount = 1 // fails -> FailedCallback -> Fail callback

	cb := &testCallback{}
	bl := &BatchLog{
		Key: BatchKey{Topic: "topic", Source: "src", FileName: "file", CallBackFun: cb},
		Log: &pb.Log{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}},
	}
	bl.setMetricsTracked(p.collectionGenerationNow())
	batch := initProducerBatch(bl, p.config)
	p.threadPool.sender.sendToServer(batch, 0) // fails -> Fail callback

	if cb.failCount != 1 {
		t.Fatalf("callback failCount=%d, want 1", cb.failCount)
	}
	st := p.collectMeasurements()
	if got, want := st.CallbackDurationMs.Count, uint64(cb.failCount); got != want {
		t.Errorf("callback count=%d, want %d (one per failed callback invocation)", got, want)
	}
}

// testCallback counts Success/Fail invocations.
type testCallback struct {
	successCount int
	failCount    int
}

func (c *testCallback) Success(result *Result) { c.successCount++ }
func (c *testCallback) Fail(result *Result)    { c.failCount++ }

func TestJSONMarshalOmitsHistogramsWhenMetricsOff(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelOff)
	st, getErr := p.GetStats(StatsOptions{View: StatsViewDebug})
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

// batch_queue_delay must use sealed/enqueued time, not createTime.
// A batch created long ago but just enqueued should have ~0 queue delay,
// not a large delay (which would prove it wrongly uses createTime).
func TestBatchQueueDelayUsesSealedTimeNotCreateTime(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelDetailed)

	batch := newBatchWithLogs(t, p, 1)
	// Simulate batch created 5 seconds ago (lingering/batching period)
	batch.createTime = time.Now().Add(-5 * time.Second)
	// Simulate just enqueued on the producer's monotonic timeline.
	batch.latencyAnchorElapsedNanos = p.stats.elapsedNanos()
	batch.latencyGeneration = p.currentDetailGeneration()

	p.threadPool.sender.sendToServer(batch, 0)

	st := p.collectMeasurements()
	if st.BatchQueueDelayMs == nil {
		t.Fatalf("batch_queue_delay histogram nil")
	}
	// If using createTime (wrong), SumMs would be ~5000.
	// If using sealedTime (correct, just enqueued), SumMs should be < 1000.
	if st.BatchQueueDelayMs.SumMs > 1000 {
		t.Errorf("batch_queue_delay SumMs=%.1f, expected < 1000 (must use sealed/enqueued time, not createTime)",
			st.BatchQueueDelayMs.SumMs)
	}
}

// If latencyAnchorElapsedNanos is zero (not enqueued via normal path), batch_queue_delay
// observation must be SKIPPED — no fallback to createTime.
func TestBatchQueueDelaySkippedWhenSealedTimestampZero(t *testing.T) {
	p, _ := newLatencyTestProducer(t, MetricsLevelDetailed)

	batch := newBatchWithLogs(t, p, 1)
	batch.createTime = time.Now().Add(-5 * time.Second)
	// Explicitly clear the compact timestamp to simulate a never-enqueued batch.
	batch.latencyAnchorElapsedNanos = 0

	p.threadPool.sender.sendToServer(batch, 0)

	st := p.collectMeasurements()
	if st.BatchQueueDelayMs == nil {
		t.Fatalf("batch_queue_delay histogram nil")
	}
	if st.BatchQueueDelayMs.Count != 0 {
		t.Errorf("batch_queue_delay count=%d, want 0 (must skip when sealed timestamp is zero, no createTime fallback)",
			st.BatchQueueDelayMs.Count)
	}
}
