package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func AddSharedConfig(t *testing.T) {
	resp, err := DefaultInstance.AddSharedConfig(&cdn.AddSharedConfigRequest{
		ConfigName: "test",
		Project:    cdn.GetStrPtr("default"),
		ConfigType: "deny_ip_access_rule",
		AllowIpAccessRule: &cdn.GlobalIPAccessRule{
			Rules: []string{"1.1.1.1"},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
