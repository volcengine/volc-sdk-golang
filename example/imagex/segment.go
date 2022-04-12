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

	param := &imagex.GetImageSegmentParam{
		ServiceId: "imageX service id",
		Class:     imagex.SEGMENT_CLASS_HUMAN_V2,
		Refine:    false,
		StoreUri:  "store uri",
		OutFormat: "out format",
		Contour: &imagex.Contour{
			Color: "#000000",
			Size:  0,
		},
		TransBg: true,
	}

	resp, err := instance.GetImageSegment(param)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
