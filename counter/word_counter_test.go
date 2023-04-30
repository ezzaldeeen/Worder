package counter

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"worder/workerpool"
)

func BenchmarkWordCounter_Run(b *testing.B) {
	logger, _ := zap.NewDevelopment()
	atomicCounter := new(atomic.Uint64)

	for _, v := range []int{1, 2, 3, 4} {
		b.Run(fmt.Sprintf("worker_pool_size_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				pathsChannel := make(chan string)
				wg := new(sync.WaitGroup)
				// todo: change path
				dispatcher := NewFileDispatcher("/Users/ezzaldeen/Hands-on/Golang/worder/data", pathsChannel)

				go func() {
					err := dispatcher.Dispatch()
					if err != nil {
						logger.Error(err.Error())
					}
				}()
				wc := NewWordCounter(pathsChannel, atomicCounter, logger)

				wp := workerpool.NewWorkerPool(v, wc)
				wp.Start(wg)

			}
		})
	}
}

func TestNewFileDispatcher(t *testing.T) {
	type args struct {
		src string
		ch  chan<- string
	}
	var tests []struct {
		name string
		args args
		want *FilesDispatcher
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileDispatcher(tt.args.src, tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileDispatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCounter_Run(t *testing.T) {
	type fields struct {
		paths   <-chan string
		counter *atomic.Uint64
		logger  *zap.Logger
	}
	type args struct {
		w   int
		ctx context.Context
		wg  *sync.WaitGroup
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := WordCounter{
				paths:   tt.fields.paths,
				counter: tt.fields.counter,
				logger:  tt.fields.logger,
			}
			wc.Run(tt.args.w, tt.args.ctx, tt.args.wg)
		})
	}
}
