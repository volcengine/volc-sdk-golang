package consumer

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/innerlogger"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type consumer struct {
	ctx          context.Context
	cancel       context.CancelFunc
	logger       log.Logger
	checkpointCh chan *checkpointInfo
	conf         *Config
	consumeFunc  func(topicID string, shardID int, l *pb.LogGroupList)
	wg           *sync.WaitGroup

	heartbeat  *heartbeatRunner
	checkpoint *checkpointManager
	workerMap  map[string]*logConsumer

	client tls.Client
}

func newConsumer(ctx context.Context, conf *Config, consumeFunc func(topicID string, shardID int, l *pb.LogGroupList)) (*consumer, error) {
	if err := validateConsumerConfig(conf); err != nil {
		return nil, err
	}

	client := tls.NewClient(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret, "", conf.Region)
	logger := innerlogger.NewLogger(&conf.LoggerConfig)
	checkpointCh := make(chan *checkpointInfo)

	return &consumer{
		ctx:          ctx,
		logger:       logger,
		checkpointCh: checkpointCh,
		conf:         conf,
		consumeFunc:  consumeFunc,
		heartbeat:    newHeartbeatRunner(logger, client, conf),
		checkpoint:   newCheckpointManager(logger, conf, client, checkpointCh),
		workerMap:    make(map[string]*logConsumer),
		client:       client,
		wg:           &sync.WaitGroup{},
	}, nil
}

func (c *consumer) Start() error {
	if c.wg == nil {
		return errors.New("consumer is not ready")
	}

	ctx, cancel := context.WithCancel(c.ctx)
	c.cancel = cancel

	if err := c.init(); err != nil {
		return err
	}

	c.wg.Add(3)
	go c.checkpoint.run(ctx, c.wg)
	go c.heartbeat.run(ctx, c.wg)
	go c.run(ctx, c.wg)

	return nil
}

func (c *consumer) run(ctx context.Context, wg *sync.WaitGroup) {
	level.Info(c.logger).Log("msg", "consumer start")
	defer wg.Done()

	fetchDataTicker := time.NewTicker(time.Duration(c.conf.DataFetchIntervalInMillisecond) * time.Millisecond)
	defer fetchDataTicker.Stop()

	for {
		select {
		case <-fetchDataTicker.C:
			c.handleShards(c.heartbeat.getShards())
		case <-ctx.Done():
			return
		}
	}
}

func (c *consumer) handleShards(shards []*tls.ConsumeShard) {
	newShardsSet := make(map[string]*tls.ConsumeShard)
	for _, shard := range shards {
		newShardsSet[shard.TopicID+strconv.Itoa(shard.ShardID)] = shard
	}

	for shardName := range c.workerMap {
		if _, ok := newShardsSet[shardName]; !ok {
			delete(c.workerMap, shardName)
		}
	}

	for shardName, shardInfo := range newShardsSet {
		_, ok := c.workerMap[shardName]
		if !ok {
			c.workerMap[shardName] = c.newLogConsumer(shardInfo)
		}
	}

	for _, worker := range c.workerMap {
		worker.run()
	}
}

func (c *consumer) newLogConsumer(consumeShard *tls.ConsumeShard) *logConsumer {
	return &logConsumer{
		ctx:              c.ctx,
		client:           c.client,
		logger:           c.logger,
		conf:             c.conf,
		statusLock:       &sync.RWMutex{},
		status:           pending,
		shard:            consumeShard,
		consumeFunc:      c.consumeFunc,
		checkpointCh:     c.checkpointCh,
		nextCheckpoint:   "",
		currLogGroupList: nil,
	}
}

func (c *consumer) init() error {
	_, err := c.client.CreateConsumerGroup(&tls.CreateConsumerGroupRequest{
		ProjectID:         c.conf.ProjectID,
		TopicIDList:       c.conf.TopicIDList,
		ConsumerGroupName: c.conf.ConsumerGroupName,
		HeartbeatTTL:      c.conf.HeartbeatIntervalInSecond * 3,
		OrderedConsume:    c.conf.OrderedConsume,
	})
	if err != nil {
		clientErr := tls.NewClientError(err)
		if clientErr.Code == tls.ErrConsumerGroupAlreadyExists {
			return nil
		}

		return clientErr
	}

	return nil
}

func (c *consumer) Stop() {
	c.workerMap = make(map[string]*logConsumer)

	c.cancel()
	c.wg.Wait()
}
