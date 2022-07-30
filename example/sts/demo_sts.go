package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/sts"
)

const (
	testAk = "testAK"
	testSk = "testSK"
)

func main() {
	sts.DefaultInstance.Client.SetAccessKey(testAk)
	sts.DefaultInstance.Client.SetSecretKey(testSk)

	list, status, err := sts.DefaultInstance.AssumeRole(&sts.AssumeRoleRequest{
		DurationSeconds: 900,
		RoleTrn:         "trn:iam::YourAccountID:role/YourRoleName",
		RoleSessionName: "jest_for_test",
	})
	fmt.Println(status, err)
	b, _ := json.Marshal(list)
	fmt.Println(string(b))
}
