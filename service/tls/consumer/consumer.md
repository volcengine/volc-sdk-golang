# TLS Go Consumer

日志服务提供消费日志的OpenAPI接口ConsumeLogs，支持实时消费采集到服务端的日志数据。
在使用ConsumeLogs接口时，需要按照日志分区维度消费日志数据，消费时自行指定日志主题ID、Shard ID和起始结束游标（Cursor），所以消费日志的进度受限于单个Shard的读写能力，还需要自行维护消费进度，在Shard自动分裂的场景下消费逻辑与流程繁琐。

日志服务通过SDK提供了消费组（ConsumerGroup）功能，支持通过消费组消费日志数据，通过消费组消费时，日志服务会自动均衡各个消费者的消费能力与进度，自动分配Shard，您无需关注消费组的内部调度细节及消费者之间的负载均衡、故障转移等，只需要专注于业务逻辑。

日志服务提供了Consumer异步日志消费库，支持消费同一个日志项目下多个日志主题，具有异步消费、高性能、失败重试、优雅关闭等特性。

关于通过消费组消费数据的基本概念和限制说明等更多信息，请您参阅[通过消费组消费数据](https://www.volcengine.com/docs/6470/1152208)。

## 示例代码

```go
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
```

## Consumer配置

### Consumer Config可配置参数

日志服务SDK支持通过参数配置消费组消费的各种细节配置，例如是否开启顺序消费、Consumer心跳上报时间间隔等，您可以通过这些配置管理消费组的各种消费逻辑。

例如，在上述Go SDK示例代码中，log_consumer.GetDefaultConsumerConfig()函数返回了消费组的默认配置consumerCfg，并向您展示了如何配置您的Endpoint、Region、AccessKeyID、AccessKeySecret等基本信息、日志项目ID和日志主题ID列表、消费组名称和消费者名称。

除此之外，您还可通过consumerCfg的其他字段进行其他自定义配置。consumerCfg支持的字段如下所示。


| 参数                             | 类型           | 示例值   | 描述                                                                                                                                                                      |
|:-------------------------------|:-------------|:------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| MaxFetchLogGroupCount          | int          | 100   | 消费者单次消费日志时，最大获取LogGroup数量，默认为100，最大为1000。                                                                                                                               |
| HeartbeatIntervalInSecond      | int          | 20    | Consumer心跳上报时间间隔，单位为秒。                                                                                                                                                  |
| DataFetchIntervalInMillisecond | int          | 200   | Consumer消费日志时间间隔，单位为毫秒。                                                                                                                                                 |
| FlushCheckpointIntervalSecond  | int          | 5     | Consumer上传消费进度的时间间隔，单位为秒。                                                                                                                                               |
| ConsumeFrom                    | str          | begin | 开始消费时的默认消费位点，与DescribeCursor的From参数一致，仅在该消费者从未上传过消费位点时有效。                                                                                                               |
| OrderedConsume                 | bool         | false | 是否开启顺序消费。开启顺序消费后，消费者会根据Shard分裂的父子关系进行消费。例如Shard0分裂为Shard1与Shard2，而Shard1又分裂为Shard3与Shard4。在开启顺序消费之后，会根据(Shard0) -> (Shard1, Shard2) -> (Shard2, Shard3, Shard4)的顺序进行消费。 |
| LoggerConfig                   | LoggerConfig |       | 日志相关配置                                                                                                                                                                  |

### LoggerConfig可配置参数

| 参数          | 类型     | 示例值   | 描述                                           |
|-------------|--------|-------|----------------------------------------------|
| LogLevel    | string | info  | 设置日志输出级别，支持设置为debug、info、warn和error，默认为info。 |
| LogFileName | string | 50    | 日志文件输出路径，默认输出到stdout。                        |
| IsJsonType  | bool   | false | 是否格式化文件输出格式，默认为false。                        |
| LogMaxSize  | int    | 10    | 单个日志存储数量，默认为10M。                             |
| LogCompress | bool   | false | 日志是否开启压缩。                                    |
