package tls

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActiveTlsAccount(t *testing.T) {
	if os.Getenv("LOG_SERVICE_ENDPOINT") == "" ||
		os.Getenv("LOG_SERVICE_AK") == "" ||
		os.Getenv("LOG_SERVICE_SK") == "" ||
		os.Getenv("LOG_SERVICE_REGION") == "" {
		t.Skip()
	}

	client := NewClientWithEnv()

	request := &ActiveTlsAccountRequest{}
	
	response, err := client.ActiveTlsAccount(request)
	
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.RequestID)
}
