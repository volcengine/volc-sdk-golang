package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func ListWorkspaceDemo() {
	v := getVerenderInstance()
	r := verender.ListWorkspaceRequest{
		WorkspaceID:        0,
		WorkspaceName:      "",
		PageNum:            0,
		PageSize:           0,
		StartPurchaseTime:  "",
		EndPurchaseTime:    "",
		OrderType:          "",
		OrderField:         "",
		FuzzyWorkspaceName: "",
		FuzzySearchContent: "",
		ListScope:          "",
	}

	resp, err := v.ListWorkspace(&r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
