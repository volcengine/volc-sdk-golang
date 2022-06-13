package upload

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/consts"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/functions"
)

func TestVod_UploadMediaMaterialWithCallback(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	spaceName := "your space"
	filePath := "material file path"

	snapShotFunc := functions.SnapshotFunc(0.0)
	getMetaFunc := functions.GetMetaFunc()
	addOptionFunc := functions.AddOptionInfoFunc(vod.OptionInfo{
		Title:       "素材测试视频",             // 标题
		Tags:        "test",               // 多个标签可用逗号隔开
		Description: "素材测试，视频文件",          // 素材描述信息
		Category:    consts.CategoryVideo, // 素材分类，在 video、audio、image、dynamic_img、subtitle、font 中枚举
		RecordType:  2,                    // 素材上传 Type 值为 2
		Format:      "mp4",                //格式。若传入 Format 的话，以您传入参数为准，否则以系统识别出的 Format 为准。若遇到特殊文件无法识别，Format 可能为空。
	})

	vodFunctions := []vod.Function{addOptionFunc, getMetaFunc, snapShotFunc}
	fbts, _ := json.Marshal(vodFunctions)

	vodUploadMaterialRequest := &request.VodUploadMaterialRequest{
		SpaceName:    spaceName,
		FilePath:     filePath,
		CallbackArgs: "my callback",
		Functions:    string(fbts),
		FileType:     consts.FileTypeMedia,
		FileName:     "",
	}

	resp, _, err := instance.UploadMaterialWithCallback(vodUploadMaterialRequest)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("\nresp = %s", bts)
	}

	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetMid())
	fmt.Printf("%d\n", int(resp.GetResult().GetData().GetSourceInfo().GetSize()))

}

func TestVod_UploadImageMaterialWithCallback(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	spaceName := "your space"
	filePath := "material file path"

	getMetaFunc := functions.GetMetaFunc()
	addOptionFunc := functions.AddOptionInfoFunc(vod.OptionInfo{
		Title:       "素材测试图片",
		Tags:        "test",
		Description: "素材测试，图片文件",
		Category:    consts.CategoryImage,
		RecordType:  2,
		Format:      "jpg",
	})

	vodFunctions := []vod.Function{addOptionFunc, getMetaFunc}
	fbts, _ := json.Marshal(vodFunctions)

	vodUploadMaterialRequest := &request.VodUploadMaterialRequest{
		SpaceName:    spaceName,
		FilePath:     filePath,
		CallbackArgs: "my callback",
		Functions:    string(fbts),
		FileType:     consts.FileTypeImage,
		FileName:     "",
	}

	resp, _, err := instance.UploadMaterialWithCallback(vodUploadMaterialRequest)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("\nresp = %s", bts)
	}

	fmt.Println(resp.String())
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetMid())

}

func TestVod_UploadObjectMaterialWithCallback(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	spaceName := "your space"
	filePath := "material file path"

	getMetaFunc := functions.GetMetaFunc()
	addOptionFunc := functions.AddOptionInfoFunc(vod.OptionInfo{
		Title:       "素材测试字幕",
		Tags:        "test",
		Description: "素材测试，字幕文件",
		Category:    consts.CategorySubtitle,
		RecordType:  2,
		Format:      "vtt",
	})

	vodFunctions := []vod.Function{addOptionFunc, getMetaFunc}
	fbts, _ := json.Marshal(vodFunctions)

	vodUploadMaterialRequest := &request.VodUploadMaterialRequest{
		SpaceName:    spaceName,
		FilePath:     filePath,
		CallbackArgs: "my callback",
		Functions:    string(fbts),
		FileType:     consts.FileTypeObject,
		FileName:     "",
	}

	resp, _, err := instance.UploadMaterialWithCallback(vodUploadMaterialRequest)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("\nresp = %s", bts)
	}

	fmt.Println(resp.String())
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetMid())

}
