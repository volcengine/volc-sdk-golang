package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func AddRenderSettingDemo() {
	v := getVerenderInstance()
	workspaceId := int64(1935)
	w, err := v.GetWorkspace(workspaceId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := v.GetCurrentUser()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rs := verender.RenderSetting{
		AccountID:   user.AccountID,
		UserID:      user.UserID,
		WorkspaceID: w.Id,
		Name:        "test-sdk-add-render-setting-1",
		Dcc:         "maya",
		DccVersion:  "2022",
		Plugins: []*verender.Plugin{
			{
				Name:           "mtoa",
				Version:        "5.0.0",
				NeedLicense:    true,
				RendererPlugin: true,
			},
		},
		RenderLayerMode:    "LegacyRenderLayer",
		ProjectPath:        "",
		FramesCountOneCell: 1,
		CellSpecID:         9,
	}

	resp, err := w.AddRenderSetting(&rs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
