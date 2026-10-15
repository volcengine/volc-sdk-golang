package consumer

import (
	"time"
)

// histogramProfile selects an internal frozen bucket set. It is not a public
// collection switch; MetricsLevelDetailed owns histogram allocation.
type histogramProfile int

const (
	histogramProfileOff histogramProfile = iota
	histogramProfileFastV1
	histogramProfileSlowV1
)

// consumerState is the lifecycle state of the consumer.
type consumerState string

const (
	stateNew     consumerState = "new"
	stateRunning consumerState = "running"
	stateClosing consumerState = "closing"
	stateClosed  consumerState = "closed"
)

// consistencyLevel indicates snapshot consistency.
type consistencyLevel string

const (
	consistencyWeak  consistencyLevel = "weak"
	consistencyFinal consistencyLevel = "final"
)

// workerRole identifies a consumer worker role for role stats.
type workerRole string

const (
	roleMainLoop   workerRole = "main_loop"
	roleHeartbeat  workerRole = "heartbeat"
	roleCheckpoint workerRole = "checkpoint"
)

// queueMeasurements is the private source for the public queue contract.
type queueMeasurements struct {
	CurrentItems       uint64
	CapacityItems      uint64
	HighWatermarkItems uint64
	Bounded            bool
}

// workerRoleMeasurements reports one bounded worker-role fact.
type workerRoleMeasurements struct {
	Role     workerRole
	Expected uint32
	Running  uint32
	InFlight uint32
}

// consumerMeasurements is the private, pre-aggregated input used to project
// the public view-specific Contract snapshot.
type consumerMeasurements struct {
	// Common snapshot facts.
	InstanceID       string
	CapturedAtUnixMs int64
	SDKName          string
	SDKLanguage      string
	SDKVersion       string
	ComponentType    string
	State            consumerState
	Consistency      consistencyLevel
	MetricsLevel     MetricsLevel
	MetricsDisabled  bool
	UptimeMs         int64
	SnapshotSequence uint64
	LastSuccessAgeMs *uint64
	LastErrorAgeMs   *uint64

	// Worker roles (only declared roles are output).
	WorkerRoles []workerRoleMeasurements

	// Assignment and shard workers.
	AssignedShards *uint32

	// Held batch references and ready-to-process work.
	HeldBatchReferences *uint64
	ReadyToProcess      *uint64

	// In-flight calls.
	FetchInFlight   *uint32
	ProcessInFlight *uint32

	// Checkpoint state.
	CheckpointPending *queueMeasurements

	// Cumulative counters.
	FetchCalls             *uint64
	FetchFailures          *uint64
	FetchedLogGroups       *uint64
	FetchedLogs            *uint64
	ProcessAttempts        *uint64
	ProcessFailures        *uint64
	ProcessedSuccessGroups *uint64
	ProcessedSuccessLogs   *uint64
	HeartbeatCalls         *uint64
	HeartbeatFailures      *uint64
	CheckpointCalls        *uint64
	CheckpointFailures     *uint64
	BackoffEvents          *uint64
	RebalanceEvents        *uint64

	// Detailed latency histograms. Nil below MetricsLevelDetailed.
	FetchCallDurationMs      *HistogramSnapshot
	DispatchDelayMs          *HistogramSnapshot
	ProcessDurationMs        *HistogramSnapshot
	CheckpointCallDurationMs *HistogramSnapshot
}

// collectMeasurements returns the pre-aggregated internal measurements used
// to project the public view-specific snapshot.
// It is O(1), performs no I/O, invokes no callbacks, and does not traverse
// business maps/queues deeply. Safe to call from any goroutine at any lifecycle state.
func (c *consumer) collectMeasurements() consumerMeasurements {
	if c == nil || c.stats == nil {
		return consumerMeasurements{MetricsDisabled: true}
	}
	return c.stats.snapshot(c)
}

// nowMs returns the current unix time in milliseconds.
func nowMs() int64 {
	return unixMillis(time.Now())
}

func unixMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
