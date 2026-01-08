package tls

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKAlarmContentTemplateTestSuite struct {
	suite.Suite

	cli          Client
	templateID   string
	templateName string
}

func TestSDKAlarmContentTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(SDKAlarmContentTemplateTestSuite))
}

func (suite *SDKAlarmContentTemplateTestSuite) SetupSuite() {
	suite.cli = NewClientWithEnv()

	token := strings.ReplaceAll(uuid.New().String(), "-", "")[:12]
	suite.templateName = "golang-sdk-test-alarm-content-template-" + token

	req := &CreateAlarmContentTemplateRequest{
		AlarmContentTemplateName: suite.templateName,
		Email: &EmailContentTemplateInfo{
			Subject: StrPtr("告警通知"),
			Content: StrPtr("告警策略：{{Alarm}}<br> 告警日志项目：{{ProjectName}}<br>"),
			Locale:  StrPtr("zh-CN"),
		},
		DingTalk: &DingTalkContentTemplateInfo{
			Title:   StrPtr("告警通知"),
			Content: StrPtr("尊敬的用户，您好！\n您的账号（主账户ID：{{AccountID}} ）的日志服务{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\n告警策略：{{Alarm}}\n告警日志主题：{{AlarmTopicName}}\n触发时间：{{StartTime}}\n触发条件：{{Condition}}\n当前查询结果：[{%-for x in TriggerParams-%}]{{-x-}} {%-endfor-%}]\n通知内容：{{NotifyMsg|escapejs}}\n日志检索详情：[查看详情]({{QueryUrl}})\n告警详情：[查看详情]({{SignInUrl}})\n\n感谢对火山引擎的支持"),
			Locale:  StrPtr("zh-CN"),
		},
		Lark: &LarkContentTemplateInfo{
			Title:   StrPtr("告警通知"),
			Content: StrPtr("尊敬的用户，您好！\n您的账号（主账户ID：{{AccountID}} ）的日志服务{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\n告警策略：{{Alarm}}\n告警日志主题：{{AlarmTopicName}}\n触发时间：{{StartTime}}\n触发条件：{{Condition}}\n当前查询结果：[{%-for x in TriggerParams-%}]{{-x-}} {%-endfor-%}]\n通知内容：{{NotifyMsg|escapejs}}\n日志检索详情：[查看详情]({{QueryUrl}})\n告警详情：[查看详情]({{SignInUrl}})\n\n感谢对火山引擎的支持"),
			Locale:  StrPtr("zh-CN"),
		},
		WeChat: &WeChatContentTemplateInfo{
			Title:   StrPtr("告警通知"),
			Content: StrPtr("尊敬的用户，您好！\n您的账号（主账户ID：{{AccountID}} ）的日志服务{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\n告警策略：{{Alarm}}\n告警日志主题：{{AlarmTopicName}}\n触发时间：{{StartTime}}\n触发条件：{{Condition}}\n当前查询结果：[{%-for x in TriggerParams-%}]{{-x-}} {%-endfor-%}]\n通知内容：{{NotifyMsg|escapejs}}\n日志检索详情：[查看详情]({{QueryUrl}})\n告警详情：[查看详情]({{SignInUrl}})\n\n感谢对火山引擎的支持"),
			Locale:  StrPtr("zh-CN"),
		},
		Sms: &SmsContentTemplateInfo{
			Content: StrPtr("告警策略{{Alarm}}， 告警日志项目：{{ProjectName}}， 告警日志主题：{{AlarmTopicName}}， 告警级别：{{Severity}}， 通知类型：{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}，触发时间：{{StartTime}}， 触发条件：{{Condition}}， 当前查询结果：[{%-for x in TriggerParams-%}]{{-x-}} {%-endfor-%}]， 通知内容：{{NotifyMsg}}"),
			Locale:  StrPtr("zh-CN"),
		},
		Vms: &VmsContentTemplateInfo{
			Content: StrPtr("通知类型：{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}"),
			Locale:  StrPtr("zh-CN"),
		},
		Webhook: &WebhookContentTemplateInfo{
			Content: StrPtr("{ \"msg_type\": \"interactive\", \"card\": { \"config\": { \"wide_screen_mode\": true }, \"elements\": [ { \"content\": \"尊敬的用户，您好！\n您的账号（主账户ID：{{AccountID}} ）的日志服务{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\n告警策略：{{Alarm}}\n告警日志主题：{{AlarmTopicName}}\n触发时间：{{StartTime}}\n触发条件：{{Condition}}\n当前查询结果：[{%-for x in TriggerParams-%}]{{-x-}} {%-endfor-%}];\n通知内容：{{NotifyMsg|escapejs}}\n\n感谢对火山引擎支持\", \"tag\": \"markdown\" } ], \"header\": { \"template\": \"{%if NotifyType==1%}red{%else%}green{%endif%}\", \"title\": { \"content\": \"【火山引擎】【日志服务】{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\", \"tag\": \"plain_text\" } } } } }"),
		},
		NeedValidContent: BoolPtr(true),
	}

	resp, err := suite.cli.CreateAlarmContentTemplate(req)
	suite.NoError(err)
	suite.NotNil(resp)
	suite.NotEmpty(resp.AlarmContentTemplateID)

	suite.templateID = resp.AlarmContentTemplateID
}

func (suite *SDKAlarmContentTemplateTestSuite) TearDownSuite() {
	if suite.templateID == "" {
		return
	}
	_, err := suite.cli.DeleteAlarmContentTemplate(&DeleteAlarmContentTemplateRequest{AlarmContentTemplateID: suite.templateID})
	suite.NoError(err)
}

func (suite *SDKAlarmContentTemplateTestSuite) TestDescribeAlarmContentTemplates() {
	request := &DescribeAlarmContentTemplatesRequest{
		PageNumber: 1,
		PageSize:   20,
	}

	response, err := suite.cli.DescribeAlarmContentTemplates(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.GreaterOrEqual(response.Total, 0)
	suite.NotNil(response.AlarmContentTemplates)

	asc := true
	requestWithParams := &DescribeAlarmContentTemplatesRequest{
		PageNumber: 1,
		PageSize:   10,
		ASC:        &asc,
	}

	responseWithParams, err := suite.cli.DescribeAlarmContentTemplates(requestWithParams)
	suite.NoError(err)
	suite.NotNil(responseWithParams)
	suite.GreaterOrEqual(responseWithParams.Total, 0)
	suite.NotNil(responseWithParams.AlarmContentTemplates)
}

func (suite *SDKAlarmContentTemplateTestSuite) TestModifyAlarmContentTemplate() {
	newName := suite.templateName + "-modified"
	resp, err := suite.cli.ModifyAlarmContentTemplate(&ModifyAlarmContentTemplateRequest{
		AlarmContentTemplateID:   suite.templateID,
		AlarmContentTemplateName: newName,
		Email: &EmailContentTemplateInfo{
			Subject: StrPtr("告警通知"),
			Content: StrPtr("告警策略：{{Alarm}}<br> 告警日志项目：{{ProjectName}}<br>"),
			Locale:  StrPtr("zh-CN,en-US"),
		},
		DingTalk: &DingTalkContentTemplateInfo{
			Title:   StrPtr("告警通知"),
			Content: StrPtr("尊敬的用户，您好！\n您的账号（主账户ID：{{AccountID}} ）的日志服务{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\n告警策略：{{Alarm}}\n告警日志主题：{{AlarmTopicName}}\n触发时间：{{StartTime}}\n触发条件：{{Condition}}\n当前查询结果：[{%-for x in TriggerParams-%}]{{-x-}} {%-endfor-%}]\n通知内容：{{NotifyMsg|escapejs}}\n日志检索详情：[查看详情]({{QueryUrl}})\n告警详情：[查看详情]({{SignInUrl}})\n\n感谢对火山引擎的支持"),
			Locale:  StrPtr("zh-CN"),
		},
		NeedValidContent: BoolPtr(true),
	})
	suite.NoError(err)
	suite.NotNil(resp)

	suite.templateName = newName
}

func (suite *SDKAlarmContentTemplateTestSuite) TestZ_DeleteAlarmContentTemplate() {
	_, err := suite.cli.DeleteAlarmContentTemplate(&DeleteAlarmContentTemplateRequest{AlarmContentTemplateID: suite.templateID})
	suite.NoError(err)

	suite.templateID = ""
}
