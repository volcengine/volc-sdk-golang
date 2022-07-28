package tls

import (
	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
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

	testProjectID := createResp.ProjectID

	// 新建topic
	// TopicName Description字段规范参考api文档
	createTopicRequest := &tls.CreateTopicRequest{
		ProjectID:   testProjectID,
		TopicName:   testPrefix + uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic desc",
	}
	topic, _ := client.CreateTopic(createTopicRequest)
	testTopicID := topic.TopicID

	//新建index，开启全文索引和kv索引
	createIndexReq := &tls.CreateIndexRequest{
		TopicID: testTopicID,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
		KeyValue: &[]tls.KeyValueInfo{
			{
				Key: "index_key",
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

	// 上传日志

	// 索引类型为kv
	_, _ = client.PutLogs(&tls.PutLogsRequest{
		TopicID:      testTopicID,
		HashKey:      "",
		CompressType: "lz4", //压缩类型，lz4或者none
		LogBody: &pb.LogGroupList{
			LogGroups: []*pb.LogGroup{
				{
					Logs: []*pb.Log{
						{
							//Time: 1630000001, 如果不填默认为当前时间戳
							Contents: []*pb.LogContent{
								{
									Key:   "index_key",
									Value: "hello world",
								},
							},
						},
						{
							//Time: 1630000002,
							Contents: []*pb.LogContent{
								{
									Key:   "index_key",
									Value: "hello world",
								},
							},
						},
					},
				},
			},
		},
	})

	// 索引类型为全文索引
	_, _ = client.PutLogs(&tls.PutLogsRequest{
		TopicID:      testTopicID,
		HashKey:      "",
		CompressType: "lz4", //压缩类型，lz4或者none
		LogBody: &pb.LogGroupList{
			LogGroups: []*pb.LogGroup{
				{
					Logs: []*pb.Log{
						{
							//Time: 1630000001, 如果不填默认为当前时间戳
							Contents: []*pb.LogContent{
								{
									Key:   tls.FullTextIndexKey, //全文索引key固定为 __content__
									Value: "hello world",
								},
							},
						},
					},
				},
			},
		},
	})

	//查询日志

	//检索语法规则（Lucene）
	_, _ = client.SearchLogs(&tls.SearchLogsRequest{
		TopicID:   testTopicID,
		Query:     "*",
		StartTime: 1630000000,
		EndTime:   2630454400,
		Limit:     100,
		HighLight: false,
		Context:   "",
		Sort:      "",
	})

	//SQL语法
	_, _ = client.SearchLogs(&tls.SearchLogsRequest{
		TopicID:   testTopicID,
		Query:     "* | select *",
		StartTime: 1630000000,
		EndTime:   2630454400,
		Limit:     100,
		HighLight: false,
		Context:   "",
		Sort:      "",
	})

	//消费日志
	beginCursorResp, _ := client.DescribeCursor(&tls.DescribeCursorRequest{
		TopicID: testTopicID,
		ShardID: 0,
		From:    "begin", //时间点（Unix时间戳，以秒为单位）或者字符串begin、end。begin对应该分区最早的一条日志，end则对应下一条将要被写入的日志
	})

	beginCursor := beginCursorResp.Cursor

	logGroupCount := 100

	_, _ = client.ConsumeLogs(&tls.ConsumeLogsRequest{
		TopicID:       testTopicID,
		ShardID:       0,
		Cursor:        beginCursor,
		EndCursor:     nil,
		LogGroupCount: &logGroupCount,
		Compression:   nil,
	})

	//查询shard
	_, _ = client.DescribeShards(&tls.DescribeShardsRequest{
		TopicID:    testTopicID,
		PageNumber: 1,
		PageSize:   10,
	})

}
