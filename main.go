package main

import (
	"fmt"
	counter2 "worder/counter"
)

func main() {
	//rd := strings.NewReader("hello, World!\nHello, World!\n")
	//reader := bufio.NewReader(rd)
	//ch := make(chan byte)
	//acc := make(chan uint)
	//
	//handler := func(b byte) uint {
	//	switch b {
	//	case ' ':
	//		return 1
	//	default:
	//		return 0
	//	}
	//}
	//
	//dispatcher := counter.NewBytesDispatcher(reader, ch)
	//receiver := counter.NewBytesReceiver(ch, acc)
	//
	//go func() {
	//	err := dispatcher.Dispatch()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//
	//go func() {
	//	receiver.Receive(handler)
	//}()
	//
	//go func() {
	//	var sum uint
	//	for a := range acc {
	//		sum += a
	//	}
	//	fmt.Println(sum)
	//}()
	//
	//time.Sleep(2 * time.Second)

	counter := counter2.NewCounter("tdata")
	err := counter.Count()
	if err != nil {
		fmt.Println(err)
	}

	//time.Sleep(15 * time.Second)

}
