package main

import (
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * get image upload token
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
	query.Set("ServiceId", "imagex service id")
	// set expire time of the upload token, default is 15min(900),
	query.Set("X-Amz-Expires", "60")

	token, err := instance.GetUploadAuthToken(query)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("token %s", token)
	}
}
