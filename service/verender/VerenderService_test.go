package verender

import (
	"os"
	"testing"
)

const (
	AK          = "Your AK"
	SK          = "Your SK"
	WORKSPACEID = 1
	JOBID       = "Your Job Id"
)

var DefaultInstance = NewInstance(AK, SK)

func init() {

}

func Test_packErrorInfo(t *testing.T) {
	resp := ResponseMetaData{
		Error: Error{
			CodeN:   10001,
			Code:    "Error1",
			Message: "Error1",
		},
		RequestId: "test1",
	}

	expect := "RequestId=test1 Error={10001 Error1 Error1}"
	err := packErrorInfo(resp)

	if expect != err.Error() {
		t.Errorf("test packErrorInfo failed")
		return
	}
}

func TestVerender_GetWorkspace(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if w == nil {
		t.Errorf("GetWorkspace Error %v", WORKSPACEID)
		return
	}
}

func TestVerender_DeleteAllMessages(t *testing.T) {
	r := MessageIdList{
		MessageIds: []int{
			1,
			2,
		},
	}

	if err := DefaultInstance.BatchDeleteMessages(&r); err != nil {
		t.Errorf("DeleteAllMessages Error %v", err)
		return
	}
}

func TestVerender_DeleteMessage(t *testing.T) {
	if err := DefaultInstance.DeleteMessage(1); err != nil {
		t.Errorf("DeleteMessage Error %v", err)
		return
	}
}

func TestVerender_DeleteWorkspace(t *testing.T) {
	w := Workspace{
		WorkspaceId: 1234,
	}
	if err := DefaultInstance.DeleteWorkspace(&w); err != nil {
		t.Errorf("DeleteWorkspace Error %v", err)
		return
	}
}

func TestVerender_BatchMarkMessagesAsRead(t *testing.T) {
	r := MessageIdList{
		MessageIds: []int{
			1,
			2,
		},
	}

	if err := DefaultInstance.BatchMarkMessagesAsRead(&r); err != nil {
		t.Errorf("BatchMarkMessagesAsRead Error %v", err)
		return
	}
}

func TestVerender_DeleteAllReadMessages(t *testing.T) {
	if err := DefaultInstance.DeleteAllReadMessages(); err != nil {
		t.Errorf("DeleteAllReadMessages Error %v", err)
		return
	}
}

func TestVerender_DownloadStatisticDetails(t *testing.T) {
	fileName := "/tmp/test.csv"
	r := StatisticsRequest{
		StartTime: "2022-05-30T00:00:00+08:00",
		EndTime:   "2022-05-31T23:59:59+08:00",
		PageNum:   1,
		PageSize:  100,
		FileName:  fileName,
	}

	if err := DefaultInstance.DownloadStatisticDetails(&r); err != nil {
		t.Errorf("DownloadStatisticDetails Error %v", err)
		return
	}

	if _, err := os.Stat(fileName); err != nil {
		t.Errorf("DownloadStatisticDetails Error %v", err)
		return
	}
}

func TestVerender_GetAccountStatisticDetails(t *testing.T) {
	r := StatisticsRequest{
		StartTime: "2022-05-30T00:00:00+08:00",
		EndTime:   "2022-05-31T23:59:59+08:00",
		PageNum:   1,
		PageSize:  100,
		Filter: StatisticsFilter{
			WorkspaceIds: []int{WORKSPACEID},
		},
	}

	resp, err := DefaultInstance.GetAccountStatisticDetails(&r)
	if err != nil || resp == nil {
		t.Errorf("GetAccountStatisticDetails Error %v", err)
		return
	}
}

func TestVerender_GetCurrentUser(t *testing.T) {
	user, err := DefaultInstance.GetCurrentUser()
	if err != nil || user == nil {
		t.Errorf("GetCurrentUser Error %v", err)
		return
	}
}

func TestVerender_GetAccountStatistics(t *testing.T) {
	r := StatisticsRequest{
		TimeType: "last_week",
	}

	resp, err := DefaultInstance.GetAccountStatistics(&r)
	if err != nil || resp == nil {
		t.Errorf("GetAccountStatistics Error %v", err)
		return
	}
}

func TestVerender_GetHardwareSpecifications(t *testing.T) {
	resp, err := DefaultInstance.GetHardwareSpecifications(WORKSPACEID)
	if err != nil || resp == nil {
		t.Errorf("GetHardwareSpecifications Error %v", err)
		return
	}
}

func TestVerender_ListWorkspaces(t *testing.T) {
	r := ListWorkspaceRequest{}
	resp, err := DefaultInstance.ListWorkspaces(&r)
	if err != nil || resp == nil {
		t.Errorf("ListWorkspace Error %v", err)
		return
	}
}

func TestVerender_GetJobOverallStatistics(t *testing.T) {
	resp, err := DefaultInstance.GetJobOverallStatistics()
	if err != nil || resp == nil || resp.Success <= 0 {
		t.Errorf("GetJobOverallStatistics Error %v", err)
		return
	}
}

func TestVerender_GetWorkspaceLimit(t *testing.T) {
	resp, err := DefaultInstance.GetWorkspaceLimit()
	if err != nil || resp == nil || resp.WorkspaceAmountUpperLimit <= 0 {
		t.Errorf("GetWorkspaceLimit Error %v", err)
		return
	}
}

func TestVerender_ListMyMessages(t *testing.T) {
	r := ListMessageRequest{}
	resp, err := DefaultInstance.ListMyMessages(&r)
	if err != nil || resp == nil || resp.Total <= 0 {
		t.Errorf("ListMyMessages Error %v", err)
		return
	}
}

func TestVerender_UpdateWorkspace(t *testing.T) {
	w := Workspace{
		WorkspaceId:   1650,
		WorkspaceName: "workspace-new-name",
		StorageTotal:  110 << 30,
	}

	resp, err := DefaultInstance.UpdateWorkspace(&w)
	if err != nil || resp == nil || resp.WorkspaceId != w.WorkspaceId {
		t.Errorf("UpdateWorkspace Error %v", err)
		return
	}
}

func TestVerender_ListResourcePools(t *testing.T) {
	resp, err := DefaultInstance.ListResourcePools()
	if err != nil || resp == nil {
		t.Errorf("ListResourcePools Error %v", err)
		return
	}
}

func TestVerender_MarkAllMessagesAsRead(t *testing.T) {
	if err := DefaultInstance.MarkAllMessagesAsRead(); err != nil {
		t.Errorf("MarkAllMessagesAsRead Error %v", err)
		return
	}
}

func TestVerender_PurchaseWorkspace(t *testing.T) {
	pools, err := DefaultInstance.ListResourcePools()
	if err != nil || pools == nil {
		t.Errorf("ListResourcePools Error %v", err)
	}

	w := Workspace{
		WorkspaceName:  "new-workspace",
		Description:    "",
		StorageTotal:   100 << 30,
		CpsId:          (*pools)[0].CpsId,
		ResourcePoolId: (*pools)[0].ResourcePoolId,
	}

	resp, err := DefaultInstance.PurchaseWorkspace(&w)
	if err != nil || resp.WorkspaceId <= 0 {
		t.Errorf("PurchaseWorkspace Error %v", err)
		return
	}
}

func TestVerender_MarkMessageAsRead(t *testing.T) {
	if err := DefaultInstance.MarkMessageAsRead(1); err != nil {
		t.Errorf("MarkMessageAsRead Error %v", err)
		return
	}
}

func TestVerender_BatchDeleteMessages(t *testing.T) {
	r := MessageIdList{
		MessageIds: []int{
			1,
			2,
		},
	}
	if err := DefaultInstance.BatchDeleteMessages(&r); err != nil {
		t.Errorf("BatchDeleteMessages Error %v", err)
		return
	}
}

func TestWorkspace_CreateJob(t *testing.T) {
	user, err := DefaultInstance.GetCurrentUser()
	if err != nil || user == nil {
		t.Errorf("GetCurrentUser Error %v", err)
		return
	}
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil || w.WorkspaceId <= 0 {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}
	job := JobInfo{
		UserName:    user.UserName,
		Title:       "test-create-render-job",
		Description: "",
		DccTool:     "maya",
		Renderer:    "redshift",
		Tryout:      false,
		SceneFile:   "test.ma",
		Cameras:     []string{"render_camera"},
		Layers:      []string{"masterLayer"},
	}

	resp, err := w.CreateJob(&job)
	if err != nil || resp == nil || resp.JobId == "" {
		t.Errorf("CreateJob Error %v", err)
		return
	}

	job.DccTool = "test"

	resp, err = w.CreateJob(&job)
	if err == nil || resp != nil {
		t.Errorf("invalid dcc create job success")
		return
	}
}

func TestWorkspace_DeleteJob(t *testing.T) {
	jobId := "test-123"

	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if err = w.DeleteJob(jobId); err != nil {
		t.Errorf("DeleteJob Error %v", err)
		return
	}
}

func TestWorkspace_DeleteJobs(t *testing.T) {
	jobIds := BatchJobRequest{
		JobIds: []string{
			"test-job-name-1",
			"test-job-name-2",
		},
	}

	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if err = w.DeleteJobs(&jobIds); err != nil {
		t.Errorf("DeleteJobs Error %v", err)
		return
	}
}

func TestWorkspace_DownloadFile(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	local := "/tmp/1.txt"
	target := "test.ma"

	if err = w.DownloadFile(local, target); err != nil {
		t.Errorf("DownloadFile Error %v", err)
		return
	}
}

func TestWorkspace_EditJob(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	job := JobInfo{
		JobId:       JOBID,
		Title:       "edit-job-name",
		Description: "edit-job-desc",
	}

	if err = w.EditJob(&job); err != nil {
		t.Errorf("EditJob Error %v", err)
		return
	}
}

func TestWorkspace_FullSpeedJob(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if err = w.FullSpeedJob(JOBID); err != nil {
		t.Errorf("FullSpeedJob Error %v", err)
		return
	}
}

func TestWorkspace_FullSpeedJobs(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}
	jobIds := BatchJobRequest{
		JobIds: []string{
			"test-job-id-1",
			"test-job-id-2",
		},
	}

	if err = w.FullSpeedJobs(&jobIds); err != nil {
		t.Errorf("FullSpeedJobs Error %v", err)
		return
	}
}

func TestWorkspace_GetFrames(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	r := GetFramesRequest{
		RenderJobId: JOBID,
		PageNum:     1,
		PageSize:    10,
	}
	resp, err := w.GetFrames(&r)
	if err != nil || resp == nil || resp.TotalFrames <= 0 {
		t.Errorf("GetFrames Error %v", err)
		return
	}
}

func TestWorkspace_GetJob(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	resp, err := w.GetJob(JOBID)
	if err != nil || resp == nil {
		t.Errorf("GetJob Error %v", err)
		return
	}
}

func TestWorkspace_GetLayerFrames(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	r := GetLayerFramesRequest{
		RenderJobId: JOBID,
		LayerRequests: []LayerRequest{
			{
				LayerIndex: 0,
			},
		},
	}

	resp, err := w.GetLayerFrames(&r)
	if err != nil || resp == nil {
		t.Errorf("GetLayerFrames Error %v", err)
		return
	}
}

func TestWorkspace_ListFile(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	_, err = w.ListFile()
	if err != nil {
		t.Errorf("ListFile Error %v", err)
		return
	}
}

func TestWorkspace_ListJobs(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	r := ListJobRequest{
		PageNum:  1,
		PageSize: 10,
	}

	resp, err := w.ListJobs(&r)
	if err != nil || resp == nil {
		t.Errorf("ListJobs Error %v", err)
		return
	}
}

func TestWorkspace_PauseJob(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if err = w.PauseJob(JOBID); err != nil {
		t.Errorf("PauseJob Error %v", err)
		return
	}
}

func TestWorkspace_PauseJobs(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	r := BatchJobRequest{
		JobIds: []string{
			"test-job-id-1",
			"test-job-id-2",
		},
	}
	if err = w.PauseJobs(&r); err != nil {
		t.Errorf("PauseJobs Error %v", err)
		return
	}
}

func TestWorkspace_ResumeJob(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if err = w.ResumeJob(JOBID); err != nil {
		t.Errorf("ResumeJob Error %v", err)
		return
	}
}

func TestWorkspace_ResumeJobs(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	r := BatchJobRequest{
		JobIds: []string{
			"test-job-id-1",
			"test-job-id-2",
		},
	}
	if err = w.ResumeJobs(&r); err != nil {
		t.Errorf("ResumeJobs Error %v", err)
		return
	}
}

func TestWorkspace_RetryJob(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if err = w.RetryJob(JOBID); err != nil {
		t.Errorf("RetryJob Error %v", err)
		return
	}
}

func TestWorkspace_StopJob(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	if err = w.StopJob(JOBID); err != nil {
		t.Errorf("StopJob Error %v", err)
		return
	}
}

func TestWorkspace_StopJobs(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorksace Error %v", err)
		return
	}

	r := BatchJobRequest{
		JobIds: []string{
			"test-job-id-1",
			"test-job-id-2",
		},
	}
	if err = w.StopJobs(&r); err != nil {
		t.Errorf("StopJobs Error %v", err)
		return
	}
}

func TestWorkspace_UploadFile(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	resp, err := w.UploadFile("/tmp/test.ma", "")
	if err != nil || resp == "" {
		t.Errorf("UploadFile Error %v", err)
		return
	}
}

func TestWorkspace_UploadFolder(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspace Error %v", err)
		return
	}

	_, err = w.UploadFolder("/tmp/123/456", "")
	if err != nil {
		t.Errorf("UploadFolder Error %v", err)
		return
	}
}

func TestWorkspace_UpdateJobPriority(t *testing.T) {
	w, err := DefaultInstance.GetWorkspace(WORKSPACEID)
	if err != nil || w == nil {
		t.Errorf("GetWorkspaceError %v", err)
		return
	}

	if err = w.UpdateJobPriority(JOBID, 10); err != nil {
		t.Errorf("UpdateJobPriority Error %v", err)
		return
	}
}
