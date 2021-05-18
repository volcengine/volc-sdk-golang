package gameProtect

import (
	"fmt"
	"testing"
)

const (
	Ak = "***REMOVED***" // write your access key
	Sk = "***REMOVED***" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func RiskResult(appId int64, service string, startTime, endTime, pageSize, pageNum int64) {
	res, err := DefaultInstance.RiskResult(&RiskResultRequest{
		AppId:     appId,   // write your app id
		Service:   service, // write business security service
		StartTime: startTime,
		EndTime:   endTime,
		Page: Page{
			PageNum:  pageNum,
			PageSize: pageSize,
		},
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestGameProtector_RiskResult(t *testing.T) {
	RiskResult(218745, "anti_plugin",1618502400, 1618545491, 2, 1)
}