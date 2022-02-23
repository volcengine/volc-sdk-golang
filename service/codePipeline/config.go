package codePipeline

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion = base.RegionCnNorth1
	ServiceName   = "cp"
)

type CodePipeline struct {
	*base.Client
	retry bool
}

var DefaultInstance = NewInstance()

func NewInstance() *CodePipeline {
	instance := &CodePipeline{
		Client: base.NewClient(ServiceInfo, ApiInfoList),
		retry:  true,
	}
	return instance
}

func (p *CodePipeline) Retry() bool {
	return p.retry
}

func (p *CodePipeline) CloseRetry() {
	p.retry = false
}

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{
			Region:  base.RegionCnNorth1,
			Service: ServiceName,
		},
	}
	ApiInfoList = map[string]*base.ApiInfo{
		"ListPipelines": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPipelines"},
				"Version": []string{"2021-03-03"},
			},
		},
		"CreatePipeline": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePipeline"},
				"Version": []string{"2021-03-03"},
			},
		},
		"DeletePipeline": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePipeline"},
				"Version": []string{"2021-03-03"},
			},
		},
		"UpdatePipeline": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePipeline"},
				"Version": []string{"2021-03-03"},
			},
		},
		"GetPipeline": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPipeline"},
				"Version": []string{"2021-03-03"},
			},
		},
		"UpdatePipelineProperties": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePipelineProperties"},
				"Version": []string{"2021-03-03"},
			},
		},
		"RunPipeline": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RunPipeline"},
				"Version": []string{"2021-03-03"},
			},
		},
		"ListPipelineRecords": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPipelineRecords"},
				"Version": []string{"2021-03-03"},
			},
		},
		"GetPipelineRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPipelineRecord"},
				"Version": []string{"2021-03-03"},
			},
		},
		"StopPipelineRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPipelineRecord"},
				"Version": []string{"2021-03-03"},
			},
		},
		"RetryPipelineRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RetryPipelineRecord"},
				"Version": []string{"2021-03-03"},
			},
		},
		"DeletePipelineRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePipelineRecord"},
				"Version": []string{"2021-03-03"},
			},
		},
	}
)
