package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/faas"
)

func main() {
	instance := faas.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "xxx",
		SecretAccessKey: "yyy",
	})

	req := &faas.ListTriggersRequest{
		FunctionId: "xnxckabc",
	}

	resp, status, err := instance.ListTriggers(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status: %d", status)
	data, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))

}
