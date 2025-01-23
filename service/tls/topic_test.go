package tls

import (
	"net/http"
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

	projectId, err := CreateProject("golang-sdk-create-project-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)

	suite.NoError(err)
	suite.project = projectId
}

func (suite *SDKTopicTestSuite) TearDownTest() {
	for _, topicID := range suite.topicList {
		_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: topicID})
		suite.NoError(err)
	}

	_, err := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(err)

	suite.topicList = nil
}

func (suite *SDKTopicTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKTopicTestSuite(t *testing.T) {
	suite.Run(t, new(SDKTopicTestSuite))
}

func (suite *SDKTopicTestSuite) TestCreateTopicNormally() {
	timeKey := "request_time"
	timeFormat := "%Y-%m-%dT%H:%M:%S,%f"
	tags := []TagInfo{
		{
			Key:   "TagKey1",
			Value: "TagValue1",
		},
		{
			Key:   "TagKey2",
			Value: "TagValue2",
		},
	}
	logPublicIP := true

	testcases := map[*CreateTopicRequest]*DescribeTopicResponse{
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
			LogPublicIP: true,
		},
		{
			ProjectID:      suite.project,
			TopicName:      "topic2",
			Ttl:            120,
			Description:    "create a topic with new arguments",
			ShardCount:     2,
			MaxSplitShard:  Int32Ptr(2),
			AutoSplit:      true,
			EnableTracking: BoolPtr(false),
			TimeKey:        &timeKey,
			TimeFormat:     &timeFormat,
			Tags:           tags,
			LogPublicIP:    &logPublicIP,
			EnableHotTtl:   BoolPtr(true),
			HotTtl:         Int32Ptr(30),
			ColdTtl:        Int32Ptr(30),
			ArchiveTtl:     Int32Ptr(60),
		}: {
			ProjectID:      suite.project,
			TopicName:      "topic2",
			Ttl:            120,
			Description:    "create a topic with new arguments",
			ShardCount:     2,
			MaxSplitShard:  2,
			AutoSplit:      true,
			EnableTracking: false,
			TimeKey:        timeKey,
			TimeFormat:     timeFormat,
			Tags:           tags,
			LogPublicIP:    logPublicIP,
			EnableHotTtl:   true,
			HotTtl:         30,
			ColdTtl:        30,
			ArchiveTtl:     60,
		},
	}

	for req, topicInfo := range testcases {
		createTopicResp, err := suite.cli.CreateTopic(req)
		suite.NoError(err)
		suite.topicList = append(suite.topicList, createTopicResp.TopicID)

		getTopicResponse, err := suite.cli.DescribeTopic(&DescribeTopicRequest{TopicID: createTopicResp.TopicID})
		suite.NoError(err)
		suite.Equal(createTopicResp.TopicID, getTopicResponse.TopicID)
		suite.Equal(topicInfo.TopicName, getTopicResponse.TopicName)
		suite.Equal(topicInfo.Ttl, getTopicResponse.Ttl)
		suite.Equal(topicInfo.ShardCount, getTopicResponse.ShardCount)
		suite.Equal(topicInfo.Description, getTopicResponse.Description)
		suite.Equal(topicInfo.TimeKey, getTopicResponse.TimeKey)
		suite.Equal(topicInfo.TimeFormat, getTopicResponse.TimeFormat)
		suite.Equal(topicInfo.LogPublicIP, getTopicResponse.LogPublicIP)
		if topicInfo.Tags != nil {
			suite.Equal(len(topicInfo.Tags), len(getTopicResponse.Tags))
		} else {
			suite.Equal(topicInfo.Tags, getTopicResponse.Tags)
		}

		if topicInfo.AutoSplit {
			suite.True(getTopicResponse.AutoSplit)
			suite.Equal(topicInfo.MaxSplitShard, getTopicResponse.MaxSplitShard)
		} else {
			suite.False(getTopicResponse.AutoSplit)
			suite.Equal(int32(0), getTopicResponse.MaxSplitShard)
		}

		suite.Equal(topicInfo.EnableTracking, getTopicResponse.EnableTracking)
		suite.Equal(topicInfo.EnableHotTtl, getTopicResponse.EnableHotTtl)
		suite.Equal(topicInfo.HotTtl, getTopicResponse.HotTtl)
		suite.Equal(topicInfo.ColdTtl, getTopicResponse.ColdTtl)
		suite.Equal(topicInfo.ArchiveTtl, getTopicResponse.ArchiveTtl)
	}
}

func (suite *SDKTopicTestSuite) TestCreateTopicAbnormally() {
	topicName := "existing-topic"
	resp, err := suite.cli.CreateTopic(&CreateTopicRequest{
		ProjectID:    suite.project,
		TopicName:    topicName,
		Ttl:          9,
		Description:  "test topic",
		ShardCount:   3,
		EnableHotTtl: BoolPtr(true),
		HotTtl:       Int32Ptr(8),
	})
	suite.NoError(err)
	suite.topicList = append(suite.topicList, resp.TopicID)

	resp, err = suite.cli.CreateTopic(&CreateTopicRequest{
		ProjectID:   suite.project,
		TopicName:   topicName,
		Ttl:         9,
		Description: "test topic",
		ShardCount:  3,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusConflict,
		Code:     "TopicAlreadyExist",
		Message:  "Topic " + topicName + " already exist",
	}
	suite.validateError(err, expectedErr)

	if err == nil {
		suite.topicList = append(suite.topicList, resp.TopicID)
	}
}

func (suite *SDKTopicTestSuite) TestDeleteTopicAbnormally() {
	topicID := uuid.New().String()
	_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{
		TopicID: topicID,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "TopicNotExist",
		Message:  "Topic does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKTopicTestSuite) TestModifyTopicNormally() {
	createTopicReq := &CreateTopicRequest{
		ProjectID:    suite.project,
		TopicName:    "golang-sdk-update-topic-" + uuid.New().String(),
		Ttl:          9,
		Description:  "test topic",
		ShardCount:   2,
		EnableHotTtl: BoolPtr(true),
		HotTtl:       Int32Ptr(8),
	}

	createTopicResp, err := suite.cli.CreateTopic(createTopicReq)
	suite.NoError(err)
	suite.topicList = append(suite.topicList, createTopicResp.TopicID)

	_, err = suite.cli.ModifyTopic(&ModifyTopicRequest{
		Ttl:        Uint16Ptr(67),
		TopicID:    createTopicResp.TopicID,
		TopicName:  StrPtr("new-topic-name"),
		HotTtl:     Int32Ptr(7),
		ColdTtl:    Int32Ptr(30),
		ArchiveTtl: Int32Ptr(30),
	})
	suite.NoError(err)
	describeTopicResp, err := suite.cli.DescribeTopic(&DescribeTopicRequest{TopicID: createTopicResp.TopicID})
	suite.NoError(err)
	suite.Equal("new-topic-name", describeTopicResp.TopicName)
	suite.Equal(67, int(describeTopicResp.Ttl))
	suite.Equal(7, int(describeTopicResp.HotTtl))
	suite.Equal(true, describeTopicResp.EnableHotTtl)
	suite.Equal("test topic", describeTopicResp.Description)
	suite.Equal(30, int(describeTopicResp.ColdTtl))
	suite.Equal(30, int(describeTopicResp.ArchiveTtl))

	_, err = suite.cli.ModifyTopic(&ModifyTopicRequest{
		TopicID:     createTopicResp.TopicID,
		TopicName:   StrPtr("new-name"),
		Ttl:         Uint16Ptr(5),
		Description: StrPtr("new desc"),
	})
	suite.NoError(err)
	describeTopicResp, err = suite.cli.DescribeTopic(&DescribeTopicRequest{TopicID: createTopicResp.TopicID})
	suite.NoError(err)
	suite.Equal("new-name", describeTopicResp.TopicName)
	suite.Equal(5, int(describeTopicResp.Ttl))
	suite.Equal("new desc", describeTopicResp.Description)
}

func (suite *SDKTopicTestSuite) TestModifyTopicAbnormally() {
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

	topicName := "_Topic-$$$-Name_"
	_, err = suite.cli.ModifyTopic(&ModifyTopicRequest{
		TopicID:   createTopicResp.TopicID,
		TopicName: &topicName,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key topic_name, value " + topicName + ", please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKTopicTestSuite) TestDescribeTopicNormally() {
	timeKey := "request_time"
	timeFormat := "%Y-%m-%dT%H:%M:%S,%f"
	tags := []TagInfo{
		{
			Key:   "TagKey1",
			Value: "TagValue1",
		},
		{
			Key:   "TagKey2",
			Value: "TagValue2",
		},
	}
	logPublicIP := true

	createTopicResp, err := suite.cli.CreateTopic(&CreateTopicRequest{
		ProjectID:   suite.project,
		TopicName:   "test-topic",
		Ttl:         30,
		Description: "create a topic with new arguments",
		ShardCount:  2,
		TimeKey:     &timeKey,
		TimeFormat:  &timeFormat,
		Tags:        tags,
		LogPublicIP: &logPublicIP,
	})
	suite.NoError(err)
	suite.topicList = append(suite.topicList, createTopicResp.TopicID)

	describeTopicResp, err := suite.cli.DescribeTopic(&DescribeTopicRequest{TopicID: createTopicResp.TopicID})
	suite.NoError(err)
	suite.Equal("test-topic", describeTopicResp.TopicName)
	suite.Equal(suite.project, describeTopicResp.ProjectID)
	suite.Equal(30, int(describeTopicResp.Ttl))
	suite.Equal(2, int(describeTopicResp.ShardCount))
	suite.Equal("create a topic with new arguments", describeTopicResp.Description)
	suite.Equal(timeKey, describeTopicResp.TimeKey)
	suite.Equal(timeFormat, describeTopicResp.TimeFormat)
	suite.Equal(logPublicIP, describeTopicResp.LogPublicIP)
	suite.Equal(len(tags), len(describeTopicResp.Tags))
}

func (suite *SDKTopicTestSuite) TestDescribeTopicAbnormally() {
	_, err := suite.cli.DescribeTopic(&DescribeTopicRequest{
		TopicID: uuid.New().String(),
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "TopicNotExist",
		Message:  "Topic does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKTopicTestSuite) TestDescribeTopicsNormally() {
	timeKey := "request_time"
	timeFormat := "%Y-%m-%dT%H:%M:%S,%f"
	tags := []TagInfo{
		{
			Key:   "TagKey1",
			Value: "TagValue1",
		},
		{
			Key:   "TagKey2",
			Value: "TagValue2",
		},
	}
	logPublicIP := true

	createTopicReqs := []*CreateTopicRequest{
		{
			ProjectID:    suite.project,
			TopicName:    "sdk-a" + uuid.New().String(),
			Ttl:          1,
			Description:  "test topic",
			ShardCount:   2,
			EnableHotTtl: BoolPtr(false),
			HotTtl:       Int32Ptr(0),
			ColdTtl:      Int32Ptr(0),
			ArchiveTtl:   Int32Ptr(0),
		},
		{
			ProjectID:    suite.project,
			TopicName:    "sdk-a" + uuid.New().String(),
			Ttl:          2,
			Description:  "test topic",
			ShardCount:   2,
			EnableHotTtl: BoolPtr(false),
			HotTtl:       Int32Ptr(0),
			ColdTtl:      Int32Ptr(0),
			ArchiveTtl:   Int32Ptr(0),
		},
		{
			ProjectID:    suite.project,
			TopicName:    "sdk-b" + uuid.New().String(),
			Ttl:          3,
			Description:  "test topic",
			ShardCount:   2,
			EnableHotTtl: BoolPtr(false),
			HotTtl:       Int32Ptr(0),
			ColdTtl:      Int32Ptr(0),
			ArchiveTtl:   Int32Ptr(0),
		},
		{
			ProjectID:    suite.project,
			TopicName:    "sdk-c" + uuid.New().String(),
			Ttl:          120,
			Description:  "create a topic with new arguments",
			ShardCount:   2,
			TimeKey:      &timeKey,
			TimeFormat:   &timeFormat,
			Tags:         tags,
			LogPublicIP:  &logPublicIP,
			EnableHotTtl: BoolPtr(true),
			HotTtl:       Int32Ptr(30),
			ColdTtl:      Int32Ptr(30),
			ArchiveTtl:   Int32Ptr(60),
		},
	}

	for _, createTopicReq := range createTopicReqs {
		createTopicResp, err := suite.cli.CreateTopic(createTopicReq)
		suite.NoError(err)
		suite.topicList = append(suite.topicList, createTopicResp.TopicID)
	}

	testcases := map[*DescribeTopicsRequest]*DescribeTopicsResponse{
		{
			ProjectID:   suite.project,
			ProjectName: "sdk-",
		}: {
			Topics: []*Topic{
				{
					TopicName:    createTopicReqs[3].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[3].Ttl,
					ShardCount:   int32(createTopicReqs[3].ShardCount),
					Description:  createTopicReqs[3].Description,
					EnableHotTtl: *createTopicReqs[3].EnableHotTtl,
					HotTtl:       *createTopicReqs[3].HotTtl,
					ColdTtl:      *createTopicReqs[3].ColdTtl,
					ArchiveTtl:   *createTopicReqs[3].ArchiveTtl,
				},
				{
					TopicName:    createTopicReqs[2].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[2].Ttl,
					ShardCount:   int32(createTopicReqs[2].ShardCount),
					Description:  createTopicReqs[2].Description,
					EnableHotTtl: *createTopicReqs[2].EnableHotTtl,
					HotTtl:       *createTopicReqs[2].HotTtl,
					ColdTtl:      *createTopicReqs[2].ColdTtl,
					ArchiveTtl:   *createTopicReqs[2].ArchiveTtl,
				},
				{
					TopicName:    createTopicReqs[1].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[1].Ttl,
					ShardCount:   int32(createTopicReqs[1].ShardCount),
					Description:  createTopicReqs[1].Description,
					EnableHotTtl: *createTopicReqs[1].EnableHotTtl,
					HotTtl:       *createTopicReqs[1].HotTtl,
					ColdTtl:      *createTopicReqs[1].ColdTtl,
					ArchiveTtl:   *createTopicReqs[1].ArchiveTtl,
				},
				{
					TopicName:    createTopicReqs[0].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[0].Ttl,
					ShardCount:   int32(createTopicReqs[0].ShardCount),
					Description:  createTopicReqs[0].Description,
					EnableHotTtl: *createTopicReqs[0].EnableHotTtl,
					HotTtl:       *createTopicReqs[0].HotTtl,
					ColdTtl:      *createTopicReqs[0].ColdTtl,
					ArchiveTtl:   *createTopicReqs[0].ArchiveTtl,
				},
			},
			Total: 4,
		},
		{
			ProjectID:  suite.project,
			PageNumber: 1,
			PageSize:   20,
		}: {
			Topics: []*Topic{
				{
					TopicName:    createTopicReqs[3].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[3].Ttl,
					ShardCount:   int32(createTopicReqs[3].ShardCount),
					Description:  createTopicReqs[3].Description,
					EnableHotTtl: *createTopicReqs[3].EnableHotTtl,
					HotTtl:       *createTopicReqs[3].HotTtl,
					ColdTtl:      *createTopicReqs[3].ColdTtl,
					ArchiveTtl:   *createTopicReqs[3].ArchiveTtl,
				},
				{
					TopicName:    createTopicReqs[2].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[2].Ttl,
					ShardCount:   int32(createTopicReqs[2].ShardCount),
					Description:  createTopicReqs[2].Description,
					EnableHotTtl: *createTopicReqs[2].EnableHotTtl,
					HotTtl:       *createTopicReqs[2].HotTtl,
					ColdTtl:      *createTopicReqs[2].ColdTtl,
					ArchiveTtl:   *createTopicReqs[2].ArchiveTtl,
				},
				{
					TopicName:    createTopicReqs[1].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[1].Ttl,
					ShardCount:   int32(createTopicReqs[1].ShardCount),
					Description:  createTopicReqs[1].Description,
					EnableHotTtl: *createTopicReqs[1].EnableHotTtl,
					HotTtl:       *createTopicReqs[1].HotTtl,
					ColdTtl:      *createTopicReqs[1].ColdTtl,
					ArchiveTtl:   *createTopicReqs[1].ArchiveTtl,
				},
				{
					TopicName:    createTopicReqs[0].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[0].Ttl,
					ShardCount:   int32(createTopicReqs[0].ShardCount),
					Description:  createTopicReqs[0].Description,
					EnableHotTtl: *createTopicReqs[0].EnableHotTtl,
					HotTtl:       *createTopicReqs[0].HotTtl,
					ColdTtl:      *createTopicReqs[0].ColdTtl,
					ArchiveTtl:   *createTopicReqs[0].ArchiveTtl,
				},
			},
			Total: 4,
		},
		{
			ProjectID: suite.project,
			PageSize:  20,
			TopicName: "sdk-a",
		}: {
			Topics: []*Topic{
				{
					TopicName:    createTopicReqs[1].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[1].Ttl,
					ShardCount:   int32(createTopicReqs[1].ShardCount),
					Description:  createTopicReqs[1].Description,
					EnableHotTtl: *createTopicReqs[1].EnableHotTtl,
					HotTtl:       *createTopicReqs[1].HotTtl,
					ColdTtl:      *createTopicReqs[1].ColdTtl,
					ArchiveTtl:   *createTopicReqs[1].ArchiveTtl,
				},
				{
					TopicName:    createTopicReqs[0].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[0].Ttl,
					ShardCount:   int32(createTopicReqs[0].ShardCount),
					Description:  createTopicReqs[0].Description,
					EnableHotTtl: *createTopicReqs[0].EnableHotTtl,
					HotTtl:       *createTopicReqs[0].HotTtl,
					ColdTtl:      *createTopicReqs[0].ColdTtl,
					ArchiveTtl:   *createTopicReqs[0].ArchiveTtl,
				},
			},
			Total: 2,
		},
		{
			ProjectID:  suite.project,
			PageNumber: 1,
			PageSize:   20,
			TopicName:  "sdk-c",
			Tags:       tags,
		}: {
			Topics: []*Topic{
				{
					TopicName:    createTopicReqs[3].TopicName,
					ProjectID:    suite.project,
					Ttl:          createTopicReqs[3].Ttl,
					ShardCount:   int32(createTopicReqs[3].ShardCount),
					Description:  createTopicReqs[3].Description,
					EnableHotTtl: *createTopicReqs[3].EnableHotTtl,
					HotTtl:       *createTopicReqs[3].HotTtl,
					ColdTtl:      *createTopicReqs[3].ColdTtl,
					ArchiveTtl:   *createTopicReqs[3].ArchiveTtl,
				},
			},
			Total: 1,
		},
	}

	for listTopicReq, expectListTopicResp := range testcases {
		actualListTopicResp, err := suite.cli.DescribeTopics(listTopicReq)
		suite.NoError(err)

		suite.Equal(expectListTopicResp.Total, actualListTopicResp.Total)
		suite.Equal(len(expectListTopicResp.Topics), len(actualListTopicResp.Topics))

		expectTopics := expectListTopicResp.Topics
		actualTopics := actualListTopicResp.Topics

		for _, expectTopic := range expectTopics {
			isMatched := false
			for _, actualTopic := range actualTopics {
				if expectTopic.TopicName == actualTopic.TopicName {
					isMatched = true
					suite.Equal(expectTopic.ProjectID, actualTopic.ProjectID)
					suite.Equal(expectTopic.Ttl, actualTopic.Ttl)
					suite.Equal(expectTopic.ShardCount, actualTopic.ShardCount)
					suite.Equal(expectTopic.Description, actualTopic.Description)
					suite.Equal(expectTopic.EnableHotTtl, actualTopic.EnableHotTtl)
					suite.Equal(expectTopic.HotTtl, actualTopic.HotTtl)
					suite.Equal(expectTopic.ColdTtl, actualTopic.ColdTtl)
					suite.Equal(expectTopic.ArchiveTtl, actualTopic.ArchiveTtl)
					break
				}
			}
			suite.Equal(true, isMatched)
		}
	}

	resp, err := suite.cli.DescribeTopics(&DescribeTopicsRequest{})
	suite.NoError(err)
	suite.GreaterOrEqual(int(resp.Total), 0)
}

func (suite *SDKTopicTestSuite) TestDescribeTopicsAbnormally() {
	_, err := suite.cli.DescribeTopics(&DescribeTopicsRequest{ProjectID: uuid.New().String()})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)
}
