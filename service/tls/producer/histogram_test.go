package producer

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestHistogram_OffReturnsNil(t *testing.T) {
	h := newHistogramState(histogramProfileOff, 4)
	if h != nil {
		t.Fatalf("Off profile must not allocate histogram, got %v", h)
	}
}

func TestHistogram_FastV1Buckets(t *testing.T) {
	h := newHistogramState(histogramProfileFastV1, 4)
	if h == nil {
		t.Fatalf("FastV1 must allocate histogram")
	}
	if len(h.bounds) != len(fastV1BoundsMs) {
		t.Fatalf("fast_v1 bounds len=%d, want %d", len(h.bounds), len(fastV1BoundsMs))
	}
	if h.numShards != 4 {
		t.Fatalf("numShards=%d, want 4", h.numShards)
	}
	// 13 bounds -> 14 buckets (13 + +Inf) per shard
	if len(h.shards[0].counts) != len(fastV1BoundsMs)+1 {
		t.Fatalf("fast_v1 shard counts len=%d, want %d", len(h.shards[0].counts), len(fastV1BoundsMs)+1)
	}
}

func TestHistogram_ObserveBucketing(t *testing.T) {
	h := newHistogramState(histogramProfileFastV1, 2)
	// fast_v1 bounds: [1,2,5,10,...]; 0.5ms -> bucket 0 (≤1), 1.5ms -> bucket 1 (≤2)
	h.observe(0.5, 0)
	h.observe(1.5, 1)
	h.observe(100000, 0) // above last bound -> +Inf bucket

	snap := h.snapshot()
	if snap == nil {
		t.Fatalf("snapshot nil")
	}
	if snap.Count != 3 {
		t.Fatalf("count=%d, want 3", snap.Count)
	}
	// cumulative: bucket0=1, bucket1=2, ..., last(+Inf)=3
	if snap.CumulativeBucketCounts[0] != 1 {
		t.Fatalf("cum[0]=%d, want 1", snap.CumulativeBucketCounts[0])
	}
	if snap.CumulativeBucketCounts[1] != 2 {
		t.Fatalf("cum[1]=%d, want 2", snap.CumulativeBucketCounts[1])
	}
	last := len(snap.CumulativeBucketCounts) - 1
	if snap.CumulativeBucketCounts[last] != 3 {
		t.Fatalf("+Inf bucket=%d, want 3 (==count)", snap.CumulativeBucketCounts[last])
	}
	if snap.SumMs != 0.5+1.5+100000 {
		t.Fatalf("sum_ms=%v, want %v", snap.SumMs, 0.5+1.5+100000)
	}
}

func TestHistogram_SnapshotCumulativeAndInfEqualsCount(t *testing.T) {
	h := newHistogramState(histogramProfileSlowV1, 3)
	for i := 0; i < 100; i++ {
		h.observe(float64(i), i%3)
	}
	snap := h.snapshot()
	if snap.Count != 100 {
		t.Fatalf("count=%d, want 100", snap.Count)
	}
	// cumulative must be non-decreasing
	for i := 1; i < len(snap.CumulativeBucketCounts); i++ {
		if snap.CumulativeBucketCounts[i] < snap.CumulativeBucketCounts[i-1] {
			t.Fatalf("cumulative not non-decreasing at %d", i)
		}
	}
	// +Inf bucket == count
	last := len(snap.CumulativeBucketCounts) - 1
	if snap.CumulativeBucketCounts[last] != snap.Count {
		t.Fatalf("+Inf=%d, count=%d, must be equal", snap.CumulativeBucketCounts[last], snap.Count)
	}
}

func TestHistogram_SnapshotIsImmutableCopy(t *testing.T) {
	h := newHistogramState(histogramProfileFastV1, 1)
	h.observe(5, 0)
	snap1 := h.snapshot()
	h.observe(5, 0) // mutate after snapshot
	snap2 := h.snapshot()
	if snap1.Count == snap2.Count {
		t.Fatalf("snapshot must be a copy; snap1=%d snap2=%d", snap1.Count, snap2.Count)
	}
	// mutating snap1 must not affect snap2 or live state
	snap1.CumulativeBucketCounts[0] = 999
	if snap2.CumulativeBucketCounts[0] == 999 {
		t.Fatalf("snapshots share mutable state")
	}
}

func TestHistogram_ObserveSince(t *testing.T) {
	h := newHistogramState(histogramProfileFastV1, 1)
	start := time.Now()
	time.Sleep(2 * time.Millisecond)
	h.observeSince(start, 0)
	snap := h.snapshot()
	if snap.Count != 1 {
		t.Fatalf("count=%d, want 1", snap.Count)
	}
	if snap.SumMs < 1.0 {
		t.Fatalf("sum_ms=%v, expected >= ~2ms", snap.SumMs)
	}
}

func TestHistogram_NilSafe(t *testing.T) {
	var h *histogramState
	h.observe(1.0, 0)             // must not panic
	h.observeSince(time.Now(), 0) // must not panic
	if snap := h.snapshot(); snap != nil {
		t.Fatalf("nil histogram snapshot must be nil, got %v", snap)
	}
}

// TestHistogram_SenderIdxModuloWrapping verifies that out-of-range senderIdx
// values map to a valid shard via modulo wrapping (not clamping to the last
// shard), so no observation is lost. Negative indices also wrap correctly.
// Asserts the exact per-shard count distribution, not just the aggregate.
func TestHistogram_SenderIdxModuloWrapping(t *testing.T) {
	h := newHistogramState(histogramProfileFastV1, 2)
	h.observe(1.0, -1) // -1 % 2 == -1, +2 == 1 -> shard 1
	h.observe(2.0, 99) // 99 % 2 == 1 -> shard 1
	h.observe(3.0, 1)  // shard 1
	h.observe(4.0, 2)  // 2 % 2 == 0 -> shard 0

	// Assert exact per-shard distribution (quiescent: no concurrent writes).
	// shard 0 should have 1 (senderIdx 2), shard 1 should have 3 (-1, 99, 1).
	shardCounts := make([]uint64, h.numShards)
	for i := range h.shards {
		h.shards[i].mu.Lock()
		shardCounts[i] = h.shards[i].count
		h.shards[i].mu.Unlock()
	}
	if shardCounts[0] != 1 {
		t.Fatalf("shard 0 count=%d, want 1 (senderIdx 2 maps here)", shardCounts[0])
	}
	if shardCounts[1] != 3 {
		t.Fatalf("shard 1 count=%d, want 3 (senderIdx -1, 99, 1 map here)", shardCounts[1])
	}

	// Snapshot invariants: aggregate count == sum of shard counts, +Inf == count.
	snap := h.snapshot()
	if snap.Count != 4 {
		t.Fatalf("count=%d, want 4", snap.Count)
	}
	last := len(snap.CumulativeBucketCounts) - 1
	if snap.CumulativeBucketCounts[last] != snap.Count {
		t.Fatalf("+Inf=%d, count=%d, must be equal", snap.CumulativeBucketCounts[last], snap.Count)
	}
}

// TestHistogram_CrossShardAggregation verifies that observations written to
// different shards are correctly summed in the snapshot.
func TestHistogram_CrossShardAggregation(t *testing.T) {
	h := newHistogramState(histogramProfileFastV1, 4)
	// Write 10 observations to each of 4 shards = 40 total
	for shard := 0; shard < 4; shard++ {
		for i := 0; i < 10; i++ {
			h.observe(1.0, shard)
		}
	}
	snap := h.snapshot()
	if snap.Count != 40 {
		t.Fatalf("count=%d, want 40", snap.Count)
	}
	if snap.SumMs != 40.0 {
		t.Fatalf("sum_ms=%v, want 40.0", snap.SumMs)
	}
	// All observations are 1.0ms -> bucket 0 (≤1ms bound)
	if snap.CumulativeBucketCounts[0] != 40 {
		t.Fatalf("cum[0]=%d, want 40", snap.CumulativeBucketCounts[0])
	}
}

// TestHistogram_ConcurrentPerShard verifies N goroutines, each writing
// exclusively to its own shard, produce an exact snapshot (Count, SumMs,
// cumulative buckets) even when snapshot runs concurrently with the writers.
func TestHistogram_ConcurrentPerShard(t *testing.T) {
	const numShards = 8
	const perShard = 1000
	h := newHistogramState(histogramProfileFastV1, numShards)

	var wg sync.WaitGroup
	for shard := 0; shard < numShards; shard++ {
		wg.Add(1)
		go func(s int) {
			defer wg.Done()
			for i := 0; i < perShard; i++ {
				h.observe(1.0, s)
			}
		}(shard)
	}

	// Concurrently take snapshots while writers are active.
	var stop int32
	var snapWg sync.WaitGroup
	snapWg.Add(1)
	go func() {
		defer snapWg.Done()
		for atomic.LoadInt32(&stop) == 0 {
			_ = h.snapshot()
		}
	}()

	wg.Wait()
	atomic.StoreInt32(&stop, 1)
	snapWg.Wait()

	snap := h.snapshot()
	wantCount := uint64(numShards * perShard)
	if snap.Count != wantCount {
		t.Fatalf("count=%d, want %d", snap.Count, wantCount)
	}
	if snap.SumMs != float64(wantCount) {
		t.Fatalf("sum_ms=%v, want %v", snap.SumMs, float64(wantCount))
	}
	// All 1.0ms observations fall in bucket 0 (≤1ms bound).
	if snap.CumulativeBucketCounts[0] != wantCount {
		t.Fatalf("cum[0]=%d, want %d", snap.CumulativeBucketCounts[0], wantCount)
	}
	// +Inf bucket must equal count (shard-internal consistency under lock).
	last := len(snap.CumulativeBucketCounts) - 1
	if snap.CumulativeBucketCounts[last] != snap.Count {
		t.Fatalf("+Inf=%d, count=%d, must be equal", snap.CumulativeBucketCounts[last], snap.Count)
	}
}

// TestHistogram_ConcurrentSameShard verifies multiple goroutines writing to the
// SAME shard (fallback/contention case) still produce an exact count and sum.
func TestHistogram_ConcurrentSameShard(t *testing.T) {
	const numGoroutines = 16
	const perGoroutine = 500
	h := newHistogramState(histogramProfileFastV1, 1) // single shard, all contend

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < perGoroutine; j++ {
				h.observe(2.0, 0) // all write shard 0
			}
		}()
	}
	wg.Wait()

	snap := h.snapshot()
	wantCount := uint64(numGoroutines * perGoroutine)
	if snap.Count != wantCount {
		t.Fatalf("count=%d, want %d", snap.Count, wantCount)
	}
	if snap.SumMs != float64(wantCount)*2.0 {
		t.Fatalf("sum_ms=%v, want %v", snap.SumMs, float64(wantCount)*2.0)
	}
}

// TestHistogram_ModuloMappingLargeSenderIdx verifies that when senderIdx exceeds
// numShards (MaxSenderCount > maxHistogramShards), observations fold into the
// fixed shard pool via modulo and none are lost. Asserts the exact per-shard
// distribution, not just the aggregate count.
func TestHistogram_ModuloMappingLargeSenderIdx(t *testing.T) {
	h := newHistogramState(histogramProfileFastV1, 4)
	// senderIdx 0..7 map to shards 0,1,2,3,0,1,2,3 via modulo.
	for i := 0; i < 8; i++ {
		h.observe(1.0, i)
	}

	// Assert exact per-shard distribution: each shard gets exactly 2.
	for i := range h.shards {
		h.shards[i].mu.Lock()
		got := h.shards[i].count
		h.shards[i].mu.Unlock()
		if got != 2 {
			t.Fatalf("shard %d count=%d, want 2 (senderIdx %d and %d map here)", i, got, i, i+4)
		}
	}

	// Snapshot invariants: aggregate == sum of shards, +Inf == count.
	snap := h.snapshot()
	if snap.Count != 8 {
		t.Fatalf("count=%d, want 8", snap.Count)
	}
	last := len(snap.CumulativeBucketCounts) - 1
	if snap.CumulativeBucketCounts[last] != snap.Count {
		t.Fatalf("+Inf=%d, count=%d, must be equal", snap.CumulativeBucketCounts[last], snap.Count)
	}
}

// TestHistogram_ShardCountBoundary verifies histogramShardCount caps correctly
// for edge cases: non-positive, exactly at cap, above cap, and a value larger
// than 32-bit int (must not overflow on 32-bit targets).
func TestHistogram_ShardCountBoundary(t *testing.T) {
	cases := []struct {
		requested int64
		want      int64
	}{
		{0, 1},
		{-1, 1},
		{-100, 1},
		{1, 1},
		{50, 50},
		{64, 64},                  // exactly at cap
		{65, 64},                  // above cap -> capped
		{1000, 64},                // well above cap -> capped
		{1 << 31, 64},             // 2^31, larger than int32 max -> capped, no overflow
		{1 << 40, 64},             // 2^40, way above cap -> capped, no overflow
		{9223372036854775807, 64}, // int64 max -> capped, no overflow
	}
	for _, c := range cases {
		got := histogramShardCount(c.requested)
		if got != c.want {
			t.Fatalf("histogramShardCount(%d) = %d, want %d", c.requested, got, c.want)
		}
		// Verify the capped value fits in int on this platform (32 or 64 bit).
		_ = int(got)
	}
}

// TestHistogram_NewStateAcceptsInt64 verifies newHistogramState accepts int64
// directly (no pre-cap int conversion) and produces the correct shard count.
func TestHistogram_NewStateAcceptsInt64(t *testing.T) {
	// 0 -> 1 shard
	h := newHistogramState(histogramProfileFastV1, 0)
	if h.numShards != 1 {
		t.Fatalf("numShards=%d, want 1", h.numShards)
	}
	// 64 -> 64 shards (at cap)
	h = newHistogramState(histogramProfileFastV1, 64)
	if h.numShards != 64 {
		t.Fatalf("numShards=%d, want 64", h.numShards)
	}
	// 100 -> 64 shards (capped)
	h = newHistogramState(histogramProfileFastV1, 100)
	if h.numShards != 64 {
		t.Fatalf("numShards=%d, want 64", h.numShards)
	}
	// 2^31 -> 64 shards (capped, safe on 32-bit)
	h = newHistogramState(histogramProfileFastV1, 1<<31)
	if h.numShards != 64 {
		t.Fatalf("numShards=%d, want 64", h.numShards)
	}
}
