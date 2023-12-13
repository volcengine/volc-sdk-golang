package businessSecurity

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type SecuritySecurityClient struct {
	*base.Client
	//retry bool
	serviceInfo *base.ServiceInfo
	apiInfo     map[string]*base.ApiInfo
}

var SecuritySourceInstance = NewSecuritySourceInstance()

func NewSecuritySourceInstance() *SecuritySecurityClient {
	s := GetServiceInfo(base.RegionCnNorth1, "contentservice.zijieapi.com", 5*time.Minute)
	a := SecuritySourceApiInfoList
	instance := &SecuritySecurityClient{
		Client:      base.NewClient(s, a),
		serviceInfo: s,
		apiInfo:     a,
		//retry:  true,
	}
	return instance
}

//func (p *SecuritySecurityClient) Retry() bool {
//return p.retry
//}

//func (p *SecuritySecurityClient) CloseRetry() {
//	p.retry = false
//}

func (p *SecuritySecurityClient) SetRegion(region string) error {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		return fmt.Errorf("region does not spport or unknown region")
	}
	p.ServiceInfo = serviceInfo
	p.SetScheme("http")
	return nil
}

var (
	SecuritySourceApiInfoList = map[string]*base.ApiInfo{
		"SecuritySource": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Minute,
			Path:    "/openapi/v1/sse/security_source",
			Query: url.Values{
				"Action":  []string{"SecuritySource"},
				"Version": []string{"2022-08-26"},
			},
		},
		"SecuritySourceStream": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Minute,
			Path:    "/openapi/v1/sse/security_source_stream",
			Query: url.Values{
				"Action":  []string{"SecuritySourceStream"},
				"Version": []string{"2022-08-26"},
			},
		},
	}
)
