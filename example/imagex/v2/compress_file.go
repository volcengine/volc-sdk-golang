package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// ZIP压缩打包文件
func main_CreateImageCompressTask() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "xx",
		SecretAccessKey: "xx",
	})

	resp, err := instance.CreateImageCompressTask(context.Background(), &imagex.CreateImageCompressTaskReq{
		CreateImageCompressTaskQuery: &imagex.CreateImageCompressTaskQuery{
			ServiceID: "xx",
		},
		CreateImageCompressTaskBody: &imagex.CreateImageCompressTaskBody{
			FileList: []*imagex.CreateImageCompressTaskBodyFileListItem{
				{
					URL:    "xx", // 需要处理的 URL
					Alias:  "",
					Folder: "",
				},
			},
			IndexFile: "",
			Output:    "",
			Callback:  "",
			ZipMode:   0,
		},
	})

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
	/*
		tresp, err := instance.GetImageCompressTaskInfo(&imagex.GetImageCompressTaskInfoReq{
			ServiceId: "xx",
			TaskId:    "xx",
		})
		if err != nil {
			fmt.Printf("error %v", err)
		} else {
			fmt.Printf("success %v", tresp)
		}
	*/
}
