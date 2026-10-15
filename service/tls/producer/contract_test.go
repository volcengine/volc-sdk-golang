package producer

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/observability"
)

func TestContractGetStatsDoesNotWaitForRetryHeapLock(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelCore)
	p.dispatcher.retryQueue.mutex.Lock()
	done := make(chan error, 1)
	go func() {
		_, err := p.GetStats(StatsOptions{View: StatsViewOperational})
		done <- err
	}()
	select {
	case err := <-done:
		p.dispatcher.retryQueue.mutex.Unlock()
		if err != nil {
			t.Fatalf("GetStats: %v", err)
		}
	case <-time.After(100 * time.Millisecond):
		p.dispatcher.retryQueue.mutex.Unlock()
		<-done
		t.Fatal("GetStats waited for the retry business lock")
	}
}

func newContractTestProducer(t *testing.T, level MetricsLevel) *producer {
	t.Helper()
	config := GetDefaultProducerConfig()
	config.Endpoint = "http://127.0.0.1:0"
	config.Region = "cn-beijing"
	config.AccessKeyID = "ak"
	config.AccessKeySecret = "sk"
	config.MetricsLevel = level
	return newProducer(config)
}

func TestContractOffReturnsOnlyHeader(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelOff)
	snapshot, err := p.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if snapshot.SchemaVersion != observability.SchemaVersion || snapshot.ComponentType != "producer" {
		t.Fatalf("unexpected header: %+v", snapshot.SnapshotHeader)
	}
	if !snapshot.MetricsDisabled || !snapshot.Partial {
		t.Fatalf("Off debug must be disabled and partial")
	}
	if snapshot.Health != nil || snapshot.Operational != nil || snapshot.WorkerDetails != nil ||
		snapshot.BatchInflightMemoryGate != nil || snapshot.ManagedMemory != nil || snapshot.RetryDetails != nil ||
		snapshot.LatencyHistograms != nil || snapshot.MemoryBreakdown != nil {
		t.Fatalf("Off must not synthesize uncollected sections")
	}
	if p.stats.batchQueueDelay != nil || p.stats.sendCallDuration != nil {
		t.Fatalf("Off must not allocate histograms")
	}
}

func TestContractViewProjectionAndPartial(t *testing.T) {
	core := newContractTestProducer(t, MetricsLevelCore)
	operational, err := core.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats operational: %v", err)
	}
	if operational.Partial || operational.Health == nil || operational.Operational == nil {
		t.Fatalf("Core operational must be complete: %+v", operational)
	}
	debug, err := core.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats debug: %v", err)
	}
	if !debug.Partial || len(debug.MissingSections) != 4 || debug.Health == nil || debug.Operational == nil {
		t.Fatalf("Core debug must contain available prefixes and four missing detailed sections: %+v", debug)
	}

	detailed := newContractTestProducer(t, MetricsLevelDetailed)
	health, err := detailed.GetStats(StatsOptions{View: StatsViewHealth})
	if err != nil {
		t.Fatalf("GetStats health: %v", err)
	}
	if health.Partial || health.Health == nil || health.BatchInflightMemoryGate == nil ||
		health.ManagedMemory == nil || health.Operational != nil || health.LatencyHistograms != nil {
		t.Fatalf("Health view must not expose wider sections")
	}
	full, err := detailed.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats debug: %v", err)
	}
	if full.Partial || full.Health == nil || full.Operational == nil || full.WorkerDetails == nil ||
		full.RetryDetails == nil || full.LatencyHistograms == nil || full.MemoryBreakdown == nil {
		t.Fatalf("Detailed debug must be complete: %+v", full)
	}
	if full.LatencyHistograms.SendCallDuration.Histogram.BucketProfile != "fast_v1" ||
		full.LatencyHistograms.BatchDeliveryDuration.Histogram.BucketProfile != "slow_v1" {
		t.Fatalf("latency profiles do not match frozen semantics")
	}
}

func TestContractMemoryPressureReportsActiveGate(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelCore)
	atomic.StoreInt64(&p.producerLogGroupSize, p.config.TotalSizeLnBytes)
	if err := p.memoryLimiter.acquirePayload(p.memoryLimiter.payloadLimit/2, 0); err != nil {
		t.Fatalf("acquire payload: %v", err)
	}
	snapshot, err := p.GetStats(StatsOptions{View: StatsViewHealth})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if snapshot.Health == nil || snapshot.Health.Signals.MemoryPressure == nil ||
		*snapshot.Health.Signals.MemoryPressure != 1 ||
		snapshot.Health.Signals.MemoryPressureSource != "batch_inflight" {
		t.Fatalf("unexpected batch pressure: %+v", snapshot.Health)
	}

	atomic.StoreInt64(&p.producerLogGroupSize, 0)
	p.memoryLimiter.releasePayload(p.memoryLimiter.payloadLimit / 2)
	managedBytes := p.memoryLimiter.payloadLimit - 1
	if err := p.memoryLimiter.acquirePayload(managedBytes, 0); err != nil {
		t.Fatalf("acquire managed payload: %v", err)
	}
	snapshot, err = p.GetStats(StatsOptions{View: StatsViewHealth})
	if err != nil {
		t.Fatalf("GetStats managed: %v", err)
	}
	if snapshot.Health == nil || snapshot.Health.Signals.MemoryPressure == nil ||
		snapshot.Health.Signals.MemoryPressureSource != "managed_payload_admission" ||
		*snapshot.Health.Signals.MemoryPressure < 0.99 {
		t.Fatalf("unexpected managed admission pressure: %+v", snapshot.Health)
	}
	p.memoryLimiter.releasePayload(managedBytes)
}

func TestContractSnapshotIsIndependent(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelDetailed)
	first, err := p.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	first.MissingSections = append(first.MissingSections, "caller_mutation")
	first.WorkerDetails.Items = append(first.WorkerDetails.Items, WorkerDetail{Role: "caller_mutation"})
	first.LatencyHistograms.SendCallDuration.Histogram.UpperBoundsMs[0] = 99
	first.LatencyHistograms.SendCallDuration.Histogram.CumulativeBucketCounts[0] = 99
	second, err := p.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if len(second.MissingSections) != 0 || len(second.WorkerDetails.Items) != 0 ||
		second.LatencyHistograms.SendCallDuration.Histogram.UpperBoundsMs[0] != 1 ||
		second.LatencyHistograms.SendCallDuration.Histogram.CumulativeBucketCounts[0] != 0 {
		t.Fatalf("snapshot leaked mutable internal/caller state")
	}
}

func TestContractManagedMemoryOwnership(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelDetailed)
	if err := p.memoryLimiter.acquirePayload(128, 0); err != nil {
		t.Fatalf("acquire payload: %v", err)
	}
	if err := p.memoryLimiter.acquireTemporary(64, 0); err != nil {
		t.Fatalf("acquire temporary: %v", err)
	}
	snapshot, err := p.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	memory := snapshot.ManagedMemory
	if memory == nil || !memory.Valid() || memory.PayloadOwnedBytes != 128 || memory.BoundedRuntimeBufferBytes != 64 {
		t.Fatalf("invalid managed memory: %+v", memory)
	}
	if !snapshot.MemoryBreakdown.Valid(*memory) {
		t.Fatalf("memory breakdown does not conserve payload ownership")
	}
	p.memoryLimiter.releaseTemporary(64)
	p.memoryLimiter.releasePayload(128)
}

func TestContractManagedMemoryStageConservation(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelDetailed)
	limiter := p.memoryLimiter
	if err := limiter.acquirePayload(128, 0); err != nil {
		t.Fatalf("acquire payload: %v", err)
	}
	assertBreakdown := func(want [4]int64) {
		t.Helper()
		ingress, aggregate, send, retry, ok := limiter.payloadBreakdown()
		if !ok || [4]int64{ingress, aggregate, send, retry} != want || ingress+aggregate+send+retry != limiter.PayloadUsed() {
			t.Fatalf("breakdown=(%d,%d,%d,%d,%v), payload=%d, want=%v", ingress, aggregate, send, retry, ok, limiter.PayloadUsed(), want)
		}
	}
	assertBreakdown([4]int64{128, 0, 0, 0})
	limiter.movePayload(payloadStageIngress, payloadStageAggregate, 128)
	assertBreakdown([4]int64{0, 128, 0, 0})
	limiter.movePayload(payloadStageAggregate, payloadStageSend, 128)
	assertBreakdown([4]int64{0, 0, 128, 0})
	limiter.movePayload(payloadStageSend, payloadStageRetry, 128)
	assertBreakdown([4]int64{0, 0, 0, 128})
	limiter.movePayload(payloadStageRetry, payloadStageSend, 128)
	assertBreakdown([4]int64{0, 0, 128, 0})
	limiter.releasePayloadFrom(payloadStageSend, 128)
	assertBreakdown([4]int64{})
}

func TestContractManagedMemoryObservabilityStopsAtOff(t *testing.T) {
	limiter := newMemoryLimiter(100, 100)
	if err := limiter.acquirePayload(10, 0); err != nil {
		t.Fatalf("off acquire: %v", err)
	}
	if limiter.Used() != 10 || limiter.PayloadUsed() != 0 || limiter.TemporaryUsed() != 0 {
		t.Fatalf("Off must keep only functional budget accounting: used=%d payload=%d temporary=%d",
			limiter.Used(), limiter.PayloadUsed(), limiter.TemporaryUsed())
	}
	if _, _, _, _, ok := limiter.payloadBreakdown(); ok {
		t.Fatal("Off must not expose an observed payload breakdown")
	}
	limiter.releasePayload(10)
	if limiter.HighWatermark() != 0 || limiter.PayloadRejected() != 0 || limiter.TemporaryRejected() != 0 {
		t.Fatalf("Off updated observability-only memory values")
	}
	limiter.setMetricsEnabled(true)
	if err := limiter.acquirePayload(60, 0); err != nil {
		t.Fatalf("enabled acquire: %v", err)
	}
	if limiter.HighWatermark() != 60 {
		t.Fatalf("high watermark = %d, want 60", limiter.HighWatermark())
	}
	if err := limiter.acquirePayload(50, 0); err == nil {
		t.Fatal("expected memory rejection")
	}
	if limiter.PayloadRejected() != 1 || limiter.TemporaryRejected() != 0 {
		t.Fatalf("rejection counts = payload:%d runtime:%d, want 1/0", limiter.PayloadRejected(), limiter.TemporaryRejected())
	}
	limiter.releasePayload(60)
}

func TestContractManagedMemoryRejectionTotalUsesCategories(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelCore)
	limiter := p.memoryLimiter
	if err := limiter.acquirePayload(limiter.payloadLimit+1, 0); err == nil {
		t.Fatal("expected payload admission rejection")
	}
	if err := limiter.acquireTemporary(limiter.limit+1, 0); err == nil {
		t.Fatal("expected runtime buffer rejection")
	}
	snapshot, err := p.GetStats(StatsOptions{View: StatsViewHealth})
	if err != nil {
		t.Fatal(err)
	}
	if snapshot.ManagedMemory == nil || snapshot.ManagedMemory.RejectedTotal != 2 ||
		snapshot.ManagedMemory.PayloadAdmissionRejectedTotal != 1 || snapshot.ManagedMemory.RuntimeBufferRejectedTotal != 1 {
		t.Fatalf("public rejection total must sum both categories: %+v", snapshot.ManagedMemory)
	}
}

func TestContractHotEnableDoesNotFabricateManagedMemory(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelOff)
	p.Start()
	t.Cleanup(p.ForceClose)
	if err := p.memoryLimiter.acquirePayload(10, 0); err != nil {
		t.Fatalf("off acquire: %v", err)
	}
	if err := p.SetMetricsLevel(MetricsLevelCore); err != nil {
		t.Fatalf("SetMetricsLevel: %v", err)
	}
	snapshot, err := p.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if !snapshot.Partial || snapshot.Operational == nil || snapshot.ManagedMemory != nil ||
		!containsSection(snapshot.MissingSections, observability.SectionManagedMemory) {
		t.Fatalf("pre-enable ownership must be absent and explicit: %+v", snapshot)
	}

	p.memoryLimiter.releasePayload(10)
	snapshot, err = p.GetStats(StatsOptions{View: StatsViewOperational})
	if err != nil {
		t.Fatalf("GetStats after drain: %v", err)
	}
	if snapshot.Partial || snapshot.ManagedMemory == nil || !snapshot.ManagedMemory.Valid() {
		t.Fatalf("managed memory must recover after unknown ownership drains: %+v", snapshot)
	}
}

func containsSection(sections []StatsSection, target StatsSection) bool {
	for _, section := range sections {
		if section == target {
			return true
		}
	}
	return false
}

func TestContractCounterSampleRequiresOperationalView(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelCore)
	health, _ := p.GetStats(StatsOptions{View: StatsViewHealth})
	if _, ok := health.CounterSample(); ok {
		t.Fatal("health view must not synthesize counters")
	}
	operational, _ := p.GetStats(StatsOptions{View: StatsViewOperational})
	if sample, ok := operational.CounterSample(); !ok || len(sample.Counters) != 19 {
		t.Fatalf("counter sample = %+v, ok=%v", sample, ok)
	}
}

func TestContractRejectsInvalidView(t *testing.T) {
	p := newContractTestProducer(t, MetricsLevelCore)
	if _, err := p.GetStats(StatsOptions{View: StatsView(99)}); err != ErrInvalidStatsView {
		t.Fatalf("invalid view error=%v", err)
	}
}

func TestContractInvalidConfiguredLevelFailsClosed(t *testing.T) {
	config := GetDefaultProducerConfig()
	config.MetricsLevel = MetricsLevel(99)
	p := newProducer(config)
	if p.metricsLevelNow() != MetricsLevelOff {
		t.Fatalf("invalid level = %v, want Off", p.metricsLevelNow())
	}
}

func TestContractQueueProjectionPreservesHighWatermarkInvariant(t *testing.T) {
	projected := projectQueue(&queueMeasurements{
		CurrentItems:       7,
		HighWatermarkItems: 6,
		CapacityItems:      10,
		Bounded:            true,
	})
	if projected.HighWatermarkItems != projected.CurrentItems {
		t.Fatalf("queue = %+v, high watermark must cover the current weak snapshot", projected)
	}
}
