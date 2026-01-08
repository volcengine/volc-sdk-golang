package tls

import (
	"os"
	"strings"
	"sync"
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

var traceTestCleanupOnce sync.Once
var traceTestCleanupErr error

func (suite *SDKTraceTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()
}

func (suite *SDKTraceTestSuite) TestDescribeTraceInstances() {
	// 测试基本的 DescribeTraceInstances 调用
	request := &DescribeTraceInstancesRequest{
		PageNumber: IntPtr(1),
		PageSize:   IntPtr(20),
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
		PageNumber: IntPtr(1),
		PageSize:   IntPtr(10),
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
		PageNumber:        IntPtr(1),
		PageSize:          IntPtr(10),
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
	traceIDs := make(map[string]struct{}, len(suite.traceList))
	for _, traceInstanceID := range suite.traceList {
		if traceInstanceID == "" {
			continue
		}
		traceIDs[traceInstanceID] = struct{}{}
	}

	if suite.project != "" {
		ids, err := suite.listTraceInstanceIDsByProject(suite.project)
		suite.NoError(err)
		for _, id := range ids {
			if id == "" {
				continue
			}
			traceIDs[id] = struct{}{}
		}
	}

	for traceInstanceID := range traceIDs {
		_, _ = suite.cli.DeleteTraceInstance(&DeleteTraceInstanceRequest{TraceInstanceId: traceInstanceID})
	}

	if suite.project != "" {
		_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
		suite.NoError(deleteProjectErr)
	}

	traceTestCleanupOnce.Do(func() {
		traceTestCleanupErr = suite.cleanupTestTraceInstances()
	})
	suite.NoError(traceTestCleanupErr)

	suite.traceList = nil
}

func (suite *SDKTraceTestSuite) listTraceInstanceIDsByProject(projectID string) ([]string, error) {
	pageSize := 100
	pageNumber := 1

	ids := make([]string, 0)
	for {
		pid := projectID
		resp, err := suite.cli.DescribeTraceInstances(&DescribeTraceInstancesRequest{
			ProjectID:  &pid,
			PageNumber: IntPtr(pageNumber),
			PageSize:   IntPtr(pageSize),
		})
		if err != nil {
			return ids, err
		}
		if resp == nil || len(resp.TraceInstances) == 0 {
			return ids, nil
		}

		for _, instance := range resp.TraceInstances {
			if instance == nil {
				continue
			}
			ids = append(ids, instance.TraceInstanceID)
		}

		if resp.Total > 0 && int64(len(ids)) >= resp.Total {
			return ids, nil
		}

		pageNumber++
	}
}

func (suite *SDKTraceTestSuite) cleanupTestTraceInstances() error {
	const testProjectPrefix = "golang-sdk-create-project-"
	const testTraceNamePrefix = "单元测试"

	pageSize := 100
	pageNumber := 1
	for {
		resp, err := suite.cli.DescribeProjects(&DescribeProjectsRequest{
			ProjectName: testProjectPrefix,
			PageNumber:  pageNumber,
			PageSize:    pageSize,
		})
		if err != nil {
			return err
		}
		if resp == nil || len(resp.Projects) == 0 {
			return nil
		}

		for _, project := range resp.Projects {
			instances, err := suite.listTraceInstancesByProject(project.ProjectID)
			if err != nil {
				return err
			}
			for _, instance := range instances {
				if instance == nil {
					continue
				}
				if !strings.HasPrefix(instance.TraceInstanceName, testTraceNamePrefix) {
					continue
				}
				if instance.TraceInstanceID == "" {
					continue
				}
				_, _ = suite.cli.DeleteTraceInstance(&DeleteTraceInstanceRequest{TraceInstanceId: instance.TraceInstanceID})
			}
		}

		if resp.Total > 0 && int64(pageNumber*pageSize) >= resp.Total {
			return nil
		}
		pageNumber++
	}
}

func (suite *SDKTraceTestSuite) listTraceInstancesByProject(projectID string) ([]*TraceInstance, error) {
	pageSize := 100
	pageNumber := 1

	instances := make([]*TraceInstance, 0)
	for {
		pid := projectID
		resp, err := suite.cli.DescribeTraceInstances(&DescribeTraceInstancesRequest{
			ProjectID:  &pid,
			PageNumber: IntPtr(pageNumber),
			PageSize:   IntPtr(pageSize),
		})
		if err != nil {
			return instances, err
		}
		if resp == nil || len(resp.TraceInstances) == 0 {
			return instances, nil
		}

		instances = append(instances, resp.TraceInstances...)

		if resp.Total > 0 && int64(len(instances)) >= resp.Total {
			return instances, nil
		}
		pageNumber++
	}
}

func TestSDKTraceTestSuite(t *testing.T) {
	suite.Run(t, new(SDKTraceTestSuite))
}

func (suite *SDKTraceTestSuite) TestCreateTraceInstanceNormally() {
	traceName := "单元测试" + uuid.New().String()[:8]

	request := &CreateTraceInstanceRequest{
		CommonRequest:     CommonRequest{},
		ProjectID:         suite.project,
		TraceInstanceName: traceName,
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
			CommonRequest:     CommonRequest{},
			ProjectID:         "",
			TraceInstanceName: "test-trace",
		}: {
			HTTPCode: 400,
			Code:     "InvalidArgument",
			Message:  "Invalid argument, empty ProjectID",
		},
		{
			CommonRequest:     CommonRequest{},
			ProjectID:         suite.project,
			TraceInstanceName: "",
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
