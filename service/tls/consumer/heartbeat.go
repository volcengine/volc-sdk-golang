package consumer

import (
	"context"
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

	heartbeatExpiredCh <-chan struct{}
	lock               *sync.RWMutex
	shards             []*tls.ConsumeShard
}

func newHeartbeatRunner(logger log.Logger, client tls.Client, conf *Config, heartbeatExpiredChan chan struct{}) *heartbeatRunner {
	return &heartbeatRunner{
		conf:               conf,
		client:             client,
		logger:             logger,
		lock:               &sync.RWMutex{},
		lastUpdateTime:     time.Now(),
		heartbeatExpiredCh: heartbeatExpiredChan,
	}
}

func (h *heartbeatRunner) run(ctx context.Context, wg *sync.WaitGroup) {
	level.Info(h.logger).Log("msg", "heartbeat start")
	defer wg.Done()

	for i := 0; i < 5; i++ {
		err := h.uploadHeartbeat()
		if err != nil {
			level.Error(h.logger).Log("error", "heartbeat runner upload heartbeat failed, err: "+err.Error())
		}

		if len(h.getShards()) > 0 {
			break
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}

	heartbeatTicker := time.NewTicker(time.Duration(h.conf.HeartbeatIntervalInSecond) * time.Second)
	defer heartbeatTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			level.Warn(h.logger).Log("msg", "heartbeatRunner quit")

			return
		case <-h.heartbeatExpiredCh:
			level.Debug(h.logger).Log("msg", "heartbeat expired, heartbeatRunner sends heartbeat at "+time.Now().String())
			if err := h.uploadHeartbeat(); err != nil {
				level.Error(h.logger).Log("error", "heartbeatRunner failed to upload heartbeat, err: "+err.Error())
			}

			level.Debug(h.logger).Log("msg", "heartbeat runner upload heartbeat done.")
		case sendTime := <-heartbeatTicker.C:
			level.Debug(h.logger).Log("msg", "heartbeatRunner sends heartbeat at "+sendTime.String())
			if err := h.uploadHeartbeat(); err != nil {
				level.Error(h.logger).Log("error", "heartbeatRunner failed to upload heartbeat, err: "+err.Error())
			}

			level.Debug(h.logger).Log("msg", "heartbeat runner upload heartbeat done.")
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

func (h *heartbeatRunner) uploadHeartbeat() error {
	heartbeatResp, err := h.client.ConsumerHeartbeat(&tls.ConsumerHeartbeatRequest{
		ProjectID:         h.conf.ProjectID,
		ConsumerGroupName: h.conf.ConsumerGroupName,
		ConsumerName:      h.conf.ConsumerName,
	})
	if err != nil {
		return err
	}

	h.setShards(heartbeatResp.Shards)

	return nil
}
