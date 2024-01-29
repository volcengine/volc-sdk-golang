package main

import (
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 通用 API 调用例程
// 您可以在 https://www.volcengine.cn/docs/508/14106 查到 API 定义
func main_common() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	// 以 GetImageUploadFile 为例
	// 这个 API 需要一个 GET 请求，查询参数需要包含 ServiceId 和 StoreUri

	query := url.Values{}
	query.Add("ServiceId", "service id") // 服务 ID
	query.Add("StoreUri", "image uri")   // 文件的 Store URI

	// 这个 API 的返回值定义
	resp := new(struct {
		ServiceId    string `json:"ServiceId"`
		StoreUri     string `json:"StoreUri"`
		LastModified string `json:"LastModified"`
		FileSize     int    `json:"FileSize"`
		Marker       int64  `json:"Marker"`
	})

	// 使用 ImageX 客户端实例发起 Get 请求
	err := instance.ImageXGet("GetImageUploadFile", query, resp)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
