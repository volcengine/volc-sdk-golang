package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DescribeSharedConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeSharedConfig(&cdn.DescribeSharedConfigRequest{
		ConfigName: cdn.GetStrPtr("test"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
