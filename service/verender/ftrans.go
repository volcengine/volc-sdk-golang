package verender

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	LowestSupportedVersion = "v1.7.2.2"

	IspCt = "ct"
	IspUn = "un"
	IspCm = "cm"

	FtransS10P2Path = "/s10/p2/ftrans"

	FtransDefaultMountPath = "/var/mnt"

	FtransOrderTypeAsc       = "asc"
	FtransOrderTypeDesc      = "desc"
	FtransOrderFieldFileName = "name"
	FtransOrderFieldMTime    = "mtime"

	FtransDefaultPageNum         = int64(1)
	FtransDefaultPageSize        = int64(10)
	FtransDefaultPartSize        = int64(4 << 20)
	FtransDefaultPartConcurrency = (4)

	FtransTaskStatusWaiting  = int8(0)
	FtransTaskStatusDoing    = int8(1)
	FtransTaskStatusFinished = int8(2)
	FtransTaskStatusFailed   = int8(3)
)

type FtransClient struct {
	BucketName          string       `json:"bucket_name"`
	FtransS10ServerAddr string       `json:"ftrans_s10_server"`
	FtransS10ServerName string       `json:"ftrans_s10_server_name"`
	FtransAclToken      string       `json:"ftrans_acl_token"`
	hc                  *http.Client `json:"-"`
}

type FtransUploadFileRequest struct {
	LocalFilePath    string
	RemoteFilePath   string
	PartSize         int64
	ConsistencyCheck bool
	Md5Check         bool
}

type FtransDownloadFileRequest struct {
	LocalFilePath    string
	RemoteFilePath   string
	PartSize         int64
	ConsistencyCheck bool
	Md5Check         bool
}

type FtransFileInfo struct {
	Name         string `json:"name"`
	Type         string `json:"type,omitempty"`
	Size         int64  `json:"size"`
	LastModified string `json:"mtime"`
	MTime        int64  `json:"-"`
	Md5          string `json:"md5,omitempty"`
}

type FtransStatFileRequest struct {
	FileName string `json:"file_name"`
	WithMd5  bool   `json:"with_md5"`
}

type FtransListFileRequest struct {
	Prefix     string `json:"prefix"`
	FilterIn   string `json:"filter_in"`
	OrderType  string `json:"order_type"`
	OrderField string `json:"order_field"`
	PageSize   int64  `json:"page_size"`
	PageNum    int64  `json:"page_num"`
}

type FtransListFileResponse struct {
	Total   int64             `json:"total"`
	Records []*FtransFileInfo `json:"records"`
}

type ftransFileTask struct {
	LocalFilePath    string
	RemoteFilePath   string
	TempFilePath     string
	FileSize         int64
	PartSize         int64
	MtimeUSec        int64
	Md5              string
	ConsistencyCheck bool
	Md5Check         bool
	File             *os.File
	PartTasks        []*ftransPartTask
	WaitingParts     int32
	DoingParts       int32
	Status           int8
}

type ftransPartTask struct {
	Offset     int64
	Size       int64
	Status     int8
	Data       []byte
	parentTask *ftransFileTask
	f          *os.File
}

func parseS10Server(s10Server, isp string) (string, error) {
	addrIspMap := map[string]string{}
	curVersion := LowestSupportedVersion

	for _, seg := range strings.Split(s10Server, ";") {
		elems := strings.Split(seg, ":")
		sz := len(elems)

		if sz < 3 {
			continue
		}

		if strings.Compare(elems[0], curVersion) < 0 {
			continue
		}

		addr := ""
		if elems[1] == IspCt || elems[1] == IspUn || elems[1] == IspCm {
			addr = strings.Join(elems[2:], ":")
			addrIspMap[elems[1]] = addr
		} else {
			addr = strings.Join(elems[1:], ":")
			addrIspMap[IspCt] = addr
			addrIspMap[IspUn] = addr
			addrIspMap[IspCm] = addr
		}

		curVersion = elems[0]
	}

	addr, ok := addrIspMap[isp]
	if !ok {
		return "", fmt.Errorf("no addr of isp %s found", isp)
	}

	return addr, nil
}

func NewFtransClient(bucketName, isp, ftransS10Server, ftransServerName, ftransACLToken, cert, key string) (
	*FtransClient, error) {

	addr, err := parseS10Server(ftransS10Server, isp)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName: ftransServerName,
		},
	}

	if cert != "" && key != "" {
		crt, err := tls.X509KeyPair([]byte(cert), []byte(key))
		if err != nil {
			return nil, fmt.Errorf("invalid cert or key %s", err.Error())
		}

		tr.TLSClientConfig.Certificates = []tls.Certificate{crt}
	}

	hc := &http.Client{
		Transport: tr,
	}

	fc := &FtransClient{
		BucketName:          bucketName,
		FtransS10ServerAddr: addr,
		FtransS10ServerName: ftransServerName,
		FtransAclToken:      ftransACLToken,
		hc:                  hc,
	}

	return fc, nil
}

func (fc *FtransClient) UploadFile(r *FtransUploadFileRequest) (*FtransFileInfo, error) {
	stat, err := os.Stat(r.LocalFilePath)
	if err != nil {
		return nil, fmt.Errorf("stat local file %s failed, %s", r.LocalFilePath, err.Error())
	}

	if stat.IsDir() {
		ffi := &FtransFileInfo{
			Name:  r.RemoteFilePath,
			Size:  0,
			MTime: stat.ModTime().UnixMicro(),
		}

		if err := fc.makeDir(ffi.Name, ffi.MTime); err != nil {
			return nil, err
		}

		return ffi, nil
	}

	partSize := FtransDefaultPartSize
	f, err := os.OpenFile(r.LocalFilePath, os.O_RDONLY, 0755)
	if err != nil {
		return nil, fmt.Errorf("open local file %s failed, %s", r.LocalFilePath, err.Error())
	}
	defer f.Close()

	fft := &ftransFileTask{
		LocalFilePath:    r.LocalFilePath,
		RemoteFilePath:   r.RemoteFilePath,
		TempFilePath:     fmt.Sprintf("%s_%s", r.RemoteFilePath, uuid.New().String()),
		FileSize:         stat.Size(),
		PartSize:         partSize,
		MtimeUSec:        stat.ModTime().UnixMicro(),
		Md5:              "",
		ConsistencyCheck: r.ConsistencyCheck,
		Md5Check:         r.Md5Check,
		File:             f,
		PartTasks:        nil,
		DoingParts:       0,
		Status:           FtransTaskStatusDoing,
	}

	if fft.FileSize <= 0 {
		fpt := ftransPartTask{
			Offset:     0,
			Size:       0,
			Status:     0,
			Data:       nil,
			parentTask: fft,
			f:          nil,
		}

		if err := fc.uploadFilePart(&fpt); err != nil {
			if err := fc.RemoveFile(fft.TempFilePath); err != nil {
				return nil, fmt.Errorf("upload empty file %s failed, %s", fft.LocalFilePath, err.Error())
			}
			return nil, fmt.Errorf("upload empty file %s failed %s", fft.LocalFilePath, err.Error())
		}

		if err := fc.renameFile(fft.TempFilePath, fft.RemoteFilePath); err != nil {
			return nil, fmt.Errorf("upload empty file %s failed, %s", fft.LocalFilePath, err.Error())
		}

		return &FtransFileInfo{
			Name:  fft.RemoteFilePath,
			Size:  0,
			MTime: fft.MtimeUSec,
		}, nil
	}

	fc.splitPartTask(fft)
	fft.DoingParts = int32(len(fft.PartTasks))

	if r.Md5Check {
		fft.Md5 = fc.getLocalFileMd5(fft.LocalFilePath)
	}

	statReq := FtransStatFileRequest{
		FileName: r.RemoteFilePath,
		WithMd5:  r.Md5Check,
	}
	fi, err := fc.StatFile(&statReq)
	if err == nil {
		if fi.Size == fft.FileSize && fi.MTime == fft.MtimeUSec {
			resp := &FtransFileInfo{
				Name:  r.RemoteFilePath,
				Size:  fi.Size,
				MTime: fi.MTime,
			}

			if !r.Md5Check {
				return resp, nil
			}

			if strings.Compare(fi.Md5, fft.Md5) == 0 {
				resp.Md5 = fi.Md5
				return resp, nil
			}
		}
	}

	freeCh := make(chan bool, FtransDefaultPartConcurrency)
	taskCh := make(chan *ftransPartTask)
	callbackCh := make(chan *ftransPartTask)

	for idx := 0; idx < FtransDefaultPartConcurrency; idx++ {
		go fc.uploadPart(freeCh, taskCh, callbackCh)
	}

	for {
		select {
		case <-freeCh:
			if fft.WaitingParts <= 0 {
				continue
			}

			if fft.ConsistencyCheck {
				fi, err := os.Stat(r.LocalFilePath)
				if err != nil {
					return nil, fmt.Errorf("stat file %s failed, %s", fft.LocalFilePath, err.Error())
				}

				if fi.Size() != fft.FileSize && fi.ModTime().UnixMicro() != fft.MtimeUSec {
					return nil, fmt.Errorf("file %s has been changed during uploading", fft.LocalFilePath)
				}
			}

			fpt := fft.PartTasks[fft.WaitingParts-1]
			_, err = fft.File.ReadAt(fpt.Data, fpt.Offset)
			if err != nil {
				return nil, fmt.Errorf("read part(%d, %d) for %s failed, %s", fpt.Offset, fpt.Size,
					fft.LocalFilePath, err.Error())
			}
			taskCh <- fpt
			fft.WaitingParts--
		case fpt := <-callbackCh:
			fft.DoingParts--
			if fpt.Status != FtransTaskStatusFinished {
				fft.Status = fpt.Status
			} else {
				if fft.DoingParts <= 0 {
					fft.Status = FtransTaskStatusFinished
					break
				}
			}
		}

		if fft.Status != FtransTaskStatusDoing {
			break
		}
	}

	if fft.Status != FtransTaskStatusFinished {
		if err := fc.RemoveFile(fft.TempFilePath); err != nil {
			return nil, fmt.Errorf("remove temp file failed, %s", err.Error())
		}
		return nil, fmt.Errorf("upload file %s failed", fft.LocalFilePath)
	}

	if fft.Md5Check {
		md5Str, err := fc.getFileMd5(fft.TempFilePath)
		if err != nil {
			if err := fc.RemoveFile(fft.TempFilePath); err != nil {
				return nil, fmt.Errorf("remove temp file failed, %s", err.Error())
			}

			return nil, fmt.Errorf("md5 check failed for %s, local=%s, remote=%s", fft.LocalFilePath,
				fft.Md5, md5Str)
		}
	}

	if err := fc.renameFile(fft.TempFilePath, fft.RemoteFilePath); err != nil {
		return nil, err
	}

	resp := &FtransFileInfo{
		Name:  fft.RemoteFilePath,
		Size:  fft.FileSize,
		MTime: fft.MtimeUSec,
		Md5:   fft.Md5,
	}

	return resp, nil
}

func (fc *FtransClient) DownloadFile(r *FtransDownloadFileRequest) (*FtransFileInfo, error) {
	partSize := FtransDefaultPartSize

	if err := fc.fileExist(r.RemoteFilePath); err != nil {
		return nil, fmt.Errorf("file %s not found", r.RemoteFilePath)
	}

	size, mtime, err := fc.getFileSize(r.RemoteFilePath)
	if err != nil {
		return nil, fmt.Errorf("stat file %s failed, %s", r.RemoteFilePath, err.Error())
	}

	md5Str := ""
	if r.Md5Check {
		md5Str, err = fc.getFileMd5(r.RemoteFilePath)
		if err != nil {
			return nil, fmt.Errorf("get md5 of file %s failed, %s", r.RemoteFilePath, err.Error())
		}
	}

	stat, err := os.Stat(r.LocalFilePath)
	if err == nil && !stat.IsDir() {
		if stat.ModTime().UnixNano() == mtime && stat.Size() == size {
			if r.Md5Check {
				localMd5 := fc.getLocalFileMd5(r.LocalFilePath)
				if localMd5 == md5Str {
					return &FtransFileInfo{
						Name:  r.RemoteFilePath,
						Size:  size,
						MTime: mtime,
						Md5:   md5Str,
					}, nil
				}
			} else {
				return &FtransFileInfo{
					Name:  r.RemoteFilePath,
					Size:  size,
					MTime: mtime,
					Md5:   "",
				}, nil
			}
		}

		_ = os.Remove(r.LocalFilePath)
	}

	fft := &ftransFileTask{
		LocalFilePath:    r.LocalFilePath,
		RemoteFilePath:   r.RemoteFilePath,
		TempFilePath:     fmt.Sprintf("%s_%s", r.LocalFilePath, uuid.New().String()),
		FileSize:         size,
		PartSize:         partSize,
		MtimeUSec:        mtime,
		Md5:              md5Str,
		ConsistencyCheck: r.ConsistencyCheck,
		Md5Check:         r.Md5Check,
		File:             nil,
		PartTasks:        nil,
		DoingParts:       0,
		Status:           FtransTaskStatusDoing,
	}

	fc.splitPartTask(fft)
	fft.DoingParts = int32(len(fft.PartTasks))

	saveDir := filepath.Dir(fft.TempFilePath)
	st, err := os.Stat(saveDir)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, 0755); err != nil {
			return nil, fmt.Errorf("make dir %s failed %s", saveDir, err.Error())
		}
	} else if !st.IsDir() {
		return nil, fmt.Errorf("%s is a file, please remove it first", saveDir)
	}

	f, err := os.OpenFile(fft.TempFilePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil, fmt.Errorf("open local file %s failed %s", fft.TempFilePath, err.Error())
	}
	fft.File = f
	defer func(fft *ftransFileTask) {
		if fft.File != nil {
			fft.File.Close()
		}
	}(fft)

	if fft.FileSize <= 0 {
		fft.File.Close()
		if err := os.Rename(fft.TempFilePath, fft.LocalFilePath); err != nil {
			return nil, fmt.Errorf("rename file %s into %s failed %s", fft.TempFilePath, fft.LocalFilePath, err.Error())
		}
		return &FtransFileInfo{
			Name:  fft.RemoteFilePath,
			Size:  fft.FileSize,
			MTime: fft.MtimeUSec,
			Md5:   fft.Md5,
		}, nil
	}

	freeCh := make(chan bool, FtransDefaultPartConcurrency)
	taskCh := make(chan *ftransPartTask)
	callbackCh := make(chan *ftransPartTask)

	for idx := 0; idx < FtransDefaultPartConcurrency; idx++ {
		go fc.downloadPart(freeCh, taskCh, callbackCh)
	}

	for {
		select {
		case <-freeCh:
			if fft.WaitingParts <= 0 {
				continue
			}

			fpt := fft.PartTasks[fft.WaitingParts-1]
			taskCh <- fpt
			fft.WaitingParts--
		case fpt := <-callbackCh:
			if fpt.Status != FtransTaskStatusFinished {
				fft.Status = fpt.Status
				continue
			}

			fft.DoingParts--

			if fft.DoingParts == 0 {
				fft.Status = FtransTaskStatusFinished
			}

			_, err := fft.File.WriteAt(fpt.Data, fpt.Offset)
			if err != nil {
				return nil, fmt.Errorf("write file %s at (%d, %d) failed, %s",
					fft.TempFilePath, fpt.Offset, fpt.Size, err.Error())
			}
		}

		if fft.Status != FtransTaskStatusDoing {
			break
		}
	}

	if fft.Status != FtransTaskStatusFinished {
		if err := os.Remove(fft.TempFilePath); err != nil {
			return nil, fmt.Errorf("remove temp file failed, %s", err.Error())
		}
		return nil, fmt.Errorf("download file %s failed", fft.LocalFilePath)
	}

	if fft.Md5Check {
		md5Str := fc.getLocalFileMd5(fft.TempFilePath)

		if strings.Compare(md5Str, fft.Md5) != 0 {
			if err := os.Remove(fft.TempFilePath); err != nil {
				return nil, fmt.Errorf("remove temp file failed, %s", err.Error())
			}
			return nil, fmt.Errorf("md5 check failed for %s, local=%s, remote=%s", fft.RemoteFilePath,
				md5Str, fft.Md5)
		}
	}

	if fft.File != nil {
		fft.File.Close()
		fft.File = nil
	}
	if err := os.Rename(fft.TempFilePath, fft.LocalFilePath); err != nil {
		return nil, fmt.Errorf("rename file %s into %s failed, %s", fft.TempFilePath, fft.LocalFilePath, err.Error())
	}

	_ = os.Chtimes(fft.LocalFilePath, time.Now(), time.UnixMicro(fft.MtimeUSec))

	resp := &FtransFileInfo{
		Name:  fft.RemoteFilePath,
		Size:  fft.FileSize,
		MTime: fft.MtimeUSec,
		Md5:   fft.Md5,
	}

	return resp, nil
}

func (fc *FtransClient) StatFile(r *FtransStatFileRequest) (*FtransFileInfo, error) {
	size, mtime, err := fc.getFileSize(r.FileName)
	if err != nil {
		return nil, fmt.Errorf("failed to get size for %s, %s", r.FileName, err.Error())
	}

	fi := FtransFileInfo{
		Name:  r.FileName,
		Size:  size,
		MTime: mtime,
		Md5:   "",
	}

	if r.WithMd5 {
		md5, err := fc.getFileMd5(r.FileName)
		if err != nil {
			return nil, fmt.Errorf("failed to get md5 for %s, %s", r.FileName, err.Error())
		}

		fi.Md5 = md5
	}

	return &fi, nil
}

func (fc *FtransClient) ListFile(r *FtransListFileRequest) (*FtransListFileResponse, error) {
	uri := FtransS10P2Path
	orderType := r.OrderType
	orderField := r.OrderField

	if orderType != FtransOrderTypeAsc && orderField != FtransOrderTypeDesc {
		orderType = FtransOrderTypeAsc
	}

	if orderField != FtransOrderFieldFileName && orderField != FtransOrderFieldMTime {
		orderField = FtransOrderFieldFileName
	}
	pageNum := r.PageNum
	if pageNum <= 0 {
		pageNum = FtransDefaultPageNum
	}
	pageSize := r.PageSize
	if pageSize <= 0 {
		pageSize = FtransDefaultPageSize
	}

	query := map[string]string{
		"op":       "list_dir",
		"src":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(r.Prefix))),
		"filterIn": base64.URLEncoding.EncodeToString([]byte(r.FilterIn)),
		"sortBy":   orderField,
		"orderBy":  orderType,
		"pageNo":   fmt.Sprintf("%d", pageNum),
		"pageSize": fmt.Sprintf("%d", pageSize),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response code %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("empty response found")
	}
	defer resp.Body.Close()

	var records []*FtransFileInfo
	var res FtransListFileResponse

	total, err := strconv.ParseInt(resp.Header.Get("matchedNum"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid total number[%s] found in response header", resp.Header.Get("matchedNum"))
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&records); err != nil {
		return nil, fmt.Errorf("invalid response body found, %s", err.Error())
	}

	res.Total = total
	res.Records = records

	return &res, nil
}

func (fc *FtransClient) getFullPath(p string) string {
	return fmt.Sprintf("%s/%s/%s", FtransDefaultMountPath, fc.BucketName, p)
}

func (fc *FtransClient) getFileSize(fileName string) (int64, int64, error) {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "size_file",
		"src":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(fileName))),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, nil)
	if err != nil {
		return 0, 0, err
	}

	if resp.StatusCode != HTTPStatusOK {
		return 0, 0, fmt.Errorf("invalid response status [%s] when stat %s", resp.StatusCode, fileName)
	}

	if resp.Body != nil {
		resp.Body.Close()
	}

	mtime, err := strconv.ParseInt(resp.Header.Get("Mtime"), 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid Mtime[%s] in response header", resp.Header.Get("Mtime"))
	}
	size, err := strconv.ParseInt(resp.Header.Get("Fsize"), 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid Fsize[%s] in response header", resp.Header.Get("Fsize"))
	}

	return size, mtime, nil
}

func (fc *FtransClient) getFileMd5(fileName string) (string, error) {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "md5_file",
		"src":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(fileName))),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, nil)
	if err != nil {
		return "", err
	}

	if resp.Body != nil {
		resp.Body.Close()
	}

	md5Str := resp.Header.Get("Md5")
	if md5Str == "" {
		return "", fmt.Errorf("empty md5 received for %s in response header", fileName)
	}

	return md5Str, nil
}

func (fc *FtransClient) RemoveFile(fileName string) error {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "remove_file",
		"des":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(fileName))),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, nil)
	if err != nil {
		return fmt.Errorf("remove file %s failed %s", fileName, err.Error())
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status code %d found when remove %s", resp.StatusCode, fileName)
	}

	return nil
}

func (fc *FtransClient) renameFile(src, dst string) error {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "rename_file",
		"src":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(src))),
		"des":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(dst))),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, nil)
	if err != nil {
		return fmt.Errorf("rename file %s into %s failed %s", src, dst, err.Error())
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status code %d found when rename %s into %s",
			resp.StatusCode, src, dst)
	}

	return nil
}

func (fc *FtransClient) doPost(uri string, query map[string]string, header http.Header, data []byte) (
	*http.Response, error) {
	var params []string
	for k, v := range query {
		params = append(params, fmt.Sprintf("%s=%s", k, v))
	}

	url := fmt.Sprintf("https://%s%s?%s", fc.FtransS10ServerAddr, uri, strings.Join(params, "&"))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("new http request for %s failed, %s", url, err.Error())
	}

	for k, v := range header {
		req.Header[k] = v
	}

	resp, err := fc.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do http request for %s failed, %s", url, err.Error())
	}

	return resp, nil
}

func (fc *FtransClient) splitPartTask(fft *ftransFileTask) {
	start := int64(0)
	end := fft.FileSize

	for ; start < end; start += fft.PartSize {
		size := fft.PartSize
		if start+size >= end {
			size = end - start
		}
		fpt := ftransPartTask{
			Offset:     start,
			Size:       size,
			Status:     FtransTaskStatusWaiting,
			Data:       make([]byte, size),
			parentTask: fft,
		}

		fft.PartTasks = append(fft.PartTasks, &fpt)
	}

	fft.DoingParts = 0
	fft.WaitingParts = int32(len(fft.PartTasks))
}

func (fc *FtransClient) uploadFilePart(fpt *ftransPartTask) error {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "upload_file",
		"des":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(fpt.parentTask.TempFilePath))),
		"offset":   fmt.Sprintf("%d", fpt.Offset),
		"size":     fmt.Sprintf("%d", fpt.Size),
		"mtime":    fmt.Sprintf("%d", fpt.parentTask.MtimeUSec),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, fpt.Data)
	if err != nil {
		return fmt.Errorf("upload part(%d, %d) for %s failed, %s", fpt.Offset, fpt.Size,
			fpt.parentTask.LocalFilePath, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code[%d] when uploading part(%d, %d) for %s",
			resp.StatusCode, fpt.Offset, fpt.Size, fpt.parentTask.LocalFilePath)
	}

	return nil
}

func (fc *FtransClient) downloadFilePart(fpt *ftransPartTask) error {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "download_file",
		"src":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(fpt.parentTask.RemoteFilePath))),
		"offset":   fmt.Sprintf("%d", fpt.Offset),
		"size":     fmt.Sprintf("%d", fpt.Size),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, fpt.Data)
	if err != nil {
		return fmt.Errorf("upload part(%d, %d) for %s failed, %s", fpt.Offset, fpt.Size,
			fpt.parentTask.LocalFilePath, err.Error())
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code[%d] when downloading part(%d, %d) for %s",
			resp.StatusCode, fpt.Offset, fpt.Size, fpt.parentTask.LocalFilePath)
	}

	fileSize, err := strconv.ParseInt(resp.Header.Get("Fsize"), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid Fsize[%s] found in response header when downloading part(%d, %d) for %s",
			resp.Header.Get("Fsize"), fpt.Offset, fpt.Size, fpt.parentTask.RemoteFilePath)
	}

	mtime, err := strconv.ParseInt(resp.Header.Get("Mtime"), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid Mtime[%s] found in response header when downloading part(%d, %d) for %s",
			resp.Header.Get("Mtime"), fpt.Offset, fpt.Size, fpt.parentTask.RemoteFilePath)
	}

	if fpt.parentTask.ConsistencyCheck {
		if fpt.parentTask.FileSize != fileSize || fpt.parentTask.MtimeUSec != mtime {
			return fmt.Errorf("file has been changed during downloading")
		}
	}

	fpt.Data, err = ioutil.ReadAll(resp.Body)

	return err
}

func (fc *FtransClient) makeDir(dir string, mtime int64) error {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "make_dir",
		"des":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(dir))),
		"mtime":    fmt.Sprintf("%d", mtime),
		"acltoken": fc.FtransAclToken,
	}
	resp, err := fc.doPost(uri, query, nil, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status [%d] when make dir %s", resp.StatusCode, dir)
	}

	return nil
}

func (fc *FtransClient) fileExist(fileName string) error {
	uri := FtransS10P2Path
	query := map[string]string{
		"op":       "exist_file",
		"src":      base64.URLEncoding.EncodeToString([]byte(fc.getFullPath(fileName))),
		"acltoken": fc.FtransAclToken,
	}

	resp, err := fc.doPost(uri, query, nil, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status [%d] when exist_file %s", resp.StatusCode, fileName)
	}

	return nil
}

func (fc *FtransClient) getLocalFileMd5(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}

	hash := md5.New()
	_, _ = io.Copy(hash, f)
	return strings.ToLower(hex.EncodeToString(hash.Sum(nil)))
}

func (fc *FtransClient) uploadPart(freeCh chan bool, taskCh chan *ftransPartTask, callbackCh chan *ftransPartTask) {
	for {
		freeCh <- true
		select {
		case fpt := <-taskCh:
			if err := fc.uploadFilePart(fpt); err != nil {
				fpt.Status = FtransTaskStatusFailed
			} else {
				fpt.Status = FtransTaskStatusFinished
			}

			callbackCh <- fpt
		}
	}
}

func (fc *FtransClient) downloadPart(freeCh chan bool, taskCh chan *ftransPartTask, callbackCh chan *ftransPartTask) {
	for {
		freeCh <- true
		select {
		case fpt := <-taskCh:
			if err := fc.downloadFilePart(fpt); err != nil {
				fpt.Status = FtransTaskStatusFailed
			} else {
				fpt.Status = FtransTaskStatusFinished
			}

			callbackCh <- fpt
		}
	}
}
