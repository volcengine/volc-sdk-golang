package producer

import (
	"errors"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls"
)

// acquirePayload reserves a test fixture through the production admission path
// when the fixture only needs the error, not the rejection classification.
func (limiter *memoryLimiter) acquirePayload(bytes int64, timeout time.Duration) error {
	_, _, err := limiter.acquirePayloadOutcome(bytes, timeout)
	return err
}

func (limiter *memoryLimiter) acquireTemporary(bytes int64, timeout time.Duration) error {
	return limiter.acquireTemporaryWithCancel(bytes, timeout, nil)
}

func (limiter *memoryLimiter) releasePayload(bytes int64) {
	limiter.releasePayloadFrom(payloadStageIngress, bytes)
}

// initProducerBatch builds a populated test batch through the production
// constructor and aggregation path instead of duplicating their defaults.
func initProducerBatch(batchLog *BatchLog, config *Config) *Batch {
	batch := newProducerBatch(batchLog, config)
	batch.tryAddLog(batchLog, batchLog.Key.CallBackFun)
	return batch
}

// mockProducerClient is a shared test double for producer lifecycle, queue,
// callback, and statistics behavior. It overrides the client methods used by
// the producer so tests do not perform network I/O.
type mockProducerClient struct {
	tls.Client
	putLogsCount  int64
	failCount     int64
	retryableFail bool
	totalLogs     int64
}

// newTestProducerWithClient constructs a normal producer and replaces its
// transport with a test double. Keeping construction on newProducer ensures
// tests exercise the same production initialization path.
func newTestProducerWithClient(producerConfig *Config, client tls.Client) *producer {
	p := newProducer(producerConfig)
	p.cli = client
	if p.threadPool != nil && p.threadPool.sender != nil {
		p.threadPool.sender.client = client
	}
	return p
}

func (m *mockProducerClient) PutLogs(request *tls.PutLogsRequest) (*tls.CommonResponse, error) {
	n := atomic.AddInt64(&m.putLogsCount, 1)
	if request.LogBody != nil {
		for _, lg := range request.LogBody.LogGroups {
			atomic.AddInt64(&m.totalLogs, int64(len(lg.Logs)))
		}
	}
	if n <= atomic.LoadInt64(&m.failCount) {
		if m.retryableFail {
			return nil, &tls.Error{HTTPCode: 500, Code: "InternalError", Message: "mock retryable failure"}
		}
		return nil, errors.New("mock non-retryable failure")
	}
	return &tls.CommonResponse{RequestID: "mock-putlogs"}, nil
}

func (m *mockProducerClient) PutLogsV2(request *tls.PutLogsV2Request) (*tls.CommonResponse, error) {
	atomic.AddInt64(&m.putLogsCount, 1)
	return &tls.CommonResponse{RequestID: "mock-putlogs-v2"}, nil
}

// Override the client configuration methods so producer setup never reaches
// the nil embedded client in tests.
func (m *mockProducerClient) GetHttpClient() *http.Client                { return nil }
func (m *mockProducerClient) SetHttpClient(*http.Client) error           { return nil }
func (m *mockProducerClient) SetTimeout(time.Duration)                   {}
func (m *mockProducerClient) GetRetryPolicy() tls.RetryPolicy            { return tls.RetryPolicy{} }
func (m *mockProducerClient) SetRetryPolicy(tls.RetryPolicy)             {}
func (m *mockProducerClient) ResetAccessKeyToken(string, string, string) {}
