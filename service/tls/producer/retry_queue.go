package producer

import (
	"container/heap"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type RetryQueue struct {
	batch []*Batch
	mutex sync.Mutex
}

func newRetryQueue() *RetryQueue {
	retryQueue := RetryQueue{}
	heap.Init(&retryQueue)

	return &retryQueue
}

func (q *RetryQueue) addToRetryQueue(batch *Batch, logger log.Logger) {
	level.Debug(logger).Log("msg", "Send to retry queue")

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if batch != nil {
		heap.Push(q, batch)
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
