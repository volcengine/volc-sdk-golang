package producer

import (
	"errors"
	"testing"

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
