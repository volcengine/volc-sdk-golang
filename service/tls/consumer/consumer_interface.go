package consumer

import (
	"context"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type Consumer interface {
	Start() error
	Stop()
}

func NewConsumer(ctx context.Context, conf *Config, f func(topicID string, shardID int, l *pb.LogGroupList)) (Consumer, error) {
	return newConsumer(ctx, conf, f)
}
