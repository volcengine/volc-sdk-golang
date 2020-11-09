

## Q&A

1. SDK读取如何加载AKSK？

##### 显示的在client中设置AKSK

```golang
client.SetAccessKey("Your-AK")
client.SetSecretKey("Your-SK")
```

##### 从环境变量加载AKSK

环境变量信息

```
VOLC_ACCESSKEY #ak
VOLC_SECRETKEY #sk
```

##### 从HOME文件加载AKSK

~/.volc/config

```json
{
	"ak":"Your-AK",
	"sk":"Your-SK"
}
```

