package mars

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "veMARS"
	ApiVersion     = "2022-04-19"
	DefaultTimeout = 10 * time.Second
	//DefaultHost    = "vemars.volcengineapi.com"
	DefaultHost = "open.volcengineapi.com"

	DefaultRegion = "cn-north-1"

	createService       = "CreateService"
	validate            = "Validate"
	genByCount          = "GenByCount"
	genByVersion        = "GenByVersion"
	genByPkg            = "GenByPkg"
	checkResponse       = "CheckResponse"
	deletePackages      = "DeletePackages"
	deleteService       = "DeleteService"
	queryPatchByService = "QueryPatchByService"

	createServiceAction       = "diffServiceCreate"
	validateAction            = "diffValidate"
	genByCountAction          = "diffGenByCount"
	genByVersionAction        = "diffGenByVersion"
	genByPkgAction            = "diffGenByPkg"
	checkResponseAction       = "diffCheckResponse"
	deletePackagesAction      = "diffDeletePackages"
	deleteServiceAction       = "diffDeleteService"
	queryPatchByServiceAction = "diffQueryPatchByService"
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
			"Accept":   []string{"application/json"},
			"x-tt-env": []string{"ppe_mars_v2"},
		},
	}

	apiInfoList = map[string]*base.ApiInfo{
		// 新建服务
		createService: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{createServiceAction},
				"Version": []string{ApiVersion},
			},
		},
		// 验证原始包和差分包是否能生成新包
		validate: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{validateAction},
				"Version": []string{ApiVersion},
			},
		},
		// 根据传入的个数生成差分包
		genByCount: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{genByCountAction},
				"Version": []string{ApiVersion},
			},
		},
		// 根据传入的版本生成差分包
		genByVersion: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{genByVersionAction},
				"Version": []string{ApiVersion},
			},
		},
		// 根据传入的新包和老包生成差分包
		genByPkg: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{genByPkgAction},
				"Version": []string{ApiVersion},
			},
		},
		// 查看任务的完成情况
		checkResponse: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{checkResponseAction},
				"Version": []string{ApiVersion},
			},
		},
		// 删除原始包
		deletePackages: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{deletePackagesAction},
				"Version": []string{ApiVersion},
			},
		},
		// 删除服务
		deleteService: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{deleteServiceAction},
				"Version": []string{ApiVersion},
			},
		},
		// 查看服务中的包信息
		queryPatchByService: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{queryPatchByServiceAction},
				"Version": []string{ApiVersion},
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
