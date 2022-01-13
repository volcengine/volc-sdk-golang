package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.NewInstance()

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "AK",
		SecretAccessKey: "sk",
	})

	//image uri
	param := &imagex.GetImageQualityParam{
		ServiceId: "sid",
		ImageUrl:  "image uri",
	}

	resp, err := instance.GetImageQuality(param)
	if err != nil {
		fmt.Printf("GetImageQuality error %v\n", err)
	} else {
		fmt.Printf("GetImageQuality success %+v\n", resp)
	}

}
