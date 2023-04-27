package counter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
)

type Counter struct {
	source string
}

func NewCounter(src string) *Counter {
	return &Counter{
		source: src,
	}
}

func (c Counter) Count() error {
	files, err := os.ReadDir(c.source)
	if err != nil {
		return errors.New(fmt.Sprintf("directory not found: %s", c.source))
	}

	var sum uint
	for _, f := range files {
		filePath := path.Join(c.source, f.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return errors.New(fmt.Sprintf("can't open file: %s", filePath))
		}
		reader := bufio.NewReader(file)

		for {
			readByte, err := reader.ReadByte()
			if err != nil {
				break
			}
			if readByte == ' ' {
				sum += 1
			}
		}
	}

	fmt.Printf("total count: %d\n", sum)
	return nil
}
