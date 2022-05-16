package tls

import (
	"os"
)

func NewClientWithEnv() Client {
	os.Setenv("LOG_SERVICE_ENDPOINT", "http://127.0.0.1:6080")
	os.Setenv("LOG_SERVICE_AK", "***REMOVED***")
	os.Setenv("LOG_SERVICE_SK", "***REMOVED***")
	os.Setenv("LOG_SERVICE_REGION", "cn-north-3")
	return NewClient(os.Getenv("LOG_SERVICE_ENDPOINT"), os.Getenv("LOG_SERVICE_AK"),
		os.Getenv("LOG_SERVICE_SK"), "", os.Getenv("LOG_SERVICE_REGION"))
}

func StrPtr(in string) *string {
	return &in
}

func IntPtr(in int) *int {
	return &in
}

func Uint16Ptr(in uint16) *uint16 {
	return &in
}

func CreateProject(projectName, description, region string, cli Client) (string, error) {
	createProjectReq := &CreateProjectRequest{
		ProjectName: projectName,
		Description: description,
		Region:      region,
	}

	createProjectResp, err := cli.CreateProject(createProjectReq)

	if err != nil {
		return "", err
	}

	return createProjectResp.ProjectId, nil
}

func CreateTopic(projectId, topicName, description string, shardCount int, ttl uint16, cli Client) (string, error) {
	createTopicReq := &CreateTopicRequest{
		ProjectID:   projectId,
		TopicName:   topicName,
		Ttl:         ttl,
		Description: description,
		ShardCount:  shardCount,
	}

	createTopicResp, err := cli.CreateTopic(createTopicReq)
	if err != nil {
		return "", err
	}

	return createTopicResp.TopicID, nil
}

func CreateIndex(topicID string, fulltextInfo *FullTextInfo, KeyValue *[]KeyValueInfo, cli Client) error {
	createIndexReq := &CreateIndexRequest{
		TopicID:  topicID,
		FullText: fulltextInfo,
		KeyValue: KeyValue,
	}
	_, err := cli.CreateIndex(createIndexReq)

	return err
}
