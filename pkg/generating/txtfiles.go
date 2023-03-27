package generating

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path"
	"strconv"
	"sync"
	"worder/pkg"
	"worder/utils"
)

// targetFileExt the extension of the generated file
const targetFileExt = "txt"

type FilesGenerator interface {
	Generate()
}

// TXTFilesGenerator is a special generator for generating text
// based on the given source, count, and size (in bytes).
// the generated files will be on a new directory based on the destination
// todo: should be FileGenerator
// todo: Classes, Structs, functions should be singular (doing one thing at a time)
// todo: the only exceptions Routes Resource, Table names should be plurals
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

// Generate multiple files from the given source
// based on the count, size of each file
// the Generate expects two works concurrently
// todo: functions should be singular concerns (this function should generate a single file)
// todo: package names should be nouns (e.g. counting -> counter, generating -> generator)
// todo: package that have two files it's not a package
func (fg *TXTFilesGenerator) Generate() {
	// todo: the user should decide whether this function is going to use concurrently or not
	// todo: use errgroup
	wg := sync.WaitGroup{}
	wg.Add(fg.count)

	source, err := os.Open(fg.source)
	defer source.Close()
		
	if err != nil {
		fg.logger.Errorw(
			// todo: shouldn't care about logging
			pkg.SrcFileLoadingErr.Error(), "path", fg.source)
		os.Exit(pkg.SrcFileLoadingErrCode)
		// todo: ask saddam in the os.Exit: deferred functions are not run, how to solve it?
	}

	reader := bufio.NewReader(source)
	content, err := utils.GetFileContent(reader)
	if err != nil {
		fg.logger.Error(pkg.SrcFileReadingErr.Error())
		os.Exit(pkg.SrcFileReadingErrCode) // todo: os.Exist should be removed, and return instead
	}

	for i := 0; i < fg.count; i++ {
		targetPath := path.Join(
			fg.destination,
			fmt.Sprintf("%s.%s", strconv.Itoa(i), targetFileExt))

		target, err := os.Create(targetPath)
		if err != nil {
			fg.logger.Errorw(pkg.FileCreatingErr.Error(), "path", fg.destination)
			os.Exit(pkg.FileCreatingErrCode)

		}

		go func(wg *sync.WaitGroup, file *os.File, payload string) {
			err := fg.write(file, payload)
			if err != nil {
				fg.logger.Error(pkg.FileWritingErr.Error())
				os.Exit(pkg.FileWritingErrCode)
			}
			wg.Done()
		}(&wg, target, content)
	}

	wg.Wait()
}

// write to file based on the given payload
// with a specific size in bytes
// todo: always deal with bytes rather than dealing with string
// todo: always return the number of written bytes (much easier to test, closer to the actual read and write implementation in the os)
func (fg *TXTFilesGenerator) write(file *os.File, payload string) error {
	defer file.Close()
	writer := bufio.NewWriter(file)
	byteWritten := 0

	for byteWritten < fg.size {
		written, err := writer.WriteString(payload)
		if err != nil {
			return err
		}
		byteWritten += written
	}
	err := writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
