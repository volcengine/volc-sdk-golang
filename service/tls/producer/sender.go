package producer

import (
	"math"
	"sync/atomic"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type CallBack interface {
	Success(result *Result)
	Fail(result *Result)
}

type Sender struct {
	stopCh               chan struct{}
	maxSender            chan int
	client               Client
	retryQueue           *RetryQueue
	logger               log.Logger
	noRetryStatusCodeMap map[int]struct{}
	producer             *producer
}

func initSender(client Client, retryQueue *RetryQueue, maxSenderCount int64, logger log.Logger, errorStatusMap map[int]struct{}, producer *producer) *Sender {
	return &Sender{
		stopCh:               make(chan struct{}),
		client:               client,
		retryQueue:           retryQueue,
		maxSender:            make(chan int, maxSenderCount),
		logger:               logger,
		noRetryStatusCodeMap: errorStatusMap,
		producer:             producer,
	}
}

func (sender *Sender) IsShutDown() bool {
	select {
	case <-sender.stopCh:
		return true
	default:
		return false
	}
}

func (sender *Sender) sendToServer(batch *Batch) {
	level.Debug(sender.logger).Log("msg", "sender send data to server")
	var err error

	putLogsReq := &PutLogsRequest{
		TopicID:      batch.topic,
		CompressType: "lz4",
		LogBody: &pb.LogGroupList{
			LogGroups: []*pb.LogGroup{
				batch.logGroup,
			},
		},
	}

	if batch.shardHash != nil {
		putLogsReq.HashKey = *batch.shardHash
	}

	resp, err := sender.client.PutLogs(putLogsReq)
	if err == nil {
		sender.handleSuccess(batch, resp)
		return
	}

	sender.handleFailure(batch, err)
}

func (sender *Sender) handleSuccess(batch *Batch, putLogsResp *CommonResponse) {
	level.Debug(sender.logger).Log("msg", "sendToServer succeeded,Execute successful callback function")
	defer atomic.AddInt64(&sender.producer.producerLogGroupSize, -batch.totalDataSize)

	batch.result.SuccessFlag = true
	if batch.attemptCount < batch.maxReservedAttempts {
		batch.result.Attempts = append(batch.result.Attempts,
			newAttempt(true, putLogsResp.RequestID, "", "", GetTimeMs(time.Now().UnixNano())))
	}

	for _, callBack := range batch.callBackList {
		callBack.Success(batch.result)
	}
}

func (sender *Sender) handleFailure(batch *Batch, err error) {
	level.Info(sender.logger).Log("msg", "sendToServer failed", "error", err)

	if sender.IsShutDown() {
		sender.addErrorMessageToBatchAttempt(batch, err, false)
		sender.FailedCallback(batch)

		return
	}

	if sdkError, ok := err.(*Error); ok {
		if _, ok := sender.noRetryStatusCodeMap[int(sdkError.HTTPCode)]; ok {
			sender.addErrorMessageToBatchAttempt(batch, err, false)
			sender.FailedCallback(batch)

			return
		}
	}

	if batch.attemptCount < batch.maxRetryTimes {
		level.Debug(sender.logger).Log("msg", "Submit to the retry queue after meeting the retry criteria")

		sender.addErrorMessageToBatchAttempt(batch, err, true)
		retryWaitTime := batch.baseRetryBackoffMs * int64(math.Pow(2, float64(batch.attemptCount)-1))

		batch.nextRetryMs = GetTimeMs(time.Now().UnixNano())
		if retryWaitTime < batch.maxRetryIntervalInMs {
			batch.nextRetryMs += retryWaitTime
		} else {
			batch.nextRetryMs += batch.maxRetryIntervalInMs
		}

		sender.retryQueue.addToRetryQueue(batch, sender.logger)

		return
	}

	sender.addErrorMessageToBatchAttempt(batch, err, false)
	sender.FailedCallback(batch)
}

func (sender *Sender) FailedCallback(batch *Batch) {
	level.Info(sender.logger).Log("msg", "sendToServer failed,Execute failed callback function")
	defer atomic.AddInt64(&sender.producer.producerLogGroupSize, -batch.totalDataSize)

	for _, callBack := range batch.callBackList {
		callBack.Fail(batch.result)
	}
}

func (sender *Sender) addErrorMessageToBatchAttempt(producerBatch *Batch, err error, retryInfo bool) {
	if producerBatch.attemptCount < producerBatch.maxReservedAttempts {
		tlsError, ok := err.(*Error)
		var attempt *Attempt
		if ok {
			if retryInfo {
				level.Info(sender.logger).Log("msg", "sendToServer failed,start retrying", "retry times", producerBatch.attemptCount, "requestId", tlsError.RequestID, "error code", tlsError.Code, "error message", tlsError.Message)
			}
			attempt = newAttempt(false, tlsError.RequestID, tlsError.Code, tlsError.Message, GetTimeMs(time.Now().UnixNano()))
		} else {
			level.Error(sender.logger).Log("msg", "putLogs internal err", err.Error())
			attempt = newAttempt(false, "", "", err.Error(), GetTimeMs(time.Now().UnixNano()))
		}
		producerBatch.result.Attempts = append(producerBatch.result.Attempts, attempt)
	}

	producerBatch.result.SuccessFlag = false
	producerBatch.attemptCount += 1
}
