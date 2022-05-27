package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func UpdateCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.UpdateCdnConfig(&cdn.UpdateCdnConfigRequest{
		Domain: exampleDomain,
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
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
