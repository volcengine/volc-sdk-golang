package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 更新文件 URL 状态
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	serviceId := "imagex service id" // 服务 ID
	urls := []string{"image url 1"}  // 文件的 Store URI

	resp, err := instance.UpdateImageUrls(serviceId, &imagex.UpdateImageUrlPayload{
		Action:    imagex.ActionRefresh, // 刷新文件
		ImageUrls: urls,
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
