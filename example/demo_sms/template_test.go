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

func TestApplyTemplateV2(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	// 构造「短链配置」
	shortUrlConfig := sms.ShortUrlConfig{
		IsEnabled:          sms.EnableStatusEnabled,
		IsNeedClickDetails: sms.EnableStatusEnabled,
		UACheckStrategy:    1,
	}

	// 构造「模板参数」
	templateParams := []sms.TemplateParamsV2{
		{
			Name:      "username",
			ParamType: 2,
			Location:  "body",
			Content:   "用户昵称",
		},
	}

	// 构造「流量驱动参数」
	trafficDriving := [][]sms.TemplateParamWithTrafficDriving{
		{
			{
				Name:      "phone1",
				ParamType: 3,
				Content:   "18800001110",
			},
		},
	}

	req := &sms.ApplyTemplateV2Request{
		SubAccounts:            []string{"sub-account-1"},
		TemplateId:             "template-123",                        // 模板 ID
		SecondTemplateId:       "second-template-456",                 // 二级文案 ID
		Content:                "您的验证码是：${verification_code}，5分钟内有效。", // 模板内容
		ChannelType:            "SMS",                                 // 渠道类型
		Area:                   "",                                    // 地区
		Name:                   "测试模板V2",                              // 模板名称
		ShortUrlConfig:         shortUrlConfig,                        // 短链配置
		Desc:                   "测试短信模板描述",                            // 模板描述
		Signatures:             []string{"签名A", "签名B"},                // 关联签名列表
		TemplateParams:         templateParams,                        // 模板参数
		TemplateTrafficDriving: trafficDriving,                        // 流量驱动参数
	}

	result, statusCode, err := sms.DefaultInstance.ApplySmsTemplateV2(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
	t.Logf("error = %v\n", result.ResponseMetadata.Error)
}

func TestListSmsTemplateV2(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.ListSmsTemplateV2Request{
		SubAccounts: []string{"sub-account-1", "sub-account-2"}, // 子账号列表
		Page:        1,                                          // 页码
		PageSize:    10,                                         // 页大小
		TemplateId:  "template-123",                             // 模板 ID
		Signature:   "示例签名",                                     // 关联签名
		Name:        "模板名称关键词",                                  // 模板名称关键词
		Content:     "验证码",                                      // 模板内容关键词
	}

	result, statusCode, err := sms.DefaultInstance.ListSmsTemplateV2(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestBindSignature(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	// 构造「流量驱动参数」
	trafficDriving := [][]sms.TemplateParamWithTrafficDriving{
		{
			{
				Name:      "phone1",      // 流量参数名
				ParamType: 3,             // 流量参数类型
				Content:   "18800001110", // 流量参数内容
			},
		},
	}

	req := &sms.BindSignatureRequest{
		SubAccounts:            []string{"sub-account-1"}, // 子账号列表
		TemplateId:             "template-123",            // 模板 ID
		Signatures:             []string{"签名A"},           // 关联签名列表
		TemplateTrafficDriving: trafficDriving,            // 流量驱动参数
	}

	result, statusCode, err := sms.DefaultInstance.BindSignatures(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestListSecondTemplate(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.ListSecondTemplateRequest{
		SubAccounts:          []string{"sub-account-1"},                          // 子账号列表
		TemplateId:           "template-123",                                     // 模板 ID
		TemplateIdList:       []string{"template-456", "template-789"},           // 模板 ID 列表
		SignatureList:        []string{"签名B", "签名C"},                             // 关联签名列表
		SecondTemplateId:     "second-template-1",                                // 二级模板 ID（精确匹配）
		SecondTemplateIdList: []string{"second-template-2", "second-template-3"}, // 二级模板 ID 列表
		StatusList:           []int32{1, 3},                                      // 状态列表（1=审核中，3=审核通过）
	}

	result, statusCode, err := sms.DefaultInstance.ListSecondTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestBindTrafficDrivingParams(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	// 构造「流量驱动参数」
	trafficDriving := [][]sms.TemplateParamWithTrafficDriving{
		{
			{
				Name:      "phone",
				ParamType: 3,
				Content:   "13030003000",
			},
		},
	}

	req := &sms.BindTrafficDrivingParamsRequest{
		SecondTemplateId:       "second-template-123", // 二级模板 ID
		SubContentId:           "sub-content-456",     // 子内容 ID
		TemplateTrafficDriving: trafficDriving,
	}

	result, statusCode, err := sms.DefaultInstance.BindTrafficDrivingParams(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
	t.Logf("error = %v\n", result.ResponseMetadata.Error)
}

func TestListSubContent(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.ListSubContentRequest{
		SecondTemplateId: "second-template-123", // 二级模板 ID
		Page:             1,                     // 页码
		PageSize:         10,                    // 页大小
		SubContentId:     "sub-content-456",     // 子内容 ID（精确查询）
	}

	result, statusCode, err := sms.DefaultInstance.ListSubContent(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
	t.Logf("error = %v\n", result.ResponseMetadata.Error)
}

func TestCreateSubContentTemplate(t *testing.T) {
	// 这个有依赖该账号的其它数据，实际测试时，需要将AKSK替换成正确的
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	templateParamsReq := make([]sms.TemplateParamsV2, 0)
	templateParamsReq = append(templateParamsReq, sms.TemplateParamsV2{
		Name:      "code",
		ParamType: 1,
	})
	templateParamsReq = append(templateParamsReq, sms.TemplateParamsV2{
		Name:      "phone",
		ParamType: 3,
	})
	templateParamsReq = append(templateParamsReq, sms.TemplateParamsV2{
		Name:      "link",
		ParamType: 2,
	})

	phoneList := make([]sms.TemplateParamWithTrafficDriving, 0)
	phoneList = append(phoneList, sms.TemplateParamWithTrafficDriving{
		Name:      "phone",
		ParamType: 3,
		Content:   "***",
	})
	phoneList = append(phoneList, sms.TemplateParamWithTrafficDriving{
		Name:      "phone",
		ParamType: 3,
		Content:   "***",
	})
	linkList := make([]sms.TemplateParamWithTrafficDriving, 0)
	linkList = append(linkList, sms.TemplateParamWithTrafficDriving{
		Name:      "link",
		ParamType: 2,
		Content:   "***.com/",
	})
	linkList = append(linkList, sms.TemplateParamWithTrafficDriving{
		Name:      "link",
		ParamType: 2,
		Content:   "https://***.com",
	})
	templateTrafficDriving := make([][]sms.TemplateParamWithTrafficDriving, 0)
	for _, phone := range phoneList {
		for _, link := range linkList {
			templateTrafficDriving = append(templateTrafficDriving, []sms.TemplateParamWithTrafficDriving{phone, link})
		}
	}
	req := &sms.CreateSubContentTemplate{
		TemplateId:             "S1T_85e6721w",                      // 一级模版ID
		Signature:              "huoshan测试",                         // 签名
		SecondTemplateId:       "S2T_8615313e",                      // 二级模版ID
		Industry:               "universal",                         // 行业
		Content:                "test-${code}-${phone}-${link}-***", // 模版内容
		TemplateParamsReq:      templateParamsReq,
		TemplateTrafficDriving: templateTrafficDriving,
	}

	result, statusCode, err := sms.DefaultInstance.ApplySmsSubContentTemplateV2(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
	t.Logf("error = %v\n", result.ResponseMetadata.Error)
}
