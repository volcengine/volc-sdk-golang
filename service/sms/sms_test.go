package sms

import (
	"encoding/json"
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"
)

const (
	testAk = "testAK"
	testSk = "testSk"
)

type Code struct {
	Code string `json:"code"`
}

func TestSMS_send(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	c := Code{Code: "111"}
	cj, _ := json.Marshal(c)
	req := &SmsRequest{
		SmsAccount:    "smsAccount",
		Sign:          "sign",
		TemplateID:    "ST_xxx",
		TemplateParam: string(cj),
		PhoneNumbers:  "188xxxxxxxx",
		Tag:           "tag",
	}
	result, statusCode, err := DefaultInstance.Send(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSMS_sendI18n(t *testing.T) {
	i18nInstance := NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	c := Code{Code: "111"}
	cj, _ := json.Marshal(c)
	req := &SmsRequest{
		SmsAccount:    "smsAccount",
		Sign:          "sign",
		TemplateID:    "ST_xxx",
		TemplateParam: string(cj),
		PhoneNumbers:  "188xxxxxxxx",
		Tag:           "tag",
	}
	result, statusCode, err := i18nInstance.Send(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSMSVerifyCode_send(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &SmsVerifyCodeRequest{
		SmsAccount:  "smsAccount",
		Sign:        "sign",
		TemplateID:  "ST_xxx",
		PhoneNumber: "188xxxxxxxx",
		Scene:       "123",
		CodeType:    6,
		ExpireTime:  240,
		TryCount:    3,
	}
	result, statusCode, err := DefaultInstance.SendVerifyCode(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSMSVerifyCode_check(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &CheckSmsVerifyCodeRequest{
		SmsAccount:  "smsAccount",
		Scene:       "123",
		Code:        "123456",
		PhoneNumber: "188xxxxxxxx",
	}
	result, statusCode, err := DefaultInstance.CheckVerifyCode(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
