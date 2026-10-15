package producer

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newRuntimeMetricsTestProducer() *producer {
	config := GetDefaultProducerConfig()
	config.Endpoint = "http://127.0.0.1:0"
	config.Region = "cn-beijing"
	return newProducer(config)
}

func TestProducerRuntimeMetricsLevelSwitch(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	assert.True(t, p.collectMeasurements().MetricsDisabled)

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	stats := p.collectMeasurements()
	assert.Equal(t, MetricsLevelCore, stats.MetricsLevel)
	assert.False(t, stats.MetricsDisabled)
	assert.NotNil(t, stats.SendQueue)

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelOff))
	stats = p.collectMeasurements()
	assert.Equal(t, MetricsLevelOff, stats.MetricsLevel)
	assert.True(t, stats.MetricsDisabled)
	assert.Nil(t, stats.SendQueue)
}

func TestProducerRuntimeMetricsLevelOffStopsGaugeCollection(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	if p.gaugesTracked() {
		p.stats.incCallbacksPending()
		p.stats.incSendCallsInFlight()
	}

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	stats := p.collectMeasurements()
	assert.NotNil(t, stats.CallbacksPending)
	assert.NotNil(t, stats.SendCallsInFlight)
	assert.Equal(t, uint64(0), *stats.CallbacksPending)
	assert.Equal(t, uint32(0), *stats.SendCallsInFlight)
}

func TestProducerRuntimeEnableInitializesBatchInflightGate(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	atomic.StoreInt64(&p.producerLogGroupSize, 1234)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	snapshot, err := p.GetStats(StatsOptions{View: StatsViewHealth})
	assert.NoError(t, err)
	if assert.NotNil(t, snapshot.BatchInflightMemoryGate) {
		assert.Equal(t, uint64(1234), snapshot.BatchInflightMemoryGate.UsedBytes)
		assert.Equal(t, uint64(1234), snapshot.BatchInflightMemoryGate.HighWatermarkBytes)
	}

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelOff))
	atomic.StoreInt64(&p.producerLogGroupSize, 2345)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	snapshot, err = p.GetStats(StatsOptions{View: StatsViewHealth})
	assert.NoError(t, err)
	if assert.NotNil(t, snapshot.BatchInflightMemoryGate) {
		assert.Equal(t, uint64(2345), snapshot.BatchInflightMemoryGate.UsedBytes)
		assert.Equal(t, uint64(2345), snapshot.BatchInflightMemoryGate.HighWatermarkBytes,
			"new generation HWM must start from current functional usage")
	}
	atomic.StoreInt64(&p.producerLogGroupSize, 0)
}

func TestProducerSendCallObservationIsCollectionGenerationBound(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	oldGeneration := p.beginSendCallObservation()
	assert.NotZero(t, oldGeneration)
	assert.Equal(t, int64(1), atomic.LoadInt64(&p.stats.sendCallsInFlight))

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelOff))
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	p.finishSendCallObservation(oldGeneration, false, 7, 2)

	stats := p.collectMeasurements()
	assert.Zero(t, *stats.SendCallsInFlight)
	assert.Zero(t, *stats.SendCallsSuccess)
	assert.Zero(t, *stats.SendCallPayloadLogs)
	assert.Zero(t, *stats.SendCallPayloadGroups)

	currentGeneration := p.beginSendCallObservation()
	p.finishSendCallObservation(currentGeneration, false, 7, 2)
	stats = p.collectMeasurements()
	assert.Equal(t, uint64(1), *stats.SendCallsSuccess)
	assert.Equal(t, uint64(7), *stats.SendCallPayloadLogs)
	assert.Equal(t, uint64(2), *stats.SendCallPayloadGroups)
}

func TestProducerRuntimeMetricsLevelResetsHealthAges(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	p.stats.setLastSuccessTime(time.Now().Add(-time.Second).UnixNano())
	p.stats.setLastErrorTime(time.Now().Add(-time.Second).UnixNano())
	before := p.collectMeasurements()
	assert.NotNil(t, before.LastSuccessAgeMs)
	assert.NotNil(t, before.LastErrorAgeMs)

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelOff))
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	after := p.collectMeasurements()
	assert.Nil(t, after.LastSuccessAgeMs)
	assert.Nil(t, after.LastErrorAgeMs)
}

func TestProducerForceCloseDoesNotAttributePreviousGenerationOutcomes(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	// Keep workers stopped so one accepted log remains in ingress while level
	// changes. ForceClose still exercises the real drain path.
	atomic.StoreInt32(&p.started, 1)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	assert.NoError(t, p.putToDispatcher(&BatchLog{
		Key: BatchKey{Topic: "topic", Source: "source", FileName: "file"},
		Log: makeLog(),
	}))
	assert.Equal(t, int64(1), atomic.LoadInt64(&p.stats.acceptedLogs))

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelOff))
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	beforeDropped := atomic.LoadInt64(&p.stats.terminalDroppedLogs)
	p.ForceClose()

	assert.Equal(t, beforeDropped, atomic.LoadInt64(&p.stats.terminalDroppedLogs),
		"an old-generation admission must not become a new-generation terminal outcome")
}

func TestProducerRuntimeMetricsLevelChangesCollectionGeneration(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	initial, err := p.GetStats(StatsOptions{View: StatsViewHealth})
	assert.NoError(t, err)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	core, err := p.GetStats(StatsOptions{View: StatsViewHealth})
	assert.NoError(t, err)
	assert.Greater(t, core.CollectionGeneration, initial.CollectionGeneration)
	assert.Greater(t, core.CollectionStartedAtUnixMs, int64(0))
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelOff))
	off, err := p.GetStats(StatsOptions{View: StatsViewHealth})
	assert.NoError(t, err)
	assert.Greater(t, off.CollectionGeneration, core.CollectionGeneration)
	assert.Zero(t, off.CollectionStartedAtUnixMs)
}

func TestProducerCoreDetailedSwitchPreservesCollectionGeneration(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	core, err := p.GetStats(StatsOptions{View: StatsViewOperational})
	assert.NoError(t, err)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelDetailed))
	detailed, err := p.GetStats(StatsOptions{View: StatsViewDebug})
	assert.NoError(t, err)
	assert.Equal(t, core.CollectionGeneration, detailed.CollectionGeneration)
	assert.Equal(t, core.CollectionStartedAtUnixMs, detailed.CollectionStartedAtUnixMs)
	assert.NotNil(t, detailed.LatencyHistograms)
	firstDetailGeneration := p.currentDetailGeneration()
	assert.NotZero(t, firstDetailGeneration)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))
	assert.False(t, p.detailGenerationEnabled(firstDetailGeneration), "an interval crossing a Core window must be discarded")
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelDetailed))
	assert.Greater(t, p.currentDetailGeneration(), firstDetailGeneration)
	after, err := p.GetStats(StatsOptions{View: StatsViewDebug})
	assert.NoError(t, err)
	assert.Equal(t, core.CollectionGeneration, after.CollectionGeneration)
}

func TestProducerRuntimeMetricsLevelMarksIngressBeforeDispatcherHandoff(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	// Keep the dispatcher under test control: SetMetricsLevel requires running,
	// but no background goroutine is needed to exercise the channel handoff.
	atomic.StoreInt32(&p.started, 1)
	assert.NoError(t, p.SetMetricsLevel(MetricsLevelCore))

	batchLog := &BatchLog{
		Key: BatchKey{Topic: "topic", Source: "source", FileName: "file"},
		Log: makeLog(),
	}
	assert.NoError(t, p.putToDispatcher(batchLog))
	p.dispatcher.handleLogs(<-p.dispatcher.newLogRecvChan)

	stats := p.collectMeasurements()
	assert.Equal(t, uint64(1), *stats.AcceptedLogs)
	assert.Equal(t, uint64(1), *stats.AcceptedGroups)
}

func TestProducerRuntimeMetricsLevelRequiresRunning(t *testing.T) {
	p := newTestProducer()
	assert.ErrorIs(t, p.SetMetricsLevel(MetricsLevelCore), ErrMetricsLevelChangeNotRunning)
	assert.ErrorIs(t, p.SetMetricsLevel(MetricsLevel(99)), ErrInvalidMetricsLevel)
}

func TestProducerRuntimeMetricsLevelConcurrentWithSnapshot(t *testing.T) {
	p := newRuntimeMetricsTestProducer()
	p.Start()
	defer p.ForceClose()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for n := 0; n < 200; n++ {
				level := MetricsLevelOff
				if (n+i)%2 == 0 {
					level = MetricsLevelCore
				}
				if err := p.SetMetricsLevel(level); err != nil && !errors.Is(err, ErrMetricsLevelChangeClosing) {
					t.Errorf("SetMetricsLevel(%v): %v", level, err)
					return
				}
			}
		}(i)
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < 200; n++ {
				snapshot, err := p.GetStats(StatsOptions{View: StatsViewOperational})
				if err != nil {
					t.Errorf("GetStats: %v", err)
					return
				}
				if snapshot.CollectedLevel == MetricsLevelOff {
					if !snapshot.MetricsDisabled || snapshot.CollectionStartedAtUnixMs != 0 || snapshot.Health != nil || snapshot.Operational != nil {
						t.Errorf("inconsistent Off snapshot: %+v", snapshot)
						return
					}
				} else if snapshot.MetricsDisabled || snapshot.CollectionStartedAtUnixMs <= 0 || snapshot.Health == nil || snapshot.Operational == nil {
					t.Errorf("inconsistent enabled snapshot: %+v", snapshot)
					return
				}
			}
		}()
	}
	wg.Wait()

	// Ensure the test actually exercised both paths rather than only snapshots.
	assert.Greater(t, atomic.LoadUint64(&p.stats.snapshotSeq), uint64(0))
}
