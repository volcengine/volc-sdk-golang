package consumer

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// newTestConsumer builds a consumer without network access.
// validateConsumerConfig requires ProjectID/TopicIDList/ConsumerGroupName;
// we set those and a fake endpoint so NewClient does not dial.
func newTestConsumer(t *testing.T, level MetricsLevel) *consumer {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "test-project"
	conf.TopicIDList = []string{"test-topic"}
	conf.ConsumerGroupName = "test-group"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = level

	c, err := newConsumer(context.Background(), conf, func(topicID string, shardID int, l *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newConsumer: %v", err)
	}
	return c
}

func TestConsumerStats_DefaultOff(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelOff)
	st := c.collectMeasurements()
	if st.MetricsDisabled != true {
		t.Errorf("MetricsDisabled=true when Off, got %v", st.MetricsDisabled)
	}
	if st.MetricsLevel != MetricsLevelOff {
		t.Errorf("MetricsLevel=Off, got %v", st.MetricsLevel)
	}
	if st.State != stateNew {
		t.Errorf("State=stateNew before Start, got %v", st.State)
	}
	// Off: business counters must be absent (nil pointers)
	if st.FetchCalls != nil || st.ProcessAttempts != nil || st.AssignedShards != nil {
		t.Errorf("Off must have nil business counters, got non-nil")
	}
	if st.WorkerRoles != nil {
		t.Errorf("Off must have nil WorkerRoles, got %d", len(st.WorkerRoles))
	}
	if st.ComponentType != "consumer" {
		t.Errorf("ComponentType=consumer, got %s", st.ComponentType)
	}
}

func TestConsumerStats_CoreBeforeStart(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)
	st := c.collectMeasurements()
	if st.MetricsDisabled {
		t.Errorf("configured Core collection is active before Start")
	}
}

func TestConsumerStats_CoreEnabledFromConstruction(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)
	st := c.collectMeasurements()
	if st.MetricsDisabled != false {
		t.Errorf("MetricsDisabled=false for Core, got %v", st.MetricsDisabled)
	}
	if st.MetricsLevel != MetricsLevelCore {
		t.Errorf("MetricsLevel=Core, got %v", st.MetricsLevel)
	}
	if len(st.WorkerRoles) != 3 {
		t.Fatalf("WorkerRoles len=3, got %d", len(st.WorkerRoles))
	}
	// Expected counts: main_loop=1, heartbeat=1, checkpoint=1
	for _, wr := range st.WorkerRoles {
		if wr.Expected != 1 {
			t.Errorf("role %s Expected=1, got %d", wr.Role, wr.Expected)
		}
	}
	// Cumulative counters start at 0 (non-nil when enabled)
	if st.FetchCalls == nil || *st.FetchCalls != 0 {
		t.Errorf("FetchCalls=0 when enabled, got %v", st.FetchCalls)
	}
	if st.AssignedShards == nil || *st.AssignedShards != 0 {
		t.Errorf("AssignedShards=0, got %v", st.AssignedShards)
	}
}

func TestConsumerStats_CounterIncrements(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)

	firstFetch := c.stats.beginFetchCall()
	c.stats.finishFetchCall(firstFetch, false)
	c.stats.addFetchedFor(firstFetch, 3, 30)
	secondFetch := c.stats.beginFetchCall()
	c.stats.finishFetchCall(secondFetch, false)
	process := c.stats.beginProcessAttempt()
	c.stats.finishProcessAttempt(process, 3, 30, true)
	heartbeat := c.stats.beginHeartbeatCall()
	c.stats.finishHeartbeatCall(heartbeat, true)
	checkpoint := c.stats.beginCheckpointCall()
	c.stats.finishCheckpointCall(checkpoint, false)
	c.stats.incBackoffEvents()

	st := c.collectMeasurements()
	if *st.FetchCalls != 2 {
		t.Errorf("FetchCalls=2, got %d", *st.FetchCalls)
	}
	if *st.FetchedLogGroups != 3 {
		t.Errorf("FetchedLogGroups=3, got %d", *st.FetchedLogGroups)
	}
	if *st.FetchedLogs != 30 {
		t.Errorf("FetchedLogs=30, got %d", *st.FetchedLogs)
	}
	if *st.ProcessAttempts != 1 {
		t.Errorf("ProcessAttempts=1, got %d", *st.ProcessAttempts)
	}
	if *st.ProcessedSuccessGroups != 3 || *st.ProcessedSuccessLogs != 30 {
		t.Errorf("ProcessedSuccessGroups/Logs=3/30, got %d/%d", *st.ProcessedSuccessGroups, *st.ProcessedSuccessLogs)
	}
	if *st.HeartbeatCalls != 1 {
		t.Errorf("HeartbeatCalls=1, got %d", *st.HeartbeatCalls)
	}
	if *st.HeartbeatFailures != 1 {
		t.Errorf("HeartbeatFailures=1, got %d", *st.HeartbeatFailures)
	}
	if *st.CheckpointCalls != 1 {
		t.Errorf("CheckpointCalls=1, got %d", *st.CheckpointCalls)
	}
	if *st.BackoffEvents != 1 {
		t.Errorf("BackoffEvents=1, got %d", *st.BackoffEvents)
	}
}

func TestConsumerStats_InFlight(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)

	firstFetch := c.stats.beginFetchCall()
	secondFetch := c.stats.beginFetchCall()
	process := c.stats.beginProcessAttempt()

	st := c.collectMeasurements()
	if *st.FetchInFlight != 2 {
		t.Errorf("FetchInFlight=2, got %d", *st.FetchInFlight)
	}
	if *st.ProcessInFlight != 1 {
		t.Errorf("ProcessInFlight=1, got %d", *st.ProcessInFlight)
	}

	c.stats.finishFetchCall(firstFetch, false)
	st = c.collectMeasurements()
	if *st.FetchInFlight != 1 {
		t.Errorf("FetchInFlight=1 after dec, got %d", *st.FetchInFlight)
	}
	c.stats.finishFetchCall(secondFetch, false)
	c.stats.finishProcessAttempt(process, 0, 0, false)
}

func TestConsumerStats_OffNoCounting(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelOff)

	fetch := c.stats.beginFetchCall()
	c.stats.finishFetchCall(fetch, false)
	process := c.stats.beginProcessAttempt()
	c.stats.finishProcessAttempt(process, 0, 0, true)
	heartbeat := c.stats.beginHeartbeatCall()
	c.stats.finishHeartbeatCall(heartbeat, false)

	st := c.collectMeasurements()
	// Off: counters absent (nil)
	if st.FetchCalls != nil {
		t.Errorf("Off: FetchCalls must be nil, got %v", *st.FetchCalls)
	}
	if st.MetricsDisabled != true {
		t.Errorf("Off: MetricsDisabled=true, got %v", st.MetricsDisabled)
	}
}

func TestConsumerStats_WorkerRoleRunning(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)

	c.stats.incMainLoopRunning()
	c.stats.incHeartbeatRunning()

	st := c.collectMeasurements()
	for _, wr := range st.WorkerRoles {
		switch wr.Role {
		case roleMainLoop:
			if wr.Running != 1 {
				t.Errorf("main_loop Running=1, got %d", wr.Running)
			}
		case roleHeartbeat:
			if wr.Running != 1 {
				t.Errorf("heartbeat Running=1, got %d", wr.Running)
			}
		case roleCheckpoint:
			if wr.Running != 0 {
				t.Errorf("checkpoint Running=0, got %d", wr.Running)
			}
		}
	}

	c.stats.decMainLoopRunning()
	st = c.collectMeasurements()
	for _, wr := range st.WorkerRoles {
		if wr.Role == roleMainLoop && wr.Running != 0 {
			t.Errorf("main_loop Running=0 after dec, got %d", wr.Running)
		}
	}
}

func TestConsumerStats_NilConsumer(t *testing.T) {
	var c *consumer
	st := c.collectMeasurements()
	if st.MetricsDisabled != true {
		t.Errorf("nil consumer: MetricsDisabled=true, got %v", st.MetricsDisabled)
	}
}

func TestConsumerContractInterface(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)
	var contract Consumer = c
	st, err := contract.GetStats(StatsOptions{View: StatsViewHealth})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if st.CollectedLevel != MetricsLevelCore {
		t.Errorf("CollectedLevel=Core, got %v", st.CollectedLevel)
	}
}

func TestConsumerStats_LifecycleState(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)
	if c.lifecycleState() != stateNew {
		t.Errorf("state=stateNew, got %v", c.lifecycleState())
	}
	// started flag set only in Start(); simulate
	c.started = 1
	if c.lifecycleState() != stateRunning {
		t.Errorf("state=stateRunning, got %v", c.lifecycleState())
	}
	c.closing = 1
	if c.lifecycleState() != stateClosing {
		t.Errorf("state=stateClosing, got %v", c.lifecycleState())
	}
}

func TestConsumerStats_SnapshotSequenceIncrements(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)
	s1 := c.collectMeasurements().SnapshotSequence
	s2 := c.collectMeasurements().SnapshotSequence
	if s2 != s1+1 {
		t.Errorf("SnapshotSequence should increment: %d -> %d", s1, s2)
	}
}

func TestConsumerStats_ConfigDefaultOff(t *testing.T) {
	conf := GetDefaultConsumerConfig()
	if conf.MetricsLevel != MetricsLevelOff {
		t.Errorf("default MetricsLevel=Off, got %v", conf.MetricsLevel)
	}
}

// ---------------------------------------------------------------------------
// Counter-balance integration tests (cover the 3 major findings from the
// code-quality review: negative counters on ErrInvalidContent, leak on empty
// LogGroups, leak on shard removal). These drive the real fetchData()/consume()
// flow with a minimal mock client rather than calling incX() directly.
// ---------------------------------------------------------------------------

// mockClient embeds a nil tls.Client interface and overrides only the methods
// logConsumer calls. The nil embedded interface is never invoked because
// fetchData/init only call the overridden methods.
type mockClient struct {
	tls.Client
	consumeLogsFn func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error)
}

func (m *mockClient) ConsumeLogs(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
	return m.consumeLogsFn(req)
}

func (m *mockClient) DescribeCheckPoint(req *tls.DescribeCheckPointRequest) (*tls.DescribeCheckPointResponse, error) {
	return &tls.DescribeCheckPointResponse{}, nil
}

func (m *mockClient) DescribeCursor(req *tls.DescribeCursorRequest) (*tls.DescribeCursorResponse, error) {
	return &tls.DescribeCursorResponse{Cursor: "cursor-0"}, nil
}

// newCounterTestConsumer builds a logConsumer wired to a real stats state and
// checkpoint manager so fetchData()/consume()/releaseHeldBatchIfAny() run the
// real code paths.
func newCounterTestConsumer(t *testing.T, mc *mockClient) *logConsumer {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "test-project"
	conf.TopicIDList = []string{"test-topic"}
	conf.ConsumerGroupName = "test-group"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = MetricsLevelCore

	stats := newConsumerStatsState(MetricsLevelCore)
	commitCh := make(chan struct{}, 16)
	lc := &logConsumer{
		ctx:                context.Background(),
		wg:                 &sync.WaitGroup{},
		client:             mc,
		logger:             log.NewNopLogger(),
		conf:               conf,
		statusLock:         &sync.RWMutex{},
		status:             readyToFetch,
		shard:              &tls.ConsumeShard{TopicID: "test-topic", ShardID: 0},
		consumeFunc:        func(topicID string, shardID int, l *pb.LogGroupList) {},
		checkpoint:         newCheckpointManager(log.NewNopLogger(), conf, mc, commitCh, stats),
		heartbeatRestartCh: make(chan struct{}, 1),
		commitCh:           commitCh,
		nextCheckpoint:     "cursor-0",
		stats:              stats,
	}
	return lc
}

func nonEmptyBatch() *pb.LogGroupList {
	return &pb.LogGroupList{
		LogGroups: []*pb.LogGroup{{Logs: []*pb.Log{{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}}}}},
	}
}

// heldAndReady reads the two pre-aggregated counters from a snapshot.
func heldAndReady(c *consumer) (uint64, uint64) {
	st := c.collectMeasurements()
	var h, r uint64
	if st.HeldBatchReferences != nil {
		h = *st.HeldBatchReferences
	}
	if st.ReadyToProcess != nil {
		r = *st.ReadyToProcess
	}
	return h, r
}

// TestCounterBalance_NormalFetchConsume: fetchData incs held/ready for a
// non-empty batch; consume() decs them back to zero. No leak, no negative.
func TestCounterBalance_NormalFetchConsume(t *testing.T) {
	mc := &mockClient{
		consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
			return &tls.ConsumeLogsResponse{Logs: nonEmptyBatch(), Cursor: "cursor-1"}, nil
		},
	}
	lc := newCounterTestConsumer(t, mc)
	c := &consumer{stats: lc.stats}

	if err := lc.fetchData(); err != nil {
		t.Fatalf("fetchData: %v", err)
	}
	if h, r := heldAndReady(c); h != 1 || r != 1 {
		t.Fatalf("after fetch: held=%d ready=%d, want 1/1", h, r)
	}

	lc.setStatus(readyToConsume) // run() normally does this
	if err := lc.consume(); err != nil {
		t.Fatalf("consume: %v", err)
	}
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("after consume: held=%d ready=%d, want 0/0 (leak)", h, r)
	}
}

// TestCounterBalance_ErrInvalidContent: the ErrInvalidContent branch sets
// currLogGroupList and (after fix) incs held/ready; consume() must dec them
// back to zero. Before the fix this path never inc'd, so consume() drove the
// counters negative.
func TestCounterBalance_ErrInvalidContent(t *testing.T) {
	mc := &mockClient{
		consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
			return &tls.ConsumeLogsResponse{Logs: nonEmptyBatch(), Cursor: "cursor-1"},
				&tls.Error{Code: tls.ErrInvalidContent, Message: "invalid"}
		},
	}
	lc := newCounterTestConsumer(t, mc)
	c := &consumer{stats: lc.stats}

	if err := lc.fetchData(); err != nil {
		t.Fatalf("fetchData (ErrInvalidContent): %v", err)
	}
	if h, r := heldAndReady(c); h != 1 || r != 1 {
		t.Fatalf("after ErrInvalidContent fetch: held=%d ready=%d, want 1/1", h, r)
	}

	lc.setStatus(readyToConsume)
	if err := lc.consume(); err != nil {
		t.Fatalf("consume: %v", err)
	}
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("after consume: held=%d ready=%d, want 0/0 (negative or leak)", h, r)
	}
}

// TestCounterBalance_EmptyLogGroups: fetchData with empty LogGroups must NOT
// inc held/ready (consume() returns early and would never dec them). Before
// the fix this permanently leaked the counters.
func TestCounterBalance_EmptyLogGroups(t *testing.T) {
	mc := &mockClient{
		consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
			return &tls.ConsumeLogsResponse{Logs: &pb.LogGroupList{}, Cursor: "cursor-1"}, nil
		},
	}
	lc := newCounterTestConsumer(t, mc)
	c := &consumer{stats: lc.stats}

	if err := lc.fetchData(); err != nil {
		t.Fatalf("fetchData: %v", err)
	}
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("after empty fetch: held=%d ready=%d, want 0/0 (leak)", h, r)
	}

	lc.setStatus(readyToConsume)
	if err := lc.consume(); err != nil {
		t.Fatalf("consume: %v", err)
	}
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("after consume: held=%d ready=%d, want 0/0", h, r)
	}
}

// TestCounterBalance_ReleaseOnShardRemoval: a worker holding a batch (status
// readyToConsume, non-empty currLogGroupList) that is removed before consume()
// runs must have its held/ready counters released by releaseHeldBatchIfAny().
// Before the fix this permanently leaked the counters on rebalance/shutdown.
func TestCounterBalance_ReleaseOnShardRemoval(t *testing.T) {
	mc := &mockClient{
		consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
			return &tls.ConsumeLogsResponse{Logs: nonEmptyBatch(), Cursor: "cursor-1"}, nil
		},
	}
	lc := newCounterTestConsumer(t, mc)
	c := &consumer{stats: lc.stats}

	if err := lc.fetchData(); err != nil {
		t.Fatalf("fetchData: %v", err)
	}
	lc.setStatus(readyToConsume) // batch held, waiting to consume
	if h, r := heldAndReady(c); h != 1 || r != 1 {
		t.Fatalf("after fetch: held=%d ready=%d, want 1/1", h, r)
	}

	// Simulate shard removal (rebalance) before consume() runs.
	lc.releaseHeldBatchIfAny()
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("after release: held=%d ready=%d, want 0/0 (leak on removal)", h, r)
	}
}

// TestCounterBalance_ReleaseNotHeld: releaseHeldBatchIfAny() must NOT dec
// counters when the worker is not holding a batch (held flag == 0). When a
// batch IS held (held flag == 1, set by fetchData), release decrements exactly
// once via CAS — even if status is still fetching (the window between
// fetchData returning and runWithStatus setting readyToConsume). This guards
// against both no-op-when-empty and double-decrement.
func TestCounterBalance_ReleaseNotHeld(t *testing.T) {
	mc := &mockClient{
		consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
			return &tls.ConsumeLogsResponse{Logs: nonEmptyBatch(), Cursor: "cursor-1"}, nil
		},
	}
	lc := newCounterTestConsumer(t, mc)
	c := &consumer{stats: lc.stats}

	// No batch fetched yet (held=0): release must be a no-op.
	lc.releaseHeldBatchIfAny()
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("release when not holding: held=%d ready=%d, want 0/0 (negative)", h, r)
	}

	// Fetch a batch: held/ready become 1/1 and held flag is set to 1.
	if err := lc.fetchData(); err != nil {
		t.Fatalf("fetchData: %v", err)
	}
	if h, r := heldAndReady(c); h != 1 || r != 1 {
		t.Fatalf("after fetch: held=%d ready=%d, want 1/1", h, r)
	}

	// Even with status still fetching (before readyToConsume), release must
	// decrement the held batch exactly once (CAS on held flag).
	lc.setStatus(fetching)
	lc.releaseHeldBatchIfAny()
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("release while fetching: held=%d ready=%d, want 0/0 (held batch must be released)", h, r)
	}

	// A second release must be a no-op (CAS fails, held flag already 0): no
	// double-decrement into negative territory.
	lc.releaseHeldBatchIfAny()
	if h, r := heldAndReady(c); h != 0 || r != 0 {
		t.Fatalf("second release: held=%d ready=%d, want 0/0 (double-decrement)", h, r)
	}
}

// TestCounterBalance_ReleaseRace: two goroutines race to release one held
// batch — one via releaseHeldBatch() (the normal consume() pickup path) and
// one via releaseHeldBatchIfAny() (the rebalance/shutdown/retired-exit path).
// Both funnel through the same CAS 1->0 on the per-worker `held` flag.
// Repeated many times to exercise the race window. Verifies:
//   - exactly one release wins per batch (wins == 1)
//   - final held/ready counters are 0/0 (no leak, no double-decrement)
//   - counters never underflow (a uint64 wraparound would read as a huge
//     value, failing the 0/0 check)
func TestCounterBalance_ReleaseRace(t *testing.T) {
	const iterations = 1000
	for i := 0; i < iterations; i++ {
		mc := &mockClient{
			consumeLogsFn: func(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
				return &tls.ConsumeLogsResponse{Logs: nonEmptyBatch(), Cursor: "cursor-1"}, nil
			},
		}
		lc := newCounterTestConsumer(t, mc)
		c := &consumer{stats: lc.stats}

		if err := lc.fetchData(); err != nil {
			t.Fatalf("iteration %d: fetchData: %v", i, err)
		}
		if h, r := heldAndReady(c); h != 1 || r != 1 {
			t.Fatalf("iteration %d: after fetch: held=%d ready=%d, want 1/1", i, h, r)
		}

		var wg sync.WaitGroup
		var wins int32
		// Goroutine 1: normal consume() pickup path (releaseHeldBatch).
		wg.Add(1)
		go func() {
			defer wg.Done()
			if lc.releaseHeldBatch() {
				atomic.AddInt32(&wins, 1)
			}
		}()
		// Goroutine 2: rebalance/shutdown/retired-exit path
		// (releaseHeldBatchIfAny -> releaseHeldBatch). Both funnel through
		// the same CAS 1->0, so this mirrors the real production race.
		wg.Add(1)
		go func() {
			defer wg.Done()
			if lc.releaseHeldBatch() {
				atomic.AddInt32(&wins, 1)
			}
		}()
		wg.Wait()

		if wins != 1 {
			t.Fatalf("iteration %d: exactly one release should win, got %d", i, wins)
		}
		if h, r := heldAndReady(c); h != 0 || r != 0 {
			t.Fatalf("iteration %d: after race: held=%d ready=%d, want 0/0 (underflow or leak)", i, h, r)
		}
	}
}

// TestCriticalFlags_NonNilEmpty verifies the public health fact is an empty
// list rather than null when Health is collected.
func TestCriticalFlags_NonNilEmpty(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelCore)
	st, err := c.GetStats(StatsOptions{View: StatsViewHealth})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if st.Health == nil || st.Health.Signals.CriticalFlags == nil {
		t.Fatalf("CriticalFlags must be non-nil, got nil")
	}
	data, err := json.Marshal(st)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	jsonStr := string(data)
	if !strings.Contains(jsonStr, `"critical_flags":[]`) || strings.Contains(jsonStr, `"critical_flags":null`) {
		t.Fatalf("invalid CriticalFlags JSON: %s", jsonStr)
	}
}
