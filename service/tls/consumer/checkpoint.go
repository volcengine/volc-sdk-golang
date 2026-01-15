package consumer

import (
	"context"
	"fmt"
	"sync"
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
}

func (c *checkpointManager) run(ctx context.Context, wg *sync.WaitGroup) {
	_ = level.Info(c.logger).Log("msg", "checkpoint manager start")
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

	c.checkpointMap[fmt.Sprintf("%s%s%d", checkpoint.shardInfo.TopicID, Delimiter, checkpoint.shardInfo.ShardID)] = checkpoint
}

func (c *checkpointManager) uploadCheckpoint() {
	checkpointSnapshot := make(map[string]checkpointInfo)
	c.mapLock.Lock()
	for k, checkpoint := range c.checkpointMap {
		checkpointSnapshot[k] = *checkpoint
	}
	c.mapLock.Unlock()

	for k, checkpoint := range checkpointSnapshot {
		if _, err := c.client.ModifyCheckPoint(&tls.ModifyCheckPointRequest{
			ProjectID:         c.conf.ProjectID,
			TopicID:           checkpoint.shardInfo.TopicID,
			ConsumerGroupName: c.conf.ConsumerGroupName,
			ShardID:           checkpoint.shardInfo.ShardID,
			Checkpoint:        checkpoint.checkpoint,
		}); err != nil {
			_ = level.Error(c.logger).Log("msg", "upload checkpoint failed", "topic", checkpoint.shardInfo.TopicID, "shard", checkpoint.shardInfo.ShardID, "checkpoint", checkpoint.checkpoint, "error", err.Error())
			delete(checkpointSnapshot, k)
		}
	}

	c.mapLock.Lock()
	for k, checkpoint := range checkpointSnapshot {
		if checkPointMapK, ok := c.checkpointMap[k]; ok {
			if checkpoint.checkpoint == checkPointMapK.checkpoint {
				delete(c.checkpointMap, k)
			}
		}
	}
	c.mapLock.Unlock()
}

type checkpointInfo struct {
	shardInfo  *tls.ConsumeShard
	checkpoint string
}

func newCheckpointManager(logger log.Logger, conf *Config, client tls.Client, commitCh chan struct{}) *checkpointManager {
	return &checkpointManager{
		logger:        logger,
		conf:          conf,
		client:        client,
		mapLock:       &sync.RWMutex{},
		checkpointMap: make(map[string]*checkpointInfo),
		commitCh:      commitCh,
		stopCh:        make(chan struct{}),
	}
}
