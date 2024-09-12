# TLS Go Producer

Go Producer用于海量数据场景下快速发送日志数据。
Producer具有异步发送、高性能、失败重试、优雅关闭等特性。
日志服务推荐您使用Producer来上报日志。

## 场景说明

Go SDK支持通过以下方式写入日志：

|   写入方式   | 说明                                                                                                                           |
|:--------:|:-----------------------------------------------------------------------------------------------------------------------------|
| PutLogs  | 不推荐。<br/>日志服务支持通过 PutLogs 接口同步请求的方式上传日志。如果选择使用 PutLogs 上传日志，建议您一次性聚合多条日志后调用一次 PutLogs 接口。相对于逐条上传日志的方式，日志聚合后上传可以提升吞吐率并避免触发限流。 |
| Producer | 推荐。<br/>在实际生产环境中，为了提高数据写入效率，建议通过 Go Producer 方式写入日志数据。Producer 用于在海量数据、高并发场景下快速发送日志数据，具有异步发送、高性能、失败重试、优雅关闭等特性。               |

## 示例代码

```go
package tls

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/producer"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	tlsProducerCfg := producer.GetDefaultProducerConfig()
	tlsProducerCfg.Endpoint = os.Getenv("VOLCENGINE_ENDPOINT")
	tlsProducerCfg.Region = os.Getenv("VOLCENGINE_REGION")
	tlsProducerCfg.AccessKeyID = os.Getenv("VOLCENGINE_ACCESS_KEY_ID")
	tlsProducerCfg.AccessKeySecret = os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET")

	// 初始化并启动Producer
	tlsProducer := producer.NewProducer(tlsProducerCfg)
	tlsProducer.Start()

	// 请根据您的需要，填写topicId、source、filename
	topicID := "your-topic-id"
	source := "your-log-source"
	filename := "your-log-filename"

	callback := &ProducerCallback{}

	// 调用Producer SendLog接口，一次提交一条日志
	// 您可根据实际需要，自行定义实现用于业务处理的CallBack，传入SendLog接口，CallBack不能阻塞返回
	err := tlsProducer.SendLog("", topicID, source, filename, &pb.Log{
		Contents: []*pb.LogContent{
			{
				Key:   "key1",
				Value: "value1",
			},
			{
				Key:   "key2",
				Value: "value2",
			},
		},
		Time: time.Now().Unix(),
	}, callback)
	if err != nil {
		// 处理错误
		fmt.Println(err.Error())
	}

	// 调用Producer SendLogs接口，一次提交多条日志
	// 您可根据实际需要，自行定义实现用于业务处理的CallBack，传入SendLogs接口，CallBack不能阻塞返回
	err = tlsProducer.SendLogs("", topicID, source, filename, &pb.LogGroup{
		Source:   source,
		FileName: filename,
		Logs: []*pb.Log{
			{
				Contents: []*pb.LogContent{
					{
						Key:   "key1",
						Value: "value1-1",
					},
					{
						Key:   "key2",
						Value: "value2-1",
					},
				},
				Time: time.Now().Unix(),
			},
			{
				Contents: []*pb.LogContent{
					{
						Key:   "key1",
						Value: "value1-2",
					},
					{
						Key:   "key2",
						Value: "value2-2",
					},
				},
				Time: time.Now().Unix(),
			},
		},
	}, callback)
	if err != nil {
		// 处理错误
		fmt.Println(err.Error())
	}

	// 关闭Producer
	tlsProducer.Close()
}

type ProducerCallback struct {
	successCount int64
	failureCount int64
}

func (sendLog *ProducerCallback) Success(result *producer.Result) {
	atomic.AddInt64(&sendLog.successCount, 1)
}

func (sendLog *ProducerCallback) Fail(result *producer.Result) {
	atomic.AddInt64(&sendLog.failureCount, 1)
	lastAttempt := result.Attempts[len(result.Attempts)-1]
	fmt.Println(lastAttempt.ErrorMessage)
	fmt.Println(lastAttempt.ErrorCode)
	fmt.Println(lastAttempt.RequestId)
}
```

## Producer配置

### Producer Config可配置参数

| 参数                    | 类型            | 示例值                     | 描述                                                                                                                                                                                                |
|:----------------------|:--------------|:------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| TotalSizeLnBytes      | int64         | 100 * 1024 * 1024       | 单个Producer实例能缓存的日志大小上限，单位为B，默认为100MB。                                                                                                                                                             |
| MaxSenderCount        | int64         | 50                      | 单个Producer能并发的最多goroutine的数量，默认为50，该参数用户可以根据自己实际服务器的性能去配置。                                                                                                                                        |
| MaxBlockSec           | int           | 60                      | 如果Producer可用空间(TotalSizeLnBytes)不足，调用者在Send方法上的最大阻塞时间，默认为60秒；如果超过这个时间后所需空间仍无法得到满足，Send方法会抛出TimeoutException；如果将该值设为0，当所需空间无法得到满足时，Send方法会立即抛出TimeoutException；如果您希望Send方法一直阻塞直到所需空间得到满足，可将该值设为负数。 |
| MaxBatchSize          | int64         | 512 * 1024              | 当一个ProducerBatch中缓存的日志大小大于MaxBatchSize时，该Batch将被发送；默认为512KB，最大可设置成5 MB（SDK会自动将超过5 MB的配置调整为5 MB）。                                                                                                  |
| MaxBatchCount         | int           | 4096                    | 当一个ProducerBatch中缓存的日志条数大于MaxBatchCount时，该Batch将被发送；如果未指定，默认为4096，最大可设置成10000（SDK会自动将超过10000的配置调整为10000）。                                                                                         |
| LingerTime            | time.Duration | 2000 * time.Millisecond | 一个ProducerBatch从创建到可发送的逗留时间，默认为2秒，最小可设置成100毫秒。                                                                                                                                                    |
| Retries               | int           | 10                      | 如果某个ProducerBatch首次发送失败，能够对其重试的次数，默认为10次；如果Retries小于等于0，该ProducerBatch首次发送失败后将直接进入失败队列。                                                                                                           |
| MaxReservedAttempts   | int           | 11                      | 每个ProducerBatch每次被尝试发送都对应着一个Attempt，此参数用来控制返回给用户的Attempt个数，默认只保留最近的11次Attempt信息；该参数越大能让您追溯更多的信息，但同时也会消耗更多的内存。                                                                                     |
| BaseRetryBackoffMs    | int64         | 100                     | 首次重试的退避时间，默认为100毫秒；Producer采取指数退避算法，第N次重试的计划等待时间为BaseRetryBackoffMs * 2^(N-1)。                                                                                                                    |
| MaxRetryBackoffMs     | int64         | 50 * 1000               | 重试的最大退避时间，默认为50秒。                                                                                                                                                                                 |
| NoRetryStatusCodeList | []int         | [400,400]               | 用户配置的不需要重试的错误码列表，当发送日志失败时返回的错误码在列表中，则不会重试；默认包含400，404两个值。                                                                                                                                         |
| LoggerConfig          | LoggerConfig  |                         | 日志相关配置                                                                                                                                                                                            |

### LoggerConfig可配置参数

| 参数          | 类型     | 示例值   | 描述                                           |
|-------------|--------|-------|----------------------------------------------|
| LogLevel    | string | info  | 设置日志输出级别，支持设置为debug、info、warn和error，默认为info。 |
| LogFileName | string | 50    | 日志文件输出路径，默认输出到stdout。                        |
| IsJsonType  | bool   | false | 是否格式化文件输出格式，默认为false。                        |
| LogMaxSize  | int    | 10    | 单个日志存储数量，默认为10M。                             |
| LogCompress | bool   | false | 日志是否开启压缩。                                    |
