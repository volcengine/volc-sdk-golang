package demo_vod_upload

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/models/vod/business"
	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_UrlUploadURLSet(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	URLSets := []*business.VodUrlUploadURLSet{
		{
			SourceUrl: "url",
		},
	}

	req := &request.VodUrlUploadRequest{
		SpaceName: "your space name",
		URLSets:   URLSets,
	}

	resp, _, err := instance.UploadMediaByUrl(req)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if resp.GetResponseMetadata().GetError() != nil {
		fmt.Println(resp.GetResponseMetadata().GetError())
		return
	}
	bts, _ := json.Marshal(resp)
	fmt.Printf("resp = %s\n", bts)
	fmt.Println(resp.GetResult().GetData()[0].GetSourceUrl())
	fmt.Println(resp.GetResult().GetData()[0].GetJobId())

}
