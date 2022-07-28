package tls

import (
	"fmt"
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
	//初始化客户端，配置AccessKeyID,AccessKeySecret,region,securityToken;securityToken可以为空
	client := tls.NewClient(testEndPoint, testAk, testSk, testSessionToken, testRegion)

	//新建project
	createResp, _ := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: testPrefix + uuid.NewString(),
		Description: "",
		Region:      testRegion,
	})

	testProjectID := createResp.ProjectID

	// 新建topic
	// TopicName Description字段规范参考api文档
	createTopicRequest := &tls.CreateTopicRequest{
		ProjectID:   testProjectID,
		TopicName:   testPrefix + uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic desc",
	}
	topic, _ := client.CreateTopic(createTopicRequest)
	testTopicID := topic.TopicID

	// 新建rule
	logType := "minimalist_log"
	logSample := "logSample"
	inputType := 1
	createRule := tls.CreateRuleRequest{
		TopicID:  testTopicID,
		RuleName: testPrefix + uuid.NewString(),
		//Paths:       &[]string{fmt.Sprintf("/tmp/%s.log", uuid.NewString())},
		LogType:     &logType,
		ExtractRule: GenExtractRuleWithLogType(MinimalistLog),
		//ExcludePaths: &[]sdk.ExcludePath{
		//	{
		//		Type:  "File",
		//		Value: "/tmp/excludeFile.log",
		//	},
		//	{
		//		Type:  "Path",
		//		Value: "/tmp/excludePath/",
		//	},
		//},
		//UserDefineRule: &sdk.UserDefineRule{
		//	ParsePathRule: &sdk.ParsePathRule{
		//		PathSample: "/var/logs/2100101862_",
		//		Regex:      "\\/var\\/logs\\/([0-9]*)_guangzhou_([a-z0-9-]*)\\/access\\.log",
		//		Keys: []string{
		//			"accountId",
		//			"podName",
		//		},
		//	},
		//},
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
	rule, _ := client.CreateRule(&createRule)
	testRuleID := rule.RuleID

	// modify rule
	modifyRuleName := testPrefix + uuid.NewString()
	modifyPath := fmt.Sprintf("/tmp/%s.log", uuid.NewString())
	_, _ = client.ModifyRule(&tls.ModifyRuleRequest{
		RuleID:   testRuleID,
		RuleName: &modifyRuleName,
		Paths:    &[]string{modifyPath},
	})

	// add hostGroup
	hostIdentifier := "label"
	hostIPList := []string{"192.168.0.1"}
	createHostGroupReq := tls.CreateHostGroupRequest{
		HostGroupName:  "mgn1",
		HostGroupType:  "IP",
		HostIdentifier: &hostIdentifier,
		HostIPList:     &hostIPList,
	}

	hostGroup, _ := client.CreateHostGroup(&createHostGroupReq)

	hostGroupID := hostGroup.HostGroupID

	// apply rule to host group
	_, _ = client.ApplyRuleToHostGroups(&tls.ApplyRuleToHostGroupsRequest{
		RuleID:       testRuleID,
		HostGroupIDs: []string{hostGroupID},
	})

	// delete rule from host group
	_, _ = client.DeleteRuleFromHostGroups(&tls.DeleteRuleFromHostGroupsRequest{
		RuleID:       testRuleID,
		HostGroupIDs: []string{hostGroupID},
	})

	// describe rule
	_, _ = client.DescribeRule(&tls.DescribeRuleRequest{
		RuleID: testRuleID,
	})

	// describe rules
	_, _ = client.DescribeRules(&tls.DescribeRulesRequest{
		ProjectID:  testProjectID,
		PageNumber: 1,
		PageSize:   10,
	})
	// delete rule
	_, _ = client.DeleteRule(&tls.DeleteRuleRequest{
		RuleID: testRuleID,
	})
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
				tls.FilterKeyRegex{
					Key:   "__content__",
					Regex: "regex",
				},
			}
		} else {
			extractRule.FilterKeyRegex = []tls.FilterKeyRegex{
				tls.FilterKeyRegex{
					Key:   "key1",
					Regex: "regex1",
				},
				tls.FilterKeyRegex{
					Key:   "key2",
					Regex: "regex2",
				},
				tls.FilterKeyRegex{
					Key:   "key3",
					Regex: "regex3",
				},
				tls.FilterKeyRegex{
					Key:   "key4",
					Regex: "regex4",
				},
				tls.FilterKeyRegex{
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
