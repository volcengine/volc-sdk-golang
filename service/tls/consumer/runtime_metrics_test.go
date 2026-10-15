package consumer

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/tls/observability"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func newRuntimeMetricsTestConsumer(t *testing.T) *consumer {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "test-project"
	conf.TopicIDList = []string{"test-topic"}
	conf.ConsumerGroupName = "test-group"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.Region = "cn-beijing"
	c, err := newConsumer(context.Background(), conf, func(topicID string, shardID int, l *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newConsumer: %v", err)
	}
	// The setter deliberately only works for a running instance. The unit test
	// does not need network goroutines, so model that lifecycle state directly.
	atomic.StoreInt32(&c.lifecycleStage, lifecycleRunning)
	return c
}

func newRuntimeMetricsStartedConsumer(t *testing.T, parent context.Context, level MetricsLevel) *consumer {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "runtime-role-project"
	conf.TopicIDList = []string{"runtime-role-topic"}
	conf.ConsumerGroupName = "runtime-role-group"
	conf.ConsumerName = "runtime-role-consumer"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = level
	// Keep the scheduler and periodic network loops out of this lifecycle
	// probe. Start still launches all three real component goroutines, and the
	// first heartbeat call completes synchronously against the mock client.
	conf.DataFetchIntervalInMillisecond = 300000
	conf.HeartbeatIntervalInSecond = 300
	conf.FlushCheckpointIntervalSecond = 300

	c, err := newTestConsumerWithClient(parent, conf, &mockConsumerClient{}, nil)
	if err != nil {
		t.Fatalf("newTestConsumerWithClient: %v", err)
	}
	c.stopTimeout = 2 * time.Second
	if err := c.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	t.Cleanup(c.Stop)
	return c
}

func waitForRuntimeRoleCounts(t *testing.T, c *consumer, main, heartbeat, checkpoint int32) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if atomic.LoadInt32(&c.stats.mainLoopRunning) == main &&
			atomic.LoadInt32(&c.stats.heartbeatRunning) == heartbeat &&
			atomic.LoadInt32(&c.stats.checkpointRunning) == checkpoint {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("role counts = main:%d heartbeat:%d checkpoint:%d, want main:%d heartbeat:%d checkpoint:%d",
		atomic.LoadInt32(&c.stats.mainLoopRunning), atomic.LoadInt32(&c.stats.heartbeatRunning),
		atomic.LoadInt32(&c.stats.checkpointRunning), main, heartbeat, checkpoint)
}

func assertRuntimeRoleSummary(t *testing.T, snapshot StatsSnapshot, main, heartbeat, checkpoint uint32) {
	t.Helper()
	if snapshot.Operational == nil {
		t.Fatalf("operational section is nil: %+v", snapshot)
	}
	running := main + heartbeat + checkpoint
	assert.Equal(t, uint32(3), snapshot.Operational.Workers.Expected)
	assert.Equal(t, running, snapshot.Operational.Workers.Running)
	assert.Equal(t, uint32(3)-running, snapshot.Operational.Workers.MissingCriticalRoles)
	if snapshot.WorkerDetails == nil {
		t.Fatalf("worker details are nil: %+v", snapshot)
	}
	if len(snapshot.WorkerDetails.Items) != 3 {
		t.Fatalf("worker details = %+v, want 3 roles", snapshot.WorkerDetails.Items)
	}
	for _, role := range snapshot.WorkerDetails.Items {
		assert.Equal(t, uint32(1), role.Expected, "role=%s", role.Role)
		switch role.Role {
		case string(roleMainLoop):
			assert.Equal(t, main, role.Running, "role=%s", role.Role)
		case string(roleHeartbeat):
			assert.Equal(t, heartbeat, role.Running, "role=%s", role.Role)
		case string(roleCheckpoint):
			assert.Equal(t, checkpoint, role.Running, "role=%s", role.Role)
		}
	}
}

func waitForRuntimeMainLoopDone(t *testing.T, c *consumer) {
	t.Helper()
	select {
	case <-c.mainLoopDone:
	case <-time.After(2 * time.Second):
		t.Fatal("main loop did not exit after parent cancellation")
	}
}

func assertRuntimeMetricsHealthMissingRoles(t *testing.T, c *consumer, wantRunning uint32) {
	t.Helper()
	first, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("first GetStats(Debug): %v", err)
	}
	second, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("second GetStats(Debug): %v", err)
	}
	assert.Equal(t, first.CollectionGeneration, second.CollectionGeneration)
	assertRuntimeRoleSummary(t, second, 0, 0, wantRunning)
	health := EvaluateHealth(&first, second, observability.DefaultHealthPolicy())
	if health.Status != observability.HealthUnhealthy || health.InsufficientData ||
		len(health.ReasonCodes) != 1 || health.ReasonCodes[0] != observability.ReasonWorkerMissing {
		t.Fatalf("consecutive same-generation health = %+v, want unhealthy/worker_missing", health)
	}
}

func TestConsumerRuntimeMetricsRolesSurviveOffOnAndStop(t *testing.T) {
	c := newRuntimeMetricsStartedConsumer(t, context.Background(), MetricsLevelOff)
	waitForRuntimeRoleCounts(t, c, 1, 1, 1)

	off, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats(Off): %v", err)
	}
	if off.Health != nil || off.Operational != nil || off.WorkerDetails != nil || !off.MetricsDisabled {
		t.Fatalf("Off must remain header-only: %+v", off)
	}

	if err := c.SetMetricsLevel(MetricsLevelDetailed); err != nil {
		t.Fatalf("SetMetricsLevel(Detailed): %v", err)
	}
	on, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats(Detailed): %v", err)
	}
	assertRuntimeRoleSummary(t, on, 1, 1, 1)

	if err := c.SetMetricsLevel(MetricsLevelOff); err != nil {
		t.Fatalf("SetMetricsLevel(Off): %v", err)
	}
	waitForRuntimeRoleCounts(t, c, 1, 1, 1)
	if err := c.SetMetricsLevel(MetricsLevelDetailed); err != nil {
		t.Fatalf("SetMetricsLevel(Detailed) second time: %v", err)
	}
	on, err = c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats(Detailed) second time: %v", err)
	}
	assertRuntimeRoleSummary(t, on, 1, 1, 1)

	c.Stop()
	waitForRuntimeRoleCounts(t, c, 0, 0, 0)
}

func TestConsumerRuntimeMetricsEnableDoesNotDuplicateStartingRoles(t *testing.T) {
	c := newRuntimeMetricsStartedConsumer(t, context.Background(), MetricsLevelOff)
	// Enable immediately after Start, before waiting for any role entry. The
	// role atomics must be owned solely by the real goroutine entry/exit pairs.
	if err := c.SetMetricsLevel(MetricsLevelDetailed); err != nil {
		t.Fatalf("SetMetricsLevel(Detailed): %v", err)
	}
	waitForRuntimeRoleCounts(t, c, 1, 1, 1)
	on, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats(Detailed): %v", err)
	}
	assertRuntimeRoleSummary(t, on, 1, 1, 1)

	c.Stop()
	waitForRuntimeRoleCounts(t, c, 0, 0, 0)
}

func TestConsumerRuntimeMetricsRolesAfterParentCancelWhileOff(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()
	c := newRuntimeMetricsStartedConsumer(t, parent, MetricsLevelOff)
	waitForRuntimeRoleCounts(t, c, 1, 1, 1)

	cancelParent()
	waitForRuntimeMainLoopDone(t, c)
	waitForRuntimeRoleCounts(t, c, 0, 0, 1)
	if got := atomic.LoadInt32(&c.lifecycleStage); got != lifecycleRunning {
		t.Fatalf("lifecycle stage after parent cancellation = %d, want running", got)
	}

	if err := c.SetMetricsLevel(MetricsLevelDetailed); err != nil {
		t.Fatalf("SetMetricsLevel(Detailed): %v", err)
	}
	assertRuntimeMetricsHealthMissingRoles(t, c, 1)

	c.Stop()
	waitForRuntimeRoleCounts(t, c, 0, 0, 0)
}

func TestConsumerRuntimeMetricsRolesAfterParentCancelDuringOffWindow(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()
	c := newRuntimeMetricsStartedConsumer(t, parent, MetricsLevelCore)
	waitForRuntimeRoleCounts(t, c, 1, 1, 1)

	if err := c.SetMetricsLevel(MetricsLevelOff); err != nil {
		t.Fatalf("SetMetricsLevel(Off): %v", err)
	}
	cancelParent()
	waitForRuntimeMainLoopDone(t, c)
	waitForRuntimeRoleCounts(t, c, 0, 0, 1)

	if err := c.SetMetricsLevel(MetricsLevelDetailed); err != nil {
		t.Fatalf("SetMetricsLevel(Detailed): %v", err)
	}
	assertRuntimeMetricsHealthMissingRoles(t, c, 1)

	c.Stop()
	waitForRuntimeRoleCounts(t, c, 0, 0, 0)
}

func TestConsumerRuntimeMetricsRolesResetAcrossRestart(t *testing.T) {
	c := newRuntimeMetricsStartedConsumer(t, context.Background(), MetricsLevelOff)
	waitForRuntimeRoleCounts(t, c, 1, 1, 1)

	c.Stop()
	waitForRuntimeRoleCounts(t, c, 0, 0, 0)
	if err := c.Start(); err != nil {
		t.Fatalf("restart Start: %v", err)
	}
	waitForRuntimeRoleCounts(t, c, 1, 1, 1)
	if err := c.SetMetricsLevel(MetricsLevelDetailed); err != nil {
		t.Fatalf("SetMetricsLevel(Detailed) after restart: %v", err)
	}
	on, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats(Detailed) after restart: %v", err)
	}
	assertRuntimeRoleSummary(t, on, 1, 1, 1)

	c.Stop()
	waitForRuntimeRoleCounts(t, c, 0, 0, 0)
}

func TestConsumerRuntimeMetricsLevelSwitch(t *testing.T) {
	c := newRuntimeMetricsTestConsumer(t)
	assert.NotNil(t, c.stats, "runtime level changes require retained component state")
	assert.True(t, c.collectMeasurements().MetricsDisabled)

	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	stats := c.collectMeasurements()
	assert.Equal(t, MetricsLevelCore, stats.MetricsLevel)
	assert.False(t, stats.MetricsDisabled)
	assert.NotNil(t, stats.FetchCalls)

	assert.NoError(t, c.SetMetricsLevel(MetricsLevelOff))
	stats = c.collectMeasurements()
	assert.Equal(t, MetricsLevelOff, stats.MetricsLevel)
	assert.True(t, stats.MetricsDisabled)
	assert.Nil(t, stats.FetchCalls)
}

func TestConsumerRuntimeMetricsLevelStartsGaugesFreshAfterOff(t *testing.T) {
	c := newRuntimeMetricsTestConsumer(t)
	fetch := c.stats.beginFetchCall()
	process := c.stats.beginProcessAttempt()
	c.stats.incHeldBatchRefs()
	c.stats.incReadyToProcess()
	c.stats.finishFetchCall(fetch, false)
	c.stats.finishProcessAttempt(process, 0, 0, true)
	c.stats.decHeldBatchRefs()
	c.stats.decReadyToProcess()

	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	stats := c.collectMeasurements()
	assert.Equal(t, uint32(0), *stats.FetchInFlight)
	assert.Equal(t, uint32(0), *stats.ProcessInFlight)
	assert.Equal(t, uint64(0), *stats.HeldBatchReferences)
	assert.Equal(t, uint64(0), *stats.ReadyToProcess)
}

func TestConsumerRuntimeMetricsLevelResetsAndReinitializesCheckpointQueue(t *testing.T) {
	c := newRuntimeMetricsTestConsumer(t)
	c.checkpoint.checkpointMap["first"] = &checkpointInfo{}
	c.checkpoint.checkpointMap["second"] = &checkpointInfo{}

	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	queue := c.checkpoint.snapshot()
	assert.Equal(t, uint64(2), queue.CurrentItems)
	assert.Equal(t, uint64(2), queue.HighWatermarkItems)

	// A peak already observed after level publication must not be overwritten
	// by the current-depth initialization step.
	atomic.StoreInt64(&c.checkpoint.highWatermark, 3)
	c.checkpoint.observeCurrentHighWatermark()
	assert.Equal(t, uint64(3), c.checkpoint.snapshot().HighWatermarkItems)

	assert.NoError(t, c.SetMetricsLevel(MetricsLevelOff))
	queue = c.checkpoint.snapshot()
	assert.Zero(t, queue.CurrentItems)
	assert.Zero(t, queue.HighWatermarkItems)

	c.checkpoint.checkpointMap["third"] = &checkpointInfo{}
	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	queue = c.checkpoint.snapshot()
	assert.Equal(t, uint64(3), queue.CurrentItems)
	assert.Equal(t, uint64(3), queue.HighWatermarkItems)
}

func TestConsumerRuntimeMetricsLevelResetsHealthAges(t *testing.T) {
	c := newRuntimeMetricsTestConsumer(t)
	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	atomic.StoreInt64(&c.stats.lastSuccessTimeUnixMs, unixMillis(time.Now().Add(-time.Second)))
	atomic.StoreInt64(&c.stats.lastErrorTimeUnixMs, unixMillis(time.Now().Add(-time.Second)))
	before := c.collectMeasurements()
	assert.NotNil(t, before.LastSuccessAgeMs)
	assert.NotNil(t, before.LastErrorAgeMs)

	assert.NoError(t, c.SetMetricsLevel(MetricsLevelOff))
	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	after := c.collectMeasurements()
	assert.Nil(t, after.LastSuccessAgeMs)
	assert.Nil(t, after.LastErrorAgeMs)
}

func TestConsumerRuntimeMetricsLevelRequiresRunning(t *testing.T) {
	c := newRuntimeMetricsTestConsumer(t)
	atomic.StoreInt32(&c.lifecycleStage, lifecycleNew)
	assert.ErrorIs(t, c.SetMetricsLevel(MetricsLevelCore), ErrMetricsLevelChangeNotRunning)
	assert.ErrorIs(t, c.SetMetricsLevel(MetricsLevel(99)), ErrInvalidMetricsLevel)
}

func TestConsumerCoreDetailedSwitchPreservesCollectionGeneration(t *testing.T) {
	c := newRuntimeMetricsTestConsumer(t)
	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	core, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	assert.NoError(t, err)
	assert.NoError(t, c.SetMetricsLevel(MetricsLevelDetailed))
	detailed, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	assert.NoError(t, err)
	assert.Equal(t, core.CollectionGeneration, detailed.CollectionGeneration)
	assert.Equal(t, core.CollectionStartedAtUnixMs, detailed.CollectionStartedAtUnixMs)
	assert.NotNil(t, detailed.LatencyHistograms)
	firstDetailGeneration := c.stats.currentDetailedGeneration()
	assert.NotZero(t, firstDetailGeneration)
	assert.NoError(t, c.SetMetricsLevel(MetricsLevelCore))
	assert.False(t, c.stats.detailedGenerationEnabled(firstDetailGeneration), "an interval crossing a Core window must be discarded")
	assert.NoError(t, c.SetMetricsLevel(MetricsLevelDetailed))
	assert.Greater(t, c.stats.currentDetailedGeneration(), firstDetailGeneration)
	after, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	assert.NoError(t, err)
	assert.Equal(t, core.CollectionGeneration, after.CollectionGeneration)
}

func TestConsumerRuntimeMetricsLevelConcurrentWithSnapshot(t *testing.T) {
	c := newRuntimeMetricsTestConsumer(t)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for n := 0; n < 200; n++ {
				level := MetricsLevelOff
				if (n+i)%2 == 0 {
					level = MetricsLevelCore
				}
				if err := c.SetMetricsLevel(level); err != nil {
					t.Errorf("SetMetricsLevel(%v): %v", level, err)
					return
				}
			}
		}(i)
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < 200; n++ {
				snapshot, err := c.GetStats(StatsOptions{View: StatsViewOperational})
				if err != nil {
					t.Errorf("GetStats: %v", err)
					return
				}
				if snapshot.CollectedLevel == MetricsLevelOff {
					if !snapshot.MetricsDisabled || snapshot.CollectionStartedAtUnixMs != 0 || snapshot.Health != nil || snapshot.Operational != nil {
						t.Errorf("inconsistent Off snapshot: %+v", snapshot)
						return
					}
				} else if snapshot.MetricsDisabled || snapshot.CollectionStartedAtUnixMs <= 0 || snapshot.Health == nil || snapshot.Operational == nil {
					t.Errorf("inconsistent enabled snapshot: %+v", snapshot)
					return
				}
			}
		}()
	}
	wg.Wait()
	assert.Greater(t, atomic.LoadUint64(&c.stats.snapshotSequence), uint64(0))
}
