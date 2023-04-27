package counter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

type WordCounter struct {
	source string
}

func NewWordCounter(source string) *WordCounter {
	return &WordCounter{
		source: source,
	}
}

func (wc WordCounter) Count() error {
	files, err := os.ReadDir(wc.source)
	if err != nil {
		return errors.New(fmt.Sprintf("directory not found: %s", wc.source))
	}

	for _, f := range files {
		dispatching := make(chan byte)
		accumulating := make(chan uint)

		fmt.Println("iter")

		filePath := path.Join(wc.source, f.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return errors.New(fmt.Sprintf("can't open file: %s", filePath))
		}

		reader := bufio.NewReader(file)
		dispatcher := NewBytesDispatcher(reader, dispatching)
		receiver := NewBytesReceiver(dispatching, accumulating)

		go func() {
			err := dispatcher.Dispatch()
			if err != nil {
				return
			}
		}()

		handler := func(b byte) uint {
			switch b {
			case ' ':
				return 1
			default:
				return 0
			}
		}

		var sum uint
		go func() {
			receiver.Receive(handler)
			for a := range accumulating {
				sum += a
			}
		}()

		time.Sleep(2 * time.Second)

		fmt.Printf("Total Words: %d - file: %s\n", sum, filePath)
	}

	return nil
}
