package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 重复图片检测（异步）
func main_GetImageDuplicateDetection() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	// 创建检测任务
	resp, err := instance.GetImageDuplicateDetection(context.Background(), &imagex.GetImageDuplicateDetectionReq{
		GetImageDuplicateDetectionQuery: &imagex.GetImageDuplicateDetectionQuery{
			ServiceID: "service id", // 服务 ID
		},
		GetImageDuplicateDetectionBody: &imagex.GetImageDuplicateDetectionBody{
			Urls:       []string{"url1", "url2"}, // 待检测的图片地址
			Async:      true,                     // 是否使用异步
			Callback:   "",                       // 回调地址
			Similarity: 0.84,                     // 相似度阈值
		},
	})
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
	resp1, err := instance.GetDedupTaskStatus(context.Background(), &imagex.GetDedupTaskStatusQuery{
		TaskID: resp.Result.TaskID, // 任务 ID
	})
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp1)
	}
}
