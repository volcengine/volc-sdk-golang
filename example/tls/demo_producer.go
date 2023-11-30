package tls

import (
	"fmt"
	"os"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/producer"
)

func main() {
	// 初始化客户端，推荐通过环境变量动态获取火山引擎密钥等身份认证信息，以免AccessKey硬编码引发数据安全风险。详细说明请参考 https://www.volcengine.com/docs/6470/1166455
	// 使用STS时，ak和sk均使用临时密钥，且设置VOLCENGINE_TOKEN；不使用STS时，VOLCENGINE_TOKEN部分传空
	tlsProducerCfg := producer.GetDefaultProducerConfig()
	tlsProducerCfg.Endpoint = os.Getenv("VOLCENGINE_ENDPOINT")
	tlsProducerCfg.Region = os.Getenv("VOLCENGINE_REGION")
	tlsProducerCfg.AccessKeyID = os.Getenv("VOLCENGINE_ACCESS_KEY_ID")
	tlsProducerCfg.AccessKeySecret = os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET")

	// 初始化并启动Producer
	tlsProducer := producer.NewProducer(tlsProducerCfg)
	tlsProducer.Start()

	// 请根据您的需要，填写topicId、source、filename
	topicID := "your-topic-id"
	source := "your-log-source"
	filename := "your-log-filename"

	// 调用Producer SendLog接口，一次提交一条日志
	// 您可根据实际需要，自行定义实现用于业务处理的CallBack，传入SendLog接口
	err := tlsProducer.SendLog("", topicID, source, filename, &pb.Log{
		Contents: []*pb.LogContent{
			{
				Key:   "key1",
				Value: "value1",
			},
			{
				Key:   "key2",
				Value: "value2",
			},
		},
		Time: time.Now().Unix(),
	}, nil)
	if err != nil {
		// 处理错误
		fmt.Println(err.Error())
	}

	// 调用Producer SendLogs接口，一次提交多条日志
	// 您可根据实际需要，自行定义实现用于业务处理的CallBack，传入SendLogs接口
	err = tlsProducer.SendLogs("", topicID, source, filename, &pb.LogGroup{
		Source:   source,
		FileName: filename,
		Logs: []*pb.Log{
			{
				Contents: []*pb.LogContent{
					{
						Key:   "key1",
						Value: "value1-1",
					},
					{
						Key:   "key2",
						Value: "value2-1",
					},
				},
				Time: time.Now().Unix(),
			},
			{
				Contents: []*pb.LogContent{
					{
						Key:   "key1",
						Value: "value1-2",
					},
					{
						Key:   "key2",
						Value: "value2-2",
					},
				},
				Time: time.Now().Unix(),
			},
		},
	}, nil)
	if err != nil {
		// 处理错误
		fmt.Println(err.Error())
	}

	// 关闭Producer
	tlsProducer.Close()
}
