package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ResetLoginCredential(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).ResetLoginCredential(&veen.ResetLoginCredentialReq{
		InstanceIdentity: "veen20020201213020320353",
		SecretType:       2,
		SecretData:       "Abcd1234&",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}
