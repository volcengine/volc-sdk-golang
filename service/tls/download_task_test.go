package tls

import (
	"os"
	"strings"
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

	projectId, err := CreateProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-index-"+uuid.New().String(),
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

func (suite *SDKDownloadTaskTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKDownloadTaskTestSuite(t *testing.T) {
	suite.Run(t, new(SDKDownloadTaskTestSuite))
}

func (suite *SDKDownloadTaskTestSuite) TestDownloadTask() {
	var TaskId string

	var startTime = time.Now().UnixNano()/1e6 - 1000
	var endTime = time.Now().UnixNano() / 1e6

	// CreateDownloadTask
	{
		testcases := map[*CreateDownloadTaskRequest]*CreateDownloadTaskResponse{
			{
				TopicID:     suite.topic,
				TaskName:    "unit-test-task",
				Query:       "*",
				StartTime:   startTime,
				EndTime:     endTime,
				Compression: "gzip",
				DataFormat:  "json",
				Limit:       100,
				Sort:        "asc",
			}: {},
		}

		for req, _ := range testcases {
			actualResp, err := suite.cli.CreateDownloadTask(req)
			suite.NoError(err)
			suite.NotEqual(actualResp.TaskId, "")
			TaskId = actualResp.TaskId
		}
	}

	// DescribeDownloadTask
	{
		testcases := map[*DescribeDownloadTasksRequest]*DescribeDownloadTasksResponse{
			{
				TopicID: suite.topic,
			}: {
				Total: 1,
				Tasks: []*DownloadTaskResp{
					{
						TaskId:     TaskId,
						TaskName:   "unit-test-task",
						TopicId:    suite.topic,
						Query:      "*",
						StartTime:  time.Unix(startTime/1e3, 0).Format("2006-01-02 15:04:05"),
						EndTime:    time.Unix(endTime/1e3, 0).Format("2006-01-02 15:04:05"),
						TaskStatus: "creating",
					},
				},
			},
		}

		for req, resp := range testcases {
			actualResp, err := suite.cli.DescribeDownloadTasks(req)
			suite.NoError(err)
			suite.Equal(resp.Total, actualResp.Total)
			suite.Equal(len(resp.Tasks), len(actualResp.Tasks))
			for i, _ := range resp.Tasks {
				suite.Equal(resp.Tasks[i].TaskId, actualResp.Tasks[i].TaskId)
				suite.Equal(resp.Tasks[i].TaskName, actualResp.Tasks[i].TaskName)
				suite.Equal(resp.Tasks[i].TopicId, actualResp.Tasks[i].TopicId)
				suite.Equal(resp.Tasks[i].Query, actualResp.Tasks[i].Query)
				suite.Equal(resp.Tasks[i].StartTime, actualResp.Tasks[i].StartTime)
				suite.Equal(resp.Tasks[i].EndTime, actualResp.Tasks[i].EndTime)
				suite.Equal(resp.Tasks[i].TaskStatus, actualResp.Tasks[i].TaskStatus)
			}
		}
	}

	// DescribeDownloadUrl
	{
		testcases := map[*DescribeDownloadUrlRequest]*DescribeDownloadUrlResponse{
			{
				TaskId: TaskId,
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.DescribeDownloadUrl(req)
			var ok bool
			{
				if strings.Contains(err.Error(), "DownloadTaskNotFinished") || err != nil {
					ok = true
				}
			}
			suite.Equal(ok, true)
		}
	}

	// CreateDownloadTaskInvalid
	{
		testcases := map[*CreateDownloadTaskRequest]*CreateDownloadTaskResponse{
			{
				TopicID:     suite.topic,
				TaskName:    "unit-test-task",
				Query:       "*",
				StartTime:   startTime,
				EndTime:     endTime,
				Compression: "lz4",
				DataFormat:  "json",
				Limit:       100,
				Sort:        "asc",
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.CreateDownloadTask(req)
			suite.Error(err)
		}
	}

	// DescribeDownloadTaskInvalid
	{
		testcases := map[*DescribeDownloadTasksRequest]*DescribeDownloadTasksResponse{
			{
				TopicID: "Test",
			}: {
				Total: 1,
				Tasks: []*DownloadTaskResp{
					{
						TaskId:    TaskId,
						TaskName:  "unit-test-task",
						TopicId:   suite.topic,
						Query:     "*",
						StartTime: time.Unix(startTime/1e3, 0).Format("2006-01-02 15:04:05"),
						EndTime:   time.Unix(endTime/1e3, 0).Format("2006-01-02 15:04:05"),
					},
				},
			},
		}

		for req, _ := range testcases {
			_, err := suite.cli.DescribeDownloadTasks(req)
			suite.Error(err)
		}
	}

	// DescribeDownloadUrlInvalid
	{
		testcases := map[*DescribeDownloadUrlRequest]*DescribeDownloadUrlResponse{
			{
				TaskId: "Test",
			}: {},
		}

		for req, _ := range testcases {
			_, err := suite.cli.DescribeDownloadUrl(req)
			suite.Error(err)
		}
	}
}
