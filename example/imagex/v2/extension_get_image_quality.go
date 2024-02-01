package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 画质评估
func main_GetImageQuality() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImageQuality(context.Background(), &imagex.GetImageQualityReq{
		GetImageQualityQuery: &imagex.GetImageQualityQuery{
			ServiceID: "service id", // 服务 ID,
		},
		GetImageQualityBody: &imagex.GetImageQualityBody{
			VqType: "", // 评估工具，使用逗号分隔的评估工具列表
		},
	})
	if err != nil {
		fmt.Printf("GetImageQuality error %v\n", err)
	} else {
		fmt.Printf("GetImageQuality success %+v\n", resp)
	}
}
