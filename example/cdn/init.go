package cdn

import (
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"time"
)

var (
	DefaultInstance = cdn.DefaultInstance
	ak              = "ak"
	sk              = "sk"
	operateDomain   = "operate.com"
	now             = time.Now()
	testStartTime   = now.Unix()
	testEndTime     = now.Add(time.Minute * 10).Unix()
	exampleDomain   = "example.com"
)

func init() {
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
}
