package tls

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKTagsTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKTagsTestSuite) SetupTest() {
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

func (suite *SDKTagsTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKTagsTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKTagsTestSuite(t *testing.T) {
	suite.Run(t, new(SDKTagsTestSuite))
}

func (suite *SDKTagsTestSuite) TestAddTagsToResourceNormally() {
	testcases := map[*AddTagsToResourceRequest]*Error{
		{
			ResourceType:  "project",
			ResourcesList: []string{suite.project},
			Tags: []TagInfo{
				{
					Key:   "TagKey1",
					Value: "TagValue1",
				},
			},
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.AddTagsToResource(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestAddTagsToResourceAbnormally() {
	testcases := map[*AddTagsToResourceRequest]*Error{
		{
			ResourceType:  "project",
			ResourcesList: []string{uuid.New().String()},
			Tags: []TagInfo{
				{
					Key:   "TagKey1",
					Value: "TagValue1",
				},
			},
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidResourcesList",
			Message:  "Part of resourcesList is not exist",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.AddTagsToResource(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestRemoveTagsFromResourceNormally() {
	_, err := suite.cli.AddTagsToResource(&AddTagsToResourceRequest{
		ResourceType:  "project",
		ResourcesList: []string{suite.project},
		Tags: []TagInfo{
			{
				Key:   "TagKey1",
				Value: "TagValue1",
			},
		},
	})
	suite.NoError(err)

	testcases := map[*RemoveTagsFromResourceRequest]*Error{
		{
			ResourceType:  "project",
			ResourcesList: []string{suite.project},
			TagKeyList:    []string{"TagKey1"},
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.RemoveTagsFromResource(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestRemoveTagsFromResourceAbnormally() {
	_, err := suite.cli.AddTagsToResource(&AddTagsToResourceRequest{
		ResourceType:  "project",
		ResourcesList: []string{suite.project},
		Tags: []TagInfo{
			{
				Key:   "TagKey1",
				Value: "TagValue1",
			},
		},
	})
	suite.NoError(err)

	testcases := map[*RemoveTagsFromResourceRequest]*Error{
		{
			ResourceType:  "project",
			ResourcesList: []string{uuid.New().String()},
			TagKeyList:    []string{"TagKey1"},
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidResourcesList",
			Message:  "Part of resourcesList is not exist",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.RemoveTagsFromResource(req)
		suite.validateError(err, expectedErr)
	}
}
