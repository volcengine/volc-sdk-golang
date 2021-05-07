package play

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	request2 "github.com/volcengine/volc-sdk-golang/service/vod/models/vod/request"
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

	query := &request2.VodGetPrivateDrmPlayAuthRequest{
		Vid: vid,
	}
	newToken, _ := instance.GetPrivateDrmAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}
