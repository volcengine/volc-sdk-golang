package producer

import (
	"errors"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	tls "github.com/volcengine/volc-sdk-golang/service/tls"
)

func TestSenderHandleFailure_RetryPolicy(t *testing.T) {
	t.Run("does not retry on non-retryable http code", func(t *testing.T) {
		retryQueue := newRetryQueue()
		p := &producer{}
		s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{}, p)

		b := &Batch{
			totalDataSize:              123,
			maxRetryTimes:              3,
			baseRetryBackoffMs:         1000,
			baseIncreaseRetryBackoffMs: 1000,
			maxRetryIntervalInMs:       10000,
			result:                     newResult(),
			maxReservedAttempts:        5,
		}
		p.producerLogGroupSize = b.totalDataSize

		s.handleFailure(b, &tls.Error{HTTPCode: 401})

		if retryQueue.Len() != 0 {
			t.Fatalf("expected retryQueue empty, got %d", retryQueue.Len())
		}
		if p.producerLogGroupSize != 0 {
			t.Fatalf("expected producerLogGroupSize decremented to 0, got %d", p.producerLogGroupSize)
		}
	})

	t.Run("retries on 429", func(t *testing.T) {
		retryQueue := newRetryQueue()
		p := &producer{}
		s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{}, p)

		b := &Batch{
			totalDataSize:              123,
			maxRetryTimes:              3,
			baseRetryBackoffMs:         1000,
			baseIncreaseRetryBackoffMs: 1000,
			maxRetryIntervalInMs:       10000,
			result:                     newResult(),
			maxReservedAttempts:        5,
		}
		p.producerLogGroupSize = b.totalDataSize

		s.handleFailure(b, &tls.Error{HTTPCode: 429})

		if retryQueue.Len() != 1 {
			t.Fatalf("expected retryQueue size 1, got %d", retryQueue.Len())
		}
		if p.producerLogGroupSize != b.totalDataSize {
			t.Fatalf("expected producerLogGroupSize unchanged, got %d", p.producerLogGroupSize)
		}
	})

	t.Run("does not retry on non tls error", func(t *testing.T) {
		retryQueue := newRetryQueue()
		p := &producer{}
		s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{}, p)

		b := &Batch{
			totalDataSize:              123,
			maxRetryTimes:              3,
			baseRetryBackoffMs:         1000,
			baseIncreaseRetryBackoffMs: 1000,
			maxRetryIntervalInMs:       10000,
			result:                     newResult(),
			maxReservedAttempts:        5,
		}
		p.producerLogGroupSize = b.totalDataSize

		s.handleFailure(b, errors.New("boom"))

		if retryQueue.Len() != 0 {
			t.Fatalf("expected retryQueue empty, got %d", retryQueue.Len())
		}
		if p.producerLogGroupSize != 0 {
			t.Fatalf("expected producerLogGroupSize decremented to 0, got %d", p.producerLogGroupSize)
		}
	})
}

func TestSenderRetryConfigSemantics(t *testing.T) {
	t.Run("no retry status overrides retryable status", func(t *testing.T) {
		retryQueue := newRetryQueue()
		p := &producer{}
		s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{429: {}}, p)

		b := &Batch{
			totalDataSize:              123,
			maxRetryTimes:              3,
			baseRetryBackoffMs:         1000,
			baseIncreaseRetryBackoffMs: 1000,
			maxRetryIntervalInMs:       10000,
			result:                     newResult(),
			maxReservedAttempts:        5,
		}
		p.producerLogGroupSize = b.totalDataSize

		s.handleFailure(b, &tls.Error{HTTPCode: 429})

		if retryQueue.Len() != 0 {
			t.Fatalf("expected retryQueue empty, got %d", retryQueue.Len())
		}
		if b.attemptCount != 1 {
			t.Fatalf("expected one recorded attempt, got %d", b.attemptCount)
		}
		if p.producerLogGroupSize != 0 {
			t.Fatalf("expected producerLogGroupSize decremented to 0, got %d", p.producerLogGroupSize)
		}
	})

	t.Run("retries means retry attempts after the first send", func(t *testing.T) {
		retryQueue := newRetryQueue()
		p := &producer{}
		s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{}, p)

		b := &Batch{
			totalDataSize:              123,
			maxRetryTimes:              2,
			baseRetryBackoffMs:         1,
			baseIncreaseRetryBackoffMs: 1,
			maxRetryIntervalInMs:       1,
			result:                     newResult(),
			maxReservedAttempts:        5,
		}
		p.producerLogGroupSize = b.totalDataSize

		s.handleFailure(b, &tls.Error{HTTPCode: 429})
		if retryQueue.Len() != 1 {
			t.Fatalf("expected first failure to schedule retry, got queue len %d", retryQueue.Len())
		}
		_ = retryQueue.getRetryBatch(true)

		s.handleFailure(b, &tls.Error{HTTPCode: 429})
		if retryQueue.Len() != 1 {
			t.Fatalf("expected second failure to schedule retry, got queue len %d", retryQueue.Len())
		}
		_ = retryQueue.getRetryBatch(true)

		s.handleFailure(b, &tls.Error{HTTPCode: 429})
		if retryQueue.Len() != 0 {
			t.Fatalf("expected third failure to stop retrying, got queue len %d", retryQueue.Len())
		}
		if b.attemptCount != 3 {
			t.Fatalf("expected total send attempts to be retries+1, got %d", b.attemptCount)
		}
		if p.producerLogGroupSize != 0 {
			t.Fatalf("expected producerLogGroupSize decremented to 0, got %d", p.producerLogGroupSize)
		}
	})

	t.Run("retry backoff is capped by config", func(t *testing.T) {
		retryQueue := newRetryQueue()
		p := &producer{}
		s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{}, p)

		b := &Batch{
			totalDataSize:              123,
			maxRetryTimes:              3,
			baseRetryBackoffMs:         1000,
			baseIncreaseRetryBackoffMs: 1000,
			maxRetryIntervalInMs:       10,
			result:                     newResult(),
			maxReservedAttempts:        5,
		}
		p.producerLogGroupSize = b.totalDataSize

		s.handleFailure(b, &tls.Error{HTTPCode: 429})
		if b.retryBackoffMs != 10 {
			t.Fatalf("expected retryBackoffMs capped to 10, got %d", b.retryBackoffMs)
		}
		if b.nextRetryMs < GetTimeMs(time.Now().UnixNano()) {
			t.Fatalf("expected nextRetryMs in the future, got %d", b.nextRetryMs)
		}
	})
}
