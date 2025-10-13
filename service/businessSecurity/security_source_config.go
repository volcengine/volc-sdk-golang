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
	retry       bool
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
		retry:       true,
	}
	return instance
}

func (p *SecuritySecurityClient) Retry() bool {
	return p.retry
}

func (p *SecuritySecurityClient) CloseRetry() {
	p.retry = false
}

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
		"TextModeration": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Minute,
			Path:    "/openapi/v1/rc_llm/text_moderation",
			Query: url.Values{
				"Action":  []string{"TextModeration"},
				"Version": []string{"2022-08-26"},
			},
		},
		"AsyncTextModeration": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/openapi/v1/rc_llm/async_text_moderation",
			Query: url.Values{
				"Action":  []string{"AsyncTextModeration"},
				"Version": []string{"2022-08-26"},
			},
		},
		"TextModerationResult": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/openapi/v1/rc_llm/text_moderation_result",
			Query: url.Values{
				"Action":  []string{"TextModerationResult"},
				"Version": []string{"2022-08-26"},
			},
		},
		"MultiModeration": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Minute,
			Path:    "/openapi/v1/rc_llm/multi_moderation",
			Query: url.Values{
				"Action":  []string{"MultiModeration"},
				"Version": []string{"2022-08-26"},
			},
		},
		"AsyncMultiModeration": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/openapi/v1/rc_llm/async_multi_moderation",
			Query: url.Values{
				"Action":  []string{"AsyncMultiModeration"},
				"Version": []string{"2022-08-26"},
			},
		},
		"MultiModerationResult": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/openapi/v1/rc_llm/multi_moderation_result",
			Query: url.Values{
				"Action":  []string{"MultiModerationResult"},
				"Version": []string{"2022-08-26"},
			},
		},
		"CustomRisk": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Minute,
			Path:    "/openapi/v1/rc_llm/custom_risk",
			Query: url.Values{
				"Action":  []string{"CustomRisk"},
				"Version": []string{"2022-08-26"},
			},
		},
		"AsyncCustomRisk": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/openapi/v1/rc_llm/async_custom_risk",
			Query: url.Values{
				"Action":  []string{"AsyncCustomRisk"},
				"Version": []string{"2022-08-26"},
			},
		},
		"CustomRiskResult": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/openapi/v1/rc_llm/custom_risk_result",
			Query: url.Values{
				"Action":  []string{"CustomRiskResult"},
				"Version": []string{"2022-08-26"},
			},
		},
	}
)
