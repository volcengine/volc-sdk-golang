package cdn

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeDistrictIspData(t *testing.T) {
	resp, err := DefaultInstance.DescribeDistrictIspData(&cdn.DescribeDistrictIspDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "bandwidth",
		Domain:    exampleDomain,
		Aggregate: cdn.GetStrPtr("disaggregate"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	tmp, _ := json.Marshal(resp)
	fmt.Println(string(tmp))
}
