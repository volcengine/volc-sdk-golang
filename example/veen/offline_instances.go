package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func OfflineInstances(t *testing.T) {
	ignoreRunning := true
	resp, err := veen.NewInstance(ak, sk).OfflineInstances(&veen.OfflineInstancesReq{
		InstanceIdentities: []string{"veen20020201213020320353", "veen12922733323321300332"},
		IgnoreRunning:      &ignoreRunning,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}
