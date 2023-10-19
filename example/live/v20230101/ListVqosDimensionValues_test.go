package live_v20230101_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/live/v20230101"
)

func Test_ListVqosDimensionValues(t *testing.T) {
	instance := live_v20230101.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &live_v20230101.ListVqosDimensionValuesReq{
		ListVqosDimensionValuesQuery: &live_v20230101.ListVqosDimensionValuesQuery{},
		ListVqosDimensionValuesBody:  &live_v20230101.ListVqosDimensionValuesBody{},
	}

	resp, err := instance.ListVqosDimensionValues(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
