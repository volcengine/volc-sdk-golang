package verender

import (
    "bytes"
    "crypto/md5"
    "crypto/tls"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path"
    "path/filepath"
    "strconv"
    "strings"
    "time"

    "github.com/google/uuid"
)

const (
    IspCt = "ct"
    IspUn = "un"
    IspCm = "cm"

    FtransClientAclToken = "4d0c1b2513cb82263814e10bf2f136ed"

    FtransProtocolTCP = "tcp"
    FtransProtocolUDP = "udp"

    FtransSwitchOn  = 1
    FtransSwitchOff = 0

    FtransClientSignatureExpiredNSec = 15

    FtransMountPath = "/var/mnt"

    FtransOrderTypeAsc    = "asc"
    FtransOrderTypeDesc   = "desc"
    FtransOrderFieldName  = "name"
    FtransOrderFieldMtime = "mtime"

    FtransTaskStatusWaiting  = int8(0)
    FtransTaskStatusDoing    = int8(1)
    FtransTaskStatusFinished = int8(2)
    FtransTaskStatusFailed   = int8(3)

    FtransPartSize        = 4 << 20
    FtransPartConcurrency = 4

    FtransS2UriPath  = "/v2/ftrans"
    FtransS10UriPath = "/s10/p2/ftrans"

    DefaultFtransPageNo   = 1
    DefaultFtransPageSize = 10

    FtransTransStatusCompleted   = "completed"
    FtransTransStatusFailed      = "failed"
    FtransTransStatusCancelled   = "cancelled"
    FtransStatusUploadFileNone   = 0x401003
    FtransStatusUploadDirNone    = 0x401202
    FtransStatusDownloadFileNone = 0x402003
    FtransStatusSucc             = 0
)

type FtransClient struct {
    FtransBucketName     string                 `json:"ftrans_bucket_name"`
    FtransServerName     string                 `json:"ftrans_server_name"`
    FtransAclToken       string                 `json:"ftrans_acl_token"`
    FtransS10Addr        string                 `json:"ftrans_s10_addr"`
    FtransClientAddr     string                 `json:"ftrans_client_addr"`
    FtransS2Addr         string                 `json:"ftrans_s2_addr"`
    FtransProxyAddr      string                 `json:"ftrans_proxy_addr"`
    FtransClientAclToken string                 `json:"ftrans_client_acl_token"`
    FtransClientConfig   map[string]interface{} `json:"ftrans_client_config"`
    FtransS10Conn        *http.Client           `json:"-"`
}

type ProxyRecord struct {
    IP           string `json:"IP"`
    Port         int    `json:"Port"`
    Type         string `json:"Type"`
    State        string `json:"State"`
    TargetDomain string `json:"TargetDomain"`
    TargetPort   int    `json:"TargetPort"`
}

type FtransFileInfo struct {
    Name         string `json:"name"`
    Type         string `json:"type,omitempty"`
    Size         int64  `json:"size"`
    LastModified string `json:"mtime"`
    MTime        int64  `json:"-"` // unit usec
    Md5          string `json:"md5,omitempty"`
}

type FtransFileRequest struct {
    LocalFilePath  string `json:"LocalFilePath"`
    RemoteFilePath string `json:"RemoteFilePath"`
    PartSize       int64  `json:"PartSize"`
}

type FtransListFileRequest struct {
    ServerName    string `json:"serverName,omitempty"`
    ServerAddress string `json:"serverAddress,omitempty"`
    AclToken      string `json:"aclToken,omitempty"`
    Prefix        string `json:"des"`
    FilterIn      string `json:"filterIn"`
    OrderType     string `json:"orderBy"`
    OrderField    string `json:"sortBy"`
    PageNum       int64  `json:"pageNo"`
    PageSize      int64  `json:"pageSize"`
}

type FtransRemoveFileRequest struct {
    ServerName    string `json:"serverName,omitempty"`
    ServerAddress string `json:"serverAddress,omitempty"`
    AclToken      string `json:"aclToken,omitempty"`
    Des           string `json:"des"`
}

type FtransListFileResponse struct {
    Total   int64             `json:"matchedNum"`
    Records []*FtransFileInfo `json:"fileNodes"`
}

type FtransFileReqNode struct {
    TransId              string `json:"transId"`
    AclToken             string `json:"aclToken"`
    Src                  string `json:"src"`
    Des                  string `json:"des"`
    LeftSize             int64  `json:"leftSize"`
    LeftFilesNum         int64  `json:"leftFileNum"`
    LeftDirsNum          int64  `json:"leftDirsNum"`
    TransferringFilesNum int64  `json:"transferringFilesNum"`
    SubmittedFilesNum    int64  `json:"submittedFilesNum"`
    FailedFilesNum       int64  `json:"failedFilesNum"`
    CancelledFilesNum    int64  `json:"cancelledFilesNum"`
    StatusCode           int64  `json:"statusCode"`
    TransStatus          string `json:"transStatus"`
}

type FtransReqNode struct {
    ServerName    string               `json:"serverName"`
    ServerAddress string               `json:"serverAddress"`
    Files         []*FtransFileReqNode `json:"files"`
    Dirs          []*FtransFileReqNode `json:"dirs"`
}

type ftransFileTask struct {
    LocalFilePath  string
    RemoteFilePath string
    TempFilePath   string
    FileSize       int64
    PartSize       int64
    MtimeUSec      int64
    File           *os.File
    PartTasks      []*ftransPartTask
    WaitingParts   int32
    DoingParts     int32
    Status         int8
}

type ftransPartTask struct {
    Offset     int64
    Size       int64
    Status     int8
    Data       []byte
    parentTask *ftransFileTask
    f          *os.File
}

func parseAddressMap(addr string) (map[string]map[string]string, error) {
    addrMap := map[string]map[string]string{}

    for _, elem := range strings.Split(addr, ";") {
        segs := strings.Split(elem, ":")

        // version:ip:port or version:isp:ip:port
        if len(segs) < 3 {
            continue
        }

        version := segs[0]
        if segs[1] != IspCt && segs[1] != IspUn && segs[1] != IspCm {
            host := strings.Join(segs[1:], ":")

            addrMap[version] = map[string]string{
                IspCt: host,
                IspUn: host,
                IspCm: host,
            }
        } else {
            isp := segs[1]
            host := strings.Join(segs[2:], ":")

            if _, ok := addrMap[version]; ok {
                addrMap[version][isp] = host
            } else {
                addrMap[version] = map[string]string{
                    isp: host,
                }
            }
        }
    }

    return addrMap, nil
}

func getFtransClientSig(op string, t int64, token string) string {
    s := fmt.Sprintf("%s@ftrans/%s@%d", token, op, t)
    return fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(s))))
}

func NewFtransClient(bucketName, serverName, aclToken, s10Server, cert, key, clientAddr, s2Server, proxyAddr, isp,
    clientAclToken string) (*FtransClient, error) {
    fc := &FtransClient{
        FtransBucketName:     bucketName,
        FtransServerName:     serverName,
        FtransAclToken:       aclToken,
        FtransS10Addr:        "",
        FtransS10Conn:        nil,
        FtransClientAddr:     "",
        FtransS2Addr:         "",
        FtransProxyAddr:      "",
        FtransClientAclToken: clientAclToken,
        FtransClientConfig:   nil,
    }

    if fc.FtransClientAclToken == "" {
        fc.FtransClientAclToken = FtransClientAclToken
    }

    // use ct as default
    if isp == "" {
        isp = IspCt
    }

    if s10Server != "" {
        s10ServerAddrMap, err := parseAddressMap(s10Server)

        if err != nil {
            return nil, err
        }

        found := false
        for _, ispMap := range s10ServerAddrMap {
            if _, ok := ispMap[isp]; ok {
                fc.FtransS10Addr = ispMap[isp]
                found = true
                break
            }
        }

        if !found {
            return nil, fmt.Errorf("ftrans: isp[%s] not found in s10Server[%s]", isp, s10Server)
        }

        if cert == "" || key == "" {
            return nil, fmt.Errorf("ftrans: cert and key must be set when s10Server specified")
        }

        crt, err := tls.X509KeyPair([]byte(cert), []byte(key))
        if err != nil {
            return nil, fmt.Errorf("ftrans: invalid cert or key, %s", err.Error())
        }

        fc.FtransS10Conn = &http.Client{
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{
                    ServerName:   fc.FtransServerName,
                    Certificates: []tls.Certificate{crt},
                },
            },
        }
    }

    if clientAddr != "" {
        fc.FtransClientAddr = clientAddr

        conf, err := fc.getFtransClientConfig()
        if err != nil {
            return nil, err
        }

        fc.FtransClientConfig = conf
        version := fc.FtransClientConfig["Version"].(string)
        protocol := FtransProtocolUDP
        if int(conf["RunTimeTransTudpSwitch"].(float64)) == FtransSwitchOff &&
            int(conf["RunTimeTransTtcpSwitch"].(float64)) == FtransSwitchOn {
            protocol = FtransProtocolTCP
        }

        if s2Server == "" {
            return nil, fmt.Errorf("ftrans: s2Server must be set when s2Server specified")
        }

        s2ServerMap, err := parseAddressMap(s2Server)
        if err != nil {
            return nil, err
        }

        if _, ok := s2ServerMap[version]; !ok {
            return nil, fmt.Errorf("ftrans: client version[%s] is too old, please upgrade it first", version)
        }

        if _, ok := s2ServerMap[version][isp]; !ok {
            return nil, fmt.Errorf("ftrans: isp [%s] not found in s2Server [%s] of version [%s]",
                isp, s2Server, version)
        }

        fc.FtransS2Addr = s2ServerMap[version][isp]

        if proxyAddr != "" {
            proxyList, err := fc.getProxyList(proxyAddr)
            if err != nil {
                return nil, err
            }

            found := false
            for _, p := range proxyList {
                addr := fmt.Sprintf("%s:%d", p.TargetDomain, p.TargetPort)
                if addr == fc.FtransS2Addr && p.Type == protocol && p.State == "运行中" {
                    fc.FtransProxyAddr = fmt.Sprintf("%s:%d", p.IP, p.Port)
                    found = true
                    break
                }
            }

            if !found {
                targetDomain := strings.Split(fc.FtransS2Addr, ":")[0]
                targetPort, _ := strconv.ParseInt(strings.Split(fc.FtransS2Addr, ":")[1], 10, 64)
                pr := []*ProxyRecord{
                    {
                        Type:         protocol,
                        TargetDomain: targetDomain,
                        TargetPort:   int(targetPort),
                    },
                }
                proxyList, err = fc.createProxy(proxyAddr, pr)
                if err != nil {
                    return nil, fmt.Errorf("ftrans: create proxy for [%s] failed %s", fc.FtransS2Addr, err.Error())
                }

                fc.FtransProxyAddr = fmt.Sprintf("%s:%d", proxyList[0].IP, proxyList[0].Port)
            }
        }
    }

    return fc, nil
}

func (fc *FtransClient) UploadFile(r *FtransFileRequest) (*FtransFileInfo, error) {
    if r.RemoteFilePath == "" {
        return nil, fmt.Errorf("invalid des %s", r.RemoteFilePath)
    }

    r.RemoteFilePath = fc.toSlash(r.RemoteFilePath)

    if fc.enableFtransClient() {
        return fc.uploadFileS2(r)
    } else {
        return fc.uploadFileS10(r)
    }
}

func (fc *FtransClient) DownloadFile(r *FtransFileRequest) (*FtransFileInfo, error) {
    if r.RemoteFilePath == "" {
        return nil, fmt.Errorf("invalid des %s", r.RemoteFilePath)
    }

    r.RemoteFilePath = fc.toSlash(r.RemoteFilePath)
    if fc.enableFtransClient() {
        return fc.downloadFileS2(r)
    } else {
        return fc.downloadFileS10(r)
    }
}

func (fc *FtransClient) ListFile(r *FtransListFileRequest) (*FtransListFileResponse, error) {
    if r.Prefix != "" {
        r.Prefix = fc.toSlash(r.Prefix)
    }
    r.AclToken = fc.FtransAclToken
    r.ServerName = fc.FtransServerName

    if r.OrderType != FtransOrderTypeDesc {
        r.OrderType = FtransOrderTypeAsc
    }

    if r.OrderField != FtransOrderFieldMtime {
        r.OrderField = FtransOrderFieldName
    }

    if r.PageNum <= 0 {
        r.PageNum = DefaultFtransPageNo
    }

    if r.PageSize <= 0 {
        r.PageSize = DefaultFtransPageSize
    }

    if fc.enableFtransClient() {
        return fc.listFileS2(r)
    } else {
        return fc.listFileS10(r)
    }
}

func (fc *FtransClient) RemoveFile(filename string) error {
    if filename == "" {
        return fmt.Errorf("invalid des %s", filename)
    }

    filename = fc.toSlash(filename)
    if fc.enableFtransClient() {
        return fc.removeFileS2(filename)
    } else {
        return fc.removeFileS10(filename)
    }
}

func (fc *FtransClient) StatFile(filename string) (*FtransFileInfo, error) {
    if filename == "" {
        return nil, fmt.Errorf("invalid des %s", filename)
    }

    filename = fc.toSlash(filename)
    if fc.enableFtransClient() {
        return fc.statFileS2(filename)
    } else {
        return fc.statFileS10(filename)
    }
}

func (fc *FtransClient) encodePath(path string) string {
    return base64.URLEncoding.EncodeToString([]byte(path))
}

func (fc *FtransClient) normalizeFilterIn(filterIn string) string {
    specialChars := map[string]string{
        ".": "\\.",
        "*": "\\*",
        "+": "\\+",
        "?": "\\?",
        "[": "\\[",
        "]": "\\]",
        "{": "\\{",
        "}": "\\}",
        "(": "\\(",
        ")": "\\)",
        ",": "\\,",
    }

    filterIn = strings.ReplaceAll(filterIn, "\\", "\\\\")
    for old, new := range specialChars {
        filterIn = strings.ReplaceAll(filterIn, old, new)
    }

    return filterIn
}

func (fc *FtransClient) doRequest(conn *http.Client, method, url string, header http.Header, body []byte) (
    *http.Response, error) {
    if conn == nil {
        conn = &http.Client{}
    }

    req, err := http.NewRequest(method, url, bytes.NewReader(body))
    if err != nil {
        return nil, fmt.Errorf("ftrans:doRequest: new request for [%s] failed [%s]", url, err.Error())
    }

    req.Header = header

    resp, err := conn.Do(req)
    if err != nil {
        return nil, fmt.Errorf("ftrans:doRequest: do http request [%s] failed [%s]", url, err.Error())
    }

    return resp, nil
}

func (fc *FtransClient) getFtransClientConfig() (map[string]interface{}, error) {
    op := "config"
    t := time.Now().Unix() + FtransClientSignatureExpiredNSec
    sig := getFtransClientSig(op, t, fc.FtransClientAclToken)

    url := fmt.Sprintf("http://%s%s?op=%s&t=%d&sig=%s", fc.FtransClientAddr, FtransS2UriPath, op, t, sig)
    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, nil)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:getFtransClientConfig: invalid http status [%d]", resp.StatusCode)
    }

    var conf map[string]interface{}
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:getFtransClientConfig: read response body failed [%s]", err.Error())
    }

    if err := json.Unmarshal(respBody, &conf); err != nil {
        return nil, fmt.Errorf("ftrans:getFtransClientConfig: unmarshal response body failed [%s]", err.Error())
    }

    return conf, nil
}

func (fc *FtransClient) getProxyList(proxyManagerAddr string) ([]*ProxyRecord, error) {
    url := fmt.Sprintf("http://%s/api/v1/proxy-list", proxyManagerAddr)

    resp, err := fc.doRequest(nil, http.MethodGet, url, nil, nil)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != HTTPStatusOK {
        return nil, fmt.Errorf("ftrans:getProxyList: invalid http response status [%d]", resp.StatusCode)
    }

    var proxyList []*ProxyRecord
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:getProxyList: read response body failed [%s]", err.Error())
    }

    if err := json.Unmarshal(respBody, &proxyList); err != nil {
        return nil, fmt.Errorf("ftrans:getProxyList: unmarshal response body failed [%s]", err.Error())
    }

    return proxyList, nil
}

func (fc *FtransClient) createProxy(proxyManagerAddr string, p []*ProxyRecord) ([]*ProxyRecord, error) {
    url := fmt.Sprintf("http://%s/api/v1/proxy", proxyManagerAddr)
    body, _ := json.Marshal(p)

    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != HTTPStatusOK {
        return nil, fmt.Errorf("ftrans:createProxy: invalid http status [%d]", resp.StatusCode)
    }

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:createProxy: read response body failed [%s]", err.Error())
    }

    var pr []*ProxyRecord
    if err := json.Unmarshal(respBody, &pr); err != nil {
        return nil, fmt.Errorf("ftrans:createProxy: unmarshal response body failed [%s]", err.Error())
    }

    return pr, nil
}

func (fc *FtransClient) enableFtransClient() bool {
    return fc.FtransClientAddr != ""
}

func (fc *FtransClient) enableFtransProxy() bool {
    return fc.FtransProxyAddr != ""
}

func (fc *FtransClient) toSlash(filename string) string {
    filename = strings.ReplaceAll(filename, ":\\", "/")
    filename = strings.ReplaceAll(filename, "\\", "/")
    if filename[0] == '/' {
        filename = filename[1:]
    }
    return filename
}

func (fc *FtransClient) getFullPath(filename string) string {
    return fmt.Sprintf("%s/%s/%s", FtransMountPath, fc.FtransBucketName, filename)
}

func (fc *FtransClient) genFtransClientSig(op, clientAclToken string, t int64) string {
    s := fmt.Sprintf("%s@ftrans/%s@%d", clientAclToken, op, t)
    return fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(s))))
}

func (fc *FtransClient) uploadFileS2(r *FtransFileRequest) (*FtransFileInfo, error) {
    op := "upload"
    t := time.Now().Unix() + FtransClientSignatureExpiredNSec
    sig := fc.genFtransClientSig(op, fc.FtransClientAclToken, t)
    url := fmt.Sprintf("http://%s%s?op=%s&t=%d&sig=%s", fc.FtransClientAddr, FtransS2UriPath, op, t, sig)
    body := FtransReqNode{
        ServerName: fc.FtransServerName,
        Files: []*FtransFileReqNode{
            {
                TransId:  uuid.New().String(),
                AclToken: fc.FtransAclToken,
                Src:      r.LocalFilePath,
                Des:      fc.getFullPath(r.RemoteFilePath),
            },
        },
        Dirs: nil,
    }

    if fc.enableFtransProxy() {
        body.ServerAddress = fc.FtransProxyAddr
    } else {
        body.ServerAddress = fc.FtransS2Addr
    }

    data, _ := json.Marshal(body)

    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, data)
    if err != nil {
        return nil, fmt.Errorf("ftrans:uploadFileS2: do request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:uploadFileS2: invalid http status [%d]", resp.StatusCode)
    }

    for {
        sr, err := fc.statusUploadFileS2(&body)
        if err != nil {
            return nil, fmt.Errorf("ftrans:uploadFileS2: status upload failed [%s]", err.Error())
        }

        if sr.StatusCode != FtransStatusSucc && sr.StatusCode != FtransStatusUploadFileNone {
            return nil, fmt.Errorf("ftrans:uploadFileS2: status upload failed [%d]", sr.StatusCode)
        }

        if sr.TransStatus == FtransTransStatusCompleted {
            break
        } else if sr.TransStatus == FtransTransStatusFailed || sr.TransStatus == FtransTransStatusCancelled {
            return nil, fmt.Errorf("ftrans:uploadFileS2: upload file [%s] failed [%s]", r.LocalFilePath, sr.TransStatus)
        }

        time.Sleep(time.Second)
    }

    st, _ := os.Stat(r.LocalFilePath)
    ffi := FtransFileInfo{
        Name:  r.RemoteFilePath,
        Size:  st.Size(),
        MTime: st.ModTime().UnixMicro(),
    }

    return &ffi, nil
}

func (fc *FtransClient) statusUploadFileS2(r *FtransReqNode) (*FtransFileReqNode, error) {
    op := "status_upload"
    t := time.Now().Unix() + FtransClientSignatureExpiredNSec
    sig := fc.genFtransClientSig(op, fc.FtransClientAclToken, t)

    url := fmt.Sprintf("http://%s%s?op=%s&t=%d&sig=%s", fc.FtransClientAddr, FtransS2UriPath, op, t, sig)
    body, _ := json.Marshal(r)

    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statusUploadFileS2: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:statusUploadFileS2: invalid http status [%d]", resp.StatusCode)
    }

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statusUploadFileS2: read response body failed [%s]", err.Error())
    }

    var res FtransReqNode
    if err := json.Unmarshal(respBody, &res); err != nil {
        return nil, fmt.Errorf("ftrans:statusUploadFileS2: unmarshal response body failed [%s]", err.Error())
    }

    return res.Files[0], nil
}

func (fc *FtransClient) statusDownloadS2(r *FtransReqNode) (*FtransFileReqNode, error) {
    op := "status_download"
    t := time.Now().Unix() + FtransClientSignatureExpiredNSec
    sig := fc.genFtransClientSig(op, fc.FtransClientAclToken, t)

    url := fmt.Sprintf("http://%s%s?op=%s&t=%d&sig=%s", fc.FtransClientAddr, FtransS2UriPath, op, t, sig)
    body, _ := json.Marshal(r)

    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statusDownloadS2: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:statusDownloadS2: invalid http status [%d]", resp.StatusCode)
    }

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statusDownloadS2: read response body failed [%s]", err.Error())
    }

    var res FtransReqNode
    if err := json.Unmarshal(respBody, &res); err != nil {
        return nil, fmt.Errorf("ftrans:statusDownloadS2: unmarshal response body failed [%s] in ", err.Error())
    }

    return res.Files[0], nil
}

func (fc *FtransClient) downloadFileS2(r *FtransFileRequest) (*FtransFileInfo, error) {
    op := "download"
    t := time.Now().Unix() + FtransClientSignatureExpiredNSec
    sig := fc.genFtransClientSig(op, fc.FtransClientAclToken, t)

    body := FtransReqNode{
        ServerName: fc.FtransServerName,
        Files: []*FtransFileReqNode{
            {
                TransId:  uuid.New().String(),
                AclToken: fc.FtransAclToken,
                Src:      fc.getFullPath(r.RemoteFilePath),
                Des:      r.LocalFilePath,
            },
        },
        Dirs: nil,
    }

    if fc.enableFtransProxy() {
        body.ServerAddress = fc.FtransProxyAddr
    } else {
        body.ServerAddress = fc.FtransS2Addr
    }

    url := fmt.Sprintf("http://%s%s?op=%s&t=%d&sig=%s", fc.FtransClientAddr, FtransS2UriPath, op, t, sig)
    data, _ := json.Marshal(body)

    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, data)
    if err != nil {
        return nil, fmt.Errorf("ftrans:downloadFileS2: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:downloadFileS2: invalid http status [%d]", resp.StatusCode)
    }

    for {
        sr, err := fc.statusDownloadS2(&body)
        if err != nil {
            return nil, fmt.Errorf("ftrans:downloadFileS2: status download failed [%s]", err.Error())
        }

        if sr.StatusCode != FtransStatusSucc && sr.StatusCode != FtransStatusDownloadFileNone {
            return nil, fmt.Errorf("ftrans:downloadFileS2: status download failed [%d]", sr.StatusCode)
        }

        if sr.TransStatus == FtransTransStatusCompleted {
            break
        } else if sr.TransStatus == FtransTransStatusFailed || sr.TransStatus == FtransTransStatusCancelled {
            return nil, fmt.Errorf("ftrans:downloadFileS2: download file [%s] failed [%s]",
                r.LocalFilePath, sr.TransStatus)
        }

        time.Sleep(time.Second)
    }

    return nil, nil
}

func (fc *FtransClient) listFileS2(r *FtransListFileRequest) (*FtransListFileResponse, error) {
    op := "list_dir"
    t := time.Now().Unix() + FtransClientSignatureExpiredNSec
    sig := fc.genFtransClientSig(op, fc.FtransClientAclToken, t)

    url := fmt.Sprintf("http://%s%s?op=%s&t=%d&sig=%s", fc.FtransClientAddr, FtransS2UriPath, op, t, sig)

    r.Prefix = fc.getFullPath(r.Prefix)
    if fc.enableFtransProxy() {
        r.ServerAddress = fc.FtransProxyAddr
    } else {
        r.ServerAddress = fc.FtransS2Addr
    }

    data, _ := json.Marshal(r)

    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, data)
    if err != nil {
        return nil, fmt.Errorf("ftrans:listFileS2: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:listFileS2: invalid http status [%d]", resp.StatusCode)
    }

    if resp.Body != nil {
        defer resp.Body.Close()
    }

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:listFileS2: read response body failed [%s]", err.Error())
    }

    var res FtransListFileResponse
    if err := json.Unmarshal(respBody, &res); err != nil {
        return nil, fmt.Errorf("ftrans:listFileS2: unmarshal response body failed [%s]", err.Error())
    }

    for _, elem := range res.Records {
        mtime, _ := time.Parse("2006-01-02 15:04:05", elem.LastModified)
        elem.MTime = mtime.UnixMicro()
    }

    return &res, nil
}

func (fc *FtransClient) statFileS2(filename string) (*FtransFileInfo, error) {
    r := &FtransListFileRequest{
        ServerName: fc.FtransServerName,
        AclToken:   fc.FtransAclToken,
        Prefix:     path.Dir(filename),
        FilterIn:   fc.normalizeFilterIn(path.Base(filename)),
        PageNum:    DefaultFtransPageNo,
        PageSize:   DefaultFtransPageSize,
        OrderType:  FtransOrderTypeAsc,
        OrderField: FtransOrderFieldName,
    }

    resp, err := fc.listFileS2(r)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statFileS2: stat file [%s] failed [%s]", filename, err.Error())
    }

    if resp.Total == 0 {
        return nil, fmt.Errorf("ftrans:statFileS2: file[%s] not found", filename)
    }

    resp.Records[0].Name = filename

    return resp.Records[0], nil
}

func (fc *FtransClient) removeFileS2(filename string) error {
    op := "remove_file"
    t := time.Now().Unix() + FtransClientSignatureExpiredNSec
    sig := fc.genFtransClientSig(op, fc.FtransClientAclToken, t)

    url := fmt.Sprintf("http://%s%s?op=%s&t=%d&sig=%s", fc.FtransClientAddr, FtransS2UriPath, op, t, sig)

    r := FtransRemoveFileRequest{
        ServerName:    fc.FtransServerName,
        ServerAddress: "",
        AclToken:      fc.FtransAclToken,
        Des:           fc.getFullPath(filename),
    }

    if fc.enableFtransProxy() {
        r.ServerAddress = fc.FtransProxyAddr
    } else {
        r.ServerAddress = fc.FtransS2Addr
    }

    data, _ := json.Marshal(r)

    resp, err := fc.doRequest(nil, http.MethodPost, url, nil, data)
    if err != nil {
        return fmt.Errorf("ftrans:removeFileS2: do http request [%s] failed [%s]", filename, err.Error())
    }

    if resp.Body != nil {
        defer resp.Body.Close()
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:removeFileS2: invalid response status [%d]", resp.StatusCode)
    }

    return nil
}

func (fc *FtransClient) removeFileS10(filename string) error {
    op := "remove_file"
    des := fc.encodePath(fc.getFullPath(filename))
    aclToken := fc.FtransAclToken

    url := fmt.Sprintf("https://%s%s?op=%s&des=%s&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, des, aclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)
    if err != nil {
        return fmt.Errorf("ftrans:removeFileS10: do http request failed [%s]", err.Error())
    }

    if resp.Body != nil {
        defer resp.Body.Close()
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:removeFileS10: invalid http status [%s]", err.Error())
    }

    return nil
}

func (fc *FtransClient) uploadPartS10(freeCh chan bool, taskCh chan *ftransPartTask, callbackCh chan *ftransPartTask) {
    for {
        freeCh <- true
        select {
        case fpt := <-taskCh:
            if err := fc.uploadFilePartS10(fpt); err != nil {
                fpt.Status = FtransTaskStatusFailed
            } else {
                fpt.Status = FtransTaskStatusFinished
            }

            callbackCh <- fpt
        }
    }
}

func (fc *FtransClient) uploadFilePartS10(fpt *ftransPartTask) interface{} {
    op := "upload_file"
    des := fc.encodePath(fc.getFullPath(fpt.parentTask.TempFilePath))
    url := fmt.Sprintf("https://%s%s?op=%s&des=%s&offset=%d&size=%d&mtime=%d&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, des, fpt.Offset, fpt.Size, fpt.parentTask.MtimeUSec, fc.FtransAclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, fpt.Data)

    if err != nil {
        return fmt.Errorf("ftrans:uploadFilePartS10: upload part(%d, %d) for [%s] failed [%s]",
            fpt.Offset, fpt.Size, fpt.parentTask.LocalFilePath, err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:uploadFilePartS10: invalid http status [%d] when uploading part(%d, %d) for [%s]",
            resp.StatusCode, fpt.Offset, fpt.Size, fpt.parentTask.LocalFilePath)
    }

    return nil
}

func (fc *FtransClient) existFileS10(filename string) error {
    op := "exist_file"
    src := fc.encodePath(fc.getFullPath(filename))

    url := fmt.Sprintf("https://%s%s?op=%s&src=%s&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, src, fc.FtransAclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)
    if err != nil {
        return fmt.Errorf("ftrans:existFileS10: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:existFileS10: invalid http status [%d]", resp.StatusCode)
    }

    return nil
}

func (fc *FtransClient) downloadPartS10(freeCh chan bool, taskCh chan *ftransPartTask, callbackCh chan *ftransPartTask) {
    for {
        freeCh <- true
        select {
        case fpt := <-taskCh:
            if err := fc.downloadFilePartS10(fpt); err != nil {
                fpt.Status = FtransTaskStatusFailed
            } else {
                fpt.Status = FtransTaskStatusFinished
            }

            callbackCh <- fpt
        }
    }
}

func (fc *FtransClient) downloadFilePartS10(fpt *ftransPartTask) interface{} {
    op := "download_file"
    src := fc.encodePath(fc.getFullPath(fpt.parentTask.RemoteFilePath))
    url := fmt.Sprintf("https://%s%s?op=%s&src=%s&offset=%d&size=%d&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, src, fpt.Offset, fpt.Size, fc.FtransAclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)

    if err != nil {
        return fmt.Errorf("ftrans:downloadFilePartS10: download part(%d, %d) for [%s] failed [%s]",
            fpt.Offset, fpt.Size, fpt.parentTask.LocalFilePath, err.Error())
    }

    if resp.Body != nil {
        defer resp.Body.Close()
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:downloadFilePartS10: invalid http status [%d] when downloading part(%d, %d) for [%s]",
            resp.StatusCode, fpt.Offset, fpt.Size, fpt.parentTask.LocalFilePath)
    }

    fileSize, err := strconv.ParseInt(resp.Header.Get("Fsize"), 10, 64)
    if err != nil {
        return fmt.Errorf("ftrans:downloadFilePartS10: invalid Fsize[%s] found when downloading part(%d, %d) for [%s]",
            resp.Header.Get("Fsize"), fpt.Offset, fpt.Size, fpt.parentTask.RemoteFilePath)
    }

    mtime, err := strconv.ParseInt(resp.Header.Get("Mtime"), 10, 64)
    if err != nil {
        return fmt.Errorf("ftrans:downloadFilePartS10: invalid Mtime[%s] found when downloading part(%d, %d) for [%s]",
            resp.Header.Get("Mtime"), fpt.Offset, fpt.Size, fpt.parentTask.RemoteFilePath)
    }

    if fpt.parentTask.FileSize != fileSize || fpt.parentTask.MtimeUSec != mtime {
        return fmt.Errorf("ftrans:downloadFilePartS10: file has been changed during downloading")
    }

    fpt.Data, err = ioutil.ReadAll(resp.Body)

    return err
}

func (fc *FtransClient) makeDirS10(dir string, mtime int64) error {
    op := "make_dir"
    des := fc.encodePath(fc.getFullPath(dir))
    url := fmt.Sprintf("https://%s%s?op=%s&des=%s&mtime=%d&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, des, mtime, fc.FtransAclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)
    if err != nil {
        return fmt.Errorf("ftrans:makeDirS10: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:makeDirS10: invalid http status [%d]", resp.StatusCode)
    }

    return nil
}

func (fc *FtransClient) renameFileS10(src, des string) error {
    op := "rename_file"
    src = fc.encodePath(fc.getFullPath(src))
    des = fc.encodePath(fc.getFullPath(des))

    url := fmt.Sprintf("https://%s%s?op=%s&src=%s&des=%s&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, src, des, fc.FtransAclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)
    if err != nil {
        return fmt.Errorf("ftrans:renameFileS10: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:renameFileS10: invalid http status [%d]", resp.StatusCode)
    }

    return nil
}

func (fc *FtransClient) touchFileS10(des string, mtime int64) error {
    op := "touch_file"
    des = fc.encodePath(fc.getFullPath(des))

    url := fmt.Sprintf("https://%s%s?op=%s&des=%s&mtime=%d&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, des, mtime, fc.FtransAclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)
    if err != nil {
        return fmt.Errorf("ftrans:touchFileS10: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ftrans:touchFileS10: invalid http status [%d]", resp.StatusCode)
    }

    return nil
}

func (fc *FtransClient) listFileS10(r *FtransListFileRequest) (*FtransListFileResponse, error) {
    op := "list_dir"
    src := fc.encodePath(fc.getFullPath(r.Prefix))
    filterIn := fc.encodePath(r.FilterIn)
    url := fmt.Sprintf("https://%s%s?op=%s&src=%s&filterIn=%s&sortBy=%s&orderBy=%s&pageNo=%d&pageSize=%d&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, src, filterIn, r.OrderField, r.OrderType, r.PageNum, r.PageSize, r.AclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)
    if err != nil {
        return nil, fmt.Errorf("ftrans:listFileS10: do http request failed [%s]", err.Error())
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:listFileS10: invalid http status [%d]", resp.StatusCode)
    }

    if resp.Body != nil {
        defer resp.Body.Close()
    }

    total, err := strconv.ParseInt(resp.Header.Get("matchedNum"), 10, 64)
    if err != nil {
        return nil, fmt.Errorf("ftrans:listFileS10: invalid total number[%s] found", resp.Header.Get("matchedNum"))
    }

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ftrans:listFileS10: read response body failed [%s]", err.Error())
    }

    var records []*FtransFileInfo
    if err := json.Unmarshal(respBody, &records); err != nil {
        return nil, fmt.Errorf("ftrans:listFileS10: unmarshal body failed [%s]", err.Error())
    }

    listResp := &FtransListFileResponse{
        Total:   total,
        Records: records,
    }

    return listResp, nil
}

func (fc *FtransClient) statFileS10(filename string) (*FtransFileInfo, error) {
    op := "size_file"
    src := fc.encodePath(fc.getFullPath(filename))
    url := fmt.Sprintf("https://%s%s?op=%s&src=%s&acltoken=%s",
        fc.FtransS10Addr, FtransS10UriPath, op, src, fc.FtransAclToken)

    resp, err := fc.doRequest(fc.FtransS10Conn, http.MethodPost, url, nil, nil)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statFileS10 do http request failed [%s]", err.Error())
    }

    if resp.Body != nil {
        resp.Body.Close()
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ftrans:statFileS10: invalid response status [%d]", resp.StatusCode)
    }

    size, err := strconv.ParseInt(resp.Header.Get("Fsize"), 10, 64)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statFileS10: invalid Fsize [%s]", resp.Header.Get("Fsize"))
    }

    mtime, err := strconv.ParseInt(resp.Header.Get("Mtime"), 10, 64)
    if err != nil {
        return nil, fmt.Errorf("ftrans:statFileS10: invalid mtime [%s]", resp.Header.Get("Mtime"))
    }

    ffi := &FtransFileInfo{
        Name:  filename,
        Size:  size,
        MTime: mtime,
    }

    return ffi, nil
}

func (fc *FtransClient) downloadFileS10(r *FtransFileRequest) (*FtransFileInfo, error) {
    if err := fc.existFileS10(r.RemoteFilePath); err != nil {
        return nil, fmt.Errorf("ftrans:downloadFileS10: file [%s] not found in server", r.RemoteFilePath)
    }

    fi, err := fc.statFileS10(r.RemoteFilePath)
    if err != nil {
        return nil, fmt.Errorf("ftrans:downloadFileS10: stat file [%s] failed [%s]", r.RemoteFilePath, err.Error())
    }

    st, err := os.Stat(r.LocalFilePath)
    // 本地文件存在且与远端一致
    if err == nil && !st.IsDir() {
        if st.ModTime().UnixMicro() == fi.MTime && st.Size() == fi.Size {
            return &FtransFileInfo{
                Name:  r.RemoteFilePath,
                Size:  fi.Size,
                MTime: fi.MTime,
                Md5:   "",
            }, nil
        }

        _ = os.Remove(r.LocalFilePath)
    }

    fft := &ftransFileTask{
        LocalFilePath:  r.LocalFilePath,
        RemoteFilePath: r.RemoteFilePath,
        TempFilePath:   fmt.Sprintf("%s_%s", r.LocalFilePath, uuid.New().String()),
        FileSize:       fi.Size,
        PartSize:       FtransPartSize,
        MtimeUSec:      fi.MTime,
        File:           nil,
        PartTasks:      nil,
        DoingParts:     0,
        Status:         FtransTaskStatusDoing,
    }

    fc.splitPartTaskS10(fft)
    fft.DoingParts = int32(len(fft.PartTasks))

    saveDir := filepath.Dir(fft.TempFilePath)
    st, err = os.Stat(saveDir)
    if err != nil && os.IsNotExist(err) {
        if err := os.MkdirAll(saveDir, 0755); err != nil {
            return nil, fmt.Errorf("ftrans:downloadFileS10: make dir [%s] failed [%s]", saveDir, err.Error())
        }
    } else if !st.IsDir() {
        return nil, fmt.Errorf("ftrans:downloadFileS10: [%s] is a file, please remove it first", saveDir)
    }

    f, err := os.OpenFile(fft.TempFilePath, os.O_CREATE|os.O_RDWR, 0755)
    if err != nil {
        return nil, fmt.Errorf("ftrans:downloadFileS10: open local file [%s] failed [%s]", fft.TempFilePath, err.Error())
    }
    fft.File = f
    defer func(fft *ftransFileTask) {
        if fft.File != nil {
            fft.File.Close()
        }
    }(fft)

    if fft.FileSize <= 0 {
        fft.File.Close()
        if err := os.Rename(fft.TempFilePath, fft.LocalFilePath); err != nil {
            return nil, fmt.Errorf("ftrans:downloadFileS10: rename file [%s] into [%s] failed [%s]",
                fft.TempFilePath, fft.LocalFilePath, err.Error())
        }
        return &FtransFileInfo{
            Name:  fft.RemoteFilePath,
            Size:  fft.FileSize,
            MTime: fft.MtimeUSec,
        }, nil
    }

    freeCh := make(chan bool, FtransPartConcurrency)
    taskCh := make(chan *ftransPartTask)
    callbackCh := make(chan *ftransPartTask)

    for idx := 0; idx < FtransPartConcurrency; idx++ {
        go fc.downloadPartS10(freeCh, taskCh, callbackCh)
    }

    for {
        select {
        case <-freeCh:
            if fft.WaitingParts <= 0 {
                continue
            }

            fpt := fft.PartTasks[fft.WaitingParts-1]
            taskCh <- fpt
            fft.WaitingParts--
        case fpt := <-callbackCh:
            if fpt.Status != FtransTaskStatusFinished {
                fft.Status = fpt.Status
                continue
            }

            fft.DoingParts--

            if fft.DoingParts == 0 {
                fft.Status = FtransTaskStatusFinished
            }

            _, err := fft.File.WriteAt(fpt.Data, fpt.Offset)
            if err != nil {
                return nil, fmt.Errorf("ftrans:downloadFileS10: write file [%s] at (%d, %d) failed [%s]",
                    fft.TempFilePath, fpt.Offset, fpt.Size, err.Error())
            }
        }

        if fft.Status != FtransTaskStatusDoing {
            break
        }
    }

    if fft.Status != FtransTaskStatusFinished {
        if err := os.Remove(fft.TempFilePath); err != nil {
            return nil, fmt.Errorf("ftrans:downloadFileS10: remove temp file failed [%s]", err.Error())
        }
        return nil, fmt.Errorf("ftrans:downloadFileS10: download file [%s] failed", fft.LocalFilePath)
    }

    if fft.File != nil {
        fft.File.Close()
        fft.File = nil
    }
    if err := os.Rename(fft.TempFilePath, fft.LocalFilePath); err != nil {
        return nil, fmt.Errorf("ftrans:downloadFileS10: rename file [%s] into [%s] failed [%s]",
            fft.TempFilePath, fft.LocalFilePath, err.Error())
    }

    _ = os.Chtimes(fft.LocalFilePath, time.Now(), time.UnixMicro(fft.MtimeUSec))

    resp := &FtransFileInfo{
        Name:  fft.RemoteFilePath,
        Size:  fft.FileSize,
        MTime: fft.MtimeUSec,
    }

    return resp, nil
}

func (fc *FtransClient) uploadFileS10(r *FtransFileRequest) (*FtransFileInfo, error) {
    st, err := os.Stat(r.LocalFilePath)
    if err != nil {
        return nil, fmt.Errorf("ftrans:uploadFileS10: stat local file [%s] failed [%s]", r.LocalFilePath, err.Error())
    }

    fi, err := fc.statFileS10(r.RemoteFilePath)
    if err != nil {
        return nil, fmt.Errorf("ftrans:uploadFileS10: stat remote file [%s] failed [%s]", r.RemoteFilePath, err.Error())
    }

    // 文件没有变化
    if fi.Size == st.Size() && fi.MTime == st.ModTime().UnixMicro() {
        return &FtransFileInfo{Name: r.RemoteFilePath, Size: fi.Size, MTime: fi.MTime}, nil
    }

    // 空目录
    if st.IsDir() {
        ffi := FtransFileInfo{
            Name:  r.RemoteFilePath,
            Size:  0,
            MTime: st.ModTime().UnixMicro(),
        }

        if err := fc.makeDirS10(ffi.Name, ffi.MTime); err != nil {
            return nil, fmt.Errorf("ftrans:uploadFileS10: make dir [%s] failed [%s]", ffi.Name, err.Error())
        }

        return &ffi, nil
    }

    // 空文件
    if st.Size() == 0 {
        if err := fc.touchFileS10(r.RemoteFilePath, st.ModTime().UnixMicro()); err != nil {
            return nil, fmt.Errorf("ftrans:uploadFileS10: touch file [%s] failed [%s]", r.RemoteFilePath, err.Error())
        }
    }

    f, err := os.OpenFile(r.LocalFilePath, os.O_RDONLY, 0755)
    if err != nil {
        return nil, fmt.Errorf("ftrans:uploadFileS10: open local file [%s] failed [%s]", r.LocalFilePath, err.Error())
    }
    defer f.Close()

    fft := &ftransFileTask{
        LocalFilePath:  r.LocalFilePath,
        RemoteFilePath: r.RemoteFilePath,
        TempFilePath:   fmt.Sprintf("%s_%s", r.RemoteFilePath, uuid.New().String()),
        FileSize:       st.Size(),
        PartSize:       FtransPartSize,
        MtimeUSec:      st.ModTime().UnixMicro(),
        File:           f,
        PartTasks:      nil,
        DoingParts:     0,
        Status:         FtransTaskStatusDoing,
    }

    fc.splitPartTaskS10(fft)

    freeCh := make(chan bool, FtransPartConcurrency)
    taskCh := make(chan *ftransPartTask)
    callbackCh := make(chan *ftransPartTask)

    for idx := 0; idx < FtransPartConcurrency; idx++ {
        go fc.uploadPartS10(freeCh, taskCh, callbackCh)
    }

    for {
        select {
        case <-freeCh:
            if fft.WaitingParts <= 0 {
                continue
            }

            // 检查文件传输过程中是否发生变化
            if true {
                fi, err := os.Stat(r.LocalFilePath)
                if err != nil {
                    return nil, fmt.Errorf("ftrans:uploadFileS10: stat file [%s] failed [%s]",
                        fft.LocalFilePath, err.Error())
                }

                if fi.Size() != fft.FileSize && fi.ModTime().UnixMicro() != fft.MtimeUSec {
                    return nil, fmt.Errorf("ftrans:uploadFileS10: file [%s] has been changed during uploading",
                        fft.LocalFilePath)
                }
            }

            fpt := fft.PartTasks[fft.WaitingParts-1]
            _, err = fft.File.ReadAt(fpt.Data, fpt.Offset)
            if err != nil {
                return nil, fmt.Errorf("ftrans:uploadFileS10: read part(%d, %d) for [%s] failed [%s]",
                    fpt.Offset, fpt.Size, fft.LocalFilePath, err.Error())
            }
            taskCh <- fpt
            fft.WaitingParts--
        case fpt := <-callbackCh:
            fft.DoingParts--
            if fpt.Status != FtransTaskStatusFinished {
                fft.Status = fpt.Status
            } else {
                if fft.DoingParts <= 0 {
                    fft.Status = FtransTaskStatusFinished
                    break
                }
            }
        }

        if fft.Status != FtransTaskStatusDoing {
            break
        }
    }

    if fft.Status != FtransTaskStatusFinished {
        if err := fc.RemoveFile(fft.TempFilePath); err != nil {
            return nil, fmt.Errorf("ftrans:uploadFileS10: remove temp file failed [%s]", err.Error())
        }
        return nil, fmt.Errorf("ftrans:uploadFileS10: upload file [%s] failed", fft.LocalFilePath)
    }

    if err := fc.renameFileS10(fft.TempFilePath, fft.RemoteFilePath); err != nil {
        return nil, err
    }

    resp := &FtransFileInfo{
        Name:  fft.RemoteFilePath,
        Size:  fft.FileSize,
        MTime: fft.MtimeUSec,
    }

    return resp, nil
}

func (fc *FtransClient) splitPartTaskS10(fft *ftransFileTask) {
    start := int64(0)
    end := fft.FileSize

    for ; start < end; start += fft.PartSize {
        size := fft.PartSize
        if start+size >= end {
            size = end - start
        }
        fpt := ftransPartTask{
            Offset:     start,
            Size:       size,
            Status:     FtransTaskStatusWaiting,
            Data:       make([]byte, size),
            parentTask: fft,
        }

        fft.PartTasks = append(fft.PartTasks, &fpt)
    }

    fft.DoingParts = 0
    fft.WaitingParts = int32(len(fft.PartTasks))
}
