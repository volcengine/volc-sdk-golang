package diff

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

const (
	ServiceName    = "vemars_gateway"
	ApiVersion     = "2020-12-25"
	Action         = "MarsApi"
	DefaultTimeout = 10 * time.Second
	DefaultHost    = "open.volcengineapi.com"
	DefaultRegion  = "cn-north-1"

	createService = "CreateService"
	validate      = "Validate"
	genByCount    = "GenByCount"
	genByVersion  = "GenByVersion"
	genByPkg      = "GenByPkg"
	checkResponse = "CheckResponse"
)

var (
	serviceInfo = &base.ServiceInfo{
		Timeout: DefaultTimeout,
		Scheme:  "https",
		Host:    DefaultHost,
		Credentials: base.Credentials{
			Service: ServiceName,
			Region:  DefaultRegion,
		},
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	apiInfoList = map[string]*base.ApiInfo{
		// 新建服务
		createService: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{Action},
				"Version": []string{ApiVersion},
			},
			Header: http.Header{
				"x-mars-path": []string{"/diff/create_service"},
			},
		},
		// 验证原始包和差分包是否能生成新包
		validate: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{Action},
				"Version": []string{ApiVersion},
			},
			Header: http.Header{
				"x-mars-path": []string{"/diff/validate"},
			},
		},
		genByCount: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{Action},
				"Version": []string{ApiVersion},
			},
			Header: http.Header{
				"x-mars-path": []string{"/diff/gen_by_count"},
			},
		},
		genByVersion: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{Action},
				"Version": []string{ApiVersion},
			},
			Header: http.Header{
				"x-mars-path": []string{"/diff/gen_by_version"},
			},
		},
		genByPkg: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{Action},
				"Version": []string{ApiVersion},
			},
			Header: http.Header{
				"x-mars-path": []string{"/diff/gen_by_pkg"},
			},
		},
		checkResponse: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{Action},
				"Version": []string{ApiVersion},
			},
			Header: http.Header{
				"x-mars-path": []string{"/diff/check_response"},
			},
		},
	}
)

type Diff struct {
	*base.Client
}

var DefaultInstance = newInstance()

func newInstance() *Diff {
	return &Diff{
		Client: base.NewClient(serviceInfo, apiInfoList),
	}
}
