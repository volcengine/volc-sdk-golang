package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 图片隐私检测
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetPrivateImageTypeParam{
		ServiceId:  "service id", // 服务 ID
		ImageUri:   "store uri",  // 文件的 Store URI
		Method:     "all",        // 检测内容
		ThresFace:  0.52,         // 人脸置信度
		ThresCloth: 0.8,          // 衣物置信度
	}

	resp, err := instance.GetPrivateImageType(param)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

}
