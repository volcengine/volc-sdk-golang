package producer

import (
	"sync/atomic"
	"time"

	"github.com/pierrec/lz4"
)

const memoryAcquirePollInterval = 100 * time.Millisecond

type memoryLimiter struct {
	limit        int64
	payloadLimit int64
	used         int64
	notifier     *broadcastNotifier
}

func newMemoryLimiter(limit, payloadLimit int64) *memoryLimiter {
	if payloadLimit <= 0 || payloadLimit > limit {
		payloadLimit = limit
	}
	return &memoryLimiter{
		limit:        limit,
		payloadLimit: payloadLimit,
		notifier:     newBroadcastNotifier(),
	}
}

func (limiter *memoryLimiter) acquirePayload(bytes int64, timeout time.Duration) error {
	if limiter == nil {
		return nil
	}
	return limiter.acquire(bytes, limiter.payloadLimit, timeout, nil)
}

func (limiter *memoryLimiter) acquirePayloadWithCancel(bytes int64, timeout time.Duration, done <-chan struct{}) error {
	if limiter == nil {
		return nil
	}
	return limiter.acquire(bytes, limiter.payloadLimit, timeout, done)
}

func (limiter *memoryLimiter) acquireTemporary(bytes int64, timeout time.Duration) error {
	if limiter == nil {
		return nil
	}
	return limiter.acquire(bytes, limiter.limit, timeout, nil)
}

func (limiter *memoryLimiter) acquireTemporaryWithCancel(bytes int64, timeout time.Duration, done <-chan struct{}) error {
	if limiter == nil {
		return nil
	}
	return limiter.acquire(bytes, limiter.limit, timeout, done)
}

func (limiter *memoryLimiter) acquire(bytes, capBytes int64, timeout time.Duration, done <-chan struct{}) error {
	if bytes <= 0 {
		return nil
	}
	if capBytes <= 0 || capBytes > limiter.limit {
		capBytes = limiter.limit
	}
	if bytes > capBytes {
		return ErrMemoryTimeout
	}
	if timeout == 0 {
		if isMemoryAcquireDone(done) {
			return ErrMemoryTimeout
		}
		if limiter.tryAcquire(bytes, capBytes) {
			return nil
		}
		return ErrMemoryTimeout
	}

	var deadline time.Time
	if timeout > 0 {
		deadline = time.Now().Add(timeout)
	}

	for {
		if isMemoryAcquireDone(done) {
			return ErrMemoryTimeout
		}
		if limiter.tryAcquire(bytes, capBytes) {
			return nil
		}

		notifyCh := limiter.notifier.beginWait()
		if limiter.tryAcquire(bytes, capBytes) {
			limiter.notifier.endWait()
			return nil
		}

		if timeout > 0 {
			now := time.Now()
			if !now.Before(deadline) {
				limiter.notifier.endWait()
				return ErrMemoryTimeout
			}
			sleepFor := memoryAcquirePollInterval
			if remaining := deadline.Sub(now); remaining < sleepFor {
				sleepFor = remaining
			}
			if !waitBroadcast(notifyCh, sleepFor, done) {
				limiter.notifier.endWait()
				return ErrMemoryTimeout
			}
			limiter.notifier.endWait()
			continue
		}

		if !waitBroadcast(notifyCh, 0, done) {
			limiter.notifier.endWait()
			return ErrMemoryTimeout
		}
		limiter.notifier.endWait()
	}
}

func (limiter *memoryLimiter) tryAcquire(bytes, capBytes int64) bool {
	for {
		used := atomic.LoadInt64(&limiter.used)
		if used > capBytes-bytes {
			return false
		}
		if atomic.CompareAndSwapInt64(&limiter.used, used, used+bytes) {
			return true
		}
	}
}

func isMemoryAcquireDone(done <-chan struct{}) bool {
	if done == nil {
		return false
	}
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func (limiter *memoryLimiter) release(bytes int64) {
	if limiter == nil || bytes <= 0 {
		return
	}
	for {
		used := atomic.LoadInt64(&limiter.used)
		next := used - bytes
		if next < 0 {
			next = 0
		}
		if atomic.CompareAndSwapInt64(&limiter.used, used, next) {
			if next < used {
				limiter.notifier.notify()
			}
			return
		}
	}
}

func (limiter *memoryLimiter) Used() int64 {
	if limiter == nil {
		return 0
	}
	return atomic.LoadInt64(&limiter.used)
}

func (limiter *memoryLimiter) reset() {
	if limiter == nil {
		return
	}
	atomic.StoreInt64(&limiter.used, 0)
	limiter.notifier.notify()
}

func estimateBatchLogReservationBytes(batchLog *BatchLog, logSize int64) int64 {
	if batchLog == nil || batchLog.Log == nil {
		return 0
	}
	if logSize <= 0 || logSize > int64(intMax) {
		logSize = int64(batchLog.size())
	}
	groupSize := sizeOfBytesField(int(logSize))
	groupSize += sizeOfStringField(batchLog.Key.Source)
	groupSize += sizeOfStringField(batchLog.Key.FileName)
	groupSize += sizeOfStringField(batchLog.Key.ContextFlow)
	return int64(sizeOfBytesField(groupSize))
}

func sizeOfStringField(value string) int {
	if value == "" {
		return 0
	}
	return sizeOfBytesField(len(value))
}

func estimatePutLogsBufferBytes(rawSize int64) int64 {
	if rawSize <= 0 {
		return 0
	}
	if rawSize > int64(intMax) {
		if rawSize > int64Max/2 {
			return int64Max
		}
		return rawSize * 2
	}
	return rawSize + int64(lz4.CompressBlockBound(int(rawSize)))
}

func defaultProducerMemoryBytes(totalSize int64) int64 {
	if totalSize <= 0 {
		return 0
	}
	if totalSize > int64Max/2 {
		return int64Max
	}
	return totalSize * 2
}

func derivePayloadMemoryLimit(config *Config) int64 {
	total := config.MaxProducerMemoryBytes
	if total <= 0 {
		return 0
	}
	perSenderScratch := estimatePutLogsBufferBytes(config.MaxBatchSize)
	var sendScratch int64
	if perSenderScratch > 0 && config.MaxSenderCount > total/perSenderScratch {
		sendScratch = total
	} else {
		sendScratch = perSenderScratch * config.MaxSenderCount
	}
	maxScratch := total / 2
	if sendScratch > maxScratch {
		sendScratch = maxScratch
	}
	payloadLimit := total - sendScratch
	if payloadLimit <= 0 {
		return total
	}
	return payloadLimit
}
