package demo_sms

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func TestGetSignatureAndOrderList(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.GetSignatureAndOrderListRequest{
		SubAccount: "smsAccount",
		Signature:  "sign",
		PageIndex:  1,
		PageSize:   10,
	}
	result, statusCode, err := sms.DefaultInstance.GetSignatureAndOrderList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsSignature(t *testing.T) {
	fileBase64String, _ := getBase64StringFromFile("/Users//Pictures/IMG_9033.JPG")

	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.ApplySmsSignatureRequest{
		Desc:       "测试 SDK",
		SubAccount: "smsAccount",
		Content:    "sign",
		Domain:     "http://www.xxx.com",
		Source:     sms.SignSourceTypeBrand,
		UploadFileList: []sms.SignAuthFile{
			{
				FileType:    sms.DocTypeThreeInOne,
				FileContent: fileBase64String,
				FileSuffix:  "jpg",
			},
		},
		Purpose: sms.SignPurposeForOwn,
	}
	result, statusCode, err := sms.DefaultInstance.ApplySmsSignature(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteSignature(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.DeleteSignatureRequest{
		SubAccount: "smsAccount",
		Id:         "idOfSignatureToDelete",
		IsOrder:    true,
	}
	result, statusCode, err := sms.DefaultInstance.DeleteSignature(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
