package innerlogger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log/level"
	"github.com/stretchr/testify/suite"
)

const UnitTestStdout = "./tls_innerlogger_test_stdout.log"
const UnitTestLogFile = "./tls_innerlogger_test_file.log"

type InnerLoggerTestSuite struct {
	suite.Suite

	debugLogger TLSSDKLogger
	errorLogger TLSSDKLogger
	jsonLogger  TLSSDKLogger
	fileLogger  TLSSDKLogger
}

func (suite *InnerLoggerTestSuite) SetupTest() {
	os.Remove(UnitTestLogFile)
	os.Remove(UnitTestStdout)
	var realStdout = os.Stdout
	os.Stdout, _ = os.OpenFile(UnitTestStdout, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	defer func() {
		os.Stdout = realStdout
	}()

	suite.debugLogger = NewLogger(&LoggerConfig{
		LogLevel: "debug",
	})
	suite.errorLogger = NewLogger(&LoggerConfig{
		LogLevel: "error",
	})
	suite.jsonLogger = NewLogger(&LoggerConfig{
		LogLevel:   "info",
		IsJsonType: true,
	})
	suite.fileLogger = NewLogger(&LoggerConfig{
		LogLevel:      "info",
		LogFileName:   UnitTestLogFile,
		LogMaxSize:    10,
		LogMaxBackups: 10,
	})
}

func TestInnerLoggerTestSuite(t *testing.T) {
	defer func() {
		os.Remove(UnitTestLogFile)
		os.Remove(UnitTestStdout)
	}()
	suite.Run(t, new(InnerLoggerTestSuite))
}

type TestStruct struct {
	A int
	B int
}

type TestCase struct {
	Case       interface{}
	WantLogfmt interface{}
	WantJson   interface{}
}

var testCases = []TestCase{
	{
		Case:       "Hello",
		WantLogfmt: "Hello",
		WantJson:   "Hello",
	},
	{
		Case:       123,
		WantLogfmt: "123",
		WantJson:   "123",
	},
	{
		Case:       1.5,
		WantLogfmt: "1.5",
		WantJson:   "1.5",
	},
	{
		Case: &TestStruct{
			A: 1,
			B: 2,
		},
		WantLogfmt: `"&{A:1 B:2}"`,
		WantJson:   `&{A:1 B:2}`,
	},
	{
		Case: TestStruct{
			A: 1,
			B: 2,
		},
		WantLogfmt: `"{A:1 B:2}"`,
		WantJson:   `{A:1 B:2}`,
	},
	{
		Case:       true,
		WantLogfmt: "true",
		WantJson:   `true`,
	},
	{
		Case:       []string{"H", "E", "L", "L", "O"},
		WantLogfmt: `"[H E L L O]"`,
		WantJson:   `[H E L L O]`,
	},
	{
		Case: map[string]string{
			"Hello": "World",
		},
		WantLogfmt: "map[Hello:World]",
		WantJson:   `map[Hello:World]`,
	},
	{
		Case:       nil,
		WantLogfmt: "<nil>",
		WantJson:   `<nil>`,
	},
	{
		Case:       "\x48",
		WantLogfmt: "H",
		WantJson:   `H`,
	},
}

func (suite *InnerLoggerTestSuite) TestPrintLogToStdout() {
	clearFileContent(UnitTestStdout)
	for _, v := range testCases {
		level.Error(suite.errorLogger).Log("msg", fmt.Sprintf("%+v", v.Case))
	}

	var stdoutBytes []byte
	stdoutBytes, _ = ioutil.ReadFile(UnitTestStdout)
	if len(stdoutBytes) > 0 && stdoutBytes[len(stdoutBytes)-1] == '\n' {
		stdoutBytes = stdoutBytes[0 : len(stdoutBytes)-1]
	}
	var logs = strings.Split(string(stdoutBytes), "\n")

	suite.Equal(len(testCases), len(logs))
	for idx, _ := range logs {
		suite.Equal(testCases[idx].WantLogfmt, GetMsgInLogfmt(logs[idx]))
	}
}

func (suite *InnerLoggerTestSuite) TestPrintLogToFile() {
	clearFileContent(UnitTestLogFile)
	for _, v := range testCases {
		level.Error(suite.fileLogger).Log("msg", fmt.Sprintf("%+v", v.Case))
	}

	var stdoutBytes []byte
	{
		stdoutBytes, _ = ioutil.ReadFile(UnitTestLogFile)
		if len(stdoutBytes) > 0 && stdoutBytes[len(stdoutBytes)-1] == '\n' {
			stdoutBytes = stdoutBytes[0 : len(stdoutBytes)-1]
		}
	}
	var logs = strings.Split(string(stdoutBytes), "\n")

	suite.Equal(len(testCases), len(logs))
	for idx, _ := range logs {
		suite.Equal(testCases[idx].WantLogfmt, GetMsgInLogfmt(logs[idx]))
	}
}

func (suite *InnerLoggerTestSuite) TestDebugLogWithErrorLogger() {
	clearFileContent(UnitTestStdout)
	var (
		stdoutBytes []byte
		err         error
	)
	for _, v := range testCases {
		level.Debug(suite.errorLogger).Log("msg", fmt.Sprintf("%+v", v.Case))
	}
	stdoutBytes, err = ioutil.ReadFile(UnitTestLogFile)
	suite.Equal((os.IsNotExist(err) || (len(stdoutBytes) == 0)), true)
}

func (suite *InnerLoggerTestSuite) TestPrintLogInJson() {
	clearFileContent(UnitTestStdout)
	for _, v := range testCases {
		level.Error(suite.jsonLogger).Log("msg", fmt.Sprintf("%+v", v.Case))
	}

	var stdoutBytes []byte
	stdoutBytes, _ = ioutil.ReadFile(UnitTestStdout)
	if len(stdoutBytes) > 0 && stdoutBytes[len(stdoutBytes)-1] == '\n' {
		stdoutBytes = stdoutBytes[0 : len(stdoutBytes)-1]
	}
	var logs = strings.Split(string(stdoutBytes), "\n")

	suite.Equal(len(testCases), len(logs))
	for idx, _ := range logs {
		suite.Equal(testCases[idx].WantJson, GetMsgInJson(logs[idx]))
	}
}

func GetMsgInLogfmt(logContent string) string {
	var pos = strings.Index(logContent, "msg=")
	if pos == -1 {
		return ""
	}
	return logContent[pos+4:]
}

func GetMsgInJson(logContent string) string {
	var v map[string]interface{}
	json.Unmarshal([]byte(logContent), &v)
	i, ok := v["msg"]
	if !ok {
		return ""
	}
	s, ok := i.(string)
	if !ok {
		return ""
	}
	return s
}

func clearFileContent(filePath string) {
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0755)
	defer file.Close()
}
