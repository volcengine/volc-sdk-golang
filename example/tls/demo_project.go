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

	// 修改project

	// projectName规范参考api文档
	modifyProjectName := testPrefix + uuid.NewString()

	// project desc规范参考api文档
	modifyProjectDesc := "test desc"
	updateProjectReq := &tls.ModifyProjectRequest{
		ProjectID:   testProjectId,
		ProjectName: &modifyProjectName,
		Description: &modifyProjectDesc,
	}

	_, _ = client.ModifyProject(updateProjectReq)

	// 查询project
	_, _ = client.DescribeProject(&tls.DescribeProjectRequest{ProjectID: testProjectId})

	// 批量查询project
	// 全量查询
	_, _ = client.DescribeProjects(&tls.DescribeProjectsRequest{})
	// 分页查询
	_, _ = client.DescribeProjects(&tls.DescribeProjectsRequest{
		PageNumber: 2,
		PageSize:   5,
	})
	// 模糊查询
	_, _ = client.DescribeProjects(&tls.DescribeProjectsRequest{
		ProjectName: "groupb",
	})

	//删除project
	_, _ = client.DeleteProject(&tls.DeleteProjectRequest{ProjectID: testProjectId})
}
