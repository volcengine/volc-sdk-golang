package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 刷新/预热/禁用/解禁
// 创建刷新/预热/禁用/解禁任务
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "aK",
		SecretAccessKey: "sk",
	})

	resp, err := instance.CreateImageContentTask(&imagex.CreateImageContentTaskReq{
		ServiceId: "",
		TaskType:  "block_url",
		Urls:      []string{}, // 需要处理的 URL 列表（可以通过浏览器访问的完整 URL）
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
