package tls

import "os"

var (
	// 出于安全考虑，建议您将账号和服务信息（包括Endpoint、Region、AccessKeyID、AccessKeySecret、SecurityToken）配置在环境变量中
	testEndPoint     = os.Getenv("VOLCENGINE_ENDPOINT")
	testRegion       = os.Getenv("VOLCENGINE_REGION")
	testAk           = os.Getenv("VOLCENGINE_ACCESS_KEY_ID")
	testSk           = os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET")
	testSessionToken = os.Getenv("VOLCENGINE_TOKEN")

	testPrefix = "sdk-test-"
)
