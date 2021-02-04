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

	query := url.Values{}
	query.Add("task_id", "")

	resp, status, err := visual.DefaultInstance.VideoInpaintQueryTask(query)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
