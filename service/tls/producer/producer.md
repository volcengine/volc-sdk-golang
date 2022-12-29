# LOG Go Producer

Producer 专用的异步发送log的类库，具有异步发送、高性能、失败重试、优雅关闭等特性

## 使用步骤

### producer的启动和关闭

```go
        // 获取producer默认配置 具体配置见Config
producerCfg := GetDefaultProducerConfig()
producerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
producerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
producerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
producerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")

producer := NewProducer(producerCfg)
//启动
producer.Start()
//安全关闭，等待producer中缓存的所有的数据全部发送完成以后在关闭producer
producer.Close()
//强制关闭
producer.ForceClose()
```

### 发送日志

参考example/tls/demo_producer

## producer配置

### config
| 参数                    | 类型            | 示例值                                  | 描述                                                                                                                                                                                                                  |
|:----------------------|:--------------|:-------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| TotalSizeLnBytes      | Int64         | 100 * 1024 * 1024                    | 单个 producer 实例能缓存的日志大小上限，单位为b，默认为 100MB。                                                                                                                                                                            |
| MaxSenderCount        | Int64         | 50                                   | 单个producer能并发的最多goroutine的数量，默认为50，该参数用户可以根据自己实际服务器的性能去配置。                                                                                                                                                          |
| MaxBlockSec           | Int           | 60 单位为秒                              | 如果 producer 可用空间(TotalSizeLnBytes)不足，调用者在 send 方法上的最大阻塞时间，默认为 60 秒 如果超过这个时间后所需空间仍无法得到满足，send 方法会抛出TimeoutException。如果将该值设为0，当所需空间无法得到满足时，send 方法会立即抛出 TimeoutException。如果您希望 send 方法一直阻塞直到所需空间得到满足，可将该值设为负数。        |
| MaxBatchSize          | Int64         | 512 * 1024                           | 当一个 ProducerBatch 中缓存的日志大小大于等于 batchSizeThresholdInBytes 时，该 batch 将被发送，默认为 512 KB，最大可设置成 5MB。                                                                                                                      |
| MaxBatchCount         | Int           | 4096                                 | 当一个 ProducerBatch 中缓存的日志条数大于等于 batchCountThreshold 时，该 batch 将被发送，如果未指定，默认为 4096，最大可设置成 40960                                                                                                                       |
| LingerTime            | time.Duration | 2000 * time.Millisecondtime.Duration | 一个 ProducerBatch 从创建到可发送的逗留时间，默认为 2 秒，最小可设置成 100 毫秒。                                                                                                                                                                |
| Retries               | Int           | 10                                   | 如果某个 ProducerBatch 首次发送失败，能够对其重试的次数，默认为 10 次。 如果 retries 小于等于 0，该 ProducerBatch 首次发送失败后将直接进入失败队列                                                                                                                    |
| MaxReservedAttempts   | Int           | 11                                   | 每个 ProducerBatch 每次被尝试发送都对应着一个 Attempt，此参数用来控制返回给用户的 attempt 个数，默认只保留最近的 11 次 attempt 信息。 该参数越大能让您追溯更多的信息，但同时也会消耗更多的内存。                                                                                             |
| BaseRetryBackoffMs    | Int64         | 100                                  | 首次重试的退避时间，默认为 100 毫秒。 Producer 采样指数退避算法，第 N 次重试的计划等待时间为 baseRetryBackoffMs * 2^(N-1)                                                                                                                                |
| MaxRetryBackoffMs     | Int64         | 50 * 1000                            | 重试的最大退避时间，默认为 50 秒。                                                                                                                                                                                                 |
| ShardCount            | Int           | 2                                    | 当且仅当 AdjustShardHashFlag 为 true 时，该参数才生效。此时，producer 会自动将 shardHash 重新分组，分组数量为 buckets。 如果两条数据的 shardHash 不同，它们是无法合并到一起发送的，会降低 producer 吞吐量。将 shardHash 重新分组后，能让数据有更多地机会被批量发送。该参数的取值范围是 [1, 256]，且必须是 2 的整数次幂，默认为 2 |
| NoRetryStatusCodeList | []int         | [400,400]                            | 用户配置的不需要重试的错误码列表，当发送日志失败时返回的错误码在列表中，则不会重试。默认包含400，404两个值。                                                                                                                                                           |
| LoggerConfig          | LoggerConfig  |                                      | 日志相关配置                                                                                                                                                                                                              |

### LoggerConfig

| 参数          | 类型     | 示例值   | 描述                                                               |
|-------------|--------|-------|------------------------------------------------------------------|
| LogLevel    | String | info  | 设置日志输出级别，默认值是Info,consumer中一共有4种日志输出级别，分别为debug,info,warn和error。 |
| LogFileName | String | 50    | 日志文件输出路径，不设置的话默认输出到stdout。                                       |
| IsJsonType  | Bool   | false | 是否格式化文件输出格式，默认为false。                                            |
| LogMaxSize  | Int    | 10    | 单个日志存储数量，默认为10M。                                                 |
| LogCompress | Bool   | false | 日志是否开启压缩。                                                        |
