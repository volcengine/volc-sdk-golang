package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

/*
 * upload local images
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.DefaultInstance

	// call below method if you dont set ak and sk in ～/.volc/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	params := &imagex.ApplyUploadImageParam{
		ServiceId: "imagex service id",
	}
	dat, err := ioutil.ReadFile("image file")
	if err != nil {
		fmt.Printf("read file from %s error %v", "", err)
		os.Exit(-1)
	}
	resp, err := instance.UploadImages(params, [][]byte{dat})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
