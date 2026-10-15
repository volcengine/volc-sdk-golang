package tls_test

import (
	"context"
	"os"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/tls/consumer"
	"github.com/volcengine/volc-sdk-golang/service/tls/observability"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/producer"
)

// TestHealthObservabilityDocumentationCompiles protects the public API shape
// used by service/tls/health-observability.md. The helpers are intentionally
// not called: this is a compile-only check and must not make network requests.
func TestHealthObservabilityDocumentationCompiles(t *testing.T) {
	_ = producer.StatsOptions{View: producer.StatsViewHealth}
	_ = consumer.StatsOptions{View: consumer.StatsViewOperational}
}

func producerHealthObservabilityDocumentationExampleCompiles() {
	cfg := producer.GetDefaultProducerConfig()
	cfg.Endpoint = os.Getenv("VOLCENGINE_ENDPOINT")
	cfg.Region = os.Getenv("VOLCENGINE_REGION")
	cfg.AccessKeyID = os.Getenv("VOLCENGINE_ACCESS_KEY_ID")
	cfg.AccessKeySecret = os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET")
	cfg.MetricsLevel = producer.MetricsLevelCore

	p := producer.NewProducer(cfg)
	p.Start()
	defer p.Close()

	stats, err := p.GetStats(producer.StatsOptions{View: producer.StatsViewOperational})
	if err == nil && stats.Operational != nil {
		_ = stats.LifecycleState
		_ = stats.Operational.Counters.SendCallsSuccessTotal
		_ = stats.Operational.Counters.SendCallsFailureTotal
		var adapter observability.DeltaAdapter
		if sample, ok := stats.CounterSample(); ok {
			_ = adapter.Observe(sample)
		}
	}
	_ = producer.EvaluateHealth(nil, stats, observability.DefaultHealthPolicy())
}

func consumerHealthObservabilityDocumentationExampleCompiles() {
	cfg := consumer.GetDefaultConsumerConfig()
	cfg.Endpoint = os.Getenv("VOLCENGINE_ENDPOINT")
	cfg.Region = os.Getenv("VOLCENGINE_REGION")
	cfg.AccessKeyID = os.Getenv("VOLCENGINE_ACCESS_KEY_ID")
	cfg.AccessKeySecret = os.Getenv("VOLCENGINE_ACCESS_KEY_SECRET")
	cfg.ProjectID = "your-project"
	cfg.TopicIDList = []string{"your-topic"}
	cfg.ConsumerGroupName = "your-group"
	cfg.MetricsLevel = consumer.MetricsLevelCore

	c, err := consumer.NewConsumer(context.Background(), cfg, func(topicID string, shardID int, logs *pb.LogGroupList) {
		// Process logs.
	})
	if err != nil {
		return
	}
	if err := c.Start(); err != nil {
		return
	}
	defer c.Stop()

	stats, err := c.GetStats(consumer.StatsOptions{View: consumer.StatsViewOperational})
	if err == nil && stats.Operational != nil {
		_ = stats.LifecycleState
		_ = stats.Operational.Counters.FetchCallsSuccessTotal
		_ = stats.Operational.AssignedShards
		var adapter observability.DeltaAdapter
		if sample, ok := stats.CounterSample(); ok {
			_ = adapter.Observe(sample)
		}
	}
	_ = consumer.EvaluateHealth(nil, stats, observability.DefaultHealthPolicy())
}
