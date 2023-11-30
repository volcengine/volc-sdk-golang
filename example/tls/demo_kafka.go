package tls

import (
	"fmt"
	"os"

	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 请填写您的TopicId
	topicID := "your-topic-id"

	// 开启Kafka消费特性
	// 请根据您的需要，填写您的TopicId
	// OpenKafkaConsumer API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/147592
	openKafkaConsumerResponse, _ := client.OpenKafkaConsumer(&tls.OpenKafkaConsumerRequest{
		TopicID: topicID,
	})
	fmt.Printf("%v\n", openKafkaConsumerResponse)

	// 描述Kafka消费者
	// 请根据您的需要，填写您的TopicId
	// DescribeKafkaConsumer API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/147594
	describeKafkaConsumerResponse, _ := client.DescribeKafkaConsumer(&tls.DescribeKafkaConsumerRequest{
		TopicID: topicID,
	})
	fmt.Printf("%v\n", describeKafkaConsumerResponse)

	// 关闭Kafka消费特性
	// 请根据您的需要，填写您的TopicId
	// CloseKafkaConsumer API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/147593
	closeKafkaConsumerResponse, _ := client.CloseKafkaConsumer(&tls.CloseKafkaConsumerRequest{
		TopicID: topicID,
	})
	fmt.Printf("%v\n", closeKafkaConsumerResponse)
}
