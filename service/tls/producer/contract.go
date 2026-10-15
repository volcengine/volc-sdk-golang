package producer

import (
	"github.com/volcengine/volc-sdk-golang/service/tls/observability"
)

// Shared contract aliases keep Producer configuration ergonomic while the
// schema and evaluator remain component-neutral.
type MetricsLevel = observability.MetricsLevel
type StatsView = observability.StatsView
type StatsOptions = observability.StatsOptions
type SnapshotHeader = observability.SnapshotHeader
type StatsSection = observability.StatsSection
type QueueSummary = observability.QueueSummary
type WorkerSummary = observability.WorkerSummary
type HealthSignals = observability.HealthSignals
type MemoryPressureSource = observability.MemoryPressureSource
type MemoryGateSummary = observability.MemoryGateSummary
type ManagedMemorySummary = observability.ManagedMemorySummary
type ManagedMemoryBreakdown = observability.ManagedMemoryBreakdown
type HealthPolicy = observability.HealthPolicy
type HealthResult = observability.HealthResult

const (
	MetricsLevelOff      = observability.MetricsOff
	MetricsLevelCore     = observability.MetricsCore
	MetricsLevelDetailed = observability.MetricsDetailed

	StatsViewHealth      = observability.ViewHealth
	StatsViewOperational = observability.ViewOperational
	StatsViewDebug       = observability.ViewDebug

	MemoryPressureSourceBatchInflight           = observability.MemoryPressureSourceBatchInflight
	MemoryPressureSourceManagedPayloadAdmission = observability.MemoryPressureSourceManagedPayloadAdmission
)

var ErrInvalidStatsView = observability.ErrInvalidStatsView

// StatsSnapshot is an immutable, view-specific Producer snapshot. Sections
// that were not requested or were not collected are absent rather than filled
// with synthetic zero values.
type StatsSnapshot struct {
	observability.SnapshotHeader
	MissingSections         []observability.StatsSection `json:"missing_sections,omitempty"`
	Health                  *HealthStats                 `json:"health,omitempty"`
	BatchInflightMemoryGate *MemoryGateSummary           `json:"batch_inflight_memory_gate,omitempty"`
	ManagedMemory           *ManagedMemorySummary        `json:"managed_memory,omitempty"`
	Operational             *OperationalStats            `json:"operational,omitempty"`
	WorkerDetails           *WorkerDetails               `json:"worker_details,omitempty"`
	RetryDetails            *RetryDetails                `json:"retry_details,omitempty"`
	LatencyHistograms       *LatencyHistograms           `json:"latency_histograms,omitempty"`
	MemoryBreakdown         *ManagedMemoryBreakdown      `json:"memory_breakdown,omitempty"`
}

type HealthStats struct {
	Signals HealthSignals `json:"signals"`
}

type ProducerCounters struct {
	AcceptedLogsTotal   uint64 `json:"accepted_logs_total"`
	AcceptedGroupsTotal uint64 `json:"accepted_groups_total"`

	TerminalSuccessLogsTotal uint64 `json:"terminal_success_logs_total"`
	TerminalFailureLogsTotal uint64 `json:"terminal_failure_logs_total"`
	TerminalDroppedLogsTotal uint64 `json:"terminal_dropped_logs_total"`

	TerminalSuccessGroupsTotal uint64 `json:"terminal_success_groups_total"`
	TerminalFailureGroupsTotal uint64 `json:"terminal_failure_groups_total"`
	TerminalDroppedGroupsTotal uint64 `json:"terminal_dropped_groups_total"`

	SendCallsSuccessTotal      uint64 `json:"send_calls_success_total"`
	SendCallsFailureTotal      uint64 `json:"send_calls_failure_total"`
	SendCallPayloadLogsTotal   uint64 `json:"send_call_payload_logs_total"`
	SendCallPayloadGroupsTotal uint64 `json:"send_call_payload_groups_total"`
	RetriesScheduledTotal      uint64 `json:"retries_scheduled_total"`

	RejectedInvalidArgumentTotal    uint64 `json:"rejected_invalid_argument_total"`
	RejectedMemoryFullTotal         uint64 `json:"rejected_memory_full_total"`
	RejectedBatchInflightLimitTotal uint64 `json:"rejected_batch_inflight_limit_total"`
	RejectedManagedAdmissionTotal   uint64 `json:"rejected_managed_admission_total"`
	RejectedClosedTotal             uint64 `json:"rejected_closed_total"`
	RejectedCircuitOpenTotal        uint64 `json:"rejected_circuit_open_total"`
}

type OperationalStats struct {
	Counters ProducerCounters `json:"counters"`

	IngressQueue *QueueSummary `json:"ingress_queue,omitempty"`
	SendQueue    *QueueSummary `json:"send_queue,omitempty"`
	RetryQueue   *QueueSummary `json:"retry_queue,omitempty"`

	AggregateBatches  uint64 `json:"aggregate_batches"`
	AggregateGroups   uint64 `json:"aggregate_groups"`
	SendCallsInFlight uint32 `json:"send_calls_in_flight"`
	CallbacksPending  uint64 `json:"callbacks_pending"`
	CallbacksInFlight uint32 `json:"callbacks_in_flight"`

	Workers WorkerSummary `json:"workers"`
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
	Queue                 *QueueSummary `json:"queue,omitempty"`
	RetriesScheduledTotal uint64        `json:"retries_scheduled_total"`
}

// LatencyHistogram freezes the semantic boundary of every observed latency.
type LatencyHistogram struct {
	Start         string             `json:"start"`
	End           string             `json:"end"`
	Unit          string             `json:"unit"`
	Object        string             `json:"object"`
	IncludesRetry bool               `json:"includes_retry"`
	Histogram     *HistogramSnapshot `json:"histogram,omitempty"`
}

type LatencyHistograms struct {
	BatchQueueDelay       LatencyHistogram `json:"batch_queue_delay"`
	SendCallDuration      LatencyHistogram `json:"send_call_duration"`
	BatchDeliveryDuration LatencyHistogram `json:"batch_delivery_duration"`
	CallbackDuration      LatencyHistogram `json:"callback_duration"`
}

// GetStats returns an O(1), no-I/O snapshot and never changes collection.
func (producer *producer) GetStats(options StatsOptions) (StatsSnapshot, error) {
	if !options.View.Valid() {
		return StatsSnapshot{}, ErrInvalidStatsView
	}
	producer.metricsMu.RLock()
	defer producer.metricsMu.RUnlock()
	raw := producer.collectMeasurements()
	batchGate := producer.batchInflightMemoryGateSnapshot(raw)
	memory := producer.managedMemorySnapshot(raw)
	missing, _ := observability.MissingSectionsForCapabilities(raw.MetricsLevel, options.View, producerContractCapabilities)
	if raw.MetricsLevel != MetricsLevelOff && memory.summary == nil {
		missing = appendMissingSection(missing, observability.SectionManagedMemory)
	}
	header := producer.snapshotHeader(raw, options.View, len(missing) != 0)
	snapshot := StatsSnapshot{SnapshotHeader: header, MissingSections: missing}

	if observability.Available(raw.MetricsLevel, observability.SectionHealth) {
		snapshot.Health = &HealthStats{Signals: producer.healthSignals(raw, batchGate, memory.summary)}
	}
	if batchGate != nil {
		snapshot.BatchInflightMemoryGate = batchGate
	}
	if memory.summary != nil {
		snapshot.ManagedMemory = memory.summary
	}
	if options.View >= StatsViewOperational && observability.Available(raw.MetricsLevel, observability.SectionOperational) {
		snapshot.Operational = producer.operationalStats(raw)
	}
	if options.View == StatsViewDebug && raw.MetricsLevel == MetricsLevelDetailed {
		snapshot.WorkerDetails = producer.workerDetails(raw)
		snapshot.RetryDetails = producer.retryDetails(raw)
		snapshot.LatencyHistograms = producer.latencyHistograms(raw)
		if memory.breakdown != nil {
			snapshot.MemoryBreakdown = memory.breakdown
		} else {
			snapshot.Partial = true
			snapshot.MissingSections = appendMissingSection(snapshot.MissingSections, observability.SectionMemoryBreakdown)
		}
	}
	return snapshot, nil
}

func (producer *producer) snapshotHeader(raw producerMeasurements, view StatsView, partial bool) SnapshotHeader {
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
		CollectionGeneration:      producer.collectionGenerationNow(),
		CollectionStartedAtUnixMs: producer.collectionStartedAtNow(),
		LifecycleState:            observability.LifecycleState(raw.State),
		Consistency:               observability.ConsistencyLevel(raw.Consistency),
		CollectedLevel:            raw.MetricsLevel,
		RequestedView:             view,
		MetricsDisabled:           raw.MetricsDisabled,
		Partial:                   partial,
		Capabilities:              producerContractCapabilities,
	}
}

func (producer *producer) healthSignals(raw producerMeasurements, batchGate *MemoryGateSummary, memory *ManagedMemorySummary) HealthSignals {
	workers := summarizeWorkers(raw.WorkerRoles)
	queuePressure := maxQueuePressure(raw.IngressQueue, raw.SendQueue)
	var managedMemoryPressure *float64
	if memory != nil {
		managedMemoryPressure = memory.Pressure()
	}
	memoryPressure, memoryPressureSource := maxMemoryPressure(batchGate, memory)
	return HealthSignals{
		LastSuccessAgeMs:      cloneUint64(raw.LastSuccessAgeMs),
		LastErrorAgeMs:        cloneUint64(raw.LastErrorAgeMs),
		QueuePressure:         queuePressure,
		ManagedMemoryPressure: managedMemoryPressure,
		MemoryPressure:        memoryPressure,
		MemoryPressureSource:  memoryPressureSource,
		WorkerSummary:         workers,
		CriticalFlags:         []string{},
	}
}

func (producer *producer) operationalStats(raw producerMeasurements) *OperationalStats {
	return &OperationalStats{
		Counters: ProducerCounters{
			AcceptedLogsTotal:               value64(raw.AcceptedLogs),
			AcceptedGroupsTotal:             value64(raw.AcceptedGroups),
			TerminalSuccessLogsTotal:        value64(raw.TerminalSuccessLogs),
			TerminalFailureLogsTotal:        value64(raw.TerminalFailureLogs),
			TerminalDroppedLogsTotal:        value64(raw.TerminalDroppedLogs),
			TerminalSuccessGroupsTotal:      value64(raw.TerminalSuccessGroups),
			TerminalFailureGroupsTotal:      value64(raw.TerminalFailureGroups),
			TerminalDroppedGroupsTotal:      value64(raw.TerminalDroppedGroups),
			SendCallsSuccessTotal:           value64(raw.SendCallsSuccess),
			SendCallsFailureTotal:           value64(raw.SendCallsFailure),
			SendCallPayloadLogsTotal:        value64(raw.SendCallPayloadLogs),
			SendCallPayloadGroupsTotal:      value64(raw.SendCallPayloadGroups),
			RetriesScheduledTotal:           value64(raw.RetriesScheduled),
			RejectedInvalidArgumentTotal:    value64(raw.RejectedInvalidArgument),
			RejectedMemoryFullTotal:         value64(raw.RejectedMemoryFull),
			RejectedBatchInflightLimitTotal: value64(raw.RejectedBatchInflightLimit),
			RejectedManagedAdmissionTotal:   value64(raw.RejectedManagedAdmission),
			RejectedClosedTotal:             value64(raw.RejectedClosed),
			RejectedCircuitOpenTotal:        value64(raw.RejectedCircuitOpen),
		},
		IngressQueue:      projectQueue(raw.IngressQueue),
		SendQueue:         projectQueue(raw.SendQueue),
		RetryQueue:        projectQueue(raw.RetryQueue),
		AggregateBatches:  value64(raw.AggregateBatches),
		AggregateGroups:   value64(raw.AggregateGroups),
		SendCallsInFlight: value32(raw.SendCallsInFlight),
		CallbacksPending:  value64(raw.CallbacksPending),
		CallbacksInFlight: value32(raw.CallbacksInFlight),
		Workers:           summarizeWorkers(raw.WorkerRoles),
	}
}

func (producer *producer) batchInflightMemoryGateSnapshot(raw producerMeasurements) *MemoryGateSummary {
	if raw.BatchInflightMemoryGate == nil {
		return nil
	}
	highWatermark := raw.BatchInflightMemoryGate.HighWatermarkBytes
	if highWatermark < raw.BatchInflightMemoryGate.UsedBytes {
		highWatermark = raw.BatchInflightMemoryGate.UsedBytes
	}
	return &MemoryGateSummary{
		LimitBytes:         raw.BatchInflightMemoryGate.LimitBytes,
		UsedBytes:          raw.BatchInflightMemoryGate.UsedBytes,
		HighWatermarkBytes: highWatermark,
		RejectedTotal:      raw.BatchInflightMemoryGate.RejectedTotal,
	}
}

func maxMemoryPressure(batchGate *MemoryGateSummary, memory *ManagedMemorySummary) (*float64, observability.MemoryPressureSource) {
	var pressure *float64
	var source observability.MemoryPressureSource
	if batchGate != nil {
		pressure = batchGate.Pressure()
		if pressure != nil && *pressure > 0 {
			source = observability.MemoryPressureSourceBatchInflight
		}
	}
	if memory != nil {
		if admissionPressure := memory.PayloadAdmissionPressure(); admissionPressure != nil &&
			(pressure == nil || *admissionPressure > *pressure) {
			pressure = admissionPressure
			if *admissionPressure > 0 {
				source = observability.MemoryPressureSourceManagedPayloadAdmission
			} else {
				source = ""
			}
		}
	}
	return pressure, source
}

type producerMemorySnapshot struct {
	summary   *ManagedMemorySummary
	breakdown *ManagedMemoryBreakdown
}

// managedMemorySnapshot captures one internally conserved memory view for all
// sections projected by a GetStats call. The retry count is fixed, preserving
// O(1) behavior independently of workload and backlog size.
func (producer *producer) managedMemorySnapshot(raw producerMeasurements) producerMemorySnapshot {
	if raw.Memory == nil {
		return producerMemorySnapshot{}
	}
	limit := raw.Memory.AccountedLimitBytes
	makeSummary := func(used, payload, temporary uint64) *ManagedMemorySummary {
		highWatermark := raw.Memory.AccountedHighWatermarkBytes
		if highWatermark < used {
			highWatermark = used
		}
		return &ManagedMemorySummary{
			LimitBytes:                    limit,
			PayloadAdmissionLimitBytes:    raw.Memory.PayloadAdmissionLimitBytes,
			UsedBytes:                     used,
			PayloadOwnedBytes:             payload,
			BoundedRuntimeBufferBytes:     temporary,
			HighWatermarkBytes:            highWatermark,
			RejectedTotal:                 raw.Memory.AccountingRejectedTotal,
			PayloadAdmissionRejectedTotal: raw.Memory.PayloadAdmissionRejectedTotal,
			RuntimeBufferRejectedTotal:    raw.Memory.RuntimeBufferRejectedTotal,
		}
	}
	if producer != nil && producer.memoryLimiter != nil {
		for attempt := 0; attempt < 8; attempt++ {
			ingress, aggregate, send, retry, ok := producer.memoryLimiter.payloadBreakdown()
			if !ok || ingress < 0 || aggregate < 0 || send < 0 || retry < 0 {
				continue
			}
			payload := ingress + aggregate + send + retry
			used := producer.memoryLimiter.Used()
			if used < payload || used < 0 || uint64(used) > limit {
				continue
			}
			temporary := producer.memoryLimiter.TemporaryUsed()
			if temporary < 0 || used != payload+temporary {
				continue
			}
			return producerMemorySnapshot{
				summary: makeSummary(uint64(used), uint64(payload), uint64(temporary)),
				breakdown: &ManagedMemoryBreakdown{
					IngressOwnedBytes:         uint64(ingress),
					AggregateOwnedBytes:       uint64(aggregate),
					SendOwnedBytes:            uint64(send),
					RetryOwnedBytes:           uint64(retry),
					BoundedRuntimeBufferBytes: uint64(temporary),
				},
			}
		}
	}

	// A summary is safe only when the independently observed ownership classes
	// conserve the functional budget. In particular, an Off -> On transition
	// with live allocations stays partial until the unknown ownership drains.
	used := raw.Memory.AccountedUsedBytes
	payload := raw.Memory.PayloadUsedBytes
	temporary := raw.Memory.TemporaryUsedBytes
	if used <= limit && used == payload+temporary {
		return producerMemorySnapshot{summary: makeSummary(used, payload, temporary)}
	}
	return producerMemorySnapshot{}
}

func (producer *producer) workerDetails(raw producerMeasurements) *WorkerDetails {
	items := make([]WorkerDetail, 0, len(raw.WorkerRoles))
	for _, role := range raw.WorkerRoles {
		items = append(items, WorkerDetail{Role: string(role.Role), Expected: role.Expected, Running: role.Running, InFlight: role.InFlight})
	}
	return &WorkerDetails{Items: items, Truncated: false}
}

func (producer *producer) retryDetails(raw producerMeasurements) *RetryDetails {
	return &RetryDetails{Queue: projectQueue(raw.RetryQueue), RetriesScheduledTotal: value64(raw.RetriesScheduled)}
}

func (producer *producer) latencyHistograms(raw producerMeasurements) *LatencyHistograms {
	return &LatencyHistograms{
		BatchQueueDelay:       LatencyHistogram{Start: "batch_ready", End: "send_attempt_start", Unit: "ms", Object: "batch", IncludesRetry: false, Histogram: cloneHistogram(raw.BatchQueueDelayMs)},
		SendCallDuration:      LatencyHistogram{Start: "transport_call_start", End: "transport_call_end", Unit: "ms", Object: "call", IncludesRetry: false, Histogram: cloneHistogram(raw.SendCallDurationMs)},
		BatchDeliveryDuration: LatencyHistogram{Start: "first_send_attempt_start", End: "terminal_outcome", Unit: "ms", Object: "batch", IncludesRetry: true, Histogram: cloneHistogram(raw.BatchDeliveryDurationMs)},
		CallbackDuration:      LatencyHistogram{Start: "callback_start", End: "callback_end", Unit: "ms", Object: "callback", IncludesRetry: false, Histogram: cloneHistogram(raw.CallbackDurationMs)},
	}
}

func appendMissingSection(sections []StatsSection, section StatsSection) []StatsSection {
	for _, existing := range sections {
		if existing == section {
			return sections
		}
	}
	return append(sections, section)
}

func summarizeWorkers(roles []workerRoleMeasurements) WorkerSummary {
	var summary WorkerSummary
	for _, role := range roles {
		summary.InFlight += role.InFlight
		// Sender goroutines are demand-driven concurrency, not long-lived
		// critical workers; idle senders must not make health unhealthy.
		if role.Role == roleSender {
			continue
		}
		summary.Expected += role.Expected
		summary.Running += role.Running
		if role.Running < role.Expected {
			summary.MissingCriticalRoles++
		}
	}
	return summary
}

func maxQueuePressure(queues ...*queueMeasurements) *float64 {
	var max float64
	found := false
	for _, queue := range queues {
		if queue == nil || !queue.Bounded || queue.CapacityItems == 0 {
			continue
		}
		value := float64(queue.CurrentItems) / float64(queue.CapacityItems)
		if !found || value > max {
			max = value
		}
		found = true
	}
	if !found {
		return nil
	}
	return &max
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

// CounterSample projects the fixed Producer cumulative-counter set consumed by
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
			"accepted_logs_total":                 counters.AcceptedLogsTotal,
			"accepted_groups_total":               counters.AcceptedGroupsTotal,
			"terminal_success_logs_total":         counters.TerminalSuccessLogsTotal,
			"terminal_failure_logs_total":         counters.TerminalFailureLogsTotal,
			"terminal_dropped_logs_total":         counters.TerminalDroppedLogsTotal,
			"terminal_success_groups_total":       counters.TerminalSuccessGroupsTotal,
			"terminal_failure_groups_total":       counters.TerminalFailureGroupsTotal,
			"terminal_dropped_groups_total":       counters.TerminalDroppedGroupsTotal,
			"send_calls_success_total":            counters.SendCallsSuccessTotal,
			"send_calls_failure_total":            counters.SendCallsFailureTotal,
			"send_call_payload_logs_total":        counters.SendCallPayloadLogsTotal,
			"send_call_payload_groups_total":      counters.SendCallPayloadGroupsTotal,
			"retries_scheduled_total":             counters.RetriesScheduledTotal,
			"rejected_invalid_argument_total":     counters.RejectedInvalidArgumentTotal,
			"rejected_memory_full_total":          counters.RejectedMemoryFullTotal,
			"rejected_batch_inflight_limit_total": counters.RejectedBatchInflightLimitTotal,
			"rejected_managed_admission_total":    counters.RejectedManagedAdmissionTotal,
			"rejected_closed_total":               counters.RejectedClosedTotal,
			"rejected_circuit_open_total":         counters.RejectedCircuitOpenTotal,
		},
	}, true
}

const producerContractCapabilities = observability.CapabilityHealth |
	observability.CapabilityOperational |
	observability.CapabilityWorkerDetails |
	observability.CapabilityRetryDetails |
	observability.CapabilityLatencyHistograms |
	observability.CapabilityMemoryBreakdown |
	observability.CapabilityManagedMemory |
	observability.CapabilityBatchInflightMemoryGate |
	observability.CapabilityRuntimeLevelChange
