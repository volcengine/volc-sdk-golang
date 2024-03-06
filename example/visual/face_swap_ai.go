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
		"req_key": "faceswap_ai",
		//"binary_data_base64": []string{"",""},
		"image_urls": []string{"http://xxx", "http://xxx"}, //List长度为2，输入两张图，将第一张图人脸融入第二张图
		"do_risk":    false,
		"gpen":       0.9,
		"skin":       0.1,
		"tou_repair": 1,
	}

	resp, status, err := visual.DefaultInstance.FaceSwapAI(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
