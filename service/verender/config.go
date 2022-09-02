package verender

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion  = "cn-north-1"
	ServiceVersion = "2021-12-31"
	ServiceName    = "verender"
	StatusOk       = 0
	MINPARTSIZE    = 5 << 20
	MAXPARTSIZE    = 5 << 30
	MAXOBJECTSIZE  = MAXPARTS * MAXPARTSIZE
	MAXPARTS       = 10000
)

var AllowedFileTypes = map[string]struct{}{"asset": {}, "scene_file": {}, "client_script": {}}
var OrderTypeMap = map[string]string{"descend": "1", "ascend": "2"}

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{AccessKeyID: "", SecretAccessKey: "", Service: "verender", Region: DefaultRegion},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"ListWorkspaces": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListWorkspaces"},
				"Version": []string{ServiceVersion},
			},
			Timeout: 15 * time.Second,
		},
		"PurchaseWorkspace": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PurchaseWorkspace"},
				"Version": []string{ServiceVersion},
			},
		},
		"UpdateWorkspace": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateWorkspace"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteWorkspace": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteWorkspace"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListResourcePools": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListResourcePools"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetWorkspaceLimit": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWorkspaceLimit"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetHardwareSpecifications": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWorkspaceHardwareSpecifications"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListJobs": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRenderJobs"},
				"Version": []string{ServiceVersion},
			},
			Timeout: 15 * time.Second,
		},
		"GetRenderJob": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"CreateJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"EditJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EditRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetJobPriority": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRenderJobPriority"},
				"Version": []string{ServiceVersion},
			},
		},
		"StartJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"PauseJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PauseRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"StopJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"FullSpeedJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FullSpeedRenderJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"RetryJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RetryJob"},
				"Version": []string{ServiceVersion},
			},
		},
		"PauseJobs": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PauseJobs"},
				"Version": []string{ServiceVersion},
			},
		},
		"ResumeJobs": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeJobs"},
				"Version": []string{ServiceVersion},
			},
		},
		"StopJobs": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopJobs"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteJobs": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteJobs"},
				"Version": []string{ServiceVersion},
			},
		},
		"FullSpeedJobs": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FullSpeedRenderJobs"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetLayerFrames": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetLayerFrames"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetAccountStatistics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAccountStatistics"},
				"Version": []string{ServiceVersion},
			},
			Timeout: 15 * time.Second,
		},
		"GetAccountStatisticDetails": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAccountStatisticDetails"},
				"Version": []string{ServiceVersion},
			},
			Timeout: 15 * time.Second,
		},
		"DownloadStatisticDetails": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DownloadStatisticDetails"},
				"Version": []string{ServiceVersion},
			},
			Timeout: 15 * time.Second,
		},
		"GetCurrentUser": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCurrentUser"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetJobOverallStatistics": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetJobOverallStatistics"},
				"Version": []string{ServiceVersion},
			},
			Timeout: 15 * time.Second,
		},
		"GetStorageAccess": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetStorageAccess"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListMyMessages": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListMyMessages"},
				"Version": []string{ServiceVersion},
			},
		},
		"MarkMessageAsRead": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MarkMessageAsRead"},
				"Version": []string{ServiceVersion},
			},
		},
		"BatchMarkMessagesAsRead": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchMarkMessagesAsRead"},
				"Version": []string{ServiceVersion},
			},
		},
		"MarkAllMessagesAsRead": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MarkAllMessagesAsRead"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteMessage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteMessage"},
				"Version": []string{ServiceVersion},
			},
		},
		"BatchDeleteMessages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchDeleteMessages"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteAllMessages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteAllMessages"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteAllReadMessages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteAllReadMessages"},
				"Version": []string{ServiceVersion},
			},
		},
	}
)

// Verender
type Verender struct {
	Client *base.Client
}

// NewInstance create an instance
func NewInstance(ak, sk string) *Verender {
	instance := &Verender{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	if ak == "" {
		instance.Client.SetAccessKey(os.Getenv("VOLCENGINE_AK"))
	} else {
		instance.Client.SetAccessKey(ak)
	}
	if sk == "" {
		instance.Client.SetSecretKey(os.Getenv("VOLCENGINE_SK"))
	} else {
		instance.Client.SetSecretKey(sk)
	}

	return instance
}

func GetMinioClient(ak, sk, token, endpoint string) (*minio.Client, error) {
	if ok := strings.HasPrefix(endpoint, "https://"); ok {
		endpoint = strings.Replace(endpoint, "https://", "", 1)
	} else {
		endpoint = strings.Replace(endpoint, "http://", "", 1)
	}

	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(ak, sk, token),
		Secure: true,
	})

	if err != nil {
		return nil, err
	}
	return cli, nil
}

// GetServiceInfo interface
func (v *Verender) GetServiceInfo(env string) *base.ServiceInfo {
	return v.Client.ServiceInfo
}

// GetAPIInfo interface
func (v *Verender) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion
func (v *Verender) SetRegion(env, region string) {
	v.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (v *Verender) SetHost(host string) {
	v.Client.ServiceInfo.Host = host
}

// SetSchema .
func (v *Verender) SetSchema(schema string) {
	v.Client.ServiceInfo.Scheme = schema
}
