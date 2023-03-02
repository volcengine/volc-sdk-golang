package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func main() {
	//testAk := "ak"
	//testSk := "sk"
	//
	//// 调用AssumeRole接口获取临时ak/sk及token
	//sts.DefaultInstance.Client.SetAccessKey(testAk)
	//sts.DefaultInstance.Client.SetSecretKey(testSk)
	//
	//list, status, err := sts.DefaultInstance.AssumeRole(&sts.AssumeRoleRequest{
	//	DurationSeconds: 900,
	//	RoleTrn:         "trn:iam::YourAccountID:role/YourRoleName",
	//	RoleSessionName: "jest_for_test",
	//})
	//fmt.Println(status, err)
	//b, _ := json.Marshal(list)
	//fmt.Println(string(b))

	visual.DefaultInstance.Client.SetAccessKey("your-sts-ak")       // sts返回的临时ak
	visual.DefaultInstance.Client.SetSecretKey("your-sts-sk")       // sts返回的临时sk
	visual.DefaultInstance.Client.SetSessionToken("your-sts-token") // sts返回的临时SessionToken

	// CertToken接口示例
	reqBody := &model.CertTokenRequest{
		ReqKey:   "cert_token", // 固定值
		StsToken: "your-sts-token",
		//TosInfo: &struct {
		//	StsAk    string `json:"sts_ak"`
		//	StsSk    string `json:"sts_sk"`
		//	StsToken string `json:"sts_token"`
		//	Bucket   string `json:"bucket"`
		//	Endpoint string `json:"endpoint"`
		//	Region   string `json:"region"`
		//}{
		//	StsAk:    "",
		//	StsSk:    "",
		//	StsToken: "",
		//	Bucket:   "",
		//	Endpoint: "",
		//	Region:   "",
		//},
		RefSource:        "1",
		LivenessType:     "",
		IdCardName:       "",
		IdCardNo:         "",
		RefImage:         "", // 无源比对时必选，输入图片的base64数组，在无源比对时需要传入1张用户的基准图，有源比对无需传入
		LivenessTimeout:  10,
		MotionList:       []string{""},
		FixedMotionList:  []string{""},
		MotionCount:      2,
		MaxLivenessTrial: 10,
		//CallbackInfo: &struct {
		//	Switch     bool   `json:"switch"`
		//	Block      bool   `json:"block"`
		//	Url        string `json:"url"`
		//	ClientName string `json:"client_name"`
		//}{
		//	Switch:     false,
		//	Block:      false,
		//	Url:        "",
		//	ClientName: "",
		//},
		//ConfigID: "",
	}

	resp, status, err := visual.DefaultInstance.CertToken(reqBody)

	// 配置Config示例
	//reqBody := &model.CertConfigInitRequest{}
	//resp, status, err := visual.DefaultInstance.CertConfigInit(reqBody)

	// Query接口示例
	//reqBody := &model.CertVerifyQueryRequest{}
	//resp, status, err := visual.DefaultInstance.CertVerifyQuery(reqBody)

	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
