package consumer

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/producer"
	"os"
	"strconv"
	"testing"
	"time"
)

type SdkConsumerTestSuite struct {
	suite.Suite

	cli       tls.Client
	projectId string
	topicId   string
	producer  producer.Producer
	timestamp int64
}

func TestSdkConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(SdkConsumerTestSuite))
}

func (suite *SdkConsumerTestSuite) SetupTest() {

	suite.timestamp = time.Now().Unix()

	suite.cli = NewClientWithEnv()

	// 创建project
	createProjectReq := &tls.CreateProjectRequest{
		ProjectName: "consumer-test-project" + uuid.New().String(),
		Description: "测试项目",
		Region:      os.Getenv("LOG_SERVICE_REGION"),
	}

	createProjectResp, err := suite.cli.CreateProject(createProjectReq)
	suite.NoError(err)
	suite.projectId = createProjectResp.ProjectID

	fmt.Printf("ProjectID is %s\n", suite.projectId)

	// 创建topic
	createTopicReq := &tls.CreateTopicRequest{
		ProjectID:   suite.projectId,
		TopicName:   "consumer-test-topic" + uuid.New().String(),
		Ttl:         1,
		Description: "测试topic",
		ShardCount:  1,
	}

	createTopicResp, err := suite.cli.CreateTopic(createTopicReq)
	suite.NoError(err)
	suite.topicId = createTopicResp.TopicID

	fmt.Printf("TopicID is %s\n", suite.topicId)

	// 写数据
	// init producer
	producerCfg := producer.GetDefaultProducerConfig()
	producerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
	producerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
	producerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
	producerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")

	suite.producer = producer.NewProducer(producerCfg)
	suite.producer.Start()

	logGroup := &pb.LogGroup{
		Logs:     nil,
		Source:   "/test/sdk/localhost",
		FileName: "test-sdk.log",
	}

	var logContents []*pb.LogContent

	for i := 0; i < 10; i++ {

		idx := strconv.Itoa(i)

		logContent := &pb.LogContent{
			Key:   "key-" + idx,
			Value: "test-message" + idx,
		}

		logContents = append(logContents, logContent)
	}

	logs := []*pb.Log{
		{
			Contents: logContents,
			Time:     time.Now().Unix(),
		},
	}

	logGroup.Logs = logs

	err = suite.producer.SendLogs("", suite.topicId, logGroup.Source, logGroup.FileName, logGroup, nil)
	suite.NoError(err)

	// 关闭Producer
	suite.producer.Close()

	fmt.Println("logs send success")
}

func (suite *SdkConsumerTestSuite) TearDownTest() {
	_, deleteTopicErr := suite.cli.DeleteTopic(&tls.DeleteTopicRequest{TopicID: suite.topicId})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&tls.DeleteProjectRequest{ProjectID: suite.projectId})
	suite.NoError(deleteProjectErr)
}

func (suite *SdkConsumerTestSuite) TestConsumer() {
	consumerCfg := &Config{
		ClientConfig: common.ClientConfig{
			Endpoint:        os.Getenv("LOG_SERVICE_ENDPOINT"),
			Region:          os.Getenv("LOG_SERVICE_REGION"),
			AccessKeyID:     os.Getenv("LOG_SERVICE_AK"),
			AccessKeySecret: os.Getenv("LOG_SERVICE_SK"),
		},
		LoggerConfig: common.LoggerConfig{
			LogLevel:      "info",
			LogFileName:   "",
			IsJsonType:    false,
			LogMaxSize:    10,
			LogMaxBackups: 10,
			LogCompress:   false,
		},
		ConsumeFrom:                    strconv.FormatInt(suite.timestamp, 10),
		HeartbeatIntervalInSecond:      20,
		DataFetchIntervalInMillisecond: 200,
		MaxFetchLogGroupCount:          100,
		FlushCheckpointIntervalSecond:  5,
		OrderedConsume:                 false,
		ProjectID:                      suite.projectId,
		TopicIDList:                    []string{suite.topicId},
		ConsumerGroupName:              "test-consumer-group-name",
	}

	var handleLogs = func(topicID string, shardID int, l *pb.LogGroupList) {
		fmt.Printf("received new logs from topic: %s, shard: %d\n", topicID, shardID)
		logCount := len(l.LogGroups)
		fmt.Printf("logCount: %d\n", logCount)
		for _, logGroup := range l.LogGroups {
			for _, log := range logGroup.Logs {
				fmt.Printf("time: %s\n", time.Unix(log.Time, 0).Format("2006-01-02 15:04:05"))
				for _, content := range log.Contents {
					fmt.Printf("%s: %s\n", content.Key, content.Value)
				}
			}
		}
	}

	// 创建消费者
	consumer, err := NewConsumer(context.TODO(), consumerCfg, handleLogs)
	suite.NoError(err)

	// 启动消费者消费
	err = consumer.Start()
	suite.NoError(err)

	// 等待消费
	time.Sleep(300 * time.Second)

	// 停止消费
	consumer.Stop()
}

func (suite *SdkConsumerTestSuite) TestStopConsumer() {
	consumerGroup := "test-consumer-group-name"
	consumerCfg := &Config{
		ClientConfig: common.ClientConfig{
			Endpoint:        os.Getenv("LOG_SERVICE_ENDPOINT"),
			Region:          os.Getenv("LOG_SERVICE_REGION"),
			AccessKeyID:     os.Getenv("LOG_SERVICE_AK"),
			AccessKeySecret: os.Getenv("LOG_SERVICE_SK"),
		},
		LoggerConfig: common.LoggerConfig{
			LogLevel:      "info",
			LogFileName:   "",
			IsJsonType:    false,
			LogMaxSize:    10,
			LogMaxBackups: 10,
			LogCompress:   false,
		},
		ConsumeFrom:                    strconv.FormatInt(suite.timestamp, 10),
		HeartbeatIntervalInSecond:      20,
		DataFetchIntervalInMillisecond: 200,
		MaxFetchLogGroupCount:          100,
		FlushCheckpointIntervalSecond:  5,
		OrderedConsume:                 false,
		ProjectID:                      suite.projectId,
		TopicIDList:                    []string{suite.topicId},
		ConsumerGroupName:              consumerGroup,
	}

	_, err := suite.cli.CreateConsumerGroup(&tls.CreateConsumerGroupRequest{
		ProjectID:         suite.projectId,
		TopicIDList:       []string{suite.topicId},
		ConsumerGroupName: consumerGroup,
		HeartbeatTTL:      60,
		OrderedConsume:    true,
	})
	suite.NoError(err)

	checkPointInfo, err := suite.cli.DescribeCheckPoint(&tls.DescribeCheckPointRequest{
		ProjectID:         suite.projectId,
		TopicID:           suite.topicId,
		ConsumerGroupName: consumerGroup,
		ShardID:           0,
	})
	suite.NoError(err)
	suite.Equal(int32(0), checkPointInfo.ShardID)
	suite.Equal("", checkPointInfo.Checkpoint)

	stopFlag := make(chan struct{})

	var handleLogs = func(topicID string, shardID int, l *pb.LogGroupList) {
		stopFlag <- struct{}{}

		fmt.Printf("mock user handle logs")

		time.Sleep(10 * time.Second)
	}
	// 创建消费者
	consumer, err := NewConsumer(context.TODO(), consumerCfg, handleLogs)
	suite.NoError(err)

	// 启动消费者消费
	err = consumer.Start()
	suite.NoError(err)

	<-stopFlag

	consumer.Stop()

	checkPointInfo, err = suite.cli.DescribeCheckPoint(&tls.DescribeCheckPointRequest{
		ProjectID:         suite.projectId,
		TopicID:           suite.topicId,
		ConsumerGroupName: consumerGroup,
		ShardID:           0,
	})
	suite.NoError(err)
	suite.Equal(int32(0), checkPointInfo.ShardID)
	suite.NotEqual("", checkPointInfo.Checkpoint)
	fmt.Printf("consumer after: Checkpoint: %s, UpdateTime: %d, Consumer: %s\n",
		checkPointInfo.Checkpoint, checkPointInfo.UpdateTime, checkPointInfo.Consumer)
}
