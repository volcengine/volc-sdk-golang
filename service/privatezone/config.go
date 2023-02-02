package privatezone

import (
	"net/http"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

var (
	DefaultRegion  string
	ServiceVersion string
	ServiceName    string
	ServiceHost    string
	Timeout        int
	ServiceInfo    *base.ServiceInfo
)

func init() {
	ServiceName = "private_zone"
	ServiceHost = "open.volcengineapi.com"
	ServiceVersion = "2022-06-01"
	DefaultRegion = "cn-beijing"
	Timeout = 15

	ServiceInfo = &base.ServiceInfo{
		Timeout: time.Duration(Timeout) * time.Second,
		Host:    ServiceHost,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Scheme: "http",
		Credentials: base.Credentials{
			Service: ServiceName,
			Region:  DefaultRegion,
		},
	}
}
