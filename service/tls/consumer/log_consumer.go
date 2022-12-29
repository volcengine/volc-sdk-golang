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
	errBusy = errors.New("server is busy")
)

type logConsumer struct {
	ctx    context.Context
	client tls.Client
	logger log.Logger

	conf         *Config
	statusLock   *sync.RWMutex
	status       int
	shard        *tls.ConsumeShard
	consumeFunc  func(topicID string, shardID int, l *pb.LogGroupList)
	checkpointCh chan<- *checkpointInfo

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
		go lc.runWithStatus(lc.consume, readyToFetch, readyToConsume, readyToConsume, nop)
	case consuming:
	case backoff:
		go lc.runWithStatus(lc.backoff, readyToFetch, nop, backoff, backoff)
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
		if errors.Is(err, errBusy) {
			lc.setStatus(ifBusy)
		} else {
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
		level.Error(lc.logger).Log("error", "init log consumer failed in getting checkpoint, err: "+err.Error())

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
		level.Error(lc.logger).Log("error", "init log consumer failed in getting cursor, err: "+err.Error())

		return err
	}

	lc.nextCheckpoint = descCursorResp.Cursor
	lc.lastBackoffTime = time.Time{}

	return nil
}

func (lc *logConsumer) fetchData() error {
	fetchResp, err := lc.client.ConsumeLogs(&tls.ConsumeLogsRequest{
		TopicID:       lc.shard.TopicID,
		ShardID:       lc.shard.ShardID,
		Cursor:        lc.nextCheckpoint,
		LogGroupCount: &lc.conf.MaxFetchLogGroupCount,
		Compression:   nil,
	})
	if err != nil {
		clientErr := tls.NewClientError(err)
		if clientErr.HTTPCode == http.StatusTooManyRequests {
			return errBusy
		}

		level.Error(lc.logger).Log("error", "fetch data failed, err: "+err.Error())

		return err
	}

	lc.currLogGroupList = fetchResp.Logs
	lc.nextCheckpoint = fetchResp.Cursor

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
	lc.checkpointCh <- &checkpointInfo{
		shardInfo:  lc.shard,
		checkpoint: lc.nextCheckpoint,
	}

	return nil
}

func (lc *logConsumer) backoff() error {
	if time.Since(lc.lastBackoffTime) > time.Second*5 {
		return nil
	}

	return errBusy
}
