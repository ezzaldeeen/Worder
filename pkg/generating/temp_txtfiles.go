package generating

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path"
	"strconv"
	"sync"
	"worder/custom"
	"worder/utils"
)

// targetFileExt the extension of the generated file
const (
	targetFileExt = "txt"
)

// TXTFilesGenerator is a special generator for generating text
// based on the given source, count, and size (in bytes).
// the generated files will be on a new directory based on the destination
type TXTFilesGenerator struct {
	logger      *zap.SugaredLogger
	source      string
	destination string
	count       int
	size        int
}

func NewTXTFilesGenerator(
	logger *zap.SugaredLogger, src, dest string, count, size int) *TXTFilesGenerator {
	return &TXTFilesGenerator{
		logger:      logger,
		source:      src,
		destination: dest,
		count:       count,
		size:        size,
	}
}

func (fg *TXTFilesGenerator) Generate() {
	wg := sync.WaitGroup{}
	wg.Add(fg.count)

	source, err := os.Open(fg.source)
	defer source.Close()
	if err != nil {
		fg.logger.Errorw(
			custom.SrcFileLoadingErr.Error(), "path", fg.source)
		os.Exit(custom.SrcFileLoadingErrCode)
		// todo: ask saddam in the os.Exit: deferred functions are not run, how to solve it?
	}

	reader := bufio.NewReader(source)
	content, err := utils.GetFileContent(reader)
	if err != nil {
		fg.logger.Error(custom.SrcFileReadingErr.Error())
		os.Exit(custom.SrcFileReadingErrCode)
	}

	for i := 0; i < fg.count; i++ {
		targetPath := path.Join(
			fg.destination,
			fmt.Sprintf("%s.%s", strconv.Itoa(i), targetFileExt))

		target, err := os.Create(targetPath)
		if err != nil {
			fg.logger.Errorw(custom.FileCreatingErr.Error(), "path", fg.destination)
			os.Exit(custom.FileCreatingErrCode)

		}

		go func(wg *sync.WaitGroup, file *os.File, payload string) {
			err := fg.write(file, payload)
			if err != nil {
				fg.logger.Error(custom.FileWritingErr.Error())
				os.Exit(custom.FileWritingErrCode)
			}
			wg.Done()
		}(&wg, target, content)
	}
}

func (fg *TXTFilesGenerator) write(file *os.File, payload string) error {
	defer file.Close()
	writer := bufio.NewWriter(file)

	for curSize := 0; curSize < fg.size; {
		written, err := writer.WriteString(payload)
		if err != nil {
			return err
		}
		curSize += written
	}
	err := writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
