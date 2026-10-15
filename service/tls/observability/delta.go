package observability

import "sync"

// CounterSample is the component-neutral input to rate calculation. Counters
// must be cumulative, non-negative values from one operational snapshot.
type CounterSample struct {
	Header   SnapshotHeader
	Counters map[string]uint64
}

type DeltaResetReason string

const (
	DeltaInitialBaseline    DeltaResetReason = "initial_baseline"
	DeltaMetricsUnavailable DeltaResetReason = "metrics_unavailable"
	DeltaInstanceChanged    DeltaResetReason = "instance_changed"
	DeltaCollectionChanged  DeltaResetReason = "collection_changed"
	DeltaSequenceRegressed  DeltaResetReason = "sequence_not_increasing"
	DeltaUptimeRegressed    DeltaResetReason = "uptime_regressed"
	DeltaClockRegressed     DeltaResetReason = "capture_time_not_increasing"
	DeltaCounterRegressed   DeltaResetReason = "counter_regressed"
	DeltaCounterSetChanged  DeltaResetReason = "counter_set_changed"
)

// DeltaResult describes one interval. Valid is false when the current sample
// became a new baseline; in that case Deltas and RatesPerSecond are absent.
type DeltaResult struct {
	Valid          bool               `json:"valid"`
	ResetReason    DeltaResetReason   `json:"reset_reason,omitempty"`
	IntervalMs     int64              `json:"interval_ms,omitempty"`
	Deltas         map[string]uint64  `json:"deltas,omitempty"`
	RatesPerSecond map[string]float64 `json:"rates_per_second,omitempty"`
}

// DeltaAdapter retains exactly one bounded counter baseline. It starts no
// goroutine and performs no I/O.
type DeltaAdapter struct {
	mu       sync.Mutex
	previous *CounterSample
}

// Observe computes a delta from the preceding valid baseline. Every invalid
// interval establishes current as the next baseline, as required by the
// observability contract.
func (adapter *DeltaAdapter) Observe(current CounterSample) DeltaResult {
	adapter.mu.Lock()
	defer adapter.mu.Unlock()

	current = cloneCounterSample(current)
	if current.Header.MetricsDisabled || current.Counters == nil {
		adapter.previous = &current
		return DeltaResult{ResetReason: DeltaMetricsUnavailable}
	}
	if adapter.previous == nil {
		adapter.previous = &current
		return DeltaResult{ResetReason: DeltaInitialBaseline}
	}

	result := ComputeDelta(*adapter.previous, current)
	adapter.previous = &current
	return result
}

// Reset explicitly discards the current baseline.
func (adapter *DeltaAdapter) Reset() {
	adapter.mu.Lock()
	adapter.previous = nil
	adapter.mu.Unlock()
}

// ComputeDelta is the stateless form of the adapter validation and math.
func ComputeDelta(previous, current CounterSample) DeltaResult {
	if current.Header.MetricsDisabled || previous.Header.MetricsDisabled ||
		current.Counters == nil || previous.Counters == nil {
		return DeltaResult{ResetReason: DeltaMetricsUnavailable}
	}
	if previous.Header.InstanceID == "" || previous.Header.InstanceID != current.Header.InstanceID {
		return DeltaResult{ResetReason: DeltaInstanceChanged}
	}
	if previous.Header.CollectionGeneration != current.Header.CollectionGeneration {
		return DeltaResult{ResetReason: DeltaCollectionChanged}
	}
	if current.Header.SnapshotSequence <= previous.Header.SnapshotSequence {
		return DeltaResult{ResetReason: DeltaSequenceRegressed}
	}
	if current.Header.UptimeMs < previous.Header.UptimeMs {
		return DeltaResult{ResetReason: DeltaUptimeRegressed}
	}
	intervalMs := current.Header.CapturedAtUnixMs - previous.Header.CapturedAtUnixMs
	if intervalMs <= 0 {
		return DeltaResult{ResetReason: DeltaClockRegressed}
	}

	deltas := make(map[string]uint64, len(current.Counters))
	rates := make(map[string]float64, len(current.Counters))
	if len(current.Counters) != len(previous.Counters) {
		return DeltaResult{ResetReason: DeltaCounterSetChanged}
	}
	for name, currentValue := range current.Counters {
		previousValue, ok := previous.Counters[name]
		if !ok {
			return DeltaResult{ResetReason: DeltaCounterSetChanged}
		}
		if currentValue < previousValue {
			return DeltaResult{ResetReason: DeltaCounterRegressed}
		}
		delta := currentValue - previousValue
		deltas[name] = delta
		rates[name] = float64(delta) * 1000 / float64(intervalMs)
	}

	return DeltaResult{
		Valid:          true,
		IntervalMs:     intervalMs,
		Deltas:         deltas,
		RatesPerSecond: rates,
	}
}

func cloneCounterSample(sample CounterSample) CounterSample {
	copy := CounterSample{Header: sample.Header}
	if sample.Counters != nil {
		copy.Counters = make(map[string]uint64, len(sample.Counters))
		for name, value := range sample.Counters {
			copy.Counters[name] = value
		}
	}
	return copy
}
