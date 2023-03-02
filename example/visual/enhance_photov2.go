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
	enhanceReq := &model.EnhancePhotoV2Request{
		ReqKey:             "lens_lqir", // 固定值
		BinaryDataBase64:   []string{"image_base64"},
		ResolutionBoundary: "720p", // 是否走超分的判断条件
		EnableHdr:          false,  // 是否开启hdr能力
		EnableWb:           false,  // 是否开启白平衡能力
		ResultFormat:       0,      //0 代表结果图片为png格式, 1 代表结果图片为jpeg格式
		JpgQuality:         95,     // 值越高，代表生成jpg图片的质量越高
	}

	resp, status, err := visual.DefaultInstance.EnhancePhotoV2(enhanceReq)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
