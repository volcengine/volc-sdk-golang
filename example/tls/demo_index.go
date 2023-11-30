package tls

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 创建日志项目
	createResp, _ := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: uuid.NewString(),
		Description: "",
		Region:      os.Getenv("VOLCENGINE_REGION"),
	})
	projectID := createResp.ProjectID

	// 创建日志主题
	topic, _ := client.CreateTopic(&tls.CreateTopicRequest{
		ProjectID:   projectID,
		TopicName:   uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic description",
	})
	topicID := topic.TopicID

	// 创建索引配置
	// 请根据您的需要，配置FullText全文索引或KeyValue键值索引
	// CreateIndex API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112187
	_, _ = client.CreateIndex(&tls.CreateIndexRequest{
		TopicID: topicID,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
		KeyValue: &[]tls.KeyValueInfo{
			{
				Key: "test1",
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

	// 修改索引配置
	// 请根据您的需要，填写TopicId和待修改的FullText或KeyValue配置
	// ModifyIndex API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112189
	_, _ = client.ModifyIndex(&tls.ModifyIndexRequest{
		TopicID: topicID,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
		KeyValue: nil,
	})

	// 查询索引配置
	// 请根据您的需要，填写待查询的TopicId
	// DescribeIndex API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112190
	describeIndexResp, _ := client.DescribeIndex(&tls.DescribeIndexRequest{
		TopicID: topicID,
	})
	fmt.Println(describeIndexResp.TopicID)

	// 删除索引配置
	// 请根据您的需要，填写待删除索引的TopicId
	// DeleteIndex API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112188
	_, _ = client.DeleteIndex(&tls.DeleteIndexRequest{
		TopicID: topicID,
	})

	// 删除日志主题
	_, _ = client.DeleteTopic(&tls.DeleteTopicRequest{
		TopicID: topicID,
	})

	// 删除日志项目
	_, _ = client.DeleteProject(&tls.DeleteProjectRequest{
		ProjectID: projectID,
	})
}
