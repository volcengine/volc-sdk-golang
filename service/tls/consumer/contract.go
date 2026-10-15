package consumer

import "github.com/volcengine/volc-sdk-golang/service/tls/observability"

type MetricsLevel = observability.MetricsLevel
type StatsView = observability.StatsView
type StatsOptions = observability.StatsOptions
type SnapshotHeader = observability.SnapshotHeader
type StatsSection = observability.StatsSection
type QueueSummary = observability.QueueSummary
type WorkerSummary = observability.WorkerSummary
type HealthSignals = observability.HealthSignals
type HealthPolicy = observability.HealthPolicy
type HealthResult = observability.HealthResult

const (
	MetricsLevelOff      = observability.MetricsOff
	MetricsLevelCore     = observability.MetricsCore
	MetricsLevelDetailed = observability.MetricsDetailed

	StatsViewHealth      = observability.ViewHealth
	StatsViewOperational = observability.ViewOperational
	StatsViewDebug       = observability.ViewDebug
)

var ErrInvalidStatsView = observability.ErrInvalidStatsView

type StatsSnapshot struct {
	observability.SnapshotHeader
	MissingSections   []observability.StatsSection `json:"missing_sections,omitempty"`
	Health            *HealthStats                 `json:"health,omitempty"`
	Operational       *OperationalStats            `json:"operational,omitempty"`
	WorkerDetails     *WorkerDetails               `json:"worker_details,omitempty"`
	RetryDetails      *RetryDetails                `json:"retry_details,omitempty"`
	LatencyHistograms *LatencyHistograms           `json:"latency_histograms,omitempty"`
}

type HealthStats struct {
	Signals HealthSignals `json:"signals"`
}

type ConsumerCounters struct {
	FetchCallsSuccessTotal      uint64 `json:"fetch_calls_success_total"`
	FetchCallsFailureTotal      uint64 `json:"fetch_calls_failure_total"`
	FetchedLogsTotal            uint64 `json:"fetched_logs_total"`
	FetchedGroupsTotal          uint64 `json:"fetched_groups_total"`
	ProcessAttemptsSuccessTotal uint64 `json:"process_attempts_success_total"`
	ProcessAttemptsFailureTotal uint64 `json:"process_attempts_failure_total"`
	ProcessedSuccessLogsTotal   uint64 `json:"processed_success_logs_total"`
	ProcessedSuccessGroupsTotal uint64 `json:"processed_success_groups_total"`
	CheckpointCallsSuccessTotal uint64 `json:"checkpoint_calls_success_total"`
	CheckpointCallsFailureTotal uint64 `json:"checkpoint_calls_failure_total"`
	HeartbeatCallsSuccessTotal  uint64 `json:"heartbeat_calls_success_total"`
	HeartbeatCallsFailureTotal  uint64 `json:"heartbeat_calls_failure_total"`
	RebalanceTotal              uint64 `json:"rebalance_total"`
}

type OperationalStats struct {
	Counters ConsumerCounters `json:"counters"`

	AssignedShards      uint32 `json:"assigned_shards"`
	HeldBatchReferences uint64 `json:"held_batch_references"`
	ReadyToProcess      uint64 `json:"ready_to_process"`
	FetchInFlight       uint32 `json:"fetch_in_flight"`
	ProcessInFlight     uint32 `json:"process_in_flight"`

	CheckpointPending *QueueSummary `json:"checkpoint_pending,omitempty"`
	Workers           WorkerSummary `json:"workers"`
}

type WorkerDetail struct {
	Role     string `json:"role"`
	Expected uint32 `json:"expected"`
	Running  uint32 `json:"running"`
	InFlight uint32 `json:"in_flight"`
}

type WorkerDetails struct {
	Items     []WorkerDetail `json:"items"`
	Truncated bool           `json:"truncated"`
}

type RetryDetails struct {
	BackoffEventsTotal uint64 `json:"backoff_events_total"`
}

type LatencyHistogram struct {
	Start         string             `json:"start"`
	End           string             `json:"end"`
	Unit          string             `json:"unit"`
	Object        string             `json:"object"`
	IncludesRetry bool               `json:"includes_retry"`
	Histogram     *HistogramSnapshot `json:"histogram,omitempty"`
}

type LatencyHistograms struct {
	FetchCallDuration      LatencyHistogram `json:"fetch_call_duration"`
	DispatchDelay          LatencyHistogram `json:"dispatch_delay"`
	ProcessDuration        LatencyHistogram `json:"process_duration"`
	CheckpointCallDuration LatencyHistogram `json:"checkpoint_call_duration"`
}

func (consumer *consumer) GetStats(options StatsOptions) (StatsSnapshot, error) {
	if !options.View.Valid() {
		return StatsSnapshot{}, ErrInvalidStatsView
	}
	consumer.metricsMu.RLock()
	defer consumer.metricsMu.RUnlock()
	raw := consumer.collectMeasurements()
	missing, _ := observability.MissingSectionsForCapabilities(raw.MetricsLevel, options.View, consumerContractCapabilities)
	snapshot := StatsSnapshot{
		SnapshotHeader:  consumer.snapshotHeader(raw, options.View, len(missing) != 0),
		MissingSections: missing,
	}
	if observability.Available(raw.MetricsLevel, observability.SectionHealth) {
		snapshot.Health = &HealthStats{Signals: consumer.healthSignals(raw)}
	}
	if options.View >= StatsViewOperational && observability.Available(raw.MetricsLevel, observability.SectionOperational) {
		snapshot.Operational = consumer.operationalStats(raw)
	}
	if options.View == StatsViewDebug && raw.MetricsLevel == MetricsLevelDetailed {
		snapshot.WorkerDetails = consumer.workerDetails(raw)
		snapshot.RetryDetails = &RetryDetails{BackoffEventsTotal: value64(raw.BackoffEvents)}
		snapshot.LatencyHistograms = consumer.latencyHistograms(raw)
	}
	return snapshot, nil
}

func (consumer *consumer) snapshotHeader(raw consumerMeasurements, view StatsView, partial bool) SnapshotHeader {
	return SnapshotHeader{
		SchemaVersion:             observability.SchemaVersion,
		InstanceID:                raw.InstanceID,
		SDKName:                   raw.SDKName,
		SDKLanguage:               raw.SDKLanguage,
		SDKVersion:                raw.SDKVersion,
		ComponentType:             raw.ComponentType,
		CapturedAtUnixMs:          raw.CapturedAtUnixMs,
		UptimeMs:                  raw.UptimeMs,
		SnapshotSequence:          raw.SnapshotSequence,
		CollectionGeneration:      consumer.collectionGenerationNow(),
		CollectionStartedAtUnixMs: consumer.collectionStartedAtNow(),
		LifecycleState:            observability.LifecycleState(raw.State),
		Consistency:               observability.ConsistencyLevel(raw.Consistency),
		CollectedLevel:            raw.MetricsLevel,
		RequestedView:             view,
		MetricsDisabled:           raw.MetricsDisabled,
		Partial:                   partial,
		Capabilities:              consumerContractCapabilities,
	}
}

func (consumer *consumer) healthSignals(raw consumerMeasurements) HealthSignals {
	workers := summarizeWorkers(raw.WorkerRoles)
	return HealthSignals{
		LastSuccessAgeMs: cloneUint64(raw.LastSuccessAgeMs),
		LastErrorAgeMs:   cloneUint64(raw.LastErrorAgeMs),
		WorkerSummary:    workers,
		CriticalFlags:    []string{},
	}
}

func (consumer *consumer) operationalStats(raw consumerMeasurements) *OperationalStats {
	fetchCalls := value64(raw.FetchCalls)
	fetchFailures := value64(raw.FetchFailures)
	processAttempts := value64(raw.ProcessAttempts)
	processFailures := value64(raw.ProcessFailures)
	heartbeatCalls := value64(raw.HeartbeatCalls)
	heartbeatFailures := value64(raw.HeartbeatFailures)
	checkpointCalls := value64(raw.CheckpointCalls)
	checkpointFailures := value64(raw.CheckpointFailures)
	return &OperationalStats{
		Counters: ConsumerCounters{
			FetchCallsSuccessTotal:      subtractFloor(fetchCalls, fetchFailures),
			FetchCallsFailureTotal:      fetchFailures,
			FetchedLogsTotal:            value64(raw.FetchedLogs),
			FetchedGroupsTotal:          value64(raw.FetchedLogGroups),
			ProcessAttemptsSuccessTotal: subtractFloor(processAttempts, processFailures),
			ProcessAttemptsFailureTotal: processFailures,
			ProcessedSuccessLogsTotal:   value64(raw.ProcessedSuccessLogs),
			ProcessedSuccessGroupsTotal: value64(raw.ProcessedSuccessGroups),
			CheckpointCallsSuccessTotal: subtractFloor(checkpointCalls, checkpointFailures),
			CheckpointCallsFailureTotal: checkpointFailures,
			HeartbeatCallsSuccessTotal:  subtractFloor(heartbeatCalls, heartbeatFailures),
			HeartbeatCallsFailureTotal:  heartbeatFailures,
			RebalanceTotal:              value64(raw.RebalanceEvents),
		},
		AssignedShards:      value32(raw.AssignedShards),
		HeldBatchReferences: value64(raw.HeldBatchReferences),
		ReadyToProcess:      value64(raw.ReadyToProcess),
		FetchInFlight:       value32(raw.FetchInFlight),
		ProcessInFlight:     value32(raw.ProcessInFlight),
		CheckpointPending:   projectQueue(raw.CheckpointPending),
		Workers:             summarizeWorkers(raw.WorkerRoles),
	}
}

func (consumer *consumer) workerDetails(raw consumerMeasurements) *WorkerDetails {
	items := make([]WorkerDetail, 0, len(raw.WorkerRoles))
	for _, role := range raw.WorkerRoles {
		items = append(items, WorkerDetail{Role: string(role.Role), Expected: role.Expected, Running: role.Running, InFlight: role.InFlight})
	}
	return &WorkerDetails{Items: items, Truncated: false}
}

func (consumer *consumer) latencyHistograms(raw consumerMeasurements) *LatencyHistograms {
	return &LatencyHistograms{
		FetchCallDuration:      LatencyHistogram{Start: "consume_logs_call_start", End: "consume_logs_call_end", Unit: "ms", Object: "call", IncludesRetry: true, Histogram: cloneHistogram(raw.FetchCallDurationMs)},
		DispatchDelay:          LatencyHistogram{Start: "batch_ready", End: "callback_start", Unit: "ms", Object: "batch", IncludesRetry: false, Histogram: cloneHistogram(raw.DispatchDelayMs)},
		ProcessDuration:        LatencyHistogram{Start: "callback_start", End: "callback_end", Unit: "ms", Object: "batch", IncludesRetry: false, Histogram: cloneHistogram(raw.ProcessDurationMs)},
		CheckpointCallDuration: LatencyHistogram{Start: "checkpoint_call_start", End: "checkpoint_call_end", Unit: "ms", Object: "call", IncludesRetry: true, Histogram: cloneHistogram(raw.CheckpointCallDurationMs)},
	}
}

func summarizeWorkers(roles []workerRoleMeasurements) WorkerSummary {
	var summary WorkerSummary
	for _, role := range roles {
		summary.Expected += role.Expected
		summary.Running += role.Running
		summary.InFlight += role.InFlight
		if role.Running < role.Expected {
			summary.MissingCriticalRoles++
		}
	}
	return summary
}

func projectQueue(queue *queueMeasurements) *QueueSummary {
	if queue == nil {
		return nil
	}
	highWatermark := queue.HighWatermarkItems
	if highWatermark < queue.CurrentItems {
		highWatermark = queue.CurrentItems
	}
	result := &QueueSummary{
		CurrentItems:       queue.CurrentItems,
		HighWatermarkItems: highWatermark,
		Bounded:            queue.Bounded,
	}
	if queue.CapacityItems != 0 || queue.Bounded {
		result.CapacityItems = uint64Ptr(queue.CapacityItems)
	}
	return result
}

func cloneHistogram(histogram *HistogramSnapshot) *HistogramSnapshot {
	if histogram == nil {
		return nil
	}
	copy := *histogram
	copy.UpperBoundsMs = append([]uint64(nil), histogram.UpperBoundsMs...)
	copy.CumulativeBucketCounts = append([]uint64(nil), histogram.CumulativeBucketCounts...)
	return &copy
}

func cloneUint64(value *uint64) *uint64 {
	if value == nil {
		return nil
	}
	copy := *value
	return &copy
}

func value64(value *uint64) uint64 {
	if value == nil {
		return 0
	}
	return *value
}

func value32(value *uint32) uint32 {
	if value == nil {
		return 0
	}
	return *value
}

func uint64Ptr(value uint64) *uint64 { return &value }

func subtractFloor(total, subset uint64) uint64 {
	if subset > total {
		return 0
	}
	return total - subset
}

func (snapshot StatsSnapshot) HealthSample() observability.HealthSample {
	var signals *HealthSignals
	if snapshot.Health != nil {
		copy := snapshot.Health.Signals
		copy.CriticalFlags = append([]string(nil), snapshot.Health.Signals.CriticalFlags...)
		signals = &copy
	}
	return observability.HealthSample{Header: snapshot.SnapshotHeader, Signals: signals}
}

func EvaluateHealth(previous *StatsSnapshot, current StatsSnapshot, policy HealthPolicy) HealthResult {
	var previousSample *observability.HealthSample
	if previous != nil {
		value := previous.HealthSample()
		previousSample = &value
	}
	return observability.EvaluateHealth(previousSample, current.HealthSample(), policy)
}

// CounterSample projects the fixed Consumer cumulative-counter set consumed by
// observability.DeltaAdapter. It returns false when Operational data was not
// requested or was not collected.
func (snapshot StatsSnapshot) CounterSample() (observability.CounterSample, bool) {
	if snapshot.Operational == nil {
		return observability.CounterSample{}, false
	}
	counters := snapshot.Operational.Counters
	return observability.CounterSample{
		Header: snapshot.SnapshotHeader,
		Counters: map[string]uint64{
			"fetch_calls_success_total":      counters.FetchCallsSuccessTotal,
			"fetch_calls_failure_total":      counters.FetchCallsFailureTotal,
			"fetched_logs_total":             counters.FetchedLogsTotal,
			"fetched_groups_total":           counters.FetchedGroupsTotal,
			"process_attempts_success_total": counters.ProcessAttemptsSuccessTotal,
			"process_attempts_failure_total": counters.ProcessAttemptsFailureTotal,
			"processed_success_logs_total":   counters.ProcessedSuccessLogsTotal,
			"processed_success_groups_total": counters.ProcessedSuccessGroupsTotal,
			"checkpoint_calls_success_total": counters.CheckpointCallsSuccessTotal,
			"checkpoint_calls_failure_total": counters.CheckpointCallsFailureTotal,
			"heartbeat_calls_success_total":  counters.HeartbeatCallsSuccessTotal,
			"heartbeat_calls_failure_total":  counters.HeartbeatCallsFailureTotal,
			"rebalance_total":                counters.RebalanceTotal,
		},
	}, true
}

const consumerContractCapabilities = observability.CapabilityHealth |
	observability.CapabilityOperational |
	observability.CapabilityWorkerDetails |
	observability.CapabilityRetryDetails |
	observability.CapabilityLatencyHistograms |
	observability.CapabilityRuntimeLevelChange
