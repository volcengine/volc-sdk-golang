package observability

type HealthStatus string

const (
	HealthUnknown   HealthStatus = "unknown"
	HealthHealthy   HealthStatus = "healthy"
	HealthDegraded  HealthStatus = "degraded"
	HealthUnhealthy HealthStatus = "unhealthy"
)

type HealthReason string

const (
	ReasonInsufficientData    HealthReason = "insufficient_data"
	ReasonMetricsDisabled     HealthReason = "metrics_disabled"
	ReasonInstanceChanged     HealthReason = "instance_changed"
	ReasonCollectionChanged   HealthReason = "collection_changed"
	ReasonLifecycleNotRunning HealthReason = "lifecycle_not_running"
	ReasonLifecycleFailed     HealthReason = "lifecycle_failed"
	ReasonCriticalFlag        HealthReason = "critical_flag"
	ReasonWorkerMissing       HealthReason = "worker_missing"
	ReasonQueuePressure       HealthReason = "queue_pressure"
	ReasonMemoryPressure      HealthReason = "memory_pressure"
)

type HealthPolicy struct {
	QueuePressureDegraded  float64
	MemoryPressureDegraded float64
}

func DefaultHealthPolicy() HealthPolicy {
	return HealthPolicy{
		QueuePressureDegraded:  0.80,
		MemoryPressureDegraded: 0.80,
	}
}

type HealthResult struct {
	Status            HealthStatus   `json:"status"`
	ReasonCodes       []HealthReason `json:"reason_codes"`
	EvaluatedAtUnixMs int64          `json:"evaluated_at_unix_ms"`
	InsufficientData  bool           `json:"insufficient_data"`
}

type HealthSample struct {
	Header  SnapshotHeader
	Signals *HealthSignals
}

func EvaluateHealth(previous *HealthSample, current HealthSample, policy HealthPolicy) HealthResult {
	result := HealthResult{
		Status:            HealthUnknown,
		ReasonCodes:       []HealthReason{},
		EvaluatedAtUnixMs: current.Header.CapturedAtUnixMs,
	}
	if current.Header.MetricsDisabled {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonMetricsDisabled)
		return result
	}
	if current.Header.LifecycleState == LifecycleNew ||
		current.Header.LifecycleState == LifecycleClosing ||
		current.Header.LifecycleState == LifecycleClosed {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonLifecycleNotRunning)
		return result
	}
	if previous == nil || previous.Header.InstanceID == "" {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonInsufficientData)
		return result
	}
	if previous.Header.MetricsDisabled {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonMetricsDisabled)
		return result
	}
	if previous.Header.InstanceID != current.Header.InstanceID {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonInstanceChanged)
		return result
	}
	if previous.Header.CollectionGeneration != current.Header.CollectionGeneration {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonCollectionChanged)
		return result
	}
	if current.Header.LifecycleState == LifecycleFailed {
		result.Status = HealthUnhealthy
		result.ReasonCodes = append(result.ReasonCodes, ReasonLifecycleFailed)
		return result
	}
	if current.Header.LifecycleState != LifecycleRunning {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonLifecycleNotRunning)
		return result
	}
	if current.Signals == nil {
		result.InsufficientData = true
		result.ReasonCodes = append(result.ReasonCodes, ReasonInsufficientData)
		return result
	}

	if len(current.Signals.CriticalFlags) > 0 {
		result.Status = HealthUnhealthy
		result.ReasonCodes = append(result.ReasonCodes, ReasonCriticalFlag)
		return result
	}
	if current.Signals.WorkerSummary.MissingCriticalRoles > 0 {
		result.Status = HealthUnhealthy
		result.ReasonCodes = append(result.ReasonCodes, ReasonWorkerMissing)
		return result
	}

	result.Status = HealthHealthy
	if current.Signals.QueuePressure != nil &&
		*current.Signals.QueuePressure >= normalizedThreshold(policy.QueuePressureDegraded, 0.80) {
		result.Status = HealthDegraded
		result.ReasonCodes = append(result.ReasonCodes, ReasonQueuePressure)
	}
	memoryPressure := current.Signals.MemoryPressure
	if memoryPressure == nil {
		memoryPressure = current.Signals.ManagedMemoryPressure
	}
	if memoryPressure != nil &&
		*memoryPressure >= normalizedThreshold(policy.MemoryPressureDegraded, 0.80) {
		result.Status = HealthDegraded
		result.ReasonCodes = append(result.ReasonCodes, ReasonMemoryPressure)
	}
	return result
}

func normalizedThreshold(value, fallback float64) float64 {
	if value <= 0 || value > 1 {
		return fallback
	}
	return value
}
