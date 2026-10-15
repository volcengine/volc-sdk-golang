# TLS Go SDK 本地可观测 Contract

Producer 和 Consumer 通过实例方法 `GetStats(StatsOptions)` 提供本地、只读的运行健康事实。
SDK 不采集进程 Heap、RSS、goroutine 或 GC；这些指标应由独立的进程监控采集。

`schema_version=2` 标识唯一当前的跨语言快照结构，`fast_v1`/`slow_v1` 标识固定 histogram 桶布局。
这些标识不是可选的 API 版本；采集级别与返回视图均使用同一套接口和快照结构。

## 1. 采集级别与返回视图

`MetricsLevel` 控制 SDK 在业务路径上采集什么，`StatsView` 只控制一次 `GetStats` 返回什么；调用
`GetStats` 不会改变采集级别。

| 采集级别 | 采集内容 |
| --- | --- |
| `MetricsLevelOff` | 默认值。停止业务 counter、queue、worker 和 latency 采集；仅保留最小生命周期状态，不创建 histogram 或观测线程 |
| `MetricsLevelCore` | lifecycle、累计 counter、queue summary、worker summary、健康事实；Producer 还包含 batch-inflight gate 与受预算控制的 managed memory summary |
| `MetricsLevelDetailed` | Core 全部内容，以及 worker/retry detail、latency histogram；Producer 额外提供 managed memory breakdown |

| 返回视图 | 返回 section |
| --- | --- |
| `StatsViewHealth` | header + health facts；Producer 还返回 batch-inflight gate 与 managed memory summary |
| `StatsViewOperational` | Health + Operational；Producer 的两个 memory section 仍为独立顶层 section |
| `StatsViewDebug` | Operational + 当前组件支持的 Detailed sections |

视图满足 `Health ⊂ Operational ⊂ Debug`。请求的 section 未采集时：

- `Partial=true`；
- `MissingSections` 明确列出缺失 section；
- section 字段为 `nil`/JSON absent，不用假 0 填充。

Consumer 不支持 managed memory，因此 capability 中不声明该能力，也不会把它列为缺失。

## 2. 初始化与热切换

默认配置关闭采集。推荐始终从默认配置取得 `Config` 后修改：

```go
cfg := producer.GetDefaultProducerConfig()
cfg.MetricsLevel = producer.MetricsLevelCore
p := producer.NewProducer(cfg)
p.Start()
```

```go
cfg := consumer.GetDefaultConsumerConfig()
cfg.MetricsLevel = consumer.MetricsLevelCore
c, err := consumer.NewConsumer(ctx, cfg, consumeFn)
```

`Config` 新增了公开字段。外部代码若使用没有字段名的 `producer.Config{...}` 或
`consumer.Config{...}`，升级后会编译失败；请改为默认配置后修改或 keyed literal。默认关闭并不能
避免 unkeyed literal 的编译期字段数量检查。

运行中的实例可直接切换采集级别：

```go
if err := p.SetMetricsLevel(producer.MetricsLevelDetailed); err != nil {
    // 处理实例尚未运行、正在关闭或 level 非法
}
defer p.SetMetricsLevel(producer.MetricsLevelOff)
```

Producer 和 Consumer 都只允许在 Running 状态切换。切换到 Off 是真正停止业务指标采集，不只是隐藏输出；
切换到 Detailed 时 histogram 在发布新 level 前惰性初始化。实例一旦进入过 Detailed，固定大小的
histogram 状态会保留供后续重新开启复用，但 Off 期间不更新时间、不写 bucket，也不创建线程。
Consumer 的三个常驻角色（main loop、heartbeat、checkpoint）仅在 goroutine 启动/退出时维护存活状态，
Off 期间也保留这一最小生命周期事实；重新开启后按实际存活角色返回，不把已退出的角色重新视为运行中。
采集级别切换中，只有跨越 Off 边界才会改变 `CollectionGeneration`；Core 与 Detailed 之间切换不会
打断共同的累计 counter。Consumer 完整 Stop 后在同一实例上显式 Start 会开始新的 collection
generation，以便 Delta 和健康判定重新建立基线；InstanceID 与累计 counter 保持连续。

Producer 的批次聚合键包含采集 generation。同一 topic/source/filename 的日志跨越 Off 边界后
可能进入不同批次；频繁 Off/On 切换可能减小批次、增加发送请求数。Core 与 Detailed 之间切换
不会因此拆批。

重新开启时，物理 queue 的 current/HWM 会从当时深度初始化；需要成对更新的 in-flight、held、
callback 和 aggregate gauge 从新 generation 的事件边界开始统计，不补算 Off 期间的历史事件。
`last_success_age`/`last_error_age` 也会在新 generation 清空，避免把关闭前的旧事件误认为当前采集期事实。
因此控制面应把切换后的第一个样本作为基线，不能与关闭前样本拼接。

## 3. 获取快照

```go
snapshot, err := p.GetStats(producer.StatsOptions{
    View: producer.StatsViewOperational,
})
if err != nil {
    return err
}
if snapshot.Operational != nil {
    counters := snapshot.Operational.Counters
    fmt.Println(counters.AcceptedLogsTotal)
    fmt.Println(counters.SendCallsSuccessTotal)
    fmt.Println(snapshot.Operational.IngressQueue)
}
```

```go
snapshot, err := c.GetStats(consumer.StatsOptions{
    View: consumer.StatsViewOperational,
})
if err != nil {
    return err
}
if snapshot.Operational != nil {
    counters := snapshot.Operational.Counters
    fmt.Println(counters.FetchCallsSuccessTotal)
    fmt.Println(counters.ProcessAttemptsFailureTotal)
    fmt.Println(snapshot.Operational.AssignedShards)
}
```

`GetStats` 的约束：

- O(1)，成本与 backlog、topic 和 shard 数量无关；
- 无网络 I/O，不调用用户 callback，不启动后台任务；
- 不改变业务状态；
- 支持并发调用，返回值不引用内部可变 slice；
- Debug 列表有固定上限并带 `Truncated` 标志；
- 不默认输出 topic、shard、group、endpoint 或 request ID。

## 4. Header 与采集边界

每个快照包含：schema、实例/组件/SDK 标识、采集时间、uptime、单调快照序号、lifecycle、
consistency、采集 level、请求 view、capabilities、partial 状态以及采集 generation。

- `MetricsDisabled=true` 表示当前 level 为 Off；此时只返回 header。
- Running/Closing 快照为 weak consistency；后台任务全部结束后 Closed 快照冻结为 final。
- `InstanceID` 是不透明、跨进程重启高概率唯一且在实例生命周期内稳定的标识；变化表示实例重建。
- `CollectionGeneration` 变化表示采集边界发生变化。
- `SnapshotSequence` 仅在同一实例、同一 generation 内用于判断顺序。

## 5. Counter 语义

所有 counter 都是在实例内单调累计的观测值，SDK 不直接输出 QPS、错误率或 5 分钟窗口。Off 期间
不更新，重新开启后保留之前的绝对值；跨越 Off 的 operation 只有在完整归属于同一采集 generation
时才进入需要成对守恒的计数。因此绝对值用于事实展示，任意速率计算都必须在同一 generation 内以
重新开启后的第一个样本为基线。

Producer 明确区分：

- accepted logs：成功进入 Producer 所有权的日志；accepted groups：这些日志在 SDK 聚合阶段实际形成的
  出站 `pb.LogGroup`，不是 `SendLogs` API 调用次数；
- terminal success/failure/dropped logs/groups：最终结果，互斥；
- send call success/failure：一次传输调用结果，不等于日志数；
- send call payload logs/groups：调用实际携带量；
- retries scheduled：进入 SDK retry 调度的次数；
- rejected：未被 Producer 接受，并按原因拆分。

Producer 的同步内存背压 counter 满足：

```text
rejected_memory_full_total
  = rejected_batch_inflight_limit_total
  + rejected_managed_admission_total
```

- `rejected_batch_inflight_limit_total`：`SendLog`/`SendLogs` 在兼容的
  `TotalSizeLnBytes` gate 等待失败；
- `rejected_managed_admission_total`：新 payload 无法取得 managed-memory admission reservation；
- `SendLogs` 在拆分逐条 admission 前失败时整次调用计 1，不虚构被拒绝日志数；
- sender 已接纳后申请序列化/压缩 buffer 失败不属于 API admission rejection，只计入
  `managed_memory.runtime_buffer_rejected_total` 并进入终态结果。

Producer 的 accepted/terminal cohort 绑定日志被接纳时的 collection generation；跨越 Off 边界的旧批次
不会在新 generation 补造 terminal outcome。`send_calls_*` 和其 payload 则以单次 `PutLogs` 调用为观察对象：
只有调用开始与结束都属于同一 generation 才记账，因此启用后发生的旧批次发送仍是有效 call fact，跨边界
未完成的调用不会污染新 generation。

Consumer 明确区分：

- fetch call success/failure 与 fetched logs/groups；
- process attempt success/failure 与 processed-success logs/groups；
- checkpoint/heartbeat call success/failure；
- rebalance 次数。

Consumer 的 `backoff_events_total` 只在实际转入 backoff 状态时累计一次；等待期间的重复调度
以及退出 backoff 都不增加该计数，Off 期间发生的状态转换不会在重新开启后补计。

Producer 的 `last_success_age_ms` 以最近一次 terminal success 为来源，`last_error_age_ms` 以最近一次
terminal failure 为来源。Consumer 的 age 信号覆盖成功/失败完成的 fetch、process callback、checkpoint
和 heartbeat；它表示“组件最近一次成功/错误事实”，不等同于最近一次消费到非空数据。

SDK 不输出 `consumer_lag`，因为当前协议没有同时掌握服务端最新 offset 和 consumer offset。

Queue summary 当前只承诺可精确、O(1) 获得的 `current_items`、可用时的 `capacity_items`、
`high_watermark_items` 和 `bounded`。HWM 作用域是当前 collection generation。SDK 不输出没有
可靠 ownership 来源的 queue bytes；`oldest_pending_age_ms` 也保持 absent，因为在现有队列结构上
精确维护它需要按 backlog 增加状态或在 GetStats 遍历集合，二者都违反本 Contract。

## 6. Delta、QPS 与错误率

组件快照可投影成固定 counter sample，再交给无后台任务的 `observability.DeltaAdapter`：

```go
snapshot, _ := p.GetStats(producer.StatsOptions{View: producer.StatsViewOperational})
sample, ok := snapshot.CounterSample()
if ok {
    result := adapter.Observe(sample)
    if result.Valid {
        qps := result.RatesPerSecond["accepted_logs_total"]
        _ = qps
    }
}
```

Delta 仅在以下条件全部满足时有效：同一 instance、同一 collection generation、sequence 严格递增、
uptime 未回退、采集时间递增且 counter 未下降。否则当前样本成为新基线，`Valid=false` 并通过
`ResetReason` 说明原因；Metrics Off/section absent 以及 counter key 集合变化同样会重新建立基线。

## 7. HealthEvaluator

SDK 快照只输出事实，`EvaluateHealth` 根据相邻样本和调用方策略输出
`unknown/healthy/degraded/unhealthy`：

```go
result := producer.EvaluateHealth(previous, current, policy)
```

首次样本、Metrics Off、实例或 collection generation 变化，以及 new/closing/closed lifecycle 都返回
`unknown` 且 `InsufficientData=true`；建立可比较基线后，明确的 failed lifecycle、critical flag 或关键 worker 缺失为 unhealthy；queue/memory
pressure 达到 policy 阈值为 degraded。Producer 没有输入且没有 backlog、Consumer 没有分片时，
不会仅因缺少成功事件判异常。

## 8. Producer Memory Gates

Producer 有两道公开配置门禁，managed lifecycle budget 内还有一个派生 admission 子门禁：

| 口径 | 来源 | 含义 |
| --- | --- | --- |
| batch-inflight gate | `TotalSizeLnBytes` | 已进入 aggregate/send/retry Batch 的逻辑 protobuf bytes；不含尚在 ingress 的日志 |
| managed payload admission | 由 `MaxProducerMemoryBytes`、`MaxSenderCount`、`MaxBatchSize` 派生 | 给 sender 临时 buffer 保留 headroom 后，新 payload admission 可使用的有效上限 |
| managed lifecycle total | `MaxProducerMemoryBytes` | SDK 明确记账的 payload 与受控运行时 buffer 总上限 |

`BatchInflightMemoryGate` 返回 `limit/used/high_watermark/rejected`。为保持历史 admission 行为，
旧 gate 在等于 limit 时仍可能再接纳并产生有限超调，因此其 pressure 可以略大于 1。

`ManagedMemorySummary` 只表示 SDK 明确拥有且受 Producer 内存预算控制的内存：

```text
used = payload_owned + bounded_runtime_buffer
used <= limit
rejected_total = payload_admission_rejected_total + runtime_buffer_rejected_total
```

其中：

- `limit_bytes` 是 `MaxProducerMemoryBytes`；
- `payload_admission_limit_bytes` 是实际 admission 子门禁，通常小于 `limit_bytes`；
- `payload_admission_rejected_total` 与 `runtime_buffer_rejected_total` 是 limiter reservation
  因容量/超时失败的尝试口径，不等同于 API admission counter；Close/ForceClose 取消等待不计为
  memory-pressure rejection。为保持既有调用错误兼容，managed admission 等待被关闭取消时仍返回
  `TimeoutException`，但统计明确归入 `rejected_closed_total`，不归入 memory-pressure counter。

同一 payload 同一时刻只属于 ingress、aggregate、send、retry 之一。Detailed breakdown 满足：

```text
payload_owned = ingress_owned + aggregate_owned + send_owned + retry_owned
```

受预算控制的压缩/序列化 buffer 计入 `BoundedRuntimeBufferBytes`，不重复计入 payload；不受预算
控制的临时 workspace 不进入 managed memory。

若 Off 期间已存在仍存活的 Producer allocation，运行期启用后 SDK 虽能知道总预算占用，却不能
无损补造其 ingress/aggregate/send/retry ownership。此时 `ManagedMemory`/breakdown 必须 absent，
快照标记 `Partial=true` 并列出 `managed_memory`；这些旧 allocation 释放后自动恢复精确 summary。
进程 Heap/RSS、metadata cache 和不受此预算约束的对象不进入 managed-memory ratio。

Health 同时返回：

- `managed_memory_pressure = managed used / managed lifecycle total limit`，保留原字段语义；
- `memory_pressure = max(batch-inflight pressure, managed payload-admission pressure)`；
- `memory_pressure_source` 固定为 `batch_inflight` 或 `managed_payload_admission`，指出当前
  会限制新 payload admission 的最大压力来源。managed lifecycle-total pressure 仍由
  `managed_memory_pressure` 独立展示，不伪装成当前公式下不可达的 effective source。

`EvaluateHealth` 优先使用有效的 `memory_pressure`；读取旧快照时回退到
`managed_memory_pressure`。

## 9. Latency 语义

Detailed 固定返回以下 histogram；每个字段同时声明 start、end、unit、object 和是否包含 retry：

- Producer：batch queue delay、send call duration、batch delivery duration、callback duration；
- Consumer：fetch call duration、dispatch delay、process callback duration、checkpoint call duration。

短调用使用 `fast_v1` bucket，端到端/回调类调用使用 `slow_v1` bucket。每个 histogram 都返回
`UpperBoundsMs` 的独立副本；`CumulativeBucketCounts` 比有限上界多一个最终 `+Inf` bucket。
bucket contract 不由用户单独配置，避免相同字段在不同实例上具有不同含义。
其中 send-call duration 的一次样本对应一次底层 `PutLogs` 调用；Producer retry queue 发起的每次重试是
独立样本，`IncludesRetry=false` 表示不包含 SDK 层重试等待和其他调用，底层 transport 在该调用内部的
行为仍包含在 start/end 边界内。

## 10. 资源边界与建议

- Off：从未进入 Detailed 的实例无 histogram；无观测 goroutine，counter/gauge mutator 只检查
  level 后返回。为支持运行期启用，保留实例级固定状态和 level gate；曾进入 Detailed 的实例还会
  保留固定大小、停止更新的 histogram state。
- Core：增加固定数量原子计数和边界事件上的更新。Producer 为按 collection generation 隔离在途
  batch，当前实现还会产生约 32 B/log 的瞬时 routing-key 分配；它不 retained，但会随吞吐线性增长。
- Detailed：增加固定 histogram 状态以及阶段边界的时间读取；不会为每条日志保留长期观测对象。
- `GetStats` 每次分配独立 snapshot；建议正常监控每 15 秒或更低频调用，故障诊断再临时提高频率。

实际 CPU、allocation、retained memory 和 `GetStats` 延迟会随吞吐、并发度与采样频率变化；上线前
应使用接近生产的业务负载验证。
