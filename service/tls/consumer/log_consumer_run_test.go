package consumer

import (
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

// newRunTestConsumer builds a minimal logConsumer for exercising the
// runWithStatus worker-completion contract. It wires a real WaitGroup, status
// lock, nop logger, and (optionally) a stats state so the real
// releaseIfRetired/releaseHeldBatch paths run. The consume shard is set so
// fetch/consume paths that touch it do not nil-panic.
func newRunTestConsumer(t *testing.T, withStats bool) *logConsumer {
	t.Helper()
	lc := &logConsumer{
		wg:         &sync.WaitGroup{},
		logger:     log.NewNopLogger(),
		statusLock: &sync.RWMutex{},
		status:     readyToFetch,
		shard:      &tls.ConsumeShard{TopicID: "test-topic", ShardID: 0},
	}
	if withStats {
		lc.stats = newConsumerStatsState(MetricsLevelCore)
	}
	return lc
}

// waitWG waits for wg to reach zero, failing the test on timeout. Used to assert
// that runWithStatus always calls Done() (success/error/panic).
func waitWG(t *testing.T, wg *sync.WaitGroup) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("WaitGroup did not reach zero: runWithStatus failed to call Done()")
	}
}

// heldReady returns the current held/ready counter values from a stats state,
// read atomically (avoids needing a *consumer for snapshot()).
func heldReady(s *consumerStatsState) (int64, int64) {
	return atomic.LoadInt64(&s.heldBatchRefs), atomic.LoadInt64(&s.readyToProcess)
}

// TestRunWithStatus_SuccessDonesWGAndSetsStatus: a nil error from f must set
// status to ifDone and the worker must call Done() so workerWG reaches zero.
func TestRunWithStatus_SuccessDonesWGAndSetsStatus(t *testing.T) {
	lc := newRunTestConsumer(t, false)
	lc.wg.Add(1)
	go lc.runWithStatus(func() error { return nil }, readyToConsume, fetching, backoff, backoff)

	waitWG(t, lc.wg)
	if got := lc.loadStatus(); got != readyToConsume {
		t.Fatalf("status after success = %d, want %d (readyToConsume)", got, readyToConsume)
	}
}

// TestRunWithStatus_BusyErrorDonesWGAndSetsIfBusy: errBusy must set status to
// ifBusy and still call Done().
func TestRunWithStatus_BusyErrorDonesWGAndSetsIfBusy(t *testing.T) {
	lc := newRunTestConsumer(t, false)
	lc.wg.Add(1)
	go lc.runWithStatus(func() error { return errBusy }, readyToConsume, fetching, backoff, backoff)

	waitWG(t, lc.wg)
	if got := lc.loadStatus(); got != backoff {
		t.Fatalf("status after errBusy = %d, want %d (backoff/ifBusy)", got, backoff)
	}
}

// TestRunWithStatus_GenericErrorDonesWGAndSetsIfErr: a non-sentinel error must
// set status to ifErr and still call Done().
func TestRunWithStatus_GenericErrorDonesWGAndSetsIfErr(t *testing.T) {
	lc := newRunTestConsumer(t, false)
	lc.wg.Add(1)
	go lc.runWithStatus(func() error { return errors.New("boom") }, readyToConsume, fetching, backoff, backoff)

	waitWG(t, lc.wg)
	if got := lc.loadStatus(); got != fetching {
		t.Fatalf("status after generic error = %d, want %d (fetching/ifErr)", got, fetching)
	}
}

// TestRunWithStatus_PanicRecoveredDonesWGAndSetsIfPanic: a panic in f must be
// recovered, status set to ifPanic, and Done() called so the goroutine does not
// leak and workerWG reaches zero.
func TestRunWithStatus_PanicRecoveredDonesWGAndSetsIfPanic(t *testing.T) {
	lc := newRunTestConsumer(t, false)
	lc.wg.Add(1)
	go lc.runWithStatus(func() error { panic("kaboom") }, readyToConsume, fetching, backoff, backoff)

	waitWG(t, lc.wg)
	if got := lc.loadStatus(); got != backoff {
		t.Fatalf("status after panic = %d, want %d (backoff/ifPanic)", got, backoff)
	}
}

// TestRunWithStatus_NoStatsStillDonesWG: when stats is nil, releaseIfRetired is
// not registered, but Done() must still run and status must be set.
func TestRunWithStatus_NoStatsStillDonesWG(t *testing.T) {
	lc := newRunTestConsumer(t, false)
	if lc.stats != nil {
		t.Fatalf("precondition: stats must be nil")
	}
	lc.wg.Add(1)
	go lc.runWithStatus(func() error { return nil }, readyToConsume, fetching, backoff, backoff)

	waitWG(t, lc.wg)
	if got := lc.loadStatus(); got != readyToConsume {
		t.Fatalf("status = %d, want %d", got, readyToConsume)
	}
}

// TestRunWithStatus_RetiredWithStatsReleasesHeldBatch: when stats is active and
// the worker is retired while holding a batch, runWithStatus must call
// releaseIfRetired on exit, balancing the held/ready counters back to zero
// before Done() completes.
func TestRunWithStatus_RetiredWithStatsReleasesHeldBatch(t *testing.T) {
	lc := newRunTestConsumer(t, true)

	// Simulate fetchData holding a non-empty batch.
	lc.stats.incHeldBatchRefs()
	lc.stats.incReadyToProcess()
	atomic.StoreUint64(&lc.heldGeneration, lc.stats.currentCollectionGeneration())
	atomic.StoreInt32(&lc.held, 1)
	if h, r := heldReady(lc.stats); h != 1 || r != 1 {
		t.Fatalf("precondition: held=%d ready=%d, want 1/1", h, r)
	}

	// Worker is removed (rebalance/shutdown) mid-flight.
	lc.markRetired()

	lc.wg.Add(1)
	go lc.runWithStatus(func() error { return nil }, readyToFetch, fetching, backoff, backoff)

	waitWG(t, lc.wg)
	if h, r := heldReady(lc.stats); h != 0 || r != 0 {
		t.Fatalf("after retired exit: held=%d ready=%d, want 0/0 (releaseIfRetired did not run)", h, r)
	}
}

// TestRunWithStatus_RetiredNoStatsNoPanic: when stats is nil the worker can be
// marked retired; runWithStatus must still complete without panicking (no
// releaseIfRetired registration, no held batch to release).
func TestRunWithStatus_RetiredNoStatsNoPanic(t *testing.T) {
	lc := newRunTestConsumer(t, false)
	lc.markRetired()
	lc.wg.Add(1)
	go lc.runWithStatus(func() error { return nil }, readyToFetch, fetching, backoff, backoff)

	waitWG(t, lc.wg)
}

// TestRun_LaunchesWorkerAndCompletes verifies the run() state machine does
// Add(1) before launching runWithStatus directly (no wrapper) and that the
// worker completes so workerWG reaches zero. Uses the pending->init path with a
// mock client so init() succeeds (DescribeCheckPoint empty -> DescribeCursor
// returns a cursor).
func TestRun_LaunchesWorkerAndCompletes(t *testing.T) {
	lc := newRunTestConsumer(t, false)
	lc.client = &mockClient{}
	lc.conf = GetDefaultConsumerConfig()
	lc.conf.ProjectID = "p"
	lc.conf.ConsumerGroupName = "g"
	lc.status = pending

	lc.run()
	waitWG(t, lc.wg)
	if got := lc.loadStatus(); got != readyToFetch {
		t.Fatalf("status after init run = %d, want %d (readyToFetch)", got, readyToFetch)
	}
}

func backoffEventCount(lc *logConsumer) int64 {
	if lc.stats == nil {
		return 0
	}
	return atomic.LoadInt64(&lc.stats.backoffEvents)
}

func runAndWait(t *testing.T, lc *logConsumer) {
	t.Helper()
	lc.run()
	waitWG(t, lc.wg)
}

// TestRun_BackoffEventsCountStateTransitions verifies that one throttled fetch
// contributes one event, repeated backoff checks do not create more events,
// expiry is not counted as another event, and a later throttled fetch creates
// the next event.
func TestRun_BackoffEventsCountStateTransitions(t *testing.T) {
	mc := &mockClient{
		consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
			return nil, &tls.Error{HTTPCode: http.StatusTooManyRequests, Message: "throttled"}
		},
	}
	lc := newCounterTestConsumer(t, mc)

	// The first ConsumeLogs 429 transitions the worker into backoff and must
	// count immediately, without waiting for another 200ms scheduler tick.
	runAndWait(t, lc)
	if got := lc.loadStatus(); got != backoff {
		t.Errorf("after first 429: status=%d, want backoff=%d", got, backoff)
	}
	if got := backoffEventCount(lc); got != 1 {
		t.Errorf("after first 429: backoff events=%d, want 1", got)
	}

	// The backoff worker can run repeatedly while the 5s window is active, but
	// those same-state checks must not create additional events.
	for i := 0; i < 3; i++ {
		runAndWait(t, lc)
	}
	if got := backoffEventCount(lc); got != 1 {
		t.Errorf("after repeated backoff checks: backoff events=%d, want 1", got)
	}

	// Force the backoff window to expire. The worker exits to readyToFetch; that
	// state transition is not a new backoff event.
	lc.lastBackoffTime = time.Now().Add(-6 * time.Second)
	runAndWait(t, lc)
	if got := lc.loadStatus(); got != readyToFetch {
		t.Errorf("after expired backoff: status=%d, want readyToFetch=%d", got, readyToFetch)
	}
	if got := backoffEventCount(lc); got != 1 {
		t.Errorf("after expired backoff: backoff events=%d, want 1", got)
	}

	// The next 429 is a fresh transition into backoff and must count once.
	runAndWait(t, lc)
	if got := lc.loadStatus(); got != backoff {
		t.Errorf("after second 429: status=%d, want backoff=%d", got, backoff)
	}
	if got := backoffEventCount(lc); got != 2 {
		t.Errorf("after second 429: backoff events=%d, want 2", got)
	}
}

// TestRun_BackoffEventsStatsBoundaries verifies nil stats safety and that an
// existing backoff state is not backfilled when metrics are re-enabled.
func TestRun_BackoffEventsStatsBoundaries(t *testing.T) {
	t.Run("nil stats", func(t *testing.T) {
		lc := newRunTestConsumer(t, false)
		lc.setStatus(backoff)
		lc.lastBackoffTime = time.Now().Add(-6 * time.Second)

		runAndWait(t, lc)
		if got := lc.loadStatus(); got != readyToFetch {
			t.Fatalf("status with nil stats=%d, want readyToFetch=%d", got, readyToFetch)
		}
	})

	t.Run("off then re-enable", func(t *testing.T) {
		mc := &mockClient{
			consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
				return nil, &tls.Error{HTTPCode: http.StatusTooManyRequests, Message: "throttled"}
			},
		}
		lc := newCounterTestConsumer(t, mc)
		lc.stats = newConsumerStatsState(MetricsLevelOff)
		c := &consumer{stats: lc.stats}
		atomic.StoreInt32(&c.lifecycleStage, lifecycleRunning)

		// A 429 while metrics are Off must not count.
		runAndWait(t, lc)
		if got := backoffEventCount(lc); got != 0 {
			t.Fatalf("while metrics Off: backoff events=%d, want 0", got)
		}

		if err := c.SetMetricsLevel(MetricsLevelCore); err != nil {
			t.Fatalf("re-enable metrics: %v", err)
		}

		// Re-enabling metrics while still in the existing backoff state must not
		// synthesize an event for that already-observed state.
		runAndWait(t, lc)
		if got := backoffEventCount(lc); got != 0 {
			t.Fatalf("after re-enable existing backoff: backoff events=%d, want 0", got)
		}
	})
}
