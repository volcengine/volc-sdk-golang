package tls

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKTraceTestSuite struct {
	suite.Suite
	cli       Client
	project   string
	traceList []string
}

func (suite *SDKTraceTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()
}

func (suite *SDKTraceTestSuite) TestDescribeTraceInstances() {
	// 测试基本的 DescribeTraceInstances 调用
	request := &DescribeTraceInstancesRequest{
		PageNumber: intPtr(1),
		PageSize:   intPtr(20),
	}

	response, err := suite.cli.DescribeTraceInstances(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.GreaterOrEqual(response.Total, int64(0))
	suite.NotNil(response.TraceInstances)
}

func (suite *SDKTraceTestSuite) TestDescribeTraceInstancesWithProjectID() {
	// 测试带 ProjectID 的 DescribeTraceInstances 调用
	projectID := "test-project-id"
	request := &DescribeTraceInstancesRequest{
		ProjectID:  &projectID,
		PageNumber: intPtr(1),
		PageSize:   intPtr(10),
	}

	response, err := suite.cli.DescribeTraceInstances(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.GreaterOrEqual(response.Total, int64(0))
}

func (suite *SDKTraceTestSuite) TestDescribeTraceInstancesWithTraceInstanceName() {
	// 测试带 TraceInstanceName 的 DescribeTraceInstances 调用
	traceInstanceName := "test-trace-instance"
	request := &DescribeTraceInstancesRequest{
		TraceInstanceName: &traceInstanceName,
		PageNumber:        intPtr(1),
		PageSize:          intPtr(10),
	}

	response, err := suite.cli.DescribeTraceInstances(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.GreaterOrEqual(response.Total, int64(0))

	projectID, err := CreateProject("golang-sdk-create-project-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectID
}

func (suite *SDKTraceTestSuite) TearDownTest() {
	for _, traceInstanceID := range suite.traceList {
		// 清理 Trace 实例（假设有对应的删除接口）
		// 这里需要根据实际的删除接口来实现
		_ = traceInstanceID
	}

	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)

	suite.traceList = nil
}

func TestSDKTraceTestSuite(t *testing.T) {
	suite.Run(t, new(SDKTraceTestSuite))
}

// 辅助函数
func intPtr(i int) *int {
	return &i
}

func (suite *SDKTraceTestSuite) TestCreateTraceInstanceNormally() {
	traceName := "单元测试" + uuid.New().String()[:8]

	request := &CreateTraceInstanceRequest{
		CommonRequest: CommonRequest{},
		Data: TraceInsCreateReq{
			ProjectID:         suite.project,
			TraceInstanceName: traceName,
		},
	}

	response, err := suite.cli.CreateTraceInstance(request)
	suite.NoError(err)
	suite.NotEmpty(response.TraceInstanceID)

	// 添加到清理列表
	suite.traceList = append(suite.traceList, response.TraceInstanceID)
}

func (suite *SDKTraceTestSuite) TestCreateTraceInstanceAbnormally() {
	testcases := map[*CreateTraceInstanceRequest]*Error{
		{
			CommonRequest: CommonRequest{},
			Data: TraceInsCreateReq{
				ProjectID:         "",
				TraceInstanceName: "test-trace",
			},
		}: {
			HTTPCode: 400,
			Code:     "InvalidArgument",
			Message:  "Invalid argument, empty ProjectID",
		},
		{
			CommonRequest: CommonRequest{},
			Data: TraceInsCreateReq{
				ProjectID:         suite.project,
				TraceInstanceName: "",
			},
		}: {
			HTTPCode: 400,
			Code:     "InvalidArgument",
			Message:  "Invalid argument, empty TraceInstanceName",
		},
	}

	for request, expectedErr := range testcases {
		_, err := suite.cli.CreateTraceInstance(request)
		suite.validateError(err, expectedErr)
	}
}

func (suite *SDKTraceTestSuite) validateError(err error, expectErr *Error) {
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
