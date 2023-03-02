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
	reqBody := &model.Img2ImgStyleRequest{
		ReqKey: "img2img_style", // 固定值
		Prompt: "",              // 引导词文本 ，控制转换图像风格
		//Strength:         0.5,
		//Seed:             -1,
		BinaryDataBase64: []string{""},
		//ImageUrl:         "",
	}

	resp, status, err := visual.DefaultInstance.Img2ImgStyle(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
