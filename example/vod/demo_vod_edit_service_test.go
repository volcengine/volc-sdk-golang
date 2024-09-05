// Code generated by protoc-gen-volcengine-sdk
// source: VodEditService
// DO NOT EDIT!

package vod

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

func Test_SubmitDirectEditTaskAsync(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

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
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

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
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetDirectEditResultRequest{
		ReqIds: []string{"your ReqId here"},
	}

	resp, status, err := instance.GetDirectEditResult(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetDirectEditProgress(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

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
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCancelDirectEditTaskRequest{
		ReqId: "your ReqId here",
	}
	resp, status, err := instance.CancelDirectEditTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
