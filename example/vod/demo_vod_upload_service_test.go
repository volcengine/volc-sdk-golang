// Code generated by protoc-gen-volcengine-sdk
// source: VodUploadService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_QueryUploadTaskInfo(t *testing.T) {
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

	query := &request.VodQueryUploadTaskInfoRequest{
		JobIds: "your JobIds",
	}

	resp, status, err := instance.QueryUploadTaskInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ApplyUploadInfo(t *testing.T) {
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

	query := &request.VodApplyUploadInfoRequest{
		SpaceName:         "your SpaceName",
		SessionKey:        "your SessionKey",
		FileSize:          0,
		FileType:          "your FileType",
		FileName:          "your FileName",
		StorageClass:      0,
		FileExtension:     "your FileExtension",
		ClientNetWorkMode: "your ClientNetWorkMode",
		ClientIDCMode:     "your ClientIDCMode",
		NeedFallback:      false,
		UploadHostPrefer:  "your UploadHostPrefer",
	}

	resp, status, err := instance.ApplyUploadInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CommitUploadInfo(t *testing.T) {
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

	query := &request.VodCommitUploadInfoRequest{
		SpaceName:       "your SpaceName",
		SessionKey:      "your SessionKey",
		CallbackArgs:    "your CallbackArgs",
		Functions:       "your Functions",
		VodUploadSource: "your VodUploadSource",
		ExpireTime:      "your ExpireTime",
	}

	resp, status, err := instance.CommitUploadInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
