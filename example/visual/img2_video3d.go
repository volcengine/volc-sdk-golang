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

	//请求入参
	reqBody := &model.Img2Video3DRequest{
		ReqKey:           "img2video3d", // 固定值
		BinaryDataBase64: []string{"image_base64"},
		//RenderSpec: &model.Img2Video3DRenderSpec{Mode: 0, LongSide: 960, FrameNum: 90, Fps: 30, UseFlow: -1, SpeedShift: []float64{0,1,0.5,4, 0.5,4,1,1}},
	}

	resp, status, err := visual.DefaultInstance.Img2Video3D(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
