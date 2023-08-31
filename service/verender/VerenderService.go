package verender

import (
    "encoding/json"
    "fmt"
    "net/url"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/volcengine/volc-sdk-golang/base"
)

func checkResponseOK(err Error) bool {
    if err.CodeN != CodeOK {
        return false
    }

    return true
}

func packError(resp *ResponseMetadata) error {
    errDetail, _ := json.Marshal(resp)
    return fmt.Errorf(string(errDetail))
}

func NewVerenderInstance() *Verender {
    v := Verender{}
    v.Client = base.NewClient(ServiceInfo, APIInfoList)
    v.Client.ServiceInfo.Credentials.Service = ServiceName
    v.Client.ServiceInfo.Credentials.Region = DefaultRegion

    return &v
}

// SetFtransClientAddr ftransClientAddr must be set when transport by ftrans client
func (v *Verender) SetFtransClientAddr(ftransClientAddr string) {
    v.ftransClientAddr = ftransClientAddr
}

// SetFtransProxyAddr ftransProxyAddr must be set when transport by ftrans proxy
func (v *Verender) SetFtransProxyAddr(ftransProxyAddr string) {
    v.ftransProxyAddr = ftransProxyAddr
}

func (v *Verender) SetFtransServerIsp(isp string) {
    v.ftransServerIsp = isp
}

func (v *Verender) SetFtransClientAclToken(aclToken string) {
    v.ftransClientAclToken = aclToken
}

func (v *Verender) queryHandler(api string, query url.Values) ([]byte, int, error) {
    return v.Client.Query(api, query)
}

func (v *Verender) jsonHandler(api string, query url.Values, body string) ([]byte, int, error) {
    return v.Client.Json(api, query, body)
}

func (w *Workspace) queryHandler(api string, query url.Values) ([]byte, int, error) {
    return w.client.Query(api, query)
}

func (w *Workspace) jsonHandler(api string, query url.Values, body string) ([]byte, int, error) {
    return w.client.Json(api, query, body)
}

func listLocalDir(dir string, depth int) ([]string, error) {
    if depth <= 0 {
        return nil, fmt.Errorf("listLocalDir: dir is too deep")
    }

    var files []string

    fs, err := os.ReadDir(dir)
    if err != nil {
        return nil, fmt.Errorf("listLocalDir: read dir failed [%s]", err.Error())
    }

    if len(fs) <= 0 {
        return []string{dir}, nil
    }

    for _, f := range fs {
        if f.IsDir() {
            subDir, err := listLocalDir(filepath.Join(dir, f.Name()), depth-1)
            if err != nil {
                return nil, fmt.Errorf("ftrans:listLocalDir: list sub dir failed [%s]", err.Error())
            }

            files = append(files, subDir...)
        } else {
            files = append(files, filepath.Join(dir, f.Name()))
        }
    }

    return files, nil
}

func (v *Verender) ListWorkspace(r *ListWorkspaceRequest) (*WorkspaceList, error) {
    if r == nil {
        return nil, ErrInvalidArgs
    }

    q := url.Values{}
    if r.WorkspaceID > 0 {
        q.Set("WorkspaceId", fmt.Sprintf("%d", r.WorkspaceID))
    }
    if r.WorkspaceName != "" {
        q.Set("WorkspaceName", r.WorkspaceName)
    }
    if r.StartPurchaseTime != "" {
        q.Set("StartPurchaseTime", r.StartPurchaseTime)
    }
    if r.EndPurchaseTime != "" {
        q.Set("EndPurchaseTime", r.EndPurchaseTime)
    }
    if r.FuzzyWorkspaceName != "" {
        q.Set("FuzzyWorkspaceName", r.FuzzyWorkspaceName)
    }
    if r.FuzzySearchContent != "" {
        q.Set("FuzzySearchContent", r.FuzzySearchContent)
    }
    if r.PageNum > 0 {
        q.Set("PageNum", fmt.Sprintf("%d", r.PageNum))
    } else {
        q.Set("PageNum", DefaultPageNum)
    }
    if r.PageSize > 0 {
        q.Set("PageSize", fmt.Sprintf("%d", r.PageSize))
    } else {
        q.Set("PageSize", DefaultPageSize)
    }
    if r.ListScope != "" {
        q.Set("ListScope", r.ListScope)
    }
    orderType, ok := ValidOrderTypeWorkspace[r.OrderType]
    if !ok {
        q.Set("OrderType", DefaultOrderTypeWorkspace)
    } else {
        q.Set("OrderType", orderType)
    }
    if _, ok := ValidOrderFieldWorkspace[r.OrderField]; !ok {
        q.Set("OrderBy", DefaultOrderFieldWorkspace)
    } else {
        q.Set("OrderBy", r.OrderField)
    }

    body, code, err := v.queryHandler("ListWorkspace", q)
    if err != nil {
        return nil, err
    }

    if code != HTTPStatusOK {
        return nil, fmt.Errorf("invalid response status: %d", code)
    }

    var resp ListWorkspaceResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return &resp.Result, nil
}

func (v *Verender) GetCurrentUser() (*User, error) {
    body, status, err := v.queryHandler("GetCurrentUser", nil)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, fmt.Errorf("invalid response status: %d", status)
    }

    var resp GetCurrentUserResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result, nil
}

func (v *Verender) GetWorkspace(workspaceID int64) (*Workspace, error) {
    r := ListWorkspaceRequest{
        WorkspaceID: workspaceID,
    }

    resp, err := v.ListWorkspace(&r)
    if err != nil {
        return nil, err
    }

    if resp.Total != 1 {
        return nil, fmt.Errorf("workspace not found")
    }

    w := resp.Workspaces[0]
    w.client = v.Client

    if w.storageAccess == nil || w.storageAccess.ExpiredNSec <= time.Now().Unix() {
        w.storageAccess, err = w.getStorageAccess(v.ftransClientAddr, v.ftransProxyAddr, v.ftransClientAclToken,
            v.ftransServerIsp)
        if err != nil {
            return nil, err
        }
        w.storageAccess.ExpiredNSec = time.Now().Unix() + StorageAccessExpiredNSec
    }

    cellSpecList, err := w.ListCellSpec()
    if err != nil {
        return nil, err
    }

    for _, cs := range cellSpecList {
        if cs.SystemInfo == VerenderOsTypeWindows {
            w.ignoreUploadCase = true
            break
        }
    }

    return w, nil
}

func (w *Workspace) getStorageAccess(ftransClientAddr, ftransProxyAddr, ftransClientAclToken, isp string) (
    *StorageAccess, error) {
    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    body, status, err := w.queryHandler("GetStorageAccess", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, fmt.Errorf("invalid response status: %d", status)
    }

    var resp GetStorageAccessResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    cli, err := NewFtransClient(
        resp.Result.BucketName, resp.Result.FtransCertName, resp.Result.FtransSecurityToken,
        resp.Result.FtransS10Server, resp.Result.FtransCert, resp.Result.FtransKey,
        ftransClientAddr, resp.Result.FtransQuicServer, ftransProxyAddr, isp, ftransClientAclToken)
    if err != nil {
        return nil, err
    }
    resp.Result.FtransClient = cli
    return &resp.Result, nil
}

func (w *Workspace) UploadFile(src, des string) (*FileInfo, error) {
    r := FtransFileRequest{
        LocalFilePath:  src,
        RemoteFilePath: des,
    }

    if w.ignoreUploadCase {
        r.RemoteFilePath = strings.ToLower(r.RemoteFilePath)
    }

    info, err := w.storageAccess.FtransClient.UploadFile(&r)
    if err != nil {
        return nil, err
    }

    fi := FileInfo{
        Key:           info.Name,
        ContentLength: info.Size,
        LastModified:  time.UnixMicro(info.MTime),
    }

    return &fi, nil
}

func (w *Workspace) UploadFolder(srcFolder, desFolder string) error {
    files, err := listLocalDir(srcFolder, 50)
    if err != nil {
        return fmt.Errorf("list local folder [%s] failed [%s]", srcFolder, err.Error())
    }

    for _, f := range files {
        des := filepath.Join(desFolder, f[len(srcFolder):])
        if w.ignoreUploadCase {
            des = strings.ToLower(des)
        }

        if _, err := w.UploadFile(f, des); err != nil {
            return fmt.Errorf("upload folder [%s] failed [%s]", srcFolder, err.Error())
        }
    }

    return nil
}

func (w *Workspace) DownloadFile(src, dst string) error {
    r := FtransFileRequest{
        LocalFilePath:  dst,
        RemoteFilePath: src,
    }

    _, err := w.storageAccess.FtransClient.DownloadFile(&r)
    if err != nil {
        return err
    }

    return nil
}

func (w *Workspace) RemoveFile(filename string) error {
    return w.storageAccess.FtransClient.RemoveFile(filename)
}

// ListFile 返回指定目录下的所有子目录以及文件列表 不递归子文件夹
func (w *Workspace) ListFile(prefix, filter, orderType, orderFiled string, pageNum, pageSize int64) (
    int64, []*FileInfo, []*FileInfo, error) {

    var files []*FileInfo
    var subDirs []*FileInfo

    cli := w.storageAccess.FtransClient
    r := FtransListFileRequest{
        Prefix:     prefix,
        FilterIn:   filter,
        OrderType:  orderType,
        OrderField: orderFiled,
        PageSize:   pageSize,
        PageNum:    pageNum,
    }
    resp, err := cli.ListFile(&r)
    if err != nil {
        return 0, nil, nil, err
    }

    for _, elem := range resp.Records {
        fileInfo := FileInfo{
            Key:           elem.Name,
            ContentLength: elem.Size,
            LastModified:  time.UnixMicro(elem.MTime),
        }
        if elem.Type == "dir" {
            subDirs = append(subDirs, &fileInfo)
        } else {
            files = append(files, &fileInfo)
        }
    }

    return resp.Total, subDirs, files, nil
}

func (w *Workspace) StatFile(filename string) (*FileInfo, error) {
    fi, err := w.storageAccess.FtransClient.StatFile(filename)

    if err != nil {
        return nil, err
    }

    fo := FileInfo{
        Key:           fi.Name,
        ContentLength: fi.Size,
        LastModified:  time.UnixMicro(fi.MTime).UTC(),
    }

    return &fo, nil
}

func (w *Workspace) ListCellSpec() ([]*CellSpec, error) {
    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    body, status, err := w.queryHandler("ListCellSpec", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, fmt.Errorf("invalid response status: %d", status)
    }

    var resp ListCellSpecResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result.CellSpecs, nil
}

func (w *Workspace) AddRenderSetting(rs *RenderSetting) (*AddRenderSettingResult, error) {
    if rs == nil {
        return nil, ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(rs)
    if err != nil {
        return nil, err
    }

    body, status, err := w.jsonHandler("AddRenderSetting", q, string(data))
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, fmt.Errorf("invalid response status: %d", status)
    }

    var resp AddRenderSettingResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result, nil
}

func (w *Workspace) UpdateRenderSetting(rs *RenderSetting) error {
    if rs == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("RenderSettingId", fmt.Sprintf("%d", rs.Id))

    data, err := json.Marshal(rs)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("UpdateRenderSetting", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return fmt.Errorf("invalid response status: %d", status)
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) DeleteRenderSetting(settingId int64) error {
    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("RenderSettingId", fmt.Sprintf("%d", settingId))

    body, status, err := w.queryHandler("DeleteRenderSetting", q)
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return fmt.Errorf("invalid response status: %d", status)
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (v *Verender) ListDccs() ([]Dcc, error) {
    q := url.Values{}

    body, status, err := v.queryHandler("ListDccs", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp ListDccsResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result.Dccs, nil
}

func (v *Verender) ListAccountDccPlugins(r *ListAccountDccPluginsReq) ([]DccPlugin, error) {
    if r == nil {
        return nil, ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("SpecTemplateId", fmt.Sprintf("%d", r.SpecTemplateId))
    if r.Dcc != "" {
        q.Set("Dcc", r.Dcc)
    }
    if r.DccVersion != "" {
        q.Set("DccVersion", r.DccVersion)
    }
    if r.RegionId > 0 {
        q.Set("RegionId", fmt.Sprintf("%d", r.RegionId))
    }

    body, status, err := v.queryHandler("ListAccountDccPlugins", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp ListAccountDccPluginsResp
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result.Plugins, nil
}

func (w *Workspace) ListRenderSetting(dcc string) ([]*RenderSetting, error) {
    q := url.Values{}
    q.Set("AccountId", fmt.Sprintf("%d", w.AccountID))
    q.Set("UserId", fmt.Sprintf("%d", w.UserID))
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("Dcc", dcc)

    body, status, err := w.queryHandler("ListRenderSetting", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp ListRenderSettingResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result[dcc], nil
}

func (w *Workspace) GetRenderSetting(settingID int64) (*RenderSetting, error) {
    q := url.Values{}
    q.Set("AccountId", fmt.Sprintf("%d", w.AccountID))
    q.Set("UserId", fmt.Sprintf("%d", w.UserID))
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("CheckUserId", "false")
    q.Set("Id", fmt.Sprintf("%d", settingID))
    q.Set("WithDeleted", "true")

    body, status, err := w.queryHandler("GetRenderSetting", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp GetRenderSettingResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    if resp.Result == nil || len(resp.Result.Settings) == 0 {
        return nil, fmt.Errorf("render setting not found")
    }

    return resp.Result.Settings[0], nil
}

func (w *Workspace) CreateRenderJob(j *RenderJob) (*RenderJob, error) {
    if j == nil {
        return nil, fmt.Errorf("invalid args j")
    }

    data, err := json.Marshal(j)
    if err != nil {
        return nil, err
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", j.WorkspaceID))
    body, status, err := w.jsonHandler("CreateRenderJob", q, string(data))
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp CreateRenderJobResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result.RenderJob, nil
}

func (w *Workspace) ListRenderJob(r *ListRenderJobRequest) (*ListRenderJobResult, error) {
    if r == nil {
        return nil, ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("AccountId", fmt.Sprintf("%d", w.AccountID))
    q.Set("UserId", fmt.Sprintf("%d", w.UserID))
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    if r.PageNum > 0 {
        q.Set("PageNum", fmt.Sprintf("%d", r.PageNum))
    } else {
        q.Set("PageNum", DefaultPageNum)
    }
    if r.PageSize > 0 {
        q.Set("PageSize", fmt.Sprintf("%d", r.PageSize))
    } else {
        q.Set("PageSize", DefaultPageSize)
    }
    orderType, ok := ValidOrderTypeRenderJob[r.OrderType]
    if !ok {
        q.Set("OrderBy", DefaultOrderTypeRenderJob)
    } else {
        q.Set("OrderBy", orderType)
    }

    if _, ok := ValidOrderFieldRenderJob[r.OrderField]; !ok {
        q.Set("OrderField", DefaultOrderFieldRenderJob)
    } else {
        q.Set("OrderField", r.OrderField)
    }
    q.Set("Keyword", r.Keyword)
    q.Set("Status", strings.Join(r.Status, ","))
    var renderSettings []string
    for _, id := range r.RenderSettings {
        renderSettings = append(renderSettings, fmt.Sprintf("%d", id))
    }
    q.Set("RenderSetting", strings.Join(renderSettings, ","))

    body, status, err := w.queryHandler("ListRenderJob", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp ListRenderJobResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result, nil
}

func (w *Workspace) GetRenderJob(jobID string) (*RenderJob, error) {
    q := url.Values{}
    q.Set("AccountId", fmt.Sprintf("%d", w.AccountID))
    q.Set("UserId", fmt.Sprintf("%d", w.UserID))
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("RenderJobId", jobID)

    body, status, err := w.queryHandler("GetRenderJob", q)
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp GetRenderJobResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result.Job, nil
}

func (w *Workspace) RetryRenderJob(r *RetryJobRequest) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("RenderJobId", r.JobID)

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("RetryRenderJob", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) ResumeRenderJobs(r *BatchJobIDList) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("ResumeRenderJobs", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) StopRenderJobs(r *BatchJobIDList) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("StopRenderJobs", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) DeleteRenderJobs(r *BatchJobIDList) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("DeleteRenderJobs", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) FullSpeedRenderJobs(r *BatchJobIDList) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("FullSpeedRenderJobs", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) AutoFullSpeedRenderJobs(r *BatchJobIDList) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("AutoFullSpeedRenderJobs", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) UpdateRenderJobsPriority(r *BatchJobPriority) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("UpdateRenderJobsPriority", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }

    return nil
}

func (w *Workspace) ListJobOutput(r *ListJobOutputRequest) (*ListJobOutputResult, error) {
    if r == nil {
        return nil, ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("AccountId", fmt.Sprintf("%d", w.AccountID))
    q.Set("UserId", fmt.Sprintf("%d", w.UserID))
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))

    data, err := json.Marshal(r)
    if err != nil {
        return nil, err
    }

    body, status, err := w.jsonHandler("ListJobOutput", q, string(data))
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp ListJobOutputResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result, nil
}

func (w *Workspace) GetJobOutput(r *GetJobOutputRequest) (*GetJobOutputResult, error) {
    if r == nil {
        return nil, ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("JobId", r.JobID)

    data, err := json.Marshal(r)
    if err != nil {
        return nil, err
    }

    body, status, err := w.jsonHandler("GetJobOutput", q, string(data))
    if err != nil {
        return nil, err
    }

    if status != HTTPStatusOK {
        return nil, ErrInvalidResponseStatus
    }

    var resp GetJobOutputResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return nil, err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return nil, packError(&resp.ResponseMetadata)
    }

    return resp.Result, nil
}

func (w *Workspace) UpdateJobOutput(r *UpdateJobOutputRequest) error {
    if r == nil {
        return ErrInvalidArgs
    }

    q := url.Values{}
    q.Set("WorkspaceId", fmt.Sprintf("%d", w.Id))
    q.Set("JobId", r.JobID)

    r.AccountID = w.AccountID
    r.UserID = w.UserID
    r.WorkspaceID = w.Id

    data, err := json.Marshal(r)
    if err != nil {
        return err
    }

    body, status, err := w.jsonHandler("UpdateJobOutput", q, string(data))
    if err != nil {
        return err
    }

    if status != HTTPStatusOK {
        return ErrInvalidResponseStatus
    }

    var resp VoidResponse
    if err := json.Unmarshal(body, &resp); err != nil {
        return err
    }

    if !checkResponseOK(resp.ResponseMetadata.Error) {
        return packError(&resp.ResponseMetadata)
    }
    return nil
}
