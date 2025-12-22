package tls

import (
	"net/http"
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

	projectId, err := CreateProject("golang-sdk-create-project-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-topic-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	keyValueList := make([]KeyValueInfo, 0)
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-1",
		Value: Value{
			ValueType:      "text",
			Delimiter:      "",
			CaseSensitive:  false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-2",
		Value: Value{
			ValueType:      "long",
			Delimiter:      "",
			CaseSensitive:  false,
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

func (suite *SDKKafkaTestSuite) validateError(err error, expectErr *Error) {
	sdkErr, ok := err.(*Error)

	if sdkErr == nil {
		suite.Nil(sdkErr)
		return
	}

	suite.Equal(true, ok)
	suite.Equal(expectErr.HTTPCode, sdkErr.HTTPCode)
	suite.Equal(expectErr.Code, sdkErr.Code)
	suite.Equal(expectErr.Message, sdkErr.Message)
}

func TestSDKKafkaTestSuite(t *testing.T) {
	suite.Run(t, new(SDKKafkaTestSuite))
}

func (suite *SDKKafkaTestSuite) TestOpenKafkaConsumerNormally() {
	_, err := suite.cli.OpenKafkaConsumer(&OpenKafkaConsumerRequest{TopicID: suite.topic})
	suite.NoError(err)
}

func (suite *SDKKafkaTestSuite) TestOpenKafkaConsumerAbnormally() {
	topicID := uuid.New().String()
	_, err := suite.cli.OpenKafkaConsumer(&OpenKafkaConsumerRequest{TopicID: topicID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "TopicNotExist",
		Message:  "Topic does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKKafkaTestSuite) TestCloseKafkaConsumerNormally() {
	_, err := suite.cli.OpenKafkaConsumer(&OpenKafkaConsumerRequest{TopicID: suite.topic})
	suite.NoError(err)
	_, err = suite.cli.CloseKafkaConsumer(&CloseKafkaConsumerRequest{TopicID: suite.topic})
	suite.NoError(err)
}

func (suite *SDKKafkaTestSuite) TestCloseKafkaConsumerAbnormally() {
	topicID := uuid.New().String()
	_, err := suite.cli.CloseKafkaConsumer(&CloseKafkaConsumerRequest{TopicID: topicID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "TopicNotExist",
		Message:  "Topic does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKKafkaTestSuite) TestDescribeKafkaConsumerNormally() {
	_, err := suite.cli.OpenKafkaConsumer(&OpenKafkaConsumerRequest{TopicID: suite.topic})
	suite.NoError(err)
	resp, err := suite.cli.DescribeKafkaConsumer(&DescribeKafkaConsumerRequest{TopicID: suite.topic})
	suite.NoError(err)
	suite.Equal(true, resp.AllowConsume)
}

func (suite *SDKKafkaTestSuite) TestDescribeKafkaConsumerAbnormally() {
	topicID := uuid.New().String()
	_, err := suite.cli.DescribeKafkaConsumer(&DescribeKafkaConsumerRequest{TopicID: topicID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "TopicNotExist",
		Message:  "Topic does not exist.",
	}
	suite.validateError(err, expectedErr)
}
