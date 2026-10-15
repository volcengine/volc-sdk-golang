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

// sendQueueLen returns the current depth of the sender task queue.
func (threadPool *ThreadPool) sendQueueLen() int {
	if threadPool == nil {
		return 0
	}
	return len(threadPool.taskChan)
}

// sendQueueCap returns the capacity of the sender task queue.
func (threadPool *ThreadPool) sendQueueCap() int {
	if threadPool == nil {
		return 0
	}
	return cap(threadPool.taskChan)
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

	// Acquire a stable sender index (0..MaxSenderCount-1) from the token pool.
	// The index is exclusive from this point until the goroutine releases it,
	// and is used as the callback shard index to eliminate true sharing.
	var senderIdx int
	select {
	case senderIdx = <-sender.maxSender:
	case <-threadPool.forceQuitCh:
		// Force-quit before acquiring a token: the batch was consumed from
		// taskChan but will never be sent. Use the centralized drop helper.
		if sender.producer != nil {
			sender.producer.dropBatchStats(batch)
			sender.producer.releaseBatchResources(batch)
		}
		return
	}
	select {
	case <-threadPool.forceQuitCh:
		sender.maxSender <- senderIdx
		// Force-quit after acquiring a token but before spawning the sender:
		// the batch will never be sent. Use the centralized drop helper.
		if sender.producer != nil {
			sender.producer.dropBatchStats(batch)
			sender.producer.releaseBatchResources(batch)
		}
		return
	default:
	}
	threadPool.ioWaitGroup.Add(1)

	go func(batch *Batch, senderIdx int) {
		// Register Done() first so it runs LAST (LIFO): the senderRunning
		// counter and the token are released before ioWaitGroup.Done()
		// unblocks Close(), matching the dispatcher pattern.
		defer func() {
			threadPool.ioWaitGroup.Done()
			sender.maxSender <- senderIdx
		}()
		// stats: count actual alive sender goroutine (0..MaxSenderCount)
		var workerGeneration uint32
		if sender.producer != nil && sender.producer.gaugesTracked() {
			workerGeneration = sender.producer.collectionGenerationToken()
			sender.producer.stats.incSenderRunning()
			defer func() {
				if sender.producer.metricsEnabled() &&
					sender.producer.collectionGenerationToken() == workerGeneration {
					sender.producer.stats.decSenderRunning()
				}
			}()
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic happens: %v", r)
				if sender.producer != nil {
					sender.producer.releaseBatchResources(batch)
				}
			}
		}()

		sender.sendToServer(batch, senderIdx)
	}(batch, senderIdx)
}

func (threadPool *ThreadPool) releasePending() {
	if threadPool == nil || threadPool.sender == nil || threadPool.sender.producer == nil {
		return
	}
	producer := threadPool.sender.producer
	for {
		select {
		case batch, ok := <-threadPool.taskChan:
			if !ok {
				return
			}
			// stats: send-queue batch dropped during ForceClose → use the
			// centralized drop helper (callbacks + dropped batch).
			producer.dropBatchStats(batch)
			producer.releaseBatchResources(batch)
		default:
			return
		}
	}
}
