package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DeleteCdnCertificate(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnCertificate(&cdn.DeleteCdnCertificateRequest{
		CertId: "cert_hosting-1234",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
