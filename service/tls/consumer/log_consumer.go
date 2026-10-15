package consumer

import (
	"context"
	"errors"
	"net/http"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

var (
	errBusy             = errors.New("server is busy")
	errHeartbeatExpired = errors.New("heartbeat expired")
)

type logConsumer struct {
	ctx    context.Context
	wg     *sync.WaitGroup
	client tls.Client
	logger log.Logger

	conf               *Config
	statusLock         *sync.RWMutex
	status             int
	shard              *tls.ConsumeShard
	consumeFunc        func(topicID string, shardID int, l *pb.LogGroupList)
	checkpoint         *checkpointManager
	heartbeatRestartCh chan<- struct{}
	commitCh           chan<- struct{}

	nextCheckpoint            string
	currLogGroupList          *pb.LogGroupList
	lastBackoffTime           time.Time
	dispatchStartTime         time.Time // set when batch handed to SDK; read by consume via statusLock happens-before
	dispatchLatencyGeneration uint64

	stats *consumerStatsState

	// retired is set to 1 (atomically) when the worker is removed from
	// workerMap by rebalance or shutdown. An in-flight task that observes
	// retired must not leave a held/ready batch; it calls
	// releaseHeldBatchIfAny() on exit. Constant-space per worker.
	retired int32

	// held is 1 when this worker currently holds a fetched, not-yet-consumed
	// batch (i.e. it has incremented the global heldBatchRefs/readyToProcess
	// counters). releaseHeldBatch() uses a CAS on held 1->0 to decrement the
	// global counters exactly once per batch; both consume() (normal pickup)
	// and releaseHeldBatchIfAny() (rebalance/shutdown/retired-exit) funnel
	// through it, so concurrent release attempts cannot double-decrement.
	held           int32
	heldGeneration uint64
}

func (lc *logConsumer) run() {
	switch lc.loadStatus() {
	case pending:
		lc.setStatus(initializing)
		lc.wg.Add(1)
		go lc.runWithStatus(lc.init, readyToFetch, pending, pending, nop)
	case initializing:
	case readyToFetch:
		lc.setStatus(fetching)
		lc.wg.Add(1)
		go lc.runWithStatus(lc.fetchData, readyToConsume, readyToFetch, readyToFetch, backoff)
	case fetching:
	case readyToConsume:
		lc.setStatus(consuming)
		lc.wg.Add(1)
		go lc.runWithStatus(lc.consume, readyToFetch, readyToConsume, readyToConsume, nop)
	case consuming:
	case backoff:
		lc.wg.Add(1)
		go lc.runWithStatus(lc.backoff, readyToFetch, nop, backoff, backoff)
	case waitForRestart:
	}
}

// markRetired flags the worker as removed from workerMap. Called under
// stateMu during rebalance/shutdown. After this is set, any in-flight task
// that holds a batch will release it on exit (see releaseIfRetired).
func (lc *logConsumer) markRetired() {
	atomic.StoreInt32(&lc.retired, 1)
}

// releaseIfRetired is deferred in every worker task goroutine. If the worker
// was retired (removed by rebalance/shutdown) while the task ran, the task
// may have left a held/ready batch that will never be consumed — release it
// idempotently so the global counters stay balanced. No-op if not retired.
func (lc *logConsumer) releaseIfRetired() {
	if atomic.LoadInt32(&lc.retired) == 1 {
		lc.releaseHeldBatchIfAny()
	}
}

// runWithStatus runs f and transitions the worker status based on the result.
// It owns worker completion: the shared workerWG is always Done()'d on exit,
// and releaseIfRetired() runs when stats is active (held/retired exists only for
// stats balancing).
//
// Defer ordering (LIFO): the body recover is registered last so it runs first,
// then releaseIfRetired, then wg.Done(). This preserves body panic recovery
// (status set to ifPanic) without swallowing WaitGroup bookkeeping: Done() is
// unconditional and outermost so it always runs, even if releaseIfRetired or the
// recover itself misbehave. releaseIfRetired completes before Done() so a
// retiring worker's held batch is balanced before workerWG reaches zero.
func (lc *logConsumer) runWithStatus(f func() error, ifDone int, ifErr int, ifPanic int, ifBusy int) {
	defer lc.wg.Done()
	if lc.stats != nil {
		defer lc.releaseIfRetired()
	}
	defer func() {
		if r := recover(); r != nil {
			lc.setStatus(ifPanic)

			level.Error(lc.logger).Log("panic", "panic happened during fetching data. info: "+string(debug.Stack()))
		}
	}()

	if err := f(); err != nil {
		switch err {
		case errBusy:
			lc.setStatus(ifBusy)
		case errHeartbeatExpired:
			lc.setStatus(waitForRestart)
		default:
			lc.setStatus(ifErr)
		}

		return
	}

	lc.setStatus(ifDone)
}

func (lc *logConsumer) setStatus(status int) {
	lc.statusLock.Lock()
	defer lc.statusLock.Unlock()

	oldStatus := lc.status
	lc.status = status
	if oldStatus != backoff && status == backoff && lc.stats != nil {
		lc.stats.incBackoffEvents()
	}
}

func (lc *logConsumer) loadStatus() int {
	lc.statusLock.RLock()
	defer lc.statusLock.RUnlock()

	return lc.status
}

func (lc *logConsumer) init() error {
	checkpointResp, err := lc.client.DescribeCheckPoint(&tls.DescribeCheckPointRequest{
		ProjectID:         lc.conf.ProjectID,
		TopicID:           lc.shard.TopicID,
		ConsumerGroupName: lc.conf.ConsumerGroupName,
		ShardID:           lc.shard.ShardID,
	})
	if err != nil {
		_ = level.Error(lc.logger).Log("msg", "init get checkpoint failed", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "error", err.Error())

		return err
	}

	if len(checkpointResp.Checkpoint) != 0 {
		lc.nextCheckpoint = checkpointResp.Checkpoint

		return nil
	}

	descCursorResp, err := lc.client.DescribeCursor(&tls.DescribeCursorRequest{
		TopicID: lc.shard.TopicID,
		ShardID: lc.shard.ShardID,
		From:    lc.conf.ConsumeFrom,
	})
	if err != nil {
		_ = level.Error(lc.logger).Log("msg", "init get cursor failed", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "from", lc.conf.ConsumeFrom, "error", err.Error())

		return err
	}

	lc.nextCheckpoint = descCursorResp.Cursor
	lc.lastBackoffTime = time.Time{}

	return nil
}

func (lc *logConsumer) fetchData() error {
	var fetchGeneration uint64
	fetchFailed := true
	if lc.stats != nil {
		fetchGeneration = lc.stats.beginFetchCall()
		defer func() { lc.stats.finishFetchCall(fetchGeneration, fetchFailed) }()
	}
	compressType := lc.conf.CompressType
	consumeReq := tls.ConsumeLogsRequest{
		TopicID:           lc.shard.TopicID,
		ShardID:           lc.shard.ShardID,
		Cursor:            lc.nextCheckpoint,
		LogGroupCount:     &lc.conf.MaxFetchLogGroupCount,
		Compression:       &compressType,
		ConsumerGroupName: &lc.conf.ConsumerGroupName,
		ConsumerName:      &lc.conf.ConsumerName,
		Original:          lc.conf.Original,
	}

	_ = level.Debug(lc.logger).Log("msg", "fetch start", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "req_cursor", consumeReq.Cursor, "count_req", lc.conf.MaxFetchLogGroupCount, "compress", compressType, "original", lc.conf.Original)

	// fetch_call_duration: time for one ConsumeLogs call (incl. internal retry).
	// time.Now() only when latency is on (off path: no clock call).
	var fetchStart time.Time
	var fetchLatencyGeneration uint64
	if lc.stats != nil {
		fetchLatencyGeneration = lc.stats.currentDetailedGeneration()
	}
	if fetchLatencyGeneration != 0 {
		fetchStart = lc.stats.nowFunc()
	}
	fetchResp, err := lc.client.ConsumeLogs(&consumeReq)
	if lc.stats != nil && lc.stats.detailedGenerationEnabled(fetchLatencyGeneration) {
		lc.stats.fetchCallDuration.observeSince(fetchStart, lc.stats.nowFunc)
	}
	if err != nil {
		clientErr := tls.NewClientError(err)
		if clientErr.HTTPCode == http.StatusTooManyRequests {
			lc.lastBackoffTime = time.Now()
			return errBusy
		}

		if clientErr.Code == tls.ErrConsumerHeartbeatExpired {
			// Both channels are unbuffered and their receivers (heartbeat runner
			// / checkpoint manager) may have exited after Stop() cancelled the
			// runtime context. A blocking send here would wedge the worker
			// forever, so workerWG never drains and Stop() never finalizes.
			// Make the sends cancel-aware: if the context is done, abort and
			// return errHeartbeatExpired so the worker exits cleanly.
			select {
			case lc.heartbeatRestartCh <- struct{}{}:
			case <-lc.ctx.Done():
				return errHeartbeatExpired
			}
			select {
			case lc.commitCh <- struct{}{}:
			case <-lc.ctx.Done():
				return errHeartbeatExpired
			}
			return errHeartbeatExpired
		}

		if clientErr.Code == tls.ErrInvalidContent {
			var groups int
			if fetchResp != nil && fetchResp.Logs != nil {
				groups = len(fetchResp.Logs.LogGroups)
			}
			_ = level.Error(lc.logger).Log("msg", "fetch invalid data", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "req_cursor", lc.nextCheckpoint, "resp_count", fetchResp.Count, "resp_cursor", fetchResp.Cursor, "resp_groups", groups, "request_id", fetchResp.RequestID)
			lc.currLogGroupList = fetchResp.Logs
			// dispatch_delay start: batch handed to SDK (ErrInvalidContent path).
			var dispatchGeneration uint64
			if lc.stats != nil {
				dispatchGeneration = lc.stats.currentDetailedGeneration()
			}
			if dispatchGeneration != 0 {
				lc.dispatchStartTime = lc.stats.nowFunc()
				lc.dispatchLatencyGeneration = dispatchGeneration
			} else {
				lc.dispatchStartTime = time.Time{}
				lc.dispatchLatencyGeneration = 0
			}
			if len(fetchResp.Cursor) != 0 {
				lc.nextCheckpoint = fetchResp.Cursor
			}
			// Mirror the success path: account for the held batch so consume()
			// can balance the counters. Only non-empty batches are held.
			if lc.stats != nil && fetchResp.Logs != nil && len(fetchResp.Logs.LogGroups) > 0 {
				if lc.stats.generationEnabled(fetchGeneration) {
					lc.stats.incHeldBatchRefs()
					lc.stats.incReadyToProcess()
					atomic.StoreUint64(&lc.heldGeneration, fetchGeneration)
					atomic.StoreInt32(&lc.held, 1)
				}
			}
			return nil
		}

		_ = level.Error(lc.logger).Log("msg", "fetch failed", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "req_cursor", lc.nextCheckpoint, "compress", compressType, "original", lc.conf.Original, "http_code", clientErr.HTTPCode, "code", clientErr.Code, "error", err.Error())

		return err
	}

	lc.currLogGroupList = fetchResp.Logs
	// dispatch_delay start: batch handed to SDK. Set before status change
	// (statusLock provides happens-before for consume to read). time.Now()
	// only when latency is on.
	var dispatchGeneration uint64
	if lc.stats != nil {
		dispatchGeneration = lc.stats.currentDetailedGeneration()
	}
	if dispatchGeneration != 0 {
		lc.dispatchStartTime = lc.stats.nowFunc()
		lc.dispatchLatencyGeneration = dispatchGeneration
	} else {
		lc.dispatchStartTime = time.Time{}
		lc.dispatchLatencyGeneration = 0
	}
	if len(fetchResp.Cursor) != 0 {
		lc.nextCheckpoint = fetchResp.Cursor
	}
	if lc.stats != nil && fetchResp.Logs != nil {
		groups := int64(len(fetchResp.Logs.LogGroups))
		var logs int64
		for _, lg := range fetchResp.Logs.LogGroups {
			logs += int64(len(lg.Logs))
		}
		lc.stats.addFetchedFor(fetchGeneration, groups, logs)
		// Worker now holds a batch and is ready to process. Only non-empty
		// batches are held; consume() returns early on empty batches and
		// would otherwise never balance these counters.
		if groups > 0 {
			if lc.stats.generationEnabled(fetchGeneration) {
				lc.stats.incHeldBatchRefs()
				lc.stats.incReadyToProcess()
				atomic.StoreUint64(&lc.heldGeneration, fetchGeneration)
				atomic.StoreInt32(&lc.held, 1)
			}
		}
	}
	fetchFailed = false
	_ = level.Debug(lc.logger).Log("msg", "fetch ok", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "req_cursor", consumeReq.Cursor, "resp_count", fetchResp.Count, "resp_cursor", fetchResp.Cursor)

	return nil
}

func (lc *logConsumer) consume() error {
	if lc.currLogGroupList == nil {
		return nil
	}

	if len(lc.currLogGroupList.LogGroups) == 0 {
		return nil
	}

	// Picking up the held batch for processing: no longer "held" or "ready".
	// Funnel through releaseHeldBatch() so consume() competes with the
	// rebalance/shutdown/retired-exit path on the same CAS 1->0. Exactly one
	// winner decrements the global counters per batch; if rebalance already
	// released the batch, this CAS fails and consume() must NOT decrement
	// again (no double-decrement / underflow).
	lc.releaseHeldBatch()

	// dispatch_delay: time from batch handed to SDK to process callback start.
	// Read dispatchStartTime (set by fetchData, published via statusLock
	// happens-before), observe, then clear to avoid stale value for next batch.
	if lc.stats != nil && lc.stats.detailedGenerationEnabled(lc.dispatchLatencyGeneration) &&
		!lc.dispatchStartTime.IsZero() {
		lc.stats.dispatchDelay.observeSince(lc.dispatchStartTime, lc.stats.nowFunc)
	}
	lc.dispatchStartTime = time.Time{}
	lc.dispatchLatencyGeneration = 0

	// process_duration: time for user callback ONLY (not terminal accounting/
	// addCheckpoint). Private method scopes the defer to consumeFunc; defer
	// still observes on panic (defer runs during panic unwind). No closure
	// allocation: method receiver is existing pointer, defer args evaluated
	// at registration (no capture).
	groups := int64(len(lc.currLogGroupList.LogGroups))
	var logs int64
	for _, group := range lc.currLogGroupList.LogGroups {
		logs += int64(len(group.Logs))
	}
	lc.processCurrentBatch(groups, logs)
	if len(lc.nextCheckpoint) != 0 {
		lc.checkpoint.addCheckpoint(&checkpointInfo{
			shardInfo:  lc.shard,
			checkpoint: lc.nextCheckpoint,
		})
		_ = level.Debug(lc.logger).Log("msg", "checkpoint enqueue", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "checkpoint", lc.nextCheckpoint, "groups", len(lc.currLogGroupList.LogGroups))
	}

	return nil
}

// processCurrentBatch scopes process in-flight and outcome facts to the user
// callback. The deferred finish also records a failure if the callback panics;
// checkpoint bookkeeping is deliberately outside this boundary.
func (lc *logConsumer) processCurrentBatch(groups, logs int64) {
	var processGeneration uint64
	processSucceeded := false
	if lc.stats != nil {
		processGeneration = lc.stats.beginProcessAttempt()
		defer func() {
			lc.stats.finishProcessAttempt(processGeneration, groups, logs, processSucceeded)
		}()
	}
	var processLatencyGeneration uint64
	if lc.stats != nil {
		processLatencyGeneration = lc.stats.currentDetailedGeneration()
	}
	lc.consumeWithTiming(processLatencyGeneration)
	processSucceeded = true
}

// consumeWithTiming invokes consumeFunc, optionally measuring process_duration.
// The defer is scoped to this method so it times ONLY the callback (not
// terminal accounting/checkpoint work which run in the caller). On panic the
// defer still observes (defer runs during unwind). No closure allocation:
// lc is the existing receiver pointer, procStart is passed to defer by value.
func (lc *logConsumer) consumeWithTiming(latencyGeneration uint64) {
	if latencyGeneration == 0 {
		lc.consumeFunc(lc.shard.TopicID, lc.shard.ShardID, lc.currLogGroupList)
		return
	}
	procStart := lc.stats.nowFunc()
	defer lc.observeProcessDuration(procStart, latencyGeneration)
	lc.consumeFunc(lc.shard.TopicID, lc.shard.ShardID, lc.currLogGroupList)
}

func (lc *logConsumer) observeProcessDuration(start time.Time, generation uint64) {
	if lc.stats != nil && lc.stats.detailedGenerationEnabled(generation) {
		lc.stats.processDuration.observeSince(start, lc.stats.nowFunc)
	}
}

// releaseHeldBatch atomically claims ownership of releasing the held batch by
// CAS-ing the per-worker `held` flag 1->0. The winner — exactly one of normal
// consume pickup, rebalance, shutdown, or a retired task's exit path —
// decrements the global heldBatchRefs/readyToProcess counters once. Returns
// true if this call won the CAS and decremented; false if held was already 0
// (batch already released by a concurrent path, or never held). This is the
// single release path: both consume() and releaseHeldBatchIfAny() funnel
// through it so concurrent release attempts cannot double-decrement (which
// would underflow the uint64 counters read by GetStats).
func (lc *logConsumer) releaseHeldBatch() bool {
	if lc.stats == nil {
		return false
	}
	if !atomic.CompareAndSwapInt32(&lc.held, 1, 0) {
		return false // no batch currently held, or already released
	}
	generation := atomic.LoadUint64(&lc.heldGeneration)
	if lc.stats.generationEnabled(generation) {
		lc.stats.decHeldBatchRefs()
		lc.stats.decReadyToProcess()
	}
	return true
}

// releaseHeldBatchIfAny balances the held-batch counters when a worker is
// removed (rebalance/shutdown) before it could consume its fetched batch, or
// when a retired worker's task exits having just fetched a batch that will
// never be consumed. Idempotent: it funnels through releaseHeldBatch(), whose
// CAS on the per-worker `held` flag ensures the global counters are
// decremented at most once per batch, so concurrent calls from the rebalance
// path, the shutdown path, and the retiring task's exit path cannot
func (lc *logConsumer) releaseHeldBatchIfAny() {
	lc.releaseHeldBatch()
}

func (lc *logConsumer) backoff() error {
	if time.Since(lc.lastBackoffTime) > time.Second*5 {
		return nil
	}

	return errBusy
}
