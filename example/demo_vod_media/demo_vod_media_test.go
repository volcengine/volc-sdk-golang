package demo_vod_media

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestVod_GetMediaInfos(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	// Media Info
	instance := vod.NewInstance()
	query := &request.VodGetMediaInfosRequest{
		Vids: vid,
	}
	resp, code, err := instance.GetMediaInfos(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_GetRecommendedPoster(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	instance := vod.NewInstance()
	query1 := &request.VodGetRecommendedPosterRequest{
		Vids: vid,
	}
	resp1, code, err := instance.GetRecommendedPoster(query1)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp1)
	fmt.Println(string(b))
}

func TestVod_UpdateMediaInfo(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	instance := vod.NewInstance()
	query2 := &request.VodUpdateMediaInfoRequest{
		Vid:   vid,
		Title: wrapperspb.String("aaaaa"),
		//Tags:  wrapperspb.String("aaa,aa:w"),
	}
	resp2, code, err := instance.UpdateMediaInfo(query2)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp2)
	fmt.Println(string(b))
}

func TestVod_UpdateMediaPublishStatus(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	instance := vod.NewInstance()
	query3 := &request.VodUpdateMediaPublishStatusRequest{
		Vid:    vid,
		Status: "Unpublished",
	}
	resp3, code, err := instance.UpdateMediaPublishStatus(query3)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp3)
	fmt.Println(string(b))
}
