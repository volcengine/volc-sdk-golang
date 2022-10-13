package sms

import (
	"testing"
)

func TestGetSignatureList(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &GetSignatureAndOrderListRequest{
		SubAccount: "smsAccount",
	}
	result, statusCode, err := DefaultInstance.GetSignatureAndOrderList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsSignature(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &ApplySmsSignatureRequest{
		Desc:       "测试 SDK",
		SubAccount: "smsAccount",
		Content:    "sign",
		Source:     "公司全称/简称",
	}
	result, statusCode, err := DefaultInstance.ApplySmsSignature(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteSignature(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &DeleteSignatureRequest{
		SubAccount: "smsAccount",
		Id:         "idOfSignatureToDelete",
		IsOrder:    true,
	}
	result, statusCode, err := DefaultInstance.DeleteSignature(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
