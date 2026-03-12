package producer

import (
	"strings"
	"sync"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestProducerDispatchDoesNotExceedMaxBatchSize(t *testing.T) {
	if MaxBatchSize != 9*1024*1024+512*1024 {
		t.Fatalf("unexpected MaxBatchSize: %d", MaxBatchSize)
	}

	makeLog := func(payloadBytes int) *pb.Log {
		return &pb.Log{
			Time: 1,
			Contents: []*pb.LogContent{
				{
					Key:   "k",
					Value: strings.Repeat("a", payloadBytes),
				},
			},
		}
	}

	producerConfig := GetDefaultProducerConfig()
	producerConfig.MaxBatchSize = MaxBatchSize
	producerConfig.MaxBatchCount = MaxBatchCount

	p := &producer{
		closeCh: make(chan struct{}),
		config:  producerConfig,
		logger:  log.NewNopLogger(),
	}

	threadPool := &ThreadPool{
		taskChan: make(chan *Batch, 1000),
	}

	sender := &Sender{
		retryQueue: newRetryQueue(),
	}

	dispatcher := initDispatcher(producerConfig, sender, log.NewNopLogger(), threadPool, p)

	key := BatchKey{
		Topic:       "t",
		Source:      "s",
		ShardHash:   "h",
		FileName:    "f",
		ContextFlow: "",
	}

	const goroutines = 8
	const logsPerGoroutine = 4
	payloadBytes := 2 * 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < logsPerGoroutine; j++ {
				dispatcher.handleLogs(&BatchLog{Key: key, Log: makeLog(payloadBytes)})
			}
		}()
	}
	wg.Wait()

	dispatcher.lock.Lock()
	for k, batch := range dispatcher.logGroupData {
		threadPool.taskChan <- batch
		delete(dispatcher.logGroupData, k)
	}
	dispatcher.lock.Unlock()

	var batches []*Batch
	for {
		select {
		case b := <-threadPool.taskChan:
			if b != nil {
				batches = append(batches, b)
			}
		default:
			goto Done
		}
	}
Done:

	if len(batches) < 2 {
		t.Fatalf("expected batches to be split, got %d", len(batches))
	}

	for _, b := range batches {
		bodySize := b.logGroupList.Size()
		if int64(bodySize) > MaxBatchSize {
			t.Fatalf("batch exceeds MaxBatchSize: %d > %d", bodySize, MaxBatchSize)
		}
		if b.getLogCount() > 32768 {
			t.Fatalf("batch exceeds server max count 32768: %d", b.getLogCount())
		}
		if b.totalDataSize != int64(bodySize) {
			t.Fatalf("unexpected batch totalDataSize: %d != %d", b.totalDataSize, bodySize)
		}
		if b.getLogCount() > MaxBatchCount {
			t.Fatalf("batch exceeds MaxBatchCount: %d > %d", b.getLogCount(), MaxBatchCount)
		}
	}
}

func TestProducerDispatchSplitsLogGroupBy10000InOneRequest(t *testing.T) {
	producerConfig := GetDefaultProducerConfig()
	producerConfig.MaxBatchSize = MaxBatchSize
	producerConfig.MaxBatchCount = MaxBatchCount

	p := &producer{
		closeCh: make(chan struct{}),
		config:  producerConfig,
		logger:  log.NewNopLogger(),
	}

	threadPool := &ThreadPool{
		taskChan: make(chan *Batch, 100),
	}

	sender := &Sender{
		retryQueue: newRetryQueue(),
	}

	dispatcher := initDispatcher(producerConfig, sender, log.NewNopLogger(), threadPool, p)

	key := BatchKey{
		Topic:       "t",
		Source:      "s",
		ShardHash:   "h",
		FileName:    "f",
		ContextFlow: "",
	}

	for i := 0; i < 25000; i++ {
		dispatcher.handleLogs(&BatchLog{Key: key, Log: &pb.Log{Time: 1}})
	}

	dispatcher.lock.Lock()
	var batch *Batch
	for _, b := range dispatcher.logGroupData {
		batch = b
		break
	}
	dispatcher.lock.Unlock()

	if batch == nil {
		t.Fatalf("expected one pending batch")
	}
	if len(batch.logGroupList.GetLogGroups()) != 3 {
		t.Fatalf("expected 3 log groups, got %d", len(batch.logGroupList.GetLogGroups()))
	}
	if batch.getLogCount() != 25000 {
		t.Fatalf("expected log count 25000, got %d", batch.getLogCount())
	}
	if batch.logGroupList.GetLogGroups()[0].GetLogs() == nil || len(batch.logGroupList.GetLogGroups()[0].GetLogs()) != 10000 {
		t.Fatalf("unexpected first log group size: %d", len(batch.logGroupList.GetLogGroups()[0].GetLogs()))
	}
	if len(batch.logGroupList.GetLogGroups()[1].GetLogs()) != 10000 {
		t.Fatalf("unexpected second log group size: %d", len(batch.logGroupList.GetLogGroups()[1].GetLogs()))
	}
	if len(batch.logGroupList.GetLogGroups()[2].GetLogs()) != 5000 {
		t.Fatalf("unexpected third log group size: %d", len(batch.logGroupList.GetLogGroups()[2].GetLogs()))
	}
	if batch.logGroupList.Size() > int(MaxBatchSize) {
		t.Fatalf("unexpected batch size: %d", batch.logGroupList.Size())
	}
}

func TestProducerRejectsSingleLogOverMaxBatchSize(t *testing.T) {
	makeLog := func(payloadBytes int) *pb.Log {
		return &pb.Log{
			Time: 1,
			Contents: []*pb.LogContent{
				{
					Key:   "k",
					Value: strings.Repeat("a", payloadBytes),
				},
			},
		}
	}

	producerConfig := GetDefaultProducerConfig()
	producerConfig.TotalSizeLnBytes = 1 << 62

	p := &producer{
		closeCh: make(chan struct{}),
		config:  producerConfig,
		logger:  log.NewNopLogger(),
	}
	threadPool := &ThreadPool{
		taskChan: make(chan *Batch, 10),
	}
	sender := &Sender{
		retryQueue: newRetryQueue(),
	}
	p.dispatcher = initDispatcher(producerConfig, sender, log.NewNopLogger(), threadPool, p)

	logSize := int(MaxBatchSize) + 1024
	err := p.putToDispatcher(&BatchLog{Key: BatchKey{Topic: "t"}, Log: makeLog(logSize)})
	if err == nil {
		t.Fatalf("expected error for log larger than MaxBatchSize")
	}
}
