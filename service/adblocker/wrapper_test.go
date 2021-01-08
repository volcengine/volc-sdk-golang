package adblocker

import (
	"fmt"
	"testing"
)

const (
	region = "region"
	testAk = "ak"
	testSk = "sk"
)

func TestAdBlock(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	res, err := DefaultInstance.AdBlock(&AdBlockRequest{
		AppID:      3332,
		Service:    "chat",
		Parameters: "{\"uid\":123411, \"operate_time\":1609818934, \"chat_text\":\"a😊\"}",
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}
