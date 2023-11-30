package tls

import (
	"fmt"
	"os"

	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 创建机器组
	// 请根据您的需要，填写HostGroupName、HostGroupType和HostIpList等参数
	// CreateHostGroup API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112206
	createHostGroupResp, _ := client.CreateHostGroup(&tls.CreateHostGroupRequest{
		HostGroupName: "my-host-group",
		HostGroupType: "IP",
		HostIPList:    &[]string{"192.168.0.1"},
	})
	hostGroupId := createHostGroupResp.HostGroupID

	// 修改机器组
	// 请根据您的需要，填写HostGroupId和待修改的机器组信息
	// ModifyHostGroup API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112209
	hostGroupName := "my-host-group-new"
	hostGroupType := "IP"
	_, _ = client.ModifyHostGroup(&tls.ModifyHostGroupRequest{
		HostGroupID:   hostGroupId,
		HostGroupName: &hostGroupName,
		HostGroupType: &hostGroupType,
		HostIPList:    &[]string{"192.168.1.1", "192.168.1.2"},
	})

	// 获取指定机器组信息
	// 请根据您的需要，填写待查询的HostGroupId
	// DescribeHostGroup API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112210
	describeHostGroupResp, _ := client.DescribeHostGroup(&tls.DescribeHostGroupRequest{
		HostGroupID: hostGroupId,
	})
	fmt.Println(describeHostGroupResp.HostGroupHostsRulesInfo.HostGroupInfo.HostGroupName)

	// 获取全部机器组信息
	// 请根据您的需要，填写HostGroupName等可选参数
	// DescribeHostGroups API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112211
	describeHostGroupsResp, _ := client.DescribeHostGroups(&tls.DescribeHostGroupsRequest{
		PageNumber: 1,
		PageSize:   10,
	})
	fmt.Println(describeHostGroupsResp.Total)

	// 查询机器组所有机器
	// 请根据您的需要，填写待查询的HostGroupId和其它参数
	// DescribeHosts API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112212
	searchIP := "192.168.1.1"
	describeHostsResp, _ := client.DescribeHosts(&tls.DescribeHostsRequest{
		HostGroupID: hostGroupId,
		IP:          &searchIP,
	})
	fmt.Println(describeHostsResp.Total)

	// 删除机器组内指定机器
	// 请根据您的需要，填写待删除机器的HostGroupId和IP
	// DeleteHost API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112213
	_, _ = client.DeleteHost(&tls.DeleteHostRequest{
		HostGroupID: hostGroupId,
		IP:          "192.168.1.1",
	})

	// 删除机器组
	// 请根据您的需要，填写待删除的HostGroupId
	// DeleteHostGroup API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112208
	_, _ = client.DeleteHostGroup(&tls.DeleteHostGroupRequest{
		HostGroupID: hostGroupId,
	})
}
