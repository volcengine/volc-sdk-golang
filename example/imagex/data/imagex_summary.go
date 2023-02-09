package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
	"github.com/volcengine/volc-sdk-golang/service/imagex/data"
)

func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance
	_ = data.AddDataModule(instance)

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	req := &data.DescribeImageXSummaryReq{
		ServiceIds: "service1,service2",
		Timestamp:  "2023-02-08T00:00:00+08:00",
	}

	resp, err := data.DescribeImageXSummary(instance, req)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
