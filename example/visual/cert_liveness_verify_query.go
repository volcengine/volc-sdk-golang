package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/visual"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func main() {

	visual.DefaultInstance.Client.SetAccessKey("your-sts-ak") // sts返回的临时ak
	visual.DefaultInstance.Client.SetSecretKey("your-sts-sk") // sts返回的临时sk

	// Query接口示例
	reqBody := &model.CertLivenessVerifyQueryRequest{
		ReqKey:        "cert_pro_liveness_verify_query", // 固定值
		BytedToken:    "",                               // 通过Token接口获取的byted_token
		OmitData:      false,
		OmitImageData: false,
		OmitVideoData: false,
	}
	resp, status, err := visual.DefaultInstance.CertLivenessVerifyQuery(reqBody)

	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
