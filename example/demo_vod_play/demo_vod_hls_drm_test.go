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
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")
	deadline := time.Now().Add(5 * time.Minute)
	token, _ := instance.CreateSha1HlsDrmAuthToken(deadline)
	fmt.Println(token)
}
