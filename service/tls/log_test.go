package tls

import (
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

func (suite *SDKLogTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKLogTestSuite(t *testing.T) {
	suite.Run(t, new(SDKLogTestSuite))
}

// TestPutLogs: test put logs
func (suite *SDKLogTestSuite) TestPutLogs() {
	var logs []*pb.LogGroupList
	for i := 0; i < 10; i++ {
		idx := strconv.Itoa(i)

		logs = append(logs, &pb.LogGroupList{
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

	// update 10 logGroupLists without compression
	for _, logGroupList := range logs {
		_, err := suite.cli.PutLogs(&PutLogsRequest{
			TopicID:      suite.topic,
			CompressType: "none",
			LogBody:      logGroupList,
		})
		suite.NoError(err)
	}

	// update 10 logGroupLists with lz4 compression
	for _, logGroupList := range logs {
		_, err := suite.cli.PutLogs(&PutLogsRequest{
			TopicID:      suite.topic,
			CompressType: "lz4",
			LogBody:      logGroupList,
		})
		suite.NoError(err)
	}

	// wait for consumption
	time.Sleep(30 * time.Second)

	// test pull logs
	beginCursorResp, err := suite.cli.DescribeCursor(&DescribeCursorRequest{
		TopicID: suite.topic,
		ShardID: 0,
		From:    "begin",
	})
	suite.NoError(err)

	beginCursor := beginCursorResp.Cursor

	pullLogsResp, err := suite.cli.ConsumeLogs(&ConsumeLogsRequest{
		TopicID:       suite.topic,
		ShardID:       0,
		Cursor:        beginCursor,
		EndCursor:     nil,
		LogGroupCount: IntPtr(100),
		Compression:   nil,
	})
	suite.NoError(err)

	suite.Equal(pullLogsResp.Count, 10)

	// wait for consumption
	time.Sleep(30 * time.Second)

	// test search logs
	searchRes, err := suite.cli.SearchLogs(&SearchLogsRequest{
		TopicID:   suite.topic,
		Query:     "*",
		StartTime: 1600000000000,
		EndTime:   2600000000000,
		Limit:     100,
	})
	suite.NoError(err)

	suite.Equal(20, searchRes.HitCount)

	logMap := make(map[string]struct{})
	for _, searchLog := range searchRes.Logs {
		for _, v := range searchLog {
			switch v.(type) {
			case string:
				logMap[v.(string)] = struct{}{}
			}
		}
	}

	suite.Contains(logMap, "localhost")

	for _, logGroupList := range logs {
		for _, logGroup := range logGroupList.LogGroups {
			suite.Contains(logMap, logGroup.FileName)

			for _, log := range logGroup.Logs {
				for _, content := range log.Contents {
					suite.Contains(logMap, content.Value)
				}
			}
		}
	}
}
