package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKAlarmTaskTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
}

func (suite *SDKAlarmTaskTestSuite) SetupTest() {
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

func (suite *SDKAlarmTaskTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKAlarmTaskTestSuite(t *testing.T) {
	suite.Run(t, new(SDKAlarmTaskTestSuite))
}

func (suite *SDKAlarmTaskTestSuite) TestAlarm() {
	var groupID string
	// CreateAlarmNotifyGroup
	{
		req := &CreateAlarmNotifyGroupRequest{
			GroupName: "golang-sdk-unit-test",
			NoticeType: []NoticeType{
				"Recovery",
				"Trigger",
			},
			Receivers: []Receiver{
				{
					ReceiverType:     "User",
					ReceiverNames:    []string{"test_xiaoping"},
					ReceiverChannels: []ReceiverChannel{"Webhook"},
					StartTime:        "00:00:00",
					EndTime:          "23:59:59",
				},
			},
		}
		req.Receivers[0].Webhook = "TestWebhook"
		resp, err := suite.cli.CreateAlarmNotifyGroup(req)
		suite.NoError(err)
		groupID = resp.NotifyGroupID
	}

	// ModifyAlarmNotifyGroup
	{
		req := &ModifyAlarmNotifyGroupRequest{
			AlarmNotifyGroupID: groupID,
			NoticeType: &NoticeTypes{
				"Recovery",
				"Trigger",
			},
		}
		_, err := suite.cli.ModifyAlarmNotifyGroup(req)
		suite.NoError(err)
	}

	// DescribeAlarmNotifyGroup
	{
		req := &DescribeAlarmNotifyGroupsRequest{}
		_, err := suite.cli.DescribeAlarmNotifyGroups(req)
		suite.NoError(err)
	}

	// DescribeAlarms
	{
		req := &DescribeAlarmsRequest{
			ProjectID: suite.project,
		}
		_, err := suite.cli.DescribeAlarms(req)
		suite.NoError(err)
	}

	// DeleteAlarmNotifyGroup
	{
		req := &DeleteAlarmNotifyGroupRequest{
			AlarmNotifyGroupID: groupID,
		}
		_, err := suite.cli.DeleteAlarmNotifyGroup(req)
		suite.NoError(err)
	}

	// CreateAlarmNotifyGroupInvalid
	{
		req := &CreateAlarmNotifyGroupRequest{
			GroupName: groupID,
			NoticeType: []NoticeType{
				"Recovery",
				"Trigger",
			},
			Receivers: []Receiver{
				{
					ReceiverType:     "User",
					ReceiverNames:    []string{"notexistuser"},
					ReceiverChannels: []ReceiverChannel{"Webhook"},
					StartTime:        "00:00:00",
					EndTime:          "23:59:59",
				},
			},
		}
		req.Receivers[0].Webhook = "TestWebhook"
		_, err := suite.cli.CreateAlarmNotifyGroup(req)
		suite.Error(err)
	}

	// ModifyAlaramNotifyGroupInvalid
	{
		req := &ModifyAlarmNotifyGroupRequest{
			AlarmNotifyGroupName: new(string),
			NoticeType: &NoticeTypes{
				"Invalid",
			},
		}
		*req.AlarmNotifyGroupName = "Test"
		_, err := suite.cli.ModifyAlarmNotifyGroup(req)
		suite.Error(err)
	}

	// DescribeAlarmNotifyGroupsInvalid
	{
		req := &DescribeAlarmNotifyGroupsRequest{
			PageSize: -1,
		}
		_, err := suite.cli.DescribeAlarmNotifyGroups(req)
		suite.Error(err)
	}

	// DescribeAlarmsInvalid
	{
		req := &DescribeAlarmsRequest{
			ProjectID: "test",
		}
		_, err := suite.cli.DescribeAlarms(req)
		suite.Error(err)
	}
}
