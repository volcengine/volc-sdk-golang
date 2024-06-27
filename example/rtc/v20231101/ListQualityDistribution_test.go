package rtc_v20231101_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/rtc/v20231101"
)

func Test_ListQualityDistribution(t *testing.T) {
	instance := rtc_v20231101.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &rtc_v20231101.ListQualityDistributionBody{}

	resp, statusCode, err := instance.ListQualityDistribution(context.Background(), param)

	if err != nil {
		if resp != nil && resp.ResponseMetadata.Error != nil {
			errStr, _ := json.Marshal(resp.ResponseMetadata.Error)
			t.Logf("statusCode: %d, error: %v", statusCode, string(errStr))
			// 网关返回的错误
			if resp.ResponseMetadata.Error.CodeN != nil && *resp.ResponseMetadata.Error.CodeN != 0 {
				switch *resp.ResponseMetadata.Error.CodeN {
				// InvalidAccessKey
				case 100009:
					t.Logf("请求的AK不合法")
				// SignatureDoesNotMatch
				case 100010:
					t.Logf("签名结果不正确")
				}
			} else {
				// 服务端返回的错误
				switch resp.ResponseMetadata.Error.Code {
				case "InvalidParameter":
					t.Logf("请求的参数错误, 请根据具体Error中的Message提示调整参数")
				}
			}
		} else {
			t.Errorf("statusCode: %d, error: %v", statusCode, err)
		}
	} else {
		b, _ := json.Marshal(resp)
		t.Logf("success %v", string(b))
	}
}
