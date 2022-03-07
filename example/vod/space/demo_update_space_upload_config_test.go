package space

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func TestVod_UpdateSpaceUploadConfig(t *testing.T) {
	instance := vod.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	query := &request.VodUpdateSpaceUploadConfigRequest{
		SpaceName:   "your space name",
		ConfigKey:   "your config key",
		ConfigValue: "your config value",
	}
	resp, code, err := instance.UpdateSpaceUploadConfig(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	resp.GetResponseMetadata().GetError().String()
	fmt.Println(string(b))
}
