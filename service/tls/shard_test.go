package tls

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func NewShardClientWithEnv() Client {
	return NewClient(os.Getenv("VOLC_ENDPOINT"), os.Getenv("VOLC_ACCESSKEY"), os.Getenv("VOLC_SECRETKEY"), os.Getenv("VOLC_SESSIONTOKEN"), os.Getenv("VOLC_REGION"))
}

func TestManualShardSplit(t *testing.T) {
	client := NewShardClientWithEnv()

	project, err := client.CreateProject(&CreateProjectRequest{
		ProjectName: fmt.Sprintf("test-project-%s", uuid.New().String()),
		Region:      os.Getenv("VOLC_REGION"),
	})

	assert.NoError(t, err)

	topic, err := client.CreateTopic(&CreateTopicRequest{
		ProjectID:  project.ProjectID,
		TopicName:  fmt.Sprintf("test-topic-%s", uuid.New().String()),
		Ttl:        1,
		ShardCount: 2,
	})

	assert.NoError(t, err)

	describeShardsResponse, err := client.DescribeShards(&DescribeShardsRequest{
		TopicID: topic.TopicID,
	})

	assert.NoError(t, err)
	assert.Len(t, describeShardsResponse.Shards, 2)

	manualShardSplitResponse, err := client.ManualShardSplit(context.Background(), &ManualShardSplitRequest{
		TopicID: topic.TopicID,
		ShardID: int(describeShardsResponse.Shards[0].ShardID),
		Number:  2,
	})

	assert.NoError(t, err)
	assert.Len(t, manualShardSplitResponse.Shards, 2)

	_, err = client.DeleteTopic(&DeleteTopicRequest{
		TopicID: topic.TopicID,
	})
	assert.NoError(t, err)

	_, err = client.DeleteProject(&DeleteProjectRequest{
		ProjectID: project.ProjectID,
	})

	assert.NoError(t, err)

}

func TestManualShardSplitRequestCheckValidation_Number(t *testing.T) {
	testcases := []struct {
		name    string
		number  int
		wantErr bool
	}{
		{"number-2", 2, false},
		{"number-4", 4, false},
		{"number-8", 8, false},
		{"number-16", 16, false},
		{"number-0", 0, true},
		{"number-3", 3, true},
		{"number-5", 5, true},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			req := &ManualShardSplitRequest{
				TopicID: "topic-id",
				ShardID: 1,
				Number:  tc.number,
			}
			err := req.CheckValidation()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
