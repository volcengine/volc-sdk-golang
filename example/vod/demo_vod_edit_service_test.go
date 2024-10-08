// Code generated by protoc-gen-volcengine-sdk
// source: VodEditService
// DO NOT EDIT!

package vod

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

func Test_SubmitDirectEditTaskAsync(t *testing.T) {
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

	params, _ := json.Marshal(map[string]interface{}{})
	query := &request.VodSubmitDirectEditTaskAsyncRequest{
		Uploader:     "your target store vod space",
		Application:  "your application",
		EditParam:    params,
		Priority:     0,
		CallbackUri:  "your CallbackUri if you need",
		CallbackArgs: "your CallbackArgs if you need",
	}
	resp, status, err := instance.SubmitDirectEditTaskAsync(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_SubmitDirectEditTaskSync(t *testing.T) {
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

	params, _ := json.Marshal(map[string]interface{}{})
	query := &request.VodSubmitDirectEditTaskSyncRequest{
		Uploader:    "your target store vod space",
		Application: "your application",
		EditParam:   params,
	}
	resp, status, err := instance.SubmitDirectEditTaskSync(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetDirectEditResult(t *testing.T) {
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

	query := &request.VodGetDirectEditResultRequest{
		ReqIds: []string{"your ReqId here"},
	}

	resp, status, err := instance.GetDirectEditResult(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetDirectEditProgress(t *testing.T) {
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

	query := &request.VodGetDirectEditProgressRequest{
		ReqId: "your ReqId here",
	}

	resp, status, err := instance.GetDirectEditProgress(query)
	fmt.Println(status)
	fmt.Println(err)
	op := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	format := op.Format(resp)
	fmt.Println(format)
}

func Test_CancelDirectEditTask(t *testing.T) {
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

	query := &request.VodCancelDirectEditTaskRequest{
		ReqId: "your ReqId here",
	}
	resp, status, err := instance.CancelDirectEditTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
