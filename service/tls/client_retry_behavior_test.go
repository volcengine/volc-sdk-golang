package tls

import (
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestPutLogsUsesClientRetryPolicyBeforeReturningToCaller(t *testing.T) {
	transport := &retryCountingRoundTripper{}
	client := NewClient("http://example.com", "ak", "sk", "", "cn-north-1")
	if err := client.SetHttpClient(&http.Client{Transport: transport, Timeout: time.Second}); err != nil {
		t.Fatalf("SetHttpClient failed: %v", err)
	}
	client.SetRetryPolicy(RetryPolicy{
		TotalTimeout:        30 * time.Second,
		InitialInterval:     100 * time.Millisecond,
		MaxInterval:         100 * time.Millisecond,
		Multiplier:          1,
		RandomizationFactor: 0.1,
		MaxAttempts:         2,
	})

	_, err := client.PutLogs(&PutLogsRequest{
		TopicID:      "topic",
		CompressType: CompressLz4,
		LogBody: &pb.LogGroupList{LogGroups: []*pb.LogGroup{{
			Logs: []*pb.Log{{Time: 1, Contents: []*pb.LogContent{{Key: "k", Value: "v"}}}},
		}}},
	})
	if err != nil {
		t.Fatalf("PutLogs failed: %v", err)
	}
	if got := atomic.LoadInt32(&transport.calls); got != 2 {
		t.Fatalf("expected one client-level retry before returning, got %d calls", got)
	}
}

type retryCountingRoundTripper struct {
	calls int32
}

func (rt *retryCountingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
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
