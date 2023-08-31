package main

import (
	"encoding/json"
	"fmt"
)

func GetRenderJobDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jobId := "r76199b2e96"
	resp, err := w.GetRenderJob(jobId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
