package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 添加盲水印
func main_CreateImageHmEmbed() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "aK",
		SecretAccessKey: "sk",
	})

	resp, err := instance.CreateImageHmEmbed(context.Background(), &imagex.CreateImageHmEmbedBody{
		ServiceID:     "service id", // 服务 ID
		StoreURI:      "store uri",  // 文件的 Store URI
		Algorithm:     "default",    // 算法模型
		Info:          "info",       // 盲水印的內容
		OutFormat:     "jpeg",       // 输出图片格式
		OutQuality:    90,           // 输出图片质量
		StrengthLevel: "medium",     // 算法强度
	})
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

}
