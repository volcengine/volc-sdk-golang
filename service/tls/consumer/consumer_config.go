package consumer

import (
	"errors"
	"github.com/go-kit/kit/log"
	"regexp"

	"github.com/volcengine/volc-sdk-golang/service/tls/common"
)

type Config struct {
	common.LoggerConfig
	common.ClientConfig
	ProjectID                      string
	TopicIDList                    []string
	ConsumerGroupName              string
	ConsumerName                   string
	ConsumeFrom                    string
	HeartbeatIntervalInSecond      int
	DataFetchIntervalInMillisecond int64
	FlushCheckpointIntervalSecond  int
	MaxFetchLogGroupCount          int
	OrderedConsume                 bool
	Logger                         *log.Logger
}

func GetDefaultConsumerConfig() *Config {
	return &Config{
		LoggerConfig: common.LoggerConfig{
			LogLevel:      "info",
			LogFileName:   "",
			IsJsonType:    false,
			LogMaxSize:    10,
			LogMaxBackups: 10,
			LogCompress:   false,
		},
		ConsumeFrom:                    ConsumeFromBegin,
		HeartbeatIntervalInSecond:      20,
		DataFetchIntervalInMillisecond: 200,
		MaxFetchLogGroupCount:          100,
		FlushCheckpointIntervalSecond:  5,
		OrderedConsume:                 false,
	}
}

func validateConsumerConfig(c *Config) error {
	if len(c.ProjectID) == 0 {
		return errors.New("empty ProjectID")
	}

	if len(c.TopicIDList) == 0 {
		return errors.New("empty TopicIDList")
	}

	if len(c.ConsumerGroupName) == 0 {
		return errors.New("empty ConsumerGroupName")
	}

	// ConsumeFrom specify the cursor position. If it's not "begin" or "end", it must be a timestamp string.
	if c.ConsumeFrom != ConsumeFromBegin && c.ConsumeFrom != ConsumeFromEnd {
		ok, _ := regexp.MatchString("\\d+", c.ConsumeFrom)
		if !ok {
			return errors.New("invalid ConsumeFrom. valid options: \"begin\", \"end\" or timestamp")
		}
	}

	if c.HeartbeatIntervalInSecond <= 0 || c.HeartbeatIntervalInSecond > 300 {
		return errors.New("invalid HeartbeatIntervalInSecond. acceptable range: (1, 300]")
	}

	if c.DataFetchIntervalInMillisecond <= 0 || c.DataFetchIntervalInMillisecond > 300000 {
		return errors.New("invalid DataFetchIntervalInMillisecond. acceptable range: (1, 300000]")
	}

	if c.FlushCheckpointIntervalSecond <= 0 || c.FlushCheckpointIntervalSecond > 300 {
		return errors.New("invalid FlushCheckpointIntervalSecond. acceptable range: (1, 300]")
	}

	if c.MaxFetchLogGroupCount <= 0 || c.MaxFetchLogGroupCount > 1000 {
		return errors.New("invalid MaxFetchLogGroupCount. acceptable range: (1, 1000]")
	}

	return nil
}
