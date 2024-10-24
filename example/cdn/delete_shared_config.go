package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DeleteSharedConfig(t *testing.T) {
	resp, err := DefaultInstance.DeleteSharedConfig(&cdn.DeleteSharedConfigRequest{
		ConfigName: "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
