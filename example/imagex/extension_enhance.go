package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 画质增强
func main_GetImageEnhanceResult() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImageEnhanceResult(context.Background(), &imagex.GetImageEnhanceResultBody{
		ServiceID:    "service id", // 服务 ID
		StoreURI:     "store uri",  // 文件的 Store URI
		Model:        0,            // 0 通用模型
		DisableAr:    false,        // 是否不优化失真
		DisableSharp: false,        // 是否不自动锐化
		//Resolution:   "720p",       // 分辨率
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
