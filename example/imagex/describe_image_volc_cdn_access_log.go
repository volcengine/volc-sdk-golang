package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 提取盲水印
func main() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.DescribeImageVolcCdnAccessLogReq{
		ServiceId: "service id",
		StartTime: 1680500000,
		EndTime:   1680515000,
		Domain:    "domain",
		Region:    "global",
		PageNum:   1,
		PageSize:  10,
	}

	resp, err := instance.DescribeImageVolcCdnAccessLog(param)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

}
