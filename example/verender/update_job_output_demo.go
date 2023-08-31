package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func UpdateJobOutputDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := verender.UpdateJobOutputRequest{
		JobID: "rb3bf60e847",
		Files: []string{"Result/test-create-render-job_rb3bf60e847/thumbs/rb3bf60e847_masterLayer_4.png"},
	}

	if err := w.UpdateJobOutput(&r); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("UpdateJobOutput done")
}
