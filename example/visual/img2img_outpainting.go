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
		"req_key":            "i2i_outpainting",
		"prompt":             "蓝色的海洋",
		"binary_data_base64": []string{"原图base64"},
		"scale":              7,
		"seed":               -1,
		"steps":              30,
		"strength":           0.8,
		"top":                0.1,
		"bottom":             0.1,
		"left":               1,
		"right":              1,
		"max_height":         1920,
		"max_width":          1920,
	}
	// 使用画布
	//reqBody := map[string]interface{}{
	//	"req_key":            "i2i_outpainting",
	//	"prompt":             "蓝色的海洋",
	//	"binary_data_base64": []string{"延边图base64", "延边图mask"},
	//	"scale":              7,
	//	"seed":               -1,
	//	"steps":              30,
	//	"strength":           0.8,
	//	"max_height":         1920,
	//	"max_width":          1920,
	//}

	resp, status, err := visual.DefaultInstance.Img2ImgOutpainting(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
