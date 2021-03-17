package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/iam"
)

const (
	testAk           = "testAK"
	testSk           = "testSK"
	testSessionToken = "testSessionToken"
)

func main() {
	iam.DefaultInstance.Client.SetAccessKey(testAk)
	iam.DefaultInstance.Client.SetSecretKey(testSk)
	iam.DefaultInstance.Client.SetSessionToken(testSessionToken)

	list, status, err := iam.DefaultInstance.ListUsers(nil, nil)
	fmt.Println(status, err)
	b, _ := json.Marshal(list)
	fmt.Println(string(b))
}
