package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func AddCertificate(t *testing.T) {
	resp, err := DefaultInstance.AddCertificate(&cdn.AddCertificateRequest{
		Certificate: "-----BEGIN CERTIFICATE-----\\r\\nHmU2w=\\r\\n-----END CERTIFICATE-----\\r\\n-----BEGIN CERTIFICATE-----\\r\\nMIIElg==\\r\\n-----END CERTIFICATE-----\\r\\n-----BEGIN CERTIFICATE-----\\r\\nMIIDV5Adg\\r\\n06O/nVsJ8dWd4=\\r\\n-----END CERTIFICATE-----",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
