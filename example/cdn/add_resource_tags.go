package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func AddResourceTags(t *testing.T) {
	resp, err := DefaultInstance.AddResourceTags(&cdn.AddResourceTagsRequest{
		Resources: []string{exampleDomain},
		ResourceTags: []cdn.ResourceTagEntry{
			{Key: "userKey", Value: "userValue"},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}
