package main

import (
	"fmt"
)

func UploadFolderDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	src := "/tmp/test_upload_folder"
	des := "/tmp/abc"
	err = w.UploadFolder(src, des)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("done")
}
