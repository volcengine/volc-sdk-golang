// 描述：演示了OpenAPI鉴权SDK的使用方式
//  	使用时请参考包github.com/volcengine/volc-sdk-golang/service/rtc或把它拷贝到自己的项目中追加所需的API
//
// service/rtc目录包含三个文件：
//	1. config.go 定义API属性,引入签名包并初始化服务
//	2. model.go 存放API的请求参数与返回参数的定义
//	3. wrappper.go API的方法体，主要分为两类即GET和POST
package main

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/rtc"
)

const (
	// ak/sk 获取方式参考:https://www.volcengine.com/docs/6348/69828
	appId  = ""
	testAk = ""
	testSk = ""
)

func TestListRoomInformation(t *testing.T) {
	// init rtc service, 初始化一次即可
	s := rtc.NewInstance()
	s.Client.SetAccessKey(testAk)
	s.Client.SetSecretKey(testSk)
	s.Client.SetHost(rtc.ServiceHost)

	query := url.Values{}
	query.Set("AppId", appId)
	query.Set("StartTime", "2022-01-22T12:00:00+08:00")
	query.Set("EndTime", "2022-05-22T12:59:00+08:00")

	res, _, err := rtc.ListRoomInformation(s, query)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if res.ResponseMetadata.Error != nil {
		fmt.Printf("response err:%v\n", res.ResponseMetadata.Error)
		return
	}
	fmt.Printf("metaData: %+v\n", res.ResponseMetadata)
	fmt.Printf("result: %+v\n", res.Result)
}

func TestListIndicators(t *testing.T) {
	// init rtc service
	s := rtc.NewInstance()
	s.Client.SetAccessKey(testAk)
	s.Client.SetSecretKey(testSk)
	s.Client.SetHost(rtc.ServiceHost)

	req := rtc.ListIndicatorsRequest{
		AppId:     appId,
		StartTime: "2021-08-17T00:00:00+08:00",
		EndTime:   "2021-08-18T00:00:00+08:00",
		Indicator: "NetworkTransDelay",
	}
	// req.OS = "android"
	// req.Network = "wifi"
	res, _, err := rtc.ListIndicators(s, &req)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if res.ResponseMetadata.Error != nil {
		fmt.Printf("response err:%v\n", res.ResponseMetadata.Error)
		return
	}
	fmt.Printf("metaData: %+v\n", res.ResponseMetadata)
	fmt.Printf("result: %+v\n", res.Result)
}
