package tls

import (
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

	projectId, err := CreateProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-index-"+uuid.New().String(),
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

func TestSDKShardTestSuite(t *testing.T) {
	suite.Run(t, new(SDKShardTestSuite))
}

// TestCreateIndex: test create index
func (suite *SDKShardTestSuite) TestDescribeShards() {

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
		suite.Equal(len(listResp.Shards), expect)
	}
}
