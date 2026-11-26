package consumer

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type processorTransport func(*http.Request) (*http.Response, error)

func (f processorTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func processorTestConfig() *Config {
	cfg := GetDefaultConsumerConfig()
	cfg.Endpoint = "http://example.com"
	cfg.Region = "cn-beijing"
	cfg.AccessKeyID = "ak"
	cfg.AccessKeySecret = "sk"
	cfg.ProjectID = "project"
	cfg.TopicIDList = []string{"topic"}
	cfg.ConsumerGroupName = "group"
	cfg.ConsumerName = "consumer"
	return cfg
}

func TestConsumerProcessorFetchContract(t *testing.T) {
	cases := []struct {
		name, processorID, path string
		original                bool
	}{
		{name: "default original", original: true, path: tls.PathConsumeOriginalLogs},
		{name: "explicit normal", original: false, path: tls.PathConsumeLogs},
		{name: "processor selects normal despite default original", processorID: "not-a-uuid", original: true, path: tls.PathConsumeLogs},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := processorTestConfig()
			cfg.Original = tc.original
			cfg.ProcessorID = tc.processorID
			consumerInstance, err := NewConsumer(context.Background(), cfg, func(string, int, *pb.LogGroupList) {})
			if err != nil {
				t.Fatalf("new consumer: %v", err)
			}
			c := consumerInstance.(*consumer)

			calls := 0
			transport := processorTransport(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != http.MethodGet || req.URL.Path != tc.path {
					t.Fatalf("method/path = %s %s, want GET %s", req.Method, req.URL.Path, tc.path)
				}
				if req.URL.Query().Get("TopicId") != "topic" || req.URL.Query().Get("ShardId") != "7" {
					t.Fatalf("query = %v", req.URL.Query())
				}
				if req.Header.Get("ConsumerGroupName") != "group" || req.Header.Get("ConsumerName") != "consumer" {
					t.Fatalf("headers = %v", req.Header)
				}
				var body map[string]interface{}
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Fatal(err)
				}
				expectedCursor := "cursor"
				if calls > 1 {
					expectedCursor = "next"
				}
				if body["Cursor"] != expectedCursor {
					t.Fatalf("cursor = %v, want %s", body["Cursor"], expectedCursor)
				}
				id, present := body["ProcessorId"]
				if present != (tc.processorID != "") || (present && id != tc.processorID) {
					t.Fatalf("ProcessorId = %v, present = %v", id, present)
				}
				headers := make(http.Header)
				headers.Set("x-tls-count", "0")
				headers.Set("x-tls-cursor", "next")
				headers.Set("x-tls-original", "false")
				headers.Set("x-tls-bodyrawsize", "0")
				return &http.Response{StatusCode: http.StatusOK, Header: headers, Body: io.NopCloser(strings.NewReader("")), Request: req}, nil
			})
			if err := c.client.SetHttpClient(&http.Client{Transport: transport}); err != nil {
				t.Fatal(err)
			}
			lc := &logConsumer{client: c.client, logger: log.NewNopLogger(), conf: c.conf, shard: &tls.ConsumeShard{TopicID: "topic", ShardID: 7}, nextCheckpoint: "cursor"}
			for i := 0; i < 2; i++ {
				if err := lc.fetchData(); err != nil {
					t.Fatalf("fetch %d: %v", i, err)
				}
			}
			if calls != 2 {
				t.Fatalf("calls = %d, want 2", calls)
			}
			if cfg.Original != tc.original {
				t.Fatalf("consumer changed Original from %v to %v", tc.original, cfg.Original)
			}
		})
	}
}
