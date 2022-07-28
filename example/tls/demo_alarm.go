package tls

import (
	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	//初始化客户端，配置AccessKeyID,AccessKeySecret,region,securityToken;securityToken可以为空
	client := tls.NewClient(testEndPoint, testAk, testSk, testSessionToken, testRegion)

	//新建project
	createProjectResp, _ := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: testPrefix + uuid.NewString(),
		Description: "",
		Region:      testRegion,
	})

	testProjectID := createProjectResp.ProjectID

	//新建topic
	topicName := testPrefix + uuid.NewString()
	topic, _ := client.CreateTopic(&tls.CreateTopicRequest{
		ProjectID:   testProjectID,
		TopicName:   topicName,
		Ttl:         30,
		ShardCount:  2,
		Description: "topic desc",
	})
	testTopicID := topic.TopicID

	//create alarm notify group
	createAlarmGroupReq := &tls.CreateAlarmNotifyGroupRequest{
		GroupName:  testPrefix + uuid.NewString(),
		NoticeType: tls.NoticeTypes{"Recovery", "Trigger"},
		Receivers: []tls.Receiver{
			{
				ReceiverType:     "User",
				ReceiverNames:    []string{"yutao001"},
				ReceiverChannels: []tls.ReceiverChannel{"Email"},
				StartTime:        "00:00:00",
				EndTime:          "23:59:59",
			},
		},
	}

	createAlarmNotifyGroupRes, _ := client.CreateAlarmNotifyGroup(createAlarmGroupReq)

	notifyGroupId := createAlarmNotifyGroupRes.NotifyGroupID

	// create alarm
	createAlarmStatus := true
	createAlarmReq := &tls.CreateAlarmRequest{
		AlarmName: testPrefix + uuid.NewString(),
		ProjectID: testProjectID,
		Status:    &createAlarmStatus,
		QueryRequest: tls.QueryRequests{
			{
				Query:           "*|select count(*) as count",
				Number:          1,
				StartTimeOffset: -3,
				TopicID:         testTopicID,
				TopicName:       topicName,
				EndTimeOffset:   0,
			},
		},
		RequestCycle: tls.RequestCycle{
			Type: "period",
			Time: 1,
		},
		Condition:        "$1.count > 0",
		TriggerPeriod:    1,
		AlarmPeriod:      10,
		AlarmNotifyGroup: []string{notifyGroupId},
	}

	resp, _ := client.CreateAlarm(createAlarmReq)
	alarmID := resp.AlarmID

	// modify alarm
	modifyUserDefineMsg := "modifyUserDefineMsg"
	modifyStatus := true
	modifyAlarmName := testPrefix + uuid.NewString()
	modifyCondition := "$1.count > 0"
	modifyTriggerPeriod := 1
	modifyAlarmPeriod := 10
	modifyNotifyGroup := []string{notifyGroupId}
	_, _ = client.ModifyAlarm(&tls.ModifyAlarmRequest{
		AlarmID:   alarmID,
		AlarmName: &modifyAlarmName,
		Status:    &modifyStatus,
		QueryRequest: &tls.QueryRequests{
			{
				Query:           "*|select count(*) as count",
				Number:          1,
				StartTimeOffset: -3,
				TopicID:         testTopicID,
				TopicName:       topicName,
				EndTimeOffset:   0,
			},
		},
		RequestCycle: &tls.RequestCycle{
			Type: "period",
			Time: 1,
		},
		Condition:        &modifyCondition,
		TriggerPeriod:    &modifyTriggerPeriod,
		AlarmPeriod:      &modifyAlarmPeriod,
		AlarmNotifyGroup: &modifyNotifyGroup,
		UserDefineMsg:    &modifyUserDefineMsg,
	})

	// describe alarms
	_, _ = client.DescribeAlarms(&tls.DescribeAlarmsRequest{
		ProjectID: testProjectID,
		TopicID:   &testTopicID,
	})

	// delete alarm
	_, _ = client.DescribeAlarms(&tls.DescribeAlarmsRequest{
		ProjectID:     testProjectID,
		AlarmPolicyID: &alarmID,
	})
}
