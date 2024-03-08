package tls

import (
	"net/http"
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

func (suite *SDKIndexTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKIndexTestSuite(t *testing.T) {
	suite.Run(t, new(SDKIndexTestSuite))
}

func (suite *SDKIndexTestSuite) TestCreateIndexNormally() {
	testcases := map[*CreateIndexRequest]*Error{
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
		}: nil,
		{
			TopicID: suite.topic,
			KeyValue: &[]KeyValueInfo{
				{
					Key: "test-key-1",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",",
						CasSensitive:   true,
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
						SQLFlag:        true,
					},
				},
			},
			UserInnerKeyValue: &[]KeyValueInfo{
				{
					Key: "__content__",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",:-/ ",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
			},
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.CreateIndex(req)
		suite.validateError(err, expectedErr)

		_, err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
		suite.NoError(err)
	}
}

func (suite *SDKIndexTestSuite) TestCreateIndexAbnormally() {
	testcases := map[*CreateIndexRequest]*Error{
		{
			TopicID: suite.topic,
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidIndexArgument",
			Message:  "At least one of full-text indexing and key-value indexing must be turned on",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.CreateIndex(req)
		suite.validateError(err, expectedErr)

		if err == nil {
			_, err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
			suite.NoError(err)
		}
	}
}

func (suite *SDKIndexTestSuite) TestDeleteIndexAbnormally() {
	_, err := suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "IndexNotExists",
		Message:  "Index does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKIndexTestSuite) TestModifyIndexNormally() {
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
		suite.Equal(expectGetIndexResp.KeyValue, actualGetIndexResp.KeyValue)
	}

	_, err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
	suite.NoError(err)
}

func (suite *SDKIndexTestSuite) TestModifyIndexAbnormally() {
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

	testcases := map[*ModifyIndexRequest]*Error{
		{
			TopicID: suite.topic,
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidIndexArgument",
			Message:  "At least one of full-text indexing and key-value indexing must be turned on",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.ModifyIndex(req)
		suite.validateError(err, expectedErr)
	}

	_, err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
	suite.NoError(err)
}

func (suite *SDKIndexTestSuite) TestDescribeIndexNormally() {
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
		{
			TopicID: suite.topic,
			UserInnerKeyValue: &[]KeyValueInfo{
				{
					Key: "__content__",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",:-/ ",
						CasSensitive:   false,
						IncludeChinese: false,
						SQLFlag:        false,
					},
				},
			},
		}: {
			TopicID: suite.topic,
			UserInnerKeyValue: &[]KeyValueInfo{
				{
					Key: "__content__",
					Value: Value{
						ValueType:      "text",
						Delimiter:      ",:-/ ",
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
		if expectGetIndexResp.KeyValue != nil {
			suite.Equal(expectGetIndexResp.KeyValue, actualGetIndexResp.KeyValue)
		} else {
			suite.Equal(0, len(*actualGetIndexResp.KeyValue))
		}
		if expectGetIndexResp.UserInnerKeyValue != nil {
			suite.Equal(expectGetIndexResp.UserInnerKeyValue, actualGetIndexResp.UserInnerKeyValue)
		} else {
			suite.Equal(0, len(*actualGetIndexResp.UserInnerKeyValue))
		}

		_, err = suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
		suite.NoError(err)
	}
}

func (suite *SDKIndexTestSuite) TestDescribeIndexAbnormally() {
	testcases := map[*DescribeIndexRequest]*Error{
		{
			TopicID: suite.topic,
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "IndexNotExists",
			Message:  "Index does not exist.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeIndex(req)
		suite.validateError(err, expectedErr)
	}
}
