package tls

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

type SDKImportTestSuite struct {
	suite.Suite

	cli               Client
	tosBucketName     string
	tosObjectKey      string
	projectId         string
	topicId           string
	iamProjectName    string
	tosTaskId         string
	KafkaTaskId       string
	createParams      CreateImportTaskRequest
	createTosParams   *CreateImportTaskRequest
	createKafkaParams *CreateImportTaskRequest
}

func (suite *SDKImportTestSuite) SetupSuite() {
	suite.cli = NewClientWithEnv()
	suite.Equal("/CreateImportTask", PathCreateImportTask)
	suite.Equal("/DeleteImportTask", PathDeleteImportTask)
	suite.Equal("/ModifyImportTask", PathModifyImportTask)
	suite.Equal("/DescribeImportTask", PathDescribeImportTask)
	suite.Equal("/DescribeImportTasks", PathDescribeImportTasks)

	suite.iamProjectName = os.Getenv("IAM_PROJECT_NAME")
	suite.tosObjectKey = os.Getenv("IMPORT_TOS_OBJECT_KEY")
	suite.tosBucketName = os.Getenv("IMPORT_TOS_BUCKET")

	createProjectReq := &CreateProjectRequest{
		ProjectName:    "import-test-project" + uuid.New().String(),
		Description:    "go sdk导入需求测试项目",
		Region:         os.Getenv("LOG_SERVICE_REGION"),
		IamProjectName: &suite.iamProjectName,
	}

	createProjectResp, err := suite.cli.CreateProject(createProjectReq)
	suite.NoError(err)
	suite.projectId = createProjectResp.ProjectID

	createTopicReq := &CreateTopicRequest{
		ProjectID:      suite.projectId,
		TopicName:      "import-test-topic" + uuid.New().String(),
		Ttl:            1,
		Description:    "go sdk导入需求测试topic",
		ShardCount:     1,
		EnableTracking: BoolPtr(true),
	}

	createTopicResp, err := suite.cli.CreateTopic(createTopicReq)
	suite.NoError(err)
	suite.topicId = createTopicResp.TopicID
}

func (suite *SDKImportTestSuite) TearDownSuite() {
	_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topicId})
	suite.NoError(err)
	_, err = suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.projectId})
	suite.NoError(err)
}

func (suite *SDKImportTestSuite) SetupTest() {
	suite.createParams = CreateImportTaskRequest{
		Description:      "go sdk test create import task function",
		ImportSourceInfo: &ImportSourceInfo{},
		ProjectID:        suite.projectId,
		SourceType:       "",
		TargetInfo: &TargetInfo{
			Region:  os.Getenv("LOG_SERVICE_REGION"),
			LogType: "minimalist_log",
			ExtractRule: &ImportExtractRule{
				TimeZone:         "GMT",
				SkipLineCount:    0,
				TimeExtractRegex: "",
				ExtractRule: &ExtractRule{
					Quote: "\"",
				},
			},
		},
		TaskName: "test-" + uuid.New().String(),
		TopicID:  suite.topicId,
	}
}

func (suite *SDKImportTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKImportTestSuite(t *testing.T) {
	suite.Run(t, new(SDKImportTestSuite))
}

func (suite *SDKImportTestSuite) TestImportTaskWholeProcessOfTos() {
	taskId := suite.testCreateImportTaskByTosNormally()
	suite.testDescribeImportTaskNormally(taskId, "tos")
	suite.testDeleteImportTaskNormally([]string{taskId})
}

func (suite *SDKImportTestSuite) TestImportTaskWholeProcessOfKafka() {
	taskId := suite.testCreateImportTaskByKafkaNormally()
	suite.testDescribeImportTaskNormally(taskId, "kafka")
	suite.testModifyImportTaskByKafkaNormally(taskId)
	suite.testDeleteImportTaskNormally([]string{taskId})
}

func (suite *SDKImportTestSuite) TestDescribeImportTasksNormally() {
	reqParams := &DescribeImportTasksRequest{
		TaskID:         "",
		TaskName:       "",
		ProjectID:      suite.projectId,
		ProjectName:    "",
		IamProjectName: "",
		TopicID:        suite.topicId,
		TopicName:      "",
		PageNumber:     0,
		PageSize:       10,
		SourceType:     "",
		Status:         "",
	}

	res, err := suite.cli.DescribeImportTasks(reqParams)
	suite.NoError(err)
	suite.Equal(0, res.Total)

	tosTaskId := suite.testCreateImportTaskByTosNormally()
	kafkaTaskId := suite.testCreateImportTaskByKafkaNormally()

	res, err = suite.cli.DescribeImportTasks(reqParams)
	suite.NoError(err)

	var taskId string
	for _, taskInfoItem := range res.TaskInfo {

		sourceType := taskInfoItem.SourceType

		switch sourceType {
		case "tos":
			taskId = tosTaskId
		case "kafka":
			taskId = kafkaTaskId
		default:
			suite.T().Errorf("Unexpected value for SourceType: %s", sourceType)
			return
		}

		suite.testDescribeImportTaskReturnParams(taskInfoItem, taskId, sourceType)
	}

	suite.Equal(2, res.Total)
}

func (suite *SDKImportTestSuite) testCreateImportTaskByTosNormally() string {
	suite.createParams.TaskName = "test-" + uuid.New().String()
	params := suite.createParams
	params.SourceType = "tos"
	params.ImportSourceInfo.TosSourceInfo = &TosSourceInfo{
		Bucket:       suite.tosBucketName,
		Prefix:       suite.tosObjectKey,
		Region:       os.Getenv("LOG_SERVICE_REGION"),
		CompressType: "none",
	}

	suite.createTosParams = &params

	res, err := suite.cli.CreateImportTask(&params)
	suite.NoError(err)

	suite.NotNil(res.TaskID)

	return res.TaskID
}

func (suite *SDKImportTestSuite) testCreateImportTaskByKafkaNormally() string {
	suite.createParams.TaskName = "test-" + uuid.New().String()
	params := suite.createParams
	params.SourceType = "kafka"
	params.ImportSourceInfo.KafkaSourceInfo = &KafkaSourceInfo{
		Host:              os.Getenv("IMPORT_KAFKA_HOST"),
		Group:             "group-test-origin",
		Topic:             "app",
		Encode:            "UTF-8",
		Password:          "",
		Protocol:          "",
		Username:          "",
		Mechanism:         "",
		InstanceID:        "",
		IsNeedAuth:        false,
		InitialOffset:     0,
		TimeSourceDefault: 0,
	}

	suite.createKafkaParams = &params

	res, err := suite.cli.CreateImportTask(&params)
	suite.NoError(err)
	suite.NotNil(res.TaskID)

	return res.TaskID
}

func (suite *SDKImportTestSuite) testModifyImportTaskByKafkaNormally(taskId string) {
	modifyImportTaskRequest := &ModifyImportTaskRequest{
		TopicID:     suite.topicId,
		TaskID:      taskId,
		ProjectID:   "just-for-test",
		TaskName:    "test-modify-taskName-" + uuid.New().String(),
		Status:      5,
		SourceType:  "kafka",
		Description: nil,
		ImportSourceInfo: &ImportSourceInfo{
			TosSourceInfo: nil,
			KafkaSourceInfo: &KafkaSourceInfo{
				Host:              os.Getenv("IMPORT_KAFKA_HOST"),
				Group:             "group-test-modify",
				Topic:             "app",
				Encode:            "UTF-8",
				Password:          "",
				Protocol:          "",
				Username:          "",
				Mechanism:         "",
				InstanceID:        "",
				IsNeedAuth:        false,
				InitialOffset:     0,
				TimeSourceDefault: 0,
			},
		},
		TargetInfo: &TargetInfo{
			Region:      os.Getenv("LOG_SERVICE_REGION"),
			LogType:     "minimalist_log",
			ExtractRule: nil,
		},
	}

	_, err := suite.cli.ModifyImportTask(modifyImportTaskRequest)
	suite.NoError(err)

	reqParams := &DescribeImportTaskRequest{
		TaskID: taskId,
	}

	res, err := suite.cli.DescribeImportTask(reqParams)
	suite.NoError(err)
	suite.Equal(suite.projectId, res.TaskInfo.ProjectID)
	suite.Equal(suite.createKafkaParams.Description, res.TaskInfo.Description)
	suite.NotNil(res.TaskInfo.ImportSourceInfo.KafkaSourceInfo)
	suite.Equal(modifyImportTaskRequest.ImportSourceInfo.KafkaSourceInfo.Group, res.TaskInfo.ImportSourceInfo.KafkaSourceInfo.Group)
}

func (suite *SDKImportTestSuite) testDescribeImportTaskNormally(taskId string, importType string) {
	res, err := suite.cli.DescribeImportTask(&DescribeImportTaskRequest{TaskID: taskId})
	suite.NoError(err)
	suite.NotNil(res.TaskInfo)
	suite.testDescribeImportTaskReturnParams(res.TaskInfo, taskId, importType)
	suite.Equal(taskId, res.TaskInfo.TaskID)
}

func (suite *SDKImportTestSuite) testDeleteImportTaskNormally(taskIds []string) {
	for _, taskId := range taskIds {
		if taskId == "" {
			continue
		}
		_, err := suite.cli.DeleteImportTask(&DeleteImportTaskRequest{TaskID: taskId})
		suite.NoError(err)
	}
}

func (suite *SDKImportTestSuite) testDescribeImportTaskReturnParams(resTaskInfo *ImportTaskInfo, taskId string, importType string) {
	var createParams *CreateImportTaskRequest
	suite.NotNil(resTaskInfo.ImportSourceInfo)
	switch importType {
	case "tos":
		createParams = suite.createTosParams
		suite.NotNil(resTaskInfo.ImportSourceInfo.TosSourceInfo)
		suite.Equal(createParams.ImportSourceInfo.TosSourceInfo.Bucket, resTaskInfo.ImportSourceInfo.TosSourceInfo.Bucket)
		suite.Equal(createParams.ImportSourceInfo.TosSourceInfo.Prefix, resTaskInfo.ImportSourceInfo.TosSourceInfo.Prefix)
		suite.Equal(createParams.ImportSourceInfo.TosSourceInfo.Region, resTaskInfo.ImportSourceInfo.TosSourceInfo.Region)
		suite.Equal(createParams.ImportSourceInfo.TosSourceInfo.CompressType, resTaskInfo.ImportSourceInfo.TosSourceInfo.CompressType)
	case "kafka":
		createParams = suite.createKafkaParams
		suite.NotNil(resTaskInfo.ImportSourceInfo.KafkaSourceInfo)
		suite.Equal(createParams.ImportSourceInfo.KafkaSourceInfo.Host, resTaskInfo.ImportSourceInfo.KafkaSourceInfo.Host)
		suite.Equal(createParams.ImportSourceInfo.KafkaSourceInfo.Group, resTaskInfo.ImportSourceInfo.KafkaSourceInfo.Group)
		suite.Equal(createParams.ImportSourceInfo.KafkaSourceInfo.Topic, resTaskInfo.ImportSourceInfo.KafkaSourceInfo.Topic)
		suite.Equal(createParams.ImportSourceInfo.KafkaSourceInfo.Encode, resTaskInfo.ImportSourceInfo.KafkaSourceInfo.Encode)
		suite.False(resTaskInfo.ImportSourceInfo.KafkaSourceInfo.IsNeedAuth)
		suite.Equal("", resTaskInfo.ImportSourceInfo.KafkaSourceInfo.Password)
		suite.Equal("", resTaskInfo.ImportSourceInfo.KafkaSourceInfo.Username)
		suite.Equal("", resTaskInfo.ImportSourceInfo.KafkaSourceInfo.Mechanism)
	default:
		suite.T().Errorf("Unexpected value for SourceType: %s", importType)
		return
	}

	suite.Equal(taskId, resTaskInfo.TaskID)
	suite.Equal(createParams.TopicID, resTaskInfo.TopicID)
	suite.Equal(createParams.TaskName, resTaskInfo.TaskName)
	suite.Equal(createParams.ProjectID, resTaskInfo.ProjectID)
	suite.NotNil(resTaskInfo.TopicName)
	suite.NotNil(resTaskInfo.CreateTime)
	suite.Equal(importType, resTaskInfo.SourceType)
	suite.Equal(createParams.Description, resTaskInfo.Description)
	suite.NotNil(resTaskInfo.ProjectName)

	suite.NotNil(resTaskInfo.TargetInfo)
	suite.Equal(os.Getenv("LOG_SERVICE_REGION"), resTaskInfo.TargetInfo.Region)
	suite.Equal(createParams.TargetInfo.LogType, resTaskInfo.TargetInfo.LogType)
	suite.NotNil(resTaskInfo.TargetInfo.ExtractRule)
	suite.Equal(createParams.TargetInfo.ExtractRule.TimeZone, resTaskInfo.TargetInfo.ExtractRule.TimeZone)

	suite.NotNil(resTaskInfo.TaskStatistics)
	suite.NotNil(resTaskInfo.TaskStatistics.TaskStatus)
	suite.NotNil(resTaskInfo.TaskStatistics.Transferred)
}

func (suite *SDKImportTestSuite) TestCreateImportTaskByTosAbnormally1() {
	params := suite.createParams
	params.SourceType = "tos1"
	params.ImportSourceInfo.TosSourceInfo = &TosSourceInfo{
		Bucket:       suite.tosBucketName,
		Prefix:       suite.tosObjectKey,
		Region:       os.Getenv("LOG_SERVICE_REGION"),
		CompressType: "none",
	}

	suite.createTosParams = &params

	res, err := suite.cli.CreateImportTask(&params)

	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "ImportTaskTypeUnimplemented",
		Message:  "Import task type: %s unimplemented",
	}

	suite.validateError(err, expectedErr)

	if err == nil {
		suite.testDeleteImportTaskNormally([]string{res.TaskID})
	}
}

func (suite *SDKImportTestSuite) TestCreateImportTaskByTosAbnormally2() {
	params := suite.createParams
	params.SourceType = "tos"
	params.ImportSourceInfo.TosSourceInfo = &TosSourceInfo{
		Bucket:       suite.tosBucketName,
		Prefix:       suite.tosObjectKey,
		Region:       os.Getenv("LOG_SERVICE_REGION"),
		CompressType: "none",
	}
	suite.createTosParams = &params
	res, err := suite.cli.CreateImportTask(&params)
	suite.NoError(err)

	taskIds := []string{res.TaskID}

	res, err = suite.cli.CreateImportTask(&params)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "DuplicateImportTask",
		Message:  "Duplicate import task name within one project",
	}

	suite.validateError(err, expectedErr)

	if err == nil {
		taskIds = append(taskIds, res.TaskID)
	}

	suite.testDeleteImportTaskNormally(taskIds)
}

func (suite *SDKImportTestSuite) TestCreateImportTaskByKafkaAbnormally1() {
	params := suite.createParams
	params.SourceType = "kafka1"
	params.ImportSourceInfo.KafkaSourceInfo = &KafkaSourceInfo{
		Host:              os.Getenv("IMPORT_KAFKA_HOST"),
		Group:             "group-test-origin",
		Topic:             "app",
		Encode:            "UTF-8",
		Password:          "",
		Protocol:          "",
		Username:          "",
		Mechanism:         "",
		InstanceID:        "",
		IsNeedAuth:        false,
		InitialOffset:     0,
		TimeSourceDefault: 0,
	}

	suite.createKafkaParams = &params

	res, err := suite.cli.CreateImportTask(&params)

	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "ImportTaskTypeUnimplemented",
		Message:  "Import task type: %s unimplemented",
	}

	suite.validateError(err, expectedErr)

	if err == nil {
		suite.testDeleteImportTaskNormally([]string{res.TaskID})
	}
}

func (suite *SDKImportTestSuite) TestDeleteImportTaskAbnormally() {
	taskId := "test-no-this-task-id"

	_, err := suite.cli.DeleteImportTask(&DeleteImportTaskRequest{TaskID: taskId})

	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key TaskID, value test-no-this-task-id, please check argument.",
	}

	suite.validateError(err, expectedErr)
}

func (suite *SDKImportTestSuite) TestModifyImportTaskAbnormally() {
	modifyImportTaskRequest := &ModifyImportTaskRequest{
		TaskID:   "No-this-task-taskId",
		TopicID:  suite.topicId,
		TaskName: "test-task-name-this-name-is-not-will-be-modify",
		Status:   5,
	}
	_, err := suite.cli.ModifyImportTask(modifyImportTaskRequest)

	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key TaskID, value No-this-task-taskId, please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKImportTestSuite) TestModifyImportTaskByTosAbnormally1() {
	createParams := suite.createParams
	createParams.SourceType = "kafka"
	createParams.ImportSourceInfo.KafkaSourceInfo = &KafkaSourceInfo{
		Host:              os.Getenv("IMPORT_KAFKA_HOST"),
		Group:             "group-test-origin",
		Topic:             "app",
		Encode:            "UTF-8",
		Password:          "",
		Protocol:          "",
		Username:          "",
		Mechanism:         "",
		InstanceID:        "",
		IsNeedAuth:        false,
		InitialOffset:     0,
		TimeSourceDefault: 0,
	}

	resCreate, err := suite.cli.CreateImportTask(&createParams)
	suite.NoError(err)

	modifyParams := &ModifyImportTaskRequest{
		TopicID:    suite.topicId,
		TaskName:   createParams.TaskName,
		TaskID:     resCreate.TaskID,
		Status:     5,
		SourceType: "tos",
		ImportSourceInfo: &ImportSourceInfo{
			TosSourceInfo: &TosSourceInfo{
				Bucket:       suite.tosBucketName,
				Prefix:       suite.tosObjectKey,
				Region:       os.Getenv("LOG_SERVICE_REGION"),
				CompressType: "none",
			},
			KafkaSourceInfo: nil,
		},
		TargetInfo: &TargetInfo{
			Region:      os.Getenv("LOG_SERVICE_REGION"),
			LogType:     "minimalist_log",
			ExtractRule: nil,
		},
	}

	_, err = suite.cli.ModifyImportTask(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key SourceType, value tos, acceptable value: kafka please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.testDeleteImportTaskNormally([]string{resCreate.TaskID})
}

func (suite *SDKImportTestSuite) TestDescribeImportTaskByTosAbnormally() {
	params := &DescribeImportTaskRequest{TaskID: "no-this-describe-task-taskId"}
	_, err := suite.cli.DescribeImportTask(params)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key TaskID, value no-this-describe-task-taskId, please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKImportTestSuite) TestDescribeImportTasksByTosAbnormally1() {
	params := &DescribeImportTasksRequest{
		TaskID:    "no-this-describe-tasks-taskId",
		ProjectID: suite.projectId,
		TopicID:   suite.topicId,
	}
	_, err := suite.cli.DescribeImportTasks(params)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key TaskID, value no-this-describe-tasks-taskId, please check argument.",
	}
	suite.validateError(err, expectedErr)
}
