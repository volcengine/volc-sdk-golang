package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * image erase
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// type: 0 for auto detect and erase, 1 for box erase
	models, err := instance.GetImageEraseModel(1)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	fmt.Printf("success %+v", models)

	param := &imagex.GetImageEraseParam{
		ServiceId: "imagex service id",
		StoreUri:  "image uri",
		Model:     models[0],
	}
	resp, err := instance.GetImageErase(param)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	} else {
		fmt.Printf("success %+v", resp)
	}
}
