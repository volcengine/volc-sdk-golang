package consumer

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/observability"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestContractGetStatsDoesNotWaitForCheckpointMapLock(t *testing.T) {
	c := newContractTestConsumer(t, MetricsLevelCore)
	c.checkpoint.mapLock.Lock()
	done := make(chan error, 1)
	go func() {
		_, err := c.GetStats(StatsOptions{View: StatsViewOperational})
		done <- err
	}()
	select {
	case err := <-done:
		c.checkpoint.mapLock.Unlock()
		if err != nil {
			t.Fatalf("GetStats: %v", err)
		}
	case <-time.After(100 * time.Millisecond):
		c.checkpoint.mapLock.Unlock()
		<-done
		t.Fatal("GetStats waited for the checkpoint business lock")
	}
}

func newContractTestConsumer(t *testing.T, level MetricsLevel) *consumer {
	t.Helper()
	config := GetDefaultConsumerConfig()
	config.ProjectID = "project"
	config.TopicIDList = []string{"topic"}
	config.ConsumerGroupName = "group"
	config.Endpoint = "http://127.0.0.1:0"
	config.Region = "cn-beijing"
	config.AccessKeyID = "ak"
	config.AccessKeySecret = "sk"
	config.MetricsLevel = level
	consumer, err := newConsumer(context.Background(), config, func(string, int, *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newConsumer: %v", err)
	}
	return consumer
}

func TestContractOffReturnsOnlyHeader(t *testing.T) {
	c := newContractTestConsumer(t, MetricsLevelOff)
	snapshot, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if snapshot.SchemaVersion != observability.SchemaVersion || snapshot.ComponentType != "consumer" ||
		!snapshot.MetricsDisabled || !snapshot.Partial {
		t.Fatalf("unexpected Off header: %+v", snapshot.SnapshotHeader)
	}
	if snapshot.Health != nil || snapshot.Operational != nil || snapshot.WorkerDetails != nil ||
		snapshot.RetryDetails != nil || snapshot.LatencyHistograms != nil {
		t.Fatalf("Off must not synthesize uncollected sections")
	}
	if c.stats.fetchCallDuration != nil {
		t.Fatalf("Off must not allocate histograms")
	}
}

func TestContractViewProjectionAndComponentCapabilities(t *testing.T) {
	core := newContractTestConsumer(t, MetricsLevelCore)
	debug, err := core.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if !debug.Partial || len(debug.MissingSections) != 3 || debug.Health == nil || debug.Operational == nil {
		t.Fatalf("Core debug must miss only supported detailed sections: %+v", debug)
	}

	detailed := newContractTestConsumer(t, MetricsLevelDetailed)
	full, err := detailed.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if full.Partial || full.Health == nil || full.Operational == nil || full.WorkerDetails == nil ||
		full.RetryDetails == nil || full.LatencyHistograms == nil {
		t.Fatalf("Detailed debug must be complete for Consumer capabilities: %+v", full)
	}
	if full.LatencyHistograms.FetchCallDuration.Histogram.BucketProfile != "fast_v1" ||
		full.LatencyHistograms.ProcessDuration.Histogram.BucketProfile != "slow_v1" {
		t.Fatalf("latency profiles do not match frozen semantics")
	}
}

func TestContractConsumerCounterOutcomes(t *testing.T) {
	c := newContractTestConsumer(t, MetricsLevelCore)
	fetchSuccess := c.stats.beginFetchCall()
	c.stats.finishFetchCall(fetchSuccess, false)
	c.stats.addFetchedFor(fetchSuccess, 2, 5)
	fetchFailure := c.stats.beginFetchCall()
	c.stats.finishFetchCall(fetchFailure, true)
	processSuccess := c.stats.beginProcessAttempt()
	c.stats.finishProcessAttempt(processSuccess, 2, 5, true)
	heartbeatFailure := c.stats.beginHeartbeatCall()
	c.stats.finishHeartbeatCall(heartbeatFailure, true)
	checkpointSuccess := c.stats.beginCheckpointCall()
	c.stats.finishCheckpointCall(checkpointSuccess, false)
	snapshot, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	counters := snapshot.Operational.Counters
	if counters.FetchCallsSuccessTotal != 1 || counters.FetchCallsFailureTotal != 1 ||
		counters.FetchedGroupsTotal != 2 || counters.FetchedLogsTotal != 5 ||
		counters.ProcessAttemptsSuccessTotal != 1 || counters.ProcessedSuccessGroupsTotal != 2 ||
		counters.ProcessedSuccessLogsTotal != 5 || counters.HeartbeatCallsFailureTotal != 1 ||
		counters.CheckpointCallsSuccessTotal != 1 {
		t.Fatalf("unexpected exact outcome counters: %+v", counters)
	}
}

func TestContractInFlightCallsAreNotReportedAsCompletedOutcomes(t *testing.T) {
	c := newContractTestConsumer(t, MetricsLevelCore)
	fetchGeneration := c.stats.beginFetchCall()
	processGeneration := c.stats.beginProcessAttempt()
	checkpointGeneration := c.stats.beginCheckpointCall()
	heartbeatGeneration := c.stats.beginHeartbeatCall()

	during, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats during calls: %v", err)
	}
	counters := during.Operational.Counters
	if counters.FetchCallsSuccessTotal != 0 || counters.FetchCallsFailureTotal != 0 ||
		counters.ProcessAttemptsSuccessTotal != 0 || counters.ProcessAttemptsFailureTotal != 0 ||
		counters.CheckpointCallsSuccessTotal != 0 || counters.CheckpointCallsFailureTotal != 0 ||
		counters.HeartbeatCallsSuccessTotal != 0 || counters.HeartbeatCallsFailureTotal != 0 {
		t.Fatalf("in-flight calls must not be projected as completed outcomes: %+v", counters)
	}
	if during.Operational.FetchInFlight != 1 || during.Operational.ProcessInFlight != 1 {
		t.Fatalf("in-flight gauges = fetch:%d process:%d, want 1/1", during.Operational.FetchInFlight, during.Operational.ProcessInFlight)
	}

	c.stats.finishFetchCall(fetchGeneration, false)
	c.stats.finishProcessAttempt(processGeneration, 2, 5, false)
	c.stats.finishCheckpointCall(checkpointGeneration, true)
	c.stats.finishHeartbeatCall(heartbeatGeneration, false)
	after, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats after calls: %v", err)
	}
	counters = after.Operational.Counters
	if counters.FetchCallsSuccessTotal != 1 || counters.ProcessAttemptsFailureTotal != 1 ||
		counters.CheckpointCallsFailureTotal != 1 || counters.HeartbeatCallsSuccessTotal != 1 ||
		after.Operational.FetchInFlight != 0 || after.Operational.ProcessInFlight != 0 {
		t.Fatalf("completed outcomes were not settled exactly once: %+v operational=%+v", counters, after.Operational)
	}
}

func TestContractSnapshotIsIndependentAndInvalidViewFails(t *testing.T) {
	c := newContractTestConsumer(t, MetricsLevelDetailed)
	first, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	first.WorkerDetails.Items = append(first.WorkerDetails.Items, WorkerDetail{Role: "caller_mutation"})
	first.LatencyHistograms.FetchCallDuration.Histogram.UpperBoundsMs[0] = 99
	first.LatencyHistograms.FetchCallDuration.Histogram.CumulativeBucketCounts[0] = 99
	second, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if len(second.WorkerDetails.Items) != 3 || second.LatencyHistograms.FetchCallDuration.Histogram.UpperBoundsMs[0] != 1 ||
		second.LatencyHistograms.FetchCallDuration.Histogram.CumulativeBucketCounts[0] != 0 {
		t.Fatalf("snapshot leaked mutable state")
	}
	if _, err := c.GetStats(StatsOptions{View: StatsView(99)}); err != ErrInvalidStatsView {
		t.Fatalf("invalid view error=%v", err)
	}
}

func TestContractOldGenerationCompletionDoesNotPolluteNewWindow(t *testing.T) {
	c := newContractTestConsumer(t, MetricsLevelCore)
	atomic.StoreInt32(&c.lifecycleStage, lifecycleRunning)
	oldGeneration := c.stats.beginFetchCall()
	if err := c.SetMetricsLevel(MetricsLevelOff); err != nil {
		t.Fatalf("disable metrics: %v", err)
	}
	if err := c.SetMetricsLevel(MetricsLevelCore); err != nil {
		t.Fatalf("enable metrics: %v", err)
	}
	c.stats.finishFetchCall(oldGeneration, true)
	snapshot, err := c.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if snapshot.Operational.Counters.FetchCallsFailureTotal != 0 || snapshot.Operational.FetchInFlight != 0 {
		t.Fatalf("old generation contaminated current window: %+v", snapshot.Operational)
	}
}

func TestContractCounterSampleRequiresOperationalView(t *testing.T) {
	c := newContractTestConsumer(t, MetricsLevelCore)
	health, _ := c.GetStats(StatsOptions{View: StatsViewHealth})
	if _, ok := health.CounterSample(); ok {
		t.Fatal("health view must not synthesize counters")
	}
	operational, _ := c.GetStats(StatsOptions{View: StatsViewOperational})
	if sample, ok := operational.CounterSample(); !ok || len(sample.Counters) != 13 {
		t.Fatalf("counter sample = %+v, ok=%v", sample, ok)
	}
}

func TestContractRejectsInvalidConfiguredLevel(t *testing.T) {
	config := GetDefaultConsumerConfig()
	config.ProjectID = "project"
	config.TopicIDList = []string{"topic"}
	config.ConsumerGroupName = "group"
	config.MetricsLevel = MetricsLevel(99)
	if _, err := newConsumer(context.Background(), config, func(string, int, *pb.LogGroupList) {}); err != ErrInvalidMetricsLevel {
		t.Fatalf("error = %v, want %v", err, ErrInvalidMetricsLevel)
	}
}

func TestContractQueueProjectionPreservesHighWatermarkInvariant(t *testing.T) {
	projected := projectQueue(&queueMeasurements{
		CurrentItems:       7,
		HighWatermarkItems: 6,
	})
	if projected.HighWatermarkItems != projected.CurrentItems {
		t.Fatalf("queue = %+v, high watermark must cover the current weak snapshot", projected)
	}
}
