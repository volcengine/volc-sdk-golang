package main

import (
	"fmt"
	"os"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

const (
	ServiceID = "6jdjx83cbx"
)

// 画质增强
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImageEnhanceParam{
		ServiceId:    "service id", // 服务 ID
		Model:        0,            // 0 通用模型
		DisableAr:    false,        // 是否不优化失真
		DisableSharp: false,        // 是否不自动锐化
		//Resolution:   "720p",       // 分辨率
		//Actions:      []string{},   // 自适应处理列表。
	}
	file, err := os.Open("/path/to/file1")
	if err != nil {
		panic(err)
	}

	resp, err := instance.GetImageEnhanceWithData(param, file)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
