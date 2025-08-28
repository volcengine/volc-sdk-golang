# 日志服务Go SDK

火山引擎日志服务 Go SDK 封装了日志服务的常用接口，您可以通过日志服务 Go SDK 调用服务端 API，实现日志采集、日志检索等功能。

## 快速开始

### 初始化客户端

初始化 Client 实例之后，才可以向 TLS 服务发送请求。初始化时推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免 AccessKey 硬编码引发数据安全风险。

初始化代码如下：

```go
client := NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
    os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))
```

### 示例代码

本文档以日志服务的基本日志采集和检索流程为例，介绍如何使用日志服务 Go SDK 管理日志服务基础资源。创建一个 TLSQuickStart.go 文件，并调用接口分别完成创建 Project、创建 Topic、创建索引、写入日志数据、消费日志和查询日志数据。

详细示例代码如下：

```go
package tls

import (
    "fmt"
    "os"
    "time"

    "github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
    // 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免 AccessKey 硬编码引发数据安全风险。详细说明请参考https://www.volcengine.com/docs/6470/1166455
    // 使用 STS 时，ak 和 sk 均使用临时密钥，且设置 VOLCENGINE_TOKEN；不使用 STS 时，VOLCENGINE_TOKEN 部分传空
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
    client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
       os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

    // 创建日志项目
    // 请根据您的需要，填写ProjectName和可选的Description；请您填写和初始化client时一致的Region；
    // CreateProject API的请求参数规范请参阅https://www.volcengine.com/docs/6470/112174
    createProjectResp, err := client.CreateProject(&tls.CreateProjectRequest{
       ProjectName: "project-name",
       Description: "project-description",
       Region:      os.Getenv("VOLCENGINE_REGION"),
    })
    if err != nil {
       // 处理错误
       fmt.Println(err.Error())
    }
    projectID := createProjectResp.ProjectID

    // 创建日志主题
    // 请根据您的需要，填写ProjectId、TopicName、Ttl、Description、ShardCount等参数值
    // CreateTopic API的请求参数规范请参阅https://www.volcengine.com/docs/6470/112180
	maxSplitShard := int32(50)
    createTopicResp, err := client.CreateTopic(&tls.CreateTopicRequest{
       ProjectID:   projectID,
       TopicName:   "topic-name",
       Ttl:         30,
       Description: "topic-description",
       ShardCount:  2,
	   AutoSplit: true,
	   MaxSplitShard: &maxSplitShard,
    })
    if err != nil {
       // 处理错误
       fmt.Println(err.Error())
    }
    topicID := createTopicResp.TopicID

    // 创建索引配置
    // 请根据您的需要，填写TopicId，开启FullText全文索引或KeyValue键值索引或同时开启二者
    // CreateIndex API的请求参数规范请参阅https://www.volcengine.com/docs/6470/112187
    _, err = client.CreateIndex(&tls.CreateIndexRequest{
       TopicID: topicID,
       FullText: &tls.FullTextInfo{
          Delimiter:      ",-;",
          CaseSensitive:  false,
          IncludeChinese: false,
       },
       KeyValue: &[]tls.KeyValueInfo{
          {
             Key: "key",
             Value: tls.Value{
                ValueType:      "text",
                Delimiter:      ", ?",
                CasSensitive:   false,
                IncludeChinese: false,
                SQLFlag:        false,
             },
          },
       },
    })
    if err != nil {
       // 处理错误
       fmt.Println(err.Error())
    }

    // （不推荐）本文档以 PutLogs 接口同步请求的方式上传日志为例
    // （推荐）在实际生产环境中，为了提高数据写入效率，建议通过 Go Producer 方式写入日志数据
     
    // 如果选择使用PutLogs上传日志的方式，建议您一次性聚合多条日志后调用一次PutLogs接口，以提升吞吐率并避免触发限流
    // 请根据您的需要，填写TopicId、Source、FileName和Logs列表，建议您使用lz4压缩
    // PutLogs API的请求参数规范和限制请参阅https://www.volcengine.com/docs/6470/112191
    _, err = client.PutLogsV2(&tls.PutLogsV2Request{
       TopicID:      topicID,
       CompressType: "lz4",
       Source:       "your-log-source",
       FileName:     "your-log-filename",
       Logs: []tls.Log{
          {
             Contents: []tls.LogContent{
                {
                   Key:   "key1",
                   Value: "value1-1",
                },
                {
                   Key:   "key2",
                   Value: "value2-1",
                },
             },
          },
          {
             Contents: []tls.LogContent{
                {
                   Key:   "key1",
                   Value: "value1-2",
                },
                {
                   Key:   "key2",
                   Value: "value2-2",
                },
             },
          },
       },
    })
    if err != nil {
       // 处理错误
       fmt.Println(err.Error())
    }
    time.Sleep(30 * time.Second)

    // 查询分析日志数据
    // 请根据您的需要，填写TopicId、Query、StartTime、EndTime、Limit等参数值
    // SearchLogs API的请求参数规范和限制请参阅https://www.volcengine.com/docs/6470/112195
    resp, err := client.SearchLogsV2(&tls.SearchLogsRequest{
       TopicID:   topicID,
       Query:     "*",
	   StartTime: 1346457600000,
	   EndTime:   1630454400000,
       Limit:     20,
    })
    if err != nil {
       // 处理错误
       fmt.Println(err.Error())
    }

    // 打印SearchLogs接口返回值中的部分基本信息
    // 请根据您的需要，自行处理返回值中的其他信息
    fmt.Println(resp.Status)
    fmt.Println(resp.HitCount)
    fmt.Println(resp.Count)
    fmt.Println(resp.Analysis)
}
```


## 通过 Producer 上报日志数据

[通过Producer上报日志数据](producer/producer.md)

## 通过 Consumer 消费日志数据

[通过Consumer消费日志数据](consumer/consumer.md)
