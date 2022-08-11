package main

import (
	"encoding/json"
	"fmt"
	url2 "net/url"
	"strings"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func main() {
	ListWorkspacesDemo()
	PurchaseWorkspaceDemo()
	UpdateWorkspaceDemo()
	DeleteWorkspaceDemo()
	GetWorkspaceLimitDemo()
	CreateRenderJobDemo()
	ListRenderJobsDemo()
	GetRenderJobDemo()
	UpdateJobPriorityDemo()
	SingleJobOperationDemo()
	BatchJobOperationDemo()
	EditRenderJobDemo()
	GetLayerFramesDemo()
	GetAccountStatisticsDemo()
	GetAccountStatisticDetailsDemo()
	DownloadAccountStatisticDetailsDemo()
	GetJobOverallStatisticsDemo()
	ListMyMessagesDemo()
	SingleMessageOperationDemo()
	BatchMessagesOperationDemo()
	DeleteAllMessagesDemo()
	UploadFolderDemo()
	DownloadFileDemo()
	ListObjectDemo()
	StatObjectDemo()
}

func GetVerenderInstance() (*verender.Verender, error) {
	v := verender.NewInstance("", "")

	return v, nil
}

func StatObjectDemo() {
	v, _ := GetVerenderInstance()
	w, _ := v.GetWorkspace(1442)

	objInfo, err := w.StatObject("2048/1.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(objInfo)
}

func ListObjectDemo() {
	v, _ := GetVerenderInstance()
	w, _ := v.GetWorkspace(1442)

	prefix := "2048/"
	startAfter := ""
	maxKeys := 1000
	for {
		files, subs, next, err := w.ListObject(prefix, startAfter, maxKeys)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(len(*files))
		fmt.Println(subs)

		if next == "" {
			break
		} else {
			startAfter = next
		}
	}
}

func ListWorkspacesDemo() {
	v, _ := GetVerenderInstance()

	r := &verender.ListWorkspaceRequest{}

	resp, err := v.ListWorkspaces(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func PurchaseWorkspaceDemo() {
	v, _ := GetVerenderInstance()

	resourcePools, err := v.ListResourcePools()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := verender.Workspace{
		WorkspaceName:  "Your Workspace Name",
		Description:    "Your Workspace Description",
		StorageTotal:   100 << 30,
		ResourcePoolId: (*resourcePools)[0].ResourcePoolId,
		CpsId:          (*resourcePools)[0].CpsId,
	}

	w, err := v.PurchaseWorkspace(&r)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(w)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func UpdateWorkspaceDemo() {
	v, _ := GetVerenderInstance()

	w := verender.Workspace{
		WorkspaceId:   1682,
		WorkspaceName: "Update Workspace Name",
		Description:   "Update Workspace Description",
		StorageTotal:  150 << 30,
	}

	resp, err := v.UpdateWorkspace(&w)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func DeleteWorkspaceDemo() {
	v, _ := GetVerenderInstance()

	w := verender.Workspace{
		WorkspaceId: 1682,
	}

	err := v.DeleteWorkspace(&w)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetWorkspaceLimitDemo() {
	v, _ := GetVerenderInstance()

	resp, err := v.GetWorkspaceLimit()
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func CreateRenderJobDemo() {
	v, _ := GetVerenderInstance()

	user, _ := v.GetCurrentUser()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sceneFile, err := w.UploadFile("/tmp/test.ma", "asset")
	if err != nil {
		fmt.Println(err)
		return
	}
	svrPath := func() string {
		url, err := url2.Parse(sceneFile)
		if err != nil {
			return sceneFile
		}
		segs := strings.Split(url.Path, "/")
		if len(segs) < 2 {
			return sceneFile
		}

		return strings.Join(segs[2:], "/")
	}()

	job := verender.JobInfo{
		UserName:           user.UserName,
		Title:              "test-job-1",
		Description:        "",
		DccTool:            "maya",
		Renderer:           "redshift",
		Tryout:             true,
		TryoutFrameNumbers: []string{"1", "10"},
		FrameSettings: verender.FrameSettings{
			Start: 1,
			End:   10,
			Step:  1,
		},
		SceneFile: svrPath,
		Cameras:   []string{"render_camera"},
		Layers:    []string{"masterLayer"},
	}

	resp, err := w.CreateJob(&job)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func ListRenderJobsDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := verender.ListJobRequest{
		JobType: "maya_redshift",
	}
	jobs, err := w.ListJobs(&r)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(jobs)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func GetRenderJobDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	job, err := w.GetJob("r2e40c22574")
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(job.JobInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func UpdateJobPriorityDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = w.UpdateJobPriority("r2e40c22574", 10); err != nil {
		fmt.Println(err)
		return
	}
}

// PauseJob/ResumeJob/RetryJob/FullSpeedJob/DeleteJob/StopJob示例
func SingleJobOperationDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = w.DeleteJob("r7b22301433"); err != nil {
		fmt.Println(err)
		return
	}
}

// PauseJobs/ResumeJobs/FullSpeedJobs/DeleteJobs/StopJobs 示例
func BatchJobOperationDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := verender.BatchJobRequest{
		JobIds: []string{
			"r4a3703fa21",
			"rb860f95b62",
		},
	}

	if err = w.DeleteJobs(&r); err != nil {
		fmt.Println(err)
		return
	}
}

// EditRenderJobDemo 目前只支持修改Title和Description
func EditRenderJobDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	job := verender.JobInfo{
		Title:       "update-job-name",
		Description: "update-job-description",
		JobId:       "r5ed716ef6c",
	}

	if err = w.EditJob(&job); err != nil {
		fmt.Println(err)
		return
	}
}

func GetLayerFramesDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := verender.GetLayerFramesRequest{
		RenderJobId: "r5ed716ef6c",
		LayerRequests: []verender.LayerRequest{
			{
				LayerIndex: 0,
				PageNum:    0,
				PageSize:   10,
				Statuses:   []int{0},
			},
		},
	}

	resp, err := w.GetLayerFrames(&r)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func GetAccountStatisticsDemo() {
	v, _ := GetVerenderInstance()

	r := verender.StatisticsRequest{
		TimeType: "last_week",
	}
	resp, err := v.GetAccountStatistics(&r)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func GetAccountStatisticDetailsDemo() {
	v, _ := GetVerenderInstance()

	r := verender.StatisticsRequest{
		StartTime: "2022-05-01T00:00:00+08:00",
		EndTime:   "2022-05-31T23:59:59+08:00",
		PageNum:   1,
		PageSize:  10,
		Filter: verender.StatisticsFilter{
			WorkspaceIds: []int{1442},
		},
	}

	resp, err := v.GetAccountStatisticDetails(&r)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func DownloadAccountStatisticDetailsDemo() {
	v, _ := GetVerenderInstance()

	r := verender.StatisticsRequest{
		StartTime: "2022-05-30T00:00:00+08:00",
		EndTime:   "2022-05-31T23:59:59+08:00",
		FileName:  "/tmp/12345.csv",
	}

	if err := v.DownloadStatisticDetails(&r); err != nil {
		fmt.Println(err)
		return
	}
}

func GetJobOverallStatisticsDemo() {
	v, _ := GetVerenderInstance()

	resp, err := v.GetJobOverallStatistics()
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func ListMyMessagesDemo() {
	v, _ := GetVerenderInstance()

	r := verender.ListMessageRequest{
		PageNum:  1,
		PageSize: 10,
	}
	messages, err := v.ListMyMessages(&r)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(messages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

// MarkMessageAsRead/DeleteMessage
func SingleMessageOperationDemo() {
	v, _ := GetVerenderInstance()

	if err := v.DeleteMessage(3007); err != nil {
		fmt.Println(err)
		return
	}
}

func BatchMessagesOperationDemo() {
	v, _ := GetVerenderInstance()

	r := verender.MessageIdList{
		MessageIds: []int{
			3006,
			3005,
		},
	}

	if err := v.BatchMarkMessagesAsRead(&r); err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteAllMessagesDemo() {
	v, _ := GetVerenderInstance()
	if err := v.DeleteAllMessages(); err != nil {
		fmt.Println(err)
		return
	}
}

func UploadFolderDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	locs, err := w.UploadFolder("/tmp/2048", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(locs)
}

func DownloadFileDemo() {
	v, _ := GetVerenderInstance()
	w, err := v.GetWorkspace(1442)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = w.DownloadFile("/tmp/d.log", "test.ma"); err != nil {
		fmt.Println(err)
		return
	}
}
