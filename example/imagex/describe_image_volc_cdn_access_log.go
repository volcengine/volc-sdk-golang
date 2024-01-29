package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

// 查询离线日志下载地址
func main_DescribeImageVolcCdnAccessLog() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.DescribeImageVolcCdnAccessLog(context.Background(), &imagex.DescribeImageVolcCdnAccessLogReq{
		DescribeImageVolcCdnAccessLogQuery: &imagex.DescribeImageVolcCdnAccessLogQuery{
			ServiceID: "",
		},
		DescribeImageVolcCdnAccessLogBody: &imagex.DescribeImageVolcCdnAccessLogBody{
			StartTime: 1680500000,
			EndTime:   1680515000,
			Domain:    "domain",
			Region:    "global",
			PageNum:   1,
			PageSize:  10,
		},
	})
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

}
