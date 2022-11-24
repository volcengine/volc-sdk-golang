package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func AddCdnCertificate(t *testing.T) {
	resp, err := DefaultInstance.AddCdnCertificate(&cdn.AddCdnCertificateRequest{
		Certificate: cdn.Certificate{
			Certificate: cdn.GetStrPtr("-----BEGIN CERTIFICATE-----\\r\\nHmU2w=\\r\\n-----END CERTIFICATE-----\\r\\n-----BEGIN CERTIFICATE-----\\r\\nMIIElg==\\r\\n-----END CERTIFICATE-----\\r\\n-----BEGIN CERTIFICATE-----\\r\\nMIIDV5Adg\\r\\n06O/nVsJ8dWd4=\\r\\n-----END CERTIFICATE-----"),
			PrivateKey:  cdn.GetStrPtr("-----BEGIN RSA PRIVATE KEY-----\nMIIE6\nBNtw==\n-----END RSA PRIVATE KEY-----"),
		},
		CertInfo: &cdn.AddCdnCertInfo{
			Desc: cdn.GetStrPtr("MyCert"),
		},
		Source: cdn.GetStrPtr("volc_cert_center"),
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
