package live

import (
	"encoding/json"
	"testing"
)

const (
	testAk = ""
	testSk = ""
)

//done
func TestUpdateCallback(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"MessageType":        "record",
		"Vhost":              "vhost",
		"CallbackDetailList": []interface{}{},
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
	bodyMap := map[string]interface{}{
		"MessageType": "",
		"Domain":      "push-rtmp-testf5go.bytedance.com",
		"App":         "app",
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
	bodyMap := map[string]interface{}{
		"MessageType": "record",
		"Vhost":       "vhost",
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
		"Domain": "push-rtmp-testf5go.bytedance.com",
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
		"Domain": "push-rtmp-testf5go.bytedance.com",
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
	bodyMap := map[string]interface{}{
		"DomainList": []string{"push-rtmp-testf5go.bytedance.com"},
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
		"Domain": "push-rtmp-testf5go.bytedance.com",
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
		"Domain": "push-rtmp-testf5go.bytedance.com",
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
	resp, statusCode, err := DefaultInstance.UpdateAuthKey(nil, string(body))
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
	bodyMap := map[string]interface{}{
		"Domain":    "domain",
		"SceneType": "push",
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
		"Domain": "domain",
		"App":    "app",
		"Stream": "stream",
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
		"Domain": "",
		"App":    "",
		"Stream": "",
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
	bodyMap := map[string]interface{}{
		"Domain":    "",
		"Available": true,
		"Expiring":  false,
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
		"UseWay":   "sign",
		"CertName": "",
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
func TestUpdateCert(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"UseWay":  "sign",
		"ChainID": "xxx",
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
	bodyMap := map[string]interface{}{
		"Domain":     "",
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
	bodyMap := map[string]interface{}{
		"ChainID": "",
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
	bodyMap := map[string]interface{}{
		"Domain": "",
		"App":    "",
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
	bodyMap := map[string]interface{}{
		"Vhost":  "",
		"App":    "",
		"Preset": "preset",
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
	bodyMap := map[string]interface{}{
		"Vhost": "vhost",
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
	bodyMap := map[string]interface{}{
		"Vhost":  "",
		"App":    "",
		"Preset": "",
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

func TestListCommonTransPresetDetail(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string][]string{
		"PresetList": []string{"Preset1", "Preset2"},
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.ListCommonTransPresetDetail(nil, string(body))
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
		"Vhost":  "",
		"App":    "",
		"Preset": "",
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

func TestDescribePullToPushBandwidthData(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"DomainList":      []string{"example.com"},
		"DstAddrTypeList": []string{"live", "Third"},
		"StartTime":       "2021-04-13T00:00:00+08:00",
		"EndTime":         "2021-04-14T00:00:00+08:00",
		"Aggregation":     300,
		"ShowDetail":      true,
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.DescribePullToPushBandwidthData(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))

}

func TestCreateSnapshotAuditPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":              "xx",
		"App":                "xx",
		"Interval":           2.3,
		"Bucket":             "xx",
		"StorageDir":         "ii",
		"CallbackDetailList": []CallbackDetail{{"XX", "http"}},
		"Description":        "xx",
		"StorageStrategy":    0,
		"Label":              []string{},
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.CreateSnapshotAuditPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateSnapshotAuditPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":              "xx",
		"App":                "xx",
		"Interval":           2.3,
		"Bucket":             "xx",
		"StorageDir":         "ii",
		"CallbackDetailList": []CallbackDetail{{"XX", "http"}},
		"Description":        "xx",
		"StorageStrategy":    1,
		"Label":              []string{},
		"Preset":             "PresetName",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.UpdateSnapshotAuditPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteSnapshotAuditPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost":  "xx",
		"App":    "xx",
		"Preset": "PresetName",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.DeleteSnapshotAuditPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListVhostSnapshotAuditPreset(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Vhost": "xx",
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.ListVhostSnapshotAuditPreset(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDescribeLiveAuditData(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"DomainList":  []string{"example.com", "example2.com"},
		"StartTime":   "2021-04-13T00:00:00+08:00",
		"EndTime":     "2021-04-14T00:00:00+08:00",
		"Aggregation": 86400,
		"DetaiField":  []string{"Domain"},
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.DescribeLiveAuditData(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
