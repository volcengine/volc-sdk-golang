package demo_sms

import (
	"encoding/json"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

type sendBatchSmsTemplateParam struct {
	Code string `json:"code"`
}

func TestSmsBatchSend(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	c := sendBatchSmsTemplateParam{Code: "111"}
	cj, _ := json.Marshal(c)
	messages := make([]*sms.SmsBatchMessages, 0)
	messages = append(messages, &sms.SmsBatchMessages{
		TemplateParam: string(cj),
		PhoneNumber:   "188xxxxxxxx",
	})

	req := &sms.SmsBatchRequest{
		SmsAccount: "smsAccount",
		Sign:       "sign",
		TemplateID: "ST_xxx",
		Tag:        "tag",
		Messages:   messages,
	}
	result, statusCode, err := sms.DefaultInstance.BatchSend(req)
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
