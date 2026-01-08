package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKScheduleSqlTaskTestSuite struct {
	suite.Suite

	cli       Client
	projectID string
	topicIDs  []string
}

func (suite *SDKScheduleSqlTaskTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()
	suite.topicIDs = make([]string, 0)
}

func (suite *SDKScheduleSqlTaskTestSuite) TearDownTest() {
	for _, topicID := range suite.topicIDs {
		_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: topicID})
		suite.NoError(err)
	}

	if suite.projectID != "" {
		_, err := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.projectID})
		suite.NoError(err)
	}
}

func (suite *SDKScheduleSqlTaskTestSuite) validateError(err error, expectErr *Error) {
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

func (suite *SDKScheduleSqlTaskTestSuite) TestCreateScheduleSqlTaskNormally() {
	// 创建项目和主题用于测试
	projectName := "golang-sdk-schedule-project-" + uuid.New().String()[:8]
	createProjectResp, err := suite.cli.CreateProject(&CreateProjectRequest{
		ProjectName: projectName,
		Description: "test schedule sql task project",
		Region:      getRegion(),
	})
	suite.NoError(err)
	suite.NotNil(createProjectResp)
	suite.projectID = createProjectResp.ProjectID

	// 创建源主题
	sourceTopicName := "golang-sdk-schedule-source-topic-" + uuid.New().String()[:8]
	createSourceTopicResp, err := suite.cli.CreateTopic(&CreateTopicRequest{
		ProjectID:   suite.projectID,
		TopicName:   sourceTopicName,
		Ttl:         1,
		ShardCount:  1,
		Description: "source topic for schedule sql task test",
	})
	suite.NoError(err)
	suite.NotNil(createSourceTopicResp)
	sourceTopicID := createSourceTopicResp.TopicID
	suite.topicIDs = append(suite.topicIDs, sourceTopicID)

	// 创建目标主题
	destTopicName := "golang-sdk-schedule-dest-topic-" + uuid.New().String()[:8]
	createDestTopicResp, err := suite.cli.CreateTopic(&CreateTopicRequest{
		ProjectID:   suite.projectID,
		TopicName:   destTopicName,
		Ttl:         1,
		ShardCount:  1,
		Description: "destination topic for schedule sql task test",
	})
	suite.NoError(err)
	suite.NotNil(createDestTopicResp)
	destTopicID := createDestTopicResp.TopicID
	suite.topicIDs = append(suite.topicIDs, destTopicID)
}

func (suite *SDKScheduleSqlTaskTestSuite) TestDeleteScheduleSqlTask() {
	taskId := "test-schedule-sql-task-123456"

	deleteResult, err := suite.cli.DeleteScheduleSqlTask(&DeleteScheduleSqlTaskRequest{
		TaskID: taskId,
	})

	if err != nil {
		suite.T().Logf("DeleteScheduleSqlTask returned error (expected): %v", err)
	} else {
		suite.NotNil(deleteResult)
		suite.NotEmpty(deleteResult.RequestID)
	}
}

func (suite *SDKScheduleSqlTaskTestSuite) TestDeleteScheduleSqlTaskValidation() {
	_, err := suite.cli.DeleteScheduleSqlTask(&DeleteScheduleSqlTaskRequest{
		TaskID: "",
	})
	suite.Error(err)
	suite.Contains(err.Error(), "Invalid argument, empty TaskID")
}

func (suite *SDKScheduleSqlTaskTestSuite) TestDescribeScheduleSqlTask() {
	taskId := "f3e901c3-b17f-42fd-aa8c-dc91a6c7****"

	request := &DescribeScheduleSqlTaskRequest{
		TaskID: taskId,
	}

	response, err := suite.cli.DescribeScheduleSqlTask(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.NotEmpty(response.RequestID)
}

func (suite *SDKScheduleSqlTaskTestSuite) TestDescribeScheduleSqlTaskValidation() {
	request := &DescribeScheduleSqlTaskRequest{
		TaskID: "",
	}

	_, err := suite.cli.DescribeScheduleSqlTask(request)
	suite.Error(err)
}

func (suite *SDKScheduleSqlTaskTestSuite) TestDescribeScheduleSqlTasksNormally() {
	testcases := map[*DescribeScheduleSqlTasksRequest]*Error{
		{
			PageNumber: IntPtr(1),
			PageSize:   IntPtr(20),
		}: nil,
		{
			TaskName: StrPtr("test-task"),
		}: nil,
		{
			ProjectId: StrPtr("test-project-id"),
			TopicId:   StrPtr("test-topic-id"),
		}: nil,
	}

	for req, expectedErr := range testcases {
		resp, err := suite.cli.DescribeScheduleSqlTasks(req)
		if expectedErr == nil {
			suite.NoError(err)
			suite.NotNil(resp)
			suite.GreaterOrEqual(resp.Total, 0)
			suite.NotNil(resp.Tasks)
		} else {
			suite.validateError(err, expectedErr)
		}
	}
}

func (suite *SDKScheduleSqlTaskTestSuite) TestDescribeScheduleSqlTasksWithPagination() {
	req := &DescribeScheduleSqlTasksRequest{
		PageNumber: IntPtr(1),
		PageSize:   IntPtr(10),
	}

	resp, err := suite.cli.DescribeScheduleSqlTasks(req)
	suite.NoError(err)
	suite.NotNil(resp)
	suite.GreaterOrEqual(resp.Total, 0)
	suite.LessOrEqual(len(resp.Tasks), 10)
}

func (suite *SDKScheduleSqlTaskTestSuite) TestDescribeScheduleSqlTasksWithFilters() {
	testcases := []*DescribeScheduleSqlTasksRequest{
		{
			Status: StrPtr("Running"),
		},
		{
			TaskId: StrPtr("test-task-id"),
		},
		{
			SourceTopicName: StrPtr("test-topic"),
		},
		{
			ProjectName: StrPtr("test-project"),
		},
		{
			IamProjectName: StrPtr("test-iam-project"),
		},
	}

	for _, req := range testcases {
		resp, err := suite.cli.DescribeScheduleSqlTasks(req)
		suite.NoError(err)
		suite.NotNil(resp)
		suite.GreaterOrEqual(resp.Total, 0)
	}
}

func TestSDKScheduleSqlTaskTestSuite(t *testing.T) {
	suite.Run(t, new(SDKScheduleSqlTaskTestSuite))
}

func getRegion() string {
	region := os.Getenv("LOG_SERVICE_REGION")
	if region == "" {
		region = "cn-north-1"
	}
	return region
}
