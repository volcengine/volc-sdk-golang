package demo_vod_upload

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/functions"
	"io/ioutil"
	"os"
	"testing"
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

	spaceName := "your space name"
	filePath := "media file path"

	snapShotFunc := functions.SnapshotFunc(1.3)
	getMetaFunc := functions.GetMeatFunc()

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file from %s error %v", filePath, err)
		os.Exit(-1)
	}

	resp, _, err := instance.UploadMediaWithCallback(dat, spaceName, "my callback", getMetaFunc, snapShotFunc)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("\nresp = %s", bts)
	}

	fmt.Println()
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetVid())

}
