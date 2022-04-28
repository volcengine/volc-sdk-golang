package test

import (
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"os"
	"testing"
)

func TestCreateTopic(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	createRequest := &CreateTopicRequest{
		ProjectID:   "parent-project-id",
		TopicName:   "topic1",
		Ttl:         1024,
		Description: "test topic",
	}

	resp, err := client.CreateTopic(createRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if resp.TopicID == "" {
		t.Error("empty topic id")
	}
}

func TestDeleteTopic(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	deleteRequest := &DeleteTopicRequest{
		TopicID: "topic-id",
	}

	_, err := client.DeleteTopic(deleteRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestUpdateTopic(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	newTopicName := "new-topic-name"
	newDescription := "new-description"
	updateRequest := &UpdateTopicRequest{
		TopicID:     "topic-id",
		TopicName:   &newTopicName,
		Description: &newDescription,
	}

	_, err := client.UpdateTopic(updateRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestGetTopic(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	getRequest := &GetTopicRequest{
		TopicID: "topic-id",
	}

	topicInfo, err := client.GetTopic(getRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if topicInfo.TopicID == "" {
		t.Error("empty topic id")
	}
}

func TestListTopic(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	listRequest := &ListTopicRequest{
		ProjectID: "parent project id",
		TopicID:   "topic-id",
		TopicName: "topic-name",
		Offset:    5,
		Limit:     5,
	}

	topicList, err := client.ListTopic(listRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(topicList.Topics) == 0 {
		t.Error("empty result")
	}
}
