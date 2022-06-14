package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ListResourceTags(t *testing.T) {
	resp, err := DefaultInstance.ListResourceTags()
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
	assert.NotEmpty(t, resp.Result.ResourceTags)
}
