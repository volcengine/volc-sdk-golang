package consumer

import (
	"context"
	"strconv"
	"sync/atomic"

	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

// mockConsumerClient is a shared test double for consumer lifecycle, polling,
// checkpoint, and statistics behavior. It keeps the tests independent of
// network I/O while exercising the consumer implementation.
type mockConsumerClient struct {
	tls.Client
	consumeLogsCount          int64
	cursor                    int64
	emptyResponse             bool
	consumeLogsErr            error
	consumeLogsInvalidContent bool
	modifyCheckPointErr       error
}

// newTestConsumerWithClient constructs a normal consumer and replaces the
// clients used by its public operations with a test double. Keeping
// construction on newConsumer ensures tests exercise the production path.
func newTestConsumerWithClient(ctx context.Context, conf *Config, client tls.Client, consumeFunc func(topicID string, shardID int, l *pb.LogGroupList)) (*consumer, error) {
	c, err := newConsumer(ctx, conf, consumeFunc)
	if err != nil {
		return nil, err
	}
	c.client = client
	c.heartbeat.client = client
	c.checkpoint.client = client
	return c, nil
}

func (m *mockConsumerClient) DescribeConsumerGroups(req *tls.DescribeConsumerGroupsRequest) (*tls.DescribeConsumerGroupsResponse, error) {
	return &tls.DescribeConsumerGroupsResponse{
		ConsumerGroups: []*tls.ConsumerGroupResp{{ConsumerGroupName: req.ConsumerGroupName}},
	}, nil
}

func (m *mockConsumerClient) CreateConsumerGroup(req *tls.CreateConsumerGroupRequest) (*tls.CreateConsumerGroupResponse, error) {
	return &tls.CreateConsumerGroupResponse{}, nil
}

func (m *mockConsumerClient) ConsumerHeartbeat(req *tls.ConsumerHeartbeatRequest) (*tls.ConsumerHeartbeatResponse, error) {
	shards := make([]*tls.ConsumeShard, 0, 50)
	for i := 0; i < 50; i++ {
		shards = append(shards, &tls.ConsumeShard{TopicID: "test-topic", ShardID: i})
	}
	return &tls.ConsumerHeartbeatResponse{Shards: shards}, nil
}

func (m *mockConsumerClient) DescribeCheckPoint(req *tls.DescribeCheckPointRequest) (*tls.DescribeCheckPointResponse, error) {
	return &tls.DescribeCheckPointResponse{Checkpoint: ""}, nil
}

func (m *mockConsumerClient) DescribeCursor(req *tls.DescribeCursorRequest) (*tls.DescribeCursorResponse, error) {
	return &tls.DescribeCursorResponse{Cursor: "cursor-0"}, nil
}

func (m *mockConsumerClient) ConsumeLogs(req *tls.ConsumeLogsRequest) (*tls.ConsumeLogsResponse, error) {
	atomic.AddInt64(&m.consumeLogsCount, 1)
	if m.consumeLogsErr != nil {
		return nil, m.consumeLogsErr
	}
	if m.consumeLogsInvalidContent {
		return &tls.ConsumeLogsResponse{
			Cursor: "cursor-invalid",
			Count:  1,
			Logs: &pb.LogGroupList{
				LogGroups: []*pb.LogGroup{{Logs: []*pb.Log{{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}}}}},
			},
		}, &tls.Error{HTTPCode: 400, Code: tls.ErrInvalidContent, Message: "invalid content"}
	}
	c := atomic.AddInt64(&m.cursor, 1)
	resp := &tls.ConsumeLogsResponse{
		Cursor: "cursor-" + strconv.FormatInt(c, 10),
		Count:  1,
	}
	if !m.emptyResponse {
		resp.Logs = &pb.LogGroupList{
			LogGroups: []*pb.LogGroup{{Logs: []*pb.Log{{Contents: []*pb.LogContent{{Key: "k", Value: "v"}}}}}},
		}
	}
	return resp, nil
}

func (m *mockConsumerClient) ModifyCheckPoint(req *tls.ModifyCheckPointRequest) (*tls.CommonResponse, error) {
	if m.modifyCheckPointErr != nil {
		return nil, m.modifyCheckPointErr
	}
	return &tls.CommonResponse{}, nil
}
