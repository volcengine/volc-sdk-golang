package consumer

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// newStartedMockConsumer builds and starts a consumer backed by mockConsumerClient
// so Stop() has a valid cancel func and wait group. Returns the consumer and a
// cleanup that stops it (safe to call even if the test stops it early).
func newStartedMockConsumer(t *testing.T, level MetricsLevel) *consumer {
	t.Helper()
	conf := GetDefaultConsumerConfig()
	conf.ProjectID = "test-project"
	conf.TopicIDList = []string{"test-topic"}
	conf.ConsumerGroupName = "test-group"
	conf.ConsumerName = "test-consumer"
	conf.Endpoint = "http://127.0.0.1:0"
	conf.AccessKeyID = "ak"
	conf.AccessKeySecret = "sk"
	conf.Region = "cn-beijing"
	conf.MetricsLevel = level
	conf.DataFetchIntervalInMillisecond = 10
	conf.HeartbeatIntervalInSecond = 1
	conf.FlushCheckpointIntervalSecond = 1

	mc := &mockConsumerClient{}
	c, err := newTestConsumerWithClient(context.Background(), conf, mc, func(topicID string, shardID int, l *pb.LogGroupList) {})
	if err != nil {
		t.Fatalf("newTestConsumerWithClient: %v", err)
	}
	if err := c.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	// Let workers spin up so Stop() has real goroutines to wait on.
	time.Sleep(100 * time.Millisecond)
	return c
}

// TestSnapshotSequence_IncrementsBeforeClose verifies the sequence advances on
// each GetStats while the consumer is running, and never carries the frozen
// high bit.
func TestSnapshotSequence_IncrementsBeforeClose(t *testing.T) {
	c := newStartedMockConsumer(t, MetricsLevelCore)
	defer c.Stop()

	var prev uint64
	for i := 0; i < 5; i++ {
		seq := c.collectMeasurements().SnapshotSequence
		if seq&snapshotSeqFrozenBit != 0 {
			t.Fatalf("snapshot %d carries frozen high bit: %d", i, seq)
		}
		if i > 0 && seq != prev+1 {
			t.Fatalf("sequence did not increment by 1: %d -> %d", prev, seq)
		}
		prev = seq
	}
}

// TestFinalSnapshot_FrozenAfterStop verifies that after Stop() the snapshot is
// immutable: stateClosed, consistencyFinal, and CapturedAt/Uptime/Sequence do
// not change across repeated GetStats calls.
func TestFinalSnapshot_FrozenAfterStop(t *testing.T) {
	c := newStartedMockConsumer(t, MetricsLevelCore)
	c.Stop()

	st1 := c.collectMeasurements()
	if st1.State != stateClosed {
		t.Fatalf("State after Stop = %s, want stateClosed", st1.State)
	}
	if st1.Consistency != consistencyFinal {
		t.Fatalf("Consistency after Stop = %s, want consistencyFinal", st1.Consistency)
	}
	if st1.SnapshotSequence&snapshotSeqFrozenBit != 0 {
		t.Fatalf("final sequence carries frozen high bit: %d", st1.SnapshotSequence)
	}

	// Subsequent snapshots must be byte-for-byte identical (frozen clock + seq).
	for i := 0; i < 5; i++ {
		st2 := c.collectMeasurements()
		if !reflect.DeepEqual(st2, st1) {
			t.Fatalf("snapshot changed after Stop (call %d):\n  got:  %+v\n  want: %+v", i, st2, st1)
		}
	}
}

// TestFinalSnapshot_RepeatedStopUnchanged verifies that calling Stop() more
// than once (idempotent via sync.Once) does not alter the final snapshot.
func TestFinalSnapshot_RepeatedStopUnchanged(t *testing.T) {
	c := newStartedMockConsumer(t, MetricsLevelCore)
	c.Stop()
	st1 := c.collectMeasurements()

	// Repeated Stop must be a no-op that does not rewrite the final clock/seq.
	c.Stop()
	c.Stop()
	st2 := c.collectMeasurements()
	if !reflect.DeepEqual(st2, st1) {
		t.Fatalf("final snapshot changed after repeated Stop:\n  got:  %+v\n  want: %+v", st2, st1)
	}
}

// TestFinalSnapshot_ConcurrentStopUnchanged verifies that concurrent Stop()
// calls (sync.Once) produce a single, stable final snapshot.
func TestFinalSnapshot_ConcurrentStopUnchanged(t *testing.T) {
	c := newStartedMockConsumer(t, MetricsLevelCore)

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Stop()
		}()
	}
	wg.Wait()

	st1 := c.collectMeasurements()
	if st1.State != stateClosed {
		t.Fatalf("State after concurrent Stop = %s, want stateClosed", st1.State)
	}
	if st1.Consistency != consistencyFinal {
		t.Fatalf("Consistency after concurrent Stop = %s, want consistencyFinal", st1.Consistency)
	}
	// Must remain stable.
	st2 := c.collectMeasurements()
	if !reflect.DeepEqual(st2, st1) {
		t.Fatalf("final snapshot changed after concurrent Stop:\n  got:  %+v\n  want: %+v", st2, st1)
	}
}

// TestConcurrentGetStatsStop_NoFrozenBitOrRollback hammers GetStats from many
// goroutines while Stop runs in another. Every observed sequence must:
//   - never carry the frozen high bit,
//   - never go backwards (monotonic),
//   - never exceed the final frozen value.
func TestConcurrentGetStatsStop_NoFrozenBitOrRollback(t *testing.T) {
	c := newStartedMockConsumer(t, MetricsLevelCore)

	var stopWg sync.WaitGroup
	stopWg.Add(1)
	go func() {
		defer stopWg.Done()
		// Small delay so some running snapshots are captured first.
		time.Sleep(20 * time.Millisecond)
		c.Stop()
	}()

	var maxSeq uint64
	var readers sync.WaitGroup
	const numReaders = 16
	for i := 0; i < numReaders; i++ {
		readers.Add(1)
		go func() {
			defer readers.Done()
			var prev uint64
			for j := 0; j < 200; j++ {
				seq := c.collectMeasurements().SnapshotSequence
				if seq&snapshotSeqFrozenBit != 0 {
					t.Errorf("reader observed frozen high bit: %d", seq)
					return
				}
				if seq < prev {
					t.Errorf("reader observed sequence rollback: %d -> %d", prev, seq)
					return
				}
				prev = seq
				atomicMax(&maxSeq, seq)
			}
		}()
	}
	readers.Wait()
	stopWg.Wait()

	finalSeq := c.collectMeasurements().SnapshotSequence
	if finalSeq&snapshotSeqFrozenBit != 0 {
		t.Fatalf("final sequence carries frozen high bit: %d", finalSeq)
	}
	if maxSeq > finalSeq {
		t.Fatalf("reader observed sequence %d exceeding final %d", maxSeq, finalSeq)
	}
}

func atomicMax(dst *uint64, v uint64) {
	for {
		old := atomic.LoadUint64(dst)
		if v <= old || atomic.CompareAndSwapUint64(dst, old, v) {
			return
		}
	}
}

func TestConsumerDetailedPublishesLatencyCapability(t *testing.T) {
	c := newTestConsumer(t, MetricsLevelDetailed)
	st, err := c.GetStats(StatsOptions{View: StatsViewDebug})
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if st.Capabilities&consumerContractCapabilities == 0 || st.LatencyHistograms == nil {
		t.Fatalf("Detailed debug snapshot must expose latency capability and section")
	}
}
