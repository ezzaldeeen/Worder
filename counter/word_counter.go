package counter

import (
	"bufio"
	"context"
	"go.uber.org/zap"
	"os"
	"path"
	"sync"
)

// FilesDispatcher sending the file paths through paths channel
type FilesDispatcher struct {
	source string
	paths  chan<- string
}

// NewFileDispatcher factory function
func NewFileDispatcher(src string, ch chan<- string) *FilesDispatcher {
	return &FilesDispatcher{
		source: src,
		paths:  ch,
	}
}

// Dispatch operation for listing the files in the given directory (source)
// and sending the file paths in that dir through a paths channel
func (fd FilesDispatcher) Dispatch() error {
	defer close(fd.paths)
	entries, err := os.ReadDir(fd.source)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		filePath := path.Join(fd.source, entry.Name())
		fd.paths <- filePath
	}
	return nil
}

type WordCounter struct {
	logger    zap.Logger
	filePaths <-chan string
}

func NewWordCounter(ch <-chan string) *WordCounter {
	return &WordCounter{
		filePaths: ch,
	}
}

func (wc WordCounter) Run(ctx context.Context, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		wc.logger.Error("context cancelled")
	default:
		filePath := <-wc.filePaths
		file, err := os.Open(filePath)
		defer file.Close()
		if err != nil {
			wc.logger.Error("unable to open the file:", zap.Field{
				Key:    "path",
				String: filePath,
			})
		}

		reader := bufio.NewReader(file)
		for {
			b, err := reader.ReadByte()
			if err != nil {
				return
			}
			if b == ' ' {
				// todo: increase counter
			}
		}
	}
}
