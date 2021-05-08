package upload

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_GetUploadAuthWithExpiredTime(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetAccessKey("")
	instance.SetSecretKey("")
	ret, _ := instance.GetUploadAuth()
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))

	ret2, _ := instance.GetUploadAuthWithExpiredTime(3 * time.Hour)
	b2, _ := json.Marshal(ret2)
	fmt.Println(string(b2))
}
