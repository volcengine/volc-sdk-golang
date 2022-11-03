package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 图像擦除
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	// 获取图像擦除模型
	// 0 表示获取自动检测并擦除模型列表，1 表示获取指定区域擦除模型列表。默认 0
	models, err := instance.GetImageEraseModel(1)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	fmt.Printf("success %+v", models)

	// 图像擦除
	param := &imagex.GetImageEraseParam{
		ServiceId: "service id", // 服务 ID
		StoreUri:  "store uri",  // 文件的 Store URI
		Model:     models[0],    // 使用的模型
		BBox: []imagex.EraseBox{ // 处理指定区域
			{
				X1: 0,
				Y1: 0,
				X2: 1,
				Y2: 1,
			},
		},
	}
	resp, err := instance.GetImageErase(param)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	} else {
		fmt.Printf("success %+v", resp)
	}
}
