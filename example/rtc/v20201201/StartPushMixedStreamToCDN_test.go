package rtc_v20201201_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/rtc/v20201201"
)

func Test_StartPushMixedStreamToCDN(t *testing.T) {
	instance := rtc_v20201201.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &rtc_v20201201.StartPushMixedStreamToCDNBody{}

	resp, statusCode, err := instance.StartPushMixedStreamToCDN(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v statusCode %d", err, statusCode)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}