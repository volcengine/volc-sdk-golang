package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 画质评估
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImageQualityParam{
		ServiceId: "service id", // 服务 ID
		ImageUrl:  "store uri",  // 文件的 Store URI
		VqType:    "vqscore",    // 评估工具，使用逗号分隔的评估工具列表
	}

	resp, err := instance.GetImageQuality(param)
	if err != nil {
		fmt.Printf("GetImageQuality error %v\n", err)
	} else {
		fmt.Printf("GetImageQuality success %+v\n", resp)
	}
}
