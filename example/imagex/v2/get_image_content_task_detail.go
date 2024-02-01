package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 刷新/预热/禁用/解禁
// 获取任务详情
func main_GetImageContentTaskDetail() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImageContentTaskDetail(context.Background(), &imagex.GetImageContentTaskDetailBody{
		TaskType:  "refresh_url",
		StartTime: 0,
		EndTime:   2147483647,
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
