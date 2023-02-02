package vms

import (
	"testing"
)

const (
	testAk = "your ak"
	testSk = "your sk"
)

func init() {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestRisk_CanCall(t *testing.T) {
	req := &RiskControlReq{
		CustomerNumberList: "188xxxx1647",
		BusinessLineId:     "abc",
		CallType:           1,
	}
	result, statusCode, err := DefaultInstance.QueryCanCall(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
