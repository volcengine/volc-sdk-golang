package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 单资源 URL 数据迁移
func main_FetchImageURL() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	body := &imagex.FetchImageURLBody{
		URL: "https://example.com/example.jpg", // 待迁移的文件

		ServiceID: "service id", // 想将文件迁移到哪一个服务
		StoreKey:  "store key",  // 迁移后的存储名为
		// Async: true,
	}
	resp, err := instance.FetchImageURL(context.Background(), body)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %v\n", resp)
	}

	if body.Async {
		req2 := &imagex.GetURLFetchTaskQuery{
			ID:        resp.Result.TaskID,
			ServiceID: body.ServiceID,
		}
		resp, err := instance.GetURLFetchTask(context.Background(), req2)
		if err != nil {
			fmt.Printf("error %v\n", err)
		} else {
			fmt.Printf("success %v\n", resp)
		}
	}
}
