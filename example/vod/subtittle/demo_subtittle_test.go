package subtittle

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestVod_GetSubtitleInfoList(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})
	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")

	// Media Info
	query := &request.VodGetSubtitleInfoListRequest{
		Vid:         "your search vid",
		FileIds:     "your search fileIds",
		Languages:   "your search languages",
		LanguageIds: "your search languageIds",
		SubtitleIds: "your search subtitleIds",
		Formats:     "your search format",
		Status:      "your search status",
		Title:       "your search title",
		Tag:         "your search tag",
		Offset:      "your search offset",
		PageSize:    "your search page size",
		Ssl:         "your search ssl type",
	}
	resp, code, err := instance.GetSubtitleInfoList(query)
	fmt.Println(code)
	fmt.Println(err)
	fmt.Println(resp.Result.Offset)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_UpdateSubtitleStatus(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})
	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")

	// Media Info
	query := &request.VodUpdateSubtitleStatusRequest{
		Vid:       "your update vid",
		FileIds:   "your update fileIds",
		Languages: "your update languages",
		Formats:   "your update formats",
		Status:    "your update status",
	}
	resp, code, err := instance.UpdateSubtitleStatus(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_UpdateSubtitleInfo(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})
	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")

	// Media Info
	query := &request.VodUpdateSubtitleInfoRequest{
		Vid:      "your update vid",
		FileId:   "your update fileId",
		Language: "your update language",
		Format:   "your update format",
		Title:    wrapperspb.String("your update title"),
		Tag:      wrapperspb.String("your update tag"),
	}
	resp, code, err := instance.UpdateSubtitleInfo(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_GetSubtitleAuthToken(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})
	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")

	// Media Info
	query := &request.VodGetSubtitleInfoListRequest{
		Vid:         "your search vid",
	}
	expireSeconds := 100 // expire time duration: (s)
	token, err := instance.GetSubtitleAuthToken(query, expireSeconds)
	fmt.Println("token ===> ", token)
	fmt.Println("err ===> ", err)
}
