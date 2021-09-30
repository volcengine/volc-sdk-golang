package adblocker

import (
	"fmt"
	"testing"
)

const (
	Ak = "AK" // write your access key
	Sk = "SK" // write your secret key
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
