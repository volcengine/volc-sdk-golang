package cdn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DeleteResourceTags(t *testing.T) {
	resp, err := DefaultInstance.DeleteResourceTags(&cdn.DeleteResourceTagsRequest{
		Resources: []string{exampleDomain},
		ResourceTags: []cdn.ResourceTag{
			{Key: cdn.GetStrPtr("userKey"), Value: cdn.GetStrPtr("userValue")},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}
