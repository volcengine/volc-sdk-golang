package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 图像擦除
func main_GetImageEraseResult() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	// 获取图像擦除模型
	// 0 表示获取自动检测并擦除模型列表，1 表示获取指定区域擦除模型列表。默认 0
	models, err := instance.GetImageEraseModels(context.Background(), &imagex.GetImageEraseModelsQuery{
		Type: 0,
	})
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	fmt.Printf("success %+v", models)

	resp, err := instance.GetImageEraseResult(context.Background(), &imagex.GetImageEraseResultBody{
		ServiceID: "service id",                              // 服务 ID
		StoreURI:  "store uri",                               // 文件的 Store URI
		Model:     "",                                        // 使用的模型
		BBox:      []*imagex.GetImageEraseResultBodyBBoxItem{ // 处理指定区域
		},
	})
	if err != nil {
		fmt.Printf("error %v", err)
		return
	} else {
		fmt.Printf("success %+v", resp)
	}
}
