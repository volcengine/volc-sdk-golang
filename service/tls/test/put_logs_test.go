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

	// kv index log
	kvLogGroupList := &pb.LogGroupList{
		LogGroups: []*pb.LogGroup{
			{
				Logs: []*pb.Log{
					{
						Contents: []*pb.LogContent{
							{
								Key:   "test-key1",
								Value: "test-value",
							},
							{
								Key:   "test-key2",
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

	kvPutLogsRequest := &PutLogsRequest{
		TopicID:      "topic-id",
		HashKey:      "7fffff",
		CompressType: CompressLz4,
		LogBody:      kvLogGroupList,
	}

	_, err := client.PutLogs(kvPutLogsRequest)
	if err != nil {
		t.Error(err.Error())
	}

	// fulltext index log

	fulltextLogGroupList := &pb.LogGroupList{
		LogGroups: []*pb.LogGroup{
			{
				Logs: []*pb.Log{
					{
						Contents: []*pb.LogContent{
							{
								Key:   FullTextIndexKey,
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

	fulltextPutLogsRequest := &PutLogsRequest{
		TopicID:      "topic-id",
		HashKey:      "7fffff",
		CompressType: CompressLz4,
		LogBody:      fulltextLogGroupList,
	}

	_, err = client.PutLogs(fulltextPutLogsRequest)
	if err != nil {
		t.Error(err.Error())
	}

}
