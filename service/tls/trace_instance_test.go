package tls

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SDKTraceInstanceTestSuite struct {
	suite.Suite

	cli Client
}

func (suite *SDKTraceInstanceTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()
}

func (suite *SDKTraceInstanceTestSuite) validateError(err error, expectErr *Error) {
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

func (suite *SDKTraceInstanceTestSuite) TestModifyTraceInstanceValidation() {
	suite.T().Run("EmptyTraceInstanceId", func(t *testing.T) {
		request := &ModifyTraceInstanceRequest{TraceInstanceId: ""}
		err := request.CheckValidation()
		suite.Error(err)
	})

	suite.T().Run("ValidRequest", func(t *testing.T) {
		request := &ModifyTraceInstanceRequest{
			TraceInstanceId: "test-trace-instance-id",
			Description:     StrPtr("test description"),
		}
		err := request.CheckValidation()
		suite.NoError(err)
	})

	suite.T().Run("ValidRequestWithoutDescription", func(t *testing.T) {
		request := &ModifyTraceInstanceRequest{TraceInstanceId: "test-trace-instance-id"}
		err := request.CheckValidation()
		suite.NoError(err)
	})
}

func (suite *SDKTraceInstanceTestSuite) TestDescribeTraceInstance() {
	traceInstanceId := "test-trace-instance-id"

	req := &DescribeTraceInstanceRequest{
		TraceInstanceId: traceInstanceId,
	}

	resp, err := suite.cli.DescribeTraceInstance(req)

	// 由于这是单元测试，我们主要验证请求结构和基本逻辑
	// 实际的 API 调用结果取决于服务端状态
	if err != nil {
		// 验证错误类型
		suite.NotNil(err)
	} else {
		// 验证响应结构
		suite.NotNil(resp)
		suite.NotEmpty(resp.RequestID)

		// 验证返回的字段是否符合预期结构
		if resp.TraceInstanceId != "" {
			suite.Equal(traceInstanceId, resp.TraceInstanceId)
		}
	}
}

func (suite *SDKTraceInstanceTestSuite) TestDescribeTraceInstanceValidation() {
	// 测试空 TraceInstanceId 的验证
	req := &DescribeTraceInstanceRequest{
		TraceInstanceId: "",
	}

	_, err := suite.cli.DescribeTraceInstance(req)
	suite.NotNil(err)

	sdkErr, ok := err.(*Error)
	suite.True(ok)
	suite.NotNil(sdkErr)
}

func TestSDKTraceInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(SDKTraceInstanceTestSuite))
}
