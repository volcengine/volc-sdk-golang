package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func ListSharedConfig(t *testing.T) {
	var num, size = int64(1), int64(100)
	resp, err := DefaultInstance.ListSharedConfig(&cdn.ListSharedConfigRequest{
		ConfigName: cdn.GetStrPtr("test"),
		PageNum:    &num,
		PageSize:   &size,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
