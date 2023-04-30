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
	Example: "worder count --path=./data",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		atomicCounter := new(atomic.Uint64)
		pathsChannel := make(chan string)
		wg := new(sync.WaitGroup)
		dispatcher := counter.NewFileDispatcher("data", pathsChannel)

		go func() {
			err := dispatcher.Dispatch()
			if err != nil {
				logger.Error(err.Error())
			}
		}()

		wc := counter.NewWordCounter(pathsChannel, atomicCounter, logger)
		wp := workerpool.NewWorkerPool(1, wc)

		wp.Start(wg)

		fmt.Println("Total Number of Words:", atomicCounter.Load())
	},
}

func init() {
	counterCmd.Flags().StringVarP(&sourcePath, "path", "p",
		"./data", "Directory destination")
	counterCmd.Flags().IntVar(&numOfWorkers, "wpSize",
		10, "Worker Pool Size")
}
