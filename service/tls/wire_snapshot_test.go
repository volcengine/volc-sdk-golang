package tls

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"
)

// Cross-SDK wire snapshot dumper.
//
// 按 cospec/.../context/l2-fixtures.json 字面量构造请求，捕获实际 wire 形态，
// 输出到 cospec/.../context/l2-snapshots/go.json。后续由 Python 比对脚本对照
// 4 SDK snapshot 生成 4×5 一致性矩阵。

type wireSnapshot struct {
	Method string                     `json:"method"`
	Path   string                     `json:"path"`
	Query  map[string]string          `json:"query"`
	Body   map[string]json.RawMessage `json:"body"`
}

func captureWire(t *testing.T, fn func(c *LsClient)) wireSnapshot {
	t.Helper()
	var snap wireSnapshot
	transport := rtFunc(func(r *http.Request) (*http.Response, error) {
		snap.Method = r.Method
		snap.Path = r.URL.Path
		snap.Query = map[string]string{}
		for k, v := range r.URL.Query() {
			if len(v) > 0 {
				snap.Query[k] = v[0]
			} else {
				snap.Query[k] = ""
			}
		}
		if r.Body != nil {
			raw, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("read body: %v", err)
			}
			r.Body = io.NopCloser(bytes.NewReader(raw))
			if len(raw) > 0 {
				if err := json.Unmarshal(raw, &snap.Body); err != nil {
					t.Fatalf("body not JSON object: %v body=%s", err, raw)
				}
			}
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"X-Tls-Requestid": []string{"req"}},
			Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
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
	fn(cli)
	return snap
}

func TestWireSnapshotDump(t *testing.T) {
	outDir := firstNonEmptyEnv("WIRE_SNAPSHOT_OUT_DIR", "L2_OUT_DIR")
	if outDir == "" {
		t.Skip("WIRE_SNAPSHOT_OUT_DIR not set; skipping")
	}

	out := map[string]interface{}{
		"sdk":        "go",
		"sdk_repo":   "volc-sdk-golang",
		"interfaces": map[string]wireSnapshot{},
	}
	ifaces := out["interfaces"].(map[string]wireSnapshot)

	ifaces["DescribeCursor"] = captureWire(t, func(c *LsClient) {
		_, _ = c.DescribeCursor(&DescribeCursorRequest{TopicID: "tid-cursor", ShardID: 1, From: "begin"})
	})
	ifaces["DescribeCheckPoint"] = captureWire(t, func(c *LsClient) {
		_, _ = c.DescribeCheckPoint(&DescribeCheckPointRequest{
			ProjectID: "pid-ck", TopicID: "tid-ck", ShardID: 2, ConsumerGroupName: "g1",
		})
	})
	ifaces["SearchLogs"] = captureWire(t, func(c *LsClient) {
		accurate, mustComplete := true, true
		var offset int64 = 0
		_, _ = c.SearchLogs(&SearchLogsRequest{
			TopicID:       "tid-search",
			Query:         "*",
			StartTime:     1700000000,
			EndTime:       1700001000,
			Limit:         20,
			Context:       "",
			Sort:          "asc",
			HighLight:     false,
			AccurateQuery: &accurate,
			MustComplete:  &mustComplete,
			Offset:        &offset,
		})
	})
	ifaces["CreateIndex"] = captureWire(t, func(c *LsClient) {
		enable := true
		maxLen := int32(2048)
		_, _ = c.CreateIndex(&CreateIndexRequest{
			TopicID:           "tid-idx",
			MaxTextLen:        &maxLen,
			EnableAutoIndex:   &enable,
			EnablePhraseIndex: &enable,
		})
	})
	ifaces["CreateAlarm"] = captureWire(t, func(c *LsClient) {
		status := true
		userMsg := "msg"
		severity := "warning"
		sendResolved := false
		_, _ = c.CreateAlarm(&CreateAlarmRequest{
			AlarmName:        "alarm-1",
			ProjectID:        "pid-alarm",
			Status:           &status,
			QueryRequest:     QueryRequests{},
			RequestCycle:     RequestCycle{Type: "Period", Time: 5},
			Condition:        "x>1",
			TriggerPeriod:    1,
			AlarmPeriod:      5,
			AlarmNotifyGroup: []string{"g-1"},
			UserDefineMsg:    &userMsg,
			Severity:         &severity,
			SendResolved:     &sendResolved,
		})
	})

	// 写到 cospec 仓库统一目录（由 WIRE_SNAPSHOT_OUT_DIR 提供）。
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	outPath := filepath.Join(outDir, "go.json")
	// 排序输出
	pretty, err := marshalSortedSnapshot(out)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if err := os.WriteFile(outPath, pretty, 0o644); err != nil {
		t.Fatalf("write %s: %v", outPath, err)
	}
	t.Logf("wire snapshot written: %s", outPath)
}

func marshalSortedSnapshot(v map[string]interface{}) ([]byte, error) {
	// Encode interfaces map sorted by key for deterministic output.
	ifaces := v["interfaces"].(map[string]wireSnapshot)
	keys := make([]string, 0, len(ifaces))
	for k := range ifaces {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	buf.WriteString("{\n  \"sdk\": ")
	enc, _ := json.Marshal(v["sdk"])
	buf.Write(enc)
	buf.WriteString(",\n  \"sdk_repo\": ")
	enc, _ = json.Marshal(v["sdk_repo"])
	buf.Write(enc)
	buf.WriteString(",\n  \"interfaces\": {")
	for i, k := range keys {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("\n    ")
		nameEnc, _ := json.Marshal(k)
		buf.Write(nameEnc)
		buf.WriteString(": ")
		snapEnc, err := json.MarshalIndent(ifaces[k], "    ", "  ")
		if err != nil {
			return nil, err
		}
		buf.Write(snapEnc)
	}
	buf.WriteString("\n  }\n}\n")
	return buf.Bytes(), nil
}
