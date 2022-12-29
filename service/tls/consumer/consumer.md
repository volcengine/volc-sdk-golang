# LOG Go Consumer

Consumer 异步日志消费库，支持消费同一个project下多个topic。具有异步消费、高性能、失败重试、优雅关闭等特性

## 使用步骤

### consumer的启动和关闭

```go
        // 获取producer默认配置 具体配置见Config
consumerCfg := consumer.GetDefaultConsumerConfig()
consumerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
consumerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
consumerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
consumerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")
consumerCfg.ProjectID = "<YOUR-PROJECT-ID>"
consumerCfg.TopicIDList = []string{"<YOUR-TOPIC-ID>"}
consumerCfg.ConsumerGroupName = "<CONSUMER-GROUP-NAME>"
consumerCfg.ConsumerName = "<CONSUMER_NAME>"
```

### 发送日志

参考example/tls/demo_consumer

## consumer配置

### config

| 参数                    | 类型            | 示例值                                  | 描述                                                                                                                                                                                                                  |
|:----------------------|:--------------|:-------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| MaxFetchLogGroupCount             | Int           | 100   | 消费者单次消费日志时，最大获取LogGroup数量，默认100，最大1000。                                                                                                                                                                            |
| HeartbeatIntervalInSecond         | Int           | 20    | consumer心跳上报时间间隔，以秒为单位。                                                                                                                                                         |
| DataFetchIntervalInMillisecond    | Int           | 200   | consumer消费日志时间间隔，以毫秒为单位。       |
| FlushCheckpointIntervalSecond     | Int           | 5     | consumer上传消费进度的时间间隔，以秒为单位。                                                                                                                   |
| ConsumeFrom                       | String        | begin | 开始消费时的默认消费位点，与DescribeCursor的From参数一致。只有在该消费者没有上传过消费位点的情况下有效。|
| OrderedConsume                    | bool          | false | 是否开启顺序消费。开启顺序消费后，消费者会根据shard分裂的父子关系进行消费。如shard0分裂至shard1与shard2，而shard1又分裂成shard3与shard4。在开启顺序消费之后，会根据(shard0) -> (shard1, shard2) -> (shard2, shard3, shard4) 的顺序进行消费。                                                                                                                                                         |
| LoggerConfig                      | LoggerConfig  |       | 日志相关配置                                                                                                                                                                                                              |

### LoggerConfig

| 参数          | 类型     | 示例值   | 描述                                                               |
|-------------|--------|-------|------------------------------------------------------------------|
| LogLevel    | String | info  | 设置日志输出级别，默认值是Info,consumer中一共有4种日志输出级别，分别为debug,info,warn和error。 |
| LogFileName | String | 50    | 日志文件输出路径，不设置的话默认输出到stdout。                                       |
| IsJsonType  | Bool   | false | 是否格式化文件输出格式，默认为false。                                            |
| LogMaxSize  | Int    | 10    | 单个日志存储数量，默认为10M。                                                 |
| LogCompress | Bool   | false | 日志是否开启压缩。                                                        |
