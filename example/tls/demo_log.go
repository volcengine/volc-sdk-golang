package tls

import (
	"fmt"
	"os"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 创建日志项目
	createProjectResp, _ := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: "project-name",
		Description: "project-description",
		Region:      os.Getenv("VOLCENGINE_REGION"),
	})
	projectID := createProjectResp.ProjectID

	// 创建日志主题
	createTopicResp, _ := client.CreateTopic(&tls.CreateTopicRequest{
		ProjectID:   projectID,
		TopicName:   "topic-name",
		Ttl:         30,
		Description: "topic-description",
		ShardCount:  2,
	})
	topicID := createTopicResp.TopicID

	// 创建索引配置
	_, _ = client.CreateIndex(&tls.CreateIndexRequest{
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

	// （不推荐）本文档以PutLogs接口同步请求的方式上传日志为例
	// （推荐）在实际生产环境中，为了提高数据写入效率，建议通过Go Producer方式写入日志数据

	// 如果选择使用PutLogs上传日志的方式，建议您一次性聚合多条日志后调用一次PutLogs接口，以提升吞吐率并避免触发限流
	// 请根据您的需要，填写TopicId、Source、FileName和Logs列表，建议您使用lz4压缩
	// PutLogs API的请求参数规范和限制请参阅 https://www.volcengine.com/docs/6470/112191
	_, _ = client.PutLogsV2(&tls.PutLogsV2Request{
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
	time.Sleep(30 * time.Second)

	// 查询分析日志数据
	// 请根据您的需要，填写TopicId、Query、StartTime、EndTime、Limit等参数值
	// SearchLogs API的请求参数规范和限制请参阅 https://www.volcengine.com/docs/6470/112195
	resp, _ := client.SearchLogsV2(&tls.SearchLogsRequest{
		TopicID:   topicID,
		Query:     "*",
		StartTime: 1672502400000,
		EndTime:   1688140800000,
		Limit:     20,
	})
	// 打印SearchLogs接口返回值中的部分基本信息
	// 请根据您的需要，自行处理返回值中的其他信息
	fmt.Println(resp.Status)
	fmt.Println(resp.HitCount)
	fmt.Println(resp.Count)
	fmt.Println(resp.Analysis)

	// 查询Shard
	// 请根据您的需要，填写TopicId等参数
	// DescribeShards API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112197
	_, _ = client.DescribeShards(&tls.DescribeShardsRequest{
		TopicID:    topicID,
		PageNumber: 1,
		PageSize:   10,
	})

	// 获取消费游标
	// 请根据您的需要，填写TopicId、ShardId和From
	// DescribeCursor API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112193
	describeCursorResp, _ := client.DescribeCursor(&tls.DescribeCursorRequest{
		TopicID: topicID,
		ShardID: 0,
		From:    "begin",
	})
	beginCursor := describeCursorResp.Cursor

	// 消费日志
	// 请根据您的需要，填写TopicId、ShardId、Cursor和LogGroupCount等参数
	// ConsumeLogs API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112194
	logGroupCount := 100
	consumeLogsResp, _ := client.ConsumeLogs(&tls.ConsumeLogsRequest{
		TopicID:       topicID,
		ShardID:       0,
		Cursor:        beginCursor,
		LogGroupCount: &logGroupCount,
	})
	handleLogs(topicID, 0, consumeLogsResp.Logs)
}

// 定义日志消费函数，您可根据业务需要，自行实现处理LogGroupList的日志消费函数
// 下面展示了逐个打印消费到的每条日志的每个键值对的代码实现示例
func handleLogs(topicID string, shardID int, l *pb.LogGroupList) {
	fmt.Printf("received new logs from topic: %s, shard: %d\n", topicID, shardID)
	for _, logGroup := range l.LogGroups {
		for _, log := range logGroup.Logs {
			for _, content := range log.Contents {
				fmt.Printf("%s: %s\n", content.Key, content.Value)
			}
		}
	}
}
