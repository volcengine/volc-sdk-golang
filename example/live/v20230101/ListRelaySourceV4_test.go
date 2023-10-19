package live_v20230101_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/live/v20230101"
)

func Test_ListRelaySourceV4(t *testing.T) {
	instance := live_v20230101.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &live_v20230101.ListRelaySourceV4Body{}

	resp, err := instance.ListRelaySourceV4(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
