package main

import (
	"fmt"
)

func DeleteRenderSettingDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	settingId := int64(1191)
	if err := w.DeleteRenderSetting(settingId); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("DeleteRenderSetting done")
}
