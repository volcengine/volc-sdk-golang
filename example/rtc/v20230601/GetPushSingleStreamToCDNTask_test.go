package rtc_v20230601_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/rtc/v20230601"
)

func Test_GetPushSingleStreamToCDNTask(t *testing.T) {
	instance := rtc_v20230601.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &rtc_v20230601.GetPushSingleStreamToCDNTaskQuery{}

	resp, statusCode, err := instance.GetPushSingleStreamToCDNTask(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v statusCode %d", err, statusCode)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
