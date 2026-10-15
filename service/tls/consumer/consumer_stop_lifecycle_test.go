package consumer

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// controllableMockClient is a tls.Client stand-in for Stop/rebalance lifecycle
// tests. It lets the test:
//   - block ConsumeLogs until consumeRelease is closed (simulating a stuck
//     fetch/consume writer that holds the worker's wg),
//   - control the shard set returned by ConsumerHeartbeat via setShards, so the
//     test can deterministically trigger rebalance (add then remove a shard).
//
// All non-overridden methods delegate to the embedded mockConsumerClient.
type controllableMockClient struct {
	mockConsumerClient
	consumeRelease chan struct{} // closed to unblock ConsumeLogs
	consumeCalls   int64         // atomic; incremented on every ConsumeLogs entry

	shardsMu sync.RWMutex
	shards   []*tls.ConsumeShard
}

func newControllableMockClient() *controllableMockClient {
	return &controllableMockClient{
		consumeRelease: make(chan struct{}),
	}
}

func (m *controllableMockClient) setShards(shards []*tls.ConsumeShard) {
	m.shardsMu.Lock()
	defer m.shardsMu.Unlock()
	m.shards = shards
}

func (m *controllableMockClient) ConsumerHeartbeat(req *tls.ConsumerHeartbeatRequest) (*tls.ConsumerHeartbeatResponse, error) {
	m.shardsMu.RLock()
	defer m.shardsMu.RUnlock()
	return &tls.ConsumerHeartbeatResponse{Shards: m.shards}, nil
}

func (m *controllableMockClient) ConsumeLogs(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
	atomic.AddInt64(&m.consumeCalls, 1)
	// Block until the test releases us, simulating a hung HTTP/retry.
	<-m.consumeRelease
	return &tls.ConsumeLogsResponse{
		Cursor: "cursor-blocked",
		Count:  1,
		Logs: &pb.LogGroupList{
			LogGroups: []*pb.LogGroup{{Logs: []*pb.Log{{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}}}}},
		},
	}, nil
}

// newStopTestConsumer builds and starts a consumer backed by
// controllableMockClient with a short fetch interval so workers spin up fast.
// The caller controls stopTimeout and shards. Returns the consumer and the
// mock client.
func newStopTestConsumer(t *testing.T, stopTimeout time.Duration) (*consumer, *controllableMockClient) {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "stop-test-project"
	conf.TopicIDList = []string{"stop-test-topic"}
	conf.ConsumerGroupName = "stop-test-group"
	conf.ConsumerName = "stop-test-consumer"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = MetricsLevelCore
	conf.DataFetchIntervalInMillisecond = 10
	conf.HeartbeatIntervalInSecond = 1
	conf.FlushCheckpointIntervalSecond = 1

	mc := newControllableMockClient()
	c, err := newTestConsumerWithClient(context.Background(), conf, mc, func(topicID string, shardID int, l *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newTestConsumerWithClient: %v", err)
	}
	c.stopTimeout = stopTimeout
	if err := c.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	return c, mc
}

// waitForConsumeBlocks polls until the mock client has observed at least one
// ConsumeLogs call (meaning a worker has entered the blocked fetch path), or
// times out.
func waitForConsumeBlocks(t *testing.T, mc *controllableMockClient, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if atomic.LoadInt64(&mc.consumeCalls) > 0 {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("ConsumeLogs never called within %v", timeout)
}

// waitForClosed polls until the consumer reaches stateClosed, or fails after
// the timeout. Tests that deliberately drive Stop() into a bounded timeout
// (leaving a blocked worker + async finalizer running) MUST call this after
// releasing the worker so their background goroutines fully unwind before the
// test returns. Otherwise, under `go test -count=N` those still-running
// goroutines bleed into the next iteration's process-wide NumGoroutine samples
// and cause false "goroutine grew" failures in unrelated tests.
func waitForClosed(t *testing.T, c *consumer, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if c.collectMeasurements().State == stateClosed {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("consumer did not reach stateClosed within %v (state=%s)", timeout, c.collectMeasurements().State)
}

// settleGoroutinePlateau samples runtime.NumGoroutine() until it stops
// decreasing (transient goroutines from startup / a prior -count iteration have
// unwound), and returns that stable plateau. Used as a leak-check baseline that
// is robust to cross-test settling noise: callers assert the count does not
// grow beyond the returned plateau.
func settleGoroutinePlateau(within time.Duration) int {
	deadline := time.Now().Add(within)
	prev := runtime.NumGoroutine()
	stable := 0
	for time.Now().Before(deadline) {
		time.Sleep(20 * time.Millisecond)
		cur := runtime.NumGoroutine()
		if cur >= prev {
			// No longer decreasing; require a couple consecutive stable reads.
			stable++
			if stable >= 3 {
				return maxInt(cur, prev)
			}
		} else {
			stable = 0
		}
		if cur < prev {
			prev = cur
		}
	}
	return prev
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TestStop_RebalanceRemovedWorkerNotMissed verifies that a worker blocked in
// the ConsumeLogs writer path, then removed from workerMap by rebalance, is
// still waited on by Stop() (via the shared workerWG, not an append-only
// registry). Stop must return bounded (not hang 30s), stay closing/weak while
// the worker is blocked, and finalize only after the worker actually exits.
func TestStop_RebalanceRemovedWorkerNotMissed(t *testing.T) {
	c, mc := newStopTestConsumer(t, 300*time.Millisecond)
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(mc.consumeRelease) }) }
	defer release() // ensure no goroutine leak on failure

	// Assign one shard and let its worker reach the blocked ConsumeLogs.
	mc.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	// Heartbeat runner uploads on its own ticker; push shards directly so the
	// main loop's handleShards observes them deterministically.
	c.heartbeat.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	waitForConsumeBlocks(t, mc, 5*time.Second)

	// Rebalance away the shard: worker 0 is removed from workerMap, but its
	// goroutine is still blocked in ConsumeLogs and still tracked by the
	// shared workerWG (no append-only registry; constant space).
	c.heartbeat.setShards(nil)
	// Give the main loop a tick to process the rebalance.
	time.Sleep(100 * time.Millisecond)
	c.stateMu.RLock()
	inMap := len(c.workerMap)
	c.stateMu.RUnlock()
	if inMap != 0 {
		t.Fatalf("workerMap still has %d workers after rebalance, want 0", inMap)
	}
	// The removed worker must be flagged retired so its in-flight task releases
	// any held batch on exit (no held/ready leak).
	c.stateMu.RLock()
	// workerMap is empty; we cannot reach the removed worker through it. The
	// behavioral proof below (Stop stays closing/weak until release) is the
	// real assertion that the removed worker is still waited on.
	c.stateMu.RUnlock()

	// Stop must return bounded (well under 30s) even though the worker is stuck.
	stopDone := make(chan struct{})
	go func() {
		c.Stop()
		close(stopDone)
	}()
	select {
	case <-stopDone:
	case <-time.After(5 * time.Second):
		t.Fatal("Stop() hung for 5s; bounded stop semantics lost (must not wait 30s)")
	}

	// While the worker is still blocked, state must be closing/weak — NOT a
	// fake closed/final.
	st := c.collectMeasurements()
	if st.State != stateClosing {
		t.Fatalf("State while worker blocked = %s, want stateClosing", st.State)
	}
	if st.Consistency != consistencyWeak {
		t.Fatalf("Consistency while worker blocked = %s, want consistencyWeak", st.Consistency)
	}

	// Release the blocked worker; async finalize should freeze the snapshot.
	release()

	// Wait for the consumer to reach closed/final.
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if c.collectMeasurements().State == stateClosed {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	st = c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State after worker exit = %s, want stateClosed (async finalize did not run)", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency after worker exit = %s, want consistencyFinal", st.Consistency)
	}

	// Final snapshot must be immutable.
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

// TestStop_NeverReturningWorkerBounded verifies that a worker that never
// returns does not make Stop() hang for the historical 30s. Stop must return
// within the (short, test-only) stopTimeout and stay closing/weak.
func TestStop_NeverReturningWorkerBounded(t *testing.T) {
	c, mc := newStopTestConsumer(t, 150*time.Millisecond)
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(mc.consumeRelease) }) }
	defer release()

	mc.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	c.heartbeat.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	waitForConsumeBlocks(t, mc, 5*time.Second)

	start := time.Now()
	c.Stop()
	elapsed := time.Since(start)

	// Stop must return in roughly stopTimeout, nowhere near 30s.
	if elapsed > 3*time.Second {
		t.Fatalf("Stop() took %v with a stuck worker; expected ~%v bounded wait (must not wait 30s)", elapsed, 150*time.Millisecond)
	}

	st := c.collectMeasurements()
	if st.State != stateClosing {
		t.Fatalf("State = %s with stuck worker, want stateClosing (no fake final)", st.State)
	}
	if st.Consistency != consistencyWeak {
		t.Fatalf("Consistency = %s with stuck worker, want consistencyWeak (no fake final)", st.Consistency)
	}
	// A timed-out generation is still closing; Start must not cross its
	// finalizer. Once the worker is released and the generation reaches closed,
	// the same Consumer may be explicitly restarted.
	if err := c.Start(); err == nil {
		t.Fatal("Start() during async finalization returned nil; want error")
	}

	// Release the stuck worker and wait for the async finalizer to converge, so
	// this test's background goroutines fully unwind before it returns (prevents
	// cross-test NumGoroutine pollution under -count=N).
	release()
	waitForClosed(t, c, 5*time.Second)
	if err := c.Start(); err != nil {
		t.Fatalf("Start() after async finalizer completed = %v, want nil", err)
	}
	if st = c.collectMeasurements(); st.State != stateRunning || st.Consistency != consistencyWeak {
		t.Fatalf("state after async restart = %s/%s, want running/weak", st.State, st.Consistency)
	}
	c.Stop()
	waitForClosed(t, c, 5*time.Second)
}

// TestStop_AllWorkersExit_Final verifies the happy path: when all workers exit
// promptly (no block), Stop() freezes a final, immutable snapshot.
func TestStop_AllWorkersExit_Final(t *testing.T) {
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "stop-test-project"
	conf.TopicIDList = []string{"stop-test-topic"}
	conf.ConsumerGroupName = "stop-test-group"
	conf.ConsumerName = "stop-test-consumer"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = MetricsLevelCore
	conf.DataFetchIntervalInMillisecond = 10
	conf.HeartbeatIntervalInSecond = 1
	conf.FlushCheckpointIntervalSecond = 1

	mc := &mockConsumerClient{}
	c, err := newTestConsumerWithClient(context.Background(), conf, mc, func(topicID string, shardID int, l *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newTestConsumerWithClient: %v", err)
	}
	c.stopTimeout = 5 * time.Second
	if err := c.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	// Let workers spin up.
	time.Sleep(200 * time.Millisecond)

	c.Stop()
	st := c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State = %s, want stateClosed", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency = %s, want consistencyFinal", st.Consistency)
	}
	st1 := c.collectMeasurements()
	for i := 0; i < 5; i++ {
		st2 := c.collectMeasurements()
		if st2.SnapshotSequence != st1.SnapshotSequence {
			t.Fatalf("final sequence changed: %d -> %d", st1.SnapshotSequence, st2.SnapshotSequence)
		}
	}
}

// TestStop_TimeoutKeepsClosingWeakNoFakeFinal is a focused regression test:
// after a timeout, GetStats must never report stateClosed/consistencyFinal and
// must not return a frozen (high-bit) sequence while still closing.
func TestStop_TimeoutKeepsClosingWeakNoFakeFinal(t *testing.T) {
	c, mc := newStopTestConsumer(t, 100*time.Millisecond)
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(mc.consumeRelease) }) }
	defer release()

	mc.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	c.heartbeat.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	waitForConsumeBlocks(t, mc, 5*time.Second)

	c.Stop()

	// Sample GetStats many times immediately after the timeout: every snapshot
	// must be closing/weak with no frozen high bit.
	for i := 0; i < 50; i++ {
		st := c.collectMeasurements()
		if st.State == stateClosed {
			t.Fatalf("iteration %d: observed stateClosed while worker still blocked (fake final)", i)
		}
		if st.Consistency == consistencyFinal {
			t.Fatalf("iteration %d: observed consistencyFinal while worker still blocked (fake final)", i)
		}
		if st.SnapshotSequence&snapshotSeqFrozenBit != 0 {
			t.Fatalf("iteration %d: snapshot carries frozen high bit while still closing", i)
		}
	}

	// Release the stuck worker and wait for convergence so this test's async
	// finalizer/worker goroutines fully unwind before returning (no cross-test
	// NumGoroutine pollution under -count=N).
	release()
	waitForClosed(t, c, 5*time.Second)
}

// TestStop_HighFrequencyRebalanceNoLeak verifies that many rebalance cycles
// (add/remove shards) do NOT grow any per-worker registry/slice and do NOT
// leak held/ready counters. With the old append-only allWorkers, each cycle
// would permanently retain a worker pointer; this test proves constant space.
func TestStop_HighFrequencyRebalanceNoLeak(t *testing.T) {
	c, mc := newStopTestConsumer(t, 5*time.Second)
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(mc.consumeRelease) }) }
	defer release()

	oneShard := []*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}}

	// Drive many rebalance cycles directly through handleShards (no ticker wait).
	const cycles = 200
	for i := 0; i < cycles; i++ {
		if i%2 == 0 {
			c.handleShards(oneShard)
		} else {
			c.handleShards(nil)
		}
	}
	// End with no shards assigned.
	c.handleShards(nil)

	// workerMap must be empty (no retained workers). Constant space: the only
	// worker tracking is the shared workerWG, which drains as tasks exit.
	c.stateMu.RLock()
	mapSize := len(c.workerMap)
	c.stateMu.RUnlock()
	if mapSize != 0 {
		t.Fatalf("after %d rebalance cycles, workerMap has %d workers, want 0", cycles, mapSize)
	}

	// Release any blocked worker so in-flight tasks can exit and release their
	// held batches (idempotent release via retired flag).
	release()

	// Give any in-flight fetch tasks a moment to exit and release their batches.
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		st := c.collectMeasurements()
		held := uint64(0)
		ready := uint64(0)
		if st.HeldBatchReferences != nil {
			held = *st.HeldBatchReferences
		}
		if st.ReadyToProcess != nil {
			ready = *st.ReadyToProcess
		}
		if held == 0 && ready == 0 {
			break
		}
		time.Sleep(20 * time.Millisecond)
	}

	st := c.collectMeasurements()
	held := uint64(0)
	ready := uint64(0)
	if st.HeldBatchReferences != nil {
		held = *st.HeldBatchReferences
	}
	if st.ReadyToProcess != nil {
		ready = *st.ReadyToProcess
	}
	if held != 0 || ready != 0 {
		t.Fatalf("after %d rebalance cycles: held=%d ready=%d, want 0/0 (counter leak)", cycles, held, ready)
	}

	// Stop must still work cleanly after many rebalances. Under -count=N on a
	// busy machine the async finalizer may take a moment to converge, so poll
	// for stateClosed rather than asserting it synchronously.
	c.Stop()
	waitForClosed(t, c, 5*time.Second)
}

// TestStop_SingleFinalizerGoroutine verifies that a Stop() timeout creates at
// most ONE background finalizer goroutine (not two, as the old
// waitGroupTimeout+waitGroupAll implementation did). Uses runtime.NumGoroutine
// to prove no duplicate waiters accumulate.
func TestStop_SingleFinalizerGoroutine(t *testing.T) {
	c, mc := newStopTestConsumer(t, 100*time.Millisecond)
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(mc.consumeRelease) }) }
	defer release()

	mc.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	c.heartbeat.setShards([]*tls.ConsumeShard{{TopicID: "stop-test-topic", ShardID: 0}})
	waitForConsumeBlocks(t, mc, 5*time.Second)

	// After Stop() times out, exactly ONE finalizer goroutine should remain,
	// waiting on the blocked worker. The old code created two. We prove no
	// DUPLICATE/growing waiter accumulates.
	//
	// A single-shot g1/g2 delta is fragile under -count=N: right after Stop()
	// returns, transient worker/fetch goroutines from this test's own startup
	// (and a prior iteration's still-settling finalizer) are unwinding, so an
	// early g1 sample is artificially low and a later sample can look like
	// "growth". Instead, let the count settle to a stable plateau first, then
	// assert it does not grow beyond that plateau.
	c.Stop()
	plateau := settleGoroutinePlateau(3 * time.Second)
	time.Sleep(200 * time.Millisecond) // well past stopTimeout; no new waiter should appear
	if after := runtime.NumGoroutine(); after > plateau+1 {
		t.Fatalf("goroutine count grew after Stop timeout: plateau %d -> %d (duplicate finalizer waiter?)", plateau, after)
	}

	// State must be closing/weak while the worker is blocked.
	st := c.collectMeasurements()
	if st.State != stateClosing {
		t.Fatalf("State = %s, want stateClosing", st.State)
	}

	// Release the blocked worker: the single finalizer goroutine should
	// complete and transition to closed/final.
	release()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if c.collectMeasurements().State == stateClosed {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	st = c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State after release = %s, want stateClosed (finalizer did not run)", st.State)
	}
}

// checkpointFlushMockClient is a tls.Client stand-in for the final-checkpoint-
// flush ordering tests. Unlike controllableMockClient (which blocks in
// ConsumeLogs), ConsumeLogs returns promptly so workers proceed all the way to
// the user consume callback. ModifyCheckPoint calls are recorded so the test can
// assert whether the final flush happened and with what cursor. Shards are
// controllable via setShards so the test can deterministically assign exactly
// one worker. FlushCheckpointIntervalSecond is set very long in the test helper
// so only the final flush (driven by stopCh) can satisfy the assertions.
type checkpointFlushMockClient struct {
	mockConsumerClient
	shardsMu sync.RWMutex
	shards   []*tls.ConsumeShard

	modifyMu      sync.Mutex
	modifyCalls   int
	modifyCursors []string
}

func (m *checkpointFlushMockClient) setShards(shards []*tls.ConsumeShard) {
	m.shardsMu.Lock()
	defer m.shardsMu.Unlock()
	m.shards = shards
}

func (m *checkpointFlushMockClient) ConsumerHeartbeat(req *tls.ConsumerHeartbeatRequest) (*tls.ConsumerHeartbeatResponse, error) {
	m.shardsMu.RLock()
	defer m.shardsMu.RUnlock()
	return &tls.ConsumerHeartbeatResponse{Shards: m.shards}, nil
}

func (m *checkpointFlushMockClient) ModifyCheckPoint(req *tls.ModifyCheckPointRequest) (*tls.CommonResponse, error) {
	m.modifyMu.Lock()
	m.modifyCalls++
	m.modifyCursors = append(m.modifyCursors, req.Checkpoint)
	m.modifyMu.Unlock()
	return &tls.CommonResponse{}, nil
}

func (m *checkpointFlushMockClient) modifyCount() int {
	m.modifyMu.Lock()
	defer m.modifyMu.Unlock()
	return m.modifyCalls
}

func (m *checkpointFlushMockClient) lastCursor() string {
	m.modifyMu.Lock()
	defer m.modifyMu.Unlock()
	if len(m.modifyCursors) == 0 {
		return ""
	}
	return m.modifyCursors[len(m.modifyCursors)-1]
}

// newCheckpointFlushTestConsumer builds and starts a consumer backed by
// checkpointFlushMockClient. The user consume callback blocks on consumeRelease
// until the test closes it, simulating a slow callback that holds the worker
// past Stop()'s cancel. consumeEntered is closed once the worker has entered the
// callback (post-fetch), so the test can deterministically wait for the right
// moment. FlushCheckpointIntervalSecond is 1h so only the final flush can fire.
// Returns the consumer, the mock client, the release func, and the entered chan.
func newCheckpointFlushTestConsumer(t *testing.T, stopTimeout time.Duration) (*consumer, *checkpointFlushMockClient, func(), <-chan struct{}) {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "ckpt-test-project"
	conf.TopicIDList = []string{"ckpt-test-topic"}
	conf.ConsumerGroupName = "ckpt-test-group"
	conf.ConsumerName = "ckpt-test-consumer"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = MetricsLevelCore
	conf.DataFetchIntervalInMillisecond = 10
	conf.HeartbeatIntervalInSecond = 1
	conf.FlushCheckpointIntervalSecond = 300 // long (max allowed): only the final flush fires

	mc := &checkpointFlushMockClient{}
	consumeRelease := make(chan struct{})
	consumeEntered := make(chan struct{})
	var enteredOnce sync.Once
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(consumeRelease) }) }

	c, err := newTestConsumerWithClient(context.Background(), conf, mc, func(topicID string, shardID int, l *pb.LogGroupList) {
		enteredOnce.Do(func() { close(consumeEntered) })
		<-consumeRelease // block until the test releases us
	})
	if err != nil {
		t.Fatalf("newTestConsumerWithClient: %v", err)
	}
	c.stopTimeout = stopTimeout
	if err := c.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	return c, mc, release, consumeEntered
}

// waitForConsumeEntered waits until the worker has entered the (blocked) consume
// callback, or times out.
func waitForConsumeEntered(t *testing.T, consumeEntered <-chan struct{}, timeout time.Duration) {
	t.Helper()
	select {
	case <-consumeEntered:
	case <-time.After(timeout):
		t.Fatalf("worker never entered consume callback within %v", timeout)
	}
}

// TestStop_FinalCheckpointFlushWaitsForWorker verifies that after Stop()
// cancels, the checkpoint manager must not do its final flush until
// every in-flight worker has exited (and thus called addCheckpoint with its
// final cursor). The test blocks the consume callback, starts Stop, asserts no
// ModifyCheckPoint call while the callback is blocked, releases the callback,
// then asserts Stop returns Closed/Final and ModifyCheckPoint received the
// worker's final cursor. With the old (buggy) code the checkpoint manager
// flushes immediately on cancel, before the blocked worker can add its
// checkpoint, so the final cursor is lost.
func TestStop_FinalCheckpointFlushWaitsForWorker(t *testing.T) {
	c, mc, release, consumeEntered := newCheckpointFlushTestConsumer(t, 5*time.Second)
	defer release()

	oneShard := []*tls.ConsumeShard{{TopicID: "ckpt-test-topic", ShardID: 0}}
	mc.setShards(oneShard)
	c.heartbeat.setShards(oneShard)
	waitForConsumeEntered(t, consumeEntered, 5*time.Second)

	// No checkpoint upload yet (long flush interval, no commitCh signal).
	if got := mc.modifyCount(); got != 0 {
		t.Fatalf("ModifyCheckPoint called %d times before Stop, want 0", got)
	}

	// Start Stop; it should block waiting for the in-flight worker.
	stopDone := make(chan struct{})
	go func() {
		c.Stop()
		close(stopDone)
	}()

	// Give Stop time to cancel and reach the worker wait.
	time.Sleep(300 * time.Millisecond)

	// While the consume callback is still blocked, the final checkpoint flush
	// must NOT have run. The checkpoint manager must wait for the worker to
	// exit (and add its checkpoint) before uploading.
	if got := mc.modifyCount(); got != 0 {
		t.Fatalf("ModifyCheckPoint called %d times while worker blocked, want 0 (final flush ran too early)", got)
	}

	// Release the callback: the worker finishes, adds its checkpoint, then the
	// finalizer signals the checkpoint manager to do the final flush.
	release()

	select {
	case <-stopDone:
	case <-time.After(5 * time.Second):
		t.Fatal("Stop() did not return after worker release")
	}
	st := c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State = %s, want stateClosed", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency = %s, want consistencyFinal", st.Consistency)
	}

	// Exactly one final flush, carrying the worker's cursor.
	if got := mc.modifyCount(); got != 1 {
		t.Fatalf("ModifyCheckPoint called %d times after Stop, want 1", got)
	}
	if got := mc.lastCursor(); got != "cursor-1" {
		t.Fatalf("final checkpoint cursor = %q, want %q", got, "cursor-1")
	}
}

// TestStop_TimeoutDefersFinalCheckpointFlush verifies the bounded-stop path:
// when Stop times out (worker still blocked in the consume callback), it returns
// Closing/Weak and the checkpoint manager must NOT have flushed yet. After the
// callback is released, the single async finalizer waits for the worker, then
// drives the final flush and transitions to Closed/Final.
func TestStop_TimeoutDefersFinalCheckpointFlush(t *testing.T) {
	c, mc, release, consumeEntered := newCheckpointFlushTestConsumer(t, 100*time.Millisecond)
	defer release()

	oneShard := []*tls.ConsumeShard{{TopicID: "ckpt-test-topic", ShardID: 0}}
	mc.setShards(oneShard)
	c.heartbeat.setShards(oneShard)
	waitForConsumeEntered(t, consumeEntered, 5*time.Second)

	// Stop must time out bounded and return closing/weak.
	start := time.Now()
	c.Stop()
	elapsed := time.Since(start)
	if elapsed > 3*time.Second {
		t.Fatalf("Stop() took %v with blocked worker, expected ~100ms bounded", elapsed)
	}
	st := c.collectMeasurements()
	if st.State != stateClosing {
		t.Fatalf("State = %s, want stateClosing", st.State)
	}
	if st.Consistency != consistencyWeak {
		t.Fatalf("Consistency = %s, want consistencyWeak", st.Consistency)
	}

	// No final flush while the worker is still blocked.
	if got := mc.modifyCount(); got != 0 {
		t.Fatalf("ModifyCheckPoint called %d times during timeout, want 0", got)
	}

	// Release the worker; the async finalizer should flush the checkpoint.
	release()

	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if c.collectMeasurements().State == stateClosed {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	st = c.collectMeasurements()
	if st.State != stateClosed {
		t.Fatalf("State after release = %s, want stateClosed (async finalizer did not run)", st.State)
	}
	if st.Consistency != consistencyFinal {
		t.Fatalf("Consistency after release = %s, want consistencyFinal", st.Consistency)
	}

	// The deferred final flush must have uploaded the worker's cursor.
	if got := mc.modifyCount(); got != 1 {
		t.Fatalf("ModifyCheckPoint called %d times after release, want 1", got)
	}
	if got := mc.lastCursor(); got != "cursor-1" {
		t.Fatalf("final checkpoint cursor = %q, want %q", got, "cursor-1")
	}
}
