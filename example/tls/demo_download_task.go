package tls

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 请填写您的TopicId
	topicID := "your-topic-id"

	// 创建下载任务
	// 请根据您的需要，填写TopicId、TaskName、Query、StartTime、EndTime、Compression、DataFormat、Limit、Sort等参数
	// StartTime/EndTime 为秒级 Unix 时间戳
	// CreateDownloadTask API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/142119
	createDownloadTaskResp, _ := client.CreateDownloadTask(&tls.CreateDownloadTaskRequest{
		TopicID:     topicID,
		TaskName:    uuid.NewString(),
		Query:       "*",
		StartTime:   1672502400,
		EndTime:     1688140800,
		Compression: "gzip",
		DataFormat:  "json",
		Limit:       100,
		Sort:        "asc",
		TaskType:    0,
	})
	fmt.Printf("%v\n", createDownloadTaskResp)

	// 获取下载任务列表
	// 请根据您的需要，填写待查看的TopicId
	// DescribeDownloadTasks API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/142120
	describeDownloadTasksResp, _ := client.DescribeDownloadTasks(&tls.DescribeDownloadTasksRequest{
		TopicID: topicID,
	})
	fmt.Printf("%v\n", describeDownloadTasksResp)

	var taskId string
	if len(describeDownloadTasksResp.Tasks) > 0 {
		taskId = describeDownloadTasksResp.Tasks[0].TaskId
	}

	// 获取任务下载链接
	// 请根据您的需要，填写待下载的TaskId
	// DescribeDownloadUrl API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/142122
	describeDownloadUrlResp, _ := client.DescribeDownloadUrl(&tls.DescribeDownloadUrlRequest{
		TaskId: taskId,
	})
	fmt.Printf("%v\n", describeDownloadUrlResp)
}
