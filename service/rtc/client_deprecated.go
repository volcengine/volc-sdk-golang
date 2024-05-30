package rtc

// client 这个文件由 Codegen 生成，但后续不再更新（如果没有，则会重新生成），包含配置结构体和创建服务实例的代码
// 开发者可以在这里给服务结构体添加自定义扩展方法

import (
	"fmt"

	common "github.com/volcengine/volc-sdk-golang/base"
)

type Rtc struct {
	*common.Client
}

// Deprecated: NewInstance is deprecated.
func NewInstance() *Rtc {
	return NewInstanceWithRegion("cn-north-1")
}

// Deprecated: NewInstanceWithRegion is deprecated.
func NewInstanceWithRegion(region string) *Rtc {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("Rtc not support region %s", region))
	}
	instance := &Rtc{
		Client: common.NewClient(&serviceInfo, ApiListInfo),
	}
	return instance
}
