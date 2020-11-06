package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	//vid := "your vid"
	vid := "v0c2c369007abu04ru8riko30uo9n73g"

	// GetPlayInfo
	instance := vod.NewInstance()

	query := &request.VodGetPlayInfoRequest{
		Vid:        vid,
		Format:     "",
		Codec:      "",
		Definition: "360p",
		FileType:   "",
		LogoType:   "",
		Base64:     "1",
		Ssl:        "",
	}
	resp, code, err := instance.GetPlayInfo(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))


	// GetOriginalPlayInfo
	query2 := &request.VodGetOriginalPlayInfoRequest{Vid: vid}
	resp2, code2, err2 := instance.GetOriginalPlayInfo(query2)
	fmt.Println(code2)
	fmt.Println(err2)
	b2, _ := json.Marshal(resp2)
	fmt.Println(string(b2))

}
