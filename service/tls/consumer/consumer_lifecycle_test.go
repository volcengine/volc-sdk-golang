package consumer

import (
	"context"
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/observability"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// assertNoGoroutineGrowthWithin asserts that the live goroutine count settles
// back to at most base+tolerance within the deadline, polling until it does.
//
// Polling avoids treating background goroutines that are still winding down
// after a supported asynchronous Stop timeout as a leak.
func assertNoGoroutineGrowthWithin(t *testing.T, base, tolerance int, within time.Duration, ctx string) {
	t.Helper()
	deadline := time.Now().Add(within)
	var last int
	for time.Now().Before(deadline) {
		last = runtime.NumGoroutine()
		if last <= base+tolerance {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("%s: goroutine count did not settle: %d still > base(%d)+tol(%d) after %v", ctx, last, base, tolerance, within)
}

// errSentinelInit is a sentinel error returned by failing init mocks.
var errSentinelInit = errors.New("sentinel init failure")

func TestStartNilConsumerReturnsError(t *testing.T) {
	var c *consumer

	if err := c.Start(); err == nil {
		t.Fatal("Start on a nil consumer must return an error")
	}
}

// initFailMockClient makes init() (DescribeConsumerGroups) return an error so
// Start() fails before launching goroutines. All other methods delegate to the
// embedded mockConsumerClient.
type initFailMockClient struct {
	mockConsumerClient
}

func (m *initFailMockClient) DescribeConsumerGroups(req *tls.DescribeConsumerGroupsRequest) (*tls.DescribeConsumerGroupsResponse, error) {
	return nil, errSentinelInit
}

// initBlockMockClient makes init() (DescribeConsumerGroups) block until
// initRelease is closed. Used to deterministically race Stop() against an
// in-flight init().
type initBlockMockClient struct {
	mockConsumerClient
	initRelease chan struct{}
}

func (m *initBlockMockClient) DescribeConsumerGroups(req *tls.DescribeConsumerGroupsRequest) (*tls.DescribeConsumerGroupsResponse, error) {
	<-m.initRelease
	return &tls.DescribeConsumerGroupsResponse{
		ConsumerGroups: []*tls.ConsumerGroupResp{{ConsumerGroupName: req.ConsumerGroupName}},
	}, nil
}

// initTransientFailMockClient makes init() fail the first N times (controlled
// by a counter) then succeed. Used to verify a transient init failure allows a
// subsequent Start() retry.
type initTransientFailMockClient struct {
	mockConsumerClient
	failRemaining int32 // atomic
}

func (m *initTransientFailMockClient) DescribeConsumerGroups(req *tls.DescribeConsumerGroupsRequest) (*tls.DescribeConsumerGroupsResponse, error) {
	if atomic.AddInt32(&m.failRemaining, -1) >= 0 {
		return nil, errSentinelInit
	}
	return &tls.DescribeConsumerGroupsResponse{
		ConsumerGroups: []*tls.ConsumerGroupResp{{ConsumerGroupName: req.ConsumerGroupName}},
	}, nil
}

// heartbeatExpiredMockClient makes ConsumeLogs block until consumeRelease is
// closed, then return ErrConsumerHeartbeatExpired. Used to reproduce the
// shutdown race where a worker tries to notify an exited heartbeat runner.
type heartbeatExpiredMockClient struct {
	mockConsumerClient
	consumeRelease chan struct{}
	consumeCalls   int64 // atomic
}

func (m *heartbeatExpiredMockClient) ConsumerHeartbeat(req *tls.ConsumerHeartbeatRequest) (*tls.ConsumerHeartbeatResponse, error) {
	return &tls.ConsumerHeartbeatResponse{Shards: []*tls.ConsumeShard{{TopicID: "hb-test-topic", ShardID: 0}}}, nil
}

func (m *heartbeatExpiredMockClient) ConsumeLogs(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
	atomic.AddInt64(&m.consumeCalls, 1)
	<-m.consumeRelease
	return nil, &tls.Error{HTTPCode: 400, Code: tls.ErrConsumerHeartbeatExpired, Message: "heartbeat expired"}
}

// newLifecycleTestConsumer builds a consumer (without starting it) backed by the given
// client with a short fetch interval and a configurable stopTimeout.
func newLifecycleTestConsumer(t *testing.T, client tls.Client, stopTimeout time.Duration) *consumer {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "lifecycle-test-project"
	conf.TopicIDList = []string{"lifecycle-test-topic"}
	conf.ConsumerGroupName = "lifecycle-test-group"
	conf.ConsumerName = "lifecycle-test-consumer"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = MetricsLevelCore
	conf.DataFetchIntervalInMillisecond = 10
	conf.HeartbeatIntervalInSecond = 1
	conf.FlushCheckpointIntervalSecond = 1

	c, err := newTestConsumerWithClient(context.Background(), conf, client, func(topicID string, shardID int, l *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newTestConsumerWithClient: %v", err)
	}
	c.stopTimeout = stopTimeout
	return c
}

// TestInitFailureStopReturnsInOneSecond verifies that after init()
// returns a sentinel error (so goroutines never start), Stop() must return
// within 1s, the final snapshot must be stateClosed + consistencyFinal, and
// repeated GetStats must be fully frozen.
func TestInitFailureStopReturnsInOneSecond(t *testing.T) {
	mc := &initFailMockClient{}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	if err := c.Start(); !errors.Is(err, errSentinelInit) {
		t.Fatalf("Start() err = %v, want sentinel init error", err)
	}

	// Stop must return well within 1s (it must NOT block on mainLoopDone, which
	// never closes because run() was never started).
	start := time.Now()
	c.Stop()
	elapsed := time.Since(start)
	if elapsed > 1*time.Second {
		t.Fatalf("Stop() took %v after init failure; want < 1s (must not block on mainLoopDone)", elapsed)
	}

	st := c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State = %s, want stateClosed", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency = %s, want consistencyFinal", st.Consistency)
	}

	// Final snapshot must be immutable across repeated reads.
	st1 := c.collectMeasurements()
	for i := 0; i < 5; i++ {
		st2 := c.collectMeasurements()
		if st2.SnapshotSequence != st1.SnapshotSequence {
			t.Fatalf("final snapshot sequence changed after close: %d -> %d", st1.SnapshotSequence, st2.SnapshotSequence)
		}
		if st2.CapturedAtUnixMs != st1.CapturedAtUnixMs {
			t.Fatalf("final snapshot clock changed after close: %d -> %d", st1.CapturedAtUnixMs, st2.CapturedAtUnixMs)
		}
	}
}

// TestStopBeforeStartNoPanic verifies that Stop() before Start()
// must not panic or block, and a later explicit Start() may begin a fresh
// lifecycle generation. Repeated / concurrent Stop() calls remain safe and
// idempotent within each generation.
func TestStopBeforeStartNoPanic(t *testing.T) {
	mc := &mockConsumerClient{}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	// Stop before Start: must not panic (cancel is nil) and must return fast.
	start := time.Now()
	c.Stop()
	if elapsed := time.Since(start); elapsed > 1*time.Second {
		t.Fatalf("Stop() before Start took %v; want fast return", elapsed)
	}

	st := c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State after Stop-before-Start = %s, want stateClosed", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency after Stop-before-Start = %s, want consistencyFinal", st.Consistency)
	}
	firstInstanceID := st.InstanceID
	firstSequence := st.SnapshotSequence

	// Repeated Stop must not panic.
	c.Stop()

	// Stop-before-Start is a completed generation, not a permanent terminal
	// object. An explicit Start publishes the next generation and thaws the
	// retained stats state.
	if err := c.Start(); err != nil {
		t.Fatalf("Start() after Stop-before-Start = %v, want nil", err)
	}
	running := c.collectMeasurements()
	if running.State != stateRunning || running.Consistency != consistencyWeak {
		t.Fatalf("state after restart = %s/%s, want running/weak", running.State, running.Consistency)
	}
	if running.InstanceID != firstInstanceID {
		t.Fatalf("InstanceID changed across restart: %q -> %q", firstInstanceID, running.InstanceID)
	}
	if running.SnapshotSequence <= firstSequence {
		t.Fatalf("SnapshotSequence did not continue across restart: %d -> %d", firstSequence, running.SnapshotSequence)
	}
	c.Stop()
	st = c.collectMeasurements()
	if st.State != stateClosed || st.Consistency != consistencyFinal {
		t.Fatalf("state after restarted Stop = %s/%s, want closed/final", st.State, st.Consistency)
	}

	// 8 concurrent Stop calls must all return safely (idempotent via the current
	// generation's stopOnce pointer).
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Stop()
		}()
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("concurrent Stop() calls did not return within 3s")
	}
}

// TestTransientInitFailureAllowsRetry verifies that a single
// transient init failure, without Stop(), allows a subsequent Start() to
// succeed.
func TestTransientInitFailureAllowsRetry(t *testing.T) {
	mc := &initTransientFailMockClient{failRemaining: 1}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	// First Start fails (init fails once).
	if err := c.Start(); !errors.Is(err, errSentinelInit) {
		t.Fatalf("first Start() err = %v, want sentinel", err)
	}

	// Second Start succeeds (init now succeeds).
	if err := c.Start(); err != nil {
		t.Fatalf("second Start() err = %v, want nil", err)
	}

	// Clean up.
	c.Stop()
	st := c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State = %s, want stateClosed", st.State)
	}
}

// TestRepeatedStartStopRestart verifies that a second Start() while a
// generation is running is rejected, while an explicit Start() after a fully
// completed Stop() begins a new generation on the same Consumer object. The
// stats object is retained across all three generations: InstanceID and
// cumulative counters stay stable, and the frozen snapshot sequence resumes
// strictly above the previous generation's final value.
func TestRepeatedStartStopRestart(t *testing.T) {
	mc := &mockConsumerClient{}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	// Baseline BEFORE any Start, so the tolerance is measured against this
	// test's own starting point rather than a process-wide instantaneous count.
	baseline := runtime.NumGoroutine()

	var instanceID string
	var previousFinalSequence uint64
	var previousFinalGeneration uint64
	var previousFetchCalls uint64
	for round := 0; round < 3; round++ {
		if err := c.Start(); err != nil {
			t.Fatalf("round %d Start() err = %v, want nil", round, err)
		}

		running := c.collectMeasurements()
		if running.State != stateRunning || running.Consistency != consistencyWeak {
			t.Fatalf("round %d running snapshot = %s/%s, want running/weak", round, running.State, running.Consistency)
		}
		if round == 0 {
			instanceID = running.InstanceID
			if generation := c.stats.collectionGenerationNow(); generation != 1 {
				t.Fatalf("first Start changed initial collection generation: %d", generation)
			}
		} else {
			if running.InstanceID != instanceID {
				t.Fatalf("round %d InstanceID changed: %q -> %q", round, instanceID, running.InstanceID)
			}
			generation := c.stats.collectionGenerationNow()
			if generation != previousFinalGeneration+1 {
				t.Fatalf("round %d collection generation = %d, want %d", round, generation, previousFinalGeneration+1)
			}
			if running.SnapshotSequence <= previousFinalSequence {
				t.Fatalf("round %d SnapshotSequence did not continue: previous final %d, running %d", round, previousFinalSequence, running.SnapshotSequence)
			}
			if running.LastSuccessAgeMs != nil || running.LastErrorAgeMs != nil {
				t.Fatalf("round %d restart retained previous health ages: success=%v error=%v", round, running.LastSuccessAgeMs, running.LastErrorAgeMs)
			}
		}

		// Model work completed by the generation to prove that cumulative
		// counters are not reset when the stats snapshot is thawed.
		atomic.AddInt64(&c.stats.fetchCalls, int64(round+1))
		if running.FetchCalls != nil && uint64(*running.FetchCalls) < previousFetchCalls {
			t.Fatalf("round %d FetchCalls rolled back: %d -> %d", round, previousFetchCalls, *running.FetchCalls)
		}

		// A second Start while running is still rejected.
		if err := c.Start(); err == nil {
			t.Fatalf("round %d second Start() returned nil; want error", round)
		}

		// Seed each generation immediately before Stop. The closed/final
		// snapshot must preserve these health ages, while the next explicit
		// restart must clear them as part of the new collection baseline.
		atomic.StoreInt64(&c.stats.lastSuccessTimeUnixMs, unixMillis(time.Now().Add(-2*time.Second)))
		atomic.StoreInt64(&c.stats.lastErrorTimeUnixMs, unixMillis(time.Now().Add(-3*time.Second)))
		c.Stop()
		final := c.collectMeasurements()
		if final.State != stateClosed || final.Consistency != consistencyFinal {
			t.Fatalf("round %d final snapshot = %s/%s, want closed/final", round, final.State, final.Consistency)
		}
		if final.InstanceID != instanceID {
			t.Fatalf("round %d final InstanceID changed: %q -> %q", round, instanceID, final.InstanceID)
		}
		if final.SnapshotSequence < running.SnapshotSequence {
			t.Fatalf("round %d final sequence rolled back: running %d -> final %d", round, running.SnapshotSequence, final.SnapshotSequence)
		}
		if final.LastSuccessAgeMs == nil || final.LastErrorAgeMs == nil {
			t.Fatalf("round %d final health ages were cleared: success=%v error=%v", round, final.LastSuccessAgeMs, final.LastErrorAgeMs)
		}
		finalAgain := c.collectMeasurements()
		if finalAgain.SnapshotSequence != final.SnapshotSequence || finalAgain.CapturedAtUnixMs != final.CapturedAtUnixMs {
			t.Fatalf("round %d final snapshot is not frozen", round)
		}
		if final.FetchCalls != nil {
			previousFetchCalls = *final.FetchCalls
		}
		previousFinalSequence = final.SnapshotSequence
		previousFinalGeneration = c.stats.collectionGenerationNow()
	}

	// All three generations' goroutines must unwind; no rejected Start attempt
	// may leave a survivor behind.
	assertNoGoroutineGrowthWithin(t, baseline, 1, 3*time.Second, "lifecycle repeated/failed Start")
}

// TestPublicAdaptersRebaselineOnRestart verifies the public projection at
// the exact closed -> running boundary. DeltaAdapter and EvaluateHealth must
// both treat the new collection generation as a fresh baseline even though
// InstanceID and cumulative counters remain unchanged.
func TestPublicAdaptersRebaselineOnRestart(t *testing.T) {
	mc := &mockConsumerClient{}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	if err := c.Start(); err != nil {
		t.Fatalf("first Start() err = %v, want nil", err)
	}
	first, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("first GetStats: %v", err)
	}
	c.Stop()
	closed, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("closed GetStats: %v", err)
	}
	if closed.LifecycleState != observability.LifecycleClosed {
		t.Fatalf("closed lifecycle state = %s, want closed", closed.LifecycleState)
	}
	if closed.CollectionGeneration != first.CollectionGeneration {
		t.Fatalf("collection generation changed during Stop: %d -> %d", first.CollectionGeneration, closed.CollectionGeneration)
	}

	adapter := &observability.DeltaAdapter{}
	adapter.Observe(observability.CounterSample{
		Header: closed.SnapshotHeader,
		Counters: map[string]uint64{
			"fetch_calls_success_total": closed.Operational.Counters.FetchCallsSuccessTotal,
		},
	})

	if err := c.Start(); err != nil {
		t.Fatalf("restart Start() err = %v, want nil", err)
	}
	restarted, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("restarted GetStats: %v", err)
	}
	if restarted.CollectionGeneration != closed.CollectionGeneration+1 {
		t.Fatalf("restart collection generation = %d, want %d", restarted.CollectionGeneration, closed.CollectionGeneration+1)
	}
	delta := adapter.Observe(observability.CounterSample{
		Header: restarted.SnapshotHeader,
		Counters: map[string]uint64{
			"fetch_calls_success_total": restarted.Operational.Counters.FetchCallsSuccessTotal,
		},
	})
	if delta.Valid || delta.ResetReason != observability.DeltaCollectionChanged {
		t.Fatalf("closed -> restart delta = %+v, want collection_changed rebaseline", delta)
	}

	var previousSignals *observability.HealthSignals
	if closed.Health != nil {
		previousSignals = &closed.Health.Signals
	}
	var currentSignals *observability.HealthSignals
	if restarted.Health != nil {
		currentSignals = &restarted.Health.Signals
	}
	health := observability.EvaluateHealth(
		&observability.HealthSample{Header: closed.SnapshotHeader, Signals: previousSignals},
		observability.HealthSample{Header: restarted.SnapshotHeader, Signals: currentSignals},
		observability.DefaultHealthPolicy(),
	)
	if !health.InsufficientData || len(health.ReasonCodes) != 1 || health.ReasonCodes[0] != observability.ReasonCollectionChanged {
		t.Fatalf("closed -> restart health = %+v, want collection_changed rebaseline", health)
	}
	c.Stop()
}

// TestConcurrentStartOnlyOneSucceeds verifies that two concurrent
// Start() calls have at most one success and must not double the goroutine
// count.
func TestConcurrentStartOnlyOneSucceeds(t *testing.T) {
	mc := &mockConsumerClient{}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	// Baseline BEFORE the concurrent Start, measured against this test's own
	// starting point (not a process-wide instantaneous count).
	time.Sleep(100 * time.Millisecond)
	gBefore := runtime.NumGoroutine()

	var wg sync.WaitGroup
	results := make(chan error, 2)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- c.Start()
		}()
	}
	wg.Wait()
	close(results)

	successes := 0
	for err := range results {
		if err == nil {
			successes++
		}
	}
	if successes != 1 {
		t.Fatalf("concurrent Start: %d succeeded, want exactly 1", successes)
	}

	// The definitive proof that the losing Start did NOT launch a second set of
	// goroutines is that after Stop() the count drains back to the baseline: a
	// double-start would leak a second checkpoint/heartbeat/run trio (or
	// double-close panic). A running consumer has transient worker/fetch tasks,
	// so the pre-Stop instantaneous count is noisy and NOT asserted; the
	// post-Stop drain is the reliable, leak-sensitive signal.
	c.Stop()
	assertNoGoroutineGrowthWithin(t, gBefore, 1, 3*time.Second, "lifecycle concurrent Start")
}

// TestStopDuringInitLinearizes verifies that while init() is
// blocked, Stop() can linearize and close; when init() later returns it must
// not launch goroutines. Stop() must NOT block waiting for the network init.
func TestStopDuringInitLinearizes(t *testing.T) {
	mc := &initBlockMockClient{initRelease: make(chan struct{})}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	// Start in a goroutine; it blocks in init() (DescribeConsumerGroups).
	startErr := make(chan error, 1)
	go func() { startErr <- c.Start() }()

	// Give Start time to enter the blocking init.
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		if atomic.LoadInt32(&c.lifecycleStage) == lifecycleStarting {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	if atomic.LoadInt32(&c.lifecycleStage) != lifecycleStarting {
		t.Fatal("Start() never reached lifecycleStarting (init block)")
	}

	// Stop must return quickly — it must NOT wait for the blocked init.
	stopStart := time.Now()
	stopDone := make(chan struct{})
	go func() { c.Stop(); close(stopDone) }()
	select {
	case <-stopDone:
	case <-time.After(2 * time.Second):
		t.Fatal("Stop() blocked for 2s during init; must not wait on network init")
	}
	if elapsed := time.Since(stopStart); elapsed > 1*time.Second {
		t.Fatalf("Stop() took %v during init; want < 1s", elapsed)
	}

	// State must already be Closed/Final even though init is still blocked.
	st := c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State during blocked init = %s, want stateClosed", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency during blocked init = %s, want consistencyFinal", st.Consistency)
	}

	// Release init; Start must observe lifecycleClosing and NOT launch
	// goroutines (no runLaunched flip, no mainLoopDone close).
	close(mc.initRelease)
	select {
	case err := <-startErr:
		if err == nil {
			t.Fatal("Start() returned nil after Stop during init; want error")
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Start() did not return after init release")
	}

	// No main loop should be running.
	if c.runLaunched {
		t.Fatal("runLaunched is true after Stop-during-init; goroutines must not be launched")
	}

	// Final state remains Closed/Final and frozen.
	st = c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("final State = %s, want stateClosed", st.State)
	}
}

// TestHeartbeatExpiredRaceDrain verifies that a worker blocked in
// ConsumeLogs, released after Stop() has exited the heartbeat runner, returns
// ErrConsumerHeartbeatExpired. The worker must drain (not block forever on the
// unbuffered heartbeatExpiredCh) so Stop() finalizes to Closed/Final with no
// goroutine leak.
func TestHeartbeatExpiredRaceDrain(t *testing.T) {
	mc := &heartbeatExpiredMockClient{consumeRelease: make(chan struct{})}
	c := newLifecycleTestConsumer(t, mc, 200*time.Millisecond)

	// Baseline before Start, for the post-finalization leak check.
	baseline := runtime.NumGoroutine()

	if err := c.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}

	oneShard := []*tls.ConsumeShard{{TopicID: "hb-test-topic", ShardID: 0}}
	c.heartbeat.setShards(oneShard)

	// Wait for the worker to enter the blocked ConsumeLogs.
	waitDeadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(waitDeadline) {
		if atomic.LoadInt64(&mc.consumeCalls) > 0 {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	if atomic.LoadInt64(&mc.consumeCalls) == 0 {
		t.Fatal("worker never entered ConsumeLogs")
	}

	// Stop: cancels runtime ctx, heartbeat runner exits (no longer reading
	// heartbeatExpiredCh). Stop should time out (worker still blocked) and
	// return closing/weak, leaving the single finalizer running async.
	stopStart := time.Now()
	c.Stop()
	stopElapsed := time.Since(stopStart)
	// Stop should time out bounded (worker still blocked in ConsumeLogs).
	if stopElapsed > 2*time.Second {
		t.Fatalf("Stop() took %v with blocked worker; expected bounded timeout", stopElapsed)
	}

	// Now release ConsumeLogs: it returns ErrConsumerHeartbeatExpired. The
	// worker tries to send to heartbeatExpiredCh (no receiver) and commitCh.
	// Without the cancel-aware select this blocks forever; with it, the worker
	// aborts via ctx.Done() and exits, draining workerWG so the finalizer runs.
	close(mc.consumeRelease)

	// The finalizer should now complete and transition to Closed/Final.
	finalDeadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(finalDeadline) {
		if c.collectMeasurements().State == stateClosed {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	st := c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State after release = %s, want stateClosed (worker did not drain; blocked on heartbeatExpiredCh?)", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency after release = %s, want consistencyFinal", st.Consistency)
	}

	// No goroutine leak: after finalization the worker, heartbeat, checkpoint and
	// the async finalizer must all have exited, so the count drains back to the
	// pre-Start baseline. Poll (rather than a single-shot delta) so a prior
	// iteration's still-settling background goroutine cannot cause a false leak.
	assertNoGoroutineGrowthWithin(t, baseline, 1, 3*time.Second, "lifecycle after finalization")
}

// assertClosedSnapshotSound asserts the stateClosed invariant: any snapshot in
// stateClosed must simultaneously be consistencyFinal, carry a valid (non-zero)
// close-time clock, and expose a frozen (never rolled-back / never re-incrementing)
// sequence. Returns the snapshot so callers can chain immutability checks.
func assertClosedSnapshotSound(t *testing.T, st consumerMeasurements, ctx string) {
	t.Helper()
	if st.State != stateClosed {
		return
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("%s: stateClosed but Consistency = %s, want consistencyFinal", ctx, st.Consistency)
	}
	if st.CapturedAtUnixMs <= 0 {
		t.Fatalf("%s: stateClosed but CapturedAtUnixMs = %d, want a valid close-time clock", ctx, st.CapturedAtUnixMs)
	}
	// The frozen high bit must never leak into the public sequence.
	if st.SnapshotSequence&snapshotSeqFrozenBit != 0 {
		t.Fatalf("%s: public SnapshotSequence carries frozen high bit: %d", ctx, st.SnapshotSequence)
	}
}

// TestStopBeforeStartConcurrentGetStatsFrozen verifies the contract for the
// unstarted close path: while many readers hammer GetStats,
// a Stop() before Start() races them. There must be NO window where a snapshot
// reports stateClosed without also being consistencyFinal, with a valid
// close-time clock, and a frozen sequence. Once any reader observes stateClosed,
// every later snapshot it reads must be byte-for-byte identical (fully frozen).
//
// This is the regression test for the publish-order bug where the lifecycleNew
// branch set c.closed=1 BEFORE freezeOnClose(), so a concurrent GetStats could
// observe stateClosed+Final with CapturedAtUnixMs=0 and a still-live sequence.
func TestStopBeforeStartConcurrentGetStatsFrozen(t *testing.T) {
	mc := &mockConsumerClient{}
	c := newLifecycleTestConsumer(t, mc, 5*time.Second)

	const numReaders = 16
	var readers sync.WaitGroup
	start := make(chan struct{})

	for i := 0; i < numReaders; i++ {
		readers.Add(1)
		go func() {
			defer readers.Done()
			<-start // line readers up so they overlap the Stop tightly
			var frozen *consumerMeasurements
			for j := 0; j < 500; j++ {
				st := c.collectMeasurements()
				assertClosedSnapshotSound(t, st, "reader")
				if st.State == stateClosed {
					// First closed observation pins the final snapshot; every
					// subsequent read must be completely unchanged.
					if frozen == nil {
						snap := st
						frozen = &snap
						continue
					}
					if st.SnapshotSequence != frozen.SnapshotSequence ||
						st.CapturedAtUnixMs != frozen.CapturedAtUnixMs ||
						st.Consistency != frozen.Consistency ||
						st.State != frozen.State {
						t.Errorf("final snapshot mutated after close: %+v -> seq=%d cap=%d cons=%s state=%s",
							*frozen, st.SnapshotSequence, st.CapturedAtUnixMs, st.Consistency, st.State)
						return
					}
				}
			}
		}()
	}

	// 8 concurrent Stop-before-Start callers, racing the readers.
	var stoppers sync.WaitGroup
	for i := 0; i < 8; i++ {
		stoppers.Add(1)
		go func() {
			defer stoppers.Done()
			<-start
			c.Stop()
		}()
	}

	close(start)
	stoppers.Wait()
	readers.Wait()

	// Terminal invariant: fully frozen and sound.
	st1 := c.collectMeasurements()
	assertClosedSnapshotSound(t, st1, "terminal")
	if st1.State != stateClosed {
		t.Fatalf("terminal State = %s, want stateClosed", st1.State)
	}
	for i := 0; i < 5; i++ {
		st2 := c.collectMeasurements()
		if st2.SnapshotSequence != st1.SnapshotSequence || st2.CapturedAtUnixMs != st1.CapturedAtUnixMs {
			t.Fatalf("terminal snapshot not frozen: seq %d->%d cap %d->%d",
				st1.SnapshotSequence, st2.SnapshotSequence, st1.CapturedAtUnixMs, st2.CapturedAtUnixMs)
		}
	}
}
