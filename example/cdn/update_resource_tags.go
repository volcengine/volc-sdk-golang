package cdn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func UpdateResourceTags(t *testing.T) {
	resp, err := DefaultInstance.UpdateResourceTags(&cdn.UpdateResourceTagsRequest{
		Resources: []string{exampleDomain},
		ResourceTags: []cdn.ResourceTag{
			{Key: cdn.GetStrPtr("userKey"), Value: cdn.GetStrPtr("userValue")},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}
