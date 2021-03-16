package sts

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	testAk = "testAK"
	testSk = "testSK"
)

func TestIAM_AssumeRole(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	req := &AssumeRoleRequest{
		DurationSeconds: 3600,
		Policy:          "",
		RoleTrn:         "test-role-trn",
		RoleSessionName: "test",
	}

	list, status, err := DefaultInstance.AssumeRole(req)
	fmt.Println(status, err)
	b, _ := json.Marshal(list)
	fmt.Println(string(b))
}
