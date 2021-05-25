package gameProtect

import (
	"fmt"
	"testing"
)

const (
	Ak = "ak" // write your access key
	Sk = "sk" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func RiskResult(appId int64, startTime, endTime, pageSize, pageNum int64) {
	res, err := DefaultInstance.RiskResult(&RiskResultRequest{
		AppId:     appId,   // write your app id
		StartTime: startTime,
		EndTime:   endTime,
		Page: Page {
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
	RiskResult(218745, 1618502400, 1618545491, 2, 1)
}