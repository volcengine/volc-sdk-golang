package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * update image urls
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.volc/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	serviceId := "imagex service id"
	urls := []string{"image url 1"}

	resp, err := instance.RefreshImageUrls(serviceId, urls)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
