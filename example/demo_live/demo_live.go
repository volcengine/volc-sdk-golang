package demo_live

import (
	"encoding/json"
	"github.com/volcengine/volc-sdk-golang/service/live"
	"testing"
)

const (
	testAk = "testAk"
	testSk = "testSk"
)

func TestUpdateCallback(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"MessageType":        "record",
		"Vhost":              "vhost",
		"CallbackDetailList": []interface{}{},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateCallback(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeCallback(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"MessageType": "",
		"Domain":      "domain",
		"App":         "app",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeCallback(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCallback(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"MessageType": "record",
		"Vhost":       "vhost",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCallback(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "domain",
		"Type":   "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "domain",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListDomainDetail(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"PageNum":  1,
		"PageSize": 10,
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.ListDomainDetail(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"DomainList": []string{"domain1", "domain2"},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestEnableDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.EnableDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDisableDomain(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DisableDomain(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestManagerPullPushDomainBind(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"PullDomain": "",
		"PushDomain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ManagerPullPushDomainBind(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateAuthKey(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":    "domain",
		"SceneType": "push",
		"AuthDetailList": []interface{}{map[string]interface{}{
			"EncryptionAlgorithm": "md5",
			"SecretKey":           "xx",
		}},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateAuthKey(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestEnableAuth(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":    "domain",
		"App":       "app",
		"SceneType": "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.EnableAuth(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDisableAuth(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":    "domain",
		"App":       "app",
		"SceneType": "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DisableAuth(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeAuth(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":    "domain",
		"SceneType": "push",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeAuth(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestForbidStream(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "domain",
		"App":    "app",
		"Stream": "stream",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ForbidStream(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestResumeStream(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "",
		"App":    "",
		"Stream": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ResumeStream(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":    "",
		"Available": true,
		"Expiring":  false,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ListCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"UseWay":   "sign",
		"CertName": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeCertDetailSecret(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"ChainID": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeCertDetailSecret(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"UseWay":  "sign",
		"ChainID": "xxx",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestBindCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain":     "",
		"CertDomain": "",
		"ChainID":    "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.BindCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUnbindCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UnbindCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCert(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"ChainID": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCert(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateReferer(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "",
		"App":    "",
		"RefererInfoList": []map[string]interface{}{
			{
				"Key":      "asd",
				"Type":     "xx",
				"Value":    "sad",
				"Priority": 0,
			},
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateReferer(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteReferer(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
		"App":   "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteReferer(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeReferer(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Domain": "",
		"App":    "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DescribeReferer(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":  "vhost",
		"App":    "app",
		"Status": "",
		"Bucket": "bb",
		"RecordTob": []map[string]interface{}{
			{
				"Format":       "hls",
				"Duration":     100,
				"Splice":       -1,
				"RecordObject": "/xx/xx",
			},
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Preset": "preset",
		"Vhost":  "vhost",
		"App":    "app",
		"Status": "",
		"Bucket": "bb",
		"RecordTob": []map[string]interface{}{
			{
				"Format":       "hls",
				"Duration":     100,
				"Splice":       -1,
				"RecordObject": "/xx/xx",
			},
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":  "",
		"App":    "",
		"Preset": "preset",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListVhostRecordPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost": "vhost",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ListVhostRecordPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateTranscodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":      "",
		"App":        "",
		"status":     "",
		"SuffixName": "",
		"Preset":     "",
		"FPS":        30,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateTranscodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateTranscodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":      "",
		"App":        "",
		"status":     "",
		"SuffixName": "",
		"Preset":     "",
		"FPS":        60,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateTranscodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteTranscodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":  "",
		"App":    "",
		"Preset": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteTranscodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListVhostTransCodePreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ListVhostTransCodePreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListCommonTransPresetDetail(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string][]string{
		"PresetList": []string{"Preset1", "Preset2"},
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.ListCommonTransPresetDetail(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":          "",
		"App":            "",
		"Status":         1,
		"Interval":       5,
		"Bucket":         "",
		"SnapshotFormat": "jpeg",
		"SnapshotObject": "xx/xx",
		"StorageDir":     "/xx",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":          "",
		"App":            "",
		"Preset":         "",
		"Status":         1,
		"Interval":       5,
		"Bucket":         "",
		"SnapshotFormat": "jpeg",
		"SnapshotObject": "xx/xx",
		"StorageDir":     "/xx",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost":  "",
		"App":    "",
		"Preset": "",
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListVhostSnapshotPreset(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]string{
		"Vhost": "",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.ListVhostSnapshotPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
