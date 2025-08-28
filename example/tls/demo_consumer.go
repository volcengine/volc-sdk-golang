package tls

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	log_consumer "github.com/volcengine/volc-sdk-golang/service/tls/consumer"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func launchConsumer() error {
	// 获取消费组的默认配置
	consumerCfg := log_consumer.GetDefaultConsumerConfig()
	// 请配置您的Endpoint、Region、AccessKeyID、AccessKeySecret等基本信息
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
	consumerCfg.Endpoint = os.Getenv("VOLCENGINE_ENDPOINT")
	consumerCfg.Region = os.Getenv("VOLCENGINE_REGION")
	consumerCfg.AccessKeyID = os.Getenv("VOLCENGINE_ACCESS_KEY_ID")
	consumerCfg.AccessKeySecret = os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET")
	// 请配置您的日志项目ID和日志主题ID列表
	consumerCfg.ProjectID = "<YOUR-PROJECT-ID>"
	consumerCfg.TopicIDList = []string{"<YOUR-TOPIC-ID>"}
	// 请配置您的消费组名称（若您未创建过消费组，SDK将默认为您创建指定名称的消费组）
	consumerCfg.ConsumerGroupName = "<CONSUMER-GROUP-NAME>"
	// 请配置消费者名称（同一个消费组的不同消费者需要保证不同名）
	consumerCfg.ConsumerName = "<CONSUMER_NAME>"

	// 定义日志消费函数，您可根据业务需要，自行实现处理LogGroupList的日志消费函数
	// 下面展示了逐个打印消费到的每条日志的每个键值对的代码实现示例
	var handleLogs = func(topicID string, shardID int, l *pb.LogGroupList) {
		fmt.Printf("received new logs from topic: %s, shard: %d\n", topicID, shardID)
		for _, logGroup := range l.LogGroups {
			for _, log := range logGroup.Logs {
				for _, content := range log.Contents {
					fmt.Printf("%s: %s\n", content.Key, content.Value)
				}
			}
		}
	}

	// 创建消费者
	consumer, err := log_consumer.NewConsumer(context.TODO(), consumerCfg, handleLogs)
	if err != nil {
		return errors.Wrap(err, "get new consumer failed.")
	}

	// 启动消费者消费
	if err := consumer.Start(); err != nil {
		return errors.Wrap(err, "start consumer failed.")
	}

	// 等待消费
	<-time.After(time.Second * 60)

	// 停止消费
	consumer.Stop()

	return nil
}

func main() {
	if err := launchConsumer(); err != nil {
		fmt.Println(err.Error())
	}
}
