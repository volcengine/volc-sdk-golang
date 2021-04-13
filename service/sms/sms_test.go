package sms

import (
	"encoding/json"
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
