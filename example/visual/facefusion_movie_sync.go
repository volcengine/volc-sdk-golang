package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	reqBody := &model.FaceFusionMovieRequest{
		ReqKey:             "facefusionmovie_standard", // _usage
		ImageUrl:           "xxx",
		VideoUrl:           "xxx",
		EnableFaceBeautify: true,
		RefImgUrl:          "xxx",
	}

	resp, status, err := visual.DefaultInstance.FaceFusionMovie(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
