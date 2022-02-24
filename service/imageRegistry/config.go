package imageRegistry

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion = base.RegionCnNorth1
	ServiceName   = "cr"
)

type ImageRegistry struct {
	*base.Client
	retry bool
}

var DefaultInstance = NewInstance()

func NewInstance() *ImageRegistry {
	instance := &ImageRegistry{
		Client: base.NewClient(ServiceInfo, ApiInfoList),
		retry:  true,
	}
	return instance
}

func (p *ImageRegistry) Retry() bool {
	return p.retry
}

func (p *ImageRegistry) CloseRetry() {
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
		"DeleteNamespaceBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteNamespaceBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"CreateNamespaceBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateNamespaceBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"ValidateNamespaceBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ValidateNamespaceBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"GetNamespaceBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetNamespaceBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"ListNamespacesBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListNamespacesBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"ListRepositoriesBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRepositoriesBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"GetRepositoryBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRepositoryBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"CreateRepositoryBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRepositoryBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"UpdateRepositoryBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRepositoryBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"DeleteRepositoryBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRepositoryBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"ValidateRepositoryBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ValidateRepositoryBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"ListTagsBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListTagsBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"GetTagBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTagBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"GetTagAdditionBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTagAdditionBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
		"DeleteTagBasic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTagBasic"},
				"Version": []string{"2021-03-03"},
			},
		},
	}
)
