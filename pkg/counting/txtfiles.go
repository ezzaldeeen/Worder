package counting

import (
	"bufio"
	"go.uber.org/zap"
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"
)

type Counter interface {
	Count()
}

type WordCounter struct {
	logger *zap.SugaredLogger
	source string
}

func NewWordCounter(logger *zap.SugaredLogger, source string) *WordCounter {
	return &WordCounter{
		logger: logger,
		source: source,
	}
}

func (wc *WordCounter) Count() uint64 {
	var counter uint64
	files, err := os.ReadDir(wc.source)
	if err != nil {
		wc.logger.Error(err)
		os.Exit(0)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(files))

	for _, file := range files {
		filePath := path.Join(wc.source, file.Name())
		currFile, err := os.Open(filePath)
		if err != nil {
			wc.logger.Error(err)
			os.Exit(0)
		}

		go func(wg *sync.WaitGroup, file *os.File) {
			defer file.Close()
			reader := bufio.NewReader(currFile)

		loop:
			for {
				buf, err := reader.ReadString('\n')
				if err != nil {
					switch err {
					case io.EOF:
						break loop
					default:
						wc.logger.Error(err)
					}
				}
				words := strings.Fields(buf)
				atomic.AddUint64(&counter, uint64(len(words)))
			}
			wg.Done()
		}(&wg, currFile)
	}

	wg.Wait()
	return atomic.LoadUint64(&counter)
}
