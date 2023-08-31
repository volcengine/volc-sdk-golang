package main

import (
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func RetryRenderJobDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := verender.RetryJobRequest{
		JobID:           "r5ad3829bef",
		AllFailedFrames: true,
	}

	if err := w.RetryRenderJob(&r); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("RetryRenderJob done")
}
