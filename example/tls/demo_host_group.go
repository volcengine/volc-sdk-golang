package tls

import (
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	//初始化客户端，配置AccessKeyID,AccessKeySecret,region,securityToken;securityToken可以为空
	client := tls.NewClient(testEndPoint, testAk, testSk, testSessionToken, testRegion)

	//create host group
	hostIdentifier := "label"
	hostIPList := []string{"192.168.0.1"}
	createHostGroupReq := tls.CreateHostGroupRequest{
		HostGroupName:  "mgn1",
		HostGroupType:  "IP",
		HostIdentifier: &hostIdentifier,
		HostIPList:     &hostIPList,
	}
	hostGroup, _ := client.CreateHostGroup(&createHostGroupReq)
	testHostGroupId := hostGroup.HostGroupID

	//modify host group
	modifyHostIdentifier := "label"
	modifyHostIPList := []string{"192.168.0.1"}
	modifyHostGroupName := "mgn1"
	modifyHostGroupType := "IP"
	modifyHostGroupReq := tls.ModifyHostGroupRequest{
		HostGroupID:    testHostGroupId,
		HostGroupName:  &modifyHostGroupName,
		HostGroupType:  &modifyHostGroupType,
		HostIdentifier: &modifyHostIdentifier,
		HostIPList:     &modifyHostIPList,
	}
	_, _ = client.ModifyHostGroup(&modifyHostGroupReq)

	//describe host group
	_, _ = client.DescribeHostGroup(&tls.DescribeHostGroupRequest{
		HostGroupID: testHostGroupId,
	})

	//describe host groups
	_, _ = client.DescribeHostGroups(&tls.DescribeHostGroupsRequest{})
	_, _ = client.DescribeHostGroups(&tls.DescribeHostGroupsRequest{
		PageNumber: 1,
		PageSize:   10,
	})

	//describe hosts
	_, _ = client.DescribeHosts(&tls.DescribeHostsRequest{
		HostGroupID: testHostGroupId,
	})

	searchIP := "192.168.0.1"
	_, _ = client.DescribeHosts(&tls.DescribeHostsRequest{
		HostGroupID: testHostGroupId,
		IP:          &searchIP,
	})

	//delete host
	_, _ = client.DeleteHost(&tls.DeleteHostRequest{
		HostGroupID: testHostGroupId,
		IP:          "192.168.0.1",
	})

	//delete host group
	_, _ = client.DeleteHostGroup(&tls.DeleteHostGroupRequest{
		HostGroupID: testHostGroupId,
	})
}
