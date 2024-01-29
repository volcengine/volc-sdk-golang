package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 检测图像是否经过处理
func main_GetImagePSDetection() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImagePSDetection(context.Background(), &imagex.GetImagePSDetectionReq{
		GetImagePSDetectionQuery: &imagex.GetImagePSDetectionQuery{
			ServiceID: "service id",
		},
		GetImagePSDetectionBody: &imagex.GetImagePSDetectionBody{
			ImageURI: "store uri",
			Method:   "he",
		},
	})
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

}
