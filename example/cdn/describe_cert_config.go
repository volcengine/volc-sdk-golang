package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DescribeCertConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeCertConfig(&cdn.DescribeCertConfigRequest{
		CertId: cdn.GetStrPtr("cert-2b12dd79c3ef441ea1e58a09248d0fd6"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
