## Example

调用代码示例均在example/visual文件夹下，以下为银行卡OCR调用示例
```
package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func main() {
	testAk := "your-ak"
	testSk := "your-sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)

	form := url.Values{}
	form.Add("image_base64", "")

	resp, status, err := visual.DefaultInstance.BankCard(form)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
```
运行代码方式，在根目录下执行
```
go run example/visual/bank_card.go
```

## 接口文档
文档链接请点击[这里](https://www.volcengine.cn/docs)
并在【视觉智能】列表查看

