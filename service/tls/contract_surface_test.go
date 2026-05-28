package tls

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestClientInterfaceIncludesOfficialContractGaps(t *testing.T) {
	assertOfficialContractGapClientSurface(nil)
}

func assertOfficialContractGapClientSurface(client Client) {
	if client == nil {
		return
	}
	_, _ = client.ModifyTraceInstance(&ModifyTraceInstanceRequest{})
	_, _ = client.CreateLogBackFlowTask(&CreateLogBackFlowTaskRequest{})
	_, _ = client.DeleteLogBackFlowTask(&DeleteLogBackFlowTaskRequest{})
	_, _ = client.DescribeLogBackFlowTasks(&DescribeLogBackFlowTasksRequest{})
	_, _ = client.ModifyLogBackFlowTask(&ModifyLogBackFlowTaskRequest{})
	_, _ = client.DescribeCursorTime(&DescribeCursorTimeRequest{})
}

func TestContractGapFieldsSerializeAndDeserialize(t *testing.T) {
	mustComplete := true
	downloadBody, err := json.Marshal(&CreateDownloadTaskRequest{
		TopicID:      "topic-id",
		TaskName:     "task-name",
		Query:        "*",
		StartTime:    1,
		EndTime:      2,
		Compression:  "gzip",
		DataFormat:   "json",
		Limit:        10,
		Sort:         "asc",
		TaskType:     1,
		MustComplete: &mustComplete,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !json.Valid(downloadBody) || !containsJSONKey(downloadBody, "MustComplete") {
		t.Fatalf("CreateDownloadTaskRequest missing MustComplete in json: %s", string(downloadBody))
	}

	traceBody := []byte(`{"TraceInstanceId":"tid","BackendConfig":{"Ttl":30},"CsAccountChannel":"cs"}`)
	var traceResp DescribeTraceInstanceResponse
	if err := json.Unmarshal(traceBody, &traceResp); err != nil {
		t.Fatal(err)
	}
	if traceResp.BackendConfig == nil || traceResp.BackendConfig.TTL != 30 {
		t.Fatalf("DescribeTraceInstanceResponse missing BackendConfig: %#v", traceResp.BackendConfig)
	}
	if traceResp.CsAccountChannel != "cs" {
		t.Fatalf("DescribeTraceInstanceResponse missing CsAccountChannel: %q", traceResp.CsAccountChannel)
	}

	allowEdit := true
	allowDelete := false
	ruleBody, err := json.Marshal(&DescribeRuleResponseV2{
		CsAccountChannel: "cs",
		AllowEdit:        &allowEdit,
		AllowDelete:      &allowDelete,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"CsAccountChannel", "AllowEdit", "AllowDelete"} {
		if !containsJSONKey(ruleBody, key) {
			t.Fatalf("DescribeRuleResponseV2 missing %s in json: %s", key, string(ruleBody))
		}
	}
}

func TestDescribeCursorTimeRequiresOfficialQueryFields(t *testing.T) {
	shardID := 0
	tests := []struct {
		name    string
		request *DescribeCursorTimeRequest
		wantErr bool
	}{
		{
			name:    "missing topic id",
			request: &DescribeCursorTimeRequest{ShardID: &shardID, Cursor: "cursor"},
			wantErr: true,
		},
		{
			name:    "missing shard id",
			request: &DescribeCursorTimeRequest{TopicID: "topic-id", Cursor: "cursor"},
			wantErr: true,
		},
		{
			name:    "missing cursor",
			request: &DescribeCursorTimeRequest{TopicID: "topic-id", ShardID: &shardID},
			wantErr: true,
		},
		{
			name:    "all required fields",
			request: &DescribeCursorTimeRequest{TopicID: "topic-id", ShardID: &shardID, Cursor: "cursor"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.CheckValidation()
			if (err != nil) != tt.wantErr {
				t.Fatalf("CheckValidation err=%v wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

func TestDescribeLogBackFlowTasksPreservesRepeatedTopicIDList(t *testing.T) {
	cap := &captured{}
	c := newWireClient(t, cap)

	_, err := c.DescribeLogBackFlowTasks(&DescribeLogBackFlowTasksRequest{
		TopicIDList: []string{"topic-a", "topic-b"},
	})
	if err != nil {
		t.Fatalf("DescribeLogBackFlowTasks returned err: %v", err)
	}
	if cap.method != http.MethodGet {
		t.Fatalf("method got %s want GET", cap.method)
	}
	got := cap.queryValues["TopicIDList"]
	want := []string{"topic-a", "topic-b"}
	if len(got) != len(want) {
		t.Fatalf("TopicIDList values got=%v want=%v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("TopicIDList values got=%v want=%v", got, want)
		}
	}
}

func containsJSONKey(data []byte, key string) bool {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return false
	}
	_, ok := m[key]
	return ok
}
