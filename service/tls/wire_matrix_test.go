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

// 真实环境矩阵 dump：5 接口 × 1 case，每接口落 wire + 服务端 status + 响应 body。
//
// 必需 env：
//   WIRE_MATRIX_OUT_DIR / LOG_SERVICE_ENDPOINT / LOG_SERVICE_AK / LOG_SERVICE_SK
//   LOG_SERVICE_REGION / WIRE_PROJECT_ID / WIRE_TOPIC_ID / WIRE_CONSUMER_GROUP
// 输出：${WIRE_MATRIX_OUT_DIR}/go.jsonl，每行 {iface, case, method, path, query, body, status, req_id, resp_body, err}。
//
// case 命名按字母序与其他 SDK 对齐。CreateAlarm 故意伪造 NotifyGroup 期望 4xx，仅比 wire。
//
// 用途：resp_body 是 4 SDK Response 反序列化回归 fixture 的唯一来源（响应由服务端定，
// 仅需 Go 一次抽样即可，4 SDK 各自跑反序列化对照即覆盖"零字段丢失"）。

type wireMatrixRecord struct {
	Iface    string                     `json:"iface"`
	Case     string                     `json:"case"`
	Method   string                     `json:"method"`
	Path     string                     `json:"path"`
	Query    map[string]string          `json:"query"`
	Body     map[string]json.RawMessage `json:"body"`
	Status   int                        `json:"status"`
	ReqID    string                     `json:"req_id"`
	RespBody json.RawMessage            `json:"resp_body,omitempty"`
	Err      string                     `json:"err,omitempty"`
}

func TestWireMatrixDump(t *testing.T) {
	required := []struct {
		name     string
		fallback string
	}{
		{"WIRE_MATRIX_OUT_DIR", "L3_OUT_DIR"},
		{"LOG_SERVICE_ENDPOINT", ""},
		{"LOG_SERVICE_AK", ""},
		{"LOG_SERVICE_SK", ""},
		{"LOG_SERVICE_REGION", ""},
		{"WIRE_PROJECT_ID", "L3_PROJECT_ID"},
		{"WIRE_TOPIC_ID", "L3_TOPIC_ID"},
		{"WIRE_CONSUMER_GROUP", "L3_CONSUMER_GROUP"},
	}
	for _, item := range required {
		if firstNonEmptyEnv(item.name, item.fallback) == "" {
			t.Skipf("wire matrix env %s not set; skip", item.name)
		}
	}
	outDir := firstNonEmptyEnv("WIRE_MATRIX_OUT_DIR", "L3_OUT_DIR")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		t.Fatalf("mkdir out: %v", err)
	}

	cli := NewClient(firstNonEmptyEnv("LOG_SERVICE_ENDPOINT"), firstNonEmptyEnv("LOG_SERVICE_AK"),
		firstNonEmptyEnv("LOG_SERVICE_SK"), "", firstNonEmptyEnv("LOG_SERVICE_REGION")).(*LsClient)

	var (
		mu      sync.Mutex
		records = make([]wireMatrixRecord, 0, 5)
		current = "" // 当前接口名（注入 Iface 字段）
		base    = http.DefaultTransport
	)
	cli.Client.Transport = rtFunc(func(r *http.Request) (*http.Response, error) {
		method, path := r.Method, r.URL.Path
		query := map[string]string{}
		for k, v := range r.URL.Query() {
			if len(v) > 0 {
				query[k] = v[0]
			}
		}
		body := map[string]json.RawMessage{}
		if r.Body != nil {
			buf, _ := io.ReadAll(r.Body)
			_ = r.Body.Close()
			r.Body = io.NopCloser(bytes.NewReader(buf))
			r.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(buf)), nil }
			if len(buf) > 0 {
				_ = json.Unmarshal(buf, &body)
			}
		}
		resp, err := base.RoundTrip(r)
		mu.Lock()
		defer mu.Unlock()
		rec := wireMatrixRecord{Iface: current, Case: "default", Method: method, Path: path, Query: query, Body: body}
		if err != nil {
			rec.Err = err.Error()
		}
		if resp != nil {
			rec.Status = resp.StatusCode
			rec.ReqID = resp.Header.Get("X-Tls-Requestid")
			if rec.ReqID == "" {
				rec.ReqID = resp.Header.Get("X-Top-Requestid")
			}
			// 拷贝响应 body 后回填 resp.Body，保证业务层仍可读
			if resp.Body != nil {
				respBuf, _ := io.ReadAll(resp.Body)
				_ = resp.Body.Close()
				resp.Body = io.NopCloser(bytes.NewReader(respBuf))
				if json.Valid(respBuf) {
					rec.RespBody = json.RawMessage(respBuf)
				} else if len(respBuf) > 0 {
					// 非 JSON 响应（如 PB / 二进制）以 base64 包裹为字符串落盘
					b, _ := json.Marshal(string(respBuf))
					rec.RespBody = json.RawMessage(b)
				}
			}
		}
		records = append(records, rec)
		return resp, err
	})

	projectID := firstNonEmptyEnv("WIRE_PROJECT_ID", "L3_PROJECT_ID")
	topicID := firstNonEmptyEnv("WIRE_TOPIC_ID", "L3_TOPIC_ID")
	cgName := firstNonEmptyEnv("WIRE_CONSUMER_GROUP", "L3_CONSUMER_GROUP")

	// DescribeCursor
	current = "DescribeCursor"
	_, _ = cli.DescribeCursor(&DescribeCursorRequest{TopicID: topicID, ShardID: 0, From: "begin"})

	// DescribeCheckPoint
	current = "DescribeCheckPoint"
	_, _ = cli.DescribeCheckPoint(&DescribeCheckPointRequest{
		ProjectID: projectID, TopicID: topicID, ShardID: 0, ConsumerGroupName: cgName,
	})

	// SearchLogs
	current = "SearchLogs"
	accurate, mustComplete := true, true
	var offset int64 = 0
	_, _ = cli.SearchLogs(&SearchLogsRequest{
		TopicID:       topicID,
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

	// CreateIndex（首次 200，重跑 Conflict；wire 始终一致）
	current = "CreateIndex"
	enable := true
	maxLen := int32(2048)
	_, _ = cli.CreateIndex(&CreateIndexRequest{
		TopicID:           topicID,
		MaxTextLen:        &maxLen,
		EnableAutoIndex:   &enable,
		EnablePhraseIndex: &enable,
	})

	// CreateAlarm（伪 NotifyGroup 期望 4xx，仅比 wire）
	current = "CreateAlarm"
	status := true
	userMsg := "msg"
	severity := "warning"
	sendResolved := false
	_, _ = cli.CreateAlarm(&CreateAlarmRequest{
		AlarmName:        "l3-sdk-align-fake-alarm",
		ProjectID:        projectID,
		Status:           &status,
		QueryRequest:     QueryRequests{},
		RequestCycle:     RequestCycle{Type: "Period", Time: 5},
		Condition:        "x>1",
		TriggerPeriod:    1,
		AlarmPeriod:      5,
		AlarmNotifyGroup: []string{"g-fake"},
		UserDefineMsg:    &userMsg,
		Severity:         &severity,
		SendResolved:     &sendResolved,
	})

	// 写 jsonl，按 iface 字母序
	target := filepath.Join(outDir, "go.jsonl")
	fh, err := os.Create(target)
	if err != nil {
		t.Fatalf("create out: %v", err)
	}
	defer fh.Close()
	enc := json.NewEncoder(fh)
	enc.SetEscapeHTML(false)
	sort.SliceStable(records, func(i, j int) bool { return records[i].Iface < records[j].Iface })
	for _, rec := range records {
		if err := enc.Encode(rec); err != nil {
			t.Fatalf("encode: %v", err)
		}
	}
	t.Logf("wire matrix dump -> %s (%d records)", target, len(records))
}
