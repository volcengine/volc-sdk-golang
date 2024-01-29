package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 图片智能剪裁
func main_GetImageSmartCropResult() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImageSmartCropResult(context.Background(), &imagex.GetImageSmartCropResultBody{
		ServiceID: "service id", // 服务 ID,
		StoreURI:  "store uri",  // 文件的 Store URI,
		Policy:    "",
		Width:     10,
		Height:    10,
	})
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}
}
