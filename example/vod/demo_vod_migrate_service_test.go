// Code generated by protoc-gen-volcengine-sdk
// source: VodMigrateService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_SetCloudMigrateJob(t *testing.T) {
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

	query := &request.VodSetCloudMigrateJobRequest{
		JobId:           0,
		JobSourceInfo:   nil,
		CallbackAddress: "your CallbackAddress",
		SpaceName:       "your SpaceName",
	}

	resp, status, err := instance.SetCloudMigrateJob(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_SubmitCloudMigrateJob(t *testing.T) {
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

	query := &request.VodSubmitCloudMigrateJobRequest{
		JobId:        0,
		SubAppId:     "your SubAppId",
		SourceVid:    "your SourceVid",
		TargetVid:    "your TargetVid",
		CallbackArgs: "your CallbackArgs",
		SpaceName:    "your SpaceName",
	}

	resp, status, err := instance.SubmitCloudMigrateJob(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetCloudMigrateJob(t *testing.T) {
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

	query := &request.VodGetCloudMigrateJobRequest{
		JobId:     0,
		SpaceName: "your SpaceName",
	}

	resp, status, err := instance.GetCloudMigrateJob(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}