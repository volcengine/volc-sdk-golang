package consumer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// lifecycleStage values for the linearized Start/Stop state machine.
const (
	lifecycleNew = iota
	lifecycleStarting
	lifecycleRunning
	lifecycleClosing
	lifecycleClosed
)

type consumer struct {
	ctx                context.Context
	cancel             context.CancelFunc
	logger             log.Logger
	conf               *Config
	consumeFunc        func(topicID string, shardID int, l *pb.LogGroupList)
	heartbeatExpiredCh chan struct{}
	commitCh           chan struct{}
	wg                 *sync.WaitGroup

	heartbeat  *heartbeatRunner
	checkpoint *checkpointManager
	workerMap  map[string]*logConsumer

	// workerWG tracks every in-flight worker task (init/fetch/consume/backoff)
	// across ALL workers, including those removed from workerMap by rebalance.
	// It is consumer-level and constant-space: a single WaitGroup, never a slice
	// that grows with rebalance count. Stop() waits on it only AFTER the main
	// loop has exited (so no new Add() calls can race the Wait), then finalizes.
	workerWG *sync.WaitGroup

	// mainLoopDone is closed by run() when it exits. Stop() waits on it before
	// waiting on workerWG: once the main loop has exited, handleShards can no
	// longer be called, so no worker task can call workerWG.Add(1). This
	// prevents the classic Add/Wait race and guarantees that after mainLoopDone
	// closes, workerWG only ever decreases.
	mainLoopDone chan struct{}

	// stopTimeout bounds how long Stop() blocks waiting for workers to exit.
	// Defaults to 30s (the historical consumerStop bound). Overridable in tests
	// so the bounded-stop semantics can be verified without waiting 30s. On
	// timeout Stop returns with the consumer still in stateClosing (no fake
	// final); an async goroutine finalizes once workers actually stop.
	stopTimeout time.Duration

	client tls.Client

	stats *consumerStatsState

	// metricsMu serializes runtime metrics-level changes. It is independent of
	// lifecycleMu and never held by the fetch/process hot path.
	metricsMu sync.RWMutex

	// stateMu protects lifecycleState and workerMap snapshot consistency.
	stateMu sync.RWMutex
	started int32
	closing int32
	closed  int32 // set once after Stop() completes all background work; freezes snapshots

	// lifecycleMu serializes Start()/Stop() transitions so they are linearized.
	// It is NEVER held across the blocking network init() call: Start() releases
	// it before init() and re-acquires after. The hot fetch/process path never
	// touches it.
	lifecycleMu sync.Mutex

	// lifecycleStage is the linearized lifecycle state, read atomically by
	// GetStats (no lock) and written under lifecycleMu by Start/Stop.
	lifecycleStage int32

	// runtimeCtx is the context derived from c.ctx by the successful Start()
	// that launched the goroutines. Workers (logConsumer) use it instead of
	// c.ctx so they sense the Start-derived cancel on Stop(). Nil until the
	// first Start() transitions to starting.
	runtimeCtx context.Context

	// runLaunched is set true exactly once, under lifecycleMu, when Start()
	// launches the three goroutines (checkpoint/heartbeat/run). Stop() uses it
	// to decide whether mainLoopDone will ever close: if false, goroutines were
	// never started (init failed/was stopped) and Stop must NOT wait on
	// mainLoopDone or the workerWG/checkpoint finalizer.
	runLaunched bool

	// stopOnce is immutable for one lifecycle generation. Stop() captures the
	// current pointer under lifecycleMu before calling Do(), while a successful
	// restart publishes a fresh pointer. This preserves the concurrent Stop()
	// contract without attempting to reset sync.Once in place.
	stopOnce *sync.Once
}

func newConsumer(ctx context.Context, conf *Config, consumeFunc func(topicID string, shardID int, l *pb.LogGroupList)) (*consumer, error) {
	client := tls.NewClient(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret, conf.SecurityToken, conf.Region)
	if err := validateConsumerConfig(conf); err != nil {
		return nil, err
	}

	var logger log.Logger
	if conf.Logger != nil {
		logger = *conf.Logger
	} else {
		logger = common.LogConfig(conf.LoggerConfig)
	}
	heartbeatExpiredCh := make(chan struct{})
	commitCh := make(chan struct{})
	stats := newConsumerStatsState(conf.MetricsLevel)

	return &consumer{
		ctx:                ctx,
		logger:             logger,
		conf:               conf,
		consumeFunc:        consumeFunc,
		heartbeatExpiredCh: heartbeatExpiredCh,
		commitCh:           commitCh,
		heartbeat:          newHeartbeatRunner(logger, client, conf, heartbeatExpiredCh, stats),
		checkpoint:         newCheckpointManager(logger, conf, client, commitCh, stats),
		workerMap:          make(map[string]*logConsumer),
		workerWG:           &sync.WaitGroup{},
		mainLoopDone:       make(chan struct{}),
		stopTimeout:        30 * time.Second,
		client:             client,
		wg:                 &sync.WaitGroup{},
		stats:              stats,
		stopOnce:           &sync.Once{},
	}, nil
}

func (c *consumer) Start() error {
	if c == nil {
		return errors.New("consumer is not ready")
	}

	// Create the derived runtime context before taking the lock so the cancel
	// is published atomically with the lifecycleStarting transition: Stop()
	// that observes lifecycleStarting/lifecycleRunning is guaranteed a non-nil
	// cancel.
	ctx, cancel := context.WithCancel(c.ctx)

	c.lifecycleMu.Lock()
	if c.wg == nil {
		c.lifecycleMu.Unlock()
		cancel()
		return errors.New("consumer is not ready")
	}
	stage := atomic.LoadInt32(&c.lifecycleStage)
	switch stage {
	case lifecycleStarting, lifecycleRunning:
		c.lifecycleMu.Unlock()
		cancel()
		return errors.New("consumer already started")
	case lifecycleClosing:
		c.lifecycleMu.Unlock()
		cancel()
		return errors.New("consumer is stopping or already stopped")
	case lifecycleNew, lifecycleClosed:
		// Both a first start and an explicit restart publish a new lifecycle
		// generation. A closed generation is safe to recycle only after its
		// finalizer has published lifecycleClosed, so the old WaitGroups and
		// goroutines cannot still be active here.
		c.metricsMu.Lock()
		if stage == lifecycleClosed {
			// Keep the same stats object so InstanceID and cumulative counters
			// survive restarts, but thaw only under the metrics write lock. Public
			// GetStats holds the matching read lock, so it cannot observe a
			// closed/frozen or running/frozen mixed snapshot during this change.
			c.stats.resumeForStart()
			// A failed final checkpoint flush may intentionally leave entries in
			// checkpointMap. Rebuild the observer-only pending/HWM facts from that
			// real map for the new collection generation instead of dropping them.
			if c.checkpoint != nil {
				c.checkpoint.resetObservation()
				if c.stats.isEnabled() {
					c.checkpoint.observeCurrentHighWatermark()
				}
			}
		}
		c.wg = &sync.WaitGroup{}
		c.workerWG = &sync.WaitGroup{}
		c.mainLoopDone = make(chan struct{})
		c.checkpoint.stopCh = make(chan struct{})
		c.stopOnce = &sync.Once{}
		c.runLaunched = false
		c.cancel = cancel
		c.runtimeCtx = ctx
		atomic.StoreInt32(&c.started, 0)
		atomic.StoreInt32(&c.closing, 0)
		atomic.StoreInt32(&c.closed, 0)
		atomic.StoreInt32(&c.lifecycleStage, lifecycleStarting)
		c.metricsMu.Unlock()
	default:
		c.lifecycleMu.Unlock()
		cancel()
		return errors.New("consumer has invalid lifecycle state")
	}
	// c.lifecycleMu remains held until all per-run fields and the lifecycle
	// stage are published. Stop() therefore cannot capture the new generation
	// half-way through its setup.
	c.lifecycleMu.Unlock()

	// init() makes blocking network calls; the lifecycle lock is NOT held here
	// so Stop() can linearize and close even while init is in flight (contract F).
	if err := c.init(); err != nil {
		c.lifecycleMu.Lock()
		if atomic.LoadInt32(&c.lifecycleStage) == lifecycleStarting {
			// Transient init failure with no concurrent Stop: reset to New so a
			// subsequent Start() can retry (contract C). Clear the cancel we
			// published so a later Stop() does not call a stale cancel.
			c.metricsMu.Lock()
			atomic.StoreInt32(&c.lifecycleStage, lifecycleNew)
			atomic.StoreInt32(&c.started, 0)
			atomic.StoreInt32(&c.closing, 0)
			atomic.StoreInt32(&c.closed, 0)
			c.cancel = nil
			c.runtimeCtx = nil
			c.metricsMu.Unlock()
		}
		// If stage is lifecycleClosing, Stop() already took over; leave it.
		c.lifecycleMu.Unlock()
		cancel()
		return err
	}

	c.lifecycleMu.Lock()
	if atomic.LoadInt32(&c.lifecycleStage) != lifecycleStarting {
		// Stop() ran during init() and transitioned to lifecycleClosing. We
		// must NOT launch goroutines (contract F). Clean up and return.
		c.lifecycleMu.Unlock()
		cancel()
		return errors.New("consumer stopped during init")
	}
	// Transition to Running and launch goroutines atomically under the lock so
	// Stop() that observes lifecycleRunning is guaranteed the goroutines (and
	// thus mainLoopDone) exist. No Add/Wait race: Stop only waits on c.wg when
	// runLaunched is true, which is set here before the lock is released.
	c.metricsMu.Lock()
	atomic.StoreInt32(&c.lifecycleStage, lifecycleRunning)
	c.runLaunched = true
	atomic.StoreInt32(&c.started, 1)
	c.wg.Add(3)
	go c.checkpoint.run(ctx, c.wg)
	go c.heartbeat.run(ctx, c.wg)
	go c.run(ctx, c.wg)
	c.metricsMu.Unlock()
	c.lifecycleMu.Unlock()

	return nil
}

func (c *consumer) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
	c.client.ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken)
}

func (c *consumer) run(ctx context.Context, wg *sync.WaitGroup) {
	_ = level.Info(c.logger).Log("msg", "consumer start")
	// mainLoopDone signals Stop() that the scheduler has exited and no new
	// worker tasks will be scheduled (no new workerWG.Add calls). This close
	// runs AFTER wg.Done() (defers are LIFO: wg.Done is registered later, so it
	// executes first). By the time mainLoopDone closes, run has already called
	// wg.Done() and performs no further scheduling/Add, so Stop() can still rely
	// on it as the "no more Adds" barrier.
	defer close(c.mainLoopDone)
	if c.stats != nil {
		c.stats.incMainLoopRunning()
		defer func() {
			c.stats.decMainLoopRunning()
			wg.Done()
		}()
	} else {
		defer wg.Done()
	}

	fetchDataTicker := time.NewTicker(time.Duration(c.conf.DataFetchIntervalInMillisecond) * time.Millisecond)
	defer fetchDataTicker.Stop()

	for {
		select {
		case <-fetchDataTicker.C:
			c.handleShards(c.heartbeat.getShards())
		case <-ctx.Done():
			c.consumerStop()
			return
		}
	}
}

func (c *consumer) consumerStop() {
	_ = level.Info(c.logger).Log("msg", "consumer stopping")
	// The main loop simply exits on cancel. It does NOT signal the checkpoint
	// manager to stop: that must happen only AFTER every in-flight worker has
	// exited (and called addCheckpoint with its final cursor), otherwise the
	// final flush runs against an empty checkpoint map and the last cursor is
	// lost (restart duplicates). Stop()'s finalizer is responsible for waiting
	// on workerWG, then sending checkpoint.stopCh, then waiting on c.wg.
	_ = level.Info(c.logger).Log("msg", "consumer stopped")
}

func (c *consumer) handleShards(shards []*tls.ConsumeShard) {
	newShardsSet := make(map[string]*tls.ConsumeShard)
	for _, shard := range shards {
		newShardsSet[fmt.Sprintf("%s%s%d", shard.TopicID, Delimiter, shard.ShardID)] = shard
	}

	c.stateMu.Lock()
	oldCount := len(c.workerMap)
	for shardName, lc := range c.workerMap {
		if _, ok := newShardsSet[shardName]; !ok {
			// Worker is being removed (rebalance). Mark it retired so its
			// in-flight task knows not to leave a held/ready batch, and
			// idempotently release any batch it currently holds. The worker
			// pointer is dropped from workerMap here; it is NOT retained in
			// any registry (no unbounded state). Its in-flight goroutine is
			// still tracked by the shared workerWG, so Stop() will wait for it.
			lc.markRetired()
			lc.releaseHeldBatchIfAny()
			delete(c.workerMap, shardName)
		}
	}

	added := 0
	for shardName, shardInfo := range newShardsSet {
		lc, ok := c.workerMap[shardName]
		if !ok || lc.loadStatus() == waitForRestart {
			c.workerMap[shardName] = c.newLogConsumer(shardInfo)
			if !ok {
				added++
			}
		}
	}
	newCount := len(c.workerMap)
	c.stateMu.Unlock()

	// Update pre-aggregated counters (no map traversal in GetStats).
	if c.stats != nil {
		c.stats.setAssignedShards(newCount)
		// A rebalance event is any change in assigned shard set.
		if added > 0 || newCount != oldCount {
			c.stats.incRebalanceEvents()
		}
	}

	c.stateMu.RLock()
	workers := make([]*logConsumer, 0, len(c.workerMap))
	for _, worker := range c.workerMap {
		workers = append(workers, worker)
	}
	c.stateMu.RUnlock()

	for _, worker := range workers {
		worker.run()
	}
}

func (c *consumer) newLogConsumer(consumeShard *tls.ConsumeShard) *logConsumer {
	// Use the Start-derived runtime context (if set) so workers sense the
	// Stop() cancel and can abort blocking sends to heartbeatExpiredCh /
	// commitCh. Fall back to the parent context if Start() has not run yet
	// (e.g. tests that call handleShards directly without Start).
	ctx := c.runtimeCtx
	if ctx == nil {
		ctx = c.ctx
	}
	lc := &logConsumer{
		ctx:                ctx,
		wg:                 c.workerWG, // shared consumer-level WaitGroup; constant space
		client:             c.client,
		logger:             c.logger,
		conf:               c.conf,
		statusLock:         &sync.RWMutex{},
		status:             pending,
		shard:              consumeShard,
		consumeFunc:        c.consumeFunc,
		checkpoint:         c.checkpoint,
		heartbeatRestartCh: c.heartbeatExpiredCh,
		commitCh:           c.commitCh,
		nextCheckpoint:     "",
		currLogGroupList:   nil,
		stats:              c.stats,
	}
	return lc
}

func (c *consumer) init() error {

	describeConsumerGroupsRes, err := c.client.DescribeConsumerGroups(&tls.DescribeConsumerGroupsRequest{
		ProjectID:         c.conf.ProjectID,
		ConsumerGroupName: c.conf.ConsumerGroupName,
	})

	if err != nil {
		return err
	}

	for _, consumerGroup := range describeConsumerGroupsRes.ConsumerGroups {
		if consumerGroup.ConsumerGroupName == c.conf.ConsumerGroupName {
			return nil
		}
	}

	_, err = c.client.CreateConsumerGroup(&tls.CreateConsumerGroupRequest{
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
	// Capture the current generation before calling sync.Once. A pointer is
	// used instead of resetting sync.Once in place: once a generation is fully
	// closed, Start() can publish a new pointer without racing an old Stop().
	c.lifecycleMu.Lock()
	stopOnce := c.stopOnce
	c.lifecycleMu.Unlock()
	if stopOnce == nil {
		return
	}
	stopOnce.Do(func() {
		c.stopGeneration(stopOnce)
	})
}

// stopGeneration performs one generation's shutdown. The generation pointer
// check is important for a Stop() that captured an old closed generation just
// before a concurrent Start() published a new one: that stale Stop must not
// stop the new run.
func (c *consumer) stopGeneration(stopOnce *sync.Once) {
	c.lifecycleMu.Lock()
	if c.stopOnce != stopOnce {
		c.lifecycleMu.Unlock()
		return
	}

	stage := atomic.LoadInt32(&c.lifecycleStage)
	switch stage {
	case lifecycleClosing, lifecycleClosed:
		// Already stopping or stopped; nothing to do.
		c.lifecycleMu.Unlock()
		return
	case lifecycleNew:
		// Stop-before-Start enters closing first and uses the same finalizer as
		// a started run. This closes the freeze window: Start() cannot pass
		// through while the final snapshot is being published.
		c.metricsMu.Lock()
		atomic.StoreInt32(&c.lifecycleStage, lifecycleClosing)
		atomic.StoreInt32(&c.started, 0)
		atomic.StoreInt32(&c.closing, 1)
		atomic.StoreInt32(&c.closed, 0)
		c.metricsMu.Unlock()
		c.lifecycleMu.Unlock()
		c.finalizeStop(stopOnce)
		return
	case lifecycleStarting, lifecycleRunning:
		// Serialize this transition with SetMetricsLevel so no setter can
		// publish a new level after Stop() has begun finalization. Capture all
		// per-run resources while lifecycleMu is held; a later restart cannot
		// replace them until this generation reaches lifecycleClosed.
		c.metricsMu.Lock()
		atomic.StoreInt32(&c.lifecycleStage, lifecycleClosing)
		atomic.StoreInt32(&c.closing, 1)
		atomic.StoreInt32(&c.closed, 0)
		launched := c.runLaunched
		cancel := c.cancel
		mainLoopDone := c.mainLoopDone
		workerWG := c.workerWG
		wg := c.wg
		checkpointStopCh := c.checkpoint.stopCh
		c.metricsMu.Unlock()
		c.lifecycleMu.Unlock()

		// Cancel the derived runtime context. Start() publishes it before the
		// starting state, so this is non-nil for both starting and running.
		if cancel != nil {
			cancel()
		}

		if !launched {
			// Stop raced init(): no goroutine was launched, so there is no main
			// loop, worker task, or checkpoint manager to wait for.
			c.finalizeStop(stopOnce)
			return
		}

		// Wait for the scheduler first. Once mainLoopDone closes, no new
		// workerWG.Add(1) calls can occur, making the subsequent Wait safe.
		<-mainLoopDone

		// Exactly one finalizer waits for workers, asks checkpoint to perform
		// its final flush, waits for all component goroutines, then publishes
		// closed/final. On timeout this goroutine remains the sole finalizer.
		finalizeDone := make(chan struct{})
		tls.GoWithRecovery(func() {
			workerWG.Wait()
			checkpointStopCh <- struct{}{}
			wg.Wait()
			c.finalizeStop(stopOnce)
			close(finalizeDone)
		})

		select {
		case <-finalizeDone:
			return
		case <-time.After(c.stopTimeout):
			_ = level.Error(c.logger).Log("msg", "consumer stop timeout: workers still running; staying closing/weak, finalizing async", "timeout", c.stopTimeout)
			return
		}
	}
	c.lifecycleMu.Unlock()
}

// finalizeStop releases held batches for all workers still in workerMap,
// clears the heartbeat assignment, freezes the snapshot clock, and marks the
// consumer closed. It is called only after the current generation's workers
// and component goroutines have stopped (or directly for Stop-before-Start).
// The generation check prevents a stale finalizer from touching a restarted
// consumer.
func (c *consumer) finalizeStop(stopOnce *sync.Once) {
	c.lifecycleMu.Lock()
	if c.stopOnce != stopOnce || atomic.LoadInt32(&c.lifecycleStage) != lifecycleClosing {
		c.lifecycleMu.Unlock()
		return
	}
	c.lifecycleMu.Unlock()

	c.stateMu.Lock()
	for _, lc := range c.workerMap {
		lc.releaseHeldBatchIfAny()
	}
	c.workerMap = make(map[string]*logConsumer)
	c.stateMu.Unlock()
	if c.heartbeat != nil {
		c.heartbeat.clearShards()
	}

	// Publish the final snapshot and the closed lifecycle state under the
	// lifecycle -> metrics lock order. GetStats holds metricsMu.RLock, so it
	// cannot observe closed before the freeze or running while the sequence is
	// still frozen.
	c.lifecycleMu.Lock()
	defer c.lifecycleMu.Unlock()
	if c.stopOnce != stopOnce || atomic.LoadInt32(&c.lifecycleStage) != lifecycleClosing {
		return
	}
	c.metricsMu.Lock()
	if c.stats != nil {
		c.stats.freezeOnClose()
		atomic.StoreInt64(&c.stats.assignedShards, 0)
	}
	c.cancel = nil
	c.runtimeCtx = nil
	c.runLaunched = false
	atomic.StoreInt32(&c.started, 0)
	atomic.StoreInt32(&c.closing, 0)
	atomic.StoreInt32(&c.closed, 1)
	atomic.StoreInt32(&c.lifecycleStage, lifecycleClosed)
	c.metricsMu.Unlock()
}

// lifecycleState returns the current lifecycle state for GetStats.
func (c *consumer) lifecycleState() consumerState {
	if atomic.LoadInt32(&c.closed) == 1 {
		return stateClosed
	}
	if atomic.LoadInt32(&c.closing) == 1 {
		return stateClosing
	}
	if atomic.LoadInt32(&c.started) == 1 {
		return stateRunning
	}
	return stateNew
}

// checkpointSnapshot returns a queueMeasurements snapshot of the pending checkpoint
// queue, or nil if the checkpoint manager has no queue.
func (c *consumer) checkpointSnapshot() *queueMeasurements {
	if c.checkpoint == nil {
		return nil
	}
	return c.checkpoint.snapshot()
}
