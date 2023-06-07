package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//请求入参
	reqBody := &model.LensVidaVideoGetResultV2Request{
		ReqKey: "lens_vida_video", // 固定值
		TaskId: "xxx",
	}

	resp, status, err := visual.DefaultInstance.LensVidaVideoGetResultV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
