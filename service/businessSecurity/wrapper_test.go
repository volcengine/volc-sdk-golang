package businessSecurity

const (
	Ak = "AK" // write your access key
	Sk = "SK" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func TestQueryNameListItem(t *testing.T) {
	params := new(QuerySystemNameListItemReq)
	params.Product = "adblocker"
	params.AppId = 428354
	params.Service = "chat"
	params.Object = "J3DLte11st09f25002"

	resp, err := DefaultInstance.QuerySystemNameListItem(params)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if resp.Code != 0 {
		t.Errorf("error %s log id %s", resp.Message, resp.RequestId)
		return
	}
	t.Logf("detail %+v", *resp.Data)
}

func TestDeleteNameListItem(t *testing.T) {
	params := new(DelSystemNameListItemReq)
	params.Product = "adblocker"
	params.AppId = 428354
	params.Service = "chat"
	params.Object = "J3DLte11st09f25002"

	resp, err := DefaultInstance.DelSystemNameListItem(params)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if resp.Code != 0 {
		t.Errorf("error %s log id %s", resp.Message, resp.RequestId)
		return
	}
	t.Logf("detail %+v", *resp.Data)
}
