package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func BatchUpdateCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.BatchUpdateCdnConfig(&cdn.BatchUpdateCdnConfigRequest{
		Domains: []string{"xx1.com", "xx2.com"},
		Origin: []cdn.OriginRule{
			{OriginAction: &cdn.OriginAction{
				OriginLines: []cdn.OriginLine{
					{
						OriginType:   cdn.GetStrPtr("primary"),
						InstanceType: cdn.GetStrPtr("ip"),
						Address:      cdn.GetStrPtr("1.1.1.1"),
						HttpPort:     cdn.GetStrPtr("80"),
						HttpsPort:    cdn.GetStrPtr("80"),
						Weight:       cdn.GetStrPtr("100"),
					},
				},
			}},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
