package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
	"time"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//提交任务-请求入参
	submitReqBody := &model.VideoOverResolutionSubmitTaskV2Request{
		ReqKey: "lens_video_nnsr", // 固定值
		Url:    "",
	}

	submitResp, status, err := visual.DefaultInstance.VideoOverResolutionSubmitTaskV2(submitReqBody)

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
		queryReqBody := &model.VideoOverResolutionQueryTaskV2Request{
			ReqKey: "lens_video_nnsr", // 固定值
			TaskId: taskId,
		}

		queryResp, status, err := visual.DefaultInstance.VideoOverResolutionQueryTaskV2(queryReqBody)
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
