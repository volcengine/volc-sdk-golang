package verender

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

func (v *Verender) commonHandler(api string, query url.Values) (string, int, error) {
	respBody, statusCode, err := v.Client.Query(api, query)
	if err != nil {
		return string(respBody), statusCode, err
	}
	return string(respBody), statusCode, nil
}

func packErrorInfo(data ResponseMetaData) error {
	return fmt.Errorf("RequestId=%s Error=%v", data.RequestId, data.Error)
}

func (v *Verender) commonHandlerJson(api string, query url.Values, body string) (string, int, error) {
	respBody, statusCode, err := v.Client.Json(api, query, body)
	if err != nil {
		return string(respBody), statusCode, err
	}
	return string(respBody), statusCode, nil
}

func (w *Workspace) commonHandler(api string, query url.Values) (string, int, error) {
	respBody, statusCode, err := w.Client.Query(api, query)
	if err != nil {
		return string(respBody), statusCode, err
	}
	return string(respBody), statusCode, nil
}

func (w *Workspace) commonHandlerJson(api string, query url.Values, body string) (string, int, error) {
	respBody, statusCode, err := w.Client.Json(api, query, body)
	if err != nil {
		return string(respBody), statusCode, err
	}
	return string(respBody), statusCode, nil
}

func (j *Job) commonHandler(api string, query url.Values) (string, int, error) {
	respBody, statusCode, err := j.Client.Query(api, query)
	if err != nil {
		return string(respBody), statusCode, err
	}
	return string(respBody), statusCode, nil
}

func (j *Job) commonHandlerJson(api string, query url.Values, body string) (string, int, error) {
	respBody, statusCode, err := j.Client.Json(api, query, body)
	if err != nil {
		return string(respBody), statusCode, err
	}
	return string(respBody), statusCode, nil
}

func (v *Verender) GetCurrentUser() (*UserInfo, error) {
	body, _, err := v.commonHandler("GetCurrentUser", nil)
	if err != nil {
		return nil, err
	}

	resp := &GetUserInfoResult{}
	err = json.Unmarshal([]byte(body), resp)
	if err != nil {
		return nil, err
	}

	return &resp.User, nil
}

// PurchaseWorkspace 需要指定WorkspaceName, Description, StorageTotal, ResourcePoolId, CpsId
func (v *Verender) PurchaseWorkspace(w *Workspace) (*Workspace, error) {
	data, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}

	body, _, err := v.commonHandlerJson("PurchaseWorkspace", nil, string(data))
	if err != nil {
		return nil, err
	}

	resp := PurchaseWorkspaceResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.WorkspaceInfo, nil
}

// UpdateWorkspace 只能修改WorkspaceName, Description, StorageTotal, 其中StorageTotal不能比当前小
func (v *Verender) UpdateWorkspace(w *Workspace) (*Workspace, error) {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	data, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}

	body, _, err := v.commonHandlerJson("UpdateWorkspace", query, string(data))
	if err != nil {
		return nil, err
	}

	resp := PurchaseWorkspaceResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.WorkspaceInfo, nil
}

func (v *Verender) DeleteWorkspace(w *Workspace) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))

	body, _, err := v.commonHandler("DeleteWorkspace", query)
	if err != nil {
		return err
	}

	resp := DeleteWorkspaceResp{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return packErrorInfo(resp.MetaData)
	}

	return nil
}

// ListResourcePools 此接口用于购买工作区时指定ResourcePoolId和CpsId参数
func (v *Verender) ListResourcePools() (*[]ResourcePool, error) {
	body, _, err := v.commonHandler("ListResourcePools", nil)
	if err != nil {
		return nil, err
	}

	resp := ListResourcePoolResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.ResourcePools, nil
}

// GetWorkspaceLimit 返回当前登录用户可创建的工作区个数和单个工作区可使用的存储空间上限
func (v *Verender) GetWorkspaceLimit() (*WorkspaceLimit, error) {
	body, _, err := v.commonHandler("GetWorkspaceLimit", nil)
	if err != nil {
		return nil, err
	}

	resp := GetWorkspaceLimitResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.WorkspaceLimit, nil
}

func (v *Verender) GetHardwareSpecifications(workspaceId int) (*[]WorkspaceHardwareSpecification, error) {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(workspaceId))
	body, _, err := v.commonHandler("GetHardwareSpecifications", query)
	if err != nil {
		return nil, err
	}

	resp := GetWorkspaceHardwareSpecificationResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.HardwareSpecs, nil
}

func checkTimeType(r *StatisticsRequest) error {
	if r.TimeType == "" {
		return fmt.Errorf("TimeType can't be null")
	}
	//TODO 根据时间范围设置合适的粒度, 需要和API协商后制定
	return nil
}

func (v *Verender) GetAccountStatistics(r *StatisticsRequest) (*AccountStatistics, error) {
	if err := checkTimeType(r); err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("TimeType", r.TimeType)
	if r.StartTime != "" && r.EndTime != "" {
		query.Set("StartTime", r.StartTime)
		query.Set("EndTime", r.EndTime)
	}

	data, err := json.Marshal(r.Filter)
	if err != nil {
		return nil, err
	}

	body, _, err := v.commonHandlerJson("GetAccountStatistics", query, string(data))
	if err != nil {
		return nil, err
	}

	resp := GetAccountStatisticsResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.AccountStat, nil
}

func (v *Verender) GetAccountStatisticDetails(r *StatisticsRequest) (*AccountStatisticDetails, error) {
	query := url.Values{}
	query.Set("StartTime", r.StartTime)
	query.Set("EndTime", r.EndTime)

	if r.PageSize > 0 {
		query.Set("PageNum", strconv.FormatInt(r.PageNum, 10))
		query.Set("PageSize", strconv.FormatInt(r.PageSize, 10))
	} else {
		query.Set("PageNum", "1")
		query.Set("PageSize", "10")
	}
	if r.OrderField != "" {
		query.Set("OrderField", r.OrderField)
	}
	if r.OrderBy != "" {
		query.Set("OrderBy", r.OrderBy)
	}

	data, err := json.Marshal(r.Filter)
	if err != nil {
		return nil, err
	}

	body, _, err := v.commonHandlerJson("GetAccountStatisticDetails", query, string(data))
	if err != nil {
		return nil, err
	}

	resp := GetAccountStatisticDetailsResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.Details, nil
}

func (v *Verender) DownloadStatisticDetails(r *StatisticsRequest) error {
	query := url.Values{}
	if r.StartTime != "" {
		query.Set("StartTime", r.StartTime)
	}
	if r.EndTime != "" {
		query.Set("EndTime", r.EndTime)
	}
	if r.PageSize > 0 {
		query.Set("PageNum", strconv.FormatInt(r.PageNum, 10))
		query.Set("PageSize", strconv.FormatInt(r.PageSize, 10))
	} else {
		query.Set("PageNum", "1")
		query.Set("PageSize", "10")
	}
	if r.OrderField != "" {
		query.Set("OrderField", r.OrderField)
	}
	if r.OrderBy != "" {
		query.Set("OrderBy", r.OrderBy)
	}

	resp, _, err := v.commonHandler("DownloadStatisticDetails", query)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(r.FileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(resp)
	if err != nil {
		return err
	}
	return nil
}

func (v *Verender) GetJobOverallStatistics() (*JobOverallStatistics, error) {
	body, _, err := v.commonHandler("GetJobOverallStatistics", nil)
	if err != nil {
		return nil, err
	}

	resp := GetJobOverallStatisticsResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.Stat, nil
}

func (v *Verender) ListWorkspaces(r *ListWorkspaceRequest) (*WorkspaceList, error) {
	query := url.Values{}
	if r.WorkspaceId > 0 {
		query.Set("WorkspaceId", strconv.FormatInt(r.WorkspaceId, 10))
	}

	if r.WorkspaceName != "" {
		query.Set("WorkspaceName", r.WorkspaceName)
	}

	if r.PageSize > 0 {
		query.Set("PageNum", strconv.FormatInt(r.PageNum, 10))
		query.Set("PageSize", strconv.FormatInt(r.PageSize, 10))
	}
	if r.OrderBy != "" {
		query.Set("OrderBy", r.OrderBy)
	}
	if r.OrderType != "" {
		query.Set("OrderType", r.OrderType)
	}
	if r.StartPurchaseTime != "" {
		query.Set("StartPurchaseTime", r.StartPurchaseTime)
	}
	if r.EndPurchaseTime != "" {
		query.Set("EndPurchaseTime", r.EndPurchaseTime)
	}
	if r.FuzzyWorkspaceName != "" {
		query.Set("FuzzyWorkspaceName", r.FuzzyWorkspaceName)
	}
	if r.FuzzySearchContent != "" {
		query.Set("FuzzySearchContent", r.FuzzySearchContent)
	}

	resBody, _, err := v.commonHandler("ListWorkspaces", query)
	if err != nil {
		return nil, err
	}

	resp := ListWorkspaceResponse{}
	if err = json.Unmarshal([]byte(resBody), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.Workspaces, nil
}

func (v *Verender) GetWorkspace(workspaceId int64) (*Workspace, error) {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.FormatInt(workspaceId, 10))
	resp, _, err := v.commonHandler("GetStorageAccess", query)
	if err != nil {
		return nil, err
	}
	res := GetStorageAccessResponse{}
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		return nil, err
	}
	if res.Error.CodeN != 0 {
		return nil, fmt.Errorf("%v", res.Error)
	}
	cli, err := GetMinioClient(res.AccessKey, res.SecretKey, res.BucketToken, res.ClusterEndpoint)
	if err != nil {
		return nil, err
	}
	return &Workspace{
		WorkspaceId:     res.WorkspaceId,
		ClusterEndpoint: res.ClusterEndpoint,
		BucketName:      res.BucketName,
		BucketToken:     res.BucketToken,
		VolumeName:      res.VolumeName,
		ClusterName:     res.ClusterName,
		AccessKey:       res.AccessKey,
		SecretKey:       res.SecretKey,
		Time:            time.Now(),
		MinioClient:     cli,
		Client:          v.Client,
	}, nil
}

func (w *Workspace) GetStorageAccess() Workspace {
	return *w
}

func listLocalDir(dir string) ([]string, error) {
	var files []string
	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	}

	err := filepath.Walk(dir, walkFunc)

	return files, err
}

func getBucketPath(localPath string, fileName string) string {
	// 将windows的目录 c:\abc\1.ma 改为 c:/abc/1.ma
	localPath = filepath.ToSlash(localPath)
	fileName = filepath.ToSlash(fileName)

	lastPath := localPath
	segs := strings.Split(localPath, "/")
	if len(segs) > 1 {
		lastPath = segs[len(segs)-1]
	}

	ok := strings.HasPrefix(fileName, localPath)
	if ok {
		fileName = fileName[len(localPath)+1:]
	}

	// 保留文件夹最后一级目录 + 子目录 + 文件名作为objectName
	if lastPath[len(lastPath)-1] != '/' && fileName[0] != '/' {
		return lastPath + "/" + fileName
	} else if lastPath[len(lastPath)-1] == '/' && fileName[0] == '/' {
		return lastPath + fileName[1:]
	}

	return lastPath + fileName
}

func getPartSize(fileName string) (uint64, error) {
	info, err := os.Stat(fileName)
	if err != nil {
		return MINPARTSIZE, fmt.Errorf("Stat file %s Error %s", fileName, err.Error())
	}

	fileSize := info.Size()
	if fileSize > MAXOBJECTSIZE {
		return MINPARTSIZE, fmt.Errorf("File %s too large, size=%d limit=%d", fileName, fileSize, MAXOBJECTSIZE)
	}
	if fileSize <= MINPARTSIZE {
		return MINPARTSIZE, nil
	}
	// 分片大小为MINPARTSIZE的整数倍
	partSize := float64(fileSize) / MAXPARTS
	partSize = math.Ceil(float64(partSize)/MINPARTSIZE) * MINPARTSIZE

	return uint64(partSize), nil
}

// UploadFolder 上传本地目录(包括子目录和文件)到边缘存储, 且保留目录层级信息
func (w *Workspace) UploadFolder(localPath string, suffix string) ([]string, error) {
	var res []string

	if w == nil {
		return res, fmt.Errorf("invalid workspace")
	}

	_, err := os.Stat(localPath)
	if err != nil {
		return res, err
	}

	files, _ := listLocalDir(localPath)

	for _, file := range files {
		if "" != suffix && !strings.HasSuffix(file, suffix) {
			continue
		}

		partSize, err := getPartSize(file)
		if err != nil {
			return res, fmt.Errorf("upload file %s error %s", file, err.Error())
		}
		object := getBucketPath(localPath, file)
		info, err := w.MinioClient.FPutObject(context.Background(), w.BucketName, object,
			file, minio.PutObjectOptions{PartSize: partSize})
		if err != nil {
			return res, fmt.Errorf("upload file %s error %s", file, err.Error())
		}

		location := info.Location
		if location == "" {
			location = w.MinioClient.EndpointURL().String() + "/" + w.BucketName + "/" + info.Key
		}
		res = append(res, location)
	}

	return res, nil
}

// UploadFile 上传本地文件到边缘存储的根目录，只保留文件名
func (w *Workspace) UploadFile(filePath string, fileType string) (string, error) {
	if fileType == "" {
		fileType = "asset"
	}
	_, ok := AllowedFileTypes[fileType]
	if !ok {
		return "", fmt.Errorf("fileType must be one of %v", AllowedFileTypes)
	}
	partSize, err := getPartSize(filePath)
	if err != nil {
		return "", fmt.Errorf("upload file %s error %s", filePath, err.Error())
	}
	object := path.Base(filepath.ToSlash(filePath))
	info, err := w.MinioClient.FPutObject(context.Background(), w.BucketName, object, filePath,
		minio.PutObjectOptions{PartSize: partSize})
	if err != nil {
		return "", err
	}

	location := info.Location
	if location == "" {
		location = w.MinioClient.EndpointURL().String() + "/" + w.BucketName + "/" + info.Key
	}

	return location, nil
}

func (w *Workspace) DownloadFile(local string, target string) error {
	err := w.MinioClient.FGetObject(context.Background(), w.BucketName, target, local, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (w *Workspace) StatObject(key string) (*ObjectInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	obj, err := w.MinioClient.StatObject(ctx, w.BucketName, key, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}

	objInfo := &ObjectInfo{
		Etag:          obj.ETag,
		Key:           obj.Key,
		ContentLength: obj.Size,
		ContentType:   obj.ContentType,
		LastModified:  obj.LastModified,
	}

	return objInfo, nil
}

func (w *Workspace) ListObject(prefix, startAfter string, maxKeys int) (*[]string, *[]string, string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var files []string
	var subDirectories []string

	objectCh := w.MinioClient.ListObjects(ctx, w.BucketName, minio.ListObjectsOptions{
		Prefix:     prefix,
		MaxKeys:    maxKeys, // don't take effect now
		Recursive:  false,
		StartAfter: startAfter,
	})

	nextStart := ""
	for obj := range objectCh {
		if obj.Err != nil {
			continue
		}
		if obj.Key[len(obj.Key)-1] == '/' {
			subDirectories = append(subDirectories, obj.Key)
		} else {
			files = append(files, obj.Key)
		}
		nextStart = obj.Key
	}

	return &files, &subDirectories, nextStart, nil
}

func (w *Workspace) GetJob(jobId string) (*Job, error) {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)
	respBody, _, err := w.commonHandler("GetRenderJob", query)
	if err != nil {
		return nil, err
	}

	res := GetJobResponse{}
	if err = json.Unmarshal([]byte(respBody), &res); err != nil {
		return nil, err
	}

	if res.Error.CodeN != StatusOk {
		return nil, packErrorInfo(res.ResponseMetaData)
	}

	job := &Job{
		JobInfo: res.Job,
		Client:  w.Client,
	}

	return job, err
}

func (w *Workspace) ListJobs(r *ListJobRequest) (*JobList, error) {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	orderBy, ok := OrderTypeMap[r.OrderBy]
	if !ok {
		orderBy = "1"
	}
	query.Set("OrderBy", orderBy)
	if r.PageSize > 0 {
		query.Set("PageNum", strconv.Itoa(r.PageNum))
		query.Set("PageSize", strconv.Itoa(r.PageSize))
	} else {
		query.Set("PageNum", "1")
		query.Set("PageSize", "10")
	}
	query.Set("Tryout", strconv.Itoa(r.Tryout))
	query.Set("Keyword", r.Keyword)
	query.Set("Status", r.Status)
	// JobType支持maya_redshift, maya_arnold, c4d_redshift, blender, octane, 多个任务类型之间使用英文逗号分隔
	query.Set("JobType", r.JobType)

	body, _, err := w.commonHandler("ListJobs", query)
	if err != nil {
		return nil, err
	}

	resp := &ListJobResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.JobsData, nil
}

// CreateJob 创建任务之前需要先上传工程文件到边缘存储
func (w *Workspace) CreateJob(job *JobInfo) (*JobInfo, error) {
	jobString, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	body, _, err := w.commonHandlerJson("CreateJob", query, string(jobString))
	if err != nil {
		return nil, err
	}

	resp := &CreateJobResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.Metadata.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.Metadata)
	}

	return &resp.JobDetail.Job, nil
}

func (w *Workspace) PauseJob(jobId string) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)

	body, _, err := w.commonHandler("PauseJob", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) ResumeJob(jobId string) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)

	body, _, err := w.commonHandler("StartJob", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) StopJob(jobId string) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)

	body, _, err := w.commonHandler("StopJob", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) DeleteJob(jobId string) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)

	body, _, err := w.commonHandler("DeleteJob", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) RetryJob(jobId string) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)

	body, _, err := w.commonHandler("RetryJob", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

// FullSpeedJob 试渲染完成后, 需要继续渲染剩余帧时使用FullSpeedJob
func (w *Workspace) FullSpeedJob(jobId string) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)

	body, _, err := w.commonHandler("FullSpeedJob", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) EditJob(job *JobInfo) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", job.JobId)

	jobString, err := json.Marshal(job)
	if err != nil {
		return err
	}

	body, _, err := w.commonHandlerJson("EditJob", query, string(jobString))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}
	return nil
}

func (w *Workspace) UpdateJobPriority(jobId string, priority int) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", jobId)
	query.Set("Priority", strconv.Itoa(priority))

	body, _, err := w.commonHandler("SetJobPriority", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}
	return nil
}

func (w *Workspace) PauseJobs(r *BatchJobRequest) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))

	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	body, _, err := w.commonHandlerJson("PauseJobs", query, string(data))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) ResumeJobs(r *BatchJobRequest) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))

	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	body, _, err := w.commonHandlerJson("ResumeJobs", query, string(data))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) StopJobs(r *BatchJobRequest) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))

	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	body, _, err := w.commonHandlerJson("StopJobs", query, string(data))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return err
	}

	return nil
}

func (w *Workspace) DeleteJobs(r *BatchJobRequest) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))

	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	body, _, err := w.commonHandlerJson("DeleteJobs", query, string(data))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) FullSpeedJobs(r *BatchJobRequest) error {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))

	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	body, _, err := w.commonHandlerJson("FullSpeedJobs", query, string(data))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (w *Workspace) GetLayerFrames(r *GetLayerFramesRequest) (*[]LayerFrameList, error) {
	query := url.Values{}
	query.Set("WorkspaceId", strconv.Itoa(w.WorkspaceId))
	query.Set("RenderJobId", r.RenderJobId)

	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	body, _, err := w.commonHandlerJson("GetLayerFrames", query, string(data))
	if err != nil {
		return nil, err
	}

	resp := GetLayerFramesResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.LayerFrames, nil
}

func (v *Verender) ListMyMessages(r *ListMessageRequest) (*MessageList, error) {
	query := url.Values{}
	if r.PageSize > 0 {
		query.Set("PageNum", strconv.FormatInt(r.PageNum, 10))
		query.Set("PageSize", strconv.FormatInt(r.PageSize, 10))
	} else {
		query.Set("PageNum", "1")
		query.Set("PageSize", "10")
	}

	body, _, err := v.commonHandler("ListMyMessages", query)
	if err != nil {
		return nil, err
	}

	resp := ListMessageResponse{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	if resp.MetaData.Error.CodeN != StatusOk {
		return nil, packErrorInfo(resp.MetaData)
	}

	return &resp.Messages, nil
}

func (v *Verender) MarkMessageAsRead(id int) error {
	query := url.Values{}
	query.Set("MessageId", strconv.Itoa(id))

	body, _, err := v.commonHandler("MarkMessageAsRead", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (v *Verender) BatchMarkMessagesAsRead(r *MessageIdList) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	body, _, err := v.commonHandlerJson("BatchMarkMessagesAsRead", nil, string(data))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (v *Verender) MarkAllMessagesAsRead() error {
	body, _, err := v.commonHandler("MarkAllMessagesAsRead", nil)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return err
	}

	return nil
}

func (v *Verender) DeleteMessage(id int) error {
	query := url.Values{}
	query.Set("MessageId", strconv.Itoa(id))
	body, _, err := v.commonHandler("DeleteMessage", query)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (v *Verender) BatchDeleteMessages(r *MessageIdList) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	body, _, err := v.commonHandlerJson("BatchDeleteMessages", nil, string(data))
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (v *Verender) DeleteAllMessages() error {
	body, _, err := v.commonHandler("DeleteAllMessages", nil)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (v *Verender) DeleteAllReadMessages() error {
	body, _, err := v.commonHandler("DeleteAllReadMessages", nil)
	if err != nil {
		return err
	}

	resp := ResponseMetaData{}
	if err = json.Unmarshal([]byte(body), &resp); err != nil {
		return err
	}

	if resp.Error.CodeN != StatusOk {
		return packErrorInfo(resp)
	}

	return nil
}

func (j *Job) GetJobInfo() Job {
	return *j
}
