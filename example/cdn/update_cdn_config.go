package cdn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func UpdateCdnConfig(t *testing.T) {
	on := bool(true)
	resp, err := DefaultInstance.UpdateCdnConfig(&cdn.UpdateCdnConfigRequest{
		Domain: &exampleDomain,
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
		RuleEngine: &cdn.RuleEngine{
			Rules: []cdn.RERule{
				{
					Name: cdn.GetStrPtr("rule2"),
					Rule: cdn.GetStrPtr("{\"IfBlock\":{\"Condition\":{\"IsGroup\":false,\"Connective\":\"and\",\"Condition\":{\"Object\":\"path\",\"Operator\":\"suffix_match\",\"IgnoreCase\":true,\"Value\":[\"txt\"]}},\"Actions\":[{\"Action\":\"response_header\",\"Stage\":\"client_response\",\"Phase\":6,\"Module\":1,\"Groups\":[{\"Dimension\":\"response_header\",\"GroupParameters\":[{\"Parameters\":[{\"Name\":\"action\",\"Values\":[\"set\"]},{\"Name\":\"header_name\",\"Values\":[\"animal\"]},{\"Name\":\"header_value\",\"Values\":[\"dog\"]}]}]}]}]}}"),
				},
			},
			Switch: &on,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
