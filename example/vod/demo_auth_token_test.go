package vod

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/model"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func TestVod_GetPlayAuthToken(t *testing.T) {
	vid := "your vid"
	tokenExpireTime := 600 // Token Expire Duration（s）
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodGetPlayInfoRequest{
		Vid:             vid,
		Format:          "mp4",
		Definition:      "360p",
		FileType:        "video",
		LogoType:        "",
		Ssl:             "1",
		NeedThumbs:      "0",
		NeedBarrageMask: "0",
		CdnType:         "0",
	}
	newToken, _ := instance.GetPlayAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}

func TestVod_GetSha1HlsDrmAuthToken(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	expireDuration := int64(6000000) //change to your expire duration (s), no default duration
	token, _ := instance.CreateSha1HlsDrmAuthToken(expireDuration)
	fmt.Println(token)
}

func TestVod_GetPrivateDrmAuthToken(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodGetPrivateDrmPlayAuthRequest{
		Vid:         "your vid",
		DrmType:     "your drmType",
		PlayAuthIds: "your playAuthIds",
		UnionInfo:   "your unionInfo",
	}
	tokenExpireTime := 6000000 // Token Expire Duration（s）
	newToken, _ := instance.GetPrivateDrmAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}

func TestVod_GetSubtitleAuthToken(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})
	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")

	// Media Info
	query := &request.VodGetSubtitleInfoListRequest{
		Vid: "your search vid",
	}
	expireSeconds := 100 // expire time duration: (s)
	token, err := instance.GetSubtitleAuthToken(query, expireSeconds)
	fmt.Println("token ===> ", token)
	fmt.Println("err ===> ", err)
}

func TestVod_GetUploadAuthWithExpiredTime(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	opts := make([]model.UploadAuthOpt, 0)
	// 使用 vod.WithUploadKeyPtn("表达式") 来限制上传的FileName路径
	//     如: "test/*" 表示上传的文件必须包含 "test/" 前缀
	// opts = append(opts, vod.WithUploadKeyPtn("表达式"))

	// 使用 vod.WithUploadSpaceNames([]string{}) 来限制允许上传的空间
	// opts = append(opts, vod.WithUploadSpaceNames([]string{}))

	// 使用 vod.WithUploadPolicy()来设置上传策略
	//opts = append(opts, vod.WithUploadPolicy(&model.UploadPolicy{}))

	// 默认过期时间为1小时
	ret, _ := instance.GetUploadAuth(opts...)
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))

	ret2, _ := instance.GetUploadAuthWithExpiredTime(3*time.Hour, opts...)
	b2, _ := json.Marshal(ret2)
	fmt.Println(string(b2))
}
