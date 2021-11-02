package demo_live

import (
	"encoding/json"
	"github.com/volcengine/volc-sdk-golang/service/live"
	"net/url"
	"testing"
)

const (
	testAk    = "testAk"
	testSk    = "testSk"
	AccountID = "accountId"
)

func TestUpdateCallback(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateCallback(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeCallback(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeCallback(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCallback(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCallback(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain": "test-sdk-push.byte.com",
		"Type":   "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateDomain(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain": "test-sdk-push.byte.com",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteDomain(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListDomainDetail(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountID": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"PageNum":  1,
		"PageSize": 10,
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.ListDomainDetail(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain": "test-sdk-push.byte.com", //这个是不存在的
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeDomain(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestEnableDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]string{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.EnableDomain(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDisableDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DisableDomain(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestManagerPullPushDomainBind(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"PullDomain": "",
		"PushDomain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ManagerPullPushDomainBind(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateAuthKey(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain":         "",
		"SceneType":      "",
		"AuthDetailList": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateAuthKey(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestEnableAuth(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain":    "",
		"SceneType": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.EnableAuth(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDisableAuth(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain":    "",
		"SceneType": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DisableAuth(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeAuth(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeAuth(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateAllAuthUnderVhost(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost":          "",
		"SceneType":      "",
		"AuthDetailList": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateAllAuthUnderVhost(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestForbidStream(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost":  "",
		"App":    "",
		"Stream": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ForbidStream(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestResumeStream(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost":    "",
		"App":      "",
		"Stream":   "",
		"LastPSM":  "",
		"LastUser": AccountID,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ResumeStream(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ListCert(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"AccountId": AccountID,
		"useway":    "",
		"rsa":       "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateCert(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeCertDetailSecret(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"ChainID": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeCertDetailSecret(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"rsa":     "",
		"ChainID": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateCert(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestBindCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain":     "",
		"CertDomain": "",
		"ChainID":    "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.BindCert(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUnbindCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UnbindCert(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"chainId": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCert(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateReferer(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain":          "",
		"RefererInfoList": []string{""},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateReferer(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteReferer(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]string{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteReferer(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeReferer(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeReferer(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost":  "",
		"App":    "",
		"Status": "",
		"Bucket": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateRecordPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Preset": "",
		"Vhost":  "",
		"App":    "",
		"Status": "",
		"Bucket": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateRecordPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeRecordPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeRecordPresetDetail(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"PresetList": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeRecordPresetDetail(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteRecordPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListVhostRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ListVhostRecordPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateTranscodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateTranscodePreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateTranscodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateTranscodePreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeTranscodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeTranscodePreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeTranscodePresetDetail(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeTranscodePresetDetail(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteTranscodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteTranscodePreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateSnapshotPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateSnapshotPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountID": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.DescribeSnapshotPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))

}

func TestDescribeSnapshotPresetDetail(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountID": []string{AccountID},
		}
	bodyMap := map[string]interface{}{
		"PresetList": []string{""},
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.DescribeSnapshotPresetDetail(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))

}

func TestDeleteSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountId": []string{AccountID},
		}
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteSnapshotPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListVhostSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	query :=
		url.Values{
			"AccountID": []string{AccountID},
		}
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.ListVhostSnapshotPreset(query, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
