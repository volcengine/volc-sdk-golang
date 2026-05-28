package tls

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
	"testing"
)

// L1 wire-contract tests.
//
// 目标：锁定 5 个核心接口经 SDK 实际发出的 HTTP wire 形态（method/path/query keys/body keys），
// 防止后续重构静默改变 wire 接口契约。基线见
// cospec/changes/check-tls-sdk-contract-alignment/context/wire-baseline.json。

type captured struct {
	method      string
	path        string
	queryKeys   []string
	queryValues map[string][]string
	bodyKeys    []string
	headers     http.Header
}

// newWireClient returns an LsClient whose http.Client captures the next outgoing
// request and returns an empty 200 OK response so that each interface's caller
// finishes successfully.
func newWireClient(t *testing.T, c *captured) *LsClient {
	t.Helper()
	transport := rtFunc(func(r *http.Request) (*http.Response, error) {
		c.method = r.Method
		c.path = r.URL.Path
		c.headers = r.Header.Clone()

		// Sorted query keys (no values) for deterministic comparison.
		qk := make([]string, 0, len(r.URL.Query()))
		c.queryValues = map[string][]string{}
		for k, values := range r.URL.Query() {
			qk = append(qk, k)
			c.queryValues[k] = append([]string(nil), values...)
		}
		sort.Strings(qk)
		c.queryKeys = qk

		// Sorted top-level body keys.
		c.bodyKeys = nil
		if r.Body != nil {
			raw, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("read body: %v", err)
			}
			r.Body = io.NopCloser(bytes.NewReader(raw))
			if len(raw) > 0 {
				var top map[string]json.RawMessage
				if err := json.Unmarshal(raw, &top); err != nil {
					t.Fatalf("body is not a JSON object: %v\nbody=%s", err, raw)
				}
				bk := make([]string, 0, len(top))
				for k := range top {
					bk = append(bk, k)
				}
				sort.Strings(bk)
				c.bodyKeys = bk
			}
		}

		// Empty 200 with valid JSON body so SDK parsers don't blow up.
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"X-Tls-Requestid": []string{"req-test"}},
			Body:       io.NopCloser(strings.NewReader("{}")),
			Request:    r,
		}, nil
	})
	cli := &LsClient{
		Client:          &http.Client{Transport: transport},
		Endpoint:        "http://example.com",
		accessLock:      &sync.RWMutex{},
		Region:          "cn-beijing",
		AccessKeyID:     "ak",
		AccessKeySecret: "sk",
		APIVersion:      "0.3.0",
	}
	cli.retryPolicy.Store(DefaultRetryPolicy())
	return cli
}

func equalSorted(t *testing.T, name string, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("%s mismatch: got=%v want=%v", name, got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("%s mismatch: got=%v want=%v", name, got, want)
		}
	}
}

func mustContextOK(t *testing.T) context.Context {
	t.Helper()
	return context.Background()
}

func TestWireContract_DescribeCursor(t *testing.T) {
	cap := &captured{}
	c := newWireClient(t, cap)

	_, err := c.DescribeCursor(&DescribeCursorRequest{
		TopicID: "tid", ShardID: 1, From: "begin",
	})
	if err != nil {
		t.Fatalf("DescribeCursor returned err: %v", err)
	}
	if cap.method != http.MethodPost {
		t.Fatalf("method got %s want POST", cap.method)
	}
	if cap.path != "/DescribeCursor" {
		t.Fatalf("path got %s", cap.path)
	}
	equalSorted(t, "DescribeCursor.queryKeys", cap.queryKeys, []string{"ShardId", "TopicId"})
	equalSorted(t, "DescribeCursor.bodyKeys", cap.bodyKeys, []string{"From"})
}

func TestWireContract_DescribeCheckPoint(t *testing.T) {
	cap := &captured{}
	c := newWireClient(t, cap)

	_, err := c.DescribeCheckPoint(&DescribeCheckPointRequest{
		ProjectID: "pid", TopicID: "tid", ShardID: 1, ConsumerGroupName: "g1",
	})
	if err != nil {
		t.Fatalf("DescribeCheckPoint returned err: %v", err)
	}
	if cap.method != http.MethodPost {
		t.Fatalf("method got %s want POST", cap.method)
	}
	if cap.path != "/DescribeCheckPoint" {
		t.Fatalf("path got %s", cap.path)
	}
	equalSorted(t, "DescribeCheckPoint.queryKeys", cap.queryKeys, []string{"ProjectId", "ShardId", "TopicId"})
	equalSorted(t, "DescribeCheckPoint.bodyKeys", cap.bodyKeys, []string{"ConsumerGroupName"})
}

func TestWireContract_SearchLogs(t *testing.T) {
	cap := &captured{}
	c := newWireClient(t, cap)

	accurate := true
	mustComplete := true
	var offset int64 = 100
	_, err := c.SearchLogs(&SearchLogsRequest{
		TopicID:       "tid",
		Query:         "*",
		StartTime:     1,
		EndTime:       2,
		Limit:         20,
		HighLight:     false,
		Context:       "",
		Sort:          "asc",
		AccurateQuery: &accurate,
		MustComplete:  &mustComplete,
		Offset:        &offset,
	})
	if err != nil {
		t.Fatalf("SearchLogs returned err: %v", err)
	}
	if cap.method != http.MethodPost {
		t.Fatalf("method got %s want POST", cap.method)
	}
	if cap.path != "/SearchLogs" {
		t.Fatalf("path got %s", cap.path)
	}
	if len(cap.queryKeys) != 0 {
		t.Fatalf("expected no query keys for SearchLogs, got %v", cap.queryKeys)
	}
	// Baseline expectation: 12 keys including AccurateQuery/MustComplete/Offset/RegionTopics.
	// Go SDK SearchLogsRequest 当前未声明 RegionTopics 字段，故此处 want 不含 RegionTopics。
	// 该差异在 L1 报告中作为 SDK 字段缺失项（与 Java 不一致）。
	want := []string{"AccurateQuery", "Context", "EndTime", "HighLight", "Limit", "MustComplete", "Offset", "Query", "Sort", "StartTime", "TopicId"}
	equalSorted(t, "SearchLogs.bodyKeys", cap.bodyKeys, want)
}

func TestWireContract_CreateIndex(t *testing.T) {
	cap := &captured{}
	c := newWireClient(t, cap)

	enable := true
	maxLen := int32(2048)
	_, err := c.CreateIndex(&CreateIndexRequest{
		TopicID:           "tid",
		FullText:          &FullTextInfo{Delimiter: " ", CaseSensitive: false, IncludeChinese: true},
		KeyValue:          &[]KeyValueInfo{},
		UserInnerKeyValue: &[]KeyValueInfo{},
		MaxTextLen:        &maxLen,
		EnableAutoIndex:   &enable,
		EnablePhraseIndex: &enable,
	})
	if err != nil {
		t.Fatalf("CreateIndex returned err: %v", err)
	}
	if cap.method != http.MethodPost {
		t.Fatalf("method got %s want POST", cap.method)
	}
	if cap.path != "/CreateIndex" {
		t.Fatalf("path got %s", cap.path)
	}
	if len(cap.queryKeys) != 0 {
		t.Fatalf("expected no query keys for CreateIndex, got %v", cap.queryKeys)
	}
	// Baseline includes LogReduce / LogReduceBlackList / LogReduceWhiteList; SDK lacks them.
	// L1 violation captured here (locks current emitted set).
	want := []string{"EnableAutoIndex", "EnablePhraseIndex", "FullText", "KeyValue", "MaxTextLen", "TopicId", "UserInnerKeyValue"}
	equalSorted(t, "CreateIndex.bodyKeys", cap.bodyKeys, want)
}

func TestWireContract_CreateAlarm(t *testing.T) {
	cap := &captured{}
	c := newWireClient(t, cap)

	status := true
	severity := "warn"
	userMsg := "x"
	sendResolved := false
	_, err := c.CreateAlarm(&CreateAlarmRequest{
		AlarmName:          "a1",
		ProjectID:          "pid",
		Status:             &status,
		QueryRequest:       QueryRequests{},
		RequestCycle:       RequestCycle{Type: "Period", Time: 5},
		Condition:          "x>1",
		TriggerPeriod:      1,
		AlarmPeriod:        1,
		AlarmNotifyGroup:   []string{"g"},
		UserDefineMsg:      &userMsg,
		Severity:           &severity,
		AlarmPeriodDetail:  &AlarmPeriodSetting{Sms: 0, Phone: 0, Email: 0, GeneralWebhook: 0},
		JoinConfigurations: []JoinConfig{{Condition: "x", SetOperationType: "AND"}},
		TriggerConditions:  []TriggerCondition{{Severity: "warn", Condition: "x>1"}},
		SendResolved:       &sendResolved,
	})
	if err != nil {
		t.Fatalf("CreateAlarm returned err: %v", err)
	}
	if cap.method != http.MethodPost {
		t.Fatalf("method got %s want POST", cap.method)
	}
	if cap.path != "/CreateAlarm" {
		t.Fatalf("path got %s", cap.path)
	}
	if len(cap.queryKeys) != 0 {
		t.Fatalf("expected no query keys for CreateAlarm, got %v", cap.queryKeys)
	}
	want := []string{
		"AlarmName", "AlarmNotifyGroup", "AlarmPeriod", "AlarmPeriodDetail",
		"Condition", "JoinConfigurations", "ProjectId", "QueryRequest",
		"RequestCycle", "SendResolved", "Severity", "Status",
		"TriggerConditions", "TriggerPeriod", "UserDefineMsg",
	}
	equalSorted(t, "CreateAlarm.bodyKeys", cap.bodyKeys, want)
}

func TestWireContract_DescribeHostGroupsV2HiddenQueryParameter(t *testing.T) {
	cap := &captured{}
	c := newWireClient(t, cap)

	hidden := true
	_, err := c.DescribeHostGroupsV2(&DescribeHostGroupsRequestV2{Hidden: &hidden})
	if err != nil {
		t.Fatalf("DescribeHostGroupsV2 returned err: %v", err)
	}
	if cap.method != http.MethodGet {
		t.Fatalf("method got %s want GET", cap.method)
	}
	if cap.path != "/DescribeHostGroupsV2" {
		t.Fatalf("path got %s", cap.path)
	}
	equalSorted(t, "DescribeHostGroupsV2.queryKeys", cap.queryKeys, []string{"Hidden"})
	if got := cap.queryValues["Hidden"]; len(got) != 1 || got[0] != "true" {
		t.Fatalf("Hidden query got %v want [true]", got)
	}
}

// silence unused linter warnings in case some interfaces evolve.
var _ = mustContextOK
