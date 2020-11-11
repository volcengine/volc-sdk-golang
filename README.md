# 火山引擎开发者Go SDK

## 安装sdk(推荐)

建议使用go >= 1.13.1

```go get -u github.com/volcengine/volc-sdk-golang```

### AK/SK 注册申请流程

主账户和有权限的子用户可以新建AK密钥，操作如下：

1.使用帐号/密码登录控制台；

2.选择一级菜单“访问控制”->选择二级菜单“密钥管理”；

3.页面中展示主账号的访问密钥列表，每个IAM用户最多可同时拥有2个访问密钥，如果当前IAM用户的访问密钥数量未达到上限，则可以点击新建密钥按钮；

4.点击新建密钥按钮，弹出新建密钥弹窗，点击查看AccessKey详情，可直接查看访问密钥信息。

### 通过API申请AK/SK
 
[生成访问密匙](https://www.volcengine.cn/docs/6291/65578)

## 显示的在client中设置AKSK (推荐)

```go
iam.DefaultInstance.Client.SetAccessKey(testAk)
iam.DefaultInstance.Client.SetSecretKey(testSk)	
```

## 从环境变量加载AKSK

环境变量信息

```bash
VOLC_ACCESSKEY="ak"
VOLC_SECRETKEY="sk"
```

## 从HOME文件加载AKSK

在本地的`~/.volc/config`中添加如下内容：

```json
{
	"ak":"Your-AK",
	"sk":"Your-SK"
}
```

## SDK服务目录及实例

-【视觉智能】请点击[这里](service/visual/README.md)