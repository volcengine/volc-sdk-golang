package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
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

	submitResp, status, err := visual.DefaultInstance.FaceFusionMovieSubmitTaskNew(reqBody)

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

		queryResp, status, err := visual.DefaultInstance.FaceFusionMovieGetResultNew(queryReqBody)
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
