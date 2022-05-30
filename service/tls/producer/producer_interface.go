package producer

import (
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type Producer interface {
	SendLog(shardHash, topic, source, filename string, log *pb.Log, callBack CallBack) error
	SendLogs(shardHash, topic, source, filename string, logs *pb.LogGroup, callBack CallBack) error
	ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string)
	ResetProducerConfig(producerConfig *Config)
	Start()
	Close()
	ForceClose()
}

func NewProducer(config *Config) Producer {
	return newProducer(config)
}
