// Code generated by protoc-gen-volcengine-sdk
// source: VodSmartStrategyLiteService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_GetSmartStrategyLitePlayInfo(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
