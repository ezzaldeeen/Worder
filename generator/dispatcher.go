package generator

import (
	"bufio"
	"io"
)

// BytesDispatcher todo
type BytesDispatcher struct {
	reader *bufio.Reader
	ch     chan<- byte
}

func NewBytesDispatcher(
	reader *bufio.Reader, ch chan<- byte) *BytesDispatcher {
	return &BytesDispatcher{
		reader: reader,
		ch:     ch,
	}
}

// Dispatch todo
func (bd *BytesDispatcher) Dispatch() error {
	defer close(bd.ch)
	for {
		readByte, err := bd.reader.ReadByte()
		if err != nil {
			switch err {
			case io.EOF:
				return nil
			default:
				return err
			}
		}
		bd.ch <- readByte
	}
}
