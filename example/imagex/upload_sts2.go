package main

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 获取 STS2 的上传密钥（离线的）
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	serviceIds := []string{"service id"} // 欲授权的 Service ID

	// 默认超时时间为 1小时，如果有需要，请调用 imagex.GetUploadAuthWithExpire() 来设置超时时间
	// 您可以使用 imagex.WithUploadKeyPtn("表达式") 来限制上传的存储名格式
	//     如: "test/*" 表示上传的文件必须包含 "test/" 前缀
	//使用imagex.WithUploadOverwrite() 来设置sts中「上传覆盖」选项
	token, err := instance.GetUploadAuth(serviceIds, imagex.WithUploadOverwrite(false))
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("token %+v", token)
	}
}
