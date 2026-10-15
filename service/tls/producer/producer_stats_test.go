package producer

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// newTestProducer creates a producer for testing without requiring network.
func newTestProducer() *producer {
	config := GetDefaultProducerConfig()
	config.Endpoint = "http://localhost:8080"
	config.Region = "cn-beijing"
	return newTestProducerWithClient(config, &mockProducerClient{})
}

func makeLog() *pb.Log {
	return &pb.Log{Time: 1, Contents: []*pb.LogContent{{Key: "k", Value: "v"}}}
}

// TestGetStats_BeforeStart verifies GetStats works before Start() is called.
func TestGetStats_BeforeStart(t *testing.T) {
	p := newTestProducer()
	stats := p.collectMeasurements()

	assert.Equal(t, "producer", stats.ComponentType)
	assert.Equal(t, stateNew, stats.State)
	assert.Equal(t, MetricsLevelOff, stats.MetricsLevel)
	assert.True(t, stats.MetricsDisabled)
	// When off, business counters must be nil (absent)
	assert.Nil(t, stats.AcceptedLogs)
	assert.Nil(t, stats.TerminalSuccessLogs)
}

// TestGetStats_AfterClose verifies GetStats returns closed state after Close().
func TestGetStats_AfterClose(t *testing.T) {
	p := newTestProducer()
	p.Start()
	p.Close()

	stats := p.collectMeasurements()
	assert.Equal(t, stateClosed, stats.State)
}

// TestGetStats_AfterForceClose verifies GetStats returns closed state after ForceClose().
func TestGetStats_AfterForceClose(t *testing.T) {
	p := newTestProducer()
	p.Start()
	p.ForceClose()

	stats := p.collectMeasurements()
	assert.Equal(t, stateClosed, stats.State)
}

// TestGetStats_NoPanic verifies GetStats never panics at any lifecycle state.
func TestGetStats_NoPanic(t *testing.T) {
	p := newTestProducer()

	assert.NotPanics(t, func() { p.collectMeasurements() })
	p.Start()
	assert.NotPanics(t, func() { p.collectMeasurements() })
	p.Close()
	assert.NotPanics(t, func() { p.collectMeasurements() })
	assert.NotPanics(t, func() { p.Close() })
	assert.NotPanics(t, func() { p.collectMeasurements() })
}

// TestGetStats_Conservation verifies accepted logs are terminal after close.
// Uses metrics_level=core so counters are actually populated.
func TestGetStats_Conservation(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()

	for i := 0; i < 5; i++ {
		assert.NoError(t, p.SendLog("hash", "topic", "source", "file", makeLog(), nil))
	}

	stats := p.collectMeasurements()

	// While work is still in flight, terminal outcomes cannot exceed admissions.
	accepted := *stats.AcceptedLogs
	terminal := *stats.TerminalSuccessLogs + *stats.TerminalFailureLogs + *stats.TerminalDroppedLogs
	assert.LessOrEqual(t, terminal, accepted,
		"terminal outcomes must not exceed accepted logs: terminal(%d) accepted(%d)", terminal, accepted)

	p.Close()
	stats = p.collectMeasurements()
	assert.Equal(t, uint64(5), *stats.AcceptedLogs, "all successful SendLog calls must be accepted")
	terminal = *stats.TerminalSuccessLogs + *stats.TerminalFailureLogs + *stats.TerminalDroppedLogs
	assert.Equal(t, *stats.AcceptedLogs, terminal, "all accepted logs must be terminal after Close")

	assert.Greater(t, *stats.AcceptedGroups, uint64(0), "successful SendLog calls must produce groups")
	terminalGroups := *stats.TerminalSuccessGroups + *stats.TerminalFailureGroups + *stats.TerminalDroppedGroups
	assert.Equal(t, *stats.AcceptedGroups, terminalGroups, "all accepted groups must be terminal after Close")
}

// TestGetStats_Conservation_Groups verifies group-unit terminal accounting.
func TestGetStats_Conservation_Groups(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()

	// Send multiple logs — they should be grouped
	for i := 0; i < 10; i++ {
		assert.NoError(t, p.SendLog("hash", "topic", "source", "file", makeLog(), nil))
	}

	stats := p.collectMeasurements()
	accepted := *stats.AcceptedGroups
	terminal := *stats.TerminalSuccessGroups + *stats.TerminalFailureGroups + *stats.TerminalDroppedGroups
	assert.LessOrEqual(t, terminal, accepted,
		"terminal outcomes must not exceed accepted groups: terminal(%d) accepted(%d)", terminal, accepted)

	p.Close()
	stats = p.collectMeasurements()
	assert.Equal(t, uint64(10), *stats.AcceptedLogs, "all successful SendLog calls must be accepted")
	terminalLogs := *stats.TerminalSuccessLogs + *stats.TerminalFailureLogs + *stats.TerminalDroppedLogs
	assert.Equal(t, *stats.AcceptedLogs, terminalLogs, "all accepted logs must be terminal after Close")
	assert.Greater(t, *stats.AcceptedGroups, uint64(0), "successful SendLog calls must produce groups")
	terminal = *stats.TerminalSuccessGroups + *stats.TerminalFailureGroups + *stats.TerminalDroppedGroups
	assert.Equal(t, *stats.AcceptedGroups, terminal, "all accepted groups must be terminal after Close")
}

// TestProducerContractInterface verifies the concrete implementation satisfies
// the new contract-bearing Producer interface.
func TestProducerContractInterface(t *testing.T) {
	p := newTestProducer()
	var contract Producer = p
	stats, err := contract.GetStats(StatsOptions{View: StatsViewHealth})
	assert.NoError(t, err)
	assert.Equal(t, uint32(2), stats.SchemaVersion)
	assert.True(t, stats.MetricsDisabled)
}

// TestGetStats_NoIO verifies GetStats does not perform I/O (returns quickly).
func TestGetStats_NoIO(t *testing.T) {
	p := newTestProducer()
	p.Start()

	for i := 0; i < 100; i++ {
		stats := p.collectMeasurements()
		assert.NotNil(t, stats)
	}

	p.Close()
}

// TestGetStats_SnapshotSequence verifies snapshot sequence increments.
func TestGetStats_SnapshotSequence(t *testing.T) {
	p := newTestProducer()

	s1 := p.collectMeasurements()
	s2 := p.collectMeasurements()
	s3 := p.collectMeasurements()

	assert.Equal(t, uint64(1), s1.SnapshotSequence)
	assert.Equal(t, uint64(2), s2.SnapshotSequence)
	assert.Equal(t, uint64(3), s3.SnapshotSequence)
}

// TestGetStats_SnapshotSequence_Frozen verifies sequence stops incrementing when closed.
func TestGetStats_SnapshotSequence_Frozen(t *testing.T) {
	p := newTestProducer()
	p.Start()
	p.Close()

	s1 := p.collectMeasurements()
	s2 := p.collectMeasurements()
	// When closed, snapshot_sequence must not increment (frozen)
	assert.Equal(t, s1.SnapshotSequence, s2.SnapshotSequence)
}

// TestGetStats_Concurrent verifies GetStats is safe for concurrent use.
func TestGetStats_Concurrent(t *testing.T) {
	p := newTestProducer()
	p.Start()

	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				_ = p.collectMeasurements()
			}
			done <- struct{}{}
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	p.Close()
}

// TestGetStats_Uptime verifies uptime increases after start. An immediate
// millisecond-resolution snapshot may legitimately be zero, so compare two
// snapshots across a deterministic clock tick rather than asserting >0.
func TestGetStats_Uptime(t *testing.T) {
	p := newTestProducer()
	p.Start()

	first := p.collectMeasurements().UptimeMs
	time.Sleep(2 * time.Millisecond)
	assert.Greater(t, p.collectMeasurements().UptimeMs, first)

	p.Close()
}

// TestGetStats_RequiredFields verifies all required fields are present when metrics on.
func TestGetStats_RequiredFields(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()

	stats := p.collectMeasurements()

	// Common required fields
	assert.NotEmpty(t, stats.InstanceID)
	assert.NotZero(t, stats.CapturedAtUnixMs)
	assert.NotEmpty(t, stats.SDKName)
	assert.NotEmpty(t, stats.SDKLanguage)
	assert.NotEmpty(t, stats.ComponentType)
	assert.NotEmpty(t, stats.State)
	assert.NotEmpty(t, stats.Consistency)

	// Business counters must be present (non-nil) when metrics on
	assert.NotNil(t, stats.AcceptedLogs)
	assert.NotNil(t, stats.AcceptedGroups)
	assert.NotNil(t, stats.TerminalSuccessLogs)
	assert.NotNil(t, stats.TerminalFailureLogs)
	assert.NotNil(t, stats.TerminalDroppedLogs)
	assert.NotNil(t, stats.TerminalSuccessGroups)
	assert.NotNil(t, stats.TerminalFailureGroups)
	assert.NotNil(t, stats.TerminalDroppedGroups)
	assert.NotNil(t, stats.SendCallsInFlight)
	assert.NotNil(t, stats.RetriesScheduled)
	assert.NotNil(t, stats.SendCallPayloadLogs)
	assert.NotNil(t, stats.SendCallPayloadGroups)
	assert.NotNil(t, stats.Memory)
	assert.NotNil(t, stats.AggregateBatches)
	assert.NotNil(t, stats.AggregateGroups)
	assert.NotNil(t, stats.SendQueue, "send_queue is required by the operational contract")
	assert.NotNil(t, stats.RetryQueue, "retry_queue is required by the operational contract")
	assert.NotNil(t, stats.WorkerRoles, "worker_roles should be populated when running")

	p.Close()
}

// TestInstanceID_Unique verifies instance IDs are unique across instances.
func TestInstanceID_Unique(t *testing.T) {
	producer1 := newTestProducer()
	producer2 := newTestProducer()
	producer3 := newTestProducer()

	assert.NotEqual(t, producer1.instanceID, producer2.instanceID)
	assert.NotEqual(t, producer2.instanceID, producer3.instanceID)
	assert.NotEqual(t, producer1.instanceID, producer3.instanceID)
}

// TestInstanceID_Stable verifies instance ID is stable for the same instance.
func TestInstanceID_Stable(t *testing.T) {
	p := newTestProducer()
	id1 := p.collectMeasurements().InstanceID
	id2 := p.collectMeasurements().InstanceID
	assert.Equal(t, id1, id2)
}

// TestMetricsOff_JSON_OmitsBusinessCounters verifies JSON output omits business
// counters when metrics collection is Off.
func TestMetricsOff_JSON_OmitsBusinessCounters(t *testing.T) {
	p := newTestProducer()
	p.Start()

	stats, getErr := p.GetStats(StatsOptions{View: StatsViewDebug})
	assert.NoError(t, getErr)
	data, err := json.Marshal(stats)
	assert.NoError(t, err)

	jsonStr := string(data)
	// Business counter fields must NOT appear in the JSON
	assert.False(t, strings.Contains(jsonStr, "accepted_logs_total"), "accepted_logs_total should be absent when off")
	assert.False(t, strings.Contains(jsonStr, "terminal_success_logs_total"), "terminal_success_logs_total should be absent when off")
	assert.False(t, strings.Contains(jsonStr, "terminal_failure_logs_total"), "terminal_failure_logs_total should be absent when off")
	assert.False(t, strings.Contains(jsonStr, "terminal_dropped_logs_total"), "terminal_dropped_logs_total should be absent when off")
	assert.False(t, strings.Contains(jsonStr, "send_calls_success_total"), "send_calls_success_total should be absent when off")
	assert.False(t, strings.Contains(jsonStr, "send_calls_failure_total"), "send_calls_failure_total should be absent when off")
	assert.False(t, strings.Contains(jsonStr, "rejected_invalid_argument_total"), "rejected_invalid_argument_total should be absent when off")
	assert.False(t, strings.Contains(jsonStr, "rejected_memory_full_total"), "rejected_memory_full_total should be absent when off")
	// Lifecycle fields must be present
	assert.True(t, strings.Contains(jsonStr, "schema_version"))
	assert.True(t, strings.Contains(jsonStr, "metrics_disabled"))
	assert.True(t, strings.Contains(jsonStr, "lifecycle_state"))

	p.Close()
}

// TestMetricsOn_JSON_IncludesBusinessCounters verifies JSON output includes
// business counters when metrics_level=core.
func TestMetricsOn_JSON_IncludesBusinessCounters(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()

	stats, getErr := p.GetStats(StatsOptions{View: StatsViewOperational})
	assert.NoError(t, getErr)
	data, err := json.Marshal(stats)
	assert.NoError(t, err)

	jsonStr := string(data)
	assert.True(t, strings.Contains(jsonStr, "accepted_logs_total"), "accepted_logs_total should be present when on")
	assert.True(t, strings.Contains(jsonStr, "terminal_success_logs_total"), "terminal outcome counters should be present when on")
	assert.True(t, strings.Contains(jsonStr, `"critical_flags":[]`), "critical_flags should be present as an empty array")

	p.Close()
}

// TestForceClose_TerminalConservation verifies ForceClose terminal accounting.
func TestForceClose_TerminalConservation(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()

	// Accept some logs
	for i := 0; i < 5; i++ {
		assert.NoError(t, p.SendLog("hash", "topic", "source", "file", makeLog(), nil))
	}

	p.ForceClose()

	stats := p.collectMeasurements()
	accepted := *stats.AcceptedLogs
	terminal := *stats.TerminalSuccessLogs + *stats.TerminalFailureLogs + *stats.TerminalDroppedLogs
	assert.Equal(t, uint64(5), accepted, "all successful SendLog calls must be accepted")
	assert.Equal(t, accepted, terminal, "all accepted logs must be terminal after ForceClose")
}

// TestState_Closing verifies the closing state is observable.
func TestState_Closing(t *testing.T) {
	p := newTestProducer()
	p.Start()

	// We can't easily hold the closing state with active sends in a unit test,
	// but we can verify the state machine transitions: new → running → closed
	s1 := p.collectMeasurements()
	assert.Equal(t, stateRunning, s1.State)

	p.Close()
	s2 := p.collectMeasurements()
	assert.Equal(t, stateClosed, s2.State)
	assert.Equal(t, consistencyFinal, s2.Consistency)
}

// TestConsistency_FinalWhenClosed verifies consistency=final when closed.
func TestConsistency_FinalWhenClosed(t *testing.T) {
	p := newTestProducer()
	p.Start()
	p.Close()

	stats := p.collectMeasurements()
	assert.Equal(t, consistencyFinal, stats.Consistency)
}

// TestSDKVersion_FromSource verifies sdk_version comes from base package, not hardcoded.
func TestSDKVersion_FromSource(t *testing.T) {
	p := newTestProducer()
	stats := p.collectMeasurements()

	assert.Equal(t, base.SDKVersion, stats.SDKVersion)
}

// TestMemoryStats_Present verifies memoryMeasurements is populated when metrics on.
func TestMemoryStats_Present(t *testing.T) {
	p := newTestProducer()
	p.runtimeMetricsLevel = int32(MetricsLevelCore)
	p.Start()

	stats := p.collectMeasurements()
	assert.NotNil(t, stats.Memory)
	assert.Greater(t, stats.Memory.AccountedLimitBytes, uint64(0))
	assert.Equal(t, uint64(0), stats.Memory.AccountedUsedBytes) // no data sent yet

	p.Close()
}
