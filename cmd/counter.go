package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"sync"
	"sync/atomic"
	"worder/counter"
	"worder/workerpool"
)

var (
	sourcePath   string
	numOfWorkers int
)
var counterCmd = &cobra.Command{
	Use:     "count",
	Short:   "Word counter for the generated files",
	Long:    "",
	Example: "worder count --path=./data --wpsize=10",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		atomicCounter := new(atomic.Uint64)
		pathsChannel := make(chan string)
		wg := new(sync.WaitGroup)
		dispatcher := counter.NewFileDispatcher(sourcePath, pathsChannel)

		go func() {
			err := dispatcher.Dispatch()
			if err != nil {
				logger.Error(err.Error())
			}
		}()

		wordCounter := counter.NewWordCounter(pathsChannel, atomicCounter, logger)
		workerPool := workerpool.NewWorkerPool(numOfWorkers, wordCounter)

		workerPool.Start(wg)

		fmt.Println("Total Number of Words:", atomicCounter.Load())
	},
}

func init() {
	counterCmd.Flags().StringVarP(&sourcePath, "path", "p",
		"./data", "Directory destination")
	counterCmd.Flags().IntVar(&numOfWorkers, "wpSize",
		10, "Worker Pool Size")
}
