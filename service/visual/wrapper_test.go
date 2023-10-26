package visual

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

const (
	testAk = ""
	testSk = ""
)

func TestVisual_FaceCompare(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := map[string]interface{}{
		"req_key": "face_compare", // 固定值
		//"binary_data_base64": []string{"/9xxx"},
		"image_urls": []string{"https://1xxx", "https://2xxx"},
	}

	resp, status, err := DefaultInstance.FaceCompare(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_ConvertPhotoV2(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := &model.ConvertPhotoV2Request{
		ReqKey:           "lens_opr", // 固定值
		BinaryDataBase64: []string{"image_base64"},
		//IsColor: false,
	}

	resp, status, err := DefaultInstance.ConvertPhotoV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_LensVidaVideoSubmitTaskV2(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := &model.LensVidaVideoSubmitTaskV2Request{
		ReqKey:   "lens_vida_video", // 固定值
		Url:      "https://xxxx",
		VidaMode: "vida_simple",
	}

	resp, status, err := DefaultInstance.LensVidaVideoSubmitTaskV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_LensVidaVideoGetResultV2(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := &model.LensVidaVideoGetResultV2Request{
		ReqKey: "lens_vida_video", // 固定值
		TaskId: "xxx",
	}

	resp, status, err := DefaultInstance.LensVidaVideoGetResultV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_FaceFusionMovie(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := &model.FaceFusionMovieRequest{
		ReqKey:             "facefusionmovie_standard", // 固定值
		ImageUrl:           "xxx",
		VideoUrl:           "xxx",
		EnableFaceBeautify: true,
		RefImgUrl:          "xxx",
	}

	resp, status, err := DefaultInstance.FaceFusionMovie(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_FaceFusionMovieAsync(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 提交任务
	reqBody := &model.FaceFusionMovieSubmitTaskRequest{
		ReqKey:           "facefusionmovie_standard", // 固定值
		ImageUrl:         "https://xxx, https://xxx",
		VideoUrl:         "https://xxx",
		RefImgUrl:        "https://xxx, https://xxx",
		SourceSimilarity: "1",
		LogoInfo: &model.FaceFusionMovieLogoInfo{
			AddLogo:  false,
			Position: 0,
			Language: 0,
			Opacity:  0.3,
		},
	}

	resp, status, err := DefaultInstance.FaceFusionMovieSubmitTask(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

	// 查询结果
	// reqBody := &model.FaceFusionMovieGetResultRequest{
	// 	ReqKey: "facefusionmovie_standard", // 固定值
	// 	TaskId: "xxxx",
	// }

	// resp, status, err := DefaultInstance.FaceFusionMovieGetResult(reqBody)
	// fmt.Println(status, err)
	// b, _ := json.Marshal(resp)
	// fmt.Println(string(b))
}
