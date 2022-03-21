// Code generated by protoc-gen-volcengine-sdk
// source: VodUploadService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_UploadMediaByUrl(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUrlUploadRequest{
		SpaceName: "your SpaceName",
		URLSets:   nil,
	}

	resp, status, err := instance.UploadMediaByUrl(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_QueryUploadTaskInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodQueryUploadTaskInfoRequest{
		JobIds: "your JobIds",
	}

	resp, status, err := instance.QueryUploadTaskInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ApplyUploadInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodApplyUploadInfoRequest{
		SpaceName:  "your SpaceName",
		SessionKey: "your SessionKey",
		FileSize:   0,
		FileType:   "your FileType",
		FileName:   "your FileName",
	}

	resp, status, err := instance.ApplyUploadInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CommitUploadInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCommitUploadInfoRequest{
		SpaceName:    "your SpaceName",
		SessionKey:   "your SessionKey",
		CallbackArgs: "your CallbackArgs",
		Functions:    "your Functions",
	}

	resp, status, err := instance.CommitUploadInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}