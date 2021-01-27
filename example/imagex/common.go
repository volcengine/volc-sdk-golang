package main

import (
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * common api request demo
 * for api definition, refer to https://www.volcengine.cn/docs/508/14106
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.volc/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	query := url.Values{}
	query.Add("ServiceId", "imagex service id")
	query.Add("StoreUri", "image uri")

	resp := new(struct {
		ServiceId    string `json:"ServiceId"`
		FileName     string `json:"FileName"`
		StoreUri     string `json:"StoreUri"`
		LastModified string `json:"LastModified"`
		FileSize     int    `json:"FileSize"`
		ThumbUrl     string `json:"ThumbUrl"`
		Marker       int    `json:"Marker"`
	})

	err := instance.ImageXGet("GetImageUploadFile", query, resp)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
