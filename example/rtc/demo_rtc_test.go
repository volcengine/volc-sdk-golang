package main

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/rtc"
)

const (
	appId  = ""
	testAk = ""
	testSk = ""
)

func TestListRooms(t *testing.T) {
	//init rtc service
	s := rtc.NewInstance()
	s.Client.SetAccessKey(testAk)
	s.Client.SetSecretKey(testSk)
	s.Client.SetHost(rtc.ServiceHost)

	query := url.Values{}
	query.Set("AppId", appId)
	//query.Set("RoomId", "zdl_room_20210818")
	res, _, err := rtc.ListRooms(s, query)
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
	//init rtc service
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
	//req.OS = "android"
	//req.Network = "wifi"
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
