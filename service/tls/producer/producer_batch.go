package producer

import (
	"sync"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type Batch struct {
	totalDataSize              int64
	lock                       sync.RWMutex
	logGroupList               *pb.LogGroupList
	logCount                   int
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
	currentTime := time.Now()
	producerBatch := &Batch{
		logGroupList:               &pb.LogGroupList{LogGroups: []*pb.LogGroup{}},
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
	producerBatch.tryAddLog(batchLog, batchLog.Key.CallBackFun)

	if batchLog.Key.CallBackFun != nil {
		producerBatch.callBackList = append(producerBatch.callBackList, batchLog.Key.CallBackFun)
	}

	return producerBatch
}

func newProducerBatch(batchLog *BatchLog, config *Config) *Batch {
	currentTime := time.Now()
	producerBatch := &Batch{
		logGroupList:               &pb.LogGroupList{LogGroups: []*pb.LogGroup{}},
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

	return batch.logCount
}

func (batch *Batch) addLogToLogGroup(log *pb.Log) {
	defer batch.lock.Unlock()
	batch.lock.Lock()

	if len(batch.logGroupList.LogGroups) == 0 {
		return
	}
	lastGroup := batch.logGroupList.LogGroups[len(batch.logGroupList.LogGroups)-1]
	lastGroup.Logs = append(lastGroup.Logs, log)
	batch.logCount++
}

func (batch *Batch) addProducerBatchCallBack(callBack CallBack) {
	defer batch.lock.Unlock()
	batch.lock.Lock()

	batch.callBackList = append(batch.callBackList, callBack)
}

func (batch *Batch) tryAddLog(batchLog *BatchLog, callBack CallBack) bool {
	group := batch.getOrCreateTailGroup(batchLog.Key.Source, batchLog.Key.FileName, batchLog.Key.ContextFlow)
	if len(group.Logs) >= MaxLogGroupCount {
		group = batch.newGroup(batchLog.Key.Source, batchLog.Key.FileName, batchLog.Key.ContextFlow)
		batch.logGroupList.LogGroups = append(batch.logGroupList.LogGroups, group)
	}

	group.Logs = append(group.Logs, batchLog.Log)
	batch.logCount++
	batch.totalDataSize = int64(batch.logGroupList.Size())

	if !batch.hasRoomFor(0, 0) {
		group.Logs = group.Logs[:len(group.Logs)-1]
		batch.logCount--
		if len(group.Logs) == 0 && len(batch.logGroupList.LogGroups) > 0 && batch.logGroupList.LogGroups[len(batch.logGroupList.LogGroups)-1] == group {
			batch.logGroupList.LogGroups = batch.logGroupList.LogGroups[:len(batch.logGroupList.LogGroups)-1]
		}
		batch.totalDataSize = int64(batch.logGroupList.Size())
		return false
	}

	if callBack != nil {
		batch.callBackList = append(batch.callBackList, callBack)
	}
	return true
}

func (batch *Batch) hasRoomFor(size int64, cnt int) bool {
	return batch.totalDataSize+size <= MaxBatchSize && batch.logCount+cnt <= MaxBatchCount
}

func (batch *Batch) meetSendCondition(config *Config) bool {
	return batch.totalDataSize >= config.MaxBatchSize || batch.logCount >= config.MaxBatchCount
}

func parseHash(inputHash string) *string {
	if inputHash == "" {
		return nil
	}

	return &inputHash
}

func (batch *Batch) newGroup(source, fileName, contextFlow string) *pb.LogGroup {
	return &pb.LogGroup{
		Logs:        []*pb.Log{},
		Source:      source,
		FileName:    fileName,
		ContextFlow: contextFlow,
	}
}

func (batch *Batch) getOrCreateTailGroup(source, fileName, contextFlow string) *pb.LogGroup {
	if len(batch.logGroupList.LogGroups) == 0 {
		group := batch.newGroup(source, fileName, contextFlow)
		batch.logGroupList.LogGroups = append(batch.logGroupList.LogGroups, group)
		return group
	}
	lastGroup := batch.logGroupList.LogGroups[len(batch.logGroupList.LogGroups)-1]
	if lastGroup.Source == source && lastGroup.FileName == fileName && lastGroup.ContextFlow == contextFlow {
		return lastGroup
	}
	group := batch.newGroup(source, fileName, contextFlow)
	batch.logGroupList.LogGroups = append(batch.logGroupList.LogGroups, group)
	return group
}
