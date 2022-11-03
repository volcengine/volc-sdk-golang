package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 检测图像是否经过处理
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImagePSDetectionParam{
		ServiceId: "service id", // 服务 ID
		ImageUri:  "store uri",  // 文件的 Store URI
		Method:    "he",         // 检查方法
	}

	resp, err := instance.GetImagePSDetection(param)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

}
