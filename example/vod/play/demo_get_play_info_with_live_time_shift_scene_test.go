package play

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func TestVod_GetPlayInfoWithLiveTimeShiftScene(t *testing.T) {
	instance := vod.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	query := &request.VodGetPlayInfoWithLiveTimeShiftSceneRequest{
		SpaceName:             "your space name",
		StoreUris:             "your storeURIs",
		Ssl:                   "1",
		NeedComposeBucketName: "0",
	}
	resp, code, err := instance.GetPlayInfoWithLiveTimeShiftScene(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	resp.GetResponseMetadata().GetError().String()
	resp.GetResult().String()
	fmt.Println(string(b))
}
