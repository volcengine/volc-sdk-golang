package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 创意魔方
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	req := &imagex.GetImageStyleResultReq{
		ServiceId: "imagex service id", // 服务 ID
		StyleId:   "style id",          // 样式编号
		// 如果您选择的样式支持参数，可以在 Params 中填写对应的参数
		Params: map[string]string{
			"参数名": "参数值", // 样式参数
		},
		OutputFormat:  "jpeg", // 输出图片格式
		OutputQuality: 90,     // 输出图片质量
	}
	resp, err := instance.GetImageStyleResult(req)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
