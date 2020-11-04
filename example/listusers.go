package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/iam"
)

const (
	region = "region"
	host   = "hosts"
	testAk = "ak"
	testSk = "sk"
)

func main() {
	iam.DefaultInstance.Client.SetAccessKey(testAk)
	iam.DefaultInstance.Client.SetSecretKey(testSk)
	iam.DefaultInstance.SetHost(host)
	iam.DefaultInstance.SetRegion(region)

	list, status, err := iam.DefaultInstance.ListUsers(nil, nil)
	fmt.Println(status, err)
	b, _ := json.Marshal(list)
	fmt.Println(string(b))
}
