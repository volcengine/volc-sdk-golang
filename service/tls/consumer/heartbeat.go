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

type heartbeatRunner struct {
	conf           *Config
	client         tls.Client
	logger         log.Logger
	lastUpdateTime time.Time

	lock   *sync.RWMutex
	shards []*tls.ConsumeShard
}

func newHeartbeatRunner(logger log.Logger, client tls.Client, conf *Config) *heartbeatRunner {
	return &heartbeatRunner{
		conf:           conf,
		client:         client,
		logger:         logger,
		lock:           &sync.RWMutex{},
		lastUpdateTime: time.Now(),
	}
}

func (h *heartbeatRunner) run(ctx context.Context, wg *sync.WaitGroup) {
	level.Info(h.logger).Log("msg", "heartbeat start")
	defer wg.Done()

	newShards, err := h.uploadHeartbeat()
	if err != nil {
		level.Error(h.logger).Log("error", "heartbeat runner upload heartbeat failed, err: "+err.Error())
	} else {
		h.setShards(newShards)
	}

	heartbeatTicker := time.NewTicker(time.Duration(h.conf.HeartbeatIntervalInSecond) * time.Second)
	defer heartbeatTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			level.Warn(h.logger).Log("msg", "heartbeatRunner quit")

			return
		case sendTime := <-heartbeatTicker.C:
			level.Debug(h.logger).Log("msg", "heartbeatRunner sends heartbeat at "+sendTime.String())

			newShards, err := h.uploadHeartbeat()
			if err != nil {
				level.Error(h.logger).Log("error", "heartbeat runner upload heartbeat failed, err: "+err.Error())

				if time.Since(h.lastUpdateTime) > 3*time.Duration(h.conf.HeartbeatIntervalInSecond)*time.Second {
					h.clearShards()
				}

				continue
			}

			h.setShards(newShards)
			level.Debug(h.logger).Log("msg", "heartbeat runner upload heartbeat done. new shards: "+fmt.Sprintf("%+v", newShards))
		}
	}
}

func (h *heartbeatRunner) getShards() []*tls.ConsumeShard {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.shards
}

func (h *heartbeatRunner) setShards(shards []*tls.ConsumeShard) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.shards = shards
}

func (h *heartbeatRunner) clearShards() {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.shards = nil
}

func (h *heartbeatRunner) uploadHeartbeat() ([]*tls.ConsumeShard, error) {
	heartbeatResp, err := h.client.ConsumerHeartbeat(&tls.ConsumerHeartbeatRequest{
		ProjectID:         h.conf.ProjectID,
		ConsumerGroupName: h.conf.ConsumerGroupName,
		ConsumerName:      h.conf.ConsumerName,
	})
	if err != nil {
		return nil, err
	}

	return heartbeatResp.Shards, nil
}
