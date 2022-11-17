package tls

import (
	"fmt"

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
	var pAllowAnonymousWrite = new(bool)
	*pAllowAnonymousWrite = true
	createTopicRequest := &tls.CreateTopicRequest{
		ProjectID:   testProjectID,
		TopicName:   testPrefix + uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic desc",
		AllowAnonymousWrite: pAllowAnonymousWrite,
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
		testProjectID string
		testTopicID string
	)

	testProjectID, testTopicID, err := initTestProject(client)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// 获取任务下载链接
	webTracksRequest := &tls.WebTracksRequest {
		TopicID: testTopicID,
		ProjectID: testProjectID,
		CompressType: "lz4",
		Source: "sdk-test",
		Logs: []map[string]string {
			{
				"tap1": "person-A",
			},
		},
	}
	webTracksResponse, err := client.WebTracks(webTracksRequest)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", webTracksResponse)
}
