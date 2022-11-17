package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKHostGroupTestSuite struct {
	suite.Suite

	cli       Client
	project   string
	topic     string
	hostGroup string
}

func (suite *SDKHostGroupTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-index-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	keyValueList := make([]KeyValueInfo, 0)
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-1",
		Value: Value{
			ValueType:      "text",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-2",
		Value: Value{
			ValueType:      "long",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	suite.NoError(CreateIndex(topicId, nil, &keyValueList, suite.cli))
}

func (suite *SDKHostGroupTestSuite) TearDownTest() {
	_, deleteHostGroupErr := suite.cli.DeleteHostGroup(&DeleteHostGroupRequest{HostGroupID: suite.hostGroup})
	suite.NoError(deleteHostGroupErr)
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKHostGroupTestSuite(t *testing.T) {
	suite.Run(t, new(SDKHostGroupTestSuite))
}

func (suite *SDKHostGroupTestSuite) TestHostGroup() {
	var hostGroup string
	// CreateHostGroup
	{
		request := &CreateHostGroupRequest{
			HostGroupName:   "unit-test-host-gorup",
			HostGroupType:   "IP",
			HostIPList:      &[]string{"127.0.0.1"},
			HostIdentifier:  new(string),
			AutoUpdate:      new(bool),
			UpdateStartTime: new(string),
			UpdateEndTime:   new(string),
		}
		*request.HostIdentifier = "unit-test-identifier"
		*request.AutoUpdate = true
		*request.UpdateStartTime = "00:00"
		*request.UpdateEndTime = "02:00"

		response, err := suite.cli.CreateHostGroup(request)
		suite.NoError(err)
		hostGroup = response.HostGroupID
		suite.hostGroup = hostGroup
	}

	// ModifyHostGroupsAutoUpdate
	{
		request := &ModifyHostGroupsAutoUpdateRequest{
			HostGroupIds:    []string{hostGroup},
			AutoUpdate:      new(bool),
			UpdateStartTime: new(string),
			UpdateEndTime:   new(string),
		}
		*request.AutoUpdate = true
		*request.UpdateStartTime = "01:00"
		*request.UpdateEndTime = "02:30"

		_, err := suite.cli.ModifyHostGroupsAutoUpdate(request)
		suite.NoError(err)
	}

	// DescribeHostGroup
	{
		request := &DescribeHostGroupRequest{
			HostGroupID: hostGroup,
		}
		_, err := suite.cli.DescribeHostGroup(request)
		suite.NoError(err)
	}

	// CreateHostGroupInvalid
	{
		request := &CreateHostGroupRequest{
			HostGroupName:   "unit-test-host-gorup",
			HostGroupType:   "IP",
			HostIPList:      &[]string{},
			HostIdentifier:  new(string),
			AutoUpdate:      new(bool),
			UpdateStartTime: new(string),
			UpdateEndTime:   new(string),
		}
		*request.HostIdentifier = "unit-test-identifier"
		*request.AutoUpdate = true
		*request.UpdateStartTime = "00:00"
		*request.UpdateEndTime = "02:00"

		_, err := suite.cli.CreateHostGroup(request)
		suite.Error(err)
	}

	// ModifyHostGroupsAutoUpdateInvalid
	{
		request := &ModifyHostGroupsAutoUpdateRequest{
			HostGroupIds:    []string{hostGroup},
			AutoUpdate:      new(bool),
			UpdateStartTime: new(string),
			UpdateEndTime:   new(string),
		}
		*request.AutoUpdate = false
		*request.UpdateStartTime = "01:00"
		*request.UpdateEndTime = "02:30"

		_, err := suite.cli.ModifyHostGroupsAutoUpdate(request)
		suite.Error(err)
	}

	// DescribeHostGroupInvalid
	{
		request := &DescribeHostGroupRequest{
			HostGroupID: "unit-test",
		}
		_, err := suite.cli.DescribeHostGroup(request)
		suite.Error(err)
	}
}
