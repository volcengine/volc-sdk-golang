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
		"req_key": "face_swap3_3",
		//"binary_data_base64": []string{"",""},
		"image_urls": []string{"https://xxx", "https://xxx"}, //输入换脸和模板图片链接数组，换脸图在前（最多三张），模板图在后（最多一张）
		"do_risk":    false,
		"face_type":  "l2r",
		"merge_infos": []map[string]int{
			{
				"location":          1,
				"template_location": 1,
			},
		},
		"logo_info": map[string]interface{}{
			"add_logo": true,
			"position": 1,
			"language": 1,
			"opacity":  0.3,
		},
		"source_similarity": "1.0",
		"gpen":              1.0,
	}

	resp, status, err := visual.DefaultInstance.FaceSwapV3(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
