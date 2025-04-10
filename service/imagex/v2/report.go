package imagex

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/base"
)

var (
	reporter     *reporterClient
	reporterOnce sync.Once
)

type reporterClient struct {
	q        chan *pair
	maxBatch int
	signals  chan os.Signal
	closed   bool
}

type pair struct {
	ins    *Imagex
	report *report
}

// *** 打点 ***
type report struct {
	Id          string    `json:"Id,required"` // Id为uuid或LogId二者选一，必填
	IsLogId     bool      `json:"IsLogId"`
	AccountName string    `json:"AccountName"`
	Module      string    `json:"Module"`
	Status      string    `json:"Status"`
	Metrics     []*metric `json:"Metrics"`
	Tags        []*tag    `json:"Tags"`
}

type metric struct {
	Type  int    `json:"Type"`
	Value string `json:"Value"`
}

type tag struct {
	EnableLog bool   `json:"EnableLog"`
	Name      string `json:"Name"`
	Value     string `json:"Value"`
}

const (
	logHeader = "X-TT-LOGID"

	tagStatusCode = "status_code"
	tagDomain     = "domain"
	tagAction     = "action"
	tagFromHost   = "from_host"
	tagErrorMsg   = "err_msg"
	tagRetryTimes = "retry_times"

	actionApplyUploadInfo  = "ApplyUploadInfo"
	actionDirectUpload     = "DirectUpload"
	actionInitChunk        = "InitChunk"
	actionChunkUpload      = "ChunkUpload"
	actionMergeChunk       = "MergeChunk"
	actionCommitUploadInfo = "CommitUploadInfo"
	actionFinishUpload     = "FinishUpload"

	actionApplyVpcUploadInfo  = "ApplyVpcUploadInfo"
	actionVpcDirectUpload     = "VpcDirectUpload"
	actionVpcChunkUpload      = "VpcChunkUpload"
	actionVpcMergeChunk       = "VpcMergeChunk"
	actionCommitVpcUploadInfo = "CommitVpcUploadInfo"

	hardThreshold = 100 * 1024
	softThreshold = hardThreshold / 2
)

func initReporterClient() {
	reporterOnce.Do(func() {
		reporter = &reporterClient{
			q:        make(chan *pair, 256),
			maxBatch: 8,
			signals:  make(chan os.Signal, 1),
		}
		signal.Notify(reporter.signals, syscall.SIGINT, syscall.SIGTERM)
		go reporter._loop()
	})
}

func (r *reporterClient) _loop() {
	if r == nil {
		return
	}
	timer := time.NewTimer(0)
	defer func() {
		timer.Stop()
		r.closed = true
		close(r.q)
		if err := recover(); err != nil {
			return
		}
	}()
	for {
		select {
		case <-r.signals:
			return
		case <-timer.C:
			batch := r.maxBatch
			var pairs = make([]*pair, 0, batch)
			for i := 0; i < batch; i++ {
				select {
				case e := <-r.q:
					if e != nil && e.ins != nil && e.report != nil {
						pairs = append(pairs, e)
					}
				default:
					break
				}
			}
			if len(pairs) > 0 {
				vodMap := make(map[*Imagex][]*report)
				size := 0
				for _, e := range pairs {
					vodMap[e.ins] = append(vodMap[e.ins], e.report)
				}
				for vod, rs := range vodMap {
					if bts, err := json.Marshal(rs); err == nil {
						if compressed, err2 := r.compressData(bts); err2 == nil {
							id := uuid.NewString()
							//fmt.Printf("consume: %s,rports: %d, batch: %d\n", id, len(rs), batch)
							r._report(vod, id, compressed)
							size += len(compressed)
						}

					}
				}

				if int(float64(batch)*0.7) <= len(pairs) {
					if size > hardThreshold {
						batch >>= 1
						if batch == 0 {
							batch = 1
						}
						r.maxBatch = batch
					} else if size < softThreshold {
						if (size << 1) < softThreshold {
							batch <<= 1
						} else {
							batch++
						}
						r.maxBatch = batch
					}
				}
			}
			//fmt.Println("队列:", len(r.q))
			timer.Stop()
			timer = time.NewTimer(time.Second * 1) // 1s send
		}
	}
}

func (r *reporterClient) report(ins *Imagex, req *report) {
	if r == nil || r.closed || ins == nil || req == nil || ins.disableLog {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("reporterClient report panic:", r)
		}
	}()
	select {
	case r.q <- &pair{ins, req}: // maybe panic
	default:
		//fmt.Println("queue full")
	}
}

func (r *reporterClient) _report(ins *Imagex, id string, reports []byte) {
	if r == nil || ins == nil {
		return
	}
	ins.reportEvent(id, reports)
}

func (r *reporterClient) compressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	_, err := gzipWriter.Write(data)
	if err != nil {
		return nil, err
	}

	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type reportEventResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata,omitempty"`
}

func (p *Imagex) reportEvent(id string, reports []byte) (*reportEventResponse, int, error) {
	fieldItem := base.CreateMultiPartItemFormField("Id", id)
	fileItem := base.CreateMultiPartItemFormFile("Reports", "reports_data", bytes.NewReader(reports))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	respBody, status, err := p.CtxMultiPart(ctx, "ReportEvent", url.Values{}, []*base.MultiPartItem{fieldItem, fileItem})

	output := &reportEventResponse{}
	errUnmarshal := json.Unmarshal(respBody, output)

	if err != nil || status != http.StatusOK {
		if errUnmarshal != nil || output.ResponseMetadata == nil || output.ResponseMetadata.Error == nil || output.ResponseMetadata.Error.Code == "" {
			if err == nil {
				err = errors.New(string(respBody))
			}
			return nil, status, err
		} else {
			return output, status, errors.New(output.ResponseMetadata.Error.Code)
		}
	}
	return output, status, nil
}

func (p *Imagex) buildDefaultUploadReport(serviceId string, cost int64, statusCode, retryTimes int, logId, action, domain, errMsg string) *report {
	if serviceId == "" || statusCode == 200 && errMsg != "" {
		return nil
	}
	r := &report{
		Id:          logId,
		IsLogId:     true,
		AccountName: serviceId,
		Module:      "upload",
		Status:      "success",
		Metrics:     []*metric{{Type: 0}},
		Tags: []*tag{
			{true, tagStatusCode, strconv.Itoa(statusCode)},
		},
	}
	if cost > 0 {
		r.Metrics = append(r.Metrics, &metric{1, strconv.Itoa(int(cost))})
	}
	if retryTimes > 0 {
		r.Tags = append(r.Tags, &tag{true, tagRetryTimes, strconv.Itoa(retryTimes)})
	}
	if logId == "" {
		r.Id = uuid.NewString()
		r.IsLogId = false
	}
	if action != "" {
		r.Tags = append(r.Tags, &tag{true, tagAction, action})
	}
	if domain != "" {
		r.Tags = append(r.Tags, &tag{true, tagDomain, domain})
	}
	if errMsg != "" {
		r.Status = "failed"
		r.Tags = append(r.Tags, &tag{true, tagErrorMsg, errMsg})
	}
	if p != nil {
		if p.ServiceInfo.Host != "" {
			r.Tags = append(r.Tags, &tag{true, tagFromHost, p.ServiceInfo.Host})
		}
	}
	return r
}
