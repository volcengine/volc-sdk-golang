package demo_sms

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func TestGetSmsTemplateAndOrderList(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.GetSmsTemplateAndOrderListRequest{
		SubAccount: "smsAccount",
		Area:       sms.AreaCN,
		PageIndex:  1,
		PageSize:   10,
	}
	result, statusCode, err := sms.DefaultInstance.GetSmsTemplateAndOrderList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsTemplate(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.ApplySmsTemplateRequest{
		SubAccount:  "smsAccount",
		Area:        sms.AreaCN,
		ChannelType: sms.SmsChannelTypeCnMKT,
		Content:     "4张4.5折全场咖啡券仅0.9元→ https://webhook.site/edd2b39a-6c8d-4161-a310-36a470c840d4 生椰拿铁券后低至13.05元 谨防泄漏取阅回T退订",
		Desc:        "测试 SDK",
		Name:        "template_name",
		ShortUrlConfig: &sms.ShortUrlConfig{
			IsEnabled:          sms.EnableStatusEnabled,
			IsNeedClickDetails: sms.EnableStatusNotEnabled,
			UACheckStrategy:    1, // Short link accessible to all devices
		},
	}
	result, statusCode, err := sms.DefaultInstance.ApplySmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteSmsTemplate(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.DeleteSmsTemplateRequest{
		SubAccount: "smsAccount",
		Id:         "idOfTemplateToDelete",
		IsOrder:    true, // 是否是审核工单
	}
	result, statusCode, err := sms.DefaultInstance.DeleteSmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
