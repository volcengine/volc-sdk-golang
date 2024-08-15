package dts_v20221001_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/dts/v20221001"
)

func Test_SetBiSyncDDLSource(t *testing.T) {
	instance := dts_v20221001.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &dts_v20221001.SetBiSyncDDLSourceBody{}

	resp, err := instance.SetBiSyncDDLSource(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
