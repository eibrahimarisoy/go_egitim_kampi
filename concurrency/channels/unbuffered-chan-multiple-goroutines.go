package main

import (
	"fmt"
	"time"
)

func main() {
	unBufferedChan := make(chan int)

	// reader goroutine
	go func(unBufChan chan int) {
		value := <-unBufChan
		fmt.Println(value)
	}(unBufferedChan)

	// writer goroutine
	go func(unBufChan chan int) {
		unBufChan <- 1
	}(unBufferedChan)

	fmt.Println("hello channels")
	time.Sleep(time.Second * 1)
}
