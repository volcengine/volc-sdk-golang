package upload

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/functions"
)

func TestVod_UploadMediaWithCallback(t *testing.T) {
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
	filePath := "media file path"

	snapShotFunc := functions.SnapshotFunc(1.3)
	getMetaFunc := functions.GetMetaFunc()
	vodFunctions := []vod.Function{snapShotFunc, getMetaFunc}
	fbts, _ := json.Marshal(vodFunctions)

	vodUploadMediaRequset := &request.VodUploadMediaRequest{
		SpaceName:    spaceName,
		FilePath:     filePath,
		CallbackArgs: "my callback",
		Functions:    string(fbts),
		FileName:     "",
	}

	resp, _, err := instance.UploadMediaWithCallback(vodUploadMediaRequset)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("resp = %s", bts)
	}

	fmt.Println()
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetVid())

}
