package demo_sms

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func TestBulkCreateTobTrafficDrivingPhone(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.BulkCreateTobTrafficDrivingPhoneRequest{
		Data: []sms.CreateTobTrafficDrivingPhoneItem{
			{
				NumberType:     sms.MobilePhone,      // 手机号类型（枚举：MobilePhone=1）
				Number:         "13800138001",        // 手机号
				Company:        "",                   // 企业名称
				NumberPerson:   "经办人姓名",              // 联系人姓名
				PersonType:     sms.IdNumber,         // 人员类型（枚举：IdNumber=1）
				PersonId:       "150203199609070812", // 身份证号
				NumberProvince: "北京市",                // 手机号归属省
				NumberCity:     "北京市",                // 手机号归属市
			},
			{
				NumberType:     sms.MobilePhone,
				Number:         "13900139002",
				Company:        "",
				NumberPerson:   "经办人姓名2",
				PersonType:     sms.IdNumber,
				PersonId:       "150203199609070812",
				NumberProvince: "上海市",
				NumberCity:     "上海市",
			},
		},
	}

	resp, statusCode, err := sms.DefaultInstance.BulkCreateTobTrafficDrivingPhone(req)
	t.Logf("resp = %+v\n", resp)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
	t.Logf("error = %v\n", resp.ResponseMetadata.Error)
}

func TestGetTobTrafficDrivingPhoneList(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.GetTobTrafficDrivingPhoneListRequest{
		Number: "13900139002", // 待查询的手机号
	}

	resp, statusCode, err := sms.DefaultInstance.GetTobTrafficDrivingPhoneList(req)
	t.Logf("resp = %+v\n", resp)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
	t.Logf("error = %v\n", resp.ResponseMetadata.Error)
}

func TestUpdateTobTrafficDrivingPhone(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.UpdateTobTrafficDrivingPhoneRequest{
		NumberType:     sms.MobilePhone,      // 手机号类型
		Number:         "13800138000",        // 待更新的手机号
		Company:        "",                   // 企业名称（更新后）
		NumberPerson:   "更新后经办人姓名",           // 联系人姓名（更新后）
		PersonType:     sms.IdNumber,         // 人员类型
		PersonId:       "150203199609070812", // 身份证号（更新后）
		NumberProvince: "北京市",                // 手机号归属省
		NumberCity:     "北京市",                // 手机号归属市
		UpdatePersonId: true,                 // 是否更新人员 ID
	}

	resp, statusCode, err := sms.DefaultInstance.UpdateTobTrafficDrivingPhone(req)
	t.Logf("resp = %+v\n", resp)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %v\n", err)
	t.Logf("error = %v\n", resp.ResponseMetadata.Error)
}

func TestDeleteTobTrafficDrivingPhone(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.DeleteTobTrafficDrivingPhoneRequest{
		Number: "13800138000", // 待删除的手机号
	}

	result, statusCode, err := sms.DefaultInstance.DeleteTobTrafficDrivingPhone(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestBulkCreateTobTrafficDrivingLink(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	useVolcLink := true
	req := &sms.BulkCreateTobTrafficDrivingLinkRequest{
		Data: []sms.BulkCreateTobTrafficDrivingLinkItem{
			{
				UseVolcLink:           &useVolcLink,                        // 是否使用火山短链
				Protocol:              "https",                             // 协议类型
				Link:                  "https://example.com/origin-link-1", // 原始链接
				LinkDomainIcp:         "京ICP备123456号",                      // 原始链接 ICP 备案号
				LinkDomainIcpBody:     "企业主体名称",                            // 原始链接 ICP 备案主体
				JumpLink:              "https://example.com/jump-link-1",   // 跳转链接
				JumpLinkDomainIcp:     "京ICP备654321号",                      // 跳转链接 ICP 备案号
				JumpLinkDomainIcpBody: "企业主体名称2",                           // 跳转链接 ICP 备案主体
			},
			{
				UseVolcLink:           &useVolcLink,
				Protocol:              "https",
				Link:                  "https://example.com/origin-link-2",
				LinkDomainIcp:         "沪ICP备789012号",
				LinkDomainIcpBody:     "企业主体名称3",
				JumpLink:              "https://example.com/jump-link-2",
				JumpLinkDomainIcp:     "沪ICP备210987号",
				JumpLinkDomainIcpBody: "企业主体名称4",
			},
		},
	}

	result, statusCode, err := sms.DefaultInstance.BulkCreateTobTrafficDrivingLink(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
	t.Logf("error = %v\n", result.ResponseMetadata.Error)
}

func TestGetTobTrafficDrivingLinkList(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.GetTobTrafficDrivingLinkListRequest{
		Link:     "",                 // 待查询的原始链接
		JumpLink: "https://xxxx.com", // 待查询的跳转链接
	}

	result, statusCode, err := sms.DefaultInstance.GetTobTrafficDrivingLinkList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteTobTrafficDrivingLink(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.DeleteTobTrafficDrivingLinkRequest{
		Link:     "",                 // 待删除的原始链接
		JumpLink: "https://xxxx.com", // 待删除的跳转链接
	}

	result, statusCode, err := sms.DefaultInstance.DeleteTobTrafficDrivingLink(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestListRelationTemplate(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.ListRelationTemplateRequest{
		Signature: "",                 // 待关联的签名
		Number:    "",                 // 关联手机号
		Link:      "",                 // 关联原始链接
		JumpLink:  "https://xxxx.com", // 关联跳转链接
	}

	result, statusCode, err := sms.DefaultInstance.GetRelationTemplateList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
