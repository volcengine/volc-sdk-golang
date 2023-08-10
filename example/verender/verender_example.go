package main

import (
    "encoding/json"
    "fmt"
    "time"

    "github.com/volcengine/volc-sdk-golang/service/verender"
)

func main() {
    ListWorkspaceDemo()
    GetCurrentUserDemo()
    UploadFileDemo()
    UploadFolderDemo()
    DownloadFileDemo()
    StatFileDemo()
    ListFileDemo()
    ListCellSpecDemo()
    AddRenderSettingDemo()
    DeleteRenderSettingDemo()
    UpdateRenderSettingDemo()
    ListDccDemo()
    ListRenderSettingDemo()
    GetRenderSettingDemo()
    CreateRenderJobDemo()
    ListRenderJobDemo()
    GetRenderJobDemo()
    RetryRenderJobDemo()
    ResumeRenderJobsDemo()
    FullSpeedRenderJobsDemo()
    AutoFullSpeedRenderJobsDemo()
    StopRenderJobsDemo()
    DeleteRenderJobsDemo()
    ListJobOutputDemo()
    GetJobOutputDemo()
    UpdateJobOutputDemo()
    UpdateRenderJobsPriorityDemo()
    ListAccountDccPluginsDemo()
    AutoFullSpeedRenderJobsDemo()
}

func getVerenderInstance() *verender.Verender {
    v := verender.NewVerenderInstance()
    v.Client.SetAccessKey("your ak")
    v.Client.SetSecretKey("your sk")

    return v
}

func ListAccountDccPluginsDemo() {
    v := getVerenderInstance()

    r := verender.ListAccountDccPluginsReq{
        SpecTemplateId: 54,
    }

    resp, err := v.ListAccountDccPlugins(&r)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(resp)
}

func ListWorkspaceDemo() {
    v := getVerenderInstance()
    r := verender.ListWorkspaceRequest{
        WorkspaceID:        0,
        WorkspaceName:      "",
        PageNum:            0,
        PageSize:           0,
        StartPurchaseTime:  "",
        EndPurchaseTime:    "",
        OrderType:          "",
        OrderField:         "",
        FuzzyWorkspaceName: "",
        FuzzySearchContent: "",
        ListScope:          "",
    }

    resp, err := v.ListWorkspace(&r)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func GetCurrentUserDemo() {
    v := getVerenderInstance()
    resp, err := v.GetCurrentUser()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func UploadFileDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fileName := "/tmp/512M.dat"
    resp, err := w.UploadFile(fileName, fileName)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func UploadFolderDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    dir := "/tmp/test_upload_folder"
    resp, err := w.UploadFolder(dir, "/tmp/abc")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func DownloadFileDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    src := "tmp/512M.dat"
    dst := "/tmp/512M_p.dat"
    if err := w.DownloadFile(src, dst); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("download file %s done\n", src)
}

func StatFileDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fo, err := w.StatFile("tmp/512M.dat")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(fo)
    fmt.Println(string(data))
}

func RemoveFileDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    if err := w.RemoveFile("tmp/512M.dat"); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("remove file done")
}

func ListCellSpecDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    resp, err := w.ListCellSpec()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func ListFileDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    files, dirs, err := w.ListFile("tmp/", "", "", "", 1, 10)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    subFiles, _ := json.Marshal(files)
    subDirs, _ := json.Marshal(dirs)
    fmt.Println(string(subFiles))
    fmt.Println(string(subDirs))
}

func AddRenderSettingDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    user, err := v.GetCurrentUser()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    rs := verender.RenderSetting{
        AccountID:   user.AccountID,
        UserID:      user.UserID,
        WorkspaceID: w.Id,
        Name:        "test-sdk-add-render-setting-1",
        Dcc:         "maya",
        DccVersion:  "2022",
        Plugins: []*verender.Plugin{
            {
                Name:           "mtoa",
                Version:        "5.0.0",
                NeedLicense:    true,
                RendererPlugin: true,
            },
        },
        RenderLayerMode:    "LegacyRenderLayer",
        ProjectPath:        "",
        FramesCountOneCell: 1,
        CellSpecID:         9,
    }

    resp, err := w.AddRenderSetting(&rs)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func DeleteRenderSettingDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    if err := w.DeleteRenderSetting(1191); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("DeleteRenderSetting done")
}

func UpdateRenderSettingDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    rs := verender.RenderSetting{
        Id:   128,
        Name: fmt.Sprintf("test-sdk-update-render-setting-%d", time.Now().Unix()),
    }

    if err := w.UpdateRenderSetting(&rs); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("UpdateRenderSettingDone")
}

func ListDccDemo() {
    v := getVerenderInstance()
    dccs, err := v.ListDccs()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(dccs)
}

func ListRenderSettingDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    resp, err := w.ListRenderSetting("maya")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)

    fmt.Println(string(data))
}

func GetRenderSettingDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    resp, err := w.GetRenderSetting(861)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func CreateRenderJobDemo() {
    v := getVerenderInstance()
    workspaceID := int64(735)
    w, err := v.GetWorkspace(workspaceID)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    rs, err := w.GetRenderSetting(157)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fileName := "/Users/bytedance/Documents/jc_test/jc_sample/Advanced_07.ma"
    obj, err := w.UploadFile(fileName, fileName)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    renderer := ""
    for _, p := range rs.Plugins {
        if p.RendererPlugin {
            renderer = p.Name
        }
    }

    user, err := v.GetCurrentUser()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    j := verender.RenderJob{
        AccountID:                user.AccountID,
        UserID:                   user.UserID,
        UserName:                 user.UserName,
        WorkspaceID:              workspaceID,
        Title:                    "test-create-render-job",
        Description:              "",
        Tryout:                   true,
        TryoutFrameNumbers:       []string{"START", "END", "MIDDLE"},
        SceneFile:                obj.Key,
        PathMapping:              nil,
        FramesCountEachCell:      8,
        TimeoutReminderEachFrame: 86400,
        TimeoutStopperEachFrame:  86400,
        LayerConfig: []*verender.LayerConfig{
            {
                LayerIndex: 0,
                LayerName:  "masterLayer",
                Frame: verender.Frame{
                    Start: 1,
                    End:   12000,
                    Step:  1,
                },
                Resolutions: verender.Resolution{
                    Width:  1920,
                    Height: 1080,
                },
                Cameras: []string{"persp"},
            },
        },
        RenderSetting: rs,
        PluginData:    "{}",
        Priority:      0,
        Renderer:      renderer,
    }

    resp, err := w.CreateRenderJob(&j)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func ListRenderJobDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(735)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.ListRenderJobRequest{}
    resp, err := w.ListRenderJob(&r)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func GetRenderJobDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    resp, err := w.GetRenderJob("r76199b2e96")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func UpdateRenderJobsPriorityDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.BatchJobPriority{
        JobIDList: []string{"r76199b2e96"},
        Priority:  2,
    }

    if err := w.UpdateRenderJobsPriority(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("UpdateRenderJobPriority done")
}

func RetryRenderJobDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.RetryJobRequest{
        JobID:           "r5ad3829bef",
        AllFailedFrames: true,
    }

    if err := w.RetryRenderJob(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("RetryRenderJob done")
}

func ResumeRenderJobsDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.BatchJobIDList{
        JobIDList: []string{"rb3bf60e847", "r24d6fbecae"},
    }

    if err := w.ResumeRenderJobs(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("ResumeRenderJobs done")
}

func FullSpeedRenderJobsDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.BatchJobIDList{
        JobIDList: []string{"rb3bf60e847", "r24d6fbecae"},
    }

    if err := w.FullSpeedRenderJobs(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("FullSpeedRenderJobs done")
}

func AutoFullSpeedRenderJobsDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(795)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.BatchJobIDList{
        JobIDList: []string{"rb3bf60e847", "r24d6fbecae"},
    }

    if err := w.AutoFullSpeedRenderJobs(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("AutoFullSpeedRenderJobs done")
}

func StopRenderJobsDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.BatchJobIDList{
        JobIDList: []string{"rb3bf60e847", "r24d6fbecae"},
    }

    if err := w.StopRenderJobs(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("StopRenderJobs done")
}

func DeleteRenderJobsDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.BatchJobIDList{
        JobIDList: []string{"rb3bf60e847", "r24d6fbecae"},
    }

    if err := w.DeleteRenderJobs(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("DeleteRenderJobs done")
}

func ListJobOutputDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.ListJobOutputRequest{
        StartTime:  "",
        EndTime:    "",
        PageNum:    0,
        PageSize:   0,
        Type:       "",
        Status:     "",
        OrderType:  "",
        OrderField: "",
        JobIDList:  nil,
    }

    resp, err := w.ListJobOutput(&r)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func GetJobOutputDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.GetJobOutputRequest{
        JobID:  "ra884188ec7",
        Filter: nil,
    }
    resp, err := w.GetJobOutput(&r)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    data, _ := json.Marshal(resp)
    fmt.Println(string(data))
}

func UpdateJobOutputDemo() {
    v := getVerenderInstance()
    w, err := v.GetWorkspace(1802)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.UpdateJobOutputRequest{
        JobID: "rb3bf60e847",
        Files: []string{"Result/test-create-render-job_rb3bf60e847/thumbs/rb3bf60e847_masterLayer_4.png"},
    }
    if err := w.UpdateJobOutput(&r); err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("UpdateJobOutput done")
}
