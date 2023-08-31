package main

import (
	"fmt"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func UpdateRenderSettingDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rs := verender.RenderSetting{
		Id:   128,
		Name: fmt.Sprintf("test-sdk-update-render-setting-%d", time.Now().Unix()),
	}

	if err := w.UpdateRenderSetting(&rs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("UpdateRenderSettingDone")
}
