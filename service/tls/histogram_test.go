package tls

import (
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKHistogramTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKHistogramTestSuite) SetupTest() {
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

func (suite *SDKHistogramTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKHistogramTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKHistogramTestSuite(t *testing.T) {
	suite.Run(t, new(SDKHistogramTestSuite))
}

func (suite *SDKHistogramTestSuite) TestDescribeHistogramNormally() {
	startTime := time.Now().Unix()

	time.Sleep(15 * time.Second)
	err := putLogs(suite.cli, suite.topic, "192.168.1.1", "sys.log", 100)
	suite.NoError(err)

	testcases := map[*DescribeHistogramRequest]*DescribeHistogramResponse{
		{
			TopicID:   suite.topic,
			Query:     "*",
			StartTime: startTime,
			EndTime:   time.Now().Unix(),
		}: {
			ResultStatus: "complete",
			TotalCount:   100,
		},
	}

	for req, expectedResp := range testcases {
		resp, err := suite.cli.DescribeHistogram(req)
		suite.NoError(err)
		suite.Equal(expectedResp.ResultStatus, resp.ResultStatus)
		suite.Equal(expectedResp.TotalCount, resp.TotalCount)
	}
}

func (suite *SDKHistogramTestSuite) TestDescribeHistogramAbnormally() {
	startTime := time.Now().Unix() - 60

	testcases := map[*DescribeHistogramRequest]*Error{
		{
			TopicID:   suite.topic,
			Query:     "*",
			StartTime: startTime,
			EndTime:   startTime,
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key StartTime, value " + strconv.FormatInt(startTime, 10) + ", please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeHistogram(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKHistogramTestSuite) TestDescribeHistogramV1Normally() {
	startTime := time.Now().Unix()

	err := putLogs(suite.cli, suite.topic, "192.168.1.1", "sys.log", 100)
	suite.NoError(err)

	testcases := map[*DescribeHistogramV1Request]*DescribeHistogramV1Response{
		{
			TopicID:   suite.topic,
			Query:     "*",
			StartTime: startTime,
			EndTime:   startTime + 15,
			Interval:  5000,
		}: {
			ResultStatus: "complete",
			TotalCount:   100,
		},
	}

	for req, expectedResp := range testcases {
		resp, err := suite.cli.DescribeHistogramV1(req)
		suite.NoError(err)
		suite.Equal(expectedResp.ResultStatus, resp.ResultStatus)
		suite.Equal(expectedResp.TotalCount, resp.TotalCount)
	}
}

func (suite *SDKHistogramTestSuite) TestDescribeHistogramV1Normally2() {
	startTime := time.Now().Unix()

	err := putLogs(suite.cli, suite.topic, "192.168.1.1", "sys.log", 100)
	suite.NoError(err)

	testcases := map[*DescribeHistogramV1Request]*DescribeHistogramV1Response{
		{
			TopicID:   suite.topic,
			Query:     "*",
			StartTime: startTime,
			EndTime:   startTime + 15,
		}: {
			ResultStatus: "complete",
			TotalCount:   100,
		},
	}

	for req, expectedResp := range testcases {
		resp, err := suite.cli.DescribeHistogramV1(req)
		suite.NoError(err)
		suite.Equal(expectedResp.ResultStatus, resp.ResultStatus)
		suite.Equal(expectedResp.TotalCount, resp.TotalCount)
		suite.Equal(15, len(resp.HistogramInfos))
	}
}

func (suite *SDKHistogramTestSuite) TestDescribeHistogramV1Abnormally() {
	startTime := time.Now().Unix() - 60

	testcases := map[*DescribeHistogramV1Request]*Error{
		{
			TopicID:   suite.topic,
			Query:     "*",
			StartTime: startTime,
			EndTime:   startTime,
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key StartTime, value " + strconv.FormatInt(startTime, 10) + ", please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeHistogramV1(req)
		suite.validateError(err, expectedErr)
	}
}
