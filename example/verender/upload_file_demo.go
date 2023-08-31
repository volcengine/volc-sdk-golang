package main

import (
	"encoding/json"
	"fmt"
)

func UploadFileDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	src := "/tmp/512M.dat"
	des := "/tmp/512M.dat"
	resp, err := w.UploadFile(src, des)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
