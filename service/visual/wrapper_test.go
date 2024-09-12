package visual

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/sts"

	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

const (
	testAk = "ak"
	testSk = "sk"
)

func TestVisual_CVProcess(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key":     "",
		"image_url":   []string{""},
		"video_url":   "",
		"ref_img_url": []string{""},
	}

	resp, status, err := DefaultInstance.CVProcess(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVSubmitTask(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key":     "",
		"image_url":   []string{""},
		"video_url":   "",
		"ref_img_url": []string{""},
	}

	resp, status, err := DefaultInstance.CVSubmitTask(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVGetResult(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key": "",
		"task_id": "",
	}

	resp, status, err := DefaultInstance.CVGetResult(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVSync2AsyncSubmitTask(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key": "",
		// "binary_data_base64": []string{""},
		"image_urls":      []string{""},
		"positive_prompt": "beautiful",
	}

	resp, status, err := DefaultInstance.CVSync2AsyncSubmitTask(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVSync2AsyncGetResult(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key": "",
		"task_id": "",
	}

	resp, status, err := DefaultInstance.CVSync2AsyncGetResult(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_HairStyleV2(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key": "hair_style",
		//"binary_data_base64": []string{""},
		"image_urls": []string{"https://xxx"},
		"hair_type":  101,
	}

	resp, status, err := DefaultInstance.HairStyleV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_EmotionPortrait(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key": "emotion_portrait",
		//"binary_data_base64": []string{""},
		"image_urls":     []string{"https://xxx"},
		"target_emotion": "jiuwo",
	}

	resp, status, err := DefaultInstance.EmotionPortrait(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestCert_GenH5ProUrl(t *testing.T) {
	// 调用AssumeRole接口获取临时ak/sk及token
	sts.DefaultInstance.Client.SetAccessKey(testAk)
	sts.DefaultInstance.Client.SetSecretKey(testSk)

	list, status, err := sts.DefaultInstance.AssumeRole(&sts.AssumeRoleRequest{
		DurationSeconds: 43200,
		RoleTrn:         "trn:iam::xxx:role/xxxx",
		RoleSessionName: "xxx",
	})
	fmt.Println(status, err)
	var tmpAk, tmpSk, tmpToken string
	if list != nil && list.Result != nil && list.Result.Credentials != nil {
		tmpAk = list.Result.Credentials.AccessKeyId
		tmpSk = list.Result.Credentials.SecretAccessKey
		tmpToken = list.Result.Credentials.SessionToken
	} else {
		fmt.Println("AssumeRole failed")
		b, _ := json.Marshal(list)
		fmt.Println(string(b))
		return
	}
	// 获取config_id
	configId := "xxx"
	// 获取byted_token
	bytedToken := "xxx"
	// 生成H5ProUrl
	h5ProUrl := fmt.Sprintf("https://h5-v2.kych5.com?accessKeyId=%v&secretAccessKey=%v&sessionToken=%v&configId=%v&bytedToken=%v&lng=%v", tmpAk, tmpSk, tmpToken, configId, bytedToken, "zh-CN")
	fmt.Println("H5认证链接：", h5ProUrl)
}

func TestCert_CertH5Token(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key":      "cert_h5_token",
		"h5_config_id": "xxx",
		"sts_token":    "1",
		"idcard_name":  "xxx",
		"idcard_no":    "xxxx",
	}

	resp, status, err := DefaultInstance.CertH5Token(reqBody)
	fmt.Println(status, err)
	if resp != nil && resp.Data != nil && resp.Data.BytedToken != "" {
		fmt.Println(resp.Data.BytedToken)
	}
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestCert_CertH5ConfigInit(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key": "cert_h5_config_init",
		"h5_config": map[string]interface{}{
			"type":          "3",
			"theme_color":   "rgba(56, 123, 255, 1)",
			"show_guide":    "1",
			"show_result":   "1",
			"enable_record": "1",
			"redirect_url":  "https://www.volcengine.com",
		},
		"liveness_config": map[string]interface{}{
			"ref_source":       "1",
			"liveness_type":    "motion",
			"liveness_timeout": 10,
			"motion_list": []string{
				"0",
				"1",
				"2",
				"3",
			},
			"fixed_motion_list": []string{
				"0",
			},
			"motion_count":       2,
			"max_liveness_trial": 10,
		},
	}

	resp, status, err := DefaultInstance.CertH5ConfigInit(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_FaceFusionMovieV3Async(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//提交任务-请求入参
	reqBody := map[string]interface{}{
		"req_key":           "facefusionmovie_standard_v2",
		"image_url":         "https://xxx,https://yyy",
		"video_url":         "https://ccc",
		"ref_img_url":       "https://xxx,https://yyy",
		"source_similarity": "1",
		"gpen":              1.0,
		"logo_info": map[string]interface{}{
			"add_logo": false,
			"position": 0,
			"language": 0,
			"opacity":  0.3,
		},
		// 智能美肤相关参数
		"enable_face_beautify": false,
		"whitening":            0,
		"dermabrasion":         0,
		"sharpening":           0,
	}

	submitResp, status, err := DefaultInstance.FaceFusionMovieSubmitTaskNew(reqBody)

	if status != 200 {
		fmt.Println(status, err)
		b, _ := json.Marshal(submitResp)
		fmt.Println("submitTask error, response is:", string(b))
	}

	// 提交任务成功后返回的taskId
	taskId := submitResp.Data.TaskId
	// 轮询任务状态
	for {
		//查询任务-请求入参
		queryReqBody := map[string]interface{}{
			"req_key": "facefusionmovie_standard_v2",
			"task_id": taskId,
		}

		queryResp, status, err := DefaultInstance.FaceFusionMovieGetResultNew(queryReqBody)
		if status != 200 {
			fmt.Println(status, err)
			b, _ := json.Marshal(queryResp)
			fmt.Println("queryTask err, response is:", string(b))

			return
		}

		if queryResp.Data.Status == "done" {
			fmt.Println(status, err)
			b, _ := json.Marshal(queryResp)
			fmt.Println("queryTask success, response is:", string(b))

			return
		}

		fmt.Println("taskStatus is not done, wait for a moment. the task's status is:", queryResp.Data.Status)
		time.Sleep(5 * time.Second)
	}
}

func TestVisual_Img2ImgOutpainting(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 使用示例
	reqBody := map[string]interface{}{
		"req_key":            "i2i_outpainting",
		"prompt":             "蓝色的海洋",
		"binary_data_base64": []string{"原图base64"},
		"scale":              7,
		"seed":               -1,
		"steps":              30,
		"strength":           0.8,
		"top":                0.1,
		"bottom":             0.1,
		"left":               1,
		"right":              1,
		"max_height":         1920,
		"max_width":          1920,
	}
	// 使用画布
	//reqBody := map[string]interface{}{
	//	"req_key":            "i2i_outpainting",
	//	"prompt":             "蓝色的海洋",
	//	"binary_data_base64": []string{"延边图base64", "延边图mask"},
	//	"scale":              7,
	//	"seed":               -1,
	//	"steps":              30,
	//	"strength":           0.8,
	//	"max_height":         1920,
	//	"max_width":          1920,
	//}

	resp, status, err := DefaultInstance.Img2ImgOutpainting(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_Img2ImgInpaintingEdit(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := map[string]interface{}{
		"req_key":            "i2i_inpainting_edit",
		"binary_data_base64": []string{"原图", "原图标注后的mask"},
		"custom_prompt":      "", // 写入Prompt，AIGC生成取代的内容
		"scale":              5,
		"seed":               -1,
		"steps":              25,
	}

	resp, status, err := DefaultInstance.Img2ImgInpaintingEdit(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_Img2ImgInpainting(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := map[string]interface{}{
		"req_key":            "i2i_inpainting",
		"binary_data_base64": []string{"原图", "原图标注后的mask"},
		"scale":              7,
		"seed":               0,
		"steps":              30,
		"strength":           0.8,
	}

	resp, status, err := DefaultInstance.Img2ImgInpainting(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_HighAesSmartDrawing(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// 高美感通用v1.1-图生图
	//reqBody := map[string]interface{}{
	//	"req_key":            "high_aes_i2i",
	//	"prompt":             "",
	//	"binary_data_base64": []string{""},
	//	"strength":           0.5,
	//	"seed":               -1,
	//	"scale":              7.5,
	//	"ddim_steps":         20,
	//	"logo_info": map[string]interface{}{
	//		"add_logo": false,
	//		"position": 0,
	//		"language": 0,
	//		"opacity":  0.3,
	//	},
	//}
	// 高美感通用v1.1-图生图
	//reqBody := map[string]interface{}{
	//	"req_key":            "high_aes_anime_i2i",
	//	"prompt":             "",
	//	"binary_data_base64": []string{""},
	//	"strength":           0.5,
	//	"seed":               -1,
	//	"scale":              7.5,
	//	"ddim_steps":         20,
	//	//"use_sq":             false, // 不支持与use_sr同时使用
	//	"use_sr":      true,
	//	"sr_scale":    2.0,
	//	"sr_strength": 0.5,
	//	"logo_info": map[string]interface{}{
	//		"add_logo": false,
	//		"position": 0,
	//		"language": 0,
	//		"opacity":  0.3,
	//	},
	//}
	// 高美感通用v1.2-文生图
	//reqBody := map[string]interface{}{
	//	"req_key":    "high_aes_t2i",
	//	"prompt":     "一个站在海边的美女",
	//	"seed":       -1,
	//	"scale":      7.5,
	//	"ddim_steps": 20,
	//	"logo_info": map[string]interface{}{
	//		"add_logo": false,
	//		"position": 0,
	//		"language": 0,
	//		"opacity":  0.3,
	//	},
	//	//生成图的长宽，取值范围[128-512], 最长边和最短边比，不能超过4倍
	//	"width":  512,
	//	"height": 512,
	//}
	// 高美感动漫v1.3-文生图/图生图
	//reqBody := map[string]interface{}{
	//	"req_key":            "high_aes", // 固定值
	//	"prompt":             "",
	//	"model_version":      "anime_v1.3", // 固定值
	//	"binary_data_base64": []string{""}, // 输入图长边与短边比例大于等于1小于1.77，超过1.77过多有极大概率生成失败
	//	"strength":           0.7,
	//	"seed":               -1,
	//	"scale":              7.0,
	//	"ddim_steps":         20,
	//	"logo_info": map[string]interface{}{
	//		"add_logo": false,
	//		"position": 0,
	//		"language": 0,
	//		"opacity":  0.3,
	//	},
	//	//生成图的长宽，取值范围[576,1728],总像素数<=1088*1088
	//	"width":  1024,
	//	"height": 1024,
	//}
	// 高美感动漫v1.3-文生图/图生图
	reqBody := map[string]interface{}{
		"req_key":            "high_aes", // 固定值
		"prompt":             "",
		"model_version":      "general_v1.3", // 固定值
		"binary_data_base64": []string{""},
		"seed":               -1,
		"scale":              3.5,
		"ddim_steps":         25,
		"logo_info": map[string]interface{}{
			"add_logo": false,
			"position": 0,
			"language": 0,
			"opacity":  0.3,
		},
		// 生成图的长宽，取值范围[128-512]
		"width":  512,
		"height": 512,
		// 是否进行2倍超分
		"use_sr":  false,
		"sr_seed": -1,
	}
	// 单图写真（pv版）
	//reqBody := map[string]interface{}{
	//	"req_key": "img2img_photoverse_american_comics", // 13种选择
	//	// 输入图片数组（仅支持一张）, 二选一，binary_data_base64优先生效
	//	"binary_data_base64": []string{""},
	//	//"image_urls":         []string{""},
	//}
	resp, status, err := DefaultInstance.HighAesSmartDrawing(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_FaceSwapAI(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := map[string]interface{}{
		"req_key": "faceswap_ai",
		//"binary_data_base64": []string{"",""},
		"image_urls": []string{"http://xxx", "http://xxx"}, //List长度为2，输入两张图，将第一张图人脸融入第二张图
		"do_risk":    false,
		"gpen":       0.9,
		"skin":       0.1,
		"tou_repair": 1,
	}

	resp, status, err := DefaultInstance.FaceSwapAI(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_FaceSwapV3(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := map[string]interface{}{
		"req_key": "face_swap3_3",
		//"binary_data_base64": []string{"",""},
		"image_urls": []string{"https://xxx", "https://xxx"}, //输入换脸和模板图片链接数组，换脸图在前（最多三张），模板图在后（最多一张）
		"do_risk":    false,
		"face_type":  "l2r",
		"merge_infos": []map[string]int{
			{
				"location":          1,
				"template_location": 1,
			},
		},
		"logo_info": map[string]interface{}{
			"add_logo": true,
			"position": 1,
			"language": 1,
			"opacity":  0.3,
		},
		"source_similarity": "1.0",
		"gpen":              1.0,
	}

	resp, status, err := DefaultInstance.FaceSwapV3(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

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
		VideoUrl:         "https://xxxx",
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
	// 	TaskId: "xxxxx",
	// }

	// resp, status, err := DefaultInstance.FaceFusionMovieGetResult(reqBody)
	// fmt.Println(status, err)
	// b, _ := json.Marshal(resp)
	// fmt.Println(string(b))
}
