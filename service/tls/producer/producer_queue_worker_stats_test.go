package producer

import (
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

var testLogger = log.NewNopLogger()

// --- Fix A: WorkerRoles.Running must reflect actual alive goroutines ---

// TestWorkerRoles_NotStarted verifies Running=0 before Start().
func TestWorkerRoles_NotStarted(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	stats := p.collectMeasurements()

	// Before start, no worker goroutines are alive
	if stats.WorkerRoles != nil {
		for _, wr := range stats.WorkerRoles {
			assert.Equal(t, uint32(0), wr.Running, "%s running should be 0 before start", wr.Role)
		}
	}
}

// TestWorkerRoles_Running verifies dispatcher Running=2 and sender Running reflects
// actual dynamic sender goroutines after Start().
func TestWorkerRoles_Running(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()
	defer p.Close()

	// Give goroutines time to start and increment counters
	time.Sleep(50 * time.Millisecond)

	stats := p.collectMeasurements()
	assert.NotNil(t, stats.WorkerRoles)

	var dispatcherWR, senderWR *workerRoleMeasurements
	for i := range stats.WorkerRoles {
		switch stats.WorkerRoles[i].Role {
		case roleDispatcher:
			dispatcherWR = &stats.WorkerRoles[i]
		case roleSender:
			senderWR = &stats.WorkerRoles[i]
		}
	}

	assert.NotNil(t, dispatcherWR, "dispatcher role should be present")
	assert.Equal(t, uint32(2), dispatcherWR.Expected)
	assert.Equal(t, uint32(2), dispatcherWR.Running, "dispatcher running should be 2 (run + checkBatchTime)")

	assert.NotNil(t, senderWR, "sender role should be present")
	assert.Equal(t, uint32(p.config.MaxSenderCount), senderWR.Expected)
	// sender running counts dynamic goroutines; with no batches sent, should be 0
	assert.Equal(t, uint32(0), senderWR.Running, "sender running should be 0 with no in-flight sends")
}

// TestWorkerRoles_AfterClose verifies Running=0 after graceful Close().
func TestWorkerRoles_AfterClose(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()
	time.Sleep(50 * time.Millisecond)
	p.Close()

	stats := p.collectMeasurements()
	if stats.WorkerRoles != nil {
		for _, wr := range stats.WorkerRoles {
			assert.Equal(t, uint32(0), wr.Running, "%s running should be 0 after close", wr.Role)
		}
	}
}

// TestWorkerRoles_AfterForceClose verifies Running=0 after ForceClose().
func TestWorkerRoles_AfterForceClose(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()
	time.Sleep(50 * time.Millisecond)
	p.ForceClose()

	stats := p.collectMeasurements()
	if stats.WorkerRoles != nil {
		for _, wr := range stats.WorkerRoles {
			assert.Equal(t, uint32(0), wr.Running, "%s running should be 0 after force close", wr.Role)
		}
	}
}

// TestWorkerRoles_NoStaleExpectedAfterTermination proves Running is not hardcoded to Expected.
func TestWorkerRoles_NoStaleExpectedAfterTermination(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()
	time.Sleep(50 * time.Millisecond)

	// While running, dispatcher running should equal expected (2)
	s1 := p.collectMeasurements()
	for _, wr := range s1.WorkerRoles {
		if wr.Role == roleDispatcher {
			assert.Equal(t, wr.Expected, wr.Running, "dispatcher running should equal expected while running")
		}
	}

	p.Close()

	// After close, running must be 0, NOT expected
	s2 := p.collectMeasurements()
	for _, wr := range s2.WorkerRoles {
		if wr.Role == roleDispatcher {
			assert.Equal(t, uint32(2), wr.Expected, "expected should remain configured count")
			assert.Equal(t, uint32(0), wr.Running, "running must be 0 after close, not expected")
		}
		if wr.Role == roleSender {
			assert.Equal(t, uint32(p.config.MaxSenderCount), wr.Expected)
			assert.Equal(t, uint32(0), wr.Running, "sender running must be 0 after close")
		}
	}
}

// --- Fix B: Retry queue exactness/conservation ---

// TestRetryQueue_DepthExact verifies physical depth matches enqueue/dequeue count.
func TestRetryQueue_DepthExact(t *testing.T) {
	q := newRetryQueue()
	q.setMetricsEnabled(true)

	// Initially empty
	d, _ := q.snapshot()
	assert.Equal(t, 0, d)

	// Enqueue 3 batches
	for i := 0; i < 3; i++ {
		b := &Batch{nextRetryMs: 0}
		q.addToRetryQueue(b, testLogger)
	}
	d, _ = q.snapshot()
	assert.Equal(t, 3, d, "depth should be 3 after 3 enqueues")

	// Dequeue all (stopFlag=true drains everything)
	batches := q.getRetryBatch(true)
	assert.Equal(t, 3, len(batches))

	d, _ = q.snapshot()
	assert.Equal(t, 0, d, "depth should be 0 after draining all")
}

// TestRetryQueue_HighWatermark verifies HWM tracks peak depth.
func TestRetryQueue_HighWatermark(t *testing.T) {
	q := newRetryQueue()
	q.setMetricsEnabled(true)

	// Enqueue 5
	for i := 0; i < 5; i++ {
		q.addToRetryQueue(&Batch{nextRetryMs: 0}, testLogger)
	}
	_, hwm := q.snapshot()
	assert.Equal(t, int64(5), hwm, "HWM should be 5")

	// Dequeue 2
	q.getRetryBatch(true)
	// Drain remaining
	q.getRetryBatch(true)

	// HWM should still be 5 (lifetime, not reset by drain)
	_, hwm = q.snapshot()
	assert.Equal(t, int64(5), hwm, "HWM should remain 5 after drain")
}

// TestRetryQueue_HWMDisabledWhenOff verifies no HWM tracking when metrics disabled.
func TestRetryQueue_HWMDisabledWhenOff(t *testing.T) {
	q := newRetryQueue()
	q.setMetricsEnabled(false)

	for i := 0; i < 5; i++ {
		q.addToRetryQueue(&Batch{nextRetryMs: 0}, testLogger)
	}

	d, hwm := q.snapshot()
	assert.Equal(t, 0, d, "Off must not maintain an observer depth")
	assert.Equal(t, int64(0), hwm, "HWM should be 0 when metrics disabled")
}

// TestRetryQueue_NeverNegative verifies depth never goes negative.
func TestRetryQueue_NeverNegative(t *testing.T) {
	q := newRetryQueue()
	q.setMetricsEnabled(true)

	// Dequeue from empty queue — should not panic or go negative
	batches := q.getRetryBatch(true)
	assert.Equal(t, 0, len(batches))

	d, _ := q.snapshot()
	assert.Equal(t, 0, d, "depth should be 0, never negative")

	// Enqueue 1, then dequeue more than available
	q.addToRetryQueue(&Batch{nextRetryMs: 0}, testLogger)
	q.getRetryBatch(true) // drains the 1
	q.getRetryBatch(true) // drains nothing

	d, _ = q.snapshot()
	assert.Equal(t, 0, d, "depth should be 0, never negative")
}

// TestRetryQueue_Concurrent verifies concurrent enqueue/dequeue is race-free.
func TestRetryQueue_Concurrent(t *testing.T) {
	q := newRetryQueue()
	q.setMetricsEnabled(true)

	done := make(chan struct{})
	// Concurrent enqueuers
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				q.addToRetryQueue(&Batch{nextRetryMs: 0}, testLogger)
			}
			done <- struct{}{}
		}()
	}
	// Concurrent dequeuers
	for i := 0; i < 3; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				q.getRetryBatch(true)
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < 8; i++ {
		<-done
	}

	// Final depth should be consistent: 5*100 - 3*100 = 200 (minus any in-flight)
	d, _ := q.snapshot()
	assert.GreaterOrEqual(t, d, 0, "depth must never be negative")
}

// --- Fix C: Send queue semantics ---

// TestSendQueue_CurrentAndCapacity verifies current_items=len(taskChan), capacity=cap(taskChan).
func TestSendQueue_CurrentAndCapacity(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)

	// Before start, taskChan is empty
	stats := p.collectMeasurements()
	assert.NotNil(t, stats.SendQueue)
	assert.Equal(t, uint64(0), stats.SendQueue.CurrentItems)
	assert.Equal(t, uint64(cap(p.threadPool.taskChan)), stats.SendQueue.CapacityItems)
	assert.True(t, stats.SendQueue.Bounded)
}

// TestSendQueue_HighWatermark verifies HWM is observed after enqueue, not at read time.
func TestSendQueue_HighWatermark(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)

	// Use the real enqueue path (sendToThreadPool) which observes HWM after send
	for i := 0; i < 5; i++ {
		ok := p.dispatcher.sendToThreadPool(&Batch{})
		assert.True(t, ok, "sendToThreadPool should succeed")
	}

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(5), stats.SendQueue.CurrentItems)
	assert.Equal(t, uint64(5), stats.SendQueue.HighWatermarkItems, "HWM should reflect peak of 5")

	// Drain the channel
	for i := 0; i < 5; i++ {
		<-p.threadPool.taskChan
	}

	stats = p.collectMeasurements()
	assert.Equal(t, uint64(0), stats.SendQueue.CurrentItems)
	assert.Equal(t, uint64(5), stats.SendQueue.HighWatermarkItems, "HWM should persist after drain")
}

// TestSendQueue_NoAliasingWithInFlight verifies send_queue depth is distinct from send_calls_in_flight.
func TestSendQueue_NoAliasingWithInFlight(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)

	// Enqueue 3 batches to taskChan (these are waiting, not in-flight)
	for i := 0; i < 3; i++ {
		p.threadPool.taskChan <- &Batch{}
	}

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(3), stats.SendQueue.CurrentItems, "send queue should have 3 waiting")
	if stats.SendCallsInFlight != nil {
		assert.Equal(t, uint32(0), *stats.SendCallsInFlight, "no sends should be in-flight")
	}

	// Clean up
	for i := 0; i < 3; i++ {
		<-p.threadPool.taskChan
	}
}

// TestSendQueue_Drain verifies queue drains to zero after close.
func TestSendQueue_Drain(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()
	defer p.Close()

	// Send some logs to build up backlog, then close
	for i := 0; i < 3; i++ {
		_ = p.SendLog("hash", "topic", "source", "file", makeLog(), nil)
	}
	time.Sleep(100 * time.Millisecond)
	p.Close()

	stats := p.collectMeasurements()
	// After close, send queue should be drained
	assert.Equal(t, uint64(0), stats.SendQueue.CurrentItems, "send queue should be empty after close")
}

// TestSendQueue_OffNoHWMUpdate verifies metrics=off does not track HWM.
func TestSendQueue_OffNoHWMUpdate(t *testing.T) {
	p := newTestProducer()
	// metrics=off (default)

	for i := 0; i < 5; i++ {
		p.threadPool.taskChan <- &Batch{}
	}

	stats := p.collectMeasurements()
	// When off, SendQueue may be absent or have zero HWM
	if stats.SendQueue != nil {
		assert.Equal(t, uint64(0), stats.SendQueue.HighWatermarkItems, "HWM should be 0 when metrics off")
	}

	// Clean up
	for i := 0; i < 5; i++ {
		<-p.threadPool.taskChan
	}
}

// --- Public MetricsLevel config entry (public contract) ---

// TestMetricsLevelConfig_DefaultIsOff verifies zero value / unset config = Off.
func TestMetricsLevelConfig_DefaultIsOff(t *testing.T) {
	config := GetDefaultProducerConfig()
	// Zero value of MetricsLevel (unset in default config) must be Off
	assert.Equal(t, MetricsLevelOff, config.MetricsLevel, "default config MetricsLevel must be Off (zero value)")

	p := newProducer(config)
	assert.Equal(t, MetricsLevelOff, p.metricsLevelNow(), "producer must read MetricsLevel=Off from config")

	stats := p.collectMeasurements()
	assert.Equal(t, MetricsLevelOff, stats.MetricsLevel)
	assert.True(t, stats.MetricsDisabled)
}

// TestMetricsLevelConfig_CoreEnablesCoreMetrics verifies public Config=Core
// enables core metrics via the formal public entry point.
func TestMetricsLevelConfig_CoreEnablesCoreMetrics(t *testing.T) {
	config := GetDefaultProducerConfig()
	config.MetricsLevel = MetricsLevelCore
	p := newProducer(config)

	assert.Equal(t, MetricsLevelCore, p.metricsLevelNow(), "producer must read MetricsLevel=Core from public config")

	stats := p.collectMeasurements()
	assert.Equal(t, MetricsLevelCore, stats.MetricsLevel)
	assert.False(t, stats.MetricsDisabled, "core metrics must not be disabled")
	// Core metrics should populate business counters
	assert.NotNil(t, stats.AcceptedLogs, "core metrics must populate accepted logs counter")
}

// TestMetricsLevelConfig_DetailedPassesThrough verifies Detailed is passed
// through the public config entry (if the enum value is defined).
func TestMetricsLevelConfig_DetailedPassesThrough(t *testing.T) {
	config := GetDefaultProducerConfig()
	config.MetricsLevel = MetricsLevelDetailed
	p := newProducer(config)

	assert.Equal(t, MetricsLevelDetailed, p.metricsLevelNow(), "producer must read MetricsLevel=Detailed from public config")

	stats := p.collectMeasurements()
	assert.Equal(t, MetricsLevelDetailed, stats.MetricsLevel)
	assert.False(t, stats.MetricsDisabled)
}

// TestMetricsLevelConfig_IsCopiedAtConstruction verifies that mutating the
// caller-owned Config is not an implicit runtime control path.
func TestMetricsLevelConfig_IsCopiedAtConstruction(t *testing.T) {
	config := GetDefaultProducerConfig()
	config.MetricsLevel = MetricsLevelCore
	p := newProducer(config)
	assert.Equal(t, MetricsLevelCore, p.metricsLevelNow())

	// Mutating the config after construction must not affect the producer
	config.MetricsLevel = MetricsLevelOff
	assert.Equal(t, MetricsLevelCore, p.metricsLevelNow(), "config mutation must not change the live level")

	stats := p.collectMeasurements()
	assert.Equal(t, MetricsLevelCore, stats.MetricsLevel)
}

// TestMetricsLevelConfig_ExistingCallersCompatible verifies that existing
// callers that do not set MetricsLevel continue to get Off behavior.
func TestMetricsLevelConfig_ExistingCallersCompatible(t *testing.T) {
	// Simulate an existing caller that uses GetDefaultProducerConfig without
	// touching MetricsLevel — must behave exactly as before (Off).
	p := newTestProducer()
	assert.Equal(t, MetricsLevelOff, p.metricsLevelNow())

	stats := p.collectMeasurements()
	assert.True(t, stats.MetricsDisabled)
	assert.Nil(t, stats.AcceptedLogs)
}

// makeLog is defined in producer_stats_test.go; ensure pb import is used
var _ = pb.Log{}
