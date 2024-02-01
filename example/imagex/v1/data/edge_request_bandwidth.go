package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	req := &imagex.DescribeImageXEdgeRequestBandwidthReq{
		StartTime: "2023-01-21T00:00:00+08:00",
		EndTime:   "2023-01-28T00:00:00+08:00",
		Interval:  "300",
	}

	resp, err := instance.DescribeImageXEdgeRequestBandwidth(req)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
