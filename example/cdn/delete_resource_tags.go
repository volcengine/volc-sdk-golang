package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DeleteResourceTags(t *testing.T) {
	resp, err := DefaultInstance.DeleteResourceTags(&cdn.DeleteResourceTagsRequest{
		Resources: []string{exampleDomain},
		ResourceTags: []cdn.ResourceTagEntry{
			{Key: "userKey", Value: "userValue"},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}
