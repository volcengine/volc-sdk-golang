// Package observability defines the public, component-neutral TLS SDK
// observability contract. It deliberately contains no process-runtime metrics
// and starts no background work.
package observability

import "errors"

// SchemaVersion identifies the current cross-language snapshot format.
const SchemaVersion uint32 = 2

type MetricsLevel uint8

const (
	MetricsOff MetricsLevel = iota
	MetricsCore
	MetricsDetailed
)

func (level MetricsLevel) Valid() bool {
	return level == MetricsOff || level == MetricsCore || level == MetricsDetailed
}

type StatsView uint8

const (
	ViewHealth StatsView = iota
	ViewOperational
	ViewDebug
)

func (view StatsView) Valid() bool {
	return view == ViewHealth || view == ViewOperational || view == ViewDebug
}

type StatsOptions struct {
	View StatsView
}

var ErrInvalidStatsView = errors.New("invalid stats view")

type LifecycleState string

const (
	LifecycleNew     LifecycleState = "new"
	LifecycleRunning LifecycleState = "running"
	LifecycleClosing LifecycleState = "closing"
	LifecycleClosed  LifecycleState = "closed"
	LifecycleFailed  LifecycleState = "failed"
)

type ConsistencyLevel string

const (
	ConsistencyWeak  ConsistencyLevel = "weak"
	ConsistencyFinal ConsistencyLevel = "final"
)

type StatsSection string

const (
	SectionHealth                  StatsSection = "health"
	SectionOperational             StatsSection = "operational"
	SectionBatchInflightMemoryGate StatsSection = "batch_inflight_memory_gate"
	SectionManagedMemory           StatsSection = "managed_memory"
	SectionWorkerDetails           StatsSection = "worker_details"
	SectionRetryDetails            StatsSection = "retry_details"
	SectionLatencyHistograms       StatsSection = "latency_histograms"
	SectionMemoryBreakdown         StatsSection = "memory_breakdown"
)

const (
	CapabilityHealth uint64 = 1 << iota
	CapabilityOperational
	CapabilityWorkerDetails
	CapabilityRetryDetails
	CapabilityLatencyHistograms
	CapabilityMemoryBreakdown
	CapabilityManagedMemory
	CapabilityRuntimeLevelChange
	CapabilityBatchInflightMemoryGate
)

type SnapshotHeader struct {
	SchemaVersion             uint32           `json:"schema_version"`
	InstanceID                string           `json:"instance_id"`
	SDKName                   string           `json:"sdk_name"`
	SDKLanguage               string           `json:"sdk_language"`
	SDKVersion                string           `json:"sdk_version"`
	ComponentType             string           `json:"component_type"`
	CapturedAtUnixMs          int64            `json:"captured_at_unix_ms"`
	UptimeMs                  int64            `json:"uptime_ms"`
	SnapshotSequence          uint64           `json:"snapshot_sequence"`
	CollectionGeneration      uint64           `json:"collection_generation"`
	CollectionStartedAtUnixMs int64            `json:"collection_started_at_unix_ms"`
	LifecycleState            LifecycleState   `json:"lifecycle_state"`
	Consistency               ConsistencyLevel `json:"consistency"`
	CollectedLevel            MetricsLevel     `json:"collected_level"`
	RequestedView             StatsView        `json:"requested_view"`
	MetricsDisabled           bool             `json:"metrics_disabled"`
	Partial                   bool             `json:"partial"`
	Capabilities              uint64           `json:"capabilities"`
}

type QueueSummary struct {
	CurrentItems       uint64  `json:"current_items"`
	CapacityItems      *uint64 `json:"capacity_items,omitempty"`
	HighWatermarkItems uint64  `json:"high_watermark_items"`
	Bounded            bool    `json:"bounded"`
}

type WorkerSummary struct {
	Expected             uint32 `json:"expected"`
	Running              uint32 `json:"running"`
	InFlight             uint32 `json:"in_flight"`
	MissingCriticalRoles uint32 `json:"missing_critical_roles"`
}

type MemoryPressureSource string

const (
	MemoryPressureSourceBatchInflight           MemoryPressureSource = "batch_inflight"
	MemoryPressureSourceManagedPayloadAdmission MemoryPressureSource = "managed_payload_admission"
)

type HealthSignals struct {
	LastSuccessAgeMs      *uint64              `json:"last_success_age_ms,omitempty"`
	LastErrorAgeMs        *uint64              `json:"last_error_age_ms,omitempty"`
	QueuePressure         *float64             `json:"queue_pressure,omitempty"`
	ManagedMemoryPressure *float64             `json:"managed_memory_pressure,omitempty"`
	MemoryPressure        *float64             `json:"memory_pressure,omitempty"`
	MemoryPressureSource  MemoryPressureSource `json:"memory_pressure_source,omitempty"`
	OldestPendingAgeMs    *uint64              `json:"oldest_pending_age_ms,omitempty"`
	WorkerSummary         WorkerSummary        `json:"worker_summary"`
	CriticalFlags         []string             `json:"critical_flags"`
}

// MemoryGateSummary exposes a single bounded logical-memory gate.
// UsedBytes may temporarily exceed LimitBytes because the admission check is
// intentionally compatible with the historical post-threshold behavior.
type MemoryGateSummary struct {
	LimitBytes         uint64 `json:"limit_bytes"`
	UsedBytes          uint64 `json:"used_bytes"`
	HighWatermarkBytes uint64 `json:"high_watermark_bytes"`
	RejectedTotal      uint64 `json:"rejected_total"`
}

func (gate MemoryGateSummary) Pressure() *float64 {
	if gate.LimitBytes == 0 {
		return nil
	}
	value := float64(gate.UsedBytes) / float64(gate.LimitBytes)
	return &value
}

type ManagedMemorySummary struct {
	LimitBytes                    uint64 `json:"limit_bytes"`
	PayloadAdmissionLimitBytes    uint64 `json:"payload_admission_limit_bytes"`
	UsedBytes                     uint64 `json:"used_bytes"`
	PayloadOwnedBytes             uint64 `json:"payload_owned_bytes"`
	BoundedRuntimeBufferBytes     uint64 `json:"bounded_runtime_buffer_bytes"`
	HighWatermarkBytes            uint64 `json:"high_watermark_bytes"`
	RejectedTotal                 uint64 `json:"rejected_total"`
	PayloadAdmissionRejectedTotal uint64 `json:"payload_admission_rejected_total"`
	RuntimeBufferRejectedTotal    uint64 `json:"runtime_buffer_rejected_total"`
}

func (memory ManagedMemorySummary) Valid() bool {
	return memory.UsedBytes == memory.PayloadOwnedBytes+memory.BoundedRuntimeBufferBytes &&
		memory.UsedBytes <= memory.LimitBytes &&
		(memory.PayloadAdmissionLimitBytes == 0 || memory.PayloadAdmissionLimitBytes <= memory.LimitBytes) &&
		memory.RejectedTotal == memory.PayloadAdmissionRejectedTotal+memory.RuntimeBufferRejectedTotal
}

// Pressure returns lifecycle-total utilization and preserves the original
// managed_memory_pressure meaning.
func (memory ManagedMemorySummary) Pressure() *float64 {
	if memory.LimitBytes == 0 || !memory.Valid() {
		return nil
	}
	value := float64(memory.UsedBytes) / float64(memory.LimitBytes)
	return &value
}

func (memory ManagedMemorySummary) PayloadAdmissionPressure() *float64 {
	if memory.PayloadAdmissionLimitBytes == 0 || !memory.Valid() {
		return nil
	}
	value := float64(memory.UsedBytes) / float64(memory.PayloadAdmissionLimitBytes)
	return &value
}

type ManagedMemoryBreakdown struct {
	IngressOwnedBytes         uint64 `json:"ingress_owned_bytes"`
	AggregateOwnedBytes       uint64 `json:"aggregate_owned_bytes"`
	SendOwnedBytes            uint64 `json:"send_owned_bytes"`
	RetryOwnedBytes           uint64 `json:"retry_owned_bytes"`
	BoundedRuntimeBufferBytes uint64 `json:"bounded_runtime_buffer_bytes"`
}

func (breakdown ManagedMemoryBreakdown) PayloadOwnedBytes() uint64 {
	return breakdown.IngressOwnedBytes + breakdown.AggregateOwnedBytes +
		breakdown.SendOwnedBytes + breakdown.RetryOwnedBytes
}

// Valid proves that the detailed ownership classes conserve the same managed
// budget represented by summary.
func (breakdown ManagedMemoryBreakdown) Valid(summary ManagedMemorySummary) bool {
	return summary.Valid() &&
		breakdown.PayloadOwnedBytes() == summary.PayloadOwnedBytes &&
		breakdown.BoundedRuntimeBufferBytes == summary.BoundedRuntimeBufferBytes
}

func RequiredLevel(section StatsSection) MetricsLevel {
	switch section {
	case SectionHealth, SectionOperational, SectionBatchInflightMemoryGate, SectionManagedMemory:
		return MetricsCore
	case SectionWorkerDetails, SectionRetryDetails, SectionLatencyHistograms, SectionMemoryBreakdown:
		return MetricsDetailed
	default:
		return MetricsDetailed
	}
}

func RequestedSections(view StatsView) ([]StatsSection, error) {
	switch view {
	case ViewHealth:
		return []StatsSection{SectionHealth, SectionBatchInflightMemoryGate, SectionManagedMemory}, nil
	case ViewOperational:
		return []StatsSection{SectionHealth, SectionOperational, SectionBatchInflightMemoryGate, SectionManagedMemory}, nil
	case ViewDebug:
		return []StatsSection{
			SectionHealth,
			SectionOperational,
			SectionBatchInflightMemoryGate,
			SectionManagedMemory,
			SectionWorkerDetails,
			SectionRetryDetails,
			SectionLatencyHistograms,
			SectionMemoryBreakdown,
		}, nil
	default:
		return nil, ErrInvalidStatsView
	}
}

// MissingSectionsForCapabilities reports requested sections that this
// component supports but the current collection level cannot satisfy.
// Unsupported sections are described by Capabilities and do not make an
// otherwise complete component snapshot partial.
func MissingSectionsForCapabilities(level MetricsLevel, view StatsView, capabilities uint64) ([]StatsSection, error) {
	sections, err := RequestedSections(view)
	if err != nil {
		return nil, err
	}
	missing := make([]StatsSection, 0, len(sections))
	for _, section := range sections {
		if capabilities&sectionCapability(section) != 0 && level < RequiredLevel(section) {
			missing = append(missing, section)
		}
	}
	return missing, nil
}

func Available(level MetricsLevel, section StatsSection) bool {
	return level.Valid() && level >= RequiredLevel(section)
}

func sectionCapability(section StatsSection) uint64 {
	switch section {
	case SectionHealth:
		return CapabilityHealth
	case SectionOperational:
		return CapabilityOperational
	case SectionBatchInflightMemoryGate:
		return CapabilityBatchInflightMemoryGate
	case SectionManagedMemory:
		return CapabilityManagedMemory
	case SectionWorkerDetails:
		return CapabilityWorkerDetails
	case SectionRetryDetails:
		return CapabilityRetryDetails
	case SectionLatencyHistograms:
		return CapabilityLatencyHistograms
	case SectionMemoryBreakdown:
		return CapabilityMemoryBreakdown
	default:
		return 0
	}
}
