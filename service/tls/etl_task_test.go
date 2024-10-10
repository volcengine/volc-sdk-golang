package tls

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKETLTestSuite struct {
	suite.Suite

	cli         Client
	project     string
	topic       string
	destTopic   string
	ETLTaskList []string
}

func (suite *SDKETLTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-project-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)

	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-topic-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	destTopic, err := CreateTopic(projectId, "golang-sdk-create-topic-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.destTopic = destTopic

}

func (suite *SDKETLTestSuite) TearDownTest() {
	for _, taskID := range suite.ETLTaskList {
		_, err := suite.cli.DeleteETLTask(&DeleteETLTaskRequest{TaskID: taskID})
		suite.NoError(err)
	}

	_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(err)

	_, err = suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.destTopic})
	suite.NoError(err)

	_, err = suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(err)

	suite.ETLTaskList = nil
}

func (suite *SDKETLTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKETLTestSuite(t *testing.T) {
	suite.Run(t, new(SDKETLTestSuite))
}

func (suite *SDKETLTestSuite) TestCreateETLTaskNormally() {

	req := &CreateETLTaskRequest{
		DSLType:       "NORMAL",
		Description:   StrPtr("ETL test"),
		Enable:        false,
		Name:          "test-etl-task",
		Script:        "f_set(\"key\",\"value\")",
		SourceTopicId: suite.topic,
		TargetResources: []TargetResource{
			{
				Alias:   "TOPIC",
				TopicID: suite.destTopic,
			},
		},
		TaskType: "Resident",
	}

	createTopicResp, err := suite.cli.CreateETLTask(req)
	suite.NoError(err)
	suite.ETLTaskList = append(suite.ETLTaskList, createTopicResp.TaskID)

	getETLTaskResponse, err := suite.cli.DescribeETLTask(&DescribeETLTaskRequest{TaskID: createTopicResp.TaskID})
	suite.NoError(err)
	suite.Equal(createTopicResp.TaskID, getETLTaskResponse.TaskID)
	suite.Equal(req.SourceTopicId, getETLTaskResponse.SourceTopicID)
	suite.Equal(req.Script, getETLTaskResponse.Script)
	suite.Equal(req.Enable, getETLTaskResponse.Enable)
	suite.Equal(*req.Description, getETLTaskResponse.Description)
	suite.Equal(req.Name, getETLTaskResponse.Name)
	if req.FromTime != nil {
		suite.Equal(*req.FromTime, *getETLTaskResponse.FromTime)
	} else {
		suite.Nil(getETLTaskResponse.FromTime)
	}
	if req.ToTime != nil {
		suite.Equal(*req.ToTime, *getETLTaskResponse.ToTime)
	} else {
		suite.Nil(getETLTaskResponse.ToTime)
	}
	suite.Equal(len(req.TargetResources), len(getETLTaskResponse.TargetResources))

}

func (suite *SDKETLTestSuite) TestCreateETLTaskAbnormally() {
	taskName := "test-etl-task-name"
	resp, err := suite.cli.CreateETLTask(&CreateETLTaskRequest{
		DSLType:       "NORMAL",
		Description:   StrPtr("ETL test 1"),
		Enable:        false,
		Name:          taskName,
		Script:        "f_set(\"key1\",\"value1\")",
		SourceTopicId: suite.topic,
		TargetResources: []TargetResource{
			{
				Alias:   "TOPIC",
				TopicID: suite.destTopic,
			},
		},
		TaskType: "Resident",
	})
	suite.NoError(err)
	suite.ETLTaskList = append(suite.ETLTaskList, resp.TaskID)

	resp, err = suite.cli.CreateETLTask(&CreateETLTaskRequest{
		DSLType:       "NORMAL",
		Description:   StrPtr("ETL test 2"),
		Enable:        false,
		Name:          taskName,
		Script:        "f_set(\"key2\",\"value2\")",
		SourceTopicId: suite.topic,
		TargetResources: []TargetResource{
			{
				Alias:   "TOPIC",
				TopicID: suite.destTopic,
			},
		},
		TaskType: "Resident",
	})
	expectedErr := &Error{
		HTTPCode: http.StatusConflict,
		Code:     "EtlTaskExist",
		Message:  "etl task exist",
	}
	suite.validateError(err, expectedErr)

	if err == nil {
		suite.ETLTaskList = append(suite.ETLTaskList, resp.TaskID)
	}

	resp, err = suite.cli.CreateETLTask(&CreateETLTaskRequest{
		DSLType:       "NORMAL1",
		Description:   StrPtr("ETL test 3"),
		Enable:        false,
		Name:          "taskName-3",
		Script:        "f_set(\"key3\",\"value3\")",
		SourceTopicId: suite.topic,
		TargetResources: []TargetResource{
			{
				Alias:   "TOPIC",
				TopicID: suite.destTopic,
			},
		},
		TaskType: "Resident",
	})
	expectedErr = &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key DSLType, value NORMAL1, please check argument.",
	}
	suite.validateError(err, expectedErr)

	if err == nil {
		suite.ETLTaskList = append(suite.ETLTaskList, resp.TaskID)
	}
}

func (suite *SDKETLTestSuite) TestDeleteETLTaskAbnormally() {
	TaskID := uuid.New().String()
	_, err := suite.cli.DeleteETLTask(&DeleteETLTaskRequest{TaskID: TaskID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "TaskNotExist",
		Message:  "Task " + TaskID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKETLTestSuite) TestModifyETLTaskNormally() {
	req := &CreateETLTaskRequest{
		DSLType:       "NORMAL",
		Description:   StrPtr("ETL test 4"),
		Enable:        false,
		Name:          "taskName-4",
		Script:        "f_set(\"key4\",\"value4\")",
		SourceTopicId: suite.topic,
		TargetResources: []TargetResource{
			{
				Alias:   "TOPIC",
				TopicID: suite.destTopic,
			},
		},
		TaskType: "Resident",
	}

	resp, err := suite.cli.CreateETLTask(req)
	suite.NoError(err)
	suite.ETLTaskList = append(suite.ETLTaskList, resp.TaskID)

	newScript := "f_set(\"key4-1\",\"value4-1\")"
	_, err = suite.cli.ModifyETLTask(&ModifyETLTaskRequest{
		TaskID: resp.TaskID,
		Script: &newScript,
	})
	suite.NoError(err)
	describeResp, err := suite.cli.DescribeETLTask(&DescribeETLTaskRequest{TaskID: resp.TaskID})
	suite.NoError(err)
	suite.Equal(req.Name, describeResp.Name)
	suite.Equal(*req.Description, describeResp.Description)
	suite.Equal(newScript, describeResp.Script)
	suite.Equal(req.Enable, describeResp.Enable)
	suite.Equal(len(req.TargetResources), len(describeResp.TargetResources))

	newDesc := "ETL test 4-1"
	_, err = suite.cli.ModifyETLTask(&ModifyETLTaskRequest{
		TaskID:      resp.TaskID,
		Description: &newDesc,
		TargetResources: []TargetResource{
			{
				Alias:   "TOPIC1",
				TopicID: suite.destTopic,
			},
		},
	})
	suite.NoError(err)
	describeResp, err = suite.cli.DescribeETLTask(&DescribeETLTaskRequest{TaskID: resp.TaskID})
	suite.NoError(err)
	suite.Equal(req.Name, describeResp.Name)
	suite.Equal(newDesc, describeResp.Description)
	suite.Equal(newScript, describeResp.Script)
	suite.Equal(req.Enable, describeResp.Enable)
	suite.Equal(len(req.TargetResources), len(describeResp.TargetResources))
}

func (suite *SDKETLTestSuite) TestModifyETLTaskAbnormally() {
	req := &CreateETLTaskRequest{
		DSLType:       "NORMAL",
		Description:   StrPtr("ETL test 5"),
		Enable:        false,
		Name:          "taskName-5",
		Script:        "f_set(\"key5\",\"value5\")",
		SourceTopicId: suite.topic,
		TargetResources: []TargetResource{
			{
				Alias:   "TOPIC",
				TopicID: suite.destTopic,
			},
		},
		TaskType: "Resident",
	}

	resp, err := suite.cli.CreateETLTask(req)
	suite.NoError(err)
	suite.ETLTaskList = append(suite.ETLTaskList, resp.TaskID)

	_, err = suite.cli.ModifyETLTask(&ModifyETLTaskRequest{
		TaskID:          resp.TaskID,
		TargetResources: []TargetResource{{}},
	})

	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key TargetResources, value , please check argument.",
	}
	suite.validateError(err, expectedErr)

	_, err = suite.cli.ModifyETLTaskStatus(&ModifyETLTaskStatusRequest{
		TaskID: resp.TaskID,
		Enable: true,
	})
	suite.NoError(err)
	describeResp, err := suite.cli.DescribeETLTask(&DescribeETLTaskRequest{TaskID: resp.TaskID})
	suite.NoError(err)
	suite.Equal(true, describeResp.Enable)
}

func (suite *SDKETLTestSuite) TestDescribeETLTaskAbnormally() {
	TaskID := uuid.New().String()
	_, err := suite.cli.DescribeETLTask(&DescribeETLTaskRequest{TaskID: TaskID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "TaskNotExist",
		Message:  "Task " + TaskID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKETLTestSuite) TestDescribeETLTasksNormally() {
	targets := []TargetResource{
		{
			Alias:   "TOPIC",
			TopicID: suite.destTopic,
		},
	}

	createReqs := []*CreateETLTaskRequest{
		{
			DSLType:         "NORMAL",
			Description:     StrPtr("ETL test"),
			Enable:          false,
			Name:            "task-namea-" + uuid.New().String(),
			Script:          "f_set(\"key\",\"value\")",
			SourceTopicId:   suite.topic,
			TargetResources: targets,
			TaskType:        "Resident",
		},
		{
			DSLType:         "NORMAL",
			Description:     StrPtr("ETL test"),
			Enable:          false,
			Name:            "task-namea-" + uuid.New().String(),
			Script:          "f_set(\"key\",\"value\")",
			SourceTopicId:   suite.topic,
			TargetResources: targets,
			TaskType:        "Resident",
		},
		{
			DSLType:         "NORMAL",
			Description:     StrPtr("ETL test"),
			Enable:          false,
			Name:            "task-name-" + uuid.New().String(),
			Script:          "f_set(\"key\",\"value\")",
			SourceTopicId:   suite.topic,
			TargetResources: targets,
			TaskType:        "Resident",
		},
		{
			DSLType:         "NORMAL",
			Description:     StrPtr("ETL test"),
			Enable:          false,
			Name:            "task-name-" + uuid.New().String(),
			Script:          "f_set(\"key\",\"value\")",
			SourceTopicId:   suite.topic,
			TargetResources: targets,
			TaskType:        "Resident",
		},
	}

	for _, createReq := range createReqs {
		createResp, err := suite.cli.CreateETLTask(createReq)
		suite.NoError(err)
		suite.ETLTaskList = append(suite.ETLTaskList, createResp.TaskID)
	}

	testcases := map[*DescribeETLTasksRequest]*DescribeETLTasksResponse{
		{
			ProjectID: suite.project,
		}: {
			Tasks: []ETLTaskResponse{
				{
					Name:        createReqs[3].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[3].Enable,
					Description: *createReqs[3].Description,
				},
				{
					Name:        createReqs[2].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[2].Enable,
					Description: *createReqs[2].Description,
				},
				{
					Name:        createReqs[1].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[1].Enable,
					Description: *createReqs[1].Description,
				},
				{
					Name:        createReqs[0].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[0].Enable,
					Description: *createReqs[0].Description,
				},
			},
			Total: 4,
		},
		{
			ProjectID: suite.project,
			PageSize:  20,
			TaskName:  "task-name-",
		}: {
			Tasks: []ETLTaskResponse{
				{
					Name:        createReqs[3].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[3].Enable,
					Description: *createReqs[3].Description,
				},
				{
					Name:        createReqs[2].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[2].Enable,
					Description: *createReqs[2].Description,
				},
			},
			Total: 2,
		},
		{
			ProjectID:       suite.project,
			PageNumber:      1,
			PageSize:        20,
			TaskName:        "task-namea-",
			SourceTopicName: "golang-sdk-create-topic",
		}: {
			Tasks: []ETLTaskResponse{
				{
					Name:        createReqs[1].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[1].Enable,
					Description: *createReqs[1].Description,
				},
				{
					Name:        createReqs[0].Name,
					ProjectID:   suite.project,
					Enable:      createReqs[0].Enable,
					Description: *createReqs[0].Description,
				},
			},
			Total: 2,
		},
	}

	for listReq, expectListResp := range testcases {
		actualListResp, err := suite.cli.DescribeETLTasks(listReq)
		suite.NoError(err)

		suite.Equal(expectListResp.Total, actualListResp.Total)
		suite.Equal(len(expectListResp.Tasks), len(actualListResp.Tasks))

		expectTasks := expectListResp.Tasks
		actualTasks := actualListResp.Tasks

		for _, expectTask := range expectTasks {
			isMatched := false
			for _, actualTask := range actualTasks {
				if expectTask.Name == actualTask.Name {
					isMatched = true
					suite.Equal(expectTask.ProjectID, actualTask.ProjectID)
					suite.Equal(expectTask.Enable, actualTask.Enable)
					suite.Equal(expectTask.Description, actualTask.Description)
					break
				}
			}
			suite.Equal(true, isMatched)
		}
	}

	resp, err := suite.cli.DescribeETLTasks(&DescribeETLTasksRequest{})
	suite.NoError(err)
	suite.GreaterOrEqual(int(resp.Total), 0)
}

func (suite *SDKETLTestSuite) TestDescribeETLTasksAbnormally() {
	_, err := suite.cli.DescribeETLTasks(&DescribeETLTasksRequest{ProjectID: uuid.New().String()})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)
}
