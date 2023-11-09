package billing

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20220101 = "2022-01-01"
	ServiceName            = "billing"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "billing.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{

		// accessKey
		"ListBillDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListBillDetail"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"ListBill": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListBill"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"ListBillOverviewByProd": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListBillOverviewByProd"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"ListBillOverviewByCategory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListBillOverviewByCategory"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"ListSplitBillDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSplitBillDetail"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"UnsubscribeInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnsubscribeInstance"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"ListAmortizedCostBillDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAmortizedCostBillDetail"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"ListAmortizedCostBillMonthly": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAmortizedCostBillMonthly"},
				"Version": []string{ServiceVersion20220101},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// Billing .
type Billing struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *Billing {
	instance := &Billing{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *Billing) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

// GetAPIInfo interface
func (p *Billing) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *Billing) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Billing) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *Billing) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
