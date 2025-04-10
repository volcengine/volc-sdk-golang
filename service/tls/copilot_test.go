package tls

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AppInstanceTestSuite struct {
	suite.Suite

	cli           Client
	project       string
	appInstanceId string
}

func (suite *AppInstanceTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	// 保证测试之前不存在AiAssistant应用实例
	instanceType := AppInstanceTypeAiAssistant
	accountId := os.Getenv("ACCOUNT_ID")
	describeRsp, err := suite.cli.DescribeAppInstances(&DescribeAppInstancesReq{
		PageNumber:   1,
		PageSize:     10,
		InstanceType: &instanceType,
		InstanceName: &accountId,
	})
	suite.NoError(err)

	existInstance := false
	var instanceId string
	for _, instance := range describeRsp.InstanceInfo {
		if instance.InstanceName == accountId {
			existInstance = true
			instanceId = instance.InstanceId
		}
	}

	if existInstance {
		_, err := suite.cli.DeleteAppInstance(&DeleteAppInstanceReq{
			InstanceId: instanceId,
		})
		suite.NoError(err)
	}
}

func (suite *AppInstanceTestSuite) TearDownTest() {
	if suite.appInstanceId != "" {
		_, err := suite.cli.DeleteAppInstance(&DeleteAppInstanceReq{
			InstanceId: suite.appInstanceId,
		})
		suite.NoError(err)
		suite.appInstanceId = ""
	}
}

func TestSDKCopilotTestSuite(t *testing.T) {
	suite.Run(t, new(AppInstanceTestSuite))
}

func (suite *AppInstanceTestSuite) TestCreateAppInstanceNormally() {
	// 测试创建
	req := &CreateAppInstanceReq{
		InstanceType: AppInstanceTypeAiAssistant,
		InstanceName: os.Getenv("ACCOUNT_ID"),
	}

	resp, err := suite.cli.CreateAppInstance(req)
	suite.NoError(err)
	suite.NotEmpty(resp.InstanceID)

	// 记录用于清理
	suite.appInstanceId = resp.InstanceID
}

func (suite *AppInstanceTestSuite) TestCreateAppInstanceAbnormally() {
	// 先创建一个实例
	req := &CreateAppInstanceReq{
		InstanceType: AppInstanceTypeAiAssistant,
		InstanceName: os.Getenv("ACCOUNT_ID"),
	}

	resp, err := suite.cli.CreateAppInstance(req)
	suite.NoError(err)
	suite.NotEmpty(resp.InstanceID)

	// 再重复创建
	testcases := map[*CreateAppInstanceReq]*Error{
		{ // 错误的实例类型
			InstanceType: AppInstanceTypeAiAssistant,
			InstanceName: os.Getenv("ACCOUNT_ID"),
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "QuotaExceed.AppInstance",
			Message:  "The quota of created App Instance exceeds.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.CreateAppInstance(req)
		var tlsErr *Error
		ok := errors.As(err, &tlsErr)
		suite.True(ok)
		suite.Contains(tlsErr.Message, expectedErr.Message)
		suite.Equal(tlsErr.Code, expectedErr.Code)
		suite.Equal(tlsErr.HTTPCode, expectedErr.HTTPCode)
	}
}

func (suite *AppInstanceTestSuite) TestDeleteAppInstanceNormally() {
	// 先创建一个实例
	req := &CreateAppInstanceReq{
		InstanceType: AppInstanceTypeAiAssistant,
		InstanceName: os.Getenv("ACCOUNT_ID"),
	}

	resp, err := suite.cli.CreateAppInstance(req)
	suite.NoError(err)
	suite.NotEmpty(resp.InstanceID)

	// 再删除
	_, err = suite.cli.DeleteAppInstance(&DeleteAppInstanceReq{
		InstanceId: resp.InstanceID,
	})
	suite.NoError(err)
}

func (suite *AppInstanceTestSuite) TestDeleteAppInstanceAbnormally() {
	testcases := map[*DeleteAppInstanceReq]*Error{
		{ // 不存在的实例id
			InstanceId: "dsafds",
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "AppInstanceNotExists",
			Message:  "App instance does not exist.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DeleteAppInstance(req)
		var tlsErr *Error
		ok := errors.As(err, &tlsErr)
		suite.True(ok)
		suite.Contains(tlsErr.Message, expectedErr.Message)
		suite.Equal(tlsErr.Code, expectedErr.Code)
		suite.Equal(tlsErr.HTTPCode, expectedErr.HTTPCode)
	}
}

type AppSceneMetaTestSuite struct {
	suite.Suite
	cli         Client
	project     string // 新增project字段
	sceneMetaId string
	instanceId  string
	topicId     string
}

func (suite *AppSceneMetaTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	// 创建测试Project
	projectResp, err := suite.cli.CreateProject(&CreateProjectRequest{
		Region:      os.Getenv("LOG_SERVICE_REGION"),
		ProjectName: "golang-sdk-create-project-" + uuid.New().String(),
		Description: "for scene meta testing",
	})
	suite.NoError(err)
	suite.project = projectResp.ProjectID // 存储projectID

	// 创建前置资源（Instance和Topic）
	// 先查询instance，如果没有就创建
	accountId := os.Getenv("ACCOUNT_ID")
	Type := AppInstanceTypeAiAssistant
	describeInsRsp, err := suite.cli.DescribeAppInstances(&DescribeAppInstancesReq{
		PageNumber:   1,
		PageSize:     10,
		InstanceType: &Type,
		InstanceName: &accountId,
	})
	suite.NoError(err)
	for _, instance := range describeInsRsp.InstanceInfo {
		if instance.InstanceName == accountId {
			suite.instanceId = instance.InstanceId
		}
	}

	if suite.instanceId == "" {
		instanceReq := &CreateAppInstanceReq{
			InstanceType: AppInstanceTypeAiAssistant,
			InstanceName: os.Getenv("ACCOUNT_ID"),
		}
		instanceResp, err := suite.cli.CreateAppInstance(instanceReq)
		suite.NoError(err)
		suite.instanceId = instanceResp.InstanceID
	}

	topicReq := &CreateTopicRequest{
		TopicName:   "test-meta-topic",
		ProjectID:   suite.project, // 使用新创建的project
		Ttl:         30,
		Description: "for scene meta testing",
		ShardCount:  1,
	}
	topicResp, err := suite.cli.CreateTopic(topicReq)
	suite.NoError(err)
	suite.topicId = topicResp.TopicID
}

func (suite *AppSceneMetaTestSuite) TearDownTest() {
	if suite.topicId != "" {
		_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topicId})
		suite.NoError(err)
	}
	if suite.project != "" { // 新增project清理
		_, err := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
		suite.NoError(err)
	}
}

func TestSDKAppSceneMetaTestSuite(t *testing.T) {
	suite.Run(t, new(AppSceneMetaTestSuite))
}

// 创建和删除场景 正常情况
func (suite *AppSceneMetaTestSuite) TestCreateAndDeleteAppSceneMetaNormally() {
	// 创建会话
	req := &CreateAppSceneMetaReq{
		InstanceId:        suite.instanceId,
		Id:                suite.topicId,
		CreateAPPMetaType: AppMetaTypeAiAssistantSession,
	}

	resp, err := suite.cli.CreateAppSceneMeta(req)
	suite.NoError(err)
	suite.NotEmpty(resp.Id)

	// 删除会话
	_, err = suite.cli.DeleteAppSceneMeta(&DeleteAppSceneMetaReq{
		InstanceId:        suite.instanceId,
		MetaId:            resp.Id,
		DeleteAPPMetaType: AppMetaTypeAiAssistantSession,
	})
	suite.NoError(err)
}

func (suite *AppSceneMetaTestSuite) TestCreateAppSceneMetaAbnormally() {
	testcases := map[*CreateAppSceneMetaReq]*Error{
		{ // 无效的InstanceId
			InstanceId:        "xx",
			Id:                suite.topicId,
			CreateAPPMetaType: AppMetaTypeAiAssistantSession,
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
		},

		{ // 无效的topicId
			InstanceId:        suite.instanceId,
			Id:                "xx",
			CreateAPPMetaType: AppMetaTypeAiAssistantSession,
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "TopicNotExist",
		},
		{ // 无效的MetaType
			InstanceId:        suite.instanceId,
			Id:                suite.topicId,
			CreateAPPMetaType: "xx",
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.CreateAppSceneMeta(req)
		var tlsErr *Error
		ok := errors.As(err, &tlsErr)
		suite.True(ok)
		suite.Equal(tlsErr.Code, expectedErr.Code)
		suite.Equal(tlsErr.HTTPCode, expectedErr.HTTPCode)
	}
}

// 删除场景 异常情况
func (suite *AppSceneMetaTestSuite) TestDeleteAppSceneMetaAbnormally() {
	testcases := map[*DeleteAppSceneMetaReq]*Error{
		{ // 无效场景ID
			DeleteAPPMetaType: "invalid--id",
			InstanceId:        suite.instanceId,
			MetaId:            "111",
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
		},

		{ // 会话id格式错误
			DeleteAPPMetaType: AppMetaTypeAiAssistantSession,
			InstanceId:        suite.instanceId,
			MetaId:            "invalid-id",
		}: {
			HTTPCode: -1,
			Code:     "ClientError",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DeleteAppSceneMeta(req)
		var tlsErr *Error
		ok := errors.As(err, &tlsErr)
		suite.True(ok)
		suite.Equal(tlsErr.Code, expectedErr.Code)
		suite.Equal(tlsErr.HTTPCode, expectedErr.HTTPCode)
	}
}

func (suite *AppSceneMetaTestSuite) TestModifyAppSceneMetaNormally() {
	// 测试反馈
	_, err := suite.cli.ModifyAppSceneMeta(&ModifyAppSceneMetaReq{
		InstanceId:        suite.instanceId,
		ModifyAPPMetaType: AppMetaTypeAiAssistantText2SqlFeedBack,
		Id:                suite.topicId,
		Record: MetaRecord{
			FeedBackMeta: &FeedBackMeta{
				SessionId:          "1",
				MessageId:          "1",
				FeedBackType:       FeedBackTypePositive,
				AiAssistantFeature: FeatureText2Sql,
			},
		},
	})
	suite.NoError(err)
}

func (suite *AppSceneMetaTestSuite) TestModifyAppSceneMetaAbnormally() {
	testcases := map[*ModifyAppSceneMetaReq]*Error{
		{ // 无效场景ID
			ModifyAPPMetaType: "invalid--type",
			InstanceId:        suite.instanceId,
			Id:                suite.topicId,
			Record: MetaRecord{
				FeedBackMeta: &FeedBackMeta{
					SessionId:          "1",
					MessageId:          "1",
					FeedBackType:       FeedBackTypePositive,
					AiAssistantFeature: FeatureText2Sql,
				},
			},
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
		},

		{ // 无效的实例id
			ModifyAPPMetaType: AppMetaTypeAiAssistantSession,
			InstanceId:        "invalid-instance-id",
			Id:                suite.topicId,
			Record: MetaRecord{
				FeedBackMeta: &FeedBackMeta{
					SessionId:          "1",
					MessageId:          "1",
					FeedBackType:       FeedBackTypePositive,
					AiAssistantFeature: FeatureText2Sql,
				},
			},
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "AppInstanceNotExists",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.ModifyAppSceneMeta(req)
		var tlsErr *Error
		ok := errors.As(err, &tlsErr)
		suite.True(ok)
		suite.Equal(tlsErr.Code, expectedErr.Code)
		suite.Equal(tlsErr.HTTPCode, expectedErr.HTTPCode)
	}
}

func (suite *AppSceneMetaTestSuite) TestDescribeAppSceneMetasNormally() {
	// 创建会话
	createRsp, err := suite.cli.CreateAppSceneMeta(&CreateAppSceneMetaReq{
		InstanceId:        suite.instanceId,
		CreateAPPMetaType: AppMetaTypeAiAssistantSession,
		Id:                suite.topicId,
	})
	suite.NoError(err)
	suite.NotEmpty(createRsp.Id)
	sessionId := createRsp.Id

	// 查会话列表
	resp, err := suite.cli.DescribeAppSceneMetas(&DescribeAppSceneMetasReq{
		InstanceId: suite.instanceId,
		// 传topicId
		Id:                  &suite.topicId,
		DescribeAPPMetaType: AppMetaTypeAiAssistantSession,
		PageNumber:          1,
		PageSize:            10,
	})
	suite.NoError(err)
	for _, item := range resp.Items {
		suite.NotNil(item.DescribeSessionMeta)
	}

	// 查历史消息
	_, err = suite.cli.DescribeAppSceneMetas(&DescribeAppSceneMetasReq{
		InstanceId: suite.instanceId,
		// 会话id
		Id:                  &sessionId,
		DescribeAPPMetaType: AppMetaTypeAiAssistantHistoryMessage,
	})
	suite.NoError(err)

	// 查问题推荐
	resp, err = suite.cli.DescribeAppSceneMetas(&DescribeAppSceneMetasReq{
		InstanceId: suite.instanceId,
		// topicId
		Id:                  &suite.topicId,
		DescribeAPPMetaType: AppMetaTypeAiAssistantText2SqlSuggestion,
	})
	suite.NoError(err)
	for _, item := range resp.Items {
		suite.NotEmpty(item.DescribeSessionSuggestion)
	}
}

type SessionAnswerTestSuite struct {
	suite.Suite

	cli        Client
	project    string
	topic      string
	instanceId string
}

func (suite *SessionAnswerTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	// 创建测试用项目
	projectId, err := CreateProject("copilot-test-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	// 创建Topic
	topicId, err := CreateTopic(projectId, "copilot-topic-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	// 创建ai应用实例
	instanceType := AppInstanceTypeAiAssistant
	accountId := os.Getenv("ACCOUNT_ID")
	describeRsp, err := suite.cli.DescribeAppInstances(&DescribeAppInstancesReq{
		PageNumber:   1,
		PageSize:     10,
		InstanceType: &instanceType,
		InstanceName: &accountId,
	})
	suite.NoError(err)
	for _, instance := range describeRsp.InstanceInfo {
		if instance.InstanceName == accountId {
			suite.instanceId = instance.InstanceId
		}
	}
	if suite.instanceId == "" {
		instanceReq := &CreateAppInstanceReq{
			InstanceType: AppInstanceTypeAiAssistant,
			InstanceName: os.Getenv("ACCOUNT_ID"),
		}
		instanceResp, err := suite.cli.CreateAppInstance(instanceReq)
		suite.NoError(err)
		suite.instanceId = instanceResp.InstanceID
	}
}

func (suite *SessionAnswerTestSuite) TearDownTest() {
	// 清理资源
	if suite.topic != "" {
		_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
		suite.NoError(err)
	}
	if suite.project != "" {
		_, err := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
		suite.NoError(err)
	}
}

func TestSessionAnswerTestSuite(t *testing.T) {
	suite.Run(t, new(SessionAnswerTestSuite))
}

func (suite *SessionAnswerTestSuite) TestDescribeSessionAnswerNormally() {
	// 创建会话
	createReq := &CreateAppSceneMetaReq{
		InstanceId:        suite.instanceId,
		Id:                suite.topic,
		CreateAPPMetaType: AppMetaTypeAiAssistantSession,
	}
	resp, err := suite.cli.CreateAppSceneMeta(createReq)
	suite.NoError(err)
	suite.NotEmpty(resp.Id)
	sessionId := resp.Id

	req := &DescribeSessionAnswerReq{
		InstanceId: suite.instanceId,
		TopicId:    suite.topic,
		SessionId:  sessionId,
		Question:   "如何配置日志采集？",
	}

	// 调用流式接口
	reader, err := suite.cli.DescribeSessionAnswer(req)
	suite.NoError(err)
	defer reader.Close()

	// 验证SSE流数据
	eventCount := 0
	var (
		answerSb    strings.Builder
		reasoningSb strings.Builder
	)
	for {
		eventCount++
		event, err := reader.ReadEvent()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}
		suite.NoError(err)

		suite.NotEmpty(event.QuestionId)
		suite.NotEmpty(event.SessionId)
		suite.NotEmpty(event.MessageId)

		if event.RspMsgType == AgentRspMsgTypeIntentRecognition {
			suite.NotNil(event.IntentInfo)
			fmt.Printf("意图识别结果: %s\n", event.IntentInfo.Name)
		}
		if event.RspMsgType == AgentRspMsgTypeToolCalling {
			suite.NotNil(event.Tools)
			fmt.Printf("工具调用结果: %s\n", event.Tools)
		}
		if event.RspMsgType == AgentRspMsgTypeInference {
			suite.NotNil(event.ModelAnswer)
			answerSb.WriteString(event.ModelAnswer.Answer)
		}
		if event.RspMsgType == AgentRspMsgTypeQuestionsSuggestions {
			suite.NotNil(event.Suggestions)
			fmt.Printf("问题建议: %s\n", event.Suggestions)
		}
		if event.RspMsgType == AgentRspMsgTypeRetrieval {
			suite.NotNil(event.RetrievalInfo)
			fmt.Printf("检索结果: %s\n", event.RetrievalInfo)
		}
		if event.RspMsgType == AgentRspMsgTypeReasoning {
			suite.NotNil(event.ReasoningContent)
			reasoningSb.WriteString(event.ReasoningContent.Answer)
		}
	}
	fmt.Printf("思考过程:\n%s\n", reasoningSb.String())
	fmt.Printf("最终答案:\n%s\n", answerSb.String())
	suite.Greater(eventCount, 0)
}

func (suite *SessionAnswerTestSuite) TestDescribeSessionAnswerAbnormally() {
	// 创建会话
	createReq := &CreateAppSceneMetaReq{
		InstanceId:        suite.instanceId,
		Id:                suite.topic,
		CreateAPPMetaType: AppMetaTypeAiAssistantSession,
	}
	resp, err := suite.cli.CreateAppSceneMeta(createReq)
	suite.NoError(err)
	suite.NotEmpty(resp.Id)
	sessionId := resp.Id

	testcases := map[*DescribeSessionAnswerReq]*Error{
		{ // 无效实例ID
			InstanceId: "invalid_instance",
			TopicId:    suite.topic,
			SessionId:  sessionId,
			Question:   "test",
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "AppInstanceNotExists",
			Message:  "App instance does not exist",
		},

		{ // 无效TopicID
			InstanceId: suite.instanceId,
			TopicId:    "invalid_topic",
			SessionId:  sessionId,
			Question:   "test",
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key topicId, value invalid_topic, please check argument.",
		},

		{ // 无效的SessionID
			InstanceId: suite.instanceId,
			TopicId:    suite.topic,
			SessionId:  uuid.New().String(),
			Question:   "test",
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key SessionId",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeSessionAnswer(req)
		var tlsErr *Error
		if errors.As(err, &tlsErr) {
			suite.Equal(expectedErr.HTTPCode, tlsErr.HTTPCode)
			suite.Equal(expectedErr.Code, tlsErr.Code)
			suite.Contains(tlsErr.Message, expectedErr.Message)
		} else {
			suite.Fail("unexpected error type")
		}
	}
}
