package producer

import (
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"sync"
	"time"
)

type Batch struct {
	totalDataSize              int64
	lock                       sync.RWMutex
	logGroup                   *pb.LogGroup
	attemptCount               int
	retryBackoffMs             int64
	baseRetryBackoffMs         int64
	baseIncreaseRetryBackoffMs int64
	nextRetryMs                int64
	maxRetryIntervalInMs       int64
	callBackList               []CallBack
	createTime                 time.Time
	maxRetryTimes              int
	topic                      string
	shardHash                  *string
	result                     *Result
	maxReservedAttempts        int
}

func initProducerBatch(batchLog *BatchLog, config *Config) *Batch {
	var logs = []*pb.Log{batchLog.Log}

	logGroup := &pb.LogGroup{
		Logs:        logs,
		Source:      batchLog.Key.Source,
		FileName:    batchLog.Key.FileName,
		ContextFlow: batchLog.Key.ContextFlow,
	}

	currentTime := time.Now()
	producerBatch := &Batch{
		logGroup:                   logGroup,
		attemptCount:               0,
		maxRetryIntervalInMs:       config.MaxRetryBackoffMs,
		callBackList:               []CallBack{},
		createTime:                 currentTime,
		maxRetryTimes:              config.Retries,
		retryBackoffMs:             0,
		baseRetryBackoffMs:         config.BaseRetryBackoffMs,
		baseIncreaseRetryBackoffMs: 1000,
		topic:                      batchLog.Key.Topic,
		result:                     newResult(),
		maxReservedAttempts:        config.MaxReservedAttempts,
	}

	producerBatch.shardHash = parseHash(batchLog.Key.ShardHash)
	producerBatch.totalDataSize = int64(producerBatch.logGroup.Size())

	if batchLog.Key.CallBackFun != nil {
		producerBatch.callBackList = append(producerBatch.callBackList, batchLog.Key.CallBackFun)
	}

	return producerBatch
}

func newProducerBatch(batchLog *BatchLog, config *Config) *Batch {

	logGroup := &pb.LogGroup{
		Source:      batchLog.Key.Source,
		FileName:    batchLog.Key.FileName,
		ContextFlow: batchLog.Key.ContextFlow,
	}

	currentTime := time.Now()
	producerBatch := &Batch{
		logGroup:                   logGroup,
		attemptCount:               0,
		maxRetryIntervalInMs:       config.MaxRetryBackoffMs,
		callBackList:               []CallBack{},
		createTime:                 currentTime,
		maxRetryTimes:              config.Retries,
		retryBackoffMs:             0,
		baseRetryBackoffMs:         config.BaseRetryBackoffMs,
		baseIncreaseRetryBackoffMs: 1000,
		topic:                      batchLog.Key.Topic,
		result:                     newResult(),
		maxReservedAttempts:        config.MaxReservedAttempts,
	}

	producerBatch.shardHash = parseHash(batchLog.Key.ShardHash)

	return producerBatch
}

func (batch *Batch) getLogCount() int {
	defer batch.lock.RUnlock()
	batch.lock.RLock()

	return len(batch.logGroup.GetLogs())
}

func (batch *Batch) addLogToLogGroup(log *pb.Log) {
	defer batch.lock.Unlock()
	batch.lock.Lock()

	batch.logGroup.Logs = append(batch.logGroup.Logs, log)
}

func (batch *Batch) addProducerBatchCallBack(callBack CallBack) {
	defer batch.lock.Unlock()
	batch.lock.Lock()

	batch.callBackList = append(batch.callBackList, callBack)
}

func (batch *Batch) tryAddLog(log *pb.Log, size int64, callBack CallBack) bool {
	if !batch.hasRoomFor(size, 1) {
		return false
	}
	batch.logGroup.Logs = append(batch.logGroup.Logs, log)
	batch.totalDataSize += size
	if callBack != nil {
		batch.callBackList = append(batch.callBackList, callBack)
	}
	return true
}

func (batch *Batch) hasRoomFor(size int64, cnt int) bool {
	return batch.totalDataSize+size <= MaxBatchSize && len(batch.logGroup.Logs)+cnt <= MaxBatchCount
}

func (batch *Batch) meetSendCondition(config *Config) bool {
	return batch.totalDataSize >= config.MaxBatchSize || len(batch.logGroup.Logs) >= config.MaxBatchCount
}

func parseHash(inputHash string) *string {
	if inputHash == "" {
		return nil
	}

	return &inputHash
}
