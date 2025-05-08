package siem

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
	"github.com/volcengine/volc-sdk-golang/service/tls/consumer"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/producer"
)

func getConsumerCfg(projectID, topicID string) *consumer.Config {
	consumerCfg := &consumer.Config{
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
		ConsumeFrom:                    consumer.ConsumeFromBegin,
		HeartbeatIntervalInSecond:      20,
		DataFetchIntervalInMillisecond: 200,
		MaxFetchLogGroupCount:          100,
		FlushCheckpointIntervalSecond:  5,
		OrderedConsume:                 false,
		ProjectID:                      projectID,
		TopicIDList:                    []string{topicID},
		ConsumerGroupName:              "test-consumer-group-name",
	}

	return consumerCfg
}

func producerSyslogRFC3164(topicID string) error {
	var logs []*pb.Log
	for i := 0; i < 10; i++ {

		logContent := []*pb.LogContent{
			{
				Key:   "pri",
				Value: "34",
			},
			{
				Key:   "timestamp",
				Value: "Apr 29 14:12:58",
			},
			{
				Key:   "hostname",
				Value: "localhost",
			},
			{
				Key:   "appName",
				Value: "go-sdk-producer",
			},
			{
				Key:   "pid",
				Value: strconv.Itoa(os.Getgid()),
			},
			{
				Key:   "msgId",
				Value: "test-msgId",
			},
			{
				Key:   "message",
				Value: "this is a syslog 3164 message" + strconv.Itoa(i),
			},
		}

		logs = append(logs, &pb.Log{
			Time:     time.Now().Unix() + int64(i),
			Contents: logContent,
		})
	}

	logGroup := &pb.LogGroup{
		Source:   "syslog-source",
		FileName: "syslog-filename",
	}
	logGroup.Logs = logs

	// 写数据
	// init producer
	producerCfg := producer.GetDefaultProducerConfig()
	producerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
	producerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
	producerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
	producerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")

	tlsProducer := producer.NewProducer(producerCfg)
	tlsProducer.Start()

	err := tlsProducer.SendLogs("", topicID, logGroup.Source, logGroup.FileName, logGroup, nil)
	if err != nil {
		return err
	}

	// 关闭Producer
	tlsProducer.Close()

	fmt.Println("logs send success")

	return nil
}

func producerSyslogRFC5424(topicID string) error {
	var logs []*pb.Log
	for i := 0; i < 10; i++ {

		logContent := []*pb.LogContent{
			{
				Key:   "pri",
				Value: "34",
			},
			{
				Key:   "version",
				Value: "1",
			},
			{
				Key:   "timestamp",
				Value: "2025-04-29T14:30:00.000Z",
			},
			{
				Key:   "hostname",
				Value: "localhost",
			},
			{
				Key:   "appName",
				Value: "go-sdk-producer",
			},
			{
				Key:   "pid",
				Value: strconv.Itoa(os.Getgid()),
			},
			{
				Key:   "msgId",
				Value: "-",
			},
			{
				Key:   "message",
				Value: "this is a syslog 5424 message" + strconv.Itoa(i),
			},
		}

		logs = append(logs, &pb.Log{
			Time:     time.Now().Unix() + int64(i),
			Contents: logContent,
		})
	}

	logGroup := &pb.LogGroup{
		Source:   "syslog-source",
		FileName: "syslog-filename",
	}
	logGroup.Logs = logs

	// 写数据
	// init producer
	producerCfg := producer.GetDefaultProducerConfig()
	producerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
	producerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
	producerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
	producerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")

	tlsProducer := producer.NewProducer(producerCfg)
	tlsProducer.Start()

	err := tlsProducer.SendLogs("", topicID, logGroup.Source, logGroup.FileName, logGroup, nil)
	if err != nil {
		return err
	}

	// 关闭Producer
	tlsProducer.Close()

	fmt.Println("logs send success")

	return nil
}

type SiemTestSuite struct {
	suite.Suite

	tlsClient   tls.Client
	consumerCfg *consumer.Config

	projectId string
	topicId   string
}

func TestSiemTestSuite(t *testing.T) {
	suite.Run(t, new(SiemTestSuite))
}

func (s *SiemTestSuite) SetupSuite() {
	s.tlsClient = tls.NewClient(
		os.Getenv("LOG_SERVICE_ENDPOINT"),
		os.Getenv("LOG_SERVICE_AK"),
		os.Getenv("LOG_SERVICE_SK"),
		"",
		os.Getenv("LOG_SERVICE_REGION"),
	)

	// 创建project
	createProjectReq := &tls.CreateProjectRequest{
		ProjectName: "consumer-test-project" + uuid.New().String(),
		Description: "测试项目",
		Region:      os.Getenv("LOG_SERVICE_REGION"),
	}

	createProjectResp, err := s.tlsClient.CreateProject(createProjectReq)
	s.NoError(err)
	s.projectId = createProjectResp.ProjectID

	fmt.Printf("ProjectID is %s\n", s.projectId)

	// 创建topic
	createTopicReq := &tls.CreateTopicRequest{
		ProjectID:   s.projectId,
		TopicName:   "consumer-test-topic" + uuid.New().String(),
		Ttl:         1,
		Description: "测试topic",
		ShardCount:  1,
	}

	createTopicResp, err := s.tlsClient.CreateTopic(createTopicReq)
	s.NoError(err)
	s.topicId = createTopicResp.TopicID
}

func (s *SiemTestSuite) TearDownSuite() {
	_, deleteTopicErr := s.tlsClient.DeleteTopic(&tls.DeleteTopicRequest{TopicID: s.topicId})
	s.NoError(deleteTopicErr)
	_, deleteProjectErr := s.tlsClient.DeleteProject(&tls.DeleteProjectRequest{ProjectID: s.projectId})
	s.NoError(deleteProjectErr)
}

func (s *SiemTestSuite) SetupTest() {}

func (s *SiemTestSuite) TearDownTest() {}

func (s *SiemTestSuite) TestNewSyslogClientConnectAbnormal() {
	tests := []struct {
		name    string
		setup   func(*SyslogClient)
		wantErr bool
		errMsg  string
	}{
		{
			name: "host not set",
			setup: func(c *SyslogClient) {
				c.Host = ""
				c.Port = 514
			},
			wantErr: true,
			errMsg:  "host or port is not set",
		},
		{
			name: "port not set",
			setup: func(c *SyslogClient) {
				c.Host = "localhost"
				c.Port = 0
			},
			wantErr: true,
			errMsg:  "host or port is not set",
		},
		{
			name: "unsupported protocol",
			setup: func(c *SyslogClient) {
				c.Host = "localhost"
				c.Port = 514
				c.Protocol = "http"
			},
			wantErr: true,
			errMsg:  "unsupported protocol: http",
		},
		{
			name: "invalid max message length",
			setup: func(c *SyslogClient) {
				c.Host = "localhost"
				c.Port = 514
				c.Protocol = "tcp"
				maxMsgLen := -1
				c.MaxMessageLength = &maxMsgLen
			},
			wantErr: true,
			errMsg:  "invalid max message length: -1",
		},
		{
			name: "invalid conn timeout",
			setup: func(c *SyslogClient) {
				c.Host = "localhost"
				c.Port = 514
				c.Protocol = "tcp"
				maxMsgLen := 1024
				c.MaxMessageLength = &maxMsgLen
				c.ConnTimeout = 0
			},
			wantErr: true,
			errMsg:  "invalid conn timeout: 0",
		},
		{
			name: "invalid send timeout",
			setup: func(c *SyslogClient) {
				c.Host = "localhost"
				c.Port = 514
				c.Protocol = "tcp"
				maxMsgLen := 1024
				c.MaxMessageLength = &maxMsgLen
				c.ConnTimeout = 10
				c.SendTimeout = 0
			},
			wantErr: true,
			errMsg:  "invalid send timeout: 0",
		},
		{
			name: "valid settings",
			setup: func(c *SyslogClient) {
				c.Host = "localhost"
				c.Port = 514
				c.Protocol = "tcp"
				maxMsgLen := 1024
				c.MaxMessageLength = &maxMsgLen
				c.ConnTimeout = 10
				c.SendTimeout = 10
			},
			wantErr: true,
			errMsg:  "failed to connect to localhost:514 over tcp: dial tcp: lookup localhost: i/o timeout",
		},
	}

	for _, tt := range tests {
		client := &SyslogClient{}
		tt.setup(client)
		err := client.Connect()

		if tt.wantErr {
			s.Error(err)
			s.Contains(err.Error(), tt.errMsg)
		} else {
			s.NoError(err)
		}
	}

}

func (s *SiemTestSuite) TestSyncDataToSyslogRFC3164WithUdpAbnormal() {
	syslogClient := NewSyslogClient()
	syslogClient.Host = os.Getenv("SPLUNK_HOST")
	syslogClient.Port, _ = strconv.Atoi(os.Getenv("SPLUNK_PORT_UDP"))
	s.NoError(syslogClient.Connect())

	tests := []struct {
		name    string
		setup   func(*SyslogFormatRFC3164)
		wantErr bool
		errMsg  string
	}{
		{
			name: "priority not set",
			setup: func(c *SyslogFormatRFC3164) {
				c.Priority = ""
			},
			wantErr: true,
			errMsg:  "syslog priority is required",
		},
		{
			name: "timestamp not set",
			setup: func(c *SyslogFormatRFC3164) {
				c.Priority = "34"
			},
			wantErr: true,
			errMsg:  "syslog timestamp is required",
		},
		{
			name: "hostname not set",
			setup: func(c *SyslogFormatRFC3164) {
				c.Priority = "34"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
			},
			wantErr: true,
			errMsg:  "syslog hostname is required",
		},
		{
			name: "pid not set",
			setup: func(c *SyslogFormatRFC3164) {
				c.Priority = "34"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
				c.Hostname = "-"
			},
			wantErr: true,
			errMsg:  "syslog pid is required",
		},
		{
			name: "process not set",
			setup: func(c *SyslogFormatRFC3164) {
				c.Priority = "34"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
				c.Hostname = "-"
				c.Pid = "-"
			},
			wantErr: true,
			errMsg:  "syslog process is required",
		},
		{
			name: "message not set",
			setup: func(c *SyslogFormatRFC3164) {
				c.Priority = "34"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
				c.Hostname = "-"
				c.Pid = "-"
				c.Process = "-"
			},
			wantErr: true,
			errMsg:  "syslog message is required",
		},
	}

	for _, tt := range tests {
		syslogRFC3164 := &SyslogFormatRFC3164{}
		tt.setup(syslogRFC3164)
		_, err := syslogRFC3164.FormatSyslogMessageWithRFC3164()

		if tt.wantErr {
			s.Error(err)
			s.Contains(err.Error(), tt.errMsg)
		} else {
			s.NoError(err)
		}
	}
}

func (s *SiemTestSuite) TestSyncDataToSyslogRFC3164WithUdp() {
	s.NoError(producerSyslogRFC3164(s.topicId))

	syslogClient := NewSyslogClient()
	syslogClient.Host = os.Getenv("SPLUNK_HOST")
	syslogClient.Port, _ = strconv.Atoi(os.Getenv("SPLUNK_PORT_UDP"))
	s.NoError(syslogClient.Connect())

	s.consumerCfg = getConsumerCfg(s.projectId, s.topicId)

	var handleLogsFunc = func(topicID string, shardID int, l *pb.LogGroupList) {
		for _, logGroup := range l.LogGroups {
			for _, log := range logGroup.Logs {

				syslogRFC3164 := NewSyslogFormatRFC3164()

				for _, content := range log.Contents {
					switch content.Key {
					case "pri":
						syslogRFC3164.Priority = content.Value
					case "timestamp":
						timestamp, err := time.Parse(time.Stamp, content.Value)
						s.NoError(err)
						syslogRFC3164.Timestamp = timestamp
					case "hostname":
						syslogRFC3164.Hostname = content.Value
					case "pid":
						syslogRFC3164.Pid = content.Value
					case "process":
						syslogRFC3164.Process = content.Value
					case "message":
						syslogRFC3164.Message = content.Value
					}
				}

				message, err := syslogRFC3164.FormatSyslogMessageWithRFC3164()
				s.NoError(err)
				s.NoError(syslogClient.Send(message))
				fmt.Println("发送成功")
			}
		}
	}

	tlsConsumer, err := consumer.NewConsumer(context.TODO(), s.consumerCfg, handleLogsFunc)
	s.NoError(err)

	s.NoError(tlsConsumer.Start())

	time.Sleep(60 * time.Second)

	s.NoError(syslogClient.Close())
	tlsConsumer.Stop()
}

func (s *SiemTestSuite) TestSyncDataToSyslogRFC5424WithTcpAbnormal() {
	syslogClient := NewSyslogClient()
	syslogClient.Host = os.Getenv("SPLUNK_HOST")
	syslogClient.Port, _ = strconv.Atoi(os.Getenv("SPLUNK_PORT_TCP"))
	s.NoError(syslogClient.Connect())

	tests := []struct {
		name    string
		setup   func(*SyslogFormatRFC5424)
		wantErr bool
		errMsg  string
	}{
		{
			name: "priority not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = ""
			},
			wantErr: true,
			errMsg:  "syslog priority is required",
		},
		{
			name: "version not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = "34"
			},
			wantErr: true,
			errMsg:  "syslog version is required",
		},
		{
			name: "timestamp not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = "34"
				c.Version = "1"
			},
			wantErr: true,
			errMsg:  "syslog timestamp is required",
		},
		{
			name: "hostname not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = "34"
				c.Version = "1"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
			},
			wantErr: true,
			errMsg:  "syslog hostname is required",
		},
		{
			name: "app name not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = "34"
				c.Version = "1"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
				c.Hostname = "-"
			},
			wantErr: true,
			errMsg:  "syslog app name is required",
		},
		{
			name: "pid not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = "34"
				c.Version = "1"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
				c.Hostname = "-"
				c.AppName = "-"
			},
			wantErr: true,
			errMsg:  "syslog pid is required",
		},
		{
			name: "message id not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = "34"
				c.Version = "1"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
				c.Hostname = "-"
				c.AppName = "-"
				c.Pid = "-"
			},
			wantErr: true,
			errMsg:  "syslog message id is required",
		},
		{
			name: "message not set",
			setup: func(c *SyslogFormatRFC5424) {
				c.Priority = "34"
				c.Version = "1"
				c.Timestamp, _ = time.Parse(time.Stamp, "Apr 28 18:56:00")
				c.Hostname = "-"
				c.AppName = "-"
				c.Pid = "-"
				c.MsgId = "-"
			},
			wantErr: true,
			errMsg:  "syslog message is required",
		},
	}

	for _, tt := range tests {
		syslogRFC5424 := &SyslogFormatRFC5424{}
		tt.setup(syslogRFC5424)
		_, err := syslogRFC5424.FormatSyslogMessageWithRFC5424()

		if tt.wantErr {
			s.Error(err)
			s.Contains(err.Error(), tt.errMsg)
		} else {
			s.NoError(err)
		}
	}
}

func (s *SiemTestSuite) TestSyncDataToSyslogRFC5424WithTcp() {
	s.NoError(producerSyslogRFC5424(s.topicId))

	syslogClient := NewSyslogClient()
	syslogClient.Host = os.Getenv("SPLUNK_HOST")
	syslogClient.Port, _ = strconv.Atoi(os.Getenv("SPLUNK_PORT_TCP"))
	syslogClient.Protocol = "tcp"
	s.NoError(syslogClient.Connect())
	s.consumerCfg = getConsumerCfg(s.projectId, s.topicId)

	var handleLogsFunc = func(topicID string, shardID int, l *pb.LogGroupList) {
		for _, logGroup := range l.LogGroups {
			for _, log := range logGroup.Logs {

				syslogRFC5424 := NewSyslogFormatRFC5424()

				for _, content := range log.Contents {
					switch content.Key {
					case "pri":
						syslogRFC5424.Priority = content.Value
					case "version":
						syslogRFC5424.Version = content.Value
					case "timestamp":
						timestamp, err := time.Parse(time.RFC3339, content.Value)
						s.NoError(err)
						syslogRFC5424.Timestamp = timestamp
					case "hostname":
						syslogRFC5424.Hostname = content.Value
					case "appName":
						syslogRFC5424.AppName = content.Value
					case "pid":
						syslogRFC5424.Pid = content.Value
					case "msgId":
						syslogRFC5424.MsgId = content.Value
					case "message":
						syslogRFC5424.Message = content.Value
					}
				}

				syslogRFC5424Format, err := syslogRFC5424.FormatSyslogMessageWithRFC5424()
				s.NoError(err)
				s.NoError(syslogClient.Send(syslogRFC5424Format))
				fmt.Println("发送成功")
			}
		}
	}

	tlsConsumer, err := consumer.NewConsumer(context.TODO(), s.consumerCfg, handleLogsFunc)
	s.NoError(err)

	s.NoError(tlsConsumer.Start())

	time.Sleep(60 * time.Second)

	s.NoError(syslogClient.Close())
	tlsConsumer.Stop()
}
