package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 上传文件
func main_SegmentedUploadImages() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	params := &imagex.ApplyUploadImageParam{
		ServiceId: "service id", // 服务 ID
		// StoreKeys: []string{"example.jpg"}, // 指定文件存储名
	}

	files := make([]io.Reader, 0)
	sizeArr := make([]int64, 0)
	addFile(params, &files, &sizeArr, "store_key1", "/path/to/file1")
	addFile(params, &files, &sizeArr, "store_key2", "/path/to/file2")

	startTime := time.Now()
	// 上传文件
	resp, err := instance.SegmentedUploadImages(context.Background(), params, files, sizeArr)
	takes := time.Since(startTime)

	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		str, _ := json.Marshal(resp)
		fmt.Printf("success %s\n", string(str))
	}

	fmt.Printf("%v\n", takes)
}

func addFile(params *imagex.ApplyUploadImageParam, files *[]io.Reader, sizeArr *[]int64, storeKey string, filePath string) {
	// 读取文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	params.StoreKeys = append(params.StoreKeys, storeKey)
	*files = append(*files, file)
	*sizeArr = append(*sizeArr, fileInfo.Size())
}
