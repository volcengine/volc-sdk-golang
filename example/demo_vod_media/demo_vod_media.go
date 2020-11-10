package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")

	//vid := "your vid"

	// GetPlayInfo
	instance := vod.NewInstance()

	query := &request.VodGetVideoInfosRequest{
		Vids: vid,
	}
	resp, code, err := instance.GetVideoInfos(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

	query1 := &request.VodGetRecommendedPosterRequest{
		Vids: vid,
	}
	resp1, code, err := instance.GetRecommendedPoster(query1)
	fmt.Println(code)
	fmt.Println(err)
	b, _ = json.Marshal(resp1)
	fmt.Println(string(b))

	query2 := &request.VodUpdateVideoInfoRequest{
		Vid:   vid,
		Title: wrapperspb.String("aaaaa"),
		//Tags:  wrapperspb.String("aaa,aa:w"),
	}
	resp2, code, err := instance.UpdateVideoInfo(query2)
	fmt.Println(code)
	fmt.Println(err)
	b, _ = json.Marshal(resp2)
	fmt.Println(string(b))

	query3 := &request.VodUpdateVideoPublishStatusRequest{
		Vid:    vid,
		Status: "Unpublished",
	}
	resp3, code, err := instance.UpdateVideoPublishStatus(query3)
	fmt.Println(code)
	fmt.Println(err)
	b, _ = json.Marshal(resp3)
	fmt.Println(string(b))

}
