package tls

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKRuleTestSuite struct {
	suite.Suite

	cli     Client
	project string
	topic   string
	rule    string
}

func (suite *SDKRuleTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectId, err := CreateProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := CreateTopic(projectId, "golang-sdk-create-index-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	keyValueList := make([]KeyValueInfo, 0)
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-1",
		Value: Value{
			ValueType:      "text",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-2",
		Value: Value{
			ValueType:      "long",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	suite.NoError(CreateIndex(topicId, nil, &keyValueList, suite.cli))
}

func (suite *SDKRuleTestSuite) TearDownTest() {
	_, deleteRuleErr := suite.cli.DeleteRule(&DeleteRuleRequest{RuleID: suite.rule})
	suite.NoError(deleteRuleErr)
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKRuleTestSuite(t *testing.T) {
	suite.Run(t, new(SDKRuleTestSuite))
}

func (suite *SDKRuleTestSuite) TestRule() {
	var ruleID string
	// CreateRule
	{
		requestContent := `{
			"TopicId": "706f****-****-****-****-****46aa1f07",
			"RuleName": "testname",
			"Paths": [
				"/data/nginx/log/*/*/*.log"
			],
			"ExcludePaths":[
				{
					"Type": "File",
					"Value": "/data/nginx/log/*/*/exclude.log"
				},
				{
					"Type": "Path",
					"Value": "/data/nginx/log/*/exclude/"
				}
			],
			"LogType": "fullregex_log",
			"ExtractRule": {
				"LogRegex": "(\\S*)\\s*-\\s*(\\S*)\\s*\\[(\\d+/\\S+/\\d+:\\d+:\\d+:\\d+)\\s+\\S+\\]\\s*\"(\\S+)\\s+(\\S+)\\s+(\\S+)\"\\s*(\\S*)\\s*(\\S*)\\s*(\\S*)\\s*(\\S*)\\s*\"([^\"]*)\"\\s*\"([^\"]*)\".*",
				"Keys": [
					"remote_addr",
					"remote_user",
					"time_local",
					"request_method",
					"request_uri",
					"protocol",
					"request_time",
					"request_length",
					"status",
					"body_bytes_sent",
					"http_referer",
					"http_user_agent"
				],
				"TimeKey": "request_time",
				"TimeFormat": "%d/%b/%Y:%H:%M:%S",
				"FilterKeyRegex": [
					{
						"Key": "status",
						"Regex": "200"
					}
				],
				"LogTemplate": {
					"Type": "Nginx",
					"Format": "format main  '$remote_addr - $remote_user [$time_local] \"$request\" $request_time $request_length $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\"';"
				},
				"UnMatchUpLoadSwitch": true,
				"UnMatchLogKey": "LogParseFailed"
			},
			"LogSample": "192.168.1.1 - - [24/Sep/2022:12:23:12 +0800] \"POST /abc.com.testfile HTTP/1.0\" 0.000 129 200 43912736 \"-\" \"Wget/1.11.4 Red Hat modified\"",
			"UserDefineRule": {
				"ParsePathRule": {
					"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
					"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
					"Keys": ["instance-id", "pod-name"]
				},
				"ShardHashKey": {
					"HashKey": "3C"
				},
				"EnableRawLog": true,
				"Fields": {
					"ClusterID":"dabaad5f-7a10-4771-b3ea-d821f73e****"
				},
				"Advanced": {
					"CloseInactive": 10
				}
			}
		}`
		request := &CreateRuleRequest{}
		json.Unmarshal([]byte(requestContent), request)
		request.TopicID = suite.topic

		request.LogType = new(string)
		*request.LogType = "fullregex_log"

		resp, err := suite.cli.CreateRule(request)
		suite.NoError(err)
		ruleID = resp.RuleID
		suite.rule = resp.RuleID
	}

	// ModifyRule
	{
		request := &ModifyRuleRequest{
			RuleID: ruleID,
			UserDefineRule: &UserDefineRule{
				Fields: map[string]string{
					"TestKey": "TestValue",
				},
			},
		}
		_, err := suite.cli.ModifyRule(request)
		suite.NoError(err)
	}

	// DescribeRule
	{
		request := &DescribeRuleRequest{
			RuleID: ruleID,
		}
		_, err := suite.cli.DescribeRule(request)
		suite.NoError(err)
	}

	// DescribesRule
	{
		request := &DescribeRulesRequest{
			ProjectID: suite.project,
		}
		_, err := suite.cli.DescribeRules(request)
		suite.NoError(err)
	}

	// CreateRuleInvalid
	{
		var request = &CreateRuleRequest{
			TopicID: suite.topic,
		}
		_, err := suite.cli.CreateRule(request)
		suite.Error(err)
	}

	// ModifyRuleInvalid
	{
		var request = &ModifyRuleRequest{
			RuleID:  ruleID,
			LogType: new(string),
		}
		*request.LogType = "test"
		_, err := suite.cli.ModifyRule(request)
		suite.Error(err)
	}

	// DescribeRuleInvalid
	{
		var request = &DescribeRuleRequest{
			RuleID: "test",
		}
		_, err := suite.cli.DescribeRule(request)
		suite.Error(err)
	}

	// DescribeRulesInvalid
	{
		var request = &DescribeRulesRequest{
			ProjectID: "test",
		}
		_, err := suite.cli.DescribeRules(request)
		suite.Error(err)
	}
}
