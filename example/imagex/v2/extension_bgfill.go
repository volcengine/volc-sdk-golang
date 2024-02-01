package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 智能图像扩展
func main_GetImageBgFillResult() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImageBgFillResult(context.Background(), &imagex.GetImageBgFillResultBody{
		ServiceID: "imagex service id", // 服务 ID
		StoreURI:  "image store uri",   // 文件的 Store URI
		Model:     1,                   // 0 表示漫画模式, 1 表示一般模式
		Top:       0.2,                 // 向上延伸 20%
		Bottom:    0.2,                 // 向下延伸 20%
		Left:      0.2,                 // 向左延伸 20%
		Right:     0.2,                 // 向右延伸 20%
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
