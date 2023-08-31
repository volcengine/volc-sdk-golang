package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func ListRenderJobDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := verender.ListRenderJobRequest{
		PageNum:        1,
		PageSize:       10,
		OrderType:      "",
		OrderField:     "",
		Keyword:        "",
		Status:         nil,
		RenderSettings: nil,
		DeviceName:     "",
		UserId:         0,
	}
	resp, err := w.ListRenderJob(&r)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
