package adblocker

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	Ak = "" // write your access key
	Sk = "" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
	// DefaultInstance.CloseRetry()   // call this if you don't want retry on error
}

func AdBlock(appId int64, service string, parameters string) {
	res, err := DefaultInstance.AdBlock(&AdBlockRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write adblocker service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func Test_AdBlock(t *testing.T) {
	AdBlock(5461, "chat", "{\"operate_time\":1630908329, \"chat_text\":\"测试文本测试文本加我微信\"}")
}

func TestSimpleRiskStat(t *testing.T) {
	params := new(SimpleProductStatisticsParams)
	params.StartDate = "2023-05-07"
	params.EndDate = "2023-05-09"
	params.NeedServiceDetail = true
	params.NeedAppDetail = true
	params.OperateTime = time.Now().Unix()
	paramsStr, err := json.Marshal(params)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	req := new(CommonProductStatisticsReq)
	req.Product = PrdAdblocker
	req.UnitType = DAILY
	req.Parameters = string(paramsStr)

	result, err := DefaultInstance.SimpleRiskStat(req)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("total: %+v", result.Total)
	for idx, detail := range result.Detail {
		t.Logf("detail %d: %+v", idx, detail)
	}
}
