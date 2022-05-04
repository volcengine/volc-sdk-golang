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

	suite.Equal(ok, true)
	suite.Equal(sdkErr.HTTPCode, expectErr.HTTPCode)
	suite.Equal(sdkErr.Code, expectErr.Code)
	suite.Equal(sdkErr.Message, expectErr.Message)
}

func TestSDKProjectTestSuite(t *testing.T) {
	suite.Run(t, new(SDKProjectTestSuite))
}

// TestCreateNormalProject: test create project with valid parameters
func (suite *SDKProjectTestSuite) TestCreateNormalProject() {
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
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.CreateProject(req)
		suite.validateError(rErr, err)

		if resp != nil {
			suite.projectList = append(suite.projectList, resp.ProjectId)
		}
	}
}

// TestCreateAbnormalProject: test create project with invalid parameters
func (suite *SDKProjectTestSuite) TestCreateAbnormalProject() {
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
			Region:      "",
		}: &Error{
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key Region, value , please check argument.",
		},
	}

	for req, err := range testcases {
		_, rErr := suite.cli.CreateProject(req)
		suite.validateError(rErr, err.(*Error))
	}
}

// TestGetProject: test get project
func (suite *SDKProjectTestSuite) TestGetProject() {
	createReq := &CreateProjectRequest{
		ProjectName: "golang-sdk-get-project-" + uuid.New().String(),
		Description: "desc",
		Region:      os.Getenv("LOG_SERVICE_REGION"),
	}

	createResp, err := suite.cli.CreateProject(createReq)
	suite.NoError(err)

	suite.projectList = append(suite.projectList, createResp.ProjectId)

	testcases := map[*DescribeProjectRequest]*Error{
		{
			ProjectID: createResp.ProjectId,
		}: nil,
		{
			ProjectID: uuid.New().String(),
		}: &Error{
			HTTPCode: http.StatusNotFound,
			Code:     "ProjectNotExists",
			Message:  "Project does not exist.",
		},
	}

	for req, err := range testcases {
		resp, rErr := suite.cli.DescribeProject(req)
		suite.validateError(rErr, err)

		if rErr == nil {
			suite.Equal(resp.ProjectID, createResp.ProjectId)
			suite.Equal(resp.ProjectName, createReq.ProjectName)
			suite.Equal(resp.Description, createReq.Description)
		}
	}
}

// TestUpdateProject: test update project
func (suite *SDKProjectTestSuite) TestUpdateProject() {
	createReq := &CreateProjectRequest{
		ProjectName: "golang-sdk-update-project-" + uuid.New().String(),
		Description: "desc",
		Region:      os.Getenv("LOG_SERVICE_REGION"),
	}

	createResp, err := suite.cli.CreateProject(createReq)
	suite.NoError(err)

	suite.projectList = append(suite.projectList, createResp.ProjectId)

	testcases := map[*ModifyProjectRequest]*CreateProjectRequest{
		{
			ProjectID: createResp.ProjectId,
		}: {
			ProjectName: createReq.ProjectName,
			Description: "desc",
		},
		{
			ProjectID:   createResp.ProjectId,
			ProjectName: StrPtr(createReq.ProjectName + "1"),
			Description: StrPtr("new desc"),
		}: {
			ProjectName: createReq.ProjectName + "1",
			Description: "new desc",
		},
	}

	for req, expect := range testcases {
		_, rErr := suite.cli.ModifyProject(req)
		suite.NoError(rErr)

		getResp, err := suite.cli.DescribeProject(&DescribeProjectRequest{ProjectID: createResp.ProjectId})
		suite.NoError(err)

		suite.Equal(getResp.ProjectName, expect.ProjectName)
		suite.Equal(getResp.Description, expect.Description)
	}
}

// TestListProject: test list project
func (suite *SDKProjectTestSuite) TestListProject() {
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
	}

	for _, createReq := range createReqs {
		createResp, err := suite.cli.CreateProject(createReq)
		suite.NoError(err)

		suite.projectList = append(suite.projectList, createResp.ProjectId)
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
			ProjectName: "sdk-a",
			PageNumber:  1,
			PageSize:    1,
		}: 1,
	}

	for req, expect := range testcases {
		listResp, err := suite.cli.DescribeProjects(req)
		suite.NoError(err)
		suite.Equal(len(listResp.Projects), expect)
	}
}
