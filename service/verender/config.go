package verender

import (
    "fmt"
    "net/http"
    "time"

    "github.com/volcengine/volc-sdk-golang/base"
)

const (
    DefaultIsp                = "ct"
    DefaultRegion             = "cn-north-1"
    ServiceVersion            = "2021-12-31"
    ServiceName               = "verender"
    DefaultTimeoutInSecond    = 5
    DefaultSchema             = "https"
    DefaultHost               = "open.volcengineapi.com"
    CodeOK                    = 0
    HTTPStatusOK              = 200
    MinPartSize               = uint64(5 << 20) // 5MB
    MaxPartSize               = uint64(5 << 30) // 5GB
    MaxParts                  = 10000
    MaxObjectSize             = int64(MaxParts * MaxPartSize)
    DefaultLinuxDownloadDir   = "Download"
    DefaultWindowsDownloadDir = "C:\\Download"
    DefaultListDirMaxDepth    = 50

    DefaultOrderFieldWorkspace = "CreatedAt"
    DefaultOrderFieldStatistic = "StartTime"
    DefaultOrderFieldRenderJob = "created_at"
    DefaultOrderTypeWorkspace  = "ascend"
    DefaultOrderTypeStatistic  = "1"
    DefaultOrderTypeRenderJob  = "1"
    DefaultPageNum             = "1"
    DefaultPageSize            = "10"
)

var ValidOrderTypeWorkspace = map[string]string{
    "asc":  "ascend",
    "desc": "descend",
}

var ValidOrderTypeStatistic = map[string]string{
    "asc":  "1",
    "desc": "2",
}

var ValidOrderTypeRenderJob = map[string]string{
    "asc":  "1",
    "desc": "2",
}

var ValidOrderFieldWorkspace = map[string]bool{
    "CreatedAt": true,
}

var ValidOrderFieldStatistic = map[string]bool{
    "StartTime": true,
    "TotalCost": true,
}

var ValidOrderFieldRenderJob = map[string]bool{
    "created_at": true,
    "priority":   true,
    "title":      true,
}

var (
    ErrInvalidResponseStatus = fmt.Errorf("invalid response status")
    ErrInvalidArgs           = fmt.Errorf("invalid args")
)

var ServiceInfo = &base.ServiceInfo{
    Timeout: DefaultTimeoutInSecond * time.Second,
    Scheme:  DefaultSchema,
    Host:    DefaultHost,
    Header: map[string][]string{
        "Accept": {"application/json"},
    },
    Credentials: base.Credentials{
        AccessKeyID:     "",
        SecretAccessKey: "",
        Service:         ServiceName,
        Region:          DefaultRegion,
        SessionToken:    "",
    },
}

var APIInfoList = map[string]*base.ApiInfo{
    "ListWorkspace": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"ListWorkspaces"},
            "Version": {ServiceVersion},
        },
    },
    "GetStorageAccess": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"GetStorageAccess"},
            "Version": {ServiceVersion},
        },
    },
    "GetCurrentUser": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"GetCurrentUser"},
            "Version": {ServiceVersion},
        },
    },
    "ListCellSpec": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"ListCellSpecs"},
            "Version": {ServiceVersion},
        },
    },
    "ListDccs": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"ListDccs"},
            "Version": {ServiceVersion},
        },
    },
    "ListAccountDccPlugins": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"ListAccountDccPlugins"},
            "Version": {ServiceVersion},
        },
    },
    "AddRenderSetting": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"AddRenderSetting"},
            "Version": {ServiceVersion},
        },
    },
    "DeleteRenderSetting": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"DeleteRenderSetting"},
            "Version": {ServiceVersion},
        },
    },
    "UpdateRenderSetting": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"UpdateRenderSetting"},
            "Version": {ServiceVersion},
        },
    },
    "ListRenderSetting": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"GetRenderSettingList"},
            "Version": {ServiceVersion},
        },
    },
    "GetRenderSetting": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"GetRenderSetting"},
            "Version": {ServiceVersion},
        },
    },
    "CreateRenderJob": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"CreateRenderJob"},
            "Version": {ServiceVersion},
        },
    },
    "ListRenderJob": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"ListRenderJobs"},
            "Version": {ServiceVersion},
        },
    },
    "GetRenderJob": {
        Method: http.MethodGet,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"GetRenderJob"},
            "Version": {ServiceVersion},
        },
    },
    "UpdateRenderJobsPriority": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"UpdateRenderJobsPriority"},
            "Version": {ServiceVersion},
        },
    },
    "RetryRenderJob": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"RetryJob"},
            "Version": {ServiceVersion},
        },
    },
    "ResumeRenderJobs": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"ResumeJobs"},
            "Version": {ServiceVersion},
        },
    },
    "StopRenderJobs": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"StopJobs"},
            "Version": {ServiceVersion},
        },
    },
    "DeleteRenderJobs": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"DeleteJobs"},
            "Version": {ServiceVersion},
        },
    },
    "FullSpeedRenderJobs": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"FullSpeedRenderJobs"},
            "Version": {ServiceVersion},
        },
    },
    "AutoFullSpeedRenderJobs": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"AutoAllRenderJobs"},
            "Version": {ServiceVersion},
        },
    },
    "ListJobOutput": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"ListJobOutput"},
            "Version": {ServiceVersion},
        },
    },
    "GetJobOutput": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"GetJobOutput"},
            "Version": {ServiceVersion},
        },
    },
    "UpdateJobOutput": {
        Method: http.MethodPost,
        Path:   "/",
        Query: map[string][]string{
            "Action":  {"UpdateJobOutput"},
            "Version": {ServiceVersion},
        },
    },
}
