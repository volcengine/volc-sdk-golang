package acep

// client 这个文件由 Codegen 生成，但后续不再更新（如果没有，则会重新生成），包含配置结构体和创建服务实例的代码
// 开发者可以在这里给服务结构体添加自定义扩展方法

import (
	"fmt"
	"net/http"
	"net/url"

	common "github.com/volcengine/volc-sdk-golang/base"
)

type ACEP struct {
	*common.Client
}

func NewInstance() *ACEP {
	return NewInstanceWithRegion("cn-north-1")
}

func NewInstanceWithRegion(region string) *ACEP {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("ACEP not support region %s", region))
	}
	instance := &ACEP{
		Client: common.NewClient(&serviceInfo, ApiListInfo),
	}
	return instance
}

func (acep *ACEP) SetProxyHost(host, proxyUser, proxyPassword string) {
	acep.Client.Client.Transport = &http.Transport{
		Proxy: http.ProxyURL(&url.URL{
			Host:   host,
			Scheme: acep.ServiceInfo.Scheme,
			User:   url.UserPassword(proxyUser, proxyPassword),
		}),
	}
}
