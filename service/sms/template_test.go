package sms

import (
	"testing"
)

func TestGetSmsTemplateList(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &GetSmsTemplateAndOrderListRequest{
		SubAccount: "smsAccount",
		Area:       AreaCN,
	}
	result, statusCode, err := DefaultInstance.GetSmsTemplateAndOrderList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsTemplate(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &ApplySmsTemplateRequest{
		SubAccount:  "smsAccount",
		Area:        AreaCN,
		ChannelType: SmsChannelTypeCnMKT,
		Content:     "4张4.5折全场咖啡券仅0.9元→ https://webhook.site/edd2b39a-6c8d-4161-a310-36a470c840d4 拿铁券后低至13.05元 谨防泄漏取阅回T退订",
		Desc:        "测试 SDK",
		Name:        "template_name",
		ShortUrlConfig: &ShortUrlConfig{
			IsEnabled:          EnableStatusEnabled,
			IsNeedClickDetails: EnableStatusNotEnabled,
		},
	}
	result, statusCode, err := DefaultInstance.ApplySmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteSmsTemplate(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &DeleteSmsTemplateRequest{
		SubAccount: "smsAccount",
		Id:         "idOfTemplateToDelete",
		IsOrder:    true, // 是否是审核工单
	}
	result, statusCode, err := DefaultInstance.DeleteSmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
