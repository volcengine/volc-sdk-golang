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

	// 创建日志主题（开启WebTracks功能）
	enableTracking := true
	topic, _ := client.CreateTopic(&tls.CreateTopicRequest{
		ProjectID:      projectID,
		TopicName:      uuid.NewString(),
		Ttl:            30,
		ShardCount:     2,
		Description:    "topic description",
		EnableTracking: &enableTracking,
	})
	topicID := topic.TopicID

	// 创建索引配置
	_, _ = client.CreateIndex(&tls.CreateIndexRequest{
		TopicID: topicID,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
	})

	// WebTracks上传日志
	// WebTracks API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/141803
	webTracksResponse, _ := client.WebTracks(&tls.WebTracksRequest{
		TopicID:      topicID,
		ProjectID:    projectID,
		CompressType: "lz4",
		Source:       "sdk-test",
		Logs: []map[string]string{
			{
				"tap1": "person-A",
			},
		},
	})
	fmt.Printf("%v\n", webTracksResponse)
}
