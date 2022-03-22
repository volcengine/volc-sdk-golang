package notify

import "testing"

func TestQuerySingleInfo(t *testing.T)  {
	DefaultInstance.Client.SetAccessKey("")
	DefaultInstance.Client.SetSecretKey("")
	DefaultInstance.Client.ServiceInfo.Host = "volcengineapi-boe.byted.org"
	response,statusCode,err := DefaultInstance.SingleInfo("832a5c905c98480ba764928f3cd98c28")
	t.Logf("response = %+v\n", response)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}