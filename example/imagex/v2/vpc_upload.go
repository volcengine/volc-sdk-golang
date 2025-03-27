package main

import (
	"context"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 上传文件
func main_VpcUploadImages() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})
	//instance.SetHost("open.volcengineapi.com")	// VPC内网调用openapi

	params := &imagex.VpcUploadRequest{
		ServiceId:     "service id",          // 服务 ID
		FilePath:      "your file path",      // 文件路径，与Data二选一
		Data:          nil,                   // 文件数据，与FilePath二选一
		StoreKey:      "your store key",      // 文件存储名
		Prefix:        "your prefix",         // 文件前缀
		FileExtension: "your file extension", // 文件后缀
		ContentType:   "your content type",   // 文件Content-Type
		StorageClass:  "your storage class",  // 文件存储类型
		PartSize:      0,                     // 偏好分片大小，单位为字节
		Overwrite:     false,                 //是否进行上传覆盖
		SkipMeta:      false,                 // 是否跳过元信息
	}

	// 上传文件
	resp, err := instance.VpcUploadImage(context.Background(), params)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
