package test

import (
	"os"
	"testing"
	"time"

	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func TestPutLogs(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	logGroupList := &pb.LogGroupList{
		LogGroups: []*pb.LogGroup{
			{
				Logs: []*pb.Log{
					{
						Contents: []*pb.LogContent{
							{
								Key:   "test-key",
								Value: "test-value",
							},
						},
						Time: time.Now().Unix(),
					},
				},
				Source:   "127.0.0.1",
				FileName: "testFile",
			},
		},
	}

	putLogsRequest := &PutLogsRequest{
		TopicID:      "topic-id",
		HashKey:      "7fffff",
		CompressType: "lz4",
		LogBody:      logGroupList,
	}

	_, err := client.PutLogs(putLogsRequest)
	if err != nil {
		t.Error(err.Error())
	}
}
