package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ListCdnCertInfo(t *testing.T) {
	resp, err := DefaultInstance.ListCdnCertInfo(nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
