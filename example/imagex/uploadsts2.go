package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * get image upload auth
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.volc/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	// service id list allowed to do upload, pass empty list if no restriction
	serviceIds := []string{"imagex service id"}

	// set expire time by GetUploadAuthWithExpire, default is 1 hour
	// set key pattern with imagex.WithUploadKeyPtn("pattern"). eg: "test/*" means upload key must have prefix "test/"
	token, err := instance.GetUploadAuth(serviceIds)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("token %+v", token)
	}
}
