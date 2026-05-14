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
	sender := threadPool.sender
	if sender == nil {
		return
	}

	select {
	case sender.maxSender <- 1:
	case <-threadPool.forceQuitCh:
		if sender.producer != nil {
			sender.producer.releaseBatchResources(batch)
		}
		return
	}
	select {
	case <-threadPool.forceQuitCh:
		<-sender.maxSender
		if sender.producer != nil {
			sender.producer.releaseBatchResources(batch)
		}
		return
	default:
	}
	threadPool.ioWaitGroup.Add(1)

	go func(batch *Batch) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic happens: %v", r)
				if sender.producer != nil {
					sender.producer.releaseBatchResources(batch)
				}
			}
		}()
		defer func() {
			threadPool.ioWaitGroup.Done()
			<-sender.maxSender
		}()

		sender.sendToServer(batch)
	}(batch)
}

func (threadPool *ThreadPool) releasePending() {
	if threadPool == nil || threadPool.sender == nil || threadPool.sender.producer == nil {
		return
	}
	for {
		select {
		case batch, ok := <-threadPool.taskChan:
			if !ok {
				return
			}
			threadPool.sender.producer.releaseBatchResources(batch)
		default:
			return
		}
	}
}
