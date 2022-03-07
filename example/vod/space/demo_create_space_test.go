package space

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func TestVod_CreateSpace(t *testing.T) {
	instance := vod.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	query := &request.VodCreateSpaceRequest{
		SpaceName:   "your space name",
		ProjectName: "your project name",
		Description: "your desc",
		Region:      "same as vod instance region",
	}
	resp, code, err := instance.CreateSpace(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	resp.GetResponseMetadata().GetError().String()
	fmt.Println(string(b))
}
