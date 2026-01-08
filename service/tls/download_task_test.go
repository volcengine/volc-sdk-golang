package tls

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKDownloadTaskTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKDownloadTaskTestSuite) SetupTest() {
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
			CaseSensitive:  false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-2",
		Value: Value{
			ValueType:      "long",
			Delimiter:      "",
			CaseSensitive:  false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	suite.NoError(CreateIndex(topicId, nil, &keyValueList, suite.cli))
}

func (suite *SDKDownloadTaskTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKDownloadTaskTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKDownloadTaskTestSuite(t *testing.T) {
	suite.Run(t, new(SDKDownloadTaskTestSuite))
}

func (suite *SDKDownloadTaskTestSuite) TestCreateDownloadTaskNormally() {
	startTime := time.Now().Unix()

	testcases := map[*CreateDownloadTaskRequest]*Error{
		{
			TopicID:     suite.topic,
			TaskName:    "go-sdk-download-task",
			Query:       "*",
			StartTime:   startTime - 60,
			EndTime:     time.Now().Unix(),
			DataFormat:  "csv",
			Sort:        "asc",
			Limit:       100,
			Compression: "gzip",
			TaskType:    0,
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.CreateDownloadTask(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKDownloadTaskTestSuite) TestCreateDownloadTaskAbnormally() {
	startTime := time.Now().Unix()

	testcases := map[*CreateDownloadTaskRequest]*Error{
		{
			TopicID:     uuid.New().String(),
			TaskName:    "go-sdk-download-task",
			Query:       "*",
			StartTime:   startTime - 60,
			EndTime:     time.Now().Unix(),
			DataFormat:  "csv",
			Sort:        "asc",
			Limit:       100,
			Compression: "gzip",
			TaskType:    0,
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "TopicNotExist",
			Message:  "Topic does not exist.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.CreateDownloadTask(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKDownloadTaskTestSuite) TestDescribeDownloadTasksNormally() {
	startTime := time.Now().Unix()

	testcases := map[*CreateDownloadTaskRequest]*DescribeDownloadTasksRequest{
		{
			TopicID:     suite.topic,
			TaskName:    "go-sdk-download-task",
			Query:       "*",
			StartTime:   startTime - 60,
			EndTime:     time.Now().Unix(),
			DataFormat:  "csv",
			Sort:        "asc",
			Limit:       100,
			Compression: "gzip",
			TaskType:    0,
		}: {
			TopicID: suite.topic,
		},
	}

	for createDownloadTaskReq, describeDownloadTasksReq := range testcases {
		_, err := suite.cli.CreateDownloadTask(createDownloadTaskReq)
		suite.NoError(err)

		describeDownloadTasksResp, err := suite.cli.DescribeDownloadTasks(describeDownloadTasksReq)
		suite.NoError(err)
		suite.Equal(1, int(describeDownloadTasksResp.Total))
		suite.Equal(createDownloadTaskReq.TaskName, describeDownloadTasksResp.Tasks[0].TaskName)
		suite.Equal(createDownloadTaskReq.Query, describeDownloadTasksResp.Tasks[0].Query)
		suite.Equal(createDownloadTaskReq.DataFormat, describeDownloadTasksResp.Tasks[0].DataFormat)
		suite.Equal(createDownloadTaskReq.Compression, describeDownloadTasksResp.Tasks[0].Compression)
		suite.Equal(createDownloadTaskReq.TaskType, describeDownloadTasksResp.Tasks[0].TaskType)
	}
}

func (suite *SDKDownloadTaskTestSuite) TestDescribeDownloadTasksAbnormally() {
	testcases := map[*DescribeDownloadTasksRequest]*Error{
		{
			TopicID: suite.topic,
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "TopicNotExist",
			Message:  "Topic does not exist.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeDownloadTasks(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKDownloadTaskTestSuite) TestDescribeDownloadUrlNormally() {
	startTime := time.Now().Unix()

	testcases := map[*CreateDownloadTaskRequest]*Error{
		{
			TopicID:     suite.topic,
			TaskName:    "go-sdk-download-task",
			Query:       "*",
			StartTime:   startTime - 60,
			EndTime:     time.Now().Unix(),
			DataFormat:  "csv",
			Sort:        "asc",
			Limit:       100,
			Compression: "gzip",
			TaskType:    0,
		}: nil,
	}

	for req, expectedErr := range testcases {
		createDownloadTaskResp, err := suite.cli.CreateDownloadTask(req)
		suite.NoError(err)
		time.Sleep(15 * time.Second)

		resp, err := suite.cli.DescribeDownloadUrl(&DescribeDownloadUrlRequest{TaskId: createDownloadTaskResp.TaskId})
		suite.validateError(err, expectedErr)
		suite.GreaterOrEqual(len(resp.DownloadUrl), 0)
	}
}

func (suite *SDKDownloadTaskTestSuite) TestDescribeDownloadUrlAbnormally() {
	nonexistentTaskID := uuid.New().String()

	testcases := map[*DescribeDownloadUrlRequest]*Error{
		{
			TaskId: nonexistentTaskID,
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "TaskNotExist",
			Message:  "Task " + nonexistentTaskID + " does not exist",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeDownloadUrl(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKDownloadTaskTestSuite) TestCancelDownloadTaskNormally() {
	startTime := time.Now().Unix()

	createDownloadTaskReq := &CreateDownloadTaskRequest{
		TopicID:     suite.topic,
		TaskName:    "go-sdk-download-task-cancel",
		Query:       "*",
		StartTime:   startTime - 60,
		EndTime:     time.Now().Unix(),
		DataFormat:  "csv",
		Sort:        "asc",
		Limit:       100,
		Compression: "gzip",
		TaskType:    0,
	}

	createDownloadTaskResp, err := suite.cli.CreateDownloadTask(createDownloadTaskReq)
	suite.NoError(err)

	_, err = suite.cli.CancelDownloadTask(&CancelDownloadTaskRequest{TaskId: createDownloadTaskResp.TaskId})
	suite.NoError(err)
}
