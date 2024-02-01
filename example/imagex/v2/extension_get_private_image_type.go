package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 图片隐私检测
func main_GetPrivateImageType() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetPrivateImageType(context.Background(), &imagex.GetPrivateImageTypeReq{
		GetPrivateImageTypeQuery: &imagex.GetPrivateImageTypeQuery{
			ServiceID: "service id", // 服务 ID,
		},
		GetPrivateImageTypeBody: &imagex.GetPrivateImageTypeBody{

			ImageURI:   "store uri", // 文件的 Store URI
			Method:     "all",       // 检测内容
			ThresFace:  0.52,        // 人脸置信度
			ThresCloth: 0.8,         // 衣物置信度
		},
	})
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

}
