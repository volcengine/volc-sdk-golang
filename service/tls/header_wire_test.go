package tls

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Header wire 抓包断言：
// 1) SearchLogs highlight=true → SourceType=front 必现；highlight=false → 不出现
// 2) CommonRequest.Headers 自定义 header 进入 wire；与 SDK 内部 header 冲突时 SDK 覆盖

func TestSearchLogsHeader_SearchLogs_HighlightTrue_InjectsSourceType(t *testing.T) {
	var captured http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.Header.Clone()
		w.Header().Set(RequestIDHeader, "rid")
		_, _ = w.Write([]byte(`{"Logs":[]}`))
	}))
	defer srv.Close()

	c := newClient(srv.URL, "ak", "sk", "", "cn-guilin-boe", "").(*LsClient)
	c.Client.Timeout = 2 * time.Second

	req := &SearchLogsRequest{
		TopicID:   "tid",
		Query:     "*",
		StartTime: 1,
		EndTime:   2,
		Limit:     1,
		HighLight: true,
	}
	_, _ = c.SearchLogs(req)

	if got := captured.Get("SourceType"); got != "front" {
		t.Fatalf("highlight=true 应注入 SourceType=front, 实际: %q", got)
	}
}

func TestSearchLogsHeader_SearchLogs_HighlightFalse_NoSourceType(t *testing.T) {
	var captured http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.Header.Clone()
		w.Header().Set(RequestIDHeader, "rid")
		_, _ = w.Write([]byte(`{"Logs":[]}`))
	}))
	defer srv.Close()

	c := newClient(srv.URL, "ak", "sk", "", "cn-guilin-boe", "").(*LsClient)
	c.Client.Timeout = 2 * time.Second

	req := &SearchLogsRequest{
		TopicID:   "tid",
		Query:     "*",
		StartTime: 1,
		EndTime:   2,
		Limit:     1,
		HighLight: false,
	}
	_, _ = c.SearchLogs(req)

	if got := captured.Get("SourceType"); got != "" {
		t.Fatalf("highlight=false 不应注入 SourceType, 实际: %q", got)
	}
}

func TestCustomHeader_CommonRequestHeaders_PassThroughToWire(t *testing.T) {
	var captured http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.Header.Clone()
		w.Header().Set(RequestIDHeader, "rid")
		_, _ = w.Write([]byte(`{"Logs":[]}`))
	}))
	defer srv.Close()

	c := newClient(srv.URL, "ak", "sk", "", "cn-guilin-boe", "").(*LsClient)
	c.Client.Timeout = 2 * time.Second

	req := &SearchLogsRequest{
		CommonRequest: CommonRequest{
			Headers: map[string]string{
				"X-Trace-Id":   "trace-abc",
				"X-Tenant-Tag": "team-foo",
			},
		},
		TopicID:   "tid",
		Query:     "*",
		StartTime: 1,
		EndTime:   2,
		Limit:     1,
	}
	_, _ = c.SearchLogs(req)

	if got := captured.Get("X-Trace-Id"); got != "trace-abc" {
		t.Fatalf("自定义 header 未抵达 wire, X-Trace-Id=%q", got)
	}
	if got := captured.Get("X-Tenant-Tag"); got != "team-foo" {
		t.Fatalf("自定义 header 未抵达 wire, X-Tenant-Tag=%q", got)
	}
}

func TestCustomHeader_SDKHeaderOverridesUserHeader(t *testing.T) {
	var captured http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.Header.Clone()
		w.Header().Set(RequestIDHeader, "rid")
		_, _ = w.Write([]byte(`{"Logs":[]}`))
	}))
	defer srv.Close()

	c := newClient(srv.URL, "ak", "sk", "", "cn-guilin-boe", "").(*LsClient)
	c.Client.Timeout = 2 * time.Second

	// 用户尝试用自定义 header 覆盖 SDK 协议字段 SourceType。
	// 同时显式开 highlight=true，触发 SDK 写入 SourceType=front。
	// 期望：SDK 后写覆盖用户值（user-first → SDK overrides）。
	req := &SearchLogsRequest{
		CommonRequest: CommonRequest{
			Headers: map[string]string{
				"SourceType": "user-injected",
			},
		},
		TopicID:   "tid",
		Query:     "*",
		StartTime: 1,
		EndTime:   2,
		Limit:     1,
		HighLight: true,
	}
	_, _ = c.SearchLogs(req)

	if got := captured.Get("SourceType"); got != "front" {
		t.Fatalf("SDK 应覆盖用户 SourceType, 实际: %q (期望 front)", got)
	}
}
