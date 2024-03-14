package dts_v20180101_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/dts/v20180101"
)

func Test_GetAsyncPreCheckResult(t *testing.T) {
	instance := dts_v20180101.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &dts_v20180101.GetAsyncPreCheckResultBody{}

	resp, err := instance.GetAsyncPreCheckResult(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
