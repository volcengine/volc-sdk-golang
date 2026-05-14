package producer

import (
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	tls "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestProducerManagedRetryConfiguresClientAsSingleAttempt(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.RetryMode = RetryModeProducerManaged
	cfg.RequestTimeout = 1500 * time.Millisecond

	p := newProducer(cfg)

	if got := p.cli.GetRetryPolicy().MaxAttempts; got != 1 {
		t.Fatalf("expected producer-managed mode to disable client retries, got MaxAttempts=%d", got)
	}
	if got := p.cli.GetHttpClient().Timeout; got != cfg.RequestTimeout {
		t.Fatalf("expected request timeout to be applied to client, got %v want %v", got, cfg.RequestTimeout)
	}
}

func TestProducerClientRetryPolicyConfiguresClientRetryPolicy(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	policy := tls.DefaultRetryPolicy()
	policy.TotalTimeout = 2 * time.Minute
	policy.InitialInterval = 250 * time.Millisecond
	policy.MaxInterval = 3 * time.Second
	policy.Multiplier = 1.5
	policy.RandomizationFactor = 0.2
	policy.MaxAttempts = 4
	cfg.ClientRetryPolicy = &policy

	p := newProducer(cfg)

	if got := p.cli.GetRetryPolicy(); got != policy.Normalize() {
		t.Fatalf("expected configured client retry policy %#v, got %#v", policy.Normalize(), got)
	}
}

func TestProducerClientRetryPolicyAppliesToPutLogs(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://example.com"
	cfg.Region = "cn-north-1"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.Retries = 0
	cfg.MaxBatchCount = 1
	cfg.LingerTime = 101 * time.Millisecond
	cfg.MaxBlockSec = 0
	logger := log.NewNopLogger()
	cfg.Logger = &logger
	policy := tls.DefaultRetryPolicy()
	policy.TotalTimeout = 30 * time.Second
	policy.InitialInterval = 100 * time.Millisecond
	policy.MaxInterval = 100 * time.Millisecond
	policy.Multiplier = 1
	policy.RandomizationFactor = 0.1
	policy.MaxAttempts = 2
	cfg.ClientRetryPolicy = &policy

	p := newProducer(cfg)
	transport := &retryThenSuccessRoundTripper{}
	if err := p.cli.SetHttpClient(&http.Client{Transport: transport, Timeout: time.Second}); err != nil {
		t.Fatalf("SetHttpClient failed: %v", err)
	}
	p.Start()
	defer p.ForceClose()

	callback := &waitCallback{done: make(chan struct{})}
	if err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(16), callback); err != nil {
		t.Fatalf("SendLog failed: %v", err)
	}

	select {
	case <-callback.done:
	case <-time.After(3 * time.Second):
		t.Fatalf("timed out waiting for producer callback; requests=%d", atomic.LoadInt32(&transport.calls))
	}
	if got := atomic.LoadInt32(&transport.calls); got != 2 {
		t.Fatalf("expected client retry policy to produce 2 HTTP requests, got %d", got)
	}
	if got := atomic.LoadInt32(&callback.successes); got != 1 {
		t.Fatalf("expected final success callback once, got %d", got)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 0 {
		t.Fatalf("expected no final failure callback, got %d", got)
	}
}

func TestProducerManagedRetryCapsConfiguredClientRetryAttempts(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.RetryMode = RetryModeProducerManaged
	policy := tls.DefaultRetryPolicy()
	policy.InitialInterval = 250 * time.Millisecond
	policy.MaxAttempts = 4
	cfg.ClientRetryPolicy = &policy

	p := newProducer(cfg)
	got := p.cli.GetRetryPolicy()

	if got.MaxAttempts != 1 {
		t.Fatalf("expected producer-managed mode to keep one client attempt, got %d", got.MaxAttempts)
	}
	if got.InitialInterval != policy.InitialInterval {
		t.Fatalf("expected producer-managed mode to preserve configured client retry policy fields, got %v want %v", got.InitialInterval, policy.InitialInterval)
	}
}

func TestProducerManagedRetryAppliesOneClientAttemptPerProducerAttempt(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://example.com"
	cfg.Region = "cn-north-1"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.RetryMode = RetryModeProducerManaged
	cfg.Retries = 2
	cfg.BaseRetryBackoffMs = 1
	cfg.MaxRetryBackoffMs = 1
	cfg.MaxBatchCount = 1
	cfg.LingerTime = 101 * time.Millisecond
	cfg.MaxBlockSec = 0
	logger := log.NewNopLogger()
	cfg.Logger = &logger

	p := newProducer(cfg)
	transport := &always429RoundTripper{}
	if err := p.cli.SetHttpClient(&http.Client{Transport: transport, Timeout: p.config.RequestTimeout}); err != nil {
		t.Fatalf("SetHttpClient failed: %v", err)
	}
	p.Start()
	defer p.ForceClose()

	callback := &waitCallback{done: make(chan struct{})}
	if err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(16), callback); err != nil {
		t.Fatalf("SendLog failed: %v", err)
	}

	select {
	case <-callback.done:
	case <-time.After(3 * time.Second):
		t.Fatalf("timed out waiting for final producer callback; requests=%d", atomic.LoadInt32(&transport.calls))
	}
	if got := atomic.LoadInt32(&transport.calls); got != 3 {
		t.Fatalf("expected 3 HTTP requests for initial attempt plus 2 producer retries, got %d", got)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 1 {
		t.Fatalf("expected final failure callback once, got %d", got)
	}
}

func TestProducerManagedRetryDefaultsRequestTimeout(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.RetryMode = RetryModeProducerManaged

	p := newProducer(cfg)

	if got := p.config.RequestTimeout; got != 10*time.Second {
		t.Fatalf("expected producer-managed default request timeout 10s, got %v", got)
	}
	if got := p.cli.GetHttpClient().Timeout; got != 10*time.Second {
		t.Fatalf("expected client timeout to use producer-managed default 10s, got %v", got)
	}
}

func TestLegacyRetryModeKeepsClientRetryPolicy(t *testing.T) {
	cfg := GetDefaultProducerConfig()

	p := newProducer(cfg)

	if got := p.cli.GetRetryPolicy().MaxAttempts; got != tls.DefaultRetryPolicy().MaxAttempts {
		t.Fatalf("expected legacy mode to keep client retry policy, got MaxAttempts=%d", got)
	}
}

func TestDefaultNoRetryStatusCodes(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	got := make(map[int]struct{}, len(cfg.NoRetryStatusCodeList))
	for _, statusCode := range cfg.NoRetryStatusCodeList {
		got[statusCode] = struct{}{}
	}

	for _, statusCode := range []int{400, 403, 404} {
		if _, ok := got[statusCode]; !ok {
			t.Fatalf("expected default NoRetryStatusCodeList to include %d, got %v", statusCode, cfg.NoRetryStatusCodeList)
		}
	}
}

func TestDefaultNoRetryStatusCodesDoNotOpenCircuit(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.MinRequests = 1
	cfg.CircuitBreaker.FailureRatio = 0.1
	cfg = validateProducerConfig(cfg)

	controller := newFailureController(cfg)
	noRetry := map[int]struct{}{400: {}, 403: {}, 404: {}}
	for _, statusCode := range []int{400, 403, 404} {
		controller.afterSend(classifyFailure(&tls.Error{HTTPCode: int32(statusCode)}, noRetry), false, 0)
	}

	controller.lock.Lock()
	defer controller.lock.Unlock()
	if controller.state != circuitStateClosed {
		t.Fatalf("expected no-retry status codes not to open circuit, got state=%d", controller.state)
	}
	if controller.failures != 0 || controller.consecutiveFailures != 0 {
		t.Fatalf("expected no-retry status codes not to count as circuit failures, failures=%d consecutive=%d", controller.failures, controller.consecutiveFailures)
	}
}

func TestClassifyFailureDoesNotTreatServerTimeoutMessageAsLocalResource(t *testing.T) {
	failure := classifyFailure(&tls.Error{
		HTTPCode: 429,
		Message:  "server message mentions " + TimeoutException,
	}, map[int]struct{}{})

	if failure.Kind != FailureKindThrottled || !failure.Retryable {
		t.Fatalf("expected server 429 to remain retryable throttled failure, got kind=%d retryable=%v", failure.Kind, failure.Retryable)
	}
}

func TestClassifyFailureDoesNotTreatServerCircuitOpenMessageAsLocalCircuitOpen(t *testing.T) {
	failure := classifyFailure(&tls.Error{
		HTTPCode: 429,
		Message:  "server message mentions " + CircuitOpenException,
	}, map[int]struct{}{})

	if failure.Kind != FailureKindThrottled || !failure.Retryable {
		t.Fatalf("expected server 429 to remain retryable throttled failure, got kind=%d retryable=%v", failure.Kind, failure.Retryable)
	}
}

func TestValidateProducerConfigAlignsZeroValueDefaults(t *testing.T) {
	cfg := validateProducerConfig(&Config{})

	if cfg.MaxRetryBackoffMs != GetDefaultProducerConfig().MaxRetryBackoffMs {
		t.Fatalf("expected MaxRetryBackoffMs default %d, got %d", GetDefaultProducerConfig().MaxRetryBackoffMs, cfg.MaxRetryBackoffMs)
	}
	if cfg.MaxBatchSize != GetDefaultProducerConfig().MaxBatchSize {
		t.Fatalf("expected MaxBatchSize default %d, got %d", GetDefaultProducerConfig().MaxBatchSize, cfg.MaxBatchSize)
	}
	if cfg.MaxBatchCount != GetDefaultProducerConfig().MaxBatchCount {
		t.Fatalf("expected MaxBatchCount default %d, got %d", GetDefaultProducerConfig().MaxBatchCount, cfg.MaxBatchCount)
	}
}

func TestSenderRetryClassifierRetriesNetworkErrors(t *testing.T) {
	retryQueue := newRetryQueue()
	p := &producer{}
	s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{}, p)
	b := &Batch{
		totalDataSize:              123,
		maxRetryTimes:              1,
		baseRetryBackoffMs:         1,
		baseIncreaseRetryBackoffMs: 1,
		maxRetryIntervalInMs:       1,
		result:                     newResult(),
		maxReservedAttempts:        2,
	}
	p.producerLogGroupSize = b.totalDataSize

	s.handleFailure(b, io.EOF)

	if retryQueue.Len() != 1 {
		t.Fatalf("expected network EOF to be scheduled for retry, got queue len %d", retryQueue.Len())
	}
	if p.producerLogGroupSize != b.totalDataSize {
		t.Fatalf("expected retryable batch memory to be retained, got %d", p.producerLogGroupSize)
	}
}

func TestSenderRetryClassifierDoesNotRetryLocalResourceFailure(t *testing.T) {
	retryQueue := newRetryQueue()
	p := &producer{}
	s := initSender(nil, retryQueue, 1, log.NewNopLogger(), map[int]struct{}{}, p)
	b := &Batch{
		totalDataSize:              123,
		maxRetryTimes:              1,
		baseRetryBackoffMs:         1,
		baseIncreaseRetryBackoffMs: 1,
		maxRetryIntervalInMs:       1,
		result:                     newResult(),
		maxReservedAttempts:        2,
	}
	p.producerLogGroupSize = b.totalDataSize

	s.handleFailure(b, ErrMemoryTimeout)

	if retryQueue.Len() != 0 {
		t.Fatalf("expected local resource failure not to retry, got queue len %d", retryQueue.Len())
	}
	if p.producerLogGroupSize != 0 {
		t.Fatalf("expected final local failure to release producer size, got %d", p.producerLogGroupSize)
	}
}

func TestCircuitBreakerFailFastRejectsBeforeMemoryReservation(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Hour
	cfg.FailurePolicy = FailurePolicyFailFast
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)

	err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(128), nil)
	if err == nil || !strings.Contains(err.Error(), CircuitOpenException) {
		t.Fatalf("expected circuit open error, got %v", err)
	}
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected fail-fast to reject before memory reservation, got used=%d", got)
	}
}

func TestCircuitBreakerDropWithCallbackRejectsBeforeMemoryReservation(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Hour
	cfg.FailurePolicy = FailurePolicyDropWithCallback
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	callback := &countingCallback{}

	err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(128), callback)
	if err != nil {
		t.Fatalf("expected drop-with-callback to report through callback, got err=%v", err)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 1 {
		t.Fatalf("expected one failure callback, got %d", got)
	}
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected drop-with-callback to reject before memory reservation, got used=%d", got)
	}
}

func TestCircuitBreakerDropWithCallbackSendLogsReportsEachLog(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Hour
	cfg.FailurePolicy = FailurePolicyDropWithCallback
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	callback := &countingCallback{}

	group := &pb.LogGroup{Logs: []*pb.Log{
		makeMemoryBudgetLog(16),
		makeMemoryBudgetLog(16),
		makeMemoryBudgetLog(16),
	}}
	err := p.SendLogs("", "topic", "source", "file", group, callback)
	if err != nil {
		t.Fatalf("expected SendLogs drop-with-callback to return nil, got %v", err)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 3 {
		t.Fatalf("expected one failure callback per log, got %d", got)
	}
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected SendLogs drop-with-callback to avoid reservation, got used=%d", got)
	}
}

func TestCircuitBreakerRetryThenCallbackDoesNotRejectBeforeQueue(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Hour
	cfg.FailurePolicy = FailurePolicyRetryThenCallback
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	defer p.ForceClose()
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)

	err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(16), nil)
	if err != nil {
		t.Fatalf("expected RetryThenCallback to keep compatibility and enqueue, got %v", err)
	}
	if got := p.memoryLimiter.Used(); got == 0 {
		t.Fatalf("expected RetryThenCallback open circuit path to enqueue and reserve memory")
	}
}

func TestCircuitBreakerOpenFailsRetryBatchAndReleasesMemory(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = 1024 * 1024
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Hour
	cfg.FailurePolicy = FailurePolicyFailFast

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	callback := &countingCallback{}
	batch := reserveBatchForTest(t, p, 256)
	batch.logGroupList = makeSingleLogGroupList(64)
	batch.logCount = 1
	batch.callBackList = []CallBack{callback}
	client := &countingPutLogsClient{}
	sender := initSender(client, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)

	sender.sendToServer(batch)

	if got := atomic.LoadInt32(&client.calls); got != 0 {
		t.Fatalf("expected open circuit to avoid client PutLogs, got %d calls", got)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 1 {
		t.Fatalf("expected open circuit to fail callback once, got %d", got)
	}
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected open circuit to release batch memory, got %d", got)
	}
}

func TestCircuitBreakerHalfOpenLimitsIngressBeforeMemoryReservation(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Millisecond
	cfg.CircuitBreaker.HalfOpenMaxRequests = 1
	cfg.FailurePolicy = FailurePolicyFailFast
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)

	if err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(128), nil); err != nil {
		t.Fatalf("expected first half-open probe to be accepted, got %v", err)
	}
	usedAfterFirst := p.memoryLimiter.Used()
	if usedAfterFirst == 0 {
		t.Fatalf("expected first half-open probe to reserve memory")
	}

	err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(128), nil)
	if err == nil || !strings.Contains(err.Error(), CircuitOpenException) {
		t.Fatalf("expected second half-open probe to be rejected, got %v", err)
	}
	if got := p.memoryLimiter.Used(); got != usedAfterFirst {
		t.Fatalf("expected rejected half-open probe not to reserve memory, got %d want %d", got, usedAfterFirst)
	}
}

func TestCircuitBreakerHalfOpenSendLogsFailFastKeepsAvailableProbes(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Millisecond
	cfg.CircuitBreaker.HalfOpenMaxRequests = 1
	cfg.FailurePolicy = FailurePolicyFailFast
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)

	group := &pb.LogGroup{Logs: []*pb.Log{
		makeMemoryBudgetLog(16),
		makeMemoryBudgetLog(16),
	}}
	err := p.SendLogs("", "topic", "source", "file", group, nil)
	if err == nil || !strings.Contains(err.Error(), CircuitOpenException) {
		t.Fatalf("expected SendLogs to return circuit-open error for the excess log, got %v", err)
	}
	if got := p.memoryLimiter.Used(); got == 0 {
		t.Fatalf("expected available half-open probe to reserve memory")
	}
	if got := len(p.dispatcher.newLogRecvChan); got != 1 {
		t.Fatalf("expected exactly one half-open probe to be enqueued, got queue len %d", got)
	}
}

func TestCircuitBreakerHalfOpenSendLogsDropWithCallbackKeepsAvailableProbes(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Millisecond
	cfg.CircuitBreaker.HalfOpenMaxRequests = 1
	cfg.FailurePolicy = FailurePolicyDropWithCallback
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)
	callback := &countingCallback{}

	group := &pb.LogGroup{Logs: []*pb.Log{
		makeMemoryBudgetLog(16),
		makeMemoryBudgetLog(16),
	}}
	err := p.SendLogs("", "topic", "source", "file", group, callback)
	if err != nil {
		t.Fatalf("expected drop-with-callback to return nil, got %v", err)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 1 {
		t.Fatalf("expected only the excess log to fail through callback, got %d", got)
	}
	if got := p.memoryLimiter.Used(); got == 0 {
		t.Fatalf("expected available half-open probe to reserve memory")
	}
	if got := len(p.dispatcher.newLogRecvChan); got != 1 {
		t.Fatalf("expected exactly one half-open probe to be enqueued, got queue len %d", got)
	}
}

func TestCircuitBreakerSendLogsWaitTimeFailsBeforeHalfOpenPermit(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 128
	cfg.MaxProducerMemoryBytes = 4096
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Millisecond
	cfg.CircuitBreaker.HalfOpenMaxRequests = 2
	cfg.FailurePolicy = FailurePolicyFailFast
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)
	p.producerLogGroupSize = cfg.TotalSizeLnBytes + 1

	group := &pb.LogGroup{Logs: []*pb.Log{
		makeMemoryBudgetLog(16),
		makeMemoryBudgetLog(16),
	}}
	err := p.SendLogs("", "topic", "source", "file", group, nil)
	if err == nil || !strings.Contains(err.Error(), TimeoutException) {
		t.Fatalf("expected waitTime timeout, got %v", err)
	}
	p.failureController.lock.Lock()
	defer p.failureController.lock.Unlock()
	if got := p.failureController.halfOpenInFlight; got != 0 {
		t.Fatalf("expected reserved half-open permits to be released after waitTime failure, got %d", got)
	}
}

func TestCircuitBreakerSendLogsReleasesCurrentHalfOpenPermitWhenPutFails(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 1
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Millisecond
	cfg.CircuitBreaker.HalfOpenMaxRequests = 1
	cfg.FailurePolicy = FailurePolicyFailFast
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)

	group := &pb.LogGroup{Logs: []*pb.Log{
		makeMemoryBudgetLog(16),
	}}
	err := p.SendLogs("", "topic", "source", "file", group, nil)
	if err == nil || !strings.Contains(err.Error(), "is larger than the total size") {
		t.Fatalf("expected putToDispatcher size validation error, got %v", err)
	}
	p.failureController.lock.Lock()
	defer p.failureController.lock.Unlock()
	if got := p.failureController.halfOpenInFlight; got != 0 {
		t.Fatalf("expected current half-open permit to be released after putToDispatcher failure, got %d", got)
	}
	if got := p.memoryLimiter.Used(); got != 0 {
		t.Fatalf("expected failed putToDispatcher path not to retain memory, got used=%d", got)
	}
	if got := len(p.dispatcher.newLogRecvChan); got != 0 {
		t.Fatalf("expected failed putToDispatcher path not to enqueue, got queue len %d", got)
	}
}

func TestCircuitBreakerOpensByFailureRatio(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.MinRequests = 4
	cfg.CircuitBreaker.FailureRatio = 0.5
	cfg.CircuitBreaker.ConsecutiveFailures = 100
	cfg = validateProducerConfig(cfg)

	controller := newFailureController(cfg)
	controller.afterSend(classifiedFailure{}, true, 0)
	controller.afterSend(classifiedFailure{}, true, 0)
	controller.afterSend(classifiedFailure{Kind: FailureKindServer, Retryable: true}, false, 0)
	controller.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)

	controller.lock.Lock()
	defer controller.lock.Unlock()
	if controller.state != circuitStateOpen {
		t.Fatalf("expected failure ratio to open circuit, got state=%d requests=%d failures=%d", controller.state, controller.requests, controller.failures)
	}
}

func TestCircuitBreakerHalfOpenSuccessClosesAndFailureReopens(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Millisecond
	cfg.CircuitBreaker.HalfOpenMaxRequests = 2
	cfg = validateProducerConfig(cfg)

	controller := newFailureController(cfg)
	controller.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)
	for i := 0; i < cfg.CircuitBreaker.HalfOpenMaxRequests; i++ {
		allowed, permit := controller.allowEnqueue()
		if !allowed || !permit {
			t.Fatalf("expected half-open probe %d to get permit, allowed=%v permit=%v", i, allowed, permit)
		}
		controller.afterSend(classifiedFailure{}, true, 1)
	}
	controller.lock.Lock()
	if controller.state != circuitStateClosed {
		state := controller.state
		controller.lock.Unlock()
		t.Fatalf("expected half-open successes to close circuit, got state=%d", state)
	}
	controller.lock.Unlock()

	controller.afterSend(classifiedFailure{Kind: FailureKindServer, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)
	allowed, permit := controller.allowEnqueue()
	if !allowed || !permit {
		t.Fatalf("expected half-open failure probe to get permit, allowed=%v permit=%v", allowed, permit)
	}
	controller.afterSend(classifiedFailure{Kind: FailureKindServer, Retryable: true}, false, 1)

	controller.lock.Lock()
	defer controller.lock.Unlock()
	if controller.state != circuitStateOpen {
		t.Fatalf("expected half-open failure to reopen circuit, got state=%d", controller.state)
	}
}

func TestCircuitBreakerHalfOpenCoalescedBatchReleasesAllPermits(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = 1024 * 1024
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.ConsecutiveFailures = 1
	cfg.CircuitBreaker.OpenDuration = time.Millisecond
	cfg.CircuitBreaker.HalfOpenMaxRequests = 5
	cfg.FailurePolicy = FailurePolicyFailFast
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.afterSend(classifiedFailure{Kind: FailureKindThrottled, Retryable: true}, false, 0)
	time.Sleep(2 * time.Millisecond)

	for i := 0; i < cfg.CircuitBreaker.HalfOpenMaxRequests; i++ {
		if err := p.SendLog("", "topic", "source", "file", makeMemoryBudgetLog(64), nil); err != nil {
			t.Fatalf("expected half-open probe %d to be accepted, got %v", i, err)
		}
		batchLog := <-p.dispatcher.newLogRecvChan
		p.dispatcher.handleLogs(batchLog)
	}

	var batch *Batch
	for _, candidate := range p.dispatcher.logGroupData {
		batch = candidate
		break
	}
	if batch == nil {
		t.Fatalf("expected coalesced half-open batch")
	}
	sender := initSender(&countingPutLogsClient{}, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)
	sender.handleSuccess(batch, &tls.CommonResponse{})

	p.failureController.lock.Lock()
	defer p.failureController.lock.Unlock()
	if got := p.failureController.halfOpenInFlight; got != 0 {
		t.Fatalf("expected all coalesced half-open permits released, got %d", got)
	}
	if got := p.failureController.halfOpenSuccesses; got != 1 {
		t.Fatalf("expected coalesced batch to count as one successful probe, got %d", got)
	}
}

func TestCircuitBreakerTemporaryMemoryFailureReleasesHalfOpenPermit(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.TotalSizeLnBytes = 4096
	cfg.MaxProducerMemoryBytes = 4096
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.HalfOpenMaxRequests = 1
	cfg.FailurePolicy = FailurePolicyFailFast
	cfg.MaxBlockSec = 0

	p := newProducer(cfg)
	p.failureController.lock.Lock()
	p.failureController.state = circuitStateHalfOpen
	p.failureController.lock.Unlock()

	callback := &countingCallback{}
	client := &countingPutLogsClient{}
	sender := initSender(client, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)
	logGroupList := makeSingleLogGroupList(4096)
	batch := &Batch{
		totalDataSize:       int64(logGroupList.Size()),
		logGroupList:        logGroupList,
		logCount:            1,
		callBackList:        []CallBack{callback},
		result:              newResult(),
		maxReservedAttempts: 1,
	}
	p.producerLogGroupSize = batch.totalDataSize

	sender.sendToServer(batch)

	if got := atomic.LoadInt32(&client.calls); got != 0 {
		t.Fatalf("expected temporary memory failure to avoid PutLogs, got %d calls", got)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 1 {
		t.Fatalf("expected local resource failure callback once, got %d", got)
	}
	if got := batch.getCircuitPermitCount(); got != 0 {
		t.Fatalf("expected batch-side circuit permit to be cleared, got %d", got)
	}
	p.failureController.lock.Lock()
	defer p.failureController.lock.Unlock()
	if got := p.failureController.halfOpenInFlight; got != 0 {
		t.Fatalf("expected half-open permit to be released, got %d", got)
	}
	if p.failureController.state != circuitStateHalfOpen {
		t.Fatalf("expected local resource failure not to reopen circuit, got state=%d", p.failureController.state)
	}
	if got := p.failureController.failures; got != 0 {
		t.Fatalf("expected local resource failure not to count as circuit failure, got %d", got)
	}
}

func TestSenderReleasesResourcesBeforeSuccessCallback(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = 1024 * 1024
	cfg.TotalSizeLnBytes = 1024 * 1024
	p := newProducer(cfg)
	batch := reserveBatchForTest(t, p, 512)
	batch.logGroupList = makeSingleLogGroupList(64)
	batch.logCount = 1
	callback := &resourceReleaseCallback{t: t, producer: p}
	batch.callBackList = []CallBack{callback}
	p.producerLogGroupSize = batch.totalDataSize
	sender := initSender(&successPutLogsClient{}, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)

	sender.sendToServer(batch)

	if got := atomic.LoadInt32(&callback.successes); got != 1 {
		t.Fatalf("expected success callback once, got %d", got)
	}
}

func TestSenderReleasesResourcesBeforeFailureCallback(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxProducerMemoryBytes = 1024 * 1024
	cfg.TotalSizeLnBytes = 1024 * 1024
	p := newProducer(cfg)
	batch := reserveBatchForTest(t, p, 512)
	batch.logGroupList = makeSingleLogGroupList(64)
	batch.logCount = 1
	callback := &resourceReleaseCallback{t: t, producer: p}
	batch.callBackList = []CallBack{callback}
	p.producerLogGroupSize = batch.totalDataSize
	sender := initSender(&permanentErrorPutLogsClient{}, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)

	sender.sendToServer(batch)

	if got := atomic.LoadInt32(&callback.failures); got != 1 {
		t.Fatalf("expected failure callback once, got %d", got)
	}
}

func TestCircuitBreakerSenderRejectDoesNotReleaseUnheldPermit(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.CircuitBreaker.Enabled = true
	cfg.CircuitBreaker.HalfOpenMaxRequests = 1
	cfg.FailurePolicy = FailurePolicyFailFast

	p := newProducer(cfg)
	p.failureController.lock.Lock()
	p.failureController.state = circuitStateHalfOpen
	p.failureController.halfOpenInFlight = 1
	p.failureController.lock.Unlock()

	callback := &countingCallback{}
	client := &countingPutLogsClient{}
	sender := initSender(client, newRetryQueue(), 1, common.LogConfig(cfg.LoggerConfig), map[int]struct{}{}, p)
	batch := &Batch{
		totalDataSize:       64,
		logGroupList:        makeSingleLogGroupList(64),
		logCount:            1,
		callBackList:        []CallBack{callback},
		result:              newResult(),
		maxReservedAttempts: 1,
	}
	p.producerLogGroupSize = batch.totalDataSize

	sender.sendToServer(batch)

	if got := atomic.LoadInt32(&client.calls); got != 0 {
		t.Fatalf("expected sender-side circuit rejection to avoid PutLogs, got %d calls", got)
	}
	if got := atomic.LoadInt32(&callback.failures); got != 1 {
		t.Fatalf("expected rejected batch to fail callback once, got %d", got)
	}
	p.failureController.lock.Lock()
	defer p.failureController.lock.Unlock()
	if got := p.failureController.halfOpenInFlight; got != 1 {
		t.Fatalf("expected unheld permit not to be released, got %d", got)
	}
}

func TestThreadPoolHandleBatchWithNilSenderDoesNotPanic(t *testing.T) {
	threadPool := &ThreadPool{forceQuitCh: make(chan struct{})}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("expected nil sender to be ignored, got panic: %v", r)
		}
	}()
	threadPool.handleBatch(&Batch{})
}

func TestThreadPoolForceCloseDoesNotBlockOnFullSenderSlots(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.MaxSenderCount = 1
	cfg.MaxProducerMemoryBytes = 1024 * 1024
	p := newProducer(cfg)
	blockClient := &blockingPutLogsClient{
		releaseFirst:  make(chan struct{}),
		firstStarted:  make(chan struct{}),
		secondStarted: make(chan struct{}),
	}
	p.cli = blockClient
	p.threadPool.sender.client = blockClient
	p.Start()

	first := reserveBatchForTest(t, p, 128)
	first.logGroupList = makeSingleLogGroupList(32)
	first.logCount = 1
	p.threadPool.taskChan <- first

	select {
	case <-blockClient.firstStarted:
	case <-time.After(time.Second):
		t.Fatalf("expected first sender to start")
	}

	second := reserveBatchForTest(t, p, 128)
	second.logGroupList = makeSingleLogGroupList(32)
	second.logCount = 1
	p.threadPool.taskChan <- second

	done := make(chan struct{})
	go func() {
		p.ForceClose()
		close(done)
	}()
	time.Sleep(20 * time.Millisecond)
	close(blockClient.releaseFirst)

	select {
	case <-done:
	case <-time.After(300 * time.Millisecond):
		t.Fatalf("expected ForceClose not to block on full sender slots")
	}
	select {
	case <-blockClient.secondStarted:
		t.Fatalf("expected second batch not to start after force close")
	default:
	}
}

func TestDispatcherCheckBatchesUnblocksOnForceQuitWhenTaskChanFull(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.LingerTime = time.Millisecond
	cfg.TotalSizeLnBytes = 1024 * 1024
	cfg.MaxProducerMemoryBytes = 1024 * 1024
	p := newProducer(cfg)
	p.threadPool.taskChan = make(chan *Batch, 1)
	p.threadPool.taskChan <- &Batch{}
	batch := reserveBatchForTest(t, p, 128)
	batch.createTime = time.Now().Add(-p.config.LingerTime - time.Second)
	p.dispatcher.logGroupData["blocked"] = batch

	done := make(chan struct{})
	go func() {
		p.dispatcher.checkBatches(p.config)
		close(done)
	}()

	time.Sleep(20 * time.Millisecond)
	close(p.dispatcher.forceQuitCh)

	select {
	case <-done:
	case <-time.After(300 * time.Millisecond):
		t.Fatalf("expected dispatcher checkBatches to unblock on force quit when taskChan is full")
	}
}

func makeSingleLogGroupList(payloadSize int) *pb.LogGroupList {
	return &pb.LogGroupList{
		LogGroups: []*pb.LogGroup{
			{Logs: []*pb.Log{makeMemoryBudgetLog(payloadSize)}},
		},
	}
}

type countingCallback struct {
	successes int32
	failures  int32
}

func (c *countingCallback) Success(result *Result) {
	atomic.AddInt32(&c.successes, 1)
}

func (c *countingCallback) Fail(result *Result) {
	atomic.AddInt32(&c.failures, 1)
}

type waitCallback struct {
	countingCallback
	done chan struct{}
}

func (c *waitCallback) Success(result *Result) {
	c.countingCallback.Success(result)
	close(c.done)
}

func (c *waitCallback) Fail(result *Result) {
	c.countingCallback.Fail(result)
	close(c.done)
}

type resourceReleaseCallback struct {
	t         *testing.T
	producer  *producer
	successes int32
	failures  int32
}

func (c *resourceReleaseCallback) Success(result *Result) {
	c.checkReleased()
	atomic.AddInt32(&c.successes, 1)
}

func (c *resourceReleaseCallback) Fail(result *Result) {
	c.checkReleased()
	atomic.AddInt32(&c.failures, 1)
}

func (c *resourceReleaseCallback) checkReleased() {
	c.t.Helper()
	if got := c.producer.memoryLimiter.Used(); got != 0 {
		c.t.Fatalf("expected memory to be released before callback, got used=%d", got)
	}
	if got := c.producer.currentProducerLogGroupSize(); got != 0 {
		c.t.Fatalf("expected producer size to be released before callback, got %d", got)
	}
}

type always429RoundTripper struct {
	calls int32
}

func (rt *always429RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	atomic.AddInt32(&rt.calls, 1)
	return &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(`{"errorCode":"TooManyRequests","errorMessage":"slow down"}`)),
		Request:    req,
	}, nil
}

type retryThenSuccessRoundTripper struct {
	calls int32
}

func (rt *retryThenSuccessRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	call := atomic.AddInt32(&rt.calls, 1)
	if call == 1 {
		return &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"errorCode":"TooManyRequests","errorMessage":"slow down"}`)),
			Request:    req,
		}, nil
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(`{}`)),
		Request:    req,
	}, nil
}

type blockingPutLogsClient struct {
	tls.Client
	calls         int32
	releaseFirst  chan struct{}
	firstStarted  chan struct{}
	secondStarted chan struct{}
}

func (c *blockingPutLogsClient) PutLogs(req *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	call := atomic.AddInt32(&c.calls, 1)
	if call == 1 {
		close(c.firstStarted)
		<-c.releaseFirst
		return &tls.CommonResponse{}, nil
	}
	close(c.secondStarted)
	select {}
}

func (c *blockingPutLogsClient) GetRetryPolicy() tls.RetryPolicy {
	return tls.DefaultRetryPolicy()
}

func (c *blockingPutLogsClient) GetHttpClient() *http.Client {
	return nil
}

func (c *blockingPutLogsClient) SetHttpClient(client *http.Client) error {
	return nil
}

func (c *blockingPutLogsClient) SetTimeout(timeout time.Duration) {
}

func (c *blockingPutLogsClient) SetRetryPolicy(policy tls.RetryPolicy) {
}

func (c *blockingPutLogsClient) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
}

func (c *blockingPutLogsClient) SetAPIVersion(version string) {
}

func (c *blockingPutLogsClient) SetCustomUserAgent(customUserAgent string) {
}

func (c *blockingPutLogsClient) PutLogsV2(request *tls.PutLogsV2Request) (*tls.CommonResponse, error) {
	return &tls.CommonResponse{}, nil
}

type countingPutLogsClient struct {
	tls.Client
	calls int32
}

func (c *countingPutLogsClient) PutLogs(req *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	atomic.AddInt32(&c.calls, 1)
	return nil, &tls.Error{HTTPCode: 429}
}

type successPutLogsClient struct {
	tls.Client
}

func (c *successPutLogsClient) PutLogs(req *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	return &tls.CommonResponse{RequestID: "success"}, nil
}

type permanentErrorPutLogsClient struct {
	tls.Client
}

func (c *permanentErrorPutLogsClient) PutLogs(req *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	return nil, &tls.Error{HTTPCode: 400}
}
