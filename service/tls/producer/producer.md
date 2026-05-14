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
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
	tlsProducerCfg := producer.GetDefaultProducerConfig()
	tlsProducerCfg.Endpoint = os.Getenv("VOLCENGINE_ENDPOINT")
	tlsProducerCfg.Region = os.Getenv("VOLCENGINE_REGION")
	tlsProducerCfg.AccessKeyID = os.Getenv("VOLCENGINE_ACCESS_KEY_ID")
	tlsProducerCfg.AccessKeySecret = os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET")
	// 如需使用 API Key 匿名鉴权写入日志，可改为配置 tlsProducerCfg.APIKey。
	// 当前 API Key 匿名鉴权仅支持 Producer 最终发送的 /PutLogs 请求；其他接口仍需要完整 AK/SK。
	// tlsProducerCfg.APIKey = os.Getenv("VOLCENGINE_TLS_API_KEY")

	// 初始化并启动Producer
	tlsProducer := producer.NewProducer(tlsProducerCfg)
	tlsProducer.Start()
	// 如需轮换 API Key，可在运行时更新：
	// tlsProducer.SetAPIKey(os.Getenv("VOLCENGINE_TLS_API_KEY_NEW"))

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
| TotalSizeLnBytes      | int64         | 100 * 1024 * 1024       | Batch 层面已聚合在途数据的兼容上限，单位为B，默认为100MB。                                                                                                                                                            |
| MaxProducerMemoryBytes | int64         | 200 * 1024 * 1024       | 单个 Producer 实例完整生命周期的日志内存预算，单位为B，默认为200MB（默认 TotalSizeLnBytes 的 2 倍）；覆盖接收队列、已聚合 Batch、重试队列 Batch，以及发送时 protobuf marshal/lz4 压缩临时 buffer 的 SDK 侧 reservation。设置为0时按当前 TotalSizeLnBytes 的 2 倍自动派生；显式设置时 SDK 会尽量保留该值。 |
| MaxSenderCount        | int64         | 50                      | 单个Producer能并发的最多goroutine的数量，默认为50，该参数用户可以根据自己实际服务器的性能去配置。                                                                                                                                        |
| MaxBlockSec           | int           | 60                      | 如果 Producer 可用内存预算不足，调用者在 Send 方法上的最大阻塞时间，默认为60秒；如果超过这个时间后所需空间仍无法得到满足，Send 方法会抛出 TimeoutException；如果将该值设为0，当所需空间无法得到满足时，Send 方法会立即抛出 TimeoutException；如果您希望 Send 方法一直阻塞直到可用空间恢复，可将该值设为负数。 |
| BatchQueueSize        | int           | 1000000                 | Producer内部接收队列的缓冲容量。SendLog/SendLogs 会在写入队列前占用 MaxProducerMemoryBytes 预算，再由后台 goroutine 聚合成 Batch 并发送；最小值为100。                                                                                             |
| MaxBatchSize          | int64         | 512 * 1024              | 当一个ProducerBatch中缓存的日志大小大于MaxBatchSize时，该Batch将被发送；默认为512KB，最大可设置成 9.5 MB（SDK会自动将超过 9.5 MB 的配置调整为 9.5 MB）。                                                                                                  |
| MaxBatchCount         | int           | 4096                    | 当一个ProducerBatch中缓存的日志条数大于MaxBatchCount时，该Batch将被发送；如果未指定，默认为4096，最大可设置成32768（SDK会自动将超过32768的配置调整为32768）。                                                                                         |
| LingerTime            | time.Duration | 2000 * time.Millisecond | 一个ProducerBatch从创建到可发送的逗留时间，默认为2秒，最小可设置成100毫秒。                                                                                                                                                    |
| Retries               | int           | 10                      | 如果某个ProducerBatch首次发送失败，能够对其重试的次数，默认为10次；如果Retries小于等于0，该ProducerBatch首次发送失败后将直接进入失败队列。                                                                                                           |
| MaxReservedAttempts   | int           | 11                      | 每个ProducerBatch每次被尝试发送都对应着一个Attempt，此参数用来控制返回给用户的Attempt个数，默认只保留最近的11次Attempt信息；该参数越大能让您追溯更多的信息，但同时也会消耗更多的内存。                                                                                     |
| BaseRetryBackoffMs    | int64         | 1000                    | 首次重试的退避时间，默认为1000毫秒；后续重试会在基础增量上加入随机退避。                                                                                                                    |
| MaxRetryBackoffMs     | int64         | 10 * 1000               | 重试的最大退避时间，默认为10秒。                                                                                                                                                                                 |
| NoRetryStatusCodeList | []int         | [400,403,404]           | 用户配置的不需要重试的错误码列表，当发送日志失败时返回的错误码在列表中，则不会重试；默认包含400、403、404三个值。                                                                                                                                         |
| RetryMode             | RetryMode     | RetryModeLegacyDoubleRetry | 重试归属模式。默认 RetryModeLegacyDoubleRetry 保持兼容，由底层 Client 和 Producer 各自按原逻辑重试；推荐新业务使用 RetryModeProducerManaged，由 Producer 统一托管重试，底层 Client 只做单次请求。 |
| RequestTimeout        | time.Duration | 10 * time.Second        | 单次 HTTP 请求超时时间。仅在设置后生效；RetryModeProducerManaged 下如果未设置，默认 10 秒。该参数不等同于 MaxBlockSec，MaxBlockSec 只控制本地内存/缓存等待。 |
| ClientRetryPolicy     | *tls.RetryPolicy | nil                  | 底层 Client 的重试策略。未设置时使用 Client 默认策略；设置后会传给 Producer 内部创建的 Client。RetryModeProducerManaged 下仍会将 MaxAttempts 归一为 1，避免 Client 和 Producer 双重重试。 |
| FailurePolicy         | FailurePolicy | FailurePolicyRetryThenCallback | 熔断打开后的降级策略。FailFast 表示新日志直接返回 CircuitOpenException；DropWithCallback 表示新日志不入队并通过 callback Fail 通知；默认 RetryThenCallback 保持兼容。 |
| CircuitBreaker        | CircuitBreakerConfig |                         | 持续 429、5xx 或网络失败时的熔断配置。默认关闭；开启后 Producer 会按失败率/连续失败次数打开熔断，并在 OpenDuration 后进入 HalfOpen 探测。 |
| LoggerConfig          | LoggerConfig  |                         | 日志相关配置                                                                                                                                                                                            |

### RetryMode 与底层 Client 的关系

- RetryModeLegacyDoubleRetry 是兼容模式：Producer 保留自身重试，同时底层 Client 使用默认 retry policy 或用户设置的 ClientRetryPolicy。该模式避免升级后行为突变，但用户需要注意实际 HTTP 请求次数可能大于 Producer 的 Retries。
- RetryModeProducerManaged 是推荐模式：Producer 创建 Client 后会将底层 Client 的 MaxAttempts 设置为 1，并由 Producer 统一处理 Retries、BaseRetryBackoffMs、MaxRetryBackoffMs、NoRetryStatusCodeList 和熔断策略；如果同时设置 ClientRetryPolicy，除 MaxAttempts 外的 Client retry policy 字段仍会传入底层 Client。
- RequestTimeout 表示单次 HTTP attempt 的 timeout；Retries 表示 Producer batch 首次发送失败后还能重试的次数，Retries=N 时最多发送 N+1 次。

### ClientRetryPolicy 与底层重试策略

ClientRetryPolicy 控制 Producer 内部 TLS Client 发起单次 PutLogs 请求时的重试策略。Producer 自身的 Retries 控制 Batch 级重试；ClientRetryPolicy 控制一次 PutLogs 请求内部的 HTTP attempt 重试。两者同时生效时，实际 HTTP 请求次数可能被放大。

如需显式控制底层 Client 重试，可设置：

```go
policy := tls.DefaultRetryPolicy()
policy.TotalTimeout = 90 * time.Second
policy.InitialInterval = 500 * time.Millisecond
policy.MaxInterval = 10 * time.Second
policy.Multiplier = 2.0
policy.RandomizationFactor = 0.25
policy.MaxAttempts = 2 // 包含第一次请求；0 表示不限制次数，仅受 TotalTimeout 限制

producerConfig.ClientRetryPolicy = &policy
```

参数含义：

- TotalTimeout：单次 PutLogs 请求含所有底层 Client 重试的总截止时间，默认 90 秒。
- InitialInterval：第一次重试前的基础等待间隔，默认 500 毫秒。
- MaxInterval：单次等待间隔上限，默认 10 秒。
- Multiplier：指数退避增长倍数，默认 2.0。
- RandomizationFactor：随机扰动因子，用于打散多个客户端的重试时间，默认 0.25。
- MaxAttempts：最大尝试次数，包含第一次请求；0 表示不按次数限制，只受 TotalTimeout 限制。

选择建议：

- 希望完全由 Producer 统一管理重试次数、退避和熔断时，使用 RetryModeProducerManaged。该模式会把底层 Client 的 MaxAttempts 归一为 1，避免 Client 和 Producer 双重重试。
- 需要保留底层 Client 重试或做兼容迁移时，使用 RetryModeLegacyDoubleRetry，并通过 ClientRetryPolicy 控制底层 Client 的 TotalTimeout、退避参数和 MaxAttempts。
- 不要只增大 ClientRetryPolicy.TotalTimeout 来应对持续 429/5xx；这会延长 Batch 占用 sender 和内存预算的时间。持续失败场景应优先使用 RetryModeProducerManaged 配合 CircuitBreaker 和 FailurePolicy。

### 熔断与降级

CircuitBreaker 默认关闭。开启后，Producer 会把 429、5xx、网络失败纳入熔断统计；本地资源不足、参数错误、鉴权错误、NoRetryStatusCodeList 命中的错误不会推动熔断。推荐和 RetryModeProducerManaged 一起使用，否则底层 Client 可能先完成自己的重试，Producer 层看到失败的时间会更晚，实际 HTTP 请求次数也更难估算。

熔断配置参数含义：

| 参数 | 默认值 | 含义 | 配置指导 |
|:---|:---|:---|:---|
| Enabled | false | 是否启用 Producer 级熔断。 | 老业务默认关闭以保持兼容；持续 429、5xx、网络失败场景建议开启。 |
| MinRequests | 100 | 失败率判断前需要累计的最小请求样本数。 | 值越小越快打开，但越容易受短抖动影响；高吞吐业务可使用 100～500，低吞吐业务可适当降低。 |
| FailureRatio | 0.8 | 当 failures / requests 达到该比例，且 requests >= MinRequests 时打开熔断。 | 0.8 表示 80% 请求失败才打开；希望更敏感可设 0.5～0.7，希望更保守可设 0.9。 |
| ConsecutiveFailures | 20 | 连续可重试失败达到该值时直接打开熔断，不等待 MinRequests。 | 用于持续 429/网络不可达的快速保护；希望更快保护可设 5～10，避免误判可保留 20 或更高。 |
| OpenDuration | 30s | 熔断打开后保持拒绝新请求的时间，到期后进入 HalfOpen 探测。 | 太短会在后端未恢复时频繁探测，太长会延迟恢复；常用 10～60 秒。 |
| HalfOpenMaxRequests | 5 | HalfOpen 状态允许的最大探测并发，也是成功关闭所需的成功探测数。 | 低吞吐或后端压力敏感场景可设 1～3；高吞吐场景可保留 5 或更高，但探测请求会真实打到服务端。 |

MinRequests、FailureRatio、ConsecutiveFailures 是两条“或”关系的打开条件：

- 连续失败触发：`consecutiveFailures >= ConsecutiveFailures`。
- 失败率触发：`requests >= MinRequests && failures / requests >= FailureRatio`。
- 两个条件任意一个满足就会打开熔断；因此默认值不是“等 100 次里失败 80 次才熔断”，而是“连续失败 20 次会先熔断；如果失败不是连续发生，但累计到 100 个样本后失败率达到 80%，也会熔断”。

这三个参数需要配合设置：

- `ConsecutiveFailures` 用来保护“服务端完全不可用、持续 429、网络完全不通”这类连续失败；如果希望这条快速保护路径生效，通常应小于 `MinRequests`。
- `MinRequests + FailureRatio` 用来保护“不是连续失败，但整体失败率很高”的场景；`MinRequests` 太小容易被短抖动误触发，太大则熔断反应变慢。
- 如果 `ConsecutiveFailures` 设置得很大，持续失败时往往会先被失败率条件打开，连续失败条件的实际价值会降低。
- 如果 `FailureRatio` 很低且 `MinRequests` 很小，短时间偶发错误也可能打开熔断。

熔断状态流转：

- Closed：正常发送；429、5xx、网络失败会计入失败统计，成功会清理连续失败计数。
- Open：在 OpenDuration 内按 FailurePolicy 降级处理新日志；不会把新日志继续放入队列消耗内存。
- HalfOpen：OpenDuration 到期后允许少量探测请求；探测成功达到 HalfOpenMaxRequests 后回到 Closed，任意可重试失败会重新回到 Open。
- NoRetryStatusCodeList 命中的 400、403、404、本地内存不足、CircuitOpenException 不会推动熔断。

建议初始配置：

```go
producerConfig.RetryMode = producer.RetryModeProducerManaged
producerConfig.RequestTimeout = 10 * time.Second
producerConfig.Retries = 3
producerConfig.BaseRetryBackoffMs = 1000
producerConfig.MaxRetryBackoffMs = 10 * 1000
producerConfig.MaxProducerMemoryBytes = 200 * 1024 * 1024
producerConfig.MaxBlockSec = 0
producerConfig.FailurePolicy = producer.FailurePolicyFailFast
producerConfig.CircuitBreaker.Enabled = true
producerConfig.CircuitBreaker.MinRequests = 100
producerConfig.CircuitBreaker.FailureRatio = 0.8
producerConfig.CircuitBreaker.ConsecutiveFailures = 20
producerConfig.CircuitBreaker.OpenDuration = 30 * time.Second
producerConfig.CircuitBreaker.HalfOpenMaxRequests = 5
```

熔断打开后：

- FailurePolicyFailFast：SendLog/SendLogs 直接返回 CircuitOpenException，不进入 Producer 内部队列，不占用 MaxProducerMemoryBytes。
- FailurePolicyDropWithCallback：SendLog/SendLogs 返回 nil，但立即触发 callback Fail，不进入 Producer 内部队列，不占用 MaxProducerMemoryBytes；SendLogs 按日志条数触发失败 callback，每条 log 对应一次 Fail。
- FailurePolicyRetryThenCallback：兼容模式，不在入口直接拒绝新日志；适合偶发抖动，不适合持续 429 或后端长时间不可用。
- HalfOpen 状态下，如果 SendLogs 的日志条数超过剩余探测 permit，已获得 permit 的日志会作为探针入队，超出 permit 的日志按 FailurePolicyFailFast 或 FailurePolicyDropWithCallback 降级处理。
- 已进入重试队列的 Batch 在熔断打开时不会继续打到服务端，会进入失败 callback 并释放已占用内存。

熔断和降级的关系：

- 熔断是判断机制：根据 429、5xx、网络失败把 Producer 从 Closed 切到 Open/HalfOpen，决定是否应该继续向服务端发送。
- 降级是处理策略：当熔断已经 Open 且拒绝新日志时，FailurePolicy 决定 SendLog/SendLogs 是直接返回错误、返回 nil 并走失败 callback，还是保持兼容行为。
- 两者是配合关系，不是两个互相独立的开关。CircuitBreaker 未开启或未打开时，FailurePolicy 不会改变正常发送路径。
- FailFast 和 DropWithCallback 的降级发生在新日志进入 Producer 队列之前，因此不会占用 MaxProducerMemoryBytes；RetryThenCallback 为兼容模式，仍可能继续进入 Producer 内部流程。

配置选择建议：

| 场景 | 推荐配置 | 说明 |
|:---|:---|:---|
| 后端持续 429、5xx 或网络不可达 | RetryModeProducerManaged + CircuitBreaker.Enabled=true + FailurePolicyFailFast + MaxBlockSec=0 | 调用方立刻得到 CircuitOpenException，可在业务侧限流、采样、旁路记录或丢弃，Producer 不继续堆积内存。 |
| 业务不希望 SendLog 返回错误，统一依赖 callback 统计结果 | RetryModeProducerManaged + CircuitBreaker.Enabled=true + FailurePolicyDropWithCallback + MaxBlockSec=0 | SendLog/SendLogs 返回 nil，但通过 callback Fail 标记降级失败；SendLogs 每条 log 触发一次 Fail，适合主链路不能被同步错误打断的场景。 |
| 后端偶发抖动，业务希望尽量发送 | RetryModeProducerManaged + CircuitBreaker.Enabled=true + FailurePolicyRetryThenCallback + MaxBlockSec=1～5 | 允许短时间等待和重试，但持续失败时仍要依赖熔断避免无限堆积。 |
| 后端健康且追求极限吞吐 | 可关闭 CircuitBreaker，或使用较保守阈值，例如 MinRequests=500、FailureRatio=0.9、ConsecutiveFailures=50 | 避免短抖动误触发熔断；仍建议保留 MaxProducerMemoryBytes 作为生命周期预算。 |

使用注意：

- 不要把 429 或 5xx 加入 NoRetryStatusCodeList，否则会绕过重试和熔断保护。
- CircuitOpenException 是 SDK 本地降级错误，不代表已经请求服务端失败。
- 开启熔断后仍应监控 SendLog 返回错误、callback Fail 数、pending callback 数和进程 RSS；MaxProducerMemoryBytes 不是进程 RSS 硬上限。
- callback 不应阻塞。Go Producer 在触发最终成功/失败 callback 前会先释放对应 Batch reservation 和发送临时 buffer，但阻塞 callback 仍会占用 sender goroutine 并影响后续回调处理。

### 功能说明与推荐使用方式

默认配置保持兼容：RetryModeLegacyDoubleRetry 会保留原有 Producer 重试和底层 Client retry policy，FailurePolicyRetryThenCallback 会保留原有最终失败后 callback 的行为，CircuitBreaker 默认关闭。这个模式适合希望无感升级、短期不改变线上行为的用户，但需要注意实际 HTTP 请求次数可能被底层 Client 重试放大。

推荐新业务显式使用 RetryModeProducerManaged。该模式下 Producer 会把底层 Client 的 MaxAttempts 设置为 1，由 Producer 统一处理 Retries、BaseRetryBackoffMs、MaxRetryBackoffMs、NoRetryStatusCodeList、RequestTimeout 和熔断策略。这样重试次数、单次请求超时、失败分类和熔断统计都在 Producer 一层收敛，更容易估算最坏耗时和资源占用。

生产环境持续 429、5xx 或网络失败时，推荐组合为 RetryModeProducerManaged + CircuitBreaker + FailurePolicyFailFast + MaxBlockSec=0。该组合会在失败集中出现时快速打开熔断；熔断打开后，新日志不会进入 Producer 队列，也不会占用 MaxProducerMemoryBytes，而是直接向 SendLog/SendLogs 返回 CircuitOpenException。调用方可以据此走本地降级、限流、采样或旁路记录。

如果业务已经把 callback 作为统一结果通道，推荐将 FailurePolicy 设置为 FailurePolicyDropWithCallback。熔断打开后，SendLog/SendLogs 返回 nil，但日志会立即通过 callback Fail 通知，不进入内部队列、不占用内存；SendLogs 会对每条 log 各触发一次 Fail。该模式适合调用方不希望 Send 方法返回错误打断主链路，但仍需要统计降级失败的场景。

如果业务追求尽可能发送且可以承受短暂阻塞，可以保留 FailurePolicyRetryThenCallback，并将 MaxBlockSec 设置为一个较小正数，例如 1 到 5 秒。该模式适合后端偶发抖动，但不适合持续 429 或长时间不可用场景；持续失败时建议开启熔断，否则重试队列会更长时间持有内存预算。

NoRetryStatusCodeList 默认包含 400、403、404。这些错误通常代表参数、权限或资源问题，继续重试无法恢复，因此不会重试，也不会推动熔断。用户如需覆盖该列表，应确保不要把可恢复的 429、5xx 放入 no-retry 列表，否则会绕过重试和熔断保护。

内存约束建议以 MaxProducerMemoryBytes 为 Producer 生命周期总预算，以 TotalSizeLnBytes 为 Batch 层面在途预算，并保持 TotalSizeLnBytes 不大于 MaxProducerMemoryBytes。默认 MaxProducerMemoryBytes 是 TotalSizeLnBytes 的 2 倍，既能给接收队列、重试队列和发送临时 buffer 留出空间，也避免只把上限设成 100MB 后过早压低吞吐。

典型场景建议：

- 后端持续限流或不可用：RetryModeProducerManaged + CircuitBreaker + FailurePolicyFailFast + MaxBlockSec=0。
- 希望主链路不被 Send 返回错误打断：RetryModeProducerManaged + CircuitBreaker + FailurePolicyDropWithCallback。
- 后端健康且追求吞吐：可以保持熔断关闭或使用较保守阈值，同时调大 TotalSizeLnBytes、MaxProducerMemoryBytes、MaxBatchSize、MaxBatchCount 和 MaxSenderCount。
- 低内存进程或 sidecar：降低 MaxProducerMemoryBytes、MaxBatchSize、MaxBatchCount 和 BatchQueueSize，MaxBlockSec 建议设置为 0 或较小值。
- 依赖最终 callback 做账：关闭时优先调用 Close，不要依赖 ForceClose 触发所有 pending 数据的 callback。

### 关闭语义

- Close 会尽量等待已进入 Producer 的数据发送完成，并执行对应成功或失败 callback。
- ForceClose 表示强制关闭：会停止后台调度并释放接收队列、聚合 Batch、重试队列和待发送任务占用的内存，但不保证这些 pending 数据都会触发 callback。业务如果依赖最终 callback，应优先使用 Close。

## 内存控制与上限估算

### 如何控制内存

Producer 的内存占用主要由三部分组成：尚未被后台聚合处理的接收队列缓存（BatchQueueSize）、已聚合但尚未最终发送完成的 Batch（包含重试队列中的 Batch），以及发送时 protobuf marshal/lz4 压缩产生的临时 buffer。

- 设置 MaxProducerMemoryBytes 控制 Producer 完整生命周期的 SDK 侧日志内存预算。默认值为 TotalSizeLnBytes 的 2 倍，且会保证不小于 TotalSizeLnBytes；显式设置时 SDK 会尽量保留该值，如果小于 TotalSizeLnBytes 会归一化为 2 * TotalSizeLnBytes 并输出 Warn 日志。如需按当前 TotalSizeLnBytes 自动派生预算，可将 MaxProducerMemoryBytes 设为 0。SendLog/SendLogs 在写入内部队列前会先申请预算；成功发送、最终失败、ForceClose 丢弃或 panic 兜底都会释放预算。
- 设置 TotalSizeLnBytes 控制 Batch 层面的兼容“在途数据”上限；建议不大于 MaxProducerMemoryBytes。
- 设置 MaxBlockSec 决定达到 MaxProducerMemoryBytes 或 TotalSizeLnBytes 上限后的行为：0 表示立即返回错误；正数表示最多阻塞 N 秒等待空间；负数表示一直阻塞直到可用空间恢复。
- 设置 BatchQueueSize 控制接收队列的条目上限；队列中的日志也会计入 MaxProducerMemoryBytes，因此该值不会再绕过内存预算。
- 通过 MaxBatchSize / MaxBatchCount / LingerTime 缩短日志在内存中停留的时间：更小的 LingerTime 或更低的 batch 触发阈值会更快落盘发送，降低瞬时在途堆积。
- 在下游不稳定时，Retries / MaxRetryBackoffMs 会影响 Batch 在重试队列中的停留时长，从而影响内存占用曲线。

### 最大内存如何估算

定义：
- S_limit = MaxProducerMemoryBytes（Producer SDK 侧日志生命周期预算）
- S_total = TotalSizeLnBytes（Batch 层面兼容“在途数据”字节上限，近似等于 PutLogs 请求体的 protobuf 大小总和）

粗略上界可按以下方式估算：
- Producer SDK 侧日志载荷、重试持有和发送临时 buffer 的 reservation 上界为：S_limit。
- 实际进程 RSS 峰值仍需要在 S_limit 基础上预留 Go 对象、切片、字符串、HTTP client、goroutine 栈和 runtime 开销。

说明：
- SDK 会为发送临时 buffer 预留一部分 MaxProducerMemoryBytes 预算，因此接收队列和 Batch 可用的 payload 预算会低于 MaxProducerMemoryBytes；默认 2 倍预算可以避免有效 payload 预算低于 TotalSizeLnBytes，降低吞吐被过早压制的风险。发送临时 buffer 申请采用 fail-fast 语义，不受 MaxBlockSec 等待控制；如果当前预算内无法立即申请成功，该 Batch 会失败释放，避免 sender 在持有 payload reservation 时自阻塞。
- MaxProducerMemoryBytes 不是 Go 进程 RSS 的绝对上限；它约束 Producer 生命周期内由日志数据、Batch、重试和发送临时 buffer 产生的 SDK 侧 reservation。生产环境建议在该值基础上预留额外裕量（例如按 1.3～2.0 倍）作为 Go 堆与运行时开销空间。
- 如果 source/fileName/contextFlow 很长或日志字段较多，单条日志的对象开销会增大，建议通过降低 MaxProducerMemoryBytes、MaxBatchSize、MaxBatchCount 或 BatchQueueSize 控制峰值。

### LoggerConfig可配置参数

| 参数          | 类型     | 示例值   | 描述                                           |
|-------------|--------|-------|----------------------------------------------|
| LogLevel    | string | info  | 设置日志输出级别，支持设置为debug、info、warn和error，默认为info。 |
| LogFileName | string | 50    | 日志文件输出路径，默认输出到stdout。                        |
| IsJsonType  | bool   | false | 是否格式化文件输出格式，默认为false。                        |
| LogMaxSize  | int    | 10    | 单个日志存储数量，默认为10M。                             |
| LogCompress | bool   | false | 日志是否开启压缩。                                    |
