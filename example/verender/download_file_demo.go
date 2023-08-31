package main

import (
	"fmt"
)

func DownloadFileDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	src := "tmp/512M.dat"
	dst := "/tmp/512M_p1.dat"
	if err := w.DownloadFile(src, dst); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("download file %s done\n", src)
}
