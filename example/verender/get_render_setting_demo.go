package main

import (
	"encoding/json"
	"fmt"
)

func GetRenderSettingDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	settingId := int64(861)
	resp, err := w.GetRenderSetting(settingId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
