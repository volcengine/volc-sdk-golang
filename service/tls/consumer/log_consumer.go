package consumer

import (
	"context"
	"errors"
	"net/http"
	"runtime/debug"
	"sync"
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

	nextCheckpoint   string
	currLogGroupList *pb.LogGroupList
	lastBackoffTime  time.Time
}

func (lc *logConsumer) run() {
	switch lc.loadStatus() {
	case pending:
		lc.setStatus(initializing)
		go lc.runWithStatus(lc.init, readyToFetch, pending, pending, nop)
	case initializing:
	case readyToFetch:
		lc.setStatus(fetching)
		go lc.runWithStatus(lc.fetchData, readyToConsume, readyToFetch, readyToFetch, backoff)
	case fetching:
	case readyToConsume:
		lc.setStatus(consuming)
		lc.wg.Add(1)
		tls.GoWithRecovery(func() {
			defer lc.wg.Done()
			lc.runWithStatus(lc.consume, readyToFetch, readyToConsume, readyToConsume, nop)
		})
	case consuming:
	case backoff:
		go lc.runWithStatus(lc.backoff, readyToFetch, nop, backoff, backoff)
	case waitForRestart:
	}
}

func (lc *logConsumer) runWithStatus(f func() error, ifDone int, ifErr int, ifPanic int, ifBusy int) {
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

	lc.status = status
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

	fetchResp, err := lc.client.ConsumeLogs(&consumeReq)
	if err != nil {
		clientErr := tls.NewClientError(err)
		if clientErr.HTTPCode == http.StatusTooManyRequests {
			lc.lastBackoffTime = time.Now()
			return errBusy
		}

		if clientErr.Code == tls.ErrConsumerHeartbeatExpired {
			lc.heartbeatRestartCh <- struct{}{}
			lc.commitCh <- struct{}{}
			return errHeartbeatExpired
		}

		if clientErr.Code == tls.ErrInvalidContent {
			var groups int
			if fetchResp != nil && fetchResp.Logs != nil {
				groups = len(fetchResp.Logs.LogGroups)
			}
			_ = level.Error(lc.logger).Log("msg", "fetch invalid data", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "req_cursor", lc.nextCheckpoint, "resp_count", fetchResp.Count, "resp_cursor", fetchResp.Cursor, "resp_groups", groups, "request_id", fetchResp.RequestID)
			lc.currLogGroupList = fetchResp.Logs
			if len(fetchResp.Cursor) != 0 {
				lc.nextCheckpoint = fetchResp.Cursor
			}
			return nil
		}

		_ = level.Error(lc.logger).Log("msg", "fetch failed", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "req_cursor", lc.nextCheckpoint, "compress", compressType, "original", lc.conf.Original, "http_code", clientErr.HTTPCode, "code", clientErr.Code, "error", err.Error())

		return err
	}

	lc.currLogGroupList = fetchResp.Logs
	if len(fetchResp.Cursor) != 0 {
		lc.nextCheckpoint = fetchResp.Cursor
	}
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

	lc.consumeFunc(lc.shard.TopicID, lc.shard.ShardID, lc.currLogGroupList)
	if len(lc.nextCheckpoint) != 0 {
		lc.checkpoint.addCheckpoint(&checkpointInfo{
			shardInfo:  lc.shard,
			checkpoint: lc.nextCheckpoint,
		})
		_ = level.Debug(lc.logger).Log("msg", "checkpoint enqueue", "topic", lc.shard.TopicID, "shard", lc.shard.ShardID, "checkpoint", lc.nextCheckpoint, "groups", len(lc.currLogGroupList.LogGroups))
	}

	return nil
}

func (lc *logConsumer) backoff() error {
	if time.Since(lc.lastBackoffTime) > time.Second*5 {
		return nil
	}

	return errBusy
}
