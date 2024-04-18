package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//请求入参
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

	resp, status, err := visual.DefaultInstance.HighAesSmartDrawing(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
