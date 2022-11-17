package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKWebtracksTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKWebtracksTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-index-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	keyValueList := make([]KeyValueInfo, 0)
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-1",
		Value: Value{
			ValueType:      "text",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-2",
		Value: Value{
			ValueType:      "long",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	suite.NoError(CreateIndex(topicId, nil, &keyValueList, suite.cli))
}

func (suite *SDKWebtracksTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKWebtracksTestSuite(t *testing.T) {
	suite.Run(t, new(SDKWebtracksTestSuite))
}

// TestWebtracks
func (suite *SDKWebtracksTestSuite) TestWebtracks() {
	{
		testcases := map[*WebTracksRequest]*WebTracksResponse{
			{
				TopicID:      suite.topic,
				ProjectID:    suite.project,
				CompressType: "lz4",
				Source:       "sdk-test",
				Logs: []map[string]string{
					{
						"tap1": "person-A",
					},
				},
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.WebTracks(req)
			suite.NoError(err)
		}
	}

	// Test invalid
	{
		testcases := map[*WebTracksRequest]*WebTracksResponse{
			{
				TopicID: "Test",
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.WebTracks(req)
			suite.Error(err)
		}
	}
}
