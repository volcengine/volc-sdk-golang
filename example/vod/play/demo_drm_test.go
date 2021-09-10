package play

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
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
	instance := vod.NewInstance()

	//instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")
	query := &request.VodGetPrivateDrmPlayAuthRequest{
		Vid: "your vid",
		DrmType: "your drmType",
		PlayAuthIds: "your playAuthIds",
		UnionInfo: "your unionInfo",
	}
	tokenExpireTime := 6000000 // Token Expire Duration（s）
	newToken, _ := instance.GetPrivateDrmAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}

func TestVod_GetPrivateDrmAuth(t *testing.T) {
	instance := vod.NewInstance()

	//instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")
	query := &request.VodGetPrivateDrmPlayAuthRequest{
		Vid: "your vid",
		DrmType: "your drmType",
		PlayAuthIds: "your playAuthIds",
		UnionInfo: "your unionInfo",
	}
	newToken, statusCode, err := instance.GetPrivateDrmPlayAuth(query)
	fmt.Println(newToken, statusCode, err)
}