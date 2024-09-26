package acep_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/acep"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
)

func Test_CreatePod(t *testing.T) {
	instance := acep.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &acep.CreatePodBody{}

	resp, err := instance.CreatePod(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
