package main

import (
	"encoding/json"
	"fmt"
)

func ListRenderSettingDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dcc := "maya"
	resp, err := w.ListRenderSetting(dcc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)

	fmt.Println(string(data))
}
