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

	stats *consumerStatsState
}

func newHeartbeatRunner(logger log.Logger, client tls.Client, conf *Config, heartbeatExpiredChan chan struct{}, stats *consumerStatsState) *heartbeatRunner {
	return &heartbeatRunner{
		conf:               conf,
		client:             client,
		logger:             logger,
		lock:               &sync.RWMutex{},
		lastUpdateTime:     time.Now(),
		heartbeatExpiredCh: heartbeatExpiredChan,
		stats:              stats,
	}
}

func (h *heartbeatRunner) run(ctx context.Context, wg *sync.WaitGroup) {
	level.Info(h.logger).Log("msg", "heartbeat start")
	if h.stats != nil {
		h.stats.incHeartbeatRunning()
		defer h.stats.decHeartbeatRunning()
	}
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// Observe cancellation before each attempt. If Stop() cancelled ctx
		// right after Start(), bail out instead of consuming the full startup
		// retry budget. Boundary: a synchronous network call already in flight
		// inside uploadHeartbeat() cannot be interrupted once entered; this
		// check only prevents *starting* another attempt after cancellation.
		select {
		case <-ctx.Done():
			level.Warn(h.logger).Log("msg", "heartbeatRunner quit during startup")

			return
		default:
		}

		err := h.uploadHeartbeat()
		if err != nil {
			level.Error(h.logger).Log("error", "heartbeat runner upload heartbeat failed, err: "+err.Error())
		}

		if len(h.getShards()) > 0 {
			break
		}

		// Cancellable retry wait: a plain time.Sleep here ignores ctx, so a
		// Stop() right after Start() would block for the full ~2.5s retry
		// budget. Use a timer + ctx select so cancellation returns promptly
		// while keeping the 500ms interval and 5-attempt semantics unchanged.
		retryTimer := time.NewTimer(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			retryTimer.Stop()
			level.Warn(h.logger).Log("msg", "heartbeatRunner quit during startup")

			return
		case <-retryTimer.C:
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
	var generation uint64
	if h.stats != nil {
		generation = h.stats.beginHeartbeatCall()
	}
	heartbeatResp, err := h.client.ConsumerHeartbeat(&tls.ConsumerHeartbeatRequest{
		ProjectID:         h.conf.ProjectID,
		ConsumerGroupName: h.conf.ConsumerGroupName,
		ConsumerName:      h.conf.ConsumerName,
	})
	if err != nil {
		if h.stats != nil {
			h.stats.finishHeartbeatCall(generation, true)
		}
		return err
	}

	h.setShards(heartbeatResp.Shards)
	if h.stats != nil {
		h.stats.finishHeartbeatCall(generation, false)
	}

	return nil
}
