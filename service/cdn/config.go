package cdn

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultRegion  = "cn-north-1"
	ServiceVersion = "2021-03-01"
	ServiceName    = "CDN"
)

var (
	ServiceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Host:    "cdn.volcengineapi.com",
			Timeout: time.Minute * 5,
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
	}
)

// DefaultInstance Package level default instance
var DefaultInstance = NewInstance()

type CDN struct {
	Client *base.Client
}

func NewInstance() *CDN {
	instance := new(CDN)
	instance.Client = base.NewClient(ServiceInfo[DefaultRegion], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

func (s *CDN) GetServiceInfo(region string) *base.ServiceInfo {
	if serviceInfo, ok := ServiceInfo[region]; ok {
		return serviceInfo
	}
	return nil
}

func (s *CDN) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

func (s *CDN) GetMethod(api string) string {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo.Method
	}
	return ""
}

func (s *CDN) SetRegion(region string) {
	s.Client.ServiceInfo.Credentials.Region = region
}

func (s *CDN) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *CDN) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}

func (s *CDN) SetMethod(api, method string) bool {
	if _, ok := ApiInfoList[api]; ok {
		ApiInfoList[api].Method = strings.ToUpper(method)
		return true
	}
	return false
}

func (s *CDN) SetAPIRetrySetting(apis []string, retryTime int, interval time.Duration) bool {
	if len(apis) <= 0 {
		return false
	}
	_retryTime, _interval := uint64(retryTime), interval
	for _, api := range apis {
		if _, ok := s.Client.ApiInfoList[api]; ok {
			s.Client.ApiInfoList[api].Retry = base.RetrySettings{
				AutoRetry:     true,
				RetryTimes:    &_retryTime,
				RetryInterval: &_interval,
			}
		} else {
			return false
		}
	}
	s.Client.SetRetrySettings(&base.RetrySettings{AutoRetry: true})
	return true
}

func (s *CDN) SetNewAPI(info base.ApiInfo) {
	if len(info.Query.Get("Action")) <= 0 {
		return
	}
	action := info.Query.Get("Action")
	if _, ok := s.Client.ApiInfoList[action]; ok {
		return
	}
	s.Client.ApiInfoList[action] = &info
}