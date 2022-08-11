package demo_live

import (
	"encoding/json"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/live"
)

func TestCreateCaster(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Name":              "test1",
		"MainStreamURL":     "vhost",
		"BackupStreamURL":   "test",
		"Provider":          "test",
		"StreamExpiredTime": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateCaster(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestListCasters(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"Status":          1,
		"CloudCasterID":   0,
		"Name":            "test",
		"CreateTimeStart": "test",
		"CreateTimeEnd":   0,
		"UpdateTimeStart": 0,
		"UpdateTimeEnd":   0,
		"PageNO":          0,
		"PageItemNO":      10,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.ListCasters(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterResource(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.GetCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterLayout(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.GetCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterConfig(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.GetCasterConfig(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestStartCaster(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.StartCaster(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestStopCaster(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.StopCaster(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResource(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"Resource": map[string]interface{}{
			"No":             3,
			"Type":           4,
			"URL":            "http://test.example.com/testMp4",
			"VodUseDownload": true,
		},
	}
	body, _ := json.Marshal(bodyMap)
	resp, statusCode, err := live.DefaultInstance.CreateCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestRemoveCasterResource(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.RemoveCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCasterLayout(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
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
				"ZIndex":     1,
				"ResourceID": 4801,
				"ResourceNO": 2,
			},
		},
		"W":             400,
		"H":             225,
		"LayoutTplID":   1,
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
		},
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterLayout(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"LayoutID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResourceImage(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
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
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateCasterResourceImage(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterResourceImage(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCasterResourceImage(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResourceImages(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"Resource": map[string]interface{}{
			"ID":   0,
			"Name": "1",
			"Type": 7,
			"ImageCollection": []map[string]interface{}{
				{
					"ImageIndexID": 0,
					"ImageID":      "Pic1",
					"Name":         "ScreenShot.png",
					"URL":          "http://test.example.com/testPic1.png",
				},
			},
		},
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateCasterResourceImages(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCasterConfig(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
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
			"Gop":      1,
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
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateCasterConfig(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterResourceOPED(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"Resource": map[string]interface{}{
			"Name": "1",
			"URL":  "http://test.example.com/open.mp4",
		},
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CreateCasterResourceOPED(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterResourceOPED(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCasterResourceOPED(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterResource(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.SwitchCasterResource(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterResourceImage(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.SwitchCasterResourceImage(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterResourceOPED(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.SwitchCasterResourceOPED(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestStartCasterOPEDArrange(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":   0,
		"ResourceID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.StartCasterOPEDArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestSwitchCasterLayout(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
		"LayoutID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.SwitchCasterLayout(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCopyCasterPVWToPGM(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.CopyCasterPVWToPGM(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterResourceVodInfo(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
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
	resp, statusCode, err := live.DefaultInstance.GetCasterResourceVodInfo(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestGetCasterArrange(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.GetCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestCreateCasterArrange(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID": 0,
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
	resp, statusCode, err := live.DefaultInstance.CreateCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestUpdateCasterArrange(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":  0,
		"ArrangeID": 0,
		"Arrange": map[string]interface{}{
			"ExecTime": 1660115568,
		},
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.UpdateCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestDeleteCasterArrange(t *testing.T) {
	live.DefaultInstance.Client.SetAccessKey(testAk)
	live.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"CasterID":  0,
		"ArrangeID": 0,
	}
	body, _ := json.Marshal(bodyMap)
	//打印请求参数
	t.Logf(string(body))
	resp, statusCode, err := live.DefaultInstance.DeleteCasterArrange(nil, string(body))
	if err != nil {
		t.Logf("error occur %v", err)
	}
	res, _ := json.Marshal(resp)
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
