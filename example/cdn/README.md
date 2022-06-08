# CDN接口调用示例

## 前置准备
### 服务开通
请确保您已开通了CDN服务。您可前往[火山引擎控制台](https://console.volcengine.com/ )，在左侧菜单中选择或在顶部搜索栏中搜索CDN服务，进入服务控制台内完成开通流程。
### 获取安全凭证
Access Key（访问密钥）是访问火山引擎服务的安全凭证，包含Access Key ID（简称为AK）和Secret Access Key（简称为SK）两部分。您可登录[火山引擎控制台](https://console.volcengine.com/ )，前往“[访问控制](https://console.volcengine.com/iam )”的“[访问密钥](https://console.volcengine.com/iam/keymanage/ )”中创建及管理您的Access Key。更多信息可参考[访问密钥帮助文档](https://www.volcengine.com/docs/6291/65568 )。
### 环境检查
Go版本需要不低于1.13.1。

## 运行方式

- 设置AK/SK
```go
// init.go
ak              = "Your AK"
sk              = "Your SK"
```
- 设置通用参数
```go

// init.go
operateDomain   = "Your domain"
exampleDomain   = "Your domain"
```

- 编写测试用例: cdn_test.go
```go
package cdn

import (
	"testing"
)

func TestCdn(t *testing.T) {
	// 调用查询域名列表接口
	ListCdnDomains(t)
}

```

- 执行测试用例
```go
go test -v cdn_test.go init.go list_cdn_domains.go

```