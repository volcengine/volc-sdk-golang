package tls

import (
	"net/http"
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

func (suite *SDKWebtracksTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKWebtracksTestSuite(t *testing.T) {
	suite.Run(t, new(SDKWebtracksTestSuite))
}

func (suite *SDKWebtracksTestSuite) TestWebTracksNormally() {
	testcases := map[*WebTracksRequest]*Error{
		{
			ProjectID:    suite.project,
			TopicID:      suite.topic,
			CompressType: "lz4",
			Source:       "192.168.1.1",
			Logs: []map[string]string{
				{
					"key1": "value11",
					"key2": "value21",
				},
				{
					"key1": "value12",
					"key2": "value22",
				},
			},
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.WebTracks(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKWebtracksTestSuite) TestWebTracksAbnormally() {
	testcases := map[*WebTracksRequest]*Error{
		{
			ProjectID:    suite.project,
			TopicID:      suite.topic,
			CompressType: "rar",
			Source:       "192.168.1.1",
			Logs: []map[string]string{
				{
					"key1": "value11",
					"key2": "value21",
				},
				{
					"key1": "value12",
					"key2": "value22",
				},
			},
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key x-tls-compresstype, value rar, please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.WebTracks(req)
		suite.validateError(err, expectedErr)
	}
}
