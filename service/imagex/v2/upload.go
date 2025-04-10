package imagex

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/avast/retry-go"
)

type uploadTaskSet struct {
	ctx       context.Context
	host      string
	info      []StoreInfo
	content   []io.Reader
	size      []int64
	cts       []string
	serviceId string

	lock        sync.Mutex
	taskChan    chan *uploadTaskElement
	successOids []string
	result      []uploadTaskResult
}

type uploadTaskResult struct {
	uri     string
	success bool
	errMsg  string
	putErr  *PutError
}

func (r *uploadTaskSet) GetResult() []Result {
	ret := make([]Result, 0)
	for _, item := range r.result {
		rst := Result{
			Uri:        item.uri,
			UriStatus:  2000,
			Encryption: Encryption{},
		}
		if item.success {
			ret = append(ret, rst)
			continue
		}
		rst.UriStatus = 2001
		rst.PutError = item.putErr
		if rst.PutError == nil {
			rst.PutError = &PutError{
				ErrorCode: -2001,
				Error:     item.errMsg,
				Message:   item.errMsg,
			}
		}
		ret = append(ret, rst)
	}
	return ret
}

func (r *uploadTaskSet) fill(result *CommitUploadImageResult) {
	m := make(map[string]int)
	for idx, item := range r.result {
		if !item.success {
			m[item.uri] = idx
		}
	}
	for idx := range result.Results {
		idx, ok := m[result.Results[idx].Uri]
		if ok {
			result.Results[idx].PutError = r.result[idx].putErr
			if result.Results[idx].PutError == nil {
				result.Results[idx].PutError = &PutError{
					ErrorCode: -2001,
					Error:     r.result[idx].errMsg,
					Message:   r.result[idx].errMsg,
				}
			}
		}
	}
}

type uploadTaskElement struct {
	ctx     context.Context
	host    string
	idx     int
	info    StoreInfo
	content io.Reader
	size    int64
	ct      string
}

func (r *uploadTaskSet) init() {
	r.lock = sync.Mutex{}
	r.successOids = make([]string, 0)
	r.result = make([]uploadTaskResult, len(r.info))
	r.taskChan = make(chan *uploadTaskElement, len(r.size))
	for idx := range r.size {
		ele := &uploadTaskElement{
			ctx:     r.ctx,
			host:    r.host,
			idx:     idx,
			info:    r.info[idx],
			content: r.content[idx],
			size:    r.size[idx],
		}
		if idx < len(r.cts) {
			ele.ct = r.cts[idx]
		}
		r.taskChan <- ele
	}
	close(r.taskChan)
}

func (r *uploadTaskSet) addSuccess(oid string) {
	r.lock.Lock()
	r.successOids = append(r.successOids, oid)
	r.lock.Unlock()
}

type uploadPartCommonResponse struct {
	Version string      `json:"Version"`
	Success int         `json:"success,omitempty"`
	Error   uploadError `json:"error"`
}

type uploadPartResponse struct {
	uploadPartCommonResponse
	PayLoad initPartPayLoad `json:"payload,omitempty"`
}

type uploadMergeResponse struct {
	uploadPartCommonResponse
	PayLoad mergePartPayLoad `json:"payload,omitempty"`
}

type mergePartPayLoad struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type uploadError struct {
	HttpCode  int    `json:"code"`
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type initPartPayLoad struct {
	UploadID string `json:"uploadID"`
}

type directUploadResponse struct {
	Success int           `json:"success"`
	Error   *uploadError  `json:"error"`
	Payload UploadPayload `json:"payload,omitempty"`
}

type UploadPayload struct {
	Hash string `json:"hash"`
}

func (c *Imagex) directUpload(ctx context.Context, host string, idx int, set *uploadTaskSet, storeInfo StoreInfo, imageBytes []byte, ct string) error {
	if len(imageBytes) == 0 {
		return fmt.Errorf("file size is zero")
	}

	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(imageBytes))
	url := fmt.Sprintf("https://%s/%s", host, getEscapePath(storeInfo.StoreUri))
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(imageBytes))
	if err != nil {
		return fmt.Errorf("fail to new put request %s, %v", url, err)
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", storeInfo.Auth)
	if ct != "" {
		req.Header.Set("Specified-Content-Type", ct)
	}
	req = req.WithContext(ctx)

	now := time.Now()
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		reporter.report(c, c.buildDefaultUploadReport(set.serviceId, time.Since(now).Microseconds(), 500, 0, "", actionDirectUpload, host, err.Error()))
		return fmt.Errorf("fail to do request to %s, %v", url, err)
	}
	defer rsp.Body.Close()

	//goland:noinspection GoDeprecation
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		reporter.report(c, c.buildDefaultUploadReport(set.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, rsp.Header.Get(logHeader), actionDirectUpload, host, err.Error()))
		return fmt.Errorf("fail to read response body %s, %v", url, err)
	}

	var putResp = directUploadResponse{
		Success: -1,
	}
	if err = json.Unmarshal(body, &putResp); err != nil {
		err = fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(body), err)
		reporter.report(c, c.buildDefaultUploadReport(set.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, rsp.Header.Get(logHeader), actionDirectUpload, host, err.Error()))
		return err
	}
	if putResp.Success != 0 {
		set.result[idx].errMsg = fmt.Sprintf("upload fail, %s", putResp.Error.Message)
		set.result[idx].putErr = &PutError{
			ErrorCode: putResp.Error.ErrorCode,
			Error:     putResp.Error.Error,
			Message:   putResp.Error.Message,
		}
		return fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(body), putResp)
	}
	if checkSum != putResp.Payload.Hash {
		return fmt.Errorf("url %s crc32 not match, got: %s, want: %s", url, putResp.Payload.Hash, checkSum)
	}
	return nil
}

type segmentedUploadParam struct {
	host string
	StoreInfo
	content     io.Reader
	size        int64
	isLargeFile bool
	idx         int
	set         *uploadTaskSet
	ct          string
	imagex      *Imagex
}

func (c *segmentedUploadParam) chunkUpload() error {
	uploadID, err := c.initUploadPart()
	if err != nil {
		return err
	}
	cur := make([]byte, MinChunkSize)
	parts := make([]string, 0)

	num := c.size / MinChunkSize
	cnt := 0
	lastNum := int(num) - 1

	// 读 n-1 片并上传上去
	var (
		part       string
		logid      string
		statusCode int
	)
	for i := 0; i < lastNum; i++ {
		n, err := io.ReadFull(c.content, cur)
		if err != nil {
			return err
		}
		cnt += n
		partNumber := i
		if c.isLargeFile {
			partNumber++
		}
		now := time.Now()
		err = retry.Do(func() error {
			part, logid, statusCode, err = c.uploadPart(uploadID, partNumber, cur)
			return err
		}, retry.Attempts(3))
		if err != nil {
			reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), statusCode, 0, logid, actionChunkUpload, c.host, err.Error()))
			return err
		}
		parts = append(parts, part)
	}

	// 读 n 和 n+1片（如有）上传上去
	//goland:noinspection GoDeprecation
	bts, err := ioutil.ReadAll(c.content)
	if err != nil {
		return err
	}
	total := len(bts) + cnt
	if total != int(c.size) {
		return fmt.Errorf(fmt.Sprintf("last part download size mismatch ,download %d , size %d", total, c.size))
	}
	if c.isLargeFile {
		lastNum++
	}
	now := time.Now()
	err = retry.Do(func() error {
		part, logid, statusCode, err = c.uploadPart(uploadID, lastNum, bts)
		return err
	}, retry.Attempts(3))
	if err != nil {
		reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), statusCode, 0, logid, actionChunkUpload, c.host, err.Error()))
		return err
	}
	parts = append(parts, part)
	return c.uploadMergePart(uploadID, parts)
}

func (c *segmentedUploadParam) initUploadPart() (string, error) {
	url := fmt.Sprintf("https://%s/%s?uploads", c.host, getEscapePath(c.StoreUri))
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", c.Auth)
	if c.isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	if c.ct != "" {
		req.Header.Set("Specified-Content-Type", c.ct)
	}

	now := time.Now()
	client := http.DefaultClient
	rsp, err := client.Do(req)
	if err != nil {
		reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), 500, 0, "", actionInitChunk, c.host, err.Error()))
		return "", err
	}
	//goland:noinspection GoDeprecation
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, rsp.Header.Get(logHeader), actionInitChunk, c.host, err.Error()))
		return "", err
	}

	res := &uploadPartResponse{}
	res.Success = -1
	if err := json.Unmarshal(b, res); err != nil {
		err = fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(b), err)
		reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, rsp.Header.Get(logHeader), actionInitChunk, c.host, err.Error()))
		return "", err
	}
	if res.Success != 0 {
		c.set.result[c.idx].putErr = &PutError{
			ErrorCode: res.Error.ErrorCode,
			Error:     res.Error.Error,
			Message:   res.Error.Message,
		}
		return "", fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(b), res)
	}
	return res.PayLoad.UploadID, nil
}

func (c *segmentedUploadParam) uploadPart(uploadID string, partNumber int, data []byte) (string, string, int, error) {
	url := fmt.Sprintf("https://%s/%s?partNumber=%d&uploadID=%s", c.host, getEscapePath(c.StoreUri), partNumber, uploadID)
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(data))
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return "", "", -1, err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", c.Auth)
	if c.isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	if c.ct != "" {
		req.Header.Set("Specified-Content-Type", c.ct)
	}

	client := http.DefaultClient
	rsp, err := client.Do(req)
	if err != nil {
		return "", "", 500, err
	}
	//goland:noinspection GoDeprecation
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", rsp.Header.Get(logHeader), rsp.StatusCode, err
	}

	res := &uploadPartResponse{}
	res.Success = -1

	if err := json.Unmarshal(b, res); err != nil {
		return "", rsp.Header.Get(logHeader), rsp.StatusCode, fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(b), err)
	}
	if res.Success != 0 {
		c.set.result[c.idx].putErr = &PutError{
			ErrorCode: res.Error.ErrorCode,
			Error:     res.Error.Error,
			Message:   res.Error.Message,
		}
		return "", rsp.Header.Get(logHeader), rsp.StatusCode, fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(b), res)
	}
	return checkSum, rsp.Header.Get(logHeader), rsp.StatusCode, nil
}

func (c *segmentedUploadParam) uploadMergePart(uploadID string, checkSum []string) error {
	url := fmt.Sprintf("https://%s/%s?uploadID=%s", c.host, getEscapePath(c.StoreUri), uploadID)
	body, err := c.genMergeBody(checkSum)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", c.Auth)
	if c.isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	if c.ct != "" {
		req.Header.Set("Specified-Content-Type", c.ct)
	}

	now := time.Now()
	client := http.DefaultClient
	rsp, err := client.Do(req)
	if err != nil {
		reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), 500, 0, "", actionMergeChunk, c.host, err.Error()))
		return err
	}
	//goland:noinspection GoDeprecation
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, rsp.Header.Get(logHeader), actionMergeChunk, c.host, err.Error()))
		return err
	}

	res := &uploadMergeResponse{}
	res.Success = -1
	if err := json.Unmarshal(b, res); err != nil {
		err = fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(b), err)
		reporter.report(c.imagex, c.imagex.buildDefaultUploadReport(c.set.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, rsp.Header.Get(logHeader), actionMergeChunk, c.host, err.Error()))
		return err
	}
	if res.Success != 0 {
		c.set.result[c.idx].putErr = &PutError{
			ErrorCode: res.Error.ErrorCode,
			Error:     res.Error.Error,
			Message:   res.Error.Message,
		}
		return fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(b), res)
	}
	return nil
}

func (c *segmentedUploadParam) genMergeBody(checkSum []string) (string, error) {
	if len(checkSum) == 0 {
		return "", fmt.Errorf("body crc32 empty")
	}
	s := make([]string, len(checkSum))
	for partNumber, crc := range checkSum {
		s[partNumber] = fmt.Sprintf("%d:%s", partNumber, crc)
	}
	return strings.Join(s, ","), nil
}

func getEscapePath(path string) string {
	elems := strings.Split(path, "/")
	for i := range elems {
		elems[i] = url.PathEscape(elems[i])
	}
	return strings.Join(elems, "/")
}

type vpcUploadDataParam struct {
	f         *os.File
	data      []byte
	size      int
	serviceId string
}

type vpcUploadPartsInfo struct {
	Parts []*vpcUploadPartInfo `json:"Parts"`
}

type vpcUploadPartInfo struct {
	PartNumber int    `json:"PartNumber"`
	ETag       string `json:"ETag"`
}

func (c *Imagex) vpcUpload(vpcUploadInfo *ApplyVpcUploadInfoResResult, dataParam *vpcUploadDataParam) error {
	if vpcUploadInfo == nil || dataParam == nil {
		return errors.New("vpc upload info is nil")
	}

	if vpcUploadInfo.UploadMode == "direct" {
		return c.vpcPut(vpcUploadInfo, dataParam)
	} else if vpcUploadInfo.UploadMode == "part" {
		return c.vpcPartUpload(vpcUploadInfo.PartUploadInfo, dataParam)
	}

	return errors.New("unknown upload mode")
}

func (c *Imagex) vpcPartUpload(partUploadInfo *ApplyVpcUploadInfoResResultPartUploadInfo, dataParam *vpcUploadDataParam) error {
	if partUploadInfo == nil {
		return errors.New("part upload info is nil")
	}

	totalNum := dataParam.size / partUploadInfo.PartSize
	lastPartSize := dataParam.size % partUploadInfo.PartSize
	if (len(partUploadInfo.PartPutURLs) != totalNum+1 && lastPartSize > 0) ||
		(lastPartSize == 0 && len(partUploadInfo.PartPutURLs) != totalNum) {
		return errors.New("mismatch part upload urls")
	}

	partsInfo := &vpcUploadPartsInfo{
		Parts: make([]*vpcUploadPartInfo, 0),
	}
	for i := 0; i < len(partUploadInfo.PartPutURLs); i++ {
		etag, err := c.vpcPartPut(partUploadInfo, dataParam, i)
		if err != nil {
			return err
		}
		partsInfo.Parts = append(partsInfo.Parts, &vpcUploadPartInfo{
			PartNumber: i + 1,
			ETag:       etag,
		})
	}

	return c.vpcPost(partUploadInfo, partsInfo, dataParam)
}

func (c *Imagex) vpcPost(partUploadInfo *ApplyVpcUploadInfoResResultPartUploadInfo, partsInfo *vpcUploadPartsInfo, param *vpcUploadDataParam) error {
	postUrl := partUploadInfo.CompletePartURL
	bodyBytes, err := json.Marshal(partsInfo)
	if err != nil {
		return err
	}
	u, err := url.Parse(postUrl)
	if err != nil {
		return errors.New("postUrl is invalid")
	}

	client := http.DefaultClient
	req, err := http.NewRequest("POST", postUrl, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	for _, putHeader := range partUploadInfo.CompletePartURLHeaders {
		req.Header.Set(putHeader.Key, putHeader.Value)
	}

	now := time.Now()
	rsp, err := client.Do(req)
	if err != nil {
		reporter.report(c, c.buildDefaultUploadReport(param.serviceId, time.Since(now).Microseconds(), 500, 0, "", actionVpcMergeChunk, u.Host, err.Error()))
		return err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		logId := rsp.Header.Get("x-tos-request-id")
		var errMsg string
		bts, err := ioutil.ReadAll(rsp.Body)
		if err == nil {
			errMsg = string(bts)
		}
		reporter.report(c, c.buildDefaultUploadReport(param.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, logId, actionVpcMergeChunk, u.Host, errMsg))
		return fmt.Errorf("post error: code %v, logId %v", rsp.StatusCode, logId)
	}

	return nil
}

func (c *Imagex) vpcPartPut(partUploadInfo *ApplyVpcUploadInfoResResultPartUploadInfo, dataParam *vpcUploadDataParam, index int) (string, error) {
	var dataReader io.Reader

	partPutUrl := partUploadInfo.PartPutURLs[index]
	chunkSize := partUploadInfo.PartSize
	offset := index * partUploadInfo.PartSize
	if index == len(partUploadInfo.PartPutURLs)-1 {
		chunkSize = dataParam.size - offset
	}

	if dataParam.f != nil {
		sectionReader := io.NewSectionReader(dataParam.f, int64(offset), int64(chunkSize))
		_, err := sectionReader.Seek(0, 0)
		if err != nil {
			return "", errors.New("reader seek")
		}
		dataReader = sectionReader
	} else {
		dataReader = bytes.NewReader(dataParam.data[offset : offset+chunkSize])
	}

	u, err := url.Parse(partPutUrl)
	if err != nil {
		return "", errors.New("partPutUrl is invalid")
	}

	client := http.DefaultClient
	req, err := http.NewRequest("PUT", partPutUrl, dataReader)
	if err != nil {
		return "", err
	}

	now := time.Now()
	rsp, err := client.Do(req)
	if err != nil {
		reporter.report(c, c.buildDefaultUploadReport(dataParam.serviceId, time.Since(now).Microseconds(), 500, 0, "", actionVpcChunkUpload, u.Host, err.Error()))
		return "", err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		logId := rsp.Header.Get("x-tos-request-id")
		var errMsg string
		bts, err := ioutil.ReadAll(rsp.Body)
		if err == nil {
			errMsg = string(bts)
		}
		reporter.report(c, c.buildDefaultUploadReport(dataParam.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, logId, actionVpcChunkUpload, u.Host, errMsg))
		return "", fmt.Errorf("put error: code %v, logId %v ", rsp.StatusCode, logId)
	}
	etag := rsp.Header.Get("ETag")

	return etag, nil
}

func (c *Imagex) vpcPut(vpcUploadInfo *ApplyVpcUploadInfoResResult, dataParam *vpcUploadDataParam) error {
	var dataReader io.Reader
	putUrl := vpcUploadInfo.PutURL
	u, err := url.Parse(putUrl)
	if err != nil {
		return errors.New("putUrl parse failed")
	}
	if dataParam.f != nil {
		_, err := dataParam.f.Seek(0, 0)
		if err != nil {
			return errors.New("file seek")
		}
		dataReader = dataParam.f
	} else {
		dataReader = bytes.NewReader(dataParam.data)
	}
	client := http.DefaultClient
	req, err := http.NewRequest("PUT", putUrl, dataReader)
	if err != nil {
		return err
	}

	for _, putHeader := range vpcUploadInfo.PutURLHeaders {
		req.Header.Set(putHeader.Key, putHeader.Value)
	}

	now := time.Now()
	rsp, err := client.Do(req)
	if err != nil {
		reporter.report(c, c.buildDefaultUploadReport(dataParam.serviceId, time.Since(now).Microseconds(), 500, 0, "", actionVpcDirectUpload, u.Host, err.Error()))
		return err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		logId := rsp.Header.Get("x-tos-request-id")
		var errMsg string
		bts, err := ioutil.ReadAll(rsp.Body)
		if err == nil {
			errMsg = string(bts)
		}
		reporter.report(c, c.buildDefaultUploadReport(dataParam.serviceId, time.Since(now).Microseconds(), rsp.StatusCode, 0, logId, actionVpcDirectUpload, u.Host, errMsg))
		return fmt.Errorf("put error: code %v, logId %v ", rsp.StatusCode, logId)
	}

	return nil
}
