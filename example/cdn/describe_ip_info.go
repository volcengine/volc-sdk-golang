package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DescribeIPInfo(t *testing.T) {
	resp, err := DefaultInstance.DescribeIPInfo(&cdn.DescribeIPInfoRequest{
		IP: "1.1.1.1",
	})
	assert.NoError(t, err)
	rsp, _ := json.Marshal(resp)
	fmt.Printf("%+v", string(rsp))
}
