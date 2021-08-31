package adblocker

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"
	"time"
)

const (
	Ak = "AK" // write your access key
	Sk = "SK" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
	err := DefaultInstance.SetRegion(base.RegionApSingapore)
	DefaultInstance.Client.SetTimeout(500 * time.Millisecond)
	if err != nil {
		panic(err)
	}
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
	AdBlock(216455, "chat", "{}")
}
