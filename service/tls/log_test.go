package tls

import (
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type SDKLogTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKLogTestSuite) SetupTest() {
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

func (suite *SDKLogTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKLogTestSuite) validateError(err error, expectErr *Error) {
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

func getLogGroupList() {
	var logGroupList []*pb.LogGroupList
	for i := 0; i < 10; i++ {
		idx := strconv.Itoa(i)
		logGroupList = append(logGroupList, &pb.LogGroupList{
			LogGroups: []*pb.LogGroup{
				{
					Source:   "localhost",
					FileName: "log" + idx,
					Logs: []*pb.Log{
						{
							Contents: []*pb.LogContent{
								{
									Key:   "key-1",
									Value: "test-message" + idx,
								},
								{
									Key:   "key-2",
									Value: idx,
								},
							},
							Time: time.Now().Unix(),
						},
					},
				},
			},
		})
	}
}

func getLogs(num int) []Log {
	logs := make([]Log, 0)

	for i := 0; i < num; i++ {
		log := Log{
			Contents: []LogContent{
				{
					Key:   "key-1",
					Value: uuid.New().String(),
				},
				{
					Key:   "key-2",
					Value: uuid.New().String(),
				},
			},
			Time: time.Now().Unix(),
		}
		logs = append(logs, log)
	}

	return logs
}

func putLogs(cli Client, topicID string, source string, filename string, num int) error {
	_, err := cli.PutLogsV2(&PutLogsV2Request{
		TopicID:      topicID,
		HashKey:      "",
		CompressType: "lz4",
		Source:       source,
		FileName:     filename,
		Logs:         getLogs(num),
	})
	if err != nil {
		return err
	}
	time.Sleep(60 * time.Second)
	return err
}

func TestSDKLogTestSuite(t *testing.T) {
	suite.Run(t, new(SDKLogTestSuite))
}

func (suite *SDKLogTestSuite) TestPutLogsNormally() {
	req := &PutLogsRequest{
		TopicID:      suite.topic,
		CompressType: "lz4",
		LogBody:      &pb.LogGroupList{},
	}
	var logGroups []*pb.LogGroup
	for i := 0; i < 3; i++ {
		idx := strconv.Itoa(i)
		if i == 1 {
			logGroups = append(logGroups, &pb.LogGroup{
				Source:   "localhost",
				FileName: "log" + idx,
				Logs:     []*pb.Log{},
			})
		} else {
			logGroups = append(logGroups, &pb.LogGroup{
				Source:   "localhost",
				FileName: "log" + idx,
				Logs: []*pb.Log{
					{
						Contents: []*pb.LogContent{
							{
								Key:   "key-1",
								Value: "test-message" + idx,
							},
							{
								Key:   "key-2",
								Value: idx,
							},
						},
					},
				},
			})
		}
	}
	req.LogBody.LogGroups = logGroups

	_, err := suite.cli.PutLogs(req)
	suite.NoError(err)
}

func (suite *SDKLogTestSuite) TestPutLogsV2Normally() {
	_, err := suite.cli.PutLogsV2(&PutLogsV2Request{
		TopicID:      suite.topic,
		HashKey:      "",
		CompressType: "lz4",
		Source:       "192.168.1.1",
		FileName:     "sys.log",
		Logs:         getLogs(100),
	})
	suite.NoError(err)
}

func (suite *SDKLogTestSuite) TestPutLogsV2Abnormally() {
	testcases := map[*PutLogsV2Request]*Error{
		{
			TopicID:      suite.topic,
			HashKey:      "",
			CompressType: "rar",
			Source:       "192.168.1.1",
			FileName:     "sys.log",
			Logs:         getLogs(100),
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key x-tls-compresstype, value rar, please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.PutLogsV2(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKLogTestSuite) TestDescribeCursorNormally() {
	testcases := map[*DescribeCursorRequest]*Error{
		{
			TopicID: suite.topic,
			ShardID: 0,
			From:    "begin",
		}: nil,
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeCursor(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKLogTestSuite) TestDescribeCursorAbnormally() {
	testcases := map[*DescribeCursorRequest]*Error{
		{
			TopicID: suite.topic,
			ShardID: 100,
			From:    "begin",
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key ShardId, value %!s(int32=100), please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeCursor(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKLogTestSuite) TestConsumeLogsNormally() {
	err := putLogs(suite.cli, suite.topic, "192.168.1.1", "sys.log", 100)
	suite.NoError(err)

	describeCursorResp, err := suite.cli.DescribeCursor(&DescribeCursorRequest{
		TopicID: suite.topic,
		ShardID: 0,
		From:    "begin",
	})
	suite.NoError(err)
	cursor := describeCursorResp.Cursor

	resp, err := suite.cli.ConsumeLogs(&ConsumeLogsRequest{
		TopicID:     suite.topic,
		ShardID:     0,
		Cursor:      cursor,
		Compression: StrPtr("lz4"),
	})
	suite.NoError(err)
	suite.Equal(1, resp.Count)
	suite.Equal(100, len(resp.Logs.LogGroups[0].Logs))
	suite.Equal("192.168.1.1", resp.Logs.LogGroups[0].Source)
	suite.Equal("sys.log", resp.Logs.LogGroups[0].FileName)
}

func (suite *SDKLogTestSuite) TestConsumeLogsAbnormally() {
	invalidCursor := uuid.New().String()
	testcases := map[*ConsumeLogsRequest]*Error{
		{
			TopicID:     suite.topic,
			ShardID:     0,
			Cursor:      invalidCursor,
			Compression: StrPtr("lz4"),
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key Cursor, value " + invalidCursor + ", please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.ConsumeLogs(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKLogTestSuite) TestSearchLogsV2Normally() {
	startTime := time.Now().Unix()

	time.Sleep(60 * time.Second)
	err := putLogs(suite.cli, suite.topic, "192.168.1.1", "sys.log", 100)
	suite.NoError(err)

	testcases := map[*SearchLogsRequest]*SearchLogsResponse{
		{
			TopicID:   suite.topic,
			Query:     "",
			StartTime: startTime,
			EndTime:   time.Now().Unix(),
			Limit:     100,
		}: {
			Status:   "complete",
			Analysis: false,
			Count:    100,
		},
	}

	for req, expectedResp := range testcases {
		resp, err := suite.cli.SearchLogsV2(req)
		suite.NoError(err)
		suite.Equal(expectedResp.Status, resp.Status)
		suite.Equal(expectedResp.Analysis, resp.Analysis)
		suite.Equal(expectedResp.Count, resp.Count)
	}
}

func (suite *SDKLogTestSuite) TestSearchLogsV2Abnormally() {
	startTime := time.Now().Unix()

	_, err := suite.cli.DeleteIndex(&DeleteIndexRequest{TopicID: suite.topic})
	suite.NoError(err)
	_, err = suite.cli.SearchLogsV2(&SearchLogsRequest{
		TopicID:   suite.topic,
		Query:     "*",
		StartTime: startTime,
		EndTime:   time.Now().Unix(),
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "IndexNotExists",
		Message:  "Index does not exist.",
	}
	suite.validateError(err, expectedErr)
}
