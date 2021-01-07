package demo_vod_play

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_GetPlayAuthToken(t *testing.T) {
	vid := "your vid"
	tokenExpireTime := 600 // Token Expire Duration（s）
	instance := vod.NewInstance()

	//instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
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
		Base64:     "1",
		Ssl:        "",
		NeedThumbs: "1",
	}
	newToken, _ := instance.GetPlayAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}
