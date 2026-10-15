package producer

import (
	"sync/atomic"
	"time"

	"github.com/pierrec/lz4"
)

const memoryAcquirePollInterval = 100 * time.Millisecond

type payloadMemoryStage int32

const (
	payloadStageIngress payloadMemoryStage = iota
	payloadStageAggregate
	payloadStageSend
	payloadStageRetry
)

type memoryLimiter struct {
	limit             int64
	payloadLimit      int64
	used              int64
	payloadUsed       int64
	temporaryUsed     int64
	ingressUsed       int64
	aggregateUsed     int64
	sendUsed          int64
	retryUsed         int64
	highWatermark     int64
	payloadRejected   int64
	temporaryRejected int64
	metricsOn         int32
	notifier          *broadcastNotifier
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

func (limiter *memoryLimiter) acquirePayloadOutcome(bytes int64, timeout time.Duration) (bool, bool, error) {
	if limiter == nil {
		return false, false, nil
	}
	rejected, err := limiter.acquire(bytes, limiter.payloadLimit, timeout, nil)
	observed := limiter.recordRejected(rejected, &limiter.payloadRejected)
	if err == nil && bytes > 0 && limiter.metricsEnabled() {
		atomic.AddInt64(&limiter.payloadUsed, bytes)
		atomic.AddInt64(&limiter.ingressUsed, bytes)
	}
	return rejected, observed, err
}

// acquirePayloadWithCancelOutcome preserves whether a failure was caused by
// capacity pressure. Cancellation deliberately returns rejected=false so the
// public Producer counters can classify it as a lifecycle rejection.
func (limiter *memoryLimiter) acquirePayloadWithCancelOutcome(bytes int64, timeout time.Duration, done <-chan struct{}) (bool, bool, error) {
	if limiter == nil {
		return false, false, nil
	}
	rejected, err := limiter.acquire(bytes, limiter.payloadLimit, timeout, done)
	observed := limiter.recordRejected(rejected, &limiter.payloadRejected)
	if err == nil && bytes > 0 && limiter.metricsEnabled() {
		atomic.AddInt64(&limiter.payloadUsed, bytes)
		atomic.AddInt64(&limiter.ingressUsed, bytes)
	}
	return rejected, observed, err
}

func (limiter *memoryLimiter) acquireTemporaryWithCancel(bytes int64, timeout time.Duration, done <-chan struct{}) error {
	if limiter == nil {
		return nil
	}
	rejected, err := limiter.acquire(bytes, limiter.limit, timeout, done)
	limiter.recordRejected(rejected, &limiter.temporaryRejected)
	if err == nil && bytes > 0 && limiter.metricsEnabled() {
		atomic.AddInt64(&limiter.temporaryUsed, bytes)
	}
	return err
}

func (limiter *memoryLimiter) acquire(bytes, capBytes int64, timeout time.Duration, done <-chan struct{}) (bool, error) {
	if bytes <= 0 {
		return false, nil
	}
	if capBytes <= 0 || capBytes > limiter.limit {
		capBytes = limiter.limit
	}
	if bytes > capBytes {
		return true, ErrMemoryTimeout
	}
	if timeout == 0 {
		if isMemoryAcquireDone(done) {
			return false, ErrMemoryTimeout
		}
		if limiter.tryAcquire(bytes, capBytes) {
			return false, nil
		}
		return true, ErrMemoryTimeout
	}

	var deadline time.Time
	if timeout > 0 {
		deadline = time.Now().Add(timeout)
	}

	for {
		if isMemoryAcquireDone(done) {
			return false, ErrMemoryTimeout
		}
		if limiter.tryAcquire(bytes, capBytes) {
			return false, nil
		}

		notifyCh := limiter.notifier.beginWait()
		if limiter.tryAcquire(bytes, capBytes) {
			limiter.notifier.endWait()
			return false, nil
		}

		if timeout > 0 {
			now := time.Now()
			if !now.Before(deadline) {
				limiter.notifier.endWait()
				return true, ErrMemoryTimeout
			}
			sleepFor := memoryAcquirePollInterval
			if remaining := deadline.Sub(now); remaining < sleepFor {
				sleepFor = remaining
			}
			if !waitBroadcast(notifyCh, sleepFor, done) {
				limiter.notifier.endWait()
				return false, ErrMemoryTimeout
			}
			limiter.notifier.endWait()
			continue
		}

		if !waitBroadcast(notifyCh, 0, done) {
			limiter.notifier.endWait()
			return false, ErrMemoryTimeout
		}
		limiter.notifier.endWait()
	}
}

func (limiter *memoryLimiter) recordRejected(rejected bool, category *int64) bool {
	if !rejected || limiter == nil || !limiter.metricsEnabled() {
		return false
	}
	atomic.AddInt64(category, 1)
	return true
}

func (limiter *memoryLimiter) tryAcquire(bytes, capBytes int64) bool {
	for {
		used := atomic.LoadInt64(&limiter.used)
		if used > capBytes-bytes {
			return false
		}
		if atomic.CompareAndSwapInt64(&limiter.used, used, used+bytes) {
			if atomic.LoadInt32(&limiter.metricsOn) != 0 {
				updatePositiveHighWatermark(&limiter.highWatermark, used+bytes)
			}
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

func (limiter *memoryLimiter) releasePayloadFrom(stage payloadMemoryStage, bytes int64) {
	if limiter == nil || bytes <= 0 {
		return
	}
	if limiter.metricsEnabled() {
		observed := subtractOwned(limiter.stageCounter(stage), bytes)
		if observed > 0 {
			subtractOwned(&limiter.payloadUsed, observed)
		}
	}
	limiter.release(bytes)
}

func (limiter *memoryLimiter) movePayload(from, to payloadMemoryStage, bytes int64) {
	if limiter == nil || bytes <= 0 || from == to || !limiter.metricsEnabled() {
		return
	}
	moved := subtractOwned(limiter.stageCounter(from), bytes)
	if moved > 0 {
		atomic.AddInt64(limiter.stageCounter(to), moved)
	}
}

func (limiter *memoryLimiter) stageCounter(stage payloadMemoryStage) *int64 {
	switch stage {
	case payloadStageAggregate:
		return &limiter.aggregateUsed
	case payloadStageSend:
		return &limiter.sendUsed
	case payloadStageRetry:
		return &limiter.retryUsed
	default:
		return &limiter.ingressUsed
	}
}

func subtractOwned(owner *int64, bytes int64) int64 {
	for {
		current := atomic.LoadInt64(owner)
		amount := bytes
		if amount > current {
			amount = current
		}
		if atomic.CompareAndSwapInt64(owner, current, current-amount) {
			return amount
		}
	}
}

func (limiter *memoryLimiter) releaseTemporary(bytes int64) {
	if limiter == nil || bytes <= 0 {
		return
	}
	if limiter.metricsEnabled() {
		subtractOwned(&limiter.temporaryUsed, bytes)
	}
	limiter.release(bytes)
}

// release updates only the functional memory budget. Observation ownership is
// maintained separately and is intentionally bypassed while metrics are Off.
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

func (limiter *memoryLimiter) PayloadUsed() int64 {
	if limiter == nil {
		return 0
	}
	return atomic.LoadInt64(&limiter.payloadUsed)
}

func (limiter *memoryLimiter) payloadBreakdown() (ingress, aggregate, send, retry int64, ok bool) {
	if limiter == nil {
		return 0, 0, 0, 0, true
	}
	if !limiter.metricsEnabled() {
		return 0, 0, 0, 0, false
	}
	for attempt := 0; attempt < 8; attempt++ {
		payloadBefore := atomic.LoadInt64(&limiter.payloadUsed)
		ingress = atomic.LoadInt64(&limiter.ingressUsed)
		aggregate = atomic.LoadInt64(&limiter.aggregateUsed)
		send = atomic.LoadInt64(&limiter.sendUsed)
		retry = atomic.LoadInt64(&limiter.retryUsed)
		payloadAfter := atomic.LoadInt64(&limiter.payloadUsed)
		if payloadBefore == payloadAfter && ingress+aggregate+send+retry == payloadAfter {
			return ingress, aggregate, send, retry, true
		}
	}
	return 0, 0, 0, 0, false
}

func (limiter *memoryLimiter) TemporaryUsed() int64 {
	if limiter == nil {
		return 0
	}
	return atomic.LoadInt64(&limiter.temporaryUsed)
}

func (limiter *memoryLimiter) Used() int64 {
	if limiter == nil {
		return 0
	}
	return atomic.LoadInt64(&limiter.used)
}

func (limiter *memoryLimiter) HighWatermark() int64 {
	if limiter == nil {
		return 0
	}
	return atomic.LoadInt64(&limiter.highWatermark)
}

func (limiter *memoryLimiter) PayloadRejected() int64 {
	if limiter == nil {
		return 0
	}
	return atomic.LoadInt64(&limiter.payloadRejected)
}

func (limiter *memoryLimiter) TemporaryRejected() int64 {
	if limiter == nil {
		return 0
	}
	return atomic.LoadInt64(&limiter.temporaryRejected)
}

func (limiter *memoryLimiter) setMetricsEnabled(enabled bool) {
	if limiter == nil {
		return
	}
	if !enabled {
		atomic.StoreInt32(&limiter.metricsOn, 0)
		return
	}
	// A new collection generation never inherits observer-only ownership from
	// the previous window. If functional memory is already live, the exact
	// split remains unavailable until those allocations drain; callers expose
	// that as a partial snapshot instead of fabricating a category.
	atomic.StoreInt64(&limiter.payloadUsed, 0)
	atomic.StoreInt64(&limiter.temporaryUsed, 0)
	atomic.StoreInt64(&limiter.ingressUsed, 0)
	atomic.StoreInt64(&limiter.aggregateUsed, 0)
	atomic.StoreInt64(&limiter.sendUsed, 0)
	atomic.StoreInt64(&limiter.retryUsed, 0)
	atomic.StoreInt64(&limiter.highWatermark, 0)
	atomic.StoreInt32(&limiter.metricsOn, 1)
	updatePositiveHighWatermark(&limiter.highWatermark, atomic.LoadInt64(&limiter.used))
}

func (limiter *memoryLimiter) metricsEnabled() bool {
	return limiter != nil && atomic.LoadInt32(&limiter.metricsOn) != 0
}

func updatePositiveHighWatermark(highWatermark *int64, value int64) {
	for {
		current := atomic.LoadInt64(highWatermark)
		if value <= current || atomic.CompareAndSwapInt64(highWatermark, current, value) {
			return
		}
	}
}

func (limiter *memoryLimiter) reset() {
	if limiter == nil {
		return
	}
	atomic.StoreInt64(&limiter.used, 0)
	atomic.StoreInt64(&limiter.payloadUsed, 0)
	atomic.StoreInt64(&limiter.temporaryUsed, 0)
	atomic.StoreInt64(&limiter.ingressUsed, 0)
	atomic.StoreInt64(&limiter.aggregateUsed, 0)
	atomic.StoreInt64(&limiter.sendUsed, 0)
	atomic.StoreInt64(&limiter.retryUsed, 0)
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
