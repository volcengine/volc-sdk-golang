package producer

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type closeRaceCallback struct {
	success int64
	failure int64
}

func (c *closeRaceCallback) Success(_ *Result) { atomic.AddInt64(&c.success, 1) }
func (c *closeRaceCallback) Fail(_ *Result)    { atomic.AddInt64(&c.failure, 1) }

// TestCloseDrainsDequeuedTailLog guards the graceful-close handoff between the
// dispatcher receive loop and the checker that performs the final drain. A log
// may already have been removed from newLogRecvChan when Close stops both
// goroutines; it must still reach the sender before the final snapshot freezes.
func TestCloseDrainsDequeuedTailLog(t *testing.T) {
	const iterations = 256
	for i := 0; i < iterations; i++ {
		client := &mockProducerClient{}
		cfg := GetDefaultProducerConfig()
		cfg.Endpoint = "http://127.0.0.1:0"
		cfg.Region = "cn-beijing"
		cfg.AccessKeyID = "ak"
		cfg.AccessKeySecret = "sk"
		cfg.MetricsLevel = MetricsLevelCore
		cfg.MaxBatchCount = 1000
		cfg.MaxSenderCount = 1
		cfg.BatchQueueSize = 100
		cfg.LingerTime = 10 * time.Second

		p := newTestProducerWithClient(cfg, client)
		p.Start()
		callback := &closeRaceCallback{}
		require.NoError(t, p.SendLog("", "topic", "source", "file", makeLog(), callback), "iteration %d", i)
		p.Close()

		stats := p.collectMeasurements()
		require.Equal(t, stateClosed, stats.State, "iteration %d: state", i)
		require.Equal(t, consistencyFinal, stats.Consistency, "iteration %d: consistency", i)
		require.Equal(t, uint64(1), closeTestValue(stats.AcceptedLogs), "iteration %d: accepted", i)
		require.Equal(t, uint64(1), closeTestValue(stats.TerminalSuccessLogs), "iteration %d: terminal success", i)
		require.Equal(t, uint64(0), closeTestValue(stats.TerminalFailureLogs), "iteration %d: terminal failure", i)
		require.Equal(t, uint64(0), closeTestValue(stats.TerminalDroppedLogs), "iteration %d: terminal dropped", i)
		terminal := closeTestValue(stats.TerminalSuccessLogs) + closeTestValue(stats.TerminalFailureLogs) + closeTestValue(stats.TerminalDroppedLogs)
		require.Equal(t, closeTestValue(stats.AcceptedLogs), terminal, "iteration %d: terminal conservation", i)
		require.Equal(t, uint64(0), closeTestValue(stats.AggregateBatches), "iteration %d: aggregate batches", i)
		require.Equal(t, uint64(0), closeTestValue(stats.AggregateGroups), "iteration %d: aggregate groups", i)
		require.Equal(t, uint64(0), closeTestQueueItems(stats.IngressQueue), "iteration %d: ingress queue", i)
		require.Equal(t, uint64(0), closeTestQueueItems(stats.SendQueue), "iteration %d: send queue", i)
		require.Equal(t, uint64(0), closeTestQueueItems(stats.RetryQueue), "iteration %d: retry queue", i)
		require.Equal(t, uint64(0), closeTestValue(stats.CallbacksPending), "iteration %d: callbacks pending", i)
		require.NotNil(t, stats.Memory, "iteration %d: memory stats", i)
		require.Equal(t, uint64(0), stats.Memory.AccountedUsedBytes, "iteration %d: accounted memory", i)
		require.Equal(t, int64(1), atomic.LoadInt64(&client.totalLogs), "iteration %d: delivered logs", i)
		require.Equal(t, int64(1), atomic.LoadInt64(&callback.success), "iteration %d: callback success", i)
		require.Equal(t, int64(0), atomic.LoadInt64(&callback.failure), "iteration %d: callback failure", i)
	}
}

func closeTestValue(value *uint64) uint64 {
	if value == nil {
		return 0
	}
	return *value
}

func closeTestQueueItems(queue *queueMeasurements) uint64 {
	if queue == nil {
		return 0
	}
	return queue.CurrentItems
}
