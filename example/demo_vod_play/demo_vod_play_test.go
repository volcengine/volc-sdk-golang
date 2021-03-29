package demo_vod_play

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_GetPlayInfo(t *testing.T) {
	vid := "your vid"

	// GetPlayInfo
	instance := vod.NewInstance()

	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	query := &request.VodGetPlayInfoRequest{
		Vid:        vid,
		Format:     "",
		Codec:      "",
		Definition: "360p",
		FileType:   "",
		LogoType:   "",
		Ssl:        "",
	}
	resp, code, err := instance.GetPlayInfo(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	resp.GetResponseMetadata().GetError().String()
	resp.GetResult().String()
	fmt.Println(string(b))
}
