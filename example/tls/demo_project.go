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
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 创建日志项目
	// 请根据您的需要，填写ProjectName和可选的Description；请您填写和初始化client时一致的Region
	// CreateProject API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112174
	createProjectResp, _ := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: uuid.NewString(),
		Description: "description",
		Region:      os.Getenv("VOLCENGINE_REGION"),
	})
	projectID := createProjectResp.ProjectID

	// 修改日志项目
	// 请根据您的需要，填写ProjectName或Description等参数
	// ModifyProject API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112177
	modifyProjectName := uuid.NewString()
	modifyProjectDesc := "new description"
	updateProjectReq := &tls.ModifyProjectRequest{
		ProjectID:   projectID,
		ProjectName: &modifyProjectName,
		Description: &modifyProjectDesc,
	}
	_, _ = client.ModifyProject(updateProjectReq)

	// 查询指定日志项目信息
	// 请根据您的需要，填写待查询的ProjectId
	// DescribeProject API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112178
	describeProjectResp, _ := client.DescribeProject(&tls.DescribeProjectRequest{
		ProjectID: projectID,
	})
	fmt.Println(describeProjectResp.ProjectName)

	// 查询所有日志项目信息
	// 请根据您的需要，填写ProjectName等参数
	// DescribeProjects API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112179
	describeProjectsResp, _ := client.DescribeProjects(&tls.DescribeProjectsRequest{})
	fmt.Println(describeProjectsResp.Total)

	// 删除日志项目
	// 请根据您的需要，填写待删除的ProjectId
	// DeleteProject API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112176
	_, _ = client.DeleteProject(&tls.DeleteProjectRequest{
		ProjectID: projectID,
	})
}
