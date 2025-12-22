package tls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActiveTlsAccount(t *testing.T) {
	client := NewClientWithEnv()

	request := &ActiveTlsAccountRequest{}
	
	response, err := client.ActiveTlsAccount(request)
	
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.RequestID)
}