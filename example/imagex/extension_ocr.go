package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 光学字符识别
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImageOCRParam{
		ServiceId: "service id", // 服务 ID
		StoreUri:  "store uri",  // 文件的 Store URI
		ImageUrl:  "",           //若StoreUri为空则访问url
	}

	// resp, err := instance.GetImageOCRGeneral(param)
	resp, err := instance.GetImageOCRLicense(param)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
