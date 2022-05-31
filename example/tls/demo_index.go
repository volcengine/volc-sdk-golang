package tls

import (
	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
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

	testProjectId = createResp.ProjectId

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

	//新建index，开启全文索引和kv索引
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

	//新建index，仅开启全文索引
	createIndexReq2 := &tls.CreateIndexRequest{
		TopicID: testTopicId,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
		KeyValue: nil,
	}
	_, _ = client.CreateIndex(createIndexReq2)

	//新建index，仅开启kv索引
	createIndexReq3 := &tls.CreateIndexRequest{
		TopicID:  testTopicId,
		FullText: nil,
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
	_, _ = client.CreateIndex(createIndexReq3)

	//修改索引

	//开启全文和kv
	updateIndexReq := &tls.ModifyIndexRequest{
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

	_, _ = client.ModifyIndex(updateIndexReq)
	//开启全文，关闭kv
	updateIndexReq2 := &tls.ModifyIndexRequest{
		TopicID: testTopicId,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
		KeyValue: nil,
	}
	_, _ = client.ModifyIndex(updateIndexReq2)

	//关闭全文，开启kv
	updateIndexReq3 := &tls.ModifyIndexRequest{
		TopicID:  testTopicId,
		FullText: nil,
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
	_, _ = client.ModifyIndex(updateIndexReq3)

	//查询索引详情
	_, _ = client.DescribeIndex(&tls.DescribeIndexRequest{
		TopicID: testTopicId,
	})

}
