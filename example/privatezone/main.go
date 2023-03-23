package main

import (
	"context"
	"fmt"
	"os"

	pz "github.com/volcengine/volc-sdk-golang/service/privatezone"
)

func main() {
	// AK, SK从环境变量获取
	os.Setenv("VOLC_ACCESSKEY", "your_ak")
	os.Setenv("VOLC_SECRETKEY", "your_sk")

	client := pz.InitDNSVolcClient()

	resp, err := client.ListRegions(context.Background(), &pz.ListRegionsRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
