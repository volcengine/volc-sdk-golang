package producer

import (
	"sync"
	"sync/atomic"
	"time"
)

type broadcastNotifier struct {
	mu      sync.Mutex
	ch      chan struct{}
	waiters int64
}

func newBroadcastNotifier() *broadcastNotifier {
	return &broadcastNotifier{ch: make(chan struct{})}
}

func (notifier *broadcastNotifier) beginWait() <-chan struct{} {
	notifier.mu.Lock()
	defer notifier.mu.Unlock()
	if notifier.ch == nil {
		notifier.ch = make(chan struct{})
	}
	atomic.AddInt64(&notifier.waiters, 1)
	return notifier.ch
}

func (notifier *broadcastNotifier) endWait() {
	atomic.AddInt64(&notifier.waiters, -1)
}

func (notifier *broadcastNotifier) notify() {
	if notifier == nil || atomic.LoadInt64(&notifier.waiters) <= 0 {
		return
	}
	notifier.mu.Lock()
	defer notifier.mu.Unlock()
	if atomic.LoadInt64(&notifier.waiters) <= 0 {
		return
	}
	if notifier.ch == nil {
		notifier.ch = make(chan struct{})
		return
	}
	close(notifier.ch)
	notifier.ch = make(chan struct{})
}

func waitBroadcast(notifyCh <-chan struct{}, duration time.Duration, done <-chan struct{}) bool {
	if duration <= 0 {
		select {
		case <-done:
			return false
		case <-notifyCh:
			return true
		}
	}
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-done:
		return false
	case <-notifyCh:
		return true
	case <-timer.C:
		return true
	}
}
