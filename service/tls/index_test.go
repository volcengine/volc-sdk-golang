package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKIndexTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKIndexTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-index-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId
}

func (suite *SDKIndexTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKIndexTestSuite(t *testing.T) {
	suite.Run(t, new(SDKIndexTestSuite))
}

// TestCreateIndex: test create index
func (suite *SDKIndexTestSuite) TestCreateIndex() {
	testcases := map[*CreateIndexRequest]*DescribeIndexResponse{
		{
			TopicID: suite.topic,
			FullText: &FullTextInfo{
				CaseSensitive:  false,
				IncludeChinese: false,
				Delimiter:      ",",
			},
			KeyValue: &[]KeyValueInfo{
				{
					Key: "test-key-1",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
				{
					Key: "test-key-2",
					Value: Value{
						ValueType:      "long",
						Delimiter:      "",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
			},
		}: {
			TopicID: suite.topic,
			FullText: &FullTextInfo{
				CaseSensitive:  false,
				IncludeChinese: false,
				Delimiter:      ",",
			},
			KeyValue: &[]KeyValueInfo{
				{
					Key: "test-key-1",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
				{
					Key: "test-key-2",
					Value: Value{
						ValueType:      "long",
						Delimiter:      "",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
			},
		},
	}

	for createIndexReq, expectGetIndexResp := range testcases {
		_, err := suite.cli.CreateIndex(createIndexReq)
		suite.NoError(err)

		actualGetIndexResp, err := suite.cli.DescribeIndex(&DescribeIndexRequest{TopicID: suite.topic})
		suite.NoError(err)

		suite.Equal(expectGetIndexResp.TopicID, actualGetIndexResp.TopicID)
		suite.Equal(expectGetIndexResp.FullText, actualGetIndexResp.FullText)

		if actualGetIndexResp.FullText != nil {
			suite.Equal(expectGetIndexResp.FullText.CaseSensitive, actualGetIndexResp.FullText.CaseSensitive)
			suite.Equal(expectGetIndexResp.FullText.IncludeChinese, actualGetIndexResp.FullText.IncludeChinese)
			suite.Equal(expectGetIndexResp.FullText.Delimiter, actualGetIndexResp.FullText.Delimiter)
		}
		suite.Equal(expectGetIndexResp.KeyValue, actualGetIndexResp.KeyValue)

		_, err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
		suite.NoError(err)
	}
}

// TestUpdateIndex: test update index
func (suite *SDKIndexTestSuite) TestUpdateIndex() {
	createIndexReq := &CreateIndexRequest{
		TopicID: suite.topic,
		FullText: &FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
		KeyValue: nil,
	}

	_, err := suite.cli.CreateIndex(createIndexReq)
	suite.NoError(err)

	testcases := map[*ModifyIndexRequest]*DescribeIndexResponse{
		{
			TopicID: suite.topic,

			FullText: &FullTextInfo{
				CaseSensitive:  false,
				IncludeChinese: false,
				Delimiter:      ",",
			},
			KeyValue: &[]KeyValueInfo{
				{
					Key: "test-key-1",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
				{
					Key: "test-key-2",
					Value: Value{
						ValueType:      "long",
						Delimiter:      "",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
			},
		}: {
			TopicID: suite.topic,
			FullText: &FullTextInfo{
				CaseSensitive:  false,
				IncludeChinese: false,
				Delimiter:      ",",
			},
			KeyValue: &[]KeyValueInfo{
				{
					Key: "test-key-1",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
				{
					Key: "test-key-2",
					Value: Value{
						ValueType:      "long",
						Delimiter:      "",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
			},
		},
	}

	for updateIndexReq, expectGetIndexResp := range testcases {
		_, err := suite.cli.ModifyIndex(updateIndexReq)
		suite.NoError(err)

		actualGetIndexResp, err := suite.cli.DescribeIndex(&DescribeIndexRequest{TopicID: suite.topic})
		suite.NoError(err)

		suite.Equal(expectGetIndexResp.TopicID, actualGetIndexResp.TopicID)
		suite.Equal(expectGetIndexResp.FullText, actualGetIndexResp.FullText)
		if actualGetIndexResp != nil {
			suite.Equal(expectGetIndexResp.FullText.CaseSensitive, actualGetIndexResp.FullText.CaseSensitive)
			suite.Equal(expectGetIndexResp.FullText.IncludeChinese, actualGetIndexResp.FullText.IncludeChinese)
			suite.Equal(expectGetIndexResp.FullText.Delimiter, actualGetIndexResp.FullText.Delimiter)
		}
		suite.Equal(expectGetIndexResp.KeyValue, actualGetIndexResp.KeyValue)

		_, err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
		suite.NoError(err)
	}
}
