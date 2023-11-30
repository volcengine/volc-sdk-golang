package tls

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 请填写您的ProjectId和TopicId
	projectID := "your-project-id"
	topicID := "your-topic-id"

	// 创建告警组
	// 请根据您的需要，填写AlarmNotifyGroupName、NotifyType和Receivers等参数
	// CreateAlarmNotifyGroup API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112220
	createAlarmNotifyGroupResp, _ := client.CreateAlarmNotifyGroup(&tls.CreateAlarmNotifyGroupRequest{
		GroupName:  uuid.NewString(),
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
	})
	alarmNotifyGroupID := createAlarmNotifyGroupResp.NotifyGroupID

	// 创建告警策略
	// 请根据您的需要，填写ProjectId、AlarmName、QueryRequest、RequestCycle、Condition、AlarmPeriod、AlarmNotifyGroup等参数
	// CreateAlarm API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112216
	createAlarmStatus := true
	createAlarmResp, _ := client.CreateAlarm(&tls.CreateAlarmRequest{
		AlarmName: uuid.NewString(),
		ProjectID: projectID,
		Status:    &createAlarmStatus,
		QueryRequest: tls.QueryRequests{
			{
				Query:           "*|select count(*) as count",
				Number:          1,
				StartTimeOffset: -3,
				TopicID:         topicID,
				TopicName:       "your-topic-name",
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
		AlarmNotifyGroup: []string{alarmNotifyGroupID},
	})
	alarmID := createAlarmResp.AlarmID

	// 修改告警策略
	// 请根据您的需要，填写待修改的AlarmId和其它参数
	// ModifyAlarm API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112218
	modifyUserDefineMsg := "modifyUserDefineMsg"
	modifyStatus := true
	modifyAlarmName := uuid.NewString()
	modifyCondition := "$1.count > 0"
	modifyTriggerPeriod := 1
	modifyAlarmPeriod := 10
	modifyNotifyGroup := []string{alarmNotifyGroupID}
	_, _ = client.ModifyAlarm(&tls.ModifyAlarmRequest{
		AlarmID:   alarmID,
		AlarmName: &modifyAlarmName,
		Status:    &modifyStatus,
		QueryRequest: &tls.QueryRequests{
			{
				Query:           "*|select count(*) as count",
				Number:          1,
				StartTimeOffset: -3,
				TopicID:         topicID,
				TopicName:       "your-topic-name",
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

	// 查询告警策略
	// 请根据您的需要，填写待查询的ProjectId
	// DescribeAlarms API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112219
	describeAlarmsResp, _ := client.DescribeAlarms(&tls.DescribeAlarmsRequest{
		ProjectID: projectID,
		TopicID:   &topicID,
	})
	fmt.Println(describeAlarmsResp.Total)

	// 删除告警策略
	// 请根据您的需要，填写待删除的AlarmId
	// DeleteAlarm API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112217
	_, _ = client.DeleteAlarm(&tls.DeleteAlarmRequest{
		AlarmID: alarmID,
	})

	// 删除告警组
	// 请根据您的需要，填写待删除的AlarmNotifyGroupId
	// DeleteAlarmNotifyGroup API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112221
	_, _ = client.DeleteAlarmNotifyGroup(&tls.DeleteAlarmNotifyGroupRequest{
		AlarmNotifyGroupID: alarmNotifyGroupID,
	})
}
