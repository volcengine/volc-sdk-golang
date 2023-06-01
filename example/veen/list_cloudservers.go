package veen

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ListCloudServers(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).ListCloudServers(&veen.ListCloudServersReq{
		Pagination: veen.Pagination{
			Page:    1,
			Limit:   10,
			OrderBy: 2,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
	json.Marshal(resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}
