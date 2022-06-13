package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnConfig(&cdn.DescribeCdnConfigRequest{
		Domain: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.Result.DomainConfig)
	domain := resp.Result.DomainConfig
	fmt.Printf("%+v\n", domain)
}
