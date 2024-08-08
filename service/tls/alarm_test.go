package tls

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKAlarmTaskTestSuite struct {
	suite.Suite

	cli                  Client
	project              string
	topic                string
	alarmNotifyGroupList []string
	alarmList            []string
}

func (suite *SDKAlarmTaskTestSuite) SetupTest() {
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

func (suite *SDKAlarmTaskTestSuite) TearDownTest() {
	for _, alarmID := range suite.alarmList {
		_, deleteAlarmErr := suite.cli.DeleteAlarm(&DeleteAlarmRequest{AlarmID: alarmID})
		suite.NoError(deleteAlarmErr)
	}
	suite.alarmList = nil
	for _, alarmNotifyGroupID := range suite.alarmNotifyGroupList {
		_, deleteAlarmNotifyGroupErr := suite.cli.DeleteAlarmNotifyGroup(&DeleteAlarmNotifyGroupRequest{AlarmNotifyGroupID: alarmNotifyGroupID})
		suite.NoError(deleteAlarmNotifyGroupErr)
	}
	suite.alarmNotifyGroupList = nil

	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKAlarmTaskTestSuite) validateError(err error, expectErr *Error) {
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

func createAlarmNotifyGroup(cli Client, groupName string) (string, error) {
	createAlarmNotifyGroupResp, err := cli.CreateAlarmNotifyGroup(&CreateAlarmNotifyGroupRequest{
		GroupName: groupName,
		NoticeType: []NoticeType{
			"Recovery",
			"Trigger",
		},
		Receivers: []Receiver{
			{
				ReceiverType:     "User",
				ReceiverNames:    []string{"zhengyu1"},
				ReceiverChannels: []ReceiverChannel{"Sms", "Phone", "Email"},
				StartTime:        "00:00:00",
				EndTime:          "23:59:59",
			},
		},
	})

	return createAlarmNotifyGroupResp.NotifyGroupID, err
}

func createAlarms(cli Client, projectID string, topicID string, alarmNotifyGroupID string) ([]string, error) {
	status := true
	createAlarmReqs := []*CreateAlarmRequest{
		{
			AlarmName: "golang-sdk-test-alarm-1",
			ProjectID: projectID,
			Status:    &status,
			QueryRequest: []QueryRequest{{
				Query:           "Failed | select count(*) as errNum",
				Number:          1,
				TopicID:         topicID,
				StartTimeOffset: -15,
				EndTimeOffset:   0,
			}},
			RequestCycle: RequestCycle{
				Type: "Period",
				Time: 10,
			},
			Condition:        "$1.errNum>0",
			TriggerPeriod:    2,
			AlarmPeriod:      60,
			AlarmNotifyGroup: []string{alarmNotifyGroupID},
			UserDefineMsg:    StrPtr("this is an alarm"),
		},
		{
			AlarmName: "golang-sdk-test-alarm-2",
			ProjectID: projectID,
			Status:    &status,
			QueryRequest: []QueryRequest{{
				Query:           "Failed | select count(*) as errNum",
				Number:          1,
				TopicID:         topicID,
				StartTimeOffset: -15,
				EndTimeOffset:   0,
			}},
			RequestCycle: RequestCycle{
				Type: "Period",
				Time: 10,
			},
			Condition:        "$1.errNum>0",
			TriggerPeriod:    2,
			AlarmPeriod:      60,
			AlarmNotifyGroup: []string{alarmNotifyGroupID},
			Severity:         StrPtr("critical"),
			AlarmPeriodDetail: &AlarmPeriodSetting{
				Sms:            10,
				Phone:          10,
				Email:          10,
				GeneralWebhook: 10,
			},
			UserDefineMsg: StrPtr("this is an alarm"),
		},
	}

	alarmList := make([]string, len(createAlarmReqs))
	for i, req := range createAlarmReqs {
		resp, err := cli.CreateAlarm(req)
		if err != nil {
			return alarmList, err
		}
		alarmList[i] = resp.AlarmID
	}

	return alarmList, nil
}

func TestSDKAlarmTaskTestSuite(t *testing.T) {
	suite.Run(t, new(SDKAlarmTaskTestSuite))
}

func (suite *SDKAlarmTaskTestSuite) TestCreateAlarmNotifyGroupNormally() {
	testcases := map[*CreateAlarmNotifyGroupRequest]*Error{
		{
			GroupName: uuid.New().String(),
			NoticeType: []NoticeType{
				"Recovery",
				"Trigger",
			},
			Receivers: []Receiver{
				{
					ReceiverType:     "User",
					ReceiverNames:    []string{"zhengyu1"},
					ReceiverChannels: []ReceiverChannel{"Sms", "Phone", "Email"},
					StartTime:        "00:00:00",
					EndTime:          "23:59:59",
				},
			},
		}: nil,
		{
			GroupName: uuid.New().String(),
			NoticeType: []NoticeType{
				"Recovery",
				"Trigger",
			},
			Receivers: []Receiver{
				{
					ReceiverType:     "User",
					ReceiverNames:    []string{"zhengyu1"},
					ReceiverChannels: []ReceiverChannel{"Phone"},
					StartTime:        "00:00:00",
					EndTime:          "23:59:59",
				},
			},
			IamProjectName: StrPtr("default"),
		}: nil,
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.CreateAlarmNotifyGroup(req)
		suite.validateError(rErr, err)
		if resp != nil {
			suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, resp.NotifyGroupID)
		}
	}
}

func (suite *SDKAlarmTaskTestSuite) TestCreateAlarmNotifyGroupAbnormally() {
	testcases := map[*CreateAlarmNotifyGroupRequest]*Error{
		{
			GroupName: uuid.New().String(),
			NoticeType: []NoticeType{
				"Recovery",
				"Trigger",
			},
			Receivers: []Receiver{
				{
					ReceiverType:     "User",
					ReceiverNames:    []string{"liangrubo"},
					ReceiverChannels: []ReceiverChannel{"Sms", "Phone", "Email"},
					StartTime:        "00:00:00",
					EndTime:          "23:59:59",
				},
			},
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "AlarmNoticeUserNotExist",
			Message:  "Alarm notice user liangrubo not exist",
		},
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.CreateAlarmNotifyGroup(req)
		suite.validateError(rErr, err)
		if resp != nil {
			suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, resp.NotifyGroupID)
		}
	}
}

func (suite *SDKAlarmTaskTestSuite) TestDeleteAlarmNotifyGroupAbnormally() {
	_, err := suite.cli.DeleteAlarmNotifyGroup(&DeleteAlarmNotifyGroupRequest{
		AlarmNotifyGroupID: uuid.New().String(),
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "AlarmNotifyGroupNotExist",
		Message:  "Alarm notify group does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKAlarmTaskTestSuite) TestModifyAlarmNotifyGroupNormally() {
	alarmNotifyGroupID, err := createAlarmNotifyGroup(suite.cli, uuid.New().String())
	suite.NoError(err)
	suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, alarmNotifyGroupID)

	_, err = suite.cli.ModifyAlarmNotifyGroup(&ModifyAlarmNotifyGroupRequest{
		AlarmNotifyGroupID: alarmNotifyGroupID,
		NoticeType:         (*NoticeTypes)(&[]NoticeType{"Trigger"}),
	})
	suite.NoError(err)
	resp, err := suite.cli.DescribeAlarmNotifyGroups(&DescribeAlarmNotifyGroupsRequest{
		NotifyGroupID: StrPtr(alarmNotifyGroupID),
	})
	alarmNotifyGroupInfo := resp.AlarmNotifyGroups[0]
	suite.Equal(1, len(alarmNotifyGroupInfo.NoticeType))
	suite.Equal("Trigger", string(alarmNotifyGroupInfo.NoticeType[0]))
}

func (suite *SDKAlarmTaskTestSuite) TestModifyAlarmNotifyGroupAbnormally() {
	alarmNotifyGroupID, err := createAlarmNotifyGroup(suite.cli, uuid.New().String())
	suite.NoError(err)
	suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, alarmNotifyGroupID)

	_, err = suite.cli.ModifyAlarmNotifyGroup(&ModifyAlarmNotifyGroupRequest{
		AlarmNotifyGroupID: alarmNotifyGroupID,
		NoticeType:         (*NoticeTypes)(&[]NoticeType{"Test"}),
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key NotifyType, value Test, please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKAlarmTaskTestSuite) TestDescribeAlarmNotifyGroupsNormally() {
	resp, err := suite.cli.DescribeAlarmNotifyGroups(&DescribeAlarmNotifyGroupsRequest{})
	suite.NoError(err)
	suite.GreaterOrEqual(int(resp.Total), 0)
}

func (suite *SDKAlarmTaskTestSuite) TestDescribeAlarmNotifyGroupsAbnormally() {
	_, err := suite.cli.DescribeAlarmNotifyGroups(&DescribeAlarmNotifyGroupsRequest{PageSize: -1})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key PageSize, value -1, please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKAlarmTaskTestSuite) TestCreateAlarmNormally() {
	alarmNotifyGroupID, err := createAlarmNotifyGroup(suite.cli, uuid.New().String())
	suite.NoError(err)
	suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, alarmNotifyGroupID)

	status := true
	testcases := map[*CreateAlarmRequest]*Error{
		{
			AlarmName: "golang-sdk-test-alarm-1",
			ProjectID: suite.project,
			Status:    &status,
			QueryRequest: []QueryRequest{{
				Query:           "Failed | select count(*) as errNum",
				Number:          1,
				TopicID:         suite.topic,
				StartTimeOffset: -15,
				EndTimeOffset:   0,
			}},
			RequestCycle: RequestCycle{
				Type: "Period",
				Time: 10,
			},
			Condition:        "$1.errNum>0",
			TriggerPeriod:    2,
			AlarmPeriod:      60,
			AlarmNotifyGroup: []string{alarmNotifyGroupID},
			UserDefineMsg:    StrPtr("this is an alarm"),
		}: nil,
		{
			AlarmName: "golang-sdk-test-alarm-2",
			ProjectID: suite.project,
			Status:    &status,
			QueryRequest: []QueryRequest{{
				Query:           "Failed | select count(*) as errNum",
				Number:          1,
				TopicID:         suite.topic,
				StartTimeOffset: -15,
				EndTimeOffset:   0,
			}},
			RequestCycle: RequestCycle{
				Type: "Period",
				Time: 10,
			},
			Condition:        "$1.errNum>0",
			TriggerPeriod:    2,
			AlarmPeriod:      60,
			AlarmNotifyGroup: []string{alarmNotifyGroupID},
			Severity:         StrPtr("critical"),
			AlarmPeriodDetail: &AlarmPeriodSetting{
				Sms:            10,
				Phone:          10,
				Email:          10,
				GeneralWebhook: 10,
			},
			UserDefineMsg: StrPtr("this is an alarm"),
		}: nil,
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.CreateAlarm(req)
		suite.validateError(rErr, err)
		if resp != nil {
			suite.alarmList = append(suite.alarmList, resp.AlarmID)
		}
	}
}

func (suite *SDKAlarmTaskTestSuite) TestCreateAlarmAbnormally() {
	alarmNotifyGroupID, err := createAlarmNotifyGroup(suite.cli, uuid.New().String())
	suite.NoError(err)
	suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, alarmNotifyGroupID)

	status := true
	testcases := map[*CreateAlarmRequest]*Error{
		{
			AlarmName: "golang-sdk-test-alarm-1",
			ProjectID: suite.project,
			Status:    &status,
			QueryRequest: []QueryRequest{{
				Query:           "Failed | select count(*) as errNum",
				Number:          1,
				TopicID:         suite.topic,
				StartTimeOffset: -15,
				EndTimeOffset:   0,
			}},
			RequestCycle: RequestCycle{
				Type: "Period",
				Time: 10,
			},
			Condition:        "$1.errNum>0",
			TriggerPeriod:    2,
			AlarmPeriod:      60,
			AlarmNotifyGroup: []string{uuid.New().String()},
			UserDefineMsg:    StrPtr("this is an alarm"),
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "AlarmNotifyGroupNotExist",
			Message:  "Alarm notify group does not exist.",
		},
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.CreateAlarm(req)
		suite.validateError(rErr, err)
		if resp != nil {
			suite.alarmList = append(suite.alarmList, resp.AlarmID)
		}
	}
}

func (suite *SDKAlarmTaskTestSuite) TestDeleteAlarmAbnormally() {
	_, err := suite.cli.DeleteAlarm(&DeleteAlarmRequest{AlarmID: uuid.New().String()})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "AlarmPolicyNotExist",
		Message:  "Alarm policy does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKAlarmTaskTestSuite) TestModifyAlarmNormally() {
	alarmNotifyGroupID, err := createAlarmNotifyGroup(suite.cli, uuid.New().String())
	suite.NoError(err)
	suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, alarmNotifyGroupID)

	alarmList, err := createAlarms(suite.cli, suite.project, suite.topic, alarmNotifyGroupID)
	suite.NoError(err)
	suite.alarmList = append(suite.alarmList, alarmList...)

	alarmID := alarmList[0]
	triggerPeriod := 10
	_, err = suite.cli.ModifyAlarm(&ModifyAlarmRequest{
		AlarmID:       alarmID,
		TriggerPeriod: &triggerPeriod,
	})
	suite.NoError(err)
	resp, err := suite.cli.DescribeAlarms(&DescribeAlarmsRequest{
		ProjectID:     suite.project,
		AlarmPolicyID: &alarmID,
	})
	suite.Equal(triggerPeriod, resp.AlarmPolicies[0].TriggerPeriod)
}

func (suite *SDKAlarmTaskTestSuite) TestModifyAlarmAbnormally() {
	alarmNotifyGroupID, err := createAlarmNotifyGroup(suite.cli, uuid.New().String())
	suite.NoError(err)
	suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, alarmNotifyGroupID)

	alarmList, err := createAlarms(suite.cli, suite.project, suite.topic, alarmNotifyGroupID)
	suite.NoError(err)
	suite.alarmList = append(suite.alarmList, alarmList...)

	alarmID := alarmList[0]
	triggerPeriod := -1
	_, err = suite.cli.ModifyAlarm(&ModifyAlarmRequest{
		AlarmID:       alarmID,
		TriggerPeriod: &triggerPeriod,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key TriggerPeriod, value -1, please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKAlarmTaskTestSuite) TestDescribeAlarmsNormally() {
	alarmNotifyGroupID, err := createAlarmNotifyGroup(suite.cli, uuid.New().String())
	suite.NoError(err)
	suite.alarmNotifyGroupList = append(suite.alarmNotifyGroupList, alarmNotifyGroupID)

	alarmList, err := createAlarms(suite.cli, suite.project, suite.topic, alarmNotifyGroupID)
	suite.NoError(err)
	suite.alarmList = append(suite.alarmList, alarmList...)

	resp, err := suite.cli.DescribeAlarms(&DescribeAlarmsRequest{ProjectID: suite.project})
	suite.NoError(err)
	suite.Equal(len(alarmList), resp.Total)
}

func (suite *SDKAlarmTaskTestSuite) TestDescribeAlarmsAbnormally() {
	_, err := suite.cli.DescribeAlarms(&DescribeAlarmsRequest{ProjectID: uuid.New().String()})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)
}
