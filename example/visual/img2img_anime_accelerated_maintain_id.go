package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func main() {
	testAk := ""
	testSk := ""

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//请求入参
	reqBody := map[string]interface{}{
		"req_key": "img2img_anime_accelerated_maintain_id",
		//"binary_data_base64": []string{""},
		"image_urls":      []string{""},
		"positive_prompt": "",
		"return_url":      true,
	}

	resp, status, err := visual.DefaultInstance.Img2ImgAnimeAcceleratedMaintainID(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
