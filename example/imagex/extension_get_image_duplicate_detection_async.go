package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 重复图片检测（异步）
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImageDuplicateDetectionAsyncParam{
		Callback: "https://example.com/callback/", // 回调地址
	}
	param.ServiceId = "service id"        // 服务 ID
	param.Urls = []string{"url1", "url2"} // 待检测的图片地址

	// 创建检测任务
	resp, err := instance.GetImageDuplicateDetectionAsync(param)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

	// // 你可以使用回调来获得结果
	// // 使用 Gin 解析回调信息的一个例子
	//  ginEngine.POST("/ping", func(c *gin.Context) {
	//      var data imagex.GetImageDuplicateDetectionResult
	//      c.BindJSON(&data)
	//      // 处理数据
	//  })

	// 也可以主动获取结果
	// 获取检测任务状态
	resp1, err := instance.GetDedupTaskStatus(resp.TaskId)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp1)
	}
}
