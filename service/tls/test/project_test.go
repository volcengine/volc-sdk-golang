package test

import (
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"os"
	"testing"
)

func TestCreateProject(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	createRequest := &CreateProjectRequest{
		ProjectName: "testa",
		Description: "ABCDEFG",
		Region:      os.Getenv("LOG_TEST_REGION"),
	}

	resp, err := client.CreateProject(createRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if resp.ProjectID == "" {
		t.Error("empty project id")
	}
}

func TestDeleteProject(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	deleteRequest := &DeleteProjectRequest{
		ProjectID: "project-id",
	}

	_, err := client.DeleteProject(deleteRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestUpdateProject(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	updateRequest := &ModifyProjectRequest{
		ProjectID:   "project-id",
		ProjectName: StrPtr("project-name"),
	}

	_, err := client.ModifyProject(updateRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestGetProject(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	getRequest := &DescribeProjectRequest{
		ProjectID: "project-id",
	}

	projectInfo, err := client.DescribeProject(getRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if projectInfo.ProjectID == "" {
		t.Error("empty project id")
	}
}

func TestListProject(t *testing.T) {
	client := NewClient(os.Getenv("LOG_TEST_ENDPOINT"), os.Getenv("LOG_TEST_ACCESS_KEY_ID"),
		os.Getenv("LOG_TEST_ACCESS_KEY_SECRET"), os.Getenv("LOG_TEST_SECURITY_TOKEN"),
		os.Getenv("LOG_TEST_REGION"))

	listRquest := &DescribeProjectsRequest{

		ProjectID:   "project-id",
		ProjectName: "project-name",
		PageNumber:  1,
		PageSize:    10,
	}

	projectList, err := client.DescribeProjects(listRquest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(projectList.Projects) == 0 {
		t.Error("empty result")
	}
}

func StrPtr(in string) *string {
	return &in
}
