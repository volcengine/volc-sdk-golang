package test

import (
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"os"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	createRequest := &CreateIndexRequest{
		TopicID:        "parent-topic-id",
		FulltextIndex:  true,
		CasSensitive:   false,
		IncludeChinese: false,
		Delimiter:      ", ?",
		KeyValueIndex:  true,
		KeyValueList: KeyValueList{
			{
				Key: "test-key",
				Value: Value{
					ValueType:      "text",
					Delimiter:      ", ?",
					CasSensitive:   false,
					IncludeChinese: false,
					SQLFlag:        true,
				},
			},
		},
	}

	resp, err := client.CreateIndex(createRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if resp.TopicID == "" {
		t.Error("empty topic id")
	}
}

func TestDeleteIndex(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	deleteRequest := &DeleteIndexRequest{
		TopicID: "topic-id",
	}

	_, err := client.DeleteIndex(deleteRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestUpdateIndex(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	updateRequest := &UpdateIndexRequest{
		TopicID:        "parent-topic-id",
		FulltextIndex:  true,
		CasSensitive:   false,
		IncludeChinese: false,
		Delimiter:      ", ?",
		KeyValueIndex:  true,
		KeyValueList: KeyValueList{
			{
				Key: "test-key",
				Value: Value{
					ValueType:      "text",
					Delimiter:      ", ?",
					CasSensitive:   false,
					IncludeChinese: false,
					SQLFlag:        true,
				},
			},
		},
	}

	_, updateIndexErr := client.UpdateIndex(updateRequest)
	if updateIndexErr != nil {
		t.Error(updateIndexErr.Error())
		return
	}
}

func TestGetIndex(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	getRequest := &GetIndexRequest{
		TopicID: "topic-id",
	}

	indexInfo, err := client.GetIndex(getRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if indexInfo.TopicID == "" {
		t.Error("empty topic id")
	}
}

func TestSearchIndex(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	searchRequest := &SearchIndexRequest{
		TopicID:   "topic-id",
		Query:     "*",
		StartTime: 1346457600,
		EndTime:   1730454400,
		Limit:     100,
		HighLight: false,
		Context:   "",
		Sort:      "asc",
	}

	searchResult, err := client.SearchIndex(searchRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(searchResult.Logs) == 0 {
		t.Error("empty result")
	}
}
