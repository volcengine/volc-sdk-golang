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
	form.Add("video_url", "")

	resp, status, err := visual.DefaultInstance.VideoCoverSelection(form)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
	// for idx, v := range resp.Data.Results {
	// 	img := base64.NewDecoder(base64.StdEncoding, strings.NewReader(v.Data))
	// 	buff, _ := ioutil.ReadAll(img)
	// 	fmt.Println(len(buff))
	// 	err := ioutil.WriteFile("./test.jpg", buff, 0666)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Printf("idx: %d, score: %f", idx, v.Score)
	// }
}
