package diff

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/diff"
	"testing"
)

const (
	testAk       = "your ak"
	testSk       = "your sk"
	appId  int64 = 12345
)

func defaultInstance() *diff.Diff {
	instance := diff.DefaultInstance

	// insert necessary params
	instance.Client.SetAccessKey(testAk)
	instance.Client.SetSecretKey(testSk)
	return instance
}

func printResult(status int, err error, response interface{}) {
	// print httpcode (and err)
	fmt.Println("status:", status, "error:", err)

	// print response
	b, _ := json.Marshal(response)
	fmt.Println("json string result:", string(b))
}

func TestCreateService(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &diff.CreateServiceReq{
		Name:    "your service name",
		Comment: "your service comment",
	}

	// do http query
	response, status, err := instance.CreateService(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestValidate(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &diff.ValidateReq{
		OldUrl:   "your old package url",
		PatchUrl: "your patch url",
	}

	// do http query
	response, status, err := instance.Validate(request)

	// printResult
	printResult(status, err, response)
}

func TestGenByPkg(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &diff.GenByPkgReq{
		ServiceId:  12345,
		Alg:        "wp",
		OldUrl:     "your old package url",
		OldVersion: "1.0.0",
		NewUrl:     "your new pacakge url",
		NewVersion: "1.1.0",
	}

	// do http query
	response, status, err := instance.GenByPkg(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestGenByCount(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &diff.GenByCountReq{
		ServiceId:  12345,
		Alg:        "wp",
		NewUrl:     "your new pacakge url",
		NewVersion: "1.1.0",
		Count:      1,
	}

	// do http query
	response, status, err := instance.GenByCount(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestGenByVersion(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &diff.GenByVersionReq{
		ServiceId:   12345,
		Alg:         "wp",
		NewUrl:      "your new pacakge url",
		NewVersion:  "1.1.0",
		OldVersions: []string{"1.0.0"},
	}

	// do http query
	response, status, err := instance.GenByVersion(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestCheckResponse(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &diff.CheckResponseReq{
		TaskId: 12345,
	}

	// do http query
	response, result, status, err := instance.CheckResponse(request)

	// printResult
	printResult(status, err, response)

	if result == nil {
		return
	}

	// 打印实际的任务结果
	switch response.Data.Api {
	case 0:
		if actualResult, ok := result.(*diff.GenResult); ok {
			b, _ := json.Marshal(actualResult)
			fmt.Println("actual gen patch data:", string(b))
		}
	case 1:
		if actualResult, ok := result.(*diff.ValidateResult); ok {
			b, _ := json.Marshal(actualResult)
			fmt.Println("actual validate data:", string(b))
		}
	}
}
