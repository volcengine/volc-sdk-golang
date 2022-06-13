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

参考tls/demo_producer

## producer配置
https://bytedance.feishu.cn/wiki/wikcnuMCDGlCumJLLL3nAT0PCAb#doxcn4s4u0oeueUUue89SFwewlg
[producer配置](https://bytedance.feishu.cn/wiki/wikcnuMCDGlCumJLLL3nAT0PCAb#doxcn4s4u0oeueUUue89SFwewlg)