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

	reqBody := map[string]interface{}{
		"req_key": "face_compare", // 固定值
		//"binary_data_base64": []string{"/9xxx"},
		"image_urls": []string{"https://1xxx", "https://2xxx"},
	}

	resp, status, err := visual.DefaultInstance.FaceCompare(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
