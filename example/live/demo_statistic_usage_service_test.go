package live

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/live"
)

func TestDescribeLiveMetricTrafficData(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey("your ak")
	live.DefaultInstance.Client.SetSecretKey("your sk")
	bodyMap := live.DescribeLiveMetricTrafficDataReq{
		DomainList:  []string{"example.com"},
		StartTime:   time.Now().Add(-time.Hour).Format(time.RFC3339),
		EndTime:     time.Now().Format(time.RFC3339),
		Aggregation: 300,
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.DescribeLiveMetricTrafficData(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
