package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKTopicTestSuite struct {
	suite.Suite

	cli       Client
	project   string
	topicList []string
}

func (suite *SDKTopicTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)

	suite.NoError(err)
	suite.project = projectId
}

func (suite *SDKTopicTestSuite) TearDownTest() {
	for _, topicID := range suite.topicList {
		suite.NoError(suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: topicID}))
	}

	suite.NoError(suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project}))

	suite.topicList = nil
}

func TestSDKTopicTestSuite(t *testing.T) {
	suite.Run(t, new(SDKTopicTestSuite))
}

// TestCreateTopic: test create topic
func (suite *SDKTopicTestSuite) TestCreateTopic() {
	testcases := map[*CreateTopicRequest]*GetTopicResponse{
		{
			ProjectID:   suite.project,
			TopicName:   "topic1",
			Ttl:         3,
			Description: "test topic",
			ShardCount:  3,
		}: {
			TopicName:   "topic1",
			ProjectID:   suite.project,
			Ttl:         3,
			ShardCount:  3,
			Description: "test topic",
		},
	}

	for req, topicInfo := range testcases {
		createTopicResp, err := suite.cli.CreateTopic(req)
		suite.NoError(err)

		suite.topicList = append(suite.topicList, createTopicResp.TopicID)

		getTopicResponse, err := suite.cli.GetTopic(&GetTopicRequest{TopicID: createTopicResp.TopicID})
		suite.NoError(err)

		suite.Equal(createTopicResp.TopicID, getTopicResponse.TopicID)
		suite.Equal(topicInfo.TopicName, getTopicResponse.TopicName)
		suite.Equal(topicInfo.Ttl, getTopicResponse.Ttl)
		suite.Equal(topicInfo.ShardCount, getTopicResponse.ShardCount)
		suite.Equal(topicInfo.Description, getTopicResponse.Description)
	}
}

// TestUpdateTopic: test update topic
func (suite *SDKTopicTestSuite) TestUpdateTopic() {
	createTopicReq := &CreateTopicRequest{
		ProjectID:   suite.project,
		TopicName:   "golang-sdk-update-topic-" + uuid.New().String(),
		Ttl:         3,
		Description: "test topic",
		ShardCount:  2,
	}

	createTopicResp, err := suite.cli.CreateTopic(createTopicReq)
	suite.NoError(err)
	suite.topicList = append(suite.topicList, createTopicResp.TopicID)

	testcases := map[*UpdateTopicRequest]*GetTopicResponse{
		{
			TopicID:     createTopicResp.TopicID,
			TopicName:   StrPtr("new-name"),
			Ttl:         Uint16Ptr(5),
			Description: StrPtr("new desc"),
		}: {
			TopicName:   "new-name",
			Ttl:         5,
			Description: "new desc",
		},
	}

	for updateTopicReq, getTopicResp := range testcases {
		err := suite.cli.UpdateTopic(updateTopicReq)
		suite.NoError(err)

		actualGetTopicResp, err := suite.cli.GetTopic(&GetTopicRequest{TopicID: createTopicResp.TopicID})
		suite.NoError(err)

		suite.Equal(getTopicResp.TopicName, actualGetTopicResp.TopicName)
		suite.Equal(getTopicResp.Ttl, actualGetTopicResp.Ttl)
		suite.Equal(getTopicResp.Description, actualGetTopicResp.Description)
	}
}

// TestListTopic: test list topic
func (suite *SDKTopicTestSuite) TestListTopic() {
	createTopicReqs := []*CreateTopicRequest{
		{
			ProjectID:   suite.project,
			TopicName:   "sdk-a" + uuid.New().String(),
			Ttl:         1,
			Description: "test topic",
			ShardCount:  2,
		},
		{
			ProjectID:   suite.project,
			TopicName:   "sdk-a" + uuid.New().String(),
			Ttl:         2,
			Description: "test topic",
			ShardCount:  2,
		},
		{
			ProjectID:   suite.project,
			TopicName:   "sdk-b" + uuid.New().String(),
			Ttl:         3,
			Description: "test topic",
			ShardCount:  2,
		},
	}

	for _, createTopicReq := range createTopicReqs {
		createTopicResp, err := suite.cli.CreateTopic(createTopicReq)
		suite.NoError(err)
		suite.topicList = append(suite.topicList, createTopicResp.TopicID)
	}

	testcases := map[*ListTopicRequest]*ListTopicResponse{
		{
			ProjectID: suite.project,
			Offset:    0,
			Limit:     20,
			TopicName: "",
			TopicID:   "",
		}: {
			Topics: []*Topic{
				{
					TopicName:   createTopicReqs[2].TopicName,
					ProjectID:   suite.project,
					Ttl:         createTopicReqs[2].Ttl,
					ShardCount:  int32(createTopicReqs[2].ShardCount),
					Description: createTopicReqs[2].Description,
				},
				{
					TopicName:   createTopicReqs[1].TopicName,
					ProjectID:   suite.project,
					Ttl:         createTopicReqs[1].Ttl,
					ShardCount:  int32(createTopicReqs[1].ShardCount),
					Description: createTopicReqs[1].Description,
				},
				{
					TopicName:   createTopicReqs[0].TopicName,
					ProjectID:   suite.project,
					Ttl:         createTopicReqs[0].Ttl,
					ShardCount:  int32(createTopicReqs[0].ShardCount),
					Description: createTopicReqs[0].Description,
				},
			},
			Total: 3,
		},
		{
			ProjectID: suite.project,
			Offset:    0,
			Limit:     20,
			TopicName: "sdk-a",
			TopicID:   "",
		}: {
			Topics: []*Topic{
				{
					TopicName:   createTopicReqs[1].TopicName,
					ProjectID:   suite.project,
					Ttl:         createTopicReqs[1].Ttl,
					ShardCount:  int32(createTopicReqs[1].ShardCount),
					Description: createTopicReqs[1].Description,
				},
				{
					TopicName:   createTopicReqs[0].TopicName,
					ProjectID:   suite.project,
					Ttl:         createTopicReqs[0].Ttl,
					ShardCount:  int32(createTopicReqs[0].ShardCount),
					Description: createTopicReqs[0].Description,
				},
			},
			Total: 2,
		},
		{
			ProjectID: suite.project,
			Offset:    2,
			Limit:     20,
			TopicName: "",
			TopicID:   "",
		}: {
			Topics: []*Topic{
				{
					TopicName:   createTopicReqs[0].TopicName,
					ProjectID:   suite.project,
					Ttl:         createTopicReqs[0].Ttl,
					ShardCount:  int32(createTopicReqs[0].ShardCount),
					Description: createTopicReqs[0].Description,
				},
			},
			Total: 3,
		},
	}

	for listTopicReq, expectListTopicResp := range testcases {
		actualListTopicResp, err := suite.cli.ListTopic(listTopicReq)
		suite.NoError(err)

		suite.Equal(expectListTopicResp.Total, actualListTopicResp.Total)
		suite.Equal(len(expectListTopicResp.Topics), len(actualListTopicResp.Topics))

		expectTopics := expectListTopicResp.Topics
		actualTopics := actualListTopicResp.Topics

		for i, actualTopic := range actualTopics {
			suite.Equal(expectTopics[i].TopicName, actualTopic.TopicName)
			suite.Equal(expectTopics[i].ProjectID, actualTopic.ProjectID)
			suite.Equal(expectTopics[i].Ttl, actualTopic.Ttl)
			suite.Equal(expectTopics[i].ShardCount, actualTopic.ShardCount)
			suite.Equal(expectTopics[i].Description, actualTopic.Description)
		}
	}
}
