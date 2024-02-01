package main

import (
	"context"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

// 创意魔方
func main_GetImageStyleResult() {
	// 默认 ImageX 实例为 `cn-north-1`，如果您想使用其他区域的实例，请使用 `imagex.NewInstanceWithRegion(区域名)` 显式指定区域
	instance := imagex.DefaultInstance

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	resp, err := instance.GetImageStyleResult(context.Background(), &imagex.GetImageStyleResultReq{
		GetImageStyleResultQuery: &imagex.GetImageStyleResultQuery{
			ServiceID: "",
		},
		GetImageStyleResultBody: &imagex.GetImageStyleResultBody{
			StyleID:       "",
			Params:        map[string]string{},
			OutputFormat:  "",
			OutputQuality: 0,
		},
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
