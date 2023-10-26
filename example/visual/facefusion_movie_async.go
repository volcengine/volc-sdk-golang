package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/visual"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//提交任务-请求入参
	logoInfo := &model.FaceFusionMovieLogoInfo{
		AddLogo:  false,
		Position: 0,
		Language: 0,
		Opacity:  0.3,
	}

	submitReqBody := &model.FaceFusionMovieSubmitTaskRequest{
		ReqKey:           "facefusionmovie_standard", // 固定值
		ImageUrl:         "https://xxx, https://xxx",
		VideoUrl:         "https://xxx",
		RefImgUrl:        "https://xxx, https://xxx",
		SourceSimilarity: "1",
		LogoInfo:         logoInfo,
	}

	submitResp, status, err := visual.DefaultInstance.FaceFusionMovieSubmitTask(submitReqBody)

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
		queryReqBody := &model.FaceFusionMovieGetResultRequest{
			ReqKey: "facefusionmovie_standard", // 固定值
			TaskId: taskId,
		}

		queryResp, status, err := visual.DefaultInstance.FaceFusionMovieGetResult(queryReqBody)
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
