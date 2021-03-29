package demo_vod_play

import (
	"fmt"
	"testing"
	"time"

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
	expireDuration := int64(600 * time.Second) //change to your expire duration (s), no default duration
	token, _ := instance.CreateSha1HlsDrmAuthToken(expireDuration)
	fmt.Println(token)
}
