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

	//Version>=2022-08-31, Content-Type是application/json类型的通用调用方式
	//请求入参(按照对应接口文档进行调用)
	reqBody := map[string]interface{}{
		"req_key": "",
		//"binary_data_base64": []string{""},
		"image_urls": []string{""},
		"return_url": true,
	}
	//action、version按照接口文档填写即可
	resp, status, err := visual.DefaultInstance.VisualCommonJSONAPI(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

	////Version=2020-08-26，Content-Type是application/x-www-form-urlencoded类型的通用调用方式
	//form := url.Values{}
	//form.Add("", "")
	////form.Add("image_base64", "")
	//
	//resp, status, err := visual.DefaultInstance.VisualCommonFormAPI("", "2020-08-26", form)
	//fmt.Println(status, err)
	//b, _ := json.Marshal(resp)
	//fmt.Println(string(b))

}
