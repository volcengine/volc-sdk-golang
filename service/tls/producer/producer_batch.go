package producer

import (
	"sync"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func normalizeLogTime(t int64) int64 {
	if t <= 0 {
		return 0
	}
	if t < int64(1e10) {
		return t * 1000
	}
	if t < int64(1e15) {
		return t
	}
	return t / int64(1e6)
}

type Batch struct {
	totalDataSize              int64
	lock                       sync.RWMutex
	logGroupList               *pb.LogGroupList
	logGroupSizes              []int
	logCount                   int
	minLogTime                 int64
	maxLogTime                 int64
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
		logGroupSizes:              []int{},
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

	return producerBatch
}

func newProducerBatch(batchLog *BatchLog, config *Config) *Batch {
	currentTime := time.Now()
	producerBatch := &Batch{
		logGroupList:               &pb.LogGroupList{LogGroups: []*pb.LogGroup{}},
		logGroupSizes:              []int{},
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

func (batch *Batch) getMinLogTime() int64 {
	defer batch.lock.RUnlock()
	batch.lock.RLock()

	return batch.minLogTime
}

func (batch *Batch) getMaxLogTime() int64 {
	defer batch.lock.RUnlock()
	batch.lock.RLock()

	return batch.maxLogTime
}

func (batch *Batch) tryAddLog(batchLog *BatchLog, callBack CallBack) (bool, int64) {
	batch.lock.Lock()
	defer batch.lock.Unlock()

	group, groupSize, newGroup := batch.getOrCreateTailGroup(batchLog.Key.Source, batchLog.Key.FileName, batchLog.Key.ContextFlow)
	if !newGroup && len(group.Logs) >= MaxLogGroupCount {
		group = nil
	}
	if group == nil {
		group = batch.newGroup(batchLog.Key.Source, batchLog.Key.FileName, batchLog.Key.ContextFlow)
		groupSize = group.Size()
		newGroup = true
	}

	logMsgSize := batchLog.Log.Size()
	logFieldSize := sizeOfBytesField(logMsgSize)
	if batch.logCount+1 > MaxBatchCount {
		return false, 0
	}

	var deltaTotal int
	var newGroupSize int
	if newGroup {
		newGroupSize = groupSize + logFieldSize
		deltaTotal = sizeOfBytesField(newGroupSize)
	} else {
		oldGroupSize := batch.logGroupSizes[len(batch.logGroupSizes)-1]
		newGroupSize = oldGroupSize + logFieldSize
		deltaTotal = logFieldSize + (varintLen(newGroupSize) - varintLen(oldGroupSize))
	}

	if batch.totalDataSize+int64(deltaTotal) > MaxBatchSize {
		return false, 0
	}

	if newGroup {
		batch.logGroupList.LogGroups = append(batch.logGroupList.LogGroups, group)
		batch.logGroupSizes = append(batch.logGroupSizes, newGroupSize)
	} else {
		batch.logGroupSizes[len(batch.logGroupSizes)-1] = newGroupSize
	}

	group.Logs = append(group.Logs, batchLog.Log)
	batch.logCount++
	batch.totalDataSize += int64(deltaTotal)

	normalizedTime := normalizeLogTime(batchLog.Log.Time)
	if batch.logCount == 1 {
		batch.minLogTime = normalizedTime
		batch.maxLogTime = normalizedTime
	} else {
		if normalizedTime < batch.minLogTime {
			batch.minLogTime = normalizedTime
		}
		if normalizedTime > batch.maxLogTime {
			batch.maxLogTime = normalizedTime
		}
	}

	if callBack != nil {
		batch.callBackList = append(batch.callBackList, callBack)
	}

	return true, int64(deltaTotal)
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

func (batch *Batch) getOrCreateTailGroup(source, fileName, contextFlow string) (*pb.LogGroup, int, bool) {
	if len(batch.logGroupList.LogGroups) == 0 {
		group := batch.newGroup(source, fileName, contextFlow)
		return group, group.Size(), true
	}
	lastGroup := batch.logGroupList.LogGroups[len(batch.logGroupList.LogGroups)-1]
	if lastGroup.Source == source && lastGroup.FileName == fileName && lastGroup.ContextFlow == contextFlow {
		return lastGroup, batch.logGroupSizes[len(batch.logGroupSizes)-1], false
	}
	group := batch.newGroup(source, fileName, contextFlow)
	return group, group.Size(), true
}

func sizeOfBytesField(payloadSize int) int {
	return 1 + varintLen(payloadSize) + payloadSize
}

func varintLen(x int) int {
	if x < 0 {
		return 10
	}
	n := 1
	for x >= 1<<7 {
		n++
		x >>= 7
	}
	return n
}
