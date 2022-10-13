package sms

import (
	"testing"
)

func TestGetSubAccountList(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &GetSubAccountListRequest{
		SubAccount: "smsAccount",
		PageIndex:  1,
		PageSize:   10,
	}
	result, statusCode, err := DefaultInstance.GetSubAccountList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestGetSubAccountDetail(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	req := &GetSubAccountDetailRequest{
		SubAccount: "smsAccount",
	}
	result, statusCode, err := DefaultInstance.GetSubAccountDetail(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
