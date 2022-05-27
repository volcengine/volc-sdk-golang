package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func AddCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.AddCdnDomain(&cdn.AddCdnDomainRequest{
		Domain:      operateDomain,
		ServiceType: "web",
		Origin: []cdn.OriginRule{
			{OriginAction: cdn.OriginAction{
				OriginLines: []cdn.OriginLine{
					{
						OriginType:   "primary",
						InstanceType: "ip",
						Address:      "1.1.1.1",
						HttpPort:     "80",
						HttpsPort:    "80",
						Weight:       "100",
					},
				},
			}},
		},
		OriginProtocol: "http",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
