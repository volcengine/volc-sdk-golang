package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func UpdateRenderJobsPriorityDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := verender.BatchJobPriority{
		JobIDList: []string{"r76199b2e96"},
		Priority:  2,
	}

	if err := w.UpdateRenderJobsPriority(&r); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("UpdateRenderJobPriority done")
}
