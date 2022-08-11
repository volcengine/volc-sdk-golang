package demo_sms

import (
	"encoding/json"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

type sendSmsTemplateParam struct {
	Code string `json:"code"`
}

/*
send same content to one/multiple phoneNumbers
*/
func TestSmsSend(t *testing.T) {
	testAk := "testAK"
	testSk := "testSk"

	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	c := sendSmsTemplateParam{Code: "111"}
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
