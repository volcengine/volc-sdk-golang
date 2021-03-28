package adblocker

import (
	"fmt"
)

const (
	Ak = "ak" // write your access key
	Sk = "sk" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
	// DefaultInstance.CloseRetry()   // call this if you don't want retry on error
}

func AdBlock(appId int64, service string, parameters string) {
	res, err := DefaultInstance.AdBlock(&AdBlockRequest{
		AppID:      appId,      // write your app id
		Service:    service,    // write adblocker service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}
