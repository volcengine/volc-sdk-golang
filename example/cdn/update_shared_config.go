package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func UpdateSharedConfig(t *testing.T) {
	resp, err := DefaultInstance.UpdateSharedConfig(&cdn.UpdateSharedConfigRequest{
		ConfigName: "test",
		AllowIpAccessRule: &cdn.GlobalIPAccessRule{
			Rules: []string{"1.1.1.1"},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
