package tls

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type SDKRuleTestSuite struct {
	suite.Suite

	ruleTypeList          []string
	jsonFormatRulesMap    map[string]string
	createRuleRequestsMap map[string]*CreateRuleRequest

	cli      Client
	project  string
	topic    string
	ruleList []string
}

func (suite *SDKRuleTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()

	projectID, err := CreateProject("golang-sdk-create-project-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectID

	topicID, err := CreateTopic(projectID, "golang-sdk-create-topic-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicID

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
	suite.NoError(CreateIndex(topicID, nil, &keyValueList, suite.cli))

	suite.jsonFormatRulesMap = getJsonFormatRulesMap()
	suite.createRuleRequestsMap, err = getCreateRuleRequestsMap(suite.jsonFormatRulesMap, suite.topic)
	suite.NoError(err)

	for key := range suite.jsonFormatRulesMap {
		suite.ruleTypeList = append(suite.ruleTypeList, key)
	}
}

func (suite *SDKRuleTestSuite) TearDownTest() {
	for _, ruleId := range suite.ruleList {
		_, deleteRuleErr := suite.cli.DeleteRule(&DeleteRuleRequest{RuleID: ruleId})
		suite.NoError(deleteRuleErr)
	}
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)

	suite.ruleList = nil
}

func (suite *SDKRuleTestSuite) validateError(err error, expectErr *Error) {
	sdkErr, ok := err.(*Error)

	if sdkErr == nil {
		suite.Nil(sdkErr)
		return
	}

	suite.Equal(true, ok)
	suite.Equal(expectErr.HTTPCode, sdkErr.HTTPCode)
	suite.Equal(expectErr.Code, sdkErr.Code)
	suite.Equal(expectErr.Message, sdkErr.Message)
}

func TestSDKRuleTestSuite(t *testing.T) {
	suite.Run(t, new(SDKRuleTestSuite))
}

func getJsonFormatRulesMap() map[string]string {
	jsonFormatRules := make(map[string]string)

	jsonFormatRules["单行全文模式"] = `{
	    "TopicId": "706f****-****-****-****-****46aa1f07",
	    "RuleName": "testname",
	    "Paths": [
	        "/data/nginx/log/*/*/*.log"
	    ],
	    "ExcludePaths": [
	        {
	            "Type": "File",
	            "Value": "/data/nginx/log/*/*/exclude.log"
	        },
	        {
	            "Type": "Path",
	            "Value": "/data/nginx/log/*/exclude/"
	        }
	    ],
	    "LogType": "minimalist_log",
	    "ExtractRule": {
	        "FilterKeyRegex": [
	            {
	                "Key": "__content__",
	                "Regex": ".*ERROR.*"
	            }
	        ]
	    },
	    "LogSample": "2018-05-22 15:35:53.850 INFO XXXX",
	    "UserDefineRule": {
	        "ParsePathRule": {
	            "PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
	            "Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
	            "Keys": [
	                "instance-id",
	                "pod-name"
	            ]
	        },
	        "ShardHashKey": {
	            "HashKey": "3C"
	        },
	        "EnableRawLog": true,
	        "RawLogKey": "raw",
	        "Fields": {
	            "ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
	        },
	        "Advanced": {
	            "CloseInactive": 10
	        }
	    }
	}`

	jsonFormatRules["多行全文模式"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"Paths": [
			"/data/nginx/log/*/*/*.log"
		],
		"ExcludePaths": [
			{
				"Type": "File",
				"Value": "/data/nginx/log/*/*/exclude.log"
			},
			{
				"Type": "Path",
				"Value": "/data/nginx/log/*/exclude/"
			}
		],
		"LogType": "multiline_log",
		"ExtractRule": {
			"BeginRegex": "\\[([^]]+)].*",
			"FilterKeyRegex": [
				{
					"Key": "__content__",
					"Regex": ".*ERROR.*"
				}
			]
		},
		"LogSample": "[2018-10-01T10:30:01,000] [INFO] java.lang.Exception: exception happened\n    at TestPrintStackTrace.f(TestPrintStackTrace.java:3)\n    at TestPrintStackTrace.g(TestPrintStackTrace.java:7)\n    at TestPrintStackTrace.main(TestPrintStackTrace.java:16)",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		}
	}`

	jsonFormatRules["分隔符模式"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"Paths": [
			"/data/nginx/log/*/*/*.log"
		],
		"ExcludePaths": [
			{
				"Type": "File",
				"Value": "/data/nginx/log/*/*/exclude.log"
			},
			{
				"Type": "Path",
				"Value": "/data/nginx/log/*/exclude/"
			}
		],
		"LogType": "delimiter_log",
		"ExtractRule": {
			"Delimiter": "&&",
			"Keys": [
				"time",
				"ip",
				"status",
				"",
				"user-agent"
			],
			"TimeKey": "time",
			"TimeFormat": "%d/%b/%Y %H:%M:%S",
			"FilterKeyRegex": [
				{
					"Key": "status",
					"Regex": "200"
				}
			],
			"UnMatchUpLoadSwitch": true,
			"UnMatchLogKey": "LogParseFailed"
		},
		"LogSample": "05/May/2016 13:30:28&&192.168.1.1&&200&&18204&&sdk-go",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		}
	}`

	jsonFormatRules["JSON模式"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"Paths": [
			"/data/nginx/log/*/*/*.log"
		],
		"ExcludePaths": [
			{
				"Type": "File",
				"Value": "/data/nginx/log/*/*/exclude.log"
			},
			{
				"Type": "Path",
				"Value": "/data/nginx/log/*/exclude/"
			}
		],
		"LogType": "json_log",
		"ExtractRule": {
			"TimeKey": "time",
			"TimeFormat": "%d/%b/%Y %H:%M:%S",
			"FilterKeyRegex": [
				{
					"Key": "status",
					"Regex": "200"
				}
			],
			"UnMatchUpLoadSwitch": true,
			"UnMatchLogKey": "LogParseFailed"
		},
		"LogSample": "05/May/2016 13:30:28&&192.168.1.1&&200&&18204&&sdk-go",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		}
	}`

	jsonFormatRules["单行完整正则模式"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"Paths": [
			"/data/nginx/log/*/*/*.log"
		],
		"ExcludePaths": [
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
			"LogRegex": "([\\s\\S]+),\\s(\\S+),\\s(\\d+),\\s(\\d+),\\s(\\S+).*",
			"Keys": [
				"time",
				"ip",
				"status",
				"size",
				"user-agent"
			],
			"TimeKey": "time",
			"TimeFormat": "%d/%b/%Y %H:%M:%S",
			"FilterKeyRegex": [
				{
					"Key": "status",
					"Regex": "200"
				}
			],
			"UnMatchUpLoadSwitch": true,
			"UnMatchLogKey": "LogParseFailed"
		},
		"LogSample": "05/May/2016 13:30:28, 192.168.1.1, 200, 18204, sdk-go",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		}
	}`

	jsonFormatRules["多行完整正则模式"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"Paths": [
			"/data/nginx/log/*/*/*.log"
		],
		"ExcludePaths": [
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
			"LogRegex": "\\[([^]]+)]\\s\\[(\\w+)]\\s([\\S\\s]+).*",
			"BeginRegex": "\\[([^]]+)].*",
			"Keys": [
				"time",
				"level",
				"message"
			],
			"TimeKey": "time",
			"TimeFormat": "%Y-%b-dT%H:%M:%S,%f",
			"FilterKeyRegex": [
				{
					"Key": "level",
					"Regex": ".*ERROR.*"
				}
			],
			"UnMatchUpLoadSwitch": true,
			"UnMatchLogKey": "LogParseFailed"
		},
		"LogSample": "[2018-10-01T10:30:01,000] [INFO] java.lang.Exception: exception happened\n    at TestPrintStackTrace.f(TestPrintStackTrace.java:3)\n    at TestPrintStackTrace.g(TestPrintStackTrace.java:7)\n    at TestPrintStackTrace.main(TestPrintStackTrace.java:16)",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		}
	}`

	jsonFormatRules["Nginx模式"] = `{
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

	jsonFormatRules["采集物理机日志"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"Paths": [
			"/data/nginx/log/*/*/*.log"
		],
		"ExcludePaths": [
			{
				"Type": "File",
				"Value": "/data/nginx/log/*/*/exclude.log"
			},
			{
				"Type": "Path",
				"Value": "/data/nginx/log/*/exclude/"
			}
		],
		"LogType": "minimalist_log",
		"ExtractRule": {
			"FilterKeyRegex": [
				{
					"Key": "__content__",
					"Regex": ".*ERROR.*"
				}
			]
		},
		"LogSample": "2018-05-22 15:35:53.850 INFO XXXX",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		}
	}`

	jsonFormatRules["标准容器输出"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"LogType": "minimalist_log",
		"ExtractRule": {
			"FilterKeyRegex": [
				{
					"Key": "__content__",
					"Regex": ".*ERROR.*"
				}
			]
		},
		"LogSample": "2018-05-22 15:35:53.850 INFO XXXX",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		},
		"InputType": 1,
		"ContainerRule": {
			"Stream": "all",
			"ContainerNameRegex": ".*test.*",
			"IncludeContainerLabelRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"ExcludeContainerLabelRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"IncludeContainerEnvRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"ExcludeContainerEnvRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"EnvTag": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"KubernetesRule": {
				"NamespaceNameRegex": ".*test.*",
				"WorkloadType": "Deployment",
				"WorkloadNameRegex": ".*test.*",
				"IncludePodLabelRegex": {
					"Key1": "Value1",
					"Key2": "Value2"
				},
				"ExcludePodLabelRegex": {
					"Key1": "Value1",
					"Key2": "Value2"
				},
				"PodNameRegex": ".*test.*",
				"LabelTag": {
					"Key1": "Value1",
					"Key2": "Value2"
				}
			}
		}
	}`

	jsonFormatRules["容器文本日志"] = `{
		"TopicId": "706f****-****-****-****-****46aa1f07",
		"RuleName": "testname",
		"Paths": [
			"/data/nginx/log/*/*/*.log"
		],
		"ExcludePaths": [
			{
				"Type": "File",
				"Value": "/data/nginx/log/*/*/exclude.log"
			},
			{
				"Type": "Path",
				"Value": "/data/nginx/log/*/exclude/"
			}
		],
		"LogType": "minimalist_log",
		"ExtractRule": {
			"FilterKeyRegex": [
				{
					"Key": "__content__",
					"Regex": ".*ERROR.*"
				}
			]
		},
		"LogSample": "2018-05-22 15:35:53.850 INFO XXXX",
		"UserDefineRule": {
			"ParsePathRule": {
				"PathSample": "/data/nginx/log/dabaad5f-7a10/tls/app.log",
				"Regex": "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log",
				"Keys": [
					"instance-id",
					"pod-name"
				]
			},
			"ShardHashKey": {
				"HashKey": "3C"
			},
			"EnableRawLog": true,
			"RawLogKey": "raw",
			"Fields": {
				"ClusterID": "dabaad5f-7a10-4771-b3ea-d821f73e****"
			},
			"Advanced": {
				"CloseInactive": 10
			}
		},
		"InputType": 2,
		"ContainerRule": {
			"ContainerNameRegex": ".*test.*",
			"IncludeContainerLabelRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"ExcludeContainerLabelRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"IncludeContainerEnvRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"ExcludeContainerEnvRegex": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"EnvTag": {
				"Key1": "Value1",
				"Key2": "Value2"
			},
			"KubernetesRule": {
				"NamespaceNameRegex": ".*test.*",
				"WorkloadType": "Deployment",
				"WorkloadNameRegex": ".*test.*",
				"IncludePodLabelRegex": {
					"Key1": "Value1",
					"Key2": "Value2"
				},
				"ExcludePodLabelRegex": {
					"Key1": "Value1",
					"Key2": "Value2"
				},
				"PodNameRegex": ".*test.*",
				"LabelTag": {
					"Key1": "Value1",
					"Key2": "Value2"
				}
			}
		}
	}`

	jsonFormatRules["启用插件采集"] = `{
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
    	"LogType": "minimalist_log",
    	"ExtractRule": {
        	"TimeKey": "time",
        	"TimeFormat": "%d/%b/%Y %H:%M:%S",
        	"FilterKeyRegex": [
            	{
					"Key": "status",
                	"Regex": "200"
            	}
        	],
        	"UnMatchUpLoadSwitch": true,
        	"UnMatchLogKey": "LogParseFailed"
    	},
    	"LogSample": "2018-05-22 15:35:53.850 INFO XXXX",
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
        	},
        	"Plugin": {
            	"processors":[
                	{
                    	"json":{
                        	"field":"__content__",
                        	"trim_keys":{
                            	"mode":"all",
                            	"chars":"#"
                        	},
                        	"trim_values":{
                            	"mode":"all",
                            	"chars":"#"
                        	},
                        	"allow_overwrite_keys":true,
                        	"allow_empty_values":true
                    	}
                	}
            	]
        	}
    	}
	}`

	return jsonFormatRules
}

func getCreateRuleRequestsMap(jsonFormatRulesMap map[string]string, topicID string) (map[string]*CreateRuleRequest, error) {
	createRuleRequestsMap := make(map[string]*CreateRuleRequest)

	for key, value := range jsonFormatRulesMap {
		req := &CreateRuleRequest{}
		if err := json.Unmarshal([]byte(value), req); err != nil {
			return nil, err
		}
		req.TopicID = topicID
		req.RuleName = uuid.New().String()
		createRuleRequestsMap[key] = req
	}

	return createRuleRequestsMap, nil
}

func createRules(cli Client, createRuleRequestsMap map[string]*CreateRuleRequest, specificRule string) (map[string]string, error) {
	ruleMap := make(map[string]string)

	if len(specificRule) > 0 {
		resp, err := cli.CreateRule(createRuleRequestsMap[specificRule])
		if err != nil {
			return ruleMap, err
		}
		ruleMap[specificRule] = resp.RuleID

		return ruleMap, nil
	}

	for key, req := range createRuleRequestsMap {
		resp, err := cli.CreateRule(req)
		if err != nil {
			return ruleMap, err
		}
		ruleMap[key] = resp.RuleID
	}

	return ruleMap, nil
}

func (suite *SDKRuleTestSuite) TestCreateRuleNormally() {
	for _, req := range suite.createRuleRequestsMap {
		resp, err := suite.cli.CreateRule(req)
		suite.NoError(err)
		if resp != nil {
			suite.ruleList = append(suite.ruleList, resp.RuleID)
		}
	}
}

func (suite *SDKRuleTestSuite) TestCreateRuleAbnormally() {
	testcases := map[*CreateRuleRequest]*Error{
		{
			TopicID:  suite.topic,
			RuleName: uuid.New().String(),
		}: {
			HTTPCode: http.StatusBadRequest,
			Code:     "InvalidArgument",
			Message:  "Invalid argument key Paths, value %!s(<nil>), please check argument.",
		},
	}

	for req, expectedErr := range testcases {
		resp, err := suite.cli.CreateRule(req)
		suite.validateError(err, expectedErr)
		if resp != nil {
			suite.ruleList = append(suite.ruleList, resp.RuleID)
		}
	}
}

func (suite *SDKRuleTestSuite) TestDeleteRuleAbnormally() {
	ruleID := uuid.New().String()
	_, err := suite.cli.DeleteRule(&DeleteRuleRequest{RuleID: ruleID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "RuleNotExist",
		Message:  "Rule " + ruleID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKRuleTestSuite) TestModifyRuleNormally() {
	specificRule := suite.ruleTypeList[0]
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, specificRule)
	suite.NoError(err)
	for _, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)
	}

	ruleID := ruleMap[specificRule]
	_, err = suite.cli.ModifyRule(&ModifyRuleRequest{
		RuleID:   ruleID,
		RuleName: StrPtr("modified-rule-name"),
	})
	suite.NoError(err)
	resp, err := suite.cli.DescribeRule(&DescribeRuleRequest{RuleID: ruleID})
	suite.Equal("modified-rule-name", resp.RuleInfo.RuleName)
}

func (suite *SDKRuleTestSuite) TestModifyRuleAbnormally() {
	specificRule := suite.ruleTypeList[0]
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, specificRule)
	suite.NoError(err)
	for _, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)
	}

	ruleID := ruleMap[specificRule]
	paths := make([]string, 20)
	for i := 0; i < 20; i++ {
		paths[i] = "/data/log" + strconv.Itoa(i+1) + "/*/*/*.log"
	}
	_, err = suite.cli.ModifyRule(&ModifyRuleRequest{
		RuleID: ruleID,
		Paths:  &paths,
	})
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "PathQuotaExceed",
		Message:  "Exceeded Path Quota",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKRuleTestSuite) TestDescribeRuleNormally() {
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, "")
	suite.NoError(err)
	for key, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)

		createRuleReq := suite.createRuleRequestsMap[key]
		resp, err := suite.cli.DescribeRule(&DescribeRuleRequest{RuleID: ruleID})
		suite.NoError(err)
		suite.Equal(suite.project, resp.ProjectID)
		suite.Equal(suite.topic, resp.TopicID)
		suite.Equal(createRuleReq.RuleName, resp.RuleInfo.RuleName)
	}
}

func (suite *SDKRuleTestSuite) TestDescribeRuleAbnormally() {
	ruleID := uuid.New().String()
	_, err := suite.cli.DescribeRule(&DescribeRuleRequest{RuleID: ruleID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "RuleNotExist",
		Message:  "Rule " + ruleID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKRuleTestSuite) TestDescribeRulesNormally() {
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, "")
	suite.NoError(err)
	for _, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)
	}

	resp, err := suite.cli.DescribeRules(&DescribeRulesRequest{
		ProjectID: suite.project,
	})
	suite.NoError(err)
	suite.Equal(len(ruleMap), int(resp.Total))
}

func (suite *SDKRuleTestSuite) TestDescribeRulesAbnormally() {
	projectID := uuid.New().String()
	_, err := suite.cli.DescribeRules(&DescribeRulesRequest{ProjectID: projectID})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "ProjectNotExists",
		Message:  "Project does not exist.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKRuleTestSuite) TestApplyRuleToHostGroupsNormally() {
	specificRule := suite.ruleTypeList[0]
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, specificRule)
	suite.NoError(err)
	for _, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)
	}

	createHostGroupResp, err := suite.cli.CreateHostGroup(&CreateHostGroupRequest{
		HostGroupName: "go-sdk-host-group",
		HostGroupType: "IP",
		HostIPList:    &[]string{"192.168.1.1", "192.168.1.2"},
	})
	suite.NoError(err)

	ruleID := ruleMap[specificRule]
	hostGroupID := createHostGroupResp.HostGroupID
	_, err = suite.cli.ApplyRuleToHostGroups(&ApplyRuleToHostGroupsRequest{
		RuleID:       ruleID,
		HostGroupIDs: []string{hostGroupID},
	})
	suite.NoError(err)

	_, err = suite.cli.DeleteHostGroup(&DeleteHostGroupRequest{HostGroupID: hostGroupID})
	suite.NoError(err)
}

func (suite *SDKRuleTestSuite) TestApplyRuleToHostGroupsAbnormally() {
	specificRule := suite.ruleTypeList[0]
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, specificRule)
	suite.NoError(err)
	for _, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)
	}

	ruleID := ruleMap[specificRule]
	hostGroupID := uuid.New().String()
	_, err = suite.cli.ApplyRuleToHostGroups(&ApplyRuleToHostGroupsRequest{
		RuleID:       ruleID,
		HostGroupIDs: []string{hostGroupID},
	})
	expectedErr := &Error{
		HTTPCode: http.StatusNotFound,
		Code:     "HostGroupNotExist",
		Message:  "HostGroup " + hostGroupID + " does not exist",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKRuleTestSuite) TestDeleteRuleFromHostGroupsNormally() {
	specificRule := suite.ruleTypeList[0]
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, specificRule)
	suite.NoError(err)
	for _, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)
	}

	createHostGroupResp, err := suite.cli.CreateHostGroup(&CreateHostGroupRequest{
		HostGroupName: "go-sdk-host-group",
		HostGroupType: "IP",
		HostIPList:    &[]string{"192.168.1.1", "192.168.1.2"},
	})
	suite.NoError(err)

	ruleID := ruleMap[specificRule]
	hostGroupID := createHostGroupResp.HostGroupID
	_, err = suite.cli.ApplyRuleToHostGroups(&ApplyRuleToHostGroupsRequest{
		RuleID:       ruleID,
		HostGroupIDs: []string{hostGroupID},
	})
	suite.NoError(err)

	_, err = suite.cli.DeleteRuleFromHostGroups(&DeleteRuleFromHostGroupsRequest{
		RuleID:       ruleID,
		HostGroupIDs: []string{hostGroupID},
	})
	suite.NoError(err)

	_, err = suite.cli.DeleteHostGroup(&DeleteHostGroupRequest{HostGroupID: hostGroupID})
	suite.NoError(err)
}

func (suite *SDKRuleTestSuite) TestDeleteRuleFromHostGroupsAbnormally() {
	specificRule := suite.ruleTypeList[0]
	ruleMap, err := createRules(suite.cli, suite.createRuleRequestsMap, specificRule)
	suite.NoError(err)
	for _, ruleID := range ruleMap {
		suite.ruleList = append(suite.ruleList, ruleID)
	}

	createHostGroupResp, err := suite.cli.CreateHostGroup(&CreateHostGroupRequest{
		HostGroupName: "go-sdk-host-group",
		HostGroupType: "IP",
		HostIPList:    &[]string{"192.168.1.1", "192.168.1.2"},
	})
	suite.NoError(err)

	ruleID := ruleMap[specificRule]
	hostGroupID := createHostGroupResp.HostGroupID

	_, rErr := suite.cli.DeleteRuleFromHostGroups(&DeleteRuleFromHostGroupsRequest{
		RuleID:       ruleID,
		HostGroupIDs: []string{hostGroupID},
	})

	_, err = suite.cli.DeleteHostGroup(&DeleteHostGroupRequest{HostGroupID: hostGroupID})
	suite.NoError(err)

	expectedErr := &Error{
		HTTPCode: http.StatusConflict,
		Code:     "RuleAlreadyUnbound",
		Message:  "Rule " + ruleID + " Already Unbound",
	}
	suite.validateError(rErr, expectedErr)
}
