package tls

const (
	testAk           = "" //替换成自己的ak
	testSk           = "" //替换成自己的sk
	testEndPoint     = "" //替换成tls endPoint,带协议头
	testRegion       = "" //替换成要请求的region，region和endPoint的对应关系 参考 https://www.volcengine.com/docs/6470/73641
	testSessionToken = ""
	testPrefix       = "sdk-test-"
)

var (
	testProjectId string
	testTopicId   string
)
