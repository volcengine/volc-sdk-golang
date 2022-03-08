package live

import (
	"encoding/json"
	"testing"
)

const (
	testAk    = "***REMOVED***"
	testSk    = "***REMOVED***"
)

//done
func TestUpdateCallback(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateCallback(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDescribeCallback(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DescribeCallback(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
func TestDeleteCallback(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteCallback(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestCreateDomain(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "test-sdk-push.byte.com",
		"Type":   "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDeleteDomain(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "test-sdk-push.byte.com",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestListDomainDetail(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"PageNum":  1,
		"PageSize": 10,
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.ListDomainDetail(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDescribeDomain(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain": "", //这个是不存在的
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DescribeDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestEnableDomain(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		//"Domain":"",
		"Domain": "test-sdk-push.byte.com", //这个是不存在的
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.EnableDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDisableDomain(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain": "", //这个是不存在的

	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DisableDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestManagerPullPushDomainBind(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"PullDomain": "",
		"PushDomain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.ManagerPullPushDomainBind(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestUpdateAuthKey(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":         "123",
		"SceneType":      "push",
		"AuthDetailList": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateAuthKey(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestEnableAuth(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":    "123",
		"SceneType": "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.EnableAuth(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDisableAuth(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":    "123",
		"SceneType": "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DisableAuth(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDescribeAuth(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DescribeAuth(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestForbidStream(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":  "123",
		"App":    "",
		"Stream": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.ForbidStream(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestResumeStream(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":    "123",
		"App":      "",
		"Stream":   "",
		"LastPSM":  "",
		"LastUser": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.ResumeStream(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestListCert(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		//"Vhost":  "123",
		//"Domain": "123",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.ListCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestCreateCert(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	bodyMap := map[string]interface{}{
		"AccountId": "",
		"useway":    "",
		"rsa":       "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDescribeCertDetailSecret(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"ChainID": "765c1319a72f49b1b34c9deaa1a8d518",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DescribeCertDetailSecret(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestUpdateCert(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"rsa":     "",
		"ChainID": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestBindCert(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain":     "123",
		"CertDomain": "",
		"ChainID":    "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.BindCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestUnbindCert(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UnbindCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDeleteCert(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"chainId": "123",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestUpdateReferer(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":          "",
		"RefererInfoList": []string{""},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateReferer(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteReferer(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteReferer(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDescribeReferer(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DescribeReferer(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestCreateRecordPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":  "",
		"App":    "",
		"Status": "",
		"Bucket": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestUpdateRecordPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Preset": "",
		"Vhost":  "",
		"App":    "",
		"Status": "",
		"Bucket": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDeleteRecordPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestListVhostRecordPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.ListVhostRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestCreateTranscodePreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateTranscodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestUpdateTranscodePreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateTranscodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDeleteTranscodePreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteTranscodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListVhostTransCodePreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.ListVhostTransCodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestCreateSnapshotPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestUpdateSnapshotPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestDeleteSnapshotPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

//done
func TestListVhostSnapshotPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.ListVhostSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListCommonTransPresetDetail(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.ListCommonTransPresetDetail(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
