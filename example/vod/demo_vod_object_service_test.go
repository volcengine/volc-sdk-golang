package vod

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"testing"
)

func TestSubmitMoveObjectTask(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "",
		SecretAccessKey: "",
	})

	query := &request.VodSubmitMoveObjectTaskRequest{
		SourceSpace:      "sourceSpace",
		SourceFileName:   "sourceFileName",
		TargetSpace:      "targetSpace",
		TargetFileName:   "",
		SaveSourceObject: false,
		ForceOverwrite:   false,
	}

	resp, status, err := instance.SubmitMoveObjectTask(query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(resp.String())

}

func Test_QueryMoveObjectTaskInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "",
		SecretAccessKey: "",
	})

	query := &request.VodQueryMoveObjectTaskInfoRequest{
		TaskId:      "dd76ff82ac254f0a8ac7b5a20def3bbe",
		SourceSpace: "sourceSpace",
		TargetSpace: "targetSpace",
	}

	resp, status, err := instance.QueryMoveObjectTaskInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func TestMoveObject(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "",
		SecretAccessKey: "",
	})

	query := &request.VodSubmitMoveObjectTaskRequest{
		SourceSpace:      "sourceSpace",
		SourceFileName:   "sourceFileName",
		TargetSpace:      "targetSpace",
		TargetFileName:   "",
		SaveSourceObject: false,
		ForceOverwrite:   false,
	}

	resp, status, err := instance.MoveObjectCrossSpace(query, 0)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
