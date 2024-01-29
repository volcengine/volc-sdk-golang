package imagex_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

func Test_GetImageDuplicateDetection(t *testing.T) {
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.GetImageDuplicateDetectionReq{
		GetImageDuplicateDetectionQuery: &imagex.GetImageDuplicateDetectionQuery{},
		GetImageDuplicateDetectionBody:  &imagex.GetImageDuplicateDetectionBody{},
	}

	resp, err := instance.GetImageDuplicateDetection(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
