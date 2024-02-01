package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 图片智能剪裁
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImageSmartCropParam{
		ServiceId: "service id",                            // 服务 ID
		StoreUri:  "store uri",                             // 文件的 Store URI
		Policy:    imagex.SMARTCROP_POLICY_DEMOTION_FGLASS, // 高斯模糊模式
		Scene:     imagex.SMARTCROP_SCENE_NORMAL,           // 普通人脸剪裁
		Sigma:     5.0,                                     // 高斯模糊半径
		Width:     480,                                     // 裁剪后的宽度
		Height:    320,                                     // 裁剪后的高度
	}

	resp, err := instance.GetImageSmartCrop(param)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}
}
