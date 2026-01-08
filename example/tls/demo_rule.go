package tls

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

type LogType = string

const (
	MinimalistLog LogType = "minimalist_log"
	MultilineLog  LogType = "multiline_log"
	DelimiterLog  LogType = "delimiter_log"
	JsonLog       LogType = "json_log"
	FullregexLog  LogType = "fullregex_log"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 请填写您的ProjectId、TopicId和HostGroupId
	projectID := "your-project-id"
	topicID := "your-topic-id"
	hostGroupID := "your-host-group-id"

	// 创建采集配置
	// 请根据您的需要，填写TopicId、RuleName和其它采集配置参数
	// CreateRule API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112199
	logType := "minimalist_log"
	logSample := "logSample"
	inputType := 1
	createRuleReq := tls.CreateRuleRequest{
		TopicID:     topicID,
		RuleName:    uuid.NewString(),
		Paths:       &[]string{fmt.Sprintf("/tmp/%s.log", uuid.NewString())},
		LogType:     &logType,
		ExtractRule: GenExtractRuleWithLogType(MinimalistLog),
		ExcludePaths: &[]tls.ExcludePath{
			{
				Type:  "File",
				Value: "/tmp/excludeFile.log",
			},
			{
				Type:  "Path",
				Value: "/tmp/excludePath/",
			},
		},
		UserDefineRule: &tls.UserDefineRule{
			ParsePathRule: &tls.ParsePathRule{
				PathSample: "/var/logs/2100101862_",
				Regex:      "\\/var\\/logs\\/([0-9]*)_guangzhou_([a-z0-9-]*)\\/access\\.log",
				Keys: []string{
					"accountId",
					"podName",
				},
			},
		},
		LogSample: &logSample,
		InputType: &inputType,
		ContainerRule: &tls.ContainerRule{
			Stream:             "stdout",
			ContainerNameRegex: "ContainerNameRegex",
			IncludeContainerLabelRegex: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			ExcludeContainerLabelRegex: map[string]string{
				"label1": "value1",
				"label2": "value2",
			},
			IncludeContainerEnvRegex: map[string]string{
				"key1": "value1",
			},
			ExcludeContainerEnvRegex: map[string]string{
				"key1": "value1",
			},
			EnvTag: map[string]string{
				"key1": "value1",
			},
			KubernetesRule: tls.KubernetesRule{
				NamespaceNameRegex: "ns",
				WorkloadType:       "Deployment",
				WorkloadNameRegex:  "WorkloadNameRegex",
				PodNameRegex:       "podNamereg",
				IncludePodLabelRegex: map[string]string{
					"key1": "value1",
				},
				ExcludePodLabelRegex: map[string]string{
					"key1": "value1",
				},
				LabelTag: map[string]string{
					"key1": "value1",
				},
			},
		},
	}
	createRuleResp, _ := client.CreateRule(&createRuleReq)
	ruleID := createRuleResp.RuleID

	// 修改采集配置
	// 请根据您的需要，填写待修改的RuleId、RuleName或其它参数
	// ModifyRule API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112201
	modifyRuleName := uuid.NewString()
	modifyPath := fmt.Sprintf("/tmp/%s.log", uuid.NewString())
	_, _ = client.ModifyRule(&tls.ModifyRuleRequest{
		RuleID:   ruleID,
		RuleName: &modifyRuleName,
		Paths:    &[]string{modifyPath},
	})

	// 应用采集配置到机器组
	// 请根据您的需要，填写RuleId和HostGroupIds列表
	// ApplyRuleToHostGroups API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112204
	_, _ = client.ApplyRuleToHostGroups(&tls.ApplyRuleToHostGroupsRequest{
		RuleID:       ruleID,
		HostGroupIDs: []string{hostGroupID},
	})

	// 删除机器组的采集配置
	// 请根据您的需要，填写RuleId和HostGroupIds列表
	// DeleteRuleFromHostGroups API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112205
	_, _ = client.DeleteRuleFromHostGroups(&tls.DeleteRuleFromHostGroupsRequest{
		RuleID:       ruleID,
		HostGroupIDs: []string{hostGroupID},
	})

	// 查询指定采集配置
	// 请根据您的需要，填写待查询的RuleId
	// DescribeRule API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112202
	describeRuleResp, _ := client.DescribeRule(&tls.DescribeRuleRequest{
		RuleID: ruleID,
	})
	fmt.Println(describeRuleResp.RuleInfo.RuleName)

	// 查询指定采集配置 V2
	// 请根据您的需要，填写待查询的RuleId
	// DescribeRule API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/2121515
	describeRuleRespV2, err := client.DescribeRuleV2(&tls.DescribeRuleRequestV2{
		RuleID: ruleID,
	})
	if err == nil && describeRuleRespV2 != nil && describeRuleRespV2.RuleInfo != nil {
		fmt.Println(describeRuleRespV2.RuleInfo.RuleName)
	}

	// 查询日志项目所有采集配置
	// 请根据您的需要，填写待查询的ProjectId
	// DescribeRules API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112203
	describeRulesResp, _ := client.DescribeRules(&tls.DescribeRulesRequest{
		ProjectID:  projectID,
		PageNumber: 1,
		PageSize:   10,
	})
	fmt.Println(describeRulesResp.Total)

	// 删除采集配置
	// 请根据您的需要，填写待删除的RuleId
	// DeleteRule API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/112200
	_, _ = client.DeleteRule(&tls.DeleteRuleRequest{
		RuleID: ruleID,
	})

	// 获取采集配置绑定的机器组信息
	// 请根据您的需要，填写待删除的RuleId
	// describeBoundHostGroupsResp API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/2121516
	describeBoundHostGroupsResp, err := client.DescribeBoundHostGroups(&tls.DescribeBoundHostGroupsRequest{
		RuleID: ruleID,
	})
	if err == nil && describeBoundHostGroupsResp != nil {
		fmt.Println(describeBoundHostGroupsResp.Total)
	}
}

func GenExtractRuleWithLogType(logType LogType) *tls.ExtractRule {
	switch logType {
	case MinimalistLog:
		return GenExtractRule(logType, false, false, false, false, false, false, true, true, true)
	case MultilineLog:
		return GenExtractRule(logType, false, true, false, false, false, false, true, true, true)
	case DelimiterLog:
		return GenExtractRule(logType, true, false, false, true, true, true, true, true, true)
	case JsonLog:
		return GenExtractRule(logType, false, false, false, false, true, true, true, true, true)
	case FullregexLog:
		return GenExtractRule(logType, false, false, true, true, true, true, true, true, true)
	default:
		return GenExtractRule(logType, false, false, true, true, true, true, true, true, true)
	}
}

func GenExtractRule(logType LogType, setDelimiter, setBeginRegex, setLogRegex, setKeys, setTimeKey, setTimeFormat, setFilterKeyRegex, setUnMatchUpLoadSwitch, setUnMatchLogKey bool) *tls.ExtractRule {
	var extractRule = tls.ExtractRule{}

	if setDelimiter {
		extractRule.Delimiter = "#"
	}
	if setBeginRegex {
		extractRule.BeginRegex = "beginRegx"
	}
	if setLogRegex {
		extractRule.LogRegex = "logRegex"
	}
	if setKeys {
		extractRule.Keys = []string{"key"}
	}
	if setTimeKey {
		extractRule.TimeKey = "time"
	}
	if setTimeFormat {
		extractRule.TimeFormat = "%Y-%m-%dT%H:%M:%S,%f"
	}
	if setFilterKeyRegex {
		if logType == MinimalistLog || logType == MultilineLog {
			extractRule.FilterKeyRegex = []tls.FilterKeyRegex{
				{
					Key:   "__content__",
					Regex: "regex",
				},
			}
		} else {
			extractRule.FilterKeyRegex = []tls.FilterKeyRegex{
				{
					Key:   "key1",
					Regex: "regex1",
				},
				{
					Key:   "key2",
					Regex: "regex2",
				},
				{
					Key:   "key3",
					Regex: "regex3",
				},
				{
					Key:   "key4",
					Regex: "regex4",
				},
				{
					Key:   "key5",
					Regex: "regex5",
				},
			}
		}
	}
	if setUnMatchUpLoadSwitch {
		extractRule.UnMatchUpLoadSwitch = true
	}

	if setUnMatchLogKey {
		extractRule.UnMatchLogKey = "LogParseFailed"
	}

	return &extractRule
}
