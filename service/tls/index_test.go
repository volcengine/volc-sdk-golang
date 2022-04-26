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
	suite.NoError(suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic}))
	suite.NoError(suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project}))
}

func TestSDKIndexTestSuite(t *testing.T) {
	suite.Run(t, new(SDKIndexTestSuite))
}

// TestCreateIndex: test create index
func (suite *SDKIndexTestSuite) TestCreateIndex() {
	testcases := map[*CreateIndexRequest]*GetIndexResponse{
		{
			TopicID:        suite.topic,
			FulltextIndex:  true,
			CasSensitive:   false,
			IncludeChinese: false,
			Delimiter:      ",",
			KeyValueIndex:  true,
			KeyValueList: KeyValueList{
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
			TopicID:        suite.topic,
			FulltextIndex:  true,
			CasSensitive:   false,
			IncludeChinese: false,
			Delimiter:      ",",
			KeyValueIndex:  true,
			KeyValueList: KeyValueList{
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

		actualGetIndexResp, err := suite.cli.GetIndex(&GetIndexRequest{TopicID: suite.topic})
		suite.NoError(err)

		suite.Equal(expectGetIndexResp.TopicID, actualGetIndexResp.TopicID)
		suite.Equal(expectGetIndexResp.FulltextIndex, actualGetIndexResp.FulltextIndex)
		suite.Equal(expectGetIndexResp.CasSensitive, actualGetIndexResp.CasSensitive)
		suite.Equal(expectGetIndexResp.IncludeChinese, actualGetIndexResp.IncludeChinese)
		suite.Equal(expectGetIndexResp.Delimiter, actualGetIndexResp.Delimiter)
		suite.Equal(expectGetIndexResp.KeyValueIndex, actualGetIndexResp.KeyValueIndex)

		actualKV := actualGetIndexResp.KeyValueList
		expectKV := expectGetIndexResp.KeyValueList

		for i, kv := range actualKV {
			suite.Equal(expectKV[i].Key, kv.Key)
			suite.Equal(expectKV[i].Value.ValueType, kv.Value.ValueType)
			suite.Equal(expectKV[i].Value.IncludeChinese, kv.Value.IncludeChinese)
			suite.Equal(expectKV[i].Value.Delimiter, kv.Value.Delimiter)
			suite.Equal(expectKV[i].Value.SQLFlag, kv.Value.SQLFlag)
			suite.Equal(expectKV[i].Value.CasSensitive, kv.Value.CasSensitive)
		}

		err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
		suite.NoError(err)
	}
}

// TestUpdateIndex: test update index
func (suite *SDKIndexTestSuite) TestUpdateIndex() {
	createIndexReq := &CreateIndexRequest{
		TopicID:        suite.topic,
		FulltextIndex:  true,
		CasSensitive:   false,
		IncludeChinese: false,
		Delimiter:      ", ?",
		KeyValueIndex:  false,
		KeyValueList:   nil,
	}

	_, err := suite.cli.CreateIndex(createIndexReq)
	suite.NoError(err)

	testcases := map[*UpdateIndexRequest]*GetIndexResponse{
		{
			TopicID:        suite.topic,
			FulltextIndex:  true,
			CasSensitive:   false,
			IncludeChinese: false,
			Delimiter:      ",",
			KeyValueIndex:  true,
			KeyValueList: KeyValueList{
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
			TopicID:        suite.topic,
			FulltextIndex:  true,
			CasSensitive:   false,
			IncludeChinese: false,
			Delimiter:      ",",
			KeyValueIndex:  true,
			KeyValueList: KeyValueList{
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
		err := suite.cli.UpdateIndex(updateIndexReq)
		suite.NoError(err)

		actualGetIndexResp, err := suite.cli.GetIndex(&GetIndexRequest{TopicID: suite.topic})
		suite.NoError(err)

		suite.Equal(expectGetIndexResp.TopicID, actualGetIndexResp.TopicID)
		suite.Equal(expectGetIndexResp.FulltextIndex, actualGetIndexResp.FulltextIndex)
		suite.Equal(expectGetIndexResp.CasSensitive, actualGetIndexResp.CasSensitive)
		suite.Equal(expectGetIndexResp.IncludeChinese, actualGetIndexResp.IncludeChinese)
		suite.Equal(expectGetIndexResp.Delimiter, actualGetIndexResp.Delimiter)
		suite.Equal(expectGetIndexResp.KeyValueIndex, actualGetIndexResp.KeyValueIndex)

		actualKV := actualGetIndexResp.KeyValueList
		expectKV := expectGetIndexResp.KeyValueList

		for i, kv := range actualKV {
			suite.Equal(expectKV[i].Key, kv.Key)
			suite.Equal(expectKV[i].Value.ValueType, kv.Value.ValueType)
			suite.Equal(expectKV[i].Value.IncludeChinese, kv.Value.IncludeChinese)
			suite.Equal(expectKV[i].Value.Delimiter, kv.Value.Delimiter)
			suite.Equal(expectKV[i].Value.SQLFlag, kv.Value.SQLFlag)
			suite.Equal(expectKV[i].Value.CasSensitive, kv.Value.CasSensitive)
		}

		err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
		suite.NoError(err)
	}
}
