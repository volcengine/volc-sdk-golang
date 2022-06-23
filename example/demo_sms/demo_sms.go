package demo_sms

import (
	"encoding/json"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

const (
	testAk = "testAK"
	testSk = "testSk"
)

type Code struct {
	Code string `json:"code"`
}

func TestSmsSend(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	c := Code{Code: "111"}
	cj, _ := json.Marshal(c)
	req := &sms.SmsRequest{
		SmsAccount:    "smsAccount",
		Sign:          "sign",
		TemplateID:    "ST_xxx",
		TemplateParam: string(cj),
		PhoneNumbers:  "188xxxxxxxx",
		Tag:           "tag",
	}
	result, statusCode, err := sms.DefaultInstance.Send(req)
	t.Logf("result = %+v\n", result)
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSmsBatchSend(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	c := Code{Code: "111"}
	cj, _ := json.Marshal(c)
	messages := make([]*sms.SmsBatchMessages, 0)
	messages = append(messages, &sms.SmsBatchMessages{
		TemplateParam: string(cj),
		PhoneNumber:   "188xxxxxxxx",
	})

	req := &sms.SmsBatchRequest{
		SmsAccount: "smsAccount",
		From:       "BytePlus",
		TemplateID: "ST_xxx",
		Tag:        "tag",
		Messages:   messages,
	}
	result, statusCode, err := sms.DefaultInstance.BatchSend(req)
	t.Logf("result = %+v\n", result)
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSmsConversion(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.ConversionRequest{
		MessageIDs: []string{"MessageID"},
	}
	result, statusCode, err := sms.DefaultInstance.Conversion(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

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
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
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
