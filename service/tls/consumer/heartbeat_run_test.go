package consumer

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

// emptyShardsHeartbeatClient is a minimal tls.Client stand-in whose
// ConsumerHeartbeat always succeeds but returns zero shards. That keeps
// heartbeatRunner.run() inside its startup retry loop for all 5 attempts,
// exercising the inter-attempt wait — the code path that used to sleep without
// observing ctx.
type emptyShardsHeartbeatClient struct {
	tls.Client
	calls int64 // atomic
}

func (c *emptyShardsHeartbeatClient) ConsumerHeartbeat(req *tls.ConsumerHeartbeatRequest) (*tls.ConsumerHeartbeatResponse, error) {
	atomic.AddInt64(&c.calls, 1)
	return &tls.ConsumerHeartbeatResponse{Shards: nil}, nil
}

// TestHeartbeatRun_StartupLoopHonorsCtxCancel verifies that when ctx is
// cancelled before run() starts (Start immediately followed by Stop), the
// startup retry loop must return promptly instead of consuming its full
// ~2.5s (5 attempts x 500ms) retry budget.
func TestHeartbeatRun_StartupLoopHonorsCtxCancel(t *testing.T) {
	client := &emptyShardsHeartbeatClient{}
	conf := GetDefaultConsumerConfig()
	conf.HeartbeatIntervalInSecond = 1

	expiredCh := make(chan struct{})
	h := newHeartbeatRunner(log.NewNopLogger(), client, conf, expiredCh, nil)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancelled before run() ever starts, as an immediate Stop() would

	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan struct{})
	start := time.Now()
	go func() {
		h.run(ctx, &wg)
		close(done)
	}()

	// The buggy implementation sleeps ~2.5s in the startup loop before it ever
	// reaches the ctx.Done() select. Give a generous-but-decisive budget: the
	// fix returns in well under 500ms; the bug needs ~2.5s.
	select {
	case <-done:
		elapsed := time.Since(start)
		if elapsed > 500*time.Millisecond {
			t.Fatalf("run() returned after %v; want < 500ms (startup loop must observe ctx cancel)", elapsed)
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("run() did not return within 2s after ctx cancel; startup retry loop ignored ctx.Done()")
	}

	wg.Wait()

	// At most one attempt may have started before the first ctx check; the loop
	// must not have run all 5 attempts.
	if n := atomic.LoadInt64(&client.calls); n > 1 {
		t.Fatalf("ConsumerHeartbeat called %d times after pre-cancelled ctx; want <= 1", n)
	}
}

// signalingHeartbeatClient succeeds with zero shards (keeping run() in the
// startup retry loop) and closes firstCall exactly once, on the first
// ConsumerHeartbeat. That gives the test a deterministic edge to synchronise
// on: once firstCall is closed, run() has finished attempt #1 and is about to
// enter — or is already inside — the inter-attempt retry wait.
type signalingHeartbeatClient struct {
	tls.Client
	calls     int64 // atomic
	firstCall chan struct{}
	once      sync.Once
}

func (c *signalingHeartbeatClient) ConsumerHeartbeat(req *tls.ConsumerHeartbeatRequest) (*tls.ConsumerHeartbeatResponse, error) {
	atomic.AddInt64(&c.calls, 1)
	c.once.Do(func() { close(c.firstCall) })
	return &tls.ConsumerHeartbeatResponse{Shards: nil}, nil
}

// TestHeartbeatRun_RetryWaitHonorsCtxCancel is the coverage the pre-cancel test
// cannot provide: it proves the 500ms inter-attempt retry wait itself is
// cancellable, not merely the ctx check at the top of the loop.
//
// It never uses sleep to guess scheduling. Instead it lets attempt #1 run for
// real, blocks on firstCall until run() has entered the retry wait, then
// cancels. Against the fixed implementation (timer + ctx select) run() returns
// essentially immediately and never starts attempt #2. Against the old
// uncancellable time.Sleep version, the sleep ignores the cancel, run() burns
// the full ~2.5s (5 attempts x 500ms) budget and calls ConsumerHeartbeat 5
// times — failing both assertions below.
func TestHeartbeatRun_RetryWaitHonorsCtxCancel(t *testing.T) {
	client := &signalingHeartbeatClient{firstCall: make(chan struct{})}
	conf := GetDefaultConsumerConfig()
	conf.HeartbeatIntervalInSecond = 1

	expiredCh := make(chan struct{})
	h := newHeartbeatRunner(log.NewNopLogger(), client, conf, expiredCh, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan struct{})
	go func() {
		h.run(ctx, &wg)
		close(done)
	}()

	// Synchronise on the real first heartbeat instead of sleeping: block until
	// attempt #1 has happened and run() is entering/inside the retry wait.
	<-client.firstCall

	// Cancel during the retry wait and time the return from that moment. A
	// cancellable wait returns near-instantly; the buggy time.Sleep would keep
	// blocking for the remainder of the ~2.5s startup budget.
	cancelTime := time.Now()
	cancel()

	select {
	case <-done:
		if elapsed := time.Since(cancelTime); elapsed > 250*time.Millisecond {
			t.Fatalf("run() returned %v after cancel; want < 250ms (retry wait must be cancellable, not a fixed sleep)", elapsed)
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("run() did not return within 2s after cancel during retry wait; the wait ignored ctx.Done()")
	}

	wg.Wait()

	// Cancelling during the retry wait must prevent attempt #2. Exactly one
	// heartbeat should have been issued (the one we synchronised on).
	if n := atomic.LoadInt64(&client.calls); n != 1 {
		t.Fatalf("ConsumerHeartbeat called %d times; want exactly 1 (cancel during retry wait must stop before attempt #2)", n)
	}
}
