package tls

import (
	"fmt"
	"os"

	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	//endpoint = "https://tls-cn-beijing.volces.com"
	//access_key_id = "AKLxxxxxxxx"
	//access_key_secret = "TUxxxxxxxxxx=="
	//region = "cn-beijing"
	client := tls.NewClient(os.Getenv("VOLCENGINE_ENDPOINT"), os.Getenv("VOLCENGINE_ACCESS_KEY_ID"),
		os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET"), os.Getenv("VOLCENGINE_TOKEN"), os.Getenv("VOLCENGINE_REGION"))

	// 请填写您的TopicId
	topicID := "your-topic-id"

	// 获取直方图
	// DescribeHistogram API的请求参数规范请参阅 https://www.volcengine.com/docs/6470/142136
	describeHistogramResponse, _ := client.DescribeHistogram(&tls.DescribeHistogramRequest{
		TopicID:   topicID,
		Query:     "*",
		StartTime: 1672502400000,
		EndTime:   1688140800000,
	})
	fmt.Printf("%v\n", describeHistogramResponse)
}
