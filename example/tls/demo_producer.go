package tls

import (
	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/producer"
	"os"
	"time"
)

func main() {
	//初始化客户端，配置AccessKeyID,AccessKeySecret,region,securityToken;securityToken可以为空
	client := tls.NewClient(testEndPoint, testAk, testSk, testSessionToken, testRegion)

	//新建project
	createResp, _ := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: testPrefix + uuid.NewString(),
		Description: "",
		Region:      testRegion,
	})

	testProjectId = createResp.ProjectID

	// 新建topic
	// TopicName Description字段规范参考api文档
	createTopicRequest := &tls.CreateTopicRequest{
		ProjectID:   testProjectId,
		TopicName:   testPrefix + uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic desc",
	}
	topic, _ := client.CreateTopic(createTopicRequest)
	testTopicId = topic.TopicID

	//新建index
	createIndexReq := &tls.CreateIndexRequest{
		TopicID: testTopicId,
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
	}
	_, _ = client.CreateIndex(createIndexReq)

	//初始化producer
	producerCfg := producer.GetDefaultProducerConfig()
	producerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
	producerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
	producerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
	producerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")

	producer := producer.NewProducer(producerCfg)
	producer.Start()

	_ = producer.SendLogs("", testTopicId, "fileSource", "fileName", &pb.LogGroup{
		Source:   "localhost",
		FileName: "logFileName",
		Logs: []*pb.Log{
			{
				Contents: []*pb.LogContent{
					{
						Key:   "key-1",
						Value: "test-message",
					},
					{
						Key:   "key-2",
						Value: "test_message",
					},
				},
				Time: time.Now().Unix(),
			},
		},
	}, nil)

	//关闭producer
	producer.Close()
	//强制关闭producer
	producer.ForceClose()
}
