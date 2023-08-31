package main

import (
	"fmt"
)

func RemoveFileDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := "tmp/512M.dat"
	if err := w.RemoveFile(filename); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("remove file done")
}
