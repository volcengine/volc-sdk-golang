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

func TestSMS_send(t *testing.T) {
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
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
