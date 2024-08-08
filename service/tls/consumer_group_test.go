package tls

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKConsumerGroupTestSuite struct {
	suite.Suite

	cli               Client
	project           string
	topic             string
	consumerGroupName string
}

func (suite *SDKConsumerGroupTestSuite) SetupTest() {
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

func (suite *SDKConsumerGroupTestSuite) TearDownTest() {
	if len(suite.consumerGroupName) > 0 {
		_, deleteConsumerGroupErr := suite.cli.DeleteConsumerGroup(&DeleteConsumerGroupRequest{
			ProjectID:         suite.project,
			ConsumerGroupName: suite.consumerGroupName,
		})
		suite.NoError(deleteConsumerGroupErr)
	}

	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKConsumerGroupTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKConsumerGroupTestSuite(t *testing.T) {
	suite.Run(t, new(SDKConsumerGroupTestSuite))
}

func createConsumerGroup(cli Client, projectID string, topicID string, consumerGroupName string) error {
	_, err := cli.CreateConsumerGroup(&CreateConsumerGroupRequest{
		ProjectID:         projectID,
		TopicIDList:       []string{topicID},
		ConsumerGroupName: consumerGroupName,
		HeartbeatTTL:      10,
		OrderedConsume:    true,
	})

	return err
}

func (suite *SDKConsumerGroupTestSuite) TestCreateConsumerGroupNormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	_, err := suite.cli.CreateConsumerGroup(&CreateConsumerGroupRequest{
		ProjectID:         suite.project,
		TopicIDList:       []string{suite.topic},
		ConsumerGroupName: consumerGroupName,
		HeartbeatTTL:      10,
		OrderedConsume:    true,
	})
	suite.NoError(err)

	suite.consumerGroupName = consumerGroupName
}

func (suite *SDKConsumerGroupTestSuite) TestCreateConsumerGroupAbnormally() {
	projectID := uuid.New().String()
	_, err := suite.cli.CreateConsumerGroup(&CreateConsumerGroupRequest{
		ProjectID:         projectID,
		TopicIDList:       []string{suite.topic},
		ConsumerGroupName: suite.consumerGroupName,
		HeartbeatTTL:      10,
		OrderedConsume:    true,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}

func (suite *SDKConsumerGroupTestSuite) TestDeleteConsumerGroupAbnormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	_, err := suite.cli.DeleteConsumerGroup(&DeleteConsumerGroupRequest{
		ProjectID:         suite.project,
		ConsumerGroupName: consumerGroupName,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ConsumerGroupName, value golang-sdk-consumer-group, please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}

func (suite *SDKConsumerGroupTestSuite) TestModifyConsumerGroupNormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	err := createConsumerGroup(suite.cli, suite.project, suite.topic, consumerGroupName)
	suite.NoError(err)

	_, err = suite.cli.ModifyConsumerGroup(&ModifyConsumerGroupRequest{
		ProjectID:         suite.project,
		ConsumerGroupName: consumerGroupName,
		OrderedConsume:    BoolPtr(false),
	})
	suite.NoError(err)

	resp, err := suite.cli.DescribeConsumerGroups(&DescribeConsumerGroupsRequest{
		ProjectID: suite.project,
	})
	suite.NoError(err)
	suite.Equal(1, len(resp.ConsumerGroups))
	consumerGroupInfo := resp.ConsumerGroups[0]
	suite.Equal(suite.project, consumerGroupInfo.ProjectID)
	suite.Equal(consumerGroupName, consumerGroupInfo.ConsumerGroupName)
	suite.Equal(false, consumerGroupInfo.OrderedConsume)

	suite.consumerGroupName = consumerGroupName
}

func (suite *SDKConsumerGroupTestSuite) TestModifyConsumerGroupAbnormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	_, err := suite.cli.ModifyConsumerGroup(&ModifyConsumerGroupRequest{
		ProjectID:         suite.project,
		ConsumerGroupName: consumerGroupName,
		OrderedConsume:    BoolPtr(false),
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  fmt.Sprintf("Invalid argument key ConsumerGroupName, value %s, please check argument.", consumerGroupName),
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}

func (suite *SDKConsumerGroupTestSuite) TestDescribeConsumerGroupsNormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	err := createConsumerGroup(suite.cli, suite.project, suite.topic, consumerGroupName)
	suite.NoError(err)

	resp, err := suite.cli.DescribeConsumerGroups(&DescribeConsumerGroupsRequest{
		ProjectID: suite.project,
	})
	suite.NoError(err)
	suite.Equal(1, len(resp.ConsumerGroups))
	consumerGroupInfo := resp.ConsumerGroups[0]
	suite.Equal(suite.project, consumerGroupInfo.ProjectID)
	suite.Equal(consumerGroupName, consumerGroupInfo.ConsumerGroupName)
	suite.Equal(10, consumerGroupInfo.HeartbeatTTL)
	suite.Equal(true, consumerGroupInfo.OrderedConsume)

	suite.consumerGroupName = consumerGroupName
}

func (suite *SDKConsumerGroupTestSuite) TestDescribeConsumerGroupsAbnormally() {
	projectID := uuid.New().String()
	_, err := suite.cli.DescribeConsumerGroups(&DescribeConsumerGroupsRequest{
		ProjectID: projectID,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}

func (suite *SDKConsumerGroupTestSuite) TestConsumerHeartbeatNormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	err := createConsumerGroup(suite.cli, suite.project, suite.topic, consumerGroupName)
	suite.NoError(err)

	consumerName := "golang-sdk-consumer"

	for i := 0; i < 5; i++ {
		resp, err := suite.cli.ConsumerHeartbeat(&ConsumerHeartbeatRequest{
			ProjectID:         suite.project,
			ConsumerGroupName: consumerGroupName,
			ConsumerName:      consumerName,
		})
		suite.NoError(err)

		if len(resp.Shards) > 0 {
			suite.Equal(1, len(resp.Shards))
			suite.Equal(suite.topic, resp.Shards[0].TopicID)
			suite.Equal(0, resp.Shards[0].ShardID)
			break
		} else {
			time.Sleep(5 * time.Second)
		}
	}

	suite.consumerGroupName = consumerGroupName
}

func (suite *SDKConsumerGroupTestSuite) TestConsumerHeartbeatAbnormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	_, err := suite.cli.ConsumerHeartbeat(&ConsumerHeartbeatRequest{
		ProjectID:         suite.project,
		ConsumerGroupName: consumerGroupName,
		ConsumerName:      "consumer",
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  fmt.Sprintf("Invalid argument key ConsumerGroupName, value %s, please check argument.", consumerGroupName),
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}

func (suite *SDKConsumerGroupTestSuite) TestDescribeCheckPointNormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	err := createConsumerGroup(suite.cli, suite.project, suite.topic, consumerGroupName)
	suite.NoError(err)

	resp, err := suite.cli.DescribeCheckPoint(&DescribeCheckPointRequest{
		ConsumerGroupName: consumerGroupName,
		ProjectID:         suite.project,
		TopicID:           suite.topic,
		ShardID:           0,
	})
	suite.NoError(err)
	suite.Equal(0, int(resp.ShardID))
	suite.Equal("", resp.Checkpoint)

	suite.consumerGroupName = consumerGroupName
}

func (suite *SDKConsumerGroupTestSuite) TestDescribeCheckPointAbnormally() {
	_, err := suite.cli.DescribeCheckPoint(&DescribeCheckPointRequest{
		ConsumerGroupName: "",
		ProjectID:         suite.project,
		TopicID:           suite.topic,
		ShardID:           0,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ConsumerGroupName, value , please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}

func (suite *SDKConsumerGroupTestSuite) TestModifyCheckPointNormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	err := createConsumerGroup(suite.cli, suite.project, suite.topic, consumerGroupName)
	suite.NoError(err)

	describeCursorResp, err := suite.cli.DescribeCursor(&DescribeCursorRequest{
		TopicID: suite.topic,
		ShardID: 0,
		From:    "begin",
	})
	suite.NoError(err)
	cursor := describeCursorResp.Cursor

	_, err = suite.cli.ModifyCheckPoint(&ModifyCheckPointRequest{
		ConsumerGroupName: consumerGroupName,
		ProjectID:         suite.project,
		TopicID:           suite.topic,
		ShardID:           0,
		Checkpoint:        cursor,
	})
	suite.NoError(err)

	describeCheckpointResp, err := suite.cli.DescribeCheckPoint(&DescribeCheckPointRequest{
		ConsumerGroupName: consumerGroupName,
		ProjectID:         suite.project,
		TopicID:           suite.topic,
		ShardID:           0,
	})
	suite.NoError(err)
	suite.Equal(describeCheckpointResp.Checkpoint, cursor)

	suite.consumerGroupName = consumerGroupName
}

func (suite *SDKConsumerGroupTestSuite) TestModifyCheckPointAbnormally() {
	consumerGroupName := "golang-sdk-consumer-group"

	describeCursorResp, err := suite.cli.DescribeCursor(&DescribeCursorRequest{
		TopicID: suite.topic,
		ShardID: 0,
		From:    "begin",
	})
	suite.NoError(err)
	cursor := describeCursorResp.Cursor

	_, err = suite.cli.ModifyCheckPoint(&ModifyCheckPointRequest{
		ConsumerGroupName: consumerGroupName,
		ProjectID:         suite.project,
		TopicID:           suite.topic,
		ShardID:           0,
		Checkpoint:        cursor,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  fmt.Sprintf("Invalid argument key ConsumerGroupName, value %s, please check argument.", consumerGroupName),
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}

func (suite *SDKConsumerGroupTestSuite) TestResetCheckPointNormally() {
	consumerGroupName := "golang-sdk-consumer-group"
	err := createConsumerGroup(suite.cli, suite.project, suite.topic, consumerGroupName)
	suite.NoError(err)

	_, err = suite.cli.ResetCheckPoint(&ResetCheckPointRequest{
		ProjectID:         suite.project,
		ConsumerGroupName: consumerGroupName,
		Position:          "begin",
	})
	suite.NoError(err)

	suite.consumerGroupName = consumerGroupName
}

func (suite *SDKConsumerGroupTestSuite) TestResetCheckPointAbnormally() {
	consumerGroupName := "golang-sdk-consumer-group"

	_, err := suite.cli.ResetCheckPoint(&ResetCheckPointRequest{
		ProjectID:         suite.project,
		ConsumerGroupName: consumerGroupName,
		Position:          "begin",
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  fmt.Sprintf("Invalid argument key ConsumerGroupName, value %s, please check argument.", consumerGroupName),
	}
	suite.validateError(err, expectedErr)

	suite.consumerGroupName = ""
}
