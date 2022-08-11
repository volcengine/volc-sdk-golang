package live

import (
	"encoding/json"
	"testing"
)

func TestCreateCaster(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Name":              "test1",
		"MainStreamURL":     "vhost",
		"BackupStreamURL":   "test",
		"Provider":          "test",
		"StreamExpiredTime": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateCaster(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListCasters(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Status": 1,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.ListCasters(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterResource(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.GetCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterLayout(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.GetCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterConfig(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.GetCasterConfig(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestStartCaster(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.StartCaster(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestStopCaster(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.StopCaster(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResource(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"Resource": map[string]interface{}{
			"No":             3,
			"Type":           7,
			"URL":            "http://test.example.com/testMp4",
			"VodUseDownload": true,
		},
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := DefaultInstance.CreateCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestRemoveCasterResource(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.RemoveCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCasterLayout(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"LayoutID": 0,
		"Name":     "布局3(横)",
		"Elements": []map[string]interface{}{
			{
				"Name":       "A",
				"W":          200,
				"H":          112.5,
				"X":          0,
				"Y":          56.25,
				"Opacity":    100,
				"ZIndex":     1,
				"ResourceID": 4359,
				"ResourceNO": 1,
			},
			{
				"Name":       "B",
				"W":          200,
				"H":          225,
				"X":          200,
				"Y":          0,
				"Opacity":    100,
				"ZIndex":     2,
				"ResourceID": 4801,
				"ResourceNO": 2,
			},
		},
		"W":             400,
		"H":             225,
		"LayoutTplID":   2,
		"ScreenType":    1,
		"SyncLayoutTpl": false,
		"LayoutTplElements": []map[string]interface{}{
			{
				"Name":    "A",
				"W":       200,
				"H":       112.5,
				"X":       0,
				"Y":       56.25,
				"Opacity": 100,
				"ZIndex":  1,
			},
			{
				"Name":    "B",
				"W":       200,
				"H":       225,
				"X":       200,
				"Y":       0,
				"Opacity": 100,
				"ZIndex":  2,
			},
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterLayout(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"LayoutID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResourceImage(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"Resource": map[string]interface{}{
			"ImageId": "test1",
			"URL":     "http://test.example.com/test1.png",
			"Name":    "LWScreenShot 2022-04-20 at 13",
		},
		"BackResource": map[string]interface{}{
			"ImageId": "test1",
			"URL":     "http://test.example.com/test1.png",
			"Name":    "LWScreenShot 2022-04-20 at 13",
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateCasterResourceImage(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterResourceImage(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteCasterResourceImage(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResourceImages(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"Resource": map[string]interface{}{
			"ID":   0,
			"Name": "1",
			"Type": 7,
			"ImageCollection": []map[string]interface{}{
				{
					"ImageIndexID": 1,
					"ImageID":      "testPic1",
					"Name":         "LWScreenShot 2022-04-20 at 13.54.00.png",
					"URL":          "http://test.example.com/testPic1.png",
				},
			},
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateCasterResourceImages(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCasterConfig(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"PushParams": map[string]interface{}{
			"Direction": 1,
			"Size":      1080,
			"Fps":       30,
			"VBitRate":  6000,
			"Delay":     0,
			"SEI":       "ts",
			"Logo": map[string]interface{}{
				"URL": "",
				"X":   0,
				"Y":   0,
				"W":   0,
				"H":   0,
			},
			"Gop":      2,
			"ABitRate": 256,
			"AChannel": "stereo",
			"Codec":    "h264",
			"Profile":  "high",
			"ASample":  44100,
			"Tune":     "",
			"AILab":    "",
			"Options":  "nal-hrd=cbr:force-cfr=1",
			"Streams": []map[string]interface{}{
				{
					"ID":  1,
					"URL": "rtmp://push.example.com/live/streamname",
				},
			},
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateCasterConfig(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResourceOPED(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"Resource": map[string]interface{}{
			"Name": "1",
			"URL":  "http://test.example.com/open.mp4",
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateCasterResourceOPED(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterResourceOPED(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteCasterResourceOPED(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterResource(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.SwitchCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterResourceImage(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.SwitchCasterResourceImage(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterResourceOPED(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.SwitchCasterResourceOPED(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestStartCasterOPEDArrange(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.StartCasterOPEDArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterLayout(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"LayoutID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.SwitchCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCopyCasterPVWToPGM(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CopyCasterPVWToPGM(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterResourceVodInfo(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"ResourceID": []int{
			0,
			1,
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.GetCasterResourceVodInfo(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterArrange(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.GetCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterArrange(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 5921,
		"Arrange": map[string]interface{}{
			"OpType":        1,
			"ExecType":      1,
			"CountdownTime": 600,
			"ExecTime":      0,
			"ResourceNo":    1,
			"MainBackup":    false,
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.CreateCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCasterArrange(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":  5921,
		"ArrangeID": 48359,
		"Arrange": map[string]interface{}{
			"ExecTime": 1660115568,
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.UpdateCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterArrange(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":  0,
		"ArrangeID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := DefaultInstance.DeleteCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
