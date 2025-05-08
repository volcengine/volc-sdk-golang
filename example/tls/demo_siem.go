package tls

import (
	"context"
	"crypto/tls"
	"fmt"
	tlsConsumer "github.com/volcengine/volc-sdk-golang/service/tls/consumer"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/siem"
	"log"
	"os"
	"strconv"
	"time"
)

func ConsumerLogsToSyslogClient() error {
	// 获取消费组的默认配置
	consumerCfg := tlsConsumer.GetDefaultConsumerConfig()
	// 请配置您的Endpoint、Region、AccessKeyID、AccessKeySecret等基本信息
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

	// 设置syslog client
	syslogClient := siem.NewSyslogClient()
	// 设置client配置
	syslogClient.Host = os.Getenv("SPLUNK_HOST")
	syslogClient.Port, _ = strconv.Atoi(os.Getenv("SPLUNK_PORT_UDP"))
	// ssl认证请配置, 注意这里的tls不要与日志的tls搞混
	syslogClient.SslSettings = &tls.Config{}

	// 启动client
	err := syslogClient.Connect()
	if err != nil {
		log.Fatalf("syslog client connect error %v:", err)
	}

	// 定义日志消费函数，您可根据业务需要，自行实现处理LogGroupList的日志消费函数
	// 以RFC5424的syslog为例 下面展示了逐个打印消费到的每条日志的每个键值对的代码实现示例
	var handleLogsFunc = func(topicID string, shardID int, l *pb.LogGroupList) {
		for _, logGroup := range l.LogGroups {
			for _, tlsLog := range logGroup.Logs {

				// 如果是RFC3164 请使用 NewSyslogFormatRFC3164
				syslogRFC5424 := siem.NewSyslogFormatRFC5424()

				for _, content := range tlsLog.Contents {
					switch content.Key {
					// 字段改成日志中真实的字段
					case "pri":
						syslogRFC5424.Priority = content.Value
					// 字段改成日志中真实的字段
					case "version":
						syslogRFC5424.Version = content.Value
					// 字段改成日志中真实的字段
					case "timestamp":
						syslogRFC5424.Timestamp, _ = time.Parse(time.RFC3339, content.Value)
					// 字段改成日志中真实的字段
					case "hostname":
						syslogRFC5424.Hostname = content.Value
					// 字段改成日志中真实的字段
					case "appName":
						syslogRFC5424.AppName = content.Value
					// 字段改成日志中真实的字段
					case "pid":
						syslogRFC5424.Pid = content.Value
					// 字段改成日志中真实的字段
					case "msgId":
						syslogRFC5424.MsgId = content.Value
					// 字段改成日志中真实的字段
					case "message":
						syslogRFC5424.Message = content.Value
					}
				}

				syslogRFC5424Format, err := syslogRFC5424.FormatSyslogMessageWithRFC5424()
				if err != nil {
					// 这里处理格式化错误
					log.Printf("NOTICE: syslogRFC5424.FormatSyslogMessageWithRFC5424() failed: %v", err)
				}

				// send失败时会主动重试发送1次
				err = syslogClient.Send(syslogRFC5424Format)
				if err != nil {
					// 这里处理发送失败的逻辑
					return
				}

				fmt.Println("发送成功")
			}
		}
	}

	tlsConsumerIns, err := tlsConsumer.NewConsumer(context.TODO(), consumerCfg, handleLogsFunc)
	if err != nil {
		log.Fatalf("NewConsumer error %v", err)
	}

	// 启动消费者消费
	if err := tlsConsumerIns.Start(); err != nil {
		return fmt.Errorf("start consumer failed, error is %v", err)
	}

	// 根据自己的场景控制消费退出
	time.Sleep(300 * time.Second)

	_ = syslogClient.Close()
	tlsConsumerIns.Stop()

	return nil
}
