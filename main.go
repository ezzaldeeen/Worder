package main

import (
	"fmt"
	"go.uber.org/zap"
	"runtime"
	"sync"
	"sync/atomic"
	"worder/counter"
	"worder/workerpool"
)

func main() {
	logger, _ := zap.NewDevelopment()
	a := new(atomic.Uint64)
	ch := make(chan string)

	wg := new(sync.WaitGroup)
	dispatcher := counter.NewFileDispatcher("data", ch)
	go func() {
		err := dispatcher.Dispatch()
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	wc := counter.NewWordCounter(ch, a, logger)

	wp := workerpool.NewWorkerPool(1, wc)
	wp.Start(wg)

	fmt.Println("# of Goroutines:", runtime.NumGoroutine())
	fmt.Println("# of Words:", a.Load())
}
