package main

import (
	"encoding/json"
	"fmt"
)

func ListFileDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prefix := "tmp/"
	filter := ""
	orderType := ""
	orderField := ""
	pageNum := int64(1)
	pageSize := int64(10)
	total, files, dirs, err := w.ListFile(prefix, filter, orderType, orderField, pageNum, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	subFiles, _ := json.Marshal(files)
	subDirs, _ := json.Marshal(dirs)
	fmt.Println(total)
	fmt.Println(string(subFiles))
	fmt.Println(string(subDirs))
}
