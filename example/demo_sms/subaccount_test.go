package demo_sms

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func TestGetSubAccountList(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.GetSubAccountListRequest{
		PageIndex: 1,
		PageSize:  10,
	}
	result, statusCode, err := sms.DefaultInstance.GetSubAccountList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestGetSubAccountDetail(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.GetSubAccountDetailRequest{
		SubAccount: "smsAccount",
	}
	result, statusCode, err := sms.DefaultInstance.GetSubAccountDetail(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestInsertSmsSubAccount(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.InsertSmsSubAccountReq{
		SubAccountName: "testsdk",
	}
	result, statusCode, err := sms.DefaultInstance.InsertSmsSubAccount(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
