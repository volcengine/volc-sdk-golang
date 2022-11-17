package tls

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func initTestProject(client tls.Client) (testProjectID string, testTopicID string, err error) {
	//新建project
	createResp, err := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: testPrefix + uuid.NewString(),
		Description: "",
		Region:      testRegion,
	})
	if err != nil {
		return "", "", err
	}
	testProjectID = createResp.ProjectID

	// 新建topic
	// TopicName Description字段规范参考api文档
	createTopicRequest := &tls.CreateTopicRequest{
		ProjectID:   testProjectID,
		TopicName:   testPrefix + uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic desc",
	}
	topic, err := client.CreateTopic(createTopicRequest)
	testTopicID = topic.TopicID
	if err != nil {
		return "", "", err
	}

	//新建index，开启全文索引和kv索引
	createIndexReq := &tls.CreateIndexRequest{
		TopicID: testTopicID,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
	}
	_, err = client.CreateIndex(createIndexReq)
	if err != nil {
		return "", "", err
	}
	return testProjectID, testTopicID, nil
}

func main() {
	//初始化客户端，配置AccessKeyID,AccessKeySecret,region,securityToken;securityToken可以为空
	client := tls.NewClient(testEndPoint, testAk, testSk, testSessionToken, testRegion)

	var (
		// testProjectID string
		testTopicID string
	)

	_, testTopicID, err := initTestProject(client)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// 创建下载任务
	createDownloadTaskReq := &tls.CreateDownloadTaskRequest{
		TopicID:     testTopicID,
		TaskName:    testPrefix + uuid.NewString(),
		Query:       "*",
		StartTime:   time.Now().UnixNano()/1e6 - 1000,
		EndTime:     time.Now().UnixNano() / 1e6,
		Compression: "gzip",
		DataFormat:  "json",
		Limit:       100,
		Sort:        "asc",
	}
	createDownloadTaskResp, err := client.CreateDownloadTask(createDownloadTaskReq)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", createDownloadTaskResp)

	// 获取下载任务列表
	describeDownloadTasksReq := &tls.DescribeDownloadTasksRequest{
		TopicID: testTopicID,
	}
	describeDownloadTasksResp, err := client.DescribeDownloadTasks(describeDownloadTasksReq)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", describeDownloadTasksResp)

	var taskId string
	if len(describeDownloadTasksResp.Tasks) > 0 {
		taskId = describeDownloadTasksResp.Tasks[0].TaskId
	}

	// 获取任务下载链接
	describeDownloadUrlReq := &tls.DescribeDownloadUrlRequest{
		TaskId: taskId,
	}
	describeDownloadUrlResp, err := client.DescribeDownloadUrl(describeDownloadUrlReq)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", describeDownloadUrlResp)
}
