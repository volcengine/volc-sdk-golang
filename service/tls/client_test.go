package tls

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"sync"
	"testing"
)

type rtFunc func(*http.Request) (*http.Response, error)

func (f rtFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

type trackingReadCloser struct {
	r      io.Reader
	closed int
}

func (t *trackingReadCloser) Read(p []byte) (int, error) {
	return t.r.Read(p)
}

func (t *trackingReadCloser) Close() error {
	t.closed++
	return nil
}

func gzipBytes(b []byte) []byte {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, _ = w.Write(b)
	_ = w.Close()
	return buf.Bytes()
}

func newTestClient(rt http.RoundTripper) *LsClient {
	return &LsClient{
		Client:          &http.Client{Transport: rt},
		Endpoint:        "http://example.com",
		accessLock:      &sync.RWMutex{},
		Region:          "cn-beijing",
		AccessKeyID:     "ak",
		AccessKeySecret: "sk",
		APIVersion:      "0.3.0",
	}
}

func TestRealRequest_GzipCloseClosesUnderlyingBody(t *testing.T) {
	body := &trackingReadCloser{r: bytes.NewReader(gzipBytes([]byte("ok")))}
	c := newTestClient(rtFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Encoding": []string{CompressGz}},
			Body:       body,
			Request:    r,
		}, nil
	}))

	resp, err := c.realRequest(context.Background(), http.MethodGet, "/DescribeTopic", map[string]string{}, nil)
	if err != nil {
		t.Fatalf("expected nil err, got %v", err)
	}
	_ = resp.Body.Close()
	if body.closed != 1 {
		t.Fatalf("expected underlying body close once, got %d", body.closed)
	}
}

func TestRealRequest_GzipNewReaderErrorClosesUnderlyingBody(t *testing.T) {
	body := &trackingReadCloser{r: bytes.NewReader([]byte("not-gzip"))}
	c := newTestClient(rtFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Encoding": []string{CompressGz}},
			Body:       body,
			Request:    r,
		}, nil
	}))

	resp, err := c.realRequest(context.Background(), http.MethodGet, "/DescribeTopic", map[string]string{}, nil)
	if err == nil || resp != nil {
		t.Fatalf("expected error and nil resp, got resp=%v err=%v", resp, err)
	}
	if body.closed != 1 {
		t.Fatalf("expected underlying body close once, got %d", body.closed)
	}
}

func TestNewClientWithAPIKeyPutLogsUsesAnonymousHeader(t *testing.T) {
	var gotHeader http.Header
	c := NewClientWithAPIKey("http://example.com", "cn-beijing", "api-key").(*LsClient)
	if err := c.SetHttpClient(&http.Client{Transport: rtFunc(func(r *http.Request) (*http.Response, error) {
		gotHeader = r.Header.Clone()
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body:       io.NopCloser(bytes.NewReader(nil)),
			Request:    r,
		}, nil
	})}); err != nil {
		t.Fatalf("unexpected SetHttpClient error: %v", err)
	}

	_, err := c.realRequest(context.Background(), http.MethodPost, PathPutLogs, map[string]string{}, []byte("body"))
	if err != nil {
		t.Fatalf("unexpected request error: %v", err)
	}
	if got := gotHeader.Get(AnonymousIdentityHeader); got != "api-key" {
		t.Fatalf("expected anonymous header, got %q", got)
	}
	if got := gotHeader.Get("Authorization"); got != "" {
		t.Fatalf("expected unsigned anonymous request, got Authorization %q", got)
	}
}

func TestClientInterfaceSetAPIKeyUpdatesAnonymousHeader(t *testing.T) {
	var gotHeader http.Header
	var c Client = NewClientWithAPIKey("http://example.com", "cn-beijing", "old-api-key")
	c.SetAPIKey("new-api-key")
	lc := c.(*LsClient)
	if err := lc.SetHttpClient(&http.Client{Transport: rtFunc(func(r *http.Request) (*http.Response, error) {
		gotHeader = r.Header.Clone()
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body:       io.NopCloser(bytes.NewReader(nil)),
			Request:    r,
		}, nil
	})}); err != nil {
		t.Fatalf("unexpected SetHttpClient error: %v", err)
	}

	_, err := lc.realRequest(context.Background(), http.MethodPost, PathPutLogs, map[string]string{}, []byte("body"))
	if err != nil {
		t.Fatalf("unexpected request error: %v", err)
	}
	if got := gotHeader.Get(AnonymousIdentityHeader); got != "new-api-key" {
		t.Fatalf("expected updated anonymous header, got %q", got)
	}
}

func TestAPIKeyOnlyNonAnonymousActionReturnsMissingCredentialsBeforeNetwork(t *testing.T) {
	c := NewClientWithAPIKey("http://example.com", "cn-beijing", "api-key").(*LsClient)
	called := false
	if err := c.SetHttpClient(&http.Client{Transport: rtFunc(func(r *http.Request) (*http.Response, error) {
		called = true
		return nil, nil
	})}); err != nil {
		t.Fatalf("unexpected SetHttpClient error: %v", err)
	}

	_, err := c.realRequest(context.Background(), http.MethodGet, "/DescribeTopic", map[string]string{}, nil)
	if err == nil {
		t.Fatalf("expected missing credentials error")
	}
	if called {
		t.Fatalf("expected request to fail before network")
	}
}

func TestAPIKeyWithCredentialsUsesAKSKForNonAnonymousAction(t *testing.T) {
	var gotHeader http.Header
	c := NewClientWithAPIKey("http://example.com", "cn-beijing", "api-key", "ak", "sk", "").(*LsClient)
	if err := c.SetHttpClient(&http.Client{Transport: rtFunc(func(r *http.Request) (*http.Response, error) {
		gotHeader = r.Header.Clone()
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body:       io.NopCloser(bytes.NewReader(nil)),
			Request:    r,
		}, nil
	})}); err != nil {
		t.Fatalf("unexpected SetHttpClient error: %v", err)
	}

	_, err := c.realRequest(context.Background(), http.MethodGet, "/DescribeTopic", map[string]string{}, nil)
	if err != nil {
		t.Fatalf("unexpected request error: %v", err)
	}
	if got := gotHeader.Get(AnonymousIdentityHeader); got != "" {
		t.Fatalf("expected non-anonymous action to omit anonymous header, got %q", got)
	}
	if got := gotHeader.Get("Authorization"); got == "" {
		t.Fatalf("expected signed request")
	}
}
