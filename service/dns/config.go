package dns

import (
	"net/http"
	"os"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion  = "cn-beijing"
	ServiceVersion = "2018-08-01"
	ServiceName    = "dns"
	Timeout        = 15

	xTopAccountID = ""
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: Timeout * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Scheme: "http",
		Credentials: base.Credentials{
			Service:         ServiceName,
			Region:          DefaultRegion,
			AccessKeyID:     os.Getenv("VOLC_ACCESSKEY"),
			SecretAccessKey: os.Getenv("VOLC_SECRETKEY"),
		},
	}
)

// SDKClient .
type SDKClient struct {
	Client *base.Client
}

// GetServiceInfo interface
func (p *SDKClient) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

func (p *SDKClient) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

func (p *SDKClient) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

func (p *SDKClient) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
