package producer

import (
	"container/heap"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type RetryQueue struct {
	batch            []*Batch
	mutex            sync.Mutex
	observedDepth    int64
	observedDepthMax int64
	metricsEnabled   bool
}

func newRetryQueue() *RetryQueue {
	retryQueue := RetryQueue{}
	heap.Init(&retryQueue)

	return &retryQueue
}

// setMetricsEnabled controls whether HWM is tracked.
// Called at Start() and by the runtime level setter.
func (q *RetryQueue) setMetricsEnabled(enabled bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if !enabled {
		q.metricsEnabled = false
		atomic.StoreInt64(&q.observedDepth, 0)
		atomic.StoreInt64(&q.observedDepthMax, 0)
		return
	}
	depth := int64(len(q.batch))
	atomic.StoreInt64(&q.observedDepth, depth)
	atomic.StoreInt64(&q.observedDepthMax, depth)
	q.metricsEnabled = true
}

// snapshot returns pre-aggregated observer facts without taking the business
// heap mutex. Its latency is therefore independent of retry backlog size and
// of an O(N) shutdown drain already holding that mutex.
func (q *RetryQueue) snapshot() (depth int, depthMax int64) {
	depth64 := atomic.LoadInt64(&q.observedDepth)
	depthMax = atomic.LoadInt64(&q.observedDepthMax)
	if depthMax < depth64 {
		depthMax = depth64
	}
	return int(depth64), depthMax
}

func (q *RetryQueue) addToRetryQueue(batch *Batch, logger log.Logger) {
	level.Debug(logger).Log("msg", "Send to retry queue")

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if batch != nil {
		heap.Push(q, batch)
		if q.metricsEnabled {
			depth := int64(len(q.batch))
			atomic.StoreInt64(&q.observedDepth, depth)
			updatePositiveHighWatermark(&q.observedDepthMax, depth)
		}
	}
}

func (q *RetryQueue) getRetryBatch(stopFlag bool) (producerBatchList []*Batch) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if !stopFlag {
		for q.Len() > 0 {
			producerBatch := heap.Pop(q)
			if producerBatch.(*Batch).nextRetryMs < GetTimeMs(time.Now().UnixNano()) {
				producerBatchList = append(producerBatchList, producerBatch.(*Batch))
			} else {
				heap.Push(q, producerBatch.(*Batch))
				break
			}
		}
	} else {
		for q.Len() > 0 {
			producerBatch := heap.Pop(q)
			producerBatchList = append(producerBatchList, producerBatch.(*Batch))
		}
	}
	if q.metricsEnabled {
		atomic.StoreInt64(&q.observedDepth, int64(len(q.batch)))
	}

	return producerBatchList
}

func (q *RetryQueue) Len() int {
	return len(q.batch)
}

func (q *RetryQueue) Less(i, j int) bool {
	return q.batch[i].nextRetryMs < q.batch[j].nextRetryMs
}

func (q *RetryQueue) Swap(i, j int) {
	q.batch[i], q.batch[j] = q.batch[j], q.batch[i]
}

func (q *RetryQueue) Push(x interface{}) {
	item := x.(*Batch)

	q.batch = append(q.batch, item)
}

func (q *RetryQueue) Pop() interface{} {
	old := q.batch
	n := len(old)
	item := old[n-1]
	q.batch = old[0 : n-1]

	return item
}
