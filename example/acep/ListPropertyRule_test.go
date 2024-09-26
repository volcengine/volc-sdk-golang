package acep_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	ACEP "github.com/volcengine/volc-sdk-golang/service/acep"
)

func StringPtr(v string) *string {
	return &v
}

func Int64Ptr(v int64) *int64 {
	return &v
}

func Test_ListPropertyRule(t *testing.T) {
	// 强烈建议不要把 VOLC_ACCESSKEY 和 VOLC_SECRETKEY 保存到工程代码里，否则可能导致 AccessKey 泄露，威胁您账号下所有资源的安全。
	// 本示例通过从环境变量中读取 VOLC_ACCESSKEY 和 VOLC_SECRETKEY，来实现 API 访问的身份验证。运行代码示例前，请配置环境变量 VOLC_ACCESSKEY 和 VOLC_SECRETKEY
	service := ACEP.NewInstance()
	service.SetCredential(base.Credentials{
		AccessKeyID:     os.Getenv("VOLC_ACCESSKEY"),
		SecretAccessKey: os.Getenv("VOLC_SECRETKEY"),
	})

	query := &ACEP.ListPropertyRuleQuery{
		// 机型参数规则名称，账户下规则名称唯一
		//RuleName: StringPtr(`Oppo1`),
		// 分页偏移量。默认值为 0。
		Offset: Int64Ptr(0),
		// 分页大小，默认值为 10，最大值为 100。
		Count: Int64Ptr(10),
	}

	resp, err := service.ListPropertyRule(context.Background(), query)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		data, _ := json.Marshal(resp)
		fmt.Printf("success %+v", string(data))
	}
}
