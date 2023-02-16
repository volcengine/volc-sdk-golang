package mcdn

import (
	"github.com/volcengine/volc-sdk-golang/service/mcdn"
	"time"
)

var (
	DefaultInstance = mcdn.DefaultInstance
	ak              = "ak"
	sk              = "sk"
	now             = time.Now()
	testStartTime   = now.Unix()
	testEndTime     = now.Add(time.Minute * 10).Unix()
)

func init() {
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
}
