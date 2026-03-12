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
