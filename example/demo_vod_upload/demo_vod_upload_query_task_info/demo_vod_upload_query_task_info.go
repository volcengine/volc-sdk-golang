package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"strings"
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

	jobIds := make([]string, 0)
	jobId := "url jobId"
	jobIds = append(jobIds, jobId)
	str := strings.Join(jobIds, ",")

	queryRequest := &request.VodQueryUploadTaskInfoRequest{JobIds: str}
	resp, _, err := instance.QueryUploadTaskInfo(queryRequest)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.ResponseMetadata.Error != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}
	fmt.Println(resp.GetResult().GetData().GetVideoInfoList()[0].GetVid())
	bts, _ := json.Marshal(resp)
	fmt.Printf("resp = %s", bts)
}
