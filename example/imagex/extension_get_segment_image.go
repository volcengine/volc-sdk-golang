package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 智能移除背景
func main_GetSegmentImage() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetSegmentImage(context.Background(), &imagex.GetSegmentImageReq{
		GetSegmentImageQuery: &imagex.GetSegmentImageQuery{
			ServiceID: "service id",
		},
		GetSegmentImageBody: &imagex.GetSegmentImageBody{
			StoreURI:  "store uri",
			Refine:    false,
			OutFormat: "out format",
			Contour: &imagex.GetSegmentImageBodyContour{
				Color: "#000000",
				Size:  0,
			},
			TransBg: true,
		},
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
