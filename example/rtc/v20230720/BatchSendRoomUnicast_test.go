package rtc_v20230720_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/rtc/v20230720"
)

func Test_BatchSendRoomUnicast(t *testing.T) {
	instance := rtc_v20230720.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &rtc_v20230720.BatchSendRoomUnicastBody{}

	resp, statusCode, err := instance.BatchSendRoomUnicast(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v statusCode %d", err, statusCode)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
