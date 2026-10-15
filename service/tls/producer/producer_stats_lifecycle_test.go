package producer

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// blockingMockClient blocks PutLogs on a channel.
type blockingMockClient struct {
	mockProducerClient
	release chan struct{}
}

func newBlockingMockClient() *blockingMockClient {
	return &blockingMockClient{release: make(chan struct{})}
}

func (m *blockingMockClient) PutLogs(request *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	atomic.AddInt64(&m.putLogsCount, 1)
	<-m.release
	return &tls.CommonResponse{RequestID: "mock-putlogs"}, nil
}

func (m *blockingMockClient) releasePutLogs() { close(m.release) }

// delayMockClient sleeps a configurable duration in PutLogs.
type delayMockClient struct {
	mockProducerClient
	delay time.Duration
}

func (m *delayMockClient) PutLogs(request *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	atomic.AddInt64(&m.putLogsCount, 1)
	if m.delay > 0 {
		time.Sleep(m.delay)
	}
	return &tls.CommonResponse{RequestID: "mock-putlogs"}, nil
}

// panicMockClient panics on every PutLogs call.
type panicMockClient struct {
	mockProducerClient
}

func (m *panicMockClient) PutLogs(request *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	atomic.AddInt64(&m.putLogsCount, 1)
	panic("mock PutLogs panic")
}

// newStatsTestProducer creates a started producer with the given mock client and metrics on.
func newStatsTestProducer(t *testing.T, mc tls.Client, opts ...func(*Config)) *producer {
	t.Helper()
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://127.0.0.1:0"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.MetricsLevel = MetricsLevelCore
	cfg.MaxBatchCount = 1
	cfg.MaxSenderCount = 2
	cfg.LingerTime = 100 * time.Millisecond
	for _, opt := range opts {
		opt(cfg)
	}
	p := newTestProducerWithClient(cfg, mc)
	p.Start()
	t.Cleanup(func() { p.ForceClose() })
	return p
}

// waitFor polls cond until it returns true or the timeout elapses.
func waitFor(t *testing.T, timeout time.Duration, cond func() bool, msg string) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("timed out waiting for: %s", msg)
}

// TestIngressQueue_Populated verifies IngressQueue capacity/HWM and that
// rejected input does not change the prior HWM value.
func TestIngressQueue_Populated(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.BatchQueueSize = 500
		c.MaxBatchCount = 100
		c.LingerTime = 5 * time.Second
	})

	stats := p.collectMeasurements()
	assert.NotNil(t, stats.IngressQueue, "IngressQueue must be populated when metrics on")
	assert.Equal(t, uint64(500), stats.IngressQueue.CapacityItems, "capacity should match BatchQueueSize")
	assert.True(t, stats.IngressQueue.Bounded, "ingress queue is bounded")

	for i := 0; i < 50; i++ {
		err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
		assert.NoError(t, err, "SendLog %d should succeed", i)
	}
	waitFor(t, 5*time.Second, func() bool {
		return p.collectMeasurements().IngressQueue.HighWatermarkItems > 0
	}, "HWM should reflect successful enqueues")
	// Wait for all 50 logs to be dequeued by the dispatcher (acceptedLogs is now
	// counted at dequeue/aggregation time) so the HWM is stable before we
	// capture it and assert it does not change on rejection.
	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.AcceptedLogs != nil && *s.AcceptedLogs >= 50
	}, "all 50 logs should be dequeued and accepted")

	hwmBefore := p.collectMeasurements().IngressQueue.HighWatermarkItems

	bigLog := &pb.Log{Contents: []*pb.LogContent{{
		Key:   "k",
		Value: string(make([]byte, 10*1024*1024)),
	}}}
	err := p.SendLog("", "topic", "src", "file", bigLog, nil)
	assert.Error(t, err, "oversized log should be rejected")

	stats = p.collectMeasurements()
	assert.Equal(t, hwmBefore, stats.IngressQueue.HighWatermarkItems,
		"HWM must not change after rejected ingress")
}

// TestWaitTimeMemoryTimeoutCountsAdmissionRejection verifies the legacy
// TotalSizeLnBytes gate is part of the public admission boundary. A rejected
// SendLogs call counts as one failed admission decision; it must not invent a
// rejected-log count or enter accepted/terminal conservation.
func TestWaitTimeMemoryTimeoutCountsAdmissionRejection(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MetricsLevel = MetricsLevelCore
	cfg.TotalSizeLnBytes = 128
	cfg.MaxProducerMemoryBytes = 4096
	cfg.MaxBlockSec = 0
	p := newTestProducerWithClient(cfg, &mockProducerClient{})
	atomic.StoreInt64(&p.producerLogGroupSize, cfg.TotalSizeLnBytes+1)

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.ErrorIs(t, err, ErrMemoryTimeout)

	group := &pb.LogGroup{Logs: []*pb.Log{makeLog(), makeLog()}}
	err = p.SendLogs("", "topic", "src", "file", group, nil)
	assert.ErrorIs(t, err, ErrMemoryTimeout)

	snapshot, err := p.GetStats(StatsOptions{View: StatsViewOperational})
	assert.NoError(t, err)
	if assert.NotNil(t, snapshot.Operational) {
		assert.Equal(t, uint64(2), snapshot.Operational.Counters.RejectedMemoryFullTotal)
		assert.Equal(t, uint64(2), snapshot.Operational.Counters.RejectedBatchInflightLimitTotal)
		assert.Zero(t, snapshot.Operational.Counters.RejectedManagedAdmissionTotal)
		assert.Zero(t, snapshot.Operational.Counters.AcceptedLogsTotal)
	}
	if assert.NotNil(t, snapshot.BatchInflightMemoryGate) {
		assert.Equal(t, uint64(2), snapshot.BatchInflightMemoryGate.RejectedTotal)
	}
}

func TestWaitTimeMemoryTimeoutHonorsRuntimeMetricsLevel(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MetricsLevel = MetricsLevelOff
	cfg.TotalSizeLnBytes = 128
	cfg.MaxProducerMemoryBytes = 4096
	cfg.MaxBlockSec = 0
	p := newTestProducerWithClient(cfg, &mockProducerClient{})
	p.Start()
	t.Cleanup(p.ForceClose)
	atomic.StoreInt64(&p.producerLogGroupSize, cfg.TotalSizeLnBytes+1)

	assert.ErrorIs(t, p.SendLog("", "topic", "src", "file", makeLog(), nil), ErrMemoryTimeout)
	assert.Zero(t, atomic.LoadInt64(&p.stats.rejectedBatchInflightLimit), "Off must not collect rejection counters")

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	assert.ErrorIs(t, p.SendLog("", "topic", "src", "file", makeLog(), nil), ErrMemoryTimeout)
	snapshot, err := p.GetStats(StatsOptions{View: StatsViewOperational})
	assert.NoError(t, err)
	if assert.NotNil(t, snapshot.Operational) {
		assert.Equal(t, uint64(1), snapshot.Operational.Counters.RejectedMemoryFullTotal)
		assert.Equal(t, uint64(1), snapshot.Operational.Counters.RejectedBatchInflightLimitTotal)
	}

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelOff))
	assert.ErrorIs(t, p.SendLog("", "topic", "src", "file", makeLog(), nil), ErrMemoryTimeout)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	snapshot, err = p.GetStats(StatsOptions{View: StatsViewOperational})
	assert.NoError(t, err)
	if assert.NotNil(t, snapshot.Operational) {
		assert.Equal(t, uint64(1), snapshot.Operational.Counters.RejectedMemoryFullTotal,
			"Off-period rejections must not increment the instance counter")
	}
}

// TestCallbacksPending_AtIngressBoundary verifies callbacksPending increments
// at the successful ingress enqueue boundary, before the dispatcher aggregates.
func TestCallbacksPending_AtIngressBoundary(t *testing.T) {
	mc := &mockProducerClient{}
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://127.0.0.1:0"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.MetricsLevel = MetricsLevelCore
	cfg.MaxBatchCount = 1000
	cfg.LingerTime = 10 * time.Second
	cfg.BatchQueueSize = 100
	p := newTestProducerWithClient(cfg, mc)
	t.Cleanup(func() { p.ForceClose() })

	cb := newBlockingCallback()
	err := p.SendLog("", "topic", "src", "file", makeLog(), cb)
	assert.NoError(t, err, "SendLog should succeed at ingress")

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(1), *stats.CallbacksPending,
		"pending must be 1 immediately after successful SendLog, before dispatcher aggregation")
	assert.Equal(t, uint64(1), stats.IngressQueue.CurrentItems,
		"ingress queue current should be 1 (log not yet consumed)")
}

// TestAdmissionStatsReadyCoversAllTrackedLogs verifies that both callback
// and callback-free tracked logs publish admission bookkeeping before the
// dispatcher is allowed to process them.
func TestAdmissionStatsReadyCoversAllTrackedLogs(t *testing.T) {
	mc := &mockProducerClient{}
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://127.0.0.1:0"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.MetricsLevel = MetricsLevelCore
	cfg.MaxBatchCount = 1000
	cfg.LingerTime = 10 * time.Second
	cfg.BatchQueueSize = 100
	// Do NOT Start: logs stay in newLogRecvChan so we can inspect their state word.
	p := newTestProducerWithClient(cfg, mc)

	// Nil callback still needs accepted/group/terminal causal ordering.
	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)
	nilLog := <-p.dispatcher.newLogRecvChan
	assert.True(t, nilLog.isAdmissionStatsReady(), "tracked nil-callback log must publish admission readiness")

	// Non-nil callback: ready flag MUST be set after pending is incremented.
	cb := newBlockingCallback()
	err = p.SendLog("", "topic", "src", "file", makeLog(), cb)
	assert.NoError(t, err)
	cbLog := <-p.dispatcher.newLogRecvChan
	assert.True(t, cbLog.isAdmissionStatsReady(), "non-nil callback must publish admission readiness")
	assert.Equal(t, uint64(1), *p.collectMeasurements().CallbacksPending,
		"pending must be 1 before dispatcher consumes the callback log")
}

// TestAggregateGauges_LiveAndDrain verifies AggregateBatches/AggregateGroups
// reflect live aggregate state and return to zero after drain/close.
func TestAggregateGauges_LiveAndDrain(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1000
		c.LingerTime = 10 * time.Second
	})

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.AggregateBatches != nil && *s.AggregateBatches == 1
	}, "one batch should be aggregating")

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(1), *stats.AggregateBatches, "one batch should be aggregating")
	assert.Equal(t, uint64(1), *stats.AggregateGroups, "one group should be aggregating")

	p.Close()
	stats = p.collectMeasurements()
	assert.Equal(t, uint64(0), *stats.AggregateBatches, "aggregate batches should return to 0 after close")
	assert.Equal(t, uint64(0), *stats.AggregateGroups, "aggregate groups should return to 0 after close")
}

// blockingCallback blocks until released, used to observe in-flight state.
type blockingCallback struct {
	started chan struct{}
	done    chan struct{}
}

func newBlockingCallback() *blockingCallback {
	return &blockingCallback{started: make(chan struct{}), done: make(chan struct{})}
}

func (b *blockingCallback) Success(result *Result) {
	close(b.started)
	<-b.done
}
func (b *blockingCallback) Fail(result *Result) {
	close(b.started)
	<-b.done
}

// TestCallbacks_PendingAndInFlight verifies exact pending→inflight→done
// transitions.
func TestCallbacks_PendingAndInFlight(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
	})

	cb := newBlockingCallback()
	err := p.SendLog("", "topic", "src", "file", makeLog(), cb)
	assert.NoError(t, err)

	select {
	case <-cb.started:
	case <-time.After(5 * time.Second):
		t.Fatal("callback never started")
	}

	stats := p.collectMeasurements()
	assert.Equal(t, uint32(1), *stats.CallbacksInFlight, "exactly one callback in flight")
	assert.Equal(t, uint64(0), *stats.CallbacksPending, "pending should be 0 while callback runs")

	close(cb.done)
	waitFor(t, 5*time.Second, func() bool {
		return *p.collectMeasurements().CallbacksInFlight == 0
	}, "in flight should return to 0 after callback completes")
}

// statsCheckCallback calls GetStats during execution and asserts pending=0,
// inflight=1. This proves the ingress-boundary race is resolved: the callback
// can never observe a negative (MaxUint64) pending value.
type statsCheckCallback struct {
	t       *testing.T
	p       *producer
	invoked chan struct{}
}

func (c *statsCheckCallback) Success(result *Result) {
	s := c.p.collectMeasurements()
	assert.Equal(c.t, uint64(0), *s.CallbacksPending, "pending must be 0 while callback runs")
	assert.Equal(c.t, uint32(1), *s.CallbacksInFlight, "inflight must be 1 while callback runs")
	close(c.invoked)
}
func (c *statsCheckCallback) Fail(result *Result) {
	s := c.p.collectMeasurements()
	assert.Equal(c.t, uint64(0), *s.CallbacksPending, "pending must be 0 while callback runs")
	assert.Equal(c.t, uint32(1), *s.CallbacksInFlight, "inflight must be 1 while callback runs")
	close(c.invoked)
}

// TestCallbackCallsGetStats verifies that during callback execution,
// CallbacksPending=0 and CallbacksInFlight=1 (no underflow from the race).
func TestCallbackCallsGetStats(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
	})

	cb := &statsCheckCallback{t: t, p: p, invoked: make(chan struct{})}
	err := p.SendLog("", "topic", "src", "file", makeLog(), cb)
	assert.NoError(t, err)

	select {
	case <-cb.invoked:
	case <-time.After(5 * time.Second):
		t.Fatal("callback never invoked")
	}
}

// TestSendCallsInFlight_OnlyPutLogs verifies SendCallsInFlight wraps only
// the PutLogs call, not callback handling.
func TestSendCallsInFlight_OnlyPutLogs(t *testing.T) {
	mc := newBlockingMockClient()
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
	})

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		return atomic.LoadInt64(&mc.putLogsCount) > 0
	}, "PutLogs should be called")

	stats := p.collectMeasurements()
	assert.GreaterOrEqual(t, *stats.SendCallsInFlight, uint32(1), "send call should be in flight during PutLogs")
	assert.Equal(t, uint32(0), *stats.CallbacksInFlight, "callbacks should not be in flight during PutLogs")

	mc.releasePutLogs()
	waitFor(t, 5*time.Second, func() bool {
		return *p.collectMeasurements().SendCallsInFlight == 0
	}, "send calls in flight should return to 0 after PutLogs returns")
}

// TestSendCallDuration_Nonzero verifies send_call_duration measures the
// actual PutLogs call duration (not near-zero) using a controlled client delay.
func TestSendCallDuration_Nonzero(t *testing.T) {
	const delay = 50 * time.Millisecond
	mc := &delayMockClient{delay: delay}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
		c.MetricsLevel = MetricsLevelDetailed
	})

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.SendCallDurationMs != nil && s.SendCallDurationMs.Count > 0
	}, "send_call_duration should have at least one observation")

	stats := p.collectMeasurements()
	assert.NotNil(t, stats.SendCallDurationMs)
	assert.Greater(t, stats.SendCallDurationMs.Count, uint64(0), "should have observations")
	// SumMs should cover the controlled delay (allow for timing jitter: use half).
	assert.Greater(t, stats.SendCallDurationMs.SumMs, float64(delay.Milliseconds())/2.0,
		"send_call_duration sum should cover the client delay")
}

// TestLifecycle_ClosingThenClosed verifies Closing becomes true once close
// begins; Closed becomes true only after background work is actually complete.
// Closing snapshots are weak (sequence increments); closed snapshots are final.
func TestLifecycle_ClosingThenClosed(t *testing.T) {
	mc := newBlockingMockClient()
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
	})

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		return atomic.LoadInt64(&mc.putLogsCount) > 0
	}, "PutLogs should be called")

	assert.Equal(t, stateRunning, p.collectMeasurements().State)

	closeDone := make(chan struct{})
	go func() {
		p.Close()
		close(closeDone)
	}()

	waitFor(t, 5*time.Second, func() bool {
		return atomic.LoadInt32(&p.closing) != 0
	}, "close should begin")

	stats := p.collectMeasurements()
	assert.Equal(t, stateClosing, stats.State, "state should be closing while background work active")
	assert.Equal(t, consistencyWeak, stats.Consistency, "closing snapshots must be weak")

	seq1 := p.collectMeasurements().SnapshotSequence
	seq2 := p.collectMeasurements().SnapshotSequence
	assert.Greater(t, seq2, seq1, "snapshot sequence must still increment during closing")

	mc.releasePutLogs()
	select {
	case <-closeDone:
	case <-time.After(5 * time.Second):
		t.Fatal("Close did not complete after PutLogs released")
	}

	stats = p.collectMeasurements()
	assert.Equal(t, stateClosed, stats.State, "state should be closed after Close completes")
	assert.Equal(t, consistencyFinal, stats.Consistency, "closed snapshots must be final")

	seq1 = p.collectMeasurements().SnapshotSequence
	seq2 = p.collectMeasurements().SnapshotSequence
	assert.Equal(t, seq1, seq2, "snapshot sequence must be frozen after true closure")
}

// TestForceClose_BatchConservation verifies ForceClose counts dropped
// batches in terminalDroppedBatches so batch conservation holds.
func TestForceClose_BatchConservation(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1000
		c.LingerTime = 10 * time.Second
	})

	for i, src := range []string{"src1", "src2", "src3"} {
		err := p.SendLog("", "topic", src, "file", makeLog(), nil)
		assert.NoError(t, err, "SendLog %d should succeed", i)
	}

	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.AggregateBatches != nil && *s.AggregateBatches == 3
	}, "3 batches should be aggregating")

	p.ForceClose()

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(3), *stats.TerminalDroppedLogs, "3 dropped logs should be counted")
	assert.Equal(t, uint64(0), *stats.AggregateBatches, "aggregate should be empty after ForceClose")
}

// TestForceClose_TaskChanConsumed verifies that a batch already consumed
// from taskChan (blocked acquiring a maxSender token) is still counted as a
// terminal dropped batch and its pending callback is unwound on ForceClose.
func TestForceClose_TaskChanConsumed(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.MaxSenderCount = 1
		c.LingerTime = 10 * time.Second
	})

	token := <-p.threadPool.sender.maxSender

	cb := newBlockingCallback()
	err := p.SendLog("", "topic", "src", "file", makeLog(), cb)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		return p.threadPool.sendQueueLen() == 0
	}, "batch should be consumed from taskChan")

	time.Sleep(50 * time.Millisecond)

	p.ForceClose()

	p.threadPool.sender.maxSender <- token

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(1), *stats.TerminalDroppedLogs,
		"log consumed from taskChan must be counted as dropped")
	assert.Equal(t, uint64(0), *stats.CallbacksPending,
		"pending callback must be unwound on ForceClose drop")
}

// TestForceClose_HandleBatchForceQuit verifies that a batch consumed from
// taskChan and blocked in handleBatch (acquiring a maxSender token) is dropped
// via the centralized helper on ForceClose: terminalDroppedBatches incremented
// and pending callbacks unwound.
func TestForceClose_HandleBatchForceQuit(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.MaxSenderCount = 1
		c.LingerTime = 10 * time.Second
	})

	// Withhold the only sender token so handleBatch blocks after consuming the
	// batch from taskChan.
	token := <-p.threadPool.sender.maxSender

	cb := newBlockingCallback()
	err := p.SendLog("", "topic", "src", "file", makeLog(), cb)
	assert.NoError(t, err)

	// Wait until the batch has left the ingress queue, the aggregate map,
	// and taskChan — meaning it is now blocked in handleBatch waiting for the
	// token.
	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.AggregateBatches != nil && *s.AggregateBatches == 0 &&
			p.threadPool.sendQueueLen() == 0 &&
			len(p.dispatcher.newLogRecvChan) == 0
	}, "batch should be in handleBatch (ingress/map/taskChan all empty)")

	p.ForceClose()

	p.threadPool.sender.maxSender <- token

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(1), *stats.TerminalDroppedLogs,
		"log dropped via handleBatch force-quit must be counted as terminal dropped")
	assert.Equal(t, uint64(0), *stats.CallbacksPending,
		"pending callback must be unwound on force-quit drop")
}

// TestCheckBatches_RetryForceQuit deterministically exercises the
// checkBatches path where getRetryBatch(false) returns due retry batches but
// sendToThreadPool fails because forceQuitCh is closed. Every extracted retry
// batch must be routed through the centralized drop helper: terminal log/group
// outcomes increment exactly once, callbacksPending unwinds to zero, and
// resources are released. Uses a direct same-package setup with no background loops.
func TestCheckBatches_RetryForceQuit(t *testing.T) {
	mc := &mockProducerClient{}
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://127.0.0.1:0"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.MetricsLevel = MetricsLevelCore
	cfg.MaxBatchCount = 1000
	cfg.LingerTime = 10 * time.Second
	// Do NOT Start: no dispatcher/checkBatchTime/threadPool goroutines run.
	// Since no goroutines are started, no ForceClose cleanup is needed.
	p := newTestProducerWithClient(cfg, mc)

	// Replace taskChan with an unbuffered channel so sendToThreadPool's send
	// always blocks (no receiver); the select then reliably takes the closed
	// forceQuitCh case and returns false.
	p.threadPool.taskChan = make(chan *Batch)

	// Build two due retry batches, each carrying a callback.
	cb1 := newBlockingCallback()
	cb2 := newBlockingCallback()
	bl1 := &BatchLog{Key: BatchKey{Topic: "topic", Source: "src1", CallBackFun: cb1}, Log: makeLog()}
	bl2 := &BatchLog{Key: BatchKey{Topic: "topic", Source: "src2", CallBackFun: cb2}, Log: makeLog()}
	bl1.setMetricsTracked(p.collectionGenerationNow())
	bl2.setMetricsTracked(p.collectionGenerationNow())
	batch1 := initProducerBatch(bl1, cfg)
	batch2 := initProducerBatch(bl2, cfg)
	// Mark both as due for retry (nextRetryMs in the past).
	batch1.nextRetryMs = 0
	batch2.nextRetryMs = 0

	// Simulate the pending callbacks that would have been counted at ingress.
	p.stats.incCallbacksPending()
	p.stats.incCallbacksPending()
	assert.Equal(t, uint64(2), *p.collectMeasurements().CallbacksPending, "two callbacks should be pending before checkBatches")

	// Add both batches to the retry queue directly.
	p.dispatcher.retryQueue.addToRetryQueue(batch1, p.dispatcher.logger)
	p.dispatcher.retryQueue.addToRetryQueue(batch2, p.dispatcher.logger)

	// Close forceQuitCh so sendToThreadPool fails for every retry batch.
	close(p.dispatcher.forceQuitCh)

	// Invoke checkBatches deterministically (no background loop needed).
	p.dispatcher.checkBatches(cfg)

	stats := p.collectMeasurements()
	// Both retry batches must transfer their logs to terminal dropped exactly once.
	assert.Equal(t, uint64(2), *stats.TerminalDroppedLogs,
		"both extracted retry batches must be counted as terminal dropped")
	assert.Equal(t, uint64(0), *stats.CallbacksPending,
		"all pending callbacks must be unwound after retry force-quit")
}

// TestSendCallsInFlight_PanicSafe verifies SendCallsInFlight decrements if
// PutLogs panics.
func TestSendCallsInFlight_PanicSafe(t *testing.T) {
	mc := &panicMockClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
	})

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		return atomic.LoadInt64(&mc.putLogsCount) > 0
	}, "PutLogs should be called (and panic)")

	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.SendCallsInFlight != nil && *s.SendCallsInFlight == 0
	}, "SendCallsInFlight must be 0 after PutLogs panic")
}

// TestPutLogsPanic_Conservation verifies that after a PutLogs panic,
// accepted logs and groups are all terminal and no callback remains pending.
func TestPutLogsPanic_Conservation(t *testing.T) {
	mc := &panicMockClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
	})

	cb := newBlockingCallback()
	err := p.SendLog("", "topic", "src", "file", makeLog(), cb)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		return atomic.LoadInt64(&mc.putLogsCount) > 0
	}, "PutLogs should be called (and panic)")

	// Wait for the panic to be recovered and accounting recorded.
	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.TerminalDroppedLogs != nil && *s.TerminalDroppedLogs >= 1
	}, "terminal dropped log should be recorded after PutLogs panic")

	stats := p.collectMeasurements()
	// Conservation (logs): accepted = terminal(success+failure+dropped).
	accepted := *stats.AcceptedLogs
	terminal := *stats.TerminalSuccessLogs + *stats.TerminalFailureLogs + *stats.TerminalDroppedLogs
	assert.Equal(t, accepted, terminal, "accepted logs must equal terminal logs after panic")
	// Conservation (groups): accepted = terminal(success+failure+dropped).
	acceptedGroups := *stats.AcceptedGroups
	terminalGroups := *stats.TerminalSuccessGroups + *stats.TerminalFailureGroups + *stats.TerminalDroppedGroups
	assert.Equal(t, acceptedGroups, terminalGroups, "accepted groups must equal terminal groups after panic")

	assert.Equal(t, uint64(0), *stats.CallbacksPending, "pending callbacks must be 0 after panic")
	assert.Equal(t, uint64(1), *stats.TerminalDroppedGroups, "one group should be terminal dropped")
	assert.Equal(t, uint64(1), *stats.SendCallsFailure, "one failed send call should be recorded")
}

// panicCallbackWithCounter panics on Success and records invocation count.
type panicCallbackWithCounter struct {
	invoked int32
}

func (c *panicCallbackWithCounter) Success(result *Result) {
	atomic.AddInt32(&c.invoked, 1)
	panic("callback success panic")
}
func (c *panicCallbackWithCounter) Fail(result *Result) {
	atomic.AddInt32(&c.invoked, 1)
	panic("callback fail panic")
}

// TestCallbacks_PanicSafe verifies that if a callback panics, the in-flight
// gauge unwinds, remaining pending callbacks do not stay stuck, and subsequent
// callbacks in the same batch are not invoked (existing behavior preserved).
func TestCallbacks_PanicSafe(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 2 // both logs fit in one batch
		c.LingerTime = 10 * time.Second
	})

	cb1 := &panicCallbackWithCounter{}
	cb2 := &panicCallbackWithCounter{}
	// Same key → same batch → both callbacks in callBackList.
	err := p.SendLog("", "topic", "src", "file", makeLog(), cb1)
	assert.NoError(t, err)
	err = p.SendLog("", "topic", "src", "file", makeLog(), cb2)
	assert.NoError(t, err)

	// Wait for the first callback to be invoked (and panic).
	waitFor(t, 5*time.Second, func() bool {
		return atomic.LoadInt32(&cb1.invoked) > 0
	}, "first callback should be invoked")

	// The second callback must never be invoked (panic stops the loop).
	assert.Equal(t, int32(0), atomic.LoadInt32(&cb2.invoked),
		"second callback must not be invoked after first panics")

	// After the panic is recovered, both gauges must be unwound.
	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.CallbacksInFlight != nil && *s.CallbacksInFlight == 0 &&
			s.CallbacksPending != nil && *s.CallbacksPending == 0
	}, "callbacks in flight and pending must both be 0 after callback panic")
}

// TestClosedSnapshot_FinalAndImmutable verifies that after stateClosed, the
// complete final measurement value remains identical across multiple reads.
func TestClosedSnapshot_FinalAndImmutable(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Millisecond
	})

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.SendCallsSuccess != nil && *s.SendCallsSuccess >= 1
	}, "log should be sent successfully")

	p.Close()

	s1 := p.collectMeasurements()
	if err := p.SendLog("", "topic", "src", "file", makeLog(), nil); err == nil {
		t.Fatal("post-close SendLog must fail")
	}
	// Sleep across multiple wall-clock ticks so any unfrozen time field would change.
	time.Sleep(100 * time.Millisecond)
	s2 := p.collectMeasurements()
	time.Sleep(100 * time.Millisecond)
	s3 := p.collectMeasurements()

	// The entire snapshot must be byte-for-byte equal across calls.
	assert.Equal(t, s1, s2, "complete closed snapshot must be immutable (s1 vs s2)")
	assert.Equal(t, s1, s3, "complete closed snapshot must be immutable (s1 vs s3)")
	assert.Equal(t, stateClosed, s1.State)
	assert.Equal(t, consistencyFinal, s1.Consistency)
}

func TestForceClosedSnapshotRemainsImmutableAfterPostCloseSend(t *testing.T) {
	p := newStatsTestProducer(t, &mockProducerClient{})
	p.ForceClose()

	s1 := p.collectMeasurements()
	if err := p.SendLog("", "topic", "src", "file", makeLog(), nil); err == nil {
		t.Fatal("post-force-close SendLog must fail")
	}
	s2 := p.collectMeasurements()
	assert.Equal(t, s1, s2, "post-force-close rejection must not mutate the final snapshot")
	assert.Equal(t, stateClosed, s2.State)
	assert.Equal(t, consistencyFinal, s2.Consistency)
}

// TestConcurrentClose_GetStatsFreeze verifies that concurrent GetStats calls
// during/after Close observe a monotonic sequence that freezes exactly once,
// and every observed sequence is <= the final frozen sequence.
func TestConcurrentClose_GetStatsFreeze(t *testing.T) {
	mc := newBlockingMockClient()
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchCount = 1
		c.LingerTime = 10 * time.Second
	})

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		return atomic.LoadInt64(&mc.putLogsCount) > 0
	}, "PutLogs should be called")

	var wg sync.WaitGroup
	stop := make(chan struct{})
	var seqs []uint64
	var seqsMu sync.Mutex

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
				}
				seq := p.collectMeasurements().SnapshotSequence
				seqsMu.Lock()
				seqs = append(seqs, seq)
				seqsMu.Unlock()
			}
		}()
	}

	mc.releasePutLogs()
	p.Close()

	time.Sleep(50 * time.Millisecond)
	close(stop)
	wg.Wait()

	finalSeq := p.collectMeasurements().SnapshotSequence
	// Every observed sequence must be <= the final frozen sequence (monotonic).
	for _, seq := range seqs {
		assert.LessOrEqual(t, seq, finalSeq,
			"observed sequence %d must be <= final sequence %d", seq, finalSeq)
	}
	// After close, the sequence must never change.
	time.Sleep(50 * time.Millisecond)
	assert.Equal(t, finalSeq, p.collectMeasurements().SnapshotSequence,
		"sequence must be frozen after close")
	assert.Equal(t, stateClosed, p.collectMeasurements().State)
}

// TestMetricsOff_NoNewAtomics verifies the Off fast path has no stats
// instrumentation (GetStats still works, fields nil).
func TestMetricsOff_NoNewAtomics(t *testing.T) {
	mc := &mockProducerClient{}
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://127.0.0.1:0"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.MetricsLevel = MetricsLevelOff
	cfg.MaxBatchCount = 1
	cfg.LingerTime = 10 * time.Millisecond
	p := newTestProducerWithClient(cfg, mc)
	p.Start()
	t.Cleanup(func() { p.ForceClose() })

	err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		return p.collectMeasurements().State == stateRunning
	}, "producer should be running")

	stats := p.collectMeasurements()
	assert.True(t, stats.MetricsDisabled)
	assert.Nil(t, stats.IngressQueue, "IngressQueue must be nil when metrics off")
	assert.Nil(t, stats.AggregateBatches, "AggregateBatches must be nil when metrics off")
	assert.Nil(t, stats.CallbacksPending, "CallbacksPending must be nil when metrics off")
	assert.Nil(t, stats.SendCallsInFlight, "SendCallsInFlight must be nil when metrics off")
}

// Compile-time interface checks.
var _ CallBack = (*blockingCallback)(nil)
var _ CallBack = (*statsCheckCallback)(nil)
var _ CallBack = (*panicCallbackWithCounter)(nil)

// TestBatchRollover_Conservation verifies that when tryAddLog fails on a
// full batch (batch rollover), the log remains accepted (counted at ingress)
// and is eventually terminal — accepted is never lost from the public count.
// Two logs with the same key that together exceed config.MaxBatchSize trigger
// the rollover: the first batch is sent (sealed by size) and a new batch
// accepts the second log. dropBatchLogStats is not exercised here because the
// tryAddLog size check uses the package MaxBatchSize constant (9.5MB); a log
// large enough to fail that check is rejected at ingress.
func TestBatchRollover_Conservation(t *testing.T) {
	mc := &mockProducerClient{}
	p := newStatsTestProducer(t, mc, func(c *Config) {
		c.MaxBatchSize = 200 // small enough that two ~150-byte logs overflow
		c.MaxBatchCount = 1000
		c.LingerTime = 10 * time.Second
	})

	log1 := &pb.Log{Contents: []*pb.LogContent{{Key: "k", Value: string(make([]byte, 150))}}}
	log2 := &pb.Log{Contents: []*pb.LogContent{{Key: "k", Value: string(make([]byte, 150))}}}
	err := p.SendLog("", "topic", "src", "file", log1, nil)
	assert.NoError(t, err)
	err = p.SendLog("", "topic", "src", "file", log2, nil)
	assert.NoError(t, err)

	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		return s.AcceptedLogs != nil && *s.AcceptedLogs >= 2
	}, "both logs should be accepted")

	waitFor(t, 5*time.Second, func() bool {
		s := p.collectMeasurements()
		if s.AcceptedLogs == nil || *s.AcceptedLogs < 2 ||
			s.TerminalSuccessLogs == nil || s.TerminalFailureLogs == nil || s.TerminalDroppedLogs == nil {
			return false
		}
		terminal := *s.TerminalSuccessLogs + *s.TerminalFailureLogs + *s.TerminalDroppedLogs
		return terminal >= 2
	}, "both logs should reach terminal state")

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(2), *stats.AcceptedLogs, "both logs must remain accepted")
	terminal := *stats.TerminalSuccessLogs + *stats.TerminalFailureLogs + *stats.TerminalDroppedLogs
	assert.Equal(t, uint64(2), terminal, "both logs must be terminal")
}

// TestForceClose_ChannelResidual_Conservation verifies that logs still in
// newLogRecvChan when ForceClose runs are accepted (counted at ingress) then
// explicitly terminal-dropped during drain, so conservation holds without
// relying on the close catch-all.
func TestForceClose_ChannelResidual_Conservation(t *testing.T) {
	mc := &mockProducerClient{}
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://127.0.0.1:0"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.MetricsLevel = MetricsLevelCore
	cfg.MaxBatchCount = 1000
	cfg.LingerTime = 10 * time.Second
	cfg.BatchQueueSize = 100
	// Do NOT Start: dispatcher goroutine not running, so logs stay in newLogRecvChan.
	p := newTestProducerWithClient(cfg, mc)

	// Send 3 logs; they sit in newLogRecvChan unprocessed.
	for i := 0; i < 3; i++ {
		err := p.SendLog("", "topic", "src", "file", makeLog(), nil)
		assert.NoError(t, err, "SendLog %d should succeed", i)
	}

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(3), *stats.AcceptedLogs, "3 logs should be accepted at ingress")
	assert.Equal(t, uint64(3), stats.IngressQueue.CurrentItems, "3 logs should be in the channel")

	// ForceClose drains newLogRecvChan; each residual log must be terminal-dropped.
	p.ForceClose()

	stats = p.collectMeasurements()
	assert.Equal(t, uint64(3), *stats.AcceptedLogs, "accepted count must not change")
	assert.Equal(t, uint64(3), *stats.TerminalDroppedLogs, "3 residual logs must be terminal dropped")
	terminal := *stats.TerminalSuccessLogs + *stats.TerminalFailureLogs + *stats.TerminalDroppedLogs
	assert.Equal(t, *stats.AcceptedLogs, terminal, "all accepted residual logs must be terminal after drain")
	assert.Equal(t, uint64(0), *stats.TerminalSuccessLogs, "no logs should be terminal success")
}

// TestSetAdmissionStatsReady_Idempotent verifies that publishing readiness
// twice sequentially does not corrupt the BatchLog state word: the ready bit
// is set exactly once and the size bits are preserved (no overflow into
// adjacent bits from a non-idempotent Add).
func TestSetAdmissionStatsReady_Idempotent(t *testing.T) {
	batchLog := &BatchLog{
		Key: BatchKey{CallBackFun: &blockingCallback{}},
		Log: makeLog(),
	}
	batchLog.setLogSize(42)

	batchLog.setAdmissionStatsReady()
	batchLog.setAdmissionStatsReady()

	assert.True(t, batchLog.isAdmissionStatsReady(), "ready bit must be set")
	assert.Equal(t, 42, batchLog.cachedLogSize(), "size must be preserved after double publish")
	assert.False(t, batchLog.hasCircuitPermit(), "circuit permit bit must be unaffected")
}

// TestSetAdmissionStatsReady_Concurrent verifies concurrent publication
// calls from multiple goroutines do not corrupt the BatchLog state word: the
// ready bit ends up set exactly once and size bits are preserved regardless of
// how many goroutines race to set it.
func TestSetAdmissionStatsReady_Concurrent(t *testing.T) {
	batchLog := &BatchLog{
		Key: BatchKey{CallBackFun: &blockingCallback{}},
		Log: makeLog(),
	}
	batchLog.setLogSize(99)

	// Hammer readiness publication from 8 goroutines concurrently.
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			batchLog.setAdmissionStatsReady()
		}()
	}
	wg.Wait()

	assert.True(t, batchLog.isAdmissionStatsReady(), "ready bit must be set after concurrent calls")
	assert.Equal(t, 99, batchLog.cachedLogSize(), "size must be preserved after concurrent publish")
	assert.False(t, batchLog.hasCircuitPermit(), "circuit permit bit must be unaffected")
}
