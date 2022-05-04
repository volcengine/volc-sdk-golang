package test

import (
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"os"
	"testing"
)

func TestListShard(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	listRequest := &DescribeShardsRequest{
		TopicID:    "topic-id",
		PageNumber: 1,
		PageSize:   5,
	}

	resp, err := client.DescribeShards(listRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(resp.Shards) == 0 {
		t.Error("empty result")
	}
}
