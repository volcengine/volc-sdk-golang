package functions

import "github.com/volcengine/volc-sdk-golang/service/vod"

func GetMeatFunc() vod.Function {
	return vod.Function{Name: "GetMeta"}
}

func SnapshotFunc(snapshotTime float64) vod.Function {
	return vod.Function{Name: "Snapshot", Input: vod.SnapshotInput{SnapshotTime: snapshotTime}}
}

func AddOptionInfoFunc(info vod.OptionInfo) vod.Function {
	return vod.Function{
		Name:  "AddOptionInfo",
		Input: info,
	}
}

func StartWorkflowFunc(templateId string) vod.Function {
	return vod.Function{
		Name:  "StartWorkflow",
		Input: vod.WorkflowInput{TemplateId: templateId},
	}
}

func EncryptionFunc(encryption vod.EntryptionInput) vod.Function {
	return vod.Function{
		Name:  "Encryption",
		Input: encryption,
	}
}
