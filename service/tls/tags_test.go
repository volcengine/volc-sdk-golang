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
	if expectErr == nil {
		suite.NoError(err)
		return
	}

	suite.Error(err)

	sdkErr, ok := err.(*Error)
	suite.True(ok)
	if !ok {
		return
	}

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

func (suite *SDKTagsTestSuite) TestTagResourcesNormally() {
	testcases := map[*TagResourcesRequest]*Error{
		{
			ResourceType: "project",
			ResourcesIds: []string{suite.project},
			Tags: []TagInfo{
				{
					Key:   "owner",
					Value: "test-user",
				},
				{
					Key:   "environment",
					Value: "test",
				},
			},
		}: nil,
		{
			ResourceType: "topic",
			ResourcesIds: []string{suite.topic},
			Tags: []TagInfo{
				{
					Key:   "topic-owner",
					Value: "test-user",
				},
				{
					Key:   "topic-type",
					Value: "test",
				},
			},
		}: nil,
		{
			ResourceType: "topic",
			ResourcesIds: []string{suite.topic},
			Tags: []TagInfo{
				{
					Key:   "multi-tag",
					Value: "multi-value",
				},
			},
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.TagResources(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestTagResourcesAbnormally() {
	testcases := map[*TagResourcesRequest]*Error{
		{
			ResourceType: "",
			ResourcesIds: []string{suite.project},
			Tags: []TagInfo{
				{
					Key:   "owner",
					Value: "test-user",
				},
			},
		}: {
			HTTPCode: -1,
			Code:     "ClientError",
			Message:  "Invalid argument, empty ResourceType",
		},
		{
			ResourceType: "project",
			ResourcesIds: []string{},
			Tags: []TagInfo{
				{
					Key:   "owner",
					Value: "test-user",
				},
			},
		}: {
			HTTPCode: -1,
			Code:     "ClientError",
			Message:  "Invalid argument, empty ResourcesIds",
		},
		{
			ResourceType: "project",
			ResourcesIds: []string{suite.project},
			Tags:         []TagInfo{},
		}: {
			HTTPCode: -1,
			Code:     "ClientError",
			Message:  "Invalid argument, empty Tags",
		},
		{
			ResourceType: "project",
			ResourcesIds: make([]string, 51),
			Tags: []TagInfo{
				{
					Key:   "owner",
					Value: "test-user",
				},
			},
		}: {
			HTTPCode: -1,
			Code:     "ClientError",
			Message:  "Invalid argument, ResourcesIds count exceeds 50",
		},
		{
			ResourceType: "project",
			ResourcesIds: []string{uuid.New().String()},
			Tags: []TagInfo{
				{
					Key:   "owner",
					Value: "test-user",
				},
			},
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidResourcesIds",
			Message:  "Part of resourcesIds is not exist",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.TagResources(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestListTagsForResourcesNormally() {
	_, err := suite.cli.AddTagsToResource(&AddTagsToResourceRequest{
		ResourceType:  "project",
		ResourcesList: []string{suite.project},
		Tags: []TagInfo{
			{
				Key:   "test-key",
				Value: "test-value",
			},
		},
	})
	suite.NoError(err)

	// Test listing project tags
	testcases := map[*ListTagsForResourcesRequest]*Error{
		{
			ResourceType: "project",
			ResourcesIds: []string{suite.project},
			MaxResults:   10,
		}: nil,
		{
			ResourceType: "project",
			MaxResults:   10,
		}: nil,
		{
			ResourceType: "topic",
			ResourcesIds: []string{suite.topic},
			MaxResults:   10,
		}: nil,
		{
			ResourceType: "project",
			TagFilters: []*FilterTag{
				{
					Key:    "test-key",
					Values: []string{"test-value"},
				},
			},
			MaxResults: 10,
		}: nil,
		{
			ResourceType: "project",
			MaxResults:   10,
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.ListTagsForResources(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestListTagsForResourcesAbnormally() {
	testcases := map[*ListTagsForResourcesRequest]*Error{
		{
			ResourceType: "",
		}: {
			HTTPCode: -1,
			Code:     "ClientError",
			Message:  "Invalid argument, empty ResourceType",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.ListTagsForResources(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestUntagResourcesNormally() {
	_, err := suite.cli.AddTagsToResource(&AddTagsToResourceRequest{
		ResourceType:  "project",
		ResourcesList: []string{suite.project},
		Tags: []TagInfo{
			{
				Key:   "TagKey1",
				Value: "TagValue1",
			},
			{
				Key:   "TagKey2",
				Value: "TagValue2",
			},
		},
	})
	suite.NoError(err)

	testcases := map[*UntagResourcesRequest]*Error{
		{
			ResourceType:  "project",
			ResourcesList: []string{suite.project},
			TagKeyList:    []string{"TagKey1", "TagKey2"},
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.UntagResources(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTagsTestSuite) TestUntagResourcesAbnormally() {
	testcases := map[*UntagResourcesRequest]*Error{
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
		_, err := suite.cli.UntagResources(req)
		suite.validateError(err, expectedErr)
	}
}
