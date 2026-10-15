package consumer

import (
	"errors"
	"sync/atomic"
	"time"
)

var (
	ErrMetricsLevelChangeNotRunning = errors.New("consumer metrics level change requires a running consumer")
	ErrMetricsLevelChangeClosing    = errors.New("consumer is closing or closed")
	ErrInvalidMetricsLevel          = errors.New("invalid consumer metrics level")
)

func validMetricsLevel(level MetricsLevel) bool {
	return level == MetricsLevelOff || level == MetricsLevelCore || level == MetricsLevelDetailed
}

func (c *consumer) SetMetricsLevel(level MetricsLevel) error {
	if c == nil {
		return ErrMetricsLevelChangeNotRunning
	}
	if !validMetricsLevel(level) {
		return ErrInvalidMetricsLevel
	}
	if c.stats == nil {
		return ErrMetricsLevelChangeNotRunning
	}
	c.metricsMu.Lock()
	defer c.metricsMu.Unlock()
	if atomic.LoadInt32(&c.closing) != 0 || atomic.LoadInt32(&c.closed) != 0 {
		return ErrMetricsLevelChangeClosing
	}
	if atomic.LoadInt32(&c.lifecycleStage) != lifecycleRunning {
		return ErrMetricsLevelChangeNotRunning
	}
	old := c.stats.metricsLevelNow()
	if old == level {
		return nil
	}
	if (old == MetricsLevelDetailed) != (level == MetricsLevelDetailed) {
		atomic.AddUint64(&c.stats.detailGeneration, 1)
	}
	if old != MetricsLevelOff && level != MetricsLevelOff {
		if level == MetricsLevelDetailed {
			c.stats.initDetailedHistograms()
		}
		atomic.StoreInt32(&c.stats.runtimeMetricsLevel, int32(level))
		return nil
	}

	if level != MetricsLevelOff {
		if level == MetricsLevelDetailed {
			c.stats.initDetailedHistograms()
		}
		// Serialize publication with shard-set mutation. Once enabled is visible,
		// later mutations update the gauges themselves; the checkpoint HWM uses
		// the same publish-then-observe ordering under its own lock.
		c.stateMu.Lock()
		assigned := len(c.workerMap)
		atomic.AddUint64(&c.stats.collectionGeneration, 1)
		c.stats.resetCollectionGauges()
		atomic.StoreInt64(&c.stats.assignedShards, int64(assigned))
		atomic.StoreInt64(&c.stats.collectionStartedAtUnixMs, unixMillis(time.Now()))
		atomic.StoreInt32(&c.stats.runtimeMetricsLevel, int32(level))
		c.stateMu.Unlock()
		if c.checkpoint != nil {
			c.checkpoint.observeCurrentHighWatermark()
		}
		return nil
	}

	atomic.AddUint64(&c.stats.collectionGeneration, 1)
	atomic.StoreInt64(&c.stats.collectionStartedAtUnixMs, 0)
	atomic.StoreInt32(&c.stats.runtimeMetricsLevel, int32(level))
	c.stats.resetCollectionGauges()
	if c.checkpoint != nil {
		c.checkpoint.resetObservation()
	}
	return nil
}

func (c *consumer) collectionGenerationNow() uint64 {
	if c == nil || c.stats == nil {
		return 0
	}
	return c.stats.collectionGenerationNow()
}

func (c *consumer) collectionStartedAtNow() int64 {
	if c == nil || c.stats == nil {
		return 0
	}
	return c.stats.collectionStartedAtNow()
}
