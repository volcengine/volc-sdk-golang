package producer

import (
	"errors"
	"sync/atomic"
	"time"
)

var (
	ErrMetricsLevelChangeNotRunning = errors.New("producer metrics level change requires a running producer")
	ErrMetricsLevelChangeClosing    = errors.New("producer is closing or closed")
	ErrInvalidMetricsLevel          = errors.New("invalid producer metrics level")
)

func validMetricsLevel(level MetricsLevel) bool {
	return level == MetricsLevelOff || level == MetricsLevelCore || level == MetricsLevelDetailed
}

func (p *producer) metricsLevelNow() MetricsLevel {
	if p == nil {
		return MetricsLevelOff
	}
	return MetricsLevel(atomic.LoadInt32(&p.runtimeMetricsLevel))
}

func (p *producer) metricsEnabled() bool {
	return p != nil && p.stats != nil && p.metricsLevelNow() != MetricsLevelOff
}

// gaugesTracked is false at Off, so collection is genuinely disabled rather
// than merely hidden from snapshots.
func (p *producer) gaugesTracked() bool {
	return p != nil && p.stats != nil && p.metricsEnabled()
}

func (p *producer) beginSendCallObservation() uint32 {
	if !p.metricsEnabled() {
		return 0
	}
	generation := p.collectionGenerationToken()
	p.stats.incSendCallsInFlight()
	return generation
}

func (p *producer) sendCallGenerationEnabled(generation uint32) bool {
	return generation != 0 && p != nil && p.metricsEnabled() &&
		p.collectionGenerationToken() == generation
}

func (p *producer) finishSendCallObservation(generation uint32, failed bool, logs, groups int64) {
	if !p.sendCallGenerationEnabled(generation) {
		return
	}
	p.stats.decSendCallsInFlight()
	if failed {
		p.stats.addSendCallsFailure(1)
	} else {
		p.stats.addSendCallsSuccess(1)
	}
	p.stats.addSendCallPayloadLogs(logs)
	p.stats.addSendCallPayloadGroups(groups)
}

func (p *producer) SetMetricsLevel(level MetricsLevel) error {
	if p == nil {
		return ErrMetricsLevelChangeNotRunning
	}
	if !validMetricsLevel(level) {
		return ErrInvalidMetricsLevel
	}
	p.metricsMu.Lock()
	defer p.metricsMu.Unlock()
	if atomic.LoadInt32(&p.closing) != 0 || atomic.LoadInt32(&p.closed) != 0 {
		return ErrMetricsLevelChangeClosing
	}
	if atomic.LoadInt32(&p.started) == 0 {
		return ErrMetricsLevelChangeNotRunning
	}

	old := p.metricsLevelNow()
	if old == level {
		return nil
	}
	if (old == MetricsLevelDetailed) != (level == MetricsLevelDetailed) {
		atomic.AddUint64(&p.detailGeneration, 1)
	}
	if old != MetricsLevelOff && level != MetricsLevelOff {
		if level == MetricsLevelDetailed && p.stats != nil {
			p.stats.initDetailedHistograms(p.config.MaxSenderCount)
		}
		// Core and Detailed share one uninterrupted cumulative-counter window.
		// Changing only the detail richness must not orphan accepted work or
		// force downstream rate adapters to re-baseline.
		atomic.StoreInt32(&p.runtimeMetricsLevel, int32(level))
		return nil
	}
	if level == MetricsLevelDetailed && p.stats != nil {
		p.stats.initDetailedHistograms(p.config.MaxSenderCount)
	}
	atomic.AddUint64(&p.collectionGeneration, 1)
	if level == MetricsLevelOff {
		if p.dispatcher != nil && p.dispatcher.retryQueue != nil {
			p.dispatcher.retryQueue.setMetricsEnabled(false)
		}
		if p.memoryLimiter != nil {
			p.memoryLimiter.setMetricsEnabled(false)
		}
		atomic.StoreInt64(&p.collectionStartedAtUnixMs, 0)
		atomic.StoreInt32(&p.runtimeMetricsLevel, int32(level))
		p.stats.resetCollectionGauges(false)
		return nil
	}

	p.stats.resetCollectionGauges(atomic.LoadInt32(&p.started) != 0)
	p.stats.observeBatchInflightMemory(p.currentProducerLogGroupSize())
	if p.dispatcher != nil && p.dispatcher.retryQueue != nil {
		p.dispatcher.retryQueue.setMetricsEnabled(true)
	}
	if p.memoryLimiter != nil {
		p.memoryLimiter.setMetricsEnabled(true)
	}
	atomic.StoreInt64(&p.collectionStartedAtUnixMs, unixMillis(time.Now()))
	// Publish enabled only after all optional state exists. GetStats is blocked
	// by metricsMu until the queue observations below complete; hot-path updates
	// that race after this store are included by their own max-CAS operations.
	atomic.StoreInt32(&p.runtimeMetricsLevel, int32(level))
	if p.dispatcher != nil {
		p.stats.observeIngressQueueDepth(int64(len(p.dispatcher.newLogRecvChan)))
		if p.threadPool != nil {
			p.stats.observeSendQueueDepth(int64(len(p.threadPool.taskChan)))
		}
	}
	return nil
}

func (p *producer) collectionGenerationNow() uint64 {
	if p == nil {
		return 0
	}
	return atomic.LoadUint64(&p.collectionGeneration)
}

func (p *producer) collectionGenerationToken() uint32 {
	const maxToken = uint64((1 << 32) - 1)
	value := p.collectionGenerationNow()
	return uint32((value-1)%maxToken + 1)
}

func (p *producer) collectionStartedAtNow() int64 {
	if p == nil {
		return 0
	}
	return atomic.LoadInt64(&p.collectionStartedAtUnixMs)
}

func (p *producer) currentDetailGeneration() uint64 {
	if p == nil || p.metricsLevelNow() != MetricsLevelDetailed {
		return 0
	}
	return atomic.LoadUint64(&p.detailGeneration)
}

func (p *producer) detailGenerationEnabled(generation uint64) bool {
	return generation != 0 && p != nil && p.metricsLevelNow() == MetricsLevelDetailed &&
		atomic.LoadUint64(&p.detailGeneration) == generation
}
