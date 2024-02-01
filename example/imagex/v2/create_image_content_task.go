package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 刷新/预热/禁用/解禁
// 创建刷新/预热/禁用/解禁任务
func main_CreateImageContentTask() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "aK",
		SecretAccessKey: "sk",
	})

	resp, err := instance.CreateImageContentTask(context.Background(), &imagex.CreateImageContentTaskReq{
		CreateImageContentTaskQuery: &imagex.CreateImageContentTaskQuery{
			ServiceID: "",
		},
		CreateImageContentTaskBody: &imagex.CreateImageContentTaskBody{
			Urls: []string{},
		},
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
