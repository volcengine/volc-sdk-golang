package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/sami"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	sami.DefaultInstance.Client.SetAccessKey(testAk)
	sami.DefaultInstance.Client.SetSecretKey(testSk)

	param := sami.GetTokenReq{
		Version:    "volc-auth-v1",
		Appkey:     "your appkey",
		Expiration: 3600,
	}

	resp, status, err := sami.DefaultInstance.GetToken(param)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
