package tls

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestPutLogsRequestValidationLogGroupCount(t *testing.T) {
	makeGroup := func(count int) *pb.LogGroup {
		logs := make([]*pb.Log, 0, count)
		for i := 0; i < count; i++ {
			logs = append(logs, &pb.Log{})
		}
		return &pb.LogGroup{Logs: logs}
	}

	request := &PutLogsRequest{
		TopicID: "topic",
		LogBody: &pb.LogGroupList{LogGroups: []*pb.LogGroup{makeGroup(10000)}},
	}
	if err := request.CheckValidation(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	request = &PutLogsRequest{
		TopicID: "topic",
		LogBody: &pb.LogGroupList{LogGroups: []*pb.LogGroup{makeGroup(10001)}},
	}
	if err := request.CheckValidation(); err == nil {
		t.Fatalf("expected error for log group size > 10000")
	}
}
