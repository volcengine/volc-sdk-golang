package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 图片降噪
func main_GetDenoisingImage() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetDenoisingImage(context.Background(), &imagex.GetDenoisingImageReq{
		GetDenoisingImageQuery: &imagex.GetDenoisingImageQuery{
			ServiceID: "service id",
		},
		GetDenoisingImageBody: &imagex.GetDenoisingImageBody{
			CanDemotion: true,
			Intensity:   0.1,
			OutFormat:   "png",
		},
	})
	if err != nil {
		fmt.Printf("GetDenoisingImage error %v\n", err)
	} else {
		fmt.Printf("GetDenoisingImage success %+v\n", resp)
	}

}
