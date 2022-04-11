package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * image segment
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	param := &imagex.GetImageSegmentParamV2{
		GetImageSegmentParam: imagex.GetImageSegmentParam{
			ServiceId: "imagex service id",
			Class:     "process class",
			Refine:    false,
			StoreUri:  "image uri",
			OutFormat: "out format",
		},
		Contour: &imagex.Contour{
			Color: "contour color",
			Size:  10,
		},
		TransBg: true,
	}

	resp, err := instance.GetImageSegmentV2(param)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
