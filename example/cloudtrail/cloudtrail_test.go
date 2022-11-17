package cloudtrail

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/cloudtrail"
)

func init() {
	testAk := os.Getenv("TEST_AK")
	testSk := os.Getenv("TEST_SK")
	cloudtrail.DefaultInstance.Client.SetAccessKey(testAk)
	cloudtrail.DefaultInstance.Client.SetSecretKey(testSk)
	cloudtrail.DefaultInstance.Client.SetRetrySettings(&base.RetrySettings{AutoRetry: true})
}

func TestCloudTrail(t *testing.T) {
	resp, statusCode, err := cloudtrail.DefaultInstance.LookupEvents(&cloudtrail.LookupEventsReq{
		NextToken:  "",
		MaxResults: 10,
		StartTime:  0,
		EndTime:    0,
		LookupConditions: []cloudtrail.LookupCondition{
			{
				LookupConditionKey:   "EventSource",
				LookupConditionValue: "iam",
			},
		},
	})
	if err != nil {
		fmt.Println("statusCode:", statusCode, "error:", err)
		return
	}
	if statusCode != http.StatusOK {
		fmt.Println("statusCode:", statusCode)
		return
	}
	r, _ := json.Marshal(resp)
	fmt.Println(string(r))
}
