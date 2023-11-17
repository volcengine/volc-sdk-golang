package tls

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKProjectTestSuite struct {
	suite.Suite

	cli         Client
	projectList []string
}

func (suite *SDKProjectTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()
}

func (suite *SDKProjectTestSuite) TearDownTest() {
	for _, projectID := range suite.projectList {
		_, err := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: projectID})
		suite.NoError(err)
	}

	suite.projectList = nil
}

func (suite *SDKProjectTestSuite) validateError(err error, expectErr *Error) {
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

func TestSDKProjectTestSuite(t *testing.T) {
	suite.Run(t, new(SDKProjectTestSuite))
}

func (suite *SDKProjectTestSuite) TestCreateProjectNormally() {
	iamProjectName := "default"

	testcases := map[*CreateProjectRequest]*Error{
		{
			ProjectName: "golang-sdk-create-project-" + uuid.New().String(),
			Description: "",
			Region:      os.Getenv("LOG_SERVICE_REGION"),
		}: nil,
		{
			ProjectName: "golang-sdk-create-project-" + uuid.New().String(),
			Description: "hello world",
			Region:      os.Getenv("LOG_SERVICE_REGION"),
		}: nil,
		{
			ProjectName:    "golang-sdk-create-project-" + uuid.New().String(),
			Description:    "create a project with the IamProjectName and Tags",
			Region:         os.Getenv("LOG_SERVICE_REGION"),
			IamProjectName: &iamProjectName,
			Tags: []TagInfo{
				{
					Key:   "TagKey1",
					Value: "TagValue1",
				},
				{
					Key:   "TagKey2",
					Value: "TagValue2",
				},
			},
		}: nil,
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.CreateProject(req)
		suite.validateError(rErr, err)

		if resp != nil {
			suite.projectList = append(suite.projectList, resp.ProjectID)
		}
	}
}

func (suite *SDKProjectTestSuite) TestCreateProjectAbnormally() {
	testcases := map[*CreateProjectRequest]error{
		{
			ProjectName: "a",
			Description: "",
			Region:      os.Getenv("LOG_SERVICE_REGION"),
		}: &Error{
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key ProjectName, value a, please check argument.",
		},
		{
			ProjectName: "golang-sdk-create-project-" + uuid.New().String(),
			Description: "",
			Region:      "usa",
		}: &Error{
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key Region, value usa, please check argument.",
		},
	}

	for req, err := range testcases {
		_, rErr := suite.cli.CreateProject(req)
		suite.validateError(rErr, err.(*Error))
	}
}

func (suite *SDKProjectTestSuite) TestDeleteProjectAbnormally() {
	_, err := suite.cli.DeleteProject(&DeleteProjectRequest{
		ProjectID: uuid.New().String(),
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKProjectTestSuite) TestModifyProjectNormally() {
	createReq := &CreateProjectRequest{
		ProjectName: "golang-sdk-update-project-" + uuid.New().String(),
		Description: "desc",
		Region:      os.Getenv("LOG_SERVICE_REGION"),
	}

	createResp, err := suite.cli.CreateProject(createReq)
	suite.NoError(err)

	suite.projectList = append(suite.projectList, createResp.ProjectID)

	_, err = suite.cli.ModifyProject(&ModifyProjectRequest{
		ProjectID: createResp.ProjectID,
	})
	suite.NoError(err)
	getResp, err := suite.cli.DescribeProject(&DescribeProjectRequest{ProjectID: createResp.ProjectID})
	suite.NoError(err)
	suite.Equal(createReq.ProjectName, getResp.ProjectName)
	suite.Equal("desc", getResp.Description)

	_, err = suite.cli.ModifyProject(&ModifyProjectRequest{
		ProjectID:   createResp.ProjectID,
		ProjectName: StrPtr(createReq.ProjectName + "1"),
		Description: StrPtr("new desc"),
	})
	suite.NoError(err)
	getResp, err = suite.cli.DescribeProject(&DescribeProjectRequest{ProjectID: createResp.ProjectID})
	suite.NoError(err)
	suite.Equal(createReq.ProjectName+"1", getResp.ProjectName)
	suite.Equal("new desc", getResp.Description)
}

func (suite *SDKProjectTestSuite) TestModifyProjectAbnormally() {
	_, err := suite.cli.ModifyProject(&ModifyProjectRequest{
		ProjectID: "",
	})
	expectedErr := &Error{
		HTTPCode: -1,
		Code:     "ClientError",
		Message:  "Invalid argument, empty ProjectID",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKProjectTestSuite) TestDescribeProjectNormally() {
	iamProjectName := "default"
	createReq := &CreateProjectRequest{
		ProjectName:    "golang-sdk-get-project-" + uuid.New().String(),
		Description:    "desc",
		Region:         os.Getenv("LOG_SERVICE_REGION"),
		IamProjectName: &iamProjectName,
		Tags: []TagInfo{
			{
				Key:   "TagKey1",
				Value: "TagValue1",
			},
			{
				Key:   "TagKey2",
				Value: "TagValue2",
			},
		},
	}

	createResp, err := suite.cli.CreateProject(createReq)
	suite.NoError(err)
	suite.projectList = append(suite.projectList, createResp.ProjectID)

	testcases := map[*DescribeProjectRequest]*Error{
		{
			ProjectID: createResp.ProjectID,
		}: nil,
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.DescribeProject(req)
		suite.validateError(rErr, err)

		if rErr == nil {
			suite.Equal(createResp.ProjectID, resp.ProjectID)
			suite.Equal(createReq.ProjectName, resp.ProjectName)
			suite.Equal(createReq.Description, resp.Description)
			suite.Equal(iamProjectName, resp.IamProjectName)
			suite.Equal(len(createReq.Tags), len(resp.Tags))
		}
	}
}

func (suite *SDKProjectTestSuite) TestDescribeProjectAbnormally() {
	_, err := suite.cli.DescribeProject(&DescribeProjectRequest{
		ProjectID: uuid.New().String(),
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKProjectTestSuite) TestDescribeProjectsNormally() {
	iamProjectName := "default"
	createReqs := []*CreateProjectRequest{
		{
			ProjectName: "sdk-a-group" + uuid.New().String() + "0",
			Description: "desc",
			Region:      os.Getenv("LOG_SERVICE_REGION"),
		},
		{
			ProjectName: "sdk-a-group" + uuid.New().String() + "1",
			Description: "desc",
			Region:      os.Getenv("LOG_SERVICE_REGION"),
		},
		{
			ProjectName: "sdk-b-group" + uuid.New().String() + "2",
			Description: "desc",
			Region:      os.Getenv("LOG_SERVICE_REGION"),
		},
		{
			ProjectName:    "sdk-c-group" + uuid.New().String(),
			Description:    "desc",
			Region:         os.Getenv("LOG_SERVICE_REGION"),
			IamProjectName: &iamProjectName,
			Tags: []TagInfo{
				{
					Key:   "TagKey1",
					Value: "TagValue1",
				},
				{
					Key:   "TagKey2",
					Value: "TagValue2",
				},
			},
		},
	}

	for _, createReq := range createReqs {
		createResp, err := suite.cli.CreateProject(createReq)
		suite.NoError(err)
		suite.projectList = append(suite.projectList, createResp.ProjectID)
	}

	testcases := map[*DescribeProjectsRequest]int{
		{
			ProjectName: "sdk-a",
			PageNumber:  1,
		}: 2,
		{
			ProjectName: "sdk-b",
			PageNumber:  1,
		}: 1,
		{
			ProjectName: "sdk-c",
			PageNumber:  1,
			PageSize:    1,
		}: 1,
		{
			ProjectName:    "sdk-c",
			PageNumber:     1,
			PageSize:       1,
			IamProjectName: &iamProjectName,
			Tags: []TagInfo{
				{
					Key:   "TagKey1",
					Value: "TagValue1",
				},
			},
		}: 1,
	}

	for req, expect := range testcases {
		listResp, err := suite.cli.DescribeProjects(req)
		suite.NoError(err)
		suite.Equal(expect, len(listResp.Projects))
	}

	resp, err := suite.cli.DescribeProjects(&DescribeProjectsRequest{})
	suite.NoError(err)
	suite.GreaterOrEqual(int(resp.Total), 4)
}

func (suite *SDKProjectTestSuite) TestDescribeProjectsAbnormally() {
	projectID := uuid.New().String()
	_, err := suite.cli.DescribeProjects(&DescribeProjectsRequest{
		ProjectName: "sdk",
		ProjectID:   projectID,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArguments",
		Message:  "Invalid argument sdk, " + projectID + ", please check argument.",
	}
	suite.validateError(err, expectedErr)
}
