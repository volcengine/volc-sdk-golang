package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DescribeCertConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeCertConfig(&cdn.DescribeCertConfigRequest{
		CertId: "cert-123",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
