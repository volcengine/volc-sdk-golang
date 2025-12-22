package tls

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKHostGroupTestSuite struct {
	suite.Suite

	cli           Client
	project       string
	topic         string
	hostGroupList []string
}

func (suite *SDKHostGroupTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-project-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-topic-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	keyValueList := make([]KeyValueInfo, 0)
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-1",
		Value: Value{
			ValueType:      "text",
			Delimiter:      "",
			CaseSensitive:  false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-2",
		Value: Value{
			ValueType:      "long",
			Delimiter:      "",
			CaseSensitive:  false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	suite.NoError(CreateIndex(topicId, nil, &keyValueList, suite.cli))
}

func (suite *SDKHostGroupTestSuite) TearDownTest() {
	for _, hostGroupID := range suite.hostGroupList {
		_, deleteHostGroupErr := suite.cli.DeleteHostGroup(&DeleteHostGroupRequest{HostGroupID: hostGroupID})
		suite.NoError(deleteHostGroupErr)
	}
	suite.hostGroupList = nil

	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func (suite *SDKHostGroupTestSuite) validateError(err error, expectErr *Error) {
	sdkErr, ok := err.(*Error)

	if sdkErr == nil {
		suite.Nil(sdkErr)
		return
	}

	suite.Equal(true, ok)
	suite.Equal(expectErr.HTTPCode, sdkErr.HTTPCode)
	suite.Equal(expectErr.Code, sdkErr.Code)
	suite.Equal(expectErr.Message, sdkErr.Message)
}

func TestSDKHostGroupTestSuite(t *testing.T) {
	suite.Run(t, new(SDKHostGroupTestSuite))
}

func createHostGroup(cli Client) (string, error) {
	resp, err := cli.CreateHostGroup(&CreateHostGroupRequest{
		HostGroupName: "go-sdk-test-host-group" + uuid.New().String(),
		HostGroupType: "IP",
		HostIPList:    &[]string{"192.168.1.1", "192.168.1.2", "192.168.1.3"},
	})
	if err != nil {
		return "", err
	}

	return resp.HostGroupID, nil
}

func (suite *SDKHostGroupTestSuite) TestCreateHostGroupNormally() {
	testcases := map[*CreateHostGroupRequest]*Error{
		{
			HostGroupName: "go-sdk-test-host-group" + uuid.New().String(),
			HostGroupType: "IP",
			HostIPList:    &[]string{"192.168.1.1"},
		}: nil,
		{
			HostGroupName:   "go-sdk-test-host-group" + uuid.New().String(),
			HostGroupType:   "IP",
			HostIPList:      &[]string{"192.168.1.1"},
			AutoUpdate:      BoolPtr(true),
			UpdateStartTime: StrPtr("00:00"),
			UpdateEndTime:   StrPtr("02:00"),
			ServiceLogging:  BoolPtr(true),
			IamProjectName:  StrPtr("default"),
		}: nil,
	}

	for req, expectedErr := range testcases {
		resp, err := suite.cli.CreateHostGroup(req)
		suite.validateError(err, expectedErr)
		if resp != nil {
			suite.hostGroupList = append(suite.hostGroupList, resp.HostGroupID)
		}
	}
}

func (suite *SDKHostGroupTestSuite) TestCreateHostGroupAbnormally() {
	testcases := map[*CreateHostGroupRequest]*Error{
		{
			HostGroupName: "go-sdk-test-host-group" + uuid.New().String(),
			HostGroupType: "IP",
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArguments",
			Message:  "Invalid argument HostGroupType, HostIpList, please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		resp, err := suite.cli.CreateHostGroup(req)
		suite.validateError(err, expectedErr)
		if resp != nil {
			suite.hostGroupList = append(suite.hostGroupList, resp.HostGroupID)
		}
	}
}

func (suite *SDKHostGroupTestSuite) TestDeleteHostGroupAbnormally() {
	hostGroupID := uuid.New().String()
	_, err := suite.cli.DeleteHostGroup(&DeleteHostGroupRequest{HostGroupID: hostGroupID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "HostGroupNotExist",
		Message:  "HostGroup " + hostGroupID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKHostGroupTestSuite) TestModifyHostGroupNormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	_, err = suite.cli.ModifyHostGroup(&ModifyHostGroupRequest{
		HostGroupID:   hostGroupID,
		HostGroupName: StrPtr("new-host-group-name"),
	})
	suite.NoError(err)
	resp, err := suite.cli.DescribeHostGroup(&DescribeHostGroupRequest{HostGroupID: hostGroupID})
	suite.NoError(err)
	suite.Equal("new-host-group-name", resp.HostGroupHostsRulesInfo.HostGroupInfo.HostGroupName)
}

func (suite *SDKHostGroupTestSuite) TestModifyHostGroupAbnormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	testcases := map[*ModifyHostGroupRequest]*Error{
		{
			HostGroupID:   hostGroupID,
			HostGroupName: StrPtr("$$new-host-group-name$$"),
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key HostGroupName, value $$new-host-group-name$$, please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err = suite.cli.ModifyHostGroup(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostGroupNormally() {
	testcases := map[*CreateHostGroupRequest]*DescribeHostGroupResponse{
		{
			HostGroupName:   "go-sdk-test-host-group-1",
			HostGroupType:   "IP",
			HostIPList:      &[]string{"192.168.1.1"},
			AutoUpdate:      BoolPtr(true),
			UpdateStartTime: StrPtr("00:00"),
			UpdateEndTime:   StrPtr("02:00"),
			ServiceLogging:  BoolPtr(true),
			IamProjectName:  StrPtr("default"),
		}: {
			HostGroupHostsRulesInfo: &HostGroupHostsRulesInfo{
				HostGroupInfo: &HostGroupInfo{
					HostGroupName:   "go-sdk-test-host-group-1",
					HostGroupType:   "IP",
					HostCount:       1,
					AutoUpdate:      true,
					UpdateStartTime: "00:00",
					UpdateEndTime:   "02:00",
					ServiceLogging:  true,
					IamProjectName:  "default",
				},
			},
		},
	}

	for createHostGroupReq, expectedHostGroupHostsRulesInfo := range testcases {
		createHostGroupResp, err := suite.cli.CreateHostGroup(createHostGroupReq)
		suite.NoError(err)
		suite.hostGroupList = append(suite.hostGroupList, createHostGroupResp.HostGroupID)

		resp, err := suite.cli.DescribeHostGroup(&DescribeHostGroupRequest{HostGroupID: createHostGroupResp.HostGroupID})
		expectedHostGroupInfo := expectedHostGroupHostsRulesInfo.HostGroupHostsRulesInfo.HostGroupInfo
		suite.Equal(expectedHostGroupInfo.HostGroupName, resp.HostGroupHostsRulesInfo.HostGroupInfo.HostGroupName)
		suite.Equal(expectedHostGroupInfo.HostGroupType, resp.HostGroupHostsRulesInfo.HostGroupInfo.HostGroupType)
		suite.Equal(expectedHostGroupInfo.HostCount, resp.HostGroupHostsRulesInfo.HostGroupInfo.HostCount)
		suite.Equal(expectedHostGroupInfo.AutoUpdate, resp.HostGroupHostsRulesInfo.HostGroupInfo.AutoUpdate)
		suite.Equal(expectedHostGroupInfo.UpdateStartTime, resp.HostGroupHostsRulesInfo.HostGroupInfo.UpdateStartTime)
		suite.Equal(expectedHostGroupInfo.UpdateEndTime, resp.HostGroupHostsRulesInfo.HostGroupInfo.UpdateEndTime)
		suite.Equal(expectedHostGroupInfo.ServiceLogging, resp.HostGroupHostsRulesInfo.HostGroupInfo.ServiceLogging)
		suite.Equal(expectedHostGroupInfo.IamProjectName, resp.HostGroupHostsRulesInfo.HostGroupInfo.IamProjectName)
	}
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostGroupAbnormally() {
	hostGroupID := uuid.New().String()
	_, err := suite.cli.DescribeHostGroup(&DescribeHostGroupRequest{HostGroupID: hostGroupID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "HostGroupNotExist",
		Message:  "HostGroup " + hostGroupID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostGroupsNormally() {
	resp, err := suite.cli.DescribeHostGroups(&DescribeHostGroupsRequest{})
	suite.NoError(err)
	suite.GreaterOrEqual(int(resp.Total), 0)
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostGroupsAbnormally() {
	testcases := map[*DescribeHostGroupsRequest]*Error{
		{
			PageSize: 200,
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key PageSize, value 200, please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeHostGroups(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostsNormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	testcases := map[*DescribeHostsRequest]map[string]bool{
		{
			HostGroupID: hostGroupID,
		}: {
			"192.168.1.1": true,
			"192.168.1.2": true,
			"192.168.1.3": true,
		},
	}

	for describeHostsReq, hostsInfo := range testcases {
		resp, err := suite.cli.DescribeHosts(describeHostsReq)
		suite.NoError(err)
		suite.Equal(len(hostsInfo), int(resp.Total))
		for _, host := range resp.HostInfos {
			suite.Equal(true, hostsInfo[host.IP])
		}
	}
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostsAbnormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	testcases := map[*DescribeHostsRequest]*Error{
		{
			HostGroupID: hostGroupID,
			PageNumber:  -1,
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key PageNumber, value -1, please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err = suite.cli.DescribeHosts(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKHostGroupTestSuite) TestDeleteHostNormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	_, err = suite.cli.DeleteHost(&DeleteHostRequest{
		HostGroupID: hostGroupID,
		IP:          "192.168.1.3",
	})
	suite.NoError(err)
	resp, err := suite.cli.DescribeHosts(&DescribeHostsRequest{HostGroupID: hostGroupID})
	suite.NoError(err)
	suite.Equal(2, int(resp.Total))
	for _, host := range resp.HostInfos {
		suite.NotEqual("192.168.1.3", host.IP)
	}
}

func (suite *SDKHostGroupTestSuite) TestDeleteHostAbnormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	testcases := map[*DeleteHostRequest]*Error{
		{
			HostGroupID: hostGroupID,
			IP:          "192.168.2.3",
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "HostNotExist",
			Message:  "Host 192.168.2.3 does not exist in HostGroup " + hostGroupID,
		},
	}

	for req, expectedErr := range testcases {
		_, err = suite.cli.DeleteHost(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostGroupRulesNormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	resp, err := suite.cli.DescribeHostGroupRules(&DescribeHostGroupRulesRequest{HostGroupID: hostGroupID})
	suite.NoError(err)
	suite.Equal(0, int(resp.Total))
}

func (suite *SDKHostGroupTestSuite) TestDescribeHostGroupRulesAbnormally() {
	nonexistentHostGroupID := uuid.New().String()
	testcases := map[*DescribeHostGroupRulesRequest]*Error{
		{
			HostGroupID: nonexistentHostGroupID,
		}: {
			HTTPCode: http.StatusNotFound,
			Code:     "HostGroupNotExist",
			Message:  "HostGroup " + nonexistentHostGroupID + " does not exist",
		},
	}

	for req, expectedErr := range testcases {
		_, err := suite.cli.DescribeHostGroupRules(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKHostGroupTestSuite) TestModifyHostGroupsAutoUpdateNormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	_, err = suite.cli.ModifyHostGroupsAutoUpdate(&ModifyHostGroupsAutoUpdateRequest{
		HostGroupIds:    []string{hostGroupID},
		AutoUpdate:      BoolPtr(true),
		UpdateStartTime: StrPtr("00:00"),
		UpdateEndTime:   StrPtr("02:00"),
	})
	suite.NoError(err)
	resp, err := suite.cli.DescribeHostGroup(&DescribeHostGroupRequest{HostGroupID: hostGroupID})
	suite.NoError(err)
	suite.Equal(true, resp.HostGroupHostsRulesInfo.HostGroupInfo.AutoUpdate)
	suite.Equal("00:00", resp.HostGroupHostsRulesInfo.HostGroupInfo.UpdateStartTime)
	suite.Equal("02:00", resp.HostGroupHostsRulesInfo.HostGroupInfo.UpdateEndTime)
}

func (suite *SDKHostGroupTestSuite) TestModifyHostGroupsAutoUpdateAbnormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	testcases := map[*ModifyHostGroupsAutoUpdateRequest]*Error{
		{
			HostGroupIds: []string{hostGroupID},
			AutoUpdate:   BoolPtr(true),
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Argument UpdateStartTime is empty, please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		_, err = suite.cli.ModifyHostGroupsAutoUpdate(req)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKHostGroupTestSuite) TestDeleteAbnormalHostsNormally() {
	hostGroupID, err := createHostGroup(suite.cli)
	suite.NoError(err)
	suite.hostGroupList = append(suite.hostGroupList, hostGroupID)

	_, err = suite.cli.DeleteAbnormalHosts(&DeleteAbnormalHostsRequest{HostGroupID: hostGroupID})
	suite.NoError(err)
}

func (suite *SDKHostGroupTestSuite) TestDeleteAbnormalHostsAbnormally() {
	nonexistentHostGroupID := uuid.New().String()
	_, err := suite.cli.DeleteAbnormalHosts(&DeleteAbnormalHostsRequest{HostGroupID: nonexistentHostGroupID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "HostGroupNotExist",
		Message:  "HostGroup " + nonexistentHostGroupID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}
