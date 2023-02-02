package demo_sms

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func TestApplyVmsTemplate(t *testing.T) {
	fileBase64String, _ := getBase64StringFromFile("/Users/Pictures/视频.mp4")
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.ApplyVmsTemplateRequest{
		SubAccount: "smsAccount",
		Name:       "SDK测试",
		Theme:      "SDK测试视频短信",
		Signature:  "sign",
		Contents: []sms.VmsElement{
			{
				SourceType: sms.SourceTypeText,
				Content:    "测试${code}",
			},
			{
				SourceType: sms.SourceTypeVideo,
				Content:    fileBase64String,
			},
		},
	}
	result, statusCode, err := sms.DefaultInstance.ApplyVmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestGetVmsTemplateStatus(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.GetVmsTemplateStatusRequest{
		SubAccount: "smsAccount",
		TemplateId: "VT_xxx",
	}
	result, statusCode, err := sms.DefaultInstance.GetVmsTemplateStatus(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

type sendVmsTemplateParam struct {
	Code string `json:"code"`
}

func TestSendVms(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	c := sendVmsTemplateParam{Code: "111"}
	cj, _ := json.Marshal(c)
	req := &sms.SendVmsRequest{
		SmsAccount:    "smsAccount",
		TemplateID:    "VT_xxx",
		TemplateParam: string(cj),
		PhoneNumbers:  "188xxxxxxxx",
		Tag:           "tag",
	}
	result, statusCode, err := sms.DefaultInstance.SendVms(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func getBase64StringFromFile(filename string) (string, error) {
	srcByte, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	res := base64.StdEncoding.EncodeToString(srcByte)
	return res, nil
}
