package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//请求入参
	reqBody := map[string]interface{}{
		"req_key":            "i2i_inpainting_edit",
		"binary_data_base64": []string{"原图", "原图标注后的mask"},
		"custom_prompt":      "", // 写入Prompt，AIGC生成取代的内容
		"scale":              5,
		"seed":               -1,
		"steps":              25,
	}

	resp, status, err := visual.DefaultInstance.Img2ImgInpaintingEdit(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
