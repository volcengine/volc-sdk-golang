package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 智能移除背景
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImageSegmentParam{
		ServiceId: "service id",                  // 服务 ID
		StoreUri:  "store uri",                   // 文件的 Store URI
		Class:     imagex.SEGMENT_CLASS_HUMAN_V2, // 模型
		Refine:    false,                         // 处理效果
		OutFormat: "out format",                  // 输出格式
		Contour: &imagex.Contour{ // 描边
			Color: "#000000",
			Size:  0,
		},
		TransBg: true, // 透明背景
	}

	resp, err := instance.GetImageSegment(param)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
