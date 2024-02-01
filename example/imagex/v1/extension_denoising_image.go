package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 图片降噪
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetDenoisingImageParam{
		ServiceId: "service id", // 服务 ID
		StoreUri:  "store uri",  // 文件的 Store URI
		// ImageUrl:    "image uri",  // 文件的网址
		OutFormat:   "png", // 输出的文件格式
		Intensity:   0.1,   // 降噪强度
		CanDemotion: true,  // 是否允许降级
	}

	resp, err := instance.GetDenoisingImage(param)
	if err != nil {
		fmt.Printf("GetImageQuality error %v\n", err)
	} else {
		fmt.Printf("GetImageQuality success %+v\n", resp)
	}

}
