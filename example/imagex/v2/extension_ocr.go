package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 光学字符识别
func main_GetImageOCRV2() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImageOCRV2(context.Background(), &imagex.GetImageOCRV2Req{
		GetImageOCRV2Query: &imagex.GetImageOCRV2Query{
			ServiceID: "service id", // 服务 ID,
		},
		GetImageOCRV2Body: &imagex.GetImageOCRV2Body{
			Scene:    "general",   // 识别场景
			StoreURI: "store uri", // 文件的 Store URI
		},
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
