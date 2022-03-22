package notify

import (
	"testing"
	"time"
	n "github.com/volcengine/volc-sdk-golang/service/notify"
)

const (
	testAk = ""
	testSk = ""
)

func TestCreateTask(t *testing.T) {
	n.DefaultInstance.Client.SetAccessKey(testAk)
	n.DefaultInstance.Client.SetSecretKey(testSk)

	phoneParam := &n.PhoneParam{Phone: "1990000000"}
	request := &n.CreateTaskRequest{Name: "kaede1",PhoneList: []*n.PhoneParam{phoneParam},Resource: "1c9c65e4ea50417b8dfe234963350218",
		StartTime: n.JsonTime(time.Now()), EndTime: n.JsonTime(time.Now().Add(time.Duration(1) * time.Hour)), Start: true, MaxRingDuration: 20, RingAgainInterval: 5,
		RingAgainTimes: 0, Unique: false, NumberPoolNo: "NP162160143010909069", SelectNumberType:0, Type: 0}

	response,statusCode,err := n.DefaultInstance.CreateTask(request)
	t.Logf("response = %+v\n", response)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQuerySingleInfo(t *testing.T)  {
	n.DefaultInstance.Client.SetAccessKey(testAk)
	n.DefaultInstance.Client.SetSecretKey(testSk)

	response,statusCode,err := n.DefaultInstance.SingleInfo("832a5c905c98480ba764928f3cd98c28")
	t.Logf("response = %+v\n", response)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}