package tls

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestConsumeLogsProcessorWireContract(t *testing.T) {
	logList := &pb.LogGroupList{LogGroups: []*pb.LogGroup{{Logs: []*pb.Log{{Contents: []*pb.LogContent{{Key: "key", Value: "value"}}}}}}}
	payload, err := proto.Marshal(logList)
	if err != nil {
		t.Fatal(err)
	}
	endCursor, compression, group, consumer := "end", CompressNone, "group", "consumer"
	count := 12
	cases := []struct {
		name, processorID, path string
		original, conflict      bool
	}{
		{name: "default", path: PathConsumeLogs},
		{name: "original", original: true, path: PathConsumeOriginalLogs},
		{name: "processor without local UUID validation", processorID: "not-a-uuid", path: PathConsumeLogs},
		{name: "processor and original conflict", processorID: "processor", original: true, conflict: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			client := newTestClient(rtFunc(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != http.MethodGet || req.URL.Path != tc.path {
					t.Fatalf("method/path = %s %s, want GET %s", req.Method, req.URL.Path, tc.path)
				}
				if got, want := req.URL.Query(), map[string][]string{"TopicId": {"topic"}, "ShardId": {"7"}}; !reflect.DeepEqual(map[string][]string(got), want) {
					t.Fatalf("query = %v, want %v", got, want)
				}
				if req.Header.Get("ConsumerGroupName") != group || req.Header.Get("ConsumerName") != consumer || req.Header.Get("Content-Type") != "application/json" || req.Header.Get("X-Custom") != "test" || req.Header.Get("ProcessorId") != "" {
					t.Fatalf("headers = %v", req.Header)
				}
				var body map[string]interface{}
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Fatal(err)
				}
				want := map[string]interface{}{"Cursor": "cursor", "EndCursor": endCursor, "LogGroupCount": float64(count), "Compression": compression}
				if tc.processorID != "" {
					want["ProcessorId"] = tc.processorID
				}
				if !reflect.DeepEqual(body, want) {
					t.Fatalf("body = %#v, want %#v", body, want)
				}
				headers := make(http.Header)
				headers.Set(tlsCountHeader, "1")
				headers.Set(tlsCursorHeader, "next")
				headers.Set(tlsOriginalHeader, "false")
				headers.Set("x-tls-bodyrawsize", strconv.Itoa(len(payload)))
				headers.Set(RequestIDHeader, "request-id")
				return &http.Response{StatusCode: http.StatusOK, Header: headers, Body: io.NopCloser(bytes.NewReader(payload)), Request: req}, nil
			}))
			request := &ConsumeLogsRequest{
				CommonRequest: CommonRequest{Headers: map[string]string{"X-Custom": "test"}},
				TopicID:       "topic", ShardID: 7, Cursor: "cursor", EndCursor: &endCursor,
				LogGroupCount: &count, Compression: &compression, ConsumerGroupName: &group, ConsumerName: &consumer,
				Original: tc.original, ProcessorID: tc.processorID,
			}
			response, err := client.ConsumeLogs(request)
			if tc.conflict {
				clientErr, ok := err.(*Error)
				if !ok || clientErr.Code != "ClientError" || !strings.Contains(clientErr.Message, "Invalid argument") || !strings.Contains(clientErr.Message, "ProcessorID") {
					t.Fatalf("error = %#v, want local Invalid argument error mentioning ProcessorID", err)
				}
				if response != nil || calls != 0 {
					t.Fatalf("conflict response = %#v, HTTP calls = %d", response, calls)
				}
				return
			}
			if err != nil {
				t.Fatalf("ConsumeLogs: %v", err)
			}
			if calls != 1 || response.Count != 1 || response.Cursor != "next" || response.RequestID != "request-id" || !proto.Equal(response.Logs, logList) {
				t.Fatalf("response = %#v, calls = %d", response, calls)
			}
		})
	}
}

func TestConsumeLogsProcessorMinimalRequest(t *testing.T) {
	client := newTestClient(rtFunc(func(req *http.Request) (*http.Response, error) {
		var body map[string]interface{}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, map[string]interface{}{"Cursor": "cursor"}) {
			t.Fatalf("minimal body = %#v", body)
		}
		if req.Header.Get("ConsumerGroupName") != "" || req.Header.Get("ConsumerName") != "" {
			t.Fatalf("unexpected consumer headers: %v", req.Header)
		}
		headers := make(http.Header)
		headers.Set(tlsCountHeader, "0")
		headers.Set(tlsCursorHeader, "next")
		return &http.Response{StatusCode: http.StatusOK, Header: headers, Body: io.NopCloser(strings.NewReader("")), Request: req}, nil
	}))
	response, err := client.ConsumeLogs(&ConsumeLogsRequest{TopicID: "topic", Cursor: "cursor"})
	if err != nil || response.Count != 0 || response.Cursor != "next" {
		t.Fatalf("response = %#v, error = %v", response, err)
	}
}

func TestConsumeLogsProcessorServerError(t *testing.T) {
	calls := 0
	client := newTestClient(rtFunc(func(req *http.Request) (*http.Response, error) {
		calls++
		headers := make(http.Header)
		headers.Set(RequestIDHeader, "server-request-id")
		return &http.Response{StatusCode: http.StatusBadRequest, Header: headers, Body: io.NopCloser(strings.NewReader(`{"errorCode":"InvalidArgument","errorMessage":"invalid processor"}`)), Request: req}, nil
	}))
	response, err := client.ConsumeLogs(&ConsumeLogsRequest{TopicID: "topic", Cursor: "cursor", ProcessorID: "not-a-uuid"})
	clientErr, ok := err.(*Error)
	if !ok || clientErr.HTTPCode != http.StatusBadRequest || clientErr.Code != ErrInvalidParam || clientErr.Message != "invalid processor" || clientErr.RequestID != "server-request-id" {
		t.Fatalf("server error = %#v", err)
	}
	if response != nil || calls != 1 {
		t.Fatalf("response = %#v, calls = %d", response, calls)
	}
}
