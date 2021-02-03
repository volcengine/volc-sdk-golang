package businessSecurity

import (
	"fmt"
	"testing"
	"time"
)

const (
	Ak = "ak" // write your access key
	Sk = "sk" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func TestBusinessSecurity_RiskDetection(t *testing.T) {
	for i := 0; i < 400; i++ {
		go RiskDetection(3332, "register", "{\"operate_time\":1612269789}")
	}

	time.Sleep(10 * time.Second)
}

func RiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.RiskDetection(&RiskDetectionRequest{
		AppID:      appId,      // write your app id
		Service:    service,    // write adblocker service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}
