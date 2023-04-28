package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	ch := make(chan string)

	entries, err := os.ReadDir("data")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 2; i++ {
		go func() {
			for _, entry := range entries {
				filePath := path.Join("data", entry.Name())
				ch <- filePath
			}
		}()
	}

	go func() {
		for filePath := range ch {
			fmt.Println(filePath)
		}
	}()

	time.Sleep(5 * time.Second)

}
