package demo_vod_play

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_GetPlayInfo(t *testing.T) {
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	vid := "your vid"

	// GetPlayInfo
	instance := vod.NewInstance()

	query := &request.VodGetPlayInfoRequest{
		Vid:        vid,
		Format:     "",
		Codec:      "",
		Definition: "360p",
		FileType:   "",
		LogoType:   "",
		Base64:     "1",
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

func TestVod_GetOriginalPlayInfo(t *testing.T) {
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	vid := "your vid"
	instance := vod.NewInstance()
	// GetOriginalPlayInfo
	query2 := &request.VodGetOriginalPlayInfoRequest{Vid: vid}
	resp2, code2, err2 := instance.GetOriginalPlayInfo(query2)
	fmt.Println(code2)
	fmt.Println(err2)
	b2, _ := json.Marshal(resp2)
	fmt.Println(string(b2))
}
