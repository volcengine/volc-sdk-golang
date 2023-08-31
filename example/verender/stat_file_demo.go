package main

import (
	"encoding/json"
	"fmt"
)

func StatFileDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := "tmp/a+b.txt"
	fo, err := w.StatFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(fo)
	fmt.Println(string(data))
}
