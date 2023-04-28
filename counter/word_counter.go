package counter

import (
	"context"
	"os"
	"path"
	"sync"
)

type FilesDispatcher struct {
	directoryPath string
	filePaths     chan<- string
}

func FileDispatcher(dirPath string, ch chan<- string) *FilesDispatcher {
	return &FilesDispatcher{
		directoryPath: dirPath,
		filePaths:     ch,
	}
}

func (fd FilesDispatcher) Dispatch() error {
	defer close(fd.filePaths)
	entries, err := os.ReadDir(fd.directoryPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		filePath := path.Join(fd.directoryPath, entry.Name())
		fd.filePaths <- filePath
	}
	return nil
}

type WordCounter struct {
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
	default:
	}
}
