package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * image superresolution
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	param := &imagex.GetImageSmartCropParam{
		ServiceId: "",
		StoreUri:  "",
		Scene:     imagex.SMARTCROP_SCENE_NORMAL,
		Policy:    imagex.SMARTCROP_POLICY_DEMOTION_CENTER,
	}
	resp, err := instance.GetImageSmartCrop(param)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
