package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func BatchDeployCert(t *testing.T) {
	resp, err := DefaultInstance.BatchDeployCert(&cdn.BatchDeployCertRequest{
		CertId: "cert-c195f679cecb4fc5yjt3dd8c54e6c0a2",
		Domain: "www.example.com,img.example.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
