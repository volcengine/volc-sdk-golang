package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func TestVod_CreateCdnRefreshTask(t *testing.T) {
	instance := vod.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	query := &request.VodCreateCdnRefreshTaskRequest{
		SpaceName: "your space name",
		Type:      "your type",
		Urls:      "your urls",
	}
	resp, code, err := instance.CreateCdnRefreshTask(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	resp.GetResponseMetadata().GetError().String()
	fmt.Println(string(b))
}
