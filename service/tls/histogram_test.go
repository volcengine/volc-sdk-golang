package tls

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKHistogramTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKHistogramTestSuite) SetupTest() {
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

func (suite *SDKHistogramTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKHistogramTestSuite(t *testing.T) {
	suite.Run(t, new(SDKHistogramTestSuite))
}

func (suite *SDKHistogramTestSuite) TestHistogram() {
	{
		testcases := map[*DescribeHistogramRequest]*DescribeHistogramResponse{
			{
				TopicID:   suite.topic,
				Query:     "*",
				StartTime: time.Now().UnixNano()/1e9 - 10,
				EndTime:   time.Now().UnixNano() / 1e9,
			}: {
				TotalCount: 0,
			},
		}

		for req, resp := range testcases {
			actualResp, err := suite.cli.DescribeHistogram(req)
			suite.NoError(err)
			suite.Equal(resp.TotalCount, actualResp.TotalCount)
		}
	}

	// Test Invalid
	{
		testcases := map[*DescribeHistogramRequest]*DescribeHistogramResponse{
			{
				TopicID:   "Test",
				Query:     "*",
				StartTime: time.Now().UnixNano()/1e9 - 10,
				EndTime:   time.Now().UnixNano() / 1e9,
			}: {
				TotalCount: 0,
			},
		}

		for req, _ := range testcases {
			_, err := suite.cli.DescribeHistogram(req)
			suite.Error(err)
		}
	}
}
