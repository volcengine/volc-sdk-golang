// 提示：本demo演示了OpenAPI鉴权SDK的使用方式
//
//	 	使用时请参考包github.com/volcengine/volc-sdk-golang/service/rtc的格式
//			把包内的文件拷贝到自己的项目中追加所需的API
//
// service/rtc目录包含三个文件：
//  1. config.go 定义API属性,引入签名包并初始化服务
//  2. datamodel.go 存放API的请求参数与返回参数的定义
//  3. wrappper.go API的方法体，主要分为两类即GET和POST
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

var (
	roomId     = "Your_roomId"
	busineseId = "Your_BusinessId"
	taskId     = appId + "_" + roomId + "_xxx"

	s *rtc.RTC
)

func TestMain(t *testing.M) {
	// init rtc service, 初始化一次即可
	s = rtc.NewInstance()
	s.Client.SetAccessKey(testAk)
	s.Client.SetSecretKey(testSk)
	s.Client.SetHost(rtc.ServiceHost)

	t.Run()
}

func TestGetRecordTask(t *testing.T) {
	query := url.Values{}
	query.Set("AppId", appId)
	query.Set("RoomId", roomId)
	query.Set("TaskId", taskId)

	res, _, err := rtc.GetRecordTask(s, query)
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

func TestStartRecord(t *testing.T) {
	req := rtc.StartRecordRequest{
		AppId:      appId,
		BusinessId: busineseId,
		RoomId:     roomId,
		TaskId:     taskId,
		Encode: &rtc.Encode{
			VideoWidth:   1920,
			VideoHeight:  1080,
			VideoFps:     15,
			VideoBitrate: 4000,
		},
		FileFormatConfig: &rtc.FileFormatConfig{
			FileFormat: []string{"MP4"},
		},
		StorageConfig: rtc.StorageConfig{
			TosConfig: &rtc.TosConfig{
				AccountId: "Your_Volc_AccountId",
				Bucket:    "Your_Bucket",
			},
		},
	}
	res, _, err := rtc.StartRecord(s, &req)
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

func TestStartWebRecord(t *testing.T) {
	req := rtc.StartWebRecordRequest{
		AppId:    appId,
		TaskId:   taskId,
		InputURL: "http://www.xxx.xxx/xxx/xxx.html",
		Bucket:   "Your_Bucket",
	}
	res, _, err := rtc.StartWebRecord(s, &req)
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

func TestGetWebRecordTask(t *testing.T) {
	//GetWebRecordTask
	query := url.Values{}
	query.Set("AppId", appId)
	query.Set("TaskId", taskId)

	res, _, err := rtc.GetWebRecordTask(s, query)
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

	//StopWebRecord
	stop := rtc.StopWebRecordRequest{
		AppId:  appId,
		TaskId: taskId,
	}
	stopResp, _, err := rtc.StopWebRecord(s, &stop)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if stopResp.ResponseMetadata.Error != nil {
		fmt.Printf("response err:%v\n", stopResp.ResponseMetadata.Error)
		return
	}
	fmt.Printf("metaData: %+v\n", stopResp.ResponseMetadata)
	fmt.Printf("result: %+v\n", stopResp.Result)

	//GetWebRecordList
	queryList := url.Values{}
	queryList.Set("AppId", appId)
	list, _, err := rtc.GetWebRecordList(s, queryList)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if list.ResponseMetadata.Error != nil {
		fmt.Printf("response err:%v\n", list.ResponseMetadata.Error)
		return
	}
	fmt.Printf("metaData: %+v\n", list.ResponseMetadata)
	fmt.Printf("result: %+v\n", list.Result)
}
