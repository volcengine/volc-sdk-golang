package consumer

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

type checkpointManager struct {
	logger log.Logger
	conf   *Config
	client tls.Client

	mapLock       *sync.RWMutex
	checkpointMap map[string]*checkpointInfo
	commitCh      <-chan struct{}
	stopCh        chan struct{}

	// Observer-only O(1) queue facts. They are updated only while collection is
	// enabled and initialized from checkpointMap on a runtime Off -> On change.
	pendingCount  int64
	highWatermark int64

	stats *consumerStatsState
}

func (c *checkpointManager) run(ctx context.Context, wg *sync.WaitGroup) {
	_ = level.Info(c.logger).Log("msg", "checkpoint manager start")
	if c.stats != nil {
		c.stats.incCheckpointRunning()
		defer c.stats.decCheckpointRunning()
	}
	defer wg.Done()

	uploadCheckpointTicker := time.NewTicker(time.Duration(c.conf.FlushCheckpointIntervalSecond) * time.Second)
	defer uploadCheckpointTicker.Stop()

	for {
		select {
		case <-c.stopCh:
			c.uploadCheckpoint()
			_ = level.Info(c.logger).Log("msg", "checkpoint manager quit")
			return
		case <-c.commitCh:
			c.uploadCheckpoint()
		case <-uploadCheckpointTicker.C:
			c.uploadCheckpoint()
		}
	}
}

func (c *checkpointManager) addCheckpoint(checkpoint *checkpointInfo) {
	c.mapLock.Lock()
	defer c.mapLock.Unlock()

	key := fmt.Sprintf("%s%s%d", checkpoint.shardInfo.TopicID, Delimiter, checkpoint.shardInfo.ShardID)
	_, existed := c.checkpointMap[key]
	c.checkpointMap[key] = checkpoint
	if !existed && c.stats != nil && c.stats.gaugesTracked() {
		current := atomic.AddInt64(&c.pendingCount, 1)
		c.observeHighWatermark(current)
	}
}

func (c *checkpointManager) uploadCheckpoint() {
	checkpointSnapshot := make(map[string]checkpointInfo)
	c.mapLock.Lock()
	for k, checkpoint := range c.checkpointMap {
		checkpointSnapshot[k] = *checkpoint
	}
	c.mapLock.Unlock()

	// checkpoint_call_duration: time for each ModifyCheckPoint call (incl. retry).
	// time.Now() only when latency is on.
	for k, checkpoint := range checkpointSnapshot {
		var generation uint64
		if c.stats != nil {
			generation = c.stats.beginCheckpointCall()
		}
		var ckptStart time.Time
		var latencyGeneration uint64
		if c.stats != nil {
			latencyGeneration = c.stats.currentDetailedGeneration()
		}
		if latencyGeneration != 0 {
			ckptStart = c.stats.nowFunc()
		}
		if _, err := c.client.ModifyCheckPoint(&tls.ModifyCheckPointRequest{
			ProjectID:         c.conf.ProjectID,
			TopicID:           checkpoint.shardInfo.TopicID,
			ConsumerGroupName: c.conf.ConsumerGroupName,
			ShardID:           checkpoint.shardInfo.ShardID,
			Checkpoint:        checkpoint.checkpoint,
		}); err != nil {
			_ = level.Error(c.logger).Log("msg", "upload checkpoint failed", "topic", checkpoint.shardInfo.TopicID, "shard", checkpoint.shardInfo.ShardID, "checkpoint", checkpoint.checkpoint, "error", err.Error())
			delete(checkpointSnapshot, k)
			if c.stats != nil {
				c.stats.finishCheckpointCall(generation, true)
			}
		} else if c.stats != nil {
			c.stats.finishCheckpointCall(generation, false)
		}
		if c.stats != nil && c.stats.detailedGenerationEnabled(latencyGeneration) {
			c.stats.checkpointCallDuration.observeSince(ckptStart, c.stats.nowFunc)
		}
	}
	c.mapLock.Lock()
	for k, checkpoint := range checkpointSnapshot {
		if checkPointMapK, ok := c.checkpointMap[k]; ok {
			if checkpoint.checkpoint == checkPointMapK.checkpoint {
				delete(c.checkpointMap, k)
				if c.stats != nil && c.stats.gaugesTracked() {
					decrementPositiveInt64(&c.pendingCount, 1)
				}
			}
		}
	}
	c.mapLock.Unlock()
}

// snapshot returns pre-aggregated queue facts without touching checkpointMap,
// so GetStats latency is independent of checkpoint backlog size.
func (c *checkpointManager) snapshot() *queueMeasurements {
	return &queueMeasurements{
		CurrentItems:       uint64(maxInt64(atomic.LoadInt64(&c.pendingCount), 0)),
		HighWatermarkItems: uint64(maxInt64(atomic.LoadInt64(&c.highWatermark), 0)),
		Bounded:            false,
	}
}

func (c *checkpointManager) observeCurrentHighWatermark() {
	if c == nil {
		return
	}
	c.mapLock.Lock()
	defer c.mapLock.Unlock()
	current := int64(len(c.checkpointMap))
	atomic.StoreInt64(&c.pendingCount, current)
	// Collection is already published at this point. An add that raced before
	// this lock may therefore have observed a transient depth greater than the
	// current map length. Preserve that valid peak instead of overwriting it.
	c.observeHighWatermark(current)
}

func (c *checkpointManager) resetObservation() {
	if c == nil {
		return
	}
	c.mapLock.Lock()
	atomic.StoreInt64(&c.pendingCount, 0)
	atomic.StoreInt64(&c.highWatermark, 0)
	c.mapLock.Unlock()
}

func (c *checkpointManager) observeHighWatermark(value int64) {
	for {
		current := atomic.LoadInt64(&c.highWatermark)
		if value <= current || atomic.CompareAndSwapInt64(&c.highWatermark, current, value) {
			return
		}
	}
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type checkpointInfo struct {
	shardInfo  *tls.ConsumeShard
	checkpoint string
}

func newCheckpointManager(logger log.Logger, conf *Config, client tls.Client, commitCh chan struct{}, stats *consumerStatsState) *checkpointManager {
	return &checkpointManager{
		logger:        logger,
		conf:          conf,
		client:        client,
		mapLock:       &sync.RWMutex{},
		checkpointMap: make(map[string]*checkpointInfo),
		commitCh:      commitCh,
		stopCh:        make(chan struct{}),
		stats:         stats,
	}
}
