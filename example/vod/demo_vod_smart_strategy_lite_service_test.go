// Code generated by protoc-gen-volcengine-sdk
// source: VodSmartStrategyLiteService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_GetSmartStrategyLitePlayInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetSmartStrategyLitePlayInfoRequest{
		PlayUrl:         "your PlayUrl",
		Format:          "your Format",
		Codec:           "your Codec",
		Definition:      "your Definition",
		FileType:        "your FileType",
		LogoType:        "your LogoType",
		Ssl:             "your Ssl",
		NeedThumbs:      "your NeedThumbs",
		NeedBarrageMask: "your NeedBarrageMask",
		UnionInfo:       "your UnionInfo",
		HDRDefinition:   "your HDRDefinition",
	}

	resp, status, err := instance.GetSmartStrategyLitePlayInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
