package businessSecurity

import (
	"fmt"
	"testing"
)

const (
	Ak = "***REMOVED***"              // write your access key
	Sk = "***REMOVED***" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func RiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.RiskDetection(&RiskDetectionRequest{
		AppID:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func AsyncRiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.AsyncRiskDetection(&AsyncRiskDetectionRequest{
		AppID:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func RiskResult(appId int64, service string, startTime, endTime, pageSize, pageNum int64) {
	res, err := DefaultInstance.RiskResult(&RiskResultRequest{
		AppId:     appId,   // write your app id
		Service:   service, // write business security service
		StartTime: startTime,
		EndTime:   endTime,
		Page: Page{
			PageSize: pageSize,
			PageNum:  pageNum,
		},
	})
	fmt.Println(err)
	if res != nil {
		fmt.Printf("result %+v\n",*res)
	}
}

func TestBusinessSecurity_RiskResult(t *testing.T) {
	RiskResult(3332, "register", 1615535000, 1615540603, 10, 1)
}