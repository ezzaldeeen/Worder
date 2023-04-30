package workerpool

import (
	"context"
	"sync"
)

// Runnable for the executable tasks by the worker pool
type Runnable interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
}

// WorkerPool the controller of the execution behavior
type WorkerPool struct {
	numOfWorkers int
	runnable     Runnable
	ctx          context.Context
	cancelFn     context.CancelFunc
}

func NewWorkerPool(numOfWorkers int, runnable Runnable) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		numOfWorkers: numOfWorkers,
		runnable:     runnable,
		ctx:          ctx,
		cancelFn:     cancel,
	}
}

// Start starting the execution based on the number of workers
func (wp *WorkerPool) Start(wg *sync.WaitGroup) {
	wg.Add(wp.numOfWorkers)
	for i := 0; i < wp.numOfWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			wp.runnable.Run(wp.ctx, wg)
		}(wg)
	}
	wg.Wait()
}

// Stop canceling the execution by cancel signal
func (wp *WorkerPool) Stop() {
	wp.cancelFn()
}
