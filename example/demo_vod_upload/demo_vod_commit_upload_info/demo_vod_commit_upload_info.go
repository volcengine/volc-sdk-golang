package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/functions"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	space := "your space name"
	session := "apply return session"

	funcs := make([]vod.Function, 0)

	snapShotFunc := functions.SnapshotFunc(2.3)
	getMetaFunc := functions.GetMeatFunc()

	funcs = append(funcs, snapShotFunc)
	funcs = append(funcs, getMetaFunc)

	fbts, err := json.Marshal(funcs)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("\n%s", fbts)

	commitRequest := &request.VodCommitUploadInfoRequest{
		SpaceName:    space,
		SessionKey:   session,
		CallbackArgs: "my callback",
		Functions:    string(fbts),
	}
	resp, _, err := instance.CommitUploadInfo(commitRequest)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.GetResponseMetadata().GetError() != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}

	fmt.Println(resp.GetResult().GetData().GetVid())
	bts, _ := json.Marshal(resp)
	fmt.Printf("\nresp = %s", bts)
}
