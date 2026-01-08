package tls

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKShardTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKShardTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-project-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-topic-"+uuid.New().String(),
		"test", 4, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId
}

func (suite *SDKShardTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKShardTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKShardTestSuite(t *testing.T) {
	suite.Run(t, new(SDKShardTestSuite))
}

func (suite *SDKShardTestSuite) TestDescribeShardsNormally() {
	testcases := map[*DescribeShardsRequest]int{
		{
			TopicID: suite.topic,
		}: 4,
		{
			TopicID:  suite.topic,
			PageSize: 2,
		}: 2,
		{
			TopicID:    suite.topic,
			PageNumber: 2,
			PageSize:   2,
		}: 2,
	}

	for req, expect := range testcases {
		listResp, err := suite.cli.DescribeShards(req)
		suite.NoError(err)
		suite.Equal(expect, len(listResp.Shards))
	}
}

func (suite *SDKShardTestSuite) TestDescribeShardsAbnormally() {
	testcases := map[*DescribeShardsRequest]*Error{
		{
			TopicID: uuid.New().String(),
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "TopicNotExist",
			Message:  "Topic does not exist.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeShards(req)
		suite.validateError(err, expectedErr)
	}
}
