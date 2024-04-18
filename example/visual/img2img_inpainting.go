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
		"req_key":            "i2i_inpainting",
		"binary_data_base64": []string{"原图", "原图标注后的mask"},
		"scale":              7,
		"seed":               0,
		"steps":              30,
		"strength":           0.8,
	}

	resp, status, err := visual.DefaultInstance.Img2ImgInpainting(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
