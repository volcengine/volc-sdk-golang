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
	// 初始化consumer
	consumerCfg := log_consumer.GetDefaultConsumerConfig()
	consumerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
	consumerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
	consumerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
	consumerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")
	consumerCfg.ProjectID = "<YOUR-PROJECT-ID>"
	consumerCfg.TopicIDList = []string{"<YOUR-TOPIC-ID>"}
	consumerCfg.ConsumerGroupName = "<CONSUMER-GROUP-NAME>"
	consumerCfg.ConsumerName = "<CONSUMER_NAME>"

	// 定义日志消费函数
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

	consumer, err := log_consumer.NewConsumer(context.TODO(), consumerCfg, handleLogs)
	if err != nil {
		return errors.Wrap(err, "get new consumer failed: ")
	}

	// 启动consumer消费
	if err := consumer.Start(); err != nil {
		return errors.Wrap(err, "start consumer failed: ")
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
