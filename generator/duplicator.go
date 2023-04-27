package generator

import (
	"bytes"
	"context"
)

type BytesDuplicator struct {
	byteSize uint
	buffSize uint
	in       <-chan byte
	out      chan<- byte
}

func NewBytesDuplicator(
	in <-chan byte, out chan<- byte) *BytesDuplicator {
	return &BytesDuplicator{
		in:  in,
		out: out,
	}
}

func (bd *BytesDuplicator) Duplicate() error {
	defer close(bd.out)
	numOfIterations := int(bd.byteSize / bd.buffSize)
	buf := bytes.NewBuffer([]byte{})
	for b := range bd.in {
		buf.WriteByte(b)
		if buf.Len() == int(bd.buffSize) {
			for i := 0; i < numOfIterations; i++ {
				buf.Write(buf.Bytes())
			}
		}
	}
	return nil
}

func (bd *BytesDuplicator) DuplicateWithContext(ctx context.Context) error {
	defer close(bd.out)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:

		}
	}
}
