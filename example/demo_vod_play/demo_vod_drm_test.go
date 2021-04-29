package demo_vod_play

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_GetSha1HlsDrmAuthToken(t *testing.T) {
	instance := vod.NewInstance()
	//instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//instance.SetAccessKey("your ak")
	//instance.SetSecretKey("your sk")
	expireDuration := int64(6000000) //change to your expire duration (s), no default duration
	token, _ := instance.CreateSha1HlsDrmAuthToken(expireDuration)
	fmt.Println(token)
}

func TestVod_GetPrivateDrmAuthToken(t *testing.T) {
	vid := "your vid"
	tokenExpireTime := 9000000 // Token Expire Duration（s）
	instance := vod.NewInstance()

	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	query := &request.VodGetPrivateDrmPlayAuthRequest{
		Vid: vid,
	}
	newToken, _ := instance.GetPrivateDrmAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}
