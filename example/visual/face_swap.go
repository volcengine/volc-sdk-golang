package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	form := url.Values{}
	form.Add("image_base64", "xxx") //2.1版本无需此参数
	//form.Add("image_url", "xxx") //2.1版本无需此参数
	form.Add("template_base64", "xxx")
	form.Add("template_url", "xxx")
	form.Add("action_id", "faceswap")
	form.Add("version", "") //1.0, 2.0, 2.1
	form.Add("do_risk", "true")
	form.Add("type", "") //仅限2.1版本
	mergeInfoMap := map[string]interface{}{
		//"image_base64": "/9xxx",
		"image_url":         "https://xxx",
		"location":          1,
		"template_location": 1,
	}
	mergeInfo, err := json.Marshal(mergeInfoMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	form.Add("merge_infos", string(mergeInfo)) //仅限2.1版本

	resp, status, err := visual.DefaultInstance.FaceSwap(form)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
