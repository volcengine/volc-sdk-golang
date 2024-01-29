package main

import (
	"fmt"
	"os"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 上传文件
func main_UploadImages() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	params := &imagex.ApplyUploadImageParam{
		ServiceId: "service id", // 服务 ID
		// StoreKeys: []string{"example.jpg"}, // 指定文件存储名
		Overwrite: false, //是否进行上传覆盖
	}

	// 读取文件
	dat, err := os.ReadFile("image file")
	if err != nil {
		fmt.Printf("read file from %s error %v", "", err)
		os.Exit(-1)
	}

	// 上传文件
	resp, err := instance.UploadImages(params, [][]byte{dat})

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
