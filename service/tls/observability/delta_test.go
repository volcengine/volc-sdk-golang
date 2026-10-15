package observability

import "testing"

func deltaSample(instance string, generation, sequence uint64, capturedAt, uptime int64, value uint64) CounterSample {
	return CounterSample{
		Header: SnapshotHeader{
			InstanceID:           instance,
			CollectionGeneration: generation,
			SnapshotSequence:     sequence,
			CapturedAtUnixMs:     capturedAt,
			UptimeMs:             uptime,
		},
		Counters: map[string]uint64{"requests_total": value},
	}
}

func TestDeltaAdapterComputesRateAndCopiesBaseline(t *testing.T) {
	adapter := &DeltaAdapter{}
	first := deltaSample("one", 1, 1, 1000, 1000, 10)
	if result := adapter.Observe(first); result.Valid || result.ResetReason != DeltaInitialBaseline {
		t.Fatalf("first result = %+v", result)
	}
	first.Counters["requests_total"] = 999 // must not mutate retained baseline
	result := adapter.Observe(deltaSample("one", 1, 2, 3000, 3000, 16))
	if !result.Valid || result.IntervalMs != 2000 || result.Deltas["requests_total"] != 6 ||
		result.RatesPerSecond["requests_total"] != 3 {
		t.Fatalf("delta result = %+v", result)
	}
}

func TestDeltaAdapterRejectsInvalidIntervalsAndRebaselines(t *testing.T) {
	tests := []struct {
		name   string
		second CounterSample
		want   DeltaResetReason
	}{
		{"instance", deltaSample("two", 1, 2, 2000, 2000, 11), DeltaInstanceChanged},
		{"generation", deltaSample("one", 2, 2, 2000, 2000, 11), DeltaCollectionChanged},
		{"sequence", deltaSample("one", 1, 1, 2000, 2000, 11), DeltaSequenceRegressed},
		{"uptime", deltaSample("one", 1, 2, 2000, 500, 11), DeltaUptimeRegressed},
		{"clock", deltaSample("one", 1, 2, 1000, 2000, 11), DeltaClockRegressed},
		{"counter", deltaSample("one", 1, 2, 2000, 2000, 9), DeltaCounterRegressed},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			adapter := &DeltaAdapter{}
			adapter.Observe(deltaSample("one", 1, 1, 1000, 1000, 10))
			result := adapter.Observe(test.second)
			if result.Valid || result.ResetReason != test.want {
				t.Fatalf("result = %+v, want reset %q", result, test.want)
			}
		})
	}
}

func TestDeltaAdapterTreatsOffAndCounterSetChangesAsNewBaselines(t *testing.T) {
	adapter := &DeltaAdapter{}
	off := deltaSample("one", 1, 1, 1000, 1000, 0)
	off.Header.MetricsDisabled = true
	off.Counters = nil
	if result := adapter.Observe(off); result.Valid || result.ResetReason != DeltaMetricsUnavailable {
		t.Fatalf("Off result = %+v", result)
	}
	on := deltaSample("one", 2, 2, 2000, 2000, 10)
	if result := adapter.Observe(on); result.Valid || result.ResetReason != DeltaMetricsUnavailable {
		t.Fatalf("first enabled result after Off = %+v", result)
	}
	changed := deltaSample("one", 2, 3, 3000, 3000, 12)
	changed.Counters["new_total"] = 1
	if result := adapter.Observe(changed); result.Valid || result.ResetReason != DeltaCounterSetChanged {
		t.Fatalf("counter-set result = %+v", result)
	}
}
