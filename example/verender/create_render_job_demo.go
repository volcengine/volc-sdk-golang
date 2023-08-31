package main

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func CreateRenderJobDemo() {
	v := getVerenderInstance()
	workspaceID := int64(1935)
	w, err := v.GetWorkspace(workspaceID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	settingId := int64(157)
	rs, err := w.GetRenderSetting(settingId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fileName := "/Users/bytedance/Documents/jc_test/jc_sample/Advanced_07.ma"
	obj, err := w.UploadFile(fileName, fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	renderer := ""
	for _, p := range rs.Plugins {
		if p.RendererPlugin {
			renderer = p.Name
		}
	}

	user, err := v.GetCurrentUser()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	job := verender.RenderJob{
		AccountID:                user.AccountID,
		UserID:                   user.UserID,
		UserName:                 user.UserName,
		WorkspaceID:              workspaceID,
		Title:                    "test-create-render-job",
		Description:              "",
		Tryout:                   true,
		TryoutFrameNumbers:       []string{"START", "END", "MIDDLE"},
		SceneFile:                obj.Key,
		PathMapping:              nil,
		FramesCountEachCell:      8,
		TimeoutReminderEachFrame: 86400,
		TimeoutStopperEachFrame:  86400,
		LayerConfig: []*verender.LayerConfig{
			{
				LayerIndex: 0,
				LayerName:  "masterLayer",
				Frame: verender.Frame{
					Start: 1,
					End:   12000,
					Step:  1,
				},
				Resolutions: verender.Resolution{
					Width:  1920,
					Height: 1080,
				},
				Cameras: []string{"persp"},
			},
		},
		RenderSetting: rs,
		PluginData:    "{}",
		Priority:      0,
		Renderer:      renderer,
	}

	resp, err := w.CreateRenderJob(&job)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
