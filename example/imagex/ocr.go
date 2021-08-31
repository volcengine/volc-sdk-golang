package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * get image ocr
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	param := &imagex.GetImageOCRParam{
		ServiceId: "xx",
		Scene:     "license",
		StoreUri:  "xx",
	}

	resp, err := instance.GetImageOCR(param)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
