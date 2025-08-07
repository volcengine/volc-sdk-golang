package producer

import (
	"fmt"
	"sync"

	"github.com/go-kit/kit/log"
)

type ThreadPool struct {
	stopCh      chan struct{}
	forceQuitCh chan struct{}
	taskChan    chan *Batch
	sender      *Sender
	logger      log.Logger
	ioWaitGroup *sync.WaitGroup
}

func initThreadPool(sender *Sender, logger log.Logger) *ThreadPool {
	return &ThreadPool{
		stopCh:      make(chan struct{}),
		forceQuitCh: make(chan struct{}),
		taskChan:    make(chan *Batch, 100000),
		sender:      sender,
		logger:      logger,
	}
}

func (threadPool *ThreadPool) start(senderWaitGroup *sync.WaitGroup, threadPoolWaitGroup *sync.WaitGroup) {
	defer threadPoolWaitGroup.Done()

	threadPool.ioWaitGroup = senderWaitGroup
	for {
		select {
		case batch := <-threadPool.taskChan:
			threadPool.handleBatch(batch)
		case <-threadPool.stopCh:
			for batch := range threadPool.taskChan {
				threadPool.handleBatch(batch)
			}

			return
		case <-threadPool.forceQuitCh:
			return
		}
	}
}

func (threadPool *ThreadPool) handleBatch(batch *Batch) {
	if batch == nil {
		return
	}

	threadPool.ioWaitGroup.Add(1)
	threadPool.sender.maxSender <- 1

	go func(batch *Batch) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic happens: %v", r)
			}
		}()
		defer func() {
			threadPool.ioWaitGroup.Done()
			<-threadPool.sender.maxSender
		}()

		threadPool.sender.sendToServer(batch)
	}(batch)
}
