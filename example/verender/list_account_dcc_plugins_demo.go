package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func ListAccountDccPluginsDemo() {
	v := getVerenderInstance()

	r := verender.ListAccountDccPluginsReq{
		SpecTemplateId: 54,
	}

	resp, err := v.ListAccountDccPlugins(&r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
