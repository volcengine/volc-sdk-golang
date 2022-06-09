package upload

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/functions"
)

func TestVod_UploadMediaWithCallback(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	spaceName := "your space"
	filePath := "media file path"

	getMetaFunc := functions.GetMetaFunc()                         // 抽取视频元信息&封面功能，如需要则添加
	snapShotFunc := functions.SnapshotFunc(1.3)                    // 抽取特定时刻的封面截图
	startWorkFlowFunc := functions.StartWorkflowFunc("templateId") // 如希望上传完成后自动执行转码工作流，可将工作流Id填写在此函数里
	optionFunc := functions.AddOptionInfoFunc(vod.OptionInfo{
		Title:            "title",     // 视频的标题
		Tags:             "Go,编程",     // 视频的标签
		Description:      "Go 语言高级编程", // 视频的描述信息
		Format:           "MP4",       // 音视频格式
		ClassificationId: 0,           // 分类 Id，上传时可以指定分类，非必须字段
	})
	vodFunctions := []vod.Function{snapShotFunc, getMetaFunc, startWorkFlowFunc, optionFunc}
	fbts, _ := json.Marshal(vodFunctions)

	vodUploadMediaRequset := &request.VodUploadMediaRequest{
		SpaceName:    spaceName,     // 空间名称
		FilePath:     filePath,      // 本地文件路径
		CallbackArgs: "my callback", // 透传信息，业务希望透传的字段可以写入，返回和回调中会返回此字段，非必须字段
		Functions:    string(fbts),  // 函数功能，具体可以参考火山引擎点播文档 开发者API-媒资上传-确认上传的 Functions 部分，可选功能字段
		FileName:     "",            // 设置文件名，无格式长度限制，用户可自定义,目前文件名不支持空格、+ 字符,如果要使用此字段，请联系技术支持配置白名单，非必须字段
	}

	resp, _, err := instance.UploadMediaWithCallback(vodUploadMediaRequset)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("resp = %s", bts)
	}

	fmt.Println()
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	fmt.Println(resp.GetResult().GetData().GetVid())

}
