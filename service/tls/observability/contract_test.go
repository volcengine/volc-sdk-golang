package observability

import (
	"reflect"
	"testing"
)

func TestMissingSectionsUsesExplicitSectionRequirements(t *testing.T) {
	tests := []struct {
		name  string
		level MetricsLevel
		view  StatsView
		want  []StatsSection
	}{
		{"off_health", MetricsOff, ViewHealth, []StatsSection{SectionHealth, SectionBatchInflightMemoryGate, SectionManagedMemory}},
		{"off_debug", MetricsOff, ViewDebug, []StatsSection{SectionHealth, SectionOperational, SectionBatchInflightMemoryGate, SectionManagedMemory, SectionWorkerDetails, SectionRetryDetails, SectionLatencyHistograms, SectionMemoryBreakdown}},
		{"core_operational", MetricsCore, ViewOperational, []StatsSection{}},
		{"core_debug", MetricsCore, ViewDebug, []StatsSection{SectionWorkerDetails, SectionRetryDetails, SectionLatencyHistograms, SectionMemoryBreakdown}},
		{"detailed_debug", MetricsDetailed, ViewDebug, []StatsSection{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := MissingSectionsForCapabilities(test.level, test.view, ^uint64(0))
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("missing sections = %v, want %v", got, test.want)
			}
		})
	}
}

func TestInvalidStatsView(t *testing.T) {
	if _, err := MissingSectionsForCapabilities(MetricsCore, StatsView(255), ^uint64(0)); err != ErrInvalidStatsView {
		t.Fatalf("error = %v, want %v", err, ErrInvalidStatsView)
	}
}

func TestManagedMemoryConservation(t *testing.T) {
	memory := ManagedMemorySummary{
		LimitBytes:                 200,
		PayloadAdmissionLimitBytes: 125,
		UsedBytes:                  100,
		PayloadOwnedBytes:          70,
		BoundedRuntimeBufferBytes:  30,
	}
	if !memory.Valid() {
		t.Fatal("expected valid memory summary")
	}
	if pressure := memory.Pressure(); pressure == nil || *pressure != 0.5 {
		t.Fatalf("pressure = %v, want 0.5", pressure)
	}
	if pressure := memory.PayloadAdmissionPressure(); pressure == nil || *pressure != 0.8 {
		t.Fatalf("payload admission pressure = %v, want 0.8", pressure)
	}
	memory.PayloadOwnedBytes++
	if memory.Valid() || memory.Pressure() != nil {
		t.Fatal("invalid conservation must not produce a pressure ratio")
	}
}

func TestMemoryGatePressureAllowsCompatibleOvershoot(t *testing.T) {
	gate := MemoryGateSummary{LimitBytes: 100, UsedBytes: 110, HighWatermarkBytes: 110}
	if pressure := gate.Pressure(); pressure == nil || *pressure != 1.1 {
		t.Fatalf("pressure = %v, want 1.1", pressure)
	}
}

func TestEvaluateHealthGatesInsufficientSamples(t *testing.T) {
	current := HealthSample{Header: SnapshotHeader{
		InstanceID:       "instance",
		LifecycleState:   LifecycleRunning,
		CapturedAtUnixMs: 10,
		CollectedLevel:   MetricsCore,
	}}
	result := EvaluateHealth(nil, current, DefaultHealthPolicy())
	if result.Status != HealthUnknown || !result.InsufficientData {
		t.Fatalf("result = %+v", result)
	}
}

func TestEvaluateHealthUsesFactSignals(t *testing.T) {
	pressure := 0.9
	header := SnapshotHeader{
		InstanceID:           "instance",
		LifecycleState:       LifecycleRunning,
		CapturedAtUnixMs:     20,
		CollectedLevel:       MetricsCore,
		CollectionGeneration: 1,
	}
	previous := &HealthSample{Header: header, Signals: &HealthSignals{}}
	current := HealthSample{Header: header, Signals: &HealthSignals{QueuePressure: &pressure}}
	result := EvaluateHealth(previous, current, DefaultHealthPolicy())
	if result.Status != HealthDegraded || !reflect.DeepEqual(result.ReasonCodes, []HealthReason{ReasonQueuePressure}) {
		t.Fatalf("result = %+v", result)
	}
}

func TestEvaluateHealthPrefersEffectiveMemoryPressure(t *testing.T) {
	managed := 0.5
	effective := 0.9
	header := SnapshotHeader{
		InstanceID: "instance", LifecycleState: LifecycleRunning, CapturedAtUnixMs: 20,
		CollectedLevel: MetricsCore, CollectionGeneration: 1,
	}
	previous := &HealthSample{Header: header, Signals: &HealthSignals{}}
	current := HealthSample{Header: header, Signals: &HealthSignals{
		ManagedMemoryPressure: &managed,
		MemoryPressure:        &effective,
	}}
	result := EvaluateHealth(previous, current, DefaultHealthPolicy())
	if result.Status != HealthDegraded || !reflect.DeepEqual(result.ReasonCodes, []HealthReason{ReasonMemoryPressure}) {
		t.Fatalf("result = %+v", result)
	}
}

func TestEvaluateHealthDistinguishesCollectionChange(t *testing.T) {
	previous := &HealthSample{Header: SnapshotHeader{
		InstanceID: "instance", CollectionGeneration: 1, LifecycleState: LifecycleRunning,
	}}
	current := HealthSample{Header: SnapshotHeader{
		InstanceID: "instance", CollectionGeneration: 2, LifecycleState: LifecycleRunning,
	}}
	result := EvaluateHealth(previous, current, DefaultHealthPolicy())
	if !result.InsufficientData || !reflect.DeepEqual(result.ReasonCodes, []HealthReason{ReasonCollectionChanged}) {
		t.Fatalf("result = %+v", result)
	}
}

func TestEvaluateHealthFirstFailedSampleStillRequiresBaseline(t *testing.T) {
	current := HealthSample{Header: SnapshotHeader{
		InstanceID: "instance", CollectionGeneration: 1, LifecycleState: LifecycleFailed, CapturedAtUnixMs: 20,
	}}
	result := EvaluateHealth(nil, current, DefaultHealthPolicy())
	if result.Status != HealthUnknown || !result.InsufficientData ||
		!reflect.DeepEqual(result.ReasonCodes, []HealthReason{ReasonInsufficientData}) {
		t.Fatalf("result = %+v", result)
	}
}

func TestEvaluateHealthFailedLifecycleIsUnhealthyAfterBaseline(t *testing.T) {
	header := SnapshotHeader{
		InstanceID: "instance", CollectionGeneration: 1, LifecycleState: LifecycleRunning, CapturedAtUnixMs: 10,
	}
	current := HealthSample{Header: header}
	current.Header.LifecycleState = LifecycleFailed
	current.Header.CapturedAtUnixMs = 20
	result := EvaluateHealth(&HealthSample{Header: header}, current, DefaultHealthPolicy())
	if result.Status != HealthUnhealthy || result.InsufficientData ||
		!reflect.DeepEqual(result.ReasonCodes, []HealthReason{ReasonLifecycleFailed}) {
		t.Fatalf("result = %+v", result)
	}
}

func TestEvaluateHealthDoesNotJudgeNonRunningLifecycle(t *testing.T) {
	for _, lifecycle := range []LifecycleState{LifecycleNew, LifecycleClosing, LifecycleClosed} {
		t.Run(string(lifecycle), func(t *testing.T) {
			current := HealthSample{Header: SnapshotHeader{
				InstanceID: "instance", LifecycleState: lifecycle, CapturedAtUnixMs: 20,
			}}
			result := EvaluateHealth(&HealthSample{Header: current.Header}, current, DefaultHealthPolicy())
			if result.Status != HealthUnknown || !result.InsufficientData ||
				!reflect.DeepEqual(result.ReasonCodes, []HealthReason{ReasonLifecycleNotRunning}) {
				t.Fatalf("result = %+v", result)
			}
		})
	}
}

func TestEvaluateHealthDoesNotInferFailureFromIdleFacts(t *testing.T) {
	header := SnapshotHeader{
		InstanceID: "instance", CollectionGeneration: 1, LifecycleState: LifecycleRunning, CapturedAtUnixMs: 20,
	}
	previous := &HealthSample{Header: header, Signals: &HealthSignals{}}
	current := HealthSample{Header: header, Signals: &HealthSignals{}}
	result := EvaluateHealth(previous, current, DefaultHealthPolicy())
	if result.Status != HealthHealthy || result.InsufficientData || len(result.ReasonCodes) != 0 {
		t.Fatalf("idle Producer / shardless Consumer facts must not imply failure: %+v", result)
	}
}
