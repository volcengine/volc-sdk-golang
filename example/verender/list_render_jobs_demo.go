package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func ListJobOutputDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := verender.ListJobOutputRequest{
		StartTime:  "",
		EndTime:    "",
		PageNum:    1,
		PageSize:   10,
		Type:       "",
		Status:     "",
		OrderType:  "",
		OrderField: "",
		JobIDList:  nil,
	}

	resp, err := w.ListJobOutput(&r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
