package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKKafkaTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKKafkaTestSuite) SetupTest() {
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

func (suite *SDKKafkaTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKKafkaTestSuite(t *testing.T) {
	suite.Run(t, new(SDKKafkaTestSuite))
}

// TestOpenKafkaConsumer
func (suite *SDKKafkaTestSuite) TestKafkaConsumer() {
	{
		testcases := map[*OpenKafkaConsumerRequest]*OpenKafkaConsumerResponse{
			{
				TopicID: suite.topic,
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.OpenKafkaConsumer(req)
			suite.NoError(err)
		}
	}

	{
		testcases := map[*DescribeKafkaConsumerRequest]*DescribeKafkaConsumerResponse{
			{
				TopicID: suite.topic,
			}: {
				AllowConsume: true,
			},
		}

		for req, resp := range testcases {
			actualResp, err := suite.cli.DescribeKafkaConsumer(req)
			suite.NoError(err)
			suite.Equal(actualResp.AllowConsume, resp.AllowConsume)
		}
	}

	{
		testcases := map[*CloseKafkaConsumerRequest]*CloseKafkaConsumerResponse{
			{
				TopicID: suite.topic,
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.CloseKafkaConsumer(req)
			suite.NoError(err)
		}
	}

	{
		testcases := map[*OpenKafkaConsumerRequest]*OpenKafkaConsumerResponse{
			{
				TopicID: "Test",
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.OpenKafkaConsumer(req)
			suite.Error(err)
		}
	}

	{
		testcases := map[*DescribeKafkaConsumerRequest]*DescribeKafkaConsumerResponse{
			{
				TopicID: "Test",
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.DescribeKafkaConsumer(req)
			suite.Error(err)
		}
	}

	{
		testcases := map[*CloseKafkaConsumerRequest]*CloseKafkaConsumerResponse{
			{
				TopicID: "Test",
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.CloseKafkaConsumer(req)
			suite.Error(err)
		}
	}
}
