package upload

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/business"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/functions"
)

func TestVod_UploadObjectWithCallback(t *testing.T) {
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
	filePath := "object file path"

	getMetaFunc := functions.GetMetaFunc()
	vodFunctions := []business.VodUploadFunction{getMetaFunc}
	fbts, _ := json.Marshal(vodFunctions)

	resp, _, err := instance.UploadObjectWithCallback(filePath, spaceName, "my callback", "a/b/c/d.ts", "", string(fbts))
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("\nresp = %s", bts)
	}

	fmt.Println()
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetSourceInfo().GetSize())
	fmt.Println(resp.GetResult().GetData().GetSourceInfo().GetStoreUri())

}

func TestVod_UploadObjectWithCallbackV2(t *testing.T) {
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
	filePath := "file path"

	getMetaFunc := functions.GetMetaFunc()
	vodFunctions := []business.VodUploadFunction{getMetaFunc}
	fbts, _ := json.Marshal(vodFunctions)

	resp, _, err := instance.UploadObjectWithCallbackV2(&request.VodUploadObjectRequest{
		FilePath:          filePath,
		SpaceName:         spaceName,
		CallbackArgs:      "my callback",
		FileName:          "a/b/c/d.ts",
		FileExtension:     "",
		ClientNetWorkMode: "",
		ClientIDCMode:     "",
		Functions:         string(fbts),
	})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("\nresp = %s", bts)
	}

	fmt.Println()
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetSourceInfo().GetSize())
	fmt.Println(resp.GetResult().GetData().GetSourceInfo().GetStoreUri())

}
