package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 删除文件
func main_DeleteImageUploadFiles() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.DeleteImageUploadFiles(context.Background(), &imagex.DeleteImageUploadFilesReq{
		DeleteImageUploadFilesQuery: &imagex.DeleteImageUploadFilesQuery{
			ServiceID: "",
		},
		DeleteImageUploadFilesBody: &imagex.DeleteImageUploadFilesBody{
			StoreUris: []string{"image uri 1"},
		},
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
