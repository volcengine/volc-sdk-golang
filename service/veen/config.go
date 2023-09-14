package veen

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"time"
)

const (
	ServiceName    = "veenedge"
	ServiceVersion = "2021-04-30"
)

var (
	// 注册服务
	ServiceInfo = &base.ServiceInfo{
		Timeout: 15 * time.Second,
		Host:    "veenedge.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{
			Service: ServiceName,
			Region:  base.RegionCnNorth1,
		},
		Scheme: "https",
	}
)

type Veen struct {
	Client *base.Client
}

func NewInstance(ak, sk string) *Veen {
	// 初始化client
	instance := &Veen{
		Client: base.NewClient(ServiceInfo, ApiInfoList),
	}
	// 设置Access Key和Secret Key
	instance.Client.SetAccessKey(ak)
	instance.Client.SetSecretKey(sk)
	return instance
}

// GetServiceInfo interface
func (v *Veen) GetServiceInfo(env string) *base.ServiceInfo {
	return v.Client.ServiceInfo
}

// GetAPIInfo interface
func (v *Veen) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion
func (v *Veen) SetRegion(env, region string) {
	v.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (v *Veen) SetHost(host string) {
	v.Client.ServiceInfo.Host = host
}

// SetSchema .
func (v *Veen) SetSchema(schema string) {
	v.Client.ServiceInfo.Scheme = schema
}
