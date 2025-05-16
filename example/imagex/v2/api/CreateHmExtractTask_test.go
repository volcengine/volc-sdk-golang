package imagex_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

func Test_CreateHmExtractTask(t *testing.T) {
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.CreateHmExtractTaskReq{
		CreateHmExtractTaskQuery: &imagex.CreateHmExtractTaskQuery{},
		CreateHmExtractTaskBody:  &imagex.CreateHmExtractTaskBody{},
	}

	resp, err := instance.CreateHmExtractTask(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
