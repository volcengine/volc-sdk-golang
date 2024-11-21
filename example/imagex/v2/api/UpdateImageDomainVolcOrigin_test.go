package imagex_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

func Test_UpdateImageDomainVolcOrigin(t *testing.T) {
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.UpdateImageDomainVolcOriginReq{
		UpdateImageDomainVolcOriginQuery: &imagex.UpdateImageDomainVolcOriginQuery{},
		UpdateImageDomainVolcOriginBody:  &imagex.UpdateImageDomainVolcOriginBody{},
	}

	resp, err := instance.UpdateImageDomainVolcOrigin(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
