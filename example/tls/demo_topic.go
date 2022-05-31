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

	// 修改topic
	updateTopicName := testPrefix + uuid.NewString()
	modifyTopicRequest := &tls.ModifyTopicRequest{
		TopicID:   testTopicId,
		TopicName: &updateTopicName,
	}

	_, _ = client.ModifyTopic(modifyTopicRequest)

	// 查找topic
	_, _ = client.DescribeTopic(&tls.DescribeTopicRequest{TopicID: testTopicId})

	// 批量查找topic
	// 查询project的全部topic，只填project Id
	_, _ = client.DescribeTopics(&tls.DescribeTopicsRequest{
		ProjectID: testProjectId,
	})
	// 分页查询
	_, _ = client.DescribeTopics(&tls.DescribeTopicsRequest{
		ProjectID:  testProjectId,
		PageNumber: 2,
		PageSize:   5,
	})
	//模糊查询
	_, _ = client.DescribeTopics(&tls.DescribeTopicsRequest{
		ProjectID: testProjectId,
		TopicName: "groupb",
	})

	//删除topic
	_, _ = client.DeleteTopic(&tls.DeleteTopicRequest{TopicID: testTopicId})

}
