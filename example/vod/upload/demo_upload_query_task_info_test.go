package upload

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"

	"strings"
	"testing"
)

func TestVod_QueryUploadTaskInfo(t *testing.T) {
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
		fmt.Printf("err:%v\n", err)
	}
	if resp.ResponseMetadata.Error != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}
	fmt.Println(resp.GetResult().GetData().GetMediaInfoList()[0].GetVid())
	bts, _ := json.Marshal(resp)
	fmt.Printf("resp = %s", bts)
}
