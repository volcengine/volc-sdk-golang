package producer

import (
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	tls "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestProducerMemoryLimitAppliesBeforeDispatcher(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = int64(8 * 1024)
	cfg.MaxProducerMemoryBytes = int64(8 * 1024)
	cfg.MaxBlockSec = 0
	cfg.BatchQueueSize = 1000000

	p := newProducer(cfg)
	successes := 0
	var err error
	for i := 0; i < 100; i++ {
		err = p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(1024), nil)
		if err != nil {
			break
		}
		successes++
	}

	if err == nil {
		t.Fatalf("expected memory limiter to reject input before dispatcher drains queue, successes=%d", successes)
	}
	if !strings.Contains(err.Error(), TimeoutException) {
		t.Fatalf("expected timeout error, got %v", err)
	}
	if successes >= 100 {
		t.Fatalf("expected limiter to stop accepted logs, got %d successes", successes)
	}
	if got := p.memoryLimiter.Used(); got > cfg.MaxProducerMemoryBytes {
		t.Fatalf("memory limiter exceeded cap: got %d want <= %d", got, cfg.MaxProducerMemoryBytes)
	}
}

func TestEstimateBatchLogReservationBytesMatchesProtobufSize(t *testing.T) {
	batchLog := &BatchLog{
		Key: BatchKey{
			Source:      "source",
			FileName:    "file",
			ContextFlow: "context",
		},
		Log: makeMemoryBudgetLog(128),
	}
	group := &pb.LogGroup{
		Logs:        []*pb.Log{batchLog.Log},
		Source:      batchLog.Key.Source,
		FileName:    batchLog.Key.FileName,
		ContextFlow: batchLog.Key.ContextFlow,
	}
	want := int64((&pb.LogGroupList{LogGroups: []*pb.LogGroup{group}}).Size())
	if got := estimateBatchLogReservationBytes(batchLog, int64(batchLog.Log.Size())); got != want {
		t.Fatalf("reservation size mismatch: got %d want %d", got, want)
	}
}

func TestPutToDispatcherUsesFinalLogSizeForReservation(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 1 << 62
	cfg.MaxProducerMemoryBytes = 1 << 62
	cfg.BatchQueueSize = 1

	p := newProducer(cfg)
	defer p.ForceClose()

	batchLog := &BatchLog{
		Key: BatchKey{
			Topic:       "topic",
			Source:      "source",
			FileName:    "file",
			ContextFlow: "context",
		},
		Log: makeMemoryBudgetLog(128),
	}
	batchLog.Log.Time = 0

	if err := p.putToDispatcher(batchLog); err != nil {
		t.Fatalf("putToDispatcher failed: %v", err)
	}

	group := &pb.LogGroup{
		Logs:        []*pb.Log{batchLog.Log},
		Source:      batchLog.Key.Source,
		FileName:    batchLog.Key.FileName,
		ContextFlow: batchLog.Key.ContextFlow,
	}
	want := int64((&pb.LogGroupList{LogGroups: []*pb.LogGroup{group}}).Size())
	if got := batchLog.reservedMemoryBytes(); got != want {
		t.Fatalf("reserved bytes mismatch: got %d want %d", got, want)
	}
}

func TestDefaultProducerMemoryBudgetKeepsPayloadHeadroom(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	if cfg.MaxProducerMemoryBytes != 2*cfg.TotalSizeLnBytes {
		t.Fatalf("expected default lifecycle budget to be twice TotalSizeLnBytes, got %d total %d", cfg.MaxProducerMemoryBytes, cfg.TotalSizeLnBytes)
	}
	if payloadLimit := derivePayloadMemoryLimit(cfg); payloadLimit < cfg.TotalSizeLnBytes {
		t.Fatalf("expected default payload budget to keep TotalSizeLnBytes headroom, payloadLimit=%d total=%d", payloadLimit, cfg.TotalSizeLnBytes)
	}

	cfg = &Config{TotalSizeLnBytes: 64 * 1024 * 1024}
	cfg = validateProducerConfig(cfg)
	if cfg.MaxProducerMemoryBytes != 2*cfg.TotalSizeLnBytes {
		t.Fatalf("expected zero lifecycle budget to default to twice TotalSizeLnBytes, got %d total %d", cfg.MaxProducerMemoryBytes, cfg.TotalSizeLnBytes)
	}

	cfg = GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 150 * 1024 * 1024
	cfg = validateProducerConfig(cfg)
	if cfg.MaxProducerMemoryBytes != 200*1024*1024 {
		t.Fatalf("expected explicit default-sized lifecycle budget to be preserved when large enough, got %d total %d", cfg.MaxProducerMemoryBytes, cfg.TotalSizeLnBytes)
	}

	cfg = GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 1 << 40
	cfg = validateProducerConfig(cfg)
	if cfg.MaxProducerMemoryBytes != 2*cfg.TotalSizeLnBytes {
		t.Fatalf("expected lifecycle budget to follow increased TotalSizeLnBytes, got %d total %d", cfg.MaxProducerMemoryBytes, cfg.TotalSizeLnBytes)
	}

	cfg = GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 10 * 1024 * 1024
	cfg = validateProducerConfig(cfg)
	if cfg.MaxProducerMemoryBytes != 200*1024*1024 {
		t.Fatalf("expected decreased TotalSizeLnBytes to retain default lifecycle budget, got %d total %d", cfg.MaxProducerMemoryBytes, cfg.TotalSizeLnBytes)
	}

	cfg = GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 10 * 1024 * 1024
	cfg.MaxProducerMemoryBytes = 0
	cfg = validateProducerConfig(cfg)
	if cfg.MaxProducerMemoryBytes != 2*cfg.TotalSizeLnBytes {
		t.Fatalf("expected zero lifecycle budget to derive from decreased TotalSizeLnBytes, got %d total %d", cfg.MaxProducerMemoryBytes, cfg.TotalSizeLnBytes)
	}
}

func TestProducerMemoryLimitReleasesOnSuccess(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(1024 * 1024)
	p := &producer{
		config:        cfg,
		memoryLimiter: newMemoryLimiter(cfg.MaxProducerMemoryBytes, cfg.MaxProducerMemoryBytes),
	}
	if err := p.memoryLimiter.acquirePayload(128, 0); err != nil {
		t.Fatalf("reserve payload bytes: %v", err)
	}

	batch := &Batch{
		totalDataSize:       128,
		reservedBytes:       128,
		result:              newResult(),
		maxReservedAttempts: 1,
	}
	sender := initSender(&observingPutLogsClient{}, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)

	sender.handleSuccess(batch, &tls.CommonResponse{})

	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected memory to be released on success, got %d", got)
	}
}

func TestProducerMemoryLimitRetainsOnRetryAndReleasesOnFinalFailure(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(1024 * 1024)
	p := &producer{
		config:        cfg,
		memoryLimiter: newMemoryLimiter(cfg.MaxProducerMemoryBytes, cfg.MaxProducerMemoryBytes),
	}
	if err := p.memoryLimiter.acquirePayload(256, 0); err != nil {
		t.Fatalf("reserve payload bytes: %v", err)
	}

	batch := &Batch{
		totalDataSize:              256,
		reservedBytes:              256,
		result:                     newResult(),
		maxRetryTimes:              2,
		baseRetryBackoffMs:         cfg.BaseRetryBackoffMs,
		baseIncreaseRetryBackoffMs: 1000,
		maxRetryIntervalInMs:       cfg.MaxRetryBackoffMs,
		maxReservedAttempts:        2,
	}
	sender := initSender(&observingPutLogsClient{}, newRetryQueue(), 2, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)

	sender.handleFailure(batch, &tls.Error{HTTPCode: 429})
	if got := p.memoryLimiter.Used(); got != 256 {
		t.Fatalf("expected retryable batch memory to stay reserved, got %d", got)
	}
	if retry := sender.retryQueue.getRetryBatch(true); len(retry) != 1 || retry[0] != batch {
		t.Fatalf("expected batch in retry queue, got %d", len(retry))
	}

	sender.handleFailure(batch, &tls.Error{HTTPCode: 429})
	if got := p.memoryLimiter.Used(); got != 256 {
		t.Fatalf("expected second retry to keep memory reserved, got %d", got)
	}
	if retry := sender.retryQueue.getRetryBatch(true); len(retry) != 1 || retry[0] != batch {
		t.Fatalf("expected batch in retry queue on second retry, got %d", len(retry))
	}

	sender.handleFailure(batch, &tls.Error{HTTPCode: 429})
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected final failure to release memory, got %d", got)
	}
}

func TestProducerTemporaryMemoryAcquireFailsFastWhenBudgetUnavailable(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(4096)
	cfg.TotalSizeLnBytes = int64(4096)
	cfg.MaxBlockSec = -1
	cfg.MaxSenderCount = 1
	cfg.MaxBatchSize = 512
	p := newProducer(cfg)

	if err := p.memoryLimiter.acquirePayload(3000, 0); err != nil {
		t.Fatalf("reserve payload bytes: %v", err)
	}

	err := p.acquireTemporaryBytes(2000)
	if err == nil || !strings.Contains(err.Error(), TimeoutException) {
		t.Fatalf("expected temporary memory acquire to fail fast, got %v", err)
	}
	if got := p.memoryLimiter.Used(); got != 3000 {
		t.Fatalf("expected failed temporary acquire not to change used bytes, got %d", got)
	}
}

func TestProducerPayloadMemoryAcquireUnblocksOnForceClose(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(4096)
	cfg.TotalSizeLnBytes = int64(4096)
	cfg.MaxBlockSec = -1
	cfg.MaxSenderCount = 1
	cfg.MaxBatchSize = 512
	p := newProducer(cfg)

	if err := p.memoryLimiter.acquirePayload(3000, 0); err != nil {
		t.Fatalf("reserve payload bytes: %v", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(512), nil)
	}()

	time.Sleep(20 * time.Millisecond)
	p.ForceClose()

	select {
	case err := <-done:
		if err == nil || !strings.Contains(err.Error(), TimeoutException) {
			t.Fatalf("expected timeout after force close, got %v", err)
		}
	case <-time.After(300 * time.Millisecond):
		t.Fatalf("expected force close to unblock payload memory acquire")
	}
}

func TestMemoryLimiterNoWaitHonorsClosedDone(t *testing.T) {
	limiter := newMemoryLimiter(1024, 1024)
	done := make(chan struct{})
	close(done)

	err := limiter.acquireTemporaryWithCancel(128, 0, done)
	if err == nil || !strings.Contains(err.Error(), TimeoutException) {
		t.Fatalf("expected timeout for closed done, got %v", err)
	}
	if got := limiter.Used(); got != 0 {
		t.Fatalf("expected closed done to avoid reservation, got used=%d", got)
	}
}

func TestMemoryLimiterAcquireWakesOnRelease(t *testing.T) {
	limiter := newMemoryLimiter(1024, 1024)
	if err := limiter.acquireTemporary(1024, 0); err != nil {
		t.Fatalf("reserve memory: %v", err)
	}

	done := make(chan time.Duration, 1)
	go func() {
		start := time.Now()
		if err := limiter.acquireTemporary(1, time.Second); err != nil {
			t.Errorf("acquire after release: %v", err)
		}
		done <- time.Since(start)
	}()

	time.Sleep(10 * time.Millisecond)
	limiter.release(1024)

	select {
	case elapsed := <-done:
		if elapsed >= memoryAcquirePollInterval {
			t.Fatalf("expected release to wake waiter before poll interval, elapsed=%s", elapsed)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("expected release to wake waiting acquire")
	}
}

func TestMemoryLimiterReleaseWithoutWaitersDoesNotRotateNotifyChannel(t *testing.T) {
	limiter := newMemoryLimiter(1024, 1024)
	if err := limiter.acquireTemporary(512, 0); err != nil {
		t.Fatalf("reserve memory: %v", err)
	}

	limiter.notifier.mu.Lock()
	notifyCh := limiter.notifier.ch
	limiter.notifier.mu.Unlock()

	limiter.release(512)

	limiter.notifier.mu.Lock()
	defer limiter.notifier.mu.Unlock()
	if limiter.notifier.ch != notifyCh {
		t.Fatalf("expected release without waiters not to rotate notify channel")
	}
}

func TestMemoryLimiterAcquireFastPathDoesNotNeedNotifyLock(t *testing.T) {
	limiter := newMemoryLimiter(1024, 1024)
	limiter.notifier.mu.Lock()
	defer limiter.notifier.mu.Unlock()

	done := make(chan error, 1)
	go func() {
		done <- limiter.acquireTemporary(128, time.Second)
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("expected fast-path acquire to succeed, got %v", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("expected fast-path acquire not to wait on notify lock")
	}
}

func TestProducerWaitTimeFastPathDoesNotNeedNotifyLock(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 1024
	cfg.MaxProducerMemoryBytes = 4096
	cfg.MaxBlockSec = 5
	p := newProducer(cfg)
	p.sizeNotifier.mu.Lock()
	defer p.sizeNotifier.mu.Unlock()

	done := make(chan error, 1)
	go func() {
		done <- p.waitTime()
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("expected waitTime fast path to succeed, got %v", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("expected waitTime fast path not to wait on notify lock")
	}
}

func TestProducerWaitTimeWakesWhenProducerSizeDrops(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = int64(128)
	cfg.MaxProducerMemoryBytes = int64(4096)
	cfg.MaxBlockSec = 5
	p := newProducer(cfg)
	batch := &Batch{
		totalDataSize:       cfg.TotalSizeLnBytes + 1,
		result:              newResult(),
		maxReservedAttempts: 1,
	}
	atomic.StoreInt64(&p.producerLogGroupSize, batch.totalDataSize)

	done := make(chan time.Duration, 1)
	go func() {
		start := time.Now()
		if err := p.waitTime(); err != nil {
			t.Errorf("waitTime returned error: %v", err)
		}
		done <- time.Since(start)
	}()

	time.Sleep(10 * time.Millisecond)
	sender := initSender(&countingPutLogsClient{}, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)
	sender.handleSuccess(batch, &tls.CommonResponse{})

	select {
	case elapsed := <-done:
		if elapsed >= memoryAcquirePollInterval {
			t.Fatalf("expected producer size release to wake waitTime before poll interval, elapsed=%s", elapsed)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("expected producer size release to wake waitTime")
	}
}

func TestProducerWaitTimeUnblocksOnForceClose(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = int64(1024)
	cfg.MaxProducerMemoryBytes = int64(4096)
	cfg.MaxBlockSec = -1
	p := newProducer(cfg)
	atomic.StoreInt64(&p.producerLogGroupSize, cfg.TotalSizeLnBytes+1)

	done := make(chan error, 1)
	go func() {
		done <- p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(64), nil)
	}()

	time.Sleep(20 * time.Millisecond)
	p.ForceClose()

	select {
	case err := <-done:
		if err == nil || !strings.Contains(err.Error(), "closed") {
			t.Fatalf("expected close error after force close, got %v", err)
		}
	case <-time.After(300 * time.Millisecond):
		t.Fatalf("expected force close to unblock waitTime")
	}
}

func TestProducerForceCloseConcurrentSendDoesNotLeakMemory(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(64 * 1024)
	cfg.TotalSizeLnBytes = int64(64 * 1024)
	cfg.MaxBlockSec = 0
	cfg.BatchQueueSize = 100
	p := newProducer(cfg)

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(256), nil)
		}()
	}

	time.Sleep(10 * time.Millisecond)
	p.ForceClose()
	wg.Wait()

	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected force close concurrent sends to release all memory, got %d", got)
	}
}

func TestProducerMemoryLimitCountsSendBuffers(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(1024 * 1024)
	p := &producer{
		config:        cfg,
		memoryLimiter: newMemoryLimiter(cfg.MaxProducerMemoryBytes, cfg.MaxProducerMemoryBytes),
	}

	logGroupList := &pb.LogGroupList{
		LogGroups: []*pb.LogGroup{
			{Logs: []*pb.Log{makeMemoryBudgetLog(64)}},
		},
	}
	payloadBytes := int64(logGroupList.Size())
	if err := p.memoryLimiter.acquirePayload(payloadBytes, 0); err != nil {
		t.Fatalf("reserve payload bytes: %v", err)
	}

	client := &observingPutLogsClient{t: t, producer: p, minUsed: payloadBytes}
	batch := &Batch{
		logGroupList:        logGroupList,
		logCount:            1,
		totalDataSize:       payloadBytes,
		reservedBytes:       payloadBytes,
		result:              newResult(),
		maxReservedAttempts: 1,
	}
	sender := initSender(client, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)

	sender.sendToServer(batch)

	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected payload and send buffers to be released, got %d", got)
	}
}

func TestProducerMemoryLimitReleasesOnDispatcherSendPanic(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(1024 * 1024)
	cfg.MaxBlockSec = 0
	p := newProducer(cfg)
	close(p.dispatcher.newLogRecvChan)

	err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(128), nil)
	if err == nil {
		t.Fatalf("expected send to closed dispatcher channel to fail")
	}
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected reservation to be released after dispatcher panic, got %d", got)
	}
}

func TestProducerMemoryLimitForceCloseReleasesPendingQueues(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(1024 * 1024)
	p := newProducer(cfg)

	pendingLog := &BatchLog{Log: makeMemoryBudgetLog(64)}
	reserveBatchLogForTest(t, p, pendingLog, 64)
	p.dispatcher.newLogRecvChan <- pendingLog

	p.dispatcher.logGroupData["map"] = reserveBatchForTest(t, p, 128)
	p.dispatcher.retryQueue.addToRetryQueue(reserveBatchForTest(t, p, 256), common.LogConfig(cfg.LoggerConfig))
	p.threadPool.taskChan <- reserveBatchForTest(t, p, 512)
	atomic.StoreInt64(&p.producerLogGroupSize, 128+256+512)

	if got := p.memoryLimiter.Used(); got == 0 {
		t.Fatalf("expected reserved memory before force close")
	}

	p.ForceClose()

	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected force close to release pending memory, got %d", got)
	}
	if got := p.currentProducerLogGroupSize(); got != 0 {
		t.Fatalf("expected force close to release pending producer size, got %d", got)
	}
}

func TestReleaseBatchResourcesIsIdempotentForProducerSize(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(4096)
	cfg.TotalSizeLnBytes = int64(4096)
	p := newProducer(cfg)
	batch := reserveBatchForTest(t, p, 512)
	atomic.StoreInt64(&p.producerLogGroupSize, batch.totalDataSize)

	p.releaseBatchResources(batch)
	p.releaseBatchResources(batch)

	if got := p.currentProducerLogGroupSize(); got != 0 {
		t.Fatalf("expected producer size to remain released after duplicate release, got %d", got)
	}
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected memory to remain released after duplicate release, got %d", got)
	}
}

func TestProducerMemoryLimitUnderContinuous429(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = int64(64 * 1024)
	cfg.TotalSizeLnBytes = int64(64 * 1024)
	cfg.MaxBlockSec = 0
	cfg.MaxSenderCount = 2
	cfg.BatchQueueSize = 1000
	cfg.MaxBatchSize = 512
	cfg.Retries = 100
	cfg.LingerTime = 100 * time.Millisecond

	p := newProducer(cfg)
	client := &retryable429Client{}
	p.cli = client
	p.threadPool.sender.client = client
	p.Start()

	accepted := 0
	rejected := 0
	for i := 0; i < 1000; i++ {
		err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(256), nil)
		if err != nil {
			rejected++
			continue
		}
		accepted++
		if got := p.memoryLimiter.Used(); got > cfg.MaxProducerMemoryBytes {
			t.Fatalf("memory exceeded during continuous 429: got %d want <= %d", got, cfg.MaxProducerMemoryBytes)
		}
	}

	if accepted == 0 || rejected == 0 {
		t.Fatalf("expected both accepted and rejected logs under memory pressure, accepted=%d rejected=%d", accepted, rejected)
	}
	if got := p.memoryLimiter.Used(); got > cfg.MaxProducerMemoryBytes {
		t.Fatalf("memory exceeded cap before force close: got %d want <= %d", got, cfg.MaxProducerMemoryBytes)
	}

	p.ForceClose()

	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected memory to be released after continuous 429 force close, got %d", got)
	}
}

func makeMemoryBudgetLog(payloadSize int) *pb.Log {
	return &pb.Log{
		Time: 1,
		Contents: []*pb.LogContent{
			{Key: "message", Value: strings.Repeat("x", payloadSize)},
		},
	}
}

func reserveBatchLogForTest(t *testing.T, p *producer, batchLog *BatchLog, bytes int64) {
	t.Helper()
	if err := p.memoryLimiter.acquirePayload(bytes, 0); err != nil {
		t.Fatalf("reserve batch log: %v", err)
	}
	batchLog.setReservedBytes(bytes)
}

func reserveBatchForTest(t *testing.T, p *producer, bytes int64) *Batch {
	t.Helper()
	if err := p.memoryLimiter.acquirePayload(bytes, 0); err != nil {
		t.Fatalf("reserve batch: %v", err)
	}
	return &Batch{
		totalDataSize:       bytes,
		reservedBytes:       bytes,
		result:              newResult(),
		maxReservedAttempts: 1,
	}
}

type observingPutLogsClient struct {
	t        *testing.T
	producer *producer
	minUsed  int64
	tls.Client
}

func (c *observingPutLogsClient) PutLogs(req *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	if c.t != nil {
		if got := c.producer.memoryLimiter.Used(); got <= c.minUsed {
			c.t.Fatalf("expected send buffer bytes to be reserved during PutLogs, got used=%d payload=%d", got, c.minUsed)
		}
	}
	return &tls.CommonResponse{}, nil
}

type retryable429Client struct {
	tls.Client
}

func (c *retryable429Client) PutLogs(req *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	return nil, &tls.Error{HTTPCode: 429}
}
