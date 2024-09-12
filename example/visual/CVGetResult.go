package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/visual"
	// "github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func main() {
	testAk := ""
	testSk := ""

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//请求入参(按照对应接口文档进行调用)
	reqBody := map[string]interface{}{
		"req_key": "",
		"task_id": "",
	}

	resp, status, err := visual.DefaultInstance.CVGetResult(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

	// 示例：将resp转化为对应response结构体后使用
	// if status != 200 {
	// 	fmt.Println("request err, response is:", string(b))
	// 	return
	// }
	// respData := new(model.ResponseStructxxx)
	// if err = json.Unmarshal(b, respData); err != nil {
	// 	fmt.Println("unmarshal err:", err)
	// 	return
	// }
	// fmt.Println(respData.Data.Status)
}
