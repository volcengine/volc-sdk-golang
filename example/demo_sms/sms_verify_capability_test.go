package demo_sms

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func TestSmsSendVerifyCode(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.SmsVerifyCodeRequest{
		SmsAccount:  "smsAccount",
		Sign:        "sign",
		TemplateID:  "ST_xxx",
		PhoneNumber: "188xxxxxxxx",
		Scene:       "myscene",
		ExpireTime:  1800,
		TryCount:    3,
		CodeType:    6,
		Tag:         "tag",
	}
	result, statusCode, err := sms.DefaultInstance.SendVerifyCode(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)

	if result == nil {
		t.Logf("result is nil")
		return
	}
	if result.ResponseMetadata.Error != nil {
		t.Logf("result.ResponseMetadata.Error = %+v\n", result.ResponseMetadata.Error)
		return
	}
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
}

func TestSmsCheckVerifyCode(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.CheckSmsVerifyCodeRequest{
		SmsAccount:  "smsAccount",
		PhoneNumber: "188xxxxxxxx",
		Scene:       "myscene",
		Code:        "111",
	}
	result, statusCode, err := sms.DefaultInstance.CheckVerifyCode(req)
	/**
	Result = "0" // OTP correct
	Result = "1" // OTP error
	Result = "2" // OTP expired
	*/
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
