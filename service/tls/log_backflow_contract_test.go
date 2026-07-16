package tls

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateLogBackFlowTaskUsesCurrentWireContract(t *testing.T) {
	req := &CreateLogBackFlowTaskRequest{
		TaskName:          "backflow-task",
		BackFlowStartTime: 1,
		LogBackFlowTaskSource: &LogBackFlowTaskSource{
			SourceType: "Topic",
			LogBackFlowTaskTopicSource: &LogBackFlowTaskTopicSource{
				ProjectID: "project-id",
				TopicID:   "source-topic-id",
			},
		},
		ETLTaskInfo: &LogBackFlowETLTaskInfo{
			Script: "f_set(\"key\", \"value\")",
			TargetResources: []TargetResource{{
				Alias: "target", TopicID: "target-topic-id", Region: "cn-beijing",
			}},
		},
	}

	if err := req.CheckValidation(); err != nil {
		t.Fatalf("current contract allows QueryParams to be omitted: %v", err)
	}
	body, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	if _, exists := got["ScheduleSqlTaskInfo"]; exists {
		t.Fatalf("retired ScheduleSqlTaskInfo must not be sent: %s", body)
	}
	if _, exists := got["ETLTaskInfo"]; !exists {
		t.Fatalf("ETLTaskInfo must be sent: %s", body)
	}
	if _, exists := got["QueryParams"]; exists {
		t.Fatalf("optional QueryParams must be omitted when unset: %s", body)
	}
}

func TestCreateLogBackFlowTaskRejectsRetiredScheduleSQL(t *testing.T) {
	req := validCreateLogBackFlowTaskRequest()
	req.ScheduleSqlTaskInfo = &LogBackFlowScheduleSqlTaskInfo{DestTopicID: "legacy-topic-id"}
	if err := req.CheckValidation(); err == nil {
		t.Fatal("expected retired ScheduleSqlTaskInfo to be rejected")
	}
	body, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	if containsJSONKey(body, "ScheduleSqlTaskInfo") {
		t.Fatalf("retired ScheduleSqlTaskInfo must never reach strict server parser: %s", body)
	}
}

func TestDescribeLogBackFlowTasksUsesCurrentQueryContract(t *testing.T) {
	status := LogBackFlowTaskStatusDone
	etlTaskID := "etl-task-id"
	cap := &captured{}
	client := newWireClient(t, cap)

	_, err := client.DescribeLogBackFlowTasks(&DescribeLogBackFlowTasksRequest{
		TopicIDList: []string{"topic-a", "topic-b"},
		Status:      &status,
		ETLTaskID:   &etlTaskID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if cap.method != http.MethodGet || cap.path != PathDescribeLogBackFlowTasks {
		t.Fatalf("unexpected wire target: %s %s", cap.method, cap.path)
	}
	if got := cap.queryValues["Status"]; len(got) != 1 || got[0] != LogBackFlowTaskStatusDone {
		t.Fatalf("Status query got %v", got)
	}
	if got := cap.queryValues["ETLTaskId"]; len(got) != 1 || got[0] != etlTaskID {
		t.Fatalf("ETLTaskId query got %v", got)
	}
	if _, exists := cap.queryValues["ScheduleSQLTaskId"]; exists {
		t.Fatalf("retired ScheduleSQLTaskId was sent: %v", cap.queryValues)
	}
}

func TestModifyLogBackFlowTaskValidatesCoupledFields(t *testing.T) {
	if err := (&ModifyLogBackFlowTaskRequest{
		TaskID:      "task-id",
		QueryParams: &LogBackFlowQueryParams{},
	}).CheckValidation(); err == nil {
		t.Fatal("QueryParams without ETLTaskInfo must be rejected")
	}
	if err := (&ModifyLogBackFlowTaskRequest{
		TaskID:                 "task-id",
		ShipperToTosInfo:       &LogBackFlowShipperToTosInfo{},
		ShipperToAgentLoopInfo: &LogBackFlowShipperToAgentLoopInfo{},
	}).CheckValidation(); err == nil {
		t.Fatal("TOS and AgentLoop shippers must be mutually exclusive")
	}
}

func TestDescribeLogBackFlowTasksDecodesCurrentAndLegacyTaskDetails(t *testing.T) {
	raw := []byte(`{
		"TaskId":"task-id",
		"ETLTaskInfo":{"Script":"f_set(\"k\",\"v\")","TargetResources":[{"Alias":"target","TopicId":"target-topic"}]},
		"ShipperToAgentLoopInfo":{"EvaluationSetShipperInfo":{"WorkspaceId":"workspace","FieldMappings":[{"Source":"message","Target":"input"}]},"ContentInfo":{"Format":"json"}},
		"ScheduleSqlTaskInfo":{"DestTopicID":"legacy-topic"},
		"RelaTasksInfo":{"ETLTaskId":"etl-id","ETLTaskName":"etl-name"}
	}`)
	var got LogBackFlowTaskInfo
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.ETLTaskInfo == nil || got.ETLTaskInfo.Script == "" {
		t.Fatalf("ETLTaskInfo not decoded: %#v", got.ETLTaskInfo)
	}
	if got.ShipperToAgentLoopInfo == nil || got.ShipperToAgentLoopInfo.EvaluationSetShipperInfo == nil {
		t.Fatalf("AgentLoop info not decoded: %#v", got.ShipperToAgentLoopInfo)
	}
	if got.ScheduleSqlTaskInfo == nil {
		t.Fatal("legacy ScheduleSqlTaskInfo response must remain decodable")
	}
	if got.RelaTasksInfo == nil || got.RelaTasksInfo.ETLTaskID != "etl-id" {
		t.Fatalf("ETL relation not decoded: %#v", got.RelaTasksInfo)
	}
}

func TestLogBackFlowQueryFilterKeepsLegacyAndMixedValues(t *testing.T) {
	legacy, err := json.Marshal(LogBackFlowQueryFilter{
		Field: "level", Operator: "IN", Values: []string{"error", "warn"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(legacy) != `{"Field":"level","Operator":"IN","Values":["error","warn"]}` {
		t.Fatalf("legacy Values wire changed: %s", legacy)
	}
	mixed, err := json.Marshal(LogBackFlowQueryFilter{
		Field: "status", Operator: "IN", ValuesAny: []interface{}{1, "error"},
	})
	if err != nil {
		t.Fatal(err)
	}
	var decoded LogBackFlowQueryFilter
	if err := json.Unmarshal(mixed, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded.ValuesAny) != 2 || decoded.ValuesAny[0].(float64) != 1 {
		t.Fatalf("mixed Values not preserved: %#v", decoded.ValuesAny)
	}
}

func validCreateLogBackFlowTaskRequest() *CreateLogBackFlowTaskRequest {
	return &CreateLogBackFlowTaskRequest{
		TaskName:          "backflow-task",
		BackFlowStartTime: 1,
		LogBackFlowTaskSource: &LogBackFlowTaskSource{
			SourceType: "Topic",
			LogBackFlowTaskTopicSource: &LogBackFlowTaskTopicSource{
				ProjectID: "project-id", TopicID: "source-topic-id",
			},
		},
		ETLTaskInfo: &LogBackFlowETLTaskInfo{
			Script: "f_set(\"key\", \"value\")",
			TargetResources: []TargetResource{{
				Alias: "target", TopicID: "target-topic-id",
			}},
		},
	}
}
