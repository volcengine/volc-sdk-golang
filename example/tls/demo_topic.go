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

	// 创建日志项目
	createProjectResp, _ := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: uuid.NewString(),
		Description: "",
		Region:      os.Getenv("VOLCENGINE_REGION"),
	})
	projectID := createProjectResp.ProjectID

	// 创建日志主题
	// 请根据您的需要，填写ProjectId、TopicName、Ttl、ShardCount和Description等参数
	// CreateTopic API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112180
	topic, _ := client.CreateTopic(&tls.CreateTopicRequest{
		ProjectID:   projectID,
		TopicName:   uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic description",
	})
	topicID := topic.TopicID

	// 修改日志主题
	// 请根据您的需要，填写TopicId以及待修改的各项参数
	// ModifyTopic API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112183
	updateTopicName := uuid.NewString()
	_, _ = client.ModifyTopic(&tls.ModifyTopicRequest{
		TopicID:   topicID,
		TopicName: &updateTopicName,
	})

	// 查询指定日志主题信息
	// 请根据您的需要，填写待查询的TopicId
	// DescribeTopic API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112184
	describeTopicResp, _ := client.DescribeTopic(&tls.DescribeTopicRequest{
		TopicID: topicID,
	})
	fmt.Println(describeTopicResp.TopicName)

	// 查询所有日志主题信息
	// 请根据您的需要，填写待查询的ProjectId
	// DescribeTopics API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112185
	describeTopicsResp, _ := client.DescribeTopics(&tls.DescribeTopicsRequest{
		ProjectID: projectID,
	})
	fmt.Println(describeTopicsResp.Total)

	// 删除日志主题
	// 请根据您的需要，填写待删除的TopicId
	// DeleteTopic API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112182
	_, _ = client.DeleteTopic(&tls.DeleteTopicRequest{
		TopicID: topicID,
	})

	// 删除日志项目
	_, _ = client.DeleteProject(&tls.DeleteProjectRequest{
		ProjectID: projectID,
	})
}
