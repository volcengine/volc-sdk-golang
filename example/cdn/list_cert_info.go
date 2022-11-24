package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func ListCertInfo(t *testing.T) {
	resp, err := DefaultInstance.ListCertInfo(&cdn.ListCertInfoRequest{
		Source: cdn.GetStrPtr("volc_cert_center"),
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
