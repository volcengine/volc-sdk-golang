package consumer

import (
	"sync"
	"time"
)

// HistogramSnapshot is a point-in-time, immutable view of a latency histogram.
// Each snapshot carries the fixed bucket profile, explicit upper bounds,
// cumulative bucket counts, count, and sum in milliseconds. UpperBoundsMs and
// CumulativeBucketCounts are independent copies safe to read after GetStats.
//
// Frozen bucket boundaries are part of the latency metric contract:
//
//	fast_v1: [1,2,5,10,20,50,100,200,500,1000,2000,5000,10000] ms (13 bounds, 14 buckets)
//	slow_v1: [10,20,50,100,200,500,1000,2000,5000,10000,20000,50000,100000,200000,300000] ms (15 bounds, 16 buckets)
//
// The profile names and boundaries are part of the fixed metric contract, and
// the bounds are carried with each snapshot for unambiguous interpretation.
type HistogramSnapshot struct {
	BucketProfile          string   `json:"bucket_profile"`
	UpperBoundsMs          []uint64 `json:"upper_bounds_ms"`
	CumulativeBucketCounts []uint64 `json:"cumulative_bucket_counts"`
	Count                  uint64   `json:"count"`
	SumMs                  float64  `json:"sum_ms"`
}

var fastV1BoundsMs = []uint64{
	1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000,
}

var slowV1BoundsMs = []uint64{
	10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 300000,
}

// histogramState holds non-cumulative bucket counts and a sum. All updates
// are serialized by mu. snapshot() produces a cumulative copy.
type histogramState struct {
	profile histogramProfile
	bounds  []uint64
	mu      sync.Mutex
	counts  []uint64 // non-cumulative, len(bounds)+1 (last is +Inf)
	sumMs   float64
	count   uint64
}

func newHistogramState(profile histogramProfile) *histogramState {
	if profile == histogramProfileOff {
		return nil
	}
	var bounds []uint64
	switch profile {
	case histogramProfileFastV1:
		bounds = fastV1BoundsMs
	case histogramProfileSlowV1:
		bounds = slowV1BoundsMs
	default:
		return nil
	}
	return &histogramState{
		profile: profile,
		bounds:  bounds,
		counts:  make([]uint64, len(bounds)+1),
	}
}

func (h *histogramState) observe(ms float64) {
	if h == nil {
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	idx := len(h.bounds) // default to +Inf
	for i, b := range h.bounds {
		if ms <= float64(b) {
			idx = i
			break
		}
	}
	h.counts[idx]++
	h.count++
	h.sumMs += ms
}

func (h *histogramState) observeSince(start time.Time, now func() time.Time) {
	if h == nil {
		return
	}
	h.observe(float64(now().Sub(start).Microseconds()) / 1000.0)
}

func (h *histogramState) snapshot() *HistogramSnapshot {
	if h == nil {
		return nil
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	n := len(h.bounds)
	cum := make([]uint64, n+1)
	var running uint64
	for i := 0; i <= n; i++ {
		running += h.counts[i]
		cum[i] = running
	}
	name := "off"
	switch h.profile {
	case histogramProfileFastV1:
		name = "fast_v1"
	case histogramProfileSlowV1:
		name = "slow_v1"
	}
	return &HistogramSnapshot{
		BucketProfile:          name,
		UpperBoundsMs:          append([]uint64(nil), h.bounds...),
		CumulativeBucketCounts: cum,
		Count:                  h.count,
		SumMs:                  h.sumMs,
	}
}
