package producer

import (
	"sync"
	"time"
)

// HistogramSnapshot is a point-in-time, immutable view of a latency histogram.
// The latency contract uses fixed buckets, cumulative counts, a +Inf bucket,
// count == sum(non_cumulative_buckets), monotonic-clock milliseconds.
// The snapshot holds a copy of the bucket counts; it shares no mutable state
// with the live histogram and is safe to read after GetStats() returns.
type HistogramSnapshot struct {
	BucketProfile          string   `json:"bucket_profile"`
	UpperBoundsMs          []uint64 `json:"upper_bounds_ms"`
	CumulativeBucketCounts []uint64 `json:"cumulative_bucket_counts"`
	Count                  uint64   `json:"count"`
	SumMs                  float64  `json:"sum_ms"`
}

// fastV1BoundsMs contains the frozen short-operation bucket boundaries.
// 13 finite bounds -> 14 buckets (incl. +Inf). Standard 1-2-5 gradation, covering
// 1ms-10s for local queue / single request. Each bound is the inclusive upper
// edge of a bucket; the final +Inf bucket captures everything above the last
// bound. The frozen finite bounds are copied into UpperBoundsMs in each
// immutable snapshot.
var fastV1BoundsMs = []uint64{
	1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000,
}

// slowV1BoundsMs contains the frozen end-to-end bucket boundaries.
// 15 finite bounds -> 16 buckets (incl. +Inf). Standard 1-2-5 gradation, covering
// 10ms-5min for delivery/process/checkpoint including retry.
var slowV1BoundsMs = []uint64{
	10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 300000,
}

// maxHistogramShards caps the number of per-sender shards so histogram memory
// and GetStats lock count stay bounded regardless of the user's MaxSenderCount
// config. Default 50 senders map 1:1 to 50 shards; configs > 64 map senderIdx
// modulo 64 (controlled contention, no lost observations). 4 histograms * 64
// shards * ~168B ≈ 42 KiB upper bound; GetStats locks at most 4*64=256 times.
const maxHistogramShards = 64

// histogramShard is a per-sender slice of a histogram. Each sender goroutine
// writes exclusively to its own shard (selected by senderIdx), so there is no
// lock contention between senders. GetStats() aggregates across all shards.
type histogramShard struct {
	mu     sync.Mutex
	counts []uint64 // non-cumulative, len(bounds)+1 (last is +Inf)
	sumMs  float64
	count  uint64
}

// histogramState holds non-cumulative bucket counts and a sum, sharded per
// sender to eliminate cross-sender lock contention. All writes are serialized
// per-shard by shard.mu; snapshot() aggregates across shards under per-shard
// locks. The bounds slice is shared (read-only after init) across all shards.
type histogramState struct {
	profile   histogramProfile
	bounds    []uint64
	shards    []histogramShard
	numShards int
}

// histogramShardCount converts a requested sender count into the actual shard
// count (as int64): non-positive -> 1, positive -> min(requested, maxHistogramShards).
// Returns int64 so callers defer the int conversion until after capping, keeping
// the path 32-bit safe. Keeps memory and GetStats lock count bounded without
// changing the public MaxSenderCount config semantics.
func histogramShardCount(requested int64) int64 {
	if requested < 1 {
		return 1
	}
	if requested > maxHistogramShards {
		return maxHistogramShards
	}
	return requested
}

// newHistogramState creates a histogram for the given profile with numShards
// per-sender shards (capped at maxHistogramShards). Returns nil if profile is
// Off creates no histogram. numShards is
// int64 (matches the MaxSenderCount config field); the cap to maxHistogramShards
// (64) is applied before the int conversion, so this is safe on 32-bit targets.
func newHistogramState(profile histogramProfile, numShards int64) *histogramState {
	var bounds []uint64
	switch profile {
	case histogramProfileFastV1:
		bounds = fastV1BoundsMs
	case histogramProfileSlowV1:
		bounds = slowV1BoundsMs
	default:
		return nil
	}
	// Cap before int conversion: n is guaranteed <= maxHistogramShards (64),
	// which fits in int on both 32-bit and 64-bit targets.
	n := int(histogramShardCount(numShards))
	shards := make([]histogramShard, n)
	for i := range shards {
		shards[i].counts = make([]uint64, len(bounds)+1)
	}
	return &histogramState{
		profile:   profile,
		bounds:    bounds,
		shards:    shards,
		numShards: n,
	}
}

// observe records a duration in milliseconds into the correct bucket on the
// shard for senderIdx. No-op if h is nil. senderIdx maps to a shard via modulo
// numShards so no observation is lost even when senderIdx >= numShards (e.g.
// MaxSenderCount > maxHistogramShards). Uses the bucket's inclusive upper edge;
// durations above the last bound fall into the +Inf bucket.
func (h *histogramState) observe(ms float64, senderIdx int) {
	if h == nil || h.numShards == 0 {
		return
	}
	// modulo mapping: negative idx wraps to a valid shard, large idx folds into
	// the fixed shard pool. No observation is dropped.
	shard := &h.shards[((senderIdx%h.numShards)+h.numShards)%h.numShards]
	shard.mu.Lock()
	defer shard.mu.Unlock()
	idx := len(h.bounds) // default to +Inf
	for i, b := range h.bounds {
		if ms <= float64(b) {
			idx = i
			break
		}
	}
	shard.counts[idx]++
	shard.count++
	shard.sumMs += ms
}

// observeSince records the duration elapsed since start (monotonic clock) on
// the shard for senderIdx.
func (h *histogramState) observeSince(start time.Time, senderIdx int) {
	if h == nil {
		return
	}
	h.observe(float64(time.Since(start).Microseconds())/1000.0, senderIdx)
}

// snapshot returns a cumulative, immutable copy aggregated across all shards.
// Returns nil if h is nil. Each shard is locked briefly to copy its counts,
// sum, and count under the SAME lock, so a single returned snapshot is internally
// consistent (count == +Inf bucket, buckets/count/sum from the same point in
// time for that shard). Cross-shard consistency is weak (a write may land in
// shard 0 before snapshot locks it and in shard 1 after), which is acceptable
// per the weak-consistency snapshot contract. count == sum(non_cumulative_counts);
// the +Inf bucket (last) == count.
func (h *histogramState) snapshot() *HistogramSnapshot {
	if h == nil {
		return nil
	}
	n := len(h.bounds)
	nc := make([]uint64, n+1) // non-cumulative totals across shards
	var totalCount uint64
	var totalSum float64
	for i := range h.shards {
		shard := &h.shards[i]
		shard.mu.Lock()
		// copy counts, sum, and count under the same lock for shard-internal
		// consistency: the returned snapshot's Count == +Inf bucket for this
		// shard's contribution, and SumMs matches Count's observations.
		for j := 0; j <= n; j++ {
			nc[j] += shard.counts[j]
		}
		totalCount += shard.count
		totalSum += shard.sumMs
		shard.mu.Unlock()
	}
	// convert non-cumulative totals to cumulative
	cum := make([]uint64, n+1)
	var running uint64
	for i := 0; i < n; i++ {
		running += nc[i]
		cum[i] = running
	}
	// The +Inf bucket equals the total count.
	cum[n] = totalCount

	profileName := "off"
	switch h.profile {
	case histogramProfileFastV1:
		profileName = "fast_v1"
	case histogramProfileSlowV1:
		profileName = "slow_v1"
	}

	return &HistogramSnapshot{
		BucketProfile:          profileName,
		UpperBoundsMs:          append([]uint64(nil), h.bounds...),
		CumulativeBucketCounts: cum,
		Count:                  totalCount,
		SumMs:                  totalSum,
	}
}
