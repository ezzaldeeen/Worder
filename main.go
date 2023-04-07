package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
	"worder/generator"
)

func main() {
	rd := strings.NewReader("Hello, World!\nHello, World!\n")
	reader := bufio.NewReaderSize(rd, 100)
	ch := make(chan byte)

	dispatcher := generator.NewBytesDispatcher(reader, ch)

	go func() {
		err := dispatcher.Dispatch()
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		for b := range ch {
			fmt.Print(string(b))
		}
	}()

	time.Sleep(2 * time.Second)
}
